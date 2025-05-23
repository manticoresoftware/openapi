/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 5.0.0
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	_"bytes"
	_"fmt"
)

// checks if the ReplaceDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceDocumentRequest{}

// ReplaceDocumentRequest Object containing the document data for replacing an existing document in a table.
type ReplaceDocumentRequest struct {
	// Object containing the new document data to replace the existing one.
	Doc map[string]interface{} `json:"doc"` 
}

type _ReplaceDocumentRequest ReplaceDocumentRequest

// NewReplaceDocumentRequest instantiates a new ReplaceDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceDocumentRequest(doc map[string]interface{}) *ReplaceDocumentRequest {
	this := ReplaceDocumentRequest{}
	this.Doc = doc
	return &this
}

// NewReplaceDocumentRequestWithDefaults instantiates a new ReplaceDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceDocumentRequestWithDefaults() *ReplaceDocumentRequest {
	this := ReplaceDocumentRequest{}
	return &this
}

// GetDoc returns the Doc field value
func (o *ReplaceDocumentRequest) GetDoc() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Doc
}

// GetDocOk returns a tuple with the Doc field value
// and a boolean to check if the value has been set.
func (o *ReplaceDocumentRequest) GetDocOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Doc, true
}

// SetDoc sets field value
func (o *ReplaceDocumentRequest) SetDoc(v map[string]interface{}) {
	o.Doc = v
}

func (o ReplaceDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["doc"] = o.Doc
	return toSerialize, nil
}

type NullableReplaceDocumentRequest struct {
	value *ReplaceDocumentRequest
	isSet bool
}

func (v NullableReplaceDocumentRequest) Get() *ReplaceDocumentRequest {
	return v.value
}

func (v *NullableReplaceDocumentRequest) Set(val *ReplaceDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceDocumentRequest(val *ReplaceDocumentRequest) *NullableReplaceDocumentRequest {
	return &NullableReplaceDocumentRequest{value: val, isSet: true}
}

func (v NullableReplaceDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


