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

// checks if the Aggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Aggregation{}

// Aggregation struct for Aggregation
type Aggregation struct {
	Terms *AggTerms
	Sort []interface{}
	Composite *AggComposite
}

// NewAggregation instantiates a new Aggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregation() *Aggregation {
	this := Aggregation{}
	return &this
}

// NewAggregationWithDefaults instantiates a new Aggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationWithDefaults() *Aggregation {
	this := Aggregation{}
	return &this
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *Aggregation) GetTerms() AggTerms {
	if o == nil || IsNil(o.Terms) {
		var ret AggTerms
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregation) GetTermsOk() (*AggTerms, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *Aggregation) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given AggTerms and assigns it to the Terms field.
func (o *Aggregation) SetTerms(v AggTerms) {
	o.Terms = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *Aggregation) GetSort() []interface{} {
	if o == nil || IsNil(o.Sort) {
		var ret []interface{}
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregation) GetSortOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *Aggregation) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []interface{} and assigns it to the Sort field.
func (o *Aggregation) SetSort(v []interface{}) {
	o.Sort = v
}

// GetComposite returns the Composite field value if set, zero value otherwise.
func (o *Aggregation) GetComposite() AggComposite {
	if o == nil || IsNil(o.Composite) {
		var ret AggComposite
		return ret
	}
	return *o.Composite
}

// GetCompositeOk returns a tuple with the Composite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregation) GetCompositeOk() (*AggComposite, bool) {
	if o == nil || IsNil(o.Composite) {
		return nil, false
	}
	return o.Composite, true
}

// HasComposite returns a boolean if a field has been set.
func (o *Aggregation) HasComposite() bool {
	if o != nil && !IsNil(o.Composite) {
		return true
	}

	return false
}

// SetComposite gets a reference to the given AggComposite and assigns it to the Composite field.
func (o *Aggregation) SetComposite(v AggComposite) {
	o.Composite = &v
}

func (o Aggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Aggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Composite) {
		toSerialize["composite"] = o.Composite
	}
	return toSerialize, nil
}

type NullableAggregation struct {
	value *Aggregation
	isSet bool
}

func (v NullableAggregation) Get() *Aggregation {
	return v.value
}

func (v *NullableAggregation) Set(val *Aggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregation(val *Aggregation) *NullableAggregation {
	return &NullableAggregation{value: val, isSet: true}
}

func (v NullableAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

