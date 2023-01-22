/*
File Data Reporting MnS

OAS 3.0.1 definition of the File Data Reporting MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FileDataReportingMnS

import (
	"encoding/json"
)

// NotifyFilePreparationErrorAllOf struct for NotifyFilePreparationErrorAllOf
type NotifyFilePreparationErrorAllOf struct {
	FileInfoList []FileInfo `json:"fileInfoList,omitempty"`
	Reason *string `json:"reason,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
}

// NewNotifyFilePreparationErrorAllOf instantiates a new NotifyFilePreparationErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyFilePreparationErrorAllOf() *NotifyFilePreparationErrorAllOf {
	this := NotifyFilePreparationErrorAllOf{}
	return &this
}

// NewNotifyFilePreparationErrorAllOfWithDefaults instantiates a new NotifyFilePreparationErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyFilePreparationErrorAllOfWithDefaults() *NotifyFilePreparationErrorAllOf {
	this := NotifyFilePreparationErrorAllOf{}
	return &this
}

// GetFileInfoList returns the FileInfoList field value if set, zero value otherwise.
func (o *NotifyFilePreparationErrorAllOf) GetFileInfoList() []FileInfo {
	if o == nil || isNil(o.FileInfoList) {
		var ret []FileInfo
		return ret
	}
	return o.FileInfoList
}

// GetFileInfoListOk returns a tuple with the FileInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyFilePreparationErrorAllOf) GetFileInfoListOk() ([]FileInfo, bool) {
	if o == nil || isNil(o.FileInfoList) {
    return nil, false
	}
	return o.FileInfoList, true
}

// HasFileInfoList returns a boolean if a field has been set.
func (o *NotifyFilePreparationErrorAllOf) HasFileInfoList() bool {
	if o != nil && !isNil(o.FileInfoList) {
		return true
	}

	return false
}

// SetFileInfoList gets a reference to the given []FileInfo and assigns it to the FileInfoList field.
func (o *NotifyFilePreparationErrorAllOf) SetFileInfoList(v []FileInfo) {
	o.FileInfoList = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *NotifyFilePreparationErrorAllOf) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyFilePreparationErrorAllOf) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *NotifyFilePreparationErrorAllOf) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *NotifyFilePreparationErrorAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyFilePreparationErrorAllOf) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyFilePreparationErrorAllOf) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
    return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyFilePreparationErrorAllOf) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyFilePreparationErrorAllOf) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

func (o NotifyFilePreparationErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FileInfoList) {
		toSerialize["fileInfoList"] = o.FileInfoList
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyFilePreparationErrorAllOf struct {
	value *NotifyFilePreparationErrorAllOf
	isSet bool
}

func (v NullableNotifyFilePreparationErrorAllOf) Get() *NotifyFilePreparationErrorAllOf {
	return v.value
}

func (v *NullableNotifyFilePreparationErrorAllOf) Set(val *NotifyFilePreparationErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyFilePreparationErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyFilePreparationErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyFilePreparationErrorAllOf(val *NotifyFilePreparationErrorAllOf) *NullableNotifyFilePreparationErrorAllOf {
	return &NullableNotifyFilePreparationErrorAllOf{value: val, isSet: true}
}

func (v NullableNotifyFilePreparationErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyFilePreparationErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


