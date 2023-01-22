/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// PlmnRestriction1 struct for PlmnRestriction1
type PlmnRestriction1 struct {
	RatRestrictions []RatType `json:"ratRestrictions,omitempty"`
	ForbiddenAreas []Area1 `json:"forbiddenAreas,omitempty"`
	ServiceAreaRestriction *ServiceAreaRestriction1 `json:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions []CoreNetworkType `json:"coreNetworkTypeRestrictions,omitempty"`
	PrimaryRatRestrictions []RatType `json:"primaryRatRestrictions,omitempty"`
	SecondaryRatRestrictions []RatType `json:"secondaryRatRestrictions,omitempty"`
}

// NewPlmnRestriction1 instantiates a new PlmnRestriction1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnRestriction1() *PlmnRestriction1 {
	this := PlmnRestriction1{}
	return &this
}

// NewPlmnRestriction1WithDefaults instantiates a new PlmnRestriction1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnRestriction1WithDefaults() *PlmnRestriction1 {
	this := PlmnRestriction1{}
	return &this
}

// GetRatRestrictions returns the RatRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction1) GetRatRestrictions() []RatType {
	if o == nil || isNil(o.RatRestrictions) {
		var ret []RatType
		return ret
	}
	return o.RatRestrictions
}

// GetRatRestrictionsOk returns a tuple with the RatRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction1) GetRatRestrictionsOk() ([]RatType, bool) {
	if o == nil || isNil(o.RatRestrictions) {
    return nil, false
	}
	return o.RatRestrictions, true
}

// HasRatRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction1) HasRatRestrictions() bool {
	if o != nil && !isNil(o.RatRestrictions) {
		return true
	}

	return false
}

// SetRatRestrictions gets a reference to the given []RatType and assigns it to the RatRestrictions field.
func (o *PlmnRestriction1) SetRatRestrictions(v []RatType) {
	o.RatRestrictions = v
}

// GetForbiddenAreas returns the ForbiddenAreas field value if set, zero value otherwise.
func (o *PlmnRestriction1) GetForbiddenAreas() []Area1 {
	if o == nil || isNil(o.ForbiddenAreas) {
		var ret []Area1
		return ret
	}
	return o.ForbiddenAreas
}

// GetForbiddenAreasOk returns a tuple with the ForbiddenAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction1) GetForbiddenAreasOk() ([]Area1, bool) {
	if o == nil || isNil(o.ForbiddenAreas) {
    return nil, false
	}
	return o.ForbiddenAreas, true
}

// HasForbiddenAreas returns a boolean if a field has been set.
func (o *PlmnRestriction1) HasForbiddenAreas() bool {
	if o != nil && !isNil(o.ForbiddenAreas) {
		return true
	}

	return false
}

// SetForbiddenAreas gets a reference to the given []Area1 and assigns it to the ForbiddenAreas field.
func (o *PlmnRestriction1) SetForbiddenAreas(v []Area1) {
	o.ForbiddenAreas = v
}

// GetServiceAreaRestriction returns the ServiceAreaRestriction field value if set, zero value otherwise.
func (o *PlmnRestriction1) GetServiceAreaRestriction() ServiceAreaRestriction1 {
	if o == nil || isNil(o.ServiceAreaRestriction) {
		var ret ServiceAreaRestriction1
		return ret
	}
	return *o.ServiceAreaRestriction
}

// GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction1) GetServiceAreaRestrictionOk() (*ServiceAreaRestriction1, bool) {
	if o == nil || isNil(o.ServiceAreaRestriction) {
    return nil, false
	}
	return o.ServiceAreaRestriction, true
}

// HasServiceAreaRestriction returns a boolean if a field has been set.
func (o *PlmnRestriction1) HasServiceAreaRestriction() bool {
	if o != nil && !isNil(o.ServiceAreaRestriction) {
		return true
	}

	return false
}

// SetServiceAreaRestriction gets a reference to the given ServiceAreaRestriction1 and assigns it to the ServiceAreaRestriction field.
func (o *PlmnRestriction1) SetServiceAreaRestriction(v ServiceAreaRestriction1) {
	o.ServiceAreaRestriction = &v
}

// GetCoreNetworkTypeRestrictions returns the CoreNetworkTypeRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction1) GetCoreNetworkTypeRestrictions() []CoreNetworkType {
	if o == nil || isNil(o.CoreNetworkTypeRestrictions) {
		var ret []CoreNetworkType
		return ret
	}
	return o.CoreNetworkTypeRestrictions
}

// GetCoreNetworkTypeRestrictionsOk returns a tuple with the CoreNetworkTypeRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction1) GetCoreNetworkTypeRestrictionsOk() ([]CoreNetworkType, bool) {
	if o == nil || isNil(o.CoreNetworkTypeRestrictions) {
    return nil, false
	}
	return o.CoreNetworkTypeRestrictions, true
}

// HasCoreNetworkTypeRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction1) HasCoreNetworkTypeRestrictions() bool {
	if o != nil && !isNil(o.CoreNetworkTypeRestrictions) {
		return true
	}

	return false
}

// SetCoreNetworkTypeRestrictions gets a reference to the given []CoreNetworkType and assigns it to the CoreNetworkTypeRestrictions field.
func (o *PlmnRestriction1) SetCoreNetworkTypeRestrictions(v []CoreNetworkType) {
	o.CoreNetworkTypeRestrictions = v
}

// GetPrimaryRatRestrictions returns the PrimaryRatRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction1) GetPrimaryRatRestrictions() []RatType {
	if o == nil || isNil(o.PrimaryRatRestrictions) {
		var ret []RatType
		return ret
	}
	return o.PrimaryRatRestrictions
}

// GetPrimaryRatRestrictionsOk returns a tuple with the PrimaryRatRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction1) GetPrimaryRatRestrictionsOk() ([]RatType, bool) {
	if o == nil || isNil(o.PrimaryRatRestrictions) {
    return nil, false
	}
	return o.PrimaryRatRestrictions, true
}

// HasPrimaryRatRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction1) HasPrimaryRatRestrictions() bool {
	if o != nil && !isNil(o.PrimaryRatRestrictions) {
		return true
	}

	return false
}

// SetPrimaryRatRestrictions gets a reference to the given []RatType and assigns it to the PrimaryRatRestrictions field.
func (o *PlmnRestriction1) SetPrimaryRatRestrictions(v []RatType) {
	o.PrimaryRatRestrictions = v
}

// GetSecondaryRatRestrictions returns the SecondaryRatRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction1) GetSecondaryRatRestrictions() []RatType {
	if o == nil || isNil(o.SecondaryRatRestrictions) {
		var ret []RatType
		return ret
	}
	return o.SecondaryRatRestrictions
}

// GetSecondaryRatRestrictionsOk returns a tuple with the SecondaryRatRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction1) GetSecondaryRatRestrictionsOk() ([]RatType, bool) {
	if o == nil || isNil(o.SecondaryRatRestrictions) {
    return nil, false
	}
	return o.SecondaryRatRestrictions, true
}

// HasSecondaryRatRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction1) HasSecondaryRatRestrictions() bool {
	if o != nil && !isNil(o.SecondaryRatRestrictions) {
		return true
	}

	return false
}

// SetSecondaryRatRestrictions gets a reference to the given []RatType and assigns it to the SecondaryRatRestrictions field.
func (o *PlmnRestriction1) SetSecondaryRatRestrictions(v []RatType) {
	o.SecondaryRatRestrictions = v
}

func (o PlmnRestriction1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RatRestrictions) {
		toSerialize["ratRestrictions"] = o.RatRestrictions
	}
	if !isNil(o.ForbiddenAreas) {
		toSerialize["forbiddenAreas"] = o.ForbiddenAreas
	}
	if !isNil(o.ServiceAreaRestriction) {
		toSerialize["serviceAreaRestriction"] = o.ServiceAreaRestriction
	}
	if !isNil(o.CoreNetworkTypeRestrictions) {
		toSerialize["coreNetworkTypeRestrictions"] = o.CoreNetworkTypeRestrictions
	}
	if !isNil(o.PrimaryRatRestrictions) {
		toSerialize["primaryRatRestrictions"] = o.PrimaryRatRestrictions
	}
	if !isNil(o.SecondaryRatRestrictions) {
		toSerialize["secondaryRatRestrictions"] = o.SecondaryRatRestrictions
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnRestriction1 struct {
	value *PlmnRestriction1
	isSet bool
}

func (v NullablePlmnRestriction1) Get() *PlmnRestriction1 {
	return v.value
}

func (v *NullablePlmnRestriction1) Set(val *PlmnRestriction1) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnRestriction1) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnRestriction1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnRestriction1(val *PlmnRestriction1) *NullablePlmnRestriction1 {
	return &NullablePlmnRestriction1{value: val, isSet: true}
}

func (v NullablePlmnRestriction1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnRestriction1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


