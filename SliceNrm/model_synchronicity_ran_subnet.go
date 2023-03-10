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

// SynchronicityRANSubnet struct for SynchronicityRANSubnet
type SynchronicityRANSubnet struct {
	Availability *SynAvailability `json:"availability,omitempty"`
	Accuracy *float32 `json:"accuracy,omitempty"`
}

// NewSynchronicityRANSubnet instantiates a new SynchronicityRANSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSynchronicityRANSubnet() *SynchronicityRANSubnet {
	this := SynchronicityRANSubnet{}
	return &this
}

// NewSynchronicityRANSubnetWithDefaults instantiates a new SynchronicityRANSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSynchronicityRANSubnetWithDefaults() *SynchronicityRANSubnet {
	this := SynchronicityRANSubnet{}
	return &this
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *SynchronicityRANSubnet) GetAvailability() SynAvailability {
	if o == nil || isNil(o.Availability) {
		var ret SynAvailability
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynchronicityRANSubnet) GetAvailabilityOk() (*SynAvailability, bool) {
	if o == nil || isNil(o.Availability) {
    return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *SynchronicityRANSubnet) HasAvailability() bool {
	if o != nil && !isNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given SynAvailability and assigns it to the Availability field.
func (o *SynchronicityRANSubnet) SetAvailability(v SynAvailability) {
	o.Availability = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *SynchronicityRANSubnet) GetAccuracy() float32 {
	if o == nil || isNil(o.Accuracy) {
		var ret float32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynchronicityRANSubnet) GetAccuracyOk() (*float32, bool) {
	if o == nil || isNil(o.Accuracy) {
    return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *SynchronicityRANSubnet) HasAccuracy() bool {
	if o != nil && !isNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given float32 and assigns it to the Accuracy field.
func (o *SynchronicityRANSubnet) SetAccuracy(v float32) {
	o.Accuracy = &v
}

func (o SynchronicityRANSubnet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !isNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	return json.Marshal(toSerialize)
}

type NullableSynchronicityRANSubnet struct {
	value *SynchronicityRANSubnet
	isSet bool
}

func (v NullableSynchronicityRANSubnet) Get() *SynchronicityRANSubnet {
	return v.value
}

func (v *NullableSynchronicityRANSubnet) Set(val *SynchronicityRANSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableSynchronicityRANSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableSynchronicityRANSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynchronicityRANSubnet(val *SynchronicityRANSubnet) *NullableSynchronicityRANSubnet {
	return &NullableSynchronicityRANSubnet{value: val, isSet: true}
}

func (v NullableSynchronicityRANSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSynchronicityRANSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


