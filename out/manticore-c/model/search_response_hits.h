/*
 * search_response_hits.h
 *
 * 
 */

#ifndef _search_response_hits_H_
#define _search_response_hits_H_

#include <string.h>
#include "../external/cJSON.h"
#include "../include/list.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"

typedef struct search_response_hits_t search_response_hits_t;

#include "object.h"



typedef struct search_response_hits_t {
    int total; //numeric
    list_t *hits; //nonprimitive container

} search_response_hits_t;

search_response_hits_t *search_response_hits_create(
    int total,
    list_t *hits
);

void search_response_hits_free(search_response_hits_t *search_response_hits);

search_response_hits_t *search_response_hits_parseFromJSON(cJSON *search_response_hitsJSON);

cJSON *search_response_hits_convertToJSON(search_response_hits_t *search_response_hits);

#endif /* _search_response_hits_H_ */

