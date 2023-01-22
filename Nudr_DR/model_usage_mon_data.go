/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"time"
)

// UsageMonData Contains remain allowed usage data for a subscriber.
type UsageMonData struct {
	LimitId string `json:"limitId"`
	// Identifies the SNSSAI and DNN combinations for remain allowed usage data for a subscriber. The S-NSSAI is the key of the map. 
	Scopes *map[string]UsageMonDataScope `json:"scopes,omitempty"`
	UmLevel *UsageMonLevel `json:"umLevel,omitempty"`
	AllowedUsage *UsageThreshold `json:"allowedUsage,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ResetTime *time.Time `json:"resetTime,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewUsageMonData instantiates a new UsageMonData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageMonData(limitId string) *UsageMonData {
	this := UsageMonData{}
	this.LimitId = limitId
	return &this
}

// NewUsageMonDataWithDefaults instantiates a new UsageMonData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageMonDataWithDefaults() *UsageMonData {
	this := UsageMonData{}
	return &this
}

// GetLimitId returns the LimitId field value
func (o *UsageMonData) GetLimitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LimitId
}

// GetLimitIdOk returns a tuple with the LimitId field value
// and a boolean to check if the value has been set.
func (o *UsageMonData) GetLimitIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LimitId, true
}

// SetLimitId sets field value
func (o *UsageMonData) SetLimitId(v string) {
	o.LimitId = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UsageMonData) GetScopes() map[string]UsageMonDataScope {
	if o == nil || isNil(o.Scopes) {
		var ret map[string]UsageMonDataScope
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMonData) GetScopesOk() (*map[string]UsageMonDataScope, bool) {
	if o == nil || isNil(o.Scopes) {
    return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *UsageMonData) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given map[string]UsageMonDataScope and assigns it to the Scopes field.
func (o *UsageMonData) SetScopes(v map[string]UsageMonDataScope) {
	o.Scopes = &v
}

// GetUmLevel returns the UmLevel field value if set, zero value otherwise.
func (o *UsageMonData) GetUmLevel() UsageMonLevel {
	if o == nil || isNil(o.UmLevel) {
		var ret UsageMonLevel
		return ret
	}
	return *o.UmLevel
}

// GetUmLevelOk returns a tuple with the UmLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMonData) GetUmLevelOk() (*UsageMonLevel, bool) {
	if o == nil || isNil(o.UmLevel) {
    return nil, false
	}
	return o.UmLevel, true
}

// HasUmLevel returns a boolean if a field has been set.
func (o *UsageMonData) HasUmLevel() bool {
	if o != nil && !isNil(o.UmLevel) {
		return true
	}

	return false
}

// SetUmLevel gets a reference to the given UsageMonLevel and assigns it to the UmLevel field.
func (o *UsageMonData) SetUmLevel(v UsageMonLevel) {
	o.UmLevel = &v
}

// GetAllowedUsage returns the AllowedUsage field value if set, zero value otherwise.
func (o *UsageMonData) GetAllowedUsage() UsageThreshold {
	if o == nil || isNil(o.AllowedUsage) {
		var ret UsageThreshold
		return ret
	}
	return *o.AllowedUsage
}

// GetAllowedUsageOk returns a tuple with the AllowedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMonData) GetAllowedUsageOk() (*UsageThreshold, bool) {
	if o == nil || isNil(o.AllowedUsage) {
    return nil, false
	}
	return o.AllowedUsage, true
}

// HasAllowedUsage returns a boolean if a field has been set.
func (o *UsageMonData) HasAllowedUsage() bool {
	if o != nil && !isNil(o.AllowedUsage) {
		return true
	}

	return false
}

// SetAllowedUsage gets a reference to the given UsageThreshold and assigns it to the AllowedUsage field.
func (o *UsageMonData) SetAllowedUsage(v UsageThreshold) {
	o.AllowedUsage = &v
}

// GetResetTime returns the ResetTime field value if set, zero value otherwise.
func (o *UsageMonData) GetResetTime() time.Time {
	if o == nil || isNil(o.ResetTime) {
		var ret time.Time
		return ret
	}
	return *o.ResetTime
}

// GetResetTimeOk returns a tuple with the ResetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMonData) GetResetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ResetTime) {
    return nil, false
	}
	return o.ResetTime, true
}

// HasResetTime returns a boolean if a field has been set.
func (o *UsageMonData) HasResetTime() bool {
	if o != nil && !isNil(o.ResetTime) {
		return true
	}

	return false
}

// SetResetTime gets a reference to the given time.Time and assigns it to the ResetTime field.
func (o *UsageMonData) SetResetTime(v time.Time) {
	o.ResetTime = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *UsageMonData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMonData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *UsageMonData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *UsageMonData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *UsageMonData) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMonData) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *UsageMonData) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *UsageMonData) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o UsageMonData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["limitId"] = o.LimitId
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.UmLevel) {
		toSerialize["umLevel"] = o.UmLevel
	}
	if !isNil(o.AllowedUsage) {
		toSerialize["allowedUsage"] = o.AllowedUsage
	}
	if !isNil(o.ResetTime) {
		toSerialize["resetTime"] = o.ResetTime
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return json.Marshal(toSerialize)
}

type NullableUsageMonData struct {
	value *UsageMonData
	isSet bool
}

func (v NullableUsageMonData) Get() *UsageMonData {
	return v.value
}

func (v *NullableUsageMonData) Set(val *UsageMonData) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMonData) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMonData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMonData(val *UsageMonData) *NullableUsageMonData {
	return &NullableUsageMonData{value: val, isSet: true}
}

func (v NullableUsageMonData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMonData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


