/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_NetworkResourceAdaptation

import (
	"encoding/json"
	"time"
)

// UnicastSubscription Represents a unicast subscription.
type UnicastSubscription struct {
	ValTgtUe ValTargetUe `json:"valTgtUe"`
	UniQosReq *string `json:"uniQosReq,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Duration *time.Time `json:"duration,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
	ReqTestNotif *bool `json:"reqTestNotif,omitempty"`
	WsNotifCfg *WebsockNotifConfig `json:"wsNotifCfg,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewUnicastSubscription instantiates a new UnicastSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnicastSubscription(valTgtUe ValTargetUe, notifUri string) *UnicastSubscription {
	this := UnicastSubscription{}
	this.ValTgtUe = valTgtUe
	this.NotifUri = notifUri
	return &this
}

// NewUnicastSubscriptionWithDefaults instantiates a new UnicastSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnicastSubscriptionWithDefaults() *UnicastSubscription {
	this := UnicastSubscription{}
	return &this
}

// GetValTgtUe returns the ValTgtUe field value
func (o *UnicastSubscription) GetValTgtUe() ValTargetUe {
	if o == nil {
		var ret ValTargetUe
		return ret
	}

	return o.ValTgtUe
}

// GetValTgtUeOk returns a tuple with the ValTgtUe field value
// and a boolean to check if the value has been set.
func (o *UnicastSubscription) GetValTgtUeOk() (*ValTargetUe, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ValTgtUe, true
}

// SetValTgtUe sets field value
func (o *UnicastSubscription) SetValTgtUe(v ValTargetUe) {
	o.ValTgtUe = v
}

// GetUniQosReq returns the UniQosReq field value if set, zero value otherwise.
func (o *UnicastSubscription) GetUniQosReq() string {
	if o == nil || isNil(o.UniQosReq) {
		var ret string
		return ret
	}
	return *o.UniQosReq
}

// GetUniQosReqOk returns a tuple with the UniQosReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnicastSubscription) GetUniQosReqOk() (*string, bool) {
	if o == nil || isNil(o.UniQosReq) {
    return nil, false
	}
	return o.UniQosReq, true
}

// HasUniQosReq returns a boolean if a field has been set.
func (o *UnicastSubscription) HasUniQosReq() bool {
	if o != nil && !isNil(o.UniQosReq) {
		return true
	}

	return false
}

// SetUniQosReq gets a reference to the given string and assigns it to the UniQosReq field.
func (o *UnicastSubscription) SetUniQosReq(v string) {
	o.UniQosReq = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UnicastSubscription) GetDuration() time.Time {
	if o == nil || isNil(o.Duration) {
		var ret time.Time
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnicastSubscription) GetDurationOk() (*time.Time, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UnicastSubscription) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given time.Time and assigns it to the Duration field.
func (o *UnicastSubscription) SetDuration(v time.Time) {
	o.Duration = &v
}

// GetNotifUri returns the NotifUri field value
func (o *UnicastSubscription) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *UnicastSubscription) GetNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *UnicastSubscription) SetNotifUri(v string) {
	o.NotifUri = v
}

// GetReqTestNotif returns the ReqTestNotif field value if set, zero value otherwise.
func (o *UnicastSubscription) GetReqTestNotif() bool {
	if o == nil || isNil(o.ReqTestNotif) {
		var ret bool
		return ret
	}
	return *o.ReqTestNotif
}

// GetReqTestNotifOk returns a tuple with the ReqTestNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnicastSubscription) GetReqTestNotifOk() (*bool, bool) {
	if o == nil || isNil(o.ReqTestNotif) {
    return nil, false
	}
	return o.ReqTestNotif, true
}

// HasReqTestNotif returns a boolean if a field has been set.
func (o *UnicastSubscription) HasReqTestNotif() bool {
	if o != nil && !isNil(o.ReqTestNotif) {
		return true
	}

	return false
}

// SetReqTestNotif gets a reference to the given bool and assigns it to the ReqTestNotif field.
func (o *UnicastSubscription) SetReqTestNotif(v bool) {
	o.ReqTestNotif = &v
}

// GetWsNotifCfg returns the WsNotifCfg field value if set, zero value otherwise.
func (o *UnicastSubscription) GetWsNotifCfg() WebsockNotifConfig {
	if o == nil || isNil(o.WsNotifCfg) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WsNotifCfg
}

// GetWsNotifCfgOk returns a tuple with the WsNotifCfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnicastSubscription) GetWsNotifCfgOk() (*WebsockNotifConfig, bool) {
	if o == nil || isNil(o.WsNotifCfg) {
    return nil, false
	}
	return o.WsNotifCfg, true
}

// HasWsNotifCfg returns a boolean if a field has been set.
func (o *UnicastSubscription) HasWsNotifCfg() bool {
	if o != nil && !isNil(o.WsNotifCfg) {
		return true
	}

	return false
}

// SetWsNotifCfg gets a reference to the given WebsockNotifConfig and assigns it to the WsNotifCfg field.
func (o *UnicastSubscription) SetWsNotifCfg(v WebsockNotifConfig) {
	o.WsNotifCfg = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *UnicastSubscription) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnicastSubscription) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *UnicastSubscription) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *UnicastSubscription) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o UnicastSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["valTgtUe"] = o.ValTgtUe
	}
	if !isNil(o.UniQosReq) {
		toSerialize["uniQosReq"] = o.UniQosReq
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["notifUri"] = o.NotifUri
	}
	if !isNil(o.ReqTestNotif) {
		toSerialize["reqTestNotif"] = o.ReqTestNotif
	}
	if !isNil(o.WsNotifCfg) {
		toSerialize["wsNotifCfg"] = o.WsNotifCfg
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableUnicastSubscription struct {
	value *UnicastSubscription
	isSet bool
}

func (v NullableUnicastSubscription) Get() *UnicastSubscription {
	return v.value
}

func (v *NullableUnicastSubscription) Set(val *UnicastSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableUnicastSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableUnicastSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnicastSubscription(val *UnicastSubscription) *NullableUnicastSubscription {
	return &NullableUnicastSubscription{value: val, isSet: true}
}

func (v NullableUnicastSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnicastSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

