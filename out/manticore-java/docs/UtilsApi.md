# UtilsApi

All URIs are relative to *https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sql**](UtilsApi.md#sql) | **POST** /sql | Perform SQL requests


<a name="sql"></a>
# **sql**
> Map&lt;String, Object&gt; sql(body)

Perform SQL requests

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UtilsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://virtserver.swaggerhub.com/adriannuta/ManticoreSeach/1.0.0");

    UtilsApi apiInstance = new UtilsApi(defaultClient);
    String body = ["mode=raw&query=SHOW TABLES"]; // String | Expects is a query parameters string that can be in two modes: - Select only query as \"query=SELECT * FROM myindex\". The query string MUST be URL encoded - any type of query in format \"mode=raw&query=SHOW TABLES\". The string must be as is (no URL encoding) and `mode` must be first. 
    try {
      Map<String, Object> result = apiInstance.sql(body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UtilsApi#sql");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **String**| Expects is a query parameters string that can be in two modes: - Select only query as \&quot;query&#x3D;SELECT * FROM myindex\&quot;. The query string MUST be URL encoded - any type of query in format \&quot;mode&#x3D;raw&amp;query&#x3D;SHOW TABLES\&quot;. The string must be as is (no URL encoding) and &#x60;mode&#x60; must be first.  |

### Return type

**Map&lt;String, Object&gt;**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | In case of SELECT-only in mode none the response schema is the same as of &#x60;search&#x60;. In case of &#x60;mode&#x3D;raw&#x60; the response depends on the query executed.  |  -  |
**0** | error |  -  |

