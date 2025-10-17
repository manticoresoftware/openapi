# QueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query_string** | Option<**String**> | Filter object defining a query string | [optional]
**r#match** | Option<[**serde_json::Value**](.md)> | Filter object defining a match keyword passed as a string or in a Match object | [optional]
**match_phrase** | Option<[**serde_json::Value**](.md)> | Filter object defining a match phrase | [optional]
**match_all** | Option<[**serde_json::Value**](.md)> | Filter object to select all documents | [optional]
**bool** | Option<[**models::BoolFilter**](boolFilter.md)> |  | [optional]
**equals** | Option<[**serde_json::Value**](.md)> |  | [optional]
**r#in** | Option<[**serde_json::Value**](.md)> | Filter to match a given set of attribute values. | [optional]
**range** | Option<[**serde_json::Value**](.md)> | Filter to match a given range of attribute values passed in Range objects | [optional]
**geo_distance** | Option<[**models::GeoDistance**](geoDistance.md)> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


