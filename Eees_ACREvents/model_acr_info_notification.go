/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_ACREvents

import (
	"encoding/json"
)

// ACRInfoNotification Notification of ACR events information.
type ACRInfoNotification struct {
	// String identifying the Individual ACR events subscription for which the ACT Information notification is delivered. 
	SubId string `json:"subId"`
	// Application identifier of the EAS.
	EasId string `json:"easId"`
	// Identity of the AC.
	AcId *string `json:"acId,omitempty"`
	EventId ACREventIDs `json:"eventId"`
	TrgtInfo *TargetInfo `json:"trgtInfo,omitempty"`
	AcrStatus *ACRCompleteEventInfo `json:"acrStatus,omitempty"`
	// Indicates the cause information for the failure.
	FailReason *string `json:"failReason,omitempty"`
	EecCtxtReloc *EecCtxtRelocStatus `json:"eecCtxtReloc,omitempty"`
}

// NewACRInfoNotification instantiates a new ACRInfoNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACRInfoNotification(subId string, easId string, eventId ACREventIDs) *ACRInfoNotification {
	this := ACRInfoNotification{}
	this.SubId = subId
	this.EasId = easId
	this.EventId = eventId
	return &this
}

// NewACRInfoNotificationWithDefaults instantiates a new ACRInfoNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACRInfoNotificationWithDefaults() *ACRInfoNotification {
	this := ACRInfoNotification{}
	return &this
}

// GetSubId returns the SubId field value
func (o *ACRInfoNotification) GetSubId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubId
}

// GetSubIdOk returns a tuple with the SubId field value
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetSubIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubId, true
}

// SetSubId sets field value
func (o *ACRInfoNotification) SetSubId(v string) {
	o.SubId = v
}

// GetEasId returns the EasId field value
func (o *ACRInfoNotification) GetEasId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetEasIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EasId, true
}

// SetEasId sets field value
func (o *ACRInfoNotification) SetEasId(v string) {
	o.EasId = v
}

// GetAcId returns the AcId field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetAcId() string {
	if o == nil || isNil(o.AcId) {
		var ret string
		return ret
	}
	return *o.AcId
}

// GetAcIdOk returns a tuple with the AcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetAcIdOk() (*string, bool) {
	if o == nil || isNil(o.AcId) {
    return nil, false
	}
	return o.AcId, true
}

// HasAcId returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasAcId() bool {
	if o != nil && !isNil(o.AcId) {
		return true
	}

	return false
}

// SetAcId gets a reference to the given string and assigns it to the AcId field.
func (o *ACRInfoNotification) SetAcId(v string) {
	o.AcId = &v
}

// GetEventId returns the EventId field value
func (o *ACRInfoNotification) GetEventId() ACREventIDs {
	if o == nil {
		var ret ACREventIDs
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetEventIdOk() (*ACREventIDs, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *ACRInfoNotification) SetEventId(v ACREventIDs) {
	o.EventId = v
}

// GetTrgtInfo returns the TrgtInfo field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetTrgtInfo() TargetInfo {
	if o == nil || isNil(o.TrgtInfo) {
		var ret TargetInfo
		return ret
	}
	return *o.TrgtInfo
}

// GetTrgtInfoOk returns a tuple with the TrgtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetTrgtInfoOk() (*TargetInfo, bool) {
	if o == nil || isNil(o.TrgtInfo) {
    return nil, false
	}
	return o.TrgtInfo, true
}

// HasTrgtInfo returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasTrgtInfo() bool {
	if o != nil && !isNil(o.TrgtInfo) {
		return true
	}

	return false
}

// SetTrgtInfo gets a reference to the given TargetInfo and assigns it to the TrgtInfo field.
func (o *ACRInfoNotification) SetTrgtInfo(v TargetInfo) {
	o.TrgtInfo = &v
}

// GetAcrStatus returns the AcrStatus field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetAcrStatus() ACRCompleteEventInfo {
	if o == nil || isNil(o.AcrStatus) {
		var ret ACRCompleteEventInfo
		return ret
	}
	return *o.AcrStatus
}

// GetAcrStatusOk returns a tuple with the AcrStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetAcrStatusOk() (*ACRCompleteEventInfo, bool) {
	if o == nil || isNil(o.AcrStatus) {
    return nil, false
	}
	return o.AcrStatus, true
}

// HasAcrStatus returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasAcrStatus() bool {
	if o != nil && !isNil(o.AcrStatus) {
		return true
	}

	return false
}

// SetAcrStatus gets a reference to the given ACRCompleteEventInfo and assigns it to the AcrStatus field.
func (o *ACRInfoNotification) SetAcrStatus(v ACRCompleteEventInfo) {
	o.AcrStatus = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetFailReason() string {
	if o == nil || isNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetFailReasonOk() (*string, bool) {
	if o == nil || isNil(o.FailReason) {
    return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasFailReason() bool {
	if o != nil && !isNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *ACRInfoNotification) SetFailReason(v string) {
	o.FailReason = &v
}

// GetEecCtxtReloc returns the EecCtxtReloc field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetEecCtxtReloc() EecCtxtRelocStatus {
	if o == nil || isNil(o.EecCtxtReloc) {
		var ret EecCtxtRelocStatus
		return ret
	}
	return *o.EecCtxtReloc
}

// GetEecCtxtRelocOk returns a tuple with the EecCtxtReloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetEecCtxtRelocOk() (*EecCtxtRelocStatus, bool) {
	if o == nil || isNil(o.EecCtxtReloc) {
    return nil, false
	}
	return o.EecCtxtReloc, true
}

// HasEecCtxtReloc returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasEecCtxtReloc() bool {
	if o != nil && !isNil(o.EecCtxtReloc) {
		return true
	}

	return false
}

// SetEecCtxtReloc gets a reference to the given EecCtxtRelocStatus and assigns it to the EecCtxtReloc field.
func (o *ACRInfoNotification) SetEecCtxtReloc(v EecCtxtRelocStatus) {
	o.EecCtxtReloc = &v
}

func (o ACRInfoNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subId"] = o.SubId
	}
	if true {
		toSerialize["easId"] = o.EasId
	}
	if !isNil(o.AcId) {
		toSerialize["acId"] = o.AcId
	}
	if true {
		toSerialize["eventId"] = o.EventId
	}
	if !isNil(o.TrgtInfo) {
		toSerialize["trgtInfo"] = o.TrgtInfo
	}
	if !isNil(o.AcrStatus) {
		toSerialize["acrStatus"] = o.AcrStatus
	}
	if !isNil(o.FailReason) {
		toSerialize["failReason"] = o.FailReason
	}
	if !isNil(o.EecCtxtReloc) {
		toSerialize["eecCtxtReloc"] = o.EecCtxtReloc
	}
	return json.Marshal(toSerialize)
}

type NullableACRInfoNotification struct {
	value *ACRInfoNotification
	isSet bool
}

func (v NullableACRInfoNotification) Get() *ACRInfoNotification {
	return v.value
}

func (v *NullableACRInfoNotification) Set(val *ACRInfoNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableACRInfoNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableACRInfoNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACRInfoNotification(val *ACRInfoNotification) *NullableACRInfoNotification {
	return &NullableACRInfoNotification{value: val, isSet: true}
}

func (v NullableACRInfoNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACRInfoNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


