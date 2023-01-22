/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"time"
)

// AccessAndMobilityData Represents Access and Mobility data for a UE.
type AccessAndMobilityData struct {
	Location *UserLocation `json:"location,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LocationTs *time.Time `json:"locationTs,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	TimeZone *string `json:"timeZone,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeZoneTs *time.Time `json:"timeZoneTs,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
	RegStates []RmInfo `json:"regStates,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegStatesTs *time.Time `json:"regStatesTs,omitempty"`
	ConnStates []CmInfo `json:"connStates,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ConnStatesTs *time.Time `json:"connStatesTs,omitempty"`
	ReachabilityStatus *UeReachability `json:"reachabilityStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ReachabilityStatusTs *time.Time `json:"reachabilityStatusTs,omitempty"`
	SmsOverNasStatus *SmsSupport `json:"smsOverNasStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SmsOverNasStatusTs *time.Time `json:"smsOverNasStatusTs,omitempty"`
	// True  The serving PLMN of the UE is different from the HPLMN of the UE; False The serving PLMN of the UE is the HPLMN of the UE. 
	RoamingStatus *bool `json:"roamingStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RoamingStatusTs *time.Time `json:"roamingStatusTs,omitempty"`
	CurrentPlmn *PlmnId1 `json:"currentPlmn,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	CurrentPlmnTs *time.Time `json:"currentPlmnTs,omitempty"`
	RatType []RatType `json:"ratType,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RatTypesTs *time.Time `json:"ratTypesTs,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewAccessAndMobilityData instantiates a new AccessAndMobilityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessAndMobilityData() *AccessAndMobilityData {
	this := AccessAndMobilityData{}
	return &this
}

// NewAccessAndMobilityDataWithDefaults instantiates a new AccessAndMobilityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessAndMobilityDataWithDefaults() *AccessAndMobilityData {
	this := AccessAndMobilityData{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetLocation() UserLocation {
	if o == nil || isNil(o.Location) {
		var ret UserLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetLocationOk() (*UserLocation, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given UserLocation and assigns it to the Location field.
func (o *AccessAndMobilityData) SetLocation(v UserLocation) {
	o.Location = &v
}

// GetLocationTs returns the LocationTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetLocationTs() time.Time {
	if o == nil || isNil(o.LocationTs) {
		var ret time.Time
		return ret
	}
	return *o.LocationTs
}

// GetLocationTsOk returns a tuple with the LocationTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetLocationTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.LocationTs) {
    return nil, false
	}
	return o.LocationTs, true
}

// HasLocationTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasLocationTs() bool {
	if o != nil && !isNil(o.LocationTs) {
		return true
	}

	return false
}

// SetLocationTs gets a reference to the given time.Time and assigns it to the LocationTs field.
func (o *AccessAndMobilityData) SetLocationTs(v time.Time) {
	o.LocationTs = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AccessAndMobilityData) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTimeZoneTs returns the TimeZoneTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetTimeZoneTs() time.Time {
	if o == nil || isNil(o.TimeZoneTs) {
		var ret time.Time
		return ret
	}
	return *o.TimeZoneTs
}

// GetTimeZoneTsOk returns a tuple with the TimeZoneTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetTimeZoneTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeZoneTs) {
    return nil, false
	}
	return o.TimeZoneTs, true
}

// HasTimeZoneTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasTimeZoneTs() bool {
	if o != nil && !isNil(o.TimeZoneTs) {
		return true
	}

	return false
}

// SetTimeZoneTs gets a reference to the given time.Time and assigns it to the TimeZoneTs field.
func (o *AccessAndMobilityData) SetTimeZoneTs(v time.Time) {
	o.TimeZoneTs = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetAccessType() AccessType {
	if o == nil || isNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || isNil(o.AccessType) {
    return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *AccessAndMobilityData) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetRegStates returns the RegStates field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetRegStates() []RmInfo {
	if o == nil || isNil(o.RegStates) {
		var ret []RmInfo
		return ret
	}
	return o.RegStates
}

// GetRegStatesOk returns a tuple with the RegStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetRegStatesOk() ([]RmInfo, bool) {
	if o == nil || isNil(o.RegStates) {
    return nil, false
	}
	return o.RegStates, true
}

// HasRegStates returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasRegStates() bool {
	if o != nil && !isNil(o.RegStates) {
		return true
	}

	return false
}

// SetRegStates gets a reference to the given []RmInfo and assigns it to the RegStates field.
func (o *AccessAndMobilityData) SetRegStates(v []RmInfo) {
	o.RegStates = v
}

