/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"encoding/json"
)

// TimeSyncExposureConfig Contains the Time Synchronization Configuration parameters.
type TimeSyncExposureConfig struct {
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UpNodeId int32 `json:"upNodeId"`
	ReqPtpIns PtpInstance `json:"reqPtpIns"`
	// Indicates that the AF requests 5GS to act as a grandmaster for PTP or gPTP if it is included and set to true. 
	GmEnable *bool `json:"gmEnable,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	GmPrio *int32 `json:"gmPrio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TimeDom int32 `json:"timeDom"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TimeSyncErrBdgt *int32 `json:"timeSyncErrBdgt,omitempty"`
	// Notification Correlation ID assigned by the NF service consumer.
	ConfigNotifId string `json:"configNotifId"`
	// String providing an URI formatted according to RFC 3986.
	ConfigNotifUri string `json:"configNotifUri"`
	TempValidity *TemporalValidity `json:"tempValidity,omitempty"`
}

// NewTimeSyncExposureConfig instantiates a new TimeSyncExposureConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSyncExposureConfig(upNodeId int32, reqPtpIns PtpInstance, timeDom int32, configNotifId string, configNotifUri string) *TimeSyncExposureConfig {
	this := TimeSyncExposureConfig{}
	this.UpNodeId = upNodeId
	this.ReqPtpIns = reqPtpIns
	this.TimeDom = timeDom
	this.ConfigNotifId = configNotifId
	this.ConfigNotifUri = configNotifUri
	return &this
}

// NewTimeSyncExposureConfigWithDefaults instantiates a new TimeSyncExposureConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSyncExposureConfigWithDefaults() *TimeSyncExposureConfig {
	this := TimeSyncExposureConfig{}
	return &this
}

// GetUpNodeId returns the UpNodeId field value
func (o *TimeSyncExposureConfig) GetUpNodeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpNodeId
}

// GetUpNodeIdOk returns a tuple with the UpNodeId field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetUpNodeIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpNodeId, true
}

// SetUpNodeId sets field value
func (o *TimeSyncExposureConfig) SetUpNodeId(v int32) {
	o.UpNodeId = v
}

// GetReqPtpIns returns the ReqPtpIns field value
func (o *TimeSyncExposureConfig) GetReqPtpIns() PtpInstance {
	if o == nil {
		var ret PtpInstance
		return ret
	}

	return o.ReqPtpIns
}

// GetReqPtpInsOk returns a tuple with the ReqPtpIns field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetReqPtpInsOk() (*PtpInstance, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReqPtpIns, true
}

// SetReqPtpIns sets field value
func (o *TimeSyncExposureConfig) SetReqPtpIns(v PtpInstance) {
	o.ReqPtpIns = v
}

// GetGmEnable returns the GmEnable field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig) GetGmEnable() bool {
	if o == nil || isNil(o.GmEnable) {
		var ret bool
		return ret
	}
	return *o.GmEnable
}

// GetGmEnableOk returns a tuple with the GmEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetGmEnableOk() (*bool, bool) {
	if o == nil || isNil(o.GmEnable) {
    return nil, false
	}
	return o.GmEnable, true
}

// HasGmEnable returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig) HasGmEnable() bool {
	if o != nil && !isNil(o.GmEnable) {
		return true
	}

	return false
}

// SetGmEnable gets a reference to the given bool and assigns it to the GmEnable field.
func (o *TimeSyncExposureConfig) SetGmEnable(v bool) {
	o.GmEnable = &v
}

// GetGmPrio returns the GmPrio field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig) GetGmPrio() int32 {
	if o == nil || isNil(o.GmPrio) {
		var ret int32
		return ret
	}
	return *o.GmPrio
}

// GetGmPrioOk returns a tuple with the GmPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetGmPrioOk() (*int32, bool) {
	if o == nil || isNil(o.GmPrio) {
    return nil, false
	}
	return o.GmPrio, true
}

// HasGmPrio returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig) HasGmPrio() bool {
	if o != nil && !isNil(o.GmPrio) {
		return true
	}

	return false
}

// SetGmPrio gets a reference to the given int32 and assigns it to the GmPrio field.
func (o *TimeSyncExposureConfig) SetGmPrio(v int32) {
	o.GmPrio = &v
}

// GetTimeDom returns the TimeDom field value
func (o *TimeSyncExposureConfig) GetTimeDom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeDom
}

// GetTimeDomOk returns a tuple with the TimeDom field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetTimeDomOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeDom, true
}

