/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// NotifyMoiCreation struct for NotifyMoiCreation
type NotifyMoiCreation struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
	SourceIndicator *SourceIndicator `json:"sourceIndicator,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	AttributeList map[string]interface{} `json:"attributeList,omitempty"`
}

// NewNotifyMoiCreation instantiates a new NotifyMoiCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyMoiCreation(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyMoiCreation {
	this := NotifyMoiCreation{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	return &this
}

// NewNotifyMoiCreationWithDefaults instantiates a new NotifyMoiCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyMoiCreationWithDefaults() *NotifyMoiCreation {
	this := NotifyMoiCreation{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyMoiCreation) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyMoiCreation) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyMoiCreation) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyMoiCreation) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyMoiCreation) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyMoiCreation) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyMoiCreation) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyMoiCreation) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyMoiCreation) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetSystemDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyMoiCreation) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *NotifyMoiCreation) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil || isNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil || isNil(o.CorrelatedNotifications) {
    return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *NotifyMoiCreation) HasCorrelatedNotifications() bool {
	if o != nil && !isNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification and assigns it to the CorrelatedNotifications field.
func (o *NotifyMoiCreation) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyMoiCreation) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
    return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyMoiCreation) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyMoiCreation) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetSourceIndicator returns the SourceIndicator field value if set, zero value otherwise.
func (o *NotifyMoiCreation) GetSourceIndicator() SourceIndicator {
	if o == nil || isNil(o.SourceIndicator) {
		var ret SourceIndicator
		return ret
	}
	return *o.SourceIndicator
}

// GetSourceIndicatorOk returns a tuple with the SourceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetSourceIndicatorOk() (*SourceIndicator, bool) {
	if o == nil || isNil(o.SourceIndicator) {
    return nil, false
	}
	return o.SourceIndicator, true
}

// HasSourceIndicator returns a boolean if a field has been set.
func (o *NotifyMoiCreation) HasSourceIndicator() bool {
	if o != nil && !isNil(o.SourceIndicator) {
		return true
	}

	return false
}

// SetSourceIndicator gets a reference to the given SourceIndicator and assigns it to the SourceIndicator field.
func (o *NotifyMoiCreation) SetSourceIndicator(v SourceIndicator) {
	o.SourceIndicator = &v
}

// GetAttributeList returns the AttributeList field value if set, zero value otherwise.
func (o *NotifyMoiCreation) GetAttributeList() map[string]interface{} {
	if o == nil || isNil(o.AttributeList) {
		var ret map[string]interface{}
		return ret
	}
	return o.AttributeList
}

// GetAttributeListOk returns a tuple with the AttributeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreation) GetAttributeListOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AttributeList) {
    return map[string]interface{}{}, false
	}
	return o.AttributeList, true
}

// HasAttributeList returns a boolean if a field has been set.
func (o *NotifyMoiCreation) HasAttributeList() bool {
	if o != nil && !isNil(o.AttributeList) {
		return true
	}

	return false
}

// SetAttributeList gets a reference to the given map[string]interface{} and assigns it to the AttributeList field.
func (o *NotifyMoiCreation) SetAttributeList(v map[string]interface{}) {
	o.AttributeList = v
}

func (o NotifyMoiCreation) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
	}
	if !isNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	if !isNil(o.SourceIndicator) {
		toSerialize["sourceIndicator"] = o.SourceIndicator
	}
	if !isNil(o.AttributeList) {
		toSerialize["attributeList"] = o.AttributeList
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyMoiCreation struct {
	value *NotifyMoiCreation
	isSet bool
}

func (v NullableNotifyMoiCreation) Get() *NotifyMoiCreation {
	return v.value
}

func (v *NullableNotifyMoiCreation) Set(val *NotifyMoiCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyMoiCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyMoiCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyMoiCreation(val *NotifyMoiCreation) *NullableNotifyMoiCreation {
	return &NullableNotifyMoiCreation{value: val, isSet: true}
}

func (v NullableNotifyMoiCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyMoiCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


