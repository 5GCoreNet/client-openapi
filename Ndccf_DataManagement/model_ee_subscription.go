/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// EeSubscription struct for EeSubscription
type EeSubscription struct {
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
	// A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations
	MonitoringConfigurations map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ReportingOptions *ReportingOptions1 `json:"reportingOptions,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	EpcAppliedInd *bool `json:"epcAppliedInd,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamHost *string `json:"scefDiamHost,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamRealm *string `json:"scefDiamRealm,omitempty"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SecondCallbackRef *string `json:"secondCallbackRef,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList []string `json:"includeGpsiList,omitempty"`
}

// NewEeSubscription instantiates a new EeSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEeSubscription(callbackReference string, monitoringConfigurations map[string]MonitoringConfiguration) *EeSubscription {
	this := EeSubscription{}
	this.CallbackReference = callbackReference
	this.MonitoringConfigurations = monitoringConfigurations
	var epcAppliedInd bool = false
	this.EpcAppliedInd = &epcAppliedInd
	return &this
}

// NewEeSubscriptionWithDefaults instantiates a new EeSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEeSubscriptionWithDefaults() *EeSubscription {
	this := EeSubscription{}
	var epcAppliedInd bool = false
	this.EpcAppliedInd = &epcAppliedInd
	return &this
}

// GetCallbackReference returns the CallbackReference field value
func (o *EeSubscription) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *EeSubscription) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetMonitoringConfigurations returns the MonitoringConfigurations field value
func (o *EeSubscription) GetMonitoringConfigurations() map[string]MonitoringConfiguration {
	if o == nil {
		var ret map[string]MonitoringConfiguration
		return ret
	}

	return o.MonitoringConfigurations
}

// GetMonitoringConfigurationsOk returns a tuple with the MonitoringConfigurations field value
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetMonitoringConfigurationsOk() (*map[string]MonitoringConfiguration, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MonitoringConfigurations, true
}

// SetMonitoringConfigurations sets field value
func (o *EeSubscription) SetMonitoringConfigurations(v map[string]MonitoringConfiguration) {
	o.MonitoringConfigurations = v
}

// GetReportingOptions returns the ReportingOptions field value if set, zero value otherwise.
func (o *EeSubscription) GetReportingOptions() ReportingOptions1 {
	if o == nil || isNil(o.ReportingOptions) {
		var ret ReportingOptions1
		return ret
	}
	return *o.ReportingOptions
}

// GetReportingOptionsOk returns a tuple with the ReportingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetReportingOptionsOk() (*ReportingOptions1, bool) {
	if o == nil || isNil(o.ReportingOptions) {
    return nil, false
	}
	return o.ReportingOptions, true
}

// HasReportingOptions returns a boolean if a field has been set.
func (o *EeSubscription) HasReportingOptions() bool {
	if o != nil && !isNil(o.ReportingOptions) {
		return true
	}

	return false
}

// SetReportingOptions gets a reference to the given ReportingOptions1 and assigns it to the ReportingOptions field.
func (o *EeSubscription) SetReportingOptions(v ReportingOptions1) {
	o.ReportingOptions = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EeSubscription) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EeSubscription) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EeSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *EeSubscription) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
    return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *EeSubscription) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *EeSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *EeSubscription) GetContextInfo() ContextInfo {
	if o == nil || isNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || isNil(o.ContextInfo) {
    return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *EeSubscription) HasContextInfo() bool {
	if o != nil && !isNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *EeSubscription) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

// GetEpcAppliedInd returns the EpcAppliedInd field value if set, zero value otherwise.
func (o *EeSubscription) GetEpcAppliedInd() bool {
	if o == nil || isNil(o.EpcAppliedInd) {
		var ret bool
		return ret
	}
	return *o.EpcAppliedInd
}

// GetEpcAppliedIndOk returns a tuple with the EpcAppliedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetEpcAppliedIndOk() (*bool, bool) {
	if o == nil || isNil(o.EpcAppliedInd) {
    return nil, false
	}
	return o.EpcAppliedInd, true
}

// HasEpcAppliedInd returns a boolean if a field has been set.
func (o *EeSubscription) HasEpcAppliedInd() bool {
	if o != nil && !isNil(o.EpcAppliedInd) {
		return true
	}

	return false
}

// SetEpcAppliedInd gets a reference to the given bool and assigns it to the EpcAppliedInd field.
func (o *EeSubscription) SetEpcAppliedInd(v bool) {
	o.EpcAppliedInd = &v
}

// GetScefDiamHost returns the ScefDiamHost field value if set, zero value otherwise.
func (o *EeSubscription) GetScefDiamHost() string {
	if o == nil || isNil(o.ScefDiamHost) {
		var ret string
		return ret
	}
	return *o.ScefDiamHost
}

// GetScefDiamHostOk returns a tuple with the ScefDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetScefDiamHostOk() (*string, bool) {
	if o == nil || isNil(o.ScefDiamHost) {
    return nil, false
	}
	return o.ScefDiamHost, true
}

// HasScefDiamHost returns a boolean if a field has been set.
func (o *EeSubscription) HasScefDiamHost() bool {
	if o != nil && !isNil(o.ScefDiamHost) {
		return true
	}

	return false
}

// SetScefDiamHost gets a reference to the given string and assigns it to the ScefDiamHost field.
func (o *EeSubscription) SetScefDiamHost(v string) {
	o.ScefDiamHost = &v
}

