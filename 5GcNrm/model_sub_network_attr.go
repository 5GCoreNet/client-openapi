/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// SubNetworkAttr struct for SubNetworkAttr
type SubNetworkAttr struct {
	DnPrefix *string `json:"dnPrefix,omitempty"`
	UserLabel *string `json:"userLabel,omitempty"`
	UserDefinedNetworkType *string `json:"userDefinedNetworkType,omitempty"`
	SetOfMcc []string `json:"setOfMcc,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
}

// NewSubNetworkAttr instantiates a new SubNetworkAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkAttr() *SubNetworkAttr {
	this := SubNetworkAttr{}
	return &this
}

// NewSubNetworkAttrWithDefaults instantiates a new SubNetworkAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkAttrWithDefaults() *SubNetworkAttr {
	this := SubNetworkAttr{}
	return &this
}

// GetDnPrefix returns the DnPrefix field value if set, zero value otherwise.
func (o *SubNetworkAttr) GetDnPrefix() string {
	if o == nil || isNil(o.DnPrefix) {
		var ret string
		return ret
	}
	return *o.DnPrefix
}

// GetDnPrefixOk returns a tuple with the DnPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkAttr) GetDnPrefixOk() (*string, bool) {
	if o == nil || isNil(o.DnPrefix) {
    return nil, false
	}
	return o.DnPrefix, true
}

// HasDnPrefix returns a boolean if a field has been set.
func (o *SubNetworkAttr) HasDnPrefix() bool {
	if o != nil && !isNil(o.DnPrefix) {
		return true
	}

	return false
}

// SetDnPrefix gets a reference to the given string and assigns it to the DnPrefix field.
func (o *SubNetworkAttr) SetDnPrefix(v string) {
	o.DnPrefix = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *SubNetworkAttr) GetUserLabel() string {
	if o == nil || isNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkAttr) GetUserLabelOk() (*string, bool) {
	if o == nil || isNil(o.UserLabel) {
    return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *SubNetworkAttr) HasUserLabel() bool {
	if o != nil && !isNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *SubNetworkAttr) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetUserDefinedNetworkType returns the UserDefinedNetworkType field value if set, zero value otherwise.
func (o *SubNetworkAttr) GetUserDefinedNetworkType() string {
	if o == nil || isNil(o.UserDefinedNetworkType) {
		var ret string
		return ret
	}
	return *o.UserDefinedNetworkType
}

// GetUserDefinedNetworkTypeOk returns a tuple with the UserDefinedNetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkAttr) GetUserDefinedNetworkTypeOk() (*string, bool) {
	if o == nil || isNil(o.UserDefinedNetworkType) {
    return nil, false
	}
	return o.UserDefinedNetworkType, true
}

// HasUserDefinedNetworkType returns a boolean if a field has been set.
func (o *SubNetworkAttr) HasUserDefinedNetworkType() bool {
	if o != nil && !isNil(o.UserDefinedNetworkType) {
		return true
	}

	return false
}

// SetUserDefinedNetworkType gets a reference to the given string and assigns it to the UserDefinedNetworkType field.
func (o *SubNetworkAttr) SetUserDefinedNetworkType(v string) {
	o.UserDefinedNetworkType = &v
}

// GetSetOfMcc returns the SetOfMcc field value if set, zero value otherwise.
func (o *SubNetworkAttr) GetSetOfMcc() []string {
	if o == nil || isNil(o.SetOfMcc) {
		var ret []string
		return ret
	}
	return o.SetOfMcc
}

// GetSetOfMccOk returns a tuple with the SetOfMcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkAttr) GetSetOfMccOk() ([]string, bool) {
	if o == nil || isNil(o.SetOfMcc) {
    return nil, false
	}
	return o.SetOfMcc, true
}

// HasSetOfMcc returns a boolean if a field has been set.
func (o *SubNetworkAttr) HasSetOfMcc() bool {
	if o != nil && !isNil(o.SetOfMcc) {
		return true
	}

	return false
}

// SetSetOfMcc gets a reference to the given []string and assigns it to the SetOfMcc field.
func (o *SubNetworkAttr) SetSetOfMcc(v []string) {
	o.SetOfMcc = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *SubNetworkAttr) GetPriorityLabel() int32 {
	if o == nil || isNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkAttr) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || isNil(o.PriorityLabel) {
    return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *SubNetworkAttr) HasPriorityLabel() bool {
	if o != nil && !isNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *SubNetworkAttr) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *SubNetworkAttr) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || isNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkAttr) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || isNil(o.SupportedPerfMetricGroups) {
    return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *SubNetworkAttr) HasSupportedPerfMetricGroups() bool {
	if o != nil && !isNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *SubNetworkAttr) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *SubNetworkAttr) GetSupportedTraceMetrics() []string {
	if o == nil || isNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkAttr) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedTraceMetrics) {
    return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *SubNetworkAttr) HasSupportedTraceMetrics() bool {
	if o != nil && !isNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *SubNetworkAttr) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

func (o SubNetworkAttr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DnPrefix) {
		toSerialize["dnPrefix"] = o.DnPrefix
	}
	if !isNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !isNil(o.UserDefinedNetworkType) {
		toSerialize["userDefinedNetworkType"] = o.UserDefinedNetworkType
	}
	if !isNil(o.SetOfMcc) {
		toSerialize["setOfMcc"] = o.SetOfMcc
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

type NullableSubNetworkAttr struct {
	value *SubNetworkAttr
	isSet bool
}

func (v NullableSubNetworkAttr) Get() *SubNetworkAttr {
	return v.value
}

func (v *NullableSubNetworkAttr) Set(val *SubNetworkAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkAttr(val *SubNetworkAttr) *NullableSubNetworkAttr {
	return &NullableSubNetworkAttr{value: val, isSet: true}
}

func (v NullableSubNetworkAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


