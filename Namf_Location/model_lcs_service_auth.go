/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Location

import (
	"encoding/json"
	"fmt"
)

// LcsServiceAuth Possible values are: - \"LOCATION_ALLOWED_WITH_NOTIFICATION\": Location allowed with notification - \"LOCATION_ALLOWED_WITHOUT_NOTIFICATION\": Location allowed without notification - \"LOCATION_ALLOWED_WITHOUT_RESPONSE\": Location with notification and privacy    verification; location allowed if no response - \"LOCATION_RESTRICTED_WITHOUT_RESPONSE\": Location with notification and privacy   verification; location restricted if no response - \"NOTIFICATION_ONLY\": Notification only - \"NOTIFICATION_AND_VERIFICATION_ONLY\": Notification and privacy verification only 
type LcsServiceAuth struct {
	LcsServiceAuthAnyOf *LcsServiceAuthAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LcsServiceAuth) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LcsServiceAuthAnyOf
	err = json.Unmarshal(data, &dst.LcsServiceAuthAnyOf);
	if err == nil {
		jsonLcsServiceAuthAnyOf, _ := json.Marshal(dst.LcsServiceAuthAnyOf)
		if string(jsonLcsServiceAuthAnyOf) == "{}" { // empty struct
			dst.LcsServiceAuthAnyOf = nil
		} else {
			return nil // data stored in dst.LcsServiceAuthAnyOf, return on the first match
		}
	} else {
		dst.LcsServiceAuthAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LcsServiceAuth)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LcsServiceAuth) MarshalJSON() ([]byte, error) {
	if src.LcsServiceAuthAnyOf != nil {
		return json.Marshal(&src.LcsServiceAuthAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLcsServiceAuth struct {
	value *LcsServiceAuth
	isSet bool
}

func (v NullableLcsServiceAuth) Get() *LcsServiceAuth {
	return v.value
}

func (v *NullableLcsServiceAuth) Set(val *LcsServiceAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsServiceAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsServiceAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsServiceAuth(val *LcsServiceAuth) *NullableLcsServiceAuth {
	return &NullableLcsServiceAuth{value: val, isSet: true}
}

func (v NullableLcsServiceAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsServiceAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


