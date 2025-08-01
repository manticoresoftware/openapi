/*
 * Manticore Search Client
 *
 * Сlient for Manticore Search. 
 *
 * The version of the OpenAPI document: 5.0.0
 * Contact: info@manticoresearch.com
 * Generated by: https://openapi-generator.tech
 */

use std::sync::Arc;
use std::borrow::Borrow;
use std::pin::Pin;
#[allow(unused_imports)]
use std::option::Option;

use hyper;
use hyper_util::client::legacy::connect::Connect;
use futures::Future;

use crate::models;
use super::{Error, configuration};
use super::request as __internal_request;

pub struct SearchApiClient<C: Connect>
    where C: Clone + std::marker::Send + Sync + 'static {
    configuration: Arc<configuration::Configuration<C>>,
}

impl<C: Connect> SearchApiClient<C>
    where C: Clone + std::marker::Send + Sync {
    pub fn new(configuration: Arc<configuration::Configuration<C>>) -> SearchApiClient<C> {
        SearchApiClient {
            configuration,
        }
    }
}

pub trait SearchApi: Send + Sync {
    fn autocomplete(&self, autocomplete_request: models::AutocompleteRequest) -> Pin<Box<dyn Future<Output = Result<Vec<serde_json::Value>, Error>> + Send>>;
    fn percolate(&self, table: &str, percolate_request: models::PercolateRequest) -> Pin<Box<dyn Future<Output = Result<models::SearchResponse, Error>> + Send>>;
    fn search(&self, search_request: models::SearchRequest) -> Pin<Box<dyn Future<Output = Result<models::SearchResponse, Error>> + Send>>;
}

impl<C: Connect>SearchApi for SearchApiClient<C>
    where C: Clone + std::marker::Send + Sync {
    #[allow(unused_mut)]
    fn autocomplete(&self, autocomplete_request: models::AutocompleteRequest) -> Pin<Box<dyn Future<Output = Result<Vec<serde_json::Value>, Error>> + Send>> {
        let mut req = __internal_request::Request::new(hyper::Method::POST, "/autocomplete".to_string())
        ;
        req = req.with_body_param(autocomplete_request);

        req.execute(self.configuration.borrow())
    }

    #[allow(unused_mut)]
    fn percolate(&self, table: &str, percolate_request: models::PercolateRequest) -> Pin<Box<dyn Future<Output = Result<models::SearchResponse, Error>> + Send>> {
        let mut req = __internal_request::Request::new(hyper::Method::POST, "/pq/{table}/search".to_string())
        ;
        req = req.with_path_param("table".to_string(), table.to_string());
        req = req.with_body_param(percolate_request);

        req.execute(self.configuration.borrow())
    }

    #[allow(unused_mut)]
    fn search(&self, search_request: models::SearchRequest) -> Pin<Box<dyn Future<Output = Result<models::SearchResponse, Error>> + Send>> {
        let mut req = __internal_request::Request::new(hyper::Method::POST, "/search".to_string())
        ;
        req = req.with_body_param(search_request);

        req.execute(self.configuration.borrow())
    }

}
