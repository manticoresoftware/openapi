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

// checks if the PercolateRequestQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PercolateRequestQuery{}

// PercolateRequestQuery struct for PercolateRequestQuery
type PercolateRequestQuery struct {
	// Object representing the document to percolate
	Percolate map[string]interface{} `json:"percolate"` 
}

type _PercolateRequestQuery PercolateRequestQuery

// NewPercolateRequestQuery instantiates a new PercolateRequestQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercolateRequestQuery(percolate map[string]interface{}) *PercolateRequestQuery {
	this := PercolateRequestQuery{}
	this.Percolate = percolate
	return &this
}

// NewPercolateRequestQueryWithDefaults instantiates a new PercolateRequestQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercolateRequestQueryWithDefaults() *PercolateRequestQuery {
	this := PercolateRequestQuery{}
	return &this
}

// GetPercolate returns the Percolate field value
func (o *PercolateRequestQuery) GetPercolate() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Percolate
}

// GetPercolateOk returns a tuple with the Percolate field value
// and a boolean to check if the value has been set.
func (o *PercolateRequestQuery) GetPercolateOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Percolate, true
}

// SetPercolate sets field value
func (o *PercolateRequestQuery) SetPercolate(v map[string]interface{}) {
	o.Percolate = v
}

func (o PercolateRequestQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PercolateRequestQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["percolate"] = o.Percolate
	return toSerialize, nil
}

type NullablePercolateRequestQuery struct {
	value *PercolateRequestQuery
	isSet bool
}

func (v NullablePercolateRequestQuery) Get() *PercolateRequestQuery {
	return v.value
}

func (v *NullablePercolateRequestQuery) Set(val *PercolateRequestQuery) {
	v.value = val
	v.isSet = true
}

func (v NullablePercolateRequestQuery) IsSet() bool {
	return v.isSet
}

func (v *NullablePercolateRequestQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercolateRequestQuery(val *PercolateRequestQuery) *NullablePercolateRequestQuery {
	return &NullablePercolateRequestQuery{value: val, isSet: true}
}

func (v NullablePercolateRequestQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercolateRequestQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


