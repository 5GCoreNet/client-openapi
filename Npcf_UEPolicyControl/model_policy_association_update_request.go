/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
)

// PolicyAssociationUpdateRequest Represents Information that the NF service consumer provides when requesting the update of a policy association. 
type PolicyAssociationUpdateRequest struct {
	// String providing an URI formatted according to RFC 3986.
	NotificationUri *string `json:"notificationUri,omitempty"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []Ipv6Addr `json:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty"`
	// Request Triggers that the NF service consumer observes.
	Triggers []RequestTrigger `json:"triggers,omitempty"`
	// Contains the UE presence status for tracking area for which changes of the UE presence occurred. The praId attribute within the PresenceInfo data type is the key of the map. 
	PraStatuses *map[string]PresenceInfo `json:"praStatuses,omitempty"`
	UserLoc *UserLocation `json:"userLoc,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	UePolDelResult *string `json:"uePolDelResult,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	UePolReq *string `json:"uePolReq,omitempty"`
	Guami *Guami `json:"guami,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	ServingNfId *string `json:"servingNfId,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	ConnectState *CmState `json:"connectState,omitempty"`
	GroupIds []string `json:"groupIds,omitempty"`
	ProSeCapab []ProSeCapability `json:"proSeCapab,omitempty"`
}

// NewPolicyAssociationUpdateRequest instantiates a new PolicyAssociationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAssociationUpdateRequest() *PolicyAssociationUpdateRequest {
	this := PolicyAssociationUpdateRequest{}
	return &this
}

// NewPolicyAssociationUpdateRequestWithDefaults instantiates a new PolicyAssociationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAssociationUpdateRequestWithDefaults() *PolicyAssociationUpdateRequest {
	this := PolicyAssociationUpdateRequest{}
	return &this
}

// GetNotificationUri returns the NotificationUri field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetNotificationUri() string {
	if o == nil || isNil(o.NotificationUri) {
		var ret string
		return ret
	}
	return *o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetNotificationUriOk() (*string, bool) {
	if o == nil || isNil(o.NotificationUri) {
    return nil, false
	}
	return o.NotificationUri, true
}

// HasNotificationUri returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasNotificationUri() bool {
	if o != nil && !isNil(o.NotificationUri) {
		return true
	}

	return false
}

// SetNotificationUri gets a reference to the given string and assigns it to the NotificationUri field.
func (o *PolicyAssociationUpdateRequest) SetNotificationUri(v string) {
	o.NotificationUri = &v
}

// GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4Addrs() []string {
	if o == nil || isNil(o.AltNotifIpv4Addrs) {
		var ret []string
		return ret
	}
	return o.AltNotifIpv4Addrs
}

// GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4AddrsOk() ([]string, bool) {
	if o == nil || isNil(o.AltNotifIpv4Addrs) {
    return nil, false
	}
	return o.AltNotifIpv4Addrs, true
}

// HasAltNotifIpv4Addrs returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv4Addrs() bool {
	if o != nil && !isNil(o.AltNotifIpv4Addrs) {
		return true
	}

	return false
}

// SetAltNotifIpv4Addrs gets a reference to the given []string and assigns it to the AltNotifIpv4Addrs field.
func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv4Addrs(v []string) {
	o.AltNotifIpv4Addrs = v
}

// GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6Addrs() []Ipv6Addr {
	if o == nil || isNil(o.AltNotifIpv6Addrs) {
		var ret []Ipv6Addr
		return ret
	}
	return o.AltNotifIpv6Addrs
}

// GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || isNil(o.AltNotifIpv6Addrs) {
    return nil, false
	}
	return o.AltNotifIpv6Addrs, true
}

// HasAltNotifIpv6Addrs returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv6Addrs() bool {
	if o != nil && !isNil(o.AltNotifIpv6Addrs) {
		return true
	}

	return false
}

// SetAltNotifIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the AltNotifIpv6Addrs field.
func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv6Addrs(v []Ipv6Addr) {
	o.AltNotifIpv6Addrs = v
}

// GetAltNotifFqdns returns the AltNotifFqdns field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdns() []string {
	if o == nil || isNil(o.AltNotifFqdns) {
		var ret []string
		return ret
	}
	return o.AltNotifFqdns
}

// GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdnsOk() ([]string, bool) {
	if o == nil || isNil(o.AltNotifFqdns) {
    return nil, false
	}
	return o.AltNotifFqdns, true
}

// HasAltNotifFqdns returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAltNotifFqdns() bool {
	if o != nil && !isNil(o.AltNotifFqdns) {
		return true
	}

	return false
}

