# OpenapiClient::IndexApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bulk**](IndexApi.md#bulk) | **POST** /json/bulk | Bulk index operations
[**delete**](IndexApi.md#delete) | **POST** /json/delete | Delete a document in an index
[**insert**](IndexApi.md#insert) | **POST** /json/insert | Create a new document in an index
[**replace**](IndexApi.md#replace) | **POST** /json/replace | Replace new document in an index
[**update**](IndexApi.md#update) | **POST** /json/update | Update a document in an index



## bulk

> SuccessResponse bulk(request_body)

Bulk index operations

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
request_body = nil # Array<Object> | 

begin
  #Bulk index operations
  result = api_instance.bulk(request_body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling IndexApi->bulk: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request_body** | [**Array&lt;Object&gt;**](Object.md)|  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-ndjson
- **Accept**: application/json


## delete

> SuccessResponse delete(delete_document_request)

Delete a document in an index

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
delete_document_request = OpenapiClient::DeleteDocumentRequest.new # DeleteDocumentRequest | 

begin
  #Delete a document in an index
  result = api_instance.delete(delete_document_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling IndexApi->delete: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delete_document_request** | [**DeleteDocumentRequest**](DeleteDocumentRequest.md)|  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## insert

> SuccessResponse insert(insert_document_request)

Create a new document in an index

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
insert_document_request = OpenapiClient::InsertDocumentRequest.new # InsertDocumentRequest | 

begin
  #Create a new document in an index
  result = api_instance.insert(insert_document_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling IndexApi->insert: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insert_document_request** | [**InsertDocumentRequest**](InsertDocumentRequest.md)|  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## replace

> SuccessResponse replace(insert_document_request)

Replace new document in an index

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
insert_document_request = OpenapiClient::InsertDocumentRequest.new # InsertDocumentRequest | 

begin
  #Replace new document in an index
  result = api_instance.replace(insert_document_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling IndexApi->replace: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insert_document_request** | [**InsertDocumentRequest**](InsertDocumentRequest.md)|  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update

> SuccessResponse update(update_document_request)

Update a document in an index

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
update_document_request = OpenapiClient::UpdateDocumentRequest.new # UpdateDocumentRequest | 

begin
  #Update a document in an index
  result = api_instance.update(update_document_request)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling IndexApi->update: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **update_document_request** | [**UpdateDocumentRequest**](UpdateDocumentRequest.md)|  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

