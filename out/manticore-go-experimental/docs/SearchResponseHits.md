# SearchResponseHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Hits** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewSearchResponseHits

`func NewSearchResponseHits() *SearchResponseHits`

NewSearchResponseHits instantiates a new SearchResponseHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseHitsWithDefaults

`func NewSearchResponseHitsWithDefaults() *SearchResponseHits`

NewSearchResponseHitsWithDefaults instantiates a new SearchResponseHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SearchResponseHits) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchResponseHits) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchResponseHits) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchResponseHits) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHits

`func (o *SearchResponseHits) GetHits() []map[string]interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchResponseHits) GetHitsOk() (*[]map[string]interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchResponseHits) SetHits(v []map[string]interface{})`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchResponseHits) HasHits() bool`

HasHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


