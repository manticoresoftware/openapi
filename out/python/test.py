from __future__ import print_function

import time
import openapi_client
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://virtserver.swaggerhub.com/ManticoreSearch/ManticoreSearch/1.0.0
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://127.0.0.1:6368"
)



# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    insert = openapi_client.Insert(id=104,index='testrt2',doc={"title":"some text","gid":1}) # Insert | Inventory item to add (optional)
   
    try:
        # insert a document
        api_response = api_instance.insert(insert=insert)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DefaultApi->insert: %s\n" % e)
