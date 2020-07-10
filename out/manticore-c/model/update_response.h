/*
 * update_response.h
 *
 * Success response
 */

#ifndef _update_response_H_
#define _update_response_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct update_response_t update_response_t;




typedef struct update_response_t {
    char *_index; // string
    int updated; //numeric
    long _id; //numeric
    char *result; // string

} update_response_t;

update_response_t *update_response_create(
    char *_index,
    int updated,
    long _id,
    char *result
);

void update_response_free(update_response_t *update_response);

update_response_t *update_response_parseFromJSON(cJSON *update_responseJSON);

cJSON *update_response_convertToJSON(update_response_t *update_response);

#endif /* _update_response_H_ */

