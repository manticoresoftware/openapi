# WWW::OpenAPIClient::IndexApi

## Load the API package
```perl
use WWW::OpenAPIClient::Object::IndexApi;
```

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete**](IndexApi.md#delete) | **POST** /json/delete | Delete a document in an index
[**insert**](IndexApi.md#insert) | **POST** /json/insert | Create a new document in an index
[**replace**](IndexApi.md#replace) | **POST** /json/replace | Replace new document in an index
[**update**](IndexApi.md#update) | **POST** /json/update | Update a document in an index


# **delete**
> SuccessResponse delete(delete_document_request => $delete_document_request)

Delete a document in an index

### Example 
```perl
use Data::Dumper;
use WWW::OpenAPIClient::IndexApi;
my $api_instance = WWW::OpenAPIClient::IndexApi->new(
);

my $delete_document_request = WWW::OpenAPIClient::Object::DeleteDocumentRequest->new(); # DeleteDocumentRequest | 

eval { 
    my $result = $api_instance->delete(delete_document_request => $delete_document_request);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling IndexApi->delete: $@\n";
}
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **insert**
> SuccessResponse insert(insert_document_request => $insert_document_request)

Create a new document in an index

### Example 
```perl
use Data::Dumper;
use WWW::OpenAPIClient::IndexApi;
my $api_instance = WWW::OpenAPIClient::IndexApi->new(
);

my $insert_document_request = WWW::OpenAPIClient::Object::InsertDocumentRequest->new(); # InsertDocumentRequest | 

eval { 
    my $result = $api_instance->insert(insert_document_request => $insert_document_request);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling IndexApi->insert: $@\n";
}
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace**
> SuccessResponse replace(insert_document_request => $insert_document_request)

Replace new document in an index

### Example 
```perl
use Data::Dumper;
use WWW::OpenAPIClient::IndexApi;
my $api_instance = WWW::OpenAPIClient::IndexApi->new(
);

my $insert_document_request = WWW::OpenAPIClient::Object::InsertDocumentRequest->new(); # InsertDocumentRequest | 

eval { 
    my $result = $api_instance->replace(insert_document_request => $insert_document_request);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling IndexApi->replace: $@\n";
}
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update**
> SuccessResponse update(update_document_request => $update_document_request)

Update a document in an index

### Example 
```perl
use Data::Dumper;
use WWW::OpenAPIClient::IndexApi;
my $api_instance = WWW::OpenAPIClient::IndexApi->new(
);

my $update_document_request = WWW::OpenAPIClient::Object::UpdateDocumentRequest->new(); # UpdateDocumentRequest | 

eval { 
    my $result = $api_instance->update(update_document_request => $update_document_request);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling IndexApi->update: $@\n";
}
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

