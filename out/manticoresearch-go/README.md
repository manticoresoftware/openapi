# Manticorer Go client

Сlient for Manticore Search.


## Installation

```shell

go get github.com/manticoresoftware/manticoresearch-go@v1.9.0

```

## Getting Started

go mod init main
go get github.com/manticoresoftware/manticoresearch-go@v1.9.0

```go

package main

import (
	"context"
	"fmt"
	Manticoresearch "github.com/manticoresoftware/manticoresearch-go"
)

func main() {

	// Create an instance of API client
	configuration := Manticoresearch.NewConfiguration()
	configuration.Servers[0].URL = "http://localhost:9308"
	apiClient := Manticoresearch.NewAPIClient(configuration)
	
	// Perform insert and search operations
	tableName := "products"
	indexDoc := map[string]interface{} {"title": "Crossbody Bag with Tassel"}
	indexReq := Manticoresearch.NewInsertDocumentRequest(tableName, indexDoc)
	indexReq.SetId(1)
	
	apiClient.IndexAPI.Insert(context.Background()).InsertDocumentRequest(*indexReq).Execute();
	
	searchRequest := Manticoresearch.NewSearchRequest(tableName)
	searchQuery := Manticoresearch.NewSearchQuery()
	searchQuery.QueryString = "@title Bag"
	searchRequest.Query = searchQuery
	queryHighlight := Manticoresearch.NewHighlight()
	queryHighlight.Fields =  map[string]interface{} {"title": map[string]interface{} {}}
	searchRequest.Highlight = queryHighlight      
	
	_, httpRes, _ := apiClient.SearchAPI.Search(context.Background()).SearchRequest(*searchRequest).Execute()
	fmt.Printf("%+v\n\n", httpRes)
}
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### Using proxy

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

More details on the use of `HTTP_PROXY` can be found [here](https://www.cyberciti.biz/faq/linux-unix-set-proxy-environment-variable/)

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:9308*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndexAPI* | [**Bulk**](docs/IndexAPI.md#bulk) | **Post** /bulk | Bulk table operations
*IndexAPI* | [**Delete**](docs/IndexAPI.md#delete) | **Post** /delete | Delete a document in a table
*IndexAPI* | [**Insert**](docs/IndexAPI.md#insert) | **Post** /insert | Create a new document in a table
*IndexAPI* | [**PartialReplace**](docs/IndexAPI.md#partialreplace) | **Post** /{table}/_update/{id} | Partially replaces a document in a table
*IndexAPI* | [**Replace**](docs/IndexAPI.md#replace) | **Post** /replace | Replace new document in a table
*IndexAPI* | [**Update**](docs/IndexAPI.md#update) | **Post** /update | Update a document in a table
*SearchAPI* | [**Autocomplete**](docs/SearchAPI.md#autocomplete) | **Post** /autocomplete | Performs an autocomplete search on a table
*SearchAPI* | [**Percolate**](docs/SearchAPI.md#percolate) | **Post** /pq/{table}/search | Perform reverse search on a percolate table
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Post** /search | Performs a search on a table
*UtilsAPI* | [**Sql**](docs/UtilsAPI.md#sql) | **Post** /sql | Perform SQL requests


## Documentation For Models

 - [AggComposite](docs/AggComposite.md)
 - [AggCompositeSource](docs/AggCompositeSource.md)
 - [AggCompositeTerm](docs/AggCompositeTerm.md)
 - [AggDateHistogram](docs/AggDateHistogram.md)
 - [AggHistogram](docs/AggHistogram.md)
 - [AggTerms](docs/AggTerms.md)
 - [Aggregation](docs/Aggregation.md)
 - [AutocompleteRequest](docs/AutocompleteRequest.md)
 - [BoolFilter](docs/BoolFilter.md)
 - [BulkResponse](docs/BulkResponse.md)
 - [DeleteDocumentRequest](docs/DeleteDocumentRequest.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [FulltextFilter](docs/FulltextFilter.md)
 - [GeoDistance](docs/GeoDistance.md)
 - [GeoDistanceLocationAnchor](docs/GeoDistanceLocationAnchor.md)
 - [Highlight](docs/Highlight.md)
 - [HighlightFieldOption](docs/HighlightFieldOption.md)
 - [HighlightFields](docs/HighlightFields.md)
 - [HitsHits](docs/HitsHits.md)
 - [InsertDocumentRequest](docs/InsertDocumentRequest.md)
 - [Join](docs/Join.md)
 - [JoinCond](docs/JoinCond.md)
 - [JoinOn](docs/JoinOn.md)
 - [KnnQuery](docs/KnnQuery.md)
 - [Match](docs/Match.md)
 - [MatchAll](docs/MatchAll.md)
 - [ModelRange](docs/ModelRange.md)
 - [PercolateRequest](docs/PercolateRequest.md)
 - [PercolateRequestQuery](docs/PercolateRequestQuery.md)
 - [QueryFilter](docs/QueryFilter.md)
 - [ReplaceDocumentRequest](docs/ReplaceDocumentRequest.md)
 - [ResponseError](docs/ResponseError.md)
 - [ResponseErrorDetails](docs/ResponseErrorDetails.md)
 - [SearchQuery](docs/SearchQuery.md)
 - [SearchRequest](docs/SearchRequest.md)
 - [SearchResponse](docs/SearchResponse.md)
 - [SearchResponseHits](docs/SearchResponseHits.md)
 - [SourceRules](docs/SourceRules.md)
 - [SqlObjResponse](docs/SqlObjResponse.md)
 - [SqlResponse](docs/SqlResponse.md)
 - [SuccessResponse](docs/SuccessResponse.md)
 - [UpdateDocumentRequest](docs/UpdateDocumentRequest.md)
 - [UpdateResponse](docs/UpdateResponse.md)


## Documentation For Authorization

Endpoints do not require authorization.

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
