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
> HASH[string,object] sql(body => $body)

Perform SQL requests

### Example 
```perl
use Data::Dumper;
use WWW::OpenAPIClient::UtilsApi;
my $api_instance = WWW::OpenAPIClient::UtilsApi->new(
);

my $body = WWW::OpenAPIClient::Object::string->new(); # string | Expects is a query parameters string that can be in two modes: - Select only query as \"query=SELECT * FROM myindex\". The query string MUST be URL encoded - any type of query in format \"mode=raw&query=SHOW TABLES\". The string must be as is (no URL encoding) and `mode` must be first. 

eval { 
    my $result = $api_instance->sql(body => $body);
    print Dumper($result);
};
if ($@) {
    warn "Exception when calling UtilsApi->sql: $@\n";
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string**| Expects is a query parameters string that can be in two modes: - Select only query as \&quot;query&#x3D;SELECT * FROM myindex\&quot;. The query string MUST be URL encoded - any type of query in format \&quot;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES\&quot;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first.  | 

### Return type

**HASH[string,object]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

