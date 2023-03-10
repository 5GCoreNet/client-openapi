/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"time"
)

// ImsSdmSubscription A subscription to notifications of data change
type ImsSdmSubscription struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
	MonitoredResourceUris []string `json:"monitoredResourceUris"`
	// string with format 'date-time' as defined in OpenAPI.
	Expires *time.Time `json:"expires,omitempty"`
}

// NewImsSdmSubscription instantiates a new ImsSdmSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImsSdmSubscription(nfInstanceId string, callbackReference string, monitoredResourceUris []string) *ImsSdmSubscription {
	this := ImsSdmSubscription{}
	this.NfInstanceId = nfInstanceId
	this.CallbackReference = callbackReference
	this.MonitoredResourceUris = monitoredResourceUris
	return &this
}

// NewImsSdmSubscriptionWithDefaults instantiates a new ImsSdmSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImsSdmSubscriptionWithDefaults() *ImsSdmSubscription {
	this := ImsSdmSubscription{}
	return &this
}

// GetNfInstanceId returns the NfInstanceId field value
func (o *ImsSdmSubscription) GetNfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value
// and a boolean to check if the value has been set.
func (o *ImsSdmSubscription) GetNfInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfInstanceId, true
}

// SetNfInstanceId sets field value
func (o *ImsSdmSubscription) SetNfInstanceId(v string) {
	o.NfInstanceId = v
}

// GetCallbackReference returns the CallbackReference field value
func (o *ImsSdmSubscription) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *ImsSdmSubscription) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *ImsSdmSubscription) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetMonitoredResourceUris returns the MonitoredResourceUris field value
func (o *ImsSdmSubscription) GetMonitoredResourceUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MonitoredResourceUris
}

// GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field value
// and a boolean to check if the value has been set.
func (o *ImsSdmSubscription) GetMonitoredResourceUrisOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MonitoredResourceUris, true
}

// SetMonitoredResourceUris sets field value
func (o *ImsSdmSubscription) SetMonitoredResourceUris(v []string) {
	o.MonitoredResourceUris = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *ImsSdmSubscription) GetExpires() time.Time {
	if o == nil || isNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsSdmSubscription) GetExpiresOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expires) {
    return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *ImsSdmSubscription) HasExpires() bool {
	if o != nil && !isNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *ImsSdmSubscription) SetExpires(v time.Time) {
	o.Expires = &v
}

func (o ImsSdmSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfInstanceId"] = o.NfInstanceId
	}
	if true {
		toSerialize["callbackReference"] = o.CallbackReference
	}
	if true {
		toSerialize["monitoredResourceUris"] = o.MonitoredResourceUris
	}
	if !isNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	return json.Marshal(toSerialize)
}

type NullableImsSdmSubscription struct {
	value *ImsSdmSubscription
	isSet bool
}

func (v NullableImsSdmSubscription) Get() *ImsSdmSubscription {
	return v.value
}

func (v *NullableImsSdmSubscription) Set(val *ImsSdmSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableImsSdmSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableImsSdmSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImsSdmSubscription(val *ImsSdmSubscription) *NullableImsSdmSubscription {
	return &NullableImsSdmSubscription{value: val, isSet: true}
}

func (v NullableImsSdmSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImsSdmSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


