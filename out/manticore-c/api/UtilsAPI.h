#include <stdlib.h>
#include <stdio.h>
#include "../include/apiClient.h"
#include "../include/list.h"
#include "../external/cJSON.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"
#include "../model/error_response.h"


// Perform SQL requests
//
// Run a query in SQL format. <br/>  Expects is a query parameters string that can be in two modes: <br/>  * Select only query as `query=SELECT * FROM myindex`. The query string MUST be URL encoded <br/> * any type of query in format `mode=raw&query=SHOW TABLES`. The string must be as is (no URL encoding) and `mode` must be first. <br/>  The response object depends on the query executed. In select mode the response has same format as `/search` operation. 
//
list_t*_t*
UtilsAPI_sql(apiClient_t *apiClient, char * body );


