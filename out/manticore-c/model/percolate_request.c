#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "percolate_request.h"



percolate_request_t *percolate_request_create(
    percolate_request_query_t *query
    ) {
    percolate_request_t *percolate_request_local_var = malloc(sizeof(percolate_request_t));
    if (!percolate_request_local_var) {
        return NULL;
    }
    percolate_request_local_var->query = query;

    return percolate_request_local_var;
}


void percolate_request_free(percolate_request_t *percolate_request) {
    if(NULL == percolate_request){
        return ;
    }
    listEntry_t *listEntry;
    percolate_request_query_free(percolate_request->query);
    free(percolate_request);
}

cJSON *percolate_request_convertToJSON(percolate_request_t *percolate_request) {
    cJSON *item = cJSON_CreateObject();

    // percolate_request->query
    if (!percolate_request->query) {
        goto fail;
    }
    
    cJSON *query_local_JSON = percolate_request_query_convertToJSON(percolate_request->query);
    if(query_local_JSON == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "query", query_local_JSON);
    if(item->child == NULL) {
    goto fail;
    }

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

percolate_request_t *percolate_request_parseFromJSON(cJSON *percolate_requestJSON){

    percolate_request_t *percolate_request_local_var = NULL;

    // percolate_request->query
    cJSON *query = cJSON_GetObjectItemCaseSensitive(percolate_requestJSON, "query");
    if (!query) {
        goto end;
    }

    percolate_request_query_t *query_local_nonprim = NULL;
    
    query_local_nonprim = percolate_request_query_parseFromJSON(query); //nonprimitive


    percolate_request_local_var = percolate_request_create (
        query_local_nonprim
        );

    return percolate_request_local_var;
end:
    return NULL;

}
