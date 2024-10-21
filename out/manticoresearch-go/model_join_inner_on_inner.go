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

// checks if the JoinInnerOnInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JoinInnerOnInner{}

// JoinInnerOnInner struct for JoinInnerOnInner
type JoinInnerOnInner struct {
	Left *JoinInnerOnInnerLeft `json:"left,omitempty"`
	Operator *string `json:"operator,omitempty"`
	Right *JoinBasicCond `json:"right,omitempty"`
}

// NewJoinInnerOnInner instantiates a new JoinInnerOnInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoinInnerOnInner() *JoinInnerOnInner {
	this := JoinInnerOnInner{}
	return &this
}

// NewJoinInnerOnInnerWithDefaults instantiates a new JoinInnerOnInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoinInnerOnInnerWithDefaults() *JoinInnerOnInner {
	this := JoinInnerOnInner{}
	return &this
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *JoinInnerOnInner) GetLeft() JoinInnerOnInnerLeft {
	if o == nil || IsNil(o.Left) {
		var ret JoinInnerOnInnerLeft
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinInnerOnInner) GetLeftOk() (*JoinInnerOnInnerLeft, bool) {
	if o == nil || IsNil(o.Left) {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *JoinInnerOnInner) HasLeft() bool {
	if o != nil && !IsNil(o.Left) {
		return true
	}

	return false
}

// SetLeft gets a reference to the given JoinInnerOnInnerLeft and assigns it to the Left field.
func (o *JoinInnerOnInner) SetLeft(v JoinInnerOnInnerLeft) {
	o.Left = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *JoinInnerOnInner) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinInnerOnInner) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *JoinInnerOnInner) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *JoinInnerOnInner) SetOperator(v string) {
	o.Operator = &v
}

// GetRight returns the Right field value if set, zero value otherwise.
func (o *JoinInnerOnInner) GetRight() JoinBasicCond {
	if o == nil || IsNil(o.Right) {
		var ret JoinBasicCond
		return ret
	}
	return *o.Right
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinInnerOnInner) GetRightOk() (*JoinBasicCond, bool) {
	if o == nil || IsNil(o.Right) {
		return nil, false
	}
	return o.Right, true
}

// HasRight returns a boolean if a field has been set.
func (o *JoinInnerOnInner) HasRight() bool {
	if o != nil && !IsNil(o.Right) {
		return true
	}

	return false
}

// SetRight gets a reference to the given JoinBasicCond and assigns it to the Right field.
func (o *JoinInnerOnInner) SetRight(v JoinBasicCond) {
	o.Right = &v
}

func (o JoinInnerOnInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JoinInnerOnInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Left) {
		toSerialize["left"] = o.Left
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Right) {
		toSerialize["right"] = o.Right
	}
	return toSerialize, nil
}

type NullableJoinInnerOnInner struct {
	value *JoinInnerOnInner
	isSet bool
}

func (v NullableJoinInnerOnInner) Get() *JoinInnerOnInner {
	return v.value
}

func (v *NullableJoinInnerOnInner) Set(val *JoinInnerOnInner) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinInnerOnInner) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinInnerOnInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinInnerOnInner(val *JoinInnerOnInner) *NullableJoinInnerOnInner {
	return &NullableJoinInnerOnInner{value: val, isSet: true}
}

func (v NullableJoinInnerOnInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinInnerOnInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


