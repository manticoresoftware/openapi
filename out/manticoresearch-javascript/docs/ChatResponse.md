# Manticoresearch.ChatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**conversationUuid** | **String** | Existing or generated conversation id | 
**userQuery** | **String** | Original user query | 
**searchQuery** | **String** | Standalone search query used for KNN retrieval | 
**response** | **String** | LLM answer as generated | 
**sources** | **String** | JSON string containing retrieved source rows used as LLM context.  | 


