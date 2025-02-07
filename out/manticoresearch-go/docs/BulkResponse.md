# BulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]map[string]interface{}** | List of results | [optional] 
**Errors** | Pointer to **bool** | Errors occurred during the bulk operation | [optional] 
**Error** | Pointer to **string** | Error message describing an error if such occurred | [optional] 
**CurrentLine** | Pointer to **int32** | Number of the row returned in the response | [optional] 
**SkippedLines** | Pointer to **int32** | Number of rows skipped in the response | [optional] 

## Methods

### NewBulkResponse

`func NewBulkResponse() *BulkResponse`

NewBulkResponse instantiates a new BulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseWithDefaults

`func NewBulkResponseWithDefaults() *BulkResponse`

NewBulkResponseWithDefaults instantiates a new BulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BulkResponse) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkResponse) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkResponse) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *BulkResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetErrors

`func (o *BulkResponse) GetErrors() bool`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkResponse) GetErrorsOk() (*bool, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkResponse) SetErrors(v bool)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetError

`func (o *BulkResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCurrentLine

`func (o *BulkResponse) GetCurrentLine() int32`

GetCurrentLine returns the CurrentLine field if non-nil, zero value otherwise.

### GetCurrentLineOk

`func (o *BulkResponse) GetCurrentLineOk() (*int32, bool)`

GetCurrentLineOk returns a tuple with the CurrentLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLine

`func (o *BulkResponse) SetCurrentLine(v int32)`

SetCurrentLine sets CurrentLine field to given value.

### HasCurrentLine

`func (o *BulkResponse) HasCurrentLine() bool`

HasCurrentLine returns a boolean if a field has been set.

### GetSkippedLines

`func (o *BulkResponse) GetSkippedLines() int32`

GetSkippedLines returns the SkippedLines field if non-nil, zero value otherwise.

### GetSkippedLinesOk

`func (o *BulkResponse) GetSkippedLinesOk() (*int32, bool)`

GetSkippedLinesOk returns a tuple with the SkippedLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedLines

`func (o *BulkResponse) SetSkippedLines(v int32)`

SetSkippedLines sets SkippedLines field to given value.

### HasSkippedLines

`func (o *BulkResponse) HasSkippedLines() bool`

HasSkippedLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


