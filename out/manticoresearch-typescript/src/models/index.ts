/* tslint:disable */
/* eslint-disable */
/**
 * Object to perform composite aggregation, i.e., grouping search results by multiple fields
 * @export
 * @interface AggComposite
 */
export interface AggComposite {
    /**
     * Maximum number of composite buckets in the result
     * @type {number}
     * @memberof AggComposite
     */
    size?: number;
    /**
     * 
     * @type {Array<{ [key: string]: AggCompositeSource; }>}
     * @memberof AggComposite
     */
    sources?: Array<{ [key: string]: AggCompositeSource; }>;
}
/**
 * Object containing terms used for composite aggregation.
 * @export
 * @interface AggCompositeSource
 */
export interface AggCompositeSource {
    /**
     * 
     * @type {AggCompositeTerm}
     * @memberof AggCompositeSource
     */
    terms: AggCompositeTerm;
}
/**
 * Object representing a term to be used in composite aggregation.
 * @export
 * @interface AggCompositeTerm
 */
export interface AggCompositeTerm {
    /**
     * Name of field to operate with
     * @type {string}
     * @memberof AggCompositeTerm
     */
    field: string;
}
/**
 * Object containing term fields to aggregate on
 * @export
 * @interface AggTerms
 */
export interface AggTerms {
    /**
     * Name of attribute to aggregate by
     * @type {string}
     * @memberof AggTerms
     */
    field: string;
    /**
     * Maximum number of buckets in the result
     * @type {number}
     * @memberof AggTerms
     */
    size?: number;
}
/**
 * 
 * @export
 * @interface Aggregation
 */
export interface Aggregation {
    /**
     * 
     * @type {AggTerms}
     * @memberof Aggregation
     */
    terms?: AggTerms;
    /**
     * 
     * @type {Array<any>}
     * @memberof Aggregation
     */
    sort?: Array<any>;
    /**
     * 
     * @type {AggComposite}
     * @memberof Aggregation
     */
    composite?: AggComposite;
}
/**
 * 
 * @export
 * @interface BoolFilter
 */
export interface BoolFilter {
    /**
     * Query clauses that must match for the document to be included
     * @type {Array<QueryFilter>}
     * @memberof BoolFilter
     */
    must?: Array<QueryFilter>;
    /**
     * Query clauses that must not match for the document to be included
     * @type {Array<QueryFilter>}
     * @memberof BoolFilter
     */
    must_not?: Array<QueryFilter>;
    /**
     * Query clauses that should be matched, but are not required
     * @type {Array<QueryFilter>}
     * @memberof BoolFilter
     */
    should?: Array<QueryFilter>;
}
/**
 * Success response for bulk search requests
 * @export
 * @interface BulkResponse
 */
export interface BulkResponse {
    /**
     * List of results
     * @type {Array<object>}
     * @memberof BulkResponse
     */
    items?: Array<object>;
    /**
     * Errors occurred during the bulk operation
     * @type {boolean}
     * @memberof BulkResponse
     */
    errors?: boolean;
    /**
     * Error message describing an error if such occurred
     * @type {string}
     * @memberof BulkResponse
     */
    error?: string;
}
/**
 * Payload for delete request.
 * Documents can be deleted either one by one by specifying the document id or by providing a query object.
 * For more information see  [Delete API](https://manual.manticoresearch.com/Deleting_documents)
 * 
 * @export
 * @interface DeleteDocumentRequest
 */
export interface DeleteDocumentRequest {
    /**
     * Index name
     * @type {string}
     * @memberof DeleteDocumentRequest
     */
    index: string;
    /**
     * Cluster name
     * @type {string}
     * @memberof DeleteDocumentRequest
     */
    cluster?: string;
    /**
     * The ID of document for deletion
     * @type {number}
     * @memberof DeleteDocumentRequest
     */
    id?: number;
    /**
     * Defines the criteria to match documents for deletion
     * @type {object}
     * @memberof DeleteDocumentRequest
     */
    query?: object;
}
/**
 * Response object for successful delete request
 * @export
 * @interface DeleteResponse
 */
