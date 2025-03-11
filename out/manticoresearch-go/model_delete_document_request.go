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

// checks if the DeleteDocumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteDocumentRequest{}

// DeleteDocumentRequest Payload for delete request. Documents can be deleted either one by one by specifying the document id or by providing a query object. For more information see  [Delete API](https://manual.manticoresearch.com/Deleting_documents) 
type DeleteDocumentRequest struct {
	// Table name
	Table string `json:"table"` 
	// Cluster name
	Cluster *string `json:"cluster"` 
	// The ID of document for deletion
	Id *int64 `json:"id"` 
	// Defines the criteria to match documents for deletion
	Query map[string]interface{} `json:"query"` 
}

type _DeleteDocumentRequest DeleteDocumentRequest

// NewDeleteDocumentRequest instantiates a new DeleteDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDocumentRequest(table string) *DeleteDocumentRequest {
	this := DeleteDocumentRequest{}
	this.Table = table
	return &this
}

// NewDeleteDocumentRequestWithDefaults instantiates a new DeleteDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDocumentRequestWithDefaults() *DeleteDocumentRequest {
	this := DeleteDocumentRequest{}
	return &this
}

// GetTable returns the Table field value
func (o *DeleteDocumentRequest) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value
func (o *DeleteDocumentRequest) SetTable(v string) {
	o.Table = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DeleteDocumentRequest) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DeleteDocumentRequest) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *DeleteDocumentRequest) SetCluster(v string) {
	o.Cluster = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteDocumentRequest) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteDocumentRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DeleteDocumentRequest) SetId(v int64) {
	o.Id = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DeleteDocumentRequest) GetQuery() map[string]interface{} {
	if o == nil || IsNil(o.Query) {
		var ret map[string]interface{}
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetQueryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Query) {
		return map[string]interface{}{}, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DeleteDocumentRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]interface{} and assigns it to the Query field.
func (o *DeleteDocumentRequest) SetQuery(v map[string]interface{}) {
	o.Query = v
}

func (o DeleteDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["table"] = o.Table
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableDeleteDocumentRequest struct {
	value *DeleteDocumentRequest
	isSet bool
}

func (v NullableDeleteDocumentRequest) Get() *DeleteDocumentRequest {
	return v.value
}

func (v *NullableDeleteDocumentRequest) Set(val *DeleteDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDocumentRequest(val *DeleteDocumentRequest) *NullableDeleteDocumentRequest {
	return &NullableDeleteDocumentRequest{value: val, isSet: true}
}

func (v NullableDeleteDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


