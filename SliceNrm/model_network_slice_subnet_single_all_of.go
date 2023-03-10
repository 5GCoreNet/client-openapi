/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// NetworkSliceSubnetSingleAllOf struct for NetworkSliceSubnetSingleAllOf
type NetworkSliceSubnetSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewNetworkSliceSubnetSingleAllOf instantiates a new NetworkSliceSubnetSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSliceSubnetSingleAllOf() *NetworkSliceSubnetSingleAllOf {
	this := NetworkSliceSubnetSingleAllOf{}
	return &this
}

// NewNetworkSliceSubnetSingleAllOfWithDefaults instantiates a new NetworkSliceSubnetSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSliceSubnetSingleAllOfWithDefaults() *NetworkSliceSubnetSingleAllOf {
	this := NetworkSliceSubnetSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *NetworkSliceSubnetSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o NetworkSliceSubnetSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkSliceSubnetSingleAllOf struct {
	value *NetworkSliceSubnetSingleAllOf
	isSet bool
}

func (v NullableNetworkSliceSubnetSingleAllOf) Get() *NetworkSliceSubnetSingleAllOf {
	return v.value
}

func (v *NullableNetworkSliceSubnetSingleAllOf) Set(val *NetworkSliceSubnetSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceSubnetSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceSubnetSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceSubnetSingleAllOf(val *NetworkSliceSubnetSingleAllOf) *NullableNetworkSliceSubnetSingleAllOf {
	return &NullableNetworkSliceSubnetSingleAllOf{value: val, isSet: true}
}

func (v NullableNetworkSliceSubnetSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceSubnetSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


