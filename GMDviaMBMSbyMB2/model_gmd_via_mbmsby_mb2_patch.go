/*
GMDviaMBMSbyMB2

API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GMDviaMBMSbyMB2

import (
	"encoding/json"
	"time"
)

// GMDViaMBMSByMb2Patch Represents a modification request of a group message delivery via MBMS by MB2.
type GMDViaMBMSByMb2Patch struct {
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId *string `json:"externalGroupId,omitempty"`
	MbmsLocArea *MbmsLocArea `json:"mbmsLocArea,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MessageDeliveryStartTime *time.Time `json:"messageDeliveryStartTime,omitempty"`
	// String with format \"byte\" as defined in OpenAPI Specification, i.e, base64-encoded characters.
	GroupMessagePayload *string `json:"groupMessagePayload,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
}

// NewGMDViaMBMSByMb2Patch instantiates a new GMDViaMBMSByMb2Patch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGMDViaMBMSByMb2Patch() *GMDViaMBMSByMb2Patch {
	this := GMDViaMBMSByMb2Patch{}
	return &this
}

// NewGMDViaMBMSByMb2PatchWithDefaults instantiates a new GMDViaMBMSByMb2Patch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGMDViaMBMSByMb2PatchWithDefaults() *GMDViaMBMSByMb2Patch {
	this := GMDViaMBMSByMb2Patch{}
	return &this
}

// GetExternalGroupId returns the ExternalGroupId field value if set, zero value otherwise.
func (o *GMDViaMBMSByMb2Patch) GetExternalGroupId() string {
	if o == nil || isNil(o.ExternalGroupId) {
		var ret string
		return ret
	}
	return *o.ExternalGroupId
}

// GetExternalGroupIdOk returns a tuple with the ExternalGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByMb2Patch) GetExternalGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalGroupId) {
    return nil, false
	}
	return o.ExternalGroupId, true
}

// HasExternalGroupId returns a boolean if a field has been set.
func (o *GMDViaMBMSByMb2Patch) HasExternalGroupId() bool {
	if o != nil && !isNil(o.ExternalGroupId) {
		return true
	}

	return false
}

// SetExternalGroupId gets a reference to the given string and assigns it to the ExternalGroupId field.
func (o *GMDViaMBMSByMb2Patch) SetExternalGroupId(v string) {
	o.ExternalGroupId = &v
}

// GetMbmsLocArea returns the MbmsLocArea field value if set, zero value otherwise.
func (o *GMDViaMBMSByMb2Patch) GetMbmsLocArea() MbmsLocArea {
	if o == nil || isNil(o.MbmsLocArea) {
		var ret MbmsLocArea
		return ret
	}
	return *o.MbmsLocArea
}

// GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByMb2Patch) GetMbmsLocAreaOk() (*MbmsLocArea, bool) {
	if o == nil || isNil(o.MbmsLocArea) {
    return nil, false
	}
	return o.MbmsLocArea, true
}

// HasMbmsLocArea returns a boolean if a field has been set.
func (o *GMDViaMBMSByMb2Patch) HasMbmsLocArea() bool {
	if o != nil && !isNil(o.MbmsLocArea) {
		return true
	}

	return false
}

// SetMbmsLocArea gets a reference to the given MbmsLocArea and assigns it to the MbmsLocArea field.
func (o *GMDViaMBMSByMb2Patch) SetMbmsLocArea(v MbmsLocArea) {
	o.MbmsLocArea = &v
}

// GetMessageDeliveryStartTime returns the MessageDeliveryStartTime field value if set, zero value otherwise.
func (o *GMDViaMBMSByMb2Patch) GetMessageDeliveryStartTime() time.Time {
	if o == nil || isNil(o.MessageDeliveryStartTime) {
		var ret time.Time
		return ret
	}
	return *o.MessageDeliveryStartTime
}

// GetMessageDeliveryStartTimeOk returns a tuple with the MessageDeliveryStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByMb2Patch) GetMessageDeliveryStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.MessageDeliveryStartTime) {
    return nil, false
	}
	return o.MessageDeliveryStartTime, true
}

// HasMessageDeliveryStartTime returns a boolean if a field has been set.
func (o *GMDViaMBMSByMb2Patch) HasMessageDeliveryStartTime() bool {
	if o != nil && !isNil(o.MessageDeliveryStartTime) {
		return true
	}

	return false
}

// SetMessageDeliveryStartTime gets a reference to the given time.Time and assigns it to the MessageDeliveryStartTime field.
func (o *GMDViaMBMSByMb2Patch) SetMessageDeliveryStartTime(v time.Time) {
	o.MessageDeliveryStartTime = &v
}

// GetGroupMessagePayload returns the GroupMessagePayload field value if set, zero value otherwise.
func (o *GMDViaMBMSByMb2Patch) GetGroupMessagePayload() string {
	if o == nil || isNil(o.GroupMessagePayload) {
		var ret string
		return ret
	}
	return *o.GroupMessagePayload
}

// GetGroupMessagePayloadOk returns a tuple with the GroupMessagePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByMb2Patch) GetGroupMessagePayloadOk() (*string, bool) {
	if o == nil || isNil(o.GroupMessagePayload) {
    return nil, false
	}
	return o.GroupMessagePayload, true
}

// HasGroupMessagePayload returns a boolean if a field has been set.
func (o *GMDViaMBMSByMb2Patch) HasGroupMessagePayload() bool {
	if o != nil && !isNil(o.GroupMessagePayload) {
		return true
	}

	return false
}

// SetGroupMessagePayload gets a reference to the given string and assigns it to the GroupMessagePayload field.
func (o *GMDViaMBMSByMb2Patch) SetGroupMessagePayload(v string) {
	o.GroupMessagePayload = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *GMDViaMBMSByMb2Patch) GetNotificationDestination() string {
	if o == nil || isNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByMb2Patch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || isNil(o.NotificationDestination) {
    return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *GMDViaMBMSByMb2Patch) HasNotificationDestination() bool {
	if o != nil && !isNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *GMDViaMBMSByMb2Patch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

func (o GMDViaMBMSByMb2Patch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalGroupId) {
		toSerialize["externalGroupId"] = o.ExternalGroupId
	}
	if !isNil(o.MbmsLocArea) {
		toSerialize["mbmsLocArea"] = o.MbmsLocArea
	}
	if !isNil(o.MessageDeliveryStartTime) {
		toSerialize["messageDeliveryStartTime"] = o.MessageDeliveryStartTime
	}
	if !isNil(o.GroupMessagePayload) {
		toSerialize["groupMessagePayload"] = o.GroupMessagePayload
	}
	if !isNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	return json.Marshal(toSerialize)
}

type NullableGMDViaMBMSByMb2Patch struct {
	value *GMDViaMBMSByMb2Patch
	isSet bool
}

func (v NullableGMDViaMBMSByMb2Patch) Get() *GMDViaMBMSByMb2Patch {
	return v.value
}

func (v *NullableGMDViaMBMSByMb2Patch) Set(val *GMDViaMBMSByMb2Patch) {
	v.value = val
	v.isSet = true
}

func (v NullableGMDViaMBMSByMb2Patch) IsSet() bool {
	return v.isSet
}

func (v *NullableGMDViaMBMSByMb2Patch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGMDViaMBMSByMb2Patch(val *GMDViaMBMSByMb2Patch) *NullableGMDViaMBMSByMb2Patch {
	return &NullableGMDViaMBMSByMb2Patch{value: val, isSet: true}
}

func (v NullableGMDViaMBMSByMb2Patch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGMDViaMBMSByMb2Patch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


