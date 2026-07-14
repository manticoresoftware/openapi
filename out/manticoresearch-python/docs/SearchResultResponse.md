# SearchResultResponse

Response from `POST /search`. Returns `searchResponse` for regular search requests and `chatResponse` when the request includes a `chat` object. 

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**took** | **int** | Time taken to execute the search | [optional] 
**timed_out** | **bool** | Indicates whether the search operation timed out | [optional] 
**aggregations** | [**Dict[str, SearchResponseAggregationsValue]**](SearchResponseAggregationsValue.md) | Aggregated search results grouped by the specified criteria. Each named aggregation typically contains a &#x60;buckets&#x60; array (or keyed map) of bucket objects with &#x60;key&#x60;, &#x60;doc_count&#x60;, and optional &#x60;status&#x60;.  | [optional] 
**hits** | [**SearchResponseHits**](SearchResponseHits.md) |  | [optional] 
**profile** | **object** | Profile information about the search execution, if profiling is enabled | [optional] 
**scroll** | **str** | Scroll token to be used fo pagination | [optional] 
**warning** | **object** | Warnings encountered during the search operation | [optional] 
**conversation_uuid** | **str** | Existing or generated conversation id | 
**user_query** | **str** | Original user query | 
**search_query** | **str** | Standalone search query used for KNN retrieval | 
**response** | **str** | LLM answer as generated | 
**sources** | **str** | JSON string containing retrieved source rows used as LLM context.  | 

## Example

```python
from manticoresearch.models.search_result_response import SearchResultResponse

# TODO update the JSON string below
json = "{}"
# create an instance of SearchResultResponse from a JSON string
search_result_response_instance = SearchResultResponse.from_json(json)
# print the JSON string representation of the object
print(SearchResultResponse.to_json())

# convert the object into a dict
search_result_response_dict = search_result_response_instance.to_dict()
# create an instance of SearchResultResponse from a dict
search_result_response_from_dict = SearchResultResponse.from_dict(search_result_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


