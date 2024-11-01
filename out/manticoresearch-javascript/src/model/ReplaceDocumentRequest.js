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
 * The ReplaceDocumentRequest model module.
 * @module model/ReplaceDocumentRequest
 * @version 5.0.0
 */
class ReplaceDocumentRequest {
    /**
     * Constructs a new <code>ReplaceDocumentRequest</code>.
     * Object containing the document data for replacing an existing document in an index.
     * @alias module:model/ReplaceDocumentRequest
     * @param doc {Object} Object containing the new document data to replace the existing one.
     */
    constructor(doc) { 
        
        ReplaceDocumentRequest.initialize(this, doc);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, doc) { 
        obj['doc'] = doc;
    }

    /**
     * Constructs a <code>ReplaceDocumentRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ReplaceDocumentRequest} obj Optional instance to populate.
     * @return {module:model/ReplaceDocumentRequest} The populated <code>ReplaceDocumentRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ReplaceDocumentRequest();

            if (data.hasOwnProperty('doc')) {
                obj['doc'] = ApiClient.convertToType(data['doc'], Object);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ReplaceDocumentRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ReplaceDocumentRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of ReplaceDocumentRequest.RequiredProperties) {
            if (!data.hasOwnProperty(property)) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }

        return true;
    }


}

ReplaceDocumentRequest.RequiredProperties = ["doc"];

/**
 * Object containing the new document data to replace the existing one.
 * @member {Object} doc
 */
ReplaceDocumentRequest.prototype['doc'] = undefined;






export default ReplaceDocumentRequest;
