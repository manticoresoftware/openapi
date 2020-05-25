# OpenapiClient::UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests



## sql

> Hash&lt;String, Object&gt; sql(query, opts)

Perform SQL requests

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::UtilsApi.new
query = 'query_example' # String | 
opts = {
  mode: 'mode_example' # String | 
}

begin
  #Perform SQL requests
  result = api_instance.sql(query, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling UtilsApi->sql: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **String**|  | 
 **mode** | **String**|  | [optional] 

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

