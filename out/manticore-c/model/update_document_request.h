/*
 * update_document_request.h
 *
 * 
 */

#ifndef _update_document_request_H_
#define _update_document_request_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct update_document_request_t update_document_request_t;

#include "object.h"



typedef struct update_document_request_t {
    char *index; // string
    object_t *doc; //object
    long id; //numeric
    object_t *query; //object

} update_document_request_t;

update_document_request_t *update_document_request_create(
    char *index,
    object_t *doc,
    long id,
    object_t *query
);

void update_document_request_free(update_document_request_t *update_document_request);

update_document_request_t *update_document_request_parseFromJSON(cJSON *update_document_requestJSON);

cJSON *update_document_request_convertToJSON(update_document_request_t *update_document_request);

#endif /* _update_document_request_H_ */

