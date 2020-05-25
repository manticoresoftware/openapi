#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include "SearchAPI.h"


#define MAX_BUFFER_LENGTH 4096
#define intToStr(dst, src) \
    do {\
    char dst[256];\
    snprintf(dst, 256, "%ld", (long int)(src));\
}while(0)


// Perform reverse search on a percolate index
//
search_response_t*
SearchAPI_percolate(apiClient_t *apiClient, percolate_request_t * percolate_request )
{
    list_t    *localVarQueryParameters = NULL;
    list_t    *localVarHeaderParameters = NULL;
    list_t    *localVarFormParameters = NULL;
    list_t *localVarHeaderType = list_create();
    list_t *localVarContentType = list_create();
    char      *localVarBodyParameters = NULL;

    // create the path
    long sizeOfPath = strlen("/json/pq/search")+1;
    char *localVarPath = malloc(sizeOfPath);
    snprintf(localVarPath, sizeOfPath, "/json/pq/search");




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
    cJSON_Delete(localVarSingleItemJSON_percolate_request);
    free(localVarBodyParameters);
    return elementToReturn;
end:
    return NULL;

}

// Performs a search
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

