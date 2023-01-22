/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
)

// SmContextStatusNotification Data within Notify SM Context Status Request
type SmContextStatusNotification struct {
	StatusInfo StatusInfo `json:"statusInfo"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus *ApnRateStatus `json:"apnRateStatus,omitempty"`
	DdnFailureStatus *bool `json:"ddnFailureStatus,omitempty"`
	NotifyCorrelationIdsForddnFailure []string `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NewIntermediateSmfId *string `json:"newIntermediateSmfId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NewSmfId *string `json:"newSmfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	NewSmfSetId *string `json:"newSmfSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	OldSmfId *string `json:"oldSmfId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	OldSmContextRef *string `json:"oldSmContextRef,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AltAnchorSmfUri *string `json:"altAnchorSmfUri,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AltAnchorSmfId *string `json:"altAnchorSmfId,omitempty"`
	TargetDnaiInfo *TargetDnaiInfo `json:"targetDnaiInfo,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	OldPduSessionRef *string `json:"oldPduSessionRef,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	InterPlmnApiRoot *string `json:"interPlmnApiRoot,omitempty"`
}

// NewSmContextStatusNotification instantiates a new SmContextStatusNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextStatusNotification(statusInfo StatusInfo) *SmContextStatusNotification {
	this := SmContextStatusNotification{}
	this.StatusInfo = statusInfo
	var ddnFailureStatus bool = false
	this.DdnFailureStatus = &ddnFailureStatus
	return &this
}

// NewSmContextStatusNotificationWithDefaults instantiates a new SmContextStatusNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextStatusNotificationWithDefaults() *SmContextStatusNotification {
	this := SmContextStatusNotification{}
	var ddnFailureStatus bool = false
	this.DdnFailureStatus = &ddnFailureStatus
	return &this
}

// GetStatusInfo returns the StatusInfo field value
func (o *SmContextStatusNotification) GetStatusInfo() StatusInfo {
	if o == nil {
		var ret StatusInfo
		return ret
	}

	return o.StatusInfo
}

// GetStatusInfoOk returns a tuple with the StatusInfo field value
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetStatusInfoOk() (*StatusInfo, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StatusInfo, true
}

// SetStatusInfo sets field value
func (o *SmContextStatusNotification) SetStatusInfo(v StatusInfo) {
	o.StatusInfo = v
}

// GetSmallDataRateStatus returns the SmallDataRateStatus field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetSmallDataRateStatus() SmallDataRateStatus {
	if o == nil || isNil(o.SmallDataRateStatus) {
		var ret SmallDataRateStatus
		return ret
	}
	return *o.SmallDataRateStatus
}

// GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool) {
	if o == nil || isNil(o.SmallDataRateStatus) {
    return nil, false
	}
	return o.SmallDataRateStatus, true
}

// HasSmallDataRateStatus returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasSmallDataRateStatus() bool {
	if o != nil && !isNil(o.SmallDataRateStatus) {
		return true
	}

	return false
}

// SetSmallDataRateStatus gets a reference to the given SmallDataRateStatus and assigns it to the SmallDataRateStatus field.
func (o *SmContextStatusNotification) SetSmallDataRateStatus(v SmallDataRateStatus) {
	o.SmallDataRateStatus = &v
}

// GetApnRateStatus returns the ApnRateStatus field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetApnRateStatus() ApnRateStatus {
	if o == nil || isNil(o.ApnRateStatus) {
		var ret ApnRateStatus
		return ret
	}
	return *o.ApnRateStatus
}

// GetApnRateStatusOk returns a tuple with the ApnRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetApnRateStatusOk() (*ApnRateStatus, bool) {
	if o == nil || isNil(o.ApnRateStatus) {
    return nil, false
	}
	return o.ApnRateStatus, true
}

// HasApnRateStatus returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasApnRateStatus() bool {
	if o != nil && !isNil(o.ApnRateStatus) {
		return true
	}

	return false
}

// SetApnRateStatus gets a reference to the given ApnRateStatus and assigns it to the ApnRateStatus field.
func (o *SmContextStatusNotification) SetApnRateStatus(v ApnRateStatus) {
	o.ApnRateStatus = &v
}

// GetDdnFailureStatus returns the DdnFailureStatus field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetDdnFailureStatus() bool {
	if o == nil || isNil(o.DdnFailureStatus) {
		var ret bool
		return ret
	}
	return *o.DdnFailureStatus
}

// GetDdnFailureStatusOk returns a tuple with the DdnFailureStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetDdnFailureStatusOk() (*bool, bool) {
	if o == nil || isNil(o.DdnFailureStatus) {
    return nil, false
	}
	return o.DdnFailureStatus, true
}

