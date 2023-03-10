/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"time"
)

// SubscriptionDataSubscriptions A subscription to notifications.
type SubscriptionDataSubscriptions struct {
	// String represents the SUPI or GPSI
	UeId *string `json:"ueId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
	// String providing an URI formatted according to RFC 3986.
	OriginalCallbackReference *string `json:"originalCallbackReference,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	SdmSubscription *SdmSubscription `json:"sdmSubscription,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	UniqueSubscription *bool `json:"uniqueSubscription,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewSubscriptionDataSubscriptions instantiates a new SubscriptionDataSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDataSubscriptions(callbackReference string, monitoredResourceUris []string) *SubscriptionDataSubscriptions {
	this := SubscriptionDataSubscriptions{}
	this.CallbackReference = callbackReference
	this.MonitoredResourceUris = monitoredResourceUris
	return &this
}

// NewSubscriptionDataSubscriptionsWithDefaults instantiates a new SubscriptionDataSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDataSubscriptionsWithDefaults() *SubscriptionDataSubscriptions {
	this := SubscriptionDataSubscriptions{}
	return &this
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetUeId() string {
	if o == nil || isNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetUeIdOk() (*string, bool) {
	if o == nil || isNil(o.UeId) {
    return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasUeId() bool {
	if o != nil && !isNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *SubscriptionDataSubscriptions) SetUeId(v string) {
	o.UeId = &v
}

// GetCallbackReference returns the CallbackReference field value
func (o *SubscriptionDataSubscriptions) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *SubscriptionDataSubscriptions) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetOriginalCallbackReference returns the OriginalCallbackReference field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetOriginalCallbackReference() string {
	if o == nil || isNil(o.OriginalCallbackReference) {
		var ret string
		return ret
	}
	return *o.OriginalCallbackReference
}

// GetOriginalCallbackReferenceOk returns a tuple with the OriginalCallbackReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetOriginalCallbackReferenceOk() (*string, bool) {
	if o == nil || isNil(o.OriginalCallbackReference) {
    return nil, false
	}
	return o.OriginalCallbackReference, true
}

// HasOriginalCallbackReference returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasOriginalCallbackReference() bool {
	if o != nil && !isNil(o.OriginalCallbackReference) {
		return true
	}

	return false
}

// SetOriginalCallbackReference gets a reference to the given string and assigns it to the OriginalCallbackReference field.
func (o *SubscriptionDataSubscriptions) SetOriginalCallbackReference(v string) {
	o.OriginalCallbackReference = &v
}

// GetMonitoredResourceUris returns the MonitoredResourceUris field value
func (o *SubscriptionDataSubscriptions) GetMonitoredResourceUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MonitoredResourceUris
}

// GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetMonitoredResourceUrisOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MonitoredResourceUris, true
}

// SetMonitoredResourceUris sets field value
func (o *SubscriptionDataSubscriptions) SetMonitoredResourceUris(v []string) {
	o.MonitoredResourceUris = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *SubscriptionDataSubscriptions) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSdmSubscription returns the SdmSubscription field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetSdmSubscription() SdmSubscription {
	if o == nil || isNil(o.SdmSubscription) {
		var ret SdmSubscription
		return ret
	}
	return *o.SdmSubscription
}

// GetSdmSubscriptionOk returns a tuple with the SdmSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetSdmSubscriptionOk() (*SdmSubscription, bool) {
	if o == nil || isNil(o.SdmSubscription) {
    return nil, false
	}
	return o.SdmSubscription, true
}

// HasSdmSubscription returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasSdmSubscription() bool {
	if o != nil && !isNil(o.SdmSubscription) {
		return true
	}

	return false
}

// SetSdmSubscription gets a reference to the given SdmSubscription and assigns it to the SdmSubscription field.
func (o *SubscriptionDataSubscriptions) SetSdmSubscription(v SdmSubscription) {
	o.SdmSubscription = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
    return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SubscriptionDataSubscriptions) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUniqueSubscription returns the UniqueSubscription field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetUniqueSubscription() bool {
	if o == nil || isNil(o.UniqueSubscription) {
		var ret bool
		return ret
	}
	return *o.UniqueSubscription
}

// GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetUniqueSubscriptionOk() (*bool, bool) {
	if o == nil || isNil(o.UniqueSubscription) {
    return nil, false
	}
	return o.UniqueSubscription, true
}

// HasUniqueSubscription returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasUniqueSubscription() bool {
	if o != nil && !isNil(o.UniqueSubscription) {
		return true
	}

	return false
}

// SetUniqueSubscription gets a reference to the given bool and assigns it to the UniqueSubscription field.
func (o *SubscriptionDataSubscriptions) SetUniqueSubscription(v bool) {
	o.UniqueSubscription = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SubscriptionDataSubscriptions) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o SubscriptionDataSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	if true {
		toSerialize["callbackReference"] = o.CallbackReference
	}
	if !isNil(o.OriginalCallbackReference) {
		toSerialize["originalCallbackReference"] = o.OriginalCallbackReference
	}
	if true {
		toSerialize["monitoredResourceUris"] = o.MonitoredResourceUris
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !isNil(o.SdmSubscription) {
		toSerialize["sdmSubscription"] = o.SdmSubscription
	}
	if !isNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !isNil(o.UniqueSubscription) {
		toSerialize["uniqueSubscription"] = o.UniqueSubscription
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionDataSubscriptions struct {
	value *SubscriptionDataSubscriptions
	isSet bool
}

func (v NullableSubscriptionDataSubscriptions) Get() *SubscriptionDataSubscriptions {
	return v.value
}

func (v *NullableSubscriptionDataSubscriptions) Set(val *SubscriptionDataSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDataSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDataSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDataSubscriptions(val *SubscriptionDataSubscriptions) *NullableSubscriptionDataSubscriptions {
	return &NullableSubscriptionDataSubscriptions{value: val, isSet: true}
}

func (v NullableSubscriptionDataSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDataSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


