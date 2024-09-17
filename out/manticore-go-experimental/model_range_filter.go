/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 3.3.1
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RangeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeFilter{}

// RangeFilter Range attribute filter
type RangeFilter struct {
	Field string `json:"field"`
	Lte NullableRangeFilterLte `json:"lte,omitempty"`
	Gte NullableRangeFilterLte `json:"gte,omitempty"`
	Lt NullableRangeFilterLte `json:"lt,omitempty"`
	Gt NullableRangeFilterLte `json:"gt,omitempty"`
}

type _RangeFilter RangeFilter

// NewRangeFilter instantiates a new RangeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeFilter(field string) *RangeFilter {
	this := RangeFilter{}
	this.Field = field
	return &this
}

// NewRangeFilterWithDefaults instantiates a new RangeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeFilterWithDefaults() *RangeFilter {
	this := RangeFilter{}
	return &this
}

// GetField returns the Field field value
func (o *RangeFilter) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *RangeFilter) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *RangeFilter) SetField(v string) {
	o.Field = v
}

// GetLte returns the Lte field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RangeFilter) GetLte() RangeFilterLte {
	if o == nil || IsNil(o.Lte.Get()) {
		var ret RangeFilterLte
		return ret
	}
	return *o.Lte.Get()
}

// GetLteOk returns a tuple with the Lte field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RangeFilter) GetLteOk() (*RangeFilterLte, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lte.Get(), o.Lte.IsSet()
}

// HasLte returns a boolean if a field has been set.
func (o *RangeFilter) HasLte() bool {
	if o != nil && o.Lte.IsSet() {
		return true
	}

	return false
}

// SetLte gets a reference to the given NullableRangeFilterLte and assigns it to the Lte field.
func (o *RangeFilter) SetLte(v RangeFilterLte) {
	o.Lte.Set(&v)
}
// SetLteNil sets the value for Lte to be an explicit nil
func (o *RangeFilter) SetLteNil() {
	o.Lte.Set(nil)
}

// UnsetLte ensures that no value is present for Lte, not even an explicit nil
func (o *RangeFilter) UnsetLte() {
	o.Lte.Unset()
}

// GetGte returns the Gte field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RangeFilter) GetGte() RangeFilterLte {
	if o == nil || IsNil(o.Gte.Get()) {
		var ret RangeFilterLte
		return ret
	}
	return *o.Gte.Get()
}

// GetGteOk returns a tuple with the Gte field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RangeFilter) GetGteOk() (*RangeFilterLte, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gte.Get(), o.Gte.IsSet()
}

// HasGte returns a boolean if a field has been set.
func (o *RangeFilter) HasGte() bool {
	if o != nil && o.Gte.IsSet() {
		return true
	}

	return false
}

// SetGte gets a reference to the given NullableRangeFilterLte and assigns it to the Gte field.
func (o *RangeFilter) SetGte(v RangeFilterLte) {
	o.Gte.Set(&v)
}
// SetGteNil sets the value for Gte to be an explicit nil
func (o *RangeFilter) SetGteNil() {
	o.Gte.Set(nil)
}

// UnsetGte ensures that no value is present for Gte, not even an explicit nil
func (o *RangeFilter) UnsetGte() {
	o.Gte.Unset()
}

// GetLt returns the Lt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RangeFilter) GetLt() RangeFilterLte {
	if o == nil || IsNil(o.Lt.Get()) {
		var ret RangeFilterLte
		return ret
	}
	return *o.Lt.Get()
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RangeFilter) GetLtOk() (*RangeFilterLte, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lt.Get(), o.Lt.IsSet()
}

// HasLt returns a boolean if a field has been set.
func (o *RangeFilter) HasLt() bool {
	if o != nil && o.Lt.IsSet() {
		return true
	}

	return false
}

// SetLt gets a reference to the given NullableRangeFilterLte and assigns it to the Lt field.
func (o *RangeFilter) SetLt(v RangeFilterLte) {
	o.Lt.Set(&v)
}
// SetLtNil sets the value for Lt to be an explicit nil
func (o *RangeFilter) SetLtNil() {
	o.Lt.Set(nil)
}

// UnsetLt ensures that no value is present for Lt, not even an explicit nil
func (o *RangeFilter) UnsetLt() {
	o.Lt.Unset()
}

// GetGt returns the Gt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RangeFilter) GetGt() RangeFilterLte {
	if o == nil || IsNil(o.Gt.Get()) {
		var ret RangeFilterLte
		return ret
	}
	return *o.Gt.Get()
}

// GetGtOk returns a tuple with the Gt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RangeFilter) GetGtOk() (*RangeFilterLte, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gt.Get(), o.Gt.IsSet()
}

// HasGt returns a boolean if a field has been set.
func (o *RangeFilter) HasGt() bool {
	if o != nil && o.Gt.IsSet() {
		return true
	}

	return false
}

// SetGt gets a reference to the given NullableRangeFilterLte and assigns it to the Gt field.
func (o *RangeFilter) SetGt(v RangeFilterLte) {
	o.Gt.Set(&v)
}
// SetGtNil sets the value for Gt to be an explicit nil
func (o *RangeFilter) SetGtNil() {
	o.Gt.Set(nil)
}

// UnsetGt ensures that no value is present for Gt, not even an explicit nil
func (o *RangeFilter) UnsetGt() {
	o.Gt.Unset()
}

func (o RangeFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	if o.Lte.IsSet() {
		toSerialize["lte"] = o.Lte.Get()
	}
	if o.Gte.IsSet() {
		toSerialize["gte"] = o.Gte.Get()
	}
	if o.Lt.IsSet() {
		toSerialize["lt"] = o.Lt.Get()
	}
	if o.Gt.IsSet() {
		toSerialize["gt"] = o.Gt.Get()
	}
	return toSerialize, nil
}

func (o *RangeFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
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

	varRangeFilter := _RangeFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRangeFilter)

	if err != nil {
		return err
	}

	*o = RangeFilter(varRangeFilter)

	return err
}

type NullableRangeFilter struct {
	value *RangeFilter
	isSet bool
}

func (v NullableRangeFilter) Get() *RangeFilter {
	return v.value
}

func (v *NullableRangeFilter) Set(val *RangeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeFilter(val *RangeFilter) *NullableRangeFilter {
	return &NullableRangeFilter{value: val, isSet: true}
}

func (v NullableRangeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


