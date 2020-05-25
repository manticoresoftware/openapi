#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "error_response.h"



error_response_t *error_response_create(
    list_t* error,
    int status
    ) {
    error_response_t *error_response_local_var = malloc(sizeof(error_response_t));
    if (!error_response_local_var) {
        return NULL;
    }
    error_response_local_var->error = error;
    error_response_local_var->status = status;

    return error_response_local_var;
}


void error_response_free(error_response_t *error_response) {
    if(NULL == error_response){
        return ;
    }
    listEntry_t *listEntry;
    list_ForEach(listEntry, error_response->error) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*) listEntry->data;
        free (localKeyValue->key);
        free (localKeyValue->value);
    }
    list_free(error_response->error);
    free(error_response);
}

cJSON *error_response_convertToJSON(error_response_t *error_response) {
    cJSON *item = cJSON_CreateObject();

    // error_response->error
    if (!error_response->error) {
        goto fail;
    }
    
    cJSON *error = cJSON_AddObjectToObject(item, "error");
    if(error == NULL) {
        goto fail; //primitive map container
    }
    cJSON *localMapObject = error;
    listEntry_t *errorListEntry;
    if (error_response->error) {
    list_ForEach(errorListEntry, error_response->error) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*)errorListEntry->data;
    }
    }


    // error_response->status
    if (!error_response->status) {
        goto fail;
    }
    
    if(cJSON_AddNumberToObject(item, "status", error_response->status) == NULL) {
    goto fail; //Numeric
    }

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

error_response_t *error_response_parseFromJSON(cJSON *error_responseJSON){

    error_response_t *error_response_local_var = NULL;

    // error_response->error
    cJSON *error = cJSON_GetObjectItemCaseSensitive(error_responseJSON, "error");
    if (!error) {
        goto end;
    }

    list_t *errorList;
    
    cJSON *error_local_map;
    if(!cJSON_IsObject(error)) {
        goto end;//primitive map container
    }
    errorList = list_create();
    keyValuePair_t *localMapKeyPair;
    cJSON_ArrayForEach(error_local_map, error)
    {
		cJSON *localMapObject = error_local_map;
        list_addElement(errorList , localMapKeyPair);
    }

    // error_response->status
    cJSON *status = cJSON_GetObjectItemCaseSensitive(error_responseJSON, "status");
    if (!status) {
        goto end;
    }

    
    if(!cJSON_IsNumber(status))
    {
    goto end; //Numeric
    }


    error_response_local_var = error_response_create (
        errorList,
        status->valuedouble
        );

    return error_response_local_var;
end:
    return NULL;

}
