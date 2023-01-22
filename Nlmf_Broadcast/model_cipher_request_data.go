/*
LMF Broadcast

LMF Broadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nlmf_Broadcast

import (
	"encoding/json"
)

// CipherRequestData Information within Ciphering Key Data request.
type CipherRequestData struct {
	// String providing an URI formatted according to RFC 3986.
	AmfCallBackURI string `json:"amfCallBackURI"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewCipherRequestData instantiates a new CipherRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCipherRequestData(amfCallBackURI string) *CipherRequestData {
	this := CipherRequestData{}
	this.AmfCallBackURI = amfCallBackURI
	return &this
}

// NewCipherRequestDataWithDefaults instantiates a new CipherRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCipherRequestDataWithDefaults() *CipherRequestData {
	this := CipherRequestData{}
	return &this
}

// GetAmfCallBackURI returns the AmfCallBackURI field value
func (o *CipherRequestData) GetAmfCallBackURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfCallBackURI
}

// GetAmfCallBackURIOk returns a tuple with the AmfCallBackURI field value
// and a boolean to check if the value has been set.
func (o *CipherRequestData) GetAmfCallBackURIOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AmfCallBackURI, true
}

// SetAmfCallBackURI sets field value
func (o *CipherRequestData) SetAmfCallBackURI(v string) {
	o.AmfCallBackURI = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *CipherRequestData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CipherRequestData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *CipherRequestData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *CipherRequestData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o CipherRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amfCallBackURI"] = o.AmfCallBackURI
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableCipherRequestData struct {
	value *CipherRequestData
	isSet bool
}

func (v NullableCipherRequestData) Get() *CipherRequestData {
	return v.value
}

func (v *NullableCipherRequestData) Set(val *CipherRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCipherRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCipherRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCipherRequestData(val *CipherRequestData) *NullableCipherRequestData {
	return &NullableCipherRequestData{value: val, isSet: true}
}

func (v NullableCipherRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCipherRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


