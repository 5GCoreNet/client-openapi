/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"time"
)

// SmsfRegistration struct for SmsfRegistration
type SmsfRegistration struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmsfInstanceId string `json:"smsfInstanceId"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	SmsfSetId *string `json:"smsfSetId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	PlmnId PlmnId `json:"plmnId"`
	SmsfMAPAddress *string `json:"smsfMAPAddress,omitempty"`
	SmsfDiameterAddress *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime *time.Time `json:"registrationTime,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri *string `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	SmsfSbiSupInd *bool `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd *bool `json:"udrRestartInd,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LastSynchronizationTime *time.Time `json:"lastSynchronizationTime,omitempty"`
}

// NewSmsfRegistration instantiates a new SmsfRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsfRegistration(smsfInstanceId string, plmnId PlmnId) *SmsfRegistration {
	this := SmsfRegistration{}
	this.SmsfInstanceId = smsfInstanceId
	this.PlmnId = plmnId
	var smsfSbiSupInd bool = false
	this.SmsfSbiSupInd = &smsfSbiSupInd
	var udrRestartInd bool = false
	this.UdrRestartInd = &udrRestartInd
	return &this
}

// NewSmsfRegistrationWithDefaults instantiates a new SmsfRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsfRegistrationWithDefaults() *SmsfRegistration {
	this := SmsfRegistration{}
	var smsfSbiSupInd bool = false
	this.SmsfSbiSupInd = &smsfSbiSupInd
	var udrRestartInd bool = false
	this.UdrRestartInd = &udrRestartInd
	return &this
}

// GetSmsfInstanceId returns the SmsfInstanceId field value
func (o *SmsfRegistration) GetSmsfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmsfInstanceId
}

// GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field value
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetSmsfInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SmsfInstanceId, true
}

// SetSmsfInstanceId sets field value
func (o *SmsfRegistration) SetSmsfInstanceId(v string) {
	o.SmsfInstanceId = v
}

// GetSmsfSetId returns the SmsfSetId field value if set, zero value otherwise.
func (o *SmsfRegistration) GetSmsfSetId() string {
	if o == nil || isNil(o.SmsfSetId) {
		var ret string
		return ret
	}
	return *o.SmsfSetId
}

// GetSmsfSetIdOk returns a tuple with the SmsfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetSmsfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.SmsfSetId) {
    return nil, false
	}
	return o.SmsfSetId, true
}

// HasSmsfSetId returns a boolean if a field has been set.
func (o *SmsfRegistration) HasSmsfSetId() bool {
	if o != nil && !isNil(o.SmsfSetId) {
		return true
	}

	return false
}

// SetSmsfSetId gets a reference to the given string and assigns it to the SmsfSetId field.
func (o *SmsfRegistration) SetSmsfSetId(v string) {
	o.SmsfSetId = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SmsfRegistration) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SmsfRegistration) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SmsfRegistration) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetPlmnId returns the PlmnId field value
func (o *SmsfRegistration) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *SmsfRegistration) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetSmsfMAPAddress returns the SmsfMAPAddress field value if set, zero value otherwise.
func (o *SmsfRegistration) GetSmsfMAPAddress() string {
	if o == nil || isNil(o.SmsfMAPAddress) {
		var ret string
		return ret
	}
	return *o.SmsfMAPAddress
}

// GetSmsfMAPAddressOk returns a tuple with the SmsfMAPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetSmsfMAPAddressOk() (*string, bool) {
	if o == nil || isNil(o.SmsfMAPAddress) {
    return nil, false
	}
	return o.SmsfMAPAddress, true
}

// HasSmsfMAPAddress returns a boolean if a field has been set.
func (o *SmsfRegistration) HasSmsfMAPAddress() bool {
	if o != nil && !isNil(o.SmsfMAPAddress) {
		return true
	}

	return false
}

// SetSmsfMAPAddress gets a reference to the given string and assigns it to the SmsfMAPAddress field.
func (o *SmsfRegistration) SetSmsfMAPAddress(v string) {
	o.SmsfMAPAddress = &v
}

