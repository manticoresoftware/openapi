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

// checks if the QueryStringFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryStringFilter{}

// QueryStringFilter struct for QueryStringFilter
type QueryStringFilter struct {
	QueryString *string `json:"query_string,omitempty"`
}

// NewQueryStringFilter instantiates a new QueryStringFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryStringFilter() *QueryStringFilter {
	this := QueryStringFilter{}
	return &this
}

// NewQueryStringFilterWithDefaults instantiates a new QueryStringFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryStringFilterWithDefaults() *QueryStringFilter {
	this := QueryStringFilter{}
	return &this
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *QueryStringFilter) GetQueryString() string {
	if o == nil || IsNil(o.QueryString) {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStringFilter) GetQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *QueryStringFilter) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *QueryStringFilter) SetQueryString(v string) {
	o.QueryString = &v
}

func (o QueryStringFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryStringFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueryString) {
		toSerialize["query_string"] = o.QueryString
	}
	return toSerialize, nil
}

type NullableQueryStringFilter struct {
	value *QueryStringFilter
	isSet bool
}

func (v NullableQueryStringFilter) Get() *QueryStringFilter {
	return v.value
}

func (v *NullableQueryStringFilter) Set(val *QueryStringFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryStringFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryStringFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryStringFilter(val *QueryStringFilter) *NullableQueryStringFilter {
	return &NullableQueryStringFilter{value: val, isSet: true}
}

func (v NullableQueryStringFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryStringFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


