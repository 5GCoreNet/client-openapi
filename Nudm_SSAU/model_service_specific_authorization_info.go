/*
Nudm_SSAU

Nudm Service Specific Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SSAU

import (
	"encoding/json"
)

// ServiceSpecificAuthorizationInfo Authorization information for a specific service
type ServiceSpecificAuthorizationInfo struct {
	Snssai *Snssai `json:"snssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AuthUpdateCallbackUri *string `json:"authUpdateCallbackUri,omitempty"`
	AfId *string `json:"afId,omitempty"`
	// Identity of the NEF
	NefId *string `json:"nefId,omitempty"`
}

// NewServiceSpecificAuthorizationInfo instantiates a new ServiceSpecificAuthorizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSpecificAuthorizationInfo() *ServiceSpecificAuthorizationInfo {
	this := ServiceSpecificAuthorizationInfo{}
	return &this
}

// NewServiceSpecificAuthorizationInfoWithDefaults instantiates a new ServiceSpecificAuthorizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSpecificAuthorizationInfoWithDefaults() *ServiceSpecificAuthorizationInfo {
	this := ServiceSpecificAuthorizationInfo{}
	return &this
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationInfo) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationInfo) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *ServiceSpecificAuthorizationInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationInfo) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationInfo) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationInfo) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *ServiceSpecificAuthorizationInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationInfo) GetMtcProviderInformation() string {
	if o == nil || isNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationInfo) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || isNil(o.MtcProviderInformation) {
    return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationInfo) HasMtcProviderInformation() bool {
	if o != nil && !isNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *ServiceSpecificAuthorizationInfo) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

// GetAuthUpdateCallbackUri returns the AuthUpdateCallbackUri field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationInfo) GetAuthUpdateCallbackUri() string {
	if o == nil || isNil(o.AuthUpdateCallbackUri) {
		var ret string
		return ret
	}
	return *o.AuthUpdateCallbackUri
}

// GetAuthUpdateCallbackUriOk returns a tuple with the AuthUpdateCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationInfo) GetAuthUpdateCallbackUriOk() (*string, bool) {
	if o == nil || isNil(o.AuthUpdateCallbackUri) {
    return nil, false
	}
	return o.AuthUpdateCallbackUri, true
}

// HasAuthUpdateCallbackUri returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationInfo) HasAuthUpdateCallbackUri() bool {
	if o != nil && !isNil(o.AuthUpdateCallbackUri) {
		return true
	}

	return false
}

// SetAuthUpdateCallbackUri gets a reference to the given string and assigns it to the AuthUpdateCallbackUri field.
func (o *ServiceSpecificAuthorizationInfo) SetAuthUpdateCallbackUri(v string) {
	o.AuthUpdateCallbackUri = &v
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationInfo) GetAfId() string {
	if o == nil || isNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationInfo) GetAfIdOk() (*string, bool) {
	if o == nil || isNil(o.AfId) {
    return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationInfo) HasAfId() bool {
	if o != nil && !isNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *ServiceSpecificAuthorizationInfo) SetAfId(v string) {
	o.AfId = &v
}

// GetNefId returns the NefId field value if set, zero value otherwise.
func (o *ServiceSpecificAuthorizationInfo) GetNefId() string {
	if o == nil || isNil(o.NefId) {
		var ret string
		return ret
	}
	return *o.NefId
}

// GetNefIdOk returns a tuple with the NefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationInfo) GetNefIdOk() (*string, bool) {
	if o == nil || isNil(o.NefId) {
    return nil, false
	}
	return o.NefId, true
}

// HasNefId returns a boolean if a field has been set.
func (o *ServiceSpecificAuthorizationInfo) HasNefId() bool {
	if o != nil && !isNil(o.NefId) {
		return true
	}

	return false
}

// SetNefId gets a reference to the given string and assigns it to the NefId field.
func (o *ServiceSpecificAuthorizationInfo) SetNefId(v string) {
	o.NefId = &v
}

func (o ServiceSpecificAuthorizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	if !isNil(o.AuthUpdateCallbackUri) {
		toSerialize["authUpdateCallbackUri"] = o.AuthUpdateCallbackUri
	}
	if !isNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	if !isNil(o.NefId) {
		toSerialize["nefId"] = o.NefId
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSpecificAuthorizationInfo struct {
	value *ServiceSpecificAuthorizationInfo
	isSet bool
}

func (v NullableServiceSpecificAuthorizationInfo) Get() *ServiceSpecificAuthorizationInfo {
	return v.value
}

func (v *NullableServiceSpecificAuthorizationInfo) Set(val *ServiceSpecificAuthorizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSpecificAuthorizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSpecificAuthorizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSpecificAuthorizationInfo(val *ServiceSpecificAuthorizationInfo) *NullableServiceSpecificAuthorizationInfo {
	return &NullableServiceSpecificAuthorizationInfo{value: val, isSet: true}
}

func (v NullableServiceSpecificAuthorizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSpecificAuthorizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


