/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
)

// AcrMgntEventSubsc Represents an ACR Management Event Subscription.
type AcrMgntEventSubsc struct {
	Event AcrMgntEvent `json:"event"`
	EventFilter *AcrMgntEventFilter `json:"eventFilter,omitempty"`
	EvtReq *ReportingInformation `json:"evtReq,omitempty"`
	TgtUeId *TargetUeIdentification `json:"tgtUeId,omitempty"`
	DnaiChgType *DnaiChangeType `json:"dnaiChgType,omitempty"`
	EasAckInd *bool `json:"easAckInd,omitempty"`
	// A list of EAS characteristics.
	EasChars []EasCharacteristics `json:"easChars,omitempty"`
}

// NewAcrMgntEventSubsc instantiates a new AcrMgntEventSubsc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcrMgntEventSubsc(event AcrMgntEvent) *AcrMgntEventSubsc {
	this := AcrMgntEventSubsc{}
	this.Event = event
	return &this
}

// NewAcrMgntEventSubscWithDefaults instantiates a new AcrMgntEventSubsc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcrMgntEventSubscWithDefaults() *AcrMgntEventSubsc {
	this := AcrMgntEventSubsc{}
	return &this
}

// GetEvent returns the Event field value
func (o *AcrMgntEventSubsc) GetEvent() AcrMgntEvent {
	if o == nil {
		var ret AcrMgntEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AcrMgntEventSubsc) GetEventOk() (*AcrMgntEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AcrMgntEventSubsc) SetEvent(v AcrMgntEvent) {
	o.Event = v
}

// GetEventFilter returns the EventFilter field value if set, zero value otherwise.
func (o *AcrMgntEventSubsc) GetEventFilter() AcrMgntEventFilter {
	if o == nil || isNil(o.EventFilter) {
		var ret AcrMgntEventFilter
		return ret
	}
	return *o.EventFilter
}

// GetEventFilterOk returns a tuple with the EventFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventSubsc) GetEventFilterOk() (*AcrMgntEventFilter, bool) {
	if o == nil || isNil(o.EventFilter) {
    return nil, false
	}
	return o.EventFilter, true
}

// HasEventFilter returns a boolean if a field has been set.
func (o *AcrMgntEventSubsc) HasEventFilter() bool {
	if o != nil && !isNil(o.EventFilter) {
		return true
	}

	return false
}

// SetEventFilter gets a reference to the given AcrMgntEventFilter and assigns it to the EventFilter field.
func (o *AcrMgntEventSubsc) SetEventFilter(v AcrMgntEventFilter) {
	o.EventFilter = &v
}

// GetEvtReq returns the EvtReq field value if set, zero value otherwise.
func (o *AcrMgntEventSubsc) GetEvtReq() ReportingInformation {
	if o == nil || isNil(o.EvtReq) {
		var ret ReportingInformation
		return ret
	}
	return *o.EvtReq
}

// GetEvtReqOk returns a tuple with the EvtReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventSubsc) GetEvtReqOk() (*ReportingInformation, bool) {
	if o == nil || isNil(o.EvtReq) {
    return nil, false
	}
	return o.EvtReq, true
}

// HasEvtReq returns a boolean if a field has been set.
func (o *AcrMgntEventSubsc) HasEvtReq() bool {
	if o != nil && !isNil(o.EvtReq) {
		return true
	}

	return false
}

// SetEvtReq gets a reference to the given ReportingInformation and assigns it to the EvtReq field.
func (o *AcrMgntEventSubsc) SetEvtReq(v ReportingInformation) {
	o.EvtReq = &v
}

// GetTgtUeId returns the TgtUeId field value if set, zero value otherwise.
func (o *AcrMgntEventSubsc) GetTgtUeId() TargetUeIdentification {
	if o == nil || isNil(o.TgtUeId) {
		var ret TargetUeIdentification
		return ret
	}
	return *o.TgtUeId
}

// GetTgtUeIdOk returns a tuple with the TgtUeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventSubsc) GetTgtUeIdOk() (*TargetUeIdentification, bool) {
	if o == nil || isNil(o.TgtUeId) {
    return nil, false
	}
	return o.TgtUeId, true
}

// HasTgtUeId returns a boolean if a field has been set.
func (o *AcrMgntEventSubsc) HasTgtUeId() bool {
	if o != nil && !isNil(o.TgtUeId) {
		return true
	}

	return false
}

// SetTgtUeId gets a reference to the given TargetUeIdentification and assigns it to the TgtUeId field.
func (o *AcrMgntEventSubsc) SetTgtUeId(v TargetUeIdentification) {
	o.TgtUeId = &v
}

