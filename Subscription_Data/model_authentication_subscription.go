/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// AuthenticationSubscription A UE's authentication data.
type AuthenticationSubscription struct {
	AuthenticationMethod AuthMethod `json:"authenticationMethod"`
	EncPermanentKey *string `json:"encPermanentKey,omitempty"`
	ProtectionParameterId *string `json:"protectionParameterId,omitempty"`
	SequenceNumber *SequenceNumber `json:"sequenceNumber,omitempty"`
	AuthenticationManagementField *string `json:"authenticationManagementField,omitempty"`
	AlgorithmId *string `json:"algorithmId,omitempty"`
	EncOpcKey *string `json:"encOpcKey,omitempty"`
	EncTopcKey *string `json:"encTopcKey,omitempty"`
	VectorGenerationInHss *bool `json:"vectorGenerationInHss,omitempty"`
	// Identifier of a group of NFs.
	HssGroupId *string `json:"hssGroupId,omitempty"`
	N5gcAuthMethod *AuthMethod `json:"n5gcAuthMethod,omitempty"`
	RgAuthenticationInd *bool `json:"rgAuthenticationInd,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	AkmaAllowed *bool `json:"akmaAllowed,omitempty"`
	RoutingId *string `json:"routingId,omitempty"`
}

// NewAuthenticationSubscription instantiates a new AuthenticationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationSubscription(authenticationMethod AuthMethod) *AuthenticationSubscription {
	this := AuthenticationSubscription{}
	this.AuthenticationMethod = authenticationMethod
	var vectorGenerationInHss bool = false
	this.VectorGenerationInHss = &vectorGenerationInHss
	var rgAuthenticationInd bool = false
	this.RgAuthenticationInd = &rgAuthenticationInd
	var akmaAllowed bool = false
	this.AkmaAllowed = &akmaAllowed
	return &this
}

// NewAuthenticationSubscriptionWithDefaults instantiates a new AuthenticationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationSubscriptionWithDefaults() *AuthenticationSubscription {
	this := AuthenticationSubscription{}
	var vectorGenerationInHss bool = false
	this.VectorGenerationInHss = &vectorGenerationInHss
	var rgAuthenticationInd bool = false
	this.RgAuthenticationInd = &rgAuthenticationInd
	var akmaAllowed bool = false
	this.AkmaAllowed = &akmaAllowed
	return &this
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *AuthenticationSubscription) GetAuthenticationMethod() AuthMethod {
	if o == nil {
		var ret AuthMethod
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetAuthenticationMethodOk() (*AuthMethod, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *AuthenticationSubscription) SetAuthenticationMethod(v AuthMethod) {
	o.AuthenticationMethod = v
}

// GetEncPermanentKey returns the EncPermanentKey field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetEncPermanentKey() string {
	if o == nil || isNil(o.EncPermanentKey) {
		var ret string
		return ret
	}
	return *o.EncPermanentKey
}

// GetEncPermanentKeyOk returns a tuple with the EncPermanentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetEncPermanentKeyOk() (*string, bool) {
	if o == nil || isNil(o.EncPermanentKey) {
    return nil, false
	}
	return o.EncPermanentKey, true
}

// HasEncPermanentKey returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasEncPermanentKey() bool {
	if o != nil && !isNil(o.EncPermanentKey) {
		return true
	}

	return false
}

// SetEncPermanentKey gets a reference to the given string and assigns it to the EncPermanentKey field.
func (o *AuthenticationSubscription) SetEncPermanentKey(v string) {
	o.EncPermanentKey = &v
}

// GetProtectionParameterId returns the ProtectionParameterId field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetProtectionParameterId() string {
	if o == nil || isNil(o.ProtectionParameterId) {
		var ret string
		return ret
	}
	return *o.ProtectionParameterId
}

// GetProtectionParameterIdOk returns a tuple with the ProtectionParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetProtectionParameterIdOk() (*string, bool) {
	if o == nil || isNil(o.ProtectionParameterId) {
    return nil, false
	}
	return o.ProtectionParameterId, true
}

// HasProtectionParameterId returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasProtectionParameterId() bool {
	if o != nil && !isNil(o.ProtectionParameterId) {
		return true
	}

	return false
}

// SetProtectionParameterId gets a reference to the given string and assigns it to the ProtectionParameterId field.
func (o *AuthenticationSubscription) SetProtectionParameterId(v string) {
	o.ProtectionParameterId = &v
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetSequenceNumber() SequenceNumber {
	if o == nil || isNil(o.SequenceNumber) {
		var ret SequenceNumber
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetSequenceNumberOk() (*SequenceNumber, bool) {
	if o == nil || isNil(o.SequenceNumber) {
    return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasSequenceNumber() bool {
	if o != nil && !isNil(o.SequenceNumber) {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given SequenceNumber and assigns it to the SequenceNumber field.
func (o *AuthenticationSubscription) SetSequenceNumber(v SequenceNumber) {
	o.SequenceNumber = &v
}

// GetAuthenticationManagementField returns the AuthenticationManagementField field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetAuthenticationManagementField() string {
	if o == nil || isNil(o.AuthenticationManagementField) {
		var ret string
		return ret
	}
	return *o.AuthenticationManagementField
}

// GetAuthenticationManagementFieldOk returns a tuple with the AuthenticationManagementField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetAuthenticationManagementFieldOk() (*string, bool) {
	if o == nil || isNil(o.AuthenticationManagementField) {
    return nil, false
	}
	return o.AuthenticationManagementField, true
}

// HasAuthenticationManagementField returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasAuthenticationManagementField() bool {
	if o != nil && !isNil(o.AuthenticationManagementField) {
		return true
	}

	return false
}

// SetAuthenticationManagementField gets a reference to the given string and assigns it to the AuthenticationManagementField field.
func (o *AuthenticationSubscription) SetAuthenticationManagementField(v string) {
	o.AuthenticationManagementField = &v
}

// GetAlgorithmId returns the AlgorithmId field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetAlgorithmId() string {
	if o == nil || isNil(o.AlgorithmId) {
		var ret string
		return ret
	}
	return *o.AlgorithmId
}

// GetAlgorithmIdOk returns a tuple with the AlgorithmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetAlgorithmIdOk() (*string, bool) {
	if o == nil || isNil(o.AlgorithmId) {
    return nil, false
	}
	return o.AlgorithmId, true
}

// HasAlgorithmId returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasAlgorithmId() bool {
	if o != nil && !isNil(o.AlgorithmId) {
		return true
	}

	return false
}

// SetAlgorithmId gets a reference to the given string and assigns it to the AlgorithmId field.
func (o *AuthenticationSubscription) SetAlgorithmId(v string) {
	o.AlgorithmId = &v
}

// GetEncOpcKey returns the EncOpcKey field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetEncOpcKey() string {
	if o == nil || isNil(o.EncOpcKey) {
		var ret string
		return ret
	}
	return *o.EncOpcKey
}

// GetEncOpcKeyOk returns a tuple with the EncOpcKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetEncOpcKeyOk() (*string, bool) {
	if o == nil || isNil(o.EncOpcKey) {
    return nil, false
	}
	return o.EncOpcKey, true
}

// HasEncOpcKey returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasEncOpcKey() bool {
	if o != nil && !isNil(o.EncOpcKey) {
		return true
	}

	return false
}

// SetEncOpcKey gets a reference to the given string and assigns it to the EncOpcKey field.
func (o *AuthenticationSubscription) SetEncOpcKey(v string) {
	o.EncOpcKey = &v
}

// GetEncTopcKey returns the EncTopcKey field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetEncTopcKey() string {
	if o == nil || isNil(o.EncTopcKey) {
		var ret string
		return ret
	}
	return *o.EncTopcKey
}

// GetEncTopcKeyOk returns a tuple with the EncTopcKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetEncTopcKeyOk() (*string, bool) {
	if o == nil || isNil(o.EncTopcKey) {
    return nil, false
	}
	return o.EncTopcKey, true
}

// HasEncTopcKey returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasEncTopcKey() bool {
	if o != nil && !isNil(o.EncTopcKey) {
		return true
	}

	return false
}

// SetEncTopcKey gets a reference to the given string and assigns it to the EncTopcKey field.
func (o *AuthenticationSubscription) SetEncTopcKey(v string) {
	o.EncTopcKey = &v
}

// GetVectorGenerationInHss returns the VectorGenerationInHss field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetVectorGenerationInHss() bool {
	if o == nil || isNil(o.VectorGenerationInHss) {
		var ret bool
		return ret
	}
	return *o.VectorGenerationInHss
}

// GetVectorGenerationInHssOk returns a tuple with the VectorGenerationInHss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetVectorGenerationInHssOk() (*bool, bool) {
	if o == nil || isNil(o.VectorGenerationInHss) {
    return nil, false
	}
	return o.VectorGenerationInHss, true
}

// HasVectorGenerationInHss returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasVectorGenerationInHss() bool {
	if o != nil && !isNil(o.VectorGenerationInHss) {
		return true
	}

	return false
}

// SetVectorGenerationInHss gets a reference to the given bool and assigns it to the VectorGenerationInHss field.
func (o *AuthenticationSubscription) SetVectorGenerationInHss(v bool) {
	o.VectorGenerationInHss = &v
}

// GetHssGroupId returns the HssGroupId field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetHssGroupId() string {
	if o == nil || isNil(o.HssGroupId) {
		var ret string
		return ret
	}
	return *o.HssGroupId
}

// GetHssGroupIdOk returns a tuple with the HssGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetHssGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.HssGroupId) {
    return nil, false
	}
	return o.HssGroupId, true
}

// HasHssGroupId returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasHssGroupId() bool {
	if o != nil && !isNil(o.HssGroupId) {
		return true
	}

	return false
}

// SetHssGroupId gets a reference to the given string and assigns it to the HssGroupId field.
func (o *AuthenticationSubscription) SetHssGroupId(v string) {
	o.HssGroupId = &v
}

// GetN5gcAuthMethod returns the N5gcAuthMethod field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetN5gcAuthMethod() AuthMethod {
	if o == nil || isNil(o.N5gcAuthMethod) {
		var ret AuthMethod
		return ret
	}
	return *o.N5gcAuthMethod
}

// GetN5gcAuthMethodOk returns a tuple with the N5gcAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetN5gcAuthMethodOk() (*AuthMethod, bool) {
	if o == nil || isNil(o.N5gcAuthMethod) {
    return nil, false
	}
	return o.N5gcAuthMethod, true
}

// HasN5gcAuthMethod returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasN5gcAuthMethod() bool {
	if o != nil && !isNil(o.N5gcAuthMethod) {
		return true
	}

	return false
}

// SetN5gcAuthMethod gets a reference to the given AuthMethod and assigns it to the N5gcAuthMethod field.
func (o *AuthenticationSubscription) SetN5gcAuthMethod(v AuthMethod) {
	o.N5gcAuthMethod = &v
}

// GetRgAuthenticationInd returns the RgAuthenticationInd field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetRgAuthenticationInd() bool {
	if o == nil || isNil(o.RgAuthenticationInd) {
		var ret bool
		return ret
	}
	return *o.RgAuthenticationInd
}

// GetRgAuthenticationIndOk returns a tuple with the RgAuthenticationInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetRgAuthenticationIndOk() (*bool, bool) {
	if o == nil || isNil(o.RgAuthenticationInd) {
    return nil, false
	}
	return o.RgAuthenticationInd, true
}

// HasRgAuthenticationInd returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasRgAuthenticationInd() bool {
	if o != nil && !isNil(o.RgAuthenticationInd) {
		return true
	}

	return false
}

// SetRgAuthenticationInd gets a reference to the given bool and assigns it to the RgAuthenticationInd field.
func (o *AuthenticationSubscription) SetRgAuthenticationInd(v bool) {
	o.RgAuthenticationInd = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *AuthenticationSubscription) SetSupi(v string) {
	o.Supi = &v
}

// GetAkmaAllowed returns the AkmaAllowed field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetAkmaAllowed() bool {
	if o == nil || isNil(o.AkmaAllowed) {
		var ret bool
		return ret
	}
	return *o.AkmaAllowed
}

// GetAkmaAllowedOk returns a tuple with the AkmaAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetAkmaAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.AkmaAllowed) {
    return nil, false
	}
	return o.AkmaAllowed, true
}

// HasAkmaAllowed returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasAkmaAllowed() bool {
	if o != nil && !isNil(o.AkmaAllowed) {
		return true
	}

	return false
}

// SetAkmaAllowed gets a reference to the given bool and assigns it to the AkmaAllowed field.
func (o *AuthenticationSubscription) SetAkmaAllowed(v bool) {
	o.AkmaAllowed = &v
}

// GetRoutingId returns the RoutingId field value if set, zero value otherwise.
func (o *AuthenticationSubscription) GetRoutingId() string {
	if o == nil || isNil(o.RoutingId) {
		var ret string
		return ret
	}
	return *o.RoutingId
}

// GetRoutingIdOk returns a tuple with the RoutingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSubscription) GetRoutingIdOk() (*string, bool) {
	if o == nil || isNil(o.RoutingId) {
    return nil, false
	}
	return o.RoutingId, true
}

// HasRoutingId returns a boolean if a field has been set.
func (o *AuthenticationSubscription) HasRoutingId() bool {
	if o != nil && !isNil(o.RoutingId) {
		return true
	}

	return false
}

// SetRoutingId gets a reference to the given string and assigns it to the RoutingId field.
func (o *AuthenticationSubscription) SetRoutingId(v string) {
	o.RoutingId = &v
}

func (o AuthenticationSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	if !isNil(o.EncPermanentKey) {
		toSerialize["encPermanentKey"] = o.EncPermanentKey
	}
	if !isNil(o.ProtectionParameterId) {
		toSerialize["protectionParameterId"] = o.ProtectionParameterId
	}
	if !isNil(o.SequenceNumber) {
		toSerialize["sequenceNumber"] = o.SequenceNumber
	}
	if !isNil(o.AuthenticationManagementField) {
		toSerialize["authenticationManagementField"] = o.AuthenticationManagementField
	}
	if !isNil(o.AlgorithmId) {
		toSerialize["algorithmId"] = o.AlgorithmId
	}
	if !isNil(o.EncOpcKey) {
		toSerialize["encOpcKey"] = o.EncOpcKey
	}
	if !isNil(o.EncTopcKey) {
		toSerialize["encTopcKey"] = o.EncTopcKey
	}
	if !isNil(o.VectorGenerationInHss) {
		toSerialize["vectorGenerationInHss"] = o.VectorGenerationInHss
	}
	if !isNil(o.HssGroupId) {
		toSerialize["hssGroupId"] = o.HssGroupId
	}
	if !isNil(o.N5gcAuthMethod) {
		toSerialize["n5gcAuthMethod"] = o.N5gcAuthMethod
	}
	if !isNil(o.RgAuthenticationInd) {
		toSerialize["rgAuthenticationInd"] = o.RgAuthenticationInd
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.AkmaAllowed) {
		toSerialize["akmaAllowed"] = o.AkmaAllowed
	}
	if !isNil(o.RoutingId) {
		toSerialize["routingId"] = o.RoutingId
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationSubscription struct {
	value *AuthenticationSubscription
	isSet bool
}

func (v NullableAuthenticationSubscription) Get() *AuthenticationSubscription {
	return v.value
}

func (v *NullableAuthenticationSubscription) Set(val *AuthenticationSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationSubscription(val *AuthenticationSubscription) *NullableAuthenticationSubscription {
	return &NullableAuthenticationSubscription{value: val, isSet: true}
}

func (v NullableAuthenticationSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


