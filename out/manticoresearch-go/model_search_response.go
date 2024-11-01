/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 5.0.0
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResponse{}

// SearchResponse Response object containing the results of a search request
type SearchResponse struct {
	// Time taken to execute the search
	Took *int32
	// Indicates whether the search operation timed out
	TimedOut *bool
	// Aggregated search results grouped by the specified criteria
	Aggregations map[string]interface{}
	Hits *SearchResponseHits
	// Profile information about the search execution, if profiling is enabled
	Profile map[string]interface{}
	// Warnings encountered during the search operation
	Warning map[string]interface{}
}

// NewSearchResponse instantiates a new SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResponse() *SearchResponse {
	this := SearchResponse{}
	return &this
}

// NewSearchResponseWithDefaults instantiates a new SearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResponseWithDefaults() *SearchResponse {
	this := SearchResponse{}
	return &this
}

// GetTook returns the Took field value if set, zero value otherwise.
func (o *SearchResponse) GetTook() int32 {
	if o == nil || IsNil(o.Took) {
		var ret int32
		return ret
	}
	return *o.Took
}

// GetTookOk returns a tuple with the Took field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetTookOk() (*int32, bool) {
	if o == nil || IsNil(o.Took) {
		return nil, false
	}
	return o.Took, true
}

// HasTook returns a boolean if a field has been set.
func (o *SearchResponse) HasTook() bool {
	if o != nil && !IsNil(o.Took) {
		return true
	}

	return false
}

// SetTook gets a reference to the given int32 and assigns it to the Took field.
func (o *SearchResponse) SetTook(v int32) {
	o.Took = &v
}

// GetTimedOut returns the TimedOut field value if set, zero value otherwise.
func (o *SearchResponse) GetTimedOut() bool {
	if o == nil || IsNil(o.TimedOut) {
		var ret bool
		return ret
	}
	return *o.TimedOut
}

// GetTimedOutOk returns a tuple with the TimedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetTimedOutOk() (*bool, bool) {
	if o == nil || IsNil(o.TimedOut) {
		return nil, false
	}
	return o.TimedOut, true
}

// HasTimedOut returns a boolean if a field has been set.
func (o *SearchResponse) HasTimedOut() bool {
	if o != nil && !IsNil(o.TimedOut) {
		return true
	}

	return false
}

// SetTimedOut gets a reference to the given bool and assigns it to the TimedOut field.
func (o *SearchResponse) SetTimedOut(v bool) {
	o.TimedOut = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *SearchResponse) GetAggregations() map[string]interface{} {
	if o == nil || IsNil(o.Aggregations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetAggregationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Aggregations) {
		return map[string]interface{}{}, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *SearchResponse) HasAggregations() bool {
	if o != nil && !IsNil(o.Aggregations) {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given map[string]interface{} and assigns it to the Aggregations field.
func (o *SearchResponse) SetAggregations(v map[string]interface{}) {
	o.Aggregations = v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *SearchResponse) GetHits() SearchResponseHits {
	if o == nil || IsNil(o.Hits) {
		var ret SearchResponseHits
		return ret
	}
	return *o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetHitsOk() (*SearchResponseHits, bool) {
	if o == nil || IsNil(o.Hits) {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *SearchResponse) HasHits() bool {
	if o != nil && !IsNil(o.Hits) {
		return true
	}

	return false
}

// SetHits gets a reference to the given SearchResponseHits and assigns it to the Hits field.
func (o *SearchResponse) SetHits(v SearchResponseHits) {
	o.Hits = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SearchResponse) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SearchResponse) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *SearchResponse) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *SearchResponse) GetWarning() map[string]interface{} {
	if o == nil || IsNil(o.Warning) {
		var ret map[string]interface{}
		return ret
	}
	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetWarningOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Warning) {
		return map[string]interface{}{}, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *SearchResponse) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given map[string]interface{} and assigns it to the Warning field.
func (o *SearchResponse) SetWarning(v map[string]interface{}) {
	o.Warning = v
}

func (o SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Took) {
		toSerialize["took"] = o.Took
	}
	if !IsNil(o.TimedOut) {
		toSerialize["timed_out"] = o.TimedOut
	}
	if !IsNil(o.Aggregations) {
		toSerialize["aggregations"] = o.Aggregations
	}
	if !IsNil(o.Hits) {
		toSerialize["hits"] = o.Hits
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	return toSerialize, nil
}

type NullableSearchResponse struct {
	value *SearchResponse
	isSet bool
}

func (v NullableSearchResponse) Get() *SearchResponse {
	return v.value
}

func (v *NullableSearchResponse) Set(val *SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponse(val *SearchResponse) *NullableSearchResponse {
	return &NullableSearchResponse{value: val, isSet: true}
}

func (v NullableSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

