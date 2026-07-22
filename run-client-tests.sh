#!/usr/bin/env bash
# Orchestrates multi-client tests against no-auth and auth Manticore instances.
#
# Host usage:
#   ./run-client-tests.sh
#   CLIENTS=python,java ./run-client-tests.sh
#   EXCLUDE_CLIENTS=java,rust ./run-client-tests.sh
#   ./run-client-tests.sh --exclude java,rust
#
# Container usage (compose entrypoint):
#   ./run-client-tests.sh --matrix

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLIENTS_SRC="${CLIENTS_SRC:-/clients}"
EXCLUDE_CLIENTS="${EXCLUDE_CLIENTS:-}"
WORK_ROOT="${WORK_ROOT:-/tmp/openapi-client-tests}"

MANTICORE_NOAUTH_URL="${MANTICORE_NOAUTH_URL:-http://localhost:19408}"
MANTICORE_AUTH_URL="${MANTICORE_AUTH_URL:-http://localhost:19508}"
MANTICORE_USER="${MANTICORE_USER:-admin}"
MANTICORE_PASSWORD="${MANTICORE_PASSWORD:-TestPass#2026}"

# Modes: noauth against 19408; basic + bearer against the auth instance (19508).
MODES=(noauth basic bearer)

ALL_CLIENTS=(
  python
  python-asyncio
  java
  javascript
  typescript
  csharp
  go
  rust
  elixir
)

declare -A CLIENT_DIRS=(
  [python]=manticoresearch-python
  [python-asyncio]=manticoresearch-python-asyncio
  [java]=manticoresearch-java
  [javascript]=manticoresearch-javascript
  [typescript]=manticoresearch-typescript
  [csharp]=manticoresearch-net
  [go]=manticoresearch-go
  [rust]=manticoresearch-rust
  [elixir]=manticoresearch-elixir
)

PASS_COUNT=0
FAIL_COUNT=0
SKIP_COUNT=0
RESULTS=()

log() { printf '[%s] %s\n' "$(date '+%H:%M:%S')" "$*"; }
hr() { printf '%s\n' "------------------------------------------------------------"; }

# Drop every table on an instance so later clients do not see leftover state.
cleanup_tables() {
  local url="$1" need_auth="${2:-0}"
  python3 - "$url" "$need_auth" "$MANTICORE_USER" "$MANTICORE_PASSWORD" <<'CLEANUP_PY'
import base64, json, sys, urllib.parse, urllib.request

url, need_auth, user, password = sys.argv[1:5]
base = url.rstrip("/")

def sql(query: str):
    req = urllib.request.Request(
        f"{base}/sql?mode=raw",
        data=urllib.parse.urlencode({"query": query}).encode(),
        method="POST",
    )
    if need_auth == "1":
        token = base64.b64encode(f"{user}:{password}".encode()).decode()
        req.add_header("Authorization", f"Basic {token}")
    with urllib.request.urlopen(req, timeout=30) as resp:
        return json.load(resp)

try:
    payload = sql("SHOW TABLES")
except Exception as exc:
    print(f"cleanup: SHOW TABLES failed on {base}: {exc}", file=sys.stderr)
    sys.exit(0)

blocks = payload if isinstance(payload, list) else [payload]
tables = []
for block in blocks:
    if not isinstance(block, dict):
        continue
    for row in block.get("data") or []:
        if not isinstance(row, dict):
            continue
        name = row.get("Table") or row.get("Index") or row.get("table")
        if name:
            tables.append(str(name))

dropped = 0
for name in tables:
    safe = name.replace("`", "``")
    try:
        sql(f"DROP TABLE IF EXISTS `{safe}`")
        dropped += 1
    except Exception as exc:
        print(f"cleanup: DROP `{safe}` failed on {base}: {exc}", file=sys.stderr)

print(f"cleanup: dropped {dropped}/{len(tables)} table(s) on {base}", flush=True)
CLEANUP_PY
}
basic_auth_header() {
  printf 'Basic %s' "$(printf '%s:%s' "$MANTICORE_USER" "$MANTICORE_PASSWORD" | base64 -w0 2>/dev/null || printf '%s:%s' "$MANTICORE_USER" "$MANTICORE_PASSWORD" | base64)"
}

