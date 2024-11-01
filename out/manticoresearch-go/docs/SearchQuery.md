# SearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **interface{}** | Filter object defining a query string | [optional] 
**Match** | Pointer to **interface{}** | Filter object defining a match keyword passed as a string or in a Match object | [optional] 
**MatchPhrase** | Pointer to **interface{}** | Filter object defining a match phrase | [optional] 
**MatchAll** | Pointer to **interface{}** | Filter object to select all documents | [optional] 
**Bool** | Pointer to [**BoolFilter**](BoolFilter.md) |  | [optional] 
**Equals** | Pointer to **interface{}** |  | [optional] 
**In** | Pointer to **interface{}** | Filter to match a given set of attribute values. | [optional] 
**Range** | Pointer to **interface{}** | Filter to match a given range of attribute values passed in Range objects | [optional] 
**GeoDistance** | Pointer to [**GeoDistance**](GeoDistance.md) |  | [optional] 
**Highlight** | Pointer to [**Highlight**](Highlight.md) |  | [optional] 

## Methods

### NewSearchQuery

`func NewSearchQuery() *SearchQuery`

NewSearchQuery instantiates a new SearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryWithDefaults

`func NewSearchQueryWithDefaults() *SearchQuery`

NewSearchQueryWithDefaults instantiates a new SearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *SearchQuery) GetQueryString() interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *SearchQuery) GetQueryStringOk() (*interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *SearchQuery) SetQueryString(v interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *SearchQuery) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### SetQueryStringNil

`func (o *SearchQuery) SetQueryStringNil(b bool)`

 SetQueryStringNil sets the value for QueryString to be an explicit nil

### UnsetQueryString
`func (o *SearchQuery) UnsetQueryString()`

UnsetQueryString ensures that no value is present for QueryString, not even an explicit nil
### GetMatch

`func (o *SearchQuery) GetMatch() interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *SearchQuery) GetMatchOk() (*interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *SearchQuery) SetMatch(v interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *SearchQuery) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### SetMatchNil

`func (o *SearchQuery) SetMatchNil(b bool)`

 SetMatchNil sets the value for Match to be an explicit nil

### UnsetMatch
`func (o *SearchQuery) UnsetMatch()`

UnsetMatch ensures that no value is present for Match, not even an explicit nil
### GetMatchPhrase

`func (o *SearchQuery) GetMatchPhrase() interface{}`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *SearchQuery) GetMatchPhraseOk() (*interface{}, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *SearchQuery) SetMatchPhrase(v interface{})`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *SearchQuery) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### SetMatchPhraseNil

`func (o *SearchQuery) SetMatchPhraseNil(b bool)`

 SetMatchPhraseNil sets the value for MatchPhrase to be an explicit nil

### UnsetMatchPhrase
`func (o *SearchQuery) UnsetMatchPhrase()`

UnsetMatchPhrase ensures that no value is present for MatchPhrase, not even an explicit nil
### GetMatchAll

`func (o *SearchQuery) GetMatchAll() interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *SearchQuery) GetMatchAllOk() (*interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *SearchQuery) SetMatchAll(v interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *SearchQuery) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### SetMatchAllNil

`func (o *SearchQuery) SetMatchAllNil(b bool)`

 SetMatchAllNil sets the value for MatchAll to be an explicit nil

### UnsetMatchAll
`func (o *SearchQuery) UnsetMatchAll()`

UnsetMatchAll ensures that no value is present for MatchAll, not even an explicit nil
### GetBool

`func (o *SearchQuery) GetBool() BoolFilter`

GetBool returns the Bool field if non-nil, zero value otherwise.

### GetBoolOk

`func (o *SearchQuery) GetBoolOk() (*BoolFilter, bool)`

GetBoolOk returns a tuple with the Bool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBool

`func (o *SearchQuery) SetBool(v BoolFilter)`

SetBool sets Bool field to given value.

### HasBool

`func (o *SearchQuery) HasBool() bool`

HasBool returns a boolean if a field has been set.

### GetEquals

`func (o *SearchQuery) GetEquals() interface{}`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *SearchQuery) GetEqualsOk() (*interface{}, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *SearchQuery) SetEquals(v interface{})`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *SearchQuery) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### SetEqualsNil

`func (o *SearchQuery) SetEqualsNil(b bool)`

 SetEqualsNil sets the value for Equals to be an explicit nil

### UnsetEquals
`func (o *SearchQuery) UnsetEquals()`

UnsetEquals ensures that no value is present for Equals, not even an explicit nil
### GetIn

`func (o *SearchQuery) GetIn() interface{}`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *SearchQuery) GetInOk() (*interface{}, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *SearchQuery) SetIn(v interface{})`

SetIn sets In field to given value.

### HasIn

`func (o *SearchQuery) HasIn() bool`

HasIn returns a boolean if a field has been set.

### SetInNil

`func (o *SearchQuery) SetInNil(b bool)`

 SetInNil sets the value for In to be an explicit nil

### UnsetIn
`func (o *SearchQuery) UnsetIn()`

UnsetIn ensures that no value is present for In, not even an explicit nil
### GetRange

`func (o *SearchQuery) GetRange() interface{}`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SearchQuery) GetRangeOk() (*interface{}, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SearchQuery) SetRange(v interface{})`

SetRange sets Range field to given value.

### HasRange

`func (o *SearchQuery) HasRange() bool`

HasRange returns a boolean if a field has been set.

### SetRangeNil

`func (o *SearchQuery) SetRangeNil(b bool)`

 SetRangeNil sets the value for Range to be an explicit nil

### UnsetRange
`func (o *SearchQuery) UnsetRange()`

UnsetRange ensures that no value is present for Range, not even an explicit nil
### GetGeoDistance

`func (o *SearchQuery) GetGeoDistance() GeoDistance`

GetGeoDistance returns the GeoDistance field if non-nil, zero value otherwise.

### GetGeoDistanceOk

`func (o *SearchQuery) GetGeoDistanceOk() (*GeoDistance, bool)`

GetGeoDistanceOk returns a tuple with the GeoDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoDistance

`func (o *SearchQuery) SetGeoDistance(v GeoDistance)`

SetGeoDistance sets GeoDistance field to given value.

### HasGeoDistance

`func (o *SearchQuery) HasGeoDistance() bool`

HasGeoDistance returns a boolean if a field has been set.

### GetHighlight

`func (o *SearchQuery) GetHighlight() Highlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *SearchQuery) GetHighlightOk() (*Highlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *SearchQuery) SetHighlight(v Highlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *SearchQuery) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

