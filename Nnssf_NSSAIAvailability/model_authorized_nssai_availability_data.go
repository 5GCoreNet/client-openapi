/*
NSSF NSSAI Availability

NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnssf_NSSAIAvailability

import (
	"encoding/json"
)

// AuthorizedNssaiAvailabilityData This contains the Nssai availability data information per TA authorized by the NSSF
type AuthorizedNssaiAvailabilityData struct {
	Tai Tai `json:"tai"`
	SupportedSnssaiList []ExtSnssai `json:"supportedSnssaiList"`
	RestrictedSnssaiList []RestrictedSnssai `json:"restrictedSnssaiList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	NsagInfos []NsagInfo `json:"nsagInfos,omitempty"`
}

// NewAuthorizedNssaiAvailabilityData instantiates a new AuthorizedNssaiAvailabilityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizedNssaiAvailabilityData(tai Tai, supportedSnssaiList []ExtSnssai) *AuthorizedNssaiAvailabilityData {
	this := AuthorizedNssaiAvailabilityData{}
	this.Tai = tai
	this.SupportedSnssaiList = supportedSnssaiList
	return &this
}

// NewAuthorizedNssaiAvailabilityDataWithDefaults instantiates a new AuthorizedNssaiAvailabilityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizedNssaiAvailabilityDataWithDefaults() *AuthorizedNssaiAvailabilityData {
	this := AuthorizedNssaiAvailabilityData{}
	return &this
}

// GetTai returns the Tai field value
func (o *AuthorizedNssaiAvailabilityData) GetTai() Tai {
	if o == nil {
		var ret Tai
		return ret
	}

	return o.Tai
}

// GetTaiOk returns a tuple with the Tai field value
// and a boolean to check if the value has been set.
func (o *AuthorizedNssaiAvailabilityData) GetTaiOk() (*Tai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tai, true
}

// SetTai sets field value
func (o *AuthorizedNssaiAvailabilityData) SetTai(v Tai) {
	o.Tai = v
}

// GetSupportedSnssaiList returns the SupportedSnssaiList field value
func (o *AuthorizedNssaiAvailabilityData) GetSupportedSnssaiList() []ExtSnssai {
	if o == nil {
		var ret []ExtSnssai
		return ret
	}

	return o.SupportedSnssaiList
}

// GetSupportedSnssaiListOk returns a tuple with the SupportedSnssaiList field value
// and a boolean to check if the value has been set.
func (o *AuthorizedNssaiAvailabilityData) GetSupportedSnssaiListOk() ([]ExtSnssai, bool) {
	if o == nil {
    return nil, false
	}
	return o.SupportedSnssaiList, true
}

// SetSupportedSnssaiList sets field value
func (o *AuthorizedNssaiAvailabilityData) SetSupportedSnssaiList(v []ExtSnssai) {
	o.SupportedSnssaiList = v
}

// GetRestrictedSnssaiList returns the RestrictedSnssaiList field value if set, zero value otherwise.
func (o *AuthorizedNssaiAvailabilityData) GetRestrictedSnssaiList() []RestrictedSnssai {
	if o == nil || isNil(o.RestrictedSnssaiList) {
		var ret []RestrictedSnssai
		return ret
	}
	return o.RestrictedSnssaiList
}

// GetRestrictedSnssaiListOk returns a tuple with the RestrictedSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNssaiAvailabilityData) GetRestrictedSnssaiListOk() ([]RestrictedSnssai, bool) {
	if o == nil || isNil(o.RestrictedSnssaiList) {
    return nil, false
	}
	return o.RestrictedSnssaiList, true
}

// HasRestrictedSnssaiList returns a boolean if a field has been set.
func (o *AuthorizedNssaiAvailabilityData) HasRestrictedSnssaiList() bool {
	if o != nil && !isNil(o.RestrictedSnssaiList) {
		return true
	}

	return false
}

// SetRestrictedSnssaiList gets a reference to the given []RestrictedSnssai and assigns it to the RestrictedSnssaiList field.
func (o *AuthorizedNssaiAvailabilityData) SetRestrictedSnssaiList(v []RestrictedSnssai) {
	o.RestrictedSnssaiList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *AuthorizedNssaiAvailabilityData) GetTaiList() []Tai {
	if o == nil || isNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNssaiAvailabilityData) GetTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaiList) {
    return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *AuthorizedNssaiAvailabilityData) HasTaiList() bool {
	if o != nil && !isNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *AuthorizedNssaiAvailabilityData) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *AuthorizedNssaiAvailabilityData) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNssaiAvailabilityData) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *AuthorizedNssaiAvailabilityData) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *AuthorizedNssaiAvailabilityData) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetNsagInfos returns the NsagInfos field value if set, zero value otherwise.
func (o *AuthorizedNssaiAvailabilityData) GetNsagInfos() []NsagInfo {
	if o == nil || isNil(o.NsagInfos) {
		var ret []NsagInfo
		return ret
	}
	return o.NsagInfos
}

// GetNsagInfosOk returns a tuple with the NsagInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedNssaiAvailabilityData) GetNsagInfosOk() ([]NsagInfo, bool) {
	if o == nil || isNil(o.NsagInfos) {
    return nil, false
	}
	return o.NsagInfos, true
}

// HasNsagInfos returns a boolean if a field has been set.
func (o *AuthorizedNssaiAvailabilityData) HasNsagInfos() bool {
	if o != nil && !isNil(o.NsagInfos) {
		return true
	}

	return false
}

// SetNsagInfos gets a reference to the given []NsagInfo and assigns it to the NsagInfos field.
func (o *AuthorizedNssaiAvailabilityData) SetNsagInfos(v []NsagInfo) {
	o.NsagInfos = v
}

func (o AuthorizedNssaiAvailabilityData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tai"] = o.Tai
	}
	if true {
		toSerialize["supportedSnssaiList"] = o.SupportedSnssaiList
	}
	if !isNil(o.RestrictedSnssaiList) {
		toSerialize["restrictedSnssaiList"] = o.RestrictedSnssaiList
	}
	if !isNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.NsagInfos) {
		toSerialize["nsagInfos"] = o.NsagInfos
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizedNssaiAvailabilityData struct {
	value *AuthorizedNssaiAvailabilityData
	isSet bool
}

func (v NullableAuthorizedNssaiAvailabilityData) Get() *AuthorizedNssaiAvailabilityData {
	return v.value
}

func (v *NullableAuthorizedNssaiAvailabilityData) Set(val *AuthorizedNssaiAvailabilityData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizedNssaiAvailabilityData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizedNssaiAvailabilityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizedNssaiAvailabilityData(val *AuthorizedNssaiAvailabilityData) *NullableAuthorizedNssaiAvailabilityData {
	return &NullableAuthorizedNssaiAvailabilityData{value: val, isSet: true}
}

func (v NullableAuthorizedNssaiAvailabilityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizedNssaiAvailabilityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


