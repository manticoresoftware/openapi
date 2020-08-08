#include <stdlib.h>
#include <stdio.h>
#include "../include/apiClient.h"
#include "../include/list.h"
#include "../external/cJSON.h"
#include "../include/keyValuePair.h"
#include "../include/binary.h"
#include "../model/error_response.h"
#include "../model/percolate_request.h"
#include "../model/search_request.h"
#include "../model/search_response.h"


// Perform reverse search on a percolate index
//
// Performs a percolate search. <br/> This method must be used only on percolate indexes. <br/> Expects two paramenters: the index name and an object with array of documents to be tested. <br/> An example of the documents object: <br/> ``` {\"query\":{\"percolate\":{\"document\":{\"content\":\"sample content\"}}}} ``` <br/> Responds with an object with matched stored queries: <br/> ``` {'timed_out':false,'hits':{'total':2,'max_score':1,'hits':[{'_index':'idx_pq_1','_type':'doc','_id':'2','_score':'1','_source':{'query':{'match':{'title':'some'},}}},{'_index':'idx_pq_1','_type':'doc','_id':'5','_score':'1','_source':{'query':{'ql':'some | none'}}}]}} ``` 
//
search_response_t*
SearchAPI_percolate(apiClient_t *apiClient, char * index , percolate_request_t * percolate_request );


// Performs a search
//
// Performs a search. <br/> Expects an object with mandatory properties: <br/> - the index name <br/> - the match query object <br/> Example : <br/> <code> {'index':'movies','query':{'bool':{'must':[{'query_string':' movie'}]}},'script_fields':{'myexpr':{'script':{'inline':'IF(rating>8,1,0)'}}},'sort':[{'myexpr':'desc'},{'_score':'desc'}],'profile':true} </code> <br/> It responds with an object with <br/> - time of execution <br/> - if the query timed out <br/> - an array with hits (matched documents) <br/> - additional, if profiling is enabled, an array with profiling information is attached <br/>  ``` {'took':10,'timed_out':false,'hits':{'total':2,'hits':[{'_id':'1','_score':1,'_source':{'gid':11}},{'_id':'2','_score':1,'_source':{'gid':12}}]}} ``` <br/> For more information about the match query syntax, additional paramaters that can be set to the input and response, please check: https://docs.manticoresearch.com/latest/html/http_reference/json_search.html. 
//
search_response_t*
SearchAPI_search(apiClient_t *apiClient, search_request_t * search_request );


