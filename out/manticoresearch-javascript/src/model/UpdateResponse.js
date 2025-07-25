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
 * The UpdateResponse model module.
 * @module model/UpdateResponse
 * @version 8.1.0
 */
class UpdateResponse {
    /**
     * Constructs a new <code>UpdateResponse</code>.
     * Success response returned after updating one or more documents
     * @alias module:model/UpdateResponse
     */
    constructor() { 
        
        UpdateResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UpdateResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UpdateResponse} obj Optional instance to populate.
     * @return {module:model/UpdateResponse} The populated <code>UpdateResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UpdateResponse();

            if (data.hasOwnProperty('table')) {
                obj['table'] = ApiClient.convertToType(data['table'], 'String');
            }
            if (data.hasOwnProperty('updated')) {
                obj['updated'] = ApiClient.convertToType(data['updated'], 'Number');
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('result')) {
                obj['result'] = ApiClient.convertToType(data['result'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>UpdateResponse</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>UpdateResponse</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['table'] && !(typeof data['table'] === 'string' || data['table'] instanceof String)) {
            throw new Error("Expected the field `table` to be a primitive type in the JSON string but got " + data['table']);
        }
        // ensure the json data is a string
        if (data['result'] && !(typeof data['result'] === 'string' || data['result'] instanceof String)) {
            throw new Error("Expected the field `result` to be a primitive type in the JSON string but got " + data['result']);
        }

        return true;
    }


}



/**
 * Name of the document table
 * @member {String} table
 */
UpdateResponse.prototype['table'] = undefined;

/**
 * Number of documents updated
 * @member {Number} updated
 */
UpdateResponse.prototype['updated'] = undefined;

/**
 * Document ID
 * @member {Number} id
 */
UpdateResponse.prototype['id'] = undefined;

/**
 * Result of the update operation, typically 'updated'
 * @member {String} result
 */
UpdateResponse.prototype['result'] = undefined;






export default UpdateResponse;

