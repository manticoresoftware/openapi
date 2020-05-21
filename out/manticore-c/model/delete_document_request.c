#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "delete_document_request.h"



delete_document_request_t *delete_document_request_create(
    char *index,
    long id,
    object_t *doc
    ) {
    delete_document_request_t *delete_document_request_local_var = malloc(sizeof(delete_document_request_t));
    if (!delete_document_request_local_var) {
        return NULL;
    }
    delete_document_request_local_var->index = index;
    delete_document_request_local_var->id = id;
    delete_document_request_local_var->doc = doc;

    return delete_document_request_local_var;
}


void delete_document_request_free(delete_document_request_t *delete_document_request) {
    if(NULL == delete_document_request){
        return ;
    }
    listEntry_t *listEntry;
    free(delete_document_request->index);
    object_free(delete_document_request->doc);
    free(delete_document_request);
}

cJSON *delete_document_request_convertToJSON(delete_document_request_t *delete_document_request) {
    cJSON *item = cJSON_CreateObject();

    // delete_document_request->index
    if (!delete_document_request->index) {
        goto fail;
    }
    
    if(cJSON_AddStringToObject(item, "index", delete_document_request->index) == NULL) {
    goto fail; //String
    }


    // delete_document_request->id
    if(delete_document_request->id) { 
    if(cJSON_AddNumberToObject(item, "id", delete_document_request->id) == NULL) {
    goto fail; //Numeric
    }
     } 


    // delete_document_request->doc
    if(delete_document_request->doc) { 
    cJSON *doc_object = object_convertToJSON(delete_document_request->doc);
    if(doc_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "doc", doc_object);
    if(item->child == NULL) {
    goto fail;
    }
     } 

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

delete_document_request_t *delete_document_request_parseFromJSON(cJSON *delete_document_requestJSON){

    delete_document_request_t *delete_document_request_local_var = NULL;

    // delete_document_request->index
    cJSON *index = cJSON_GetObjectItemCaseSensitive(delete_document_requestJSON, "index");
    if (!index) {
        goto end;
    }

    
    if(!cJSON_IsString(index))
    {
    goto end; //String
    }

    // delete_document_request->id
    cJSON *id = cJSON_GetObjectItemCaseSensitive(delete_document_requestJSON, "id");
    if (id) { 
    if(!cJSON_IsNumber(id))
    {
    goto end; //Numeric
    }
    }

    // delete_document_request->doc
    cJSON *doc = cJSON_GetObjectItemCaseSensitive(delete_document_requestJSON, "doc");
    object_t *doc_local_object = NULL;
    if (doc) { 
    doc_local_object = object_parseFromJSON(doc); //object
    }


    delete_document_request_local_var = delete_document_request_create (
        strdup(index->valuestring),
        id ? id->valuedouble : 0,
        doc ? doc_local_object : NULL
        );

    return delete_document_request_local_var;
end:
    return NULL;

}
