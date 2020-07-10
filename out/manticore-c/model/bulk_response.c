#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "bulk_response.h"



bulk_response_t *bulk_response_create(
    object_t *items,
    int error
    ) {
    bulk_response_t *bulk_response_local_var = malloc(sizeof(bulk_response_t));
    if (!bulk_response_local_var) {
        return NULL;
    }
    bulk_response_local_var->items = items;
    bulk_response_local_var->error = error;

    return bulk_response_local_var;
}


void bulk_response_free(bulk_response_t *bulk_response) {
    if(NULL == bulk_response){
        return ;
    }
    listEntry_t *listEntry;
    object_free(bulk_response->items);
    free(bulk_response);
}

cJSON *bulk_response_convertToJSON(bulk_response_t *bulk_response) {
    cJSON *item = cJSON_CreateObject();

    // bulk_response->items
    if(bulk_response->items) { 
    cJSON *items_object = object_convertToJSON(bulk_response->items);
    if(items_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "items", items_object);
    if(item->child == NULL) {
    goto fail;
    }
     } 


    // bulk_response->error
    if(bulk_response->error) { 
    if(cJSON_AddBoolToObject(item, "error", bulk_response->error) == NULL) {
    goto fail; //Bool
    }
     } 

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

bulk_response_t *bulk_response_parseFromJSON(cJSON *bulk_responseJSON){

    bulk_response_t *bulk_response_local_var = NULL;

    // bulk_response->items
    cJSON *items = cJSON_GetObjectItemCaseSensitive(bulk_responseJSON, "items");
    object_t *items_local_object = NULL;
    if (items) { 
    items_local_object = object_parseFromJSON(items); //object
    }

    // bulk_response->error
    cJSON *error = cJSON_GetObjectItemCaseSensitive(bulk_responseJSON, "error");
    if (error) { 
    if(!cJSON_IsBool(error))
    {
    goto end; //Bool
    }
    }


    bulk_response_local_var = bulk_response_create (
        items ? items_local_object : NULL,
        error ? error->valueint : 0
        );

    return bulk_response_local_var;
end:
    return NULL;

}
