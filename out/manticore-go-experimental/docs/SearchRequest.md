# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | 
**Query** | **map[string]map[string]interface{}** |  | 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**MaxMatches** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ScriptFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Highlight** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to **[]string** |  | [optional] 
**Profile** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest(index string, query map[string]map[string]interface{}, ) *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *SearchRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SearchRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SearchRequest) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetQuery

`func (o *SearchRequest) GetQuery() map[string]map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequest) GetQueryOk() (*map[string]map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequest) SetQuery(v map[string]map[string]interface{})`

SetQuery sets Query field to given value.


### GetLimit

`func (o *SearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *SearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMaxMatches

`func (o *SearchRequest) GetMaxMatches() int32`

GetMaxMatches returns the MaxMatches field if non-nil, zero value otherwise.

### GetMaxMatchesOk

`func (o *SearchRequest) GetMaxMatchesOk() (*int32, bool)`

GetMaxMatchesOk returns a tuple with the MaxMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMatches

`func (o *SearchRequest) SetMaxMatches(v int32)`

SetMaxMatches sets MaxMatches field to given value.

### HasMaxMatches

`func (o *SearchRequest) HasMaxMatches() bool`

HasMaxMatches returns a boolean if a field has been set.

### GetSort

`func (o *SearchRequest) GetSort() []map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchRequest) GetSortOk() (*[]map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchRequest) SetSort(v []map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetScriptFields

`func (o *SearchRequest) GetScriptFields() map[string]interface{}`

GetScriptFields returns the ScriptFields field if non-nil, zero value otherwise.

### GetScriptFieldsOk

`func (o *SearchRequest) GetScriptFieldsOk() (*map[string]interface{}, bool)`

GetScriptFieldsOk returns a tuple with the ScriptFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptFields

`func (o *SearchRequest) SetScriptFields(v map[string]interface{})`

SetScriptFields sets ScriptFields field to given value.

### HasScriptFields

`func (o *SearchRequest) HasScriptFields() bool`

HasScriptFields returns a boolean if a field has been set.

### GetHighlight

`func (o *SearchRequest) GetHighlight() map[string]interface{}`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *SearchRequest) GetHighlightOk() (*map[string]interface{}, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *SearchRequest) SetHighlight(v map[string]interface{})`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *SearchRequest) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetSource

`func (o *SearchRequest) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchRequest) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchRequest) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetProfile

`func (o *SearchRequest) GetProfile() bool`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SearchRequest) GetProfileOk() (*bool, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SearchRequest) SetProfile(v bool)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SearchRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


