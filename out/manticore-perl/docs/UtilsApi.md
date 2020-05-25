# WWW::OpenAPIClient::UtilsApi

## Load the API package
```perl
use WWW::OpenAPIClient::Object::UtilsApi;
```

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests


# **sql**
> HASH[string,object] sql(query => $query, mode => $mode)

Perform SQL requests

### Example 
```perl
use Data::Dumper;
use WWW::OpenAPIClient::UtilsApi;
my $api_instance = WWW::OpenAPIClient::UtilsApi->new(
);

my $query = "query_example"; # string | 
my $mode = "mode_example"; # string | 

eval { 
    my $result = $api_instance->sql(query => $query, mode => $mode);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling UtilsApi->sql: $@\n";
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string**|  | 
 **mode** | **string**|  | [optional] 

### Return type

**HASH[string,object]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

