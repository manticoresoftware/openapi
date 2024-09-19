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
import KnnDocIdRequest from './KnnDocIdRequest';
import KnnQueryVectorRequest from './KnnQueryVectorRequest';

/**
 * The KnnSearchRequestAllOfKnn model module.
 * @module model/KnnSearchRequestAllOfKnn
 * @version 4.0.0
 */
class KnnSearchRequestAllOfKnn {
    /**
     * Constructs a new <code>KnnSearchRequestAllOfKnn</code>.
     * @alias module:model/KnnSearchRequestAllOfKnn
     * @param {(module:model/KnnDocIdRequest|module:model/KnnQueryVectorRequest)} instance The actual instance to initialize KnnSearchRequestAllOfKnn.
     */
    constructor(instance = null) {
        if (instance === null) {
            this.actualInstance = null;
            return;
        }
        var match = 0;
        var errorMessages = [];
        try {
            if (typeof instance === "KnnDocIdRequest") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                KnnDocIdRequest.validateJSON(instance); // throw an exception if no match
                // create KnnDocIdRequest from JS object
                this.actualInstance = KnnDocIdRequest.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into KnnDocIdRequest
            errorMessages.push("Failed to construct KnnDocIdRequest: " + err)
        }

        try {
            if (typeof instance === "KnnQueryVectorRequest") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                KnnQueryVectorRequest.validateJSON(instance); // throw an exception if no match
                // create KnnQueryVectorRequest from JS object
                this.actualInstance = KnnQueryVectorRequest.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into KnnQueryVectorRequest
            errorMessages.push("Failed to construct KnnQueryVectorRequest: " + err)
        }

        if (match > 1) {
            throw new Error("Multiple matches found constructing `KnnSearchRequestAllOfKnn` with oneOf schemas KnnDocIdRequest, KnnQueryVectorRequest. Input: " + JSON.stringify(instance));
        } else if (match === 0) {
            this.actualInstance = null; // clear the actual instance in case there are multiple matches
            throw new Error("No match found constructing `KnnSearchRequestAllOfKnn` with oneOf schemas KnnDocIdRequest, KnnQueryVectorRequest. Details: " +
                            errorMessages.join(", "));
        } else { // only 1 match
            // the input is valid
        }
    }

    /**
     * Constructs a <code>KnnSearchRequestAllOfKnn</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/KnnSearchRequestAllOfKnn} obj Optional instance to populate.
     * @return {module:model/KnnSearchRequestAllOfKnn} The populated <code>KnnSearchRequestAllOfKnn</code> instance.
     */
    static constructFromObject(data, obj) {
        return new KnnSearchRequestAllOfKnn(data);
    }

    /**
     * Gets the actual instance, which can be <code>KnnDocIdRequest</code>, <code>KnnQueryVectorRequest</code>.
     * @return {(module:model/KnnDocIdRequest|module:model/KnnQueryVectorRequest)} The actual instance.
     */
    getActualInstance() {
        return this.actualInstance;
    }

    /**
     * Sets the actual instance, which can be <code>KnnDocIdRequest</code>, <code>KnnQueryVectorRequest</code>.
     * @param {(module:model/KnnDocIdRequest|module:model/KnnQueryVectorRequest)} obj The actual instance.
     */
    setActualInstance(obj) {
       this.actualInstance = KnnSearchRequestAllOfKnn.constructFromObject(obj).getActualInstance();
    }

    /**
     * Returns the JSON representation of the actual instance.
     * @return {string}
     */
    toJSON = function(){
        return this.getActualInstance();
    }

    /**
     * Create an instance of KnnSearchRequestAllOfKnn from a JSON string.
     * @param {string} json_string JSON string.
     * @return {module:model/KnnSearchRequestAllOfKnn} An instance of KnnSearchRequestAllOfKnn.
     */
    static fromJSON = function(json_string){
        return KnnSearchRequestAllOfKnn.constructFromObject(JSON.parse(json_string));
    }
}

/**
 * @member {Number} doc_id
 */
KnnSearchRequestAllOfKnn.prototype['doc_id'] = undefined;

/**
 * @member {String} field
 */
KnnSearchRequestAllOfKnn.prototype['field'] = undefined;

/**
 * @member {Number} k
 */
KnnSearchRequestAllOfKnn.prototype['k'] = undefined;

/**
 * @member {Number} ef
 */
KnnSearchRequestAllOfKnn.prototype['ef'] = undefined;

/**
 * @member {Array.<Number>} query_vector
 */
KnnSearchRequestAllOfKnn.prototype['query_vector'] = undefined;


KnnSearchRequestAllOfKnn.OneOf = ["KnnDocIdRequest", "KnnQueryVectorRequest"];

export default KnnSearchRequestAllOfKnn;

