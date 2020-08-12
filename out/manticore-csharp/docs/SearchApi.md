# Org.OpenAPITools.Api.SearchApi

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Percolate**](SearchApi.md#percolate) | **POST** /json/pq/{index}/search | Perform reverse search on a percolate index
[**Search**](SearchApi.md#search) | **POST** /json/search | Performs a search



## Percolate

> SearchResponse Percolate (string index, PercolateRequest percolateRequest)

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

```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Org.OpenAPITools.Api;
using Org.OpenAPITools.Client;
using Org.OpenAPITools.Model;

namespace Example
{
    public class PercolateExample
    {
        public static void Main()
        {
            Configuration.Default.BasePath = "http://127.0.0.1:9308";
            var apiInstance = new SearchApi(Configuration.Default);
            var index = index_example;  // string | Name of the percolate index
            var percolateRequest = new PercolateRequest(); // PercolateRequest | 

            try
            {
                // Perform reverse search on a percolate index
                SearchResponse result = apiInstance.Percolate(index, percolateRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException e)
            {
                Debug.Print("Exception when calling SearchApi.Percolate: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **string**| Name of the percolate index | 
 **percolateRequest** | [**PercolateRequest**](PercolateRequest.md)|  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | items found |  -  |
| **0** | error |  -  |

[[Back to top]](#)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResponse Search (SearchRequest searchRequest)

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

```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Org.OpenAPITools.Api;
using Org.OpenAPITools.Client;
using Org.OpenAPITools.Model;

namespace Example
{
    public class SearchExample
    {
        public static void Main()
        {
            Configuration.Default.BasePath = "http://127.0.0.1:9308";
            var apiInstance = new SearchApi(Configuration.Default);
            var searchRequest = new SearchRequest(); // SearchRequest | 

            try
            {
                // Performs a search
                SearchResponse result = apiInstance.Search(searchRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException e)
            {
                Debug.Print("Exception when calling SearchApi.Search: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md)|  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Ok |  -  |
| **0** | error |  -  |

[[Back to top]](#)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

