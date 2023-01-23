/*
3gpp-eas-deployment

API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EASDeployment

import (
	"encoding/json"
)

// DnnSnssaiInformation Represents a (DNN, SNSSAI) combination.
type DnnSnssaiInformation struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
}

// NewDnnSnssaiInformation instantiates a new DnnSnssaiInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnSnssaiInformation() *DnnSnssaiInformation {
	this := DnnSnssaiInformation{}
	return &this
}

// NewDnnSnssaiInformationWithDefaults instantiates a new DnnSnssaiInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnSnssaiInformationWithDefaults() *DnnSnssaiInformation {
	this := DnnSnssaiInformation{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *DnnSnssaiInformation) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnSnssaiInformation) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *DnnSnssaiInformation) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *DnnSnssaiInformation) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *DnnSnssaiInformation) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnSnssaiInformation) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *DnnSnssaiInformation) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *DnnSnssaiInformation) SetSnssai(v Snssai) {
	o.Snssai = &v
}

func (o DnnSnssaiInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	return json.Marshal(toSerialize)
}

type NullableDnnSnssaiInformation struct {
	value *DnnSnssaiInformation
	isSet bool
}

func (v NullableDnnSnssaiInformation) Get() *DnnSnssaiInformation {
	return v.value
}

func (v *NullableDnnSnssaiInformation) Set(val *DnnSnssaiInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnSnssaiInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnSnssaiInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnSnssaiInformation(val *DnnSnssaiInformation) *NullableDnnSnssaiInformation {
	return &NullableDnnSnssaiInformation{value: val, isSet: true}
}

func (v NullableDnnSnssaiInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnSnssaiInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


