#!/usr/bin/env python3
"""Apply UUID companion field edits to generated Go client (run on post-uint64 tree)."""
from pathlib import Path
import re

root = Path("out/manticoresearch-go")

replacements = [
    ("model_hits_hits.go", "Id", "_id"),
    ("model_insert_document_request.go", "Id", "id"),
    ("model_delete_document_request.go", "Id", "id"),
    ("model_delete_response.go", "Id", "id"),
    ("model_success_response.go", "Id", "id"),
    ("model_update_document_request.go", "Id", "id"),
    ("model_update_response.go", "Id", "id"),
    ("model_knn.go", "DocId", "doc_id"),
]

for fname, field, wire in replacements:
    p = root / fname
    t = p.read_text()
    if field == "Id":
        t = t.replace(f'Id *uint64 `json:"{wire},omitempty"`', 'Id *uint64 `json:"-"`')
        t = t.replace(
            f'\tif !IsNil(o.Id) {{\n\t\ttoSerialize["{wire}"] = o.Id\n\t}}',
            f'\tif w := wireDocumentID(o.Id, o.Uuid); w != nil {{\n\t\ttoSerialize["{wire}"] = w\n\t}}',
        )
    else:
        t = t.replace(f'DocId *uint64 `json:"{wire},omitempty"`', 'DocId *uint64 `json:"-"`')
        t = t.replace(
            f'\tif !IsNil(o.DocId) {{\n\t\ttoSerialize["{wire}"] = o.DocId\n\t}}',
            f'\tif w := wireDocumentID(o.DocId, o.Uuid); w != nil {{\n\t\ttoSerialize["{wire}"] = w\n\t}}',
        )

    m = re.search(r"func New(\w+)\(", t)
    if not m:
        raise SystemExit(f"no New* in {fname}")
    typ = m.group(1)

    if f"func (o *{typ}) UnmarshalJSON" not in t:
        assign = "o.DocId = id" if field == "DocId" else "o.Id = id"
        unmarshal = f"""
func (o *{typ}) UnmarshalJSON(data []byte) error {{
	type Alias {typ}
	aux := struct {{
		RawId json.RawMessage `json:"{wire}"`
		*Alias
	}}{{
		Alias: (*Alias)(o),
	}}
	if err := json.Unmarshal(data, &aux); err != nil {{
		return err
	}}
	id, uuid, err := splitDocumentIDJSON(aux.RawId)
	if err != nil {{
		return err
	}}
	{assign}
	o.Uuid = uuid
	return nil
}}

"""
        idx = t.find(f"type Nullable{typ}")
        if idx < 0:
            t = t + unmarshal
        else:
            t = t[:idx] + unmarshal + t[idx:]

    if '"encoding/json"' not in t:
        t = t.replace("import (\n", 'import (\n\t"encoding/json"\n', 1)

    if f"func (o *{typ}) GetUuid(" not in t:
        helpers = f"""
// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *{typ}) GetUuid() string {{
	if o == nil || IsNil(o.Uuid) {{
		var ret string
		return ret
	}}
	return *o.Uuid
}}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *{typ}) GetUuidOk() (*string, bool) {{
	if o == nil || IsNil(o.Uuid) {{
		return nil, false
	}}
	return o.Uuid, true
}}

// HasUuid returns a boolean if a Uuid field has been set.
func (o *{typ}) HasUuid() bool {{
	if o != nil && !IsNil(o.Uuid) {{
		return true
	}}
	return false
}}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *{typ}) SetUuid(v string) {{
	o.Uuid = &v
}}

"""
        setname = "SetDocId" if field == "DocId" else "SetId"
        m = re.search(
            rf"func \(o \*{re.escape(typ)}\) {setname}\(v uint64\) \{{[^}}]+\}}\n",
            t,
        )
        if m:
            t = t[: m.end()] + helpers + t[m.end() :]
        else:
            print("WARN no SetId insert point", fname)

    p.write_text(t)
    print("updated", fname, typ)

api = root / "api_index.go"
t = api.read_text()
if "PartialReplaceUuid" not in t:
    m = re.search(
        r"func \(a \*IndexAPIService\) PartialReplaceExecute\(r ApiPartialReplaceRequest\) \(\*UpdateResponse, \*http\.Response, error\) \{.*?\n\}\n",
        t,
        re.S,
    )
    if not m:
        raise SystemExit("cannot find PartialReplaceExecute body")
    exec_fn = m.group(0)
    uuid_exec = exec_fn.replace(
        "PartialReplaceExecute(r ApiPartialReplaceRequest)",
        "PartialReplaceUuidExecute(r ApiPartialReplaceUuidRequest)",
    )
    uuid_exec = uuid_exec.replace(
        'IndexAPIService.PartialReplace"', 'IndexAPIService.PartialReplaceUuid"'
    )
    uuid_exec = uuid_exec.replace(
        'parameterValueToString(r.id, "id")', 'parameterValueToString(r.uuid, "id")'
    )
    block = (
        """
type ApiPartialReplaceUuidRequest struct {
	ctx context.Context
	ApiService *IndexAPIService
	table string
	uuid string
	replaceDocumentRequest *ReplaceDocumentRequest
}

func (r ApiPartialReplaceUuidRequest) ReplaceDocumentRequest(replaceDocumentRequest ReplaceDocumentRequest) ApiPartialReplaceUuidRequest {
	r.replaceDocumentRequest = &replaceDocumentRequest
	return r
}

func (r ApiPartialReplaceUuidRequest) Execute() (*UpdateResponse, *http.Response, error) {
	return r.ApiService.PartialReplaceUuidExecute(r)
}

// PartialReplaceUuid Partially replaces a document in a table by UUID
func (a *IndexAPIService) PartialReplaceUuid(ctx context.Context, table string, uuid string) ApiPartialReplaceUuidRequest {
	return ApiPartialReplaceUuidRequest{
		ApiService: a,
		ctx: ctx,
		table: table,
		uuid: uuid,
	}
}

"""
        + uuid_exec
        + "\n"
    )
    t = t[: m.end()] + block + t[m.end() :]
    api.write_text(t)
    print("api PartialReplaceUuid added")

print("done")