// GetDnaiChgType returns the DnaiChgType field value if set, zero value otherwise.
func (o *AcrMgntEventSubsc) GetDnaiChgType() DnaiChangeType {
	if o == nil || isNil(o.DnaiChgType) {
		var ret DnaiChangeType
		return ret
	}
	return *o.DnaiChgType
}

// GetDnaiChgTypeOk returns a tuple with the DnaiChgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventSubsc) GetDnaiChgTypeOk() (*DnaiChangeType, bool) {
	if o == nil || isNil(o.DnaiChgType) {
    return nil, false
	}
	return o.DnaiChgType, true
}

// HasDnaiChgType returns a boolean if a field has been set.
func (o *AcrMgntEventSubsc) HasDnaiChgType() bool {
	if o != nil && !isNil(o.DnaiChgType) {
		return true
	}

	return false
}

// SetDnaiChgType gets a reference to the given DnaiChangeType and assigns it to the DnaiChgType field.
func (o *AcrMgntEventSubsc) SetDnaiChgType(v DnaiChangeType) {
	o.DnaiChgType = &v
}

// GetEasAckInd returns the EasAckInd field value if set, zero value otherwise.
func (o *AcrMgntEventSubsc) GetEasAckInd() bool {
	if o == nil || isNil(o.EasAckInd) {
		var ret bool
		return ret
	}
	return *o.EasAckInd
}

// GetEasAckIndOk returns a tuple with the EasAckInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventSubsc) GetEasAckIndOk() (*bool, bool) {
	if o == nil || isNil(o.EasAckInd) {
    return nil, false
	}
	return o.EasAckInd, true
}

// HasEasAckInd returns a boolean if a field has been set.
func (o *AcrMgntEventSubsc) HasEasAckInd() bool {
	if o != nil && !isNil(o.EasAckInd) {
		return true
	}

	return false
}

// SetEasAckInd gets a reference to the given bool and assigns it to the EasAckInd field.
func (o *AcrMgntEventSubsc) SetEasAckInd(v bool) {
	o.EasAckInd = &v
}

// GetEasChars returns the EasChars field value if set, zero value otherwise.
func (o *AcrMgntEventSubsc) GetEasChars() []EasCharacteristics {
	if o == nil || isNil(o.EasChars) {
		var ret []EasCharacteristics
		return ret
	}
	return o.EasChars
}

// GetEasCharsOk returns a tuple with the EasChars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventSubsc) GetEasCharsOk() ([]EasCharacteristics, bool) {
	if o == nil || isNil(o.EasChars) {
    return nil, false
	}
	return o.EasChars, true
}

// HasEasChars returns a boolean if a field has been set.
func (o *AcrMgntEventSubsc) HasEasChars() bool {
	if o != nil && !isNil(o.EasChars) {
		return true
	}

	return false
}

// SetEasChars gets a reference to the given []EasCharacteristics and assigns it to the EasChars field.
func (o *AcrMgntEventSubsc) SetEasChars(v []EasCharacteristics) {
	o.EasChars = v
}

func (o AcrMgntEventSubsc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if !isNil(o.EventFilter) {
		toSerialize["eventFilter"] = o.EventFilter
	}
	if !isNil(o.EvtReq) {
		toSerialize["evtReq"] = o.EvtReq
	}
	if !isNil(o.TgtUeId) {
		toSerialize["tgtUeId"] = o.TgtUeId
	}
	if !isNil(o.DnaiChgType) {
		toSerialize["dnaiChgType"] = o.DnaiChgType
	}
	if !isNil(o.EasAckInd) {
		toSerialize["easAckInd"] = o.EasAckInd
	}
	if !isNil(o.EasChars) {
		toSerialize["easChars"] = o.EasChars
	}
	return json.Marshal(toSerialize)
}

type NullableAcrMgntEventSubsc struct {
	value *AcrMgntEventSubsc
	isSet bool
}

func (v NullableAcrMgntEventSubsc) Get() *AcrMgntEventSubsc {
	return v.value
}

func (v *NullableAcrMgntEventSubsc) Set(val *AcrMgntEventSubsc) {
	v.value = val
	v.isSet = true
}

func (v NullableAcrMgntEventSubsc) IsSet() bool {
	return v.isSet
}

func (v *NullableAcrMgntEventSubsc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcrMgntEventSubsc(val *AcrMgntEventSubsc) *NullableAcrMgntEventSubsc {
	return &NullableAcrMgntEventSubsc{value: val, isSet: true}
}

func (v NullableAcrMgntEventSubsc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcrMgntEventSubsc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


