/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FaultMnS

import (
	"encoding/json"
)

// NotifyChangedAlarm struct for NotifyChangedAlarm
type NotifyChangedAlarm struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	AlarmId string `json:"alarmId"`
	AlarmType AlarmType `json:"alarmType"`
	ProbableCause ProbableCause `json:"probableCause"`
	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity"`
}

// NewNotifyChangedAlarm instantiates a new NotifyChangedAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyChangedAlarm(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity) *NotifyChangedAlarm {
	this := NotifyChangedAlarm{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.AlarmId = alarmId
	this.AlarmType = alarmType
	this.ProbableCause = probableCause
	this.PerceivedSeverity = perceivedSeverity
	return &this
}

// NewNotifyChangedAlarmWithDefaults instantiates a new NotifyChangedAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyChangedAlarmWithDefaults() *NotifyChangedAlarm {
	this := NotifyChangedAlarm{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyChangedAlarm) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyChangedAlarm) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyChangedAlarm) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyChangedAlarm) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyChangedAlarm) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyChangedAlarm) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyChangedAlarm) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyChangedAlarm) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyChangedAlarm) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetSystemDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyChangedAlarm) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyChangedAlarm) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetAlarmIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyChangedAlarm) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetAlarmType returns the AlarmType field value
func (o *NotifyChangedAlarm) GetAlarmType() AlarmType {
	if o == nil {
		var ret AlarmType
		return ret
	}

	return o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlarmType, true
}

// SetAlarmType sets field value
func (o *NotifyChangedAlarm) SetAlarmType(v AlarmType) {
	o.AlarmType = v
}

// GetProbableCause returns the ProbableCause field value
func (o *NotifyChangedAlarm) GetProbableCause() ProbableCause {
	if o == nil {
		var ret ProbableCause
		return ret
	}

	return o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProbableCause, true
}

// SetProbableCause sets field value
func (o *NotifyChangedAlarm) SetProbableCause(v ProbableCause) {
	o.ProbableCause = v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value
func (o *NotifyChangedAlarm) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil {
		var ret PerceivedSeverity
		return ret
	}

	return o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PerceivedSeverity, true
}

// SetPerceivedSeverity sets field value
func (o *NotifyChangedAlarm) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = v
}

func (o NotifyChangedAlarm) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["alarmId"] = o.AlarmId
	}
	if true {
		toSerialize["alarmType"] = o.AlarmType
	}
	if true {
		toSerialize["probableCause"] = o.ProbableCause
	}
	if true {
		toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyChangedAlarm struct {
	value *NotifyChangedAlarm
	isSet bool
}

func (v NullableNotifyChangedAlarm) Get() *NotifyChangedAlarm {
	return v.value
}

func (v *NullableNotifyChangedAlarm) Set(val *NotifyChangedAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyChangedAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyChangedAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyChangedAlarm(val *NotifyChangedAlarm) *NullableNotifyChangedAlarm {
	return &NullableNotifyChangedAlarm{value: val, isSet: true}
}

func (v NullableNotifyChangedAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyChangedAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


