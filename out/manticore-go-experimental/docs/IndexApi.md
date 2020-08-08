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

Sends multiple operatons like inserts, updates, replaces or deletes. 
For each operation it's object must have same format as in their dedicated method. 
The method expects a raw string as the batch in NDJSON.
 Each operation object needs to be serialized to 
 JSON and separated by endline (\n). 
 
  An example of raw input:
  
  ```
  {"insert": {"index": "movies", "doc": {"plot": "A secret team goes to North Pole", "rating": 9.5, "language": [2, 3], "title": "This is an older movie", "lon": 51.99, "meta": {"keywords":["travel","ice"],"genre":["adventure"]}, "year": 1950, "lat": 60.4, "advise": "PG-13"}}}
  \n
  {"delete": {"index": "movies","id":700}}
  ```
  
  Responds with an object telling whenever any errors occured and an array with status for each operation:
  
  ```
  {'items':[{'update':{'_index':'products','_id':1,'result':'updated'}},{'update':{'_index':'products','_id':2,'result':'updated'}}],'errors':false}
  ```
 


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

Delete one or several documents.
The method has 2 ways of deleting: either by id, in case only one document is deleted or by using a  match query, in which case multiple documents can be delete .
Example of input to delete by id:

  ```
  {'index':'movies','id':100}
  ```

Example of input to delete using a query:

  ```
  {'index':'movies','query':{'bool':{'must':[{'query_string':'new movie'}]}}}
  ```

The match query has same syntax as in for searching.
Responds with an object telling how many documents got deleted: 

  ```
  {'_index':'products','updated':1}
  ```


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

Insert a document. 
Expects an object like:
 
  ```
  {'index':'movies','id':701,'doc':{'title':'This is an old movie','plot':'A secret team goes to North Pole','year':1950,'rating':9.5,'lat':60.4,'lon':51.99,'advise':'PG-13','meta':'{"keywords":{"travel","ice"},"genre":{"adventure"}}','language':[2,3]}}
  ```
 
The document id can also be missing, in which case an autogenerated one will be used:
         
  ```
  {'index':'movies','doc':{'title':'This is a new movie','plot':'A secret team goes to North Pole','year':2020,'rating':9.5,'lat':60.4,'lon':51.99,'advise':'PG-13','meta':'{"keywords":{"travel","ice"},"genre":{"adventure"}}','language':[2,3]}}
  ```
 
It responds with an object in format:
  
  ```
  {'_index':'products','_id':701,'created':true,'result':'created','status':201}
  ```


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

Replace an existing document. Input has same format as `insert` operation. <br/>
Responds with an object in format: <br/>

  ```
  {'_index':'products','_id':1,'created':false,'result':'updated','status':200}
  ```


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

Update one or several documents.
The update can be made by passing the id or by using a match query in case multiple documents can be updated.  For example update a document using document id:

  ```
  {'index':'movies','doc':{'rating':9.49},'id':100}
  ```

And update by using a match query:

  ```
  {'index':'movies','doc':{'rating':9.49},'query':{'bool':{'must':[{'query_string':'new movie'}]}}}
  ``` 

The match query has same syntax as for searching.
Responds with an object that tells how many documents where updated in format: 

  ```
  {'_index':'products','updated':1}
  ```


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

