/*
File Data Reporting MnS

OAS 3.0.1 definition of the File Data Reporting MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FileDataReportingMnS

import (
	"encoding/json"
)

// NotifyFileReady struct for NotifyFileReady
type NotifyFileReady struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	FileInfoList []FileInfo `json:"fileInfoList,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
}

// NewNotifyFileReady instantiates a new NotifyFileReady object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyFileReady(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyFileReady {
	this := NotifyFileReady{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	return &this
}

// NewNotifyFileReadyWithDefaults instantiates a new NotifyFileReady object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyFileReadyWithDefaults() *NotifyFileReady {
	this := NotifyFileReady{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyFileReady) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyFileReady) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyFileReady) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyFileReady) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyFileReady) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyFileReady) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyFileReady) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyFileReady) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyFileReady) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyFileReady) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyFileReady) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyFileReady) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyFileReady) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyFileReady) GetSystemDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyFileReady) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetFileInfoList returns the FileInfoList field value if set, zero value otherwise.
func (o *NotifyFileReady) GetFileInfoList() []FileInfo {
	if o == nil || isNil(o.FileInfoList) {
		var ret []FileInfo
		return ret
	}
	return o.FileInfoList
}

// GetFileInfoListOk returns a tuple with the FileInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyFileReady) GetFileInfoListOk() ([]FileInfo, bool) {
	if o == nil || isNil(o.FileInfoList) {
    return nil, false
	}
	return o.FileInfoList, true
}

// HasFileInfoList returns a boolean if a field has been set.
func (o *NotifyFileReady) HasFileInfoList() bool {
	if o != nil && !isNil(o.FileInfoList) {
		return true
	}

	return false
}

// SetFileInfoList gets a reference to the given []FileInfo and assigns it to the FileInfoList field.
func (o *NotifyFileReady) SetFileInfoList(v []FileInfo) {
	o.FileInfoList = v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyFileReady) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyFileReady) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
    return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyFileReady) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyFileReady) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

func (o NotifyFileReady) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["notificationId"] = o.NotificationId
	}
	if true {
		toSerialize["notificationType"] = o.NotificationType
	}
	if true {
		toSerialize["eventTime"] = o.EventTime
	}
	if true {
		toSerialize["systemDN"] = o.SystemDN
	}
	if !isNil(o.FileInfoList) {
		toSerialize["fileInfoList"] = o.FileInfoList
	}
	if !isNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyFileReady struct {
	value *NotifyFileReady
	isSet bool
}

func (v NullableNotifyFileReady) Get() *NotifyFileReady {
	return v.value
}

func (v *NullableNotifyFileReady) Set(val *NotifyFileReady) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyFileReady) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyFileReady) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyFileReady(val *NotifyFileReady) *NullableNotifyFileReady {
	return &NullableNotifyFileReady{value: val, isSet: true}
}

func (v NullableNotifyFileReady) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyFileReady) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