// GetRegStatesTs returns the RegStatesTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetRegStatesTs() time.Time {
	if o == nil || isNil(o.RegStatesTs) {
		var ret time.Time
		return ret
	}
	return *o.RegStatesTs
}

// GetRegStatesTsOk returns a tuple with the RegStatesTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetRegStatesTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.RegStatesTs) {
    return nil, false
	}
	return o.RegStatesTs, true
}

// HasRegStatesTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasRegStatesTs() bool {
	if o != nil && !isNil(o.RegStatesTs) {
		return true
	}

	return false
}

// SetRegStatesTs gets a reference to the given time.Time and assigns it to the RegStatesTs field.
func (o *AccessAndMobilityData) SetRegStatesTs(v time.Time) {
	o.RegStatesTs = &v
}

// GetConnStates returns the ConnStates field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetConnStates() []CmInfo {
	if o == nil || isNil(o.ConnStates) {
		var ret []CmInfo
		return ret
	}
	return o.ConnStates
}

// GetConnStatesOk returns a tuple with the ConnStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetConnStatesOk() ([]CmInfo, bool) {
	if o == nil || isNil(o.ConnStates) {
    return nil, false
	}
	return o.ConnStates, true
}

// HasConnStates returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasConnStates() bool {
	if o != nil && !isNil(o.ConnStates) {
		return true
	}

	return false
}

// SetConnStates gets a reference to the given []CmInfo and assigns it to the ConnStates field.
func (o *AccessAndMobilityData) SetConnStates(v []CmInfo) {
	o.ConnStates = v
}

// GetConnStatesTs returns the ConnStatesTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetConnStatesTs() time.Time {
	if o == nil || isNil(o.ConnStatesTs) {
		var ret time.Time
		return ret
	}
	return *o.ConnStatesTs
}

// GetConnStatesTsOk returns a tuple with the ConnStatesTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetConnStatesTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.ConnStatesTs) {
    return nil, false
	}
	return o.ConnStatesTs, true
}

// HasConnStatesTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasConnStatesTs() bool {
	if o != nil && !isNil(o.ConnStatesTs) {
		return true
	}

	return false
}

// SetConnStatesTs gets a reference to the given time.Time and assigns it to the ConnStatesTs field.
func (o *AccessAndMobilityData) SetConnStatesTs(v time.Time) {
	o.ConnStatesTs = &v
}

// GetReachabilityStatus returns the ReachabilityStatus field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetReachabilityStatus() UeReachability {
	if o == nil || isNil(o.ReachabilityStatus) {
		var ret UeReachability
		return ret
	}
	return *o.ReachabilityStatus
}

// GetReachabilityStatusOk returns a tuple with the ReachabilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetReachabilityStatusOk() (*UeReachability, bool) {
	if o == nil || isNil(o.ReachabilityStatus) {
    return nil, false
	}
	return o.ReachabilityStatus, true
}

// HasReachabilityStatus returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasReachabilityStatus() bool {
	if o != nil && !isNil(o.ReachabilityStatus) {
		return true
	}

	return false
}

// SetReachabilityStatus gets a reference to the given UeReachability and assigns it to the ReachabilityStatus field.
func (o *AccessAndMobilityData) SetReachabilityStatus(v UeReachability) {
	o.ReachabilityStatus = &v
}

// GetReachabilityStatusTs returns the ReachabilityStatusTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetReachabilityStatusTs() time.Time {
	if o == nil || isNil(o.ReachabilityStatusTs) {
		var ret time.Time
		return ret
	}
	return *o.ReachabilityStatusTs
}

// GetReachabilityStatusTsOk returns a tuple with the ReachabilityStatusTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetReachabilityStatusTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.ReachabilityStatusTs) {
    return nil, false
	}
	return o.ReachabilityStatusTs, true
}

// HasReachabilityStatusTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasReachabilityStatusTs() bool {
	if o != nil && !isNil(o.ReachabilityStatusTs) {
		return true
	}

	return false
}

// SetReachabilityStatusTs gets a reference to the given time.Time and assigns it to the ReachabilityStatusTs field.
func (o *AccessAndMobilityData) SetReachabilityStatusTs(v time.Time) {
	o.ReachabilityStatusTs = &v
}

// GetSmsOverNasStatus returns the SmsOverNasStatus field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetSmsOverNasStatus() SmsSupport {
	if o == nil || isNil(o.SmsOverNasStatus) {
		var ret SmsSupport
		return ret
	}
	return *o.SmsOverNasStatus
}

