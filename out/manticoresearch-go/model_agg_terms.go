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

// checks if the AggTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggTerms{}

// AggTerms Object containing term fields to aggregate on
type AggTerms struct {
	// Name of attribute to aggregate by
	Field string `json:"field"` 
	// Maximum number of buckets in the result
	Size *int32 `json:"size"` 
}

type _AggTerms AggTerms

// NewAggTerms instantiates a new AggTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggTerms(field string) *AggTerms {
	this := AggTerms{}
	this.Field = field
	return &this
}

// NewAggTermsWithDefaults instantiates a new AggTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggTermsWithDefaults() *AggTerms {
	this := AggTerms{}
	return &this
}

// GetField returns the Field field value
func (o *AggTerms) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *AggTerms) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *AggTerms) SetField(v string) {
	o.Field = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *AggTerms) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggTerms) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *AggTerms) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *AggTerms) SetSize(v int32) {
	o.Size = &v
}

func (o AggTerms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableAggTerms struct {
	value *AggTerms
	isSet bool
}

func (v NullableAggTerms) Get() *AggTerms {
	return v.value
}

func (v *NullableAggTerms) Set(val *AggTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableAggTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableAggTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggTerms(val *AggTerms) *NullableAggTerms {
	return &NullableAggTerms{value: val, isSet: true}
}

func (v NullableAggTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