// HasDdnFailureStatus returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasDdnFailureStatus() bool {
	if o != nil && !isNil(o.DdnFailureStatus) {
		return true
	}

	return false
}

// SetDdnFailureStatus gets a reference to the given bool and assigns it to the DdnFailureStatus field.
func (o *SmContextStatusNotification) SetDdnFailureStatus(v bool) {
	o.DdnFailureStatus = &v
}

// GetNotifyCorrelationIdsForddnFailure returns the NotifyCorrelationIdsForddnFailure field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetNotifyCorrelationIdsForddnFailure() []string {
	if o == nil || isNil(o.NotifyCorrelationIdsForddnFailure) {
		var ret []string
		return ret
	}
	return o.NotifyCorrelationIdsForddnFailure
}

// GetNotifyCorrelationIdsForddnFailureOk returns a tuple with the NotifyCorrelationIdsForddnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetNotifyCorrelationIdsForddnFailureOk() ([]string, bool) {
	if o == nil || isNil(o.NotifyCorrelationIdsForddnFailure) {
    return nil, false
	}
	return o.NotifyCorrelationIdsForddnFailure, true
}

// HasNotifyCorrelationIdsForddnFailure returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasNotifyCorrelationIdsForddnFailure() bool {
	if o != nil && !isNil(o.NotifyCorrelationIdsForddnFailure) {
		return true
	}

	return false
}

// SetNotifyCorrelationIdsForddnFailure gets a reference to the given []string and assigns it to the NotifyCorrelationIdsForddnFailure field.
func (o *SmContextStatusNotification) SetNotifyCorrelationIdsForddnFailure(v []string) {
	o.NotifyCorrelationIdsForddnFailure = v
}

// GetNewIntermediateSmfId returns the NewIntermediateSmfId field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetNewIntermediateSmfId() string {
	if o == nil || isNil(o.NewIntermediateSmfId) {
		var ret string
		return ret
	}
	return *o.NewIntermediateSmfId
}

// GetNewIntermediateSmfIdOk returns a tuple with the NewIntermediateSmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetNewIntermediateSmfIdOk() (*string, bool) {
	if o == nil || isNil(o.NewIntermediateSmfId) {
    return nil, false
	}
	return o.NewIntermediateSmfId, true
}

// HasNewIntermediateSmfId returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasNewIntermediateSmfId() bool {
	if o != nil && !isNil(o.NewIntermediateSmfId) {
		return true
	}

	return false
}

// SetNewIntermediateSmfId gets a reference to the given string and assigns it to the NewIntermediateSmfId field.
func (o *SmContextStatusNotification) SetNewIntermediateSmfId(v string) {
	o.NewIntermediateSmfId = &v
}

// GetNewSmfId returns the NewSmfId field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetNewSmfId() string {
	if o == nil || isNil(o.NewSmfId) {
		var ret string
		return ret
	}
	return *o.NewSmfId
}

// GetNewSmfIdOk returns a tuple with the NewSmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetNewSmfIdOk() (*string, bool) {
	if o == nil || isNil(o.NewSmfId) {
    return nil, false
	}
	return o.NewSmfId, true
}

// HasNewSmfId returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasNewSmfId() bool {
	if o != nil && !isNil(o.NewSmfId) {
		return true
	}

	return false
}

// SetNewSmfId gets a reference to the given string and assigns it to the NewSmfId field.
func (o *SmContextStatusNotification) SetNewSmfId(v string) {
	o.NewSmfId = &v
}

// GetNewSmfSetId returns the NewSmfSetId field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetNewSmfSetId() string {
	if o == nil || isNil(o.NewSmfSetId) {
		var ret string
		return ret
	}
	return *o.NewSmfSetId
}

// GetNewSmfSetIdOk returns a tuple with the NewSmfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetNewSmfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.NewSmfSetId) {
    return nil, false
	}
	return o.NewSmfSetId, true
}

// HasNewSmfSetId returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasNewSmfSetId() bool {
	if o != nil && !isNil(o.NewSmfSetId) {
		return true
	}

	return false
}

// SetNewSmfSetId gets a reference to the given string and assigns it to the NewSmfSetId field.
func (o *SmContextStatusNotification) SetNewSmfSetId(v string) {
	o.NewSmfSetId = &v
}

// GetOldSmfId returns the OldSmfId field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetOldSmfId() string {
	if o == nil || isNil(o.OldSmfId) {
		var ret string
		return ret
	}
	return *o.OldSmfId
}

