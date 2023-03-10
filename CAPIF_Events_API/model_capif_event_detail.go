/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Events_API

import (
	"encoding/json"
)

// CAPIFEventDetail Represents a CAPIF event details.
type CAPIFEventDetail struct {
	// Description of the service API as published by the APF.
	ServiceAPIDescriptions []ServiceAPIDescription `json:"serviceAPIDescriptions,omitempty"`
	// Identifier of the service API
	ApiIds []string `json:"apiIds,omitempty"`
	// Identity of the API invoker
	ApiInvokerIds []string `json:"apiInvokerIds,omitempty"`
	AccCtrlPolList *AccessControlPolicyListExt `json:"accCtrlPolList,omitempty"`
	// Invocation logs.
	InvocationLogs []InvocationLog `json:"invocationLogs,omitempty"`
	ApiTopoHide *TopologyHiding `json:"apiTopoHide,omitempty"`
}

// NewCAPIFEventDetail instantiates a new CAPIFEventDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCAPIFEventDetail() *CAPIFEventDetail {
	this := CAPIFEventDetail{}
	return &this
}

// NewCAPIFEventDetailWithDefaults instantiates a new CAPIFEventDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCAPIFEventDetailWithDefaults() *CAPIFEventDetail {
	this := CAPIFEventDetail{}
	return &this
}

// GetServiceAPIDescriptions returns the ServiceAPIDescriptions field value if set, zero value otherwise.
func (o *CAPIFEventDetail) GetServiceAPIDescriptions() []ServiceAPIDescription {
	if o == nil || isNil(o.ServiceAPIDescriptions) {
		var ret []ServiceAPIDescription
		return ret
	}
	return o.ServiceAPIDescriptions
}

// GetServiceAPIDescriptionsOk returns a tuple with the ServiceAPIDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPIFEventDetail) GetServiceAPIDescriptionsOk() ([]ServiceAPIDescription, bool) {
	if o == nil || isNil(o.ServiceAPIDescriptions) {
    return nil, false
	}
	return o.ServiceAPIDescriptions, true
}

// HasServiceAPIDescriptions returns a boolean if a field has been set.
func (o *CAPIFEventDetail) HasServiceAPIDescriptions() bool {
	if o != nil && !isNil(o.ServiceAPIDescriptions) {
		return true
	}

	return false
}

// SetServiceAPIDescriptions gets a reference to the given []ServiceAPIDescription and assigns it to the ServiceAPIDescriptions field.
func (o *CAPIFEventDetail) SetServiceAPIDescriptions(v []ServiceAPIDescription) {
	o.ServiceAPIDescriptions = v
}

// GetApiIds returns the ApiIds field value if set, zero value otherwise.
func (o *CAPIFEventDetail) GetApiIds() []string {
	if o == nil || isNil(o.ApiIds) {
		var ret []string
		return ret
	}
	return o.ApiIds
}

// GetApiIdsOk returns a tuple with the ApiIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPIFEventDetail) GetApiIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ApiIds) {
    return nil, false
	}
	return o.ApiIds, true
}

// HasApiIds returns a boolean if a field has been set.
func (o *CAPIFEventDetail) HasApiIds() bool {
	if o != nil && !isNil(o.ApiIds) {
		return true
	}

	return false
}

// SetApiIds gets a reference to the given []string and assigns it to the ApiIds field.
func (o *CAPIFEventDetail) SetApiIds(v []string) {
	o.ApiIds = v
}

// GetApiInvokerIds returns the ApiInvokerIds field value if set, zero value otherwise.
func (o *CAPIFEventDetail) GetApiInvokerIds() []string {
	if o == nil || isNil(o.ApiInvokerIds) {
		var ret []string
		return ret
	}
	return o.ApiInvokerIds
}

// GetApiInvokerIdsOk returns a tuple with the ApiInvokerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPIFEventDetail) GetApiInvokerIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ApiInvokerIds) {
    return nil, false
	}
	return o.ApiInvokerIds, true
}

// HasApiInvokerIds returns a boolean if a field has been set.
func (o *CAPIFEventDetail) HasApiInvokerIds() bool {
	if o != nil && !isNil(o.ApiInvokerIds) {
		return true
	}

	return false
}

// SetApiInvokerIds gets a reference to the given []string and assigns it to the ApiInvokerIds field.
func (o *CAPIFEventDetail) SetApiInvokerIds(v []string) {
	o.ApiInvokerIds = v
}

