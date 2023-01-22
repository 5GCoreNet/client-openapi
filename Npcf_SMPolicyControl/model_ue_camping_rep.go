/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
)

// UeCampingRep Contains the current applicable values corresponding to the policy control request triggers.
type UeCampingRep struct {
	AccessType *AccessType `json:"accessType,omitempty"`
	RatType *RatType `json:"ratType,omitempty"`
	ServNfId *ServingNfIdentity `json:"servNfId,omitempty"`
	ServingNetwork *PlmnIdNid `json:"servingNetwork,omitempty"`
	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UeTimeZone *string `json:"ueTimeZone,omitempty"`
	NetLocAccSupp *NetLocAccessSupport `json:"netLocAccSupp,omitempty"`
}

// NewUeCampingRep instantiates a new UeCampingRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeCampingRep() *UeCampingRep {
	this := UeCampingRep{}
	return &this
}

// NewUeCampingRepWithDefaults instantiates a new UeCampingRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeCampingRepWithDefaults() *UeCampingRep {
	this := UeCampingRep{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *UeCampingRep) GetAccessType() AccessType {
	if o == nil || isNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCampingRep) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || isNil(o.AccessType) {
    return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *UeCampingRep) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *UeCampingRep) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *UeCampingRep) GetRatType() RatType {
	if o == nil || isNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCampingRep) GetRatTypeOk() (*RatType, bool) {
	if o == nil || isNil(o.RatType) {
    return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *UeCampingRep) HasRatType() bool {
	if o != nil && !isNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *UeCampingRep) SetRatType(v RatType) {
	o.RatType = &v
}

// GetServNfId returns the ServNfId field value if set, zero value otherwise.
func (o *UeCampingRep) GetServNfId() ServingNfIdentity {
	if o == nil || isNil(o.ServNfId) {
		var ret ServingNfIdentity
		return ret
	}
	return *o.ServNfId
}

// GetServNfIdOk returns a tuple with the ServNfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCampingRep) GetServNfIdOk() (*ServingNfIdentity, bool) {
	if o == nil || isNil(o.ServNfId) {
    return nil, false
	}
	return o.ServNfId, true
}

// HasServNfId returns a boolean if a field has been set.
func (o *UeCampingRep) HasServNfId() bool {
	if o != nil && !isNil(o.ServNfId) {
		return true
	}

	return false
}

// SetServNfId gets a reference to the given ServingNfIdentity and assigns it to the ServNfId field.
func (o *UeCampingRep) SetServNfId(v ServingNfIdentity) {
	o.ServNfId = &v
}

// GetServingNetwork returns the ServingNetwork field value if set, zero value otherwise.
func (o *UeCampingRep) GetServingNetwork() PlmnIdNid {
	if o == nil || isNil(o.ServingNetwork) {
		var ret PlmnIdNid
		return ret
	}
	return *o.ServingNetwork
}

// GetServingNetworkOk returns a tuple with the ServingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCampingRep) GetServingNetworkOk() (*PlmnIdNid, bool) {
	if o == nil || isNil(o.ServingNetwork) {
    return nil, false
	}
	return o.ServingNetwork, true
}

// HasServingNetwork returns a boolean if a field has been set.
func (o *UeCampingRep) HasServingNetwork() bool {
	if o != nil && !isNil(o.ServingNetwork) {
		return true
	}

	return false
}

// SetServingNetwork gets a reference to the given PlmnIdNid and assigns it to the ServingNetwork field.
func (o *UeCampingRep) SetServingNetwork(v PlmnIdNid) {
	o.ServingNetwork = &v
}

// GetUserLocationInfo returns the UserLocationInfo field value if set, zero value otherwise.
func (o *UeCampingRep) GetUserLocationInfo() UserLocation {
	if o == nil || isNil(o.UserLocationInfo) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationInfo
}

