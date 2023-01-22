/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_DataManagement

import (
	"encoding/json"
)

// ChangeOfSupiPeiAssociationReport struct for ChangeOfSupiPeiAssociationReport
type ChangeOfSupiPeiAssociationReport struct {
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	NewPei string `json:"newPei"`
}

// NewChangeOfSupiPeiAssociationReport instantiates a new ChangeOfSupiPeiAssociationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeOfSupiPeiAssociationReport(newPei string) *ChangeOfSupiPeiAssociationReport {
	this := ChangeOfSupiPeiAssociationReport{}
	this.NewPei = newPei
	return &this
}

// NewChangeOfSupiPeiAssociationReportWithDefaults instantiates a new ChangeOfSupiPeiAssociationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeOfSupiPeiAssociationReportWithDefaults() *ChangeOfSupiPeiAssociationReport {
	this := ChangeOfSupiPeiAssociationReport{}
	return &this
}

// GetNewPei returns the NewPei field value
func (o *ChangeOfSupiPeiAssociationReport) GetNewPei() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPei
}

// GetNewPeiOk returns a tuple with the NewPei field value
// and a boolean to check if the value has been set.
func (o *ChangeOfSupiPeiAssociationReport) GetNewPeiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NewPei, true
}

// SetNewPei sets field value
func (o *ChangeOfSupiPeiAssociationReport) SetNewPei(v string) {
	o.NewPei = v
}

func (o ChangeOfSupiPeiAssociationReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["newPei"] = o.NewPei
	}
	return json.Marshal(toSerialize)
}

type NullableChangeOfSupiPeiAssociationReport struct {
	value *ChangeOfSupiPeiAssociationReport
	isSet bool
}

func (v NullableChangeOfSupiPeiAssociationReport) Get() *ChangeOfSupiPeiAssociationReport {
	return v.value
}

func (v *NullableChangeOfSupiPeiAssociationReport) Set(val *ChangeOfSupiPeiAssociationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeOfSupiPeiAssociationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeOfSupiPeiAssociationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeOfSupiPeiAssociationReport(val *ChangeOfSupiPeiAssociationReport) *NullableChangeOfSupiPeiAssociationReport {
	return &NullableChangeOfSupiPeiAssociationReport{value: val, isSet: true}
}

func (v NullableChangeOfSupiPeiAssociationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeOfSupiPeiAssociationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


