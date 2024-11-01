# Manticore .Net client

❗ WARNING: this is a development version of the client. The latest release's readme is https://github.com/manticoresoftware/manticoresearch-net/tree/5.0.0

- API version: 5.0.0
- Build package: org.openapitools.codegen.languages.CSharpClientCodegen
    For more information, please visit [https://manticoresearch.com/contact-us/](https://manticoresearch.com/contact-us/)

<a id="frameworks-supported"></a>
## Frameworks supported

<a id="dependencies"></a>
## Dependencies

- [Json.NET](https://www.nuget.org/packages/Newtonsoft.Json/) - 13.0.2 or later
- [JsonSubTypes](https://www.nuget.org/packages/JsonSubTypes/) - 1.8.0 or later
- [System.ComponentModel.Annotations](https://www.nuget.org/packages/System.ComponentModel.Annotations) - 5.0.0 or later

The DLLs included in the package may not be the latest version. We recommend using [NuGet](https://docs.nuget.org/consume/installing-nuget) to obtain the latest version of the packages:

| Manticore Search  | manticoresearch-net     |
| ----------------- | ----------------------- |
| >= 6.2.0          | >= 3.3.1                |
| >= 2.5.1          | >= 1.0.x                |

```
Install-Package Newtonsoft.Json
Install-Package JsonSubTypes
Install-Package System.ComponentModel.Annotations
```
<a id="installation"></a>
## Installation
Run the following command to generate the DLL
- [Mac/Linux] `/bin/sh build.sh`
- [Windows] `build.bat`

Then include the DLL (under the `bin` folder) in the C# project, and use the namespaces:
```csharp
using ManticoreSearch.Api;
using ManticoreSearch.Client;
using ManticoreSearch.Model;
```

<a id="usage"></a>
## Usage

To use the API client with a HTTP proxy, setup a `System.Net.WebProxy`
```csharp
Configuration c = new Configuration();
System.Net.WebProxy webProxy = new System.Net.WebProxy("http://myProxyUrl:80/");
webProxy.Credentials = System.Net.CredentialCache.DefaultCredentials;
c.Proxy = webProxy;
```

### Connections
Each ApiClass (properly the ApiClient inside it) will create an instance of HttpClient. It will use that for the entire lifecycle and dispose it when called the Dispose method.

To better manager the connections it's a common practice to reuse the HttpClient and HttpClientHandler (see [here](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/implement-resilient-applications/use-httpclientfactory-to-implement-resilient-http-requests#issues-with-the-original-httpclient-class-available-in-net) for details). To use your own HttpClient instance just pass it to the ApiClass constructor.

```csharp
HttpClientHandler yourHandler = new HttpClientHandler();
HttpClient yourHttpClient = new HttpClient(yourHandler);
var api = new YourApiClass(yourHttpClient, yourHandler);
```

If you want to use an HttpClient and don't have access to the handler, for example in a DI context in Asp.net Core when using IHttpClientFactory.

```csharp
HttpClient yourHttpClient = new HttpClient();
var api = new YourApiClass(yourHttpClient);
```
You'll loose some configuration settings, the features affected are: Setting and Retrieving Cookies, Client Certificates, Proxy settings. You need to either manually handle those in your setup of the HttpClient or they won't be available.

Here an example of DI setup in a sample web project:

```csharp
services.AddHttpClient<YourApiClass>(httpClient =>
   new PetApi(httpClient));
```


<a id="getting-started"></a>
## Getting Started

```csharp
using System.Collections.Generic;
using System.Diagnostics;
using System.Net.Http;
using ManticoreSearch.Api;
using ManticoreSearch.Client;
using ManticoreSearch.Model;

namespace Example
{
    public class Example
    {
        public static void Main()
        {

            Configuration config = new Configuration();
            config.BasePath = "http://127.0.0.1:9308";
            // create instances of HttpClient, HttpClientHandler to be reused later with different Api classes
            HttpClient httpClient = new HttpClient();
            HttpClientHandler httpClientHandler = new HttpClientHandler();
            var apiInstance = new IndexApi(httpClient, config, httpClientHandler);
            var body = "body_example";  // string | 

            try
            {
                // Bulk index operations
                BulkResponse result = apiInstance.Bulk(body);
                Debug.WriteLine(result);
            }
            catch (ApiException e)
            {
                Debug.Print("Exception when calling IndexApi.Bulk: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
            
            apiInstance = new SearchApi(httpClient, config, httpClientHandler);
            
            // Create SearchRequest
            var basicSearchRequest = new BasicSearchRequest(index: "test", query: "Title 1");
            var searchRequest = new SearchRequest(basicSearchRequest);

            try
            {
                // Perform a search
                SearchResponse result = apiInstance.Search(searchRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling SearchApi.Search: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
            

        }
    }
}
```

<a id="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:9308*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndexApi* | [**Bulk**](docs/IndexApi.md#bulk) | **POST** /bulk | Bulk index operations
*IndexApi* | [**Delete**](docs/IndexApi.md#delete) | **POST** /delete | Delete a document in an index
*IndexApi* | [**Insert**](docs/IndexApi.md#insert) | **POST** /insert | Create a new document in an index
*IndexApi* | [**PartialReplace**](docs/IndexApi.md#partialreplace) | **POST** /{index}/_update/{id} | Partially replaces a document in an index
*IndexApi* | [**Replace**](docs/IndexApi.md#replace) | **POST** /replace | Replace new document in an index
*IndexApi* | [**Update**](docs/IndexApi.md#update) | **POST** /update | Update a document in an index
*SearchApi* | [**Percolate**](docs/SearchApi.md#percolate) | **POST** /pq/{index}/search | Perform reverse search on a percolate index
*SearchApi* | [**Search**](docs/SearchApi.md#search) | **POST** /search | Performs a search on an index
*UtilsApi* | [**Sql**](docs/UtilsApi.md#sql) | **POST** /sql | Perform SQL requests


<a id="documentation-for-models"></a>
## Documentation for Models

 - [Model.AggComposite](docs/AggComposite.md)
 - [Model.AggCompositeSource](docs/AggCompositeSource.md)
 - [Model.AggCompositeTerm](docs/AggCompositeTerm.md)
 - [Model.AggTerms](docs/AggTerms.md)
 - [Model.Aggregation](docs/Aggregation.md)
 - [Model.BoolFilter](docs/BoolFilter.md)
 - [Model.BulkResponse](docs/BulkResponse.md)
 - [Model.DeleteDocumentRequest](docs/DeleteDocumentRequest.md)
 - [Model.DeleteResponse](docs/DeleteResponse.md)
 - [Model.ErrorResponse](docs/ErrorResponse.md)
 - [Model.FulltextFilter](docs/FulltextFilter.md)
 - [Model.GeoDistance](docs/GeoDistance.md)
 - [Model.GeoDistanceLocationAnchor](docs/GeoDistanceLocationAnchor.md)
 - [Model.Highlight](docs/Highlight.md)
 - [Model.HighlightFieldOption](docs/HighlightFieldOption.md)
 - [Model.InsertDocumentRequest](docs/InsertDocumentRequest.md)
 - [Model.Join](docs/Join.md)
 - [Model.JoinCond](docs/JoinCond.md)
 - [Model.JoinOn](docs/JoinOn.md)
 - [Model.KnnQuery](docs/KnnQuery.md)
 - [Model.PercolateRequest](docs/PercolateRequest.md)
 - [Model.PercolateRequestQuery](docs/PercolateRequestQuery.md)
 - [Model.QueryFilter](docs/QueryFilter.md)
 - [Model.ReplaceDocumentRequest](docs/ReplaceDocumentRequest.md)
 - [Model.ResponseError](docs/ResponseError.md)
 - [Model.ResponseErrorDetails](docs/ResponseErrorDetails.md)
 - [Model.SearchQuery](docs/SearchQuery.md)
 - [Model.SearchRequest](docs/SearchRequest.md)
 - [Model.SearchResponse](docs/SearchResponse.md)
 - [Model.SearchResponseHits](docs/SearchResponseHits.md)
 - [Model.SourceRules](docs/SourceRules.md)
 - [Model.SuccessResponse](docs/SuccessResponse.md)
 - [Model.UpdateDocumentRequest](docs/UpdateDocumentRequest.md)
 - [Model.UpdateResponse](docs/UpdateResponse.md)


<a id="documentation-for-authorization"></a>
## Documentation for Authorization

Endpoints do not require authorization.