wait_http() {
  local url="$1" label="$2" expect_auth="${3:-0}"
  log "Waiting for $label at $url ..."
  local i code
  for i in $(seq 1 60); do
    code="$(curl -s -o /dev/null -w '%{http_code}' "$url/" || true)"
    if [[ "$expect_auth" == "1" ]]; then
      # Up once we get any HTTP response (often 401 before/after bootstrap).
      if [[ "$code" != "000" && -n "$code" ]]; then
        log "$label is up (HTTP $code)"
        return 0
      fi
    else
      if [[ "$code" =~ ^(200|400|404|405)$ ]]; then
        log "$label is up (HTTP $code)"
        return 0
      fi
    fi
    sleep 2
  done
  log "ERROR: timed out waiting for $label"
  return 1
}

bootstrap_auth() {
  log "Bootstrapping auth admin user on manticore-auth ..."
  # Idempotent-ish: if bootstrap already done, continue.
  # /sql without mode=raw only accepts SELECT; SHOW VERSION needs raw mode.
  if curl -sf -u "${MANTICORE_USER}:${MANTICORE_PASSWORD}" \
      "${MANTICORE_AUTH_URL}/sql?mode=raw" --data-urlencode "query=SHOW VERSION" >/dev/null 2>&1; then
    log "Auth admin already usable"
    return 0
  fi

  # Docker image applies searchd_* env overrides via manticore.conf.sh (executable
  # config). Bootstrap must use that same config — the static manticore.conf does
  # not include auth=1 and fails with "authentication disabled".
  # Run as manticore so auth.json is owned by the daemon user (not root).
  local bootstrap_cmd
  bootstrap_cmd=$(cat <<EOF
set -e
CONFIG=/etc/manticoresearch/manticore.conf.sh
if [ ! -e "\$CONFIG" ]; then
  CONFIG=/etc/manticoresearch/manticore.conf
fi
bash "\$CONFIG" 2>/dev/null | grep -E 'auth\\s*=' || true
printf '%s\\n%s\\n%s\\n' '$MANTICORE_USER' '$MANTICORE_PASSWORD' '$MANTICORE_PASSWORD' | searchd --config "\$CONFIG" --auth-non-interactive
EOF
)

  if docker compose -f "$ROOT_DIR/docker-compose.yml" exec -T -u manticore manticore-auth \
      bash -lc "$bootstrap_cmd"; then
    log "Auth bootstrap completed"
  else
    sleep 3
    docker compose -f "$ROOT_DIR/docker-compose.yml" exec -T -u manticore manticore-auth \
      bash -lc "$bootstrap_cmd" \
      || log "WARNING: auth bootstrap command returned non-zero (store may already exist)"
  fi

  local i
  for i in $(seq 1 30); do
    if curl -sf -u "${MANTICORE_USER}:${MANTICORE_PASSWORD}" \
        "${MANTICORE_AUTH_URL}/sql?mode=raw" --data-urlencode "query=SHOW VERSION" >/dev/null 2>&1; then
      log "Authenticated SQL works"
      return 0
    fi
    sleep 1
  done
  log "ERROR: auth admin login failed after bootstrap"
  return 1
}

