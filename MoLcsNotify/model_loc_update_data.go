/*
3gpp-mo-lcs-notify

API for UE updated location information notification.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MoLcsNotify

import (
	"encoding/json"
)

// LocUpdateData Represents a UE updated location information.
type LocUpdateData struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	LocInfo LocationInfo `json:"locInfo"`
	LcsQosClass LcsQosClass `json:"lcsQosClass"`
	// Contains the service identity
	SvcId *string `json:"svcId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewLocUpdateData instantiates a new LocUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocUpdateData(gpsi string, locInfo LocationInfo, lcsQosClass LcsQosClass, suppFeat string) *LocUpdateData {
	this := LocUpdateData{}
	this.Gpsi = gpsi
	this.LocInfo = locInfo
	this.LcsQosClass = lcsQosClass
	this.SuppFeat = suppFeat
	return &this
}

// NewLocUpdateDataWithDefaults instantiates a new LocUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocUpdateDataWithDefaults() *LocUpdateData {
	this := LocUpdateData{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *LocUpdateData) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetGpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *LocUpdateData) SetGpsi(v string) {
	o.Gpsi = v
}

// GetLocInfo returns the LocInfo field value
func (o *LocUpdateData) GetLocInfo() LocationInfo {
	if o == nil {
		var ret LocationInfo
		return ret
	}

	return o.LocInfo
}

// GetLocInfoOk returns a tuple with the LocInfo field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetLocInfoOk() (*LocationInfo, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocInfo, true
}

// SetLocInfo sets field value
func (o *LocUpdateData) SetLocInfo(v LocationInfo) {
	o.LocInfo = v
}

// GetLcsQosClass returns the LcsQosClass field value
func (o *LocUpdateData) GetLcsQosClass() LcsQosClass {
	if o == nil {
		var ret LcsQosClass
		return ret
	}

	return o.LcsQosClass
}

// GetLcsQosClassOk returns a tuple with the LcsQosClass field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetLcsQosClassOk() (*LcsQosClass, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LcsQosClass, true
}

// SetLcsQosClass sets field value
func (o *LocUpdateData) SetLcsQosClass(v LcsQosClass) {
	o.LcsQosClass = v
}

// GetSvcId returns the SvcId field value if set, zero value otherwise.
func (o *LocUpdateData) GetSvcId() string {
	if o == nil || isNil(o.SvcId) {
		var ret string
		return ret
	}
	return *o.SvcId
}

// GetSvcIdOk returns a tuple with the SvcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetSvcIdOk() (*string, bool) {
	if o == nil || isNil(o.SvcId) {
    return nil, false
	}
	return o.SvcId, true
}

// HasSvcId returns a boolean if a field has been set.
func (o *LocUpdateData) HasSvcId() bool {
	if o != nil && !isNil(o.SvcId) {
		return true
	}

	return false
}

// SetSvcId gets a reference to the given string and assigns it to the SvcId field.
func (o *LocUpdateData) SetSvcId(v string) {
	o.SvcId = &v
}

// GetSuppFeat returns the SuppFeat field value
func (o *LocUpdateData) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *LocUpdateData) GetSuppFeatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *LocUpdateData) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o LocUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gpsi"] = o.Gpsi
	}
	if true {
		toSerialize["locInfo"] = o.LocInfo
	}
	if true {
		toSerialize["lcsQosClass"] = o.LcsQosClass
	}
	if !isNil(o.SvcId) {
		toSerialize["svcId"] = o.SvcId
	}
	if true {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableLocUpdateData struct {
	value *LocUpdateData
	isSet bool
}

func (v NullableLocUpdateData) Get() *LocUpdateData {
	return v.value
}

func (v *NullableLocUpdateData) Set(val *LocUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableLocUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableLocUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocUpdateData(val *LocUpdateData) *NullableLocUpdateData {
	return &NullableLocUpdateData{value: val, isSet: true}
}

func (v NullableLocUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


