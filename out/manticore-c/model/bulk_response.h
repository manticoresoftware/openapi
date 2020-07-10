/*
 * bulk_response.h
 *
 * Success bulk response
 */

#ifndef _bulk_response_H_
#define _bulk_response_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct bulk_response_t bulk_response_t;

#include "object.h"



typedef struct bulk_response_t {
    object_t *items; //object
    int error; //boolean

} bulk_response_t;

bulk_response_t *bulk_response_create(
    object_t *items,
    int error
);

void bulk_response_free(bulk_response_t *bulk_response);

bulk_response_t *bulk_response_parseFromJSON(cJSON *bulk_responseJSON);

cJSON *bulk_response_convertToJSON(bulk_response_t *bulk_response);

#endif /* _bulk_response_H_ */

