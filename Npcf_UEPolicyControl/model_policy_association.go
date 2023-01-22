/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_UEPolicyControl

import (
	"encoding/json"
)

// PolicyAssociation Contains the description of a policy association that is returned by the PCF when a policy Association is created, updated, or read. 
type PolicyAssociation struct {
	Request *PolicyAssociationRequest `json:"request,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	UePolicy *string `json:"uePolicy,omitempty"`
	N2Pc5Pol *N2InfoContent `json:"n2Pc5Pol,omitempty"`
	N2Pc5ProSePol *N2InfoContent `json:"n2Pc5ProSePol,omitempty"`
	// Request Triggers that the PCF subscribes. Only values \"LOC_CH\" and \"PRA_CH\" are permitted. 
	Triggers []RequestTrigger `json:"triggers,omitempty"`
	// Contains the presence reporting area(s) for which reporting was requested. The praId attribute within the PresenceInfo data type is the key of the map. 
	Pras *map[string]PresenceInfo `json:"pras,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewPolicyAssociation instantiates a new PolicyAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAssociation(suppFeat string) *PolicyAssociation {
	this := PolicyAssociation{}
	this.SuppFeat = suppFeat
	return &this
}

// NewPolicyAssociationWithDefaults instantiates a new PolicyAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAssociationWithDefaults() *PolicyAssociation {
	this := PolicyAssociation{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *PolicyAssociation) GetRequest() PolicyAssociationRequest {
	if o == nil || isNil(o.Request) {
		var ret PolicyAssociationRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociation) GetRequestOk() (*PolicyAssociationRequest, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *PolicyAssociation) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given PolicyAssociationRequest and assigns it to the Request field.
func (o *PolicyAssociation) SetRequest(v PolicyAssociationRequest) {
	o.Request = &v
}

// GetUePolicy returns the UePolicy field value if set, zero value otherwise.
func (o *PolicyAssociation) GetUePolicy() string {
	if o == nil || isNil(o.UePolicy) {
		var ret string
		return ret
	}
	return *o.UePolicy
}

// GetUePolicyOk returns a tuple with the UePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociation) GetUePolicyOk() (*string, bool) {
	if o == nil || isNil(o.UePolicy) {
    return nil, false
	}
	return o.UePolicy, true
}

// HasUePolicy returns a boolean if a field has been set.
func (o *PolicyAssociation) HasUePolicy() bool {
	if o != nil && !isNil(o.UePolicy) {
		return true
	}

	return false
}

// SetUePolicy gets a reference to the given string and assigns it to the UePolicy field.
func (o *PolicyAssociation) SetUePolicy(v string) {
	o.UePolicy = &v
}

// GetN2Pc5Pol returns the N2Pc5Pol field value if set, zero value otherwise.
func (o *PolicyAssociation) GetN2Pc5Pol() N2InfoContent {
	if o == nil || isNil(o.N2Pc5Pol) {
		var ret N2InfoContent
		return ret
	}
	return *o.N2Pc5Pol
}

// GetN2Pc5PolOk returns a tuple with the N2Pc5Pol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociation) GetN2Pc5PolOk() (*N2InfoContent, bool) {
	if o == nil || isNil(o.N2Pc5Pol) {
    return nil, false
	}
	return o.N2Pc5Pol, true
}

// HasN2Pc5Pol returns a boolean if a field has been set.
func (o *PolicyAssociation) HasN2Pc5Pol() bool {
	if o != nil && !isNil(o.N2Pc5Pol) {
		return true
	}

	return false
}

// SetN2Pc5Pol gets a reference to the given N2InfoContent and assigns it to the N2Pc5Pol field.
func (o *PolicyAssociation) SetN2Pc5Pol(v N2InfoContent) {
	o.N2Pc5Pol = &v
}

// GetN2Pc5ProSePol returns the N2Pc5ProSePol field value if set, zero value otherwise.
func (o *PolicyAssociation) GetN2Pc5ProSePol() N2InfoContent {
	if o == nil || isNil(o.N2Pc5ProSePol) {
		var ret N2InfoContent
		return ret
	}
	return *o.N2Pc5ProSePol
}

// GetN2Pc5ProSePolOk returns a tuple with the N2Pc5ProSePol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociation) GetN2Pc5ProSePolOk() (*N2InfoContent, bool) {
	if o == nil || isNil(o.N2Pc5ProSePol) {
    return nil, false
	}
	return o.N2Pc5ProSePol, true
}

// HasN2Pc5ProSePol returns a boolean if a field has been set.
func (o *PolicyAssociation) HasN2Pc5ProSePol() bool {
	if o != nil && !isNil(o.N2Pc5ProSePol) {
		return true
	}

	return false
}

// SetN2Pc5ProSePol gets a reference to the given N2InfoContent and assigns it to the N2Pc5ProSePol field.
func (o *PolicyAssociation) SetN2Pc5ProSePol(v N2InfoContent) {
	o.N2Pc5ProSePol = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *PolicyAssociation) GetTriggers() []RequestTrigger {
	if o == nil || isNil(o.Triggers) {
		var ret []RequestTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociation) GetTriggersOk() ([]RequestTrigger, bool) {
	if o == nil || isNil(o.Triggers) {
    return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *PolicyAssociation) HasTriggers() bool {
	if o != nil && !isNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []RequestTrigger and assigns it to the Triggers field.
func (o *PolicyAssociation) SetTriggers(v []RequestTrigger) {
	o.Triggers = v
}

// GetPras returns the Pras field value if set, zero value otherwise.
func (o *PolicyAssociation) GetPras() map[string]PresenceInfo {
	if o == nil || isNil(o.Pras) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.Pras
}

// GetPrasOk returns a tuple with the Pras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociation) GetPrasOk() (*map[string]PresenceInfo, bool) {
	if o == nil || isNil(o.Pras) {
    return nil, false
	}
	return o.Pras, true
}

// HasPras returns a boolean if a field has been set.
func (o *PolicyAssociation) HasPras() bool {
	if o != nil && !isNil(o.Pras) {
		return true
	}

	return false
}

// SetPras gets a reference to the given map[string]PresenceInfo and assigns it to the Pras field.
func (o *PolicyAssociation) SetPras(v map[string]PresenceInfo) {
	o.Pras = &v
}

// GetSuppFeat returns the SuppFeat field value
func (o *PolicyAssociation) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *PolicyAssociation) GetSuppFeatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *PolicyAssociation) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o PolicyAssociation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.UePolicy) {
		toSerialize["uePolicy"] = o.UePolicy
	}
	if !isNil(o.N2Pc5Pol) {
		toSerialize["n2Pc5Pol"] = o.N2Pc5Pol
	}
	if !isNil(o.N2Pc5ProSePol) {
		toSerialize["n2Pc5ProSePol"] = o.N2Pc5ProSePol
	}
	if !isNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !isNil(o.Pras) {
		toSerialize["pras"] = o.Pras
	}
	if true {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyAssociation struct {
	value *PolicyAssociation
	isSet bool
}

func (v NullablePolicyAssociation) Get() *PolicyAssociation {
	return v.value
}

func (v *NullablePolicyAssociation) Set(val *PolicyAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAssociation(val *PolicyAssociation) *NullablePolicyAssociation {
	return &NullablePolicyAssociation{value: val, isSet: true}
}

func (v NullablePolicyAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


