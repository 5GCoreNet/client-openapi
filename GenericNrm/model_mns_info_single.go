/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package GenericNrm

import (
	"encoding/json"
)

// MnsInfoSingle struct for MnsInfoSingle
type MnsInfoSingle struct {
	MnsLabel *string `json:"mnsLabel,omitempty"`
	MnsType *string `json:"mnsType,omitempty"`
	MnsVersion *string `json:"mnsVersion,omitempty"`
	MnsAddress *string `json:"mnsAddress,omitempty"`
	// List of the managed object instances that can be accessed using the MnS. If a complete SubNetwork can be accessed using the MnS, this attribute may contain the DN of the SubNetwork instead of the DNs of the individual managed entities within the SubNetwork.
	MnsScope []string `json:"mnsScope,omitempty"`
}

// NewMnsInfoSingle instantiates a new MnsInfoSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnsInfoSingle() *MnsInfoSingle {
	this := MnsInfoSingle{}
	return &this
}

// NewMnsInfoSingleWithDefaults instantiates a new MnsInfoSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnsInfoSingleWithDefaults() *MnsInfoSingle {
	this := MnsInfoSingle{}
	return &this
}

// GetMnsLabel returns the MnsLabel field value if set, zero value otherwise.
func (o *MnsInfoSingle) GetMnsLabel() string {
	if o == nil || isNil(o.MnsLabel) {
		var ret string
		return ret
	}
	return *o.MnsLabel
}

// GetMnsLabelOk returns a tuple with the MnsLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsInfoSingle) GetMnsLabelOk() (*string, bool) {
	if o == nil || isNil(o.MnsLabel) {
    return nil, false
	}
	return o.MnsLabel, true
}

// HasMnsLabel returns a boolean if a field has been set.
func (o *MnsInfoSingle) HasMnsLabel() bool {
	if o != nil && !isNil(o.MnsLabel) {
		return true
	}

	return false
}

// SetMnsLabel gets a reference to the given string and assigns it to the MnsLabel field.
func (o *MnsInfoSingle) SetMnsLabel(v string) {
	o.MnsLabel = &v
}

// GetMnsType returns the MnsType field value if set, zero value otherwise.
func (o *MnsInfoSingle) GetMnsType() string {
	if o == nil || isNil(o.MnsType) {
		var ret string
		return ret
	}
	return *o.MnsType
}

// GetMnsTypeOk returns a tuple with the MnsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsInfoSingle) GetMnsTypeOk() (*string, bool) {
	if o == nil || isNil(o.MnsType) {
    return nil, false
	}
	return o.MnsType, true
}

// HasMnsType returns a boolean if a field has been set.
func (o *MnsInfoSingle) HasMnsType() bool {
	if o != nil && !isNil(o.MnsType) {
		return true
	}

	return false
}

// SetMnsType gets a reference to the given string and assigns it to the MnsType field.
func (o *MnsInfoSingle) SetMnsType(v string) {
	o.MnsType = &v
}

// GetMnsVersion returns the MnsVersion field value if set, zero value otherwise.
func (o *MnsInfoSingle) GetMnsVersion() string {
	if o == nil || isNil(o.MnsVersion) {
		var ret string
		return ret
	}
	return *o.MnsVersion
}

// GetMnsVersionOk returns a tuple with the MnsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsInfoSingle) GetMnsVersionOk() (*string, bool) {
	if o == nil || isNil(o.MnsVersion) {
    return nil, false
	}
	return o.MnsVersion, true
}

// HasMnsVersion returns a boolean if a field has been set.
func (o *MnsInfoSingle) HasMnsVersion() bool {
	if o != nil && !isNil(o.MnsVersion) {
		return true
	}

	return false
}

// SetMnsVersion gets a reference to the given string and assigns it to the MnsVersion field.
func (o *MnsInfoSingle) SetMnsVersion(v string) {
	o.MnsVersion = &v
}

// GetMnsAddress returns the MnsAddress field value if set, zero value otherwise.
func (o *MnsInfoSingle) GetMnsAddress() string {
	if o == nil || isNil(o.MnsAddress) {
		var ret string
		return ret
	}
	return *o.MnsAddress
}

// GetMnsAddressOk returns a tuple with the MnsAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsInfoSingle) GetMnsAddressOk() (*string, bool) {
	if o == nil || isNil(o.MnsAddress) {
    return nil, false
	}
	return o.MnsAddress, true
}

// HasMnsAddress returns a boolean if a field has been set.
func (o *MnsInfoSingle) HasMnsAddress() bool {
	if o != nil && !isNil(o.MnsAddress) {
		return true
	}

	return false
}

// SetMnsAddress gets a reference to the given string and assigns it to the MnsAddress field.
func (o *MnsInfoSingle) SetMnsAddress(v string) {
	o.MnsAddress = &v
}

// GetMnsScope returns the MnsScope field value if set, zero value otherwise.
func (o *MnsInfoSingle) GetMnsScope() []string {
	if o == nil || isNil(o.MnsScope) {
		var ret []string
		return ret
	}
	return o.MnsScope
}

// GetMnsScopeOk returns a tuple with the MnsScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsInfoSingle) GetMnsScopeOk() ([]string, bool) {
	if o == nil || isNil(o.MnsScope) {
    return nil, false
	}
	return o.MnsScope, true
}

// HasMnsScope returns a boolean if a field has been set.
func (o *MnsInfoSingle) HasMnsScope() bool {
	if o != nil && !isNil(o.MnsScope) {
		return true
	}

	return false
}

// SetMnsScope gets a reference to the given []string and assigns it to the MnsScope field.
func (o *MnsInfoSingle) SetMnsScope(v []string) {
	o.MnsScope = v
}

func (o MnsInfoSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MnsLabel) {
		toSerialize["mnsLabel"] = o.MnsLabel
	}
	if !isNil(o.MnsType) {
		toSerialize["mnsType"] = o.MnsType
	}
	if !isNil(o.MnsVersion) {
		toSerialize["mnsVersion"] = o.MnsVersion
	}
	if !isNil(o.MnsAddress) {
		toSerialize["mnsAddress"] = o.MnsAddress
	}
	if !isNil(o.MnsScope) {
		toSerialize["mnsScope"] = o.MnsScope
	}
	return json.Marshal(toSerialize)
}

type NullableMnsInfoSingle struct {
	value *MnsInfoSingle
	isSet bool
}

func (v NullableMnsInfoSingle) Get() *MnsInfoSingle {
	return v.value
}

func (v *NullableMnsInfoSingle) Set(val *MnsInfoSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableMnsInfoSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableMnsInfoSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnsInfoSingle(val *MnsInfoSingle) *NullableMnsInfoSingle {
	return &NullableMnsInfoSingle{value: val, isSet: true}
}

func (v NullableMnsInfoSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnsInfoSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