// GetSmsfDiameterAddress returns the SmsfDiameterAddress field value if set, zero value otherwise.
func (o *SmsfRegistration) GetSmsfDiameterAddress() NetworkNodeDiameterAddress {
	if o == nil || isNil(o.SmsfDiameterAddress) {
		var ret NetworkNodeDiameterAddress
		return ret
	}
	return *o.SmsfDiameterAddress
}

// GetSmsfDiameterAddressOk returns a tuple with the SmsfDiameterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetSmsfDiameterAddressOk() (*NetworkNodeDiameterAddress, bool) {
	if o == nil || isNil(o.SmsfDiameterAddress) {
    return nil, false
	}
	return o.SmsfDiameterAddress, true
}

// HasSmsfDiameterAddress returns a boolean if a field has been set.
func (o *SmsfRegistration) HasSmsfDiameterAddress() bool {
	if o != nil && !isNil(o.SmsfDiameterAddress) {
		return true
	}

	return false
}

// SetSmsfDiameterAddress gets a reference to the given NetworkNodeDiameterAddress and assigns it to the SmsfDiameterAddress field.
func (o *SmsfRegistration) SetSmsfDiameterAddress(v NetworkNodeDiameterAddress) {
	o.SmsfDiameterAddress = &v
}

// GetRegistrationTime returns the RegistrationTime field value if set, zero value otherwise.
func (o *SmsfRegistration) GetRegistrationTime() time.Time {
	if o == nil || isNil(o.RegistrationTime) {
		var ret time.Time
		return ret
	}
	return *o.RegistrationTime
}

// GetRegistrationTimeOk returns a tuple with the RegistrationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetRegistrationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RegistrationTime) {
    return nil, false
	}
	return o.RegistrationTime, true
}

// HasRegistrationTime returns a boolean if a field has been set.
func (o *SmsfRegistration) HasRegistrationTime() bool {
	if o != nil && !isNil(o.RegistrationTime) {
		return true
	}

	return false
}

// SetRegistrationTime gets a reference to the given time.Time and assigns it to the RegistrationTime field.
func (o *SmsfRegistration) SetRegistrationTime(v time.Time) {
	o.RegistrationTime = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *SmsfRegistration) GetContextInfo() ContextInfo {
	if o == nil || isNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || isNil(o.ContextInfo) {
    return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *SmsfRegistration) HasContextInfo() bool {
	if o != nil && !isNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *SmsfRegistration) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

// GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field value if set, zero value otherwise.
func (o *SmsfRegistration) GetDataRestorationCallbackUri() string {
	if o == nil || isNil(o.DataRestorationCallbackUri) {
		var ret string
		return ret
	}
	return *o.DataRestorationCallbackUri
}

// GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetDataRestorationCallbackUriOk() (*string, bool) {
	if o == nil || isNil(o.DataRestorationCallbackUri) {
    return nil, false
	}
	return o.DataRestorationCallbackUri, true
}

// HasDataRestorationCallbackUri returns a boolean if a field has been set.
func (o *SmsfRegistration) HasDataRestorationCallbackUri() bool {
	if o != nil && !isNil(o.DataRestorationCallbackUri) {
		return true
	}

	return false
}

// SetDataRestorationCallbackUri gets a reference to the given string and assigns it to the DataRestorationCallbackUri field.
func (o *SmsfRegistration) SetDataRestorationCallbackUri(v string) {
	o.DataRestorationCallbackUri = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *SmsfRegistration) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *SmsfRegistration) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *SmsfRegistration) SetResetIds(v []string) {
	o.ResetIds = v
}

// GetSmsfSbiSupInd returns the SmsfSbiSupInd field value if set, zero value otherwise.
func (o *SmsfRegistration) GetSmsfSbiSupInd() bool {
	if o == nil || isNil(o.SmsfSbiSupInd) {
		var ret bool
		return ret
	}
	return *o.SmsfSbiSupInd
}

// GetSmsfSbiSupIndOk returns a tuple with the SmsfSbiSupInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetSmsfSbiSupIndOk() (*bool, bool) {
	if o == nil || isNil(o.SmsfSbiSupInd) {
    return nil, false
	}
	return o.SmsfSbiSupInd, true
}