fetch_bearer_token() {
  local response token
  response="$(curl -sf -u "${MANTICORE_USER}:${MANTICORE_PASSWORD}" \
      -X POST "${MANTICORE_AUTH_URL}/token" -H 'Content-Type: application/json' -d '{}')"
  token="$(printf '%s' "$response" | python3 -c '
import json,sys
raw=sys.stdin.read().strip()
try:
  data=json.loads(raw)
  if isinstance(data, str):
    print(data)
  elif isinstance(data, dict):
    print(data.get("token") or data.get("access_token") or data.get("id") or "")
  else:
    print(raw)
except Exception:
  print(raw.strip().strip("\""))
')"
  if [[ -z "$token" ]]; then
    log "ERROR: empty bearer token from /token (response=$response)"
    return 1
  fi
  printf '%s' "$token"
}

normalize_client_name() {
  local c
  c="$(printf '%s' "$1" | tr '[:upper:]' '[:lower:]' | xargs)"
  case "$c" in
    net|dotnet|c#) c=csharp ;;
    js) c=javascript ;;
    ts) c=typescript ;;
    py) c=python ;;
    py-asyncio|asyncio) c=python-asyncio ;;
  esac
  printf '%s' "$c"
}

# Parse a comma/space-separated client list into normalized names (one per line).
parse_client_list() {
  local raw="${1:-}" c
  raw="${raw//,/ }"
  for c in $raw; do
    [[ -z "$c" ]] && continue
    c="$(normalize_client_name "$c")"
    if [[ -z "${CLIENT_DIRS[$c]+x}" ]]; then
      log "ERROR: unknown client '$c'. Valid: ${ALL_CLIENTS[*]}"
      exit 2
    fi
    printf '%s\n' "$c"
  done
}

selected_clients() {
  local -a chosen=()
  local -A exclude=()
  local c

  if [[ -n "${EXCLUDE_CLIENTS:-}" ]]; then
    while IFS= read -r c; do
      [[ -z "$c" ]] && continue
      exclude["$c"]=1
    done < <(parse_client_list "$EXCLUDE_CLIENTS")
  fi

  if [[ -z "${CLIENTS:-}" ]]; then
    chosen=("${ALL_CLIENTS[@]}")
  else
    while IFS= read -r c; do
      [[ -z "$c" ]] && continue
      chosen+=("$c")
    done < <(parse_client_list "$CLIENTS")
  fi

  for c in "${chosen[@]}"; do
    if [[ -n "${exclude[$c]+x}" ]]; then
      continue
    fi
    printf '%s\n' "$c"
  done
}

mode_url() {
  case "$1" in
    noauth) printf '%s' "$MANTICORE_NOAUTH_URL" ;;
    basic|bearer) printf '%s' "$MANTICORE_AUTH_URL" ;;
    *) return 1 ;;
  esac
}

record_result() {
  local client="$1" mode="$2" status="$3" detail="${4:-}"
  RESULTS+=("$client|$mode|$status|$detail")
  case "$status" in
    PASS) PASS_COUNT=$((PASS_COUNT + 1)) ;;
    FAIL) FAIL_COUNT=$((FAIL_COUNT + 1)) ;;
    SKIP) SKIP_COUNT=$((SKIP_COUNT + 1)) ;;
  esac
}

