# Manticoresearch.SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**took** | **Number** | Time taken to execute the search | [optional] 
**timedOut** | **Boolean** | Indicates whether the search operation timed out | [optional] 
**aggregations** | [**{String: AggBucketsResult}**](AggBucketsResult.md) | Aggregated search results grouped by the specified criteria. Each named aggregation typically contains a &#x60;buckets&#x60; array (or keyed map) of bucket objects with &#x60;key&#x60;, &#x60;doc_count&#x60;, and optional &#x60;status&#x60;.  | [optional] 
**hits** | [**SearchResponseHits**](SearchResponseHits.md) |  | [optional] 
**profile** | **Object** | Profile information about the search execution, if profiling is enabled | [optional] 
**scroll** | **String** | Scroll token to be used fo pagination | [optional] 
**warning** | **Object** | Warnings encountered during the search operation | [optional] 


