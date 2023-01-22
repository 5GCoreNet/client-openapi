/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
	"time"
)

// LocationData Information within Determine Location Response.
type LocationData struct {
	LocationEstimate GeographicArea `json:"locationEstimate"`
	AccuracyFulfilmentIndicator *AccuracyFulfilmentIndicator `json:"accuracyFulfilmentIndicator,omitempty"`
	// Indicates value of the age of the location estimate.
	AgeOfLocationEstimate *int32 `json:"ageOfLocationEstimate,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimestampOfLocationEstimate *time.Time `json:"timestampOfLocationEstimate,omitempty"`
	VelocityEstimate *VelocityEstimate `json:"velocityEstimate,omitempty"`
	CivicAddress *CivicAddress `json:"civicAddress,omitempty"`
	LocalLocationEstimate *LocalArea `json:"localLocationEstimate,omitempty"`
	PositioningDataList []PositioningMethodAndUsage `json:"positioningDataList,omitempty"`
	GnssPositioningDataList []GnssPositioningMethodAndUsage `json:"gnssPositioningDataList,omitempty"`
	Ecgi *Ecgi `json:"ecgi,omitempty"`
	Ncgi *Ncgi `json:"ncgi,omitempty"`
	// Indicates value of altitude.
	Altitude *float64 `json:"altitude,omitempty"`
	// Specifies the measured uncompensated atmospheric pressure.
	BarometricPressure *int32 `json:"barometricPressure,omitempty"`
	// LMF identification.
	ServingLMFIdentification *string `json:"servingLMFIdentification,omitempty"`
	// Positioning capabilities supported by the UE. A string encoding the \"ProvideCapabilities-r9-IEs\" IE as specified in clause 6.3 of 3GPP TS 37.355 (start from octet 1).
	UePositioningCap *string `json:"uePositioningCap,omitempty"`
	UeAreaInd *UeAreaIndication `json:"ueAreaInd,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	AchievedQos *MinorLocationQoS `json:"achievedQos,omitempty"`
}

// NewLocationData instantiates a new LocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationData(locationEstimate GeographicArea) *LocationData {
	this := LocationData{}
	this.LocationEstimate = locationEstimate
	return &this
}

// NewLocationDataWithDefaults instantiates a new LocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationDataWithDefaults() *LocationData {
	this := LocationData{}
	return &this
}

// GetLocationEstimate returns the LocationEstimate field value
func (o *LocationData) GetLocationEstimate() GeographicArea {
	if o == nil {
		var ret GeographicArea
		return ret
	}

	return o.LocationEstimate
}

// GetLocationEstimateOk returns a tuple with the LocationEstimate field value
// and a boolean to check if the value has been set.
func (o *LocationData) GetLocationEstimateOk() (*GeographicArea, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocationEstimate, true
}

// SetLocationEstimate sets field value
func (o *LocationData) SetLocationEstimate(v GeographicArea) {
	o.LocationEstimate = v
}

// GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field value if set, zero value otherwise.
func (o *LocationData) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator {
	if o == nil || isNil(o.AccuracyFulfilmentIndicator) {
		var ret AccuracyFulfilmentIndicator
		return ret
	}
	return *o.AccuracyFulfilmentIndicator
}

// GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool) {
	if o == nil || isNil(o.AccuracyFulfilmentIndicator) {
    return nil, false
	}
	return o.AccuracyFulfilmentIndicator, true
}

// HasAccuracyFulfilmentIndicator returns a boolean if a field has been set.
func (o *LocationData) HasAccuracyFulfilmentIndicator() bool {
	if o != nil && !isNil(o.AccuracyFulfilmentIndicator) {
		return true
	}

	return false
}

// SetAccuracyFulfilmentIndicator gets a reference to the given AccuracyFulfilmentIndicator and assigns it to the AccuracyFulfilmentIndicator field.
func (o *LocationData) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator) {
	o.AccuracyFulfilmentIndicator = &v
}

// GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field value if set, zero value otherwise.
func (o *LocationData) GetAgeOfLocationEstimate() int32 {
	if o == nil || isNil(o.AgeOfLocationEstimate) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationEstimate
}

// GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetAgeOfLocationEstimateOk() (*int32, bool) {
	if o == nil || isNil(o.AgeOfLocationEstimate) {
    return nil, false
	}
	return o.AgeOfLocationEstimate, true
}

// HasAgeOfLocationEstimate returns a boolean if a field has been set.
func (o *LocationData) HasAgeOfLocationEstimate() bool {
	if o != nil && !isNil(o.AgeOfLocationEstimate) {
		return true
	}

	return false
}

// SetAgeOfLocationEstimate gets a reference to the given int32 and assigns it to the AgeOfLocationEstimate field.
func (o *LocationData) SetAgeOfLocationEstimate(v int32) {
	o.AgeOfLocationEstimate = &v
}

// GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field value if set, zero value otherwise.
func (o *LocationData) GetTimestampOfLocationEstimate() time.Time {
	if o == nil || isNil(o.TimestampOfLocationEstimate) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOfLocationEstimate
}

// GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetTimestampOfLocationEstimateOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimestampOfLocationEstimate) {
    return nil, false
	}
	return o.TimestampOfLocationEstimate, true
}

// HasTimestampOfLocationEstimate returns a boolean if a field has been set.
func (o *LocationData) HasTimestampOfLocationEstimate() bool {
	if o != nil && !isNil(o.TimestampOfLocationEstimate) {
		return true
	}

	return false
}

// SetTimestampOfLocationEstimate gets a reference to the given time.Time and assigns it to the TimestampOfLocationEstimate field.
func (o *LocationData) SetTimestampOfLocationEstimate(v time.Time) {
	o.TimestampOfLocationEstimate = &v
}

// GetVelocityEstimate returns the VelocityEstimate field value if set, zero value otherwise.
func (o *LocationData) GetVelocityEstimate() VelocityEstimate {
	if o == nil || isNil(o.VelocityEstimate) {
		var ret VelocityEstimate
		return ret
	}
	return *o.VelocityEstimate
}

// GetVelocityEstimateOk returns a tuple with the VelocityEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetVelocityEstimateOk() (*VelocityEstimate, bool) {
	if o == nil || isNil(o.VelocityEstimate) {
    return nil, false
	}
	return o.VelocityEstimate, true
}

// HasVelocityEstimate returns a boolean if a field has been set.
func (o *LocationData) HasVelocityEstimate() bool {
	if o != nil && !isNil(o.VelocityEstimate) {
		return true
	}

	return false
}

// SetVelocityEstimate gets a reference to the given VelocityEstimate and assigns it to the VelocityEstimate field.
func (o *LocationData) SetVelocityEstimate(v VelocityEstimate) {
	o.VelocityEstimate = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *LocationData) GetCivicAddress() CivicAddress {
	if o == nil || isNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddress) {
    return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *LocationData) HasCivicAddress() bool {
	if o != nil && !isNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given CivicAddress and assigns it to the CivicAddress field.
func (o *LocationData) SetCivicAddress(v CivicAddress) {
	o.CivicAddress = &v
}

// GetLocalLocationEstimate returns the LocalLocationEstimate field value if set, zero value otherwise.
func (o *LocationData) GetLocalLocationEstimate() LocalArea {
	if o == nil || isNil(o.LocalLocationEstimate) {
		var ret LocalArea
		return ret
	}
	return *o.LocalLocationEstimate
}

// GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetLocalLocationEstimateOk() (*LocalArea, bool) {
	if o == nil || isNil(o.LocalLocationEstimate) {
    return nil, false
	}
	return o.LocalLocationEstimate, true
}

// HasLocalLocationEstimate returns a boolean if a field has been set.
func (o *LocationData) HasLocalLocationEstimate() bool {
	if o != nil && !isNil(o.LocalLocationEstimate) {
		return true
	}

	return false
}

// SetLocalLocationEstimate gets a reference to the given LocalArea and assigns it to the LocalLocationEstimate field.
func (o *LocationData) SetLocalLocationEstimate(v LocalArea) {
	o.LocalLocationEstimate = &v
}

// GetPositioningDataList returns the PositioningDataList field value if set, zero value otherwise.
func (o *LocationData) GetPositioningDataList() []PositioningMethodAndUsage {
	if o == nil || isNil(o.PositioningDataList) {
		var ret []PositioningMethodAndUsage
		return ret
	}
	return o.PositioningDataList
}

// GetPositioningDataListOk returns a tuple with the PositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetPositioningDataListOk() ([]PositioningMethodAndUsage, bool) {
	if o == nil || isNil(o.PositioningDataList) {
    return nil, false
	}
	return o.PositioningDataList, true
}

// HasPositioningDataList returns a boolean if a field has been set.
func (o *LocationData) HasPositioningDataList() bool {
	if o != nil && !isNil(o.PositioningDataList) {
		return true
	}

	return false
}

// SetPositioningDataList gets a reference to the given []PositioningMethodAndUsage and assigns it to the PositioningDataList field.
func (o *LocationData) SetPositioningDataList(v []PositioningMethodAndUsage) {
	o.PositioningDataList = v
}

// GetGnssPositioningDataList returns the GnssPositioningDataList field value if set, zero value otherwise.
func (o *LocationData) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage {
	if o == nil || isNil(o.GnssPositioningDataList) {
		var ret []GnssPositioningMethodAndUsage
		return ret
	}
	return o.GnssPositioningDataList
}

// GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetGnssPositioningDataListOk() ([]GnssPositioningMethodAndUsage, bool) {
	if o == nil || isNil(o.GnssPositioningDataList) {
    return nil, false
	}
	return o.GnssPositioningDataList, true
}

// HasGnssPositioningDataList returns a boolean if a field has been set.
func (o *LocationData) HasGnssPositioningDataList() bool {
	if o != nil && !isNil(o.GnssPositioningDataList) {
		return true
	}

	return false
}

// SetGnssPositioningDataList gets a reference to the given []GnssPositioningMethodAndUsage and assigns it to the GnssPositioningDataList field.
func (o *LocationData) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage) {
	o.GnssPositioningDataList = v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *LocationData) GetEcgi() Ecgi {
	if o == nil || isNil(o.Ecgi) {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || isNil(o.Ecgi) {
    return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *LocationData) HasEcgi() bool {
	if o != nil && !isNil(o.Ecgi) {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *LocationData) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

// GetNcgi returns the Ncgi field value if set, zero value otherwise.
func (o *LocationData) GetNcgi() Ncgi {
	if o == nil || isNil(o.Ncgi) {
		var ret Ncgi
		return ret
	}
	return *o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetNcgiOk() (*Ncgi, bool) {
	if o == nil || isNil(o.Ncgi) {
    return nil, false
	}
	return o.Ncgi, true
}

// HasNcgi returns a boolean if a field has been set.
func (o *LocationData) HasNcgi() bool {
	if o != nil && !isNil(o.Ncgi) {
		return true
	}

	return false
}

// SetNcgi gets a reference to the given Ncgi and assigns it to the Ncgi field.
func (o *LocationData) SetNcgi(v Ncgi) {
	o.Ncgi = &v
}

// GetAltitude returns the Altitude field value if set, zero value otherwise.
func (o *LocationData) GetAltitude() float64 {
	if o == nil || isNil(o.Altitude) {
		var ret float64
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetAltitudeOk() (*float64, bool) {
	if o == nil || isNil(o.Altitude) {
    return nil, false
	}
	return o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *LocationData) HasAltitude() bool {
	if o != nil && !isNil(o.Altitude) {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given float64 and assigns it to the Altitude field.
func (o *LocationData) SetAltitude(v float64) {
	o.Altitude = &v
}

// GetBarometricPressure returns the BarometricPressure field value if set, zero value otherwise.
func (o *LocationData) GetBarometricPressure() int32 {
	if o == nil || isNil(o.BarometricPressure) {
		var ret int32
		return ret
	}
	return *o.BarometricPressure
}

// GetBarometricPressureOk returns a tuple with the BarometricPressure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetBarometricPressureOk() (*int32, bool) {
	if o == nil || isNil(o.BarometricPressure) {
    return nil, false
	}
	return o.BarometricPressure, true
}

// HasBarometricPressure returns a boolean if a field has been set.
func (o *LocationData) HasBarometricPressure() bool {
	if o != nil && !isNil(o.BarometricPressure) {
		return true
	}

	return false
}

// SetBarometricPressure gets a reference to the given int32 and assigns it to the BarometricPressure field.
func (o *LocationData) SetBarometricPressure(v int32) {
	o.BarometricPressure = &v
}

// GetServingLMFIdentification returns the ServingLMFIdentification field value if set, zero value otherwise.
func (o *LocationData) GetServingLMFIdentification() string {
	if o == nil || isNil(o.ServingLMFIdentification) {
		var ret string
		return ret
	}
	return *o.ServingLMFIdentification
}

// GetServingLMFIdentificationOk returns a tuple with the ServingLMFIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetServingLMFIdentificationOk() (*string, bool) {
	if o == nil || isNil(o.ServingLMFIdentification) {
    return nil, false
	}
	return o.ServingLMFIdentification, true
}

// HasServingLMFIdentification returns a boolean if a field has been set.
func (o *LocationData) HasServingLMFIdentification() bool {
	if o != nil && !isNil(o.ServingLMFIdentification) {
		return true
	}

	return false
}

// SetServingLMFIdentification gets a reference to the given string and assigns it to the ServingLMFIdentification field.
func (o *LocationData) SetServingLMFIdentification(v string) {
	o.ServingLMFIdentification = &v
}

// GetUePositioningCap returns the UePositioningCap field value if set, zero value otherwise.
func (o *LocationData) GetUePositioningCap() string {
	if o == nil || isNil(o.UePositioningCap) {
		var ret string
		return ret
	}
	return *o.UePositioningCap
}

// GetUePositioningCapOk returns a tuple with the UePositioningCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetUePositioningCapOk() (*string, bool) {
	if o == nil || isNil(o.UePositioningCap) {
    return nil, false
	}
	return o.UePositioningCap, true
}

// HasUePositioningCap returns a boolean if a field has been set.
func (o *LocationData) HasUePositioningCap() bool {
	if o != nil && !isNil(o.UePositioningCap) {
		return true
	}

	return false
}

// SetUePositioningCap gets a reference to the given string and assigns it to the UePositioningCap field.
func (o *LocationData) SetUePositioningCap(v string) {
	o.UePositioningCap = &v
}

// GetUeAreaInd returns the UeAreaInd field value if set, zero value otherwise.
func (o *LocationData) GetUeAreaInd() UeAreaIndication {
	if o == nil || isNil(o.UeAreaInd) {
		var ret UeAreaIndication
		return ret
	}
	return *o.UeAreaInd
}

// GetUeAreaIndOk returns a tuple with the UeAreaInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetUeAreaIndOk() (*UeAreaIndication, bool) {
	if o == nil || isNil(o.UeAreaInd) {
    return nil, false
	}
	return o.UeAreaInd, true
}

// HasUeAreaInd returns a boolean if a field has been set.
func (o *LocationData) HasUeAreaInd() bool {
	if o != nil && !isNil(o.UeAreaInd) {
		return true
	}

	return false
}

// SetUeAreaInd gets a reference to the given UeAreaIndication and assigns it to the UeAreaInd field.
func (o *LocationData) SetUeAreaInd(v UeAreaIndication) {
	o.UeAreaInd = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *LocationData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *LocationData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *LocationData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAchievedQos returns the AchievedQos field value if set, zero value otherwise.
func (o *LocationData) GetAchievedQos() MinorLocationQoS {
	if o == nil || isNil(o.AchievedQos) {
		var ret MinorLocationQoS
		return ret
	}
	return *o.AchievedQos
}

// GetAchievedQosOk returns a tuple with the AchievedQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationData) GetAchievedQosOk() (*MinorLocationQoS, bool) {
	if o == nil || isNil(o.AchievedQos) {
    return nil, false
	}
	return o.AchievedQos, true
}

// HasAchievedQos returns a boolean if a field has been set.
func (o *LocationData) HasAchievedQos() bool {
	if o != nil && !isNil(o.AchievedQos) {
		return true
	}

	return false
}

// SetAchievedQos gets a reference to the given MinorLocationQoS and assigns it to the AchievedQos field.
func (o *LocationData) SetAchievedQos(v MinorLocationQoS) {
	o.AchievedQos = &v
}

func (o LocationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locationEstimate"] = o.LocationEstimate
	}
	if !isNil(o.AccuracyFulfilmentIndicator) {
		toSerialize["accuracyFulfilmentIndicator"] = o.AccuracyFulfilmentIndicator
	}
	if !isNil(o.AgeOfLocationEstimate) {
		toSerialize["ageOfLocationEstimate"] = o.AgeOfLocationEstimate
	}
	if !isNil(o.TimestampOfLocationEstimate) {
		toSerialize["timestampOfLocationEstimate"] = o.TimestampOfLocationEstimate
	}
	if !isNil(o.VelocityEstimate) {
		toSerialize["velocityEstimate"] = o.VelocityEstimate
	}
	if !isNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	if !isNil(o.LocalLocationEstimate) {
		toSerialize["localLocationEstimate"] = o.LocalLocationEstimate
	}
	if !isNil(o.PositioningDataList) {
		toSerialize["positioningDataList"] = o.PositioningDataList
	}
	if !isNil(o.GnssPositioningDataList) {
		toSerialize["gnssPositioningDataList"] = o.GnssPositioningDataList
	}
	if !isNil(o.Ecgi) {
		toSerialize["ecgi"] = o.Ecgi
	}
	if !isNil(o.Ncgi) {
		toSerialize["ncgi"] = o.Ncgi
	}
	if !isNil(o.Altitude) {
		toSerialize["altitude"] = o.Altitude
	}
	if !isNil(o.BarometricPressure) {
		toSerialize["barometricPressure"] = o.BarometricPressure
	}
	if !isNil(o.ServingLMFIdentification) {
		toSerialize["servingLMFIdentification"] = o.ServingLMFIdentification
	}
	if !isNil(o.UePositioningCap) {
		toSerialize["uePositioningCap"] = o.UePositioningCap
	}
	if !isNil(o.UeAreaInd) {
		toSerialize["ueAreaInd"] = o.UeAreaInd
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.AchievedQos) {
		toSerialize["achievedQos"] = o.AchievedQos
	}
	return json.Marshal(toSerialize)
}

type NullableLocationData struct {
	value *LocationData
	isSet bool
}

func (v NullableLocationData) Get() *LocationData {
	return v.value
}

func (v *NullableLocationData) Set(val *LocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationData(val *LocationData) *NullableLocationData {
	return &NullableLocationData{value: val, isSet: true}
}

func (v NullableLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


