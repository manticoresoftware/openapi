# OpenapiClient::SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**took** | **Integer** |  | [optional] 
**timed_out** | **Boolean** |  | [optional] 
**hits** | [**Hash&lt;String, SearchResponseHits&gt;**](SearchResponseHits.md) |  | [optional] 
**profile** | **Object** |  | [optional] 

## Code Sample

```ruby
require 'OpenapiClient'

instance = OpenapiClient::SearchResponse.new(took: null,
                                 timed_out: null,
                                 hits: {&quot;total&quot;:2,&quot;hits&quot;:[{&quot;_id&quot;:1,&quot;_score&quot;:1,&quot;_source&quot;:{&quot;gid&quot;:11}},{&quot;_id&quot;:2,&quot;_score&quot;:1,&quot;_source&quot;:{&quot;gid&quot;:20}}]},
                                 profile: null)
```


