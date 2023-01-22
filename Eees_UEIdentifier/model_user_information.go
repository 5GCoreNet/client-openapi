/*
EES UE Identifier Service

EES UE Identifier Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_UEIdentifier

import (
	"encoding/json"
)

// UserInformation Represents information about the User or the UE, that used by EES to use 3GPP CN capability  to retrieve the EAS specific UE identifier. 
type UserInformation struct {
	EasId string `json:"easId"`
	EasProviderId *string `json:"easProviderId,omitempty"`
	IpAddr IpAddr `json:"ipAddr"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewUserInformation instantiates a new UserInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInformation(easId string, ipAddr IpAddr) *UserInformation {
	this := UserInformation{}
	this.EasId = easId
	this.IpAddr = ipAddr
	return &this
}

// NewUserInformationWithDefaults instantiates a new UserInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInformationWithDefaults() *UserInformation {
	this := UserInformation{}
	return &this
}

// GetEasId returns the EasId field value
func (o *UserInformation) GetEasId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value
// and a boolean to check if the value has been set.
func (o *UserInformation) GetEasIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EasId, true
}

// SetEasId sets field value
func (o *UserInformation) SetEasId(v string) {
	o.EasId = v
}

// GetEasProviderId returns the EasProviderId field value if set, zero value otherwise.
func (o *UserInformation) GetEasProviderId() string {
	if o == nil || isNil(o.EasProviderId) {
		var ret string
		return ret
	}
	return *o.EasProviderId
}

// GetEasProviderIdOk returns a tuple with the EasProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformation) GetEasProviderIdOk() (*string, bool) {
	if o == nil || isNil(o.EasProviderId) {
    return nil, false
	}
	return o.EasProviderId, true
}

// HasEasProviderId returns a boolean if a field has been set.
func (o *UserInformation) HasEasProviderId() bool {
	if o != nil && !isNil(o.EasProviderId) {
		return true
	}

	return false
}

// SetEasProviderId gets a reference to the given string and assigns it to the EasProviderId field.
func (o *UserInformation) SetEasProviderId(v string) {
	o.EasProviderId = &v
}

// GetIpAddr returns the IpAddr field value
func (o *UserInformation) GetIpAddr() IpAddr {
	if o == nil {
		var ret IpAddr
		return ret
	}

	return o.IpAddr
}

// GetIpAddrOk returns a tuple with the IpAddr field value
// and a boolean to check if the value has been set.
func (o *UserInformation) GetIpAddrOk() (*IpAddr, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IpAddr, true
}

// SetIpAddr sets field value
func (o *UserInformation) SetIpAddr(v IpAddr) {
	o.IpAddr = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *UserInformation) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformation) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *UserInformation) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *UserInformation) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o UserInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["easId"] = o.EasId
	}
	if !isNil(o.EasProviderId) {
		toSerialize["easProviderId"] = o.EasProviderId
	}
	if true {
		toSerialize["ipAddr"] = o.IpAddr
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableUserInformation struct {
	value *UserInformation
	isSet bool
}

func (v NullableUserInformation) Get() *UserInformation {
	return v.value
}

func (v *NullableUserInformation) Set(val *UserInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInformation(val *UserInformation) *NullableUserInformation {
	return &NullableUserInformation{value: val, isSet: true}
}

func (v NullableUserInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