export interface DeleteResponse {
    /**
     * The name of the index from which the document was deleted
     * @type {string}
     * @memberof DeleteResponse
     */
    _index?: string;
    /**
     * Number of documents deleted
     * @type {number}
     * @memberof DeleteResponse
     */
    deleted?: number;
    /**
     * The ID of the deleted document. If multiple documents are deleted, the ID of the first deleted document is returned
     * @type {number}
     * @memberof DeleteResponse
     */
    _id?: number;
    /**
     * Indicates whether any documents to be deleted were found
     * @type {boolean}
     * @memberof DeleteResponse
     */
    found?: boolean;
    /**
     * Result of the delete operation, typically 'deleted'
     * @type {string}
     * @memberof DeleteResponse
     */
    result?: string;
}
/**
 * Error response object containing information about the error and a status code
 * @export
 * @interface ErrorResponse
 */
export interface ErrorResponse {
    /**
     * 
     * @type {ResponseError}
     * @memberof ErrorResponse
     */
    error: ResponseError;
    /**
     * HTTP status code of the error response
     * @type {number}
     * @memberof ErrorResponse
     */
    status?: number;
}
/**
 * Defines a type of filter for full-text search queries
 * @export
 * @interface FulltextFilter
 */
export interface FulltextFilter {
    /**
     * Filter object defining a query string
     * @type {string}
     * @memberof FulltextFilter
     */
    query_string?: string;
    /**
     * Filter object defining a match keyword
     * @type {object}
     * @memberof FulltextFilter
     */
    match?: object;
    /**
     * Filter object defining a match phrase
     * @type {object}
     * @memberof FulltextFilter
     */
    match_phrase?: object;
    /**
     * Filter object to select all documents
     * @type {object}
     * @memberof FulltextFilter
     */
    match_all?: object;
}
/**
 * Object to perform geo-distance based filtering on queries
 * @export
 * @interface GeoDistance
 */
export interface GeoDistance {
    [key: string]: any | any;
    /**
     * 
     * @type {GeoDistanceLocationAnchor}
     * @memberof GeoDistance
     */
    location_anchor?: GeoDistanceLocationAnchor;
    /**
     * Field name in the document that contains location data
     * @type {any}
     * @memberof GeoDistance
     */
    location_source?: any | null;
    /**
     * Algorithm used to calculate the distance
     * @type {any}
     * @memberof GeoDistance
     */
    distance_type?: GeoDistanceDistanceTypeEnum | null;
    /**
     * The distance from the anchor point to filter results by
     * @type {any}
     * @memberof GeoDistance
     */
    distance?: any | null;
}

/**
* @export
* @enum {string}
*/
export enum GeoDistanceDistanceTypeEnum {
    adaptive = 'adaptive',
    haversine = 'haversine'
}

/**
 * Specifies the location of the pin point used for search
 * @export
 * @interface GeoDistanceLocationAnchor
 */
export interface GeoDistanceLocationAnchor {
    /**
     * Latitude of the anchor point
     * @type {any}
     * @memberof GeoDistanceLocationAnchor
     */
    lat?: any | null;
    /**
     * Longitude of the anchor point
     * @type {any}
     * @memberof GeoDistanceLocationAnchor
     */
    lon?: any | null;
}
/**
 * 
 * @export
 * @interface Highlight
 */
