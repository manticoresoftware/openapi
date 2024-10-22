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
 * The MatchAll model module.
 * @module model/MatchAll
 * @version 5.0.0
 */
class MatchAll {
    /**
     * Constructs a new <code>MatchAll</code>.
     * Filter helper object defining the &#39;match all&#39;&#39; condition
     * @alias module:model/MatchAll
     * @param all {module:model/MatchAll.AllEnum} 
     */
    constructor(all) { 
        
        MatchAll.initialize(this, all);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, all) { 
        obj['_all'] = all;
    }

    /**
     * Constructs a <code>MatchAll</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MatchAll} obj Optional instance to populate.
     * @return {module:model/MatchAll} The populated <code>MatchAll</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new MatchAll();

            if (data.hasOwnProperty('_all')) {
                obj['_all'] = ApiClient.convertToType(data['_all'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>MatchAll</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>MatchAll</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of MatchAll.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['_all'] && !(typeof data['_all'] === 'string' || data['_all'] instanceof String)) {
            throw new Error("Expected the field `_all` to be a primitive type in the JSON string but got " + data['_all']);
        }

        return true;
    }


}

MatchAll.RequiredProperties = ["_all"];

/**
 * @member {module:model/MatchAll.AllEnum} _all
 */
MatchAll.prototype['_all'] = undefined;





/**
 * Allowed values for the <code>_all</code> property.
 * @enum {String}
 * @readonly
 */
MatchAll['AllEnum'] = {

    /**
     * value: "{}"
     * @const
     */
    "{}": "{}"
};



export default MatchAll;

