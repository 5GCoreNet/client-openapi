/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_DNSContext

import (
	"encoding/json"
)

// DnsContextNotification Data within DNS Context Notify
type DnsContextNotification struct {
	EventreportList []DnsContextEventReport `json:"eventreportList,omitempty"`
}

// NewDnsContextNotification instantiates a new DnsContextNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsContextNotification() *DnsContextNotification {
	this := DnsContextNotification{}
	return &this
}

// NewDnsContextNotificationWithDefaults instantiates a new DnsContextNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsContextNotificationWithDefaults() *DnsContextNotification {
	this := DnsContextNotification{}
	return &this
}

// GetEventreportList returns the EventreportList field value if set, zero value otherwise.
func (o *DnsContextNotification) GetEventreportList() []DnsContextEventReport {
	if o == nil || isNil(o.EventreportList) {
		var ret []DnsContextEventReport
		return ret
	}
	return o.EventreportList
}

// GetEventreportListOk returns a tuple with the EventreportList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsContextNotification) GetEventreportListOk() ([]DnsContextEventReport, bool) {
	if o == nil || isNil(o.EventreportList) {
    return nil, false
	}
	return o.EventreportList, true
}

// HasEventreportList returns a boolean if a field has been set.
func (o *DnsContextNotification) HasEventreportList() bool {
	if o != nil && !isNil(o.EventreportList) {
		return true
	}

	return false
}

// SetEventreportList gets a reference to the given []DnsContextEventReport and assigns it to the EventreportList field.
func (o *DnsContextNotification) SetEventreportList(v []DnsContextEventReport) {
	o.EventreportList = v
}

func (o DnsContextNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventreportList) {
		toSerialize["eventreportList"] = o.EventreportList
	}
	return json.Marshal(toSerialize)
}

type NullableDnsContextNotification struct {
	value *DnsContextNotification
	isSet bool
}

func (v NullableDnsContextNotification) Get() *DnsContextNotification {
	return v.value
}

func (v *NullableDnsContextNotification) Set(val *DnsContextNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsContextNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsContextNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsContextNotification(val *DnsContextNotification) *NullableDnsContextNotification {
	return &NullableDnsContextNotification{value: val, isSet: true}
}

func (v NullableDnsContextNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsContextNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


