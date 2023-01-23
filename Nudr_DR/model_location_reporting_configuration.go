/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// LocationReportingConfiguration struct for LocationReportingConfiguration
type LocationReportingConfiguration struct {
	CurrentLocation bool `json:"currentLocation"`
	OneTime *bool `json:"oneTime,omitempty"`
	Accuracy *LocationAccuracy `json:"accuracy,omitempty"`
	N3gppAccuracy *LocationAccuracy `json:"n3gppAccuracy,omitempty"`
}

// NewLocationReportingConfiguration instantiates a new LocationReportingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationReportingConfiguration(currentLocation bool) *LocationReportingConfiguration {
	this := LocationReportingConfiguration{}
	this.CurrentLocation = currentLocation
	return &this
}

// NewLocationReportingConfigurationWithDefaults instantiates a new LocationReportingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationReportingConfigurationWithDefaults() *LocationReportingConfiguration {
	this := LocationReportingConfiguration{}
	return &this
}

// GetCurrentLocation returns the CurrentLocation field value
func (o *LocationReportingConfiguration) GetCurrentLocation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CurrentLocation
}

// GetCurrentLocationOk returns a tuple with the CurrentLocation field value
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration) GetCurrentLocationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CurrentLocation, true
}

// SetCurrentLocation sets field value
func (o *LocationReportingConfiguration) SetCurrentLocation(v bool) {
	o.CurrentLocation = v
}

// GetOneTime returns the OneTime field value if set, zero value otherwise.
func (o *LocationReportingConfiguration) GetOneTime() bool {
	if o == nil || isNil(o.OneTime) {
		var ret bool
		return ret
	}
	return *o.OneTime
}

// GetOneTimeOk returns a tuple with the OneTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration) GetOneTimeOk() (*bool, bool) {
	if o == nil || isNil(o.OneTime) {
    return nil, false
	}
	return o.OneTime, true
}

// HasOneTime returns a boolean if a field has been set.
func (o *LocationReportingConfiguration) HasOneTime() bool {
	if o != nil && !isNil(o.OneTime) {
		return true
	}

	return false
}

// SetOneTime gets a reference to the given bool and assigns it to the OneTime field.
func (o *LocationReportingConfiguration) SetOneTime(v bool) {
	o.OneTime = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *LocationReportingConfiguration) GetAccuracy() LocationAccuracy {
	if o == nil || isNil(o.Accuracy) {
		var ret LocationAccuracy
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration) GetAccuracyOk() (*LocationAccuracy, bool) {
	if o == nil || isNil(o.Accuracy) {
    return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *LocationReportingConfiguration) HasAccuracy() bool {
	if o != nil && !isNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given LocationAccuracy and assigns it to the Accuracy field.
func (o *LocationReportingConfiguration) SetAccuracy(v LocationAccuracy) {
	o.Accuracy = &v
}

// GetN3gppAccuracy returns the N3gppAccuracy field value if set, zero value otherwise.
func (o *LocationReportingConfiguration) GetN3gppAccuracy() LocationAccuracy {
	if o == nil || isNil(o.N3gppAccuracy) {
		var ret LocationAccuracy
		return ret
	}
	return *o.N3gppAccuracy
}

// GetN3gppAccuracyOk returns a tuple with the N3gppAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration) GetN3gppAccuracyOk() (*LocationAccuracy, bool) {
	if o == nil || isNil(o.N3gppAccuracy) {
    return nil, false
	}
	return o.N3gppAccuracy, true
}

// HasN3gppAccuracy returns a boolean if a field has been set.
func (o *LocationReportingConfiguration) HasN3gppAccuracy() bool {
	if o != nil && !isNil(o.N3gppAccuracy) {
		return true
	}

	return false
}

// SetN3gppAccuracy gets a reference to the given LocationAccuracy and assigns it to the N3gppAccuracy field.
func (o *LocationReportingConfiguration) SetN3gppAccuracy(v LocationAccuracy) {
	o.N3gppAccuracy = &v
}

func (o LocationReportingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currentLocation"] = o.CurrentLocation
	}
	if !isNil(o.OneTime) {
		toSerialize["oneTime"] = o.OneTime
	}
	if !isNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	if !isNil(o.N3gppAccuracy) {
		toSerialize["n3gppAccuracy"] = o.N3gppAccuracy
	}
	return json.Marshal(toSerialize)
}

type NullableLocationReportingConfiguration struct {
	value *LocationReportingConfiguration
	isSet bool
}

func (v NullableLocationReportingConfiguration) Get() *LocationReportingConfiguration {
	return v.value
}

func (v *NullableLocationReportingConfiguration) Set(val *LocationReportingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationReportingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationReportingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationReportingConfiguration(val *LocationReportingConfiguration) *NullableLocationReportingConfiguration {
	return &NullableLocationReportingConfiguration{value: val, isSet: true}
}

func (v NullableLocationReportingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationReportingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


