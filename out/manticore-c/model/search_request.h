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
#include "one_ofstringobject.h"



typedef struct search_request_t {
    char *index; // string
    object_t *query; //object
    int limit; //numeric
    int offset; //numeric
    int max_matches; //numeric
    list_t *sort; //nonprimitive container
    object_t *script_fields; //object
    object_t *highlight; //object
    struct one_ofstringobject_t *_source; //model
    int profile; //boolean

} search_request_t;

search_request_t *search_request_create(
    char *index,
    object_t *query,
    int limit,
    int offset,
    int max_matches,
    list_t *sort,
    object_t *script_fields,
    object_t *highlight,
    one_ofstringobject_t *_source,
    int profile
);

void search_request_free(search_request_t *search_request);

search_request_t *search_request_parseFromJSON(cJSON *search_requestJSON);

cJSON *search_request_convertToJSON(search_request_t *search_request);

#endif /* _search_request_H_ */

