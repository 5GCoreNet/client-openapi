/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// PLMNTargetType The PLMN for which sessions shall be selected in the Trace Session in case of management based activation when several PLMNs are supported in the RAN (this means that shared cells and not shared cells are allowed for the specified PLMN. Note that the PLMN Target might differ from the PLMN specified in the Trace Reference, as that specifies the PLMN that is containing the management system requesting the Trace Session from the NE. See 3GPP TS 32.422 clause 5.9b for additional details.
type PLMNTargetType struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}

// NewPLMNTargetType instantiates a new PLMNTargetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPLMNTargetType(mcc string, mnc string) *PLMNTargetType {
	this := PLMNTargetType{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPLMNTargetTypeWithDefaults instantiates a new PLMNTargetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPLMNTargetTypeWithDefaults() *PLMNTargetType {
	this := PLMNTargetType{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PLMNTargetType) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PLMNTargetType) GetMccOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PLMNTargetType) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PLMNTargetType) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PLMNTargetType) GetMncOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PLMNTargetType) SetMnc(v string) {
	o.Mnc = v
}

func (o PLMNTargetType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mcc"] = o.Mcc
	}
	if true {
		toSerialize["mnc"] = o.Mnc
	}
	return json.Marshal(toSerialize)
}

type NullablePLMNTargetType struct {
	value *PLMNTargetType
	isSet bool
}

func (v NullablePLMNTargetType) Get() *PLMNTargetType {
	return v.value
}

func (v *NullablePLMNTargetType) Set(val *PLMNTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullablePLMNTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullablePLMNTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePLMNTargetType(val *PLMNTargetType) *NullablePLMNTargetType {
	return &NullablePLMNTargetType{value: val, isSet: true}
}

func (v NullablePLMNTargetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePLMNTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


