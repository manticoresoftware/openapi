# OpenapiClient::SearchApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**percolate**](SearchApi.md#percolate) | **POST** /json/pq/search | Perform reverse search on a percolate index
[**search**](SearchApi.md#search) | **POST** /json/search | Performs a search



## percolate

> SearchResponse percolate(percolate_request)

Perform reverse search on a percolate index

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::SearchApi.new
percolate_request = OpenapiClient::PercolateRequest.new # PercolateRequest | 

begin
  #Perform reverse search on a percolate index
  result = api_instance.percolate(percolate_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling SearchApi->percolate: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **percolate_request** | [**PercolateRequest**](PercolateRequest.md)|  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


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

