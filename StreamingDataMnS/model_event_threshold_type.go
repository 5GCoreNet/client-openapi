/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package StreamingDataMnS

import (
	"encoding/json"
)

// EventThresholdType See details in 3GPP TS 32.422 clause 5.10.7, 5.10.7a, 5.10.13 and 5.10.14.
type EventThresholdType struct {
	EventThresholdRSRP *EventThresholdTypeEventThresholdRSRP `json:"EventThresholdRSRP,omitempty"`
	EventThresholdRSRQ *EventThresholdTypeEventThresholdRSRQ `json:"EventThresholdRSRQ,omitempty"`
	EventThreshold1F *EventThresholdTypeEventThreshold1F `json:"EventThreshold1F,omitempty"`
	EventThreshold1I *int32 `json:"EventThreshold1I,omitempty"`
}

// NewEventThresholdType instantiates a new EventThresholdType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventThresholdType() *EventThresholdType {
	this := EventThresholdType{}
	return &this
}

// NewEventThresholdTypeWithDefaults instantiates a new EventThresholdType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventThresholdTypeWithDefaults() *EventThresholdType {
	this := EventThresholdType{}
	return &this
}

// GetEventThresholdRSRP returns the EventThresholdRSRP field value if set, zero value otherwise.
func (o *EventThresholdType) GetEventThresholdRSRP() EventThresholdTypeEventThresholdRSRP {
	if o == nil || isNil(o.EventThresholdRSRP) {
		var ret EventThresholdTypeEventThresholdRSRP
		return ret
	}
	return *o.EventThresholdRSRP
}

// GetEventThresholdRSRPOk returns a tuple with the EventThresholdRSRP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdType) GetEventThresholdRSRPOk() (*EventThresholdTypeEventThresholdRSRP, bool) {
	if o == nil || isNil(o.EventThresholdRSRP) {
    return nil, false
	}
	return o.EventThresholdRSRP, true
}

// HasEventThresholdRSRP returns a boolean if a field has been set.
func (o *EventThresholdType) HasEventThresholdRSRP() bool {
	if o != nil && !isNil(o.EventThresholdRSRP) {
		return true
	}

	return false
}

// SetEventThresholdRSRP gets a reference to the given EventThresholdTypeEventThresholdRSRP and assigns it to the EventThresholdRSRP field.
func (o *EventThresholdType) SetEventThresholdRSRP(v EventThresholdTypeEventThresholdRSRP) {
	o.EventThresholdRSRP = &v
}

// GetEventThresholdRSRQ returns the EventThresholdRSRQ field value if set, zero value otherwise.
func (o *EventThresholdType) GetEventThresholdRSRQ() EventThresholdTypeEventThresholdRSRQ {
	if o == nil || isNil(o.EventThresholdRSRQ) {
		var ret EventThresholdTypeEventThresholdRSRQ
		return ret
	}
	return *o.EventThresholdRSRQ
}

// GetEventThresholdRSRQOk returns a tuple with the EventThresholdRSRQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdType) GetEventThresholdRSRQOk() (*EventThresholdTypeEventThresholdRSRQ, bool) {
	if o == nil || isNil(o.EventThresholdRSRQ) {
    return nil, false
	}
	return o.EventThresholdRSRQ, true
}

// HasEventThresholdRSRQ returns a boolean if a field has been set.
func (o *EventThresholdType) HasEventThresholdRSRQ() bool {
	if o != nil && !isNil(o.EventThresholdRSRQ) {
		return true
	}

	return false
}

// SetEventThresholdRSRQ gets a reference to the given EventThresholdTypeEventThresholdRSRQ and assigns it to the EventThresholdRSRQ field.
func (o *EventThresholdType) SetEventThresholdRSRQ(v EventThresholdTypeEventThresholdRSRQ) {
	o.EventThresholdRSRQ = &v
}

// GetEventThreshold1F returns the EventThreshold1F field value if set, zero value otherwise.
func (o *EventThresholdType) GetEventThreshold1F() EventThresholdTypeEventThreshold1F {
	if o == nil || isNil(o.EventThreshold1F) {
		var ret EventThresholdTypeEventThreshold1F
		return ret
	}
	return *o.EventThreshold1F
}

// GetEventThreshold1FOk returns a tuple with the EventThreshold1F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdType) GetEventThreshold1FOk() (*EventThresholdTypeEventThreshold1F, bool) {
	if o == nil || isNil(o.EventThreshold1F) {
    return nil, false
	}
	return o.EventThreshold1F, true
}

// HasEventThreshold1F returns a boolean if a field has been set.
func (o *EventThresholdType) HasEventThreshold1F() bool {
	if o != nil && !isNil(o.EventThreshold1F) {
		return true
	}

	return false
}

// SetEventThreshold1F gets a reference to the given EventThresholdTypeEventThreshold1F and assigns it to the EventThreshold1F field.
func (o *EventThresholdType) SetEventThreshold1F(v EventThresholdTypeEventThreshold1F) {
	o.EventThreshold1F = &v
}

// GetEventThreshold1I returns the EventThreshold1I field value if set, zero value otherwise.
func (o *EventThresholdType) GetEventThreshold1I() int32 {
	if o == nil || isNil(o.EventThreshold1I) {
		var ret int32
		return ret
	}
	return *o.EventThreshold1I
}

// GetEventThreshold1IOk returns a tuple with the EventThreshold1I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventThresholdType) GetEventThreshold1IOk() (*int32, bool) {
	if o == nil || isNil(o.EventThreshold1I) {
    return nil, false
	}
	return o.EventThreshold1I, true
}

// HasEventThreshold1I returns a boolean if a field has been set.
func (o *EventThresholdType) HasEventThreshold1I() bool {
	if o != nil && !isNil(o.EventThreshold1I) {
		return true
	}

	return false
}

// SetEventThreshold1I gets a reference to the given int32 and assigns it to the EventThreshold1I field.
func (o *EventThresholdType) SetEventThreshold1I(v int32) {
	o.EventThreshold1I = &v
}

func (o EventThresholdType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventThresholdRSRP) {
		toSerialize["EventThresholdRSRP"] = o.EventThresholdRSRP
	}
	if !isNil(o.EventThresholdRSRQ) {
		toSerialize["EventThresholdRSRQ"] = o.EventThresholdRSRQ
	}
	if !isNil(o.EventThreshold1F) {
		toSerialize["EventThreshold1F"] = o.EventThreshold1F
	}
	if !isNil(o.EventThreshold1I) {
		toSerialize["EventThreshold1I"] = o.EventThreshold1I
	}
	return json.Marshal(toSerialize)
}

type NullableEventThresholdType struct {
	value *EventThresholdType
	isSet bool
}

func (v NullableEventThresholdType) Get() *EventThresholdType {
	return v.value
}

func (v *NullableEventThresholdType) Set(val *EventThresholdType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventThresholdType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventThresholdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventThresholdType(val *EventThresholdType) *NullableEventThresholdType {
	return &NullableEventThresholdType{value: val, isSet: true}
}

func (v NullableEventThresholdType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventThresholdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


