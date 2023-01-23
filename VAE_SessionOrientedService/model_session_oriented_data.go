/*
VAE_SessionOrientedService

API for VAE_SessionOrientedService   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_SessionOrientedService

import (
	"encoding/json"
)

// SessionOrientedData Represents an Individual Session Oriented Service Subscription resource.
type SessionOrientedData struct {
	// Represents the identifier of the V2X UE.
	UeId string `json:"ueId"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
	// Represents the V2X service ID to which a V2X message belongs.
	ServiceId string `json:"serviceId"`
	// Represents the V2X application specific server identifier.
	AppSerId string `json:"appSerId"`
	AppQosReq *AppplicationQosRequirement `json:"appQosReq,omitempty"`
	// Set to true by the NF service consumer to request the VAE server to send a test notification as defined in clause 6.3.5.3. Set to false or omitted otherwise. 
	RequestTestNotification *bool `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewSessionOrientedData instantiates a new SessionOrientedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionOrientedData(ueId string, notifUri string, serviceId string, appSerId string) *SessionOrientedData {
	this := SessionOrientedData{}
	this.UeId = ueId
	this.NotifUri = notifUri
	this.ServiceId = serviceId
	this.AppSerId = appSerId
	return &this
}

// NewSessionOrientedDataWithDefaults instantiates a new SessionOrientedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionOrientedDataWithDefaults() *SessionOrientedData {
	this := SessionOrientedData{}
	return &this
}

// GetUeId returns the UeId field value
func (o *SessionOrientedData) GetUeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetUeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UeId, true
}

// SetUeId sets field value
func (o *SessionOrientedData) SetUeId(v string) {
	o.UeId = v
}

// GetNotifUri returns the NotifUri field value
func (o *SessionOrientedData) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *SessionOrientedData) SetNotifUri(v string) {
	o.NotifUri = v
}

// GetServiceId returns the ServiceId field value
func (o *SessionOrientedData) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetServiceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *SessionOrientedData) SetServiceId(v string) {
	o.ServiceId = v
}

// GetAppSerId returns the AppSerId field value
func (o *SessionOrientedData) GetAppSerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppSerId
}

// GetAppSerIdOk returns a tuple with the AppSerId field value
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetAppSerIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppSerId, true
}

// SetAppSerId sets field value
func (o *SessionOrientedData) SetAppSerId(v string) {
	o.AppSerId = v
}

// GetAppQosReq returns the AppQosReq field value if set, zero value otherwise.
func (o *SessionOrientedData) GetAppQosReq() AppplicationQosRequirement {
	if o == nil || isNil(o.AppQosReq) {
		var ret AppplicationQosRequirement
		return ret
	}
	return *o.AppQosReq
}

// GetAppQosReqOk returns a tuple with the AppQosReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetAppQosReqOk() (*AppplicationQosRequirement, bool) {
	if o == nil || isNil(o.AppQosReq) {
    return nil, false
	}
	return o.AppQosReq, true
}

// HasAppQosReq returns a boolean if a field has been set.
func (o *SessionOrientedData) HasAppQosReq() bool {
	if o != nil && !isNil(o.AppQosReq) {
		return true
	}

	return false
}

// SetAppQosReq gets a reference to the given AppplicationQosRequirement and assigns it to the AppQosReq field.
func (o *SessionOrientedData) SetAppQosReq(v AppplicationQosRequirement) {
	o.AppQosReq = &v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *SessionOrientedData) GetRequestTestNotification() bool {
	if o == nil || isNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || isNil(o.RequestTestNotification) {
    return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *SessionOrientedData) HasRequestTestNotification() bool {
	if o != nil && !isNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *SessionOrientedData) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *SessionOrientedData) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || isNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || isNil(o.WebsockNotifConfig) {
    return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *SessionOrientedData) HasWebsockNotifConfig() bool {
	if o != nil && !isNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *SessionOrientedData) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *SessionOrientedData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionOrientedData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *SessionOrientedData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *SessionOrientedData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o SessionOrientedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ueId"] = o.UeId
	}
	if true {
		toSerialize["notifUri"] = o.NotifUri
	}
	if true {
		toSerialize["serviceId"] = o.ServiceId
	}
	if true {
		toSerialize["appSerId"] = o.AppSerId
	}
	if !isNil(o.AppQosReq) {
		toSerialize["appQosReq"] = o.AppQosReq
	}
	if !isNil(o.RequestTestNotification) {
		toSerialize["requestTestNotification"] = o.RequestTestNotification
	}
	if !isNil(o.WebsockNotifConfig) {
		toSerialize["websockNotifConfig"] = o.WebsockNotifConfig
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableSessionOrientedData struct {
	value *SessionOrientedData
	isSet bool
}

func (v NullableSessionOrientedData) Get() *SessionOrientedData {
	return v.value
}

func (v *NullableSessionOrientedData) Set(val *SessionOrientedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionOrientedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionOrientedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionOrientedData(val *SessionOrientedData) *NullableSessionOrientedData {
	return &NullableSessionOrientedData{value: val, isSet: true}
}

func (v NullableSessionOrientedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionOrientedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


