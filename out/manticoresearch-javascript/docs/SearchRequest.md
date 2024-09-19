# Manticoresearch.SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**knn** | [**KnnSearchRequestAllOfKnn**](KnnSearchRequestAllOfKnn.md) |  | 
**aggs** | [**Aggregation**](Aggregation.md) |  | [optional] 
**expressions** | **{String: String}** |  | [optional] 
**join** | [**[JoinInner]**](JoinInner.md) |  | [optional] 
**highlight** | [**Highlight**](Highlight.md) |  | [optional] 
**index** | **String** |  | 
**limit** | **Number** |  | [optional] 
**maxMatches** | **Number** |  | [optional] 
**offset** | **Number** |  | [optional] 
**options** | **Object** |  | [optional] 
**profile** | **Boolean** |  | [optional] 
**sort** | [**[SearchRequestParametersSortInner]**](SearchRequestParametersSortInner.md) |  | [optional] 
**source** | [**SearchRequestParametersSource**](SearchRequestParametersSource.md) |  | [optional] 
**trackScores** | **Boolean** |  | [optional] 
**query** | [**QueryFilter**](QueryFilter.md) |  | 


