/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// AusfInfo Information of an AUSF NF Instance
type AusfInfo struct {
	// Identifier of a group of NFs.
	GroupId *string `json:"groupId,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	RoutingIndicators []string `json:"routingIndicators,omitempty"`
	SuciInfos []SuciInfo `json:"suciInfos,omitempty"`
}

// NewAusfInfo instantiates a new AusfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAusfInfo() *AusfInfo {
	this := AusfInfo{}
	return &this
}

// NewAusfInfoWithDefaults instantiates a new AusfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAusfInfoWithDefaults() *AusfInfo {
	this := AusfInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AusfInfo) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AusfInfo) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AusfInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *AusfInfo) GetSupiRanges() []SupiRange {
	if o == nil || isNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || isNil(o.SupiRanges) {
    return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *AusfInfo) HasSupiRanges() bool {
	if o != nil && !isNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *AusfInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetRoutingIndicators returns the RoutingIndicators field value if set, zero value otherwise.
func (o *AusfInfo) GetRoutingIndicators() []string {
	if o == nil || isNil(o.RoutingIndicators) {
		var ret []string
		return ret
	}
	return o.RoutingIndicators
}

// GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfInfo) GetRoutingIndicatorsOk() ([]string, bool) {
	if o == nil || isNil(o.RoutingIndicators) {
    return nil, false
	}
	return o.RoutingIndicators, true
}

// HasRoutingIndicators returns a boolean if a field has been set.
func (o *AusfInfo) HasRoutingIndicators() bool {
	if o != nil && !isNil(o.RoutingIndicators) {
		return true
	}

	return false
}

// SetRoutingIndicators gets a reference to the given []string and assigns it to the RoutingIndicators field.
func (o *AusfInfo) SetRoutingIndicators(v []string) {
	o.RoutingIndicators = v
}

// GetSuciInfos returns the SuciInfos field value if set, zero value otherwise.
func (o *AusfInfo) GetSuciInfos() []SuciInfo {
	if o == nil || isNil(o.SuciInfos) {
		var ret []SuciInfo
		return ret
	}
	return o.SuciInfos
}

// GetSuciInfosOk returns a tuple with the SuciInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfInfo) GetSuciInfosOk() ([]SuciInfo, bool) {
	if o == nil || isNil(o.SuciInfos) {
    return nil, false
	}
	return o.SuciInfos, true
}

// HasSuciInfos returns a boolean if a field has been set.
func (o *AusfInfo) HasSuciInfos() bool {
	if o != nil && !isNil(o.SuciInfos) {
		return true
	}

	return false
}

// SetSuciInfos gets a reference to the given []SuciInfo and assigns it to the SuciInfos field.
func (o *AusfInfo) SetSuciInfos(v []SuciInfo) {
	o.SuciInfos = v
}

func (o AusfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !isNil(o.RoutingIndicators) {
		toSerialize["routingIndicators"] = o.RoutingIndicators
	}
	if !isNil(o.SuciInfos) {
		toSerialize["suciInfos"] = o.SuciInfos
	}
	return json.Marshal(toSerialize)
}

type NullableAusfInfo struct {
	value *AusfInfo
	isSet bool
}

func (v NullableAusfInfo) Get() *AusfInfo {
	return v.value
}

func (v *NullableAusfInfo) Set(val *AusfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAusfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAusfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAusfInfo(val *AusfInfo) *NullableAusfInfo {
	return &NullableAusfInfo{value: val, isSet: true}
}

func (v NullableAusfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAusfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


