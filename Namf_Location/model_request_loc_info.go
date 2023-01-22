/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Location

import (
	"encoding/json"
)

// RequestLocInfo Data within Provide Location Information Request
type RequestLocInfo struct {
	Req5gsLoc *bool `json:"req5gsLoc,omitempty"`
	ReqCurrentLoc *bool `json:"reqCurrentLoc,omitempty"`
	ReqRatType *bool `json:"reqRatType,omitempty"`
	ReqTimeZone *bool `json:"reqTimeZone,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewRequestLocInfo instantiates a new RequestLocInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestLocInfo() *RequestLocInfo {
	this := RequestLocInfo{}
	var req5gsLoc bool = false
	this.Req5gsLoc = &req5gsLoc
	var reqCurrentLoc bool = false
	this.ReqCurrentLoc = &reqCurrentLoc
	var reqRatType bool = false
	this.ReqRatType = &reqRatType
	var reqTimeZone bool = false
	this.ReqTimeZone = &reqTimeZone
	return &this
}

// NewRequestLocInfoWithDefaults instantiates a new RequestLocInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestLocInfoWithDefaults() *RequestLocInfo {
	this := RequestLocInfo{}
	var req5gsLoc bool = false
	this.Req5gsLoc = &req5gsLoc
	var reqCurrentLoc bool = false
	this.ReqCurrentLoc = &reqCurrentLoc
	var reqRatType bool = false
	this.ReqRatType = &reqRatType
	var reqTimeZone bool = false
	this.ReqTimeZone = &reqTimeZone
	return &this
}

// GetReq5gsLoc returns the Req5gsLoc field value if set, zero value otherwise.
func (o *RequestLocInfo) GetReq5gsLoc() bool {
	if o == nil || isNil(o.Req5gsLoc) {
		var ret bool
		return ret
	}
	return *o.Req5gsLoc
}

// GetReq5gsLocOk returns a tuple with the Req5gsLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLocInfo) GetReq5gsLocOk() (*bool, bool) {
	if o == nil || isNil(o.Req5gsLoc) {
    return nil, false
	}
	return o.Req5gsLoc, true
}

// HasReq5gsLoc returns a boolean if a field has been set.
func (o *RequestLocInfo) HasReq5gsLoc() bool {
	if o != nil && !isNil(o.Req5gsLoc) {
		return true
	}

	return false
}

// SetReq5gsLoc gets a reference to the given bool and assigns it to the Req5gsLoc field.
func (o *RequestLocInfo) SetReq5gsLoc(v bool) {
	o.Req5gsLoc = &v
}

// GetReqCurrentLoc returns the ReqCurrentLoc field value if set, zero value otherwise.
func (o *RequestLocInfo) GetReqCurrentLoc() bool {
	if o == nil || isNil(o.ReqCurrentLoc) {
		var ret bool
		return ret
	}
	return *o.ReqCurrentLoc
}

// GetReqCurrentLocOk returns a tuple with the ReqCurrentLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLocInfo) GetReqCurrentLocOk() (*bool, bool) {
	if o == nil || isNil(o.ReqCurrentLoc) {
    return nil, false
	}
	return o.ReqCurrentLoc, true
}

// HasReqCurrentLoc returns a boolean if a field has been set.
func (o *RequestLocInfo) HasReqCurrentLoc() bool {
	if o != nil && !isNil(o.ReqCurrentLoc) {
		return true
	}

	return false
}

// SetReqCurrentLoc gets a reference to the given bool and assigns it to the ReqCurrentLoc field.
func (o *RequestLocInfo) SetReqCurrentLoc(v bool) {
	o.ReqCurrentLoc = &v
}

// GetReqRatType returns the ReqRatType field value if set, zero value otherwise.
func (o *RequestLocInfo) GetReqRatType() bool {
	if o == nil || isNil(o.ReqRatType) {
		var ret bool
		return ret
	}
	return *o.ReqRatType
}

// GetReqRatTypeOk returns a tuple with the ReqRatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLocInfo) GetReqRatTypeOk() (*bool, bool) {
	if o == nil || isNil(o.ReqRatType) {
    return nil, false
	}
	return o.ReqRatType, true
}

// HasReqRatType returns a boolean if a field has been set.
func (o *RequestLocInfo) HasReqRatType() bool {
	if o != nil && !isNil(o.ReqRatType) {
		return true
	}

	return false
}

// SetReqRatType gets a reference to the given bool and assigns it to the ReqRatType field.
func (o *RequestLocInfo) SetReqRatType(v bool) {
	o.ReqRatType = &v
}

// GetReqTimeZone returns the ReqTimeZone field value if set, zero value otherwise.
func (o *RequestLocInfo) GetReqTimeZone() bool {
	if o == nil || isNil(o.ReqTimeZone) {
		var ret bool
		return ret
	}
	return *o.ReqTimeZone
}

// GetReqTimeZoneOk returns a tuple with the ReqTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLocInfo) GetReqTimeZoneOk() (*bool, bool) {
	if o == nil || isNil(o.ReqTimeZone) {
    return nil, false
	}
	return o.ReqTimeZone, true
}

// HasReqTimeZone returns a boolean if a field has been set.
func (o *RequestLocInfo) HasReqTimeZone() bool {
	if o != nil && !isNil(o.ReqTimeZone) {
		return true
	}

	return false
}

// SetReqTimeZone gets a reference to the given bool and assigns it to the ReqTimeZone field.
func (o *RequestLocInfo) SetReqTimeZone(v bool) {
	o.ReqTimeZone = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *RequestLocInfo) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLocInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *RequestLocInfo) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *RequestLocInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o RequestLocInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Req5gsLoc) {
		toSerialize["req5gsLoc"] = o.Req5gsLoc
	}
	if !isNil(o.ReqCurrentLoc) {
		toSerialize["reqCurrentLoc"] = o.ReqCurrentLoc
	}
	if !isNil(o.ReqRatType) {
		toSerialize["reqRatType"] = o.ReqRatType
	}
	if !isNil(o.ReqTimeZone) {
		toSerialize["reqTimeZone"] = o.ReqTimeZone
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableRequestLocInfo struct {
	value *RequestLocInfo
	isSet bool
}

func (v NullableRequestLocInfo) Get() *RequestLocInfo {
	return v.value
}

func (v *NullableRequestLocInfo) Set(val *RequestLocInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestLocInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestLocInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestLocInfo(val *RequestLocInfo) *NullableRequestLocInfo {
	return &NullableRequestLocInfo{value: val, isSet: true}
}

func (v NullableRequestLocInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestLocInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


