#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "search_response_hits.h"



search_response_hits_t *search_response_hits_create(
    int total,
    list_t *hits
    ) {
    search_response_hits_t *search_response_hits_local_var = malloc(sizeof(search_response_hits_t));
    if (!search_response_hits_local_var) {
        return NULL;
    }
    search_response_hits_local_var->total = total;
    search_response_hits_local_var->hits = hits;

    return search_response_hits_local_var;
}


void search_response_hits_free(search_response_hits_t *search_response_hits) {
    if(NULL == search_response_hits){
        return ;
    }
    listEntry_t *listEntry;
    list_ForEach(listEntry, search_response_hits->hits) {
        object_free(listEntry->data);
    }
    list_free(search_response_hits->hits);
    free(search_response_hits);
}

cJSON *search_response_hits_convertToJSON(search_response_hits_t *search_response_hits) {
    cJSON *item = cJSON_CreateObject();

    // search_response_hits->total
    if(search_response_hits->total) { 
    if(cJSON_AddNumberToObject(item, "total", search_response_hits->total) == NULL) {
    goto fail; //Numeric
    }
     } 


    // search_response_hits->hits
    if(search_response_hits->hits) { 
    cJSON *hits = cJSON_AddArrayToObject(item, "hits");
    if(hits == NULL) {
    goto fail; //nonprimitive container
    }

    listEntry_t *hitsListEntry;
    if (search_response_hits->hits) {
    list_ForEach(hitsListEntry, search_response_hits->hits) {
    cJSON *itemLocal = object_convertToJSON(hitsListEntry->data);
    if(itemLocal == NULL) {
    goto fail;
    }
    cJSON_AddItemToArray(hits, itemLocal);
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

search_response_hits_t *search_response_hits_parseFromJSON(cJSON *search_response_hitsJSON){

    search_response_hits_t *search_response_hits_local_var = NULL;

    // search_response_hits->total
    cJSON *total = cJSON_GetObjectItemCaseSensitive(search_response_hitsJSON, "total");
    if (total) { 
    if(!cJSON_IsNumber(total))
    {
    goto end; //Numeric
    }
    }

    // search_response_hits->hits
    cJSON *hits = cJSON_GetObjectItemCaseSensitive(search_response_hitsJSON, "hits");
    list_t *hitsList;
    if (hits) { 
    cJSON *hits_local_nonprimitive;
    if(!cJSON_IsArray(hits)){
        goto end; //nonprimitive container
    }

    hitsList = list_create();

    cJSON_ArrayForEach(hits_local_nonprimitive,hits )
    {
        if(!cJSON_IsObject(hits_local_nonprimitive)){
            goto end;
        }
        object_t *hitsItem = object_parseFromJSON(hits_local_nonprimitive);

        list_addElement(hitsList, hitsItem);
    }
    }


    search_response_hits_local_var = search_response_hits_create (
        total ? total->valuedouble : 0,
        hits ? hitsList : NULL
        );

    return search_response_hits_local_var;
end:
    return NULL;

}
