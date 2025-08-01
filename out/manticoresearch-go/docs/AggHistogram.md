# AggHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field to group by | 
**Interval** | **int32** | Interval of the histogram values | 
**Offset** | Pointer to **int32** | Offset of the histogram values. Default value is 0. | [optional] 
**Keyed** | Pointer to **bool** | Flag that defines if a search response will be a dictionary with the bucket keys. Default value is false. | [optional] 

## Methods

### NewAggHistogram

`func NewAggHistogram(field string, interval int32, ) *AggHistogram`

NewAggHistogram instantiates a new AggHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggHistogramWithDefaults

`func NewAggHistogramWithDefaults() *AggHistogram`

NewAggHistogramWithDefaults instantiates a new AggHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggHistogram) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggHistogram) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggHistogram) SetField(v string)`

SetField sets Field field to given value.


### GetInterval

`func (o *AggHistogram) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AggHistogram) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AggHistogram) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetOffset

`func (o *AggHistogram) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AggHistogram) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AggHistogram) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AggHistogram) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetKeyed

`func (o *AggHistogram) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *AggHistogram) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *AggHistogram) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *AggHistogram) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


