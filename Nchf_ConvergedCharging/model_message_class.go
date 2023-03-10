/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// MessageClass struct for MessageClass
type MessageClass struct {
	ClassIdentifier *ClassIdentifier `json:"classIdentifier,omitempty"`
	TokenText *string `json:"tokenText,omitempty"`
}

// NewMessageClass instantiates a new MessageClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageClass() *MessageClass {
	this := MessageClass{}
	return &this
}

// NewMessageClassWithDefaults instantiates a new MessageClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageClassWithDefaults() *MessageClass {
	this := MessageClass{}
	return &this
}

// GetClassIdentifier returns the ClassIdentifier field value if set, zero value otherwise.
func (o *MessageClass) GetClassIdentifier() ClassIdentifier {
	if o == nil || isNil(o.ClassIdentifier) {
		var ret ClassIdentifier
		return ret
	}
	return *o.ClassIdentifier
}

// GetClassIdentifierOk returns a tuple with the ClassIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageClass) GetClassIdentifierOk() (*ClassIdentifier, bool) {
	if o == nil || isNil(o.ClassIdentifier) {
    return nil, false
	}
	return o.ClassIdentifier, true
}

// HasClassIdentifier returns a boolean if a field has been set.
func (o *MessageClass) HasClassIdentifier() bool {
	if o != nil && !isNil(o.ClassIdentifier) {
		return true
	}

	return false
}

// SetClassIdentifier gets a reference to the given ClassIdentifier and assigns it to the ClassIdentifier field.
func (o *MessageClass) SetClassIdentifier(v ClassIdentifier) {
	o.ClassIdentifier = &v
}

// GetTokenText returns the TokenText field value if set, zero value otherwise.
func (o *MessageClass) GetTokenText() string {
	if o == nil || isNil(o.TokenText) {
		var ret string
		return ret
	}
	return *o.TokenText
}

// GetTokenTextOk returns a tuple with the TokenText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageClass) GetTokenTextOk() (*string, bool) {
	if o == nil || isNil(o.TokenText) {
    return nil, false
	}
	return o.TokenText, true
}

// HasTokenText returns a boolean if a field has been set.
func (o *MessageClass) HasTokenText() bool {
	if o != nil && !isNil(o.TokenText) {
		return true
	}

	return false
}

// SetTokenText gets a reference to the given string and assigns it to the TokenText field.
func (o *MessageClass) SetTokenText(v string) {
	o.TokenText = &v
}

func (o MessageClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClassIdentifier) {
		toSerialize["classIdentifier"] = o.ClassIdentifier
	}
	if !isNil(o.TokenText) {
		toSerialize["tokenText"] = o.TokenText
	}
	return json.Marshal(toSerialize)
}

type NullableMessageClass struct {
	value *MessageClass
	isSet bool
}

func (v NullableMessageClass) Get() *MessageClass {
	return v.value
}

func (v *NullableMessageClass) Set(val *MessageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageClass(val *MessageClass) *NullableMessageClass {
	return &NullableMessageClass{value: val, isSet: true}
}

func (v NullableMessageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


