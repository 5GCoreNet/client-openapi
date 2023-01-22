/*
Nudm_MT

UDM MT Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_MT

import (
	"encoding/json"
)

// LocationInfoResult Represents the requested location information.
type LocationInfoResult struct {
	VPlmnId *PlmnId `json:"vPlmnId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfInstanceId *string `json:"amfInstanceId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmsfInstanceId *string `json:"smsfInstanceId,omitempty"`
	Ncgi *Ncgi `json:"ncgi,omitempty"`
	Ecgi *Ecgi `json:"ecgi,omitempty"`
	Tai *Tai `json:"tai,omitempty"`
	CurrentLoc *bool `json:"currentLoc,omitempty"`
	GeoInfo *GeographicArea `json:"geoInfo,omitempty"`
	// Indicates value of the age of the location estimate.
	LocationAge *int32 `json:"locationAge,omitempty"`
	RatType *RatType `json:"ratType,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	Timezone *string `json:"timezone,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewLocationInfoResult instantiates a new LocationInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationInfoResult() *LocationInfoResult {
	this := LocationInfoResult{}
	return &this
}

// NewLocationInfoResultWithDefaults instantiates a new LocationInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationInfoResultWithDefaults() *LocationInfoResult {
	this := LocationInfoResult{}
	return &this
}

// GetVPlmnId returns the VPlmnId field value if set, zero value otherwise.
func (o *LocationInfoResult) GetVPlmnId() PlmnId {
	if o == nil || isNil(o.VPlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.VPlmnId
}

// GetVPlmnIdOk returns a tuple with the VPlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetVPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.VPlmnId) {
    return nil, false
	}
	return o.VPlmnId, true
}

// HasVPlmnId returns a boolean if a field has been set.
func (o *LocationInfoResult) HasVPlmnId() bool {
	if o != nil && !isNil(o.VPlmnId) {
		return true
	}

	return false
}

// SetVPlmnId gets a reference to the given PlmnId and assigns it to the VPlmnId field.
func (o *LocationInfoResult) SetVPlmnId(v PlmnId) {
	o.VPlmnId = &v
}

// GetAmfInstanceId returns the AmfInstanceId field value if set, zero value otherwise.
func (o *LocationInfoResult) GetAmfInstanceId() string {
	if o == nil || isNil(o.AmfInstanceId) {
		var ret string
		return ret
	}
	return *o.AmfInstanceId
}

// GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetAmfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfInstanceId) {
    return nil, false
	}
	return o.AmfInstanceId, true
}

// HasAmfInstanceId returns a boolean if a field has been set.
func (o *LocationInfoResult) HasAmfInstanceId() bool {
	if o != nil && !isNil(o.AmfInstanceId) {
		return true
	}

	return false
}

// SetAmfInstanceId gets a reference to the given string and assigns it to the AmfInstanceId field.
func (o *LocationInfoResult) SetAmfInstanceId(v string) {
	o.AmfInstanceId = &v
}

// GetSmsfInstanceId returns the SmsfInstanceId field value if set, zero value otherwise.
func (o *LocationInfoResult) GetSmsfInstanceId() string {
	if o == nil || isNil(o.SmsfInstanceId) {
		var ret string
		return ret
	}
	return *o.SmsfInstanceId
}

// GetSmsfInstanceIdOk returns a tuple with the SmsfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetSmsfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.SmsfInstanceId) {
    return nil, false
	}
	return o.SmsfInstanceId, true
}

// HasSmsfInstanceId returns a boolean if a field has been set.
func (o *LocationInfoResult) HasSmsfInstanceId() bool {
	if o != nil && !isNil(o.SmsfInstanceId) {
		return true
	}

	return false
}

// SetSmsfInstanceId gets a reference to the given string and assigns it to the SmsfInstanceId field.
func (o *LocationInfoResult) SetSmsfInstanceId(v string) {
	o.SmsfInstanceId = &v
}

// GetNcgi returns the Ncgi field value if set, zero value otherwise.
func (o *LocationInfoResult) GetNcgi() Ncgi {
	if o == nil || isNil(o.Ncgi) {
		var ret Ncgi
		return ret
	}
	return *o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetNcgiOk() (*Ncgi, bool) {
	if o == nil || isNil(o.Ncgi) {
    return nil, false
	}
	return o.Ncgi, true
}

// HasNcgi returns a boolean if a field has been set.
func (o *LocationInfoResult) HasNcgi() bool {
	if o != nil && !isNil(o.Ncgi) {
		return true
	}

	return false
}

// SetNcgi gets a reference to the given Ncgi and assigns it to the Ncgi field.
func (o *LocationInfoResult) SetNcgi(v Ncgi) {
	o.Ncgi = &v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *LocationInfoResult) GetEcgi() Ecgi {
	if o == nil || isNil(o.Ecgi) {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || isNil(o.Ecgi) {
    return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *LocationInfoResult) HasEcgi() bool {
	if o != nil && !isNil(o.Ecgi) {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *LocationInfoResult) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

// GetTai returns the Tai field value if set, zero value otherwise.
func (o *LocationInfoResult) GetTai() Tai {
	if o == nil || isNil(o.Tai) {
		var ret Tai
		return ret
	}
	return *o.Tai
}

// GetTaiOk returns a tuple with the Tai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetTaiOk() (*Tai, bool) {
	if o == nil || isNil(o.Tai) {
    return nil, false
	}
	return o.Tai, true
}

// HasTai returns a boolean if a field has been set.
func (o *LocationInfoResult) HasTai() bool {
	if o != nil && !isNil(o.Tai) {
		return true
	}

	return false
}

// SetTai gets a reference to the given Tai and assigns it to the Tai field.
func (o *LocationInfoResult) SetTai(v Tai) {
	o.Tai = &v
}

// GetCurrentLoc returns the CurrentLoc field value if set, zero value otherwise.
func (o *LocationInfoResult) GetCurrentLoc() bool {
	if o == nil || isNil(o.CurrentLoc) {
		var ret bool
		return ret
	}
	return *o.CurrentLoc
}

// GetCurrentLocOk returns a tuple with the CurrentLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetCurrentLocOk() (*bool, bool) {
	if o == nil || isNil(o.CurrentLoc) {
    return nil, false
	}
	return o.CurrentLoc, true
}

// HasCurrentLoc returns a boolean if a field has been set.
func (o *LocationInfoResult) HasCurrentLoc() bool {
	if o != nil && !isNil(o.CurrentLoc) {
		return true
	}

	return false
}

// SetCurrentLoc gets a reference to the given bool and assigns it to the CurrentLoc field.
func (o *LocationInfoResult) SetCurrentLoc(v bool) {
	o.CurrentLoc = &v
}

// GetGeoInfo returns the GeoInfo field value if set, zero value otherwise.
func (o *LocationInfoResult) GetGeoInfo() GeographicArea {
	if o == nil || isNil(o.GeoInfo) {
		var ret GeographicArea
		return ret
	}
	return *o.GeoInfo
}

// GetGeoInfoOk returns a tuple with the GeoInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetGeoInfoOk() (*GeographicArea, bool) {
	if o == nil || isNil(o.GeoInfo) {
    return nil, false
	}
	return o.GeoInfo, true
}

// HasGeoInfo returns a boolean if a field has been set.
func (o *LocationInfoResult) HasGeoInfo() bool {
	if o != nil && !isNil(o.GeoInfo) {
		return true
	}

	return false
}

// SetGeoInfo gets a reference to the given GeographicArea and assigns it to the GeoInfo field.
func (o *LocationInfoResult) SetGeoInfo(v GeographicArea) {
	o.GeoInfo = &v
}

// GetLocationAge returns the LocationAge field value if set, zero value otherwise.
func (o *LocationInfoResult) GetLocationAge() int32 {
	if o == nil || isNil(o.LocationAge) {
		var ret int32
		return ret
	}
	return *o.LocationAge
}

// GetLocationAgeOk returns a tuple with the LocationAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetLocationAgeOk() (*int32, bool) {
	if o == nil || isNil(o.LocationAge) {
    return nil, false
	}
	return o.LocationAge, true
}

// HasLocationAge returns a boolean if a field has been set.
func (o *LocationInfoResult) HasLocationAge() bool {
	if o != nil && !isNil(o.LocationAge) {
		return true
	}

	return false
}

// SetLocationAge gets a reference to the given int32 and assigns it to the LocationAge field.
func (o *LocationInfoResult) SetLocationAge(v int32) {
	o.LocationAge = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *LocationInfoResult) GetRatType() RatType {
	if o == nil || isNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetRatTypeOk() (*RatType, bool) {
	if o == nil || isNil(o.RatType) {
    return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *LocationInfoResult) HasRatType() bool {
	if o != nil && !isNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *LocationInfoResult) SetRatType(v RatType) {
	o.RatType = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *LocationInfoResult) GetTimezone() string {
	if o == nil || isNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetTimezoneOk() (*string, bool) {
	if o == nil || isNil(o.Timezone) {
    return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *LocationInfoResult) HasTimezone() bool {
	if o != nil && !isNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *LocationInfoResult) SetTimezone(v string) {
	o.Timezone = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *LocationInfoResult) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfoResult) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *LocationInfoResult) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *LocationInfoResult) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o LocationInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VPlmnId) {
		toSerialize["vPlmnId"] = o.VPlmnId
	}
	if !isNil(o.AmfInstanceId) {
		toSerialize["amfInstanceId"] = o.AmfInstanceId
	}
	if !isNil(o.SmsfInstanceId) {
		toSerialize["smsfInstanceId"] = o.SmsfInstanceId
	}
	if !isNil(o.Ncgi) {
		toSerialize["ncgi"] = o.Ncgi
	}
	if !isNil(o.Ecgi) {
		toSerialize["ecgi"] = o.Ecgi
	}
	if !isNil(o.Tai) {
		toSerialize["tai"] = o.Tai
	}
	if !isNil(o.CurrentLoc) {
		toSerialize["currentLoc"] = o.CurrentLoc
	}
	if !isNil(o.GeoInfo) {
		toSerialize["geoInfo"] = o.GeoInfo
	}
	if !isNil(o.LocationAge) {
		toSerialize["locationAge"] = o.LocationAge
	}
	if !isNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !isNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableLocationInfoResult struct {
	value *LocationInfoResult
	isSet bool
}

func (v NullableLocationInfoResult) Get() *LocationInfoResult {
	return v.value
}

func (v *NullableLocationInfoResult) Set(val *LocationInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationInfoResult(val *LocationInfoResult) *NullableLocationInfoResult {
	return &NullableLocationInfoResult{value: val, isSet: true}
}

func (v NullableLocationInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


