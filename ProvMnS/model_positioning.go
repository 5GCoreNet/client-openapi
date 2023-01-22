/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// Positioning struct for Positioning
type Positioning struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	Availability []string `json:"availability,omitempty"`
	Predictionfrequency *Predictionfrequency `json:"predictionfrequency,omitempty"`
	Accuracy *float32 `json:"accuracy,omitempty"`
}

// NewPositioning instantiates a new Positioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositioning() *Positioning {
	this := Positioning{}
	return &this
}

// NewPositioningWithDefaults instantiates a new Positioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositioningWithDefaults() *Positioning {
	this := Positioning{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *Positioning) GetServAttrCom() ServAttrCom {
	if o == nil || isNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Positioning) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || isNil(o.ServAttrCom) {
    return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *Positioning) HasServAttrCom() bool {
	if o != nil && !isNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *Positioning) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *Positioning) GetAvailability() []string {
	if o == nil || isNil(o.Availability) {
		var ret []string
		return ret
	}
	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Positioning) GetAvailabilityOk() ([]string, bool) {
	if o == nil || isNil(o.Availability) {
    return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *Positioning) HasAvailability() bool {
	if o != nil && !isNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given []string and assigns it to the Availability field.
func (o *Positioning) SetAvailability(v []string) {
	o.Availability = v
}

// GetPredictionfrequency returns the Predictionfrequency field value if set, zero value otherwise.
func (o *Positioning) GetPredictionfrequency() Predictionfrequency {
	if o == nil || isNil(o.Predictionfrequency) {
		var ret Predictionfrequency
		return ret
	}
	return *o.Predictionfrequency
}

// GetPredictionfrequencyOk returns a tuple with the Predictionfrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Positioning) GetPredictionfrequencyOk() (*Predictionfrequency, bool) {
	if o == nil || isNil(o.Predictionfrequency) {
    return nil, false
	}
	return o.Predictionfrequency, true
}

// HasPredictionfrequency returns a boolean if a field has been set.
func (o *Positioning) HasPredictionfrequency() bool {
	if o != nil && !isNil(o.Predictionfrequency) {
		return true
	}

	return false
}

// SetPredictionfrequency gets a reference to the given Predictionfrequency and assigns it to the Predictionfrequency field.
func (o *Positioning) SetPredictionfrequency(v Predictionfrequency) {
	o.Predictionfrequency = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *Positioning) GetAccuracy() float32 {
	if o == nil || isNil(o.Accuracy) {
		var ret float32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Positioning) GetAccuracyOk() (*float32, bool) {
	if o == nil || isNil(o.Accuracy) {
    return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *Positioning) HasAccuracy() bool {
	if o != nil && !isNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given float32 and assigns it to the Accuracy field.
func (o *Positioning) SetAccuracy(v float32) {
	o.Accuracy = &v
}

func (o Positioning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !isNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !isNil(o.Predictionfrequency) {
		toSerialize["predictionfrequency"] = o.Predictionfrequency
	}
	if !isNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	return json.Marshal(toSerialize)
}

type NullablePositioning struct {
	value *Positioning
	isSet bool
}

func (v NullablePositioning) Get() *Positioning {
	return v.value
}

func (v *NullablePositioning) Set(val *Positioning) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioning) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioning(val *Positioning) *NullablePositioning {
	return &NullablePositioning{value: val, isSet: true}
}

func (v NullablePositioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