// SetAltNotifFqdns gets a reference to the given []string and assigns it to the AltNotifFqdns field.
func (o *PolicyAssociationUpdateRequest) SetAltNotifFqdns(v []string) {
	o.AltNotifFqdns = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetTriggers() []RequestTrigger {
	if o == nil || isNil(o.Triggers) {
		var ret []RequestTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetTriggersOk() ([]RequestTrigger, bool) {
	if o == nil || isNil(o.Triggers) {
    return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasTriggers() bool {
	if o != nil && !isNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []RequestTrigger and assigns it to the Triggers field.
func (o *PolicyAssociationUpdateRequest) SetTriggers(v []RequestTrigger) {
	o.Triggers = v
}

// GetPraStatuses returns the PraStatuses field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetPraStatuses() map[string]PresenceInfo {
	if o == nil || isNil(o.PraStatuses) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.PraStatuses
}

// GetPraStatusesOk returns a tuple with the PraStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetPraStatusesOk() (*map[string]PresenceInfo, bool) {
	if o == nil || isNil(o.PraStatuses) {
    return nil, false
	}
	return o.PraStatuses, true
}

// HasPraStatuses returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasPraStatuses() bool {
	if o != nil && !isNil(o.PraStatuses) {
		return true
	}

	return false
}

// SetPraStatuses gets a reference to the given map[string]PresenceInfo and assigns it to the PraStatuses field.
func (o *PolicyAssociationUpdateRequest) SetPraStatuses(v map[string]PresenceInfo) {
	o.PraStatuses = &v
}

// GetUserLoc returns the UserLoc field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetUserLoc() UserLocation {
	if o == nil || isNil(o.UserLoc) {
		var ret UserLocation
		return ret
	}
	return *o.UserLoc
}

// GetUserLocOk returns a tuple with the UserLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetUserLocOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UserLoc) {
    return nil, false
	}
	return o.UserLoc, true
}

// HasUserLoc returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasUserLoc() bool {
	if o != nil && !isNil(o.UserLoc) {
		return true
	}

	return false
}

// SetUserLoc gets a reference to the given UserLocation and assigns it to the UserLoc field.
func (o *PolicyAssociationUpdateRequest) SetUserLoc(v UserLocation) {
	o.UserLoc = &v
}

// GetUePolDelResult returns the UePolDelResult field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetUePolDelResult() string {
	if o == nil || isNil(o.UePolDelResult) {
		var ret string
		return ret
	}
	return *o.UePolDelResult
}

// GetUePolDelResultOk returns a tuple with the UePolDelResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetUePolDelResultOk() (*string, bool) {
	if o == nil || isNil(o.UePolDelResult) {
    return nil, false
	}
	return o.UePolDelResult, true
}

// HasUePolDelResult returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasUePolDelResult() bool {
	if o != nil && !isNil(o.UePolDelResult) {
		return true
	}

	return false
}

// SetUePolDelResult gets a reference to the given string and assigns it to the UePolDelResult field.
func (o *PolicyAssociationUpdateRequest) SetUePolDelResult(v string) {
	o.UePolDelResult = &v
}

// GetUePolTransFailNotif returns the UePolTransFailNotif field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetUePolTransFailNotif() UePolicyTransferFailureNotification {
	if o == nil || isNil(o.UePolTransFailNotif) {
		var ret UePolicyTransferFailureNotification
		return ret
	}
	return *o.UePolTransFailNotif
}

// GetUePolTransFailNotifOk returns a tuple with the UePolTransFailNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetUePolTransFailNotifOk() (*UePolicyTransferFailureNotification, bool) {
	if o == nil || isNil(o.UePolTransFailNotif) {
    return nil, false
	}
	return o.UePolTransFailNotif, true
}

// HasUePolTransFailNotif returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasUePolTransFailNotif() bool {
	if o != nil && !isNil(o.UePolTransFailNotif) {
		return true
	}

	return false
}

// SetUePolTransFailNotif gets a reference to the given UePolicyTransferFailureNotification and assigns it to the UePolTransFailNotif field.
func (o *PolicyAssociationUpdateRequest) SetUePolTransFailNotif(v UePolicyTransferFailureNotification) {
	o.UePolTransFailNotif = &v
}

// GetUePolReq returns the UePolReq field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetUePolReq() string {
	if o == nil || isNil(o.UePolReq) {
		var ret string
		return ret
	}
	return *o.UePolReq
}

// GetUePolReqOk returns a tuple with the UePolReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetUePolReqOk() (*string, bool) {
	if o == nil || isNil(o.UePolReq) {
    return nil, false
	}
	return o.UePolReq, true
}

// HasUePolReq returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasUePolReq() bool {
	if o != nil && !isNil(o.UePolReq) {
		return true
	}

	return false
}

// SetUePolReq gets a reference to the given string and assigns it to the UePolReq field.
func (o *PolicyAssociationUpdateRequest) SetUePolReq(v string) {
	o.UePolReq = &v
}

// GetGuami returns the Guami field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetGuami() Guami {
	if o == nil || isNil(o.Guami) {
		var ret Guami
		return ret
	}
	return *o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetGuamiOk() (*Guami, bool) {
	if o == nil || isNil(o.Guami) {
    return nil, false
	}
	return o.Guami, true
}

// HasGuami returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasGuami() bool {
	if o != nil && !isNil(o.Guami) {
		return true
	}

	return false
}

// SetGuami gets a reference to the given Guami and assigns it to the Guami field.
func (o *PolicyAssociationUpdateRequest) SetGuami(v Guami) {
	o.Guami = &v
}

