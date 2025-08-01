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

// checks if the MatchAll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchAll{}

// MatchAll Filter helper object defining the 'match all'' condition
type MatchAll struct {
	All string `json:"_all"`
}

type _MatchAll MatchAll

// NewMatchAll instantiates a new MatchAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchAll(all string) *MatchAll {
	this := MatchAll{}
	this.All = all
	return &this
}

// NewMatchAllWithDefaults instantiates a new MatchAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchAllWithDefaults() *MatchAll {
	this := MatchAll{}
	return &this
}

// GetAll returns the All field value
func (o *MatchAll) GetAll() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.All
}

// GetAllOk returns a tuple with the All field value
// and a boolean to check if the value has been set.
func (o *MatchAll) GetAllOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.All, true
}

// SetAll sets field value
func (o *MatchAll) SetAll(v string) {
	o.All = v
}

func (o MatchAll) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_all"] = o.All
	return toSerialize, nil
}

func (o *MatchAll) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_all",
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

	varMatchAll := _MatchAll{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatchAll)

	if err != nil {
		return err
	}

	*o = MatchAll(varMatchAll)

	return err
}

type NullableMatchAll struct {
	value *MatchAll
	isSet bool
}

func (v NullableMatchAll) Get() *MatchAll {
	return v.value
}

func (v *NullableMatchAll) Set(val *MatchAll) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchAll) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchAll(val *MatchAll) *NullableMatchAll {
	return &NullableMatchAll{value: val, isSet: true}
}

func (v NullableMatchAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


