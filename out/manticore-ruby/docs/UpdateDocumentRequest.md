# OpenapiClient::UpdateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **String** |  | 
**doc** | **Hash&lt;String, Object&gt;** | Index name | 
**id** | **Integer** | Document ID | [optional] 
**query** | **Hash&lt;String, Object&gt;** | Query tree object | [optional] 

## Code Sample

```ruby
require 'OpenapiClient'

instance = OpenapiClient::UpdateDocumentRequest.new(index: null,
                                 doc: {&quot;gid&quot;:10},
                                 id: null,
                                 query: {&quot;query&quot;:{&quot;match&quot;:{&quot;title&quot;:&quot;match me&quot;}}})
```