// GetSmsOverNasStatusOk returns a tuple with the SmsOverNasStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetSmsOverNasStatusOk() (*SmsSupport, bool) {
	if o == nil || isNil(o.SmsOverNasStatus) {
    return nil, false
	}
	return o.SmsOverNasStatus, true
}

// HasSmsOverNasStatus returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasSmsOverNasStatus() bool {
	if o != nil && !isNil(o.SmsOverNasStatus) {
		return true
	}

	return false
}

// SetSmsOverNasStatus gets a reference to the given SmsSupport and assigns it to the SmsOverNasStatus field.
func (o *AccessAndMobilityData) SetSmsOverNasStatus(v SmsSupport) {
	o.SmsOverNasStatus = &v
}

// GetSmsOverNasStatusTs returns the SmsOverNasStatusTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetSmsOverNasStatusTs() time.Time {
	if o == nil || isNil(o.SmsOverNasStatusTs) {
		var ret time.Time
		return ret
	}
	return *o.SmsOverNasStatusTs
}

// GetSmsOverNasStatusTsOk returns a tuple with the SmsOverNasStatusTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetSmsOverNasStatusTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.SmsOverNasStatusTs) {
    return nil, false
	}
	return o.SmsOverNasStatusTs, true
}

// HasSmsOverNasStatusTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasSmsOverNasStatusTs() bool {
	if o != nil && !isNil(o.SmsOverNasStatusTs) {
		return true
	}

	return false
}

// SetSmsOverNasStatusTs gets a reference to the given time.Time and assigns it to the SmsOverNasStatusTs field.
func (o *AccessAndMobilityData) SetSmsOverNasStatusTs(v time.Time) {
	o.SmsOverNasStatusTs = &v
}

// GetRoamingStatus returns the RoamingStatus field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetRoamingStatus() bool {
	if o == nil || isNil(o.RoamingStatus) {
		var ret bool
		return ret
	}
	return *o.RoamingStatus
}

// GetRoamingStatusOk returns a tuple with the RoamingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetRoamingStatusOk() (*bool, bool) {
	if o == nil || isNil(o.RoamingStatus) {
    return nil, false
	}
	return o.RoamingStatus, true
}

// HasRoamingStatus returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasRoamingStatus() bool {
	if o != nil && !isNil(o.RoamingStatus) {
		return true
	}

	return false
}

// SetRoamingStatus gets a reference to the given bool and assigns it to the RoamingStatus field.
func (o *AccessAndMobilityData) SetRoamingStatus(v bool) {
	o.RoamingStatus = &v
}

// GetRoamingStatusTs returns the RoamingStatusTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetRoamingStatusTs() time.Time {
	if o == nil || isNil(o.RoamingStatusTs) {
		var ret time.Time
		return ret
	}
	return *o.RoamingStatusTs
}

// GetRoamingStatusTsOk returns a tuple with the RoamingStatusTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetRoamingStatusTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.RoamingStatusTs) {
    return nil, false
	}
	return o.RoamingStatusTs, true
}

// HasRoamingStatusTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasRoamingStatusTs() bool {
	if o != nil && !isNil(o.RoamingStatusTs) {
		return true
	}

	return false
}

// SetRoamingStatusTs gets a reference to the given time.Time and assigns it to the RoamingStatusTs field.
func (o *AccessAndMobilityData) SetRoamingStatusTs(v time.Time) {
	o.RoamingStatusTs = &v
}

// GetCurrentPlmn returns the CurrentPlmn field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetCurrentPlmn() PlmnId1 {
	if o == nil || isNil(o.CurrentPlmn) {
		var ret PlmnId1
		return ret
	}
	return *o.CurrentPlmn
}

// GetCurrentPlmnOk returns a tuple with the CurrentPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetCurrentPlmnOk() (*PlmnId1, bool) {
	if o == nil || isNil(o.CurrentPlmn) {
    return nil, false
	}
	return o.CurrentPlmn, true
}

// HasCurrentPlmn returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasCurrentPlmn() bool {
	if o != nil && !isNil(o.CurrentPlmn) {
		return true
	}

	return false
}

// SetCurrentPlmn gets a reference to the given PlmnId1 and assigns it to the CurrentPlmn field.
func (o *AccessAndMobilityData) SetCurrentPlmn(v PlmnId1) {
	o.CurrentPlmn = &v
}

// GetCurrentPlmnTs returns the CurrentPlmnTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetCurrentPlmnTs() time.Time {
	if o == nil || isNil(o.CurrentPlmnTs) {
		var ret time.Time
		return ret
	}
	return *o.CurrentPlmnTs
}