# Rewrite hardcoded hosts/ports and inject credentials into a working copy.
prepare_workdir() {
  local client="$1" mode="$2" src="$3" dst="$4" url="$5" token="${6:-}"
  rm -rf "$dst"
  mkdir -p "$(dirname "$dst")"
  rsync -a --delete \
    --exclude node_modules \
    --exclude target \
    --exclude .git \
    --exclude dist \
    --exclude bin \
    --exclude obj \
    --exclude _build \
    --exclude deps \
    --exclude .venv \
    --exclude __pycache__ \
    "$src/" "$dst/"

  python3 - "$client" "$mode" "$dst" "$url" "$MANTICORE_USER" "$MANTICORE_PASSWORD" "$token" <<'PY'
import base64, pathlib, re, sys

client, mode, root, url, user, password, token = sys.argv[1:8]
root = pathlib.Path(root)
url = url.rstrip("/")
auth_header_basic = "Basic " + base64.b64encode(f"{user}:{password}".encode()).decode()
auth_header_bearer = f"Bearer {token}" if token else ""

url_re = re.compile(
    r"https?://(?:localhost|127\.0\.0\.1|manticoresearch-manticore)(?::\d+)?"
)

def rewrite_urls(text: str) -> str:
    return url_re.sub(url, text)

def write(path: pathlib.Path, text: str):
    path.write_text(text, encoding="utf-8")

if client == "python":
    p = root / "test" / "test.py"
    text = p.read_text(encoding="utf-8")
    text = rewrite_urls(text)
    text = re.sub(
        r"configuration\s*=\s*manticoresearch\.Configuration\([\s\S]*?\)\n",
        "",
        text,
        count=1,
    )
    if mode == "noauth":
        cfg = f'configuration = manticoresearch.Configuration(\n    host = "{url}"\n)\n'
    elif mode == "basic":
        cfg = (
            f'configuration = manticoresearch.Configuration(\n'
            f'    host = "{url}",\n'
            f'    username = {user!r},\n'
            f'    password = {password!r}\n'
            f')\n'
        )
    else:
        cfg = (
            f'configuration = manticoresearch.Configuration(\n'
            f'    host = "{url}",\n'
            f'    access_token = {token!r}\n'
            f')\n'
        )
    text = text.replace(
        "suite = unittest.TestSuite()",
        cfg + "\nsuite = unittest.TestSuite()",
        1,
    )
    if "configuration = manticoresearch.Configuration" not in text:
        text = cfg + "\n" + text
    write(p, text)

elif client == "python-asyncio":
    p = root / "test" / "test_manual.py"
    text = rewrite_urls(p.read_text(encoding="utf-8"))
    if mode == "noauth":
        repl = (
            f"self.configuration = manticoresearch.Configuration(\n"
            f'            host = "{url}"\n'
            f"        )"
        )
    elif mode == "basic":
        repl = (
            f"self.configuration = manticoresearch.Configuration(\n"
            f'            host = "{url}",\n'
            f"            username = {user!r},\n"
            f"            password = {password!r}\n"
            f"        )"
        )
    else:
        repl = (
            f"self.configuration = manticoresearch.Configuration(\n"
            f'            host = "{url}",\n'
            f"            access_token = {token!r}\n"
            f"        )"
        )
    text = re.sub(
        r"self\.configuration\s*=\s*manticoresearch\.Configuration\([\s\S]*?\)",
        repl,
        text,
        count=1,
    )
    write(p, text)

elif client == "javascript":
    p = root / "test" / "api" / "Manual.spec.js"
    text = p.read_text(encoding="utf-8")
    block = [
        "  beforeEach(function() {",
        "    var client= new Manticoresearch.ApiClient()",
        f"    client.basePath={url!r};",
    ]
    if mode == "basic":
        block += [
            f"    client.authentications.basicAuth.username={user!r};",
            f"    client.authentications.basicAuth.password={password!r};",
        ]
    elif mode == "bearer":
        block += [
            f"    client.authentications.bearerAuth.accessToken={token!r};",
        ]
    block += [
        "    indexApi = new Manticoresearch.IndexApi(client);",
        "    searchApi = new Manticoresearch.SearchApi(client);",
        "    utilsApi = new Manticoresearch.UtilsApi(client);",
        "  });",
    ]
    text = re.sub(
        r"  beforeEach\(function\(\) \{[\s\S]*?\n  \}\);",
        "\n".join(block),
        text,
        count=1,
    )
    write(p, text)

elif client == "typescript":
    p = root / "test" / "manual.test.ts"
    text = rewrite_urls(p.read_text(encoding="utf-8"))
    if mode == "noauth":
        auth_snip = ""
        create = (
            "  const config = Manticoresearch.createConfiguration({\n"
            "    baseServer: serverConfig,\n"
            "  });"
        )
    else:
        header = auth_header_basic if mode == "basic" else auth_header_bearer
        auth_snip = (
            "  const defaultAuth = {\n"
            "    getName() { return 'default'; },\n"
            f"    applySecurityAuthentication(context: any) {{ context.setHeaderParam('Authorization', {header!r}); }}\n"
            "  };\n"
        )
        create = (
            "  const config = Manticoresearch.createConfiguration({\n"
            "    baseServer: serverConfig,\n"
            "    authMethods: { default: defaultAuth },\n"
            "  });"
        )
    text = re.sub(
        r"beforeEach\(function \(\) \{[\s\S]*?\n\}\);",
        "beforeEach(function () {\n"
        f'  const serverConfig = new Manticoresearch.ServerConfiguration("{url}", {{}})\n'
        + auth_snip
        + create
        + "\n"
        "  indexApi = new Manticoresearch.IndexApi(config);\n"
        "  searchApi = new Manticoresearch.SearchApi(config);\n"
        "  utilsApi = new Manticoresearch.UtilsApi(config);\n"
        "});",
        text,
        count=1,
    )
    write(p, text)

elif client == "java":
    for p in (root / "src" / "test").rglob("*ApiTest.java"):
        text = rewrite_urls(p.read_text(encoding="utf-8"))
        text = re.sub(
            r'private static final String BASE_PATH = "[^"]+";',
            f'private static final String BASE_PATH = "{url}";',
            text,
        )
        if mode == "basic":
            inject = (
                "\t\t    client.setBasePath(BASE_PATH);\n"
                f'\t\t    client.setUsername("{user}");\n'
                f'\t\t    client.setPassword("{password}");'
            )
        elif mode == "bearer":
            inject = (
                "\t\t    client.setBasePath(BASE_PATH);\n"
                f'\t\t    client.setBearerToken("{token}");'
            )
        else:
            inject = "\t\t    client.setBasePath(BASE_PATH);"
        text = re.sub(
            r"\t\t    client\.setBasePath\(BASE_PATH\);(?:\n\t\t    client\.set[A-Za-z]+\([^;]*\);)*",
            inject,
            text,
        )
        write(p, text)

elif client == "csharp":
    for p in (root / "src" / "ManticoreSearch.Test").rglob("*Tests.cs"):
        text = rewrite_urls(p.read_text(encoding="utf-8"))
        text = re.sub(
            r'config\.BasePath\s*=\s*"[^"]+";',
            f'config.BasePath = "{url}";',
            text,
        )
        text = re.sub(r"\s*config\.(Username|Password|AccessToken)\s*=\s*[^;]+;\n", "", text)
        extra = ""
        if mode == "basic":
            extra = (
                f'            config.Username = "{user}";\n'
                f'            config.Password = "{password}";\n'
            )
        elif mode == "bearer":
            extra = f'            config.AccessToken = "{token}";\n'
        text = text.replace(
            f'config.BasePath = "{url}";',
            f'config.BasePath = "{url}";\n{extra}',
            1,
        )
        write(p, text)

elif client == "go":
    for p in (root / "test").glob("*.go"):
        text = rewrite_urls(p.read_text(encoding="utf-8"))
        text = re.sub(
            r'configuration\.Servers\[0\]\.URL\s*=\s*"[^"]+"',
            f'configuration.Servers[0].URL = "{url}"',
            text,
        )
        text = re.sub(
            r'\n\tconfiguration\.AddDefaultHeader\("Authorization", "[^"]*"\)\n',
            "\n",
            text,
        )
        if mode != "noauth":
            header = auth_header_basic if mode == "basic" else auth_header_bearer
            text = text.replace(
                f'configuration.Servers[0].URL = "{url}"\n',
                f'configuration.Servers[0].URL = "{url}"\n'
                f'\tconfiguration.AddDefaultHeader("Authorization", "{header}")\n',
                1,
            )
        write(p, text)

elif client == "rust":
    for p in (root / "tests").glob("*.rs"):
        text = p.read_text(encoding="utf-8")
        cfg_lines = [
            "let mut __cfg = Configuration::new();",
            f'__cfg.base_path = "{url}".to_owned();',
        ]
        if mode == "basic":
            cfg_lines.append(
                f'__cfg.basic_auth = Some(("{user}".into(), Some("{password}".into())));'
            )
        elif mode == "bearer":
            cfg_lines.append(f'__cfg.oauth_access_token = Some("{token}".into());')
        cfg_lines.append("let api_config = Arc::new(__cfg);")
        replacement = "\n    ".join(cfg_lines)
        text = text.replace(
            "let api_config = Arc::new(Configuration::new());",
            replacement,
        )
        write(p, text)

elif client == "elixir":
    for env_file in ("config/test.exs", "config/dev.exs", "config/config.exs"):
        p = root / env_file
        if not p.exists():
            continue
        text = rewrite_urls(p.read_text(encoding="utf-8"))
        text = re.sub(
            r'config :manticoresearch,\s*base_url:\s*"[^"]+"',
            f'config :manticoresearch, base_url: "{url}"',
            text,
        )
        if "base_url:" not in text and env_file.endswith("test.exs"):
            text += f'\nconfig :manticoresearch, base_url: "{url}"\n'
        write(p, text)

    conn = root / "lib" / "manticoresearch" / "connection.ex"
    if conn.exists():
        text = conn.read_text(encoding="utf-8")
        text = re.sub(
            r'plug Tesla\.Middleware\.Headers, \[.*?\]\n',
            'plug Tesla.Middleware.Headers, [{"user-agent", "Elixir"}]\n',
            text,
            count=1,
        )
        if mode != "noauth":
            header = auth_header_basic if mode == "basic" else auth_header_bearer
            text = text.replace(
                'plug Tesla.Middleware.Headers, [{"user-agent", "Elixir"}]',
                'plug Tesla.Middleware.Headers, [{"user-agent", "Elixir"}, '
                f'{{"Authorization", {header!r}}}]',
            )
        write(conn, text)

print(f"prepared {client} mode={mode} url={url}", flush=True)
PY
}

