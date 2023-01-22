/*
Naf_EventExposure

AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Naf_EventExposure

import (
	"encoding/json"
)

// NetworkAreaInfo Describes a network area information in which the NF service consumer requests the number of UEs. 
type NetworkAreaInfo struct {
	// Contains a list of E-UTRA cell identities.
	Ecgis []Ecgi `json:"ecgis,omitempty"`
	// Contains a list of NR cell identities.
	Ncgis []Ncgi `json:"ncgis,omitempty"`
	// Contains a list of NG RAN nodes.
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`
	// Contains a list of tracking area identities.
	Tais []Tai `json:"tais,omitempty"`
}

// NewNetworkAreaInfo instantiates a new NetworkAreaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAreaInfo() *NetworkAreaInfo {
	this := NetworkAreaInfo{}
	return &this
}

// NewNetworkAreaInfoWithDefaults instantiates a new NetworkAreaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAreaInfoWithDefaults() *NetworkAreaInfo {
	this := NetworkAreaInfo{}
	return &this
}

// GetEcgis returns the Ecgis field value if set, zero value otherwise.
func (o *NetworkAreaInfo) GetEcgis() []Ecgi {
	if o == nil || isNil(o.Ecgis) {
		var ret []Ecgi
		return ret
	}
	return o.Ecgis
}

// GetEcgisOk returns a tuple with the Ecgis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo) GetEcgisOk() ([]Ecgi, bool) {
	if o == nil || isNil(o.Ecgis) {
    return nil, false
	}
	return o.Ecgis, true
}

// HasEcgis returns a boolean if a field has been set.
func (o *NetworkAreaInfo) HasEcgis() bool {
	if o != nil && !isNil(o.Ecgis) {
		return true
	}

	return false
}

// SetEcgis gets a reference to the given []Ecgi and assigns it to the Ecgis field.
func (o *NetworkAreaInfo) SetEcgis(v []Ecgi) {
	o.Ecgis = v
}

// GetNcgis returns the Ncgis field value if set, zero value otherwise.
func (o *NetworkAreaInfo) GetNcgis() []Ncgi {
	if o == nil || isNil(o.Ncgis) {
		var ret []Ncgi
		return ret
	}
	return o.Ncgis
}

// GetNcgisOk returns a tuple with the Ncgis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo) GetNcgisOk() ([]Ncgi, bool) {
	if o == nil || isNil(o.Ncgis) {
    return nil, false
	}
	return o.Ncgis, true
}

// HasNcgis returns a boolean if a field has been set.
func (o *NetworkAreaInfo) HasNcgis() bool {
	if o != nil && !isNil(o.Ncgis) {
		return true
	}

	return false
}

// SetNcgis gets a reference to the given []Ncgi and assigns it to the Ncgis field.
func (o *NetworkAreaInfo) SetNcgis(v []Ncgi) {
	o.Ncgis = v
}

// GetGRanNodeIds returns the GRanNodeIds field value if set, zero value otherwise.
func (o *NetworkAreaInfo) GetGRanNodeIds() []GlobalRanNodeId {
	if o == nil || isNil(o.GRanNodeIds) {
		var ret []GlobalRanNodeId
		return ret
	}
	return o.GRanNodeIds
}

// GetGRanNodeIdsOk returns a tuple with the GRanNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo) GetGRanNodeIdsOk() ([]GlobalRanNodeId, bool) {
	if o == nil || isNil(o.GRanNodeIds) {
    return nil, false
	}
	return o.GRanNodeIds, true
}

// HasGRanNodeIds returns a boolean if a field has been set.
func (o *NetworkAreaInfo) HasGRanNodeIds() bool {
	if o != nil && !isNil(o.GRanNodeIds) {
		return true
	}

	return false
}

// SetGRanNodeIds gets a reference to the given []GlobalRanNodeId and assigns it to the GRanNodeIds field.
func (o *NetworkAreaInfo) SetGRanNodeIds(v []GlobalRanNodeId) {
	o.GRanNodeIds = v
}

// GetTais returns the Tais field value if set, zero value otherwise.
func (o *NetworkAreaInfo) GetTais() []Tai {
	if o == nil || isNil(o.Tais) {
		var ret []Tai
		return ret
	}
	return o.Tais
}

// GetTaisOk returns a tuple with the Tais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAreaInfo) GetTaisOk() ([]Tai, bool) {
	if o == nil || isNil(o.Tais) {
    return nil, false
	}
	return o.Tais, true
}

// HasTais returns a boolean if a field has been set.
func (o *NetworkAreaInfo) HasTais() bool {
	if o != nil && !isNil(o.Tais) {
		return true
	}

	return false
}

// SetTais gets a reference to the given []Tai and assigns it to the Tais field.
func (o *NetworkAreaInfo) SetTais(v []Tai) {
	o.Tais = v
}

func (o NetworkAreaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ecgis) {
		toSerialize["ecgis"] = o.Ecgis
	}
	if !isNil(o.Ncgis) {
		toSerialize["ncgis"] = o.Ncgis
	}
	if !isNil(o.GRanNodeIds) {
		toSerialize["gRanNodeIds"] = o.GRanNodeIds
	}
	if !isNil(o.Tais) {
		toSerialize["tais"] = o.Tais
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAreaInfo struct {
	value *NetworkAreaInfo
	isSet bool
}

func (v NullableNetworkAreaInfo) Get() *NetworkAreaInfo {
	return v.value
}

func (v *NullableNetworkAreaInfo) Set(val *NetworkAreaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAreaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAreaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAreaInfo(val *NetworkAreaInfo) *NullableNetworkAreaInfo {
	return &NullableNetworkAreaInfo{value: val, isSet: true}
}

func (v NullableNetworkAreaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAreaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


