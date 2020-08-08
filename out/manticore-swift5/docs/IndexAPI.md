# IndexAPI

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bulk**](IndexAPI.md#bulk) | **POST** /json/bulk | Bulk index operations
[**delete**](IndexAPI.md#delete) | **POST** /json/delete | Delete a document in an index
[**insert**](IndexAPI.md#insert) | **POST** /json/insert | Create a new document in an index
[**replace**](IndexAPI.md#replace) | **POST** /json/replace | Replace new document in an index
[**update**](IndexAPI.md#update) | **POST** /json/update | Update a document in an index


# **bulk**
```swift
    open class func bulk(body: String, completion: @escaping (_ data: BulkResponse?, _ error: Error?) -> Void)
```

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
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let body = "body_example" // String | 

// Bulk index operations
IndexAPI.bulk(body: body) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **String** |  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-ndjson
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete**
```swift
    open class func delete(deleteDocumentRequest: DeleteDocumentRequest, completion: @escaping (_ data: DeleteResponse?, _ error: Error?) -> Void)
```

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
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let deleteDocumentRequest = deleteDocumentRequest(index: "index_example", id: 123, query: 123) // DeleteDocumentRequest | 

// Delete a document in an index
IndexAPI.delete(deleteDocumentRequest: deleteDocumentRequest) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDocumentRequest** | [**DeleteDocumentRequest**](DeleteDocumentRequest.md) |  | 

### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **insert**
```swift
    open class func insert(insertDocumentRequest: InsertDocumentRequest, completion: @escaping (_ data: SuccessResponse?, _ error: Error?) -> Void)
```

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
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let insertDocumentRequest = insertDocumentRequest(index: "index_example", id: 123, doc: "TODO") // InsertDocumentRequest | 

// Create a new document in an index
IndexAPI.insert(insertDocumentRequest: insertDocumentRequest) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace**
```swift
    open class func replace(insertDocumentRequest: InsertDocumentRequest, completion: @escaping (_ data: SuccessResponse?, _ error: Error?) -> Void)
```

Replace new document in an index

Replace an existing document. Input has same format as `insert` operation. <br/>
Responds with an object in format: <br/>

  ```
  {'_index':'products','_id':1,'created':false,'result':'updated','status':200}
  ```


### Example 
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let insertDocumentRequest = insertDocumentRequest(index: "index_example", id: 123, doc: "TODO") // InsertDocumentRequest | 

// Replace new document in an index
IndexAPI.replace(insertDocumentRequest: insertDocumentRequest) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md) |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update**
```swift
    open class func update(updateDocumentRequest: UpdateDocumentRequest, completion: @escaping (_ data: UpdateResponse?, _ error: Error?) -> Void)
```

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
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let updateDocumentRequest = updateDocumentRequest(index: "index_example", doc: "TODO", id: 123, query: "TODO") // UpdateDocumentRequest | 

// Update a document in an index
IndexAPI.update(updateDocumentRequest: updateDocumentRequest) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDocumentRequest** | [**UpdateDocumentRequest**](UpdateDocumentRequest.md) |  | 

### Return type

[**UpdateResponse**](UpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