// HasSmsfSbiSupInd returns a boolean if a field has been set.
func (o *SmsfRegistration) HasSmsfSbiSupInd() bool {
	if o != nil && !isNil(o.SmsfSbiSupInd) {
		return true
	}

	return false
}

// SetSmsfSbiSupInd gets a reference to the given bool and assigns it to the SmsfSbiSupInd field.
func (o *SmsfRegistration) SetSmsfSbiSupInd(v bool) {
	o.SmsfSbiSupInd = &v
}

// GetUdrRestartInd returns the UdrRestartInd field value if set, zero value otherwise.
func (o *SmsfRegistration) GetUdrRestartInd() bool {
	if o == nil || isNil(o.UdrRestartInd) {
		var ret bool
		return ret
	}
	return *o.UdrRestartInd
}

// GetUdrRestartIndOk returns a tuple with the UdrRestartInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetUdrRestartIndOk() (*bool, bool) {
	if o == nil || isNil(o.UdrRestartInd) {
    return nil, false
	}
	return o.UdrRestartInd, true
}

// HasUdrRestartInd returns a boolean if a field has been set.
func (o *SmsfRegistration) HasUdrRestartInd() bool {
	if o != nil && !isNil(o.UdrRestartInd) {
		return true
	}

	return false
}

// SetUdrRestartInd gets a reference to the given bool and assigns it to the UdrRestartInd field.
func (o *SmsfRegistration) SetUdrRestartInd(v bool) {
	o.UdrRestartInd = &v
}

// GetLastSynchronizationTime returns the LastSynchronizationTime field value if set, zero value otherwise.
func (o *SmsfRegistration) GetLastSynchronizationTime() time.Time {
	if o == nil || isNil(o.LastSynchronizationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastSynchronizationTime
}

// GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfRegistration) GetLastSynchronizationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastSynchronizationTime) {
    return nil, false
	}
	return o.LastSynchronizationTime, true
}

// HasLastSynchronizationTime returns a boolean if a field has been set.
func (o *SmsfRegistration) HasLastSynchronizationTime() bool {
	if o != nil && !isNil(o.LastSynchronizationTime) {
		return true
	}

	return false
}

// SetLastSynchronizationTime gets a reference to the given time.Time and assigns it to the LastSynchronizationTime field.
func (o *SmsfRegistration) SetLastSynchronizationTime(v time.Time) {
	o.LastSynchronizationTime = &v
}

func (o SmsfRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["smsfInstanceId"] = o.SmsfInstanceId
	}
	if !isNil(o.SmsfSetId) {
		toSerialize["smsfSetId"] = o.SmsfSetId
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.SmsfMAPAddress) {
		toSerialize["smsfMAPAddress"] = o.SmsfMAPAddress
	}
	if !isNil(o.SmsfDiameterAddress) {
		toSerialize["smsfDiameterAddress"] = o.SmsfDiameterAddress
	}
	if !isNil(o.RegistrationTime) {
		toSerialize["registrationTime"] = o.RegistrationTime
	}
	if !isNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	if !isNil(o.DataRestorationCallbackUri) {
		toSerialize["dataRestorationCallbackUri"] = o.DataRestorationCallbackUri
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	if !isNil(o.SmsfSbiSupInd) {
		toSerialize["smsfSbiSupInd"] = o.SmsfSbiSupInd
	}
	if !isNil(o.UdrRestartInd) {
		toSerialize["udrRestartInd"] = o.UdrRestartInd
	}
	if !isNil(o.LastSynchronizationTime) {
		toSerialize["lastSynchronizationTime"] = o.LastSynchronizationTime
	}
	return json.Marshal(toSerialize)
}

type NullableSmsfRegistration struct {
	value *SmsfRegistration
	isSet bool
}

func (v NullableSmsfRegistration) Get() *SmsfRegistration {
	return v.value
}

func (v *NullableSmsfRegistration) Set(val *SmsfRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsfRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsfRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsfRegistration(val *SmsfRegistration) *NullableSmsfRegistration {
	return &NullableSmsfRegistration{value: val, isSet: true}
}

func (v NullableSmsfRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsfRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


