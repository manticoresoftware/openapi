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
	"bytes"
	"fmt"
)

// checks if the AggDateHistogram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggDateHistogram{}

// AggDateHistogram Object to use histograms in aggregation, i.e., grouping search results by histogram values
type AggDateHistogram struct {
	// Field to group by
	Field string `json:"field"`
	// Interval of the histogram values
	Interval int32 `json:"interval"`
	// Offset of the histogram values. Default value is 0.
	Offset *int32 `json:"offset,omitempty"`
	// Flag that defines if a search response will be a dictionary with the bucket keys. Default value is false.
	Keyed *bool `json:"keyed,omitempty"`
}

type _AggDateHistogram AggDateHistogram

// NewAggDateHistogram instantiates a new AggDateHistogram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggDateHistogram(field string, interval int32) *AggDateHistogram {
	this := AggDateHistogram{}
	this.Field = field
	this.Interval = interval
	return &this
}

// NewAggDateHistogramWithDefaults instantiates a new AggDateHistogram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggDateHistogramWithDefaults() *AggDateHistogram {
	this := AggDateHistogram{}
	return &this
}

// GetField returns the Field field value
func (o *AggDateHistogram) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *AggDateHistogram) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *AggDateHistogram) SetField(v string) {
	o.Field = v
}

// GetInterval returns the Interval field value
func (o *AggDateHistogram) GetInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *AggDateHistogram) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *AggDateHistogram) SetInterval(v int32) {
	o.Interval = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *AggDateHistogram) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggDateHistogram) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *AggDateHistogram) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *AggDateHistogram) SetOffset(v int32) {
	o.Offset = &v
}

// GetKeyed returns the Keyed field value if set, zero value otherwise.
func (o *AggDateHistogram) GetKeyed() bool {
	if o == nil || IsNil(o.Keyed) {
		var ret bool
		return ret
	}
	return *o.Keyed
}

// GetKeyedOk returns a tuple with the Keyed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggDateHistogram) GetKeyedOk() (*bool, bool) {
	if o == nil || IsNil(o.Keyed) {
		return nil, false
	}
	return o.Keyed, true
}

// HasKeyed returns a boolean if a field has been set.
func (o *AggDateHistogram) HasKeyed() bool {
	if o != nil && !IsNil(o.Keyed) {
		return true
	}

	return false
}

// SetKeyed gets a reference to the given bool and assigns it to the Keyed field.
func (o *AggDateHistogram) SetKeyed(v bool) {
	o.Keyed = &v
}

func (o AggDateHistogram) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggDateHistogram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["interval"] = o.Interval
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Keyed) {
		toSerialize["keyed"] = o.Keyed
	}
	return toSerialize, nil
}

func (o *AggDateHistogram) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
		"interval",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAggDateHistogram := _AggDateHistogram{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAggDateHistogram)

	if err != nil {
		return err
	}

	*o = AggDateHistogram(varAggDateHistogram)

	return err
}

type NullableAggDateHistogram struct {
	value *AggDateHistogram
	isSet bool
}

func (v NullableAggDateHistogram) Get() *AggDateHistogram {
	return v.value
}

func (v *NullableAggDateHistogram) Set(val *AggDateHistogram) {
	v.value = val
	v.isSet = true
}

func (v NullableAggDateHistogram) IsSet() bool {
	return v.isSet
}

func (v *NullableAggDateHistogram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggDateHistogram(val *AggDateHistogram) *NullableAggDateHistogram {
	return &NullableAggDateHistogram{value: val, isSet: true}
}

func (v NullableAggDateHistogram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggDateHistogram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