# Sync and asyncio clients both import `manticoresearch`; editable installs fight
# over the same package name when the full matrix runs python before python-asyncio.
install_python_test_deps() {
  local variant="$1"
  local -a pip_args=(python3 -m pip install -q)
  if python3 -m pip install --help 2>&1 | grep -q break-system-packages; then
    pip_args+=(--break-system-packages)
  fi
  "${pip_args[@]}" \
    'urllib3>=1.25.3,<3' \
    python-dateutil \
    'pydantic>=2' \
    'typing-extensions>=4.7.1'
  if [[ "$variant" == "asyncio" ]]; then
    "${pip_args[@]}" 'aiohttp>=3.0.0' 'aiohttp-retry>=2.8.3'
  fi
}

run_one_client() {
  local client="$1" mode="$2" url="$3" token="${4:-}"
  local dir_name="${CLIENT_DIRS[$client]}"
  local src="$CLIENTS_SRC/$dir_name"
  local dst="$WORK_ROOT/$client-$mode"

  hr
  log "TEST $client / $mode  →  $url"

  if [[ ! -d "$src" ]]; then
    log "SKIP: missing $src"
    record_result "$client" "$mode" "SKIP" "missing source"
    return 0
  fi

  prepare_workdir "$client" "$mode" "$src" "$dst" "$url" "$token"

  local logfile="$WORK_ROOT/logs/${client}-${mode}.log"
  mkdir -p "$WORK_ROOT/logs"
  set +e
  (
    cd "$dst"
    case "$client" in
      python)
        install_python_test_deps sync
        cd test
        PYTHONPATH="$dst" python3 test.py
        ;;
      python-asyncio)
        install_python_test_deps asyncio
        cd test
        PYTHONPATH="$dst" python3 -m unittest test_manual -v
        ;;
      java)
        mvn -q -DskipTests=false test
        ;;
      javascript)
        npm install --silent
        npx mocha --require @babel/register --timeout 10000 test/api/Manual.spec.js
        ;;
      typescript)
        npm install --silent
        npm test
        ;;
      csharp)
        dotnet test ManticoreSearch.sln --nologo
        ;;
      go)
        # Pull testify for tests and fill go.sum with its transitive deps
        # (go-spew, etc.); without tidy, `go test` fails on missing sum entries.
        go get github.com/stretchr/testify@v1.9.0
        go mod tidy
        go test ./test/ -count=1 -v
        ;;
      rust)
        cargo test --tests -- --nocapture
        ;;
      elixir)
        mix local.hex --force --if-missing
        mix local.rebar --force --if-missing
        mix deps.get
        MIX_ENV=test mix test
        ;;
      *)
        echo "unknown client $client" >&2
        exit 2
        ;;
    esac
  ) >"$logfile" 2>&1
  local rc=$?
  set -e

  if [[ $rc -eq 0 ]]; then
    log "PASS $client / $mode"
    record_result "$client" "$mode" "PASS"
  else
    log "FAIL $client / $mode (exit $rc) — log: $logfile"
    tail -n 40 "$logfile" || true
    record_result "$client" "$mode" "FAIL" "exit $rc"
  fi

  # Always wipe tables so the next client/mode starts from a clean instance.
  local need_auth=0
  [[ "$mode" != "noauth" ]] && need_auth=1
  cleanup_tables "$url" "$need_auth" || true
  return 0
}

