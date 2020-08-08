# ManticoreSearchApi.SearchApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**percolate**](SearchApi.md#percolate) | **POST** /json/pq/{index}/search | Perform reverse search on a percolate index
[**search**](SearchApi.md#search) | **POST** /json/search | Performs a search



## percolate

> SearchResponse percolate(index, percolateRequest)

Perform reverse search on a percolate index

Performs a percolate search. &lt;br/&gt; This method must be used only on percolate indexes. &lt;br/&gt; Expects two paramenters: the index name and an object with array of documents to be tested. &lt;br/&gt; An example of the documents object: &lt;br/&gt; &#x60;&#x60;&#x60; {\&quot;query\&quot;:{\&quot;percolate\&quot;:{\&quot;document\&quot;:{\&quot;content\&quot;:\&quot;sample content\&quot;}}}} &#x60;&#x60;&#x60; &lt;br/&gt; Responds with an object with matched stored queries: &lt;br/&gt; &#x60;&#x60;&#x60; {&#39;timed_out&#39;:false,&#39;hits&#39;:{&#39;total&#39;:2,&#39;max_score&#39;:1,&#39;hits&#39;:[{&#39;_index&#39;:&#39;idx_pq_1&#39;,&#39;_type&#39;:&#39;doc&#39;,&#39;_id&#39;:&#39;2&#39;,&#39;_score&#39;:&#39;1&#39;,&#39;_source&#39;:{&#39;query&#39;:{&#39;match&#39;:{&#39;title&#39;:&#39;some&#39;},}}},{&#39;_index&#39;:&#39;idx_pq_1&#39;,&#39;_type&#39;:&#39;doc&#39;,&#39;_id&#39;:&#39;5&#39;,&#39;_score&#39;:&#39;1&#39;,&#39;_source&#39;:{&#39;query&#39;:{&#39;ql&#39;:&#39;some | none&#39;}}}]}} &#x60;&#x60;&#x60; 

### Example

```javascript
import ManticoreSearchApi from 'manticore_search_api';

let apiInstance = new ManticoreSearchApi.SearchApi();
let index = "index_example"; // String | Name of the percolate index
let percolateRequest = {"query":{"percolate":{"document":{"title":"some text to match"}}}}; // PercolateRequest | 
apiInstance.percolate(index, percolateRequest, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
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

Performs a search

Performs a search. &lt;br/&gt; Expects an object with mandatory properties: &lt;br/&gt; - the index name &lt;br/&gt; - the match query object &lt;br/&gt; Example : &lt;br/&gt; &lt;code&gt; {&#39;index&#39;:&#39;movies&#39;,&#39;query&#39;:{&#39;bool&#39;:{&#39;must&#39;:[{&#39;query_string&#39;:&#39; movie&#39;}]}},&#39;script_fields&#39;:{&#39;myexpr&#39;:{&#39;script&#39;:{&#39;inline&#39;:&#39;IF(rating&gt;8,1,0)&#39;}}},&#39;sort&#39;:[{&#39;myexpr&#39;:&#39;desc&#39;},{&#39;_score&#39;:&#39;desc&#39;}],&#39;profile&#39;:true} &lt;/code&gt; &lt;br/&gt; It responds with an object with &lt;br/&gt; - time of execution &lt;br/&gt; - if the query timed out &lt;br/&gt; - an array with hits (matched documents) &lt;br/&gt; - additional, if profiling is enabled, an array with profiling information is attached &lt;br/&gt;  &#x60;&#x60;&#x60; {&#39;took&#39;:10,&#39;timed_out&#39;:false,&#39;hits&#39;:{&#39;total&#39;:2,&#39;hits&#39;:[{&#39;_id&#39;:&#39;1&#39;,&#39;_score&#39;:1,&#39;_source&#39;:{&#39;gid&#39;:11}},{&#39;_id&#39;:&#39;2&#39;,&#39;_score&#39;:1,&#39;_source&#39;:{&#39;gid&#39;:12}}]}} &#x60;&#x60;&#x60; &lt;br/&gt; For more information about the match query syntax, additional paramaters that can be set to the input and response, please check: https://docs.manticoresearch.com/latest/html/http_reference/json_search.html. 

### Example

```javascript
import ManticoreSearchApi from 'manticore_search_api';

let apiInstance = new ManticoreSearchApi.SearchApi();
let searchRequest = new ManticoreSearchApi.SearchRequest(); // SearchRequest | 
apiInstance.search(searchRequest, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
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

