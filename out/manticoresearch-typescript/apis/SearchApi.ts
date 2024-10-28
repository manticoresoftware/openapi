// TODO: better import syntax?
import {BaseAPIRequestFactory, RequiredError, COLLECTION_FORMATS} from './baseapi';
import {Configuration} from '../configuration';
import {RequestContext, HttpMethod, ResponseContext, HttpFile, HttpInfo} from '../http/http';
import {ObjectSerializer} from '../models/ObjectSerializer';
import {ApiException} from './exception';
import {canConsumeForm, isCodeInRange} from '../util';
import {SecurityAuthentication} from '../auth/auth';


import { ErrorResponse } from '../models/ErrorResponse';
import { PercolateRequest } from '../models/PercolateRequest';
import { SearchRequest } from '../models/SearchRequest';
import { SearchResponse } from '../models/SearchResponse';

/**
 * no description
 */
export class SearchApiRequestFactory extends BaseAPIRequestFactory {

    /**
     * Performs a percolate search.  This method must be used only on percolate indexes.  Expects two parameters: the index name and an object with array of documents to be tested. An example of the documents object:    ```   {     \"query\":     {       \"percolate\":       {         \"document\":         {           \"content\":\"sample content\"         }       }     }   }   ```  Responds with an object with matched stored queries:     ```   {     \'timed_out\':false,     \'hits\':     {       \'total\':2,       \'max_score\':1,       \'hits\':       [         {           \'_index\':\'idx_pq_1\',           \'_type\':\'doc\',           \'_id\':\'2\',           \'_score\':\'1\',           \'_source\':           {             \'query\':             {               \'match\':{\'title\':\'some\'}             }           }         },         {           \'_index\':\'idx_pq_1\',           \'_type\':\'doc\',           \'_id\':\'5\',           \'_score\':\'1\',           \'_source\':           {             \'query\':             {               \'ql\':\'some | none\'             }           }         }       ]     }   }   ``` 
     * Perform reverse search on a percolate index
     * @param index Name of the percolate index
     * @param percolateRequest 
     */
    public async percolate(index: string, percolateRequest: PercolateRequest, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;

        // verify required parameter 'index' is not null or undefined
        if (index === null || index === undefined) {
            throw new RequiredError("SearchApi", "percolate", "index");
        }


        // verify required parameter 'percolateRequest' is not null or undefined
        if (percolateRequest === null || percolateRequest === undefined) {
            throw new RequiredError("SearchApi", "percolate", "percolateRequest");
        }


        // Path Params
        const localVarPath = '/pq/{index}/search'
            .replace('{' + 'index' + '}', encodeURIComponent(String(index)));

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.POST);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")


