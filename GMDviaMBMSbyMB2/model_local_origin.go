/*
GMDviaMBMSbyMB2

API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package GMDviaMBMSbyMB2

import (
	"encoding/json"
)

// LocalOrigin Indicates a Local origin in a reference system
type LocalOrigin struct {
	CoordinateId *string `json:"coordinateId,omitempty"`
	Point *GeographicalCoordinates `json:"point,omitempty"`
}

// NewLocalOrigin instantiates a new LocalOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalOrigin() *LocalOrigin {
	this := LocalOrigin{}
	return &this
}

// NewLocalOriginWithDefaults instantiates a new LocalOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalOriginWithDefaults() *LocalOrigin {
	this := LocalOrigin{}
	return &this
}

// GetCoordinateId returns the CoordinateId field value if set, zero value otherwise.
func (o *LocalOrigin) GetCoordinateId() string {
	if o == nil || isNil(o.CoordinateId) {
		var ret string
		return ret
	}
	return *o.CoordinateId
}

// GetCoordinateIdOk returns a tuple with the CoordinateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalOrigin) GetCoordinateIdOk() (*string, bool) {
	if o == nil || isNil(o.CoordinateId) {
    return nil, false
	}
	return o.CoordinateId, true
}

// HasCoordinateId returns a boolean if a field has been set.
func (o *LocalOrigin) HasCoordinateId() bool {
	if o != nil && !isNil(o.CoordinateId) {
		return true
	}

	return false
}

// SetCoordinateId gets a reference to the given string and assigns it to the CoordinateId field.
func (o *LocalOrigin) SetCoordinateId(v string) {
	o.CoordinateId = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *LocalOrigin) GetPoint() GeographicalCoordinates {
	if o == nil || isNil(o.Point) {
		var ret GeographicalCoordinates
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalOrigin) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil || isNil(o.Point) {
    return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *LocalOrigin) HasPoint() bool {
	if o != nil && !isNil(o.Point) {
		return true
	}

	return false
}

// SetPoint gets a reference to the given GeographicalCoordinates and assigns it to the Point field.
func (o *LocalOrigin) SetPoint(v GeographicalCoordinates) {
	o.Point = &v
}

func (o LocalOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CoordinateId) {
		toSerialize["coordinateId"] = o.CoordinateId
	}
	if !isNil(o.Point) {
		toSerialize["point"] = o.Point
	}
	return json.Marshal(toSerialize)
}

type NullableLocalOrigin struct {
	value *LocalOrigin
	isSet bool
}

func (v NullableLocalOrigin) Get() *LocalOrigin {
	return v.value
}

func (v *NullableLocalOrigin) Set(val *LocalOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalOrigin(val *LocalOrigin) *NullableLocalOrigin {
	return &NullableLocalOrigin{value: val, isSet: true}
}

func (v NullableLocalOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


