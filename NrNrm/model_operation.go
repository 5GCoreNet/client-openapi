/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// Operation struct for Operation
type Operation struct {
	Name *string `json:"name,omitempty"`
	AllowedNFTypes *NFType `json:"allowedNFTypes,omitempty"`
	OperationSemantics *OperationSemantics `json:"operationSemantics,omitempty"`
}

// NewOperation instantiates a new Operation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperation() *Operation {
	this := Operation{}
	return &this
}

// NewOperationWithDefaults instantiates a new Operation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationWithDefaults() *Operation {
	this := Operation{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Operation) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Operation) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Operation) SetName(v string) {
	o.Name = &v
}

// GetAllowedNFTypes returns the AllowedNFTypes field value if set, zero value otherwise.
func (o *Operation) GetAllowedNFTypes() NFType {
	if o == nil || isNil(o.AllowedNFTypes) {
		var ret NFType
		return ret
	}
	return *o.AllowedNFTypes
}

// GetAllowedNFTypesOk returns a tuple with the AllowedNFTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetAllowedNFTypesOk() (*NFType, bool) {
	if o == nil || isNil(o.AllowedNFTypes) {
    return nil, false
	}
	return o.AllowedNFTypes, true
}

// HasAllowedNFTypes returns a boolean if a field has been set.
func (o *Operation) HasAllowedNFTypes() bool {
	if o != nil && !isNil(o.AllowedNFTypes) {
		return true
	}

	return false
}

// SetAllowedNFTypes gets a reference to the given NFType and assigns it to the AllowedNFTypes field.
func (o *Operation) SetAllowedNFTypes(v NFType) {
	o.AllowedNFTypes = &v
}

// GetOperationSemantics returns the OperationSemantics field value if set, zero value otherwise.
func (o *Operation) GetOperationSemantics() OperationSemantics {
	if o == nil || isNil(o.OperationSemantics) {
		var ret OperationSemantics
		return ret
	}
	return *o.OperationSemantics
}

// GetOperationSemanticsOk returns a tuple with the OperationSemantics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operation) GetOperationSemanticsOk() (*OperationSemantics, bool) {
	if o == nil || isNil(o.OperationSemantics) {
    return nil, false
	}
	return o.OperationSemantics, true
}

// HasOperationSemantics returns a boolean if a field has been set.
func (o *Operation) HasOperationSemantics() bool {
	if o != nil && !isNil(o.OperationSemantics) {
		return true
	}

	return false
}

// SetOperationSemantics gets a reference to the given OperationSemantics and assigns it to the OperationSemantics field.
func (o *Operation) SetOperationSemantics(v OperationSemantics) {
	o.OperationSemantics = &v
}

func (o Operation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.AllowedNFTypes) {
		toSerialize["allowedNFTypes"] = o.AllowedNFTypes
	}
	if !isNil(o.OperationSemantics) {
		toSerialize["operationSemantics"] = o.OperationSemantics
	}
	return json.Marshal(toSerialize)
}

type NullableOperation struct {
	value *Operation
	isSet bool
}

func (v NullableOperation) Get() *Operation {
	return v.value
}

func (v *NullableOperation) Set(val *Operation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation(val *Operation) *NullableOperation {
	return &NullableOperation{value: val, isSet: true}
}

func (v NullableOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


