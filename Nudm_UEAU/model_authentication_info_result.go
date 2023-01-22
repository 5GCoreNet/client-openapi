/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_UEAU

import (
	"encoding/json"
)

// AuthenticationInfoResult struct for AuthenticationInfoResult
type AuthenticationInfoResult struct {
	AuthType AuthType `json:"authType"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	AuthenticationVector *AuthenticationVector `json:"authenticationVector,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	AkmaInd *bool `json:"akmaInd,omitempty"`
	AuthAaa *bool `json:"authAaa,omitempty"`
	RoutingId *string `json:"routingId,omitempty"`
	PvsInfo []ServerAddressingInfo `json:"pvsInfo,omitempty"`
}

// NewAuthenticationInfoResult instantiates a new AuthenticationInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationInfoResult(authType AuthType) *AuthenticationInfoResult {
	this := AuthenticationInfoResult{}
	this.AuthType = authType
	var akmaInd bool = false
	this.AkmaInd = &akmaInd
	var authAaa bool = false
	this.AuthAaa = &authAaa
	return &this
}

// NewAuthenticationInfoResultWithDefaults instantiates a new AuthenticationInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationInfoResultWithDefaults() *AuthenticationInfoResult {
	this := AuthenticationInfoResult{}
	var akmaInd bool = false
	this.AkmaInd = &akmaInd
	var authAaa bool = false
	this.AuthAaa = &authAaa
	return &this
}

// GetAuthType returns the AuthType field value
func (o *AuthenticationInfoResult) GetAuthType() AuthType {
	if o == nil {
		var ret AuthType
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetAuthTypeOk() (*AuthType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *AuthenticationInfoResult) SetAuthType(v AuthType) {
	o.AuthType = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AuthenticationInfoResult) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAuthenticationVector returns the AuthenticationVector field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetAuthenticationVector() AuthenticationVector {
	if o == nil || isNil(o.AuthenticationVector) {
		var ret AuthenticationVector
		return ret
	}
	return *o.AuthenticationVector
}

// GetAuthenticationVectorOk returns a tuple with the AuthenticationVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetAuthenticationVectorOk() (*AuthenticationVector, bool) {
	if o == nil || isNil(o.AuthenticationVector) {
    return nil, false
	}
	return o.AuthenticationVector, true
}

// HasAuthenticationVector returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasAuthenticationVector() bool {
	if o != nil && !isNil(o.AuthenticationVector) {
		return true
	}

	return false
}

// SetAuthenticationVector gets a reference to the given AuthenticationVector and assigns it to the AuthenticationVector field.
func (o *AuthenticationInfoResult) SetAuthenticationVector(v AuthenticationVector) {
	o.AuthenticationVector = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *AuthenticationInfoResult) SetSupi(v string) {
	o.Supi = &v
}

// GetAkmaInd returns the AkmaInd field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetAkmaInd() bool {
	if o == nil || isNil(o.AkmaInd) {
		var ret bool
		return ret
	}
	return *o.AkmaInd
}

// GetAkmaIndOk returns a tuple with the AkmaInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetAkmaIndOk() (*bool, bool) {
	if o == nil || isNil(o.AkmaInd) {
    return nil, false
	}
	return o.AkmaInd, true
}

// HasAkmaInd returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasAkmaInd() bool {
	if o != nil && !isNil(o.AkmaInd) {
		return true
	}

	return false
}

// SetAkmaInd gets a reference to the given bool and assigns it to the AkmaInd field.
func (o *AuthenticationInfoResult) SetAkmaInd(v bool) {
	o.AkmaInd = &v
}

// GetAuthAaa returns the AuthAaa field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetAuthAaa() bool {
	if o == nil || isNil(o.AuthAaa) {
		var ret bool
		return ret
	}
	return *o.AuthAaa
}

// GetAuthAaaOk returns a tuple with the AuthAaa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetAuthAaaOk() (*bool, bool) {
	if o == nil || isNil(o.AuthAaa) {
    return nil, false
	}
	return o.AuthAaa, true
}

// HasAuthAaa returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasAuthAaa() bool {
	if o != nil && !isNil(o.AuthAaa) {
		return true
	}

	return false
}

// SetAuthAaa gets a reference to the given bool and assigns it to the AuthAaa field.
func (o *AuthenticationInfoResult) SetAuthAaa(v bool) {
	o.AuthAaa = &v
}

// GetRoutingId returns the RoutingId field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetRoutingId() string {
	if o == nil || isNil(o.RoutingId) {
		var ret string
		return ret
	}
	return *o.RoutingId
}

// GetRoutingIdOk returns a tuple with the RoutingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetRoutingIdOk() (*string, bool) {
	if o == nil || isNil(o.RoutingId) {
    return nil, false
	}
	return o.RoutingId, true
}

// HasRoutingId returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasRoutingId() bool {
	if o != nil && !isNil(o.RoutingId) {
		return true
	}

	return false
}

// SetRoutingId gets a reference to the given string and assigns it to the RoutingId field.
func (o *AuthenticationInfoResult) SetRoutingId(v string) {
	o.RoutingId = &v
}

// GetPvsInfo returns the PvsInfo field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetPvsInfo() []ServerAddressingInfo {
	if o == nil || isNil(o.PvsInfo) {
		var ret []ServerAddressingInfo
		return ret
	}
	return o.PvsInfo
}

// GetPvsInfoOk returns a tuple with the PvsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetPvsInfoOk() ([]ServerAddressingInfo, bool) {
	if o == nil || isNil(o.PvsInfo) {
    return nil, false
	}
	return o.PvsInfo, true
}

// HasPvsInfo returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasPvsInfo() bool {
	if o != nil && !isNil(o.PvsInfo) {
		return true
	}

	return false
}

// SetPvsInfo gets a reference to the given []ServerAddressingInfo and assigns it to the PvsInfo field.
func (o *AuthenticationInfoResult) SetPvsInfo(v []ServerAddressingInfo) {
	o.PvsInfo = v
}

func (o AuthenticationInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authType"] = o.AuthType
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.AuthenticationVector) {
		toSerialize["authenticationVector"] = o.AuthenticationVector
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.AkmaInd) {
		toSerialize["akmaInd"] = o.AkmaInd
	}
	if !isNil(o.AuthAaa) {
		toSerialize["authAaa"] = o.AuthAaa
	}
	if !isNil(o.RoutingId) {
		toSerialize["routingId"] = o.RoutingId
	}
	if !isNil(o.PvsInfo) {
		toSerialize["pvsInfo"] = o.PvsInfo
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationInfoResult struct {
	value *AuthenticationInfoResult
	isSet bool
}

func (v NullableAuthenticationInfoResult) Get() *AuthenticationInfoResult {
	return v.value
}

func (v *NullableAuthenticationInfoResult) Set(val *AuthenticationInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationInfoResult(val *AuthenticationInfoResult) *NullableAuthenticationInfoResult {
	return &NullableAuthenticationInfoResult{value: val, isSet: true}
}

func (v NullableAuthenticationInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


