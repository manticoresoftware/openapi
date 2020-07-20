# OpenapiClient::UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests



## sql

> Hash&lt;String, Object&gt; sql(body)

Perform SQL requests

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::UtilsApi.new
body = ["mode=raw&query=SHOW TABLES"] # String | Expects is a query parameters string that can be in two modes: - Select only query as \"query=SELECT * FROM myindex\". The query string MUST be URL encoded - any type of query in format \"mode=raw&query=SHOW TABLES\". The string must be as is (no URL encoding) and `mode` must be first. 

begin
  #Perform SQL requests
  result = api_instance.sql(body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling UtilsApi->sql: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **String**| Expects is a query parameters string that can be in two modes: - Select only query as \&quot;query&#x3D;SELECT * FROM myindex\&quot;. The query string MUST be URL encoded - any type of query in format \&quot;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES\&quot;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first.  | 

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

