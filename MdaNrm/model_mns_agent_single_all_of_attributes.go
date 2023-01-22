/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MdaNrm

import (
	"encoding/json"
)

// MnsAgentSingleAllOfAttributes struct for MnsAgentSingleAllOfAttributes
type MnsAgentSingleAllOfAttributes struct {
	SystemDN *string `json:"systemDN,omitempty"`
}

// NewMnsAgentSingleAllOfAttributes instantiates a new MnsAgentSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnsAgentSingleAllOfAttributes() *MnsAgentSingleAllOfAttributes {
	this := MnsAgentSingleAllOfAttributes{}
	return &this
}

// NewMnsAgentSingleAllOfAttributesWithDefaults instantiates a new MnsAgentSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnsAgentSingleAllOfAttributesWithDefaults() *MnsAgentSingleAllOfAttributes {
	this := MnsAgentSingleAllOfAttributes{}
	return &this
}

// GetSystemDN returns the SystemDN field value if set, zero value otherwise.
func (o *MnsAgentSingleAllOfAttributes) GetSystemDN() string {
	if o == nil || isNil(o.SystemDN) {
		var ret string
		return ret
	}
	return *o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsAgentSingleAllOfAttributes) GetSystemDNOk() (*string, bool) {
	if o == nil || isNil(o.SystemDN) {
    return nil, false
	}
	return o.SystemDN, true
}

// HasSystemDN returns a boolean if a field has been set.
func (o *MnsAgentSingleAllOfAttributes) HasSystemDN() bool {
	if o != nil && !isNil(o.SystemDN) {
		return true
	}

	return false
}

// SetSystemDN gets a reference to the given string and assigns it to the SystemDN field.
func (o *MnsAgentSingleAllOfAttributes) SetSystemDN(v string) {
	o.SystemDN = &v
}

func (o MnsAgentSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SystemDN) {
		toSerialize["systemDN"] = o.SystemDN
	}
	return json.Marshal(toSerialize)
}

type NullableMnsAgentSingleAllOfAttributes struct {
	value *MnsAgentSingleAllOfAttributes
	isSet bool
}

func (v NullableMnsAgentSingleAllOfAttributes) Get() *MnsAgentSingleAllOfAttributes {
	return v.value
}

func (v *NullableMnsAgentSingleAllOfAttributes) Set(val *MnsAgentSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMnsAgentSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMnsAgentSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnsAgentSingleAllOfAttributes(val *MnsAgentSingleAllOfAttributes) *NullableMnsAgentSingleAllOfAttributes {
	return &NullableMnsAgentSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableMnsAgentSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnsAgentSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


