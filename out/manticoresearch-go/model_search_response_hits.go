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

// checks if the SearchResponseHits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResponseHits{}

// SearchResponseHits Object containing the search hits, which represent the documents that matched the query.
type SearchResponseHits struct {
	// Maximum score among the matched documents
	MaxScore *int32
	// Total number of matched documents
	Total *int32
	// Indicates whether the total number of hits is accurate or an estimate
	TotalRelation *string
	// Array of hit objects, each representing a matched document
	Hits []map[string]interface{}
}

// NewSearchResponseHits instantiates a new SearchResponseHits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResponseHits() *SearchResponseHits {
	this := SearchResponseHits{}
	return &this
}

// NewSearchResponseHitsWithDefaults instantiates a new SearchResponseHits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResponseHitsWithDefaults() *SearchResponseHits {
	this := SearchResponseHits{}
	return &this
}

// GetMaxScore returns the MaxScore field value if set, zero value otherwise.
func (o *SearchResponseHits) GetMaxScore() int32 {
	if o == nil || IsNil(o.MaxScore) {
		var ret int32
		return ret
	}
	return *o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseHits) GetMaxScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxScore) {
		return nil, false
	}
	return o.MaxScore, true
}

// HasMaxScore returns a boolean if a field has been set.
func (o *SearchResponseHits) HasMaxScore() bool {
	if o != nil && !IsNil(o.MaxScore) {
		return true
	}

	return false
}

// SetMaxScore gets a reference to the given int32 and assigns it to the MaxScore field.
func (o *SearchResponseHits) SetMaxScore(v int32) {
	o.MaxScore = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SearchResponseHits) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseHits) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SearchResponseHits) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *SearchResponseHits) SetTotal(v int32) {
	o.Total = &v
}

// GetTotalRelation returns the TotalRelation field value if set, zero value otherwise.
func (o *SearchResponseHits) GetTotalRelation() string {
	if o == nil || IsNil(o.TotalRelation) {
		var ret string
		return ret
	}
	return *o.TotalRelation
}

// GetTotalRelationOk returns a tuple with the TotalRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseHits) GetTotalRelationOk() (*string, bool) {
	if o == nil || IsNil(o.TotalRelation) {
		return nil, false
	}
	return o.TotalRelation, true
}

// HasTotalRelation returns a boolean if a field has been set.
func (o *SearchResponseHits) HasTotalRelation() bool {
	if o != nil && !IsNil(o.TotalRelation) {
		return true
	}

	return false
}

// SetTotalRelation gets a reference to the given string and assigns it to the TotalRelation field.
func (o *SearchResponseHits) SetTotalRelation(v string) {
	o.TotalRelation = &v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *SearchResponseHits) GetHits() []map[string]interface{} {
	if o == nil || IsNil(o.Hits) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseHits) GetHitsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Hits) {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *SearchResponseHits) HasHits() bool {
	if o != nil && !IsNil(o.Hits) {
		return true
	}

	return false
}

// SetHits gets a reference to the given []map[string]interface{} and assigns it to the Hits field.
func (o *SearchResponseHits) SetHits(v []map[string]interface{}) {
	o.Hits = v
}

func (o SearchResponseHits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResponseHits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxScore) {
		toSerialize["max_score"] = o.MaxScore
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.TotalRelation) {
		toSerialize["total_relation"] = o.TotalRelation
	}
	if !IsNil(o.Hits) {
		toSerialize["hits"] = o.Hits
	}
	return toSerialize, nil
}

type NullableSearchResponseHits struct {
	value *SearchResponseHits
	isSet bool
}

func (v NullableSearchResponseHits) Get() *SearchResponseHits {
	return v.value
}

func (v *NullableSearchResponseHits) Set(val *SearchResponseHits) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponseHits) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponseHits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponseHits(val *SearchResponseHits) *NullableSearchResponseHits {
	return &NullableSearchResponseHits{value: val, isSet: true}
}

func (v NullableSearchResponseHits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponseHits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


