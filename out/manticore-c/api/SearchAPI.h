#include <stdlib.h>
#include <stdio.h>
#include "../include/apiClient.h"
#include "../include/list.h"
#include "../external/cJSON.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"
#include "../model/error_response.h"
#include "../model/search_request.h"
#include "../model/search_response.h"


// Performs a search
//
search_response_t*
SearchAPI_search(apiClient_t *apiClient, search_request_t * search_request );


