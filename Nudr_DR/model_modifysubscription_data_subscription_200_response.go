/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// ModifysubscriptionDataSubscription200Response - struct for ModifysubscriptionDataSubscription200Response
type ModifysubscriptionDataSubscription200Response struct {
	PatchResult *PatchResult
	SubscriptionDataSubscriptions *SubscriptionDataSubscriptions
}

// PatchResultAsModifysubscriptionDataSubscription200Response is a convenience function that returns PatchResult wrapped in ModifysubscriptionDataSubscription200Response
func PatchResultAsModifysubscriptionDataSubscription200Response(v *PatchResult) ModifysubscriptionDataSubscription200Response {
	return ModifysubscriptionDataSubscription200Response{
		PatchResult: v,
	}
}

// SubscriptionDataSubscriptionsAsModifysubscriptionDataSubscription200Response is a convenience function that returns SubscriptionDataSubscriptions wrapped in ModifysubscriptionDataSubscription200Response
func SubscriptionDataSubscriptionsAsModifysubscriptionDataSubscription200Response(v *SubscriptionDataSubscriptions) ModifysubscriptionDataSubscription200Response {
	return ModifysubscriptionDataSubscription200Response{
		SubscriptionDataSubscriptions: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ModifysubscriptionDataSubscription200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PatchResult
	err = newStrictDecoder(data).Decode(&dst.PatchResult)
	if err == nil {
		jsonPatchResult, _ := json.Marshal(dst.PatchResult)
		if string(jsonPatchResult) == "{}" { // empty struct
			dst.PatchResult = nil
		} else {
			match++
		}
	} else {
		dst.PatchResult = nil
	}

	// try to unmarshal data into SubscriptionDataSubscriptions
	err = newStrictDecoder(data).Decode(&dst.SubscriptionDataSubscriptions)
	if err == nil {
		jsonSubscriptionDataSubscriptions, _ := json.Marshal(dst.SubscriptionDataSubscriptions)
		if string(jsonSubscriptionDataSubscriptions) == "{}" { // empty struct
			dst.SubscriptionDataSubscriptions = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionDataSubscriptions = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PatchResult = nil
		dst.SubscriptionDataSubscriptions = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ModifysubscriptionDataSubscription200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ModifysubscriptionDataSubscription200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ModifysubscriptionDataSubscription200Response) MarshalJSON() ([]byte, error) {
	if src.PatchResult != nil {
		return json.Marshal(&src.PatchResult)
	}

	if src.SubscriptionDataSubscriptions != nil {
		return json.Marshal(&src.SubscriptionDataSubscriptions)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ModifysubscriptionDataSubscription200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PatchResult != nil {
		return obj.PatchResult
	}

	if obj.SubscriptionDataSubscriptions != nil {
		return obj.SubscriptionDataSubscriptions
	}

	// all schemas are nil
	return nil
}

type NullableModifysubscriptionDataSubscription200Response struct {
	value *ModifysubscriptionDataSubscription200Response
	isSet bool
}

func (v NullableModifysubscriptionDataSubscription200Response) Get() *ModifysubscriptionDataSubscription200Response {
	return v.value
}

func (v *NullableModifysubscriptionDataSubscription200Response) Set(val *ModifysubscriptionDataSubscription200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifysubscriptionDataSubscription200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifysubscriptionDataSubscription200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifysubscriptionDataSubscription200Response(val *ModifysubscriptionDataSubscription200Response) *NullableModifysubscriptionDataSubscription200Response {
	return &NullableModifysubscriptionDataSubscription200Response{value: val, isSet: true}
}

func (v NullableModifysubscriptionDataSubscription200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifysubscriptionDataSubscription200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


