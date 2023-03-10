/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// RANSecondaryRATUsageReport struct for RANSecondaryRATUsageReport
type RANSecondaryRATUsageReport struct {
	RANSecondaryRATType *RatType `json:"rANSecondaryRATType,omitempty"`
	QosFlowsUsageReports []QosFlowsUsageReport `json:"qosFlowsUsageReports,omitempty"`
}

// NewRANSecondaryRATUsageReport instantiates a new RANSecondaryRATUsageReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRANSecondaryRATUsageReport() *RANSecondaryRATUsageReport {
	this := RANSecondaryRATUsageReport{}
	return &this
}

// NewRANSecondaryRATUsageReportWithDefaults instantiates a new RANSecondaryRATUsageReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRANSecondaryRATUsageReportWithDefaults() *RANSecondaryRATUsageReport {
	this := RANSecondaryRATUsageReport{}
	return &this
}

// GetRANSecondaryRATType returns the RANSecondaryRATType field value if set, zero value otherwise.
func (o *RANSecondaryRATUsageReport) GetRANSecondaryRATType() RatType {
	if o == nil || isNil(o.RANSecondaryRATType) {
		var ret RatType
		return ret
	}
	return *o.RANSecondaryRATType
}

// GetRANSecondaryRATTypeOk returns a tuple with the RANSecondaryRATType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RANSecondaryRATUsageReport) GetRANSecondaryRATTypeOk() (*RatType, bool) {
	if o == nil || isNil(o.RANSecondaryRATType) {
    return nil, false
	}
	return o.RANSecondaryRATType, true
}

// HasRANSecondaryRATType returns a boolean if a field has been set.
func (o *RANSecondaryRATUsageReport) HasRANSecondaryRATType() bool {
	if o != nil && !isNil(o.RANSecondaryRATType) {
		return true
	}

	return false
}

// SetRANSecondaryRATType gets a reference to the given RatType and assigns it to the RANSecondaryRATType field.
func (o *RANSecondaryRATUsageReport) SetRANSecondaryRATType(v RatType) {
	o.RANSecondaryRATType = &v
}

// GetQosFlowsUsageReports returns the QosFlowsUsageReports field value if set, zero value otherwise.
func (o *RANSecondaryRATUsageReport) GetQosFlowsUsageReports() []QosFlowsUsageReport {
	if o == nil || isNil(o.QosFlowsUsageReports) {
		var ret []QosFlowsUsageReport
		return ret
	}
	return o.QosFlowsUsageReports
}

// GetQosFlowsUsageReportsOk returns a tuple with the QosFlowsUsageReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RANSecondaryRATUsageReport) GetQosFlowsUsageReportsOk() ([]QosFlowsUsageReport, bool) {
	if o == nil || isNil(o.QosFlowsUsageReports) {
    return nil, false
	}
	return o.QosFlowsUsageReports, true
}

// HasQosFlowsUsageReports returns a boolean if a field has been set.
func (o *RANSecondaryRATUsageReport) HasQosFlowsUsageReports() bool {
	if o != nil && !isNil(o.QosFlowsUsageReports) {
		return true
	}

	return false
}

// SetQosFlowsUsageReports gets a reference to the given []QosFlowsUsageReport and assigns it to the QosFlowsUsageReports field.
func (o *RANSecondaryRATUsageReport) SetQosFlowsUsageReports(v []QosFlowsUsageReport) {
	o.QosFlowsUsageReports = v
}

func (o RANSecondaryRATUsageReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RANSecondaryRATType) {
		toSerialize["rANSecondaryRATType"] = o.RANSecondaryRATType
	}
	if !isNil(o.QosFlowsUsageReports) {
		toSerialize["qosFlowsUsageReports"] = o.QosFlowsUsageReports
	}
	return json.Marshal(toSerialize)
}

type NullableRANSecondaryRATUsageReport struct {
	value *RANSecondaryRATUsageReport
	isSet bool
}

func (v NullableRANSecondaryRATUsageReport) Get() *RANSecondaryRATUsageReport {
	return v.value
}

func (v *NullableRANSecondaryRATUsageReport) Set(val *RANSecondaryRATUsageReport) {
	v.value = val
	v.isSet = true
}

func (v NullableRANSecondaryRATUsageReport) IsSet() bool {
	return v.isSet
}

func (v *NullableRANSecondaryRATUsageReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRANSecondaryRATUsageReport(val *RANSecondaryRATUsageReport) *NullableRANSecondaryRATUsageReport {
	return &NullableRANSecondaryRATUsageReport{value: val, isSet: true}
}

func (v NullableRANSecondaryRATUsageReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRANSecondaryRATUsageReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


