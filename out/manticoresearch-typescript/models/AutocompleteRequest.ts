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
* Object containing the data for performing an autocomplete search.
*/
export class AutocompleteRequest {
    /**
    * The table to perform the search on
    */
    'table': string;
    /**
    * The beginning of the string to autocomplete
    */
    'query': string;
    /**
    * Autocomplete options   - layouts: A comma-separated string of keyboard layout codes to validate and check for spell correction. Available options - us, ru, ua, se, pt, no, it, gr, uk, fr, es, dk, de, ch, br, bg, be. By default, all are enabled.   - fuzziness: (0,1 or 2) Maximum Levenshtein distance for finding typos. Set to 0 to disable fuzzy matching. Default is 2   - prepend: true/false If true, adds an asterisk before the last word for prefix expansion (e.g., *word )   - append:  true/false If true, adds an asterisk after the last word for prefix expansion (e.g., word* )   - expansion_len: Number of characters to expand in the last word. Default is 10. 
    */
    'options'?: any;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "table",
            "baseName": "table",
            "type": "string",
            "format": ""
        },
        {
            "name": "query",
            "baseName": "query",
            "type": "string",
            "format": ""
        },
        {
            "name": "options",
            "baseName": "options",
            "type": "any",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return AutocompleteRequest.attributeTypeMap;
    }

    public constructor() {
    }
}
