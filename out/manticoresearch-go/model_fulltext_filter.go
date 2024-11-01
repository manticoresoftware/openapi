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

// checks if the FulltextFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulltextFilter{}

// FulltextFilter Defines a type of filter for full-text search queries
type FulltextFilter struct {
	// Filter object defining a query string
	QueryString *string
	// Filter object defining a match keyword passed as a string or in a Match object
	Match map[string]interface{}
	// Filter object defining a match phrase
	MatchPhrase map[string]interface{}
	// Filter object to select all documents
	MatchAll map[string]interface{}
}

// NewFulltextFilter instantiates a new FulltextFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulltextFilter() *FulltextFilter {
	this := FulltextFilter{}
	return &this
}

// NewFulltextFilterWithDefaults instantiates a new FulltextFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulltextFilterWithDefaults() *FulltextFilter {
	this := FulltextFilter{}
	return &this
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *FulltextFilter) GetQueryString() string {
	if o == nil || IsNil(o.QueryString) {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulltextFilter) GetQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *FulltextFilter) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *FulltextFilter) SetQueryString(v string) {
	o.QueryString = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *FulltextFilter) GetMatch() map[string]interface{} {
	if o == nil || IsNil(o.Match) {
		var ret map[string]interface{}
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulltextFilter) GetMatchOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Match) {
		return map[string]interface{}{}, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *FulltextFilter) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given map[string]interface{} and assigns it to the Match field.
func (o *FulltextFilter) SetMatch(v map[string]interface{}) {
	o.Match = v
}

// GetMatchPhrase returns the MatchPhrase field value if set, zero value otherwise.
func (o *FulltextFilter) GetMatchPhrase() map[string]interface{} {
	if o == nil || IsNil(o.MatchPhrase) {
		var ret map[string]interface{}
		return ret
	}
	return o.MatchPhrase
}

// GetMatchPhraseOk returns a tuple with the MatchPhrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulltextFilter) GetMatchPhraseOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MatchPhrase) {
		return map[string]interface{}{}, false
	}
	return o.MatchPhrase, true
}

// HasMatchPhrase returns a boolean if a field has been set.
func (o *FulltextFilter) HasMatchPhrase() bool {
	if o != nil && !IsNil(o.MatchPhrase) {
		return true
	}

	return false
}

// SetMatchPhrase gets a reference to the given map[string]interface{} and assigns it to the MatchPhrase field.
func (o *FulltextFilter) SetMatchPhrase(v map[string]interface{}) {
	o.MatchPhrase = v
}

// GetMatchAll returns the MatchAll field value if set, zero value otherwise.
func (o *FulltextFilter) GetMatchAll() map[string]interface{} {
	if o == nil || IsNil(o.MatchAll) {
		var ret map[string]interface{}
		return ret
	}
	return o.MatchAll
}

// GetMatchAllOk returns a tuple with the MatchAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulltextFilter) GetMatchAllOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MatchAll) {
		return map[string]interface{}{}, false
	}
	return o.MatchAll, true
}

// HasMatchAll returns a boolean if a field has been set.
func (o *FulltextFilter) HasMatchAll() bool {
	if o != nil && !IsNil(o.MatchAll) {
		return true
	}

	return false
}

// SetMatchAll gets a reference to the given map[string]interface{} and assigns it to the MatchAll field.
func (o *FulltextFilter) SetMatchAll(v map[string]interface{}) {
	o.MatchAll = v
}

func (o FulltextFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulltextFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueryString) {
		toSerialize["query_string"] = o.QueryString
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.MatchPhrase) {
		toSerialize["match_phrase"] = o.MatchPhrase
	}
	if !IsNil(o.MatchAll) {
		toSerialize["match_all"] = o.MatchAll
	}
	return toSerialize, nil
}

type NullableFulltextFilter struct {
	value *FulltextFilter
	isSet bool
}

func (v NullableFulltextFilter) Get() *FulltextFilter {
	return v.value
}

func (v *NullableFulltextFilter) Set(val *FulltextFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableFulltextFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFulltextFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulltextFilter(val *FulltextFilter) *NullableFulltextFilter {
	return &NullableFulltextFilter{value: val, isSet: true}
}

func (v NullableFulltextFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulltextFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