export interface Highlight {
    /**
     * Maximum size of the text fragments in highlighted snippets per field
     * @type {any}
     * @memberof Highlight
     */
    fragment_size?: any | null;
    /**
     * Maximum size of snippets per field
     * @type {any}
     * @memberof Highlight
     */
    limit?: any | null;
    /**
     * Maximum number of snippets per field
     * @type {any}
     * @memberof Highlight
     */
    limit_snippets?: any | null;
    /**
     * Maximum number of words per field
     * @type {any}
     * @memberof Highlight
     */
    limit_words?: any | null;
    /**
     * Total number of highlighted fragments per field
     * @type {any}
     * @memberof Highlight
     */
    number_of_fragments?: any | null;
    /**
     * Text inserted after the matched term, typically used for HTML formatting
     * @type {string}
     * @memberof Highlight
     */
    after_match?: string;
    /**
     * Permits an empty string to be returned as the highlighting result. Otherwise, the beginning of the original text would be returned
     * @type {boolean}
     * @memberof Highlight
     */
    allow_empty?: boolean;
    /**
     * Number of words around the match to include in the highlight
     * @type {number}
     * @memberof Highlight
     */
    around?: number;
    /**
     * Text inserted before the match, typically used for HTML formatting
     * @type {string}
     * @memberof Highlight
     */
    before_match?: string;
    /**
     * Emits an HTML tag with the enclosing zone name before each highlighted snippet
     * @type {boolean}
     * @memberof Highlight
     */
    emit_zones?: boolean;
    /**
     * If set to 'html', retains HTML markup when highlighting
     * @type {string}
     * @memberof Highlight
     */
    encoder?: HighlightEncoderEnum;
    /**
     * 
     * @type {object}
     * @memberof Highlight
     */
    fields?: object | null;
    /**
     * Ignores the length limit until the result includes all keywords
     * @type {boolean}
     * @memberof Highlight
     */
    force_all_words?: boolean;
    /**
     * Forces snippet generation even if limits allow highlighting the entire text
     * @type {boolean}
     * @memberof Highlight
     */
    force_snippets?: boolean;
    /**
     * 
     * @type {QueryFilter}
     * @memberof Highlight
     */
    highlight_query?: QueryFilter | null;
    /**
     * Defines the mode for handling HTML markup in the highlight
     * @type {string}
     * @memberof Highlight
     */
    html_strip_mode?: HighlightHtmlStripModeEnum;
    /**
     * Determines whether the 'limit', 'limit_words', and 'limit_snippets' options operate as individual limits in each field of the document
     * @type {boolean}
     * @memberof Highlight
     */
    limits_per_field?: boolean;
    /**
     * If set to 1, allows an empty string to be returned as a highlighting result
     * @type {number}
     * @memberof Highlight
     */
    no_match_size?: HighlightNoMatchSizeEnum;
    /**
     * Sets the sorting order of highlighted snippets
     * @type {string}
     * @memberof Highlight
     */
    order?: HighlightOrderEnum;
    /**
     * Text inserted before each highlighted snippet
     * @type {string}
     * @memberof Highlight
     */
    pre_tags?: string;
    /**
     * Text inserted after each highlighted snippet
     * @type {string}
     * @memberof Highlight
     */
    post_tags?: string;
    /**
     * Sets the starting value of the %SNIPPET_ID% macro
     * @type {number}
     * @memberof Highlight
     */
    start_snippet_id?: number;
    /**
     * Defines whether to additionally break snippets by phrase boundary characters
     * @type {boolean}
     * @memberof Highlight
     */
    use_boundaries?: boolean;
}

/**
* @export
* @enum {string}
*/
export enum HighlightEncoderEnum {
    default = 'default',
    html = 'html'
}
/**
* @export
* @enum {string}
*/
export enum HighlightHtmlStripModeEnum {
    none = 'none',
    strip = 'strip',
    index = 'index',
    retain = 'retain'
}
/**
* @export
* @enum {string}
*/
export enum HighlightNoMatchSizeEnum {
    NUMBER_0 = 0,
    NUMBER_1 = 1
}
/**
* @export
* @enum {string}
*/
export enum HighlightOrderEnum {
    asc = 'asc',
    desc = 'desc',
    score = 'score'
}

/**
 * Options for controlling the behavior of highlighting on a per-field basis
 * @export
 * @interface HighlightFieldOption
 */
export interface HighlightFieldOption {
    /**
     * Maximum size of the text fragments in highlighted snippets per field
     * @type {number}
     * @memberof HighlightFieldOption
     */
    fragment_size?: number;
    /**
     * Maximum size of snippets per field
     * @type {number}
     * @memberof HighlightFieldOption
     */
    limit?: number;
    /**
     * Maximum number of snippets per field
     * @type {number}
     * @memberof HighlightFieldOption
     */
    limit_snippets?: number;
    /**
     * Maximum number of words per field
     * @type {number}
     * @memberof HighlightFieldOption
     */
    limit_words?: number;
    /**
     * Total number of highlighted fragments per field
     * @type {number}
     * @memberof HighlightFieldOption
     */
    number_of_fragments?: number;
}
/**
 * Object containing data for inserting a new document into the index
 * 
 * @export
 * @interface InsertDocumentRequest
 */
export interface InsertDocumentRequest {
    /**
     * Name of the index to insert the document into
     * @type {string}
     * @memberof InsertDocumentRequest
     */
    index: string;
    /**
     * Name of the cluster to insert the document into
     * @type {string}
     * @memberof InsertDocumentRequest
     */
    cluster?: string;
    /**
     * Document ID. If not provided, an ID will be auto-generated
     * 
     * @type {number}
     * @memberof InsertDocumentRequest
     */
    id?: number;
    /**
     * Object containing document data
     * 
     * @type {object}
     * @memberof InsertDocumentRequest
     */
    doc: object;
}
/**
 * 
 * @export
 * @interface Join
 */
