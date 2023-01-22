/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_DataRepository

import (
	"encoding/json"
)

// ReportItem indicates performed modivications.
type ReportItem struct {
	// Contains a JSON pointer value (as defined in IETF RFC 6901) that references a  location of a resource to which the modification is subject. 
	Path string `json:"path"`
	// A human-readable reason providing details on the reported modification failure.  The reason string should identify the operation that failed using the operation's  array index to assist in correlation of the invalid parameter with the failed  operation, e.g. \"Replacement value invalid for attribute (failed operation index= 4)\". 
	Reason *string `json:"reason,omitempty"`
}

// NewReportItem instantiates a new ReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportItem(path string) *ReportItem {
	this := ReportItem{}
	this.Path = path
	return &this
}

// NewReportItemWithDefaults instantiates a new ReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportItemWithDefaults() *ReportItem {
	this := ReportItem{}
	return &this
}

// GetPath returns the Path field value
func (o *ReportItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ReportItem) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ReportItem) SetPath(v string) {
	o.Path = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ReportItem) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportItem) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ReportItem) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ReportItem) SetReason(v string) {
	o.Reason = &v
}

func (o ReportItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableReportItem struct {
	value *ReportItem
	isSet bool
}

func (v NullableReportItem) Get() *ReportItem {
	return v.value
}

func (v *NullableReportItem) Set(val *ReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportItem(val *ReportItem) *NullableReportItem {
	return &NullableReportItem{value: val, isSet: true}
}

func (v NullableReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


