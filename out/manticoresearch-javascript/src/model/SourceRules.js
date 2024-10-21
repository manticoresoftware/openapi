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
 * The SourceRules model module.
 * @module model/SourceRules
 * @version 5.0.0
 */
class SourceRules {
    /**
     * Constructs a new <code>SourceRules</code>.
     * Defines which fields to include or exclude in the response for a search query
     * @alias module:model/SourceRules
     * @extends Object
     */
    constructor() { 
        
        SourceRules.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SourceRules</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SourceRules} obj Optional instance to populate.
     * @return {module:model/SourceRules} The populated <code>SourceRules</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SourceRules();

            ApiClient.constructFromObject(data, obj, 'Object');
            

            if (data.hasOwnProperty('includes')) {
                obj['includes'] = ApiClient.convertToType(data['includes'], Object);
            }
            if (data.hasOwnProperty('excludes')) {
                obj['excludes'] = ApiClient.convertToType(data['excludes'], Object);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>SourceRules</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>SourceRules</code>.
     */
    static validateJSON(data) {

        return true;
    }


}



/**
 * List of fields to include in the response
 * @member {Object} includes
 */
SourceRules.prototype['includes'] = undefined;

/**
 * List of fields to exclude from the response
 * @member {Object} excludes
 */
SourceRules.prototype['excludes'] = undefined;






export default SourceRules;

