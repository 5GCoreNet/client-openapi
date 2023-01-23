/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// MeContextSingleAllOfAttributes struct for MeContextSingleAllOfAttributes
type MeContextSingleAllOfAttributes struct {
	DnPrefix *string `json:"dnPrefix,omitempty"`
}

// NewMeContextSingleAllOfAttributes instantiates a new MeContextSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeContextSingleAllOfAttributes() *MeContextSingleAllOfAttributes {
	this := MeContextSingleAllOfAttributes{}
	return &this
}

// NewMeContextSingleAllOfAttributesWithDefaults instantiates a new MeContextSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeContextSingleAllOfAttributesWithDefaults() *MeContextSingleAllOfAttributes {
	this := MeContextSingleAllOfAttributes{}
	return &this
}

// GetDnPrefix returns the DnPrefix field value if set, zero value otherwise.
func (o *MeContextSingleAllOfAttributes) GetDnPrefix() string {
	if o == nil || isNil(o.DnPrefix) {
		var ret string
		return ret
	}
	return *o.DnPrefix
}

// GetDnPrefixOk returns a tuple with the DnPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeContextSingleAllOfAttributes) GetDnPrefixOk() (*string, bool) {
	if o == nil || isNil(o.DnPrefix) {
    return nil, false
	}
	return o.DnPrefix, true
}

// HasDnPrefix returns a boolean if a field has been set.
func (o *MeContextSingleAllOfAttributes) HasDnPrefix() bool {
	if o != nil && !isNil(o.DnPrefix) {
		return true
	}

	return false
}

// SetDnPrefix gets a reference to the given string and assigns it to the DnPrefix field.
func (o *MeContextSingleAllOfAttributes) SetDnPrefix(v string) {
	o.DnPrefix = &v
}

func (o MeContextSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DnPrefix) {
		toSerialize["dnPrefix"] = o.DnPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableMeContextSingleAllOfAttributes struct {
	value *MeContextSingleAllOfAttributes
	isSet bool
}

func (v NullableMeContextSingleAllOfAttributes) Get() *MeContextSingleAllOfAttributes {
	return v.value
}

func (v *NullableMeContextSingleAllOfAttributes) Set(val *MeContextSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMeContextSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMeContextSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeContextSingleAllOfAttributes(val *MeContextSingleAllOfAttributes) *NullableMeContextSingleAllOfAttributes {
	return &NullableMeContextSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableMeContextSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeContextSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


