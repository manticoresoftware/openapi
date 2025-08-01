/*
 * Manticore Search Client
 *
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 5.0.0
 * Contact: info@manticoresearch.com
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// HitsHits : Search hit representing a matched document
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct HitsHits {
    /// The ID of the matched document
    #[serde(rename = "_id", skip_serializing_if = "Option::is_none")]
    pub _id: Option<i32>,
    /// The score of the matched document
    #[serde(rename = "_score", skip_serializing_if = "Option::is_none")]
    pub _score: Option<i32>,
    /// The source data of the matched document
    #[serde(rename = "_source", skip_serializing_if = "Option::is_none")]
    pub _source: Option<serde_json::Value>,
    /// The knn distance of the matched document returned for knn queries
    #[serde(rename = "_knn_dist", skip_serializing_if = "Option::is_none")]
    pub _knn_dist: Option<f64>,
    /// The highlighting-related data of the matched document
    #[serde(rename = "highlight", skip_serializing_if = "Option::is_none")]
    pub highlight: Option<serde_json::Value>,
    /// The table name of the matched document returned for percolate queries
    #[serde(rename = "table", skip_serializing_if = "Option::is_none")]
    pub table: Option<String>,
    /// The type of the matched document returned for percolate queries
    #[serde(rename = "_type:", skip_serializing_if = "Option::is_none")]
    pub _type_colon: Option<String>,
    /// The percolate-related fields of the matched document returned for percolate queries
    #[serde(rename = "fields", skip_serializing_if = "Option::is_none")]
    pub fields: Option<serde_json::Value>,
}

impl HitsHits {
    /// Search hit representing a matched document
    pub fn new() -> HitsHits {
        HitsHits {
            _id: None,
            _score: None,
            _source: None,
            _knn_dist: None,
            highlight: None,
            table: None,
            _type_colon: None,
            fields: None,
        }
    }
}

