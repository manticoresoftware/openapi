import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { AggComposite } from '../models/AggComposite';
import { AggCompositeSource } from '../models/AggCompositeSource';
import { AggCompositeTerm } from '../models/AggCompositeTerm';
import { AggTerms } from '../models/AggTerms';
import { Aggregation } from '../models/Aggregation';
import { BoolFilter } from '../models/BoolFilter';
import { BulkResponse } from '../models/BulkResponse';
import { DeleteDocumentRequest } from '../models/DeleteDocumentRequest';
import { DeleteResponse } from '../models/DeleteResponse';
import { ErrorResponse } from '../models/ErrorResponse';
import { FulltextFilter } from '../models/FulltextFilter';
import { GeoDistance } from '../models/GeoDistance';
import { GeoDistanceLocationAnchor } from '../models/GeoDistanceLocationAnchor';
import { Highlight } from '../models/Highlight';
import { HighlightFieldOption } from '../models/HighlightFieldOption';
import { InsertDocumentRequest } from '../models/InsertDocumentRequest';
import { Join } from '../models/Join';
import { JoinCond } from '../models/JoinCond';
import { JoinOn } from '../models/JoinOn';
import { KnnQuery } from '../models/KnnQuery';
import { Match } from '../models/Match';
import { MatchAll } from '../models/MatchAll';
import { PercolateRequest } from '../models/PercolateRequest';
import { PercolateRequestQuery } from '../models/PercolateRequestQuery';
import { QueryFilter } from '../models/QueryFilter';
import { Range } from '../models/Range';
import { ReplaceDocumentRequest } from '../models/ReplaceDocumentRequest';
import { ResponseError } from '../models/ResponseError';
import { ResponseErrorDetails } from '../models/ResponseErrorDetails';
import { SearchQuery } from '../models/SearchQuery';
import { SearchRequest } from '../models/SearchRequest';
import { SearchResponse } from '../models/SearchResponse';
import { SearchResponseHits } from '../models/SearchResponseHits';
import { SourceRules } from '../models/SourceRules';
import { SqlResponse } from '../models/SqlResponse';
import { SuccessResponse } from '../models/SuccessResponse';
import { UpdateDocumentRequest } from '../models/UpdateDocumentRequest';
import { UpdateResponse } from '../models/UpdateResponse';
import { ObservableIndexApi } from './ObservableAPI';

import { IndexApiRequestFactory, IndexApiResponseProcessor} from "../apis/IndexApi";
export class PromiseIndexApi {
    private api: ObservableIndexApi

