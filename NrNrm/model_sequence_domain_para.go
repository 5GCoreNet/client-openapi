/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// SequenceDomainPara struct for SequenceDomainPara
type SequenceDomainPara struct {
	NrofRIMRSSequenceCandidatesofRS1 *int32 `json:"nrofRIMRSSequenceCandidatesofRS1,omitempty"`
	RimRSScrambleIdListofRS1 []int32 `json:"rimRSScrambleIdListofRS1,omitempty"`
	NrofRIMRSSequenceCandidatesofRS2 *int32 `json:"nrofRIMRSSequenceCandidatesofRS2,omitempty"`
	RimRSScrambleIdListofRS2 []int32 `json:"rimRSScrambleIdListofRS2,omitempty"`
	EnableEnoughNotEnoughIndication *string `json:"enableEnoughNotEnoughIndication,omitempty"`
	RIMRSScrambleTimerMultiplier *int32 `json:"RIMRSScrambleTimerMultiplier,omitempty"`
	RIMRSScrambleTimerOffset *int32 `json:"RIMRSScrambleTimerOffset,omitempty"`
}

// NewSequenceDomainPara instantiates a new SequenceDomainPara object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSequenceDomainPara() *SequenceDomainPara {
	this := SequenceDomainPara{}
	return &this
}

// NewSequenceDomainParaWithDefaults instantiates a new SequenceDomainPara object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSequenceDomainParaWithDefaults() *SequenceDomainPara {
	this := SequenceDomainPara{}
	return &this
}

// GetNrofRIMRSSequenceCandidatesofRS1 returns the NrofRIMRSSequenceCandidatesofRS1 field value if set, zero value otherwise.
func (o *SequenceDomainPara) GetNrofRIMRSSequenceCandidatesofRS1() int32 {
	if o == nil || isNil(o.NrofRIMRSSequenceCandidatesofRS1) {
		var ret int32
		return ret
	}
	return *o.NrofRIMRSSequenceCandidatesofRS1
}

// GetNrofRIMRSSequenceCandidatesofRS1Ok returns a tuple with the NrofRIMRSSequenceCandidatesofRS1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceDomainPara) GetNrofRIMRSSequenceCandidatesofRS1Ok() (*int32, bool) {
	if o == nil || isNil(o.NrofRIMRSSequenceCandidatesofRS1) {
    return nil, false
	}
	return o.NrofRIMRSSequenceCandidatesofRS1, true
}

// HasNrofRIMRSSequenceCandidatesofRS1 returns a boolean if a field has been set.
func (o *SequenceDomainPara) HasNrofRIMRSSequenceCandidatesofRS1() bool {
	if o != nil && !isNil(o.NrofRIMRSSequenceCandidatesofRS1) {
		return true
	}

	return false
}

// SetNrofRIMRSSequenceCandidatesofRS1 gets a reference to the given int32 and assigns it to the NrofRIMRSSequenceCandidatesofRS1 field.
func (o *SequenceDomainPara) SetNrofRIMRSSequenceCandidatesofRS1(v int32) {
	o.NrofRIMRSSequenceCandidatesofRS1 = &v
}

// GetRimRSScrambleIdListofRS1 returns the RimRSScrambleIdListofRS1 field value if set, zero value otherwise.
func (o *SequenceDomainPara) GetRimRSScrambleIdListofRS1() []int32 {
	if o == nil || isNil(o.RimRSScrambleIdListofRS1) {
		var ret []int32
		return ret
	}
	return o.RimRSScrambleIdListofRS1
}

// GetRimRSScrambleIdListofRS1Ok returns a tuple with the RimRSScrambleIdListofRS1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceDomainPara) GetRimRSScrambleIdListofRS1Ok() ([]int32, bool) {
	if o == nil || isNil(o.RimRSScrambleIdListofRS1) {
    return nil, false
	}
	return o.RimRSScrambleIdListofRS1, true
}

// HasRimRSScrambleIdListofRS1 returns a boolean if a field has been set.
func (o *SequenceDomainPara) HasRimRSScrambleIdListofRS1() bool {
	if o != nil && !isNil(o.RimRSScrambleIdListofRS1) {
		return true
	}

	return false
}

// SetRimRSScrambleIdListofRS1 gets a reference to the given []int32 and assigns it to the RimRSScrambleIdListofRS1 field.
func (o *SequenceDomainPara) SetRimRSScrambleIdListofRS1(v []int32) {
	o.RimRSScrambleIdListofRS1 = v
}

