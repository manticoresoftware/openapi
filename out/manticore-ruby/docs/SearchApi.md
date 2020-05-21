# OpenapiClient::SearchApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**search**](SearchApi.md#search) | **POST** /json/search | Performs a search



## search

> SearchResponse search(search_request)

Performs a search

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::SearchApi.new
search_request = OpenapiClient::SearchRequest.new # SearchRequest | 

begin
  #Performs a search
  result = api_instance.search(search_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling SearchApi->search: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search_request** | [**SearchRequest**](SearchRequest.md)|  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

