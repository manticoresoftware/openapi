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

// checks if the GeoDistance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoDistance{}

// GeoDistance Object to perform geo-distance based filtering on queries
type GeoDistance struct {
	LocationAnchor *GeoDistanceLocationAnchor
	// Field name in the document that contains location data
	LocationSource interface{}
	// Algorithm used to calculate the distance
	DistanceType interface{}
	// The distance from the anchor point to filter results by
	Distance interface{}
	AdditionalProperties map[string]interface{}
}

type _GeoDistance GeoDistance

// NewGeoDistance instantiates a new GeoDistance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoDistance() *GeoDistance {
	this := GeoDistance{}
	return &this
}

// NewGeoDistanceWithDefaults instantiates a new GeoDistance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoDistanceWithDefaults() *GeoDistance {
	this := GeoDistance{}
	return &this
}

// GetLocationAnchor returns the LocationAnchor field value if set, zero value otherwise.
func (o *GeoDistance) GetLocationAnchor() GeoDistanceLocationAnchor {
	if o == nil || IsNil(o.LocationAnchor) {
		var ret GeoDistanceLocationAnchor
		return ret
	}
	return *o.LocationAnchor
}

// GetLocationAnchorOk returns a tuple with the LocationAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoDistance) GetLocationAnchorOk() (*GeoDistanceLocationAnchor, bool) {
	if o == nil || IsNil(o.LocationAnchor) {
		return nil, false
	}
	return o.LocationAnchor, true
}

// HasLocationAnchor returns a boolean if a field has been set.
func (o *GeoDistance) HasLocationAnchor() bool {
	if o != nil && !IsNil(o.LocationAnchor) {
		return true
	}

	return false
}

// SetLocationAnchor gets a reference to the given GeoDistanceLocationAnchor and assigns it to the LocationAnchor field.
func (o *GeoDistance) SetLocationAnchor(v GeoDistanceLocationAnchor) {
	o.LocationAnchor = &v
}

// GetLocationSource returns the LocationSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeoDistance) GetLocationSource() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LocationSource
}

// GetLocationSourceOk returns a tuple with the LocationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoDistance) GetLocationSourceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LocationSource) {
		return nil, false
	}
	return &o.LocationSource, true
}

// HasLocationSource returns a boolean if a field has been set.
func (o *GeoDistance) HasLocationSource() bool {
	if o != nil && !IsNil(o.LocationSource) {
		return true
	}

	return false
}

// SetLocationSource gets a reference to the given interface{} and assigns it to the LocationSource field.
func (o *GeoDistance) SetLocationSource(v interface{}) {
	o.LocationSource = v
}

// GetDistanceType returns the DistanceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeoDistance) GetDistanceType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DistanceType
}

// GetDistanceTypeOk returns a tuple with the DistanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoDistance) GetDistanceTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DistanceType) {
		return nil, false
	}
	return &o.DistanceType, true
}

// HasDistanceType returns a boolean if a field has been set.
func (o *GeoDistance) HasDistanceType() bool {
	if o != nil && !IsNil(o.DistanceType) {
		return true
	}

	return false
}

// SetDistanceType gets a reference to the given interface{} and assigns it to the DistanceType field.
func (o *GeoDistance) SetDistanceType(v interface{}) {
	o.DistanceType = v
}

// GetDistance returns the Distance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeoDistance) GetDistance() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeoDistance) GetDistanceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return &o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *GeoDistance) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given interface{} and assigns it to the Distance field.
func (o *GeoDistance) SetDistance(v interface{}) {
	o.Distance = v
}

func (o GeoDistance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoDistance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocationAnchor) {
		toSerialize["location_anchor"] = o.LocationAnchor
	}
	if o.LocationSource != nil {
		toSerialize["location_source"] = o.LocationSource
	}
	if o.DistanceType != nil {
		toSerialize["distance_type"] = o.DistanceType
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableGeoDistance struct {
	value *GeoDistance
	isSet bool
}

func (v NullableGeoDistance) Get() *GeoDistance {
	return v.value
}

func (v *NullableGeoDistance) Set(val *GeoDistance) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoDistance) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoDistance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoDistance(val *GeoDistance) *NullableGeoDistance {
	return &NullableGeoDistance{value: val, isSet: true}
}

func (v NullableGeoDistance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoDistance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

