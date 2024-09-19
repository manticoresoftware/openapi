# Manticoresearch.Highlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fragmentSize** | **Number** |  | [optional] 
**limit** | **Number** |  | [optional] 
**limitSnippets** | **Number** |  | [optional] 
**limitWords** | **Number** |  | [optional] 
**numberOfFragments** | **Number** |  | [optional] 
**afterMatch** | **String** |  | [optional] [default to &#39;&lt;/strong&gt;&#39;]
**allowEmpty** | **Boolean** |  | [optional] 
**around** | **Number** |  | [optional] 
**beforeMatch** | **String** |  | [optional] [default to &#39;&lt;strong&gt;&#39;]
**emitZones** | **Boolean** |  | [optional] 
**encoder** | **String** |  | [optional] 
**fields** | [**HighlightAllOfFields**](HighlightAllOfFields.md) |  | [optional] 
**forceAllWords** | **Boolean** |  | [optional] 
**forceSnippets** | **Boolean** |  | [optional] 
**highlightQuery** | [**QueryFilter**](QueryFilter.md) |  | [optional] 
**htmlStripMode** | **String** |  | [optional] 
**limitsPerField** | **Boolean** |  | [optional] 
**noMatchSize** | **Number** |  | [optional] 
**order** | **String** |  | [optional] 
**preTags** | **String** |  | [optional] [default to &#39;&lt;strong&gt;&#39;]
**postTags** | **String** |  | [optional] [default to &#39;&lt;/strong&gt;&#39;]
**startSnippetId** | **Number** |  | [optional] 
**useBoundaries** | **Boolean** |  | [optional] 



## Enum: EncoderEnum


* `default` (value: `"default"`)

* `html` (value: `"html"`)





## Enum: HtmlStripModeEnum


* `none` (value: `"none"`)

* `strip` (value: `"strip"`)

* `index` (value: `"index"`)

* `retain` (value: `"retain"`)





## Enum: NoMatchSizeEnum


* `0` (value: `0`)

* `1` (value: `1`)





## Enum: OrderEnum


* `asc` (value: `"asc"`)

* `desc` (value: `"desc"`)

* `score` (value: `"score"`)




