/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// PC5ContainerInformation struct for PC5ContainerInformation
type PC5ContainerInformation struct {
	CoverageInfoList []CoverageInfo `json:"coverageInfoList,omitempty"`
	RadioParameterSetInfoList []RadioParameterSetInfo `json:"radioParameterSetInfoList,omitempty"`
	TransmitterInfoList []TransmitterInfo `json:"transmitterInfoList,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeOfFirstTransmission *time.Time `json:"timeOfFirst Transmission,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeOfFirstReception *time.Time `json:"timeOfFirst Reception,omitempty"`
}

// NewPC5ContainerInformation instantiates a new PC5ContainerInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPC5ContainerInformation() *PC5ContainerInformation {
	this := PC5ContainerInformation{}
	return &this
}

// NewPC5ContainerInformationWithDefaults instantiates a new PC5ContainerInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPC5ContainerInformationWithDefaults() *PC5ContainerInformation {
	this := PC5ContainerInformation{}
	return &this
}

// GetCoverageInfoList returns the CoverageInfoList field value if set, zero value otherwise.
func (o *PC5ContainerInformation) GetCoverageInfoList() []CoverageInfo {
	if o == nil || isNil(o.CoverageInfoList) {
		var ret []CoverageInfo
		return ret
	}
	return o.CoverageInfoList
}

// GetCoverageInfoListOk returns a tuple with the CoverageInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5ContainerInformation) GetCoverageInfoListOk() ([]CoverageInfo, bool) {
	if o == nil || isNil(o.CoverageInfoList) {
    return nil, false
	}
	return o.CoverageInfoList, true
}

// HasCoverageInfoList returns a boolean if a field has been set.
func (o *PC5ContainerInformation) HasCoverageInfoList() bool {
	if o != nil && !isNil(o.CoverageInfoList) {
		return true
	}

	return false
}

// SetCoverageInfoList gets a reference to the given []CoverageInfo and assigns it to the CoverageInfoList field.
func (o *PC5ContainerInformation) SetCoverageInfoList(v []CoverageInfo) {
	o.CoverageInfoList = v
}

// GetRadioParameterSetInfoList returns the RadioParameterSetInfoList field value if set, zero value otherwise.
func (o *PC5ContainerInformation) GetRadioParameterSetInfoList() []RadioParameterSetInfo {
	if o == nil || isNil(o.RadioParameterSetInfoList) {
		var ret []RadioParameterSetInfo
		return ret
	}
	return o.RadioParameterSetInfoList
}

// GetRadioParameterSetInfoListOk returns a tuple with the RadioParameterSetInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5ContainerInformation) GetRadioParameterSetInfoListOk() ([]RadioParameterSetInfo, bool) {
	if o == nil || isNil(o.RadioParameterSetInfoList) {
    return nil, false
	}
	return o.RadioParameterSetInfoList, true
}

// HasRadioParameterSetInfoList returns a boolean if a field has been set.
func (o *PC5ContainerInformation) HasRadioParameterSetInfoList() bool {
	if o != nil && !isNil(o.RadioParameterSetInfoList) {
		return true
	}

	return false
}

// SetRadioParameterSetInfoList gets a reference to the given []RadioParameterSetInfo and assigns it to the RadioParameterSetInfoList field.
func (o *PC5ContainerInformation) SetRadioParameterSetInfoList(v []RadioParameterSetInfo) {
	o.RadioParameterSetInfoList = v
}

// GetTransmitterInfoList returns the TransmitterInfoList field value if set, zero value otherwise.
func (o *PC5ContainerInformation) GetTransmitterInfoList() []TransmitterInfo {
	if o == nil || isNil(o.TransmitterInfoList) {
		var ret []TransmitterInfo
		return ret
	}
	return o.TransmitterInfoList
}

// GetTransmitterInfoListOk returns a tuple with the TransmitterInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5ContainerInformation) GetTransmitterInfoListOk() ([]TransmitterInfo, bool) {
	if o == nil || isNil(o.TransmitterInfoList) {
    return nil, false
	}
	return o.TransmitterInfoList, true
}

