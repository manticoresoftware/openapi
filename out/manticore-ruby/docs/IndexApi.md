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

> BulkResponse bulk(body)

Bulk index operations

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
body = 'body_example' # String | 

begin
  #Bulk index operations
  result = api_instance.bulk(body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Exception when calling IndexApi->bulk: #{e}"
end
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **String**|  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-ndjson
- **Accept**: application/json


## delete

> DeleteResponse delete(delete_document_request)

Delete a document in an index

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
delete_document_request = {"index":"test","query":{"match":{"title":"apple"}}} # DeleteDocumentRequest | 

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

[**DeleteResponse**](DeleteResponse.md)

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
insert_document_request = {"index":"test","id":1,"doc":{"title":"sample title","gid":10}} # InsertDocumentRequest | 

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
insert_document_request = {"index":"test","id":1,"doc":{"title":"updated title","gid":15}} # InsertDocumentRequest | 

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

> UpdateResponse update(update_document_request)

Update a document in an index

### Example

```ruby
# load the gem
require 'openapi_client'

api_instance = OpenapiClient::IndexApi.new
update_document_request = {"index":"test","doc":{"gid":20},"query":{"equals":{"cat_id":2}}} # UpdateDocumentRequest | 

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

[**UpdateResponse**](UpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