// GetUserLocationInfoOk returns a tuple with the UserLocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCampingRep) GetUserLocationInfoOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UserLocationInfo) {
    return nil, false
	}
	return o.UserLocationInfo, true
}

// HasUserLocationInfo returns a boolean if a field has been set.
func (o *UeCampingRep) HasUserLocationInfo() bool {
	if o != nil && !isNil(o.UserLocationInfo) {
		return true
	}

	return false
}

// SetUserLocationInfo gets a reference to the given UserLocation and assigns it to the UserLocationInfo field.
func (o *UeCampingRep) SetUserLocationInfo(v UserLocation) {
	o.UserLocationInfo = &v
}

// GetUeTimeZone returns the UeTimeZone field value if set, zero value otherwise.
func (o *UeCampingRep) GetUeTimeZone() string {
	if o == nil || isNil(o.UeTimeZone) {
		var ret string
		return ret
	}
	return *o.UeTimeZone
}

// GetUeTimeZoneOk returns a tuple with the UeTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCampingRep) GetUeTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.UeTimeZone) {
    return nil, false
	}
	return o.UeTimeZone, true
}

// HasUeTimeZone returns a boolean if a field has been set.
func (o *UeCampingRep) HasUeTimeZone() bool {
	if o != nil && !isNil(o.UeTimeZone) {
		return true
	}

	return false
}

// SetUeTimeZone gets a reference to the given string and assigns it to the UeTimeZone field.
func (o *UeCampingRep) SetUeTimeZone(v string) {
	o.UeTimeZone = &v
}

// GetNetLocAccSupp returns the NetLocAccSupp field value if set, zero value otherwise.
func (o *UeCampingRep) GetNetLocAccSupp() NetLocAccessSupport {
	if o == nil || isNil(o.NetLocAccSupp) {
		var ret NetLocAccessSupport
		return ret
	}
	return *o.NetLocAccSupp
}

// GetNetLocAccSuppOk returns a tuple with the NetLocAccSupp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCampingRep) GetNetLocAccSuppOk() (*NetLocAccessSupport, bool) {
	if o == nil || isNil(o.NetLocAccSupp) {
    return nil, false
	}
	return o.NetLocAccSupp, true
}

// HasNetLocAccSupp returns a boolean if a field has been set.
func (o *UeCampingRep) HasNetLocAccSupp() bool {
	if o != nil && !isNil(o.NetLocAccSupp) {
		return true
	}

	return false
}

// SetNetLocAccSupp gets a reference to the given NetLocAccessSupport and assigns it to the NetLocAccSupp field.
func (o *UeCampingRep) SetNetLocAccSupp(v NetLocAccessSupport) {
	o.NetLocAccSupp = &v
}

func (o UeCampingRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !isNil(o.ServNfId) {
		toSerialize["servNfId"] = o.ServNfId
	}
	if !isNil(o.ServingNetwork) {
		toSerialize["servingNetwork"] = o.ServingNetwork
	}
	if !isNil(o.UserLocationInfo) {
		toSerialize["userLocationInfo"] = o.UserLocationInfo
	}
	if !isNil(o.UeTimeZone) {
		toSerialize["ueTimeZone"] = o.UeTimeZone
	}
	if !isNil(o.NetLocAccSupp) {
		toSerialize["netLocAccSupp"] = o.NetLocAccSupp
	}
	return json.Marshal(toSerialize)
}

type NullableUeCampingRep struct {
	value *UeCampingRep
	isSet bool
}

func (v NullableUeCampingRep) Get() *UeCampingRep {
	return v.value
}

func (v *NullableUeCampingRep) Set(val *UeCampingRep) {
	v.value = val
	v.isSet = true
}

func (v NullableUeCampingRep) IsSet() bool {
	return v.isSet
}

func (v *NullableUeCampingRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeCampingRep(val *UeCampingRep) *NullableUeCampingRep {
	return &NullableUeCampingRep{value: val, isSet: true}
}

func (v NullableUeCampingRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeCampingRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


