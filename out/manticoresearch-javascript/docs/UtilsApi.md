# Manticoresearch.UtilsApi

All URIs are relative to *http://127.0.0.1:9308*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests



## sql

> [Object] sql(body, opts)

Perform SQL requests

Run a query in SQL format. Expects a query string passed through &#x60;body&#x60; parameter and optional &#x60;raw_response&#x60; parameter that defines a format of response. &#x60;raw_response&#x60; can be set to &#x60;False&#x60; for Select queries only, e.g., &#x60;SELECT * FROM myindex&#x60; The query string must stay as it is, no URL encoding is needed. The response object depends on the query executed. In select mode the response has same format as &#x60;/search&#x60; operation. 

### Example

```javascript
import Manticoresearch from 'manticoresearch';

let apiInstance = new Manticoresearch.UtilsApi();
let body = SHOW TABLES; // String | A query parameter string. 
let opts = {
  'rawResponse': true // Boolean | Optional parameter, defines a format of response. Can be set to `False` for Select only queries and set to `True` or omitted for any type of queries: 
};
apiInstance.sql(body, opts).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **String**| A query parameter string.  | 
 **rawResponse** | **Boolean**| Optional parameter, defines a format of response. Can be set to &#x60;False&#x60; for Select only queries and set to &#x60;True&#x60; or omitted for any type of queries:  | [optional] [default to true]

### Return type

**[Object]**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

