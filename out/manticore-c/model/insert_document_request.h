/*
 * insert_document_request.h
 *
 * Object with document data. 
 */

#ifndef _insert_document_request_H_
#define _insert_document_request_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct insert_document_request_t insert_document_request_t;

#include "object.h"



typedef struct insert_document_request_t {
    char *index; // string
    long id; //numeric
    list_t* doc; //map

} insert_document_request_t;

insert_document_request_t *insert_document_request_create(
    char *index,
    long id,
    list_t* doc
);

void insert_document_request_free(insert_document_request_t *insert_document_request);

insert_document_request_t *insert_document_request_parseFromJSON(cJSON *insert_document_requestJSON);

cJSON *insert_document_request_convertToJSON(insert_document_request_t *insert_document_request);

#endif /* _insert_document_request_H_ */

