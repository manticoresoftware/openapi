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

// checks if the JoinInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JoinInner{}

// JoinInner struct for JoinInner
type JoinInner struct {
	On []JoinInnerOnInner `json:"on"`
	Query *FulltextFilter `json:"query,omitempty"`
	Table string `json:"table"`
	Type string `json:"type"`
}

type _JoinInner JoinInner

// NewJoinInner instantiates a new JoinInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoinInner(on []JoinInnerOnInner, table string, type_ string) *JoinInner {
	this := JoinInner{}
	this.On = on
	this.Table = table
	this.Type = type_
	return &this
}

// NewJoinInnerWithDefaults instantiates a new JoinInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoinInnerWithDefaults() *JoinInner {
	this := JoinInner{}
	return &this
}

// GetOn returns the On field value
func (o *JoinInner) GetOn() []JoinInnerOnInner {
	if o == nil {
		var ret []JoinInnerOnInner
		return ret
	}

	return o.On
}

// GetOnOk returns a tuple with the On field value
// and a boolean to check if the value has been set.
func (o *JoinInner) GetOnOk() ([]JoinInnerOnInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.On, true
}

// SetOn sets field value
func (o *JoinInner) SetOn(v []JoinInnerOnInner) {
	o.On = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *JoinInner) GetQuery() FulltextFilter {
	if o == nil || IsNil(o.Query) {
		var ret FulltextFilter
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinInner) GetQueryOk() (*FulltextFilter, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *JoinInner) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given FulltextFilter and assigns it to the Query field.
func (o *JoinInner) SetQuery(v FulltextFilter) {
	o.Query = &v
}

// GetTable returns the Table field value
func (o *JoinInner) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *JoinInner) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value
func (o *JoinInner) SetTable(v string) {
	o.Table = v
}

// GetType returns the Type field value
func (o *JoinInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JoinInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JoinInner) SetType(v string) {
	o.Type = v
}

func (o JoinInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JoinInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["on"] = o.On
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	toSerialize["table"] = o.Table
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *JoinInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"on",
		"table",
		"type",
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

	varJoinInner := _JoinInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJoinInner)

	if err != nil {
		return err
	}

	*o = JoinInner(varJoinInner)

	return err
}

type NullableJoinInner struct {
	value *JoinInner
	isSet bool
}

func (v NullableJoinInner) Get() *JoinInner {
	return v.value
}

func (v *NullableJoinInner) Set(val *JoinInner) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinInner) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinInner(val *JoinInner) *NullableJoinInner {
	return &NullableJoinInner{value: val, isSet: true}
}

func (v NullableJoinInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