export interface Join {
    /**
     * Type of the join operation
     * @type {string}
     * @memberof Join
     */
    type: JoinTypeEnum;
    /**
     * List of objects defining joined tables
     * @type {Array<JoinOn>}
     * @memberof Join
     */
    on: Array<JoinOn>;
    /**
     * 
     * @type {FulltextFilter}
     * @memberof Join
     */
    query?: FulltextFilter;
    /**
     * Basic table of the join operation
     * @type {string}
     * @memberof Join
     */
    table: string;
}

/**
* @export
* @enum {string}
*/
export enum JoinTypeEnum {
    inner = 'inner',
    left = 'left'
}

/**
 * Object representing the conditions used to perform the join operation
 * @export
 * @interface JoinCond
 */
export interface JoinCond {
    /**
     * Field to join on
     * @type {string}
     * @memberof JoinCond
     */
    field: string;
    /**
     * Joined table
     * @type {string}
     * @memberof JoinCond
     */
    table: string;
    /**
     * 
     * @type {any}
     * @memberof JoinCond
     */
    type?: any | null;
}
/**
 * 
 * @export
 * @interface JoinOn
 */
export interface JoinOn {
    /**
     * 
     * @type {JoinCond}
     * @memberof JoinOn
     */
    right?: JoinCond;
    /**
     * 
     * @type {JoinCond}
     * @memberof JoinOn
     */
    left?: JoinCond;
    /**
     * 
     * @type {string}
     * @memberof JoinOn
     */
    operator?: JoinOnOperatorEnum;
}

/**
* @export
* @enum {string}
*/
export enum JoinOnOperatorEnum {
    eq = 'eq'
}

/**
 * Object representing a k-nearest neighbor search query
 * @export
 * @interface KnnQuery
 */
export interface KnnQuery {
    /**
     * Field to perform the k-nearest neighbor search on
     * @type {string}
     * @memberof KnnQuery
     */
    field: string;
    /**
     * The number of nearest neighbors to return
     * @type {number}
     * @memberof KnnQuery
     */
    k: number;
    /**
     * The vector used as input for the KNN search
     * @type {Array<number>}
     * @memberof KnnQuery
     */
    query_vector?: Array<number>;
    /**
     * The docuemnt ID used as input for the KNN search
     * @type {number}
     * @memberof KnnQuery
     */
    doc_id?: number;
    /**
     * Optional parameter controlling the accuracy of the search
     * @type {number}
     * @memberof KnnQuery
     */
    ef?: number;
    /**
     * 
     * @type {QueryFilter}
     * @memberof KnnQuery
     */
    filter?: QueryFilter;
}
/**
 * Object containing the query for percolating documents against stored queries in a percolate index
 * @export
 * @interface PercolateRequest
 */
export interface PercolateRequest {
    /**
     * 
     * @type {PercolateRequestQuery}
     * @memberof PercolateRequest
     */
    query: PercolateRequestQuery;
}
/**
 * 
 * @export
 * @interface PercolateRequestQuery
 */
export interface PercolateRequestQuery {
    /**
     * Object representing the document to percolate
     * @type {object}
     * @memberof PercolateRequestQuery
     */
    percolate: object;
}
/**
 * Object used to apply various conditions, such as full-text matching or attribute filtering, to a search query
 * @export
 * @interface QueryFilter
 */
export interface QueryFilter {
    /**
     * Filter object defining a query string
     * @type {any}
     * @memberof QueryFilter
     */
    query_string?: any | null;
    /**
     * Filter object defining a match keyword
     * @type {any}
     * @memberof QueryFilter
     */
    match?: any | null;
    /**
     * Filter object defining a match phrase
     * @type {any}
     * @memberof QueryFilter
     */
    match_phrase?: any | null;
    /**
     * Filter object to select all documents
     * @type {any}
     * @memberof QueryFilter
     */
    match_all?: any | null;
    /**
     * 
     * @type {BoolFilter}
     * @memberof QueryFilter
     */
    bool?: BoolFilter;
    /**
     * 
     * @type {any}
     * @memberof QueryFilter
     */
    equals?: any | null;
    /**
     * Filter to match a given set of attribute values.
     * @type {object}
     * @memberof QueryFilter
     */
    _in?: object;
    /**
     * Filter to match a given range of attribute values.
     * @type {object}
     * @memberof QueryFilter
     */
    range?: object;
    /**
     * 
     * @type {GeoDistance}
     * @memberof QueryFilter
     */
    geo_distance?: GeoDistance;
}
/**
 * Object containing the document data for replacing an existing document in an index.
 * @export
 * @interface ReplaceDocumentRequest
 */
