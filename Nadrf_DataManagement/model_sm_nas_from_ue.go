/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"time"
)

// SmNasFromUe Represents information on the SM NAS messages that SMF receives from UE for PDU Session. 
type SmNasFromUe struct {
	SmNasType string `json:"smNasType"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}

// NewSmNasFromUe instantiates a new SmNasFromUe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmNasFromUe(smNasType string, timeStamp time.Time) *SmNasFromUe {
	this := SmNasFromUe{}
	this.SmNasType = smNasType
	this.TimeStamp = timeStamp
	return &this
}

// NewSmNasFromUeWithDefaults instantiates a new SmNasFromUe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmNasFromUeWithDefaults() *SmNasFromUe {
	this := SmNasFromUe{}
	return &this
}

// GetSmNasType returns the SmNasType field value
func (o *SmNasFromUe) GetSmNasType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmNasType
}

// GetSmNasTypeOk returns a tuple with the SmNasType field value
// and a boolean to check if the value has been set.
func (o *SmNasFromUe) GetSmNasTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SmNasType, true
}

// SetSmNasType sets field value
func (o *SmNasFromUe) SetSmNasType(v string) {
	o.SmNasType = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *SmNasFromUe) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *SmNasFromUe) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *SmNasFromUe) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

func (o SmNasFromUe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["smNasType"] = o.SmNasType
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	return json.Marshal(toSerialize)
}

type NullableSmNasFromUe struct {
	value *SmNasFromUe
	isSet bool
}

func (v NullableSmNasFromUe) Get() *SmNasFromUe {
	return v.value
}

func (v *NullableSmNasFromUe) Set(val *SmNasFromUe) {
	v.value = val
	v.isSet = true
}

func (v NullableSmNasFromUe) IsSet() bool {
	return v.isSet
}

func (v *NullableSmNasFromUe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmNasFromUe(val *SmNasFromUe) *NullableSmNasFromUe {
	return &NullableSmNasFromUe{value: val, isSet: true}
}

func (v NullableSmNasFromUe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmNasFromUe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


