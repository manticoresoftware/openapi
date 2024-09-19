/**
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 4.0.0
 * Contact: info@manticoresearch.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The MatchPhraseFilter model module.
 * @module model/MatchPhraseFilter
 * @version 4.0.0
 */
class MatchPhraseFilter {
    /**
     * Constructs a new <code>MatchPhraseFilter</code>.
     * @alias module:model/MatchPhraseFilter
     */
    constructor() { 
        
        MatchPhraseFilter.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>MatchPhraseFilter</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MatchPhraseFilter} obj Optional instance to populate.
     * @return {module:model/MatchPhraseFilter} The populated <code>MatchPhraseFilter</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new MatchPhraseFilter();

            if (data.hasOwnProperty('match_phrase')) {
                obj['match_phrase'] = ApiClient.convertToType(data['match_phrase'], {'String': 'String'});
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>MatchPhraseFilter</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>MatchPhraseFilter</code>.
     */
    static validateJSON(data) {

        return true;
    }


}



/**
 * @member {Object.<String, String>} match_phrase
 */
MatchPhraseFilter.prototype['match_phrase'] = undefined;






export default MatchPhraseFilter;

