# ManticoreSearchApi.SearchApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**search**](SearchApi.md#search) | **POST** /json/search | Performs a search



## search

> SearchResponse search(searchRequest)

Performs a search

### Example

```javascript
import ManticoreSearchApi from 'manticore_search_api';

let apiInstance = new ManticoreSearchApi.SearchApi();
let searchRequest = new ManticoreSearchApi.SearchRequest(); // SearchRequest | 
apiInstance.search(searchRequest, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md)|  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

