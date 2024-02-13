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
)

// checks if the AggregationTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationTerms{}

// AggregationTerms struct for AggregationTerms
type AggregationTerms struct {
	// Attribute Name to Aggregate
	Field *string `json:"field,omitempty"`
	// Maximum Number of Buckets in the Result
	Size *int32 `json:"size,omitempty"`
}

// NewAggregationTerms instantiates a new AggregationTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationTerms() *AggregationTerms {
	this := AggregationTerms{}
	return &this
}

// NewAggregationTermsWithDefaults instantiates a new AggregationTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationTermsWithDefaults() *AggregationTerms {
	this := AggregationTerms{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *AggregationTerms) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationTerms) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *AggregationTerms) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *AggregationTerms) SetField(v string) {
	o.Field = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *AggregationTerms) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationTerms) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *AggregationTerms) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *AggregationTerms) SetSize(v int32) {
	o.Size = &v
}

func (o AggregationTerms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregationTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableAggregationTerms struct {
	value *AggregationTerms
	isSet bool
}

func (v NullableAggregationTerms) Get() *AggregationTerms {
	return v.value
}

func (v *NullableAggregationTerms) Set(val *AggregationTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationTerms(val *AggregationTerms) *NullableAggregationTerms {
	return &NullableAggregationTerms{value: val, isSet: true}
}

func (v NullableAggregationTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


