#!/usr/bin/env bash
# Build one or more OpenAPI clients via ./build.sh, then test each one.
#
# Usage:
#   ./build-clients.sh
#   ./build-clients.sh --build-version release
#   ./build-clients.sh --clients python,java
#   ./build-clients.sh --build-version dev --manticore-version 18.0.0 --clients python
#
# Defaults:
#   --build-version dev
#   --clients       all clients present under ./out
#
# Flow:
#   1. Start Manticore + client-tests image once (--prepare)
#   2. For each client: ./build.sh ..., optional compat-table update, then tests
#   3. docker compose down on exit
#
# Stops on the first build or test failure.

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT_DIR"

GENERATOR_VERSIONS_FILE="${GENERATOR_VERSIONS_FILE:-$ROOT_DIR/generator-versions}"
VERSIONS_DIR="${VERSIONS_DIR:-$ROOT_DIR/versions}"

DEFAULT_CLIENTS=(
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

declare -A CLIENT_OUT_DIRS=(
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

declare -A CLIENT_README_TEMPLATES=(
  [python]=templates/python/README.mustache
  [python-asyncio]=templates/python-asyncio/README.mustache
  [java]=templates/Java/README.mustache
  [javascript]=templates/Javascript/README.mustache
  [typescript]=templates/typescript/README.mustache
  [csharp]=templates/csharp/README.mustache
  [go]=templates/go/README.mustache
  [rust]=templates/rust/README.mustache
  [elixir]=templates/elixir/README.md.mustache
)

BUILD_VERSION="dev"
CLIENTS_ARG=""
MANTICORE_VERSION=""
SERVICES_STARTED=0

# Prefer the invoking user when this script is run under sudo.
OWNER_UID="${SUDO_UID:-$(id -u)}"
OWNER_GID="${SUDO_GID:-$(id -g)}"
OWNER_SPEC="$OWNER_UID:$OWNER_GID"

usage() {
  cat <<'EOF'
Usage: ./build-clients.sh [options]

  --build-version dev|release   Package branch mode for build.sh (default: dev).
                                Package versions always come from versions/<client>.
  --clients LIST                Comma-separated clients (default: all clients
                                currently under ./out).
  --manticore-version VER       With --build-version dev, update each client's
                                README compatibility table (template + out/) for
                                this Manticore Search version. Ignored for release.
                                If omitted, skip the table update.
  -h, --help                    Show this help.

After each successful build, runs that client's tests via:
  CLIENTS=<client> docker compose --profile test run --rm client-tests

Stops on the first build or test failure.
EOF
}

log() { printf '[%s] %s\n' "$(date '+%H:%M:%S')" "$*"; }
die() { log "ERROR: $*"; exit 1; }

cleanup() {
  local rc=$?
  if [[ "$SERVICES_STARTED" -eq 1 ]]; then
    log "Stopping Manticore / client-test services (docker compose down) ..."
    docker compose -f "$ROOT_DIR/docker-compose.yml" --project-directory "$ROOT_DIR" down || true
  fi
  exit "$rc"
}
trap cleanup EXIT

while [[ $# -gt 0 ]]; do
  case "$1" in
    --build-version)
      [[ $# -ge 2 ]] || die "$1 requires an argument"
      BUILD_VERSION="$2"
      shift 2
      ;;
    --clients)
      [[ $# -ge 2 ]] || die "$1 requires an argument"
      CLIENTS_ARG="$2"
      shift 2
      ;;
    --manticore-version)
      [[ $# -ge 2 ]] || die "$1 requires an argument"
      MANTICORE_VERSION="$2"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      die "unknown argument: $1 (see --help)"
      ;;
  esac
done

case "$BUILD_VERSION" in
  dev|release) ;;
  *) die "--build-version must be 'dev' or 'release' (got: $BUILD_VERSION)" ;;
esac

if [[ -n "$MANTICORE_VERSION" && "$BUILD_VERSION" != "dev" ]]; then
  log "WARNING: --manticore-version is ignored for --build-version $BUILD_VERSION"
  MANTICORE_VERSION=""
fi

[[ -f "$GENERATOR_VERSIONS_FILE" ]] || die "generator versions file not found: $GENERATOR_VERSIONS_FILE"
[[ -f "$ROOT_DIR/build.sh" ]] || die "build.sh not found in $ROOT_DIR"
[[ -f "$ROOT_DIR/run-client-tests.sh" ]] || die "run-client-tests.sh not found in $ROOT_DIR"
[[ -f "$ROOT_DIR/update_compat_table.py" ]] || die "update_compat_table.py not found in $ROOT_DIR"

declare -A GENERATOR_VERSION=()
while IFS= read -r line || [[ -n "$line" ]]; do
  # Skip blank lines and comments.
  [[ -z "$line" || "$line" =~ ^[[:space:]]*# ]] && continue
  local_client=""
  local_gen=""
  read -r local_client local_gen _ <<< "$line"
  [[ -n "$local_client" && -n "$local_gen" ]] \
    || die "invalid line in $GENERATOR_VERSIONS_FILE: $line"
  GENERATOR_VERSION["$local_client"]="$local_gen"
done < "$GENERATOR_VERSIONS_FILE"

parse_client_list() {
  local raw="$1"
  local part
  IFS=',' read -ra parts <<< "$raw"
  for part in "${parts[@]}"; do
    part="${part#"${part%%[![:space:]]*}"}"
    part="${part%"${part##*[![:space:]]}"}"
    [[ -n "$part" ]] || continue
    printf '%s\n' "$part"
  done
}

selected=()
if [[ -z "$CLIENTS_ARG" ]]; then
  selected=("${DEFAULT_CLIENTS[@]}")
else
  while IFS= read -r c; do
    [[ -n "$c" ]] || continue
    selected+=("$c")
  done < <(parse_client_list "$CLIENTS_ARG")
  [[ ${#selected[@]} -gt 0 ]] || die "--clients produced an empty client list"
fi

to_build=()
for c in "${selected[@]}"; do
  [[ -n "${GENERATOR_VERSION[$c]+x}" ]] \
    || die "unknown client '$c' (no entry in $(basename "$GENERATOR_VERSIONS_FILE"))"
  to_build+=("$c")
done

[[ ${#to_build[@]} -gt 0 ]] || die "no clients left to build after filtering"

# Ensure out/ exists before --prepare (image mount / clients check).
mkdir -p "$ROOT_DIR/out"

log "build-version=$BUILD_VERSION clients=${to_build[*]}"
if [[ -n "$MANTICORE_VERSION" ]]; then
  log "manticore-version=$MANTICORE_VERSION (compat tables will be updated after each build)"
fi
log "Preparing test services ..."
# Mark before --prepare so a partial start still gets compose down.
SERVICES_STARTED=1
"$ROOT_DIR/run-client-tests.sh" --prepare

run_client_matrix() {
  local client="$1"
  log "Testing $client ..."
  CLIENTS="$client" docker compose -f "$ROOT_DIR/docker-compose.yml" --project-directory "$ROOT_DIR" \
    --profile test run --rm \
    -e CLIENTS="$client" \
    -e MANTICORE_NOAUTH_URL="${MANTICORE_NOAUTH_URL:-http://localhost:19408}" \
    -e MANTICORE_AUTH_URL="${MANTICORE_AUTH_URL:-http://localhost:19508}" \
    -e MANTICORE_USER="${MANTICORE_USER:-admin}" \
    -e MANTICORE_PASSWORD="${MANTICORE_PASSWORD:-TestPass#2026}" \
    client-tests
}

update_compat_tables() {
  local client="$1" pkg_version="$2" out_path="$3"
  local template_rel="${CLIENT_README_TEMPLATES[$client]:-}"
  [[ -n "$template_rel" ]] || die "no README template mapping for client '$client'"
  local template_path="$ROOT_DIR/$template_rel"
  local out_readme="$out_path/README.md"
  [[ -f "$template_path" ]] || die "README template missing: $template_path"
  [[ -f "$out_readme" ]] || die "generated README missing: $out_readme"

  log "Updating compatibility tables for $client (client=$pkg_version, manticore=$MANTICORE_VERSION)"
  python3 "$ROOT_DIR/update_compat_table.py" \
    --client-version "$pkg_version" \
    --manticore-version "$MANTICORE_VERSION" \
    "$template_path" \
    "$out_readme"
}

for client in "${to_build[@]}"; do
  gen="${GENERATOR_VERSION[$client]:-}"
  [[ -n "$gen" ]] || die "no OpenAPI Generator version for '$client' in $(basename "$GENERATOR_VERSIONS_FILE")"

  version_file="$VERSIONS_DIR/$client"
  if [[ ! -f "$version_file" ]]; then
    die "missing package version file for '$client': $version_file"
  fi
  pkg_version="$(tr -d '[:space:]' < "$version_file")"
  [[ -n "$pkg_version" ]] || die "empty package version in $version_file"

  log "Building $client (generator=$gen, package=$pkg_version, branch=$BUILD_VERSION)"
  "$ROOT_DIR/build.sh" "$client" "$gen" "$BUILD_VERSION"

  out_dir="${CLIENT_OUT_DIRS[$client]:-}"
  [[ -n "$out_dir" ]] || die "no out/ directory mapping for client '$client'"
  out_path="$ROOT_DIR/out/$out_dir"
  [[ -d "$out_path" ]] || die "expected build output missing: $out_path"
  log "Fixing ownership of $out_path ($OWNER_SPEC) ..."
  chown -R "$OWNER_SPEC" "$out_path"

  if [[ -n "$MANTICORE_VERSION" ]]; then
    update_compat_tables "$client" "$pkg_version" "$out_path"
  fi

  log "Finished building $client"

  run_client_matrix "$client"
  log "Finished testing $client"
done

log "All requested clients built and tested successfully."
