#include <stdlib.h>
#include <stdio.h>
#include "../include/apiClient.h"
#include "../include/list.h"
#include "../external/cJSON.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"
#include "../model/error_response.h"
#include "../model/percolate_request.h"
#include "../model/search_request.h"
#include "../model/search_response.h"


// Perform reverse search on a percolate index
//
search_response_t*
SearchAPI_percolate(apiClient_t *apiClient, char * index , percolate_request_t * percolate_request );


// Performs a search
//
search_response_t*
SearchAPI_search(apiClient_t *apiClient, search_request_t * search_request );


