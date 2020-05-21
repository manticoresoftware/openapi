#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include "UtilsAPI.h"


#define MAX_BUFFER_LENGTH 4096
#define intToStr(dst, src) \
    do {\
    char dst[256];\
    snprintf(dst, 256, "%ld", (long int)(src));\
}while(0)


// Perform SQL requests
//
object_t*
UtilsAPI_sql(apiClient_t *apiClient, char * query , char * mode )
{
    list_t    *localVarQueryParameters = NULL;
    list_t    *localVarHeaderParameters = NULL;
    list_t    *localVarFormParameters = list_create();
    list_t *localVarHeaderType = list_create();
    list_t *localVarContentType = list_create();
    char      *localVarBodyParameters = NULL;

    // create the path
    long sizeOfPath = strlen("/sql")+1;
    char *localVarPath = malloc(sizeOfPath);
    snprintf(localVarPath, sizeOfPath, "/sql");




    // form parameters
    char *keyForm_query;
    char * valueForm_query;
    keyValuePair_t *keyPairForm_query = 0;
    if (query != NULL)
    {
        keyForm_query = strdup("query");
        valueForm_query = strdup((query));
        keyPairForm_query = keyValuePair_create(keyForm_query,valueForm_query);
        list_addElement(localVarFormParameters,keyPairForm_query);
    }

    // form parameters
    char *keyForm_mode;
    char * valueForm_mode;
    keyValuePair_t *keyPairForm_mode = 0;
    if (mode != NULL)
    {
        keyForm_mode = strdup("mode");
        valueForm_mode = strdup((mode));
        keyPairForm_mode = keyValuePair_create(keyForm_mode,valueForm_mode);
        list_addElement(localVarFormParameters,keyPairForm_mode);
    }
    list_addElement(localVarHeaderType,"application/json"); //produces
    list_addElement(localVarContentType,"application/x-www-form-urlencoded"); //consumes
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
        printf("%s\n","item updated");
    }
    if (apiClient->response_code == 0) {
        printf("%s\n","error");
    }
    //nonprimitive not container
    cJSON *UtilsAPIlocalVarJSON = cJSON_Parse(apiClient->dataReceived);
    object_t *elementToReturn = object_parseFromJSON(UtilsAPIlocalVarJSON);
    cJSON_Delete(UtilsAPIlocalVarJSON);
    if(elementToReturn == NULL) {
        // return 0;
    }

    //return type
    if (apiClient->dataReceived) {
        free(apiClient->dataReceived);
        apiClient->dataReceived = NULL;
        apiClient->dataReceivedLen = 0;
    }
    
    
    list_free(localVarFormParameters);
    list_free(localVarHeaderType);
    list_free(localVarContentType);
    free(localVarPath);
    free(keyForm_query);
    free(valueForm_query);
    free(keyPairForm_query);
    free(keyForm_mode);
    free(valueForm_mode);
    free(keyPairForm_mode);
    return elementToReturn;
end:
    return NULL;

}

