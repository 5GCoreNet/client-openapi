/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_DataRepository

import (
	"encoding/json"
)

// RecordSearchResult Count and collection of Record references matching the providing filter.
type RecordSearchResult struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Count int32 `json:"count"`
	References []string `json:"references,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// A map (list of key-value pairs where recordId serves as key) of Records
	MatchingRecords *map[string]Record `json:"matchingRecords,omitempty"`
}

// NewRecordSearchResult instantiates a new RecordSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordSearchResult(count int32) *RecordSearchResult {
	this := RecordSearchResult{}
	this.Count = count
	return &this
}

// NewRecordSearchResultWithDefaults instantiates a new RecordSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordSearchResultWithDefaults() *RecordSearchResult {
	this := RecordSearchResult{}
	return &this
}

// GetCount returns the Count field value
func (o *RecordSearchResult) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *RecordSearchResult) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *RecordSearchResult) SetCount(v int32) {
	o.Count = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *RecordSearchResult) GetReferences() []string {
	if o == nil || isNil(o.References) {
		var ret []string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordSearchResult) GetReferencesOk() ([]string, bool) {
	if o == nil || isNil(o.References) {
    return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *RecordSearchResult) HasReferences() bool {
	if o != nil && !isNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []string and assigns it to the References field.
func (o *RecordSearchResult) SetReferences(v []string) {
	o.References = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *RecordSearchResult) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordSearchResult) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *RecordSearchResult) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *RecordSearchResult) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetMatchingRecords returns the MatchingRecords field value if set, zero value otherwise.
func (o *RecordSearchResult) GetMatchingRecords() map[string]Record {
	if o == nil || isNil(o.MatchingRecords) {
		var ret map[string]Record
		return ret
	}
	return *o.MatchingRecords
}

// GetMatchingRecordsOk returns a tuple with the MatchingRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordSearchResult) GetMatchingRecordsOk() (*map[string]Record, bool) {
	if o == nil || isNil(o.MatchingRecords) {
    return nil, false
	}
	return o.MatchingRecords, true
}

// HasMatchingRecords returns a boolean if a field has been set.
func (o *RecordSearchResult) HasMatchingRecords() bool {
	if o != nil && !isNil(o.MatchingRecords) {
		return true
	}

	return false
}

// SetMatchingRecords gets a reference to the given map[string]Record and assigns it to the MatchingRecords field.
func (o *RecordSearchResult) SetMatchingRecords(v map[string]Record) {
	o.MatchingRecords = &v
}

func (o RecordSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.MatchingRecords) {
		toSerialize["matchingRecords"] = o.MatchingRecords
	}
	return json.Marshal(toSerialize)
}

type NullableRecordSearchResult struct {
	value *RecordSearchResult
	isSet bool
}

func (v NullableRecordSearchResult) Get() *RecordSearchResult {
	return v.value
}

func (v *NullableRecordSearchResult) Set(val *RecordSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordSearchResult(val *RecordSearchResult) *NullableRecordSearchResult {
	return &NullableRecordSearchResult{value: val, isSet: true}
}

func (v NullableRecordSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


