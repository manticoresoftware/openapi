#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "update_document_request.h"



update_document_request_t *update_document_request_create(
    char *index,
    list_t* doc,
    long id,
    list_t* query
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
    list_ForEach(listEntry, update_document_request->doc) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*) listEntry->data;
        free (localKeyValue->key);
        free (localKeyValue->value);
    }
    list_free(update_document_request->doc);
    list_ForEach(listEntry, update_document_request->query) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*) listEntry->data;
        free (localKeyValue->key);
        free (localKeyValue->value);
    }
    list_free(update_document_request->query);
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
    
    cJSON *doc = cJSON_AddObjectToObject(item, "doc");
    if(doc == NULL) {
        goto fail; //primitive map container
    }
    cJSON *localMapObject = doc;
    listEntry_t *docListEntry;
    if (update_document_request->doc) {
    list_ForEach(docListEntry, update_document_request->doc) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*)docListEntry->data;
    }
    }


    // update_document_request->id
    if(update_document_request->id) { 
    if(cJSON_AddNumberToObject(item, "id", update_document_request->id) == NULL) {
    goto fail; //Numeric
    }
     } 


    // update_document_request->query
    if(update_document_request->query) { 
    cJSON *query = cJSON_AddObjectToObject(item, "query");
    if(query == NULL) {
        goto fail; //primitive map container
    }
    cJSON *localMapObject = query;
    listEntry_t *queryListEntry;
    if (update_document_request->query) {
    list_ForEach(queryListEntry, update_document_request->query) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*)queryListEntry->data;
    }
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

    list_t *docList;
    
    cJSON *doc_local_map;
    if(!cJSON_IsObject(doc)) {
        goto end;//primitive map container
    }
    docList = list_create();
    keyValuePair_t *localMapKeyPair;
    cJSON_ArrayForEach(doc_local_map, doc)
    {
		cJSON *localMapObject = doc_local_map;
        list_addElement(docList , localMapKeyPair);
    }

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
    list_t *queryList;
    if (query) { 
    cJSON *query_local_map;
    if(!cJSON_IsObject(query)) {
        goto end;//primitive map container
    }
    queryList = list_create();
    keyValuePair_t *localMapKeyPair;
    cJSON_ArrayForEach(query_local_map, query)
    {
		cJSON *localMapObject = query_local_map;
        list_addElement(queryList , localMapKeyPair);
    }
    }


    update_document_request_local_var = update_document_request_create (
        strdup(index->valuestring),
        docList,
        id ? id->valuedouble : 0,
        query ? queryList : NULL
        );

    return update_document_request_local_var;
end:
    return NULL;

}
