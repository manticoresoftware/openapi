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

// checks if the GeoFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoFilter{}

// GeoFilter struct for GeoFilter
type GeoFilter struct {
	GeoDistance *GeoFilterGeoDistance `json:"geo_distance,omitempty"`
}

// NewGeoFilter instantiates a new GeoFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoFilter() *GeoFilter {
	this := GeoFilter{}
	return &this
}

// NewGeoFilterWithDefaults instantiates a new GeoFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoFilterWithDefaults() *GeoFilter {
	this := GeoFilter{}
	return &this
}

// GetGeoDistance returns the GeoDistance field value if set, zero value otherwise.
func (o *GeoFilter) GetGeoDistance() GeoFilterGeoDistance {
	if o == nil || IsNil(o.GeoDistance) {
		var ret GeoFilterGeoDistance
		return ret
	}
	return *o.GeoDistance
}

// GetGeoDistanceOk returns a tuple with the GeoDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoFilter) GetGeoDistanceOk() (*GeoFilterGeoDistance, bool) {
	if o == nil || IsNil(o.GeoDistance) {
		return nil, false
	}
	return o.GeoDistance, true
}

// HasGeoDistance returns a boolean if a field has been set.
func (o *GeoFilter) HasGeoDistance() bool {
	if o != nil && !IsNil(o.GeoDistance) {
		return true
	}

	return false
}

// SetGeoDistance gets a reference to the given GeoFilterGeoDistance and assigns it to the GeoDistance field.
func (o *GeoFilter) SetGeoDistance(v GeoFilterGeoDistance) {
	o.GeoDistance = &v
}

func (o GeoFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeoDistance) {
		toSerialize["geo_distance"] = o.GeoDistance
	}
	return toSerialize, nil
}

type NullableGeoFilter struct {
	value *GeoFilter
	isSet bool
}

func (v NullableGeoFilter) Get() *GeoFilter {
	return v.value
}

func (v *NullableGeoFilter) Set(val *GeoFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoFilter(val *GeoFilter) *NullableGeoFilter {
	return &NullableGeoFilter{value: val, isSet: true}
}

func (v NullableGeoFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


