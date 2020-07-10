#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "success_response.h"



success_response_t *success_response_create(
    char *_index,
    long _id,
    int created,
    char *result,
    int found
    ) {
    success_response_t *success_response_local_var = malloc(sizeof(success_response_t));
    if (!success_response_local_var) {
        return NULL;
    }
    success_response_local_var->_index = _index;
    success_response_local_var->_id = _id;
    success_response_local_var->created = created;
    success_response_local_var->result = result;
    success_response_local_var->found = found;

    return success_response_local_var;
}


void success_response_free(success_response_t *success_response) {
    if(NULL == success_response){
        return ;
    }
    listEntry_t *listEntry;
    free(success_response->_index);
    free(success_response->result);
    free(success_response);
}

cJSON *success_response_convertToJSON(success_response_t *success_response) {
    cJSON *item = cJSON_CreateObject();

    // success_response->_index
    if(success_response->_index) { 
    if(cJSON_AddStringToObject(item, "_index", success_response->_index) == NULL) {
    goto fail; //String
    }
     } 


    // success_response->_id
    if(success_response->_id) { 
    if(cJSON_AddNumberToObject(item, "_id", success_response->_id) == NULL) {
    goto fail; //Numeric
    }
     } 


    // success_response->created
    if(success_response->created) { 
    if(cJSON_AddBoolToObject(item, "created", success_response->created) == NULL) {
    goto fail; //Bool
    }
     } 


    // success_response->result
    if(success_response->result) { 
    if(cJSON_AddStringToObject(item, "result", success_response->result) == NULL) {
    goto fail; //String
    }
     } 


    // success_response->found
    if(success_response->found) { 
    if(cJSON_AddBoolToObject(item, "found", success_response->found) == NULL) {
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

success_response_t *success_response_parseFromJSON(cJSON *success_responseJSON){

    success_response_t *success_response_local_var = NULL;

    // success_response->_index
    cJSON *_index = cJSON_GetObjectItemCaseSensitive(success_responseJSON, "_index");
    if (_index) { 
    if(!cJSON_IsString(_index))
    {
    goto end; //String
    }
    }

    // success_response->_id
    cJSON *_id = cJSON_GetObjectItemCaseSensitive(success_responseJSON, "_id");
    if (_id) { 
    if(!cJSON_IsNumber(_id))
    {
    goto end; //Numeric
    }
    }

    // success_response->created
    cJSON *created = cJSON_GetObjectItemCaseSensitive(success_responseJSON, "created");
    if (created) { 
    if(!cJSON_IsBool(created))
    {
    goto end; //Bool
    }
    }

    // success_response->result
    cJSON *result = cJSON_GetObjectItemCaseSensitive(success_responseJSON, "result");
    if (result) { 
    if(!cJSON_IsString(result))
    {
    goto end; //String
    }
    }

    // success_response->found
    cJSON *found = cJSON_GetObjectItemCaseSensitive(success_responseJSON, "found");
    if (found) { 
    if(!cJSON_IsBool(found))
    {
    goto end; //Bool
    }
    }


    success_response_local_var = success_response_create (
        _index ? strdup(_index->valuestring) : NULL,
        _id ? _id->valuedouble : 0,
        created ? created->valueint : 0,
        result ? strdup(result->valuestring) : NULL,
        found ? found->valueint : 0
        );

    return success_response_local_var;
end:
    return NULL;

}
