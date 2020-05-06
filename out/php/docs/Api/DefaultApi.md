# OpenAPI\Client\DefaultApi

All URIs are relative to *https://virtserver.swaggerhub.com/ManticoreSearch/ManticoreSearch/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**insert**](DefaultApi.md#insert) | **POST** /json/insert | insert a document



## insert

> \OpenAPI\Client\Model\InsertResponse insert($insert)

insert a document

insert a document in a manticore index

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


$apiInstance = new OpenAPI\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$insert = new \OpenAPI\Client\Model\Insert(); // \OpenAPI\Client\Model\Insert | payload for insert

try {
    $result = $apiInstance->insert($insert);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->insert: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insert** | [**\OpenAPI\Client\Model\Insert**](../Model/Insert.md)| payload for insert | [optional]

### Return type

[**\OpenAPI\Client\Model\InsertResponse**](../Model/InsertResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)

