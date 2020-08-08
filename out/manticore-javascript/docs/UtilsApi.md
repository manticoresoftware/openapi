# ManticoreSearchApi.UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests



## sql

> {String: Object} sql(body)

Perform SQL requests

Run a query in SQL format. &lt;br/&gt; Expects is a query parameters string that can be in two modes: &lt;br/&gt; * Select only query as &#x60;query&#x3D;SELECT * FROM myindex&#x60;. The query string MUST be URL encoded &lt;br/&gt; * any type of query in format &#x60;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES&#x60;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first. &lt;br/&gt; The response object depends on the query executed. In select mode the response has same format as &#x60;/search&#x60; operation. 

### Example

```javascript
import ManticoreSearchApi from 'manticore_search_api';

let apiInstance = new ManticoreSearchApi.UtilsApi();
let body = ["mode=raw&query=SHOW TABLES"]; // String | Expects is a query parameters string that can be in two modes: - Select only query as `query=SELECT * FROM myindex`. The query string MUST be URL encoded - any type of query in format `mode=raw&query=SHOW TABLES`. The string must be as is (no URL encoding) and `mode` must be first. 
apiInstance.sql(body, (error, data, response) => {
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
 **body** | **String**| Expects is a query parameters string that can be in two modes: - Select only query as &#x60;query&#x3D;SELECT * FROM myindex&#x60;. The query string MUST be URL encoded - any type of query in format &#x60;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES&#x60;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first.  | 

### Return type

**{String: Object}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