// GetServingNfId returns the ServingNfId field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetServingNfId() string {
	if o == nil || isNil(o.ServingNfId) {
		var ret string
		return ret
	}
	return *o.ServingNfId
}

// GetServingNfIdOk returns a tuple with the ServingNfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetServingNfIdOk() (*string, bool) {
	if o == nil || isNil(o.ServingNfId) {
    return nil, false
	}
	return o.ServingNfId, true
}

// HasServingNfId returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasServingNfId() bool {
	if o != nil && !isNil(o.ServingNfId) {
		return true
	}

	return false
}

// SetServingNfId gets a reference to the given string and assigns it to the ServingNfId field.
func (o *PolicyAssociationUpdateRequest) SetServingNfId(v string) {
	o.ServingNfId = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetPlmnId() PlmnId {
	if o == nil || isNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.PlmnId) {
    return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasPlmnId() bool {
	if o != nil && !isNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *PolicyAssociationUpdateRequest) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetConnectState returns the ConnectState field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetConnectState() CmState {
	if o == nil || isNil(o.ConnectState) {
		var ret CmState
		return ret
	}
	return *o.ConnectState
}

// GetConnectStateOk returns a tuple with the ConnectState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetConnectStateOk() (*CmState, bool) {
	if o == nil || isNil(o.ConnectState) {
    return nil, false
	}
	return o.ConnectState, true
}

// HasConnectState returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasConnectState() bool {
	if o != nil && !isNil(o.ConnectState) {
		return true
	}

	return false
}

// SetConnectState gets a reference to the given CmState and assigns it to the ConnectState field.
func (o *PolicyAssociationUpdateRequest) SetConnectState(v CmState) {
	o.ConnectState = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetGroupIds() []string {
	if o == nil || isNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.GroupIds) {
    return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasGroupIds() bool {
	if o != nil && !isNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *PolicyAssociationUpdateRequest) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetProSeCapab returns the ProSeCapab field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetProSeCapab() []ProSeCapability {
	if o == nil || isNil(o.ProSeCapab) {
		var ret []ProSeCapability
		return ret
	}
	return o.ProSeCapab
}

// GetProSeCapabOk returns a tuple with the ProSeCapab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetProSeCapabOk() ([]ProSeCapability, bool) {
	if o == nil || isNil(o.ProSeCapab) {
    return nil, false
	}
	return o.ProSeCapab, true
}

// HasProSeCapab returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasProSeCapab() bool {
	if o != nil && !isNil(o.ProSeCapab) {
		return true
	}

	return false
}

// SetProSeCapab gets a reference to the given []ProSeCapability and assigns it to the ProSeCapab field.
func (o *PolicyAssociationUpdateRequest) SetProSeCapab(v []ProSeCapability) {
	o.ProSeCapab = v
}

func (o PolicyAssociationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NotificationUri) {
		toSerialize["notificationUri"] = o.NotificationUri
	}
	if !isNil(o.AltNotifIpv4Addrs) {
		toSerialize["altNotifIpv4Addrs"] = o.AltNotifIpv4Addrs
	}
	if !isNil(o.AltNotifIpv6Addrs) {
		toSerialize["altNotifIpv6Addrs"] = o.AltNotifIpv6Addrs
	}
	if !isNil(o.AltNotifFqdns) {
		toSerialize["altNotifFqdns"] = o.AltNotifFqdns
	}
	if !isNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !isNil(o.PraStatuses) {
		toSerialize["praStatuses"] = o.PraStatuses
	}
	if !isNil(o.UserLoc) {
		toSerialize["userLoc"] = o.UserLoc
	}
	if !isNil(o.UePolDelResult) {
		toSerialize["uePolDelResult"] = o.UePolDelResult
	}
	if !isNil(o.UePolTransFailNotif) {
		toSerialize["uePolTransFailNotif"] = o.UePolTransFailNotif
	}
	if !isNil(o.UePolReq) {
		toSerialize["uePolReq"] = o.UePolReq
	}
	if !isNil(o.Guami) {
		toSerialize["guami"] = o.Guami
	}
	if !isNil(o.ServingNfId) {
		toSerialize["servingNfId"] = o.ServingNfId
	}
	if !isNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.ConnectState) {
		toSerialize["connectState"] = o.ConnectState
	}
	if !isNil(o.GroupIds) {
		toSerialize["groupIds"] = o.GroupIds
	}
	if !isNil(o.ProSeCapab) {
		toSerialize["proSeCapab"] = o.ProSeCapab
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyAssociationUpdateRequest struct {
	value *PolicyAssociationUpdateRequest
	isSet bool
}

func (v NullablePolicyAssociationUpdateRequest) Get() *PolicyAssociationUpdateRequest {
	return v.value
}

func (v *NullablePolicyAssociationUpdateRequest) Set(val *PolicyAssociationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAssociationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAssociationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAssociationUpdateRequest(val *PolicyAssociationUpdateRequest) *NullablePolicyAssociationUpdateRequest {
	return &NullablePolicyAssociationUpdateRequest{value: val, isSet: true}
}

func (v NullablePolicyAssociationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAssociationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


