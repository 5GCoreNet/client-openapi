/*
M1_EventDataProcessingProvisioning

5GMS AF M1 Event Data Processing Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_EventDataProcessingProvisioning

import (
	"encoding/json"
)

// PointAltitudeUncertaintyAllOf struct for PointAltitudeUncertaintyAllOf
type PointAltitudeUncertaintyAllOf struct {
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`
	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewPointAltitudeUncertaintyAllOf instantiates a new PointAltitudeUncertaintyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointAltitudeUncertaintyAllOf(point GeographicalCoordinates, altitude float64, uncertaintyEllipse UncertaintyEllipse, uncertaintyAltitude float32, confidence int32) *PointAltitudeUncertaintyAllOf {
	this := PointAltitudeUncertaintyAllOf{}
	this.Point = point
	this.Altitude = altitude
	this.UncertaintyEllipse = uncertaintyEllipse
	this.UncertaintyAltitude = uncertaintyAltitude
	this.Confidence = confidence
	return &this
}

// NewPointAltitudeUncertaintyAllOfWithDefaults instantiates a new PointAltitudeUncertaintyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointAltitudeUncertaintyAllOfWithDefaults() *PointAltitudeUncertaintyAllOf {
	this := PointAltitudeUncertaintyAllOf{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointAltitudeUncertaintyAllOf) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertaintyAllOf) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointAltitudeUncertaintyAllOf) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetAltitude returns the Altitude field value
func (o *PointAltitudeUncertaintyAllOf) GetAltitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertaintyAllOf) GetAltitudeOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Altitude, true
}

// SetAltitude sets field value
func (o *PointAltitudeUncertaintyAllOf) SetAltitude(v float64) {
	o.Altitude = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *PointAltitudeUncertaintyAllOf) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetUncertaintyAltitude returns the UncertaintyAltitude field value
func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyAltitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UncertaintyAltitude
}

// GetUncertaintyAltitudeOk returns a tuple with the UncertaintyAltitude field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyAltitudeOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UncertaintyAltitude, true
}

// SetUncertaintyAltitude sets field value
func (o *PointAltitudeUncertaintyAllOf) SetUncertaintyAltitude(v float32) {
	o.UncertaintyAltitude = v
}

// GetConfidence returns the Confidence field value
func (o *PointAltitudeUncertaintyAllOf) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertaintyAllOf) GetConfidenceOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *PointAltitudeUncertaintyAllOf) SetConfidence(v int32) {
	o.Confidence = v
}

func (o PointAltitudeUncertaintyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["point"] = o.Point
	}
	if true {
		toSerialize["altitude"] = o.Altitude
	}
	if true {
		toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	}
	if true {
		toSerialize["uncertaintyAltitude"] = o.UncertaintyAltitude
	}
	if true {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullablePointAltitudeUncertaintyAllOf struct {
	value *PointAltitudeUncertaintyAllOf
	isSet bool
}

func (v NullablePointAltitudeUncertaintyAllOf) Get() *PointAltitudeUncertaintyAllOf {
	return v.value
}

func (v *NullablePointAltitudeUncertaintyAllOf) Set(val *PointAltitudeUncertaintyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePointAltitudeUncertaintyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePointAltitudeUncertaintyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointAltitudeUncertaintyAllOf(val *PointAltitudeUncertaintyAllOf) *NullablePointAltitudeUncertaintyAllOf {
	return &NullablePointAltitudeUncertaintyAllOf{value: val, isSet: true}
}

func (v NullablePointAltitudeUncertaintyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointAltitudeUncertaintyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


