/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// DispersionRequirement Represents the dispersion analytics requirements.
type DispersionRequirement struct {
	DisperType DispersionType `json:"disperType"`
	ClassCriters []ClassCriterion `json:"classCriters,omitempty"`
	RankCriters []RankingCriterion `json:"rankCriters,omitempty"`
	DispOrderCriter *DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`
	Order *MatchingDirection `json:"order,omitempty"`
}

// NewDispersionRequirement instantiates a new DispersionRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDispersionRequirement(disperType DispersionType) *DispersionRequirement {
	this := DispersionRequirement{}
	this.DisperType = disperType
	return &this
}

// NewDispersionRequirementWithDefaults instantiates a new DispersionRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDispersionRequirementWithDefaults() *DispersionRequirement {
	this := DispersionRequirement{}
	return &this
}

// GetDisperType returns the DisperType field value
func (o *DispersionRequirement) GetDisperType() DispersionType {
	if o == nil {
		var ret DispersionType
		return ret
	}

	return o.DisperType
}

// GetDisperTypeOk returns a tuple with the DisperType field value
// and a boolean to check if the value has been set.
func (o *DispersionRequirement) GetDisperTypeOk() (*DispersionType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DisperType, true
}

// SetDisperType sets field value
func (o *DispersionRequirement) SetDisperType(v DispersionType) {
	o.DisperType = v
}

// GetClassCriters returns the ClassCriters field value if set, zero value otherwise.
func (o *DispersionRequirement) GetClassCriters() []ClassCriterion {
	if o == nil || isNil(o.ClassCriters) {
		var ret []ClassCriterion
		return ret
	}
	return o.ClassCriters
}

// GetClassCritersOk returns a tuple with the ClassCriters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionRequirement) GetClassCritersOk() ([]ClassCriterion, bool) {
	if o == nil || isNil(o.ClassCriters) {
    return nil, false
	}
	return o.ClassCriters, true
}

// HasClassCriters returns a boolean if a field has been set.
func (o *DispersionRequirement) HasClassCriters() bool {
	if o != nil && !isNil(o.ClassCriters) {
		return true
	}

	return false
}

// SetClassCriters gets a reference to the given []ClassCriterion and assigns it to the ClassCriters field.
func (o *DispersionRequirement) SetClassCriters(v []ClassCriterion) {
	o.ClassCriters = v
}

// GetRankCriters returns the RankCriters field value if set, zero value otherwise.
func (o *DispersionRequirement) GetRankCriters() []RankingCriterion {
	if o == nil || isNil(o.RankCriters) {
		var ret []RankingCriterion
		return ret
	}
	return o.RankCriters
}

// GetRankCritersOk returns a tuple with the RankCriters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionRequirement) GetRankCritersOk() ([]RankingCriterion, bool) {
	if o == nil || isNil(o.RankCriters) {
    return nil, false
	}
	return o.RankCriters, true
}

// HasRankCriters returns a boolean if a field has been set.
func (o *DispersionRequirement) HasRankCriters() bool {
	if o != nil && !isNil(o.RankCriters) {
		return true
	}

	return false
}

// SetRankCriters gets a reference to the given []RankingCriterion and assigns it to the RankCriters field.
func (o *DispersionRequirement) SetRankCriters(v []RankingCriterion) {
	o.RankCriters = v
}

// GetDispOrderCriter returns the DispOrderCriter field value if set, zero value otherwise.
func (o *DispersionRequirement) GetDispOrderCriter() DispersionOrderingCriterion {
	if o == nil || isNil(o.DispOrderCriter) {
		var ret DispersionOrderingCriterion
		return ret
	}
	return *o.DispOrderCriter
}

// GetDispOrderCriterOk returns a tuple with the DispOrderCriter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionRequirement) GetDispOrderCriterOk() (*DispersionOrderingCriterion, bool) {
	if o == nil || isNil(o.DispOrderCriter) {
    return nil, false
	}
	return o.DispOrderCriter, true
}

// HasDispOrderCriter returns a boolean if a field has been set.
func (o *DispersionRequirement) HasDispOrderCriter() bool {
	if o != nil && !isNil(o.DispOrderCriter) {
		return true
	}

	return false
}

// SetDispOrderCriter gets a reference to the given DispersionOrderingCriterion and assigns it to the DispOrderCriter field.
func (o *DispersionRequirement) SetDispOrderCriter(v DispersionOrderingCriterion) {
	o.DispOrderCriter = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DispersionRequirement) GetOrder() MatchingDirection {
	if o == nil || isNil(o.Order) {
		var ret MatchingDirection
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionRequirement) GetOrderOk() (*MatchingDirection, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DispersionRequirement) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given MatchingDirection and assigns it to the Order field.
func (o *DispersionRequirement) SetOrder(v MatchingDirection) {
	o.Order = &v
}

func (o DispersionRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["disperType"] = o.DisperType
	}
	if !isNil(o.ClassCriters) {
		toSerialize["classCriters"] = o.ClassCriters
	}
	if !isNil(o.RankCriters) {
		toSerialize["rankCriters"] = o.RankCriters
	}
	if !isNil(o.DispOrderCriter) {
		toSerialize["dispOrderCriter"] = o.DispOrderCriter
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableDispersionRequirement struct {
	value *DispersionRequirement
	isSet bool
}

func (v NullableDispersionRequirement) Get() *DispersionRequirement {
	return v.value
}

func (v *NullableDispersionRequirement) Set(val *DispersionRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionRequirement(val *DispersionRequirement) *NullableDispersionRequirement {
	return &NullableDispersionRequirement{value: val, isSet: true}
}

func (v NullableDispersionRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


