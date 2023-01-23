/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"time"
)

// AlarmListSingleAllOfAttributes struct for AlarmListSingleAllOfAttributes
type AlarmListSingleAllOfAttributes struct {
	AdministrativeState *AdministrativeState `json:"administrativeState,omitempty"`
	OperationalState *OperationalState `json:"operationalState,omitempty"`
	NumOfAlarmRecords *int32 `json:"numOfAlarmRecords,omitempty"`
	LastModification *time.Time `json:"lastModification,omitempty"`
	// This resource represents a map of alarm records. The alarmIds are used as keys in the map.
	AlarmRecords *map[string]AlarmRecord `json:"alarmRecords,omitempty"`
}

// NewAlarmListSingleAllOfAttributes instantiates a new AlarmListSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmListSingleAllOfAttributes() *AlarmListSingleAllOfAttributes {
	this := AlarmListSingleAllOfAttributes{}
	return &this
}

// NewAlarmListSingleAllOfAttributesWithDefaults instantiates a new AlarmListSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmListSingleAllOfAttributesWithDefaults() *AlarmListSingleAllOfAttributes {
	this := AlarmListSingleAllOfAttributes{}
	return &this
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *AlarmListSingleAllOfAttributes) GetAdministrativeState() AdministrativeState {
	if o == nil || isNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || isNil(o.AdministrativeState) {
    return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *AlarmListSingleAllOfAttributes) HasAdministrativeState() bool {
	if o != nil && !isNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *AlarmListSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *AlarmListSingleAllOfAttributes) GetOperationalState() OperationalState {
	if o == nil || isNil(o.OperationalState) {
		var ret OperationalState
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool) {
	if o == nil || isNil(o.OperationalState) {
    return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *AlarmListSingleAllOfAttributes) HasOperationalState() bool {
	if o != nil && !isNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given OperationalState and assigns it to the OperationalState field.
func (o *AlarmListSingleAllOfAttributes) SetOperationalState(v OperationalState) {
	o.OperationalState = &v
}

// GetNumOfAlarmRecords returns the NumOfAlarmRecords field value if set, zero value otherwise.
func (o *AlarmListSingleAllOfAttributes) GetNumOfAlarmRecords() int32 {
	if o == nil || isNil(o.NumOfAlarmRecords) {
		var ret int32
		return ret
	}
	return *o.NumOfAlarmRecords
}

// GetNumOfAlarmRecordsOk returns a tuple with the NumOfAlarmRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListSingleAllOfAttributes) GetNumOfAlarmRecordsOk() (*int32, bool) {
	if o == nil || isNil(o.NumOfAlarmRecords) {
    return nil, false
	}
	return o.NumOfAlarmRecords, true
}

// HasNumOfAlarmRecords returns a boolean if a field has been set.
func (o *AlarmListSingleAllOfAttributes) HasNumOfAlarmRecords() bool {
	if o != nil && !isNil(o.NumOfAlarmRecords) {
		return true
	}

	return false
}

// SetNumOfAlarmRecords gets a reference to the given int32 and assigns it to the NumOfAlarmRecords field.
func (o *AlarmListSingleAllOfAttributes) SetNumOfAlarmRecords(v int32) {
	o.NumOfAlarmRecords = &v
}

// GetLastModification returns the LastModification field value if set, zero value otherwise.
func (o *AlarmListSingleAllOfAttributes) GetLastModification() time.Time {
	if o == nil || isNil(o.LastModification) {
		var ret time.Time
		return ret
	}
	return *o.LastModification
}

// GetLastModificationOk returns a tuple with the LastModification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListSingleAllOfAttributes) GetLastModificationOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModification) {
    return nil, false
	}
	return o.LastModification, true
}

// HasLastModification returns a boolean if a field has been set.
func (o *AlarmListSingleAllOfAttributes) HasLastModification() bool {
	if o != nil && !isNil(o.LastModification) {
		return true
	}

	return false
}

// SetLastModification gets a reference to the given time.Time and assigns it to the LastModification field.
func (o *AlarmListSingleAllOfAttributes) SetLastModification(v time.Time) {
	o.LastModification = &v
}

// GetAlarmRecords returns the AlarmRecords field value if set, zero value otherwise.
func (o *AlarmListSingleAllOfAttributes) GetAlarmRecords() map[string]AlarmRecord {
	if o == nil || isNil(o.AlarmRecords) {
		var ret map[string]AlarmRecord
		return ret
	}
	return *o.AlarmRecords
}

// GetAlarmRecordsOk returns a tuple with the AlarmRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListSingleAllOfAttributes) GetAlarmRecordsOk() (*map[string]AlarmRecord, bool) {
	if o == nil || isNil(o.AlarmRecords) {
    return nil, false
	}
	return o.AlarmRecords, true
}

// HasAlarmRecords returns a boolean if a field has been set.
func (o *AlarmListSingleAllOfAttributes) HasAlarmRecords() bool {
	if o != nil && !isNil(o.AlarmRecords) {
		return true
	}

	return false
}

// SetAlarmRecords gets a reference to the given map[string]AlarmRecord and assigns it to the AlarmRecords field.
func (o *AlarmListSingleAllOfAttributes) SetAlarmRecords(v map[string]AlarmRecord) {
	o.AlarmRecords = &v
}

func (o AlarmListSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !isNil(o.OperationalState) {
		toSerialize["operationalState"] = o.OperationalState
	}
	if !isNil(o.NumOfAlarmRecords) {
		toSerialize["numOfAlarmRecords"] = o.NumOfAlarmRecords
	}
	if !isNil(o.LastModification) {
		toSerialize["lastModification"] = o.LastModification
	}
	if !isNil(o.AlarmRecords) {
		toSerialize["alarmRecords"] = o.AlarmRecords
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmListSingleAllOfAttributes struct {
	value *AlarmListSingleAllOfAttributes
	isSet bool
}

func (v NullableAlarmListSingleAllOfAttributes) Get() *AlarmListSingleAllOfAttributes {
	return v.value
}

func (v *NullableAlarmListSingleAllOfAttributes) Set(val *AlarmListSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmListSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmListSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmListSingleAllOfAttributes(val *AlarmListSingleAllOfAttributes) *NullableAlarmListSingleAllOfAttributes {
	return &NullableAlarmListSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableAlarmListSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmListSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


