/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// BroadcastDeliveryStatus Broadcast MBS Session's Delivery Status
type BroadcastDeliveryStatus struct {
	BroadcastDeliveryStatusAnyOf *BroadcastDeliveryStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BroadcastDeliveryStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BroadcastDeliveryStatusAnyOf
	err = json.Unmarshal(data, &dst.BroadcastDeliveryStatusAnyOf);
	if err == nil {
		jsonBroadcastDeliveryStatusAnyOf, _ := json.Marshal(dst.BroadcastDeliveryStatusAnyOf)
		if string(jsonBroadcastDeliveryStatusAnyOf) == "{}" { // empty struct
			dst.BroadcastDeliveryStatusAnyOf = nil
		} else {
			return nil // data stored in dst.BroadcastDeliveryStatusAnyOf, return on the first match
		}
	} else {
		dst.BroadcastDeliveryStatusAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(BroadcastDeliveryStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BroadcastDeliveryStatus) MarshalJSON() ([]byte, error) {
	if src.BroadcastDeliveryStatusAnyOf != nil {
		return json.Marshal(&src.BroadcastDeliveryStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBroadcastDeliveryStatus struct {
	value *BroadcastDeliveryStatus
	isSet bool
}

func (v NullableBroadcastDeliveryStatus) Get() *BroadcastDeliveryStatus {
	return v.value
}

func (v *NullableBroadcastDeliveryStatus) Set(val *BroadcastDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastDeliveryStatus(val *BroadcastDeliveryStatus) *NullableBroadcastDeliveryStatus {
	return &NullableBroadcastDeliveryStatus{value: val, isSet: true}
}

func (v NullableBroadcastDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


