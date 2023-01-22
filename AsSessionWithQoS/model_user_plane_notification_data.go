/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AsSessionWithQoS

import (
	"encoding/json"
)

// UserPlaneNotificationData Represents the parameters to be conveyed in a user plane event(s) notification.
type UserPlaneNotificationData struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Transaction string `json:"transaction"`
	// Contains the reported event and applicable information
	EventReports []UserPlaneEventReport `json:"eventReports"`
}

// NewUserPlaneNotificationData instantiates a new UserPlaneNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPlaneNotificationData(transaction string, eventReports []UserPlaneEventReport) *UserPlaneNotificationData {
	this := UserPlaneNotificationData{}
	this.Transaction = transaction
	this.EventReports = eventReports
	return &this
}

// NewUserPlaneNotificationDataWithDefaults instantiates a new UserPlaneNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPlaneNotificationDataWithDefaults() *UserPlaneNotificationData {
	this := UserPlaneNotificationData{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *UserPlaneNotificationData) GetTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *UserPlaneNotificationData) GetTransactionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *UserPlaneNotificationData) SetTransaction(v string) {
	o.Transaction = v
}

// GetEventReports returns the EventReports field value
func (o *UserPlaneNotificationData) GetEventReports() []UserPlaneEventReport {
	if o == nil {
		var ret []UserPlaneEventReport
		return ret
	}

	return o.EventReports
}

// GetEventReportsOk returns a tuple with the EventReports field value
// and a boolean to check if the value has been set.
func (o *UserPlaneNotificationData) GetEventReportsOk() ([]UserPlaneEventReport, bool) {
	if o == nil {
    return nil, false
	}
	return o.EventReports, true
}

// SetEventReports sets field value
func (o *UserPlaneNotificationData) SetEventReports(v []UserPlaneEventReport) {
	o.EventReports = v
}

func (o UserPlaneNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transaction"] = o.Transaction
	}
	if true {
		toSerialize["eventReports"] = o.EventReports
	}
	return json.Marshal(toSerialize)
}

type NullableUserPlaneNotificationData struct {
	value *UserPlaneNotificationData
	isSet bool
}

func (v NullableUserPlaneNotificationData) Get() *UserPlaneNotificationData {
	return v.value
}

func (v *NullableUserPlaneNotificationData) Set(val *UserPlaneNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPlaneNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPlaneNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPlaneNotificationData(val *UserPlaneNotificationData) *NullableUserPlaneNotificationData {
	return &NullableUserPlaneNotificationData{value: val, isSet: true}
}

func (v NullableUserPlaneNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPlaneNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


