/**
 * Manticore Search Client
 * Сlient for Manticore Search. 
 *
 * OpenAPI spec version: 5.0.0
 * Contact: info@manticoresearch.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

/**
* Filter helper object defining the \'range\' condition
*/
export class Range {
    'lt'?: any | null;
    'lte'?: any | null;
    'gt'?: any | null;
    'gte'?: any | null;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "lt",
            "baseName": "lt",
            "type": "any",
            "format": ""
        },
        {
            "name": "lte",
            "baseName": "lte",
            "type": "any",
            "format": ""
        },
        {
            "name": "gt",
            "baseName": "gt",
            "type": "any",
            "format": ""
        },
        {
            "name": "gte",
            "baseName": "gte",
            "type": "any",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Range.attributeTypeMap;
    }

    public constructor() {
    }
}
