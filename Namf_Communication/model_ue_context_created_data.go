/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// UeContextCreatedData Data within a successful response for creating an individual ueContext resource
type UeContextCreatedData struct {
	UeContext UeContext `json:"ueContext"`
	TargetToSourceData N2InfoContent `json:"targetToSourceData"`
	PduSessionList []N2SmInformation `json:"pduSessionList"`
	FailedSessionList []N2SmInformation `json:"failedSessionList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	PcfReselectedInd *bool `json:"pcfReselectedInd,omitempty"`
	AnalyticsNotUsedList []string `json:"analyticsNotUsedList,omitempty"`
}

// NewUeContextCreatedData instantiates a new UeContextCreatedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextCreatedData(ueContext UeContext, targetToSourceData N2InfoContent, pduSessionList []N2SmInformation) *UeContextCreatedData {
	this := UeContextCreatedData{}
	this.UeContext = ueContext
	this.TargetToSourceData = targetToSourceData
	this.PduSessionList = pduSessionList
	return &this
}

// NewUeContextCreatedDataWithDefaults instantiates a new UeContextCreatedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextCreatedDataWithDefaults() *UeContextCreatedData {
	this := UeContextCreatedData{}
	return &this
}

// GetUeContext returns the UeContext field value
func (o *UeContextCreatedData) GetUeContext() UeContext {
	if o == nil {
		var ret UeContext
		return ret
	}

	return o.UeContext
}

// GetUeContextOk returns a tuple with the UeContext field value
// and a boolean to check if the value has been set.
func (o *UeContextCreatedData) GetUeContextOk() (*UeContext, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UeContext, true
}

// SetUeContext sets field value
func (o *UeContextCreatedData) SetUeContext(v UeContext) {
	o.UeContext = v
}

// GetTargetToSourceData returns the TargetToSourceData field value
func (o *UeContextCreatedData) GetTargetToSourceData() N2InfoContent {
	if o == nil {
		var ret N2InfoContent
		return ret
	}

	return o.TargetToSourceData
}

// GetTargetToSourceDataOk returns a tuple with the TargetToSourceData field value
// and a boolean to check if the value has been set.
func (o *UeContextCreatedData) GetTargetToSourceDataOk() (*N2InfoContent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TargetToSourceData, true
}

// SetTargetToSourceData sets field value
func (o *UeContextCreatedData) SetTargetToSourceData(v N2InfoContent) {
	o.TargetToSourceData = v
}

// GetPduSessionList returns the PduSessionList field value
func (o *UeContextCreatedData) GetPduSessionList() []N2SmInformation {
	if o == nil {
		var ret []N2SmInformation
		return ret
	}

	return o.PduSessionList
}

// GetPduSessionListOk returns a tuple with the PduSessionList field value
// and a boolean to check if the value has been set.
func (o *UeContextCreatedData) GetPduSessionListOk() ([]N2SmInformation, bool) {
	if o == nil {
    return nil, false
	}
	return o.PduSessionList, true
}

// SetPduSessionList sets field value
func (o *UeContextCreatedData) SetPduSessionList(v []N2SmInformation) {
	o.PduSessionList = v
}

// GetFailedSessionList returns the FailedSessionList field value if set, zero value otherwise.
func (o *UeContextCreatedData) GetFailedSessionList() []N2SmInformation {
	if o == nil || isNil(o.FailedSessionList) {
		var ret []N2SmInformation
		return ret
	}
	return o.FailedSessionList
}

// GetFailedSessionListOk returns a tuple with the FailedSessionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreatedData) GetFailedSessionListOk() ([]N2SmInformation, bool) {
	if o == nil || isNil(o.FailedSessionList) {
    return nil, false
	}
	return o.FailedSessionList, true
}

// HasFailedSessionList returns a boolean if a field has been set.
func (o *UeContextCreatedData) HasFailedSessionList() bool {
	if o != nil && !isNil(o.FailedSessionList) {
		return true
	}

	return false
}

// SetFailedSessionList gets a reference to the given []N2SmInformation and assigns it to the FailedSessionList field.
func (o *UeContextCreatedData) SetFailedSessionList(v []N2SmInformation) {
	o.FailedSessionList = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *UeContextCreatedData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreatedData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *UeContextCreatedData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *UeContextCreatedData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetPcfReselectedInd returns the PcfReselectedInd field value if set, zero value otherwise.
func (o *UeContextCreatedData) GetPcfReselectedInd() bool {
	if o == nil || isNil(o.PcfReselectedInd) {
		var ret bool
		return ret
	}
	return *o.PcfReselectedInd
}

// GetPcfReselectedIndOk returns a tuple with the PcfReselectedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreatedData) GetPcfReselectedIndOk() (*bool, bool) {
	if o == nil || isNil(o.PcfReselectedInd) {
    return nil, false
	}
	return o.PcfReselectedInd, true
}

// HasPcfReselectedInd returns a boolean if a field has been set.
func (o *UeContextCreatedData) HasPcfReselectedInd() bool {
	if o != nil && !isNil(o.PcfReselectedInd) {
		return true
	}

	return false
}

// SetPcfReselectedInd gets a reference to the given bool and assigns it to the PcfReselectedInd field.
func (o *UeContextCreatedData) SetPcfReselectedInd(v bool) {
	o.PcfReselectedInd = &v
}

// GetAnalyticsNotUsedList returns the AnalyticsNotUsedList field value if set, zero value otherwise.
func (o *UeContextCreatedData) GetAnalyticsNotUsedList() []string {
	if o == nil || isNil(o.AnalyticsNotUsedList) {
		var ret []string
		return ret
	}
	return o.AnalyticsNotUsedList
}

// GetAnalyticsNotUsedListOk returns a tuple with the AnalyticsNotUsedList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextCreatedData) GetAnalyticsNotUsedListOk() ([]string, bool) {
	if o == nil || isNil(o.AnalyticsNotUsedList) {
    return nil, false
	}
	return o.AnalyticsNotUsedList, true
}

// HasAnalyticsNotUsedList returns a boolean if a field has been set.
func (o *UeContextCreatedData) HasAnalyticsNotUsedList() bool {
	if o != nil && !isNil(o.AnalyticsNotUsedList) {
		return true
	}

	return false
}

// SetAnalyticsNotUsedList gets a reference to the given []string and assigns it to the AnalyticsNotUsedList field.
func (o *UeContextCreatedData) SetAnalyticsNotUsedList(v []string) {
	o.AnalyticsNotUsedList = v
}

func (o UeContextCreatedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ueContext"] = o.UeContext
	}
	if true {
		toSerialize["targetToSourceData"] = o.TargetToSourceData
	}
	if true {
		toSerialize["pduSessionList"] = o.PduSessionList
	}
	if !isNil(o.FailedSessionList) {
		toSerialize["failedSessionList"] = o.FailedSessionList
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.PcfReselectedInd) {
		toSerialize["pcfReselectedInd"] = o.PcfReselectedInd
	}
	if !isNil(o.AnalyticsNotUsedList) {
		toSerialize["analyticsNotUsedList"] = o.AnalyticsNotUsedList
	}
	return json.Marshal(toSerialize)
}

type NullableUeContextCreatedData struct {
	value *UeContextCreatedData
	isSet bool
}

func (v NullableUeContextCreatedData) Get() *UeContextCreatedData {
	return v.value
}

func (v *NullableUeContextCreatedData) Set(val *UeContextCreatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextCreatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextCreatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextCreatedData(val *UeContextCreatedData) *NullableUeContextCreatedData {
	return &NullableUeContextCreatedData{value: val, isSet: true}
}

func (v NullableUeContextCreatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextCreatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


