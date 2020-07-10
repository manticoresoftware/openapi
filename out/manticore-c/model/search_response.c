#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "search_response.h"



search_response_t *search_response_create(
    int took,
    int timed_out,
    search_response_hits_t *hits,
    object_t *profile
    ) {
    search_response_t *search_response_local_var = malloc(sizeof(search_response_t));
    if (!search_response_local_var) {
        return NULL;
    }
    search_response_local_var->took = took;
    search_response_local_var->timed_out = timed_out;
    search_response_local_var->hits = hits;
    search_response_local_var->profile = profile;

    return search_response_local_var;
}


void search_response_free(search_response_t *search_response) {
    if(NULL == search_response){
        return ;
    }
    listEntry_t *listEntry;
    search_response_hits_free(search_response->hits);
    object_free(search_response->profile);
    free(search_response);
}

cJSON *search_response_convertToJSON(search_response_t *search_response) {
    cJSON *item = cJSON_CreateObject();

    // search_response->took
    if(search_response->took) { 
    if(cJSON_AddNumberToObject(item, "took", search_response->took) == NULL) {
    goto fail; //Numeric
    }
     } 


    // search_response->timed_out
    if(search_response->timed_out) { 
    if(cJSON_AddBoolToObject(item, "timed_out", search_response->timed_out) == NULL) {
    goto fail; //Bool
    }
     } 


    // search_response->hits
    if(search_response->hits) { 
    cJSON *hits_local_JSON = search_response_hits_convertToJSON(search_response->hits);
    if(hits_local_JSON == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "hits", hits_local_JSON);
    if(item->child == NULL) {
    goto fail;
    }
     } 


    // search_response->profile
    if(search_response->profile) { 
    cJSON *profile_object = object_convertToJSON(search_response->profile);
    if(profile_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "profile", profile_object);
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

search_response_t *search_response_parseFromJSON(cJSON *search_responseJSON){

    search_response_t *search_response_local_var = NULL;

    // search_response->took
    cJSON *took = cJSON_GetObjectItemCaseSensitive(search_responseJSON, "took");
    if (took) { 
    if(!cJSON_IsNumber(took))
    {
    goto end; //Numeric
    }
    }

    // search_response->timed_out
    cJSON *timed_out = cJSON_GetObjectItemCaseSensitive(search_responseJSON, "timed_out");
    if (timed_out) { 
    if(!cJSON_IsBool(timed_out))
    {
    goto end; //Bool
    }
    }

    // search_response->hits
    cJSON *hits = cJSON_GetObjectItemCaseSensitive(search_responseJSON, "hits");
    search_response_hits_t *hits_local_nonprim = NULL;
    if (hits) { 
    hits_local_nonprim = search_response_hits_parseFromJSON(hits); //nonprimitive
    }

    // search_response->profile
    cJSON *profile = cJSON_GetObjectItemCaseSensitive(search_responseJSON, "profile");
    object_t *profile_local_object = NULL;
    if (profile) { 
    profile_local_object = object_parseFromJSON(profile); //object
    }


    search_response_local_var = search_response_create (
        took ? took->valuedouble : 0,
        timed_out ? timed_out->valueint : 0,
        hits ? hits_local_nonprim : NULL,
        profile ? profile_local_object : NULL
        );

    return search_response_local_var;
end:
    return NULL;

}
