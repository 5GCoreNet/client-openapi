/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// PlmnRestriction struct for PlmnRestriction
type PlmnRestriction struct {
	RatRestrictions []RatType `json:"ratRestrictions,omitempty"`
	ForbiddenAreas []Area `json:"forbiddenAreas,omitempty"`
	ServiceAreaRestriction *ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions []CoreNetworkType `json:"coreNetworkTypeRestrictions,omitempty"`
	PrimaryRatRestrictions []RatType `json:"primaryRatRestrictions,omitempty"`
	SecondaryRatRestrictions []RatType `json:"secondaryRatRestrictions,omitempty"`
}

// NewPlmnRestriction instantiates a new PlmnRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnRestriction() *PlmnRestriction {
	this := PlmnRestriction{}
	return &this
}

// NewPlmnRestrictionWithDefaults instantiates a new PlmnRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnRestrictionWithDefaults() *PlmnRestriction {
	this := PlmnRestriction{}
	return &this
}

// GetRatRestrictions returns the RatRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction) GetRatRestrictions() []RatType {
	if o == nil || isNil(o.RatRestrictions) {
		var ret []RatType
		return ret
	}
	return o.RatRestrictions
}

// GetRatRestrictionsOk returns a tuple with the RatRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction) GetRatRestrictionsOk() ([]RatType, bool) {
	if o == nil || isNil(o.RatRestrictions) {
    return nil, false
	}
	return o.RatRestrictions, true
}

// HasRatRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction) HasRatRestrictions() bool {
	if o != nil && !isNil(o.RatRestrictions) {
		return true
	}

	return false
}

// SetRatRestrictions gets a reference to the given []RatType and assigns it to the RatRestrictions field.
func (o *PlmnRestriction) SetRatRestrictions(v []RatType) {
	o.RatRestrictions = v
}

// GetForbiddenAreas returns the ForbiddenAreas field value if set, zero value otherwise.
func (o *PlmnRestriction) GetForbiddenAreas() []Area {
	if o == nil || isNil(o.ForbiddenAreas) {
		var ret []Area
		return ret
	}
	return o.ForbiddenAreas
}

// GetForbiddenAreasOk returns a tuple with the ForbiddenAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction) GetForbiddenAreasOk() ([]Area, bool) {
	if o == nil || isNil(o.ForbiddenAreas) {
    return nil, false
	}
	return o.ForbiddenAreas, true
}

// HasForbiddenAreas returns a boolean if a field has been set.
func (o *PlmnRestriction) HasForbiddenAreas() bool {
	if o != nil && !isNil(o.ForbiddenAreas) {
		return true
	}

	return false
}

// SetForbiddenAreas gets a reference to the given []Area and assigns it to the ForbiddenAreas field.
func (o *PlmnRestriction) SetForbiddenAreas(v []Area) {
	o.ForbiddenAreas = v
}

// GetServiceAreaRestriction returns the ServiceAreaRestriction field value if set, zero value otherwise.
func (o *PlmnRestriction) GetServiceAreaRestriction() ServiceAreaRestriction {
	if o == nil || isNil(o.ServiceAreaRestriction) {
		var ret ServiceAreaRestriction
		return ret
	}
	return *o.ServiceAreaRestriction
}

// GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction) GetServiceAreaRestrictionOk() (*ServiceAreaRestriction, bool) {
	if o == nil || isNil(o.ServiceAreaRestriction) {
    return nil, false
	}
	return o.ServiceAreaRestriction, true
}

// HasServiceAreaRestriction returns a boolean if a field has been set.
func (o *PlmnRestriction) HasServiceAreaRestriction() bool {
	if o != nil && !isNil(o.ServiceAreaRestriction) {
		return true
	}

	return false
}

// SetServiceAreaRestriction gets a reference to the given ServiceAreaRestriction and assigns it to the ServiceAreaRestriction field.
func (o *PlmnRestriction) SetServiceAreaRestriction(v ServiceAreaRestriction) {
	o.ServiceAreaRestriction = &v
}

