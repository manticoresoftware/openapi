/*
 * delete_document_request.h
 *
 * Payload for delete request. Documents can be deleted either one by one by specifying the document id or by providing a query object. For more information see  [Delete API](https://docs.manticoresearch.com/latest/html/http_reference/json_delete.html) 
 */

#ifndef _delete_document_request_H_
#define _delete_document_request_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct delete_document_request_t delete_document_request_t;

#include "object.h"



typedef struct delete_document_request_t {
    char *index; // string
    long id; //numeric
    object_t *query; //object

} delete_document_request_t;

delete_document_request_t *delete_document_request_create(
    char *index,
    long id,
    object_t *query
);

void delete_document_request_free(delete_document_request_t *delete_document_request);

delete_document_request_t *delete_document_request_parseFromJSON(cJSON *delete_document_requestJSON);

cJSON *delete_document_request_convertToJSON(delete_document_request_t *delete_document_request);

#endif /* _delete_document_request_H_ */

