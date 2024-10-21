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
)

// checks if the BoolFilterBool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoolFilterBool{}

// BoolFilterBool struct for BoolFilterBool
type BoolFilterBool struct {
	Must []QueryFilter `json:"must,omitempty"`
	MustNot []QueryFilter `json:"must_not,omitempty"`
	Should []QueryFilter `json:"should,omitempty"`
}

// NewBoolFilterBool instantiates a new BoolFilterBool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoolFilterBool() *BoolFilterBool {
	this := BoolFilterBool{}
	return &this
}

// NewBoolFilterBoolWithDefaults instantiates a new BoolFilterBool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoolFilterBoolWithDefaults() *BoolFilterBool {
	this := BoolFilterBool{}
	return &this
}

// GetMust returns the Must field value if set, zero value otherwise.
func (o *BoolFilterBool) GetMust() []QueryFilter {
	if o == nil || IsNil(o.Must) {
		var ret []QueryFilter
		return ret
	}
	return o.Must
}

// GetMustOk returns a tuple with the Must field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoolFilterBool) GetMustOk() ([]QueryFilter, bool) {
	if o == nil || IsNil(o.Must) {
		return nil, false
	}
	return o.Must, true
}

// HasMust returns a boolean if a field has been set.
func (o *BoolFilterBool) HasMust() bool {
	if o != nil && !IsNil(o.Must) {
		return true
	}

	return false
}

// SetMust gets a reference to the given []QueryFilter and assigns it to the Must field.
func (o *BoolFilterBool) SetMust(v []QueryFilter) {
	o.Must = v
}

// GetMustNot returns the MustNot field value if set, zero value otherwise.
func (o *BoolFilterBool) GetMustNot() []QueryFilter {
	if o == nil || IsNil(o.MustNot) {
		var ret []QueryFilter
		return ret
	}
	return o.MustNot
}

// GetMustNotOk returns a tuple with the MustNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoolFilterBool) GetMustNotOk() ([]QueryFilter, bool) {
	if o == nil || IsNil(o.MustNot) {
		return nil, false
	}
	return o.MustNot, true
}

// HasMustNot returns a boolean if a field has been set.
func (o *BoolFilterBool) HasMustNot() bool {
	if o != nil && !IsNil(o.MustNot) {
		return true
	}

	return false
}

// SetMustNot gets a reference to the given []QueryFilter and assigns it to the MustNot field.
func (o *BoolFilterBool) SetMustNot(v []QueryFilter) {
	o.MustNot = v
}

// GetShould returns the Should field value if set, zero value otherwise.
func (o *BoolFilterBool) GetShould() []QueryFilter {
	if o == nil || IsNil(o.Should) {
		var ret []QueryFilter
		return ret
	}
	return o.Should
}

// GetShouldOk returns a tuple with the Should field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoolFilterBool) GetShouldOk() ([]QueryFilter, bool) {
	if o == nil || IsNil(o.Should) {
		return nil, false
	}
	return o.Should, true
}

// HasShould returns a boolean if a field has been set.
func (o *BoolFilterBool) HasShould() bool {
	if o != nil && !IsNil(o.Should) {
		return true
	}

	return false
}

// SetShould gets a reference to the given []QueryFilter and assigns it to the Should field.
func (o *BoolFilterBool) SetShould(v []QueryFilter) {
	o.Should = v
}

func (o BoolFilterBool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoolFilterBool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Must) {
		toSerialize["must"] = o.Must
	}
	if !IsNil(o.MustNot) {
		toSerialize["must_not"] = o.MustNot
	}
	if !IsNil(o.Should) {
		toSerialize["should"] = o.Should
	}
	return toSerialize, nil
}

type NullableBoolFilterBool struct {
	value *BoolFilterBool
	isSet bool
}

func (v NullableBoolFilterBool) Get() *BoolFilterBool {
	return v.value
}

func (v *NullableBoolFilterBool) Set(val *BoolFilterBool) {
	v.value = val
	v.isSet = true
}

func (v NullableBoolFilterBool) IsSet() bool {
	return v.isSet
}

func (v *NullableBoolFilterBool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoolFilterBool(val *BoolFilterBool) *NullableBoolFilterBool {
	return &NullableBoolFilterBool{value: val, isSet: true}
}

func (v NullableBoolFilterBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoolFilterBool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


