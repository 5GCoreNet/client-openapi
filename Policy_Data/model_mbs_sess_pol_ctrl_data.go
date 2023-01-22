/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Policy_Data

import (
	"encoding/json"
)

// MbsSessPolCtrlData Represents MBS Session Policy Control Data.
type MbsSessPolCtrlData struct {
	Var5qis []int32 `json:"5qis,omitempty"`
	// nullable true shall not be used for this attribute. Unsigned integer indicating the ARP Priority Level (see clause 5.7.2.2 of 3GPP TS 23.501, within the range 1 to 15.Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.  
	MaxMbsArpLevel NullableInt32 `json:"maxMbsArpLevel,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxMbsSessionAmbr *string `json:"maxMbsSessionAmbr,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxGbr *string `json:"maxGbr,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewMbsSessPolCtrlData instantiates a new MbsSessPolCtrlData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessPolCtrlData() *MbsSessPolCtrlData {
	this := MbsSessPolCtrlData{}
	return &this
}

// NewMbsSessPolCtrlDataWithDefaults instantiates a new MbsSessPolCtrlData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessPolCtrlDataWithDefaults() *MbsSessPolCtrlData {
	this := MbsSessPolCtrlData{}
	return &this
}

// GetVar5qis returns the Var5qis field value if set, zero value otherwise.
func (o *MbsSessPolCtrlData) GetVar5qis() []int32 {
	if o == nil || isNil(o.Var5qis) {
		var ret []int32
		return ret
	}
	return o.Var5qis
}

// GetVar5qisOk returns a tuple with the Var5qis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessPolCtrlData) GetVar5qisOk() ([]int32, bool) {
	if o == nil || isNil(o.Var5qis) {
    return nil, false
	}
	return o.Var5qis, true
}

// HasVar5qis returns a boolean if a field has been set.
func (o *MbsSessPolCtrlData) HasVar5qis() bool {
	if o != nil && !isNil(o.Var5qis) {
		return true
	}

	return false
}

// SetVar5qis gets a reference to the given []int32 and assigns it to the Var5qis field.
func (o *MbsSessPolCtrlData) SetVar5qis(v []int32) {
	o.Var5qis = v
}

// GetMaxMbsArpLevel returns the MaxMbsArpLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbsSessPolCtrlData) GetMaxMbsArpLevel() int32 {
	if o == nil || isNil(o.MaxMbsArpLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxMbsArpLevel.Get()
}

// GetMaxMbsArpLevelOk returns a tuple with the MaxMbsArpLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbsSessPolCtrlData) GetMaxMbsArpLevelOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxMbsArpLevel.Get(), o.MaxMbsArpLevel.IsSet()
}

// HasMaxMbsArpLevel returns a boolean if a field has been set.
func (o *MbsSessPolCtrlData) HasMaxMbsArpLevel() bool {
	if o != nil && o.MaxMbsArpLevel.IsSet() {
		return true
	}

	return false
}

// SetMaxMbsArpLevel gets a reference to the given NullableInt32 and assigns it to the MaxMbsArpLevel field.
func (o *MbsSessPolCtrlData) SetMaxMbsArpLevel(v int32) {
	o.MaxMbsArpLevel.Set(&v)
}
// SetMaxMbsArpLevelNil sets the value for MaxMbsArpLevel to be an explicit nil
func (o *MbsSessPolCtrlData) SetMaxMbsArpLevelNil() {
	o.MaxMbsArpLevel.Set(nil)
}

// UnsetMaxMbsArpLevel ensures that no value is present for MaxMbsArpLevel, not even an explicit nil
func (o *MbsSessPolCtrlData) UnsetMaxMbsArpLevel() {
	o.MaxMbsArpLevel.Unset()
}

// GetMaxMbsSessionAmbr returns the MaxMbsSessionAmbr field value if set, zero value otherwise.
func (o *MbsSessPolCtrlData) GetMaxMbsSessionAmbr() string {
	if o == nil || isNil(o.MaxMbsSessionAmbr) {
		var ret string
		return ret
	}
	return *o.MaxMbsSessionAmbr
}

// GetMaxMbsSessionAmbrOk returns a tuple with the MaxMbsSessionAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessPolCtrlData) GetMaxMbsSessionAmbrOk() (*string, bool) {
	if o == nil || isNil(o.MaxMbsSessionAmbr) {
    return nil, false
	}
	return o.MaxMbsSessionAmbr, true
}

// HasMaxMbsSessionAmbr returns a boolean if a field has been set.
func (o *MbsSessPolCtrlData) HasMaxMbsSessionAmbr() bool {
	if o != nil && !isNil(o.MaxMbsSessionAmbr) {
		return true
	}

	return false
}

// SetMaxMbsSessionAmbr gets a reference to the given string and assigns it to the MaxMbsSessionAmbr field.
func (o *MbsSessPolCtrlData) SetMaxMbsSessionAmbr(v string) {
	o.MaxMbsSessionAmbr = &v
}

// GetMaxGbr returns the MaxGbr field value if set, zero value otherwise.
func (o *MbsSessPolCtrlData) GetMaxGbr() string {
	if o == nil || isNil(o.MaxGbr) {
		var ret string
		return ret
	}
	return *o.MaxGbr
}

// GetMaxGbrOk returns a tuple with the MaxGbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessPolCtrlData) GetMaxGbrOk() (*string, bool) {
	if o == nil || isNil(o.MaxGbr) {
    return nil, false
	}
	return o.MaxGbr, true
}

// HasMaxGbr returns a boolean if a field has been set.
func (o *MbsSessPolCtrlData) HasMaxGbr() bool {
	if o != nil && !isNil(o.MaxGbr) {
		return true
	}

	return false
}

// SetMaxGbr gets a reference to the given string and assigns it to the MaxGbr field.
func (o *MbsSessPolCtrlData) SetMaxGbr(v string) {
	o.MaxGbr = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *MbsSessPolCtrlData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessPolCtrlData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *MbsSessPolCtrlData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *MbsSessPolCtrlData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o MbsSessPolCtrlData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Var5qis) {
		toSerialize["5qis"] = o.Var5qis
	}
	if o.MaxMbsArpLevel.IsSet() {
		toSerialize["maxMbsArpLevel"] = o.MaxMbsArpLevel.Get()
	}
	if !isNil(o.MaxMbsSessionAmbr) {
		toSerialize["maxMbsSessionAmbr"] = o.MaxMbsSessionAmbr
	}
	if !isNil(o.MaxGbr) {
		toSerialize["maxGbr"] = o.MaxGbr
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableMbsSessPolCtrlData struct {
	value *MbsSessPolCtrlData
	isSet bool
}

func (v NullableMbsSessPolCtrlData) Get() *MbsSessPolCtrlData {
	return v.value
}

func (v *NullableMbsSessPolCtrlData) Set(val *MbsSessPolCtrlData) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessPolCtrlData) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessPolCtrlData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessPolCtrlData(val *MbsSessPolCtrlData) *NullableMbsSessPolCtrlData {
	return &NullableMbsSessPolCtrlData{value: val, isSet: true}
}

func (v NullableMbsSessPolCtrlData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessPolCtrlData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

