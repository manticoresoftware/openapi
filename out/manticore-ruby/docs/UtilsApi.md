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
body = 'body_example' # String | 

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
 **body** | **String**|  | 

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

