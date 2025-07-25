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

/**
 * The AutocompleteRequest model module.
 * @module model/AutocompleteRequest
 * @version 8.1.0
 */
class AutocompleteRequest {
    /**
     * Constructs a new <code>AutocompleteRequest</code>.
     * Object containing the data for performing an autocomplete search.
     * @alias module:model/AutocompleteRequest
     * @param table {String} The table to perform the search on
     * @param query {String} The beginning of the string to autocomplete
     */
    constructor(table, query) { 
        
        AutocompleteRequest.initialize(this, table, query);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, table, query) { 
        obj['table'] = table;
        obj['query'] = query;
    }

    /**
     * Constructs a <code>AutocompleteRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AutocompleteRequest} obj Optional instance to populate.
     * @return {module:model/AutocompleteRequest} The populated <code>AutocompleteRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AutocompleteRequest();

            if (data.hasOwnProperty('table')) {
                obj['table'] = ApiClient.convertToType(data['table'], 'String');
            }
            if (data.hasOwnProperty('query')) {
                obj['query'] = ApiClient.convertToType(data['query'], 'String');
            }
            if (data.hasOwnProperty('options')) {
                obj['options'] = ApiClient.convertToType(data['options'], Object);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AutocompleteRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AutocompleteRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of AutocompleteRequest.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['table'] && !(typeof data['table'] === 'string' || data['table'] instanceof String)) {
            throw new Error("Expected the field `table` to be a primitive type in the JSON string but got " + data['table']);
        }
        // ensure the json data is a string
        if (data['query'] && !(typeof data['query'] === 'string' || data['query'] instanceof String)) {
            throw new Error("Expected the field `query` to be a primitive type in the JSON string but got " + data['query']);
        }

        return true;
    }


}

AutocompleteRequest.RequiredProperties = ["table", "query"];

/**
 * The table to perform the search on
 * @member {String} table
 */
AutocompleteRequest.prototype['table'] = undefined;

/**
 * The beginning of the string to autocomplete
 * @member {String} query
 */
AutocompleteRequest.prototype['query'] = undefined;

/**
 * Autocomplete options   - layouts: A comma-separated string of keyboard layout codes to validate and check for spell correction. Available options - us, ru, ua, se, pt, no, it, gr, uk, fr, es, dk, de, ch, br, bg, be. By default, all are enabled.   - fuzziness: (0,1 or 2) Maximum Levenshtein distance for finding typos. Set to 0 to disable fuzzy matching. Default is 2   - prepend: true/false If true, adds an asterisk before the last word for prefix expansion (e.g., *word )   - append:  true/false If true, adds an asterisk after the last word for prefix expansion (e.g., word* )   - expansion_len: Number of characters to expand in the last word. Default is 10. 
 * @member {Object} options
 */
AutocompleteRequest.prototype['options'] = undefined;






export default AutocompleteRequest;

