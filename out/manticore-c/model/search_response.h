/*
 * search_response.h
 *
 * Response object of a search request
 */

#ifndef _search_response_H_
#define _search_response_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct search_response_t search_response_t;

#include "object.h"
#include "search_response_hits.h"



typedef struct search_response_t {
    int took; //numeric
    int timed_out; //boolean
    list_t* hits; //map
    object_t *profile; //object

} search_response_t;

search_response_t *search_response_create(
    int took,
    int timed_out,
    list_t* hits,
    object_t *profile
);

void search_response_free(search_response_t *search_response);

search_response_t *search_response_parseFromJSON(cJSON *search_responseJSON);

cJSON *search_response_convertToJSON(search_response_t *search_response);

#endif /* _search_response_H_ */

