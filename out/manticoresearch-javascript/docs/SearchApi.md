# Manticoresearch.SearchApi

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**percolate**](SearchApi.md#percolate) | **POST** /pq/{index}/search | Perform reverse search on a percolate index
[**search**](SearchApi.md#search) | **POST** /search | Performs a search on an index



## percolate

> SearchResponse percolate(index, percolateRequest)

Perform reverse search on a percolate index

Performs a percolate search.  This method must be used only on percolate indexes.  Expects two parameters: the index name and an object with array of documents to be tested. An example of the documents object:    &#x60;&#x60;&#x60;   {     \&quot;query\&quot;:     {       \&quot;percolate\&quot;:       {         \&quot;document\&quot;:         {           \&quot;content\&quot;:\&quot;sample content\&quot;         }       }     }   }   &#x60;&#x60;&#x60;  Responds with an object with matched stored queries:     &#x60;&#x60;&#x60;   {     &#39;timed_out&#39;:false,     &#39;hits&#39;:     {       &#39;total&#39;:2,       &#39;max_score&#39;:1,       &#39;hits&#39;:       [         {           &#39;_index&#39;:&#39;idx_pq_1&#39;,           &#39;_type&#39;:&#39;doc&#39;,           &#39;_id&#39;:&#39;2&#39;,           &#39;_score&#39;:&#39;1&#39;,           &#39;_source&#39;:           {             &#39;query&#39;:             {               &#39;match&#39;:{&#39;title&#39;:&#39;some&#39;}             }           }         },         {           &#39;_index&#39;:&#39;idx_pq_1&#39;,           &#39;_type&#39;:&#39;doc&#39;,           &#39;_id&#39;:&#39;5&#39;,           &#39;_score&#39;:&#39;1&#39;,           &#39;_source&#39;:           {             &#39;query&#39;:             {               &#39;ql&#39;:&#39;some | none&#39;             }           }         }       ]     }   }   &#x60;&#x60;&#x60; 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.SearchApi();
let index = "index_example"; // String | Name of the percolate index
let percolateRequest = {"query":{"percolate":{"document":{"title":"some text to match"}}}}; // PercolateRequest | 
apiInstance.percolate(index, percolateRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **String**| Name of the percolate index | 
 **percolateRequest** | [**PercolateRequest**](PercolateRequest.md)|  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## search

> SearchResponse search(searchRequest)

Performs a search on an index

 The method expects an object with the following mandatory properties: * the name of the index to search * the match query object For details, see the documentation on [**SearchRequest**](SearchRequest.md) The method returns an object with the following properties: - took: the time taken to execute the search query. - timed_out: a boolean indicating whether the query timed out. - hits: an object with the following properties:    - total: the total number of hits found.    - hits: an array of hit objects, where each hit object represents a matched document. Each hit object has the following properties:      - _id: the ID of the matched document.      - _score: the score of the matched document.      - _source: the source data of the matched document.  In addition, if profiling is enabled, the response will include an additional array with profiling information attached. Here is an example search response:    &#x60;&#x60;&#x60;   {     &#39;took&#39;:10,     &#39;timed_out&#39;:false,     &#39;hits&#39;:     {       &#39;total&#39;:2,       &#39;hits&#39;:       [         {&#39;_id&#39;:&#39;1&#39;,&#39;_score&#39;:1,&#39;_source&#39;:{&#39;gid&#39;:11}},         {&#39;_id&#39;:&#39;2&#39;,&#39;_score&#39;:1,&#39;_source&#39;:{&#39;gid&#39;:12}}       ]     }   }   &#x60;&#x60;&#x60;  For more information about the match query syntax and additional parameters that can be added to request and response, please see the documentation [here](https://manual.manticoresearch.com/Searching/Full_text_matching/Basic_usage#HTTP-JSON). 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.SearchApi();
let searchRequest = ["'asas'"]; // SearchRequest | 
apiInstance.search(searchRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

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

