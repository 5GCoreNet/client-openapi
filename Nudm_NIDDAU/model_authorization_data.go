/*
Nudm_NIDDAU

Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_NIDDAU

import (
	"encoding/json"
	"time"
)

// AuthorizationData Represents NIDD authorization data.
type AuthorizationData struct {
	AuthorizationData []UserIdentifier `json:"authorizationData"`
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