// GetOldSmfIdOk returns a tuple with the OldSmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetOldSmfIdOk() (*string, bool) {
	if o == nil || isNil(o.OldSmfId) {
    return nil, false
	}
	return o.OldSmfId, true
}

// HasOldSmfId returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasOldSmfId() bool {
	if o != nil && !isNil(o.OldSmfId) {
		return true
	}

	return false
}

// SetOldSmfId gets a reference to the given string and assigns it to the OldSmfId field.
func (o *SmContextStatusNotification) SetOldSmfId(v string) {
	o.OldSmfId = &v
}

// GetOldSmContextRef returns the OldSmContextRef field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetOldSmContextRef() string {
	if o == nil || isNil(o.OldSmContextRef) {
		var ret string
		return ret
	}
	return *o.OldSmContextRef
}

// GetOldSmContextRefOk returns a tuple with the OldSmContextRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetOldSmContextRefOk() (*string, bool) {
	if o == nil || isNil(o.OldSmContextRef) {
    return nil, false
	}
	return o.OldSmContextRef, true
}

// HasOldSmContextRef returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasOldSmContextRef() bool {
	if o != nil && !isNil(o.OldSmContextRef) {
		return true
	}

	return false
}

// SetOldSmContextRef gets a reference to the given string and assigns it to the OldSmContextRef field.
func (o *SmContextStatusNotification) SetOldSmContextRef(v string) {
	o.OldSmContextRef = &v
}

// GetAltAnchorSmfUri returns the AltAnchorSmfUri field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetAltAnchorSmfUri() string {
	if o == nil || isNil(o.AltAnchorSmfUri) {
		var ret string
		return ret
	}
	return *o.AltAnchorSmfUri
}

// GetAltAnchorSmfUriOk returns a tuple with the AltAnchorSmfUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetAltAnchorSmfUriOk() (*string, bool) {
	if o == nil || isNil(o.AltAnchorSmfUri) {
    return nil, false
	}
	return o.AltAnchorSmfUri, true
}

// HasAltAnchorSmfUri returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasAltAnchorSmfUri() bool {
	if o != nil && !isNil(o.AltAnchorSmfUri) {
		return true
	}

	return false
}

// SetAltAnchorSmfUri gets a reference to the given string and assigns it to the AltAnchorSmfUri field.
func (o *SmContextStatusNotification) SetAltAnchorSmfUri(v string) {
	o.AltAnchorSmfUri = &v
}

// GetAltAnchorSmfId returns the AltAnchorSmfId field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetAltAnchorSmfId() string {
	if o == nil || isNil(o.AltAnchorSmfId) {
		var ret string
		return ret
	}
	return *o.AltAnchorSmfId
}

// GetAltAnchorSmfIdOk returns a tuple with the AltAnchorSmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetAltAnchorSmfIdOk() (*string, bool) {
	if o == nil || isNil(o.AltAnchorSmfId) {
    return nil, false
	}
	return o.AltAnchorSmfId, true
}

// HasAltAnchorSmfId returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasAltAnchorSmfId() bool {
	if o != nil && !isNil(o.AltAnchorSmfId) {
		return true
	}

	return false
}

// SetAltAnchorSmfId gets a reference to the given string and assigns it to the AltAnchorSmfId field.
func (o *SmContextStatusNotification) SetAltAnchorSmfId(v string) {
	o.AltAnchorSmfId = &v
}

// GetTargetDnaiInfo returns the TargetDnaiInfo field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetTargetDnaiInfo() TargetDnaiInfo {
	if o == nil || isNil(o.TargetDnaiInfo) {
		var ret TargetDnaiInfo
		return ret
	}
	return *o.TargetDnaiInfo
}

// GetTargetDnaiInfoOk returns a tuple with the TargetDnaiInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetTargetDnaiInfoOk() (*TargetDnaiInfo, bool) {
	if o == nil || isNil(o.TargetDnaiInfo) {
    return nil, false
	}
	return o.TargetDnaiInfo, true
}

// HasTargetDnaiInfo returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasTargetDnaiInfo() bool {
	if o != nil && !isNil(o.TargetDnaiInfo) {
		return true
	}

	return false
}

// SetTargetDnaiInfo gets a reference to the given TargetDnaiInfo and assigns it to the TargetDnaiInfo field.
func (o *SmContextStatusNotification) SetTargetDnaiInfo(v TargetDnaiInfo) {
	o.TargetDnaiInfo = &v
}

// GetOldPduSessionRef returns the OldPduSessionRef field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetOldPduSessionRef() string {
	if o == nil || isNil(o.OldPduSessionRef) {
		var ret string
		return ret
	}
	return *o.OldPduSessionRef
}

