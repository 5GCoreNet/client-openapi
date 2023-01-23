/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMeasJobCtrlMnS

import (
	"encoding/json"
)

// MeasJobsRetrievalResponseType struct for MeasJobsRetrievalResponseType
type MeasJobsRetrievalResponseType struct {
	JobInfoList []MeasJobInfoResourceType `json:"jobInfoList,omitempty"`
}

// NewMeasJobsRetrievalResponseType instantiates a new MeasJobsRetrievalResponseType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasJobsRetrievalResponseType() *MeasJobsRetrievalResponseType {
	this := MeasJobsRetrievalResponseType{}
	return &this
}

// NewMeasJobsRetrievalResponseTypeWithDefaults instantiates a new MeasJobsRetrievalResponseType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasJobsRetrievalResponseTypeWithDefaults() *MeasJobsRetrievalResponseType {
	this := MeasJobsRetrievalResponseType{}
	return &this
}

// GetJobInfoList returns the JobInfoList field value if set, zero value otherwise.
func (o *MeasJobsRetrievalResponseType) GetJobInfoList() []MeasJobInfoResourceType {
	if o == nil || isNil(o.JobInfoList) {
		var ret []MeasJobInfoResourceType
		return ret
	}
	return o.JobInfoList
}

// GetJobInfoListOk returns a tuple with the JobInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobsRetrievalResponseType) GetJobInfoListOk() ([]MeasJobInfoResourceType, bool) {
	if o == nil || isNil(o.JobInfoList) {
    return nil, false
	}
	return o.JobInfoList, true
}

// HasJobInfoList returns a boolean if a field has been set.
func (o *MeasJobsRetrievalResponseType) HasJobInfoList() bool {
	if o != nil && !isNil(o.JobInfoList) {
		return true
	}

	return false
}

// SetJobInfoList gets a reference to the given []MeasJobInfoResourceType and assigns it to the JobInfoList field.
func (o *MeasJobsRetrievalResponseType) SetJobInfoList(v []MeasJobInfoResourceType) {
	o.JobInfoList = v
}

func (o MeasJobsRetrievalResponseType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JobInfoList) {
		toSerialize["jobInfoList"] = o.JobInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableMeasJobsRetrievalResponseType struct {
	value *MeasJobsRetrievalResponseType
	isSet bool
}

func (v NullableMeasJobsRetrievalResponseType) Get() *MeasJobsRetrievalResponseType {
	return v.value
}

func (v *NullableMeasJobsRetrievalResponseType) Set(val *MeasJobsRetrievalResponseType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasJobsRetrievalResponseType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasJobsRetrievalResponseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasJobsRetrievalResponseType(val *MeasJobsRetrievalResponseType) *NullableMeasJobsRetrievalResponseType {
	return &NullableMeasJobsRetrievalResponseType{value: val, isSet: true}
}

func (v NullableMeasJobsRetrievalResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasJobsRetrievalResponseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