// GetAccCtrlPolList returns the AccCtrlPolList field value if set, zero value otherwise.
func (o *CAPIFEventDetail) GetAccCtrlPolList() AccessControlPolicyListExt {
	if o == nil || isNil(o.AccCtrlPolList) {
		var ret AccessControlPolicyListExt
		return ret
	}
	return *o.AccCtrlPolList
}

// GetAccCtrlPolListOk returns a tuple with the AccCtrlPolList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPIFEventDetail) GetAccCtrlPolListOk() (*AccessControlPolicyListExt, bool) {
	if o == nil || isNil(o.AccCtrlPolList) {
    return nil, false
	}
	return o.AccCtrlPolList, true
}

// HasAccCtrlPolList returns a boolean if a field has been set.
func (o *CAPIFEventDetail) HasAccCtrlPolList() bool {
	if o != nil && !isNil(o.AccCtrlPolList) {
		return true
	}

	return false
}

// SetAccCtrlPolList gets a reference to the given AccessControlPolicyListExt and assigns it to the AccCtrlPolList field.
func (o *CAPIFEventDetail) SetAccCtrlPolList(v AccessControlPolicyListExt) {
	o.AccCtrlPolList = &v
}

// GetInvocationLogs returns the InvocationLogs field value if set, zero value otherwise.
func (o *CAPIFEventDetail) GetInvocationLogs() []InvocationLog {
	if o == nil || isNil(o.InvocationLogs) {
		var ret []InvocationLog
		return ret
	}
	return o.InvocationLogs
}

// GetInvocationLogsOk returns a tuple with the InvocationLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPIFEventDetail) GetInvocationLogsOk() ([]InvocationLog, bool) {
	if o == nil || isNil(o.InvocationLogs) {
    return nil, false
	}
	return o.InvocationLogs, true
}

// HasInvocationLogs returns a boolean if a field has been set.
func (o *CAPIFEventDetail) HasInvocationLogs() bool {
	if o != nil && !isNil(o.InvocationLogs) {
		return true
	}

	return false
}

// SetInvocationLogs gets a reference to the given []InvocationLog and assigns it to the InvocationLogs field.
func (o *CAPIFEventDetail) SetInvocationLogs(v []InvocationLog) {
	o.InvocationLogs = v
}

// GetApiTopoHide returns the ApiTopoHide field value if set, zero value otherwise.
func (o *CAPIFEventDetail) GetApiTopoHide() TopologyHiding {
	if o == nil || isNil(o.ApiTopoHide) {
		var ret TopologyHiding
		return ret
	}
	return *o.ApiTopoHide
}

// GetApiTopoHideOk returns a tuple with the ApiTopoHide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPIFEventDetail) GetApiTopoHideOk() (*TopologyHiding, bool) {
	if o == nil || isNil(o.ApiTopoHide) {
    return nil, false
	}
	return o.ApiTopoHide, true
}

// HasApiTopoHide returns a boolean if a field has been set.
func (o *CAPIFEventDetail) HasApiTopoHide() bool {
	if o != nil && !isNil(o.ApiTopoHide) {
		return true
	}

	return false
}

// SetApiTopoHide gets a reference to the given TopologyHiding and assigns it to the ApiTopoHide field.
func (o *CAPIFEventDetail) SetApiTopoHide(v TopologyHiding) {
	o.ApiTopoHide = &v
}

func (o CAPIFEventDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceAPIDescriptions) {
		toSerialize["serviceAPIDescriptions"] = o.ServiceAPIDescriptions
	}
	if !isNil(o.ApiIds) {
		toSerialize["apiIds"] = o.ApiIds
	}
	if !isNil(o.ApiInvokerIds) {
		toSerialize["apiInvokerIds"] = o.ApiInvokerIds
	}
	if !isNil(o.AccCtrlPolList) {
		toSerialize["accCtrlPolList"] = o.AccCtrlPolList
	}
	if !isNil(o.InvocationLogs) {
		toSerialize["invocationLogs"] = o.InvocationLogs
	}
	if !isNil(o.ApiTopoHide) {
		toSerialize["apiTopoHide"] = o.ApiTopoHide
	}
	return json.Marshal(toSerialize)
}

type NullableCAPIFEventDetail struct {
	value *CAPIFEventDetail
	isSet bool
}

func (v NullableCAPIFEventDetail) Get() *CAPIFEventDetail {
	return v.value
}

func (v *NullableCAPIFEventDetail) Set(val *CAPIFEventDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCAPIFEventDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCAPIFEventDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCAPIFEventDetail(val *CAPIFEventDetail) *NullableCAPIFEventDetail {
	return &NullableCAPIFEventDetail{value: val, isSet: true}
}

func (v NullableCAPIFEventDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCAPIFEventDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


