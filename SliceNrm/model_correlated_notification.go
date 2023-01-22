/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
)

// CorrelatedNotification struct for CorrelatedNotification
type CorrelatedNotification struct {
	SourceObjectInstance string `json:"sourceObjectInstance"`
	NotificationIds []int32 `json:"notificationIds"`
}

// NewCorrelatedNotification instantiates a new CorrelatedNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrelatedNotification(sourceObjectInstance string, notificationIds []int32) *CorrelatedNotification {
	this := CorrelatedNotification{}
	this.SourceObjectInstance = sourceObjectInstance
	this.NotificationIds = notificationIds
	return &this
}

// NewCorrelatedNotificationWithDefaults instantiates a new CorrelatedNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrelatedNotificationWithDefaults() *CorrelatedNotification {
	this := CorrelatedNotification{}
	return &this
}

// GetSourceObjectInstance returns the SourceObjectInstance field value
func (o *CorrelatedNotification) GetSourceObjectInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceObjectInstance
}

// GetSourceObjectInstanceOk returns a tuple with the SourceObjectInstance field value
// and a boolean to check if the value has been set.
func (o *CorrelatedNotification) GetSourceObjectInstanceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceObjectInstance, true
}

// SetSourceObjectInstance sets field value
func (o *CorrelatedNotification) SetSourceObjectInstance(v string) {
	o.SourceObjectInstance = v
}

// GetNotificationIds returns the NotificationIds field value
func (o *CorrelatedNotification) GetNotificationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.NotificationIds
}

// GetNotificationIdsOk returns a tuple with the NotificationIds field value
// and a boolean to check if the value has been set.
func (o *CorrelatedNotification) GetNotificationIdsOk() ([]int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NotificationIds, true
}

// SetNotificationIds sets field value
func (o *CorrelatedNotification) SetNotificationIds(v []int32) {
	o.NotificationIds = v
}

func (o CorrelatedNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceObjectInstance"] = o.SourceObjectInstance
	}
	if true {
		toSerialize["notificationIds"] = o.NotificationIds
	}
	return json.Marshal(toSerialize)
}

type NullableCorrelatedNotification struct {
	value *CorrelatedNotification
	isSet bool
}

func (v NullableCorrelatedNotification) Get() *CorrelatedNotification {
	return v.value
}

func (v *NullableCorrelatedNotification) Set(val *CorrelatedNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrelatedNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrelatedNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrelatedNotification(val *CorrelatedNotification) *NullableCorrelatedNotification {
	return &NullableCorrelatedNotification{value: val, isSet: true}
}

func (v NullableCorrelatedNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrelatedNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