export interface ReplaceDocumentRequest {
    /**
     * Object containing the new document data to replace the existing one.
     * @type {object}
     * @memberof ReplaceDocumentRequest
     */
    doc: object;
}
/**
 * @type ResponseError
 * 
 * @export
 */
export type ResponseError = ResponseErrorDetails | string;
/**
 * Detailed error information returned in case of an error response
 * @export
 * @interface ResponseErrorDetails
 */
export interface ResponseErrorDetails {
    /**
     * Type or category of the error
     * @type {string}
     * @memberof ResponseErrorDetails
     */
    type: string;
    /**
     * Detailed explanation of why the error occurred
     * @type {string}
     * @memberof ResponseErrorDetails
     */
    reason?: string | null;
    /**
     * The index related to the error, if applicable
     * @type {string}
     * @memberof ResponseErrorDetails
     */
    index?: string | null;
}
/**
 * Defines a query structure for performing search operations
 * @export
 * @interface SearchQuery
 */
export interface SearchQuery {
    /**
     * Filter object defining a query string
     * @type {any}
     * @memberof SearchQuery
     */
    query_string?: any | null;
    /**
     * Filter object defining a match keyword
     * @type {any}
     * @memberof SearchQuery
     */
    match?: any | null;
    /**
     * Filter object defining a match phrase
     * @type {any}
     * @memberof SearchQuery
     */
    match_phrase?: any | null;
    /**
     * Filter object to select all documents
     * @type {any}
     * @memberof SearchQuery
     */
    match_all?: any | null;
    /**
     * 
     * @type {BoolFilter}
     * @memberof SearchQuery
     */
    bool?: BoolFilter;
    /**
     * 
     * @type {any}
     * @memberof SearchQuery
     */
    equals?: any | null;
    /**
     * Filter to match a given set of attribute values.
     * @type {any}
     * @memberof SearchQuery
     */
    _in?: any | null;
    /**
     * Filter to match a given range of attribute values.
     * @type {any}
     * @memberof SearchQuery
     */
    range?: any | null;
    /**
     * 
     * @type {GeoDistance}
     * @memberof SearchQuery
     */
    geo_distance?: GeoDistance;
    /**
     * 
     * @type {Highlight}
     * @memberof SearchQuery
     */
    highlight?: Highlight;
}
/**
 * Request object for search operation
 * @export
 * @interface SearchRequest
 */
export interface SearchRequest {
    /**
     * The index to perform the search on
     * @type {string}
     * @memberof SearchRequest
     */
    index: string;
    /**
     * 
     * @type {SearchQuery}
     * @memberof SearchRequest
     */
    query?: SearchQuery;
    /**
     * Join clause to combine search data from multiple tables
     * @type {Array<Join>}
     * @memberof SearchRequest
     */
    join?: Array<Join>;
    /**
     * 
     * @type {Highlight}
     * @memberof SearchRequest
     */
    highlight?: Highlight;
    /**
     * Maximum number of results to return
     * @type {number}
     * @memberof SearchRequest
     */
    limit?: number;
    /**
     * 
     * @type {KnnQuery}
     * @memberof SearchRequest
     */
    knn?: KnnQuery;
    /**
     * Defines aggregation settings for grouping results
     * @type {{ [key: string]: Aggregation; }}
     * @memberof SearchRequest
     */
    aggs?: { [key: string]: Aggregation; } | null;
    /**
     * Expressions to calculate additional values for the result
     * @type {{ [key: string]: string; }}
     * @memberof SearchRequest
     */
    expressions?: { [key: string]: string; } | null;
    /**
     * Maximum number of matches allowed in the result
     * @type {number}
     * @memberof SearchRequest
     */
    max_matches?: number;
    /**
     * Starting point for pagination of the result
     * @type {number}
     * @memberof SearchRequest
     */
    offset?: number;
    /**
     * Additional search options
     * @type {object}
     * @memberof SearchRequest
     */
    options?: object;
    /**
     * Enable or disable profiling of the search request
     * @type {boolean}
     * @memberof SearchRequest
     */
    profile?: boolean;
    /**
     * 
     * @type {any}
     * @memberof SearchRequest
     */
    sort?: any | null;
    /**
     * 
     * @type {any}
     * @memberof SearchRequest
     */
    _source?: any | null;
    /**
     * Enable or disable result weight calculation used for sorting
     * @type {boolean}
     * @memberof SearchRequest
     */
    track_scores?: boolean;
}
/**
 * Response object containing the results of a search request
 * @export
 * @interface SearchResponse
 */
