/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 3.3.1
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDocumentRequest{}

// UpdateDocumentRequest Payload for update document
type UpdateDocumentRequest struct {
	Index string `json:"index"`
	// Index name
	Doc map[string]interface{} `json:"doc"`
	// Document ID
	Id *int64 `json:"id,omitempty"`
	// Query tree object
	Query map[string]interface{} `json:"query,omitempty"`
}

type _UpdateDocumentRequest UpdateDocumentRequest

// NewUpdateDocumentRequest instantiates a new UpdateDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDocumentRequest(index string, doc map[string]interface{}) *UpdateDocumentRequest {
	this := UpdateDocumentRequest{}
	this.Index = index
	this.Doc = doc
	return &this
}

// NewUpdateDocumentRequestWithDefaults instantiates a new UpdateDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDocumentRequestWithDefaults() *UpdateDocumentRequest {
	this := UpdateDocumentRequest{}
	return &this
}

// GetIndex returns the Index field value
func (o *UpdateDocumentRequest) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *UpdateDocumentRequest) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *UpdateDocumentRequest) SetIndex(v string) {
	o.Index = v
}

// GetDoc returns the Doc field value
func (o *UpdateDocumentRequest) GetDoc() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Doc
}

// GetDocOk returns a tuple with the Doc field value
// and a boolean to check if the value has been set.
func (o *UpdateDocumentRequest) GetDocOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Doc, true
}

// SetDoc sets field value
func (o *UpdateDocumentRequest) SetDoc(v map[string]interface{}) {
	o.Doc = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateDocumentRequest) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDocumentRequest) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateDocumentRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UpdateDocumentRequest) SetId(v int64) {
	o.Id = &v
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDocumentRequest) GetQuery() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDocumentRequest) GetQueryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Query) {
		return map[string]interface{}{}, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *UpdateDocumentRequest) HasQuery() bool {
	if o != nil && IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]interface{} and assigns it to the Query field.
func (o *UpdateDocumentRequest) SetQuery(v map[string]interface{}) {
	o.Query = v
}

func (o UpdateDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["doc"] = o.Doc
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

func (o *UpdateDocumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"doc",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateDocumentRequest := _UpdateDocumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateDocumentRequest)

	if err != nil {
		return err
	}

	*o = UpdateDocumentRequest(varUpdateDocumentRequest)

	return err
}

type NullableUpdateDocumentRequest struct {
	value *UpdateDocumentRequest
	isSet bool
}

func (v NullableUpdateDocumentRequest) Get() *UpdateDocumentRequest {
	return v.value
}

func (v *NullableUpdateDocumentRequest) Set(val *UpdateDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDocumentRequest(val *UpdateDocumentRequest) *NullableUpdateDocumentRequest {
	return &NullableUpdateDocumentRequest{value: val, isSet: true}
}

func (v NullableUpdateDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

