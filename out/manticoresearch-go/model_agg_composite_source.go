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

// checks if the AggCompositeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggCompositeSource{}

// AggCompositeSource Object containing terms used for composite aggregation.
type AggCompositeSource struct {
	Terms AggCompositeTerm
}

type _AggCompositeSource AggCompositeSource

// NewAggCompositeSource instantiates a new AggCompositeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggCompositeSource(terms AggCompositeTerm) *AggCompositeSource {
	this := AggCompositeSource{}
	this.Terms = terms
	return &this
}

// NewAggCompositeSourceWithDefaults instantiates a new AggCompositeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggCompositeSourceWithDefaults() *AggCompositeSource {
	this := AggCompositeSource{}
	return &this
}

// GetTerms returns the Terms field value
func (o *AggCompositeSource) GetTerms() AggCompositeTerm {
	if o == nil {
		var ret AggCompositeTerm
		return ret
	}

	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value
// and a boolean to check if the value has been set.
func (o *AggCompositeSource) GetTermsOk() (*AggCompositeTerm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terms, true
}

// SetTerms sets field value
func (o *AggCompositeSource) SetTerms(v AggCompositeTerm) {
	o.Terms = v
}

func (o AggCompositeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggCompositeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["terms"] = o.Terms
	return toSerialize, nil
}

type NullableAggCompositeSource struct {
	value *AggCompositeSource
	isSet bool
}

func (v NullableAggCompositeSource) Get() *AggCompositeSource {
	return v.value
}

func (v *NullableAggCompositeSource) Set(val *AggCompositeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableAggCompositeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableAggCompositeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggCompositeSource(val *AggCompositeSource) *NullableAggCompositeSource {
	return &NullableAggCompositeSource{value: val, isSet: true}
}

func (v NullableAggCompositeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggCompositeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


