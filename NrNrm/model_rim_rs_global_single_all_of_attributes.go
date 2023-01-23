/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// RimRSGlobalSingleAllOfAttributes struct for RimRSGlobalSingleAllOfAttributes
type RimRSGlobalSingleAllOfAttributes struct {
	FrequencyDomainPara *FrequencyDomainPara `json:"frequencyDomainPara,omitempty"`
	SequenceDomainPara *SequenceDomainPara `json:"sequenceDomainPara,omitempty"`
	TimeDomainPara *TimeDomainPara `json:"timeDomainPara,omitempty"`
}

// NewRimRSGlobalSingleAllOfAttributes instantiates a new RimRSGlobalSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRimRSGlobalSingleAllOfAttributes() *RimRSGlobalSingleAllOfAttributes {
	this := RimRSGlobalSingleAllOfAttributes{}
	return &this
}

// NewRimRSGlobalSingleAllOfAttributesWithDefaults instantiates a new RimRSGlobalSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRimRSGlobalSingleAllOfAttributesWithDefaults() *RimRSGlobalSingleAllOfAttributes {
	this := RimRSGlobalSingleAllOfAttributes{}
	return &this
}

// GetFrequencyDomainPara returns the FrequencyDomainPara field value if set, zero value otherwise.
func (o *RimRSGlobalSingleAllOfAttributes) GetFrequencyDomainPara() FrequencyDomainPara {
	if o == nil || isNil(o.FrequencyDomainPara) {
		var ret FrequencyDomainPara
		return ret
	}
	return *o.FrequencyDomainPara
}

// GetFrequencyDomainParaOk returns a tuple with the FrequencyDomainPara field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSGlobalSingleAllOfAttributes) GetFrequencyDomainParaOk() (*FrequencyDomainPara, bool) {
	if o == nil || isNil(o.FrequencyDomainPara) {
    return nil, false
	}
	return o.FrequencyDomainPara, true
}

// HasFrequencyDomainPara returns a boolean if a field has been set.
func (o *RimRSGlobalSingleAllOfAttributes) HasFrequencyDomainPara() bool {
	if o != nil && !isNil(o.FrequencyDomainPara) {
		return true
	}

	return false
}

// SetFrequencyDomainPara gets a reference to the given FrequencyDomainPara and assigns it to the FrequencyDomainPara field.
func (o *RimRSGlobalSingleAllOfAttributes) SetFrequencyDomainPara(v FrequencyDomainPara) {
	o.FrequencyDomainPara = &v
}

// GetSequenceDomainPara returns the SequenceDomainPara field value if set, zero value otherwise.
func (o *RimRSGlobalSingleAllOfAttributes) GetSequenceDomainPara() SequenceDomainPara {
	if o == nil || isNil(o.SequenceDomainPara) {
		var ret SequenceDomainPara
		return ret
	}
	return *o.SequenceDomainPara
}

// GetSequenceDomainParaOk returns a tuple with the SequenceDomainPara field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSGlobalSingleAllOfAttributes) GetSequenceDomainParaOk() (*SequenceDomainPara, bool) {
	if o == nil || isNil(o.SequenceDomainPara) {
    return nil, false
	}
	return o.SequenceDomainPara, true
}

// HasSequenceDomainPara returns a boolean if a field has been set.
func (o *RimRSGlobalSingleAllOfAttributes) HasSequenceDomainPara() bool {
	if o != nil && !isNil(o.SequenceDomainPara) {
		return true
	}

	return false
}

// SetSequenceDomainPara gets a reference to the given SequenceDomainPara and assigns it to the SequenceDomainPara field.
func (o *RimRSGlobalSingleAllOfAttributes) SetSequenceDomainPara(v SequenceDomainPara) {
	o.SequenceDomainPara = &v
}

// GetTimeDomainPara returns the TimeDomainPara field value if set, zero value otherwise.
func (o *RimRSGlobalSingleAllOfAttributes) GetTimeDomainPara() TimeDomainPara {
	if o == nil || isNil(o.TimeDomainPara) {
		var ret TimeDomainPara
		return ret
	}
	return *o.TimeDomainPara
}

// GetTimeDomainParaOk returns a tuple with the TimeDomainPara field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSGlobalSingleAllOfAttributes) GetTimeDomainParaOk() (*TimeDomainPara, bool) {
	if o == nil || isNil(o.TimeDomainPara) {
    return nil, false
	}
	return o.TimeDomainPara, true
}

// HasTimeDomainPara returns a boolean if a field has been set.
func (o *RimRSGlobalSingleAllOfAttributes) HasTimeDomainPara() bool {
	if o != nil && !isNil(o.TimeDomainPara) {
		return true
	}

	return false
}

// SetTimeDomainPara gets a reference to the given TimeDomainPara and assigns it to the TimeDomainPara field.
func (o *RimRSGlobalSingleAllOfAttributes) SetTimeDomainPara(v TimeDomainPara) {
	o.TimeDomainPara = &v
}

func (o RimRSGlobalSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FrequencyDomainPara) {
		toSerialize["frequencyDomainPara"] = o.FrequencyDomainPara
	}
	if !isNil(o.SequenceDomainPara) {
		toSerialize["sequenceDomainPara"] = o.SequenceDomainPara
	}
	if !isNil(o.TimeDomainPara) {
		toSerialize["timeDomainPara"] = o.TimeDomainPara
	}
	return json.Marshal(toSerialize)
}

type NullableRimRSGlobalSingleAllOfAttributes struct {
	value *RimRSGlobalSingleAllOfAttributes
	isSet bool
}

func (v NullableRimRSGlobalSingleAllOfAttributes) Get() *RimRSGlobalSingleAllOfAttributes {
	return v.value
}

func (v *NullableRimRSGlobalSingleAllOfAttributes) Set(val *RimRSGlobalSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableRimRSGlobalSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableRimRSGlobalSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRimRSGlobalSingleAllOfAttributes(val *RimRSGlobalSingleAllOfAttributes) *NullableRimRSGlobalSingleAllOfAttributes {
	return &NullableRimRSGlobalSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableRimRSGlobalSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRimRSGlobalSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


