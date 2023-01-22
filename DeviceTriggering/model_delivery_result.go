/*
3gpp-device-triggering

API for device trigger.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DeviceTriggering

import (
	"encoding/json"
	"fmt"
)

// DeliveryResult Possible values are - SUCCESS: This value indicates that the device action request was successfully completed. - UNKNOWN: This value indicates any unspecified errors. - FAILURE: This value indicates that this trigger encountered a delivery error and is deemed permanently undeliverable. - TRIGGERED: This value indicates that device triggering request is accepted by the SCEF. - EXPIRED: This value indicates that the validity period expired before the trigger could be delivered. - UNCONFIRMED: This value indicates that the delivery of the device action request is not confirmed. - REPLACED: This value indicates that the device triggering replacement request is accepted by the SCEF. - TERMINATE: This value indicates that the delivery of the device action request is terminated by the SCS/AS. 
type DeliveryResult struct {
	DeliveryResultAnyOf *DeliveryResultAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DeliveryResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DeliveryResultAnyOf
	err = json.Unmarshal(data, &dst.DeliveryResultAnyOf);
	if err == nil {
		jsonDeliveryResultAnyOf, _ := json.Marshal(dst.DeliveryResultAnyOf)
		if string(jsonDeliveryResultAnyOf) == "{}" { // empty struct
			dst.DeliveryResultAnyOf = nil
		} else {
			return nil // data stored in dst.DeliveryResultAnyOf, return on the first match
		}
	} else {
		dst.DeliveryResultAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DeliveryResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DeliveryResult) MarshalJSON() ([]byte, error) {
	if src.DeliveryResultAnyOf != nil {
		return json.Marshal(&src.DeliveryResultAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDeliveryResult struct {
	value *DeliveryResult
	isSet bool
}

func (v NullableDeliveryResult) Get() *DeliveryResult {
	return v.value
}

func (v *NullableDeliveryResult) Set(val *DeliveryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryResult(val *DeliveryResult) *NullableDeliveryResult {
	return &NullableDeliveryResult{value: val, isSet: true}
}

func (v NullableDeliveryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


