# OpenAPI\Client\UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests



## sql

> map[string,object] sql($body)

Perform SQL requests

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


$apiInstance = new OpenAPI\Client\Api\UtilsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$body = ["mode=raw&query=SHOW TABLES"]; // string | Expects is a query parameters string that can be in two modes: - Select only query as \"query=SELECT * FROM myindex\". The query string MUST be URL encoded - any type of query in format \"mode=raw&query=SHOW TABLES\". The string must be as is (no URL encoding) and `mode` must be first.

try {
    $result = $apiInstance->sql($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling UtilsApi->sql: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string**| Expects is a query parameters string that can be in two modes: - Select only query as \&quot;query&#x3D;SELECT * FROM myindex\&quot;. The query string MUST be URL encoded - any type of query in format \&quot;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES\&quot;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first. |

### Return type

**map[string,object]**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)

