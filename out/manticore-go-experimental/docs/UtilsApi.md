# \UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sql**](UtilsApi.md#Sql) | **Post** /sql | Perform SQL requests



## Sql

> map[string]map[string]interface{} Sql(ctx).Body(body).Execute()

Perform SQL requests

Run a query in SQL format.
Expects is a query parameters string that can be in two modes:
* Select only query as `query=SELECT * FROM myindex`. The query string MUST be URL encoded
* any type of query in format `mode=raw&query=SHOW TABLES`. The string must be as is (no URL encoding) and `mode` must be first.
The response object depends on the query executed. In select mode the response has same format as `/search` operation.


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := "body_example" // string | Expects is a query parameters string that can be in two modes:    * Select only query as `query=SELECT * FROM myindex`. The query string MUST be URL encoded    * any type of query in format `mode=raw&query=SHOW TABLES`. The string must be as is (no URL encoding) and `mode` must be first. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UtilsApi.Sql(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsApi.Sql``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sql`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UtilsApi.Sql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Expects is a query parameters string that can be in two modes:    * Select only query as &#x60;query&#x3D;SELECT * FROM myindex&#x60;. The query string MUST be URL encoded    * any type of query in format &#x60;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES&#x60;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first.  | 

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

