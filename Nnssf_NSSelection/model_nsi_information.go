/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"encoding/json"
)

// NsiInformation Contains the API URIs of NRF services to be used to discover NFs/services, subscribe to NF status changes and/or request access tokens within the selected Network Slice instance and optional the Identifier of the selected Network Slice instance 
type NsiInformation struct {
	// String providing an URI formatted according to RFC 3986.
	NrfId string `json:"nrfId"`
	// Contains the Identifier of the selected Network Slice instance
	NsiId *string `json:"nsiId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NrfNfMgtUri *string `json:"nrfNfMgtUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NrfAccessTokenUri *string `json:"nrfAccessTokenUri,omitempty"`
	// Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \"nnrf-nfm\" or \"nnrf-disc\" 
	NrfOauth2Required *map[string]bool `json:"nrfOauth2Required,omitempty"`
}

// NewNsiInformation instantiates a new NsiInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsiInformation(nrfId string) *NsiInformation {
	this := NsiInformation{}
	this.NrfId = nrfId
	return &this
}

// NewNsiInformationWithDefaults instantiates a new NsiInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsiInformationWithDefaults() *NsiInformation {
	this := NsiInformation{}
	return &this
}

// GetNrfId returns the NrfId field value
func (o *NsiInformation) GetNrfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NrfId
}

// GetNrfIdOk returns a tuple with the NrfId field value
// and a boolean to check if the value has been set.
func (o *NsiInformation) GetNrfIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NrfId, true
}

// SetNrfId sets field value
func (o *NsiInformation) SetNrfId(v string) {
	o.NrfId = v
}

// GetNsiId returns the NsiId field value if set, zero value otherwise.
func (o *NsiInformation) GetNsiId() string {
	if o == nil || isNil(o.NsiId) {
		var ret string
		return ret
	}
	return *o.NsiId
}

// GetNsiIdOk returns a tuple with the NsiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiInformation) GetNsiIdOk() (*string, bool) {
	if o == nil || isNil(o.NsiId) {
    return nil, false
	}
	return o.NsiId, true
}

// HasNsiId returns a boolean if a field has been set.
func (o *NsiInformation) HasNsiId() bool {
	if o != nil && !isNil(o.NsiId) {
		return true
	}

	return false
}

// SetNsiId gets a reference to the given string and assigns it to the NsiId field.
func (o *NsiInformation) SetNsiId(v string) {
	o.NsiId = &v
}

// GetNrfNfMgtUri returns the NrfNfMgtUri field value if set, zero value otherwise.
func (o *NsiInformation) GetNrfNfMgtUri() string {
	if o == nil || isNil(o.NrfNfMgtUri) {
		var ret string
		return ret
	}
	return *o.NrfNfMgtUri
}

// GetNrfNfMgtUriOk returns a tuple with the NrfNfMgtUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiInformation) GetNrfNfMgtUriOk() (*string, bool) {
	if o == nil || isNil(o.NrfNfMgtUri) {
    return nil, false
	}
	return o.NrfNfMgtUri, true
}

// HasNrfNfMgtUri returns a boolean if a field has been set.
func (o *NsiInformation) HasNrfNfMgtUri() bool {
	if o != nil && !isNil(o.NrfNfMgtUri) {
		return true
	}

	return false
}

// SetNrfNfMgtUri gets a reference to the given string and assigns it to the NrfNfMgtUri field.
func (o *NsiInformation) SetNrfNfMgtUri(v string) {
	o.NrfNfMgtUri = &v
}

// GetNrfAccessTokenUri returns the NrfAccessTokenUri field value if set, zero value otherwise.
func (o *NsiInformation) GetNrfAccessTokenUri() string {
	if o == nil || isNil(o.NrfAccessTokenUri) {
		var ret string
		return ret
	}
	return *o.NrfAccessTokenUri
}

// GetNrfAccessTokenUriOk returns a tuple with the NrfAccessTokenUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiInformation) GetNrfAccessTokenUriOk() (*string, bool) {
	if o == nil || isNil(o.NrfAccessTokenUri) {
    return nil, false
	}
	return o.NrfAccessTokenUri, true
}

// HasNrfAccessTokenUri returns a boolean if a field has been set.
func (o *NsiInformation) HasNrfAccessTokenUri() bool {
	if o != nil && !isNil(o.NrfAccessTokenUri) {
		return true
	}

	return false
}

// SetNrfAccessTokenUri gets a reference to the given string and assigns it to the NrfAccessTokenUri field.
func (o *NsiInformation) SetNrfAccessTokenUri(v string) {
	o.NrfAccessTokenUri = &v
}

// GetNrfOauth2Required returns the NrfOauth2Required field value if set, zero value otherwise.
func (o *NsiInformation) GetNrfOauth2Required() map[string]bool {
	if o == nil || isNil(o.NrfOauth2Required) {
		var ret map[string]bool
		return ret
	}
	return *o.NrfOauth2Required
}

// GetNrfOauth2RequiredOk returns a tuple with the NrfOauth2Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiInformation) GetNrfOauth2RequiredOk() (*map[string]bool, bool) {
	if o == nil || isNil(o.NrfOauth2Required) {
    return nil, false
	}
	return o.NrfOauth2Required, true
}

// HasNrfOauth2Required returns a boolean if a field has been set.
func (o *NsiInformation) HasNrfOauth2Required() bool {
	if o != nil && !isNil(o.NrfOauth2Required) {
		return true
	}

	return false
}

// SetNrfOauth2Required gets a reference to the given map[string]bool and assigns it to the NrfOauth2Required field.
func (o *NsiInformation) SetNrfOauth2Required(v map[string]bool) {
	o.NrfOauth2Required = &v
}

func (o NsiInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nrfId"] = o.NrfId
	}
	if !isNil(o.NsiId) {
		toSerialize["nsiId"] = o.NsiId
	}
	if !isNil(o.NrfNfMgtUri) {
		toSerialize["nrfNfMgtUri"] = o.NrfNfMgtUri
	}
	if !isNil(o.NrfAccessTokenUri) {
		toSerialize["nrfAccessTokenUri"] = o.NrfAccessTokenUri
	}
	if !isNil(o.NrfOauth2Required) {
		toSerialize["nrfOauth2Required"] = o.NrfOauth2Required
	}
	return json.Marshal(toSerialize)
}

type NullableNsiInformation struct {
	value *NsiInformation
	isSet bool
}

func (v NullableNsiInformation) Get() *NsiInformation {
	return v.value
}

func (v *NullableNsiInformation) Set(val *NsiInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNsiInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNsiInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsiInformation(val *NsiInformation) *NullableNsiInformation {
	return &NullableNsiInformation{value: val, isSet: true}
}

func (v NullableNsiInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsiInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


