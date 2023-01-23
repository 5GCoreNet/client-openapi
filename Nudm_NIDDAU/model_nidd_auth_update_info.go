/*
Nudm_NIDDAU

Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_NIDDAU

import (
	"encoding/json"
)

// NiddAuthUpdateInfo Represents NIDD authorization update information.
type NiddAuthUpdateInfo struct {
	AuthorizationData AuthorizationData `json:"authorizationData"`
	InvalidityInd *bool `json:"invalidityInd,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	NiddCause *NiddCause `json:"niddCause,omitempty"`
}

// NewNiddAuthUpdateInfo instantiates a new NiddAuthUpdateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiddAuthUpdateInfo(authorizationData AuthorizationData) *NiddAuthUpdateInfo {
	this := NiddAuthUpdateInfo{}
	this.AuthorizationData = authorizationData
	return &this
}

// NewNiddAuthUpdateInfoWithDefaults instantiates a new NiddAuthUpdateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiddAuthUpdateInfoWithDefaults() *NiddAuthUpdateInfo {
	this := NiddAuthUpdateInfo{}
	return &this
}

// GetAuthorizationData returns the AuthorizationData field value
func (o *NiddAuthUpdateInfo) GetAuthorizationData() AuthorizationData {
	if o == nil {
		var ret AuthorizationData
		return ret
	}

	return o.AuthorizationData
}

// GetAuthorizationDataOk returns a tuple with the AuthorizationData field value
// and a boolean to check if the value has been set.
func (o *NiddAuthUpdateInfo) GetAuthorizationDataOk() (*AuthorizationData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthorizationData, true
}

// SetAuthorizationData sets field value
func (o *NiddAuthUpdateInfo) SetAuthorizationData(v AuthorizationData) {
	o.AuthorizationData = v
}

// GetInvalidityInd returns the InvalidityInd field value if set, zero value otherwise.
func (o *NiddAuthUpdateInfo) GetInvalidityInd() bool {
	if o == nil || isNil(o.InvalidityInd) {
		var ret bool
		return ret
	}
	return *o.InvalidityInd
}

// GetInvalidityIndOk returns a tuple with the InvalidityInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddAuthUpdateInfo) GetInvalidityIndOk() (*bool, bool) {
	if o == nil || isNil(o.InvalidityInd) {
    return nil, false
	}
	return o.InvalidityInd, true
}

// HasInvalidityInd returns a boolean if a field has been set.
func (o *NiddAuthUpdateInfo) HasInvalidityInd() bool {
	if o != nil && !isNil(o.InvalidityInd) {
		return true
	}

	return false
}

// SetInvalidityInd gets a reference to the given bool and assigns it to the InvalidityInd field.
func (o *NiddAuthUpdateInfo) SetInvalidityInd(v bool) {
	o.InvalidityInd = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *NiddAuthUpdateInfo) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddAuthUpdateInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *NiddAuthUpdateInfo) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *NiddAuthUpdateInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *NiddAuthUpdateInfo) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddAuthUpdateInfo) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *NiddAuthUpdateInfo) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *NiddAuthUpdateInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetNiddCause returns the NiddCause field value if set, zero value otherwise.
func (o *NiddAuthUpdateInfo) GetNiddCause() NiddCause {
	if o == nil || isNil(o.NiddCause) {
		var ret NiddCause
		return ret
	}
	return *o.NiddCause
}

// GetNiddCauseOk returns a tuple with the NiddCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddAuthUpdateInfo) GetNiddCauseOk() (*NiddCause, bool) {
	if o == nil || isNil(o.NiddCause) {
    return nil, false
	}
	return o.NiddCause, true
}

// HasNiddCause returns a boolean if a field has been set.
func (o *NiddAuthUpdateInfo) HasNiddCause() bool {
	if o != nil && !isNil(o.NiddCause) {
		return true
	}

	return false
}

// SetNiddCause gets a reference to the given NiddCause and assigns it to the NiddCause field.
func (o *NiddAuthUpdateInfo) SetNiddCause(v NiddCause) {
	o.NiddCause = &v
}

func (o NiddAuthUpdateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authorizationData"] = o.AuthorizationData
	}
	if !isNil(o.InvalidityInd) {
		toSerialize["invalidityInd"] = o.InvalidityInd
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.NiddCause) {
		toSerialize["niddCause"] = o.NiddCause
	}
	return json.Marshal(toSerialize)
}

type NullableNiddAuthUpdateInfo struct {
	value *NiddAuthUpdateInfo
	isSet bool
}

func (v NullableNiddAuthUpdateInfo) Get() *NiddAuthUpdateInfo {
	return v.value
}

func (v *NullableNiddAuthUpdateInfo) Set(val *NiddAuthUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNiddAuthUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNiddAuthUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiddAuthUpdateInfo(val *NiddAuthUpdateInfo) *NullableNiddAuthUpdateInfo {
	return &NullableNiddAuthUpdateInfo{value: val, isSet: true}
}

func (v NullableNiddAuthUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiddAuthUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


