/*
5GMS Common Data Types

5GMS Common Data Types © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
)

// ChargingSpecification struct for ChargingSpecification
type ChargingSpecification struct {
	SponId *string `json:"sponId,omitempty"`
	SponStatus *SponsoringStatus `json:"sponStatus,omitempty"`
	Gpsi []string `json:"gpsi,omitempty"`
}

// NewChargingSpecification instantiates a new ChargingSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargingSpecification() *ChargingSpecification {
	this := ChargingSpecification{}
	return &this
}

// NewChargingSpecificationWithDefaults instantiates a new ChargingSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargingSpecificationWithDefaults() *ChargingSpecification {
	this := ChargingSpecification{}
	return &this
}

// GetSponId returns the SponId field value if set, zero value otherwise.
func (o *ChargingSpecification) GetSponId() string {
	if o == nil || isNil(o.SponId) {
		var ret string
		return ret
	}
	return *o.SponId
}

// GetSponIdOk returns a tuple with the SponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingSpecification) GetSponIdOk() (*string, bool) {
	if o == nil || isNil(o.SponId) {
    return nil, false
	}
	return o.SponId, true
}

// HasSponId returns a boolean if a field has been set.
func (o *ChargingSpecification) HasSponId() bool {
	if o != nil && !isNil(o.SponId) {
		return true
	}

	return false
}

// SetSponId gets a reference to the given string and assigns it to the SponId field.
func (o *ChargingSpecification) SetSponId(v string) {
	o.SponId = &v
}

// GetSponStatus returns the SponStatus field value if set, zero value otherwise.
func (o *ChargingSpecification) GetSponStatus() SponsoringStatus {
	if o == nil || isNil(o.SponStatus) {
		var ret SponsoringStatus
		return ret
	}
	return *o.SponStatus
}

// GetSponStatusOk returns a tuple with the SponStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingSpecification) GetSponStatusOk() (*SponsoringStatus, bool) {
	if o == nil || isNil(o.SponStatus) {
    return nil, false
	}
	return o.SponStatus, true
}

// HasSponStatus returns a boolean if a field has been set.
func (o *ChargingSpecification) HasSponStatus() bool {
	if o != nil && !isNil(o.SponStatus) {
		return true
	}

	return false
}

// SetSponStatus gets a reference to the given SponsoringStatus and assigns it to the SponStatus field.
func (o *ChargingSpecification) SetSponStatus(v SponsoringStatus) {
	o.SponStatus = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *ChargingSpecification) GetGpsi() []string {
	if o == nil || isNil(o.Gpsi) {
		var ret []string
		return ret
	}
	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingSpecification) GetGpsiOk() ([]string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *ChargingSpecification) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given []string and assigns it to the Gpsi field.
func (o *ChargingSpecification) SetGpsi(v []string) {
	o.Gpsi = v
}

func (o ChargingSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SponId) {
		toSerialize["sponId"] = o.SponId
	}
	if !isNil(o.SponStatus) {
		toSerialize["sponStatus"] = o.SponStatus
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	return json.Marshal(toSerialize)
}

type NullableChargingSpecification struct {
	value *ChargingSpecification
	isSet bool
}

func (v NullableChargingSpecification) Get() *ChargingSpecification {
	return v.value
}

func (v *NullableChargingSpecification) Set(val *ChargingSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingSpecification(val *ChargingSpecification) *NullableChargingSpecification {
	return &NullableChargingSpecification{value: val, isSet: true}
}

func (v NullableChargingSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


