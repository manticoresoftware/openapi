# SearchAPI

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**percolate**](SearchAPI.md#percolate) | **POST** /json/pq/{index}/search | Perform reverse search on a percolate index
[**search**](SearchAPI.md#search) | **POST** /json/search | Performs a search


# **percolate**
```swift
    open class func percolate(index: String, percolateRequest: PercolateRequest, completion: @escaping (_ data: SearchResponse?, _ error: Error?) -> Void)
```

Perform reverse search on a percolate index

Performs a percolate search. 
This method must be used only on percolate indexes.

Expects two paramenters: the index name and an object with array of documents to be tested.
An example of the documents object:

  ```
  {"query":{"percolate":{"document":{"content":"sample content"}}}}
  ```

Responds with an object with matched stored queries: 

  ```
  {'timed_out':false,'hits':{'total':2,'max_score':1,'hits':[{'_index':'idx_pq_1','_type':'doc','_id':'2','_score':'1','_source':{'query':{'match':{'title':'some'},}}},{'_index':'idx_pq_1','_type':'doc','_id':'5','_score':'1','_source':{'query':{'ql':'some | none'}}}]}}
  ```


### Example 
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let index = "index_example" // String | Name of the percolate index
let percolateRequest = percolateRequest(query: "TODO") // PercolateRequest | 

// Perform reverse search on a percolate index
SearchAPI.percolate(index: index, percolateRequest: percolateRequest) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **String** | Name of the percolate index | 
 **percolateRequest** | [**PercolateRequest**](PercolateRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **search**
```swift
    open class func search(searchRequest: SearchRequest, completion: @escaping (_ data: SearchResponse?, _ error: Error?) -> Void)
```

Performs a search

Performs a search. 
Expects an object with mandatory properties:
* the index name
* the match query object
Example :

  ```
  {'index':'movies','query':{'bool':{'must':[{'query_string':' movie'}]}},'script_fields':{'myexpr':{'script':{'inline':'IF(rating>8,1,0)'}}},'sort':[{'myexpr':'desc'},{'_score':'desc'}],'profile':true}
  ```

It responds with an object with:
- time of execution
- if the query timed out
- an array with hits (matched documents)
- additional, if profiling is enabled, an array with profiling information is attached


  ```
  {'took':10,'timed_out':false,'hits':{'total':2,'hits':[{'_id':'1','_score':1,'_source':{'gid':11}},{'_id':'2','_score':1,'_source':{'gid':12}}]}}
  ```

For more information about the match query syntax, additional paramaters that can be set to the input and response, please check: https://docs.manticoresearch.com/latest/html/http_reference/json_search.html.


### Example 
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let searchRequest = searchRequest(index: "index_example", query: "TODO", limit: 123, offset: 123, maxMatches: 123, sort: [123], scriptFields: 123, highlight: 123, source: ["source_example"], profile: false) // SearchRequest | 

// Performs a search
SearchAPI.search(searchRequest: searchRequest) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

