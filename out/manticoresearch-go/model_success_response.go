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

// checks if the SuccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessResponse{}

// SuccessResponse Response object indicating the success of an operation, such as inserting or updating a document
type SuccessResponse struct {
	// Name of the document index
	Index *string
	// Name of the document table (alias of index)
	Table *string
	// ID of the document affected by the request operation
	Id *int64
	// Indicates whether the document was created as a result of the operation
	Created *bool
	// Result of the operation, typically 'created', 'updated', or 'deleted'
	Result *string
	// Indicates whether the document was found in the index
	Found *bool
	// HTTP status code representing the result of the operation
	Status *int32
}

// NewSuccessResponse instantiates a new SuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResponse() *SuccessResponse {
	this := SuccessResponse{}
	return &this
}

// NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResponseWithDefaults() *SuccessResponse {
	this := SuccessResponse{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SuccessResponse) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SuccessResponse) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *SuccessResponse) SetIndex(v string) {
	o.Index = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *SuccessResponse) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *SuccessResponse) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *SuccessResponse) SetTable(v string) {
	o.Table = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SuccessResponse) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SuccessResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SuccessResponse) SetId(v int64) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SuccessResponse) GetCreated() bool {
	if o == nil || IsNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SuccessResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *SuccessResponse) SetCreated(v bool) {
	o.Created = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SuccessResponse) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SuccessResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *SuccessResponse) SetResult(v string) {
	o.Result = &v
}

// GetFound returns the Found field value if set, zero value otherwise.
func (o *SuccessResponse) GetFound() bool {
	if o == nil || IsNil(o.Found) {
		var ret bool
		return ret
	}
	return *o.Found
}

// GetFoundOk returns a tuple with the Found field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetFoundOk() (*bool, bool) {
	if o == nil || IsNil(o.Found) {
		return nil, false
	}
	return o.Found, true
}

// HasFound returns a boolean if a field has been set.
func (o *SuccessResponse) HasFound() bool {
	if o != nil && !IsNil(o.Found) {
		return true
	}

	return false
}

// SetFound gets a reference to the given bool and assigns it to the Found field.
func (o *SuccessResponse) SetFound(v bool) {
	o.Found = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SuccessResponse) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SuccessResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *SuccessResponse) SetStatus(v int32) {
	o.Status = &v
}

func (o SuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["_index"] = o.Index
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Found) {
		toSerialize["found"] = o.Found
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSuccessResponse struct {
	value *SuccessResponse
	isSet bool
}

func (v NullableSuccessResponse) Get() *SuccessResponse {
	return v.value
}

func (v *NullableSuccessResponse) Set(val *SuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResponse(val *SuccessResponse) *NullableSuccessResponse {
	return &NullableSuccessResponse{value: val, isSet: true}
}

func (v NullableSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


