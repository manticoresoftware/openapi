#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "inline_object.h"



inline_object_t *inline_object_create(
    char *query,
    char *mode
    ) {
    inline_object_t *inline_object_local_var = malloc(sizeof(inline_object_t));
    if (!inline_object_local_var) {
        return NULL;
    }
    inline_object_local_var->query = query;
    inline_object_local_var->mode = mode;

    return inline_object_local_var;
}


void inline_object_free(inline_object_t *inline_object) {
    if(NULL == inline_object){
        return ;
    }
    listEntry_t *listEntry;
    free(inline_object->query);
    free(inline_object->mode);
    free(inline_object);
}

cJSON *inline_object_convertToJSON(inline_object_t *inline_object) {
    cJSON *item = cJSON_CreateObject();

    // inline_object->query
    if (!inline_object->query) {
        goto fail;
    }
    
    if(cJSON_AddStringToObject(item, "query", inline_object->query) == NULL) {
    goto fail; //String
    }


    // inline_object->mode
    if(inline_object->mode) { 
    if(cJSON_AddStringToObject(item, "mode", inline_object->mode) == NULL) {
    goto fail; //String
    }
     } 

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

inline_object_t *inline_object_parseFromJSON(cJSON *inline_objectJSON){

    inline_object_t *inline_object_local_var = NULL;

    // inline_object->query
    cJSON *query = cJSON_GetObjectItemCaseSensitive(inline_objectJSON, "query");
    if (!query) {
        goto end;
    }

    
    if(!cJSON_IsString(query))
    {
    goto end; //String
    }

    // inline_object->mode
    cJSON *mode = cJSON_GetObjectItemCaseSensitive(inline_objectJSON, "mode");
    if (mode) { 
    if(!cJSON_IsString(mode))
    {
    goto end; //String
    }
    }


    inline_object_local_var = inline_object_create (
        strdup(query->valuestring),
        mode ? strdup(mode->valuestring) : NULL
        );

    return inline_object_local_var;
end:
    return NULL;

}
