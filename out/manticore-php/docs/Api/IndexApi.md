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

