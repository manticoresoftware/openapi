/**
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 5.0.0
 * Contact: info@manticoresearch.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import BoolFilter from './BoolFilter';
import GeoDistance from './GeoDistance';
import Highlight from './Highlight';
import QueryFilter from './QueryFilter';

/**
 * The SearchQuery model module.
 * @module model/SearchQuery
 * @version 8.1.0
 */
class SearchQuery {
    /**
     * Constructs a new <code>SearchQuery</code>.
     * Defines a query structure for performing search operations
     * @alias module:model/SearchQuery
     * @implements module:model/QueryFilter
     */
    constructor() { 
        QueryFilter.initialize(this);
        SearchQuery.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SearchQuery</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SearchQuery} obj Optional instance to populate.
     * @return {module:model/SearchQuery} The populated <code>SearchQuery</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SearchQuery();
            QueryFilter.constructFromObject(data, obj);

            if (data.hasOwnProperty('query_string')) {
                obj['query_string'] = ApiClient.convertToType(data['query_string'], 'String');
            }
            if (data.hasOwnProperty('match')) {
                obj['match'] = ApiClient.convertToType(data['match'], Object);
            }
            if (data.hasOwnProperty('match_phrase')) {
                obj['match_phrase'] = ApiClient.convertToType(data['match_phrase'], Object);
            }
            if (data.hasOwnProperty('match_all')) {
                obj['match_all'] = ApiClient.convertToType(data['match_all'], Object);
            }
            if (data.hasOwnProperty('bool')) {
                obj['bool'] = BoolFilter.constructFromObject(data['bool']);
            }
            if (data.hasOwnProperty('equals')) {
                obj['equals'] = ApiClient.convertToType(data['equals'], Object);
            }
            if (data.hasOwnProperty('in')) {
                obj['in'] = ApiClient.convertToType(data['in'], Object);
            }
            if (data.hasOwnProperty('range')) {
                obj['range'] = ApiClient.convertToType(data['range'], Object);
            }
            if (data.hasOwnProperty('geo_distance')) {
                obj['geo_distance'] = GeoDistance.constructFromObject(data['geo_distance']);
            }
            if (data.hasOwnProperty('highlight')) {
                obj['highlight'] = Highlight.constructFromObject(data['highlight']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>SearchQuery</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>SearchQuery</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['query_string'] && !(typeof data['query_string'] === 'string' || data['query_string'] instanceof String)) {
            throw new Error("Expected the field `query_string` to be a primitive type in the JSON string but got " + data['query_string']);
        }
        // validate the optional field `bool`
        if (data['bool']) { // data not null
          BoolFilter.validateJSON(data['bool']);
        }
        // validate the optional field `highlight`
        if (data['highlight']) { // data not null
          Highlight.validateJSON(data['highlight']);
        }

        return true;
    }


}



/**
 * Filter object defining a query string
 * @member {String} query_string
 */
SearchQuery.prototype['query_string'] = undefined;

/**
 * Filter object defining a match keyword passed as a string or in a Match object
 * @member {Object} match
 */
SearchQuery.prototype['match'] = undefined;

/**
 * Filter object defining a match phrase
 * @member {Object} match_phrase
 */
SearchQuery.prototype['match_phrase'] = undefined;

/**
 * Filter object to select all documents
 * @member {Object} match_all
 */
SearchQuery.prototype['match_all'] = undefined;

/**
 * @member {module:model/BoolFilter} bool
 */
SearchQuery.prototype['bool'] = undefined;

/**
 * @member {Object} equals
 */
SearchQuery.prototype['equals'] = undefined;

/**
 * Filter to match a given set of attribute values.
 * @member {Object} in
 */
SearchQuery.prototype['in'] = undefined;

/**
 * Filter to match a given range of attribute values passed in Range objects
 * @member {Object} range
 */
SearchQuery.prototype['range'] = undefined;

/**
 * @member {module:model/GeoDistance} geo_distance
 */
SearchQuery.prototype['geo_distance'] = undefined;

/**
 * @member {module:model/Highlight} highlight
 */
SearchQuery.prototype['highlight'] = undefined;


// Implement QueryFilter interface:
/**
 * Filter object defining a query string
 * @member {String} query_string
 */
QueryFilter.prototype['query_string'] = undefined;
/**
 * Filter object defining a match keyword passed as a string or in a Match object
 * @member {Object} match
 */
QueryFilter.prototype['match'] = undefined;
/**
 * Filter object defining a match phrase
 * @member {Object} match_phrase
 */
QueryFilter.prototype['match_phrase'] = undefined;
/**
 * Filter object to select all documents
 * @member {Object} match_all
 */
QueryFilter.prototype['match_all'] = undefined;
/**
 * @member {module:model/BoolFilter} bool
 */
QueryFilter.prototype['bool'] = undefined;
/**
 * @member {Object} equals
 */
QueryFilter.prototype['equals'] = undefined;
/**
 * Filter to match a given set of attribute values.
 * @member {Object} in
 */
QueryFilter.prototype['in'] = undefined;
/**
 * Filter to match a given range of attribute values passed in Range objects
 * @member {Object} range
 */
QueryFilter.prototype['range'] = undefined;
/**
 * @member {module:model/GeoDistance} geo_distance
 */
QueryFilter.prototype['geo_distance'] = undefined;




export default SearchQuery;

