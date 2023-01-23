/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// VsDataContainerSingleAttributes struct for VsDataContainerSingleAttributes
type VsDataContainerSingleAttributes struct {
	VsDataType *string `json:"vsDataType,omitempty"`
	VsDataFormatVersion *string `json:"vsDataFormatVersion,omitempty"`
	VsData interface{} `json:"vsData,omitempty"`
}

// NewVsDataContainerSingleAttributes instantiates a new VsDataContainerSingleAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsDataContainerSingleAttributes() *VsDataContainerSingleAttributes {
	this := VsDataContainerSingleAttributes{}
	return &this
}

// NewVsDataContainerSingleAttributesWithDefaults instantiates a new VsDataContainerSingleAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsDataContainerSingleAttributesWithDefaults() *VsDataContainerSingleAttributes {
	this := VsDataContainerSingleAttributes{}
	return &this
}

// GetVsDataType returns the VsDataType field value if set, zero value otherwise.
func (o *VsDataContainerSingleAttributes) GetVsDataType() string {
	if o == nil || isNil(o.VsDataType) {
		var ret string
		return ret
	}
	return *o.VsDataType
}

// GetVsDataTypeOk returns a tuple with the VsDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingleAttributes) GetVsDataTypeOk() (*string, bool) {
	if o == nil || isNil(o.VsDataType) {
    return nil, false
	}
	return o.VsDataType, true
}

// HasVsDataType returns a boolean if a field has been set.
func (o *VsDataContainerSingleAttributes) HasVsDataType() bool {
	if o != nil && !isNil(o.VsDataType) {
		return true
	}

	return false
}

// SetVsDataType gets a reference to the given string and assigns it to the VsDataType field.
func (o *VsDataContainerSingleAttributes) SetVsDataType(v string) {
	o.VsDataType = &v
}

// GetVsDataFormatVersion returns the VsDataFormatVersion field value if set, zero value otherwise.
func (o *VsDataContainerSingleAttributes) GetVsDataFormatVersion() string {
	if o == nil || isNil(o.VsDataFormatVersion) {
		var ret string
		return ret
	}
	return *o.VsDataFormatVersion
}

// GetVsDataFormatVersionOk returns a tuple with the VsDataFormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingleAttributes) GetVsDataFormatVersionOk() (*string, bool) {
	if o == nil || isNil(o.VsDataFormatVersion) {
    return nil, false
	}
	return o.VsDataFormatVersion, true
}

// HasVsDataFormatVersion returns a boolean if a field has been set.
func (o *VsDataContainerSingleAttributes) HasVsDataFormatVersion() bool {
	if o != nil && !isNil(o.VsDataFormatVersion) {
		return true
	}

	return false
}

// SetVsDataFormatVersion gets a reference to the given string and assigns it to the VsDataFormatVersion field.
func (o *VsDataContainerSingleAttributes) SetVsDataFormatVersion(v string) {
	o.VsDataFormatVersion = &v
}

// GetVsData returns the VsData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsDataContainerSingleAttributes) GetVsData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.VsData
}

// GetVsDataOk returns a tuple with the VsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsDataContainerSingleAttributes) GetVsDataOk() (*interface{}, bool) {
	if o == nil || isNil(o.VsData) {
    return nil, false
	}
	return &o.VsData, true
}

// HasVsData returns a boolean if a field has been set.
func (o *VsDataContainerSingleAttributes) HasVsData() bool {
	if o != nil && isNil(o.VsData) {
		return true
	}

	return false
}

// SetVsData gets a reference to the given interface{} and assigns it to the VsData field.
func (o *VsDataContainerSingleAttributes) SetVsData(v interface{}) {
	o.VsData = v
}

func (o VsDataContainerSingleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VsDataType) {
		toSerialize["vsDataType"] = o.VsDataType
	}
	if !isNil(o.VsDataFormatVersion) {
		toSerialize["vsDataFormatVersion"] = o.VsDataFormatVersion
	}
	if o.VsData != nil {
		toSerialize["vsData"] = o.VsData
	}
	return json.Marshal(toSerialize)
}

type NullableVsDataContainerSingleAttributes struct {
	value *VsDataContainerSingleAttributes
	isSet bool
}

func (v NullableVsDataContainerSingleAttributes) Get() *VsDataContainerSingleAttributes {
	return v.value
}

func (v *NullableVsDataContainerSingleAttributes) Set(val *VsDataContainerSingleAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableVsDataContainerSingleAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableVsDataContainerSingleAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsDataContainerSingleAttributes(val *VsDataContainerSingleAttributes) *NullableVsDataContainerSingleAttributes {
	return &NullableVsDataContainerSingleAttributes{value: val, isSet: true}
}

func (v NullableVsDataContainerSingleAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsDataContainerSingleAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


