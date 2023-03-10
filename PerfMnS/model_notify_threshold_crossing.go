/*
TS 28.532 Performance Threshold Monitoring MnS

OAS 3.0.1 definition of the Performance Threshold Monitoring MnS © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMnS

import (
	"encoding/json"
)

// NotifyThresholdCrossing struct for NotifyThresholdCrossing
type NotifyThresholdCrossing struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	ObservedPerfMetricName *string `json:"observedPerfMetricName,omitempty"`
	ObservedPerfMetricValue *PerfMetricValue `json:"observedPerfMetricValue,omitempty"`
	ObservedPerfMetricDirection *PerfMetricDirection `json:"observedPerfMetricDirection,omitempty"`
	ThresholdValue *PerfMetricValue `json:"thresholdValue,omitempty"`
	Hysteresis *PerfMetricValue `json:"hysteresis,omitempty"`
	MonitorGranularityPeriod *int32 `json:"monitorGranularityPeriod,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
}

// NewNotifyThresholdCrossing instantiates a new NotifyThresholdCrossing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyThresholdCrossing(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyThresholdCrossing {
	this := NotifyThresholdCrossing{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	return &this
}

// NewNotifyThresholdCrossingWithDefaults instantiates a new NotifyThresholdCrossing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyThresholdCrossingWithDefaults() *NotifyThresholdCrossing {
	this := NotifyThresholdCrossing{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyThresholdCrossing) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyThresholdCrossing) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyThresholdCrossing) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyThresholdCrossing) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyThresholdCrossing) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyThresholdCrossing) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyThresholdCrossing) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyThresholdCrossing) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyThresholdCrossing) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetSystemDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyThresholdCrossing) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetObservedPerfMetricName returns the ObservedPerfMetricName field value if set, zero value otherwise.
func (o *NotifyThresholdCrossing) GetObservedPerfMetricName() string {
	if o == nil || isNil(o.ObservedPerfMetricName) {
		var ret string
		return ret
	}
	return *o.ObservedPerfMetricName
}

// GetObservedPerfMetricNameOk returns a tuple with the ObservedPerfMetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetObservedPerfMetricNameOk() (*string, bool) {
	if o == nil || isNil(o.ObservedPerfMetricName) {
    return nil, false
	}
	return o.ObservedPerfMetricName, true
}

// HasObservedPerfMetricName returns a boolean if a field has been set.
func (o *NotifyThresholdCrossing) HasObservedPerfMetricName() bool {
	if o != nil && !isNil(o.ObservedPerfMetricName) {
		return true
	}

	return false
}

// SetObservedPerfMetricName gets a reference to the given string and assigns it to the ObservedPerfMetricName field.
func (o *NotifyThresholdCrossing) SetObservedPerfMetricName(v string) {
	o.ObservedPerfMetricName = &v
}

// GetObservedPerfMetricValue returns the ObservedPerfMetricValue field value if set, zero value otherwise.
func (o *NotifyThresholdCrossing) GetObservedPerfMetricValue() PerfMetricValue {
	if o == nil || isNil(o.ObservedPerfMetricValue) {
		var ret PerfMetricValue
		return ret
	}
	return *o.ObservedPerfMetricValue
}

// GetObservedPerfMetricValueOk returns a tuple with the ObservedPerfMetricValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetObservedPerfMetricValueOk() (*PerfMetricValue, bool) {
	if o == nil || isNil(o.ObservedPerfMetricValue) {
    return nil, false
	}
	return o.ObservedPerfMetricValue, true
}

// HasObservedPerfMetricValue returns a boolean if a field has been set.
func (o *NotifyThresholdCrossing) HasObservedPerfMetricValue() bool {
	if o != nil && !isNil(o.ObservedPerfMetricValue) {
		return true
	}

	return false
}

// SetObservedPerfMetricValue gets a reference to the given PerfMetricValue and assigns it to the ObservedPerfMetricValue field.
func (o *NotifyThresholdCrossing) SetObservedPerfMetricValue(v PerfMetricValue) {
	o.ObservedPerfMetricValue = &v
}

// GetObservedPerfMetricDirection returns the ObservedPerfMetricDirection field value if set, zero value otherwise.
func (o *NotifyThresholdCrossing) GetObservedPerfMetricDirection() PerfMetricDirection {
	if o == nil || isNil(o.ObservedPerfMetricDirection) {
		var ret PerfMetricDirection
		return ret
	}
	return *o.ObservedPerfMetricDirection
}

// GetObservedPerfMetricDirectionOk returns a tuple with the ObservedPerfMetricDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetObservedPerfMetricDirectionOk() (*PerfMetricDirection, bool) {
	if o == nil || isNil(o.ObservedPerfMetricDirection) {
    return nil, false
	}
	return o.ObservedPerfMetricDirection, true
}

// HasObservedPerfMetricDirection returns a boolean if a field has been set.
func (o *NotifyThresholdCrossing) HasObservedPerfMetricDirection() bool {
	if o != nil && !isNil(o.ObservedPerfMetricDirection) {
		return true
	}

	return false
}

// SetObservedPerfMetricDirection gets a reference to the given PerfMetricDirection and assigns it to the ObservedPerfMetricDirection field.
func (o *NotifyThresholdCrossing) SetObservedPerfMetricDirection(v PerfMetricDirection) {
	o.ObservedPerfMetricDirection = &v
}

// GetThresholdValue returns the ThresholdValue field value if set, zero value otherwise.
func (o *NotifyThresholdCrossing) GetThresholdValue() PerfMetricValue {
	if o == nil || isNil(o.ThresholdValue) {
		var ret PerfMetricValue
		return ret
	}
	return *o.ThresholdValue
}

// GetThresholdValueOk returns a tuple with the ThresholdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetThresholdValueOk() (*PerfMetricValue, bool) {
	if o == nil || isNil(o.ThresholdValue) {
    return nil, false
	}
	return o.ThresholdValue, true
}

// HasThresholdValue returns a boolean if a field has been set.
func (o *NotifyThresholdCrossing) HasThresholdValue() bool {
	if o != nil && !isNil(o.ThresholdValue) {
		return true
	}

	return false
}

// SetThresholdValue gets a reference to the given PerfMetricValue and assigns it to the ThresholdValue field.
func (o *NotifyThresholdCrossing) SetThresholdValue(v PerfMetricValue) {
	o.ThresholdValue = &v
}

// GetHysteresis returns the Hysteresis field value if set, zero value otherwise.
func (o *NotifyThresholdCrossing) GetHysteresis() PerfMetricValue {
	if o == nil || isNil(o.Hysteresis) {
		var ret PerfMetricValue
		return ret
	}
	return *o.Hysteresis
}

// GetHysteresisOk returns a tuple with the Hysteresis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetHysteresisOk() (*PerfMetricValue, bool) {
	if o == nil || isNil(o.Hysteresis) {
    return nil, false
	}
	return o.Hysteresis, true
}

// HasHysteresis returns a boolean if a field has been set.
func (o *NotifyThresholdCrossing) HasHysteresis() bool {
	if o != nil && !isNil(o.Hysteresis) {
		return true
	}

	return false
}

// SetHysteresis gets a reference to the given PerfMetricValue and assigns it to the Hysteresis field.
func (o *NotifyThresholdCrossing) SetHysteresis(v PerfMetricValue) {
	o.Hysteresis = &v
}

// GetMonitorGranularityPeriod returns the MonitorGranularityPeriod field value if set, zero value otherwise.
func (o *NotifyThresholdCrossing) GetMonitorGranularityPeriod() int32 {
	if o == nil || isNil(o.MonitorGranularityPeriod) {
		var ret int32
		return ret
	}
	return *o.MonitorGranularityPeriod
}

// GetMonitorGranularityPeriodOk returns a tuple with the MonitorGranularityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetMonitorGranularityPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.MonitorGranularityPeriod) {
    return nil, false
	}
	return o.MonitorGranularityPeriod, true
}

// HasMonitorGranularityPeriod returns a boolean if a field has been set.
func (o *NotifyThresholdCrossing) HasMonitorGranularityPeriod() bool {
	if o != nil && !isNil(o.MonitorGranularityPeriod) {
		return true
	}

	return false
}

// SetMonitorGranularityPeriod gets a reference to the given int32 and assigns it to the MonitorGranularityPeriod field.
func (o *NotifyThresholdCrossing) SetMonitorGranularityPeriod(v int32) {
	o.MonitorGranularityPeriod = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyThresholdCrossing) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossing) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
    return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyThresholdCrossing) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyThresholdCrossing) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

func (o NotifyThresholdCrossing) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ObservedPerfMetricName) {
		toSerialize["observedPerfMetricName"] = o.ObservedPerfMetricName
	}
	if !isNil(o.ObservedPerfMetricValue) {
		toSerialize["observedPerfMetricValue"] = o.ObservedPerfMetricValue
	}
	if !isNil(o.ObservedPerfMetricDirection) {
		toSerialize["observedPerfMetricDirection"] = o.ObservedPerfMetricDirection
	}
	if !isNil(o.ThresholdValue) {
		toSerialize["thresholdValue"] = o.ThresholdValue
	}
	if !isNil(o.Hysteresis) {
		toSerialize["hysteresis"] = o.Hysteresis
	}
	if !isNil(o.MonitorGranularityPeriod) {
		toSerialize["monitorGranularityPeriod"] = o.MonitorGranularityPeriod
	}
	if !isNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyThresholdCrossing struct {
	value *NotifyThresholdCrossing
	isSet bool
}

func (v NullableNotifyThresholdCrossing) Get() *NotifyThresholdCrossing {
	return v.value
}

func (v *NullableNotifyThresholdCrossing) Set(val *NotifyThresholdCrossing) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyThresholdCrossing) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyThresholdCrossing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyThresholdCrossing(val *NotifyThresholdCrossing) *NullableNotifyThresholdCrossing {
	return &NullableNotifyThresholdCrossing{value: val, isSet: true}
}

func (v NullableNotifyThresholdCrossing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyThresholdCrossing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


