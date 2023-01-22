/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// SmfEvent Possible values are: - AC_TY_CH: Access Type Change - UP_PATH_CH: UP Path Change - PDU_SES_REL: PDU Session Release - PLMN_CH: PLMN Change - UE_IP_CH: UE IP address change - RAT_TY_CH: RAT Type Change - DDDS: Downlink data delivery status - COMM_FAIL: Communication Failure - PDU_SES_EST: PDU Session Establishment - QFI_ALLOC: QFI allocation - QOS_MON: QoS Monitoring - SMCC_EXP: SM congestion control experience for PDU Session - DISPERSION: Session Management transaction dispersion - RED_TRANS_EXP: Redundant transmission experience for PDU Session - WLAN_INFO: WLAN information on PDU session for which Access Type is NON_3GPP_ACCESS and   RAT Type is TRUSTED_WLAN - UPF_INFO: The UPF information, including the UPF ID/address/FQDN information. 
type SmfEvent struct {
	SmfEventAnyOf *SmfEventAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SmfEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SmfEventAnyOf
	err = json.Unmarshal(data, &dst.SmfEventAnyOf);
	if err == nil {
		jsonSmfEventAnyOf, _ := json.Marshal(dst.SmfEventAnyOf)
		if string(jsonSmfEventAnyOf) == "{}" { // empty struct
			dst.SmfEventAnyOf = nil
		} else {
			return nil // data stored in dst.SmfEventAnyOf, return on the first match
		}
	} else {
		dst.SmfEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SmfEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SmfEvent) MarshalJSON() ([]byte, error) {
	if src.SmfEventAnyOf != nil {
		return json.Marshal(&src.SmfEventAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSmfEvent struct {
	value *SmfEvent
	isSet bool
}

func (v NullableSmfEvent) Get() *SmfEvent {
	return v.value
}

func (v *NullableSmfEvent) Set(val *SmfEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfEvent(val *SmfEvent) *NullableSmfEvent {
	return &NullableSmfEvent{value: val, isSet: true}
}

func (v NullableSmfEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


