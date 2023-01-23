/*
Eecs_ServiceProvisioning

API for ECS Service Provisioning. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_ServiceProvisioning

import (
	"encoding/json"
)

// EESInfo Represents EES information.
type EESInfo struct {
	// Identity of the EES
	EesId string `json:"eesId"`
	EndPt *EndPoint `json:"endPt,omitempty"`
	// Application identities of the Edge Application Servers registered with the EES.
	EasIds []string `json:"easIds,omitempty"`
	// Represents an ECSP Information.
	EcspInfo *string `json:"ecspInfo,omitempty"`
	SvcArea *LocationArea5G `json:"svcArea,omitempty"`
	// Represents list of Data network access identifier.
	Dnais []string `json:"dnais,omitempty"`
	// Indicates if the EES supports service continuity or not, also indicates which ACR scenarios are supported by the EES. 
	EesSvcContSupp []ACRScenario `json:"eesSvcContSupp,omitempty"`
	// Indicates whether the EEC is required to register on the EES to use edge services or not. 
	EecRegConf bool `json:"eecRegConf"`
}

// NewEESInfo instantiates a new EESInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEESInfo(eesId string, eecRegConf bool) *EESInfo {
	this := EESInfo{}
	this.EesId = eesId
	this.EecRegConf = eecRegConf
	return &this
}

// NewEESInfoWithDefaults instantiates a new EESInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEESInfoWithDefaults() *EESInfo {
	this := EESInfo{}
	return &this
}

// GetEesId returns the EesId field value
func (o *EESInfo) GetEesId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EesId
}

// GetEesIdOk returns a tuple with the EesId field value
// and a boolean to check if the value has been set.
func (o *EESInfo) GetEesIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EesId, true
}

// SetEesId sets field value
func (o *EESInfo) SetEesId(v string) {
	o.EesId = v
}

// GetEndPt returns the EndPt field value if set, zero value otherwise.
func (o *EESInfo) GetEndPt() EndPoint {
	if o == nil || isNil(o.EndPt) {
		var ret EndPoint
		return ret
	}
	return *o.EndPt
}

// GetEndPtOk returns a tuple with the EndPt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESInfo) GetEndPtOk() (*EndPoint, bool) {
	if o == nil || isNil(o.EndPt) {
    return nil, false
	}
	return o.EndPt, true
}

// HasEndPt returns a boolean if a field has been set.
func (o *EESInfo) HasEndPt() bool {
	if o != nil && !isNil(o.EndPt) {
		return true
	}

	return false
}

// SetEndPt gets a reference to the given EndPoint and assigns it to the EndPt field.
func (o *EESInfo) SetEndPt(v EndPoint) {
	o.EndPt = &v
}

// GetEasIds returns the EasIds field value if set, zero value otherwise.
func (o *EESInfo) GetEasIds() []string {
	if o == nil || isNil(o.EasIds) {
		var ret []string
		return ret
	}
	return o.EasIds
}

// GetEasIdsOk returns a tuple with the EasIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESInfo) GetEasIdsOk() ([]string, bool) {
	if o == nil || isNil(o.EasIds) {
    return nil, false
	}
	return o.EasIds, true
}

// HasEasIds returns a boolean if a field has been set.
func (o *EESInfo) HasEasIds() bool {
	if o != nil && !isNil(o.EasIds) {
		return true
	}

	return false
}

// SetEasIds gets a reference to the given []string and assigns it to the EasIds field.
func (o *EESInfo) SetEasIds(v []string) {
	o.EasIds = v
}

// GetEcspInfo returns the EcspInfo field value if set, zero value otherwise.
func (o *EESInfo) GetEcspInfo() string {
	if o == nil || isNil(o.EcspInfo) {
		var ret string
		return ret
	}
	return *o.EcspInfo
}

// GetEcspInfoOk returns a tuple with the EcspInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESInfo) GetEcspInfoOk() (*string, bool) {
	if o == nil || isNil(o.EcspInfo) {
    return nil, false
	}
	return o.EcspInfo, true
}

// HasEcspInfo returns a boolean if a field has been set.
func (o *EESInfo) HasEcspInfo() bool {
	if o != nil && !isNil(o.EcspInfo) {
		return true
	}

	return false
}

// SetEcspInfo gets a reference to the given string and assigns it to the EcspInfo field.
func (o *EESInfo) SetEcspInfo(v string) {
	o.EcspInfo = &v
}

// GetSvcArea returns the SvcArea field value if set, zero value otherwise.
func (o *EESInfo) GetSvcArea() LocationArea5G {
	if o == nil || isNil(o.SvcArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.SvcArea
}

// GetSvcAreaOk returns a tuple with the SvcArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESInfo) GetSvcAreaOk() (*LocationArea5G, bool) {
	if o == nil || isNil(o.SvcArea) {
    return nil, false
	}
	return o.SvcArea, true
}

// HasSvcArea returns a boolean if a field has been set.
func (o *EESInfo) HasSvcArea() bool {
	if o != nil && !isNil(o.SvcArea) {
		return true
	}

	return false
}

// SetSvcArea gets a reference to the given LocationArea5G and assigns it to the SvcArea field.
func (o *EESInfo) SetSvcArea(v LocationArea5G) {
	o.SvcArea = &v
}

// GetDnais returns the Dnais field value if set, zero value otherwise.
func (o *EESInfo) GetDnais() []string {
	if o == nil || isNil(o.Dnais) {
		var ret []string
		return ret
	}
	return o.Dnais
}

// GetDnaisOk returns a tuple with the Dnais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESInfo) GetDnaisOk() ([]string, bool) {
	if o == nil || isNil(o.Dnais) {
    return nil, false
	}
	return o.Dnais, true
}

// HasDnais returns a boolean if a field has been set.
func (o *EESInfo) HasDnais() bool {
	if o != nil && !isNil(o.Dnais) {
		return true
	}

	return false
}

// SetDnais gets a reference to the given []string and assigns it to the Dnais field.
func (o *EESInfo) SetDnais(v []string) {
	o.Dnais = v
}

// GetEesSvcContSupp returns the EesSvcContSupp field value if set, zero value otherwise.
func (o *EESInfo) GetEesSvcContSupp() []ACRScenario {
	if o == nil || isNil(o.EesSvcContSupp) {
		var ret []ACRScenario
		return ret
	}
	return o.EesSvcContSupp
}

// GetEesSvcContSuppOk returns a tuple with the EesSvcContSupp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESInfo) GetEesSvcContSuppOk() ([]ACRScenario, bool) {
	if o == nil || isNil(o.EesSvcContSupp) {
    return nil, false
	}
	return o.EesSvcContSupp, true
}

// HasEesSvcContSupp returns a boolean if a field has been set.
func (o *EESInfo) HasEesSvcContSupp() bool {
	if o != nil && !isNil(o.EesSvcContSupp) {
		return true
	}

	return false
}

// SetEesSvcContSupp gets a reference to the given []ACRScenario and assigns it to the EesSvcContSupp field.
func (o *EESInfo) SetEesSvcContSupp(v []ACRScenario) {
	o.EesSvcContSupp = v
}

// GetEecRegConf returns the EecRegConf field value
func (o *EESInfo) GetEecRegConf() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EecRegConf
}

// GetEecRegConfOk returns a tuple with the EecRegConf field value
// and a boolean to check if the value has been set.
func (o *EESInfo) GetEecRegConfOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EecRegConf, true
}

// SetEecRegConf sets field value
func (o *EESInfo) SetEecRegConf(v bool) {
	o.EecRegConf = v
}

func (o EESInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eesId"] = o.EesId
	}
	if !isNil(o.EndPt) {
		toSerialize["endPt"] = o.EndPt
	}
	if !isNil(o.EasIds) {
		toSerialize["easIds"] = o.EasIds
	}
	if !isNil(o.EcspInfo) {
		toSerialize["ecspInfo"] = o.EcspInfo
	}
	if !isNil(o.SvcArea) {
		toSerialize["svcArea"] = o.SvcArea
	}
	if !isNil(o.Dnais) {
		toSerialize["dnais"] = o.Dnais
	}
	if !isNil(o.EesSvcContSupp) {
		toSerialize["eesSvcContSupp"] = o.EesSvcContSupp
	}
	if true {
		toSerialize["eecRegConf"] = o.EecRegConf
	}
	return json.Marshal(toSerialize)
}

type NullableEESInfo struct {
	value *EESInfo
	isSet bool
}

func (v NullableEESInfo) Get() *EESInfo {
	return v.value
}

func (v *NullableEESInfo) Set(val *EESInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEESInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEESInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEESInfo(val *EESInfo) *NullableEESInfo {
	return &NullableEESInfo{value: val, isSet: true}
}

func (v NullableEESInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEESInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


