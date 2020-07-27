# OpenAPI\Client\IndexApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bulk**](IndexApi.md#bulk) | **POST** /json/bulk | Bulk index operations
[**delete**](IndexApi.md#delete) | **POST** /json/delete | Delete a document in an index
[**insert**](IndexApi.md#insert) | **POST** /json/insert | Create a new document in an index
[**replace**](IndexApi.md#replace) | **POST** /json/replace | Replace new document in an index
[**update**](IndexApi.md#update) | **POST** /json/update | Update a document in an index



## bulk

> \OpenAPI\Client\Model\BulkResponse bulk($body)

Bulk index operations

Sends multiple operatons like inserts, updates, replaces or deletes. <br/>  For each operation it's object must have same format as in their dedicated method. <br/>  The method expects a raw string as the batch in NDJSON.  Each operation object needs to be serialized to   JSON and separated by endline (\\n). <br/>     An example of raw input:      `{\"insert\": {\"index\": \"movies\", \"doc\": {\"plot\": \"A secret team goes to North Pole\", \"rating\": 9.5, \"language\": [2, 3], \"title\": \"This is an older movie\", \"lon\": 51.99, \"meta\": {\"keywords\":[\"travel\",\"ice\"],\"genre\":[\"adventure\"]}, \"year\": 1950, \"lat\": 60.4, \"advise\": \"PG-13\"}}}\\n{\"delete\": {\"index\": \"movies\",\"id\":700}}`      Responds with an object telling whenever any errors occured and an array with status for each operation:<br/>   ```   {'items':[{'update':{'_index':'products','_id':1,'result':'updated'}},{'update':{'_index':'products','_id':2,'result':'updated'}}],'errors':false}   ```

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


$apiInstance = new OpenAPI\Client\Api\IndexApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$body = 'body_example'; // string | 

try {
    $result = $apiInstance->bulk($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IndexApi->bulk: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string**|  |

### Return type

[**\OpenAPI\Client\Model\BulkResponse**](../Model/BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-ndjson
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## delete

> \OpenAPI\Client\Model\DeleteResponse delete($delete_document_request)

Delete a document in an index

Delete one or several documents. <br/>  The method has 2 ways of deleting: either by id, in case only one document is deleted or by using a  match query, in which case multiple documents can be delete . <br/> Example of input to delete by id: <br/> ```{'index':'movies','id':100}``` <br/> Example of input to delete using a query: ```{'index':'movies','query':{'bool':{'must':[{'query_string':'new movie'}]}}}``` <br/> The match query has same syntax as in for searching. <br/> Responds with an object telling how many documents got deleted: <br/>  ```{'_index':'products','updated':1}```

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


$apiInstance = new OpenAPI\Client\Api\IndexApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$delete_document_request = {"index":"test","query":{"match":{"title":"apple"}}}; // \OpenAPI\Client\Model\DeleteDocumentRequest | 

try {
    $result = $apiInstance->delete($delete_document_request);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IndexApi->delete: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delete_document_request** | [**\OpenAPI\Client\Model\DeleteDocumentRequest**](../Model/DeleteDocumentRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\DeleteResponse**](../Model/DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## insert

> \OpenAPI\Client\Model\SuccessResponse insert($insert_document_request)

Create a new document in an index

Insert a document. <br/> Expects an object like: <br/>  ```{'index':'movies','id':701,'doc':{'title':'This is an old movie','plot':'A secret team goes to North Pole','year':1950,'rating':9.5,'lat':60.4,'lon':51.99,'advise':'PG-13','meta':'{\"keywords\":{\"travel\",\"ice\"},\"genre\":{\"adventure\"}}','language':[2,3]}}```. <br/>  The document id can also be missing, in which case an autogenerated one will be used: <br/> ```{'index':'movies','doc':{'title':'This is a new movie','plot':'A secret team goes to North Pole','year':2020,'rating':9.5,'lat':60.4,'lon':51.99,'advise':'PG-13','meta':'{\"keywords\":{\"travel\",\"ice\"},\"genre\":{\"adventure\"}}','language':[2,3]}}``` <br/>  It responds with an object in format: <br/>  ```{'_index':'products','_id':701,'created':true,'result':'created','status':201}```

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


$apiInstance = new OpenAPI\Client\Api\IndexApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$insert_document_request = {"index":"test","id":1,"doc":{"title":"sample title","gid":10}}; // \OpenAPI\Client\Model\InsertDocumentRequest | 

try {
    $result = $apiInstance->insert($insert_document_request);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IndexApi->insert: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insert_document_request** | [**\OpenAPI\Client\Model\InsertDocumentRequest**](../Model/InsertDocumentRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\SuccessResponse**](../Model/SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## replace

> \OpenAPI\Client\Model\SuccessResponse replace($insert_document_request)

Replace new document in an index

Replace an existing document. Input has same format as `insert` operation. <br/>  Responds with an object in format: <br/>  `{'_index':'products','_id':1,'created':false,'result':'updated','status':200}`

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


$apiInstance = new OpenAPI\Client\Api\IndexApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$insert_document_request = {"index":"test","id":1,"doc":{"title":"updated title","gid":15}}; // \OpenAPI\Client\Model\InsertDocumentRequest | 

try {
    $result = $apiInstance->replace($insert_document_request);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IndexApi->replace: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insert_document_request** | [**\OpenAPI\Client\Model\InsertDocumentRequest**](../Model/InsertDocumentRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\SuccessResponse**](../Model/SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## update

> \OpenAPI\Client\Model\UpdateResponse update($update_document_request)

Update a document in an index

Update one or several documents. <br/>  The update can be made by passing the id or by using a match query in case multiple documents can be updated. <br/> For example update a document using document id: <br/> <code> {'index':'movies','doc':{'rating':9.49},'id':100} </code> <br/> And update by using a match query: <br/> ``` {'index':'movies','doc':{'rating':9.49},'query':{'bool':{'must':[{'query_string':'new movie'}]}}} ```  <br/> The match query has same syntax as for searching.  Responds with an object that tells how many documents where updated in format: <br/> ```{'_index':'products','updated':1}```

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


$apiInstance = new OpenAPI\Client\Api\IndexApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$update_document_request = {"index":"test","doc":{"gid":20},"query":{"equals":{"cat_id":2}}}; // \OpenAPI\Client\Model\UpdateDocumentRequest | 

try {
    $result = $apiInstance->update($update_document_request);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IndexApi->update: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_document_request** | [**\OpenAPI\Client\Model\UpdateDocumentRequest**](../Model/UpdateDocumentRequest.md)|  |

### Return type

[**\OpenAPI\Client\Model\UpdateResponse**](../Model/UpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)

