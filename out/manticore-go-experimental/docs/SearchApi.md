# ::SearchApi

## Load the API package
```perl
use ::Object::SearchApi;
```

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Percolate**](SearchApi.md#Percolate) | **Post** /json/pq/{index}/search | Perform reverse search on a percolate index
[**Search**](SearchApi.md#Search) | **Post** /json/search | Performs a search


# **Percolate**
> SearchResponse Percolate(index => $index, percolateRequest => $percolateRequest)

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
```perl
use Data::Dumper;
use ::SearchApi;
my $api_instance = ::SearchApi->new(
);

my $index = index_example; # string | Name of the percolate index
my $percolateRequest = ::Object::PercolateRequest->new(); # PercolateRequest | 

eval { 
    my $result = $api_instance->Percolate(index => $index, percolateRequest => $percolateRequest);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling SearchApi->Percolate: $@\n";
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **string**| Name of the percolate index | 
 **percolateRequest** | [**PercolateRequest**](PercolateRequest.md)|  | 

### Return type

[**SearchResponse**](searchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search**
> SearchResponse Search(searchRequest => $searchRequest)

Performs a search


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
```perl
use Data::Dumper;
use ::SearchApi;
my $api_instance = ::SearchApi->new(
);

my $searchRequest = ::Object::SearchRequest->new(); # SearchRequest | 

eval { 
    my $result = $api_instance->Search(searchRequest => $searchRequest);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling SearchApi->Search: $@\n";
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md)|  | 

### Return type

[**SearchResponse**](searchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

