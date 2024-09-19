# Manticoresearch.QueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**equals** | [**EqualsFilterEquals**](EqualsFilterEquals.md) |  | [optional] 
**_in** | **{String: [EqualsFilterEquals]}** |  | [optional] 
**range** | [**{String: RangeFilterRangeValue}**](RangeFilterRangeValue.md) |  | [optional] 
**geoDistance** | [**GeoFilterGeoDistance**](GeoFilterGeoDistance.md) |  | [optional] 
**bool** | [**BoolFilterBool**](BoolFilterBool.md) |  | [optional] 
**match** | [**MatchFilterMatch**](MatchFilterMatch.md) |  | [optional] 
**matchAll** | **String** |  | [optional] 
**matchPhrase** | **{String: String}** |  | [optional] 
**queryString** | **String** |  | [optional] 



## Enum: MatchAllEnum


* `{}` (value: `"{}"`)




