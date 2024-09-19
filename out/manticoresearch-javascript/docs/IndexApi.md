# Manticoresearch.IndexApi

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bulk**](IndexApi.md#bulk) | **POST** /bulk | Bulk index operations
[**callDelete**](IndexApi.md#callDelete) | **POST** /delete | Delete a document in an index
[**insert**](IndexApi.md#insert) | **POST** /insert | Create a new document in an index
[**partialReplace**](IndexApi.md#partialReplace) | **POST** /{index}/_update/{id} | Partially replaces a document in an index
[**replace**](IndexApi.md#replace) | **POST** /replace | Replace new document in an index
[**update**](IndexApi.md#update) | **POST** /update | Update a document in an index



## bulk

> BulkResponse bulk(body)

Bulk index operations

Sends multiple operatons like inserts, updates, replaces or deletes.  For each operation it&#39;s object must have same format as in their dedicated method.  The method expects a raw string as the batch in NDJSON.  Each operation object needs to be serialized to   JSON and separated by endline (\\n).      An example of raw input:      &#x60;&#x60;&#x60;   {\&quot;insert\&quot;: {\&quot;index\&quot;: \&quot;movies\&quot;, \&quot;doc\&quot;: {\&quot;plot\&quot;: \&quot;A secret team goes to North Pole\&quot;, \&quot;rating\&quot;: 9.5, \&quot;language\&quot;: [2, 3], \&quot;title\&quot;: \&quot;This is an older movie\&quot;, \&quot;lon\&quot;: 51.99, \&quot;meta\&quot;: {\&quot;keywords\&quot;:[\&quot;travel\&quot;,\&quot;ice\&quot;],\&quot;genre\&quot;:[\&quot;adventure\&quot;]}, \&quot;year\&quot;: 1950, \&quot;lat\&quot;: 60.4, \&quot;advise\&quot;: \&quot;PG-13\&quot;}}}   \\n   {\&quot;delete\&quot;: {\&quot;index\&quot;: \&quot;movies\&quot;,\&quot;id\&quot;:700}}   &#x60;&#x60;&#x60;      Responds with an object telling whenever any errors occured and an array with status for each operation:      &#x60;&#x60;&#x60;   {     &#39;items&#39;:     [       {         &#39;update&#39;:{&#39;_index&#39;:&#39;products&#39;,&#39;_id&#39;:1,&#39;result&#39;:&#39;updated&#39;}       },       {         &#39;update&#39;:{&#39;_index&#39;:&#39;products&#39;,&#39;_id&#39;:2,&#39;result&#39;:&#39;updated&#39;}       }     ],     &#39;errors&#39;:false   }   &#x60;&#x60;&#x60; 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.IndexApi();
let body = "body_example"; // String | 
apiInstance.bulk(body).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **String**|  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-ndjson
- **Accept**: application/json


## callDelete

> DeleteResponse callDelete(deleteDocumentRequest)

Delete a document in an index

Delete one or several documents. The method has 2 ways of deleting: either by id, in case only one document is deleted or by using a  match query, in which case multiple documents can be delete . Example of input to delete by id:    &#x60;&#x60;&#x60;   {&#39;index&#39;:&#39;movies&#39;,&#39;id&#39;:100}   &#x60;&#x60;&#x60;  Example of input to delete using a query:    &#x60;&#x60;&#x60;   {     &#39;index&#39;:&#39;movies&#39;,     &#39;query&#39;:     {       &#39;bool&#39;:       {         &#39;must&#39;:         [           {&#39;query_string&#39;:&#39;new movie&#39;}         ]       }     }   }   &#x60;&#x60;&#x60;  The match query has same syntax as in for searching. Responds with an object telling how many documents got deleted:     &#x60;&#x60;&#x60;   {&#39;_index&#39;:&#39;products&#39;,&#39;updated&#39;:1}   &#x60;&#x60;&#x60; 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.IndexApi();
let deleteDocumentRequest = {"index":"test","query":{"match":{"title":"apple"}}}; // DeleteDocumentRequest | 
apiInstance.callDelete(deleteDocumentRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDocumentRequest** | [**DeleteDocumentRequest**](DeleteDocumentRequest.md)|  | 

### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## insert

> SuccessResponse insert(insertDocumentRequest)

Create a new document in an index

Insert a document.  Expects an object like:     &#x60;&#x60;&#x60;   {     &#39;index&#39;:&#39;movies&#39;,     &#39;id&#39;:701,     &#39;doc&#39;:     {       &#39;title&#39;:&#39;This is an old movie&#39;,       &#39;plot&#39;:&#39;A secret team goes to North Pole&#39;,       &#39;year&#39;:1950,       &#39;rating&#39;:9.5,       &#39;lat&#39;:60.4,       &#39;lon&#39;:51.99,       &#39;advise&#39;:&#39;PG-13&#39;,       &#39;meta&#39;:&#39;{\&quot;keywords\&quot;:{\&quot;travel\&quot;,\&quot;ice\&quot;},\&quot;genre\&quot;:{\&quot;adventure\&quot;}}&#39;,       &#39;language&#39;:[2,3]     }   }   &#x60;&#x60;&#x60;   The document id can also be missing, in which case an autogenerated one will be used:             &#x60;&#x60;&#x60;   {     &#39;index&#39;:&#39;movies&#39;,     &#39;doc&#39;:     {       &#39;title&#39;:&#39;This is a new movie&#39;,       &#39;plot&#39;:&#39;A secret team goes to North Pole&#39;,       &#39;year&#39;:2020,       &#39;rating&#39;:9.5,       &#39;lat&#39;:60.4,       &#39;lon&#39;:51.99,       &#39;advise&#39;:&#39;PG-13&#39;,       &#39;meta&#39;:&#39;{\&quot;keywords\&quot;:{\&quot;travel\&quot;,\&quot;ice\&quot;},\&quot;genre\&quot;:{\&quot;adventure\&quot;}}&#39;,       &#39;language&#39;:[2,3]     }   }   &#x60;&#x60;&#x60;   It responds with an object in format:      &#x60;&#x60;&#x60;   {&#39;_index&#39;:&#39;products&#39;,&#39;_id&#39;:701,&#39;created&#39;:true,&#39;result&#39;:&#39;created&#39;,&#39;status&#39;:201}   &#x60;&#x60;&#x60; 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.IndexApi();
let insertDocumentRequest = {"index":"test","id":1,"doc":{"title":"sample title","gid":10}}; // InsertDocumentRequest | 
apiInstance.insert(insertDocumentRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md)|  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## partialReplace

> UpdateResponse partialReplace(index, id, replaceDocumentRequest)

Partially replaces a document in an index

Partially replaces a document with given id in an index Responds with an object of the following format:     &#x60;&#x60;&#x60;   {&#39;_index&#39;:&#39;products&#39;,&#39;updated&#39;:1}   &#x60;&#x60;&#x60; 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.IndexApi();
let index = "index_example"; // String | Name of the percolate index
let id = 3.4; // Number | Id of the document to replace
let replaceDocumentRequest = {"doc":{"price":20}}; // ReplaceDocumentRequest | 
apiInstance.partialReplace(index, id, replaceDocumentRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **String**| Name of the percolate index | 
 **id** | **Number**| Id of the document to replace | 
 **replaceDocumentRequest** | [**ReplaceDocumentRequest**](ReplaceDocumentRequest.md)|  | 

### Return type

[**UpdateResponse**](UpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## replace

> SuccessResponse replace(insertDocumentRequest)

Replace new document in an index

Replace an existing document. Input has same format as &#x60;insert&#x60; operation. &lt;br/&gt; Responds with an object in format: &lt;br/&gt;    &#x60;&#x60;&#x60;   {&#39;_index&#39;:&#39;products&#39;,&#39;_id&#39;:1,&#39;created&#39;:false,&#39;result&#39;:&#39;updated&#39;,&#39;status&#39;:200}   &#x60;&#x60;&#x60; 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.IndexApi();
let insertDocumentRequest = {"index":"test","id":1,"doc":{"title":"updated title","gid":15}}; // InsertDocumentRequest | 
apiInstance.replace(insertDocumentRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md)|  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update

> UpdateResponse update(updateDocumentRequest)

Update a document in an index

Update one or several documents. The update can be made by passing the id or by using a match query in case multiple documents can be updated.  For example update a document using document id:    &#x60;&#x60;&#x60;   {&#39;index&#39;:&#39;movies&#39;,&#39;doc&#39;:{&#39;rating&#39;:9.49},&#39;id&#39;:100}   &#x60;&#x60;&#x60;  And update by using a match query:    &#x60;&#x60;&#x60;   {     &#39;index&#39;:&#39;movies&#39;,     &#39;doc&#39;:{&#39;rating&#39;:9.49},     &#39;query&#39;:     {       &#39;bool&#39;:       {         &#39;must&#39;:         [           {&#39;query_string&#39;:&#39;new movie&#39;}         ]       }     }   }   &#x60;&#x60;&#x60;   The match query has same syntax as for searching. Responds with an object that tells how many documents where updated in format:     &#x60;&#x60;&#x60;   {&#39;_index&#39;:&#39;products&#39;,&#39;updated&#39;:1}   &#x60;&#x60;&#x60; 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.IndexApi();
let updateDocumentRequest = {"index":"test","doc":{"gid":20},"query":{"equals":{"cat_id":2}}}; // UpdateDocumentRequest | 
apiInstance.update(updateDocumentRequest).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDocumentRequest** | [**UpdateDocumentRequest**](UpdateDocumentRequest.md)|  | 

### Return type

[**UpdateResponse**](UpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

