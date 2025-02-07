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

// checks if the BulkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkResponse{}

// BulkResponse Success response for bulk search requests
type BulkResponse struct {
	// List of results
	Items []map[string]interface{}
	// Errors occurred during the bulk operation
	Errors *bool
	// Error message describing an error if such occurred
	Error *string
	// Number of the row returned in the response
	CurrentLine *int32
	// Number of rows skipped in the response
	SkippedLines *int32
}

// NewBulkResponse instantiates a new BulkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkResponse() *BulkResponse {
	this := BulkResponse{}
	return &this
}

// NewBulkResponseWithDefaults instantiates a new BulkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkResponseWithDefaults() *BulkResponse {
	this := BulkResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BulkResponse) GetItems() []map[string]interface{} {
	if o == nil || IsNil(o.Items) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResponse) GetItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BulkResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []map[string]interface{} and assigns it to the Items field.
func (o *BulkResponse) SetItems(v []map[string]interface{}) {
	o.Items = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BulkResponse) GetErrors() bool {
	if o == nil || IsNil(o.Errors) {
		var ret bool
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResponse) GetErrorsOk() (*bool, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BulkResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given bool and assigns it to the Errors field.
func (o *BulkResponse) SetErrors(v bool) {
	o.Errors = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BulkResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BulkResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BulkResponse) SetError(v string) {
	o.Error = &v
}

// GetCurrentLine returns the CurrentLine field value if set, zero value otherwise.
func (o *BulkResponse) GetCurrentLine() int32 {
	if o == nil || IsNil(o.CurrentLine) {
		var ret int32
		return ret
	}
	return *o.CurrentLine
}

// GetCurrentLineOk returns a tuple with the CurrentLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResponse) GetCurrentLineOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentLine) {
		return nil, false
	}
	return o.CurrentLine, true
}

// HasCurrentLine returns a boolean if a field has been set.
func (o *BulkResponse) HasCurrentLine() bool {
	if o != nil && !IsNil(o.CurrentLine) {
		return true
	}

	return false
}

// SetCurrentLine gets a reference to the given int32 and assigns it to the CurrentLine field.
func (o *BulkResponse) SetCurrentLine(v int32) {
	o.CurrentLine = &v
}

// GetSkippedLines returns the SkippedLines field value if set, zero value otherwise.
func (o *BulkResponse) GetSkippedLines() int32 {
	if o == nil || IsNil(o.SkippedLines) {
		var ret int32
		return ret
	}
	return *o.SkippedLines
}

// GetSkippedLinesOk returns a tuple with the SkippedLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResponse) GetSkippedLinesOk() (*int32, bool) {
	if o == nil || IsNil(o.SkippedLines) {
		return nil, false
	}
	return o.SkippedLines, true
}

// HasSkippedLines returns a boolean if a field has been set.
func (o *BulkResponse) HasSkippedLines() bool {
	if o != nil && !IsNil(o.SkippedLines) {
		return true
	}

	return false
}

// SetSkippedLines gets a reference to the given int32 and assigns it to the SkippedLines field.
func (o *BulkResponse) SetSkippedLines(v int32) {
	o.SkippedLines = &v
}

func (o BulkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.CurrentLine) {
		toSerialize["current_line"] = o.CurrentLine
	}
	if !IsNil(o.SkippedLines) {
		toSerialize["skipped_lines"] = o.SkippedLines
	}
	return toSerialize, nil
}

type NullableBulkResponse struct {
	value *BulkResponse
	isSet bool
}

func (v NullableBulkResponse) Get() *BulkResponse {
	return v.value
}

func (v *NullableBulkResponse) Set(val *BulkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkResponse(val *BulkResponse) *NullableBulkResponse {
	return &NullableBulkResponse{value: val, isSet: true}
}

func (v NullableBulkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


