# ManticoreSearchApi.UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests



## sql

> Object sql(query, opts)

Perform SQL requests

### Example

```javascript
import ManticoreSearchApi from 'manticore_search_api';

let apiInstance = new ManticoreSearchApi.UtilsApi();
let query = "query_example"; // String | 
let opts = {
  'mode': "mode_example" // String | 
};
apiInstance.sql(query, opts, (error, data, response) => {
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
 **query** | **String**|  | 
 **mode** | **String**|  | [optional] 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

