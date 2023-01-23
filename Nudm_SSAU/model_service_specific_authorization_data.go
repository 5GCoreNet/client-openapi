/*
Nudm_SSAU

Nudm Service Specific Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SSAU

import (
	"encoding/json"
	"time"
)

// ServiceSpecificAuthorizationData Authorization Response for a specific service.
type ServiceSpecificAuthorizationData struct {
	AuthorizationUeId *AuthorizationUeId `json:"authorizationUeId,omitempty"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExtGroupId *string `json:"extGroupId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
}

// NewServiceSpecificAuthorizationData instantiates a new ServiceSpecificAuthorizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSpecificAuthorizationData() *ServiceSpecificAuthorizationData {
	this := ServiceSpecificAuthorizationData{}
	return &this
}

// NewServiceSpecificAuthorizationDataWithDefaults instantiates a new ServiceSpecificAuthorizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSpecificAuthorizationDataWithDefaults() *ServiceSpecificAuthorizationData {
	this := ServiceSpecificAuthorizationData{}
	return &this
}

// GetAuthorizationUeId returns the AuthorizationUeId field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationData) GetAuthorizationUeId() AuthorizationUeId {
	if o == nil || isNil(o.AuthorizationUeId) {
		var ret AuthorizationUeId
		return ret
	}
	return *o.AuthorizationUeId
}

// GetAuthorizationUeIdOk returns a tuple with the AuthorizationUeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationData) GetAuthorizationUeIdOk() (*AuthorizationUeId, bool) {
	if o == nil || isNil(o.AuthorizationUeId) {
    return nil, false
	}
	return o.AuthorizationUeId, true
}

// HasAuthorizationUeId returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationData) HasAuthorizationUeId() bool {
	if o != nil && !isNil(o.AuthorizationUeId) {
		return true
	}

	return false
}

// SetAuthorizationUeId gets a reference to the given AuthorizationUeId and assigns it to the AuthorizationUeId field.
func (o *ServiceSpecificAuthorizationData) SetAuthorizationUeId(v AuthorizationUeId) {
	o.AuthorizationUeId = &v
}

// GetExtGroupId returns the ExtGroupId field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationData) GetExtGroupId() string {
	if o == nil || isNil(o.ExtGroupId) {
		var ret string
		return ret
	}
	return *o.ExtGroupId
}

// GetExtGroupIdOk returns a tuple with the ExtGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationData) GetExtGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ExtGroupId) {
    return nil, false
	}
	return o.ExtGroupId, true
}

// HasExtGroupId returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationData) HasExtGroupId() bool {
	if o != nil && !isNil(o.ExtGroupId) {
		return true
	}

	return false
}

// SetExtGroupId gets a reference to the given string and assigns it to the ExtGroupId field.
func (o *ServiceSpecificAuthorizationData) SetExtGroupId(v string) {
	o.ExtGroupId = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationData) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationData) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
    return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationData) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *ServiceSpecificAuthorizationData) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

func (o ServiceSpecificAuthorizationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthorizationUeId) {
		toSerialize["authorizationUeId"] = o.AuthorizationUeId
	}
	if !isNil(o.ExtGroupId) {
		toSerialize["extGroupId"] = o.ExtGroupId
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSpecificAuthorizationData struct {
	value *ServiceSpecificAuthorizationData
	isSet bool
}

func (v NullableServiceSpecificAuthorizationData) Get() *ServiceSpecificAuthorizationData {
	return v.value
}

func (v *NullableServiceSpecificAuthorizationData) Set(val *ServiceSpecificAuthorizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSpecificAuthorizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSpecificAuthorizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSpecificAuthorizationData(val *ServiceSpecificAuthorizationData) *NullableServiceSpecificAuthorizationData {
	return &NullableServiceSpecificAuthorizationData{value: val, isSet: true}
}

func (v NullableServiceSpecificAuthorizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSpecificAuthorizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


