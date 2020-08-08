#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include "SearchAPI.h"

#define MAX_NUMBER_LENGTH 16
#define MAX_BUFFER_LENGTH 4096
#define intToStr(dst, src) \
    do {\
    char dst[256];\
    snprintf(dst, 256, "%ld", (long int)(src));\
}while(0)


// Perform reverse search on a percolate index
//
// Performs a percolate search. <br/> This method must be used only on percolate indexes. <br/> Expects two paramenters: the index name and an object with array of documents to be tested. <br/> An example of the documents object: <br/> ``` {\"query\":{\"percolate\":{\"document\":{\"content\":\"sample content\"}}}} ``` <br/> Responds with an object with matched stored queries: <br/> ``` {'timed_out':false,'hits':{'total':2,'max_score':1,'hits':[{'_index':'idx_pq_1','_type':'doc','_id':'2','_score':'1','_source':{'query':{'match':{'title':'some'},}}},{'_index':'idx_pq_1','_type':'doc','_id':'5','_score':'1','_source':{'query':{'ql':'some | none'}}}]}} ``` 
//
search_response_t*
SearchAPI_percolate(apiClient_t *apiClient, char * index , percolate_request_t * percolate_request )
{
    list_t    *localVarQueryParameters = NULL;
    list_t    *localVarHeaderParameters = NULL;
    list_t    *localVarFormParameters = NULL;
    list_t *localVarHeaderType = list_create();
    list_t *localVarContentType = list_create();
    char      *localVarBodyParameters = NULL;

    // create the path
    long sizeOfPath = strlen("/json/pq/{index}/search")+1;
    char *localVarPath = malloc(sizeOfPath);
    snprintf(localVarPath, sizeOfPath, "/json/pq/{index}/search");


    // Path Params
    long sizeOfPathParams_index = strlen(index)+3 + strlen("{ index }");
    if(index == NULL) {
        goto end;
    }
    char* localVarToReplace_index = malloc(sizeOfPathParams_index);
    sprintf(localVarToReplace_index, "{%s}", "index");

    localVarPath = strReplace(localVarPath, localVarToReplace_index, index);



    // Body Param
    cJSON *localVarSingleItemJSON_percolate_request;
    if (percolate_request != NULL)
    {
        //string
        localVarSingleItemJSON_percolate_request = percolate_request_convertToJSON(percolate_request);
        localVarBodyParameters = cJSON_Print(localVarSingleItemJSON_percolate_request);
    }
    list_addElement(localVarHeaderType,"application/json"); //produces
    list_addElement(localVarContentType,"application/json"); //consumes
    apiClient_invoke(apiClient,
                    localVarPath,
                    localVarQueryParameters,
                    localVarHeaderParameters,
                    localVarFormParameters,
                    localVarHeaderType,
                    localVarContentType,
                    localVarBodyParameters,
                    "POST");

    if (apiClient->response_code == 200) {
        printf("%s\n","items found");
    }
    if (apiClient->response_code == 0) {
        printf("%s\n","error");
    }
    //nonprimitive not container
    cJSON *SearchAPIlocalVarJSON = cJSON_Parse(apiClient->dataReceived);
    search_response_t *elementToReturn = search_response_parseFromJSON(SearchAPIlocalVarJSON);
    cJSON_Delete(SearchAPIlocalVarJSON);
    if(elementToReturn == NULL) {
        // return 0;
    }

    //return type
    if (apiClient->dataReceived) {
        free(apiClient->dataReceived);
        apiClient->dataReceived = NULL;
        apiClient->dataReceivedLen = 0;
    }
    
    
    
    list_free(localVarHeaderType);
    list_free(localVarContentType);
    free(localVarPath);
    free(localVarToReplace_index);
    cJSON_Delete(localVarSingleItemJSON_percolate_request);
    free(localVarBodyParameters);
    return elementToReturn;
end:
    return NULL;

}

// Performs a search
//
// Performs a search. <br/> Expects an object with mandatory properties: <br/> - the index name <br/> - the match query object <br/> Example : <br/> <code> {'index':'movies','query':{'bool':{'must':[{'query_string':' movie'}]}},'script_fields':{'myexpr':{'script':{'inline':'IF(rating>8,1,0)'}}},'sort':[{'myexpr':'desc'},{'_score':'desc'}],'profile':true} </code> <br/> It responds with an object with <br/> - time of execution <br/> - if the query timed out <br/> - an array with hits (matched documents) <br/> - additional, if profiling is enabled, an array with profiling information is attached <br/>  ``` {'took':10,'timed_out':false,'hits':{'total':2,'hits':[{'_id':'1','_score':1,'_source':{'gid':11}},{'_id':'2','_score':1,'_source':{'gid':12}}]}} ``` <br/> For more information about the match query syntax, additional paramaters that can be set to the input and response, please check: https://docs.manticoresearch.com/latest/html/http_reference/json_search.html. 
//
search_response_t*
SearchAPI_search(apiClient_t *apiClient, search_request_t * search_request )
{
    list_t    *localVarQueryParameters = NULL;
    list_t    *localVarHeaderParameters = NULL;
    list_t    *localVarFormParameters = NULL;
    list_t *localVarHeaderType = list_create();
    list_t *localVarContentType = list_create();
    char      *localVarBodyParameters = NULL;

    // create the path
    long sizeOfPath = strlen("/json/search")+1;
    char *localVarPath = malloc(sizeOfPath);
    snprintf(localVarPath, sizeOfPath, "/json/search");




    // Body Param
    cJSON *localVarSingleItemJSON_search_request;
    if (search_request != NULL)
    {
        //string
        localVarSingleItemJSON_search_request = search_request_convertToJSON(search_request);
        localVarBodyParameters = cJSON_Print(localVarSingleItemJSON_search_request);
    }
    list_addElement(localVarHeaderType,"application/json"); //produces
    list_addElement(localVarContentType,"application/json"); //consumes
    apiClient_invoke(apiClient,
                    localVarPath,
                    localVarQueryParameters,
                    localVarHeaderParameters,
                    localVarFormParameters,
                    localVarHeaderType,
                    localVarContentType,
                    localVarBodyParameters,
                    "POST");

    if (apiClient->response_code == 200) {
        printf("%s\n","Ok");
    }
    if (apiClient->response_code == 0) {
        printf("%s\n","error");
    }
    //nonprimitive not container
    cJSON *SearchAPIlocalVarJSON = cJSON_Parse(apiClient->dataReceived);
    search_response_t *elementToReturn = search_response_parseFromJSON(SearchAPIlocalVarJSON);
    cJSON_Delete(SearchAPIlocalVarJSON);
    if(elementToReturn == NULL) {
        // return 0;
    }

    //return type
    if (apiClient->dataReceived) {
        free(apiClient->dataReceived);
        apiClient->dataReceived = NULL;
        apiClient->dataReceivedLen = 0;
    }
    
    
    
    list_free(localVarHeaderType);
    list_free(localVarContentType);
    free(localVarPath);
    cJSON_Delete(localVarSingleItemJSON_search_request);
    free(localVarBodyParameters);
    return elementToReturn;
end:
    return NULL;

}

