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

// ReportingCtrlOneOf2 struct for ReportingCtrlOneOf2
type ReportingCtrlOneOf2 struct {
	StreamTarget *string `json:"streamTarget,omitempty"`
}

// NewReportingCtrlOneOf2 instantiates a new ReportingCtrlOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingCtrlOneOf2() *ReportingCtrlOneOf2 {
	this := ReportingCtrlOneOf2{}
	return &this
}

// NewReportingCtrlOneOf2WithDefaults instantiates a new ReportingCtrlOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingCtrlOneOf2WithDefaults() *ReportingCtrlOneOf2 {
	this := ReportingCtrlOneOf2{}
	return &this
}

// GetStreamTarget returns the StreamTarget field value if set, zero value otherwise.
func (o *ReportingCtrlOneOf2) GetStreamTarget() string {
	if o == nil || isNil(o.StreamTarget) {
		var ret string
		return ret
	}
	return *o.StreamTarget
}

// GetStreamTargetOk returns a tuple with the StreamTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCtrlOneOf2) GetStreamTargetOk() (*string, bool) {
	if o == nil || isNil(o.StreamTarget) {
    return nil, false
	}
	return o.StreamTarget, true
}

// HasStreamTarget returns a boolean if a field has been set.
func (o *ReportingCtrlOneOf2) HasStreamTarget() bool {
	if o != nil && !isNil(o.StreamTarget) {
		return true
	}

	return false
}

// SetStreamTarget gets a reference to the given string and assigns it to the StreamTarget field.
func (o *ReportingCtrlOneOf2) SetStreamTarget(v string) {
	o.StreamTarget = &v
}

func (o ReportingCtrlOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StreamTarget) {
		toSerialize["streamTarget"] = o.StreamTarget
	}
	return json.Marshal(toSerialize)
}

type NullableReportingCtrlOneOf2 struct {
	value *ReportingCtrlOneOf2
	isSet bool
}

func (v NullableReportingCtrlOneOf2) Get() *ReportingCtrlOneOf2 {
	return v.value
}

func (v *NullableReportingCtrlOneOf2) Set(val *ReportingCtrlOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingCtrlOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingCtrlOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingCtrlOneOf2(val *ReportingCtrlOneOf2) *NullableReportingCtrlOneOf2 {
	return &NullableReportingCtrlOneOf2{value: val, isSet: true}
}

func (v NullableReportingCtrlOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingCtrlOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