print_summary() {
  hr
  log "TEST SUMMARY"
  printf '%-18s %-8s %-6s %s\n' "CLIENT" "MODE" "RESULT" "DETAIL"
  local row client mode status detail
  for row in "${RESULTS[@]:-}"; do
    IFS='|' read -r client mode status detail <<<"$row"
    printf '%-18s %-8s %-6s %s\n' "$client" "$mode" "$status" "$detail"
  done
  hr
  log "Passed: $PASS_COUNT  Failed: $FAIL_COUNT  Skipped: $SKIP_COUNT"
  if [[ "$FAIL_COUNT" -gt 0 ]]; then
    return 1
  fi
  return 0
}

run_matrix() {
  log "Client test matrix starting"
  log "No-auth URL: $MANTICORE_NOAUTH_URL"
  log "Auth URL:    $MANTICORE_AUTH_URL"
  log "Clients:     $(selected_clients | tr '\n' ' ')"
  if [[ -n "${EXCLUDE_CLIENTS:-}" ]]; then
    log "Excluded:    $EXCLUDE_CLIENTS"
  fi

  wait_http "$MANTICORE_NOAUTH_URL" "manticore(noauth)" 0
  wait_http "$MANTICORE_AUTH_URL" "manticore(auth)" 1

  if ! curl -sf -u "${MANTICORE_USER}:${MANTICORE_PASSWORD}" \
      "${MANTICORE_AUTH_URL}/sql?mode=raw" --data-urlencode "query=SHOW VERSION" >/dev/null 2>&1; then
    log "ERROR: auth instance is not ready for admin login."
    log "Host orchestrator should have bootstrapped auth before starting the matrix."
    exit 1
  fi

  local token
  token="$(fetch_bearer_token)"
  log "Fetched bearer token (${#token} chars)"

  cleanup_tables "$MANTICORE_NOAUTH_URL" 0 || true
  cleanup_tables "$MANTICORE_AUTH_URL" 1 || true

  mkdir -p "$WORK_ROOT"
  local client mode url
  while IFS= read -r client; do
    for mode in "${MODES[@]}"; do
      url="$(mode_url "$mode")"
      run_one_client "$client" "$mode" "$url" "$token"
    done
  done < <(selected_clients)

  print_summary
}

