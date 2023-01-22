/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
)

// ReportingCtrlOneOf1 struct for ReportingCtrlOneOf1
type ReportingCtrlOneOf1 struct {
	FileReportingPeriod *int32 `json:"fileReportingPeriod,omitempty"`
	FileLocation *string `json:"fileLocation,omitempty"`
}

// NewReportingCtrlOneOf1 instantiates a new ReportingCtrlOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingCtrlOneOf1() *ReportingCtrlOneOf1 {
	this := ReportingCtrlOneOf1{}
	return &this
}

// NewReportingCtrlOneOf1WithDefaults instantiates a new ReportingCtrlOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingCtrlOneOf1WithDefaults() *ReportingCtrlOneOf1 {
	this := ReportingCtrlOneOf1{}
	return &this
}

// GetFileReportingPeriod returns the FileReportingPeriod field value if set, zero value otherwise.
func (o *ReportingCtrlOneOf1) GetFileReportingPeriod() int32 {
	if o == nil || isNil(o.FileReportingPeriod) {
		var ret int32
		return ret
	}
	return *o.FileReportingPeriod
}

// GetFileReportingPeriodOk returns a tuple with the FileReportingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCtrlOneOf1) GetFileReportingPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.FileReportingPeriod) {
    return nil, false
	}
	return o.FileReportingPeriod, true
}

// HasFileReportingPeriod returns a boolean if a field has been set.
func (o *ReportingCtrlOneOf1) HasFileReportingPeriod() bool {
	if o != nil && !isNil(o.FileReportingPeriod) {
		return true
	}

	return false
}

// SetFileReportingPeriod gets a reference to the given int32 and assigns it to the FileReportingPeriod field.
func (o *ReportingCtrlOneOf1) SetFileReportingPeriod(v int32) {
	o.FileReportingPeriod = &v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *ReportingCtrlOneOf1) GetFileLocation() string {
	if o == nil || isNil(o.FileLocation) {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCtrlOneOf1) GetFileLocationOk() (*string, bool) {
	if o == nil || isNil(o.FileLocation) {
    return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *ReportingCtrlOneOf1) HasFileLocation() bool {
	if o != nil && !isNil(o.FileLocation) {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *ReportingCtrlOneOf1) SetFileLocation(v string) {
	o.FileLocation = &v
}

func (o ReportingCtrlOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FileReportingPeriod) {
		toSerialize["fileReportingPeriod"] = o.FileReportingPeriod
	}
	if !isNil(o.FileLocation) {
		toSerialize["fileLocation"] = o.FileLocation
	}
	return json.Marshal(toSerialize)
}

type NullableReportingCtrlOneOf1 struct {
	value *ReportingCtrlOneOf1
	isSet bool
}

func (v NullableReportingCtrlOneOf1) Get() *ReportingCtrlOneOf1 {
	return v.value
}

func (v *NullableReportingCtrlOneOf1) Set(val *ReportingCtrlOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingCtrlOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingCtrlOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingCtrlOneOf1(val *ReportingCtrlOneOf1) *NullableReportingCtrlOneOf1 {
	return &NullableReportingCtrlOneOf1{value: val, isSet: true}
}

func (v NullableReportingCtrlOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingCtrlOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


