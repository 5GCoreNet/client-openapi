/*
M5_ServiceAccessInformation

5GMS AF M5 Service Access Information API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M5_ServiceAccessInformation

import (
	"encoding/json"
)

// ServiceAccessInformationResourceStreamingAccess struct for ServiceAccessInformationResourceStreamingAccess
type ServiceAccessInformationResourceStreamingAccess struct {
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	MediaPlayerEntry *string `json:"mediaPlayerEntry,omitempty"`
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	EMBMSServiceAnnouncementLocator *string `json:"eMBMSServiceAnnouncementLocator,omitempty"`
}

// NewServiceAccessInformationResourceStreamingAccess instantiates a new ServiceAccessInformationResourceStreamingAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccessInformationResourceStreamingAccess() *ServiceAccessInformationResourceStreamingAccess {
	this := ServiceAccessInformationResourceStreamingAccess{}
	return &this
}

// NewServiceAccessInformationResourceStreamingAccessWithDefaults instantiates a new ServiceAccessInformationResourceStreamingAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccessInformationResourceStreamingAccessWithDefaults() *ServiceAccessInformationResourceStreamingAccess {
	this := ServiceAccessInformationResourceStreamingAccess{}
	return &this
}

// GetMediaPlayerEntry returns the MediaPlayerEntry field value if set, zero value otherwise.
func (o *ServiceAccessInformationResourceStreamingAccess) GetMediaPlayerEntry() string {
	if o == nil || isNil(o.MediaPlayerEntry) {
		var ret string
		return ret
	}
	return *o.MediaPlayerEntry
}

// GetMediaPlayerEntryOk returns a tuple with the MediaPlayerEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceStreamingAccess) GetMediaPlayerEntryOk() (*string, bool) {
	if o == nil || isNil(o.MediaPlayerEntry) {
    return nil, false
	}
	return o.MediaPlayerEntry, true
}

// HasMediaPlayerEntry returns a boolean if a field has been set.
func (o *ServiceAccessInformationResourceStreamingAccess) HasMediaPlayerEntry() bool {
	if o != nil && !isNil(o.MediaPlayerEntry) {
		return true
	}

	return false
}

// SetMediaPlayerEntry gets a reference to the given string and assigns it to the MediaPlayerEntry field.
func (o *ServiceAccessInformationResourceStreamingAccess) SetMediaPlayerEntry(v string) {
	o.MediaPlayerEntry = &v
}

// GetEMBMSServiceAnnouncementLocator returns the EMBMSServiceAnnouncementLocator field value if set, zero value otherwise.
func (o *ServiceAccessInformationResourceStreamingAccess) GetEMBMSServiceAnnouncementLocator() string {
	if o == nil || isNil(o.EMBMSServiceAnnouncementLocator) {
		var ret string
		return ret
	}
	return *o.EMBMSServiceAnnouncementLocator
}

// GetEMBMSServiceAnnouncementLocatorOk returns a tuple with the EMBMSServiceAnnouncementLocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceStreamingAccess) GetEMBMSServiceAnnouncementLocatorOk() (*string, bool) {
	if o == nil || isNil(o.EMBMSServiceAnnouncementLocator) {
    return nil, false
	}
	return o.EMBMSServiceAnnouncementLocator, true
}

// HasEMBMSServiceAnnouncementLocator returns a boolean if a field has been set.
func (o *ServiceAccessInformationResourceStreamingAccess) HasEMBMSServiceAnnouncementLocator() bool {
	if o != nil && !isNil(o.EMBMSServiceAnnouncementLocator) {
		return true
	}

	return false
}

// SetEMBMSServiceAnnouncementLocator gets a reference to the given string and assigns it to the EMBMSServiceAnnouncementLocator field.
func (o *ServiceAccessInformationResourceStreamingAccess) SetEMBMSServiceAnnouncementLocator(v string) {
	o.EMBMSServiceAnnouncementLocator = &v
}

func (o ServiceAccessInformationResourceStreamingAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MediaPlayerEntry) {
		toSerialize["mediaPlayerEntry"] = o.MediaPlayerEntry
	}
	if !isNil(o.EMBMSServiceAnnouncementLocator) {
		toSerialize["eMBMSServiceAnnouncementLocator"] = o.EMBMSServiceAnnouncementLocator
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccessInformationResourceStreamingAccess struct {
	value *ServiceAccessInformationResourceStreamingAccess
	isSet bool
}

func (v NullableServiceAccessInformationResourceStreamingAccess) Get() *ServiceAccessInformationResourceStreamingAccess {
	return v.value
}

func (v *NullableServiceAccessInformationResourceStreamingAccess) Set(val *ServiceAccessInformationResourceStreamingAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccessInformationResourceStreamingAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccessInformationResourceStreamingAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccessInformationResourceStreamingAccess(val *ServiceAccessInformationResourceStreamingAccess) *NullableServiceAccessInformationResourceStreamingAccess {
	return &NullableServiceAccessInformationResourceStreamingAccess{value: val, isSet: true}
}

func (v NullableServiceAccessInformationResourceStreamingAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccessInformationResourceStreamingAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


