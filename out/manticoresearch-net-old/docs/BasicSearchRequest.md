# ManticoreSearch.Model.BasicSearchRequest
Request object for search operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggs** | [**Aggregation**](Aggregation.md) |  | [optional] 
**Expressions** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**Join** | [**List&lt;JoinInner&gt;**](JoinInner.md) |  | [optional] 
**Highlight** | [**Highlight**](Highlight.md) |  | [optional] 
**Index** | **string** |  | 
**Limit** | **int** |  | [optional] 
**MaxMatches** | **int** |  | [optional] 
**Offset** | **int** |  | [optional] 
**Options** | **Object** |  | [optional] 
**Profile** | **bool** |  | [optional] 
**Sort** | [**List&lt;SearchRequestParametersSortInner&gt;**](SearchRequestParametersSortInner.md) |  | [optional] 
**Source** | [**SearchRequestParametersSource**](SearchRequestParametersSource.md) |  | [optional] 
**TrackScores** | **bool** |  | [optional] 
**Query** | [**QueryFilter**](QueryFilter.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

