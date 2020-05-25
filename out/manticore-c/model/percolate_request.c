#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "percolate_request.h"



percolate_request_t *percolate_request_create(
    list_t* query
    ) {
    percolate_request_t *percolate_request_local_var = malloc(sizeof(percolate_request_t));
    if (!percolate_request_local_var) {
        return NULL;
    }
    percolate_request_local_var->query = query;

    return percolate_request_local_var;
}


void percolate_request_free(percolate_request_t *percolate_request) {
    if(NULL == percolate_request){
        return ;
    }
    listEntry_t *listEntry;
    list_ForEach(listEntry, percolate_request->query) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*) listEntry->data;
        free (localKeyValue->key);
        free (localKeyValue->value);
    }
    list_free(percolate_request->query);
    free(percolate_request);
}

cJSON *percolate_request_convertToJSON(percolate_request_t *percolate_request) {
    cJSON *item = cJSON_CreateObject();

    // percolate_request->query
    if (!percolate_request->query) {
        goto fail;
    }
    
    cJSON *query = cJSON_AddObjectToObject(item, "query");
    if(query == NULL) {
        goto fail; //primitive map container
    }
    cJSON *localMapObject = query;
    listEntry_t *queryListEntry;
    if (percolate_request->query) {
    list_ForEach(queryListEntry, percolate_request->query) {
        keyValuePair_t *localKeyValue = (keyValuePair_t*)queryListEntry->data;
    }
    }

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

percolate_request_t *percolate_request_parseFromJSON(cJSON *percolate_requestJSON){

    percolate_request_t *percolate_request_local_var = NULL;

    // percolate_request->query
    cJSON *query = cJSON_GetObjectItemCaseSensitive(percolate_requestJSON, "query");
    if (!query) {
        goto end;
    }

    list_t *queryList;
    
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


    percolate_request_local_var = percolate_request_create (
        queryList
        );

    return percolate_request_local_var;
end:
    return NULL;

}
