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

// AccessNetworkInfoChange struct for AccessNetworkInfoChange
type AccessNetworkInfoChange struct {
	AccessNetworkInformation []string `json:"accessNetworkInformation,omitempty"`
	CellularNetworkInformation *string `json:"cellularNetworkInformation,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ChangeTime *time.Time `json:"changeTime,omitempty"`
}

// NewAccessNetworkInfoChange instantiates a new AccessNetworkInfoChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessNetworkInfoChange() *AccessNetworkInfoChange {
	this := AccessNetworkInfoChange{}
	return &this
}

// NewAccessNetworkInfoChangeWithDefaults instantiates a new AccessNetworkInfoChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessNetworkInfoChangeWithDefaults() *AccessNetworkInfoChange {
	this := AccessNetworkInfoChange{}
	return &this
}

// GetAccessNetworkInformation returns the AccessNetworkInformation field value if set, zero value otherwise.
func (o *AccessNetworkInfoChange) GetAccessNetworkInformation() []string {
	if o == nil || isNil(o.AccessNetworkInformation) {
		var ret []string
		return ret
	}
	return o.AccessNetworkInformation
}

// GetAccessNetworkInformationOk returns a tuple with the AccessNetworkInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessNetworkInfoChange) GetAccessNetworkInformationOk() ([]string, bool) {
	if o == nil || isNil(o.AccessNetworkInformation) {
    return nil, false
	}
	return o.AccessNetworkInformation, true
}

// HasAccessNetworkInformation returns a boolean if a field has been set.
func (o *AccessNetworkInfoChange) HasAccessNetworkInformation() bool {
	if o != nil && !isNil(o.AccessNetworkInformation) {
		return true
	}

	return false
}

// SetAccessNetworkInformation gets a reference to the given []string and assigns it to the AccessNetworkInformation field.
func (o *AccessNetworkInfoChange) SetAccessNetworkInformation(v []string) {
	o.AccessNetworkInformation = v
}

// GetCellularNetworkInformation returns the CellularNetworkInformation field value if set, zero value otherwise.
func (o *AccessNetworkInfoChange) GetCellularNetworkInformation() string {
	if o == nil || isNil(o.CellularNetworkInformation) {
		var ret string
		return ret
	}
	return *o.CellularNetworkInformation
}

// GetCellularNetworkInformationOk returns a tuple with the CellularNetworkInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessNetworkInfoChange) GetCellularNetworkInformationOk() (*string, bool) {
	if o == nil || isNil(o.CellularNetworkInformation) {
    return nil, false
	}
	return o.CellularNetworkInformation, true
}

// HasCellularNetworkInformation returns a boolean if a field has been set.
func (o *AccessNetworkInfoChange) HasCellularNetworkInformation() bool {
	if o != nil && !isNil(o.CellularNetworkInformation) {
		return true
	}

	return false
}

// SetCellularNetworkInformation gets a reference to the given string and assigns it to the CellularNetworkInformation field.
func (o *AccessNetworkInfoChange) SetCellularNetworkInformation(v string) {
	o.CellularNetworkInformation = &v
}

// GetChangeTime returns the ChangeTime field value if set, zero value otherwise.
func (o *AccessNetworkInfoChange) GetChangeTime() time.Time {
	if o == nil || isNil(o.ChangeTime) {
		var ret time.Time
		return ret
	}
	return *o.ChangeTime
}

// GetChangeTimeOk returns a tuple with the ChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessNetworkInfoChange) GetChangeTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ChangeTime) {
    return nil, false
	}
	return o.ChangeTime, true
}

// HasChangeTime returns a boolean if a field has been set.
func (o *AccessNetworkInfoChange) HasChangeTime() bool {
	if o != nil && !isNil(o.ChangeTime) {
		return true
	}

	return false
}

// SetChangeTime gets a reference to the given time.Time and assigns it to the ChangeTime field.
func (o *AccessNetworkInfoChange) SetChangeTime(v time.Time) {
	o.ChangeTime = &v
}

func (o AccessNetworkInfoChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessNetworkInformation) {
		toSerialize["accessNetworkInformation"] = o.AccessNetworkInformation
	}
	if !isNil(o.CellularNetworkInformation) {
		toSerialize["cellularNetworkInformation"] = o.CellularNetworkInformation
	}
	if !isNil(o.ChangeTime) {
		toSerialize["changeTime"] = o.ChangeTime
	}
	return json.Marshal(toSerialize)
}

type NullableAccessNetworkInfoChange struct {
	value *AccessNetworkInfoChange
	isSet bool
}

func (v NullableAccessNetworkInfoChange) Get() *AccessNetworkInfoChange {
	return v.value
}

func (v *NullableAccessNetworkInfoChange) Set(val *AccessNetworkInfoChange) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessNetworkInfoChange) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessNetworkInfoChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessNetworkInfoChange(val *AccessNetworkInfoChange) *NullableAccessNetworkInfoChange {
	return &NullableAccessNetworkInfoChange{value: val, isSet: true}
}

func (v NullableAccessNetworkInfoChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessNetworkInfoChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


