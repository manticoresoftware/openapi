/*
 * delete_response.h
 *
 * Success response
 */

#ifndef _delete_response_H_
#define _delete_response_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct delete_response_t delete_response_t;




typedef struct delete_response_t {
    char *_index; // string
    int deleted; //numeric
    long _id; //numeric
    char *result; // string

} delete_response_t;

delete_response_t *delete_response_create(
    char *_index,
    int deleted,
    long _id,
    char *result
);

void delete_response_free(delete_response_t *delete_response);

delete_response_t *delete_response_parseFromJSON(cJSON *delete_responseJSON);

cJSON *delete_response_convertToJSON(delete_response_t *delete_response);

#endif /* _delete_response_H_ */

