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
 * The SqlResponse model module.
 * @module model/SqlResponse
 * @version 7.0.0
 */
class SqlResponse {
    /**
     * Constructs a new <code>SqlResponse</code>.
     * List of responses from executed SQL queries
     * @alias module:model/SqlResponse
     * @param {(module:model/Object|module:model/[Object])} instance The actual instance to initialize SqlResponse.
     */
    constructor(instance = null) {
        if (instance === null) {
            this.actualInstance = null;
            return;
        }
        var match = 0;
        var errorMessages = [];
        try {
            // validate array data type
            if (!Array.isArray(instance)) {
                throw new Error("Invalid data type. Expecting array. Input: " + instance);
            }
            this.actualInstance = instance;
            match++;
        } catch(err) {
            // json data failed to deserialize into [Object]
            errorMessages.push("Failed to construct [Object]: " + err)
        }

        try {
            this.actualInstance = instance;
            match++;
        } catch(err) {
            // json data failed to deserialize into Object
            errorMessages.push("Failed to construct Object: " + err)
        }

        if (match > 1) {
            throw new Error("Multiple matches found constructing `SqlResponse` with oneOf schemas Object, [Object]. Input: " + JSON.stringify(instance));
        } else if (match === 0) {
            this.actualInstance = null; // clear the actual instance in case there are multiple matches
            throw new Error("No match found constructing `SqlResponse` with oneOf schemas Object, [Object]. Details: " +
                            errorMessages.join(", "));
        } else { // only 1 match
            // the input is valid
        }
    }

    /**
     * Constructs a <code>SqlResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SqlResponse} obj Optional instance to populate.
     * @return {module:model/SqlResponse} The populated <code>SqlResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        return new SqlResponse(data);
    }

    /**
     * Gets the actual instance, which can be <code>Object</code>, <code>[Object]</code>.
     * @return {(module:model/Object|module:model/[Object])} The actual instance.
     */
    getActualInstance() {
        return this.actualInstance;
    }

    /**
     * Sets the actual instance, which can be <code>Object</code>, <code>[Object]</code>.
     * @param {(module:model/Object|module:model/[Object])} obj The actual instance.
     */
    setActualInstance(obj) {
       this.actualInstance = SqlResponse.constructFromObject(obj).getActualInstance();
    }

    /**
     * Returns the JSON representation of the actual instance.
     * @return {string}
     */
    toJSON = function(){
        return this.getActualInstance();
    }

    /**
     * Create an instance of SqlResponse from a JSON string.
     * @param {string} json_string JSON string.
     * @return {module:model/SqlResponse} An instance of SqlResponse.
     */
    static fromJSON = function(json_string){
        return SqlResponse.constructFromObject(JSON.parse(json_string));
    }
}


SqlResponse.OneOf = ["Object", "[Object]"];

export default SqlResponse;

