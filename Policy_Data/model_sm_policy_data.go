/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Policy_Data

import (
	"encoding/json"
)

// SmPolicyData Contains the SM policy data for a given subscriber.
type SmPolicyData struct {
	// Contains Session Management Policy data per S-NSSAI for all the SNSSAIs of the subscriber. The key of the map is the S-NSSAI. 
	SmPolicySnssaiData map[string]SmPolicySnssaiData `json:"smPolicySnssaiData"`
	// Contains a list of usage monitoring profiles associated with the subscriber. The limit identifier is used as the key of the map. 
	UmDataLimits *map[string]UsageMonDataLimit `json:"umDataLimits,omitempty"`
	// Contains the remaining allowed usage data associated with the subscriber. The limit identifier is used as the key of the map. 
	UmData *map[string]UsageMonData `json:"umData,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewSmPolicyData instantiates a new SmPolicyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmPolicyData(smPolicySnssaiData map[string]SmPolicySnssaiData) *SmPolicyData {
	this := SmPolicyData{}
	this.SmPolicySnssaiData = smPolicySnssaiData
	return &this
}

// NewSmPolicyDataWithDefaults instantiates a new SmPolicyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmPolicyDataWithDefaults() *SmPolicyData {
	this := SmPolicyData{}
	return &this
}

// GetSmPolicySnssaiData returns the SmPolicySnssaiData field value
func (o *SmPolicyData) GetSmPolicySnssaiData() map[string]SmPolicySnssaiData {
	if o == nil {
		var ret map[string]SmPolicySnssaiData
		return ret
	}

	return o.SmPolicySnssaiData
}

// GetSmPolicySnssaiDataOk returns a tuple with the SmPolicySnssaiData field value
// and a boolean to check if the value has been set.
func (o *SmPolicyData) GetSmPolicySnssaiDataOk() (*map[string]SmPolicySnssaiData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SmPolicySnssaiData, true
}

// SetSmPolicySnssaiData sets field value
func (o *SmPolicyData) SetSmPolicySnssaiData(v map[string]SmPolicySnssaiData) {
	o.SmPolicySnssaiData = v
}

// GetUmDataLimits returns the UmDataLimits field value if set, zero value otherwise.
func (o *SmPolicyData) GetUmDataLimits() map[string]UsageMonDataLimit {
	if o == nil || isNil(o.UmDataLimits) {
		var ret map[string]UsageMonDataLimit
		return ret
	}
	return *o.UmDataLimits
}

// GetUmDataLimitsOk returns a tuple with the UmDataLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyData) GetUmDataLimitsOk() (*map[string]UsageMonDataLimit, bool) {
	if o == nil || isNil(o.UmDataLimits) {
    return nil, false
	}
	return o.UmDataLimits, true
}

// HasUmDataLimits returns a boolean if a field has been set.
func (o *SmPolicyData) HasUmDataLimits() bool {
	if o != nil && !isNil(o.UmDataLimits) {
		return true
	}

	return false
}

// SetUmDataLimits gets a reference to the given map[string]UsageMonDataLimit and assigns it to the UmDataLimits field.
func (o *SmPolicyData) SetUmDataLimits(v map[string]UsageMonDataLimit) {
	o.UmDataLimits = &v
}

// GetUmData returns the UmData field value if set, zero value otherwise.
func (o *SmPolicyData) GetUmData() map[string]UsageMonData {
	if o == nil || isNil(o.UmData) {
		var ret map[string]UsageMonData
		return ret
	}
	return *o.UmData
}

// GetUmDataOk returns a tuple with the UmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyData) GetUmDataOk() (*map[string]UsageMonData, bool) {
	if o == nil || isNil(o.UmData) {
    return nil, false
	}
	return o.UmData, true
}

// HasUmData returns a boolean if a field has been set.
func (o *SmPolicyData) HasUmData() bool {
	if o != nil && !isNil(o.UmData) {
		return true
	}

	return false
}

// SetUmData gets a reference to the given map[string]UsageMonData and assigns it to the UmData field.
func (o *SmPolicyData) SetUmData(v map[string]UsageMonData) {
	o.UmData = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *SmPolicyData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *SmPolicyData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *SmPolicyData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *SmPolicyData) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyData) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *SmPolicyData) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *SmPolicyData) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o SmPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["smPolicySnssaiData"] = o.SmPolicySnssaiData
	}
	if !isNil(o.UmDataLimits) {
		toSerialize["umDataLimits"] = o.UmDataLimits
	}
	if !isNil(o.UmData) {
		toSerialize["umData"] = o.UmData
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return json.Marshal(toSerialize)
}

type NullableSmPolicyData struct {
	value *SmPolicyData
	isSet bool
}

func (v NullableSmPolicyData) Get() *SmPolicyData {
	return v.value
}

func (v *NullableSmPolicyData) Set(val *SmPolicyData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmPolicyData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmPolicyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmPolicyData(val *SmPolicyData) *NullableSmPolicyData {
	return &NullableSmPolicyData{value: val, isSet: true}
}

func (v NullableSmPolicyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmPolicyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


