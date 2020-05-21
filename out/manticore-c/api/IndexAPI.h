#include <stdlib.h>
#include <stdio.h>
#include "../include/apiClient.h"
#include "../include/list.h"
#include "../external/cJSON.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"
#include "../model/delete_document_request.h"
#include "../model/error_response.h"
#include "../model/insert_document_request.h"
#include "../model/success_response.h"
#include "../model/update_document_request.h"


// Delete a document in an index
//
success_response_t*
IndexAPI_delete(apiClient_t *apiClient, delete_document_request_t * delete_document_request );


// Create a new document in an index
//
success_response_t*
IndexAPI_insert(apiClient_t *apiClient, insert_document_request_t * insert_document_request );


// Replace new document in an index
//
success_response_t*
IndexAPI_replace(apiClient_t *apiClient, insert_document_request_t * insert_document_request );


// Update a document in an index
//
success_response_t*
IndexAPI_update(apiClient_t *apiClient, update_document_request_t * update_document_request );


