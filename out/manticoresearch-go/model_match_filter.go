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

// checks if the MatchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchFilter{}

// MatchFilter struct for MatchFilter
type MatchFilter struct {
	Match *MatchFilterMatch `json:"match,omitempty"`
}

// NewMatchFilter instantiates a new MatchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchFilter() *MatchFilter {
	this := MatchFilter{}
	return &this
}

// NewMatchFilterWithDefaults instantiates a new MatchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchFilterWithDefaults() *MatchFilter {
	this := MatchFilter{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *MatchFilter) GetMatch() MatchFilterMatch {
	if o == nil || IsNil(o.Match) {
		var ret MatchFilterMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchFilter) GetMatchOk() (*MatchFilterMatch, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *MatchFilter) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given MatchFilterMatch and assigns it to the Match field.
func (o *MatchFilter) SetMatch(v MatchFilterMatch) {
	o.Match = &v
}

func (o MatchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	return toSerialize, nil
}

type NullableMatchFilter struct {
	value *MatchFilter
	isSet bool
}

func (v NullableMatchFilter) Get() *MatchFilter {
	return v.value
}

func (v *NullableMatchFilter) Set(val *MatchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchFilter(val *MatchFilter) *NullableMatchFilter {
	return &NullableMatchFilter{value: val, isSet: true}
}

func (v NullableMatchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


