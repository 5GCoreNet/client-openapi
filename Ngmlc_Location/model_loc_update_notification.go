/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"time"
)

// LocUpdateNotification Location Update Notification
type LocUpdateNotification struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	LocationRequestType LocationRequestType `json:"locationRequestType"`
	LocationEstimate GeographicArea `json:"locationEstimate"`
	// Indicates value of the age of the location estimate.
	AgeOfLocationEstimate int32 `json:"ageOfLocationEstimate"`
	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfLocationEstimate *time.Time `json:"timestampOfLocationEstimate,omitempty"`
	AccuracyFulfilmentIndicator AccuracyFulfilmentIndicator `json:"accuracyFulfilmentIndicator"`
	CivicAddress *CivicAddress `json:"civicAddress,omitempty"`
	LcsQosClass LcsQosClass `json:"lcsQosClass"`
	AfId *string `json:"afId,omitempty"`
	// Contains the service identity
	ServiceIdentity *string `json:"serviceIdentity,omitempty"`
}

// NewLocUpdateNotification instantiates a new LocUpdateNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocUpdateNotification(locationRequestType LocationRequestType, locationEstimate GeographicArea, ageOfLocationEstimate int32, accuracyFulfilmentIndicator AccuracyFulfilmentIndicator, lcsQosClass LcsQosClass) *LocUpdateNotification {
	this := LocUpdateNotification{}
	this.LocationRequestType = locationRequestType
	this.LocationEstimate = locationEstimate
	this.AgeOfLocationEstimate = ageOfLocationEstimate
	this.AccuracyFulfilmentIndicator = accuracyFulfilmentIndicator
	this.LcsQosClass = lcsQosClass
	return &this
}

// NewLocUpdateNotificationWithDefaults instantiates a new LocUpdateNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocUpdateNotificationWithDefaults() *LocUpdateNotification {
	this := LocUpdateNotification{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *LocUpdateNotification) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *LocUpdateNotification) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *LocUpdateNotification) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *LocUpdateNotification) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *LocUpdateNotification) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *LocUpdateNotification) SetSupi(v string) {
	o.Supi = &v
}

// GetLocationRequestType returns the LocationRequestType field value
func (o *LocUpdateNotification) GetLocationRequestType() LocationRequestType {
	if o == nil {
		var ret LocationRequestType
		return ret
	}

	return o.LocationRequestType
}

// GetLocationRequestTypeOk returns a tuple with the LocationRequestType field value
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetLocationRequestTypeOk() (*LocationRequestType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocationRequestType, true
}

// SetLocationRequestType sets field value
func (o *LocUpdateNotification) SetLocationRequestType(v LocationRequestType) {
	o.LocationRequestType = v
}

// GetLocationEstimate returns the LocationEstimate field value
func (o *LocUpdateNotification) GetLocationEstimate() GeographicArea {
	if o == nil {
		var ret GeographicArea
		return ret
	}

	return o.LocationEstimate
}

// GetLocationEstimateOk returns a tuple with the LocationEstimate field value
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetLocationEstimateOk() (*GeographicArea, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocationEstimate, true
}

// SetLocationEstimate sets field value
func (o *LocUpdateNotification) SetLocationEstimate(v GeographicArea) {
	o.LocationEstimate = v
}

// GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field value
func (o *LocUpdateNotification) GetAgeOfLocationEstimate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgeOfLocationEstimate
}

// GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field value
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetAgeOfLocationEstimateOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AgeOfLocationEstimate, true
}

// SetAgeOfLocationEstimate sets field value
func (o *LocUpdateNotification) SetAgeOfLocationEstimate(v int32) {
	o.AgeOfLocationEstimate = v
}

// GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field value if set, zero value otherwise.
func (o *LocUpdateNotification) GetTimestampOfLocationEstimate() time.Time {
	if o == nil || isNil(o.TimestampOfLocationEstimate) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOfLocationEstimate
}

// GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetTimestampOfLocationEstimateOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimestampOfLocationEstimate) {
    return nil, false
	}
	return o.TimestampOfLocationEstimate, true
}

// HasTimestampOfLocationEstimate returns a boolean if a field has been set.
func (o *LocUpdateNotification) HasTimestampOfLocationEstimate() bool {
	if o != nil && !isNil(o.TimestampOfLocationEstimate) {
		return true
	}

	return false
}

