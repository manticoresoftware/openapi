# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**took** | Option<**i32**> | Time taken to execute the search | [optional]
**timed_out** | Option<**bool**> | Indicates whether the search operation timed out | [optional]
**aggregations** | Option<[**serde_json::Value**](.md)> | Aggregated search results grouped by the specified criteria | [optional]
**hits** | Option<[**models::SearchResponseHits**](searchResponse_hits.md)> |  | [optional]
**profile** | Option<[**serde_json::Value**](.md)> | Profile information about the search execution, if profiling is enabled | [optional]
**scroll** | Option<**String**> | Scroll token to be used fo pagination | [optional]
**warning** | Option<[**serde_json::Value**](.md)> | Warnings encountered during the search operation | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


