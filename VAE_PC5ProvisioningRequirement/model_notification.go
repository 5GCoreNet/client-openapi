/*
VAE_PC5ProvisioningRequirement

API for VAE_PC5ProvisioningRequirement   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_PC5ProvisioningRequirement

import (
	"encoding/json"
)

// Notification Represents a notificaton of result of PC5 Provisioning Requirement.
type Notification struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri"`
	Result Result `json:"result"`
}

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification(resourceUri string, result Result) *Notification {
	this := Notification{}
	this.ResourceUri = resourceUri
	this.Result = result
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetResourceUri returns the ResourceUri field value
func (o *Notification) GetResourceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value
// and a boolean to check if the value has been set.
func (o *Notification) GetResourceUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ResourceUri, true
}

// SetResourceUri sets field value
func (o *Notification) SetResourceUri(v string) {
	o.ResourceUri = v
}

// GetResult returns the Result field value
func (o *Notification) GetResult() Result {
	if o == nil {
		var ret Result
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *Notification) GetResultOk() (*Result, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *Notification) SetResult(v Result) {
	o.Result = v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceUri"] = o.ResourceUri
	}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


