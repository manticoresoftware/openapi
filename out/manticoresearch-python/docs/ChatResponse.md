# ChatResponse

Response object for conversational search. For more information see [Conversational search](https://manual.manticoresearch.com/Searching/Conversational_search#Response) 

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**conversation_uuid** | **str** | Existing or generated conversation id | 
**user_query** | **str** | Original user query | 
**search_query** | **str** | Standalone search query used for KNN retrieval | 
**response** | **str** | LLM answer as generated | 
**sources** | **str** | JSON string containing retrieved source rows used as LLM context.  | 

## Example

```python
from manticoresearch.models.chat_response import ChatResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ChatResponse from a JSON string
chat_response_instance = ChatResponse.from_json(json)
# print the JSON string representation of the object
print(ChatResponse.to_json())

# convert the object into a dict
chat_response_dict = chat_response_instance.to_dict()
# create an instance of ChatResponse from a dict
chat_response_from_dict = ChatResponse.from_dict(chat_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


