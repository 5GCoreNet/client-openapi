/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
)

// ConfirmationData Contains the result of the authentication.
type ConfirmationData struct {
	// Contains the RES*.
	ResStar NullableString `json:"resStar"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewConfirmationData instantiates a new ConfirmationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmationData(resStar NullableString) *ConfirmationData {
	this := ConfirmationData{}
	this.ResStar = resStar
	return &this
}

// NewConfirmationDataWithDefaults instantiates a new ConfirmationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmationDataWithDefaults() *ConfirmationData {
	this := ConfirmationData{}
	return &this
}

// GetResStar returns the ResStar field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConfirmationData) GetResStar() string {
	if o == nil || o.ResStar.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResStar.Get()
}

// GetResStarOk returns a tuple with the ResStar field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfirmationData) GetResStarOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ResStar.Get(), o.ResStar.IsSet()
}

// SetResStar sets field value
func (o *ConfirmationData) SetResStar(v string) {
	o.ResStar.Set(&v)
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ConfirmationData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfirmationData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ConfirmationData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ConfirmationData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o ConfirmationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resStar"] = o.ResStar.Get()
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableConfirmationData struct {
	value *ConfirmationData
	isSet bool
}

func (v NullableConfirmationData) Get() *ConfirmationData {
	return v.value
}

func (v *NullableConfirmationData) Set(val *ConfirmationData) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmationData) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmationData(val *ConfirmationData) *NullableConfirmationData {
	return &NullableConfirmationData{value: val, isSet: true}
}

func (v NullableConfirmationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


