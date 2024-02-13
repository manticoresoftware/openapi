# ManticoreSearch.Model.KnnSearchRequestByVector
Request object for knn search operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | [default to ""]
**Field** | **string** |  | [default to ""]
**QueryVector** | **List&lt;decimal&gt;** |  | 
**K** | **int** |  | 
**FulltextFilter** | **Object** |  | [optional] 
**AttrFilter** | **Object** |  | [optional] 
**Limit** | **int** |  | [optional] 
**Offset** | **int** |  | [optional] 
**MaxMatches** | **int** |  | [optional] 
**Sort** | **List&lt;Object&gt;** |  | [optional] 
**Aggs** | [**Dictionary&lt;string, Aggregation&gt;**](Aggregation.md) |  | [optional] 
**Expressions** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**Highlight** | [**Highlight**](Highlight.md) |  | [optional] 
**Source** | **Object** |  | [optional] 
**Options** | **Dictionary&lt;string, Object&gt;** |  | [optional] 
**Profile** | **bool** |  | [optional] 
**TrackScores** | **bool** |  | [optional] 



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

