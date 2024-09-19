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
import QueryFilter from './QueryFilter';

/**
 * The BoolFilterBool model module.
 * @module model/BoolFilterBool
 * @version 4.0.0
 */
class BoolFilterBool {
    /**
     * Constructs a new <code>BoolFilterBool</code>.
     * @alias module:model/BoolFilterBool
     */
    constructor() { 
        
        BoolFilterBool.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BoolFilterBool</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BoolFilterBool} obj Optional instance to populate.
     * @return {module:model/BoolFilterBool} The populated <code>BoolFilterBool</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BoolFilterBool();

            if (data.hasOwnProperty('must')) {
                obj['must'] = ApiClient.convertToType(data['must'], [QueryFilter]);
            }
            if (data.hasOwnProperty('must_not')) {
                obj['must_not'] = ApiClient.convertToType(data['must_not'], [QueryFilter]);
            }
            if (data.hasOwnProperty('should')) {
                obj['should'] = ApiClient.convertToType(data['should'], [QueryFilter]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BoolFilterBool</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BoolFilterBool</code>.
     */
    static validateJSON(data) {
        if (data['must']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['must'])) {
                throw new Error("Expected the field `must` to be an array in the JSON data but got " + data['must']);
            }
            // validate the optional field `must` (array)
            for (const item of data['must']) {
                QueryFilter.validateJSON(item);
            };
        }
        if (data['must_not']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['must_not'])) {
                throw new Error("Expected the field `must_not` to be an array in the JSON data but got " + data['must_not']);
            }
            // validate the optional field `must_not` (array)
            for (const item of data['must_not']) {
                QueryFilter.validateJSON(item);
            };
        }
        if (data['should']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['should'])) {
                throw new Error("Expected the field `should` to be an array in the JSON data but got " + data['should']);
            }
            // validate the optional field `should` (array)
            for (const item of data['should']) {
                QueryFilter.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Array.<module:model/QueryFilter>} must
 */
BoolFilterBool.prototype['must'] = undefined;

/**
 * @member {Array.<module:model/QueryFilter>} must_not
 */
BoolFilterBool.prototype['must_not'] = undefined;

/**
 * @member {Array.<module:model/QueryFilter>} should
 */
BoolFilterBool.prototype['should'] = undefined;






export default BoolFilterBool;