// GetCurrentPlmnTsOk returns a tuple with the CurrentPlmnTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetCurrentPlmnTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.CurrentPlmnTs) {
    return nil, false
	}
	return o.CurrentPlmnTs, true
}

// HasCurrentPlmnTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasCurrentPlmnTs() bool {
	if o != nil && !isNil(o.CurrentPlmnTs) {
		return true
	}

	return false
}

// SetCurrentPlmnTs gets a reference to the given time.Time and assigns it to the CurrentPlmnTs field.
func (o *AccessAndMobilityData) SetCurrentPlmnTs(v time.Time) {
	o.CurrentPlmnTs = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetRatType() []RatType {
	if o == nil || isNil(o.RatType) {
		var ret []RatType
		return ret
	}
	return o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetRatTypeOk() ([]RatType, bool) {
	if o == nil || isNil(o.RatType) {
    return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasRatType() bool {
	if o != nil && !isNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given []RatType and assigns it to the RatType field.
func (o *AccessAndMobilityData) SetRatType(v []RatType) {
	o.RatType = v
}

// GetRatTypesTs returns the RatTypesTs field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetRatTypesTs() time.Time {
	if o == nil || isNil(o.RatTypesTs) {
		var ret time.Time
		return ret
	}
	return *o.RatTypesTs
}

// GetRatTypesTsOk returns a tuple with the RatTypesTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetRatTypesTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.RatTypesTs) {
    return nil, false
	}
	return o.RatTypesTs, true
}

// HasRatTypesTs returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasRatTypesTs() bool {
	if o != nil && !isNil(o.RatTypesTs) {
		return true
	}

	return false
}

// SetRatTypesTs gets a reference to the given time.Time and assigns it to the RatTypesTs field.
func (o *AccessAndMobilityData) SetRatTypesTs(v time.Time) {
	o.RatTypesTs = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *AccessAndMobilityData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *AccessAndMobilityData) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAndMobilityData) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *AccessAndMobilityData) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *AccessAndMobilityData) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o AccessAndMobilityData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.LocationTs) {
		toSerialize["locationTs"] = o.LocationTs
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !isNil(o.TimeZoneTs) {
		toSerialize["timeZoneTs"] = o.TimeZoneTs
	}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.RegStates) {
		toSerialize["regStates"] = o.RegStates
	}
	if !isNil(o.RegStatesTs) {
		toSerialize["regStatesTs"] = o.RegStatesTs
	}
	if !isNil(o.ConnStates) {
		toSerialize["connStates"] = o.ConnStates
	}
	if !isNil(o.ConnStatesTs) {
		toSerialize["connStatesTs"] = o.ConnStatesTs
	}
	if !isNil(o.ReachabilityStatus) {
		toSerialize["reachabilityStatus"] = o.ReachabilityStatus
	}
	if !isNil(o.ReachabilityStatusTs) {
		toSerialize["reachabilityStatusTs"] = o.ReachabilityStatusTs
	}
	if !isNil(o.SmsOverNasStatus) {
		toSerialize["smsOverNasStatus"] = o.SmsOverNasStatus
	}
	if !isNil(o.SmsOverNasStatusTs) {
		toSerialize["smsOverNasStatusTs"] = o.SmsOverNasStatusTs
	}
	if !isNil(o.RoamingStatus) {
		toSerialize["roamingStatus"] = o.RoamingStatus
	}
	if !isNil(o.RoamingStatusTs) {
		toSerialize["roamingStatusTs"] = o.RoamingStatusTs
	}
	if !isNil(o.CurrentPlmn) {
		toSerialize["currentPlmn"] = o.CurrentPlmn
	}
	if !isNil(o.CurrentPlmnTs) {
		toSerialize["currentPlmnTs"] = o.CurrentPlmnTs
	}
	if !isNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !isNil(o.RatTypesTs) {
		toSerialize["ratTypesTs"] = o.RatTypesTs
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return json.Marshal(toSerialize)
}

type NullableAccessAndMobilityData struct {
	value *AccessAndMobilityData
	isSet bool
}

func (v NullableAccessAndMobilityData) Get() *AccessAndMobilityData {
	return v.value
}

func (v *NullableAccessAndMobilityData) Set(val *AccessAndMobilityData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessAndMobilityData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessAndMobilityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessAndMobilityData(val *AccessAndMobilityData) *NullableAccessAndMobilityData {
	return &NullableAccessAndMobilityData{value: val, isSet: true}
}

func (v NullableAccessAndMobilityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessAndMobilityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


