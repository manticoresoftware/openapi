

# KnnSearchRequestByVector

Request object for knn search operation

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**index** | **String** |  |  |
|**field** | **String** |  |  |
|**queryVector** | **List&lt;BigDecimal&gt;** |  |  |
|**k** | **Integer** |  |  |
|**fulltextFilter** | **Object** |  |  [optional] |
|**attrFilter** | **Object** |  |  [optional] |
|**limit** | **Integer** |  |  [optional] |
|**offset** | **Integer** |  |  [optional] |
|**maxMatches** | **Integer** |  |  [optional] |
|**sort** | **List&lt;Object&gt;** |  |  [optional] |
|**aggs** | [**Map&lt;String, Aggregation&gt;**](Aggregation.md) |  |  [optional] |
|**expressions** | **Map&lt;String, String&gt;** |  |  [optional] |
|**highlight** | [**Highlight**](Highlight.md) |  |  [optional] |
|**source** | **Object** |  |  [optional] |
|**options** | **Map&lt;String, Object&gt;** |  |  [optional] |
|**profile** | **Boolean** |  |  [optional] |
|**trackScores** | **Boolean** |  |  [optional] |





