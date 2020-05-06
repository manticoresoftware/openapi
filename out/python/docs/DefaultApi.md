# openapi_client.DefaultApi

All URIs are relative to *https://virtserver.swaggerhub.com/ManticoreSearch/ManticoreSearch/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_inventory**](DefaultApi.md#add_inventory) | **POST** /json/insert | insert a document


# **add_inventory**
> InsertResponse add_inventory(insert=insert)

insert a document

insert a document in a manticore index

### Example

```python
from __future__ import print_function
import time
import openapi_client
from openapi_client.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://virtserver.swaggerhub.com/ManticoreSearch/ManticoreSearch/1.0.0
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://virtserver.swaggerhub.com/ManticoreSearch/ManticoreSearch/1.0.0"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    insert = openapi_client.Insert() # Insert | Inventory item to add (optional)

    try:
        # insert a document
        api_response = api_instance.add_inventory(insert=insert)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DefaultApi->add_inventory: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insert** | [**Insert**](Insert.md)| Inventory item to add | [optional] 

### Return type

[**InsertResponse**](InsertResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | item inserted |  -  |
**500** | error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

