/*
EES ACR Status Update Service

EES ACR Status Update Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRStatusUpdate

import (
	"encoding/json"
)

// ACRDataStatus Represents the ACR status information.
type ACRDataStatus struct {
	E3SubscsStatus E3SubscsStatus `json:"e3SubscsStatus"`
	E3SubscIds []string `json:"e3SubscIds,omitempty"`
}

// NewACRDataStatus instantiates a new ACRDataStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACRDataStatus(e3SubscsStatus E3SubscsStatus) *ACRDataStatus {
	this := ACRDataStatus{}
	this.E3SubscsStatus = e3SubscsStatus
	return &this
}

// NewACRDataStatusWithDefaults instantiates a new ACRDataStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACRDataStatusWithDefaults() *ACRDataStatus {
	this := ACRDataStatus{}
	return &this
}

// GetE3SubscsStatus returns the E3SubscsStatus field value
func (o *ACRDataStatus) GetE3SubscsStatus() E3SubscsStatus {
	if o == nil {
		var ret E3SubscsStatus
		return ret
	}

	return o.E3SubscsStatus
}

// GetE3SubscsStatusOk returns a tuple with the E3SubscsStatus field value
// and a boolean to check if the value has been set.
func (o *ACRDataStatus) GetE3SubscsStatusOk() (*E3SubscsStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.E3SubscsStatus, true
}

// SetE3SubscsStatus sets field value
func (o *ACRDataStatus) SetE3SubscsStatus(v E3SubscsStatus) {
	o.E3SubscsStatus = v
}

// GetE3SubscIds returns the E3SubscIds field value if set, zero value otherwise.
func (o *ACRDataStatus) GetE3SubscIds() []string {
	if o == nil || isNil(o.E3SubscIds) {
		var ret []string
		return ret
	}
	return o.E3SubscIds
}

// GetE3SubscIdsOk returns a tuple with the E3SubscIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRDataStatus) GetE3SubscIdsOk() ([]string, bool) {
	if o == nil || isNil(o.E3SubscIds) {
    return nil, false
	}
	return o.E3SubscIds, true
}

// HasE3SubscIds returns a boolean if a field has been set.
func (o *ACRDataStatus) HasE3SubscIds() bool {
	if o != nil && !isNil(o.E3SubscIds) {
		return true
	}

	return false
}

// SetE3SubscIds gets a reference to the given []string and assigns it to the E3SubscIds field.
func (o *ACRDataStatus) SetE3SubscIds(v []string) {
	o.E3SubscIds = v
}

func (o ACRDataStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["e3SubscsStatus"] = o.E3SubscsStatus
	}
	if !isNil(o.E3SubscIds) {
		toSerialize["e3SubscIds"] = o.E3SubscIds
	}
	return json.Marshal(toSerialize)
}

type NullableACRDataStatus struct {
	value *ACRDataStatus
	isSet bool
}

func (v NullableACRDataStatus) Get() *ACRDataStatus {
	return v.value
}

func (v *NullableACRDataStatus) Set(val *ACRDataStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableACRDataStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableACRDataStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACRDataStatus(val *ACRDataStatus) *NullableACRDataStatus {
	return &NullableACRDataStatus{value: val, isSet: true}
}

func (v NullableACRDataStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACRDataStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


