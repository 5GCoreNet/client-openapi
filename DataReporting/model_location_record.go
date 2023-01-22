/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
)

// LocationRecord struct for LocationRecord
type LocationRecord struct {
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
	Location LocationData `json:"location"`
}

// NewLocationRecord instantiates a new LocationRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRecord(timestamp time.Time, location LocationData) *LocationRecord {
	this := LocationRecord{}
	this.Timestamp = timestamp
	this.Location = location
	return &this
}

// NewLocationRecordWithDefaults instantiates a new LocationRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRecordWithDefaults() *LocationRecord {
	this := LocationRecord{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *LocationRecord) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *LocationRecord) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *LocationRecord) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetLocation returns the Location field value
func (o *LocationRecord) GetLocation() LocationData {
	if o == nil {
		var ret LocationData
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *LocationRecord) GetLocationOk() (*LocationData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *LocationRecord) SetLocation(v LocationData) {
	o.Location = v
}

func (o LocationRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableLocationRecord struct {
	value *LocationRecord
	isSet bool
}

func (v NullableLocationRecord) Get() *LocationRecord {
	return v.value
}

func (v *NullableLocationRecord) Set(val *LocationRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRecord(val *LocationRecord) *NullableLocationRecord {
	return &NullableLocationRecord{value: val, isSet: true}
}

func (v NullableLocationRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


