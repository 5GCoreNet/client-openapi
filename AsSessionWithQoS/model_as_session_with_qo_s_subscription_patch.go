/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AsSessionWithQoS

import (
	"encoding/json"
)

// AsSessionWithQoSSubscriptionPatch Represents parameters to modify an AS session with specific QoS subscription.
type AsSessionWithQoSSubscriptionPatch struct {
	// Identifies the external Application Identifier.
	ExterAppId *string `json:"exterAppId,omitempty"`
	// Describe the IP data flow which requires QoS.
	FlowInfo []FlowInfo `json:"flowInfo,omitempty"`
	// Identifies Ethernet packet flows.
	EthFlowInfo []EthFlowDescription `json:"ethFlowInfo,omitempty"`
	// Identifies the Ethernet flows which require QoS. Each Ethernet flow consists of a flow idenifer and the corresponding UL and/or DL flows. 
	EnEthFlowInfo []EthFlowInfo `json:"enEthFlowInfo,omitempty"`
	// Pre-defined QoS reference
	QosReference *string `json:"qosReference,omitempty"`
	// Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority.
	AltQoSReferences []string `json:"altQoSReferences,omitempty"`
	// Identifies an ordered list of alternative service requirements that include individual QoS parameter sets. The lower the index of the array for a given entry, the higher the priority.
	AltQosReqs []AlternativeServiceRequirementsData `json:"altQosReqs,omitempty"`
	// Indicates whether the QoS flow parameters signalling to the UE when the SMF is notified by the NG-RAN of changes in the fulfilled QoS situation is disabled (true) or not (false). The fulfilled situation is either the QoS profile or an Alternative QoS Profile. 
	DisUeNotif *bool `json:"disUeNotif,omitempty"`
	UsageThreshold NullableUsageThresholdRm `json:"usageThreshold,omitempty"`
	QosMonInfo *QosMonitoringInformationRm `json:"qosMonInfo,omitempty"`
	// Indicates whether the direct event notification is requested (true) or not (false). 
	DirectNotifInd *bool `json:"directNotifInd,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	TscQosReq *TscQosRequirementRm `json:"tscQosReq,omitempty"`
	// Represents the updated list of user plane event(s) to which the SCS/AS requests to subscribe to.
	Events []UserPlaneEvent `json:"events,omitempty"`
}

// NewAsSessionWithQoSSubscriptionPatch instantiates a new AsSessionWithQoSSubscriptionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsSessionWithQoSSubscriptionPatch() *AsSessionWithQoSSubscriptionPatch {
	this := AsSessionWithQoSSubscriptionPatch{}
	return &this
}

// NewAsSessionWithQoSSubscriptionPatchWithDefaults instantiates a new AsSessionWithQoSSubscriptionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsSessionWithQoSSubscriptionPatchWithDefaults() *AsSessionWithQoSSubscriptionPatch {
	this := AsSessionWithQoSSubscriptionPatch{}
	return &this
}

// GetExterAppId returns the ExterAppId field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetExterAppId() string {
	if o == nil || isNil(o.ExterAppId) {
		var ret string
		return ret
	}
	return *o.ExterAppId
}

// GetExterAppIdOk returns a tuple with the ExterAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetExterAppIdOk() (*string, bool) {
	if o == nil || isNil(o.ExterAppId) {
    return nil, false
	}
	return o.ExterAppId, true
}

// HasExterAppId returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasExterAppId() bool {
	if o != nil && !isNil(o.ExterAppId) {
		return true
	}

	return false
}

// SetExterAppId gets a reference to the given string and assigns it to the ExterAppId field.
func (o *AsSessionWithQoSSubscriptionPatch) SetExterAppId(v string) {
	o.ExterAppId = &v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetFlowInfo() []FlowInfo {
	if o == nil || isNil(o.FlowInfo) {
		var ret []FlowInfo
		return ret
	}
	return o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetFlowInfoOk() ([]FlowInfo, bool) {
	if o == nil || isNil(o.FlowInfo) {
    return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasFlowInfo() bool {
	if o != nil && !isNil(o.FlowInfo) {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given []FlowInfo and assigns it to the FlowInfo field.
func (o *AsSessionWithQoSSubscriptionPatch) SetFlowInfo(v []FlowInfo) {
	o.FlowInfo = v
}

// GetEthFlowInfo returns the EthFlowInfo field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetEthFlowInfo() []EthFlowDescription {
	if o == nil || isNil(o.EthFlowInfo) {
		var ret []EthFlowDescription
		return ret
	}
	return o.EthFlowInfo
}

// GetEthFlowInfoOk returns a tuple with the EthFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetEthFlowInfoOk() ([]EthFlowDescription, bool) {
	if o == nil || isNil(o.EthFlowInfo) {
    return nil, false
	}
	return o.EthFlowInfo, true
}

// HasEthFlowInfo returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasEthFlowInfo() bool {
	if o != nil && !isNil(o.EthFlowInfo) {
		return true
	}

	return false
}

// SetEthFlowInfo gets a reference to the given []EthFlowDescription and assigns it to the EthFlowInfo field.
func (o *AsSessionWithQoSSubscriptionPatch) SetEthFlowInfo(v []EthFlowDescription) {
	o.EthFlowInfo = v
}

// GetEnEthFlowInfo returns the EnEthFlowInfo field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetEnEthFlowInfo() []EthFlowInfo {
	if o == nil || isNil(o.EnEthFlowInfo) {
		var ret []EthFlowInfo
		return ret
	}
	return o.EnEthFlowInfo
}

// GetEnEthFlowInfoOk returns a tuple with the EnEthFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetEnEthFlowInfoOk() ([]EthFlowInfo, bool) {
	if o == nil || isNil(o.EnEthFlowInfo) {
    return nil, false
	}
	return o.EnEthFlowInfo, true
}

// HasEnEthFlowInfo returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasEnEthFlowInfo() bool {
	if o != nil && !isNil(o.EnEthFlowInfo) {
		return true
	}

	return false
}

// SetEnEthFlowInfo gets a reference to the given []EthFlowInfo and assigns it to the EnEthFlowInfo field.
func (o *AsSessionWithQoSSubscriptionPatch) SetEnEthFlowInfo(v []EthFlowInfo) {
	o.EnEthFlowInfo = v
}

// GetQosReference returns the QosReference field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetQosReference() string {
	if o == nil || isNil(o.QosReference) {
		var ret string
		return ret
	}
	return *o.QosReference
}

// GetQosReferenceOk returns a tuple with the QosReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetQosReferenceOk() (*string, bool) {
	if o == nil || isNil(o.QosReference) {
    return nil, false
	}
	return o.QosReference, true
}

// HasQosReference returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasQosReference() bool {
	if o != nil && !isNil(o.QosReference) {
		return true
	}

	return false
}

// SetQosReference gets a reference to the given string and assigns it to the QosReference field.
func (o *AsSessionWithQoSSubscriptionPatch) SetQosReference(v string) {
	o.QosReference = &v
}

// GetAltQoSReferences returns the AltQoSReferences field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetAltQoSReferences() []string {
	if o == nil || isNil(o.AltQoSReferences) {
		var ret []string
		return ret
	}
	return o.AltQoSReferences
}

// GetAltQoSReferencesOk returns a tuple with the AltQoSReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetAltQoSReferencesOk() ([]string, bool) {
	if o == nil || isNil(o.AltQoSReferences) {
    return nil, false
	}
	return o.AltQoSReferences, true
}

// HasAltQoSReferences returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasAltQoSReferences() bool {
	if o != nil && !isNil(o.AltQoSReferences) {
		return true
	}

	return false
}

// SetAltQoSReferences gets a reference to the given []string and assigns it to the AltQoSReferences field.
func (o *AsSessionWithQoSSubscriptionPatch) SetAltQoSReferences(v []string) {
	o.AltQoSReferences = v
}

// GetAltQosReqs returns the AltQosReqs field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetAltQosReqs() []AlternativeServiceRequirementsData {
	if o == nil || isNil(o.AltQosReqs) {
		var ret []AlternativeServiceRequirementsData
		return ret
	}
	return o.AltQosReqs
}

// GetAltQosReqsOk returns a tuple with the AltQosReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetAltQosReqsOk() ([]AlternativeServiceRequirementsData, bool) {
	if o == nil || isNil(o.AltQosReqs) {
    return nil, false
	}
	return o.AltQosReqs, true
}

// HasAltQosReqs returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasAltQosReqs() bool {
	if o != nil && !isNil(o.AltQosReqs) {
		return true
	}

	return false
}

// SetAltQosReqs gets a reference to the given []AlternativeServiceRequirementsData and assigns it to the AltQosReqs field.
func (o *AsSessionWithQoSSubscriptionPatch) SetAltQosReqs(v []AlternativeServiceRequirementsData) {
	o.AltQosReqs = v
}

// GetDisUeNotif returns the DisUeNotif field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetDisUeNotif() bool {
	if o == nil || isNil(o.DisUeNotif) {
		var ret bool
		return ret
	}
	return *o.DisUeNotif
}

// GetDisUeNotifOk returns a tuple with the DisUeNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetDisUeNotifOk() (*bool, bool) {
	if o == nil || isNil(o.DisUeNotif) {
    return nil, false
	}
	return o.DisUeNotif, true
}

// HasDisUeNotif returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasDisUeNotif() bool {
	if o != nil && !isNil(o.DisUeNotif) {
		return true
	}

	return false
}

// SetDisUeNotif gets a reference to the given bool and assigns it to the DisUeNotif field.
func (o *AsSessionWithQoSSubscriptionPatch) SetDisUeNotif(v bool) {
	o.DisUeNotif = &v
}

// GetUsageThreshold returns the UsageThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsSessionWithQoSSubscriptionPatch) GetUsageThreshold() UsageThresholdRm {
	if o == nil || isNil(o.UsageThreshold.Get()) {
		var ret UsageThresholdRm
		return ret
	}
	return *o.UsageThreshold.Get()
}

// GetUsageThresholdOk returns a tuple with the UsageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsSessionWithQoSSubscriptionPatch) GetUsageThresholdOk() (*UsageThresholdRm, bool) {
	if o == nil {
    return nil, false
	}
	return o.UsageThreshold.Get(), o.UsageThreshold.IsSet()
}

// HasUsageThreshold returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasUsageThreshold() bool {
	if o != nil && o.UsageThreshold.IsSet() {
		return true
	}

	return false
}

// SetUsageThreshold gets a reference to the given NullableUsageThresholdRm and assigns it to the UsageThreshold field.
func (o *AsSessionWithQoSSubscriptionPatch) SetUsageThreshold(v UsageThresholdRm) {
	o.UsageThreshold.Set(&v)
}
// SetUsageThresholdNil sets the value for UsageThreshold to be an explicit nil
func (o *AsSessionWithQoSSubscriptionPatch) SetUsageThresholdNil() {
	o.UsageThreshold.Set(nil)
}

// UnsetUsageThreshold ensures that no value is present for UsageThreshold, not even an explicit nil
func (o *AsSessionWithQoSSubscriptionPatch) UnsetUsageThreshold() {
	o.UsageThreshold.Unset()
}

// GetQosMonInfo returns the QosMonInfo field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetQosMonInfo() QosMonitoringInformationRm {
	if o == nil || isNil(o.QosMonInfo) {
		var ret QosMonitoringInformationRm
		return ret
	}
	return *o.QosMonInfo
}

// GetQosMonInfoOk returns a tuple with the QosMonInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetQosMonInfoOk() (*QosMonitoringInformationRm, bool) {
	if o == nil || isNil(o.QosMonInfo) {
    return nil, false
	}
	return o.QosMonInfo, true
}

// HasQosMonInfo returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasQosMonInfo() bool {
	if o != nil && !isNil(o.QosMonInfo) {
		return true
	}

	return false
}

// SetQosMonInfo gets a reference to the given QosMonitoringInformationRm and assigns it to the QosMonInfo field.
func (o *AsSessionWithQoSSubscriptionPatch) SetQosMonInfo(v QosMonitoringInformationRm) {
	o.QosMonInfo = &v
}

// GetDirectNotifInd returns the DirectNotifInd field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetDirectNotifInd() bool {
	if o == nil || isNil(o.DirectNotifInd) {
		var ret bool
		return ret
	}
	return *o.DirectNotifInd
}

// GetDirectNotifIndOk returns a tuple with the DirectNotifInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetDirectNotifIndOk() (*bool, bool) {
	if o == nil || isNil(o.DirectNotifInd) {
    return nil, false
	}
	return o.DirectNotifInd, true
}

// HasDirectNotifInd returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasDirectNotifInd() bool {
	if o != nil && !isNil(o.DirectNotifInd) {
		return true
	}

	return false
}

// SetDirectNotifInd gets a reference to the given bool and assigns it to the DirectNotifInd field.
func (o *AsSessionWithQoSSubscriptionPatch) SetDirectNotifInd(v bool) {
	o.DirectNotifInd = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetNotificationDestination() string {
	if o == nil || isNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || isNil(o.NotificationDestination) {
    return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasNotificationDestination() bool {
	if o != nil && !isNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *AsSessionWithQoSSubscriptionPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetTscQosReq returns the TscQosReq field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetTscQosReq() TscQosRequirementRm {
	if o == nil || isNil(o.TscQosReq) {
		var ret TscQosRequirementRm
		return ret
	}
	return *o.TscQosReq
}

// GetTscQosReqOk returns a tuple with the TscQosReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetTscQosReqOk() (*TscQosRequirementRm, bool) {
	if o == nil || isNil(o.TscQosReq) {
    return nil, false
	}
	return o.TscQosReq, true
}

// HasTscQosReq returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasTscQosReq() bool {
	if o != nil && !isNil(o.TscQosReq) {
		return true
	}

	return false
}

// SetTscQosReq gets a reference to the given TscQosRequirementRm and assigns it to the TscQosReq field.
func (o *AsSessionWithQoSSubscriptionPatch) SetTscQosReq(v TscQosRequirementRm) {
	o.TscQosReq = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AsSessionWithQoSSubscriptionPatch) GetEvents() []UserPlaneEvent {
	if o == nil || isNil(o.Events) {
		var ret []UserPlaneEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsSessionWithQoSSubscriptionPatch) GetEventsOk() ([]UserPlaneEvent, bool) {
	if o == nil || isNil(o.Events) {
    return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AsSessionWithQoSSubscriptionPatch) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []UserPlaneEvent and assigns it to the Events field.
func (o *AsSessionWithQoSSubscriptionPatch) SetEvents(v []UserPlaneEvent) {
	o.Events = v
}

func (o AsSessionWithQoSSubscriptionPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExterAppId) {
		toSerialize["exterAppId"] = o.ExterAppId
	}
	if !isNil(o.FlowInfo) {
		toSerialize["flowInfo"] = o.FlowInfo
	}
	if !isNil(o.EthFlowInfo) {
		toSerialize["ethFlowInfo"] = o.EthFlowInfo
	}
	if !isNil(o.EnEthFlowInfo) {
		toSerialize["enEthFlowInfo"] = o.EnEthFlowInfo
	}
	if !isNil(o.QosReference) {
		toSerialize["qosReference"] = o.QosReference
	}
	if !isNil(o.AltQoSReferences) {
		toSerialize["altQoSReferences"] = o.AltQoSReferences
	}
	if !isNil(o.AltQosReqs) {
		toSerialize["altQosReqs"] = o.AltQosReqs
	}
	if !isNil(o.DisUeNotif) {
		toSerialize["disUeNotif"] = o.DisUeNotif
	}
	if o.UsageThreshold.IsSet() {
		toSerialize["usageThreshold"] = o.UsageThreshold.Get()
	}
	if !isNil(o.QosMonInfo) {
		toSerialize["qosMonInfo"] = o.QosMonInfo
	}
	if !isNil(o.DirectNotifInd) {
		toSerialize["directNotifInd"] = o.DirectNotifInd
	}
	if !isNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !isNil(o.TscQosReq) {
		toSerialize["tscQosReq"] = o.TscQosReq
	}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableAsSessionWithQoSSubscriptionPatch struct {
	value *AsSessionWithQoSSubscriptionPatch
	isSet bool
}

func (v NullableAsSessionWithQoSSubscriptionPatch) Get() *AsSessionWithQoSSubscriptionPatch {
	return v.value
}

func (v *NullableAsSessionWithQoSSubscriptionPatch) Set(val *AsSessionWithQoSSubscriptionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAsSessionWithQoSSubscriptionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAsSessionWithQoSSubscriptionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsSessionWithQoSSubscriptionPatch(val *AsSessionWithQoSSubscriptionPatch) *NullableAsSessionWithQoSSubscriptionPatch {
	return &NullableAsSessionWithQoSSubscriptionPatch{value: val, isSet: true}
}

func (v NullableAsSessionWithQoSSubscriptionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsSessionWithQoSSubscriptionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


