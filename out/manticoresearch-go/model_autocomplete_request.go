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

// checks if the AutocompleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutocompleteRequest{}

// AutocompleteRequest Object containing the data for performing an autocomplete search.
type AutocompleteRequest struct {
	// The table to perform the search on
	Table string `json:"table"`
	// The beginning of the string to autocomplete
	Query string `json:"query"`
	// Autocomplete options   - layouts: A comma-separated string of keyboard layout codes to validate and check for spell correction. Available options - us, ru, ua, se, pt, no, it, gr, uk, fr, es, dk, de, ch, br, bg, be. By default, all are enabled.   - fuzziness: (0,1 or 2) Maximum Levenshtein distance for finding typos. Set to 0 to disable fuzzy matching. Default is 2   - prepend: true/false If true, adds an asterisk before the last word for prefix expansion (e.g., *word )   - append:  true/false If true, adds an asterisk after the last word for prefix expansion (e.g., word* )   - expansion_len: Number of characters to expand in the last word. Default is 10. 
	Options map[string]interface{} `json:"options,omitempty"`
}

type _AutocompleteRequest AutocompleteRequest

// NewAutocompleteRequest instantiates a new AutocompleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteRequest(table string, query string) *AutocompleteRequest {
	this := AutocompleteRequest{}
	this.Table = table
	this.Query = query
	return &this
}

// NewAutocompleteRequestWithDefaults instantiates a new AutocompleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteRequestWithDefaults() *AutocompleteRequest {
	this := AutocompleteRequest{}
	return &this
}

// GetTable returns the Table field value
func (o *AutocompleteRequest) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *AutocompleteRequest) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value
func (o *AutocompleteRequest) SetTable(v string) {
	o.Table = v
}

// GetQuery returns the Query field value
func (o *AutocompleteRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *AutocompleteRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *AutocompleteRequest) SetQuery(v string) {
	o.Query = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AutocompleteRequest) GetOptions() map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteRequest) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AutocompleteRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *AutocompleteRequest) SetOptions(v map[string]interface{}) {
	o.Options = v
}

func (o AutocompleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutocompleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["table"] = o.Table
	toSerialize["query"] = o.Query
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *AutocompleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"table",
		"query",
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

	varAutocompleteRequest := _AutocompleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutocompleteRequest)

	if err != nil {
		return err
	}

	*o = AutocompleteRequest(varAutocompleteRequest)

	return err
}

type NullableAutocompleteRequest struct {
	value *AutocompleteRequest
	isSet bool
}

func (v NullableAutocompleteRequest) Get() *AutocompleteRequest {
	return v.value
}

func (v *NullableAutocompleteRequest) Set(val *AutocompleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteRequest(val *AutocompleteRequest) *NullableAutocompleteRequest {
	return &NullableAutocompleteRequest{value: val, isSet: true}
}

func (v NullableAutocompleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


