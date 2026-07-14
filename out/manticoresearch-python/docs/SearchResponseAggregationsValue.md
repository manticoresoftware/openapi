# SearchResponseAggregationsValue


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**buckets** | **object** |  | [optional] 
**after_key** | **object** | Pagination cursor returned by composite aggregations | [optional] 

## Example

```python
from manticoresearch.models.search_response_aggregations_value import SearchResponseAggregationsValue

# TODO update the JSON string below
json = "{}"
# create an instance of SearchResponseAggregationsValue from a JSON string
search_response_aggregations_value_instance = SearchResponseAggregationsValue.from_json(json)
# print the JSON string representation of the object
print(SearchResponseAggregationsValue.to_json())

# convert the object into a dict
search_response_aggregations_value_dict = search_response_aggregations_value_instance.to_dict()
# create an instance of SearchResponseAggregationsValue from a dict
search_response_aggregations_value_from_dict = SearchResponseAggregationsValue.from_dict(search_response_aggregations_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


