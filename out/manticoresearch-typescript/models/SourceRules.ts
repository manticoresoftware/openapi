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
* Defines which fields to include or exclude in the response for a search query
*/
export class SourceRules {
    /**
    * List of fields to include in the response
    */
    'includes'?: any | null;
    /**
    * List of fields to exclude from the response
    */
    'excludes'?: any | null;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "includes",
            "baseName": "includes",
            "type": "any",
            "format": ""
        },
        {
            "name": "excludes",
            "baseName": "excludes",
            "type": "any",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return SourceRules.attributeTypeMap;
    }

    public constructor() {
    }
}
