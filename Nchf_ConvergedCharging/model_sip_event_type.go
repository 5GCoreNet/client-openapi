/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
)

// SIPEventType struct for SIPEventType
type SIPEventType struct {
	SIPMethod *string `json:"sIPMethod,omitempty"`
	EventHeader *string `json:"eventHeader,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	ExpiresHeader *int32 `json:"expiresHeader,omitempty"`
}

// NewSIPEventType instantiates a new SIPEventType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSIPEventType() *SIPEventType {
	this := SIPEventType{}
	return &this
}

// NewSIPEventTypeWithDefaults instantiates a new SIPEventType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSIPEventTypeWithDefaults() *SIPEventType {
	this := SIPEventType{}
	return &this
}

// GetSIPMethod returns the SIPMethod field value if set, zero value otherwise.
func (o *SIPEventType) GetSIPMethod() string {
	if o == nil || isNil(o.SIPMethod) {
		var ret string
		return ret
	}
	return *o.SIPMethod
}

// GetSIPMethodOk returns a tuple with the SIPMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIPEventType) GetSIPMethodOk() (*string, bool) {
	if o == nil || isNil(o.SIPMethod) {
    return nil, false
	}
	return o.SIPMethod, true
}

// HasSIPMethod returns a boolean if a field has been set.
func (o *SIPEventType) HasSIPMethod() bool {
	if o != nil && !isNil(o.SIPMethod) {
		return true
	}

	return false
}

// SetSIPMethod gets a reference to the given string and assigns it to the SIPMethod field.
func (o *SIPEventType) SetSIPMethod(v string) {
	o.SIPMethod = &v
}

// GetEventHeader returns the EventHeader field value if set, zero value otherwise.
func (o *SIPEventType) GetEventHeader() string {
	if o == nil || isNil(o.EventHeader) {
		var ret string
		return ret
	}
	return *o.EventHeader
}

// GetEventHeaderOk returns a tuple with the EventHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIPEventType) GetEventHeaderOk() (*string, bool) {
	if o == nil || isNil(o.EventHeader) {
    return nil, false
	}
	return o.EventHeader, true
}

// HasEventHeader returns a boolean if a field has been set.
func (o *SIPEventType) HasEventHeader() bool {
	if o != nil && !isNil(o.EventHeader) {
		return true
	}

	return false
}

// SetEventHeader gets a reference to the given string and assigns it to the EventHeader field.
func (o *SIPEventType) SetEventHeader(v string) {
	o.EventHeader = &v
}

// GetExpiresHeader returns the ExpiresHeader field value if set, zero value otherwise.
func (o *SIPEventType) GetExpiresHeader() int32 {
	if o == nil || isNil(o.ExpiresHeader) {
		var ret int32
		return ret
	}
	return *o.ExpiresHeader
}

// GetExpiresHeaderOk returns a tuple with the ExpiresHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIPEventType) GetExpiresHeaderOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiresHeader) {
    return nil, false
	}
	return o.ExpiresHeader, true
}

// HasExpiresHeader returns a boolean if a field has been set.
func (o *SIPEventType) HasExpiresHeader() bool {
	if o != nil && !isNil(o.ExpiresHeader) {
		return true
	}

	return false
}

// SetExpiresHeader gets a reference to the given int32 and assigns it to the ExpiresHeader field.
func (o *SIPEventType) SetExpiresHeader(v int32) {
	o.ExpiresHeader = &v
}

func (o SIPEventType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SIPMethod) {
		toSerialize["sIPMethod"] = o.SIPMethod
	}
	if !isNil(o.EventHeader) {
		toSerialize["eventHeader"] = o.EventHeader
	}
	if !isNil(o.ExpiresHeader) {
		toSerialize["expiresHeader"] = o.ExpiresHeader
	}
	return json.Marshal(toSerialize)
}

type NullableSIPEventType struct {
	value *SIPEventType
	isSet bool
}

func (v NullableSIPEventType) Get() *SIPEventType {
	return v.value
}

func (v *NullableSIPEventType) Set(val *SIPEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableSIPEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableSIPEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSIPEventType(val *SIPEventType) *NullableSIPEventType {
	return &NullableSIPEventType{value: val, isSet: true}
}

func (v NullableSIPEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSIPEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


