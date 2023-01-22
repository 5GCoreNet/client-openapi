/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
)

// LocalityDescription Locality description
type LocalityDescription struct {
	LocalityType LocalityType `json:"localityType"`
	LocalityValue string `json:"localityValue"`
	AddlLocDescrItems []LocalityDescriptionItem `json:"addlLocDescrItems,omitempty"`
}

// NewLocalityDescription instantiates a new LocalityDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalityDescription(localityType LocalityType, localityValue string) *LocalityDescription {
	this := LocalityDescription{}
	this.LocalityType = localityType
	this.LocalityValue = localityValue
	return &this
}

// NewLocalityDescriptionWithDefaults instantiates a new LocalityDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalityDescriptionWithDefaults() *LocalityDescription {
	this := LocalityDescription{}
	return &this
}

// GetLocalityType returns the LocalityType field value
func (o *LocalityDescription) GetLocalityType() LocalityType {
	if o == nil {
		var ret LocalityType
		return ret
	}

	return o.LocalityType
}

// GetLocalityTypeOk returns a tuple with the LocalityType field value
// and a boolean to check if the value has been set.
func (o *LocalityDescription) GetLocalityTypeOk() (*LocalityType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocalityType, true
}

// SetLocalityType sets field value
func (o *LocalityDescription) SetLocalityType(v LocalityType) {
	o.LocalityType = v
}

// GetLocalityValue returns the LocalityValue field value
func (o *LocalityDescription) GetLocalityValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalityValue
}

// GetLocalityValueOk returns a tuple with the LocalityValue field value
// and a boolean to check if the value has been set.
func (o *LocalityDescription) GetLocalityValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocalityValue, true
}

// SetLocalityValue sets field value
func (o *LocalityDescription) SetLocalityValue(v string) {
	o.LocalityValue = v
}

// GetAddlLocDescrItems returns the AddlLocDescrItems field value if set, zero value otherwise.
func (o *LocalityDescription) GetAddlLocDescrItems() []LocalityDescriptionItem {
	if o == nil || isNil(o.AddlLocDescrItems) {
		var ret []LocalityDescriptionItem
		return ret
	}
	return o.AddlLocDescrItems
}

// GetAddlLocDescrItemsOk returns a tuple with the AddlLocDescrItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalityDescription) GetAddlLocDescrItemsOk() ([]LocalityDescriptionItem, bool) {
	if o == nil || isNil(o.AddlLocDescrItems) {
    return nil, false
	}
	return o.AddlLocDescrItems, true
}

// HasAddlLocDescrItems returns a boolean if a field has been set.
func (o *LocalityDescription) HasAddlLocDescrItems() bool {
	if o != nil && !isNil(o.AddlLocDescrItems) {
		return true
	}

	return false
}

// SetAddlLocDescrItems gets a reference to the given []LocalityDescriptionItem and assigns it to the AddlLocDescrItems field.
func (o *LocalityDescription) SetAddlLocDescrItems(v []LocalityDescriptionItem) {
	o.AddlLocDescrItems = v
}

func (o LocalityDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["localityType"] = o.LocalityType
	}
	if true {
		toSerialize["localityValue"] = o.LocalityValue
	}
	if !isNil(o.AddlLocDescrItems) {
		toSerialize["addlLocDescrItems"] = o.AddlLocDescrItems
	}
	return json.Marshal(toSerialize)
}

type NullableLocalityDescription struct {
	value *LocalityDescription
	isSet bool
}

func (v NullableLocalityDescription) Get() *LocalityDescription {
	return v.value
}

func (v *NullableLocalityDescription) Set(val *LocalityDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalityDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalityDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalityDescription(val *LocalityDescription) *NullableLocalityDescription {
	return &NullableLocalityDescription{value: val, isSet: true}
}

func (v NullableLocalityDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalityDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


