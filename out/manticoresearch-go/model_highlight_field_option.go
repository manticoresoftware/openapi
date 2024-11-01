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

// checks if the HighlightFieldOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighlightFieldOption{}

// HighlightFieldOption Options for controlling the behavior of highlighting on a per-field basis
type HighlightFieldOption struct {
	// Maximum size of the text fragments in highlighted snippets per field
	FragmentSize *int32
	// Maximum size of snippets per field
	Limit *int32
	// Maximum number of snippets per field
	LimitSnippets *int32
	// Maximum number of words per field
	LimitWords *int32
	// Total number of highlighted fragments per field
	NumberOfFragments *int32
}

// NewHighlightFieldOption instantiates a new HighlightFieldOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighlightFieldOption() *HighlightFieldOption {
	this := HighlightFieldOption{}
	return &this
}

// NewHighlightFieldOptionWithDefaults instantiates a new HighlightFieldOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighlightFieldOptionWithDefaults() *HighlightFieldOption {
	this := HighlightFieldOption{}
	return &this
}

// GetFragmentSize returns the FragmentSize field value if set, zero value otherwise.
func (o *HighlightFieldOption) GetFragmentSize() int32 {
	if o == nil || IsNil(o.FragmentSize) {
		var ret int32
		return ret
	}
	return *o.FragmentSize
}

// GetFragmentSizeOk returns a tuple with the FragmentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighlightFieldOption) GetFragmentSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FragmentSize) {
		return nil, false
	}
	return o.FragmentSize, true
}

// HasFragmentSize returns a boolean if a field has been set.
func (o *HighlightFieldOption) HasFragmentSize() bool {
	if o != nil && !IsNil(o.FragmentSize) {
		return true
	}

	return false
}

// SetFragmentSize gets a reference to the given int32 and assigns it to the FragmentSize field.
func (o *HighlightFieldOption) SetFragmentSize(v int32) {
	o.FragmentSize = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *HighlightFieldOption) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighlightFieldOption) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *HighlightFieldOption) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *HighlightFieldOption) SetLimit(v int32) {
	o.Limit = &v
}

// GetLimitSnippets returns the LimitSnippets field value if set, zero value otherwise.
func (o *HighlightFieldOption) GetLimitSnippets() int32 {
	if o == nil || IsNil(o.LimitSnippets) {
		var ret int32
		return ret
	}
	return *o.LimitSnippets
}

// GetLimitSnippetsOk returns a tuple with the LimitSnippets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighlightFieldOption) GetLimitSnippetsOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitSnippets) {
		return nil, false
	}
	return o.LimitSnippets, true
}

// HasLimitSnippets returns a boolean if a field has been set.
func (o *HighlightFieldOption) HasLimitSnippets() bool {
	if o != nil && !IsNil(o.LimitSnippets) {
		return true
	}

	return false
}

// SetLimitSnippets gets a reference to the given int32 and assigns it to the LimitSnippets field.
func (o *HighlightFieldOption) SetLimitSnippets(v int32) {
	o.LimitSnippets = &v
}

// GetLimitWords returns the LimitWords field value if set, zero value otherwise.
func (o *HighlightFieldOption) GetLimitWords() int32 {
	if o == nil || IsNil(o.LimitWords) {
		var ret int32
		return ret
	}
	return *o.LimitWords
}

// GetLimitWordsOk returns a tuple with the LimitWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighlightFieldOption) GetLimitWordsOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitWords) {
		return nil, false
	}
	return o.LimitWords, true
}

// HasLimitWords returns a boolean if a field has been set.
func (o *HighlightFieldOption) HasLimitWords() bool {
	if o != nil && !IsNil(o.LimitWords) {
		return true
	}

	return false
}

// SetLimitWords gets a reference to the given int32 and assigns it to the LimitWords field.
func (o *HighlightFieldOption) SetLimitWords(v int32) {
	o.LimitWords = &v
}

// GetNumberOfFragments returns the NumberOfFragments field value if set, zero value otherwise.
func (o *HighlightFieldOption) GetNumberOfFragments() int32 {
	if o == nil || IsNil(o.NumberOfFragments) {
		var ret int32
		return ret
	}
	return *o.NumberOfFragments
}

// GetNumberOfFragmentsOk returns a tuple with the NumberOfFragments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighlightFieldOption) GetNumberOfFragmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfFragments) {
		return nil, false
	}
	return o.NumberOfFragments, true
}

// HasNumberOfFragments returns a boolean if a field has been set.
func (o *HighlightFieldOption) HasNumberOfFragments() bool {
	if o != nil && !IsNil(o.NumberOfFragments) {
		return true
	}

	return false
}

// SetNumberOfFragments gets a reference to the given int32 and assigns it to the NumberOfFragments field.
func (o *HighlightFieldOption) SetNumberOfFragments(v int32) {
	o.NumberOfFragments = &v
}

func (o HighlightFieldOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighlightFieldOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FragmentSize) {
		toSerialize["fragment_size"] = o.FragmentSize
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.LimitSnippets) {
		toSerialize["limit_snippets"] = o.LimitSnippets
	}
	if !IsNil(o.LimitWords) {
		toSerialize["limit_words"] = o.LimitWords
	}
	if !IsNil(o.NumberOfFragments) {
		toSerialize["number_of_fragments"] = o.NumberOfFragments
	}
	return toSerialize, nil
}

type NullableHighlightFieldOption struct {
	value *HighlightFieldOption
	isSet bool
}

func (v NullableHighlightFieldOption) Get() *HighlightFieldOption {
	return v.value
}

func (v *NullableHighlightFieldOption) Set(val *HighlightFieldOption) {
	v.value = val
	v.isSet = true
}

func (v NullableHighlightFieldOption) IsSet() bool {
	return v.isSet
}

func (v *NullableHighlightFieldOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighlightFieldOption(val *HighlightFieldOption) *NullableHighlightFieldOption {
	return &NullableHighlightFieldOption{value: val, isSet: true}
}

func (v NullableHighlightFieldOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighlightFieldOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