// GetCoreNetworkTypeRestrictions returns the CoreNetworkTypeRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction) GetCoreNetworkTypeRestrictions() []CoreNetworkType {
	if o == nil || isNil(o.CoreNetworkTypeRestrictions) {
		var ret []CoreNetworkType
		return ret
	}
	return o.CoreNetworkTypeRestrictions
}

// GetCoreNetworkTypeRestrictionsOk returns a tuple with the CoreNetworkTypeRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction) GetCoreNetworkTypeRestrictionsOk() ([]CoreNetworkType, bool) {
	if o == nil || isNil(o.CoreNetworkTypeRestrictions) {
    return nil, false
	}
	return o.CoreNetworkTypeRestrictions, true
}

// HasCoreNetworkTypeRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction) HasCoreNetworkTypeRestrictions() bool {
	if o != nil && !isNil(o.CoreNetworkTypeRestrictions) {
		return true
	}

	return false
}

// SetCoreNetworkTypeRestrictions gets a reference to the given []CoreNetworkType and assigns it to the CoreNetworkTypeRestrictions field.
func (o *PlmnRestriction) SetCoreNetworkTypeRestrictions(v []CoreNetworkType) {
	o.CoreNetworkTypeRestrictions = v
}

// GetPrimaryRatRestrictions returns the PrimaryRatRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction) GetPrimaryRatRestrictions() []RatType {
	if o == nil || isNil(o.PrimaryRatRestrictions) {
		var ret []RatType
		return ret
	}
	return o.PrimaryRatRestrictions
}

// GetPrimaryRatRestrictionsOk returns a tuple with the PrimaryRatRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction) GetPrimaryRatRestrictionsOk() ([]RatType, bool) {
	if o == nil || isNil(o.PrimaryRatRestrictions) {
    return nil, false
	}
	return o.PrimaryRatRestrictions, true
}

// HasPrimaryRatRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction) HasPrimaryRatRestrictions() bool {
	if o != nil && !isNil(o.PrimaryRatRestrictions) {
		return true
	}

	return false
}

// SetPrimaryRatRestrictions gets a reference to the given []RatType and assigns it to the PrimaryRatRestrictions field.
func (o *PlmnRestriction) SetPrimaryRatRestrictions(v []RatType) {
	o.PrimaryRatRestrictions = v
}

// GetSecondaryRatRestrictions returns the SecondaryRatRestrictions field value if set, zero value otherwise.
func (o *PlmnRestriction) GetSecondaryRatRestrictions() []RatType {
	if o == nil || isNil(o.SecondaryRatRestrictions) {
		var ret []RatType
		return ret
	}
	return o.SecondaryRatRestrictions
}

// GetSecondaryRatRestrictionsOk returns a tuple with the SecondaryRatRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnRestriction) GetSecondaryRatRestrictionsOk() ([]RatType, bool) {
	if o == nil || isNil(o.SecondaryRatRestrictions) {
    return nil, false
	}
	return o.SecondaryRatRestrictions, true
}

// HasSecondaryRatRestrictions returns a boolean if a field has been set.
func (o *PlmnRestriction) HasSecondaryRatRestrictions() bool {
	if o != nil && !isNil(o.SecondaryRatRestrictions) {
		return true
	}

	return false
}

// SetSecondaryRatRestrictions gets a reference to the given []RatType and assigns it to the SecondaryRatRestrictions field.
func (o *PlmnRestriction) SetSecondaryRatRestrictions(v []RatType) {
	o.SecondaryRatRestrictions = v
}

func (o PlmnRestriction) MarshalJSON() ([]byte, error) {
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

type NullablePlmnRestriction struct {
	value *PlmnRestriction
	isSet bool
}

func (v NullablePlmnRestriction) Get() *PlmnRestriction {
	return v.value
}

func (v *NullablePlmnRestriction) Set(val *PlmnRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnRestriction(val *PlmnRestriction) *NullablePlmnRestriction {
	return &NullablePlmnRestriction{value: val, isSet: true}
}

func (v NullablePlmnRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


