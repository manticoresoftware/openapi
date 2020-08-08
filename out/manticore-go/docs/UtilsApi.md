# \UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sql**](UtilsApi.md#Sql) | **Post** /sql | Perform SQL requests



## Sql

> map[string]map[string]interface{} Sql(ctx, body)

Perform SQL requests

Run a query in SQL format. <br/> Expects is a query parameters string that can be in two modes: <br/> * Select only query as `query=SELECT * FROM myindex`. The query string MUST be URL encoded <br/> * any type of query in format `mode=raw&query=SHOW TABLES`. The string must be as is (no URL encoding) and `mode` must be first. <br/> The response object depends on the query executed. In select mode the response has same format as `/search` operation. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**| Expects is a query parameters string that can be in two modes: - Select only query as &#x60;query&#x3D;SELECT * FROM myindex&#x60;. The query string MUST be URL encoded - any type of query in format &#x60;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES&#x60;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first.  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

