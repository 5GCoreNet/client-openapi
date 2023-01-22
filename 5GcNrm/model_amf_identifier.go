/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
)

// AmfIdentifier AmfIdentifier comprise of amfRegionId, amfSetId and amfPointer
type AmfIdentifier struct {
	// AmfRegionId is defined in TS 23.003
	AmfRegionId *int32 `json:"amfRegionId,omitempty"`
	// AmfSetId is defined in TS 23.003
	AmfSetId *string `json:"amfSetId,omitempty"`
	// AmfPointer is defined in TS 23.003
	AmfPointer *int32 `json:"amfPointer,omitempty"`
}

// NewAmfIdentifier instantiates a new AmfIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfIdentifier() *AmfIdentifier {
	this := AmfIdentifier{}
	return &this
}

// NewAmfIdentifierWithDefaults instantiates a new AmfIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfIdentifierWithDefaults() *AmfIdentifier {
	this := AmfIdentifier{}
	return &this
}

// GetAmfRegionId returns the AmfRegionId field value if set, zero value otherwise.
func (o *AmfIdentifier) GetAmfRegionId() int32 {
	if o == nil || isNil(o.AmfRegionId) {
		var ret int32
		return ret
	}
	return *o.AmfRegionId
}

// GetAmfRegionIdOk returns a tuple with the AmfRegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfIdentifier) GetAmfRegionIdOk() (*int32, bool) {
	if o == nil || isNil(o.AmfRegionId) {
    return nil, false
	}
	return o.AmfRegionId, true
}

// HasAmfRegionId returns a boolean if a field has been set.
func (o *AmfIdentifier) HasAmfRegionId() bool {
	if o != nil && !isNil(o.AmfRegionId) {
		return true
	}

	return false
}

// SetAmfRegionId gets a reference to the given int32 and assigns it to the AmfRegionId field.
func (o *AmfIdentifier) SetAmfRegionId(v int32) {
	o.AmfRegionId = &v
}

// GetAmfSetId returns the AmfSetId field value if set, zero value otherwise.
func (o *AmfIdentifier) GetAmfSetId() string {
	if o == nil || isNil(o.AmfSetId) {
		var ret string
		return ret
	}
	return *o.AmfSetId
}

// GetAmfSetIdOk returns a tuple with the AmfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfIdentifier) GetAmfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfSetId) {
    return nil, false
	}
	return o.AmfSetId, true
}

// HasAmfSetId returns a boolean if a field has been set.
func (o *AmfIdentifier) HasAmfSetId() bool {
	if o != nil && !isNil(o.AmfSetId) {
		return true
	}

	return false
}

// SetAmfSetId gets a reference to the given string and assigns it to the AmfSetId field.
func (o *AmfIdentifier) SetAmfSetId(v string) {
	o.AmfSetId = &v
}

// GetAmfPointer returns the AmfPointer field value if set, zero value otherwise.
func (o *AmfIdentifier) GetAmfPointer() int32 {
	if o == nil || isNil(o.AmfPointer) {
		var ret int32
		return ret
	}
	return *o.AmfPointer
}

// GetAmfPointerOk returns a tuple with the AmfPointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfIdentifier) GetAmfPointerOk() (*int32, bool) {
	if o == nil || isNil(o.AmfPointer) {
    return nil, false
	}
	return o.AmfPointer, true
}

// HasAmfPointer returns a boolean if a field has been set.
func (o *AmfIdentifier) HasAmfPointer() bool {
	if o != nil && !isNil(o.AmfPointer) {
		return true
	}

	return false
}

// SetAmfPointer gets a reference to the given int32 and assigns it to the AmfPointer field.
func (o *AmfIdentifier) SetAmfPointer(v int32) {
	o.AmfPointer = &v
}

func (o AmfIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmfRegionId) {
		toSerialize["amfRegionId"] = o.AmfRegionId
	}
	if !isNil(o.AmfSetId) {
		toSerialize["amfSetId"] = o.AmfSetId
	}
	if !isNil(o.AmfPointer) {
		toSerialize["amfPointer"] = o.AmfPointer
	}
	return json.Marshal(toSerialize)
}

type NullableAmfIdentifier struct {
	value *AmfIdentifier
	isSet bool
}

func (v NullableAmfIdentifier) Get() *AmfIdentifier {
	return v.value
}

func (v *NullableAmfIdentifier) Set(val *AmfIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfIdentifier(val *AmfIdentifier) *NullableAmfIdentifier {
	return &NullableAmfIdentifier{value: val, isSet: true}
}

func (v NullableAmfIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


