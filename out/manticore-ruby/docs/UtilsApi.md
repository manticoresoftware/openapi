# OpenapiClient::UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests



## sql

> Hash&lt;String, Object&gt; sql(body)

Perform SQL requests

Run a query in SQL format. <br/>  Expects is a query parameters string that can be in two modes: <br/>  * Select only query as `query=SELECT * FROM myindex`. The query string MUST be URL encoded <br/> * any type of query in format `mode=raw&query=SHOW TABLES`. The string must be as is (no URL encoding) and `mode` must be first. <br/>  The response object depends on the query executed. In select mode the response has same format as `/search` operation. 

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::UtilsApi.new
body = ["mode=raw&query=SHOW TABLES"] # String | Expects is a query parameters string that can be in two modes: - Select only query as `query=SELECT * FROM myindex`. The query string MUST be URL encoded - any type of query in format `mode=raw&query=SHOW TABLES`. The string must be as is (no URL encoding) and `mode` must be first. 

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
 **body** | **String**| Expects is a query parameters string that can be in two modes: - Select only query as &#x60;query&#x3D;SELECT * FROM myindex&#x60;. The query string MUST be URL encoded - any type of query in format &#x60;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES&#x60;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first.  | 

### Return type

**Hash&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

