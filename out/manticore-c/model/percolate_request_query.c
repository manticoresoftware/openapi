#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "percolate_request_query.h"



percolate_request_query_t *percolate_request_query_create(
    object_t *percolate
    ) {
    percolate_request_query_t *percolate_request_query_local_var = malloc(sizeof(percolate_request_query_t));
    if (!percolate_request_query_local_var) {
        return NULL;
    }
    percolate_request_query_local_var->percolate = percolate;

    return percolate_request_query_local_var;
}


void percolate_request_query_free(percolate_request_query_t *percolate_request_query) {
    if(NULL == percolate_request_query){
        return ;
    }
    listEntry_t *listEntry;
    object_free(percolate_request_query->percolate);
    free(percolate_request_query);
}

cJSON *percolate_request_query_convertToJSON(percolate_request_query_t *percolate_request_query) {
    cJSON *item = cJSON_CreateObject();

    // percolate_request_query->percolate
    if (!percolate_request_query->percolate) {
        goto fail;
    }
    
    cJSON *percolate_object = object_convertToJSON(percolate_request_query->percolate);
    if(percolate_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "percolate", percolate_object);
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

percolate_request_query_t *percolate_request_query_parseFromJSON(cJSON *percolate_request_queryJSON){

    percolate_request_query_t *percolate_request_query_local_var = NULL;

    // percolate_request_query->percolate
    cJSON *percolate = cJSON_GetObjectItemCaseSensitive(percolate_request_queryJSON, "percolate");
    if (!percolate) {
        goto end;
    }

    object_t *percolate_local_object = NULL;
    
    percolate_local_object = object_parseFromJSON(percolate); //object


    percolate_request_query_local_var = percolate_request_query_create (
        percolate_local_object
        );

    return percolate_request_query_local_var;
end:
    return NULL;

}
