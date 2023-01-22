/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package EdgeNrm

import (
	"encoding/json"
)

// ManagedFunctionAttr struct for ManagedFunctionAttr
type ManagedFunctionAttr struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
}

// NewManagedFunctionAttr instantiates a new ManagedFunctionAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedFunctionAttr() *ManagedFunctionAttr {
	this := ManagedFunctionAttr{}
	return &this
}

// NewManagedFunctionAttrWithDefaults instantiates a new ManagedFunctionAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedFunctionAttrWithDefaults() *ManagedFunctionAttr {
	this := ManagedFunctionAttr{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ManagedFunctionAttr) GetUserLabel() string {
	if o == nil || isNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedFunctionAttr) GetUserLabelOk() (*string, bool) {
	if o == nil || isNil(o.UserLabel) {
    return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ManagedFunctionAttr) HasUserLabel() bool {
	if o != nil && !isNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ManagedFunctionAttr) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *ManagedFunctionAttr) GetVnfParametersList() []VnfParameter {
	if o == nil || isNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedFunctionAttr) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || isNil(o.VnfParametersList) {
    return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *ManagedFunctionAttr) HasVnfParametersList() bool {
	if o != nil && !isNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *ManagedFunctionAttr) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *ManagedFunctionAttr) GetPeeParametersList() []PeeParameter {
	if o == nil || isNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedFunctionAttr) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || isNil(o.PeeParametersList) {
    return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *ManagedFunctionAttr) HasPeeParametersList() bool {
	if o != nil && !isNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *ManagedFunctionAttr) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *ManagedFunctionAttr) GetPriorityLabel() int32 {
	if o == nil || isNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedFunctionAttr) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || isNil(o.PriorityLabel) {
    return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *ManagedFunctionAttr) HasPriorityLabel() bool {
	if o != nil && !isNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *ManagedFunctionAttr) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *ManagedFunctionAttr) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || isNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedFunctionAttr) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || isNil(o.SupportedPerfMetricGroups) {
    return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *ManagedFunctionAttr) HasSupportedPerfMetricGroups() bool {
	if o != nil && !isNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *ManagedFunctionAttr) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *ManagedFunctionAttr) GetSupportedTraceMetrics() []string {
	if o == nil || isNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedFunctionAttr) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedTraceMetrics) {
    return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *ManagedFunctionAttr) HasSupportedTraceMetrics() bool {
	if o != nil && !isNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *ManagedFunctionAttr) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

func (o ManagedFunctionAttr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !isNil(o.VnfParametersList) {
		toSerialize["vnfParametersList"] = o.VnfParametersList
	}
	if !isNil(o.PeeParametersList) {
		toSerialize["peeParametersList"] = o.PeeParametersList
	}
	if !isNil(o.PriorityLabel) {
		toSerialize["priorityLabel"] = o.PriorityLabel
	}
	if !isNil(o.SupportedPerfMetricGroups) {
		toSerialize["supportedPerfMetricGroups"] = o.SupportedPerfMetricGroups
	}
	if !isNil(o.SupportedTraceMetrics) {
		toSerialize["supportedTraceMetrics"] = o.SupportedTraceMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableManagedFunctionAttr struct {
	value *ManagedFunctionAttr
	isSet bool
}

func (v NullableManagedFunctionAttr) Get() *ManagedFunctionAttr {
	return v.value
}

func (v *NullableManagedFunctionAttr) Set(val *ManagedFunctionAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedFunctionAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedFunctionAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedFunctionAttr(val *ManagedFunctionAttr) *NullableManagedFunctionAttr {
	return &NullableManagedFunctionAttr{value: val, isSet: true}
}

func (v NullableManagedFunctionAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedFunctionAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