// SetTimestampOfLocationEstimate gets a reference to the given time.Time and assigns it to the TimestampOfLocationEstimate field.
func (o *LocUpdateNotification) SetTimestampOfLocationEstimate(v time.Time) {
	o.TimestampOfLocationEstimate = &v
}

// GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field value
func (o *LocUpdateNotification) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator {
	if o == nil {
		var ret AccuracyFulfilmentIndicator
		return ret
	}

	return o.AccuracyFulfilmentIndicator
}

// GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field value
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccuracyFulfilmentIndicator, true
}

// SetAccuracyFulfilmentIndicator sets field value
func (o *LocUpdateNotification) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator) {
	o.AccuracyFulfilmentIndicator = v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *LocUpdateNotification) GetCivicAddress() CivicAddress {
	if o == nil || isNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddress) {
    return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *LocUpdateNotification) HasCivicAddress() bool {
	if o != nil && !isNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given CivicAddress and assigns it to the CivicAddress field.
func (o *LocUpdateNotification) SetCivicAddress(v CivicAddress) {
	o.CivicAddress = &v
}

// GetLcsQosClass returns the LcsQosClass field value
func (o *LocUpdateNotification) GetLcsQosClass() LcsQosClass {
	if o == nil {
		var ret LcsQosClass
		return ret
	}

	return o.LcsQosClass
}

// GetLcsQosClassOk returns a tuple with the LcsQosClass field value
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetLcsQosClassOk() (*LcsQosClass, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LcsQosClass, true
}

// SetLcsQosClass sets field value
func (o *LocUpdateNotification) SetLcsQosClass(v LcsQosClass) {
	o.LcsQosClass = v
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *LocUpdateNotification) GetAfId() string {
	if o == nil || isNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetAfIdOk() (*string, bool) {
	if o == nil || isNil(o.AfId) {
    return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *LocUpdateNotification) HasAfId() bool {
	if o != nil && !isNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *LocUpdateNotification) SetAfId(v string) {
	o.AfId = &v
}

// GetServiceIdentity returns the ServiceIdentity field value if set, zero value otherwise.
func (o *LocUpdateNotification) GetServiceIdentity() string {
	if o == nil || isNil(o.ServiceIdentity) {
		var ret string
		return ret
	}
	return *o.ServiceIdentity
}

// GetServiceIdentityOk returns a tuple with the ServiceIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocUpdateNotification) GetServiceIdentityOk() (*string, bool) {
	if o == nil || isNil(o.ServiceIdentity) {
    return nil, false
	}
	return o.ServiceIdentity, true
}

// HasServiceIdentity returns a boolean if a field has been set.
func (o *LocUpdateNotification) HasServiceIdentity() bool {
	if o != nil && !isNil(o.ServiceIdentity) {
		return true
	}

	return false
}

// SetServiceIdentity gets a reference to the given string and assigns it to the ServiceIdentity field.
func (o *LocUpdateNotification) SetServiceIdentity(v string) {
	o.ServiceIdentity = &v
}

func (o LocUpdateNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if true {
		toSerialize["locationRequestType"] = o.LocationRequestType
	}
	if true {
		toSerialize["locationEstimate"] = o.LocationEstimate
	}
	if true {
		toSerialize["ageOfLocationEstimate"] = o.AgeOfLocationEstimate
	}
	if !isNil(o.TimestampOfLocationEstimate) {
		toSerialize["timestampOfLocationEstimate"] = o.TimestampOfLocationEstimate
	}
	if true {
		toSerialize["accuracyFulfilmentIndicator"] = o.AccuracyFulfilmentIndicator
	}
	if !isNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	if true {
		toSerialize["lcsQosClass"] = o.LcsQosClass
	}
	if !isNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	if !isNil(o.ServiceIdentity) {
		toSerialize["serviceIdentity"] = o.ServiceIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableLocUpdateNotification struct {
	value *LocUpdateNotification
	isSet bool
}

func (v NullableLocUpdateNotification) Get() *LocUpdateNotification {
	return v.value
}

func (v *NullableLocUpdateNotification) Set(val *LocUpdateNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableLocUpdateNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableLocUpdateNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocUpdateNotification(val *LocUpdateNotification) *NullableLocUpdateNotification {
	return &NullableLocUpdateNotification{value: val, isSet: true}
}

func (v NullableLocUpdateNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocUpdateNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


