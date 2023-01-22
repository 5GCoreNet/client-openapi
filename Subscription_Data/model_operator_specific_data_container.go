/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// OperatorSpecificDataContainer Container for operator specific data.
type OperatorSpecificDataContainer struct {
	DataType string `json:"dataType"`
	DataTypeDefinition *string `json:"dataTypeDefinition,omitempty"`
	Value OperatorSpecificDataContainerValue `json:"value"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewOperatorSpecificDataContainer instantiates a new OperatorSpecificDataContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorSpecificDataContainer(dataType string, value OperatorSpecificDataContainerValue) *OperatorSpecificDataContainer {
	this := OperatorSpecificDataContainer{}
	this.DataType = dataType
	this.Value = value
	return &this
}

// NewOperatorSpecificDataContainerWithDefaults instantiates a new OperatorSpecificDataContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorSpecificDataContainerWithDefaults() *OperatorSpecificDataContainer {
	this := OperatorSpecificDataContainer{}
	return &this
}

// GetDataType returns the DataType field value
func (o *OperatorSpecificDataContainer) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *OperatorSpecificDataContainer) GetDataTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *OperatorSpecificDataContainer) SetDataType(v string) {
	o.DataType = v
}

// GetDataTypeDefinition returns the DataTypeDefinition field value if set, zero value otherwise.
func (o *OperatorSpecificDataContainer) GetDataTypeDefinition() string {
	if o == nil || isNil(o.DataTypeDefinition) {
		var ret string
		return ret
	}
	return *o.DataTypeDefinition
}

// GetDataTypeDefinitionOk returns a tuple with the DataTypeDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorSpecificDataContainer) GetDataTypeDefinitionOk() (*string, bool) {
	if o == nil || isNil(o.DataTypeDefinition) {
    return nil, false
	}
	return o.DataTypeDefinition, true
}

// HasDataTypeDefinition returns a boolean if a field has been set.
func (o *OperatorSpecificDataContainer) HasDataTypeDefinition() bool {
	if o != nil && !isNil(o.DataTypeDefinition) {
		return true
	}

	return false
}

// SetDataTypeDefinition gets a reference to the given string and assigns it to the DataTypeDefinition field.
func (o *OperatorSpecificDataContainer) SetDataTypeDefinition(v string) {
	o.DataTypeDefinition = &v
}

// GetValue returns the Value field value
func (o *OperatorSpecificDataContainer) GetValue() OperatorSpecificDataContainerValue {
	if o == nil {
		var ret OperatorSpecificDataContainerValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OperatorSpecificDataContainer) GetValueOk() (*OperatorSpecificDataContainerValue, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OperatorSpecificDataContainer) SetValue(v OperatorSpecificDataContainerValue) {
	o.Value = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *OperatorSpecificDataContainer) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorSpecificDataContainer) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *OperatorSpecificDataContainer) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *OperatorSpecificDataContainer) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *OperatorSpecificDataContainer) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorSpecificDataContainer) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *OperatorSpecificDataContainer) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *OperatorSpecificDataContainer) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o OperatorSpecificDataContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataType"] = o.DataType
	}
	if !isNil(o.DataTypeDefinition) {
		toSerialize["dataTypeDefinition"] = o.DataTypeDefinition
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorSpecificDataContainer struct {
	value *OperatorSpecificDataContainer
	isSet bool
}

func (v NullableOperatorSpecificDataContainer) Get() *OperatorSpecificDataContainer {
	return v.value
}

func (v *NullableOperatorSpecificDataContainer) Set(val *OperatorSpecificDataContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorSpecificDataContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorSpecificDataContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorSpecificDataContainer(val *OperatorSpecificDataContainer) *NullableOperatorSpecificDataContainer {
	return &NullableOperatorSpecificDataContainer{value: val, isSet: true}
}

func (v NullableOperatorSpecificDataContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorSpecificDataContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


