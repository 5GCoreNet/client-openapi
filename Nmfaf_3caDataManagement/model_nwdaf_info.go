/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// NwdafInfo Information of a NWDAF NF Instance
type NwdafInfo struct {
	EventIds []EventId `json:"eventIds,omitempty"`
	NwdafEvents []NwdafEvent `json:"nwdafEvents,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	NwdafCapability *NwdafCapability `json:"nwdafCapability,omitempty"`
	// indicating a time in seconds.
	AnalyticsDelay *int32 `json:"analyticsDelay,omitempty"`
	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`
	ServingNfTypeList []NFType `json:"servingNfTypeList,omitempty"`
	MlAnalyticsList []MlAnalyticsInfo `json:"mlAnalyticsList,omitempty"`
}

// NewNwdafInfo instantiates a new NwdafInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafInfo() *NwdafInfo {
	this := NwdafInfo{}
	return &this
}

// NewNwdafInfoWithDefaults instantiates a new NwdafInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafInfoWithDefaults() *NwdafInfo {
	this := NwdafInfo{}
	return &this
}

// GetEventIds returns the EventIds field value if set, zero value otherwise.
func (o *NwdafInfo) GetEventIds() []EventId {
	if o == nil || isNil(o.EventIds) {
		var ret []EventId
		return ret
	}
	return o.EventIds
}

// GetEventIdsOk returns a tuple with the EventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetEventIdsOk() ([]EventId, bool) {
	if o == nil || isNil(o.EventIds) {
    return nil, false
	}
	return o.EventIds, true
}

// HasEventIds returns a boolean if a field has been set.
func (o *NwdafInfo) HasEventIds() bool {
	if o != nil && !isNil(o.EventIds) {
		return true
	}

	return false
}

// SetEventIds gets a reference to the given []EventId and assigns it to the EventIds field.
func (o *NwdafInfo) SetEventIds(v []EventId) {
	o.EventIds = v
}

// GetNwdafEvents returns the NwdafEvents field value if set, zero value otherwise.
func (o *NwdafInfo) GetNwdafEvents() []NwdafEvent {
	if o == nil || isNil(o.NwdafEvents) {
		var ret []NwdafEvent
		return ret
	}
	return o.NwdafEvents
}

// GetNwdafEventsOk returns a tuple with the NwdafEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetNwdafEventsOk() ([]NwdafEvent, bool) {
	if o == nil || isNil(o.NwdafEvents) {
    return nil, false
	}
	return o.NwdafEvents, true
}

// HasNwdafEvents returns a boolean if a field has been set.
func (o *NwdafInfo) HasNwdafEvents() bool {
	if o != nil && !isNil(o.NwdafEvents) {
		return true
	}

	return false
}

// SetNwdafEvents gets a reference to the given []NwdafEvent and assigns it to the NwdafEvents field.
func (o *NwdafInfo) SetNwdafEvents(v []NwdafEvent) {
	o.NwdafEvents = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NwdafInfo) GetTaiList() []Tai {
	if o == nil || isNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaiList) {
    return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NwdafInfo) HasTaiList() bool {
	if o != nil && !isNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NwdafInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NwdafInfo) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NwdafInfo) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NwdafInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetNwdafCapability returns the NwdafCapability field value if set, zero value otherwise.
func (o *NwdafInfo) GetNwdafCapability() NwdafCapability {
	if o == nil || isNil(o.NwdafCapability) {
		var ret NwdafCapability
		return ret
	}
	return *o.NwdafCapability
}

// GetNwdafCapabilityOk returns a tuple with the NwdafCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetNwdafCapabilityOk() (*NwdafCapability, bool) {
	if o == nil || isNil(o.NwdafCapability) {
    return nil, false
	}
	return o.NwdafCapability, true
}

// HasNwdafCapability returns a boolean if a field has been set.
func (o *NwdafInfo) HasNwdafCapability() bool {
	if o != nil && !isNil(o.NwdafCapability) {
		return true
	}

	return false
}

// SetNwdafCapability gets a reference to the given NwdafCapability and assigns it to the NwdafCapability field.
func (o *NwdafInfo) SetNwdafCapability(v NwdafCapability) {
	o.NwdafCapability = &v
}

// GetAnalyticsDelay returns the AnalyticsDelay field value if set, zero value otherwise.
func (o *NwdafInfo) GetAnalyticsDelay() int32 {
	if o == nil || isNil(o.AnalyticsDelay) {
		var ret int32
		return ret
	}
	return *o.AnalyticsDelay
}

// GetAnalyticsDelayOk returns a tuple with the AnalyticsDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetAnalyticsDelayOk() (*int32, bool) {
	if o == nil || isNil(o.AnalyticsDelay) {
    return nil, false
	}
	return o.AnalyticsDelay, true
}

