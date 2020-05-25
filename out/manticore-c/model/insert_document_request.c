#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "insert_document_request.h"



insert_document_request_t *insert_document_request_create(
    char *index,
    long id,
    list_t* doc
    ) {
    insert_document_request_t *insert_document_request_local_var = malloc(sizeof(insert_document_request_t));
    if (!insert_document_request_local_var) {
        return NULL;
    }
    insert_document_request_local_var->index = index;
    insert_document_request_local_var->id = id;
    insert_document_request_local_var->doc = doc;

    return insert_document_request_local_var;
}


void insert_document_request_free(insert_document_request_t *insert_document_request) {
    if(NULL == insert_document_request){
        return ;
    }
    listEntry_t *listEntry;
    free(insert_document_request->index);
    list_ForEach(listEntry, insert_document_request->doc) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*) listEntry->data;
        free (localKeyValue->key);
        free (localKeyValue->value);
    }
    list_free(insert_document_request->doc);
    free(insert_document_request);
}

cJSON *insert_document_request_convertToJSON(insert_document_request_t *insert_document_request) {
    cJSON *item = cJSON_CreateObject();

    // insert_document_request->index
    if (!insert_document_request->index) {
        goto fail;
    }
    
    if(cJSON_AddStringToObject(item, "index", insert_document_request->index) == NULL) {
    goto fail; //String
    }


    // insert_document_request->id
    if(insert_document_request->id) { 
    if(cJSON_AddNumberToObject(item, "id", insert_document_request->id) == NULL) {
    goto fail; //Numeric
    }
     } 


    // insert_document_request->doc
    if (!insert_document_request->doc) {
        goto fail;
    }
    
    cJSON *doc = cJSON_AddObjectToObject(item, "doc");
    if(doc == NULL) {
        goto fail; //primitive map container
    }
    cJSON *localMapObject = doc;
    listEntry_t *docListEntry;
    if (insert_document_request->doc) {
    list_ForEach(docListEntry, insert_document_request->doc) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*)docListEntry->data;
    }
    }

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

insert_document_request_t *insert_document_request_parseFromJSON(cJSON *insert_document_requestJSON){

    insert_document_request_t *insert_document_request_local_var = NULL;

    // insert_document_request->index
    cJSON *index = cJSON_GetObjectItemCaseSensitive(insert_document_requestJSON, "index");
    if (!index) {
        goto end;
    }

    
    if(!cJSON_IsString(index))
    {
    goto end; //String
    }

    // insert_document_request->id
    cJSON *id = cJSON_GetObjectItemCaseSensitive(insert_document_requestJSON, "id");
    if (id) { 
    if(!cJSON_IsNumber(id))
    {
    goto end; //Numeric
    }
    }

    // insert_document_request->doc
    cJSON *doc = cJSON_GetObjectItemCaseSensitive(insert_document_requestJSON, "doc");
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


    insert_document_request_local_var = insert_document_request_create (
        strdup(index->valuestring),
        id ? id->valuedouble : 0,
        docList
        );

    return insert_document_request_local_var;
end:
    return NULL;

}
