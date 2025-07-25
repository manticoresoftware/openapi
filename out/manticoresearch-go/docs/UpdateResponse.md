# UpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | Pointer to **string** | Name of the document table | [optional] 
**Updated** | Pointer to **int32** | Number of documents updated | [optional] 
**Id** | Pointer to **uint64** | Document ID | [optional]
**Result** | Pointer to **string** | Result of the update operation, typically &#39;updated&#39; | [optional] 

## Methods

### NewUpdateResponse

`func NewUpdateResponse() *UpdateResponse`

NewUpdateResponse instantiates a new UpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResponseWithDefaults

`func NewUpdateResponseWithDefaults() *UpdateResponse`

NewUpdateResponseWithDefaults instantiates a new UpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *UpdateResponse) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *UpdateResponse) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *UpdateResponse) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *UpdateResponse) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetUpdated

`func (o *UpdateResponse) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpdateResponse) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpdateResponse) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UpdateResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *UpdateResponse) GetId() uint64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateResponse) GetIdOk() (*uint64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateResponse) SetId(v uint64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResult

`func (o *UpdateResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


