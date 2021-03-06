# manticoresearch


Low-level client for Manticore Search.


## Requirements

Building the API client library requires:

1. Java 1.8+
2. Maven/Gradle

Minimum Manticore Search version is 2.5.1 with HTTP protocol enabled.

| Manticore Search  | manticoresearch-java   |     Java      |
| ----------------- | ---------------------- | ------------- |
| >= 4.2.1          | 3.0.x                  | >= 1.8        |
| >= 4.0.2          | 2.0.3                  | >= 1.8        |
| >= 2.5.1          | 2.0.2                  | >= 1.8        |


## Installation

Add this dependency to your project's POM:

```xml
<dependency>
  <groupId>com.manticoresearch</groupId>
  <artifactId>manticoresearch</artifactId>
  <version>3.2.0</version>
</dependency>
```

### Gradle users

Add this dependency to your project's build file:

```groovy
compile "com.manticoresearch:manticoresearch:3.2.0"
```

### Others

At first generate the JAR by executing:

```shell
mvn clean package
```

Then manually install the following JARs:

- `target/manticoresearch-3.2.0.jar`
- `target/lib/*.jar`

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```java

import com.manticoresearch.client.*;
import com.manticoresearch.client.auth.*;
import com.manticoresearch.client.model.*;
import com.manticoresearch.client.api.IndexApi;

public class IndexApiExample {

    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://127.0.0.1:9308");
        
        IndexApi apiInstance = new IndexApi(defaultClient);
        String body = ["'{\"insert\": {\"index\": \"test\", \"id\": 1, \"doc\": {\"title\": \"Title 1\"}}},\\n{\"insert\": {\"index\": \"test\", \"id\": 2, \"doc\": {\"title\": \"Title 2\"}}}'"]; // String | 
        try {
            BulkResponse result = apiInstance.bulk(body);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling IndexApi#bulk");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}

```

## Documentation 

Full documentation is available in  [docs](https://github.com/manticoresoftware/manticoresearch-java/tree/master/docs) folder.

Manticore Search server documentation: https://manual.manticoresearch.com.

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:9308*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndexApi* | [**bulk**](docs/IndexApi.md#bulk) | **POST** /json/bulk | Bulk index operations
*IndexApi* | [**delete**](docs/IndexApi.md#delete) | **POST** /json/delete | Delete a document in an index
*IndexApi* | [**insert**](docs/IndexApi.md#insert) | **POST** /json/insert | Create a new document in an index
*IndexApi* | [**replace**](docs/IndexApi.md#replace) | **POST** /json/replace | Replace new document in an index
*IndexApi* | [**update**](docs/IndexApi.md#update) | **POST** /json/update | Update a document in an index
*SearchApi* | [**percolate**](docs/SearchApi.md#percolate) | **POST** /json/pq/{index}/search | Perform reverse search on a percolate index
*SearchApi* | [**search**](docs/SearchApi.md#search) | **POST** /json/search | Performs a search
*UtilsApi* | [**sql**](docs/UtilsApi.md#sql) | **POST** /sql | Perform SQL requests


## Documentation for Models

 - [BulkResponse](docs/BulkResponse.md)
 - [DeleteDocumentRequest](docs/DeleteDocumentRequest.md)
 - [DeleteResponse](docs/DeleteResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [InsertDocumentRequest](docs/InsertDocumentRequest.md)
 - [PercolateRequest](docs/PercolateRequest.md)
 - [PercolateRequestQuery](docs/PercolateRequestQuery.md)
 - [SearchRequest](docs/SearchRequest.md)
 - [SearchResponse](docs/SearchResponse.md)
 - [SearchResponseHits](docs/SearchResponseHits.md)
 - [SuccessResponse](docs/SuccessResponse.md)
 - [UpdateDocumentRequest](docs/UpdateDocumentRequest.md)
 - [UpdateResponse](docs/UpdateResponse.md)


## Documentation for Authorization

All endpoints do not require authorization.
Authentication schemes defined for the API:

## Recommendation

It's recommended to create an instance of `ApiClient` per thread in a multithreaded environment to avoid any potential issues.

## Author

info@manticoresearch.com

