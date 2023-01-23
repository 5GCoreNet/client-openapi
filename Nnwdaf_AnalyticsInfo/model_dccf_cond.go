/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// DccfCond Subscription to a set of NF Instances (DCCFs), identified by NF types, NF Set Id(s) or DCCF Serving Area information, i.e. list of TAIs served by the DCCF 
type DccfCond struct {
	ConditionType string `json:"conditionType"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	ServingNfTypeList []NFType `json:"servingNfTypeList,omitempty"`
	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`
}

// NewDccfCond instantiates a new DccfCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDccfCond(conditionType string) *DccfCond {
	this := DccfCond{}
	this.ConditionType = conditionType
	return &this
}

// NewDccfCondWithDefaults instantiates a new DccfCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDccfCondWithDefaults() *DccfCond {
	this := DccfCond{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *DccfCond) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *DccfCond) GetConditionTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *DccfCond) SetConditionType(v string) {
	o.ConditionType = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *DccfCond) GetTaiList() []Tai {
	if o == nil || isNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfCond) GetTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaiList) {
    return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *DccfCond) HasTaiList() bool {
	if o != nil && !isNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *DccfCond) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *DccfCond) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfCond) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *DccfCond) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *DccfCond) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetServingNfTypeList returns the ServingNfTypeList field value if set, zero value otherwise.
func (o *DccfCond) GetServingNfTypeList() []NFType {
	if o == nil || isNil(o.ServingNfTypeList) {
		var ret []NFType
		return ret
	}
	return o.ServingNfTypeList
}

// GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfCond) GetServingNfTypeListOk() ([]NFType, bool) {
	if o == nil || isNil(o.ServingNfTypeList) {
    return nil, false
	}
	return o.ServingNfTypeList, true
}

// HasServingNfTypeList returns a boolean if a field has been set.
func (o *DccfCond) HasServingNfTypeList() bool {
	if o != nil && !isNil(o.ServingNfTypeList) {
		return true
	}

	return false
}

// SetServingNfTypeList gets a reference to the given []NFType and assigns it to the ServingNfTypeList field.
func (o *DccfCond) SetServingNfTypeList(v []NFType) {
	o.ServingNfTypeList = v
}

// GetServingNfSetIdList returns the ServingNfSetIdList field value if set, zero value otherwise.
func (o *DccfCond) GetServingNfSetIdList() []string {
	if o == nil || isNil(o.ServingNfSetIdList) {
		var ret []string
		return ret
	}
	return o.ServingNfSetIdList
}

// GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfCond) GetServingNfSetIdListOk() ([]string, bool) {
	if o == nil || isNil(o.ServingNfSetIdList) {
    return nil, false
	}
	return o.ServingNfSetIdList, true
}

// HasServingNfSetIdList returns a boolean if a field has been set.
func (o *DccfCond) HasServingNfSetIdList() bool {
	if o != nil && !isNil(o.ServingNfSetIdList) {
		return true
	}

	return false
}

// SetServingNfSetIdList gets a reference to the given []string and assigns it to the ServingNfSetIdList field.
func (o *DccfCond) SetServingNfSetIdList(v []string) {
	o.ServingNfSetIdList = v
}

func (o DccfCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["conditionType"] = o.ConditionType
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
	return json.Marshal(toSerialize)
}

type NullableDccfCond struct {
	value *DccfCond
	isSet bool
}

func (v NullableDccfCond) Get() *DccfCond {
	return v.value
}

func (v *NullableDccfCond) Set(val *DccfCond) {
	v.value = val
	v.isSet = true
}

func (v NullableDccfCond) IsSet() bool {
	return v.isSet
}

func (v *NullableDccfCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDccfCond(val *DccfCond) *NullableDccfCond {
	return &NullableDccfCond{value: val, isSet: true}
}

func (v NullableDccfCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDccfCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


