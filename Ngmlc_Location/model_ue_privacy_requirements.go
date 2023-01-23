/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
)

// UePrivacyRequirements UE privacy requirements from (H)GMLC to the serving AMF or VGMLC(in the roaming case) for the target UE
type UePrivacyRequirements struct {
	LcsServiceAuthInfo *LcsServiceAuth `json:"lcsServiceAuthInfo,omitempty"`
	CodeWordCheck *bool `json:"codeWordCheck,omitempty"`
}

// NewUePrivacyRequirements instantiates a new UePrivacyRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUePrivacyRequirements() *UePrivacyRequirements {
	this := UePrivacyRequirements{}
	return &this
}

// NewUePrivacyRequirementsWithDefaults instantiates a new UePrivacyRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUePrivacyRequirementsWithDefaults() *UePrivacyRequirements {
	this := UePrivacyRequirements{}
	return &this
}

// GetLcsServiceAuthInfo returns the LcsServiceAuthInfo field value if set, zero value otherwise.
func (o *UePrivacyRequirements) GetLcsServiceAuthInfo() LcsServiceAuth {
	if o == nil || isNil(o.LcsServiceAuthInfo) {
		var ret LcsServiceAuth
		return ret
	}
	return *o.LcsServiceAuthInfo
}

// GetLcsServiceAuthInfoOk returns a tuple with the LcsServiceAuthInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UePrivacyRequirements) GetLcsServiceAuthInfoOk() (*LcsServiceAuth, bool) {
	if o == nil || isNil(o.LcsServiceAuthInfo) {
    return nil, false
	}
	return o.LcsServiceAuthInfo, true
}

// HasLcsServiceAuthInfo returns a boolean if a field has been set.
func (o *UePrivacyRequirements) HasLcsServiceAuthInfo() bool {
	if o != nil && !isNil(o.LcsServiceAuthInfo) {
		return true
	}

	return false
}

// SetLcsServiceAuthInfo gets a reference to the given LcsServiceAuth and assigns it to the LcsServiceAuthInfo field.
func (o *UePrivacyRequirements) SetLcsServiceAuthInfo(v LcsServiceAuth) {
	o.LcsServiceAuthInfo = &v
}

// GetCodeWordCheck returns the CodeWordCheck field value if set, zero value otherwise.
func (o *UePrivacyRequirements) GetCodeWordCheck() bool {
	if o == nil || isNil(o.CodeWordCheck) {
		var ret bool
		return ret
	}
	return *o.CodeWordCheck
}

// GetCodeWordCheckOk returns a tuple with the CodeWordCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UePrivacyRequirements) GetCodeWordCheckOk() (*bool, bool) {
	if o == nil || isNil(o.CodeWordCheck) {
    return nil, false
	}
	return o.CodeWordCheck, true
}

// HasCodeWordCheck returns a boolean if a field has been set.
func (o *UePrivacyRequirements) HasCodeWordCheck() bool {
	if o != nil && !isNil(o.CodeWordCheck) {
		return true
	}

	return false
}

// SetCodeWordCheck gets a reference to the given bool and assigns it to the CodeWordCheck field.
func (o *UePrivacyRequirements) SetCodeWordCheck(v bool) {
	o.CodeWordCheck = &v
}

func (o UePrivacyRequirements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LcsServiceAuthInfo) {
		toSerialize["lcsServiceAuthInfo"] = o.LcsServiceAuthInfo
	}
	if !isNil(o.CodeWordCheck) {
		toSerialize["codeWordCheck"] = o.CodeWordCheck
	}
	return json.Marshal(toSerialize)
}

type NullableUePrivacyRequirements struct {
	value *UePrivacyRequirements
	isSet bool
}

func (v NullableUePrivacyRequirements) Get() *UePrivacyRequirements {
	return v.value
}

func (v *NullableUePrivacyRequirements) Set(val *UePrivacyRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableUePrivacyRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableUePrivacyRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUePrivacyRequirements(val *UePrivacyRequirements) *NullableUePrivacyRequirements {
	return &NullableUePrivacyRequirements{value: val, isSet: true}
}

func (v NullableUePrivacyRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUePrivacyRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


