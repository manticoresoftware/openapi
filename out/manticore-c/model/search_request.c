#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "search_request.h"



search_request_t *search_request_create(
    char *index,
    object_t *query,
    int limit,
    int offset,
    int max_matches,
    list_t *sort,
    object_t *script_fields,
    object_t *highlight,
    list_t *_source,
    int profile
    ) {
    search_request_t *search_request_local_var = malloc(sizeof(search_request_t));
    if (!search_request_local_var) {
        return NULL;
    }
    search_request_local_var->index = index;
    search_request_local_var->query = query;
    search_request_local_var->limit = limit;
    search_request_local_var->offset = offset;
    search_request_local_var->max_matches = max_matches;
    search_request_local_var->sort = sort;
    search_request_local_var->script_fields = script_fields;
    search_request_local_var->highlight = highlight;
    search_request_local_var->_source = _source;
    search_request_local_var->profile = profile;

    return search_request_local_var;
}


void search_request_free(search_request_t *search_request) {
    if(NULL == search_request){
        return ;
    }
    listEntry_t *listEntry;
    free(search_request->index);
    object_free(search_request->query);
    list_ForEach(listEntry, search_request->sort) {
        object_free(listEntry->data);
    }
    list_free(search_request->sort);
    object_free(search_request->script_fields);
    object_free(search_request->highlight);
    list_ForEach(listEntry, search_request->_source) {
        free(listEntry->data);
    }
    list_free(search_request->_source);
    free(search_request);
}

cJSON *search_request_convertToJSON(search_request_t *search_request) {
    cJSON *item = cJSON_CreateObject();

    // search_request->index
    if (!search_request->index) {
        goto fail;
    }
    
    if(cJSON_AddStringToObject(item, "index", search_request->index) == NULL) {
    goto fail; //String
    }


    // search_request->query
    if (!search_request->query) {
        goto fail;
    }
    
    cJSON *query_object = object_convertToJSON(search_request->query);
    if(query_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "query", query_object);
    if(item->child == NULL) {
    goto fail;
    }


    // search_request->limit
    if(search_request->limit) { 
    if(cJSON_AddNumberToObject(item, "limit", search_request->limit) == NULL) {
    goto fail; //Numeric
    }
     } 


    // search_request->offset
    if(search_request->offset) { 
    if(cJSON_AddNumberToObject(item, "offset", search_request->offset) == NULL) {
    goto fail; //Numeric
    }
     } 


    // search_request->max_matches
    if(search_request->max_matches) { 
    if(cJSON_AddNumberToObject(item, "max_matches", search_request->max_matches) == NULL) {
    goto fail; //Numeric
    }
     } 


    // search_request->sort
    if(search_request->sort) { 
    cJSON *sort = cJSON_AddArrayToObject(item, "sort");
    if(sort == NULL) {
    goto fail; //nonprimitive container
    }

    listEntry_t *sortListEntry;
    if (search_request->sort) {
    list_ForEach(sortListEntry, search_request->sort) {
    cJSON *itemLocal = object_convertToJSON(sortListEntry->data);
    if(itemLocal == NULL) {
    goto fail;
    }
    cJSON_AddItemToArray(sort, itemLocal);
    }
    }
     } 


    // search_request->script_fields
    if(search_request->script_fields) { 
    cJSON *script_fields_object = object_convertToJSON(search_request->script_fields);
    if(script_fields_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "script_fields", script_fields_object);
    if(item->child == NULL) {
    goto fail;
    }
     } 


    // search_request->highlight
    if(search_request->highlight) { 
    cJSON *highlight_object = object_convertToJSON(search_request->highlight);
    if(highlight_object == NULL) {
    goto fail; //model
    }
    cJSON_AddItemToObject(item, "highlight", highlight_object);
    if(item->child == NULL) {
    goto fail;
    }
     } 


    // search_request->_source
    if(search_request->_source) { 
    cJSON *_source = cJSON_AddArrayToObject(item, "_source");
    if(_source == NULL) {
        goto fail; //primitive container
    }

    listEntry_t *_sourceListEntry;
    list_ForEach(_sourceListEntry, search_request->_source) {
    if(cJSON_AddStringToObject(_source, "", (char*)_sourceListEntry->data) == NULL)
    {
        goto fail;
    }
    }
     } 


    // search_request->profile
    if(search_request->profile) { 
    if(cJSON_AddBoolToObject(item, "profile", search_request->profile) == NULL) {
    goto fail; //Bool
    }
     } 

    return item;
fail:
    if (item) {
        cJSON_Delete(item);
    }
    return NULL;
}

