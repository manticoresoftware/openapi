/*
 * search_request.h
 *
 * Payload for search operation
 */

#ifndef _search_request_H_
#define _search_request_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct search_request_t search_request_t;

#include "object.h"



typedef struct search_request_t {
    char *index; // string
    list_t* query; //map
    int limit; //numeric
    int offset; //numeric
    int max_matches; //numeric
    list_t *sort; //nonprimitive container
    object_t *script_fields; //object
    object_t *highlight; //object
    list_t *_source; //primitive container
    int profile; //boolean

} search_request_t;

search_request_t *search_request_create(
    char *index,
    list_t* query,
    int limit,
    int offset,
    int max_matches,
    list_t *sort,
    object_t *script_fields,
    object_t *highlight,
    list_t *_source,
    int profile
);

void search_request_free(search_request_t *search_request);

search_request_t *search_request_parseFromJSON(cJSON *search_requestJSON);

cJSON *search_request_convertToJSON(search_request_t *search_request);

#endif /* _search_request_H_ */