// GetNrofRIMRSSequenceCandidatesofRS2 returns the NrofRIMRSSequenceCandidatesofRS2 field value if set, zero value otherwise.
func (o *SequenceDomainPara) GetNrofRIMRSSequenceCandidatesofRS2() int32 {
	if o == nil || isNil(o.NrofRIMRSSequenceCandidatesofRS2) {
		var ret int32
		return ret
	}
	return *o.NrofRIMRSSequenceCandidatesofRS2
}

// GetNrofRIMRSSequenceCandidatesofRS2Ok returns a tuple with the NrofRIMRSSequenceCandidatesofRS2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceDomainPara) GetNrofRIMRSSequenceCandidatesofRS2Ok() (*int32, bool) {
	if o == nil || isNil(o.NrofRIMRSSequenceCandidatesofRS2) {
    return nil, false
	}
	return o.NrofRIMRSSequenceCandidatesofRS2, true
}

// HasNrofRIMRSSequenceCandidatesofRS2 returns a boolean if a field has been set.
func (o *SequenceDomainPara) HasNrofRIMRSSequenceCandidatesofRS2() bool {
	if o != nil && !isNil(o.NrofRIMRSSequenceCandidatesofRS2) {
		return true
	}

	return false
}

// SetNrofRIMRSSequenceCandidatesofRS2 gets a reference to the given int32 and assigns it to the NrofRIMRSSequenceCandidatesofRS2 field.
func (o *SequenceDomainPara) SetNrofRIMRSSequenceCandidatesofRS2(v int32) {
	o.NrofRIMRSSequenceCandidatesofRS2 = &v
}

// GetRimRSScrambleIdListofRS2 returns the RimRSScrambleIdListofRS2 field value if set, zero value otherwise.
func (o *SequenceDomainPara) GetRimRSScrambleIdListofRS2() []int32 {
	if o == nil || isNil(o.RimRSScrambleIdListofRS2) {
		var ret []int32
		return ret
	}
	return o.RimRSScrambleIdListofRS2
}

// GetRimRSScrambleIdListofRS2Ok returns a tuple with the RimRSScrambleIdListofRS2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceDomainPara) GetRimRSScrambleIdListofRS2Ok() ([]int32, bool) {
	if o == nil || isNil(o.RimRSScrambleIdListofRS2) {
    return nil, false
	}
	return o.RimRSScrambleIdListofRS2, true
}

// HasRimRSScrambleIdListofRS2 returns a boolean if a field has been set.
func (o *SequenceDomainPara) HasRimRSScrambleIdListofRS2() bool {
	if o != nil && !isNil(o.RimRSScrambleIdListofRS2) {
		return true
	}

	return false
}

// SetRimRSScrambleIdListofRS2 gets a reference to the given []int32 and assigns it to the RimRSScrambleIdListofRS2 field.
func (o *SequenceDomainPara) SetRimRSScrambleIdListofRS2(v []int32) {
	o.RimRSScrambleIdListofRS2 = v
}

// GetEnableEnoughNotEnoughIndication returns the EnableEnoughNotEnoughIndication field value if set, zero value otherwise.
func (o *SequenceDomainPara) GetEnableEnoughNotEnoughIndication() string {
	if o == nil || isNil(o.EnableEnoughNotEnoughIndication) {
		var ret string
		return ret
	}
	return *o.EnableEnoughNotEnoughIndication
}

// GetEnableEnoughNotEnoughIndicationOk returns a tuple with the EnableEnoughNotEnoughIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceDomainPara) GetEnableEnoughNotEnoughIndicationOk() (*string, bool) {
	if o == nil || isNil(o.EnableEnoughNotEnoughIndication) {
    return nil, false
	}
	return o.EnableEnoughNotEnoughIndication, true
}

// HasEnableEnoughNotEnoughIndication returns a boolean if a field has been set.
func (o *SequenceDomainPara) HasEnableEnoughNotEnoughIndication() bool {
	if o != nil && !isNil(o.EnableEnoughNotEnoughIndication) {
		return true
	}

	return false
}

// SetEnableEnoughNotEnoughIndication gets a reference to the given string and assigns it to the EnableEnoughNotEnoughIndication field.
func (o *SequenceDomainPara) SetEnableEnoughNotEnoughIndication(v string) {
	o.EnableEnoughNotEnoughIndication = &v
}

// GetRIMRSScrambleTimerMultiplier returns the RIMRSScrambleTimerMultiplier field value if set, zero value otherwise.
func (o *SequenceDomainPara) GetRIMRSScrambleTimerMultiplier() int32 {
	if o == nil || isNil(o.RIMRSScrambleTimerMultiplier) {
		var ret int32
		return ret
	}
	return *o.RIMRSScrambleTimerMultiplier
}