// SetTimeDom sets field value
func (o *TimeSyncExposureConfig) SetTimeDom(v int32) {
	o.TimeDom = v
}

// GetTimeSyncErrBdgt returns the TimeSyncErrBdgt field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig) GetTimeSyncErrBdgt() int32 {
	if o == nil || isNil(o.TimeSyncErrBdgt) {
		var ret int32
		return ret
	}
	return *o.TimeSyncErrBdgt
}

// GetTimeSyncErrBdgtOk returns a tuple with the TimeSyncErrBdgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetTimeSyncErrBdgtOk() (*int32, bool) {
	if o == nil || isNil(o.TimeSyncErrBdgt) {
    return nil, false
	}
	return o.TimeSyncErrBdgt, true
}

// HasTimeSyncErrBdgt returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig) HasTimeSyncErrBdgt() bool {
	if o != nil && !isNil(o.TimeSyncErrBdgt) {
		return true
	}

	return false
}

// SetTimeSyncErrBdgt gets a reference to the given int32 and assigns it to the TimeSyncErrBdgt field.
func (o *TimeSyncExposureConfig) SetTimeSyncErrBdgt(v int32) {
	o.TimeSyncErrBdgt = &v
}

// GetConfigNotifId returns the ConfigNotifId field value
func (o *TimeSyncExposureConfig) GetConfigNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigNotifId
}

// GetConfigNotifIdOk returns a tuple with the ConfigNotifId field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetConfigNotifIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConfigNotifId, true
}

// SetConfigNotifId sets field value
func (o *TimeSyncExposureConfig) SetConfigNotifId(v string) {
	o.ConfigNotifId = v
}

// GetConfigNotifUri returns the ConfigNotifUri field value
func (o *TimeSyncExposureConfig) GetConfigNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigNotifUri
}

// GetConfigNotifUriOk returns a tuple with the ConfigNotifUri field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetConfigNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConfigNotifUri, true
}

// SetConfigNotifUri sets field value
func (o *TimeSyncExposureConfig) SetConfigNotifUri(v string) {
	o.ConfigNotifUri = v
}

// GetTempValidity returns the TempValidity field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig) GetTempValidity() TemporalValidity {
	if o == nil || isNil(o.TempValidity) {
		var ret TemporalValidity
		return ret
	}
	return *o.TempValidity
}

// GetTempValidityOk returns a tuple with the TempValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig) GetTempValidityOk() (*TemporalValidity, bool) {
	if o == nil || isNil(o.TempValidity) {
    return nil, false
	}
	return o.TempValidity, true
}

// HasTempValidity returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig) HasTempValidity() bool {
	if o != nil && !isNil(o.TempValidity) {
		return true
	}

	return false
}

// SetTempValidity gets a reference to the given TemporalValidity and assigns it to the TempValidity field.
func (o *TimeSyncExposureConfig) SetTempValidity(v TemporalValidity) {
	o.TempValidity = &v
}

func (o TimeSyncExposureConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["upNodeId"] = o.UpNodeId
	}
	if true {
		toSerialize["reqPtpIns"] = o.ReqPtpIns
	}
	if !isNil(o.GmEnable) {
		toSerialize["gmEnable"] = o.GmEnable
	}
	if !isNil(o.GmPrio) {
		toSerialize["gmPrio"] = o.GmPrio
	}
	if true {
		toSerialize["timeDom"] = o.TimeDom
	}
	if !isNil(o.TimeSyncErrBdgt) {
		toSerialize["timeSyncErrBdgt"] = o.TimeSyncErrBdgt
	}
	if true {
		toSerialize["configNotifId"] = o.ConfigNotifId
	}
	if true {
		toSerialize["configNotifUri"] = o.ConfigNotifUri
	}
	if !isNil(o.TempValidity) {
		toSerialize["tempValidity"] = o.TempValidity
	}
	return json.Marshal(toSerialize)
}

type NullableTimeSyncExposureConfig struct {
	value *TimeSyncExposureConfig
	isSet bool
}

func (v NullableTimeSyncExposureConfig) Get() *TimeSyncExposureConfig {
	return v.value
}

func (v *NullableTimeSyncExposureConfig) Set(val *TimeSyncExposureConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSyncExposureConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSyncExposureConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSyncExposureConfig(val *TimeSyncExposureConfig) *NullableTimeSyncExposureConfig {
	return &NullableTimeSyncExposureConfig{value: val, isSet: true}
}

func (v NullableTimeSyncExposureConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSyncExposureConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


