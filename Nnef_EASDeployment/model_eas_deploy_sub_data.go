/*
Nnef_EASDeployment

NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_EASDeployment

import (
	"encoding/json"
)

// EasDeploySubData Represents an Individual EAS Deployment Event Subscription resource.
type EasDeploySubData struct {
	AppId *string `json:"appId,omitempty"`
	// Each of the element identifies a (DNN, S-NSSAI) combination.
	DnnSnssaiInfos []DnnSnssaiInformation `json:"dnnSnssaiInfos,omitempty"`
	EventId EasEvent `json:"eventId"`
	// Represents the EAS Deployment Information changes event(s) to be reported. Shall only be present if the \"immRep\" attribute is included and sets to true, and the current status of EAS Deployment Information is available. 
	EventsNotifs []EasDeployInfoData `json:"eventsNotifs,omitempty"`
	// Indication of immediate reporting. Set to true: requires the immediate reporting of the  current status of EAS Deployment Information, if available. Set to false (default): EAS  Deployment Information event report occurs when the event is met. 
	ImmRep *bool `json:"immRep,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGroupId *string `json:"interGroupId,omitempty"`
	NotifId string `json:"notifId"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
}

// NewEasDeploySubData instantiates a new EasDeploySubData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDeploySubData(eventId EasEvent, notifId string, notifUri string) *EasDeploySubData {
	this := EasDeploySubData{}
	this.EventId = eventId
	this.NotifId = notifId
	this.NotifUri = notifUri
	return &this
}

// NewEasDeploySubDataWithDefaults instantiates a new EasDeploySubData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDeploySubDataWithDefaults() *EasDeploySubData {
	this := EasDeploySubData{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *EasDeploySubData) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *EasDeploySubData) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *EasDeploySubData) SetAppId(v string) {
	o.AppId = &v
}

// GetDnnSnssaiInfos returns the DnnSnssaiInfos field value if set, zero value otherwise.
func (o *EasDeploySubData) GetDnnSnssaiInfos() []DnnSnssaiInformation {
	if o == nil || isNil(o.DnnSnssaiInfos) {
		var ret []DnnSnssaiInformation
		return ret
	}
	return o.DnnSnssaiInfos
}

// GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetDnnSnssaiInfosOk() ([]DnnSnssaiInformation, bool) {
	if o == nil || isNil(o.DnnSnssaiInfos) {
    return nil, false
	}
	return o.DnnSnssaiInfos, true
}

// HasDnnSnssaiInfos returns a boolean if a field has been set.
func (o *EasDeploySubData) HasDnnSnssaiInfos() bool {
	if o != nil && !isNil(o.DnnSnssaiInfos) {
		return true
	}

	return false
}

// SetDnnSnssaiInfos gets a reference to the given []DnnSnssaiInformation and assigns it to the DnnSnssaiInfos field.
func (o *EasDeploySubData) SetDnnSnssaiInfos(v []DnnSnssaiInformation) {
	o.DnnSnssaiInfos = v
}

// GetEventId returns the EventId field value
func (o *EasDeploySubData) GetEventId() EasEvent {
	if o == nil {
		var ret EasEvent
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetEventIdOk() (*EasEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *EasDeploySubData) SetEventId(v EasEvent) {
	o.EventId = v
}

// GetEventsNotifs returns the EventsNotifs field value if set, zero value otherwise.
func (o *EasDeploySubData) GetEventsNotifs() []EasDeployInfoData {
	if o == nil || isNil(o.EventsNotifs) {
		var ret []EasDeployInfoData
		return ret
	}
	return o.EventsNotifs
}

// GetEventsNotifsOk returns a tuple with the EventsNotifs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetEventsNotifsOk() ([]EasDeployInfoData, bool) {
	if o == nil || isNil(o.EventsNotifs) {
    return nil, false
	}
	return o.EventsNotifs, true
}

// HasEventsNotifs returns a boolean if a field has been set.
func (o *EasDeploySubData) HasEventsNotifs() bool {
	if o != nil && !isNil(o.EventsNotifs) {
		return true
	}

	return false
}

// SetEventsNotifs gets a reference to the given []EasDeployInfoData and assigns it to the EventsNotifs field.
func (o *EasDeploySubData) SetEventsNotifs(v []EasDeployInfoData) {
	o.EventsNotifs = v
}

// GetImmRep returns the ImmRep field value if set, zero value otherwise.
func (o *EasDeploySubData) GetImmRep() bool {
	if o == nil || isNil(o.ImmRep) {
		var ret bool
		return ret
	}
	return *o.ImmRep
}

// GetImmRepOk returns a tuple with the ImmRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetImmRepOk() (*bool, bool) {
	if o == nil || isNil(o.ImmRep) {
    return nil, false
	}
	return o.ImmRep, true
}

// HasImmRep returns a boolean if a field has been set.
func (o *EasDeploySubData) HasImmRep() bool {
	if o != nil && !isNil(o.ImmRep) {
		return true
	}

	return false
}

// SetImmRep gets a reference to the given bool and assigns it to the ImmRep field.
func (o *EasDeploySubData) SetImmRep(v bool) {
	o.ImmRep = &v
}

// GetInterGroupId returns the InterGroupId field value if set, zero value otherwise.
func (o *EasDeploySubData) GetInterGroupId() string {
	if o == nil || isNil(o.InterGroupId) {
		var ret string
		return ret
	}
	return *o.InterGroupId
}

// GetInterGroupIdOk returns a tuple with the InterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetInterGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.InterGroupId) {
    return nil, false
	}
	return o.InterGroupId, true
}

// HasInterGroupId returns a boolean if a field has been set.
func (o *EasDeploySubData) HasInterGroupId() bool {
	if o != nil && !isNil(o.InterGroupId) {
		return true
	}

	return false
}

// SetInterGroupId gets a reference to the given string and assigns it to the InterGroupId field.
func (o *EasDeploySubData) SetInterGroupId(v string) {
	o.InterGroupId = &v
}

// GetNotifId returns the NotifId field value
func (o *EasDeploySubData) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetNotifIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *EasDeploySubData) SetNotifId(v string) {
	o.NotifId = v
}

// GetNotifUri returns the NotifUri field value
func (o *EasDeploySubData) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *EasDeploySubData) GetNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *EasDeploySubData) SetNotifUri(v string) {
	o.NotifUri = v
}

func (o EasDeploySubData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.DnnSnssaiInfos) {
		toSerialize["dnnSnssaiInfos"] = o.DnnSnssaiInfos
	}
	if true {
		toSerialize["eventId"] = o.EventId
	}
	if !isNil(o.EventsNotifs) {
		toSerialize["eventsNotifs"] = o.EventsNotifs
	}
	if !isNil(o.ImmRep) {
		toSerialize["immRep"] = o.ImmRep
	}
	if !isNil(o.InterGroupId) {
		toSerialize["interGroupId"] = o.InterGroupId
	}
	if true {
		toSerialize["notifId"] = o.NotifId
	}
	if true {
		toSerialize["notifUri"] = o.NotifUri
	}
	return json.Marshal(toSerialize)
}

type NullableEasDeploySubData struct {
	value *EasDeploySubData
	isSet bool
}

func (v NullableEasDeploySubData) Get() *EasDeploySubData {
	return v.value
}

func (v *NullableEasDeploySubData) Set(val *EasDeploySubData) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDeploySubData) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDeploySubData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDeploySubData(val *EasDeploySubData) *NullableEasDeploySubData {
	return &NullableEasDeploySubData{value: val, isSet: true}
}

func (v NullableEasDeploySubData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDeploySubData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


