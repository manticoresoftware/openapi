# \IndexApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Bulk**](IndexApi.md#Bulk) | **Post** /json/bulk | Bulk index operations
[**Delete**](IndexApi.md#Delete) | **Post** /json/delete | Delete a document in an index
[**Insert**](IndexApi.md#Insert) | **Post** /json/insert | Create a new document in an index
[**Replace**](IndexApi.md#Replace) | **Post** /json/replace | Replace new document in an index
[**Update**](IndexApi.md#Update) | **Post** /json/update | Update a document in an index



## Bulk

> BulkResponse Bulk(ctx).Body(body).Execute()

Bulk index operations

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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndexApi.Bulk(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.Bulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Bulk`: BulkResponse
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.Bulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**BulkResponse**](bulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-ndjson
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> DeleteResponse Delete(ctx).DeleteDocumentRequest(deleteDocumentRequest).Execute()

Delete a document in an index

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
    deleteDocumentRequest := openapiclient.deleteDocumentRequest{Index: "Index_example", Id: int64(123), Query: 123} // DeleteDocumentRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndexApi.Delete(context.Background(), deleteDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.Delete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDocumentRequest** | [**DeleteDocumentRequest**](DeleteDocumentRequest.md) |  | 

### Return type

[**DeleteResponse**](deleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Insert

> SuccessResponse Insert(ctx).InsertDocumentRequest(insertDocumentRequest).Execute()

Create a new document in an index

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
    insertDocumentRequest := openapiclient.insertDocumentRequest{Index: "Index_example", Id: int64(123), Doc: map[string]string{ "Key" = "Value" }} // InsertDocumentRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndexApi.Insert(context.Background(), insertDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.Insert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Insert`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.Insert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md) |  | 

### Return type

[**SuccessResponse**](successResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Replace

> SuccessResponse Replace(ctx).InsertDocumentRequest(insertDocumentRequest).Execute()

Replace new document in an index

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
    insertDocumentRequest := openapiclient.insertDocumentRequest{Index: "Index_example", Id: int64(123), Doc: map[string]string{ "Key" = "Value" }} // InsertDocumentRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndexApi.Replace(context.Background(), insertDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.Replace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Replace`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.Replace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md) |  | 

### Return type

[**SuccessResponse**](successResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateResponse Update(ctx).UpdateDocumentRequest(updateDocumentRequest).Execute()

Update a document in an index

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
    updateDocumentRequest := openapiclient.updateDocumentRequest{Index: "Index_example", Doc: map[string]string{ "Key" = "Value" }, Id: int64(123), Query: map[string]string{ "Key" = "Value" }} // UpdateDocumentRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndexApi.Update(context.Background(), updateDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: UpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDocumentRequest** | [**UpdateDocumentRequest**](UpdateDocumentRequest.md) |  | 

### Return type

[**UpdateResponse**](updateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

