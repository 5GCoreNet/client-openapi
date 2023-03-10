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

// GeoAreaToCellMapping struct for GeoAreaToCellMapping
type GeoAreaToCellMapping struct {
	ConvexGeoPolygon []GeoCoordinate `json:"convexGeoPolygon,omitempty"`
	AssociationThreshold *int32 `json:"associationThreshold,omitempty"`
}

// NewGeoAreaToCellMapping instantiates a new GeoAreaToCellMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoAreaToCellMapping() *GeoAreaToCellMapping {
	this := GeoAreaToCellMapping{}
	return &this
}

// NewGeoAreaToCellMappingWithDefaults instantiates a new GeoAreaToCellMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoAreaToCellMappingWithDefaults() *GeoAreaToCellMapping {
	this := GeoAreaToCellMapping{}
	return &this
}

// GetConvexGeoPolygon returns the ConvexGeoPolygon field value if set, zero value otherwise.
func (o *GeoAreaToCellMapping) GetConvexGeoPolygon() []GeoCoordinate {
	if o == nil || isNil(o.ConvexGeoPolygon) {
		var ret []GeoCoordinate
		return ret
	}
	return o.ConvexGeoPolygon
}

// GetConvexGeoPolygonOk returns a tuple with the ConvexGeoPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoAreaToCellMapping) GetConvexGeoPolygonOk() ([]GeoCoordinate, bool) {
	if o == nil || isNil(o.ConvexGeoPolygon) {
    return nil, false
	}
	return o.ConvexGeoPolygon, true
}

// HasConvexGeoPolygon returns a boolean if a field has been set.
func (o *GeoAreaToCellMapping) HasConvexGeoPolygon() bool {
	if o != nil && !isNil(o.ConvexGeoPolygon) {
		return true
	}

	return false
}

// SetConvexGeoPolygon gets a reference to the given []GeoCoordinate and assigns it to the ConvexGeoPolygon field.
func (o *GeoAreaToCellMapping) SetConvexGeoPolygon(v []GeoCoordinate) {
	o.ConvexGeoPolygon = v
}

// GetAssociationThreshold returns the AssociationThreshold field value if set, zero value otherwise.
func (o *GeoAreaToCellMapping) GetAssociationThreshold() int32 {
	if o == nil || isNil(o.AssociationThreshold) {
		var ret int32
		return ret
	}
	return *o.AssociationThreshold
}

// GetAssociationThresholdOk returns a tuple with the AssociationThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoAreaToCellMapping) GetAssociationThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.AssociationThreshold) {
    return nil, false
	}
	return o.AssociationThreshold, true
}

// HasAssociationThreshold returns a boolean if a field has been set.
func (o *GeoAreaToCellMapping) HasAssociationThreshold() bool {
	if o != nil && !isNil(o.AssociationThreshold) {
		return true
	}

	return false
}

// SetAssociationThreshold gets a reference to the given int32 and assigns it to the AssociationThreshold field.
func (o *GeoAreaToCellMapping) SetAssociationThreshold(v int32) {
	o.AssociationThreshold = &v
}

func (o GeoAreaToCellMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConvexGeoPolygon) {
		toSerialize["convexGeoPolygon"] = o.ConvexGeoPolygon
	}
	if !isNil(o.AssociationThreshold) {
		toSerialize["associationThreshold"] = o.AssociationThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableGeoAreaToCellMapping struct {
	value *GeoAreaToCellMapping
	isSet bool
}

func (v NullableGeoAreaToCellMapping) Get() *GeoAreaToCellMapping {
	return v.value
}

func (v *NullableGeoAreaToCellMapping) Set(val *GeoAreaToCellMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoAreaToCellMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoAreaToCellMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoAreaToCellMapping(val *GeoAreaToCellMapping) *NullableGeoAreaToCellMapping {
	return &NullableGeoAreaToCellMapping{value: val, isSet: true}
}

func (v NullableGeoAreaToCellMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoAreaToCellMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


