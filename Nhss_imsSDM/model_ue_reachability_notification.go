/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsSDM

import (
	"encoding/json"
)

// UeReachabilityNotification Represents the contents of a notification of UE reachability for IP sent by the HSS 
type UeReachabilityNotification struct {
	ReachabilityIndicator bool `json:"reachabilityIndicator"`
	DetectingNode DetectingNode `json:"detectingNode"`
	AccessType AccessType `json:"accessType"`
}

// NewUeReachabilityNotification instantiates a new UeReachabilityNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeReachabilityNotification(reachabilityIndicator bool, detectingNode DetectingNode, accessType AccessType) *UeReachabilityNotification {
	this := UeReachabilityNotification{}
	this.ReachabilityIndicator = reachabilityIndicator
	this.DetectingNode = detectingNode
	this.AccessType = accessType
	return &this
}

// NewUeReachabilityNotificationWithDefaults instantiates a new UeReachabilityNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeReachabilityNotificationWithDefaults() *UeReachabilityNotification {
	this := UeReachabilityNotification{}
	return &this
}

// GetReachabilityIndicator returns the ReachabilityIndicator field value
func (o *UeReachabilityNotification) GetReachabilityIndicator() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReachabilityIndicator
}

// GetReachabilityIndicatorOk returns a tuple with the ReachabilityIndicator field value
// and a boolean to check if the value has been set.
func (o *UeReachabilityNotification) GetReachabilityIndicatorOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReachabilityIndicator, true
}

// SetReachabilityIndicator sets field value
func (o *UeReachabilityNotification) SetReachabilityIndicator(v bool) {
	o.ReachabilityIndicator = v
}

// GetDetectingNode returns the DetectingNode field value
func (o *UeReachabilityNotification) GetDetectingNode() DetectingNode {
	if o == nil {
		var ret DetectingNode
		return ret
	}

	return o.DetectingNode
}

// GetDetectingNodeOk returns a tuple with the DetectingNode field value
// and a boolean to check if the value has been set.
func (o *UeReachabilityNotification) GetDetectingNodeOk() (*DetectingNode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DetectingNode, true
}

// SetDetectingNode sets field value
func (o *UeReachabilityNotification) SetDetectingNode(v DetectingNode) {
	o.DetectingNode = v
}

// GetAccessType returns the AccessType field value
func (o *UeReachabilityNotification) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *UeReachabilityNotification) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *UeReachabilityNotification) SetAccessType(v AccessType) {
	o.AccessType = v
}

func (o UeReachabilityNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reachabilityIndicator"] = o.ReachabilityIndicator
	}
	if true {
		toSerialize["detectingNode"] = o.DetectingNode
	}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableUeReachabilityNotification struct {
	value *UeReachabilityNotification
	isSet bool
}

func (v NullableUeReachabilityNotification) Get() *UeReachabilityNotification {
	return v.value
}

func (v *NullableUeReachabilityNotification) Set(val *UeReachabilityNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableUeReachabilityNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableUeReachabilityNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeReachabilityNotification(val *UeReachabilityNotification) *NullableUeReachabilityNotification {
	return &NullableUeReachabilityNotification{value: val, isSet: true}
}

func (v NullableUeReachabilityNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeReachabilityNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


