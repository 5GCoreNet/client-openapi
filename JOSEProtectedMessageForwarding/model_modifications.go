/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package JOSEProtectedMessageForwarding

import (
	"encoding/json"
)

// Modifications Information on inserting of the modifications entry
type Modifications struct {
	// Fully Qualified Domain Name
	Identity string `json:"identity"`
	Operations []PatchItem `json:"operations,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// NewModifications instantiates a new Modifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifications(identity string) *Modifications {
	this := Modifications{}
	this.Identity = identity
	return &this
}

// NewModificationsWithDefaults instantiates a new Modifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModificationsWithDefaults() *Modifications {
	this := Modifications{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *Modifications) GetIdentity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *Modifications) GetIdentityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *Modifications) SetIdentity(v string) {
	o.Identity = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *Modifications) GetOperations() []PatchItem {
	if o == nil || isNil(o.Operations) {
		var ret []PatchItem
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Modifications) GetOperationsOk() ([]PatchItem, bool) {
	if o == nil || isNil(o.Operations) {
    return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *Modifications) HasOperations() bool {
	if o != nil && !isNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []PatchItem and assigns it to the Operations field.
func (o *Modifications) SetOperations(v []PatchItem) {
	o.Operations = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Modifications) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Modifications) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Modifications) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *Modifications) SetTag(v string) {
	o.Tag = &v
}

func (o Modifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableModifications struct {
	value *Modifications
	isSet bool
}

func (v NullableModifications) Get() *Modifications {
	return v.value
}

func (v *NullableModifications) Set(val *Modifications) {
	v.value = val
	v.isSet = true
}

func (v NullableModifications) IsSet() bool {
	return v.isSet
}

func (v *NullableModifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifications(val *Modifications) *NullableModifications {
	return &NullableModifications{value: val, isSet: true}
}

func (v NullableModifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


