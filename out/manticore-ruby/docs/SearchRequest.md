# OpenapiClient::SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index** | **String** |  | 
**query** | **Hash&lt;String, Object&gt;** |  | 
**limit** | **Integer** |  | [optional] 
**offset** | **Integer** |  | [optional] 
**max_matches** | **Integer** |  | [optional] 
**sort** | **Array&lt;Object&gt;** |  | [optional] 
**script_fields** | **Object** |  | [optional] 
**highlight** | **Object** |  | [optional] 
**_source** | **Array&lt;String&gt;** |  | [optional] 
**profile** | **Boolean** |  | [optional] 

## Code Sample

```ruby
require 'OpenapiClient'

instance = OpenapiClient::SearchRequest.new(index: test,
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


