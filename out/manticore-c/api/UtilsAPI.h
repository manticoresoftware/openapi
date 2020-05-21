#include <stdlib.h>
#include <stdio.h>
#include "../include/apiClient.h"
#include "../include/list.h"
#include "../external/cJSON.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"
#include "../model/error_response.h"
#include "../model/object.h"


// Perform SQL requests
//
object_t*
UtilsAPI_sql(apiClient_t *apiClient, char * query , char * mode );


