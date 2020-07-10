#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "delete_response.h"



delete_response_t *delete_response_create(
    char *_index,
    int deleted,
    long _id,
    char *result
    ) {
    delete_response_t *delete_response_local_var = malloc(sizeof(delete_response_t));
    if (!delete_response_local_var) {
        return NULL;
    }
    delete_response_local_var->_index = _index;
    delete_response_local_var->deleted = deleted;
    delete_response_local_var->_id = _id;
    delete_response_local_var->result = result;

    return delete_response_local_var;
}


void delete_response_free(delete_response_t *delete_response) {
    if(NULL == delete_response){
        return ;
    }
    listEntry_t *listEntry;
    free(delete_response->_index);
    free(delete_response->result);
    free(delete_response);
}

cJSON *delete_response_convertToJSON(delete_response_t *delete_response) {
    cJSON *item = cJSON_CreateObject();

    // delete_response->_index
    if(delete_response->_index) { 
    if(cJSON_AddStringToObject(item, "_index", delete_response->_index) == NULL) {
    goto fail; //String
    }
     } 


    // delete_response->deleted
    if(delete_response->deleted) { 
    if(cJSON_AddNumberToObject(item, "deleted", delete_response->deleted) == NULL) {
    goto fail; //Numeric
    }
     } 


    // delete_response->_id
    if(delete_response->_id) { 
    if(cJSON_AddNumberToObject(item, "_id", delete_response->_id) == NULL) {
    goto fail; //Numeric
    }
     } 


    // delete_response->result
    if(delete_response->result) { 
    if(cJSON_AddStringToObject(item, "result", delete_response->result) == NULL) {
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

delete_response_t *delete_response_parseFromJSON(cJSON *delete_responseJSON){

    delete_response_t *delete_response_local_var = NULL;

    // delete_response->_index
    cJSON *_index = cJSON_GetObjectItemCaseSensitive(delete_responseJSON, "_index");
    if (_index) { 
    if(!cJSON_IsString(_index))
    {
    goto end; //String
    }
    }

    // delete_response->deleted
    cJSON *deleted = cJSON_GetObjectItemCaseSensitive(delete_responseJSON, "deleted");
    if (deleted) { 
    if(!cJSON_IsNumber(deleted))
    {
    goto end; //Numeric
    }
    }

    // delete_response->_id
    cJSON *_id = cJSON_GetObjectItemCaseSensitive(delete_responseJSON, "_id");
    if (_id) { 
    if(!cJSON_IsNumber(_id))
    {
    goto end; //Numeric
    }
    }

    // delete_response->result
    cJSON *result = cJSON_GetObjectItemCaseSensitive(delete_responseJSON, "result");
    if (result) { 
    if(!cJSON_IsString(result))
    {
    goto end; //String
    }
    }


    delete_response_local_var = delete_response_create (
        _index ? strdup(_index->valuestring) : NULL,
        deleted ? deleted->valuedouble : 0,
        _id ? _id->valuedouble : 0,
        result ? strdup(result->valuestring) : NULL
        );

    return delete_response_local_var;
end:
    return NULL;

}
