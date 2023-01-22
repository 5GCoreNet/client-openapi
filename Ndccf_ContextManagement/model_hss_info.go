/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
)

// HssInfo Information of an HSS NF Instance
type HssInfo struct {
	// Identifier of a group of NFs.
	GroupId *string `json:"groupId,omitempty"`
	ImsiRanges []ImsiRange `json:"imsiRanges,omitempty"`
	ImsPrivateIdentityRanges []IdentityRange `json:"imsPrivateIdentityRanges,omitempty"`
	ImsPublicIdentityRanges []IdentityRange `json:"imsPublicIdentityRanges,omitempty"`
	MsisdnRanges []IdentityRange `json:"msisdnRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	HssDiameterAddress *NetworkNodeDiameterAddress `json:"hssDiameterAddress,omitempty"`
	AdditionalDiamAddresses []NetworkNodeDiameterAddress `json:"additionalDiamAddresses,omitempty"`
}

// NewHssInfo instantiates a new HssInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHssInfo() *HssInfo {
	this := HssInfo{}
	return &this
}

// NewHssInfoWithDefaults instantiates a new HssInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHssInfoWithDefaults() *HssInfo {
	this := HssInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *HssInfo) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *HssInfo) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *HssInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetImsiRanges returns the ImsiRanges field value if set, zero value otherwise.
func (o *HssInfo) GetImsiRanges() []ImsiRange {
	if o == nil || isNil(o.ImsiRanges) {
		var ret []ImsiRange
		return ret
	}
	return o.ImsiRanges
}

// GetImsiRangesOk returns a tuple with the ImsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetImsiRangesOk() ([]ImsiRange, bool) {
	if o == nil || isNil(o.ImsiRanges) {
    return nil, false
	}
	return o.ImsiRanges, true
}

// HasImsiRanges returns a boolean if a field has been set.
func (o *HssInfo) HasImsiRanges() bool {
	if o != nil && !isNil(o.ImsiRanges) {
		return true
	}

	return false
}

// SetImsiRanges gets a reference to the given []ImsiRange and assigns it to the ImsiRanges field.
func (o *HssInfo) SetImsiRanges(v []ImsiRange) {
	o.ImsiRanges = v
}

// GetImsPrivateIdentityRanges returns the ImsPrivateIdentityRanges field value if set, zero value otherwise.
func (o *HssInfo) GetImsPrivateIdentityRanges() []IdentityRange {
	if o == nil || isNil(o.ImsPrivateIdentityRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ImsPrivateIdentityRanges
}

// GetImsPrivateIdentityRangesOk returns a tuple with the ImsPrivateIdentityRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetImsPrivateIdentityRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.ImsPrivateIdentityRanges) {
    return nil, false
	}
	return o.ImsPrivateIdentityRanges, true
}

// HasImsPrivateIdentityRanges returns a boolean if a field has been set.
func (o *HssInfo) HasImsPrivateIdentityRanges() bool {
	if o != nil && !isNil(o.ImsPrivateIdentityRanges) {
		return true
	}

	return false
}

// SetImsPrivateIdentityRanges gets a reference to the given []IdentityRange and assigns it to the ImsPrivateIdentityRanges field.
func (o *HssInfo) SetImsPrivateIdentityRanges(v []IdentityRange) {
	o.ImsPrivateIdentityRanges = v
}

// GetImsPublicIdentityRanges returns the ImsPublicIdentityRanges field value if set, zero value otherwise.
func (o *HssInfo) GetImsPublicIdentityRanges() []IdentityRange {
	if o == nil || isNil(o.ImsPublicIdentityRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ImsPublicIdentityRanges
}

// GetImsPublicIdentityRangesOk returns a tuple with the ImsPublicIdentityRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetImsPublicIdentityRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.ImsPublicIdentityRanges) {
    return nil, false
	}
	return o.ImsPublicIdentityRanges, true
}

// HasImsPublicIdentityRanges returns a boolean if a field has been set.
func (o *HssInfo) HasImsPublicIdentityRanges() bool {
	if o != nil && !isNil(o.ImsPublicIdentityRanges) {
		return true
	}

	return false
}

// SetImsPublicIdentityRanges gets a reference to the given []IdentityRange and assigns it to the ImsPublicIdentityRanges field.
func (o *HssInfo) SetImsPublicIdentityRanges(v []IdentityRange) {
	o.ImsPublicIdentityRanges = v
}

// GetMsisdnRanges returns the MsisdnRanges field value if set, zero value otherwise.
func (o *HssInfo) GetMsisdnRanges() []IdentityRange {
	if o == nil || isNil(o.MsisdnRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.MsisdnRanges
}

// GetMsisdnRangesOk returns a tuple with the MsisdnRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetMsisdnRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.MsisdnRanges) {
    return nil, false
	}
	return o.MsisdnRanges, true
}

// HasMsisdnRanges returns a boolean if a field has been set.
func (o *HssInfo) HasMsisdnRanges() bool {
	if o != nil && !isNil(o.MsisdnRanges) {
		return true
	}

	return false
}

// SetMsisdnRanges gets a reference to the given []IdentityRange and assigns it to the MsisdnRanges field.
func (o *HssInfo) SetMsisdnRanges(v []IdentityRange) {
	o.MsisdnRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *HssInfo) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || isNil(o.ExternalGroupIdentifiersRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.ExternalGroupIdentifiersRanges) {
    return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *HssInfo) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && !isNil(o.ExternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *HssInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetHssDiameterAddress returns the HssDiameterAddress field value if set, zero value otherwise.
func (o *HssInfo) GetHssDiameterAddress() NetworkNodeDiameterAddress {
	if o == nil || isNil(o.HssDiameterAddress) {
		var ret NetworkNodeDiameterAddress
		return ret
	}
	return *o.HssDiameterAddress
}

// GetHssDiameterAddressOk returns a tuple with the HssDiameterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetHssDiameterAddressOk() (*NetworkNodeDiameterAddress, bool) {
	if o == nil || isNil(o.HssDiameterAddress) {
    return nil, false
	}
	return o.HssDiameterAddress, true
}

// HasHssDiameterAddress returns a boolean if a field has been set.
func (o *HssInfo) HasHssDiameterAddress() bool {
	if o != nil && !isNil(o.HssDiameterAddress) {
		return true
	}

	return false
}

// SetHssDiameterAddress gets a reference to the given NetworkNodeDiameterAddress and assigns it to the HssDiameterAddress field.
func (o *HssInfo) SetHssDiameterAddress(v NetworkNodeDiameterAddress) {
	o.HssDiameterAddress = &v
}

// GetAdditionalDiamAddresses returns the AdditionalDiamAddresses field value if set, zero value otherwise.
func (o *HssInfo) GetAdditionalDiamAddresses() []NetworkNodeDiameterAddress {
	if o == nil || isNil(o.AdditionalDiamAddresses) {
		var ret []NetworkNodeDiameterAddress
		return ret
	}
	return o.AdditionalDiamAddresses
}

// GetAdditionalDiamAddressesOk returns a tuple with the AdditionalDiamAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssInfo) GetAdditionalDiamAddressesOk() ([]NetworkNodeDiameterAddress, bool) {
	if o == nil || isNil(o.AdditionalDiamAddresses) {
    return nil, false
	}
	return o.AdditionalDiamAddresses, true
}

// HasAdditionalDiamAddresses returns a boolean if a field has been set.
func (o *HssInfo) HasAdditionalDiamAddresses() bool {
	if o != nil && !isNil(o.AdditionalDiamAddresses) {
		return true
	}

	return false
}

// SetAdditionalDiamAddresses gets a reference to the given []NetworkNodeDiameterAddress and assigns it to the AdditionalDiamAddresses field.
func (o *HssInfo) SetAdditionalDiamAddresses(v []NetworkNodeDiameterAddress) {
	o.AdditionalDiamAddresses = v
}

func (o HssInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.ImsiRanges) {
		toSerialize["imsiRanges"] = o.ImsiRanges
	}
	if !isNil(o.ImsPrivateIdentityRanges) {
		toSerialize["imsPrivateIdentityRanges"] = o.ImsPrivateIdentityRanges
	}
	if !isNil(o.ImsPublicIdentityRanges) {
		toSerialize["imsPublicIdentityRanges"] = o.ImsPublicIdentityRanges
	}
	if !isNil(o.MsisdnRanges) {
		toSerialize["msisdnRanges"] = o.MsisdnRanges
	}
	if !isNil(o.ExternalGroupIdentifiersRanges) {
		toSerialize["externalGroupIdentifiersRanges"] = o.ExternalGroupIdentifiersRanges
	}
	if !isNil(o.HssDiameterAddress) {
		toSerialize["hssDiameterAddress"] = o.HssDiameterAddress
	}
	if !isNil(o.AdditionalDiamAddresses) {
		toSerialize["additionalDiamAddresses"] = o.AdditionalDiamAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableHssInfo struct {
	value *HssInfo
	isSet bool
}

func (v NullableHssInfo) Get() *HssInfo {
	return v.value
}

func (v *NullableHssInfo) Set(val *HssInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHssInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHssInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHssInfo(val *HssInfo) *NullableHssInfo {
	return &NullableHssInfo{value: val, isSet: true}
}

func (v NullableHssInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHssInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


