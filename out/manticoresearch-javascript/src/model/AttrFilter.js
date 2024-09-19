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
import EqualsFilter from './EqualsFilter';
import EqualsFilterEquals from './EqualsFilterEquals';
import GeoFilter from './GeoFilter';
import GeoFilterGeoDistance from './GeoFilterGeoDistance';
import InFilter from './InFilter';
import RangeFilter from './RangeFilter';
import RangeFilterRangeValue from './RangeFilterRangeValue';

/**
 * The AttrFilter model module.
 * @module model/AttrFilter
 * @version 4.0.0
 */
class AttrFilter {
    /**
     * Constructs a new <code>AttrFilter</code>.
     * @alias module:model/AttrFilter
     * @param {(module:model/EqualsFilter|module:model/GeoFilter|module:model/InFilter|module:model/RangeFilter)} instance The actual instance to initialize AttrFilter.
     */
    constructor(instance = null) {
        if (instance === null) {
            this.actualInstance = null;
            return;
        }
        var match = 0;
        var errorMessages = [];
        try {
            if (typeof instance === "EqualsFilter") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                EqualsFilter.validateJSON(instance); // throw an exception if no match
                // create EqualsFilter from JS object
                this.actualInstance = EqualsFilter.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into EqualsFilter
            errorMessages.push("Failed to construct EqualsFilter: " + err)
        }

        try {
            if (typeof instance === "InFilter") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                InFilter.validateJSON(instance); // throw an exception if no match
                // create InFilter from JS object
                this.actualInstance = InFilter.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into InFilter
            errorMessages.push("Failed to construct InFilter: " + err)
        }

        try {
            if (typeof instance === "RangeFilter") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                RangeFilter.validateJSON(instance); // throw an exception if no match
                // create RangeFilter from JS object
                this.actualInstance = RangeFilter.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into RangeFilter
            errorMessages.push("Failed to construct RangeFilter: " + err)
        }

        try {
            if (typeof instance === "GeoFilter") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                GeoFilter.validateJSON(instance); // throw an exception if no match
                // create GeoFilter from JS object
                this.actualInstance = GeoFilter.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into GeoFilter
            errorMessages.push("Failed to construct GeoFilter: " + err)
        }

        if (match > 1) {
            throw new Error("Multiple matches found constructing `AttrFilter` with oneOf schemas EqualsFilter, GeoFilter, InFilter, RangeFilter. Input: " + JSON.stringify(instance));
        } else if (match === 0) {
            this.actualInstance = null; // clear the actual instance in case there are multiple matches
            throw new Error("No match found constructing `AttrFilter` with oneOf schemas EqualsFilter, GeoFilter, InFilter, RangeFilter. Details: " +
                            errorMessages.join(", "));
        } else { // only 1 match
            // the input is valid
        }
    }

    /**
     * Constructs a <code>AttrFilter</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AttrFilter} obj Optional instance to populate.
     * @return {module:model/AttrFilter} The populated <code>AttrFilter</code> instance.
     */
    static constructFromObject(data, obj) {
        return new AttrFilter(data);
    }

    /**
     * Gets the actual instance, which can be <code>EqualsFilter</code>, <code>GeoFilter</code>, <code>InFilter</code>, <code>RangeFilter</code>.
     * @return {(module:model/EqualsFilter|module:model/GeoFilter|module:model/InFilter|module:model/RangeFilter)} The actual instance.
     */
    getActualInstance() {
        return this.actualInstance;
    }

    /**
     * Sets the actual instance, which can be <code>EqualsFilter</code>, <code>GeoFilter</code>, <code>InFilter</code>, <code>RangeFilter</code>.
     * @param {(module:model/EqualsFilter|module:model/GeoFilter|module:model/InFilter|module:model/RangeFilter)} obj The actual instance.
     */
    setActualInstance(obj) {
       this.actualInstance = AttrFilter.constructFromObject(obj).getActualInstance();
    }

    /**
     * Returns the JSON representation of the actual instance.
     * @return {string}
     */
    toJSON = function(){
        return this.getActualInstance();
    }

    /**
     * Create an instance of AttrFilter from a JSON string.
     * @param {string} json_string JSON string.
     * @return {module:model/AttrFilter} An instance of AttrFilter.
     */
    static fromJSON = function(json_string){
        return AttrFilter.constructFromObject(JSON.parse(json_string));
    }
}

/**
 * @member {module:model/EqualsFilterEquals} equals
 */
AttrFilter.prototype['equals'] = undefined;

/**
 * @member {Object.<String, Array.<module:model/EqualsFilterEquals>>} in
 */
AttrFilter.prototype['in'] = undefined;

/**
 * @member {Object.<String, module:model/RangeFilterRangeValue>} range
 */
AttrFilter.prototype['range'] = undefined;

/**
 * @member {module:model/GeoFilterGeoDistance} geo_distance
 */
AttrFilter.prototype['geo_distance'] = undefined;


AttrFilter.OneOf = ["EqualsFilter", "GeoFilter", "InFilter", "RangeFilter"];

export default AttrFilter;