export interface SearchResponse {
    /**
     * Time taken to execute the search
     * @type {number}
     * @memberof SearchResponse
     */
    took?: number;
    /**
     * Indicates whether the search operation timed out
     * @type {boolean}
     * @memberof SearchResponse
     */
    timed_out?: boolean;
    /**
     * Aggregated search results grouped by the specified criteria
     * @type {object}
     * @memberof SearchResponse
     */
    aggregations?: object;
    /**
     * 
     * @type {SearchResponseHits}
     * @memberof SearchResponse
     */
    hits?: SearchResponseHits;
    /**
     * Profile information about the search execution, if profiling is enabled
     * @type {object}
     * @memberof SearchResponse
     */
    profile?: object;
    /**
     * Warnings encountered during the search operation
     * @type {object}
     * @memberof SearchResponse
     */
    warning?: object;
}
/**
 * Object containing the search hits, which represent the documents that matched the query.
 * @export
 * @interface SearchResponseHits
 */
export interface SearchResponseHits {
    /**
     * Maximum score among the matched documents
     * @type {number}
     * @memberof SearchResponseHits
     */
    max_score?: number;
    /**
     * Total number of matched documents
     * @type {number}
     * @memberof SearchResponseHits
     */
    total?: number;
    /**
     * Indicates whether the total number of hits is accurate or an estimate
     * @type {string}
     * @memberof SearchResponseHits
     */
    total_relation?: string;
    /**
     * Array of hit objects, each representing a matched document
     * @type {Array<object>}
     * @memberof SearchResponseHits
     */
    hits?: Array<object>;
}
/**
 * Defines which fields to include or exclude in the response for a search query
 * @export
 * @interface SourceRules
 */
export interface SourceRules {
    [key: string]: any | any;
    /**
     * List of fields to include in the response
     * @type {any}
     * @memberof SourceRules
     */
    includes?: any | null;
    /**
     * List of fields to exclude from the response
     * @type {any}
     * @memberof SourceRules
     */
    excludes?: any | null;
}
/**
 * Response object indicating the success of an operation, such as inserting or updating a document
 * @export
 * @interface SuccessResponse
 */
export interface SuccessResponse {
    /**
     * Name of the document index
     * @type {string}
     * @memberof SuccessResponse
     */
    _index?: string;
    /**
     * ID of the document affected by the request operation
     * @type {number}
     * @memberof SuccessResponse
     */
    _id?: number;
    /**
     * Indicates whether the document was created as a result of the operation
     * @type {boolean}
     * @memberof SuccessResponse
     */
    created?: boolean;
    /**
     * Result of the operation, typically 'created', 'updated', or 'deleted'
     * @type {string}
     * @memberof SuccessResponse
     */
    result?: string;
    /**
     * Indicates whether the document was found in the index
     * @type {boolean}
     * @memberof SuccessResponse
     */
    found?: boolean;
    /**
     * HTTP status code representing the result of the operation
     * @type {number}
     * @memberof SuccessResponse
     */
    status?: number;
}
/**
 * Payload for updating a document or multiple documents in an index
 * @export
 * @interface UpdateDocumentRequest
 */
export interface UpdateDocumentRequest {
    /**
     * Name of the document index
     * @type {string}
     * @memberof UpdateDocumentRequest
     */
    index: string;
    /**
     * Name of the document cluster
     * @type {string}
     * @memberof UpdateDocumentRequest
     */
    cluster?: string;
    /**
     * Object containing the document fields to update
     * @type {object}
     * @memberof UpdateDocumentRequest
     */
    doc: object;
    /**
     * Document ID
     * @type {number}
     * @memberof UpdateDocumentRequest
     */
    id?: number;
    /**
     * 
     * @type {QueryFilter}
     * @memberof UpdateDocumentRequest
     */
    query?: QueryFilter | null;
}
/**
 * Success response returned after updating one or more documents
 * @export
 * @interface UpdateResponse
 */
export interface UpdateResponse {
    /**
     * Name of the document index
     * @type {string}
     * @memberof UpdateResponse
     */
    _index?: string;
    /**
     * Number of documents updated
     * @type {number}
     * @memberof UpdateResponse
     */
    updated?: number;
    /**
     * Document ID
     * @type {number}
     * @memberof UpdateResponse
     */
    _id?: number;
    /**
     * Result of the update operation, typically 'updated'
     * @type {string}
     * @memberof UpdateResponse
     */
    result?: string;
}
