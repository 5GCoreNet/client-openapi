/*
Naf_EventExposure

AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_EventExposure

import (
	"encoding/json"
)

// ServiceExperienceInfoPerFlow Contains service experience information associated with a service flow.
type ServiceExperienceInfoPerFlow struct {
	SvcExprc *SvcExperience `json:"svcExprc,omitempty"`
	TimeIntev *TimeWindow `json:"timeIntev,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai *string `json:"dnai,omitempty"`
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty"`
	EthTrafficFilter *EthFlowDescription `json:"ethTrafficFilter,omitempty"`
}

// NewServiceExperienceInfoPerFlow instantiates a new ServiceExperienceInfoPerFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceExperienceInfoPerFlow() *ServiceExperienceInfoPerFlow {
	this := ServiceExperienceInfoPerFlow{}
	return &this
}

// NewServiceExperienceInfoPerFlowWithDefaults instantiates a new ServiceExperienceInfoPerFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceExperienceInfoPerFlowWithDefaults() *ServiceExperienceInfoPerFlow {
	this := ServiceExperienceInfoPerFlow{}
	return &this
}

// GetSvcExprc returns the SvcExprc field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerFlow) GetSvcExprc() SvcExperience {
	if o == nil || isNil(o.SvcExprc) {
		var ret SvcExperience
		return ret
	}
	return *o.SvcExprc
}

// GetSvcExprcOk returns a tuple with the SvcExprc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerFlow) GetSvcExprcOk() (*SvcExperience, bool) {
	if o == nil || isNil(o.SvcExprc) {
    return nil, false
	}
	return o.SvcExprc, true
}

// HasSvcExprc returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerFlow) HasSvcExprc() bool {
	if o != nil && !isNil(o.SvcExprc) {
		return true
	}

	return false
}

// SetSvcExprc gets a reference to the given SvcExperience and assigns it to the SvcExprc field.
func (o *ServiceExperienceInfoPerFlow) SetSvcExprc(v SvcExperience) {
	o.SvcExprc = &v
}

// GetTimeIntev returns the TimeIntev field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerFlow) GetTimeIntev() TimeWindow {
	if o == nil || isNil(o.TimeIntev) {
		var ret TimeWindow
		return ret
	}
	return *o.TimeIntev
}

// GetTimeIntevOk returns a tuple with the TimeIntev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerFlow) GetTimeIntevOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.TimeIntev) {
    return nil, false
	}
	return o.TimeIntev, true
}

// HasTimeIntev returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerFlow) HasTimeIntev() bool {
	if o != nil && !isNil(o.TimeIntev) {
		return true
	}

	return false
}

// SetTimeIntev gets a reference to the given TimeWindow and assigns it to the TimeIntev field.
func (o *ServiceExperienceInfoPerFlow) SetTimeIntev(v TimeWindow) {
	o.TimeIntev = &v
}

// GetDnai returns the Dnai field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerFlow) GetDnai() string {
	if o == nil || isNil(o.Dnai) {
		var ret string
		return ret
	}
	return *o.Dnai
}

// GetDnaiOk returns a tuple with the Dnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerFlow) GetDnaiOk() (*string, bool) {
	if o == nil || isNil(o.Dnai) {
    return nil, false
	}
	return o.Dnai, true
}

// HasDnai returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerFlow) HasDnai() bool {
	if o != nil && !isNil(o.Dnai) {
		return true
	}

	return false
}

// SetDnai gets a reference to the given string and assigns it to the Dnai field.
func (o *ServiceExperienceInfoPerFlow) SetDnai(v string) {
	o.Dnai = &v
}

// GetIpTrafficFilter returns the IpTrafficFilter field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerFlow) GetIpTrafficFilter() FlowInfo {
	if o == nil || isNil(o.IpTrafficFilter) {
		var ret FlowInfo
		return ret
	}
	return *o.IpTrafficFilter
}

// GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerFlow) GetIpTrafficFilterOk() (*FlowInfo, bool) {
	if o == nil || isNil(o.IpTrafficFilter) {
    return nil, false
	}
	return o.IpTrafficFilter, true
}

// HasIpTrafficFilter returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerFlow) HasIpTrafficFilter() bool {
	if o != nil && !isNil(o.IpTrafficFilter) {
		return true
	}

	return false
}

// SetIpTrafficFilter gets a reference to the given FlowInfo and assigns it to the IpTrafficFilter field.
func (o *ServiceExperienceInfoPerFlow) SetIpTrafficFilter(v FlowInfo) {
	o.IpTrafficFilter = &v
}

// GetEthTrafficFilter returns the EthTrafficFilter field value if set, zero value otherwise.
func (o *ServiceExperienceInfoPerFlow) GetEthTrafficFilter() EthFlowDescription {
	if o == nil || isNil(o.EthTrafficFilter) {
		var ret EthFlowDescription
		return ret
	}
	return *o.EthTrafficFilter
}

// GetEthTrafficFilterOk returns a tuple with the EthTrafficFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfoPerFlow) GetEthTrafficFilterOk() (*EthFlowDescription, bool) {
	if o == nil || isNil(o.EthTrafficFilter) {
    return nil, false
	}
	return o.EthTrafficFilter, true
}

// HasEthTrafficFilter returns a boolean if a field has been set.
func (o *ServiceExperienceInfoPerFlow) HasEthTrafficFilter() bool {
	if o != nil && !isNil(o.EthTrafficFilter) {
		return true
	}

	return false
}

// SetEthTrafficFilter gets a reference to the given EthFlowDescription and assigns it to the EthTrafficFilter field.
func (o *ServiceExperienceInfoPerFlow) SetEthTrafficFilter(v EthFlowDescription) {
	o.EthTrafficFilter = &v
}

func (o ServiceExperienceInfoPerFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SvcExprc) {
		toSerialize["svcExprc"] = o.SvcExprc
	}
	if !isNil(o.TimeIntev) {
		toSerialize["timeIntev"] = o.TimeIntev
	}
	if !isNil(o.Dnai) {
		toSerialize["dnai"] = o.Dnai
	}
	if !isNil(o.IpTrafficFilter) {
		toSerialize["ipTrafficFilter"] = o.IpTrafficFilter
	}
	if !isNil(o.EthTrafficFilter) {
		toSerialize["ethTrafficFilter"] = o.EthTrafficFilter
	}
	return json.Marshal(toSerialize)
}

type NullableServiceExperienceInfoPerFlow struct {
	value *ServiceExperienceInfoPerFlow
	isSet bool
}

func (v NullableServiceExperienceInfoPerFlow) Get() *ServiceExperienceInfoPerFlow {
	return v.value
}

func (v *NullableServiceExperienceInfoPerFlow) Set(val *ServiceExperienceInfoPerFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceInfoPerFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceInfoPerFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceInfoPerFlow(val *ServiceExperienceInfoPerFlow) *NullableServiceExperienceInfoPerFlow {
	return &NullableServiceExperienceInfoPerFlow{value: val, isSet: true}
}

func (v NullableServiceExperienceInfoPerFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceInfoPerFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