        // Body Params
        const contentType = ObjectSerializer.getPreferredMediaType([
            "application/json"
        ]);
        requestContext.setHeaderParam("Content-Type", contentType);
        const serializedBody = ObjectSerializer.stringify(
            ObjectSerializer.serialize(percolateRequest, "PercolateRequest", ""),
            contentType
        );
        requestContext.setBody(serializedBody);

        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
    }

    /**
     *  The method expects an object with the following mandatory properties: * the name of the index to search * the match query object For details, see the documentation on [**SearchRequest**](SearchRequest.md) The method returns an object with the following properties: - took: the time taken to execute the search query. - timed_out: a boolean indicating whether the query timed out. - hits: an object with the following properties:    - total: the total number of hits found.    - hits: an array of hit objects, where each hit object represents a matched document. Each hit object has the following properties:      - _id: the ID of the matched document.      - _score: the score of the matched document.      - _source: the source data of the matched document.  In addition, if profiling is enabled, the response will include an additional array with profiling information attached. Here is an example search response:    ```   {     \'took\':10,     \'timed_out\':false,     \'hits\':     {       \'total\':2,       \'hits\':       [         {\'_id\':\'1\',\'_score\':1,\'_source\':{\'gid\':11}},         {\'_id\':\'2\',\'_score\':1,\'_source\':{\'gid\':12}}       ]     }   }   ```  For more information about the match query syntax and additional parameters that can be added to request and response, please see the documentation [here](https://manual.manticoresearch.com/Searching/Full_text_matching/Basic_usage#HTTP-JSON). 
     * Performs a search on an index
     * @param searchRequest 
     */
    public async search(searchRequest: SearchRequest, _options?: Configuration): Promise<RequestContext> {
        let _config = _options || this.configuration;

        // verify required parameter 'searchRequest' is not null or undefined
        if (searchRequest === null || searchRequest === undefined) {
            throw new RequiredError("SearchApi", "search", "searchRequest");
        }


        // Path Params
        const localVarPath = '/search';

        // Make Request Context
        const requestContext = _config.baseServer.makeRequestContext(localVarPath, HttpMethod.POST);
        requestContext.setHeaderParam("Accept", "application/json, */*;q=0.8")


        // Body Params
        const contentType = ObjectSerializer.getPreferredMediaType([
            "application/json"
        ]);
        requestContext.setHeaderParam("Content-Type", contentType);
        const serializedBody = ObjectSerializer.stringify(
            ObjectSerializer.serialize(searchRequest, "SearchRequest", ""),
            contentType
        );
        requestContext.setBody(serializedBody);

        
        const defaultAuth: SecurityAuthentication | undefined = _options?.authMethods?.default || this.configuration?.authMethods?.default
        if (defaultAuth?.applySecurityAuthentication) {
            await defaultAuth?.applySecurityAuthentication(requestContext);
        }

        return requestContext;
    }

}

export class SearchApiResponseProcessor {

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to percolate
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async percolateWithHttpInfo(response: ResponseContext): Promise<HttpInfo<SearchResponse >> {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: SearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SearchResponse", ""
            ) as SearchResponse;
            return new HttpInfo(response.httpStatusCode, response.headers, response.body, body);
        }
        if (isCodeInRange("0", response.httpStatusCode)) {
            const body: ErrorResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "ErrorResponse", ""
            ) as ErrorResponse;
            throw new ApiException<ErrorResponse>(response.httpStatusCode, "error", body, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: SearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SearchResponse", ""
            ) as SearchResponse;
            return new HttpInfo(response.httpStatusCode, response.headers, response.body, body);
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

    /**
     * Unwraps the actual response sent by the server from the response context and deserializes the response content
     * to the expected objects
     *
     * @params response Response returned by the server for a request to search
     * @throws ApiException if the response code was not in [200, 299]
     */
     public async searchWithHttpInfo(response: ResponseContext): Promise<HttpInfo<SearchResponse >> {
        const contentType = ObjectSerializer.normalizeMediaType(response.headers["content-type"]);
        if (isCodeInRange("200", response.httpStatusCode)) {
            const body: SearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SearchResponse", ""
            ) as SearchResponse;
            return new HttpInfo(response.httpStatusCode, response.headers, response.body, body);
        }
        if (isCodeInRange("0", response.httpStatusCode)) {
            const body: ErrorResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "ErrorResponse", ""
            ) as ErrorResponse;
            throw new ApiException<ErrorResponse>(response.httpStatusCode, "error", body, response.headers);
        }

        // Work around for missing responses in specification, e.g. for petstore.yaml
        if (response.httpStatusCode >= 200 && response.httpStatusCode <= 299) {
            const body: SearchResponse = ObjectSerializer.deserialize(
                ObjectSerializer.parse(await response.body.text(), contentType),
                "SearchResponse", ""
            ) as SearchResponse;
            return new HttpInfo(response.httpStatusCode, response.headers, response.body, body);
        }

        throw new ApiException<string | Blob | undefined>(response.httpStatusCode, "Unknown API Status Code!", await response.getBodyAsAny(), response.headers);
    }

}
