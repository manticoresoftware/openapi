#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "update_document_request.h"



update_document_request_t *update_document_request_create(
    char *index,
    object_t *doc,
    long id,
    object_t *query
    ) {
    update_document_request_t *update_document_request_local_var = malloc(sizeof(update_document_request_t));
    if (!update_document_request_local_var) {
        return NULL;
    }
    update_document_request_local_var->index = index;
    update_document_request_local_var->doc = doc;
    update_document_request_local_var->id = id;
    update_document_request_local_var->query = query;

    return update_document_request_local_var;
}


void update_document_request_free(update_document_request_t *update_document_request) {
    if(NULL == update_document_request){
        return ;
    }
    listEntry_t *listEntry;
    free(update_document_request->index);
    object_free(update_document_request->doc);
    object_free(update_document_request->query);
    free(update_document_request);
}

cJSON *update_document_request_convertToJSON(update_document_request_t *update_document_request) {
    cJSON *item = cJSON_CreateObject();

    // update_document_request->index
    if (!update_document_request->index) {
        goto fail;
    }
    
    if(cJSON_AddStringToObject(item, "index", update_document_request->index) == NULL) {
    goto fail; //String
    }


    // update_document_request->doc
    if (!update_document_request->doc) {
        goto fail;
    }
    
    cJSON *doc_object = object_convertToJSON(update_document_request->doc);
    if(doc_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "doc", doc_object);
    if(item->child == NULL) {
    goto fail;
    }


    // update_document_request->id
    if(update_document_request->id) { 
    if(cJSON_AddNumberToObject(item, "id", update_document_request->id) == NULL) {
    goto fail; //Numeric
    }
     } 


    // update_document_request->query
    if(update_document_request->query) { 
    cJSON *query_object = object_convertToJSON(update_document_request->query);
    if(query_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "query", query_object);
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

update_document_request_t *update_document_request_parseFromJSON(cJSON *update_document_requestJSON){

    update_document_request_t *update_document_request_local_var = NULL;

    // update_document_request->index
    cJSON *index = cJSON_GetObjectItemCaseSensitive(update_document_requestJSON, "index");
    if (!index) {
        goto end;
    }

    
    if(!cJSON_IsString(index))
    {
    goto end; //String
    }

    // update_document_request->doc
    cJSON *doc = cJSON_GetObjectItemCaseSensitive(update_document_requestJSON, "doc");
    if (!doc) {
        goto end;
    }

    object_t *doc_local_object = NULL;
    
    doc_local_object = object_parseFromJSON(doc); //object

    // update_document_request->id
    cJSON *id = cJSON_GetObjectItemCaseSensitive(update_document_requestJSON, "id");
    if (id) { 
    if(!cJSON_IsNumber(id))
    {
    goto end; //Numeric
    }
    }

    // update_document_request->query
    cJSON *query = cJSON_GetObjectItemCaseSensitive(update_document_requestJSON, "query");
    object_t *query_local_object = NULL;
    if (query) { 
    query_local_object = object_parseFromJSON(query); //object
    }


    update_document_request_local_var = update_document_request_create (
        strdup(index->valuestring),
        doc_local_object,
        id ? id->valuedouble : 0,
        query ? query_local_object : NULL
        );

    return update_document_request_local_var;
end:
    return NULL;

}