// GetScefDiamRealm returns the ScefDiamRealm field value if set, zero value otherwise.
func (o *EeSubscription) GetScefDiamRealm() string {
	if o == nil || isNil(o.ScefDiamRealm) {
		var ret string
		return ret
	}
	return *o.ScefDiamRealm
}

// GetScefDiamRealmOk returns a tuple with the ScefDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetScefDiamRealmOk() (*string, bool) {
	if o == nil || isNil(o.ScefDiamRealm) {
    return nil, false
	}
	return o.ScefDiamRealm, true
}

// HasScefDiamRealm returns a boolean if a field has been set.
func (o *EeSubscription) HasScefDiamRealm() bool {
	if o != nil && !isNil(o.ScefDiamRealm) {
		return true
	}

	return false
}

// SetScefDiamRealm gets a reference to the given string and assigns it to the ScefDiamRealm field.
func (o *EeSubscription) SetScefDiamRealm(v string) {
	o.ScefDiamRealm = &v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *EeSubscription) GetNotifyCorrelationId() string {
	if o == nil || isNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifyCorrelationId) {
    return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *EeSubscription) HasNotifyCorrelationId() bool {
	if o != nil && !isNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *EeSubscription) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

// GetSecondCallbackRef returns the SecondCallbackRef field value if set, zero value otherwise.
func (o *EeSubscription) GetSecondCallbackRef() string {
	if o == nil || isNil(o.SecondCallbackRef) {
		var ret string
		return ret
	}
	return *o.SecondCallbackRef
}

// GetSecondCallbackRefOk returns a tuple with the SecondCallbackRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetSecondCallbackRefOk() (*string, bool) {
	if o == nil || isNil(o.SecondCallbackRef) {
    return nil, false
	}
	return o.SecondCallbackRef, true
}

// HasSecondCallbackRef returns a boolean if a field has been set.
func (o *EeSubscription) HasSecondCallbackRef() bool {
	if o != nil && !isNil(o.SecondCallbackRef) {
		return true
	}

	return false
}

// SetSecondCallbackRef gets a reference to the given string and assigns it to the SecondCallbackRef field.
func (o *EeSubscription) SetSecondCallbackRef(v string) {
	o.SecondCallbackRef = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *EeSubscription) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *EeSubscription) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *EeSubscription) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetExcludeGpsiList returns the ExcludeGpsiList field value if set, zero value otherwise.
func (o *EeSubscription) GetExcludeGpsiList() []string {
	if o == nil || isNil(o.ExcludeGpsiList) {
		var ret []string
		return ret
	}
	return o.ExcludeGpsiList
}

// GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetExcludeGpsiListOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeGpsiList) {
    return nil, false
	}
	return o.ExcludeGpsiList, true
}

// HasExcludeGpsiList returns a boolean if a field has been set.
func (o *EeSubscription) HasExcludeGpsiList() bool {
	if o != nil && !isNil(o.ExcludeGpsiList) {
		return true
	}

	return false
}

// SetExcludeGpsiList gets a reference to the given []string and assigns it to the ExcludeGpsiList field.
func (o *EeSubscription) SetExcludeGpsiList(v []string) {
	o.ExcludeGpsiList = v
}

// GetIncludeGpsiList returns the IncludeGpsiList field value if set, zero value otherwise.
func (o *EeSubscription) GetIncludeGpsiList() []string {
	if o == nil || isNil(o.IncludeGpsiList) {
		var ret []string
		return ret
	}
	return o.IncludeGpsiList
}

// GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetIncludeGpsiListOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeGpsiList) {
    return nil, false
	}
	return o.IncludeGpsiList, true
}

// HasIncludeGpsiList returns a boolean if a field has been set.
func (o *EeSubscription) HasIncludeGpsiList() bool {
	if o != nil && !isNil(o.IncludeGpsiList) {
		return true
	}

	return false
}

// SetIncludeGpsiList gets a reference to the given []string and assigns it to the IncludeGpsiList field.
func (o *EeSubscription) SetIncludeGpsiList(v []string) {
	o.IncludeGpsiList = v
}

func (o EeSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["callbackReference"] = o.CallbackReference
	}
	if true {
		toSerialize["monitoringConfigurations"] = o.MonitoringConfigurations
	}
	if !isNil(o.ReportingOptions) {
		toSerialize["reportingOptions"] = o.ReportingOptions
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !isNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	if !isNil(o.EpcAppliedInd) {
		toSerialize["epcAppliedInd"] = o.EpcAppliedInd
	}
	if !isNil(o.ScefDiamHost) {
		toSerialize["scefDiamHost"] = o.ScefDiamHost
	}
	if !isNil(o.ScefDiamRealm) {
		toSerialize["scefDiamRealm"] = o.ScefDiamRealm
	}
	if !isNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	if !isNil(o.SecondCallbackRef) {
		toSerialize["secondCallbackRef"] = o.SecondCallbackRef
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.ExcludeGpsiList) {
		toSerialize["excludeGpsiList"] = o.ExcludeGpsiList
	}
	if !isNil(o.IncludeGpsiList) {
		toSerialize["includeGpsiList"] = o.IncludeGpsiList
	}
	return json.Marshal(toSerialize)
}

type NullableEeSubscription struct {
	value *EeSubscription
	isSet bool
}

func (v NullableEeSubscription) Get() *EeSubscription {
	return v.value
}

func (v *NullableEeSubscription) Set(val *EeSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableEeSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableEeSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEeSubscription(val *EeSubscription) *NullableEeSubscription {
	return &NullableEeSubscription{value: val, isSet: true}
}

func (v NullableEeSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEeSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


