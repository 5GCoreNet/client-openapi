/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
)

// OperationEvent Operation Event for a Broadcast MBS Session.
type OperationEvent struct {
	OpEventType OpEventType `json:"opEventType"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfId *string `json:"amfId,omitempty"`
	NgranFailureEventList []NgranFailureEvent `json:"ngranFailureEventList,omitempty"`
}

// NewOperationEvent instantiates a new OperationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationEvent(opEventType OpEventType) *OperationEvent {
	this := OperationEvent{}
	this.OpEventType = opEventType
	return &this
}

// NewOperationEventWithDefaults instantiates a new OperationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationEventWithDefaults() *OperationEvent {
	this := OperationEvent{}
	return &this
}

// GetOpEventType returns the OpEventType field value
func (o *OperationEvent) GetOpEventType() OpEventType {
	if o == nil {
		var ret OpEventType
		return ret
	}

	return o.OpEventType
}

// GetOpEventTypeOk returns a tuple with the OpEventType field value
// and a boolean to check if the value has been set.
func (o *OperationEvent) GetOpEventTypeOk() (*OpEventType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OpEventType, true
}

// SetOpEventType sets field value
func (o *OperationEvent) SetOpEventType(v OpEventType) {
	o.OpEventType = v
}

// GetAmfId returns the AmfId field value if set, zero value otherwise.
func (o *OperationEvent) GetAmfId() string {
	if o == nil || isNil(o.AmfId) {
		var ret string
		return ret
	}
	return *o.AmfId
}

// GetAmfIdOk returns a tuple with the AmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationEvent) GetAmfIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfId) {
    return nil, false
	}
	return o.AmfId, true
}

// HasAmfId returns a boolean if a field has been set.
func (o *OperationEvent) HasAmfId() bool {
	if o != nil && !isNil(o.AmfId) {
		return true
	}

	return false
}

// SetAmfId gets a reference to the given string and assigns it to the AmfId field.
func (o *OperationEvent) SetAmfId(v string) {
	o.AmfId = &v
}

// GetNgranFailureEventList returns the NgranFailureEventList field value if set, zero value otherwise.
func (o *OperationEvent) GetNgranFailureEventList() []NgranFailureEvent {
	if o == nil || isNil(o.NgranFailureEventList) {
		var ret []NgranFailureEvent
		return ret
	}
	return o.NgranFailureEventList
}

// GetNgranFailureEventListOk returns a tuple with the NgranFailureEventList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationEvent) GetNgranFailureEventListOk() ([]NgranFailureEvent, bool) {
	if o == nil || isNil(o.NgranFailureEventList) {
    return nil, false
	}
	return o.NgranFailureEventList, true
}

// HasNgranFailureEventList returns a boolean if a field has been set.
func (o *OperationEvent) HasNgranFailureEventList() bool {
	if o != nil && !isNil(o.NgranFailureEventList) {
		return true
	}

	return false
}

// SetNgranFailureEventList gets a reference to the given []NgranFailureEvent and assigns it to the NgranFailureEventList field.
func (o *OperationEvent) SetNgranFailureEventList(v []NgranFailureEvent) {
	o.NgranFailureEventList = v
}

func (o OperationEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["opEventType"] = o.OpEventType
	}
	if !isNil(o.AmfId) {
		toSerialize["amfId"] = o.AmfId
	}
	if !isNil(o.NgranFailureEventList) {
		toSerialize["ngranFailureEventList"] = o.NgranFailureEventList
	}
	return json.Marshal(toSerialize)
}

type NullableOperationEvent struct {
	value *OperationEvent
	isSet bool
}

func (v NullableOperationEvent) Get() *OperationEvent {
	return v.value
}

func (v *NullableOperationEvent) Set(val *OperationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationEvent(val *OperationEvent) *NullableOperationEvent {
	return &NullableOperationEvent{value: val, isSet: true}
}

func (v NullableOperationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


