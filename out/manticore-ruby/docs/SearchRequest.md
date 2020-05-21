# OpenapiClient::SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **String** |  | 
**query** | **Object** |  | 
**limit** | **Integer** |  | [optional] 
**offset** | **Integer** |  | [optional] 
**max_matches** | **Integer** |  | [optional] 
**sort** | [**Array&lt;OneOfstringobject&gt;**](OneOfstringobject.md) |  | [optional] 
**script_fields** | **Object** |  | [optional] 
**highlight** | **Object** |  | [optional] 
**_source** | [**OneOfstringobject**](OneOfstringobject.md) |  | [optional] 
**profile** | **Boolean** |  | [optional] 

## Code Sample

```ruby
require 'OpenapiClient'

instance = OpenapiClient::SearchRequest.new(index: null,
                                 query: null,
                                 limit: null,
                                 offset: null,
                                 max_matches: null,
                                 sort: null,
                                 script_fields: null,
                                 highlight: null,
                                 _source: null,
                                 profile: null)
```


