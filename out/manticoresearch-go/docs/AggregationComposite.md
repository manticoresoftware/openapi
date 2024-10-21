# AggregationComposite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | Maximum number of composite buckets in the result | [optional] 
**Sources** | Pointer to [**[]map[string]AggregationCompositeSourcesInnerValue**](map[string]AggregationCompositeSourcesInnerValue.md) |  | [optional] 

## Methods

### NewAggregationComposite

`func NewAggregationComposite() *AggregationComposite`

NewAggregationComposite instantiates a new AggregationComposite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationCompositeWithDefaults

`func NewAggregationCompositeWithDefaults() *AggregationComposite`

NewAggregationCompositeWithDefaults instantiates a new AggregationComposite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *AggregationComposite) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AggregationComposite) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AggregationComposite) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AggregationComposite) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSources

`func (o *AggregationComposite) GetSources() []map[string]AggregationCompositeSourcesInnerValue`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AggregationComposite) GetSourcesOk() (*[]map[string]AggregationCompositeSourcesInnerValue, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AggregationComposite) SetSources(v []map[string]AggregationCompositeSourcesInnerValue)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AggregationComposite) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