// GetRIMRSScrambleTimerMultiplierOk returns a tuple with the RIMRSScrambleTimerMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceDomainPara) GetRIMRSScrambleTimerMultiplierOk() (*int32, bool) {
	if o == nil || isNil(o.RIMRSScrambleTimerMultiplier) {
    return nil, false
	}
	return o.RIMRSScrambleTimerMultiplier, true
}

// HasRIMRSScrambleTimerMultiplier returns a boolean if a field has been set.
func (o *SequenceDomainPara) HasRIMRSScrambleTimerMultiplier() bool {
	if o != nil && !isNil(o.RIMRSScrambleTimerMultiplier) {
		return true
	}

	return false
}

// SetRIMRSScrambleTimerMultiplier gets a reference to the given int32 and assigns it to the RIMRSScrambleTimerMultiplier field.
func (o *SequenceDomainPara) SetRIMRSScrambleTimerMultiplier(v int32) {
	o.RIMRSScrambleTimerMultiplier = &v
}

// GetRIMRSScrambleTimerOffset returns the RIMRSScrambleTimerOffset field value if set, zero value otherwise.
func (o *SequenceDomainPara) GetRIMRSScrambleTimerOffset() int32 {
	if o == nil || isNil(o.RIMRSScrambleTimerOffset) {
		var ret int32
		return ret
	}
	return *o.RIMRSScrambleTimerOffset
}

// GetRIMRSScrambleTimerOffsetOk returns a tuple with the RIMRSScrambleTimerOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SequenceDomainPara) GetRIMRSScrambleTimerOffsetOk() (*int32, bool) {
	if o == nil || isNil(o.RIMRSScrambleTimerOffset) {
    return nil, false
	}
	return o.RIMRSScrambleTimerOffset, true
}

// HasRIMRSScrambleTimerOffset returns a boolean if a field has been set.
func (o *SequenceDomainPara) HasRIMRSScrambleTimerOffset() bool {
	if o != nil && !isNil(o.RIMRSScrambleTimerOffset) {
		return true
	}

	return false
}

// SetRIMRSScrambleTimerOffset gets a reference to the given int32 and assigns it to the RIMRSScrambleTimerOffset field.
func (o *SequenceDomainPara) SetRIMRSScrambleTimerOffset(v int32) {
	o.RIMRSScrambleTimerOffset = &v
}

func (o SequenceDomainPara) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NrofRIMRSSequenceCandidatesofRS1) {
		toSerialize["nrofRIMRSSequenceCandidatesofRS1"] = o.NrofRIMRSSequenceCandidatesofRS1
	}
	if !isNil(o.RimRSScrambleIdListofRS1) {
		toSerialize["rimRSScrambleIdListofRS1"] = o.RimRSScrambleIdListofRS1
	}
	if !isNil(o.NrofRIMRSSequenceCandidatesofRS2) {
		toSerialize["nrofRIMRSSequenceCandidatesofRS2"] = o.NrofRIMRSSequenceCandidatesofRS2
	}
	if !isNil(o.RimRSScrambleIdListofRS2) {
		toSerialize["rimRSScrambleIdListofRS2"] = o.RimRSScrambleIdListofRS2
	}
	if !isNil(o.EnableEnoughNotEnoughIndication) {
		toSerialize["enableEnoughNotEnoughIndication"] = o.EnableEnoughNotEnoughIndication
	}
	if !isNil(o.RIMRSScrambleTimerMultiplier) {
		toSerialize["RIMRSScrambleTimerMultiplier"] = o.RIMRSScrambleTimerMultiplier
	}
	if !isNil(o.RIMRSScrambleTimerOffset) {
		toSerialize["RIMRSScrambleTimerOffset"] = o.RIMRSScrambleTimerOffset
	}
	return json.Marshal(toSerialize)
}

type NullableSequenceDomainPara struct {
	value *SequenceDomainPara
	isSet bool
}

func (v NullableSequenceDomainPara) Get() *SequenceDomainPara {
	return v.value
}

func (v *NullableSequenceDomainPara) Set(val *SequenceDomainPara) {
	v.value = val
	v.isSet = true
}

func (v NullableSequenceDomainPara) IsSet() bool {
	return v.isSet
}

func (v *NullableSequenceDomainPara) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSequenceDomainPara(val *SequenceDomainPara) *NullableSequenceDomainPara {
	return &NullableSequenceDomainPara{value: val, isSet: true}
}

func (v NullableSequenceDomainPara) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSequenceDomainPara) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