// HasTransmitterInfoList returns a boolean if a field has been set.
func (o *PC5ContainerInformation) HasTransmitterInfoList() bool {
	if o != nil && !isNil(o.TransmitterInfoList) {
		return true
	}

	return false
}

// SetTransmitterInfoList gets a reference to the given []TransmitterInfo and assigns it to the TransmitterInfoList field.
func (o *PC5ContainerInformation) SetTransmitterInfoList(v []TransmitterInfo) {
	o.TransmitterInfoList = v
}

// GetTimeOfFirstTransmission returns the TimeOfFirstTransmission field value if set, zero value otherwise.
func (o *PC5ContainerInformation) GetTimeOfFirstTransmission() time.Time {
	if o == nil || isNil(o.TimeOfFirstTransmission) {
		var ret time.Time
		return ret
	}
	return *o.TimeOfFirstTransmission
}

// GetTimeOfFirstTransmissionOk returns a tuple with the TimeOfFirstTransmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5ContainerInformation) GetTimeOfFirstTransmissionOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeOfFirstTransmission) {
    return nil, false
	}
	return o.TimeOfFirstTransmission, true
}

// HasTimeOfFirstTransmission returns a boolean if a field has been set.
func (o *PC5ContainerInformation) HasTimeOfFirstTransmission() bool {
	if o != nil && !isNil(o.TimeOfFirstTransmission) {
		return true
	}

	return false
}

// SetTimeOfFirstTransmission gets a reference to the given time.Time and assigns it to the TimeOfFirstTransmission field.
func (o *PC5ContainerInformation) SetTimeOfFirstTransmission(v time.Time) {
	o.TimeOfFirstTransmission = &v
}

// GetTimeOfFirstReception returns the TimeOfFirstReception field value if set, zero value otherwise.
func (o *PC5ContainerInformation) GetTimeOfFirstReception() time.Time {
	if o == nil || isNil(o.TimeOfFirstReception) {
		var ret time.Time
		return ret
	}
	return *o.TimeOfFirstReception
}

// GetTimeOfFirstReceptionOk returns a tuple with the TimeOfFirstReception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5ContainerInformation) GetTimeOfFirstReceptionOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeOfFirstReception) {
    return nil, false
	}
	return o.TimeOfFirstReception, true
}

// HasTimeOfFirstReception returns a boolean if a field has been set.
func (o *PC5ContainerInformation) HasTimeOfFirstReception() bool {
	if o != nil && !isNil(o.TimeOfFirstReception) {
		return true
	}

	return false
}

// SetTimeOfFirstReception gets a reference to the given time.Time and assigns it to the TimeOfFirstReception field.
func (o *PC5ContainerInformation) SetTimeOfFirstReception(v time.Time) {
	o.TimeOfFirstReception = &v
}

func (o PC5ContainerInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CoverageInfoList) {
		toSerialize["coverageInfoList"] = o.CoverageInfoList
	}
	if !isNil(o.RadioParameterSetInfoList) {
		toSerialize["radioParameterSetInfoList"] = o.RadioParameterSetInfoList
	}
	if !isNil(o.TransmitterInfoList) {
		toSerialize["transmitterInfoList"] = o.TransmitterInfoList
	}
	if !isNil(o.TimeOfFirstTransmission) {
		toSerialize["timeOfFirst Transmission"] = o.TimeOfFirstTransmission
	}
	if !isNil(o.TimeOfFirstReception) {
		toSerialize["timeOfFirst Reception"] = o.TimeOfFirstReception
	}
	return json.Marshal(toSerialize)
}

type NullablePC5ContainerInformation struct {
	value *PC5ContainerInformation
	isSet bool
}

func (v NullablePC5ContainerInformation) Get() *PC5ContainerInformation {
	return v.value
}

func (v *NullablePC5ContainerInformation) Set(val *PC5ContainerInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePC5ContainerInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePC5ContainerInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePC5ContainerInformation(val *PC5ContainerInformation) *NullablePC5ContainerInformation {
	return &NullablePC5ContainerInformation{value: val, isSet: true}
}

func (v NullablePC5ContainerInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePC5ContainerInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