host_orchestrate() {
  cd "$ROOT_DIR"

  if [[ ! -d "$ROOT_DIR/out" ]]; then
    log "ERROR: expected generated clients under $ROOT_DIR/out"
    exit 1
  fi

  log "Building client-tests image and starting Manticore services ..."
  docker compose build client-tests
  docker compose up -d manticore manticore-auth

  wait_http "$MANTICORE_NOAUTH_URL" "manticore(noauth)" 0
  wait_http "$MANTICORE_AUTH_URL" "manticore(auth)" 1
  bootstrap_auth

  log "Running test matrix in client-tests container ..."
  set +e
  docker compose --profile test run --rm \
    -e CLIENTS="${CLIENTS:-}" \
    -e EXCLUDE_CLIENTS="${EXCLUDE_CLIENTS:-}" \
    -e MANTICORE_NOAUTH_URL="$MANTICORE_NOAUTH_URL" \
    -e MANTICORE_AUTH_URL="$MANTICORE_AUTH_URL" \
    -e MANTICORE_USER="$MANTICORE_USER" \
    -e MANTICORE_PASSWORD="$MANTICORE_PASSWORD" \
    client-tests
  local rc=$?
  set -e
  exit "$rc"
}

usage() {
  cat <<'EOF'
Usage:
  ./run-client-tests.sh                         Start Manticore services and run all client tests
  ./run-client-tests.sh --matrix                Run the test matrix only (container entrypoint)
  ./run-client-tests.sh --exclude java,rust     Run all clients except the listed ones
  ./run-client-tests.sh --matrix --exclude go
  ./run-client-tests.sh --help

Options:
  --exclude LIST        Comma-separated clients to skip (aliases: js, ts, py, ...)
  --matrix              Run matrix only (no compose up)
  --help                Show this help

Environment:
  CLIENTS               Comma-separated client list (default: all)
                        python,python-asyncio,java,javascript,typescript,csharp,go,rust,elixir
  EXCLUDE_CLIENTS       Comma-separated clients to skip (same as --exclude)
  MANTICORE_NOAUTH_URL  Default http://localhost:19408
  MANTICORE_AUTH_URL    Default http://localhost:19508
  MANTICORE_USER        Default admin
  MANTICORE_PASSWORD    Default TestPass#2026
EOF
}

