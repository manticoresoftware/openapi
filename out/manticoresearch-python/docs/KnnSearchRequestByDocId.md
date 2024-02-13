# KnnSearchRequestByDocId

Request object for knn search operation
## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **str** |  | [default to ""]
**field** | **str** |  | [default to ""]
**doc_id** | **int** |  | 
**k** | **int** |  | 
**fulltext_filter** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**attr_filter** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**limit** | **int** |  | [optional] 
**offset** | **int** |  | [optional] 
**max_matches** | **int** |  | [optional] 
**sort** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]** |  | [optional] 
**aggs** | [**{str: (Aggregation,)}**](Aggregation.md) |  | [optional] 
**expressions** | **{str: (str,)}** |  | [optional] 
**highlight** | [**Highlight**](Highlight.md) |  | [optional] 
**source** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**profile** | **bool** |  | [optional] 
**track_scores** | **bool** |  | [optional] 

## Building a knn search request

[[Docs on knn search options in Manticore Search Manual]](https://manual.manticoresearch.com/Searching/KNN)
```python
    #Creating a request for knn search by document id:
    
    #Setting a search index:
    search_req = manticoresearch.model.KnnSearchRequestByVector()
    search_req.index = 'test'
    
    api_response = api_instance.search(search_req)
    pprint(api_response)
    
    #Setting knn settings:
    search.req.field = "test_field"
    search.req.doc_id = 10
    search.req.k = 5
    
    #Setting extra search settings (optional):
    
    search_req.offset = 0
    search_req.limit = 10
    search_req.profile = True
    search_req.options = {'ranker': 'bm25'}
    search_req.options['retry_count'] = 2
    
    api_response = api_instance.search(search_req)
    pprint(api_response)
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


