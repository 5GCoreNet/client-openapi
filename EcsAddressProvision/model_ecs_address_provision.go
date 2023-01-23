/*
3gpp-ecs-address-provision

API for ECS Address Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EcsAddressProvision

import (
	"encoding/json"
)

// EcsAddressProvision Represents ECS address provision configuration.
type EcsAddressProvision struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	EcsServerAddr EcsServerAddr `json:"ecsServerAddr"`
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
	TgtUe *TargetUeId `json:"tgtUe,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewEcsAddressProvision instantiates a new EcsAddressProvision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcsAddressProvision(ecsServerAddr EcsServerAddr, suppFeat string) *EcsAddressProvision {
	this := EcsAddressProvision{}
	this.EcsServerAddr = ecsServerAddr
	this.SuppFeat = suppFeat
	return &this
}

// NewEcsAddressProvisionWithDefaults instantiates a new EcsAddressProvision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcsAddressProvisionWithDefaults() *EcsAddressProvision {
	this := EcsAddressProvision{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EcsAddressProvision) GetSelf() string {
	if o == nil || isNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsAddressProvision) GetSelfOk() (*string, bool) {
	if o == nil || isNil(o.Self) {
    return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EcsAddressProvision) HasSelf() bool {
	if o != nil && !isNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *EcsAddressProvision) SetSelf(v string) {
	o.Self = &v
}

// GetEcsServerAddr returns the EcsServerAddr field value
func (o *EcsAddressProvision) GetEcsServerAddr() EcsServerAddr {
	if o == nil {
		var ret EcsServerAddr
		return ret
	}

	return o.EcsServerAddr
}

// GetEcsServerAddrOk returns a tuple with the EcsServerAddr field value
// and a boolean to check if the value has been set.
func (o *EcsAddressProvision) GetEcsServerAddrOk() (*EcsServerAddr, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EcsServerAddr, true
}

// SetEcsServerAddr sets field value
func (o *EcsAddressProvision) SetEcsServerAddr(v EcsServerAddr) {
	o.EcsServerAddr = v
}

// GetSpatialValidityCond returns the SpatialValidityCond field value if set, zero value otherwise.
func (o *EcsAddressProvision) GetSpatialValidityCond() SpatialValidityCond {
	if o == nil || isNil(o.SpatialValidityCond) {
		var ret SpatialValidityCond
		return ret
	}
	return *o.SpatialValidityCond
}

// GetSpatialValidityCondOk returns a tuple with the SpatialValidityCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsAddressProvision) GetSpatialValidityCondOk() (*SpatialValidityCond, bool) {
	if o == nil || isNil(o.SpatialValidityCond) {
    return nil, false
	}
	return o.SpatialValidityCond, true
}

// HasSpatialValidityCond returns a boolean if a field has been set.
func (o *EcsAddressProvision) HasSpatialValidityCond() bool {
	if o != nil && !isNil(o.SpatialValidityCond) {
		return true
	}

	return false
}

// SetSpatialValidityCond gets a reference to the given SpatialValidityCond and assigns it to the SpatialValidityCond field.
func (o *EcsAddressProvision) SetSpatialValidityCond(v SpatialValidityCond) {
	o.SpatialValidityCond = &v
}

// GetTgtUe returns the TgtUe field value if set, zero value otherwise.
func (o *EcsAddressProvision) GetTgtUe() TargetUeId {
	if o == nil || isNil(o.TgtUe) {
		var ret TargetUeId
		return ret
	}
	return *o.TgtUe
}

// GetTgtUeOk returns a tuple with the TgtUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsAddressProvision) GetTgtUeOk() (*TargetUeId, bool) {
	if o == nil || isNil(o.TgtUe) {
    return nil, false
	}
	return o.TgtUe, true
}

// HasTgtUe returns a boolean if a field has been set.
func (o *EcsAddressProvision) HasTgtUe() bool {
	if o != nil && !isNil(o.TgtUe) {
		return true
	}

	return false
}

// SetTgtUe gets a reference to the given TargetUeId and assigns it to the TgtUe field.
func (o *EcsAddressProvision) SetTgtUe(v TargetUeId) {
	o.TgtUe = &v
}

// GetSuppFeat returns the SuppFeat field value
func (o *EcsAddressProvision) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *EcsAddressProvision) GetSuppFeatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *EcsAddressProvision) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o EcsAddressProvision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if true {
		toSerialize["ecsServerAddr"] = o.EcsServerAddr
	}
	if !isNil(o.SpatialValidityCond) {
		toSerialize["spatialValidityCond"] = o.SpatialValidityCond
	}
	if !isNil(o.TgtUe) {
		toSerialize["tgtUe"] = o.TgtUe
	}
	if true {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableEcsAddressProvision struct {
	value *EcsAddressProvision
	isSet bool
}

func (v NullableEcsAddressProvision) Get() *EcsAddressProvision {
	return v.value
}

func (v *NullableEcsAddressProvision) Set(val *EcsAddressProvision) {
	v.value = val
	v.isSet = true
}

func (v NullableEcsAddressProvision) IsSet() bool {
	return v.isSet
}

func (v *NullableEcsAddressProvision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcsAddressProvision(val *EcsAddressProvision) *NullableEcsAddressProvision {
	return &NullableEcsAddressProvision{value: val, isSet: true}
}

func (v NullableEcsAddressProvision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcsAddressProvision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