    public constructor(
        configuration: Configuration,
        requestFactory?: IndexApiRequestFactory,
        responseProcessor?: IndexApiResponseProcessor
    ) {
        this.api = new ObservableIndexApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * Sends multiple operatons like inserts, updates, replaces or deletes.  For each operation it\'s object must have same format as in their dedicated method.  The method expects a raw string as the batch in NDJSON.  Each operation object needs to be serialized to   JSON and separated by endline (\\n).      An example of raw input:      ```   {\"insert\": {\"index\": \"movies\", \"doc\": {\"plot\": \"A secret team goes to North Pole\", \"rating\": 9.5, \"language\": [2, 3], \"title\": \"This is an older movie\", \"lon\": 51.99, \"meta\": {\"keywords\":[\"travel\",\"ice\"],\"genre\":[\"adventure\"]}, \"year\": 1950, \"lat\": 60.4, \"advise\": \"PG-13\"}}}   \\n   {\"delete\": {\"index\": \"movies\",\"id\":700}}   ```      Responds with an object telling whenever any errors occured and an array with status for each operation:      ```   {     \'items\':     [       {         \'update\':{\'_index\':\'products\',\'_id\':1,\'result\':\'updated\'}       },       {         \'update\':{\'_index\':\'products\',\'_id\':2,\'result\':\'updated\'}       }     ],     \'errors\':false   }   ``` 
     * Bulk index operations
     * @param body
     */
    public bulkWithHttpInfo(body: string, _options?: Configuration): Promise<HttpInfo<BulkResponse>> {
        const result = this.api.bulkWithHttpInfo(body, _options);
        return result.toPromise();
    }

    /**
     * Sends multiple operatons like inserts, updates, replaces or deletes.  For each operation it\'s object must have same format as in their dedicated method.  The method expects a raw string as the batch in NDJSON.  Each operation object needs to be serialized to   JSON and separated by endline (\\n).      An example of raw input:      ```   {\"insert\": {\"index\": \"movies\", \"doc\": {\"plot\": \"A secret team goes to North Pole\", \"rating\": 9.5, \"language\": [2, 3], \"title\": \"This is an older movie\", \"lon\": 51.99, \"meta\": {\"keywords\":[\"travel\",\"ice\"],\"genre\":[\"adventure\"]}, \"year\": 1950, \"lat\": 60.4, \"advise\": \"PG-13\"}}}   \\n   {\"delete\": {\"index\": \"movies\",\"id\":700}}   ```      Responds with an object telling whenever any errors occured and an array with status for each operation:      ```   {     \'items\':     [       {         \'update\':{\'_index\':\'products\',\'_id\':1,\'result\':\'updated\'}       },       {         \'update\':{\'_index\':\'products\',\'_id\':2,\'result\':\'updated\'}       }     ],     \'errors\':false   }   ``` 
     * Bulk index operations
     * @param body
     */
    public bulk(body: string, _options?: Configuration): Promise<BulkResponse> {
        const result = this.api.bulk(body, _options);
        return result.toPromise();
    }

    /**
     * Delete one or several documents. The method has 2 ways of deleting: either by id, in case only one document is deleted or by using a  match query, in which case multiple documents can be delete . Example of input to delete by id:    ```   {\'index\':\'movies\',\'id\':100}   ```  Example of input to delete using a query:    ```   {     \'index\':\'movies\',     \'query\':     {       \'bool\':       {         \'must\':         [           {\'query_string\':\'new movie\'}         ]       }     }   }   ```  The match query has same syntax as in for searching. Responds with an object telling how many documents got deleted:     ```   {\'_index\':\'products\',\'updated\':1}   ``` 
     * Delete a document in an index
     * @param deleteDocumentRequest
     */
    public deleteWithHttpInfo(deleteDocumentRequest: DeleteDocumentRequest, _options?: Configuration): Promise<HttpInfo<DeleteResponse>> {
        const result = this.api.deleteWithHttpInfo(deleteDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Delete one or several documents. The method has 2 ways of deleting: either by id, in case only one document is deleted or by using a  match query, in which case multiple documents can be delete . Example of input to delete by id:    ```   {\'index\':\'movies\',\'id\':100}   ```  Example of input to delete using a query:    ```   {     \'index\':\'movies\',     \'query\':     {       \'bool\':       {         \'must\':         [           {\'query_string\':\'new movie\'}         ]       }     }   }   ```  The match query has same syntax as in for searching. Responds with an object telling how many documents got deleted:     ```   {\'_index\':\'products\',\'updated\':1}   ``` 
     * Delete a document in an index
     * @param deleteDocumentRequest
     */
    public delete(deleteDocumentRequest: DeleteDocumentRequest, _options?: Configuration): Promise<DeleteResponse> {
        const result = this.api.delete(deleteDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Insert a document.  Expects an object like:     ```   {     \'index\':\'movies\',     \'id\':701,     \'doc\':     {       \'title\':\'This is an old movie\',       \'plot\':\'A secret team goes to North Pole\',       \'year\':1950,       \'rating\':9.5,       \'lat\':60.4,       \'lon\':51.99,       \'advise\':\'PG-13\',       \'meta\':\'{\"keywords\":{\"travel\",\"ice\"},\"genre\":{\"adventure\"}}\',       \'language\':[2,3]     }   }   ```   The document id can also be missing, in which case an autogenerated one will be used:             ```   {     \'index\':\'movies\',     \'doc\':     {       \'title\':\'This is a new movie\',       \'plot\':\'A secret team goes to North Pole\',       \'year\':2020,       \'rating\':9.5,       \'lat\':60.4,       \'lon\':51.99,       \'advise\':\'PG-13\',       \'meta\':\'{\"keywords\":{\"travel\",\"ice\"},\"genre\":{\"adventure\"}}\',       \'language\':[2,3]     }   }   ```   It responds with an object in format:      ```   {\'_index\':\'products\',\'_id\':701,\'created\':true,\'result\':\'created\',\'status\':201}   ``` 
     * Create a new document in an index
     * @param insertDocumentRequest
     */
    public insertWithHttpInfo(insertDocumentRequest: InsertDocumentRequest, _options?: Configuration): Promise<HttpInfo<SuccessResponse>> {
        const result = this.api.insertWithHttpInfo(insertDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Insert a document.  Expects an object like:     ```   {     \'index\':\'movies\',     \'id\':701,     \'doc\':     {       \'title\':\'This is an old movie\',       \'plot\':\'A secret team goes to North Pole\',       \'year\':1950,       \'rating\':9.5,       \'lat\':60.4,       \'lon\':51.99,       \'advise\':\'PG-13\',       \'meta\':\'{\"keywords\":{\"travel\",\"ice\"},\"genre\":{\"adventure\"}}\',       \'language\':[2,3]     }   }   ```   The document id can also be missing, in which case an autogenerated one will be used:             ```   {     \'index\':\'movies\',     \'doc\':     {       \'title\':\'This is a new movie\',       \'plot\':\'A secret team goes to North Pole\',       \'year\':2020,       \'rating\':9.5,       \'lat\':60.4,       \'lon\':51.99,       \'advise\':\'PG-13\',       \'meta\':\'{\"keywords\":{\"travel\",\"ice\"},\"genre\":{\"adventure\"}}\',       \'language\':[2,3]     }   }   ```   It responds with an object in format:      ```   {\'_index\':\'products\',\'_id\':701,\'created\':true,\'result\':\'created\',\'status\':201}   ``` 
     * Create a new document in an index
     * @param insertDocumentRequest
     */
    public insert(insertDocumentRequest: InsertDocumentRequest, _options?: Configuration): Promise<SuccessResponse> {
        const result = this.api.insert(insertDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Partially replaces a document with given id in an index Responds with an object of the following format:     ```   {\'_index\':\'products\',\'updated\':1}   ``` 
     * Partially replaces a document in an index
     * @param index Name of the percolate index
     * @param id Id of the document to replace
     * @param replaceDocumentRequest
     */
    public partialReplaceWithHttpInfo(index: string, id: number, replaceDocumentRequest: ReplaceDocumentRequest, _options?: Configuration): Promise<HttpInfo<UpdateResponse>> {
        const result = this.api.partialReplaceWithHttpInfo(index, id, replaceDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Partially replaces a document with given id in an index Responds with an object of the following format:     ```   {\'_index\':\'products\',\'updated\':1}   ``` 
     * Partially replaces a document in an index
     * @param index Name of the percolate index
     * @param id Id of the document to replace
     * @param replaceDocumentRequest
     */
    public partialReplace(index: string, id: number, replaceDocumentRequest: ReplaceDocumentRequest, _options?: Configuration): Promise<UpdateResponse> {
        const result = this.api.partialReplace(index, id, replaceDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Replace an existing document. Input has same format as `insert` operation. <br/> Responds with an object in format: <br/>    ```   {\'_index\':\'products\',\'_id\':1,\'created\':false,\'result\':\'updated\',\'status\':200}   ``` 
     * Replace new document in an index
     * @param insertDocumentRequest
     */
    public replaceWithHttpInfo(insertDocumentRequest: InsertDocumentRequest, _options?: Configuration): Promise<HttpInfo<SuccessResponse>> {
        const result = this.api.replaceWithHttpInfo(insertDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Replace an existing document. Input has same format as `insert` operation. <br/> Responds with an object in format: <br/>    ```   {\'_index\':\'products\',\'_id\':1,\'created\':false,\'result\':\'updated\',\'status\':200}   ``` 
     * Replace new document in an index
     * @param insertDocumentRequest
     */
    public replace(insertDocumentRequest: InsertDocumentRequest, _options?: Configuration): Promise<SuccessResponse> {
        const result = this.api.replace(insertDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Update one or several documents. The update can be made by passing the id or by using a match query in case multiple documents can be updated.  For example update a document using document id:    ```   {\'index\':\'movies\',\'doc\':{\'rating\':9.49},\'id\':100}   ```  And update by using a match query:    ```   {     \'index\':\'movies\',     \'doc\':{\'rating\':9.49},     \'query\':     {       \'bool\':       {         \'must\':         [           {\'query_string\':\'new movie\'}         ]       }     }   }   ```   The match query has same syntax as for searching. Responds with an object that tells how many documents where updated in format:     ```   {\'_index\':\'products\',\'updated\':1}   ``` 
     * Update a document in an index
     * @param updateDocumentRequest
     */
    public updateWithHttpInfo(updateDocumentRequest: UpdateDocumentRequest, _options?: Configuration): Promise<HttpInfo<UpdateResponse>> {
        const result = this.api.updateWithHttpInfo(updateDocumentRequest, _options);
        return result.toPromise();
    }

    /**
     * Update one or several documents. The update can be made by passing the id or by using a match query in case multiple documents can be updated.  For example update a document using document id:    ```   {\'index\':\'movies\',\'doc\':{\'rating\':9.49},\'id\':100}   ```  And update by using a match query:    ```   {     \'index\':\'movies\',     \'doc\':{\'rating\':9.49},     \'query\':     {       \'bool\':       {         \'must\':         [           {\'query_string\':\'new movie\'}         ]       }     }   }   ```   The match query has same syntax as for searching. Responds with an object that tells how many documents where updated in format:     ```   {\'_index\':\'products\',\'updated\':1}   ``` 
     * Update a document in an index
     * @param updateDocumentRequest
     */
    public update(updateDocumentRequest: UpdateDocumentRequest, _options?: Configuration): Promise<UpdateResponse> {
        const result = this.api.update(updateDocumentRequest, _options);
        return result.toPromise();
    }


}



import { ObservableSearchApi } from './ObservableAPI';

import { SearchApiRequestFactory, SearchApiResponseProcessor} from "../apis/SearchApi";
export class PromiseSearchApi {
    private api: ObservableSearchApi

    public constructor(
        configuration: Configuration,
        requestFactory?: SearchApiRequestFactory,
        responseProcessor?: SearchApiResponseProcessor
    ) {
        this.api = new ObservableSearchApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * Performs a percolate search.  This method must be used only on percolate indexes.  Expects two parameters: the index name and an object with array of documents to be tested. An example of the documents object:    ```   {     \"query\":     {       \"percolate\":       {         \"document\":         {           \"content\":\"sample content\"         }       }     }   }   ```  Responds with an object with matched stored queries:     ```   {     \'timed_out\':false,     \'hits\':     {       \'total\':2,       \'max_score\':1,       \'hits\':       [         {           \'_index\':\'idx_pq_1\',           \'_type\':\'doc\',           \'_id\':\'2\',           \'_score\':\'1\',           \'_source\':           {             \'query\':             {               \'match\':{\'title\':\'some\'}             }           }         },         {           \'_index\':\'idx_pq_1\',           \'_type\':\'doc\',           \'_id\':\'5\',           \'_score\':\'1\',           \'_source\':           {             \'query\':             {               \'ql\':\'some | none\'             }           }         }       ]     }   }   ``` 
     * Perform reverse search on a percolate index
     * @param index Name of the percolate index
     * @param percolateRequest
     */
    public percolateWithHttpInfo(index: string, percolateRequest: PercolateRequest, _options?: Configuration): Promise<HttpInfo<SearchResponse>> {
        const result = this.api.percolateWithHttpInfo(index, percolateRequest, _options);
        return result.toPromise();
    }

    /**
     * Performs a percolate search.  This method must be used only on percolate indexes.  Expects two parameters: the index name and an object with array of documents to be tested. An example of the documents object:    ```   {     \"query\":     {       \"percolate\":       {         \"document\":         {           \"content\":\"sample content\"         }       }     }   }   ```  Responds with an object with matched stored queries:     ```   {     \'timed_out\':false,     \'hits\':     {       \'total\':2,       \'max_score\':1,       \'hits\':       [         {           \'_index\':\'idx_pq_1\',           \'_type\':\'doc\',           \'_id\':\'2\',           \'_score\':\'1\',           \'_source\':           {             \'query\':             {               \'match\':{\'title\':\'some\'}             }           }         },         {           \'_index\':\'idx_pq_1\',           \'_type\':\'doc\',           \'_id\':\'5\',           \'_score\':\'1\',           \'_source\':           {             \'query\':             {               \'ql\':\'some | none\'             }           }         }       ]     }   }   ``` 
     * Perform reverse search on a percolate index
     * @param index Name of the percolate index
     * @param percolateRequest
     */
    public percolate(index: string, percolateRequest: PercolateRequest, _options?: Configuration): Promise<SearchResponse> {
        const result = this.api.percolate(index, percolateRequest, _options);
        return result.toPromise();
    }

    /**
     *  The method expects an object with the following mandatory properties: * the name of the index to search * the match query object For details, see the documentation on [**SearchRequest**](SearchRequest.md) The method returns an object with the following properties: - took: the time taken to execute the search query. - timed_out: a boolean indicating whether the query timed out. - hits: an object with the following properties:    - total: the total number of hits found.    - hits: an array of hit objects, where each hit object represents a matched document. Each hit object has the following properties:      - _id: the ID of the matched document.      - _score: the score of the matched document.      - _source: the source data of the matched document.  In addition, if profiling is enabled, the response will include an additional array with profiling information attached. Here is an example search response:    ```   {     \'took\':10,     \'timed_out\':false,     \'hits\':     {       \'total\':2,       \'hits\':       [         {\'_id\':\'1\',\'_score\':1,\'_source\':{\'gid\':11}},         {\'_id\':\'2\',\'_score\':1,\'_source\':{\'gid\':12}}       ]     }   }   ```  For more information about the match query syntax and additional parameters that can be added to request and response, please see the documentation [here](https://manual.manticoresearch.com/Searching/Full_text_matching/Basic_usage#HTTP-JSON). 
     * Performs a search on an index
     * @param searchRequest
     */
    public searchWithHttpInfo(searchRequest: SearchRequest, _options?: Configuration): Promise<HttpInfo<SearchResponse>> {
        const result = this.api.searchWithHttpInfo(searchRequest, _options);
        return result.toPromise();
    }

    /**
     *  The method expects an object with the following mandatory properties: * the name of the index to search * the match query object For details, see the documentation on [**SearchRequest**](SearchRequest.md) The method returns an object with the following properties: - took: the time taken to execute the search query. - timed_out: a boolean indicating whether the query timed out. - hits: an object with the following properties:    - total: the total number of hits found.    - hits: an array of hit objects, where each hit object represents a matched document. Each hit object has the following properties:      - _id: the ID of the matched document.      - _score: the score of the matched document.      - _source: the source data of the matched document.  In addition, if profiling is enabled, the response will include an additional array with profiling information attached. Here is an example search response:    ```   {     \'took\':10,     \'timed_out\':false,     \'hits\':     {       \'total\':2,       \'hits\':       [         {\'_id\':\'1\',\'_score\':1,\'_source\':{\'gid\':11}},         {\'_id\':\'2\',\'_score\':1,\'_source\':{\'gid\':12}}       ]     }   }   ```  For more information about the match query syntax and additional parameters that can be added to request and response, please see the documentation [here](https://manual.manticoresearch.com/Searching/Full_text_matching/Basic_usage#HTTP-JSON). 
     * Performs a search on an index
     * @param searchRequest
     */
    public search(searchRequest: SearchRequest, _options?: Configuration): Promise<SearchResponse> {
        const result = this.api.search(searchRequest, _options);
        return result.toPromise();
    }


}



import { ObservableUtilsApi } from './ObservableAPI';

import { UtilsApiRequestFactory, UtilsApiResponseProcessor} from "../apis/UtilsApi";
export class PromiseUtilsApi {
    private api: ObservableUtilsApi

    public constructor(
        configuration: Configuration,
        requestFactory?: UtilsApiRequestFactory,
        responseProcessor?: UtilsApiResponseProcessor
    ) {
        this.api = new ObservableUtilsApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * Run a query in SQL format. Expects a query string passed through `body` parameter and optional `raw_response` parameter that defines a format of response. `raw_response` can be set to `False` for Select queries only, e.g., `SELECT * FROM myindex` The query string must stay as it is, no URL encoding is needed. The response object depends on the query executed. In select mode the response has same format as `/search` operation. 
     * Perform SQL requests
     * @param body A query parameter string. 
     * @param [rawResponse] Optional parameter, defines a format of response. Can be set to &#x60;False&#x60; for Select only queries and set to &#x60;True&#x60; for any type of queries. Default value is \&#39;True\&#39;. 
     */
    public sqlWithHttpInfo(body: string, rawResponse?: boolean, _options?: Configuration): Promise<HttpInfo<SqlResponse>> {
        const result = this.api.sqlWithHttpInfo(body, rawResponse, _options);
        return result.toPromise();
    }

    /**
     * Run a query in SQL format. Expects a query string passed through `body` parameter and optional `raw_response` parameter that defines a format of response. `raw_response` can be set to `False` for Select queries only, e.g., `SELECT * FROM myindex` The query string must stay as it is, no URL encoding is needed. The response object depends on the query executed. In select mode the response has same format as `/search` operation. 
     * Perform SQL requests
     * @param body A query parameter string. 
     * @param [rawResponse] Optional parameter, defines a format of response. Can be set to &#x60;False&#x60; for Select only queries and set to &#x60;True&#x60; for any type of queries. Default value is \&#39;True\&#39;. 
     */
    public sql(body: string, rawResponse?: boolean, _options?: Configuration): Promise<SqlResponse> {
        const result = this.api.sql(body, rawResponse, _options);
        return result.toPromise();
    }


}


