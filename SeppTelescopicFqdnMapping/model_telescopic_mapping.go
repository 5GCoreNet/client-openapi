/*
SEPP Telescopic FQDN Mapping API

SEPP Telescopic FQDN Mapping Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SeppTelescopicFqdnMapping

import (
	"encoding/json"
)

// TelescopicMapping Contains the Telescopic mapping data
type TelescopicMapping struct {
	TelescopicLabel *string `json:"telescopicLabel,omitempty"`
	// Fully Qualified Domain Name
	SeppDomain *string `json:"seppDomain,omitempty"`
	// Fully Qualified Domain Name
	ForeignFqdn *string `json:"foreignFqdn,omitempty"`
}

// NewTelescopicMapping instantiates a new TelescopicMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelescopicMapping() *TelescopicMapping {
	this := TelescopicMapping{}
	return &this
}

// NewTelescopicMappingWithDefaults instantiates a new TelescopicMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelescopicMappingWithDefaults() *TelescopicMapping {
	this := TelescopicMapping{}
	return &this
}

// GetTelescopicLabel returns the TelescopicLabel field value if set, zero value otherwise.
func (o *TelescopicMapping) GetTelescopicLabel() string {
	if o == nil || isNil(o.TelescopicLabel) {
		var ret string
		return ret
	}
	return *o.TelescopicLabel
}

// GetTelescopicLabelOk returns a tuple with the TelescopicLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelescopicMapping) GetTelescopicLabelOk() (*string, bool) {
	if o == nil || isNil(o.TelescopicLabel) {
    return nil, false
	}
	return o.TelescopicLabel, true
}

// HasTelescopicLabel returns a boolean if a field has been set.
func (o *TelescopicMapping) HasTelescopicLabel() bool {
	if o != nil && !isNil(o.TelescopicLabel) {
		return true
	}

	return false
}

// SetTelescopicLabel gets a reference to the given string and assigns it to the TelescopicLabel field.
func (o *TelescopicMapping) SetTelescopicLabel(v string) {
	o.TelescopicLabel = &v
}

// GetSeppDomain returns the SeppDomain field value if set, zero value otherwise.
func (o *TelescopicMapping) GetSeppDomain() string {
	if o == nil || isNil(o.SeppDomain) {
		var ret string
		return ret
	}
	return *o.SeppDomain
}

// GetSeppDomainOk returns a tuple with the SeppDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelescopicMapping) GetSeppDomainOk() (*string, bool) {
	if o == nil || isNil(o.SeppDomain) {
    return nil, false
	}
	return o.SeppDomain, true
}

// HasSeppDomain returns a boolean if a field has been set.
func (o *TelescopicMapping) HasSeppDomain() bool {
	if o != nil && !isNil(o.SeppDomain) {
		return true
	}

	return false
}

// SetSeppDomain gets a reference to the given string and assigns it to the SeppDomain field.
func (o *TelescopicMapping) SetSeppDomain(v string) {
	o.SeppDomain = &v
}

// GetForeignFqdn returns the ForeignFqdn field value if set, zero value otherwise.
func (o *TelescopicMapping) GetForeignFqdn() string {
	if o == nil || isNil(o.ForeignFqdn) {
		var ret string
		return ret
	}
	return *o.ForeignFqdn
}

// GetForeignFqdnOk returns a tuple with the ForeignFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelescopicMapping) GetForeignFqdnOk() (*string, bool) {
	if o == nil || isNil(o.ForeignFqdn) {
    return nil, false
	}
	return o.ForeignFqdn, true
}

// HasForeignFqdn returns a boolean if a field has been set.
func (o *TelescopicMapping) HasForeignFqdn() bool {
	if o != nil && !isNil(o.ForeignFqdn) {
		return true
	}

	return false
}

// SetForeignFqdn gets a reference to the given string and assigns it to the ForeignFqdn field.
func (o *TelescopicMapping) SetForeignFqdn(v string) {
	o.ForeignFqdn = &v
}

func (o TelescopicMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TelescopicLabel) {
		toSerialize["telescopicLabel"] = o.TelescopicLabel
	}
	if !isNil(o.SeppDomain) {
		toSerialize["seppDomain"] = o.SeppDomain
	}
	if !isNil(o.ForeignFqdn) {
		toSerialize["foreignFqdn"] = o.ForeignFqdn
	}
	return json.Marshal(toSerialize)
}

type NullableTelescopicMapping struct {
	value *TelescopicMapping
	isSet bool
}

func (v NullableTelescopicMapping) Get() *TelescopicMapping {
	return v.value
}

func (v *NullableTelescopicMapping) Set(val *TelescopicMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableTelescopicMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableTelescopicMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelescopicMapping(val *TelescopicMapping) *NullableTelescopicMapping {
	return &NullableTelescopicMapping{value: val, isSet: true}
}

func (v NullableTelescopicMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelescopicMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


