/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
)

// NetworkSliceSubnetProviderCapabilitiesSingleAllOf struct for NetworkSliceSubnetProviderCapabilitiesSingleAllOf
type NetworkSliceSubnetProviderCapabilitiesSingleAllOf struct {
	Attributes *NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewNetworkSliceSubnetProviderCapabilitiesSingleAllOf instantiates a new NetworkSliceSubnetProviderCapabilitiesSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSliceSubnetProviderCapabilitiesSingleAllOf() *NetworkSliceSubnetProviderCapabilitiesSingleAllOf {
	this := NetworkSliceSubnetProviderCapabilitiesSingleAllOf{}
	return &this
}

// NewNetworkSliceSubnetProviderCapabilitiesSingleAllOfWithDefaults instantiates a new NetworkSliceSubnetProviderCapabilitiesSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSliceSubnetProviderCapabilitiesSingleAllOfWithDefaults() *NetworkSliceSubnetProviderCapabilitiesSingleAllOf {
	this := NetworkSliceSubnetProviderCapabilitiesSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NetworkSliceSubnetProviderCapabilitiesSingleAllOf) GetAttributes() NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingleAllOf) GetAttributesOk() (*NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NetworkSliceSubnetProviderCapabilitiesSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NetworkSliceSubnetProviderCapabilitiesSingleAllOf) SetAttributes(v NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o NetworkSliceSubnetProviderCapabilitiesSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf struct {
	value *NetworkSliceSubnetProviderCapabilitiesSingleAllOf
	isSet bool
}

func (v NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf) Get() *NetworkSliceSubnetProviderCapabilitiesSingleAllOf {
	return v.value
}

func (v *NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf) Set(val *NetworkSliceSubnetProviderCapabilitiesSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf(val *NetworkSliceSubnetProviderCapabilitiesSingleAllOf) *NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf {
	return &NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf{value: val, isSet: true}
}

func (v NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceSubnetProviderCapabilitiesSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


