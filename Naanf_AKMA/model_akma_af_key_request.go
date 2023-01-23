/*
3gpp-akma

API for Naanf_AKMA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naanf_AKMA

import (
	"encoding/json"
)

// AkmaAfKeyRequest Represents the parameters to request the retrieval of AKMA Application Key information. 
type AkmaAfKeyRequest struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	// Represents an AF identifier.
	AfId string `json:"afId"`
	// Represents an AKMA Key Identifier.
	AKId string `json:"aKId"`
	// Indicates whether an anonymous user access. Set to \"true\" if an anonymous user access is  requested; otherwise set to \"false\". Default value is \"false\" if omitted. 
	AnonInd *bool `json:"anonInd,omitempty"`
}

// NewAkmaAfKeyRequest instantiates a new AkmaAfKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAkmaAfKeyRequest(afId string, aKId string) *AkmaAfKeyRequest {
	this := AkmaAfKeyRequest{}
	this.AfId = afId
	this.AKId = aKId
	var anonInd bool = false
	this.AnonInd = &anonInd
	return &this
}

// NewAkmaAfKeyRequestWithDefaults instantiates a new AkmaAfKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAkmaAfKeyRequestWithDefaults() *AkmaAfKeyRequest {
	this := AkmaAfKeyRequest{}
	var anonInd bool = false
	this.AnonInd = &anonInd
	return &this
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *AkmaAfKeyRequest) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyRequest) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *AkmaAfKeyRequest) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *AkmaAfKeyRequest) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetAfId returns the AfId field value
func (o *AkmaAfKeyRequest) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyRequest) GetAfIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *AkmaAfKeyRequest) SetAfId(v string) {
	o.AfId = v
}

// GetAKId returns the AKId field value
func (o *AkmaAfKeyRequest) GetAKId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AKId
}

// GetAKIdOk returns a tuple with the AKId field value
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyRequest) GetAKIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AKId, true
}

// SetAKId sets field value
func (o *AkmaAfKeyRequest) SetAKId(v string) {
	o.AKId = v
}

// GetAnonInd returns the AnonInd field value if set, zero value otherwise.
func (o *AkmaAfKeyRequest) GetAnonInd() bool {
	if o == nil || isNil(o.AnonInd) {
		var ret bool
		return ret
	}
	return *o.AnonInd
}

// GetAnonIndOk returns a tuple with the AnonInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyRequest) GetAnonIndOk() (*bool, bool) {
	if o == nil || isNil(o.AnonInd) {
    return nil, false
	}
	return o.AnonInd, true
}

// HasAnonInd returns a boolean if a field has been set.
func (o *AkmaAfKeyRequest) HasAnonInd() bool {
	if o != nil && !isNil(o.AnonInd) {
		return true
	}

	return false
}

// SetAnonInd gets a reference to the given bool and assigns it to the AnonInd field.
func (o *AkmaAfKeyRequest) SetAnonInd(v bool) {
	o.AnonInd = &v
}

func (o AkmaAfKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if true {
		toSerialize["afId"] = o.AfId
	}
	if true {
		toSerialize["aKId"] = o.AKId
	}
	if !isNil(o.AnonInd) {
		toSerialize["anonInd"] = o.AnonInd
	}
	return json.Marshal(toSerialize)
}

type NullableAkmaAfKeyRequest struct {
	value *AkmaAfKeyRequest
	isSet bool
}

func (v NullableAkmaAfKeyRequest) Get() *AkmaAfKeyRequest {
	return v.value
}

func (v *NullableAkmaAfKeyRequest) Set(val *AkmaAfKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAkmaAfKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAkmaAfKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAkmaAfKeyRequest(val *AkmaAfKeyRequest) *NullableAkmaAfKeyRequest {
	return &NullableAkmaAfKeyRequest{value: val, isSet: true}
}

func (v NullableAkmaAfKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAkmaAfKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