// HasAnalyticsDelay returns a boolean if a field has been set.
func (o *NwdafInfo) HasAnalyticsDelay() bool {
	if o != nil && !isNil(o.AnalyticsDelay) {
		return true
	}

	return false
}

// SetAnalyticsDelay gets a reference to the given int32 and assigns it to the AnalyticsDelay field.
func (o *NwdafInfo) SetAnalyticsDelay(v int32) {
	o.AnalyticsDelay = &v
}

// GetServingNfSetIdList returns the ServingNfSetIdList field value if set, zero value otherwise.
func (o *NwdafInfo) GetServingNfSetIdList() []string {
	if o == nil || isNil(o.ServingNfSetIdList) {
		var ret []string
		return ret
	}
	return o.ServingNfSetIdList
}

// GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetServingNfSetIdListOk() ([]string, bool) {
	if o == nil || isNil(o.ServingNfSetIdList) {
    return nil, false
	}
	return o.ServingNfSetIdList, true
}

// HasServingNfSetIdList returns a boolean if a field has been set.
func (o *NwdafInfo) HasServingNfSetIdList() bool {
	if o != nil && !isNil(o.ServingNfSetIdList) {
		return true
	}

	return false
}

// SetServingNfSetIdList gets a reference to the given []string and assigns it to the ServingNfSetIdList field.
func (o *NwdafInfo) SetServingNfSetIdList(v []string) {
	o.ServingNfSetIdList = v
}

// GetServingNfTypeList returns the ServingNfTypeList field value if set, zero value otherwise.
func (o *NwdafInfo) GetServingNfTypeList() []NFType {
	if o == nil || isNil(o.ServingNfTypeList) {
		var ret []NFType
		return ret
	}
	return o.ServingNfTypeList
}

// GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetServingNfTypeListOk() ([]NFType, bool) {
	if o == nil || isNil(o.ServingNfTypeList) {
    return nil, false
	}
	return o.ServingNfTypeList, true
}

// HasServingNfTypeList returns a boolean if a field has been set.
func (o *NwdafInfo) HasServingNfTypeList() bool {
	if o != nil && !isNil(o.ServingNfTypeList) {
		return true
	}

	return false
}

// SetServingNfTypeList gets a reference to the given []NFType and assigns it to the ServingNfTypeList field.
func (o *NwdafInfo) SetServingNfTypeList(v []NFType) {
	o.ServingNfTypeList = v
}

// GetMlAnalyticsList returns the MlAnalyticsList field value if set, zero value otherwise.
func (o *NwdafInfo) GetMlAnalyticsList() []MlAnalyticsInfo {
	if o == nil || isNil(o.MlAnalyticsList) {
		var ret []MlAnalyticsInfo
		return ret
	}
	return o.MlAnalyticsList
}

// GetMlAnalyticsListOk returns a tuple with the MlAnalyticsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetMlAnalyticsListOk() ([]MlAnalyticsInfo, bool) {
	if o == nil || isNil(o.MlAnalyticsList) {
    return nil, false
	}
	return o.MlAnalyticsList, true
}

// HasMlAnalyticsList returns a boolean if a field has been set.
func (o *NwdafInfo) HasMlAnalyticsList() bool {
	if o != nil && !isNil(o.MlAnalyticsList) {
		return true
	}

	return false
}

// SetMlAnalyticsList gets a reference to the given []MlAnalyticsInfo and assigns it to the MlAnalyticsList field.
func (o *NwdafInfo) SetMlAnalyticsList(v []MlAnalyticsInfo) {
	o.MlAnalyticsList = v
}

func (o NwdafInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventIds) {
		toSerialize["eventIds"] = o.EventIds
	}
	if !isNil(o.NwdafEvents) {
		toSerialize["nwdafEvents"] = o.NwdafEvents
	}
	if !isNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.NwdafCapability) {
		toSerialize["nwdafCapability"] = o.NwdafCapability
	}
	if !isNil(o.AnalyticsDelay) {
		toSerialize["analyticsDelay"] = o.AnalyticsDelay
	}
	if !isNil(o.ServingNfSetIdList) {
		toSerialize["servingNfSetIdList"] = o.ServingNfSetIdList
	}
	if !isNil(o.ServingNfTypeList) {
		toSerialize["servingNfTypeList"] = o.ServingNfTypeList
	}
	if !isNil(o.MlAnalyticsList) {
		toSerialize["mlAnalyticsList"] = o.MlAnalyticsList
	}
	return json.Marshal(toSerialize)
}

type NullableNwdafInfo struct {
	value *NwdafInfo
	isSet bool
}

func (v NullableNwdafInfo) Get() *NwdafInfo {
	return v.value
}

func (v *NullableNwdafInfo) Set(val *NwdafInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafInfo(val *NwdafInfo) *NullableNwdafInfo {
	return &NullableNwdafInfo{value: val, isSet: true}
}

func (v NullableNwdafInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