// GetOldPduSessionRefOk returns a tuple with the OldPduSessionRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetOldPduSessionRefOk() (*string, bool) {
	if o == nil || isNil(o.OldPduSessionRef) {
    return nil, false
	}
	return o.OldPduSessionRef, true
}

// HasOldPduSessionRef returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasOldPduSessionRef() bool {
	if o != nil && !isNil(o.OldPduSessionRef) {
		return true
	}

	return false
}

// SetOldPduSessionRef gets a reference to the given string and assigns it to the OldPduSessionRef field.
func (o *SmContextStatusNotification) SetOldPduSessionRef(v string) {
	o.OldPduSessionRef = &v
}

// GetInterPlmnApiRoot returns the InterPlmnApiRoot field value if set, zero value otherwise.
func (o *SmContextStatusNotification) GetInterPlmnApiRoot() string {
	if o == nil || isNil(o.InterPlmnApiRoot) {
		var ret string
		return ret
	}
	return *o.InterPlmnApiRoot
}

// GetInterPlmnApiRootOk returns a tuple with the InterPlmnApiRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextStatusNotification) GetInterPlmnApiRootOk() (*string, bool) {
	if o == nil || isNil(o.InterPlmnApiRoot) {
    return nil, false
	}
	return o.InterPlmnApiRoot, true
}

// HasInterPlmnApiRoot returns a boolean if a field has been set.
func (o *SmContextStatusNotification) HasInterPlmnApiRoot() bool {
	if o != nil && !isNil(o.InterPlmnApiRoot) {
		return true
	}

	return false
}

// SetInterPlmnApiRoot gets a reference to the given string and assigns it to the InterPlmnApiRoot field.
func (o *SmContextStatusNotification) SetInterPlmnApiRoot(v string) {
	o.InterPlmnApiRoot = &v
}

func (o SmContextStatusNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["statusInfo"] = o.StatusInfo
	}
	if !isNil(o.SmallDataRateStatus) {
		toSerialize["smallDataRateStatus"] = o.SmallDataRateStatus
	}
	if !isNil(o.ApnRateStatus) {
		toSerialize["apnRateStatus"] = o.ApnRateStatus
	}
	if !isNil(o.DdnFailureStatus) {
		toSerialize["ddnFailureStatus"] = o.DdnFailureStatus
	}
	if !isNil(o.NotifyCorrelationIdsForddnFailure) {
		toSerialize["notifyCorrelationIdsForddnFailure"] = o.NotifyCorrelationIdsForddnFailure
	}
	if !isNil(o.NewIntermediateSmfId) {
		toSerialize["newIntermediateSmfId"] = o.NewIntermediateSmfId
	}
	if !isNil(o.NewSmfId) {
		toSerialize["newSmfId"] = o.NewSmfId
	}
	if !isNil(o.NewSmfSetId) {
		toSerialize["newSmfSetId"] = o.NewSmfSetId
	}
	if !isNil(o.OldSmfId) {
		toSerialize["oldSmfId"] = o.OldSmfId
	}
	if !isNil(o.OldSmContextRef) {
		toSerialize["oldSmContextRef"] = o.OldSmContextRef
	}
	if !isNil(o.AltAnchorSmfUri) {
		toSerialize["altAnchorSmfUri"] = o.AltAnchorSmfUri
	}
	if !isNil(o.AltAnchorSmfId) {
		toSerialize["altAnchorSmfId"] = o.AltAnchorSmfId
	}
	if !isNil(o.TargetDnaiInfo) {
		toSerialize["targetDnaiInfo"] = o.TargetDnaiInfo
	}
	if !isNil(o.OldPduSessionRef) {
		toSerialize["oldPduSessionRef"] = o.OldPduSessionRef
	}
	if !isNil(o.InterPlmnApiRoot) {
		toSerialize["interPlmnApiRoot"] = o.InterPlmnApiRoot
	}
	return json.Marshal(toSerialize)
}

type NullableSmContextStatusNotification struct {
	value *SmContextStatusNotification
	isSet bool
}

func (v NullableSmContextStatusNotification) Get() *SmContextStatusNotification {
	return v.value
}

func (v *NullableSmContextStatusNotification) Set(val *SmContextStatusNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextStatusNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextStatusNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextStatusNotification(val *SmContextStatusNotification) *NullableSmContextStatusNotification {
	return &NullableSmContextStatusNotification{value: val, isSet: true}
}

func (v NullableSmContextStatusNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextStatusNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

