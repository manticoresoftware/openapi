# \SearchApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Percolate**](SearchApi.md#Percolate) | **Post** /json/pq/{index}/search | Perform reverse search on a percolate index
[**Search**](SearchApi.md#Search) | **Post** /json/search | Performs a search



## Percolate

> SearchResponse Percolate(ctx, index, percolateRequest)

Perform reverse search on a percolate index

Performs a percolate search. <br/> This method must be used only on percolate indexes. <br/> Expects two paramenters: the index name and an object with array of documents to be tested. <br/> An example of the documents object: <br/> ``` {\"query\":{\"percolate\":{\"document\":{\"content\":\"sample content\"}}}} ``` <br/> Responds with an object with matched stored queries: <br/> ``` {'timed_out':false,'hits':{'total':2,'max_score':1,'hits':[{'_index':'idx_pq_1','_type':'doc','_id':'2','_score':'1','_source':{'query':{'match':{'title':'some'},}}},{'_index':'idx_pq_1','_type':'doc','_id':'5','_score':'1','_source':{'query':{'ql':'some | none'}}}]}} ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**index** | **string**| Name of the percolate index | 
**percolateRequest** | [**PercolateRequest**](PercolateRequest.md)|  | 

### Return type

[**SearchResponse**](searchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResponse Search(ctx, searchRequest)

Performs a search

Performs a search. <br/> Expects an object with mandatory properties: <br/> - the index name <br/> - the match query object <br/> Example : <br/> <code> {'index':'movies','query':{'bool':{'must':[{'query_string':' movie'}]}},'script_fields':{'myexpr':{'script':{'inline':'IF(rating>8,1,0)'}}},'sort':[{'myexpr':'desc'},{'_score':'desc'}],'profile':true} </code> <br/> It responds with an object with <br/> - time of execution <br/> - if the query timed out <br/> - an array with hits (matched documents) <br/> - additional, if profiling is enabled, an array with profiling information is attached <br/>  ``` {'took':10,'timed_out':false,'hits':{'total':2,'hits':[{'_id':'1','_score':1,'_source':{'gid':11}},{'_id':'2','_score':1,'_source':{'gid':12}}]}} ``` <br/> For more information about the match query syntax, additional paramaters that can be set to the input and response, please check: https://docs.manticoresearch.com/latest/html/http_reference/json_search.html. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchRequest** | [**SearchRequest**](SearchRequest.md)|  | 

### Return type

[**SearchResponse**](searchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

