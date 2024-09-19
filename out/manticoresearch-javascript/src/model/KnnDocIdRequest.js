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
import KnnSearchParameters from './KnnSearchParameters';

/**
 * The KnnDocIdRequest model module.
 * @module model/KnnDocIdRequest
 * @version 4.0.0
 */
class KnnDocIdRequest {
    /**
     * Constructs a new <code>KnnDocIdRequest</code>.
     * @alias module:model/KnnDocIdRequest
     * @implements module:model/KnnSearchParameters
     * @param field {String} 
     * @param docId {Number} 
     */
    constructor(field, docId) { 
        KnnSearchParameters.initialize(this, field);
        KnnDocIdRequest.initialize(this, field, docId);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, field, docId) { 
        obj['field'] = field;
        obj['doc_id'] = docId;
    }

    /**
     * Constructs a <code>KnnDocIdRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/KnnDocIdRequest} obj Optional instance to populate.
     * @return {module:model/KnnDocIdRequest} The populated <code>KnnDocIdRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new KnnDocIdRequest();
            KnnSearchParameters.constructFromObject(data, obj);

            if (data.hasOwnProperty('field')) {
                obj['field'] = ApiClient.convertToType(data['field'], 'String');
            }
            if (data.hasOwnProperty('k')) {
                obj['k'] = ApiClient.convertToType(data['k'], 'Number');
            }
            if (data.hasOwnProperty('ef')) {
                obj['ef'] = ApiClient.convertToType(data['ef'], 'Number');
            }
            if (data.hasOwnProperty('doc_id')) {
                obj['doc_id'] = ApiClient.convertToType(data['doc_id'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>KnnDocIdRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>KnnDocIdRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of KnnDocIdRequest.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['field'] && !(typeof data['field'] === 'string' || data['field'] instanceof String)) {
            throw new Error("Expected the field `field` to be a primitive type in the JSON string but got " + data['field']);
        }

        return true;
    }


}

KnnDocIdRequest.RequiredProperties = ["field", "doc_id"];

/**
 * @member {String} field
 */
KnnDocIdRequest.prototype['field'] = undefined;

/**
 * @member {Number} k
 */
KnnDocIdRequest.prototype['k'] = undefined;

/**
 * @member {Number} ef
 */
KnnDocIdRequest.prototype['ef'] = undefined;

/**
 * @member {Number} doc_id
 */
KnnDocIdRequest.prototype['doc_id'] = undefined;


// Implement KnnSearchParameters interface:
/**
 * @member {String} field
 */
KnnSearchParameters.prototype['field'] = undefined;
/**
 * @member {Number} k
 */
KnnSearchParameters.prototype['k'] = undefined;
/**
 * @member {Number} ef
 */
KnnSearchParameters.prototype['ef'] = undefined;




export default KnnDocIdRequest;

