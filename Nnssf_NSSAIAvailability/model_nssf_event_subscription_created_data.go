/*
NSSF NSSAI Availability

NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnssf_NSSAIAvailability

import (
	"encoding/json"
	"time"
)

// NssfEventSubscriptionCreatedData This contains the information for created event subscription
type NssfEventSubscriptionCreatedData struct {
	SubscriptionId string `json:"subscriptionId"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	AuthorizedNssaiAvailabilityData []AuthorizedNssaiAvailabilityData `json:"authorizedNssaiAvailabilityData,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewNssfEventSubscriptionCreatedData instantiates a new NssfEventSubscriptionCreatedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssfEventSubscriptionCreatedData(subscriptionId string) *NssfEventSubscriptionCreatedData {
	this := NssfEventSubscriptionCreatedData{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewNssfEventSubscriptionCreatedDataWithDefaults instantiates a new NssfEventSubscriptionCreatedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssfEventSubscriptionCreatedDataWithDefaults() *NssfEventSubscriptionCreatedData {
	this := NssfEventSubscriptionCreatedData{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *NssfEventSubscriptionCreatedData) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreatedData) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *NssfEventSubscriptionCreatedData) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreatedData) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreatedData) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreatedData) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *NssfEventSubscriptionCreatedData) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetAuthorizedNssaiAvailabilityData returns the AuthorizedNssaiAvailabilityData field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreatedData) GetAuthorizedNssaiAvailabilityData() []AuthorizedNssaiAvailabilityData {
	if o == nil || isNil(o.AuthorizedNssaiAvailabilityData) {
		var ret []AuthorizedNssaiAvailabilityData
		return ret
	}
	return o.AuthorizedNssaiAvailabilityData
}

// GetAuthorizedNssaiAvailabilityDataOk returns a tuple with the AuthorizedNssaiAvailabilityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreatedData) GetAuthorizedNssaiAvailabilityDataOk() ([]AuthorizedNssaiAvailabilityData, bool) {
	if o == nil || isNil(o.AuthorizedNssaiAvailabilityData) {
    return nil, false
	}
	return o.AuthorizedNssaiAvailabilityData, true
}

// HasAuthorizedNssaiAvailabilityData returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreatedData) HasAuthorizedNssaiAvailabilityData() bool {
	if o != nil && !isNil(o.AuthorizedNssaiAvailabilityData) {
		return true
	}

	return false
}

// SetAuthorizedNssaiAvailabilityData gets a reference to the given []AuthorizedNssaiAvailabilityData and assigns it to the AuthorizedNssaiAvailabilityData field.
func (o *NssfEventSubscriptionCreatedData) SetAuthorizedNssaiAvailabilityData(v []AuthorizedNssaiAvailabilityData) {
	o.AuthorizedNssaiAvailabilityData = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreatedData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreatedData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreatedData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NssfEventSubscriptionCreatedData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o NssfEventSubscriptionCreatedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !isNil(o.AuthorizedNssaiAvailabilityData) {
		toSerialize["authorizedNssaiAvailabilityData"] = o.AuthorizedNssaiAvailabilityData
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableNssfEventSubscriptionCreatedData struct {
	value *NssfEventSubscriptionCreatedData
	isSet bool
}

func (v NullableNssfEventSubscriptionCreatedData) Get() *NssfEventSubscriptionCreatedData {
	return v.value
}

func (v *NullableNssfEventSubscriptionCreatedData) Set(val *NssfEventSubscriptionCreatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableNssfEventSubscriptionCreatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableNssfEventSubscriptionCreatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssfEventSubscriptionCreatedData(val *NssfEventSubscriptionCreatedData) *NullableNssfEventSubscriptionCreatedData {
	return &NullableNssfEventSubscriptionCreatedData{value: val, isSet: true}
}

func (v NullableNssfEventSubscriptionCreatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssfEventSubscriptionCreatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


