/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// AssuranceReportSingleAllOf struct for AssuranceReportSingleAllOf
type AssuranceReportSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewAssuranceReportSingleAllOf instantiates a new AssuranceReportSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceReportSingleAllOf() *AssuranceReportSingleAllOf {
	this := AssuranceReportSingleAllOf{}
	return &this
}

// NewAssuranceReportSingleAllOfWithDefaults instantiates a new AssuranceReportSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceReportSingleAllOfWithDefaults() *AssuranceReportSingleAllOf {
	this := AssuranceReportSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AssuranceReportSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceReportSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssuranceReportSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *AssuranceReportSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o AssuranceReportSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableAssuranceReportSingleAllOf struct {
	value *AssuranceReportSingleAllOf
	isSet bool
}

func (v NullableAssuranceReportSingleAllOf) Get() *AssuranceReportSingleAllOf {
	return v.value
}

func (v *NullableAssuranceReportSingleAllOf) Set(val *AssuranceReportSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceReportSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceReportSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceReportSingleAllOf(val *AssuranceReportSingleAllOf) *NullableAssuranceReportSingleAllOf {
	return &NullableAssuranceReportSingleAllOf{value: val, isSet: true}
}

func (v NullableAssuranceReportSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceReportSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


