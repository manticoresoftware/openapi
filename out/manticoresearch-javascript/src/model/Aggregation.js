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
import AggComposite from './AggComposite';
import AggTerms from './AggTerms';

/**
 * The Aggregation model module.
 * @module model/Aggregation
 * @version 5.0.0
 */
class Aggregation {
    /**
     * Constructs a new <code>Aggregation</code>.
     * @alias module:model/Aggregation
     */
    constructor() { 
        
        Aggregation.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Aggregation</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Aggregation} obj Optional instance to populate.
     * @return {module:model/Aggregation} The populated <code>Aggregation</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Aggregation();

            if (data.hasOwnProperty('terms')) {
                obj['terms'] = AggTerms.constructFromObject(data['terms']);
            }
            if (data.hasOwnProperty('sort')) {
                obj['sort'] = ApiClient.convertToType(data['sort'], [Object]);
            }
            if (data.hasOwnProperty('composite')) {
                obj['composite'] = AggComposite.constructFromObject(data['composite']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>Aggregation</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>Aggregation</code>.
     */
    static validateJSON(data) {
        // validate the optional field `terms`
        if (data['terms']) { // data not null
          AggTerms.validateJSON(data['terms']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['sort'])) {
            throw new Error("Expected the field `sort` to be an array in the JSON data but got " + data['sort']);
        }
        // validate the optional field `composite`
        if (data['composite']) { // data not null
          AggComposite.validateJSON(data['composite']);
        }

        return true;
    }


}



/**
 * @member {module:model/AggTerms} terms
 */
Aggregation.prototype['terms'] = undefined;

/**
 * @member {Array.<Object>} sort
 */
Aggregation.prototype['sort'] = undefined;

/**
 * @member {module:model/AggComposite} composite
 */
Aggregation.prototype['composite'] = undefined;






export default Aggregation;
