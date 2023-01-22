/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_PP

import (
	"encoding/json"
	"time"
)

// PpDlPacketCountExt struct for PpDlPacketCountExt
type PpDlPacketCountExt struct {
	AfInstanceId string `json:"afInstanceId"`
	ReferenceId int32 `json:"referenceId"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	SingleNssai *Snssai `json:"singleNssai,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
}

// NewPpDlPacketCountExt instantiates a new PpDlPacketCountExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPpDlPacketCountExt(afInstanceId string, referenceId int32) *PpDlPacketCountExt {
	this := PpDlPacketCountExt{}
	this.AfInstanceId = afInstanceId
	this.ReferenceId = referenceId
	return &this
}

// NewPpDlPacketCountExtWithDefaults instantiates a new PpDlPacketCountExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPpDlPacketCountExtWithDefaults() *PpDlPacketCountExt {
	this := PpDlPacketCountExt{}
	return &this
}

// GetAfInstanceId returns the AfInstanceId field value
func (o *PpDlPacketCountExt) GetAfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfInstanceId
}

// GetAfInstanceIdOk returns a tuple with the AfInstanceId field value
// and a boolean to check if the value has been set.
func (o *PpDlPacketCountExt) GetAfInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfInstanceId, true
}

// SetAfInstanceId sets field value
func (o *PpDlPacketCountExt) SetAfInstanceId(v string) {
	o.AfInstanceId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *PpDlPacketCountExt) GetReferenceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *PpDlPacketCountExt) GetReferenceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *PpDlPacketCountExt) SetReferenceId(v int32) {
	o.ReferenceId = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PpDlPacketCountExt) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDlPacketCountExt) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PpDlPacketCountExt) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PpDlPacketCountExt) SetDnn(v string) {
	o.Dnn = &v
}

// GetSingleNssai returns the SingleNssai field value if set, zero value otherwise.
func (o *PpDlPacketCountExt) GetSingleNssai() Snssai {
	if o == nil || isNil(o.SingleNssai) {
		var ret Snssai
		return ret
	}
	return *o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDlPacketCountExt) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.SingleNssai) {
    return nil, false
	}
	return o.SingleNssai, true
}

// HasSingleNssai returns a boolean if a field has been set.
func (o *PpDlPacketCountExt) HasSingleNssai() bool {
	if o != nil && !isNil(o.SingleNssai) {
		return true
	}

	return false
}

// SetSingleNssai gets a reference to the given Snssai and assigns it to the SingleNssai field.
func (o *PpDlPacketCountExt) SetSingleNssai(v Snssai) {
	o.SingleNssai = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *PpDlPacketCountExt) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDlPacketCountExt) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
    return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *PpDlPacketCountExt) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *PpDlPacketCountExt) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *PpDlPacketCountExt) GetMtcProviderInformation() string {
	if o == nil || isNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDlPacketCountExt) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || isNil(o.MtcProviderInformation) {
    return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *PpDlPacketCountExt) HasMtcProviderInformation() bool {
	if o != nil && !isNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *PpDlPacketCountExt) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

func (o PpDlPacketCountExt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["afInstanceId"] = o.AfInstanceId
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.SingleNssai) {
		toSerialize["singleNssai"] = o.SingleNssai
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !isNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	return json.Marshal(toSerialize)
}

type NullablePpDlPacketCountExt struct {
	value *PpDlPacketCountExt
	isSet bool
}

func (v NullablePpDlPacketCountExt) Get() *PpDlPacketCountExt {
	return v.value
}

func (v *NullablePpDlPacketCountExt) Set(val *PpDlPacketCountExt) {
	v.value = val
	v.isSet = true
}

func (v NullablePpDlPacketCountExt) IsSet() bool {
	return v.isSet
}

func (v *NullablePpDlPacketCountExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpDlPacketCountExt(val *PpDlPacketCountExt) *NullablePpDlPacketCountExt {
	return &NullablePpDlPacketCountExt{value: val, isSet: true}
}

func (v NullablePpDlPacketCountExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpDlPacketCountExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


