/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// NwdafCond Subscription to a set of NF Instances (NWDAFs), identified by Analytics ID(s), S-NSSAI(s) or NWDAF Serving Area information, i.e. list of TAIs for which the NWDAF can provide analytics. 
type NwdafCond struct {
	ConditionType string `json:"conditionType"`
	AnalyticsIds []string `json:"analyticsIds,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	ServingNfTypeList []NFType `json:"servingNfTypeList,omitempty"`
	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`
	MlAnalyticsList []MlAnalyticsInfo `json:"mlAnalyticsList,omitempty"`
}

// NewNwdafCond instantiates a new NwdafCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafCond(conditionType string) *NwdafCond {
	this := NwdafCond{}
	this.ConditionType = conditionType
	return &this
}

// NewNwdafCondWithDefaults instantiates a new NwdafCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafCondWithDefaults() *NwdafCond {
	this := NwdafCond{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *NwdafCond) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetConditionTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *NwdafCond) SetConditionType(v string) {
	o.ConditionType = v
}

// GetAnalyticsIds returns the AnalyticsIds field value if set, zero value otherwise.
func (o *NwdafCond) GetAnalyticsIds() []string {
	if o == nil || isNil(o.AnalyticsIds) {
		var ret []string
		return ret
	}
	return o.AnalyticsIds
}

// GetAnalyticsIdsOk returns a tuple with the AnalyticsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetAnalyticsIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AnalyticsIds) {
    return nil, false
	}
	return o.AnalyticsIds, true
}

// HasAnalyticsIds returns a boolean if a field has been set.
func (o *NwdafCond) HasAnalyticsIds() bool {
	if o != nil && !isNil(o.AnalyticsIds) {
		return true
	}

	return false
}

// SetAnalyticsIds gets a reference to the given []string and assigns it to the AnalyticsIds field.
func (o *NwdafCond) SetAnalyticsIds(v []string) {
	o.AnalyticsIds = v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *NwdafCond) GetSnssaiList() []Snssai {
	if o == nil || isNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.SnssaiList) {
    return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *NwdafCond) HasSnssaiList() bool {
	if o != nil && !isNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *NwdafCond) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NwdafCond) GetTaiList() []Tai {
	if o == nil || isNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaiList) {
    return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NwdafCond) HasTaiList() bool {
	if o != nil && !isNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NwdafCond) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NwdafCond) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NwdafCond) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NwdafCond) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetServingNfTypeList returns the ServingNfTypeList field value if set, zero value otherwise.
func (o *NwdafCond) GetServingNfTypeList() []NFType {
	if o == nil || isNil(o.ServingNfTypeList) {
		var ret []NFType
		return ret
	}
	return o.ServingNfTypeList
}

// GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetServingNfTypeListOk() ([]NFType, bool) {
	if o == nil || isNil(o.ServingNfTypeList) {
    return nil, false
	}
	return o.ServingNfTypeList, true
}

// HasServingNfTypeList returns a boolean if a field has been set.
func (o *NwdafCond) HasServingNfTypeList() bool {
	if o != nil && !isNil(o.ServingNfTypeList) {
		return true
	}

	return false
}

// SetServingNfTypeList gets a reference to the given []NFType and assigns it to the ServingNfTypeList field.
func (o *NwdafCond) SetServingNfTypeList(v []NFType) {
	o.ServingNfTypeList = v
}

// GetServingNfSetIdList returns the ServingNfSetIdList field value if set, zero value otherwise.
func (o *NwdafCond) GetServingNfSetIdList() []string {
	if o == nil || isNil(o.ServingNfSetIdList) {
		var ret []string
		return ret
	}
	return o.ServingNfSetIdList
}

// GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetServingNfSetIdListOk() ([]string, bool) {
	if o == nil || isNil(o.ServingNfSetIdList) {
    return nil, false
	}
	return o.ServingNfSetIdList, true
}

// HasServingNfSetIdList returns a boolean if a field has been set.
func (o *NwdafCond) HasServingNfSetIdList() bool {
	if o != nil && !isNil(o.ServingNfSetIdList) {
		return true
	}

	return false
}

// SetServingNfSetIdList gets a reference to the given []string and assigns it to the ServingNfSetIdList field.
func (o *NwdafCond) SetServingNfSetIdList(v []string) {
	o.ServingNfSetIdList = v
}

// GetMlAnalyticsList returns the MlAnalyticsList field value if set, zero value otherwise.
func (o *NwdafCond) GetMlAnalyticsList() []MlAnalyticsInfo {
	if o == nil || isNil(o.MlAnalyticsList) {
		var ret []MlAnalyticsInfo
		return ret
	}
	return o.MlAnalyticsList
}

// GetMlAnalyticsListOk returns a tuple with the MlAnalyticsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetMlAnalyticsListOk() ([]MlAnalyticsInfo, bool) {
	if o == nil || isNil(o.MlAnalyticsList) {
    return nil, false
	}
	return o.MlAnalyticsList, true
}

// HasMlAnalyticsList returns a boolean if a field has been set.
func (o *NwdafCond) HasMlAnalyticsList() bool {
	if o != nil && !isNil(o.MlAnalyticsList) {
		return true
	}

	return false
}

// SetMlAnalyticsList gets a reference to the given []MlAnalyticsInfo and assigns it to the MlAnalyticsList field.
func (o *NwdafCond) SetMlAnalyticsList(v []MlAnalyticsInfo) {
	o.MlAnalyticsList = v
}

func (o NwdafCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["conditionType"] = o.ConditionType
	}
	if !isNil(o.AnalyticsIds) {
		toSerialize["analyticsIds"] = o.AnalyticsIds
	}
	if !isNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !isNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.ServingNfTypeList) {
		toSerialize["servingNfTypeList"] = o.ServingNfTypeList
	}
	if !isNil(o.ServingNfSetIdList) {
		toSerialize["servingNfSetIdList"] = o.ServingNfSetIdList
	}
	if !isNil(o.MlAnalyticsList) {
		toSerialize["mlAnalyticsList"] = o.MlAnalyticsList
	}
	return json.Marshal(toSerialize)
}

type NullableNwdafCond struct {
	value *NwdafCond
	isSet bool
}

func (v NullableNwdafCond) Get() *NwdafCond {
	return v.value
}

func (v *NullableNwdafCond) Set(val *NwdafCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafCond(val *NwdafCond) *NullableNwdafCond {
	return &NullableNwdafCond{value: val, isSet: true}
}

func (v NullableNwdafCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


