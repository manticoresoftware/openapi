# manticoresearch

Low-level client for Manticore Search.


## Installation

```shell
npm install manticoresearch-ts
```
## Requirements

Node v12.20.0.

Minimum Manticore Search version is > 4.2.0 with HTTP protocol enabled.

## Documentation

Full documentation is available in  [docs](https://github.com/manticoresoftware/manticoresearch-typescript/tree/master/docs) folder.

Manticore Search server documentation: https://manual.manticoresearch.com.

## Getting Started

A simple search case:

```javascript
import { SearchApi, SortOrderOrderEnum } from 'manticoresearch-ts'

const searchApi = new SearchApi()
searchApi
  .search({
    index: 'forum',
    query: { match_all: {}, bool: { must: [{ equals: { author_id: 123 } }, { in: { forum_id: [1, 3, 7] } }] } },
    sort: [{ post_date: SortOrderOrderEnum.desc }],
  })
  .then((res) => console.log(JSON.stringify(res)))

```

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:9308*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Manticoresearch.IndexApi* | [**bulk**](docs/IndexApi.md#bulk) | **POST** /bulk | Bulk index operations
*Manticoresearch.IndexApi* | [**delete**](docs/IndexApi.md#delete) | **POST** /delete | Delete a document in an index
*Manticoresearch.IndexApi* | [**insert**](docs/IndexApi.md#insert) | **POST** /insert | Create a new document in an index
*Manticoresearch.IndexApi* | [**replace**](docs/IndexApi.md#replace) | **POST** /replace | Replace new document in an index
*Manticoresearch.IndexApi* | [**update**](docs/IndexApi.md#update) | **POST** /update | Update a document in an index
*Manticoresearch.SearchApi* | [**percolate**](docs/SearchApi.md#percolate) | **POST** /pq/{index}/search | Perform reverse search on a percolate index
*Manticoresearch.SearchApi* | [**search**](docs/SearchApi.md#search) | **POST** /search | Performs a search
*Manticoresearch.UtilsApi* | [**sql**](docs/UtilsApi.md#sql) | **POST** /sql | Perform SQL requests


## Documentation for Models

 - [Manticoresearch.BulkResponse](docs/BulkResponse.md)
 - [Manticoresearch.DeleteDocumentRequest](docs/DeleteDocumentRequest.md)
 - [Manticoresearch.DeleteResponse](docs/DeleteResponse.md)
 - [Manticoresearch.ErrorResponse](docs/ErrorResponse.md)
 - [Manticoresearch.InsertDocumentRequest](docs/InsertDocumentRequest.md)
 - [Manticoresearch.PercolateRequest](docs/PercolateRequest.md)
 - [Manticoresearch.SearchRequest](docs/SearchRequest.md)
 - [Manticoresearch.SearchResponse](docs/SearchResponse.md)
 - [Manticoresearch.SearchResponseHits](docs/SearchResponseHits.md)
 - [Manticoresearch.SuccessResponse](docs/SuccessResponse.md)
 - [Manticoresearch.UpdateDocumentRequest](docs/UpdateDocumentRequest.md)
 - [Manticoresearch.UpdateResponse](docs/UpdateResponse.md)


## Documentation for Authorization

All endpoints do not require authorization.
