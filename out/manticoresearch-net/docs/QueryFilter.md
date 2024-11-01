# ManticoreSearch.Model.QueryFilter
Object used to apply various conditions, such as full-text matching or attribute filtering, to a search query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **Object** | Filter object defining a query string | [optional] 
**Match** | **Object** | Filter object defining a match keyword | [optional] 
**MatchPhrase** | **Object** | Filter object defining a match phrase | [optional] 
**MatchAll** | **Object** | Filter object to select all documents | [optional] 
**VarBool** | [**BoolFilter**](BoolFilter.md) |  | [optional] 
**PropertyEquals** | **Object** |  | [optional] 
**VarIn** | **Object** | Filter to match a given set of attribute values. | [optional] 
**Range** | **Object** | Filter to match a given range of attribute values. | [optional] 
**GeoDistance** | [**GeoDistance**](GeoDistance.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
