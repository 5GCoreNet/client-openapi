/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FaultMnS

import (
	"encoding/json"
)

// MergePatchAcknowledgeAlarm Patch document acknowledging or unacknowledging a single alarm. For acknowleding an alarm the value of ackState is ACKNOWLEDGED, for unacknowleding an alarm the value of ackState is UNACKNOWLEDGED.
type MergePatchAcknowledgeAlarm struct {
	AckUserId string `json:"ackUserId"`
	AckSystemId *string `json:"ackSystemId,omitempty"`
	AckState AckState `json:"ackState"`
}

// NewMergePatchAcknowledgeAlarm instantiates a new MergePatchAcknowledgeAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergePatchAcknowledgeAlarm(ackUserId string, ackState AckState) *MergePatchAcknowledgeAlarm {
	this := MergePatchAcknowledgeAlarm{}
	this.AckUserId = ackUserId
	this.AckState = ackState
	return &this
}

// NewMergePatchAcknowledgeAlarmWithDefaults instantiates a new MergePatchAcknowledgeAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergePatchAcknowledgeAlarmWithDefaults() *MergePatchAcknowledgeAlarm {
	this := MergePatchAcknowledgeAlarm{}
	return &this
}

// GetAckUserId returns the AckUserId field value
func (o *MergePatchAcknowledgeAlarm) GetAckUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AckUserId
}

// GetAckUserIdOk returns a tuple with the AckUserId field value
// and a boolean to check if the value has been set.
func (o *MergePatchAcknowledgeAlarm) GetAckUserIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AckUserId, true
}

// SetAckUserId sets field value
func (o *MergePatchAcknowledgeAlarm) SetAckUserId(v string) {
	o.AckUserId = v
}

// GetAckSystemId returns the AckSystemId field value if set, zero value otherwise.
func (o *MergePatchAcknowledgeAlarm) GetAckSystemId() string {
	if o == nil || isNil(o.AckSystemId) {
		var ret string
		return ret
	}
	return *o.AckSystemId
}

// GetAckSystemIdOk returns a tuple with the AckSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergePatchAcknowledgeAlarm) GetAckSystemIdOk() (*string, bool) {
	if o == nil || isNil(o.AckSystemId) {
    return nil, false
	}
	return o.AckSystemId, true
}

// HasAckSystemId returns a boolean if a field has been set.
func (o *MergePatchAcknowledgeAlarm) HasAckSystemId() bool {
	if o != nil && !isNil(o.AckSystemId) {
		return true
	}

	return false
}

// SetAckSystemId gets a reference to the given string and assigns it to the AckSystemId field.
func (o *MergePatchAcknowledgeAlarm) SetAckSystemId(v string) {
	o.AckSystemId = &v
}

// GetAckState returns the AckState field value
func (o *MergePatchAcknowledgeAlarm) GetAckState() AckState {
	if o == nil {
		var ret AckState
		return ret
	}

	return o.AckState
}

// GetAckStateOk returns a tuple with the AckState field value
// and a boolean to check if the value has been set.
func (o *MergePatchAcknowledgeAlarm) GetAckStateOk() (*AckState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AckState, true
}

// SetAckState sets field value
func (o *MergePatchAcknowledgeAlarm) SetAckState(v AckState) {
	o.AckState = v
}

func (o MergePatchAcknowledgeAlarm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ackUserId"] = o.AckUserId
	}
	if !isNil(o.AckSystemId) {
		toSerialize["ackSystemId"] = o.AckSystemId
	}
	if true {
		toSerialize["ackState"] = o.AckState
	}
	return json.Marshal(toSerialize)
}

type NullableMergePatchAcknowledgeAlarm struct {
	value *MergePatchAcknowledgeAlarm
	isSet bool
}

func (v NullableMergePatchAcknowledgeAlarm) Get() *MergePatchAcknowledgeAlarm {
	return v.value
}

func (v *NullableMergePatchAcknowledgeAlarm) Set(val *MergePatchAcknowledgeAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableMergePatchAcknowledgeAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableMergePatchAcknowledgeAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergePatchAcknowledgeAlarm(val *MergePatchAcknowledgeAlarm) *NullableMergePatchAcknowledgeAlarm {
	return &NullableMergePatchAcknowledgeAlarm{value: val, isSet: true}
}

func (v NullableMergePatchAcknowledgeAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergePatchAcknowledgeAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


