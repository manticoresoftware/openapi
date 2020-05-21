/*
 * success_response.h
 *
 * 
 */

#ifndef _success_response_H_
#define _success_response_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct success_response_t success_response_t;




typedef struct success_response_t {
    char *_index; // string
    long _id; //numeric
    int created; //boolean
    char *result; // string
    int found; //boolean

} success_response_t;

success_response_t *success_response_create(
    char *_index,
    long _id,
    int created,
    char *result,
    int found
);

void success_response_free(success_response_t *success_response);

success_response_t *success_response_parseFromJSON(cJSON *success_responseJSON);

cJSON *success_response_convertToJSON(success_response_t *success_response);

#endif /* _success_response_H_ */

