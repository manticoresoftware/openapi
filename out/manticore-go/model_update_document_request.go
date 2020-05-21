/*
 * Manticore Search API
 *
 * This is the API for Manticore Search HTTP protocol 
 *
 * API version: 1.0.0
 * Contact: adrian.nuta@manticoresearch.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateDocumentRequest struct for UpdateDocumentRequest
type UpdateDocumentRequest struct {
	Index string `json:"index"`
	Doc map[string]interface{} `json:"doc"`
	Id int64 `json:"id,omitempty"`
	Query map[string]interface{} `json:"query,omitempty"`
}