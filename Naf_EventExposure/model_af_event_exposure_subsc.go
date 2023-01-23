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

// AfEventExposureSubsc Represents an Individual Application Event Subscription resource.
type AfEventExposureSubsc struct {
	DataAccProfId *string `json:"dataAccProfId,omitempty"`
	EventsSubs []EventsSubs `json:"eventsSubs"`
	EventsRepInfo ReportingInformation `json:"eventsRepInfo"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
	NotifId string `json:"notifId"`
	EventNotifs []AfEventNotification `json:"eventNotifs,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewAfEventExposureSubsc instantiates a new AfEventExposureSubsc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfEventExposureSubsc(eventsSubs []EventsSubs, eventsRepInfo ReportingInformation, notifUri string, notifId string) *AfEventExposureSubsc {
	this := AfEventExposureSubsc{}
	this.EventsSubs = eventsSubs
	this.EventsRepInfo = eventsRepInfo
	this.NotifUri = notifUri
	this.NotifId = notifId
	return &this
}

// NewAfEventExposureSubscWithDefaults instantiates a new AfEventExposureSubsc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfEventExposureSubscWithDefaults() *AfEventExposureSubsc {
	this := AfEventExposureSubsc{}
	return &this
}

// GetDataAccProfId returns the DataAccProfId field value if set, zero value otherwise.
func (o *AfEventExposureSubsc) GetDataAccProfId() string {
	if o == nil || isNil(o.DataAccProfId) {
		var ret string
		return ret
	}
	return *o.DataAccProfId
}

// GetDataAccProfIdOk returns a tuple with the DataAccProfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventExposureSubsc) GetDataAccProfIdOk() (*string, bool) {
	if o == nil || isNil(o.DataAccProfId) {
    return nil, false
	}
	return o.DataAccProfId, true
}

// HasDataAccProfId returns a boolean if a field has been set.
func (o *AfEventExposureSubsc) HasDataAccProfId() bool {
	if o != nil && !isNil(o.DataAccProfId) {
		return true
	}

	return false
}

// SetDataAccProfId gets a reference to the given string and assigns it to the DataAccProfId field.
func (o *AfEventExposureSubsc) SetDataAccProfId(v string) {
	o.DataAccProfId = &v
}

// GetEventsSubs returns the EventsSubs field value
func (o *AfEventExposureSubsc) GetEventsSubs() []EventsSubs {
	if o == nil {
		var ret []EventsSubs
		return ret
	}

	return o.EventsSubs
}

// GetEventsSubsOk returns a tuple with the EventsSubs field value
// and a boolean to check if the value has been set.
func (o *AfEventExposureSubsc) GetEventsSubsOk() ([]EventsSubs, bool) {
	if o == nil {
    return nil, false
	}
	return o.EventsSubs, true
}

// SetEventsSubs sets field value
func (o *AfEventExposureSubsc) SetEventsSubs(v []EventsSubs) {
	o.EventsSubs = v
}

// GetEventsRepInfo returns the EventsRepInfo field value
func (o *AfEventExposureSubsc) GetEventsRepInfo() ReportingInformation {
	if o == nil {
		var ret ReportingInformation
		return ret
	}

	return o.EventsRepInfo
}

// GetEventsRepInfoOk returns a tuple with the EventsRepInfo field value
// and a boolean to check if the value has been set.
func (o *AfEventExposureSubsc) GetEventsRepInfoOk() (*ReportingInformation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventsRepInfo, true
}

// SetEventsRepInfo sets field value
func (o *AfEventExposureSubsc) SetEventsRepInfo(v ReportingInformation) {
	o.EventsRepInfo = v
}

// GetNotifUri returns the NotifUri field value
func (o *AfEventExposureSubsc) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *AfEventExposureSubsc) GetNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *AfEventExposureSubsc) SetNotifUri(v string) {
	o.NotifUri = v
}

// GetNotifId returns the NotifId field value
func (o *AfEventExposureSubsc) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *AfEventExposureSubsc) GetNotifIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *AfEventExposureSubsc) SetNotifId(v string) {
	o.NotifId = v
}

// GetEventNotifs returns the EventNotifs field value if set, zero value otherwise.
func (o *AfEventExposureSubsc) GetEventNotifs() []AfEventNotification {
	if o == nil || isNil(o.EventNotifs) {
		var ret []AfEventNotification
		return ret
	}
	return o.EventNotifs
}

// GetEventNotifsOk returns a tuple with the EventNotifs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventExposureSubsc) GetEventNotifsOk() ([]AfEventNotification, bool) {
	if o == nil || isNil(o.EventNotifs) {
    return nil, false
	}
	return o.EventNotifs, true
}

// HasEventNotifs returns a boolean if a field has been set.
func (o *AfEventExposureSubsc) HasEventNotifs() bool {
	if o != nil && !isNil(o.EventNotifs) {
		return true
	}

	return false
}

// SetEventNotifs gets a reference to the given []AfEventNotification and assigns it to the EventNotifs field.
func (o *AfEventExposureSubsc) SetEventNotifs(v []AfEventNotification) {
	o.EventNotifs = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *AfEventExposureSubsc) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventExposureSubsc) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *AfEventExposureSubsc) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *AfEventExposureSubsc) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o AfEventExposureSubsc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DataAccProfId) {
		toSerialize["dataAccProfId"] = o.DataAccProfId
	}
	if true {
		toSerialize["eventsSubs"] = o.EventsSubs
	}
	if true {
		toSerialize["eventsRepInfo"] = o.EventsRepInfo
	}
	if true {
		toSerialize["notifUri"] = o.NotifUri
	}
	if true {
		toSerialize["notifId"] = o.NotifId
	}
	if !isNil(o.EventNotifs) {
		toSerialize["eventNotifs"] = o.EventNotifs
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableAfEventExposureSubsc struct {
	value *AfEventExposureSubsc
	isSet bool
}

func (v NullableAfEventExposureSubsc) Get() *AfEventExposureSubsc {
	return v.value
}

func (v *NullableAfEventExposureSubsc) Set(val *AfEventExposureSubsc) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventExposureSubsc) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventExposureSubsc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventExposureSubsc(val *AfEventExposureSubsc) *NullableAfEventExposureSubsc {
	return &NullableAfEventExposureSubsc{value: val, isSet: true}
}

func (v NullableAfEventExposureSubsc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventExposureSubsc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