main() {
  local run_mode=""
  while [[ $# -gt 0 ]]; do
    case "$1" in
      --help|-h)
        usage
        return 0
        ;;
      --matrix)
        run_mode=matrix
        shift
        ;;
      --exclude)
        shift
        if [[ $# -lt 1 || "$1" == --* ]]; then
          log "ERROR: --exclude requires a comma-separated client list"
          exit 2
        fi
        if [[ -n "${EXCLUDE_CLIENTS:-}" ]]; then
          EXCLUDE_CLIENTS="${EXCLUDE_CLIENTS},$1"
        else
          EXCLUDE_CLIENTS="$1"
        fi
        shift
        ;;
      --exclude=*)
        local value="${1#--exclude=}"
        if [[ -z "$value" ]]; then
          log "ERROR: --exclude requires a comma-separated client list"
          exit 2
        fi
        if [[ -n "${EXCLUDE_CLIENTS:-}" ]]; then
          EXCLUDE_CLIENTS="${EXCLUDE_CLIENTS},$value"
        else
          EXCLUDE_CLIENTS="$value"
        fi
        shift
        ;;
      "")
        shift
        ;;
      *)
        log "Unknown argument: $1"
        usage
        exit 2
        ;;
    esac
  done

  case "$run_mode" in
    matrix) run_matrix ;;
    "") host_orchestrate ;;
    *)
      log "Unknown run mode: $run_mode"
      exit 2
      ;;
  esac
}

main "$@"
