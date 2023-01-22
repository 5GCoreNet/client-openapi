/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
	"time"
)

// SmPolicyDeleteData Contains the parameters to be sent to the PCF when an individual SM policy is deleted.
type SmPolicyDeleteData struct {
	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UeTimeZone *string `json:"ueTimeZone,omitempty"`
	ServingNetwork *PlmnIdNid `json:"servingNetwork,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	UserLocationInfoTime *time.Time `json:"userLocationInfoTime,omitempty"`
	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`
	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty"`
	PduSessRelCause *PduSessionRelCause `json:"pduSessRelCause,omitempty"`
	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`
}

// NewSmPolicyDeleteData instantiates a new SmPolicyDeleteData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmPolicyDeleteData() *SmPolicyDeleteData {
	this := SmPolicyDeleteData{}
	return &this
}

// NewSmPolicyDeleteDataWithDefaults instantiates a new SmPolicyDeleteData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmPolicyDeleteDataWithDefaults() *SmPolicyDeleteData {
	this := SmPolicyDeleteData{}
	return &this
}

// GetUserLocationInfo returns the UserLocationInfo field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetUserLocationInfo() UserLocation {
	if o == nil || isNil(o.UserLocationInfo) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationInfo
}

// GetUserLocationInfoOk returns a tuple with the UserLocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetUserLocationInfoOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UserLocationInfo) {
    return nil, false
	}
	return o.UserLocationInfo, true
}

// HasUserLocationInfo returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasUserLocationInfo() bool {
	if o != nil && !isNil(o.UserLocationInfo) {
		return true
	}

	return false
}

// SetUserLocationInfo gets a reference to the given UserLocation and assigns it to the UserLocationInfo field.
func (o *SmPolicyDeleteData) SetUserLocationInfo(v UserLocation) {
	o.UserLocationInfo = &v
}

// GetUeTimeZone returns the UeTimeZone field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetUeTimeZone() string {
	if o == nil || isNil(o.UeTimeZone) {
		var ret string
		return ret
	}
	return *o.UeTimeZone
}

// GetUeTimeZoneOk returns a tuple with the UeTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetUeTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.UeTimeZone) {
    return nil, false
	}
	return o.UeTimeZone, true
}

// HasUeTimeZone returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasUeTimeZone() bool {
	if o != nil && !isNil(o.UeTimeZone) {
		return true
	}

	return false
}

// SetUeTimeZone gets a reference to the given string and assigns it to the UeTimeZone field.
func (o *SmPolicyDeleteData) SetUeTimeZone(v string) {
	o.UeTimeZone = &v
}

// GetServingNetwork returns the ServingNetwork field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetServingNetwork() PlmnIdNid {
	if o == nil || isNil(o.ServingNetwork) {
		var ret PlmnIdNid
		return ret
	}
	return *o.ServingNetwork
}

// GetServingNetworkOk returns a tuple with the ServingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetServingNetworkOk() (*PlmnIdNid, bool) {
	if o == nil || isNil(o.ServingNetwork) {
    return nil, false
	}
	return o.ServingNetwork, true
}

// HasServingNetwork returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasServingNetwork() bool {
	if o != nil && !isNil(o.ServingNetwork) {
		return true
	}

	return false
}

// SetServingNetwork gets a reference to the given PlmnIdNid and assigns it to the ServingNetwork field.
func (o *SmPolicyDeleteData) SetServingNetwork(v PlmnIdNid) {
	o.ServingNetwork = &v
}

// GetUserLocationInfoTime returns the UserLocationInfoTime field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetUserLocationInfoTime() time.Time {
	if o == nil || isNil(o.UserLocationInfoTime) {
		var ret time.Time
		return ret
	}
	return *o.UserLocationInfoTime
}

// GetUserLocationInfoTimeOk returns a tuple with the UserLocationInfoTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetUserLocationInfoTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.UserLocationInfoTime) {
    return nil, false
	}
	return o.UserLocationInfoTime, true
}

// HasUserLocationInfoTime returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasUserLocationInfoTime() bool {
	if o != nil && !isNil(o.UserLocationInfoTime) {
		return true
	}

	return false
}

// SetUserLocationInfoTime gets a reference to the given time.Time and assigns it to the UserLocationInfoTime field.
func (o *SmPolicyDeleteData) SetUserLocationInfoTime(v time.Time) {
	o.UserLocationInfoTime = &v
}

// GetRanNasRelCauses returns the RanNasRelCauses field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetRanNasRelCauses() []RanNasRelCause {
	if o == nil || isNil(o.RanNasRelCauses) {
		var ret []RanNasRelCause
		return ret
	}
	return o.RanNasRelCauses
}

// GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetRanNasRelCausesOk() ([]RanNasRelCause, bool) {
	if o == nil || isNil(o.RanNasRelCauses) {
    return nil, false
	}
	return o.RanNasRelCauses, true
}

// HasRanNasRelCauses returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasRanNasRelCauses() bool {
	if o != nil && !isNil(o.RanNasRelCauses) {
		return true
	}

	return false
}

// SetRanNasRelCauses gets a reference to the given []RanNasRelCause and assigns it to the RanNasRelCauses field.
func (o *SmPolicyDeleteData) SetRanNasRelCauses(v []RanNasRelCause) {
	o.RanNasRelCauses = v
}

// GetAccuUsageReports returns the AccuUsageReports field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetAccuUsageReports() []AccuUsageReport {
	if o == nil || isNil(o.AccuUsageReports) {
		var ret []AccuUsageReport
		return ret
	}
	return o.AccuUsageReports
}

// GetAccuUsageReportsOk returns a tuple with the AccuUsageReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetAccuUsageReportsOk() ([]AccuUsageReport, bool) {
	if o == nil || isNil(o.AccuUsageReports) {
    return nil, false
	}
	return o.AccuUsageReports, true
}

// HasAccuUsageReports returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasAccuUsageReports() bool {
	if o != nil && !isNil(o.AccuUsageReports) {
		return true
	}

	return false
}

// SetAccuUsageReports gets a reference to the given []AccuUsageReport and assigns it to the AccuUsageReports field.
func (o *SmPolicyDeleteData) SetAccuUsageReports(v []AccuUsageReport) {
	o.AccuUsageReports = v
}

// GetPduSessRelCause returns the PduSessRelCause field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetPduSessRelCause() PduSessionRelCause {
	if o == nil || isNil(o.PduSessRelCause) {
		var ret PduSessionRelCause
		return ret
	}
	return *o.PduSessRelCause
}

// GetPduSessRelCauseOk returns a tuple with the PduSessRelCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetPduSessRelCauseOk() (*PduSessionRelCause, bool) {
	if o == nil || isNil(o.PduSessRelCause) {
    return nil, false
	}
	return o.PduSessRelCause, true
}

// HasPduSessRelCause returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasPduSessRelCause() bool {
	if o != nil && !isNil(o.PduSessRelCause) {
		return true
	}

	return false
}

// SetPduSessRelCause gets a reference to the given PduSessionRelCause and assigns it to the PduSessRelCause field.
func (o *SmPolicyDeleteData) SetPduSessRelCause(v PduSessionRelCause) {
	o.PduSessRelCause = &v
}

// GetQosMonReports returns the QosMonReports field value if set, zero value otherwise.
func (o *SmPolicyDeleteData) GetQosMonReports() []QosMonitoringReport {
	if o == nil || isNil(o.QosMonReports) {
		var ret []QosMonitoringReport
		return ret
	}
	return o.QosMonReports
}

// GetQosMonReportsOk returns a tuple with the QosMonReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmPolicyDeleteData) GetQosMonReportsOk() ([]QosMonitoringReport, bool) {
	if o == nil || isNil(o.QosMonReports) {
    return nil, false
	}
	return o.QosMonReports, true
}

// HasQosMonReports returns a boolean if a field has been set.
func (o *SmPolicyDeleteData) HasQosMonReports() bool {
	if o != nil && !isNil(o.QosMonReports) {
		return true
	}

	return false
}

// SetQosMonReports gets a reference to the given []QosMonitoringReport and assigns it to the QosMonReports field.
func (o *SmPolicyDeleteData) SetQosMonReports(v []QosMonitoringReport) {
	o.QosMonReports = v
}

func (o SmPolicyDeleteData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserLocationInfo) {
		toSerialize["userLocationInfo"] = o.UserLocationInfo
	}
	if !isNil(o.UeTimeZone) {
		toSerialize["ueTimeZone"] = o.UeTimeZone
	}
	if !isNil(o.ServingNetwork) {
		toSerialize["servingNetwork"] = o.ServingNetwork
	}
	if !isNil(o.UserLocationInfoTime) {
		toSerialize["userLocationInfoTime"] = o.UserLocationInfoTime
	}
	if !isNil(o.RanNasRelCauses) {
		toSerialize["ranNasRelCauses"] = o.RanNasRelCauses
	}
	if !isNil(o.AccuUsageReports) {
		toSerialize["accuUsageReports"] = o.AccuUsageReports
	}
	if !isNil(o.PduSessRelCause) {
		toSerialize["pduSessRelCause"] = o.PduSessRelCause
	}
	if !isNil(o.QosMonReports) {
		toSerialize["qosMonReports"] = o.QosMonReports
	}
	return json.Marshal(toSerialize)
}

type NullableSmPolicyDeleteData struct {
	value *SmPolicyDeleteData
	isSet bool
}

func (v NullableSmPolicyDeleteData) Get() *SmPolicyDeleteData {
	return v.value
}

func (v *NullableSmPolicyDeleteData) Set(val *SmPolicyDeleteData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmPolicyDeleteData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmPolicyDeleteData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmPolicyDeleteData(val *SmPolicyDeleteData) *NullableSmPolicyDeleteData {
	return &NullableSmPolicyDeleteData{value: val, isSet: true}
}

func (v NullableSmPolicyDeleteData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmPolicyDeleteData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


