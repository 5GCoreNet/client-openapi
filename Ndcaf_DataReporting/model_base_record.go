/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
	"time"
)

// BaseRecord struct for BaseRecord
type BaseRecord struct {
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
}

// NewBaseRecord instantiates a new BaseRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseRecord(timestamp time.Time) *BaseRecord {
	this := BaseRecord{}
	this.Timestamp = timestamp
	return &this
}

// NewBaseRecordWithDefaults instantiates a new BaseRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseRecordWithDefaults() *BaseRecord {
	this := BaseRecord{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *BaseRecord) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BaseRecord) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BaseRecord) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o BaseRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableBaseRecord struct {
	value *BaseRecord
	isSet bool
}

func (v NullableBaseRecord) Get() *BaseRecord {
	return v.value
}

func (v *NullableBaseRecord) Set(val *BaseRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseRecord(val *BaseRecord) *NullableBaseRecord {
	return &NullableBaseRecord{value: val, isSet: true}
}

func (v NullableBaseRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


