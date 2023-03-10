/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"time"
)

// AuthorizationData NIDD Authorization Information
type AuthorizationData struct {
	AuthorizationData []UserIdentifier `json:"authorizationData"`
	AllowedDnnList []AccessAndMobilitySubscriptionDataSubscribedDnnListInner `json:"allowedDnnList,omitempty"`
	AllowedSnssaiList []Snssai `json:"allowedSnssaiList,omitempty"`
	AllowedMtcProviders []MtcProvider `json:"allowedMtcProviders,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
}

// NewAuthorizationData instantiates a new AuthorizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationData(authorizationData []UserIdentifier) *AuthorizationData {
	this := AuthorizationData{}
	this.AuthorizationData = authorizationData
	return &this
}

// NewAuthorizationDataWithDefaults instantiates a new AuthorizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataWithDefaults() *AuthorizationData {
	this := AuthorizationData{}
	return &this
}

// GetAuthorizationData returns the AuthorizationData field value
func (o *AuthorizationData) GetAuthorizationData() []UserIdentifier {
	if o == nil {
		var ret []UserIdentifier
		return ret
	}

	return o.AuthorizationData
}

// GetAuthorizationDataOk returns a tuple with the AuthorizationData field value
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetAuthorizationDataOk() ([]UserIdentifier, bool) {
	if o == nil {
    return nil, false
	}
	return o.AuthorizationData, true
}

// SetAuthorizationData sets field value
func (o *AuthorizationData) SetAuthorizationData(v []UserIdentifier) {
	o.AuthorizationData = v
}

// GetAllowedDnnList returns the AllowedDnnList field value if set, zero value otherwise.
func (o *AuthorizationData) GetAllowedDnnList() []AccessAndMobilitySubscriptionDataSubscribedDnnListInner {
	if o == nil || isNil(o.AllowedDnnList) {
		var ret []AccessAndMobilitySubscriptionDataSubscribedDnnListInner
		return ret
	}
	return o.AllowedDnnList
}

// GetAllowedDnnListOk returns a tuple with the AllowedDnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetAllowedDnnListOk() ([]AccessAndMobilitySubscriptionDataSubscribedDnnListInner, bool) {
	if o == nil || isNil(o.AllowedDnnList) {
    return nil, false
	}
	return o.AllowedDnnList, true
}

// HasAllowedDnnList returns a boolean if a field has been set.
func (o *AuthorizationData) HasAllowedDnnList() bool {
	if o != nil && !isNil(o.AllowedDnnList) {
		return true
	}

	return false
}

// SetAllowedDnnList gets a reference to the given []AccessAndMobilitySubscriptionDataSubscribedDnnListInner and assigns it to the AllowedDnnList field.
func (o *AuthorizationData) SetAllowedDnnList(v []AccessAndMobilitySubscriptionDataSubscribedDnnListInner) {
	o.AllowedDnnList = v
}

// GetAllowedSnssaiList returns the AllowedSnssaiList field value if set, zero value otherwise.
func (o *AuthorizationData) GetAllowedSnssaiList() []Snssai {
	if o == nil || isNil(o.AllowedSnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.AllowedSnssaiList
}

// GetAllowedSnssaiListOk returns a tuple with the AllowedSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetAllowedSnssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.AllowedSnssaiList) {
    return nil, false
	}
	return o.AllowedSnssaiList, true
}

// HasAllowedSnssaiList returns a boolean if a field has been set.
func (o *AuthorizationData) HasAllowedSnssaiList() bool {
	if o != nil && !isNil(o.AllowedSnssaiList) {
		return true
	}

	return false
}

// SetAllowedSnssaiList gets a reference to the given []Snssai and assigns it to the AllowedSnssaiList field.
func (o *AuthorizationData) SetAllowedSnssaiList(v []Snssai) {
	o.AllowedSnssaiList = v
}

// GetAllowedMtcProviders returns the AllowedMtcProviders field value if set, zero value otherwise.
func (o *AuthorizationData) GetAllowedMtcProviders() []MtcProvider {
	if o == nil || isNil(o.AllowedMtcProviders) {
		var ret []MtcProvider
		return ret
	}
	return o.AllowedMtcProviders
}

// GetAllowedMtcProvidersOk returns a tuple with the AllowedMtcProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetAllowedMtcProvidersOk() ([]MtcProvider, bool) {
	if o == nil || isNil(o.AllowedMtcProviders) {
    return nil, false
	}
	return o.AllowedMtcProviders, true
}

// HasAllowedMtcProviders returns a boolean if a field has been set.
func (o *AuthorizationData) HasAllowedMtcProviders() bool {
	if o != nil && !isNil(o.AllowedMtcProviders) {
		return true
	}

	return false
}

// SetAllowedMtcProviders gets a reference to the given []MtcProvider and assigns it to the AllowedMtcProviders field.
func (o *AuthorizationData) SetAllowedMtcProviders(v []MtcProvider) {
	o.AllowedMtcProviders = v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *AuthorizationData) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationData) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
    return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *AuthorizationData) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *AuthorizationData) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

func (o AuthorizationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authorizationData"] = o.AuthorizationData
	}
	if !isNil(o.AllowedDnnList) {
		toSerialize["allowedDnnList"] = o.AllowedDnnList
	}
	if !isNil(o.AllowedSnssaiList) {
		toSerialize["allowedSnssaiList"] = o.AllowedSnssaiList
	}
	if !isNil(o.AllowedMtcProviders) {
		toSerialize["allowedMtcProviders"] = o.AllowedMtcProviders
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationData struct {
	value *AuthorizationData
	isSet bool
}

func (v NullableAuthorizationData) Get() *AuthorizationData {
	return v.value
}

func (v *NullableAuthorizationData) Set(val *AuthorizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationData(val *AuthorizationData) *NullableAuthorizationData {
	return &NullableAuthorizationData{value: val, isSet: true}
}

func (v NullableAuthorizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