search_request_t *search_request_parseFromJSON(cJSON *search_requestJSON){

    search_request_t *search_request_local_var = NULL;

    // search_request->index
    cJSON *index = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "index");
    if (!index) {
        goto end;
    }

    
    if(!cJSON_IsString(index))
    {
    goto end; //String
    }

    // search_request->query
    cJSON *query = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "query");
    if (!query) {
        goto end;
    }

    object_t *query_local_object = NULL;
    
    query_local_object = object_parseFromJSON(query); //object

    // search_request->limit
    cJSON *limit = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "limit");
    if (limit) { 
    if(!cJSON_IsNumber(limit))
    {
    goto end; //Numeric
    }
    }

    // search_request->offset
    cJSON *offset = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "offset");
    if (offset) { 
    if(!cJSON_IsNumber(offset))
    {
    goto end; //Numeric
    }
    }

    // search_request->max_matches
    cJSON *max_matches = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "max_matches");
    if (max_matches) { 
    if(!cJSON_IsNumber(max_matches))
    {
    goto end; //Numeric
    }
    }

    // search_request->sort
    cJSON *sort = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "sort");
    list_t *sortList;
    if (sort) { 
    cJSON *sort_local_nonprimitive;
    if(!cJSON_IsArray(sort)){
        goto end; //nonprimitive container
    }

    sortList = list_create();

    cJSON_ArrayForEach(sort_local_nonprimitive,sort )
    {
        if(!cJSON_IsObject(sort_local_nonprimitive)){
            goto end;
        }
        object_t *sortItem = object_parseFromJSON(sort_local_nonprimitive);

        list_addElement(sortList, sortItem);
    }
    }

    // search_request->script_fields
    cJSON *script_fields = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "script_fields");
    object_t *script_fields_local_object = NULL;
    if (script_fields) { 
    script_fields_local_object = object_parseFromJSON(script_fields); //object
    }

    // search_request->highlight
    cJSON *highlight = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "highlight");
    object_t *highlight_local_object = NULL;
    if (highlight) { 
    highlight_local_object = object_parseFromJSON(highlight); //object
    }

    // search_request->_source
    cJSON *_source = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "_source");
    list_t *_sourceList;
    if (_source) { 
    cJSON *_source_local;
    if(!cJSON_IsArray(_source)) {
        goto end;//primitive container
    }
    _sourceList = list_create();

    cJSON_ArrayForEach(_source_local, _source)
    {
        if(!cJSON_IsString(_source_local))
        {
            goto end;
        }
        list_addElement(_sourceList , strdup(_source_local->valuestring));
    }
    }

    // search_request->profile
    cJSON *profile = cJSON_GetObjectItemCaseSensitive(search_requestJSON, "profile");
    if (profile) { 
    if(!cJSON_IsBool(profile))
    {
    goto end; //Bool
    }
    }


    search_request_local_var = search_request_create (
        strdup(index->valuestring),
        query_local_object,
        limit ? limit->valuedouble : 0,
        offset ? offset->valuedouble : 0,
        max_matches ? max_matches->valuedouble : 0,
        sort ? sortList : NULL,
        script_fields ? script_fields_local_object : NULL,
        highlight ? highlight_local_object : NULL,
        _source ? _sourceList : NULL,
        profile ? profile->valueint : 0
        );

    return search_request_local_var;
end:
    return NULL;

}
