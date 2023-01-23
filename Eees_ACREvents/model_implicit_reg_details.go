/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACREvents

import (
	"encoding/json"
	"time"
)

// ImplicitRegDetails Represents the EEC implicit registration details.
type ImplicitRegDetails struct {
	// Identifier of the EEC registration.
	RegId string `json:"regId"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime *time.Time `json:"expTime,omitempty"`
}

// NewImplicitRegDetails instantiates a new ImplicitRegDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImplicitRegDetails(regId string) *ImplicitRegDetails {
	this := ImplicitRegDetails{}
	this.RegId = regId
	return &this
}

// NewImplicitRegDetailsWithDefaults instantiates a new ImplicitRegDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImplicitRegDetailsWithDefaults() *ImplicitRegDetails {
	this := ImplicitRegDetails{}
	return &this
}

// GetRegId returns the RegId field value
func (o *ImplicitRegDetails) GetRegId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegId
}

// GetRegIdOk returns a tuple with the RegId field value
// and a boolean to check if the value has been set.
func (o *ImplicitRegDetails) GetRegIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegId, true
}

// SetRegId sets field value
func (o *ImplicitRegDetails) SetRegId(v string) {
	o.RegId = v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *ImplicitRegDetails) GetExpTime() time.Time {
	if o == nil || isNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImplicitRegDetails) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpTime) {
    return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *ImplicitRegDetails) HasExpTime() bool {
	if o != nil && !isNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *ImplicitRegDetails) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

func (o ImplicitRegDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regId"] = o.RegId
	}
	if !isNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	return json.Marshal(toSerialize)
}

type NullableImplicitRegDetails struct {
	value *ImplicitRegDetails
	isSet bool
}

func (v NullableImplicitRegDetails) Get() *ImplicitRegDetails {
	return v.value
}

func (v *NullableImplicitRegDetails) Set(val *ImplicitRegDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableImplicitRegDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableImplicitRegDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImplicitRegDetails(val *ImplicitRegDetails) *NullableImplicitRegDetails {
	return &NullableImplicitRegDetails{value: val, isSet: true}
}

func (v NullableImplicitRegDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImplicitRegDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


