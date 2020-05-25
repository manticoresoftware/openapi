/*
 * percolate_request_query.h
 *
 * 
 */

#ifndef _percolate_request_query_H_
#define _percolate_request_query_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct percolate_request_query_t percolate_request_query_t;

#include "object.h"



typedef struct percolate_request_query_t {
    object_t *percolate; //object

} percolate_request_query_t;

percolate_request_query_t *percolate_request_query_create(
    object_t *percolate
);

void percolate_request_query_free(percolate_request_query_t *percolate_request_query);

percolate_request_query_t *percolate_request_query_parseFromJSON(cJSON *percolate_request_queryJSON);

cJSON *percolate_request_query_convertToJSON(percolate_request_query_t *percolate_request_query);

#endif /* _percolate_request_query_H_ */

