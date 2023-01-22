/*
VAE_HDMapDynamicInfo

API for VAE HDMapDynamicInfo Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package VAE_HDMapDynamicInfo

import (
	"encoding/json"
	"time"
)

// NrLocation Contains the NR user location.
type NrLocation struct {
	Tai Tai `json:"tai"`
	Ncgi Ncgi `json:"ncgi"`
	IgnoreNcgi *bool `json:"ignoreNcgi,omitempty"`
	// The value represents the elapsed time in minutes since the last network contact of the mobile station. Value \"0\" indicates that the location information was obtained after a successful paging procedure for Active Location Retrieval when the UE is in idle mode or after a successful  NG-RAN location reporting procedure with the eNB when the UE is in connected mode. Any other value than \"0\" indicates that the location information is the last known one. See 3GPP TS 29.002 clause 17.7.8. 
	AgeOfLocationInformation *int32 `json:"ageOfLocationInformation,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	UeLocationTimestamp *time.Time `json:"ueLocationTimestamp,omitempty"`
	// Refer to geographical Information. See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeographicalInformation *string `json:"geographicalInformation,omitempty"`
	// Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) [24] clause 3.88.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeodeticInformation *string `json:"geodeticInformation,omitempty"`
	GlobalGnbId *GlobalRanNodeId `json:"globalGnbId,omitempty"`
}

// NewNrLocation instantiates a new NrLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrLocation(tai Tai, ncgi Ncgi) *NrLocation {
	this := NrLocation{}
	this.Tai = tai
	this.Ncgi = ncgi
	var ignoreNcgi bool = false
	this.IgnoreNcgi = &ignoreNcgi
	return &this
}

// NewNrLocationWithDefaults instantiates a new NrLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrLocationWithDefaults() *NrLocation {
	this := NrLocation{}
	var ignoreNcgi bool = false
	this.IgnoreNcgi = &ignoreNcgi
	return &this
}

// GetTai returns the Tai field value
func (o *NrLocation) GetTai() Tai {
	if o == nil {
		var ret Tai
		return ret
	}

	return o.Tai
}

// GetTaiOk returns a tuple with the Tai field value
// and a boolean to check if the value has been set.
func (o *NrLocation) GetTaiOk() (*Tai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tai, true
}

// SetTai sets field value
func (o *NrLocation) SetTai(v Tai) {
	o.Tai = v
}

// GetNcgi returns the Ncgi field value
func (o *NrLocation) GetNcgi() Ncgi {
	if o == nil {
		var ret Ncgi
		return ret
	}

	return o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value
// and a boolean to check if the value has been set.
func (o *NrLocation) GetNcgiOk() (*Ncgi, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ncgi, true
}

// SetNcgi sets field value
func (o *NrLocation) SetNcgi(v Ncgi) {
	o.Ncgi = v
}

// GetIgnoreNcgi returns the IgnoreNcgi field value if set, zero value otherwise.
func (o *NrLocation) GetIgnoreNcgi() bool {
	if o == nil || isNil(o.IgnoreNcgi) {
		var ret bool
		return ret
	}
	return *o.IgnoreNcgi
}

// GetIgnoreNcgiOk returns a tuple with the IgnoreNcgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetIgnoreNcgiOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreNcgi) {
    return nil, false
	}
	return o.IgnoreNcgi, true
}

// HasIgnoreNcgi returns a boolean if a field has been set.
func (o *NrLocation) HasIgnoreNcgi() bool {
	if o != nil && !isNil(o.IgnoreNcgi) {
		return true
	}

	return false
}

// SetIgnoreNcgi gets a reference to the given bool and assigns it to the IgnoreNcgi field.
func (o *NrLocation) SetIgnoreNcgi(v bool) {
	o.IgnoreNcgi = &v
}

// GetAgeOfLocationInformation returns the AgeOfLocationInformation field value if set, zero value otherwise.
func (o *NrLocation) GetAgeOfLocationInformation() int32 {
	if o == nil || isNil(o.AgeOfLocationInformation) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationInformation
}

// GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetAgeOfLocationInformationOk() (*int32, bool) {
	if o == nil || isNil(o.AgeOfLocationInformation) {
    return nil, false
	}
	return o.AgeOfLocationInformation, true
}

// HasAgeOfLocationInformation returns a boolean if a field has been set.
func (o *NrLocation) HasAgeOfLocationInformation() bool {
	if o != nil && !isNil(o.AgeOfLocationInformation) {
		return true
	}

	return false
}

// SetAgeOfLocationInformation gets a reference to the given int32 and assigns it to the AgeOfLocationInformation field.
func (o *NrLocation) SetAgeOfLocationInformation(v int32) {
	o.AgeOfLocationInformation = &v
}

// GetUeLocationTimestamp returns the UeLocationTimestamp field value if set, zero value otherwise.
func (o *NrLocation) GetUeLocationTimestamp() time.Time {
	if o == nil || isNil(o.UeLocationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.UeLocationTimestamp
}

// GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetUeLocationTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.UeLocationTimestamp) {
    return nil, false
	}
	return o.UeLocationTimestamp, true
}

// HasUeLocationTimestamp returns a boolean if a field has been set.
func (o *NrLocation) HasUeLocationTimestamp() bool {
	if o != nil && !isNil(o.UeLocationTimestamp) {
		return true
	}

	return false
}

// SetUeLocationTimestamp gets a reference to the given time.Time and assigns it to the UeLocationTimestamp field.
func (o *NrLocation) SetUeLocationTimestamp(v time.Time) {
	o.UeLocationTimestamp = &v
}

// GetGeographicalInformation returns the GeographicalInformation field value if set, zero value otherwise.
func (o *NrLocation) GetGeographicalInformation() string {
	if o == nil || isNil(o.GeographicalInformation) {
		var ret string
		return ret
	}
	return *o.GeographicalInformation
}

// GetGeographicalInformationOk returns a tuple with the GeographicalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetGeographicalInformationOk() (*string, bool) {
	if o == nil || isNil(o.GeographicalInformation) {
    return nil, false
	}
	return o.GeographicalInformation, true
}

// HasGeographicalInformation returns a boolean if a field has been set.
func (o *NrLocation) HasGeographicalInformation() bool {
	if o != nil && !isNil(o.GeographicalInformation) {
		return true
	}

	return false
}

// SetGeographicalInformation gets a reference to the given string and assigns it to the GeographicalInformation field.
func (o *NrLocation) SetGeographicalInformation(v string) {
	o.GeographicalInformation = &v
}

// GetGeodeticInformation returns the GeodeticInformation field value if set, zero value otherwise.
func (o *NrLocation) GetGeodeticInformation() string {
	if o == nil || isNil(o.GeodeticInformation) {
		var ret string
		return ret
	}
	return *o.GeodeticInformation
}

// GetGeodeticInformationOk returns a tuple with the GeodeticInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetGeodeticInformationOk() (*string, bool) {
	if o == nil || isNil(o.GeodeticInformation) {
    return nil, false
	}
	return o.GeodeticInformation, true
}

// HasGeodeticInformation returns a boolean if a field has been set.
func (o *NrLocation) HasGeodeticInformation() bool {
	if o != nil && !isNil(o.GeodeticInformation) {
		return true
	}

	return false
}

// SetGeodeticInformation gets a reference to the given string and assigns it to the GeodeticInformation field.
func (o *NrLocation) SetGeodeticInformation(v string) {
	o.GeodeticInformation = &v
}

// GetGlobalGnbId returns the GlobalGnbId field value if set, zero value otherwise.
func (o *NrLocation) GetGlobalGnbId() GlobalRanNodeId {
	if o == nil || isNil(o.GlobalGnbId) {
		var ret GlobalRanNodeId
		return ret
	}
	return *o.GlobalGnbId
}

// GetGlobalGnbIdOk returns a tuple with the GlobalGnbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrLocation) GetGlobalGnbIdOk() (*GlobalRanNodeId, bool) {
	if o == nil || isNil(o.GlobalGnbId) {
    return nil, false
	}
	return o.GlobalGnbId, true
}

// HasGlobalGnbId returns a boolean if a field has been set.
func (o *NrLocation) HasGlobalGnbId() bool {
	if o != nil && !isNil(o.GlobalGnbId) {
		return true
	}

	return false
}

// SetGlobalGnbId gets a reference to the given GlobalRanNodeId and assigns it to the GlobalGnbId field.
func (o *NrLocation) SetGlobalGnbId(v GlobalRanNodeId) {
	o.GlobalGnbId = &v
}

func (o NrLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tai"] = o.Tai
	}
	if true {
		toSerialize["ncgi"] = o.Ncgi
	}
	if !isNil(o.IgnoreNcgi) {
		toSerialize["ignoreNcgi"] = o.IgnoreNcgi
	}
	if !isNil(o.AgeOfLocationInformation) {
		toSerialize["ageOfLocationInformation"] = o.AgeOfLocationInformation
	}
	if !isNil(o.UeLocationTimestamp) {
		toSerialize["ueLocationTimestamp"] = o.UeLocationTimestamp
	}
	if !isNil(o.GeographicalInformation) {
		toSerialize["geographicalInformation"] = o.GeographicalInformation
	}
	if !isNil(o.GeodeticInformation) {
		toSerialize["geodeticInformation"] = o.GeodeticInformation
	}
	if !isNil(o.GlobalGnbId) {
		toSerialize["globalGnbId"] = o.GlobalGnbId
	}
	return json.Marshal(toSerialize)
}

type NullableNrLocation struct {
	value *NrLocation
	isSet bool
}

func (v NullableNrLocation) Get() *NrLocation {
	return v.value
}

func (v *NullableNrLocation) Set(val *NrLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableNrLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableNrLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrLocation(val *NrLocation) *NullableNrLocation {
	return &NullableNrLocation{value: val, isSet: true}
}

func (v NullableNrLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


