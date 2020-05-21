/*
 * delete_document_request.h
 *
 * Payload for delete request.
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
    object_t *doc; //object

} delete_document_request_t;

delete_document_request_t *delete_document_request_create(
    char *index,
    long id,
    object_t *doc
);

void delete_document_request_free(delete_document_request_t *delete_document_request);

delete_document_request_t *delete_document_request_parseFromJSON(cJSON *delete_document_requestJSON);

cJSON *delete_document_request_convertToJSON(delete_document_request_t *delete_document_request);

#endif /* _delete_document_request_H_ */

