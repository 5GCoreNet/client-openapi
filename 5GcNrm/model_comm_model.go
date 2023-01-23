/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// CommModel struct for CommModel
type CommModel struct {
	GroupId *int32 `json:"groupId,omitempty"`
	CommModelType *CommModelType `json:"commModelType,omitempty"`
	TargetNFServiceList []string `json:"targetNFServiceList,omitempty"`
	CommModelConfiguration *string `json:"commModelConfiguration,omitempty"`
}

// NewCommModel instantiates a new CommModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommModel() *CommModel {
	this := CommModel{}
	return &this
}

// NewCommModelWithDefaults instantiates a new CommModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommModelWithDefaults() *CommModel {
	this := CommModel{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CommModel) GetGroupId() int32 {
	if o == nil || isNil(o.GroupId) {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommModel) GetGroupIdOk() (*int32, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CommModel) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *CommModel) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetCommModelType returns the CommModelType field value if set, zero value otherwise.
func (o *CommModel) GetCommModelType() CommModelType {
	if o == nil || isNil(o.CommModelType) {
		var ret CommModelType
		return ret
	}
	return *o.CommModelType
}

// GetCommModelTypeOk returns a tuple with the CommModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommModel) GetCommModelTypeOk() (*CommModelType, bool) {
	if o == nil || isNil(o.CommModelType) {
    return nil, false
	}
	return o.CommModelType, true
}

// HasCommModelType returns a boolean if a field has been set.
func (o *CommModel) HasCommModelType() bool {
	if o != nil && !isNil(o.CommModelType) {
		return true
	}

	return false
}

// SetCommModelType gets a reference to the given CommModelType and assigns it to the CommModelType field.
func (o *CommModel) SetCommModelType(v CommModelType) {
	o.CommModelType = &v
}

// GetTargetNFServiceList returns the TargetNFServiceList field value if set, zero value otherwise.
func (o *CommModel) GetTargetNFServiceList() []string {
	if o == nil || isNil(o.TargetNFServiceList) {
		var ret []string
		return ret
	}
	return o.TargetNFServiceList
}

// GetTargetNFServiceListOk returns a tuple with the TargetNFServiceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommModel) GetTargetNFServiceListOk() ([]string, bool) {
	if o == nil || isNil(o.TargetNFServiceList) {
    return nil, false
	}
	return o.TargetNFServiceList, true
}

// HasTargetNFServiceList returns a boolean if a field has been set.
func (o *CommModel) HasTargetNFServiceList() bool {
	if o != nil && !isNil(o.TargetNFServiceList) {
		return true
	}

	return false
}

// SetTargetNFServiceList gets a reference to the given []string and assigns it to the TargetNFServiceList field.
func (o *CommModel) SetTargetNFServiceList(v []string) {
	o.TargetNFServiceList = v
}

// GetCommModelConfiguration returns the CommModelConfiguration field value if set, zero value otherwise.
func (o *CommModel) GetCommModelConfiguration() string {
	if o == nil || isNil(o.CommModelConfiguration) {
		var ret string
		return ret
	}
	return *o.CommModelConfiguration
}

// GetCommModelConfigurationOk returns a tuple with the CommModelConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommModel) GetCommModelConfigurationOk() (*string, bool) {
	if o == nil || isNil(o.CommModelConfiguration) {
    return nil, false
	}
	return o.CommModelConfiguration, true
}

// HasCommModelConfiguration returns a boolean if a field has been set.
func (o *CommModel) HasCommModelConfiguration() bool {
	if o != nil && !isNil(o.CommModelConfiguration) {
		return true
	}

	return false
}

// SetCommModelConfiguration gets a reference to the given string and assigns it to the CommModelConfiguration field.
func (o *CommModel) SetCommModelConfiguration(v string) {
	o.CommModelConfiguration = &v
}

func (o CommModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.CommModelType) {
		toSerialize["commModelType"] = o.CommModelType
	}
	if !isNil(o.TargetNFServiceList) {
		toSerialize["targetNFServiceList"] = o.TargetNFServiceList
	}
	if !isNil(o.CommModelConfiguration) {
		toSerialize["commModelConfiguration"] = o.CommModelConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableCommModel struct {
	value *CommModel
	isSet bool
}

func (v NullableCommModel) Get() *CommModel {
	return v.value
}

func (v *NullableCommModel) Set(val *CommModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCommModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCommModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommModel(val *CommModel) *NullableCommModel {
	return &NullableCommModel{value: val, isSet: true}
}

func (v NullableCommModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


