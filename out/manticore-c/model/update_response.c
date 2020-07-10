#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "update_response.h"



update_response_t *update_response_create(
    char *_index,
    int updated,
    long _id,
    char *result
    ) {
    update_response_t *update_response_local_var = malloc(sizeof(update_response_t));
    if (!update_response_local_var) {
        return NULL;
    }
    update_response_local_var->_index = _index;
    update_response_local_var->updated = updated;
    update_response_local_var->_id = _id;
    update_response_local_var->result = result;

    return update_response_local_var;
}


void update_response_free(update_response_t *update_response) {
    if(NULL == update_response){
        return ;
    }
    listEntry_t *listEntry;
    free(update_response->_index);
    free(update_response->result);
    free(update_response);
}

cJSON *update_response_convertToJSON(update_response_t *update_response) {
    cJSON *item = cJSON_CreateObject();

    // update_response->_index
    if(update_response->_index) { 
    if(cJSON_AddStringToObject(item, "_index", update_response->_index) == NULL) {
    goto fail; //String
    }
     } 


    // update_response->updated
    if(update_response->updated) { 
    if(cJSON_AddNumberToObject(item, "updated", update_response->updated) == NULL) {
    goto fail; //Numeric
    }
     } 


    // update_response->_id
    if(update_response->_id) { 
    if(cJSON_AddNumberToObject(item, "_id", update_response->_id) == NULL) {
    goto fail; //Numeric
    }
     } 


    // update_response->result
    if(update_response->result) { 
    if(cJSON_AddStringToObject(item, "result", update_response->result) == NULL) {
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

update_response_t *update_response_parseFromJSON(cJSON *update_responseJSON){

    update_response_t *update_response_local_var = NULL;

    // update_response->_index
    cJSON *_index = cJSON_GetObjectItemCaseSensitive(update_responseJSON, "_index");
    if (_index) { 
    if(!cJSON_IsString(_index))
    {
    goto end; //String
    }
    }

    // update_response->updated
    cJSON *updated = cJSON_GetObjectItemCaseSensitive(update_responseJSON, "updated");
    if (updated) { 
    if(!cJSON_IsNumber(updated))
    {
    goto end; //Numeric
    }
    }

    // update_response->_id
    cJSON *_id = cJSON_GetObjectItemCaseSensitive(update_responseJSON, "_id");
    if (_id) { 
    if(!cJSON_IsNumber(_id))
    {
    goto end; //Numeric
    }
    }

    // update_response->result
    cJSON *result = cJSON_GetObjectItemCaseSensitive(update_responseJSON, "result");
    if (result) { 
    if(!cJSON_IsString(result))
    {
    goto end; //String
    }
    }


    update_response_local_var = update_response_create (
        _index ? strdup(_index->valuestring) : NULL,
        updated ? updated->valuedouble : 0,
        _id ? _id->valuedouble : 0,
        result ? strdup(result->valuestring) : NULL
        );

    return update_response_local_var;
end:
    return NULL;

}
