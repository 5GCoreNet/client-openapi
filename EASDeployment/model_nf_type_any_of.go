/*
3gpp-eas-deployment

API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EASDeployment

import (
	"encoding/json"
	"fmt"
)

// NFTypeAnyOf the model 'NFTypeAnyOf'
type NFTypeAnyOf string

// List of NFType_anyOf
const (
	NRF NFTypeAnyOf = "NRF"
	UDM NFTypeAnyOf = "UDM"
	AMF NFTypeAnyOf = "AMF"
	SMF NFTypeAnyOf = "SMF"
	AUSF NFTypeAnyOf = "AUSF"
	NEF NFTypeAnyOf = "NEF"
	PCF NFTypeAnyOf = "PCF"
	SMSF NFTypeAnyOf = "SMSF"
	NSSF NFTypeAnyOf = "NSSF"
	UDR NFTypeAnyOf = "UDR"
	LMF NFTypeAnyOf = "LMF"
	GMLC NFTypeAnyOf = "GMLC"
	_5_G_EIR NFTypeAnyOf = "5G_EIR"
	SEPP NFTypeAnyOf = "SEPP"
	UPF NFTypeAnyOf = "UPF"
	N3_IWF NFTypeAnyOf = "N3IWF"
	AF NFTypeAnyOf = "AF"
	UDSF NFTypeAnyOf = "UDSF"
	BSF NFTypeAnyOf = "BSF"
	CHF NFTypeAnyOf = "CHF"
	NWDAF NFTypeAnyOf = "NWDAF"
	PCSCF NFTypeAnyOf = "PCSCF"
	CBCF NFTypeAnyOf = "CBCF"
	HSS NFTypeAnyOf = "HSS"
	UCMF NFTypeAnyOf = "UCMF"
	SOR_AF NFTypeAnyOf = "SOR_AF"
	SPAF NFTypeAnyOf = "SPAF"
	MME NFTypeAnyOf = "MME"
	SCSAS NFTypeAnyOf = "SCSAS"
	SCEF NFTypeAnyOf = "SCEF"
	SCP NFTypeAnyOf = "SCP"
	NSSAAF NFTypeAnyOf = "NSSAAF"
	ICSCF NFTypeAnyOf = "ICSCF"
	SCSCF NFTypeAnyOf = "SCSCF"
	DRA NFTypeAnyOf = "DRA"
	IMS_AS NFTypeAnyOf = "IMS_AS"
	AANF NFTypeAnyOf = "AANF"
	_5_G_DDNMF NFTypeAnyOf = "5G_DDNMF"
	NSACF NFTypeAnyOf = "NSACF"
	MFAF NFTypeAnyOf = "MFAF"
	EASDF NFTypeAnyOf = "EASDF"
	DCCF NFTypeAnyOf = "DCCF"
	MB_SMF NFTypeAnyOf = "MB_SMF"
	TSCTSF NFTypeAnyOf = "TSCTSF"
	ADRF NFTypeAnyOf = "ADRF"
	GBA_BSF NFTypeAnyOf = "GBA_BSF"
	CEF NFTypeAnyOf = "CEF"
	MB_UPF NFTypeAnyOf = "MB_UPF"
	NSWOF NFTypeAnyOf = "NSWOF"
	PKMF NFTypeAnyOf = "PKMF"
	MNPF NFTypeAnyOf = "MNPF"
	SMS_GMSC NFTypeAnyOf = "SMS_GMSC"
	SMS_IWMSC NFTypeAnyOf = "SMS_IWMSC"
	MBSF NFTypeAnyOf = "MBSF"
	MBSTF NFTypeAnyOf = "MBSTF"
	PANF NFTypeAnyOf = "PANF"
)

// All allowed values of NFTypeAnyOf enum
var AllowedNFTypeAnyOfEnumValues = []NFTypeAnyOf{
	"NRF",
	"UDM",
	"AMF",
	"SMF",
	"AUSF",
	"NEF",
	"PCF",
	"SMSF",
	"NSSF",
	"UDR",
	"LMF",
	"GMLC",
	"5G_EIR",
	"SEPP",
	"UPF",
	"N3IWF",
	"AF",
	"UDSF",
	"BSF",
	"CHF",
	"NWDAF",
	"PCSCF",
	"CBCF",
	"HSS",
	"UCMF",
	"SOR_AF",
	"SPAF",
	"MME",
	"SCSAS",
	"SCEF",
	"SCP",
	"NSSAAF",
	"ICSCF",
	"SCSCF",
	"DRA",
	"IMS_AS",
	"AANF",
	"5G_DDNMF",
	"NSACF",
	"MFAF",
	"EASDF",
	"DCCF",
	"MB_SMF",
	"TSCTSF",
	"ADRF",
	"GBA_BSF",
	"CEF",
	"MB_UPF",
	"NSWOF",
	"PKMF",
	"MNPF",
	"SMS_GMSC",
	"SMS_IWMSC",
	"MBSF",
	"MBSTF",
	"PANF",
}

func (v *NFTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NFTypeAnyOf(value)
	for _, existing := range AllowedNFTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NFTypeAnyOf", value)
}

// NewNFTypeAnyOfFromValue returns a pointer to a valid NFTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNFTypeAnyOfFromValue(v string) (*NFTypeAnyOf, error) {
	ev := NFTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NFTypeAnyOf: valid values are %v", v, AllowedNFTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NFTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNFTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NFType_anyOf value
func (v NFTypeAnyOf) Ptr() *NFTypeAnyOf {
	return &v
}

type NullableNFTypeAnyOf struct {
	value *NFTypeAnyOf
	isSet bool
}

func (v NullableNFTypeAnyOf) Get() *NFTypeAnyOf {
	return v.value
}

func (v *NullableNFTypeAnyOf) Set(val *NFTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTypeAnyOf(val *NFTypeAnyOf) *NullableNFTypeAnyOf {
	return &NullableNFTypeAnyOf{value: val, isSet: true}
}

func (v NullableNFTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

