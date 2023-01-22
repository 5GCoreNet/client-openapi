/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MdaNrm

import (
	"encoding/json"
	"time"
)

// FileDownloadJobProcessMonitor This data type is the \"ProcessMonitor\" data type with specialisations for usage in the \"FileDownloadJob\".
type FileDownloadJobProcessMonitor struct {
	JobId *string `json:"jobId,omitempty"`
	Status *string `json:"status,omitempty"`
	ProgressPercentage *int32 `json:"progressPercentage,omitempty"`
	ProgressStateInfo *string `json:"progressStateInfo,omitempty"`
	ResultStateInfo *FileDownloadJobProcessMonitorResultStateInfo `json:"resultStateInfo,omitempty"`
	StartTime *time.Time `json:"startTime,omitempty"`
	EndTime *time.Time `json:"endTime,omitempty"`
	Timer *int32 `json:"timer,omitempty"`
}

// NewFileDownloadJobProcessMonitor instantiates a new FileDownloadJobProcessMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDownloadJobProcessMonitor() *FileDownloadJobProcessMonitor {
	this := FileDownloadJobProcessMonitor{}
	return &this
}

// NewFileDownloadJobProcessMonitorWithDefaults instantiates a new FileDownloadJobProcessMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDownloadJobProcessMonitorWithDefaults() *FileDownloadJobProcessMonitor {
	this := FileDownloadJobProcessMonitor{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetJobId() string {
	if o == nil || isNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetJobIdOk() (*string, bool) {
	if o == nil || isNil(o.JobId) {
    return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasJobId() bool {
	if o != nil && !isNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *FileDownloadJobProcessMonitor) SetJobId(v string) {
	o.JobId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FileDownloadJobProcessMonitor) SetStatus(v string) {
	o.Status = &v
}

// GetProgressPercentage returns the ProgressPercentage field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetProgressPercentage() int32 {
	if o == nil || isNil(o.ProgressPercentage) {
		var ret int32
		return ret
	}
	return *o.ProgressPercentage
}

// GetProgressPercentageOk returns a tuple with the ProgressPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetProgressPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.ProgressPercentage) {
    return nil, false
	}
	return o.ProgressPercentage, true
}

// HasProgressPercentage returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasProgressPercentage() bool {
	if o != nil && !isNil(o.ProgressPercentage) {
		return true
	}

	return false
}

// SetProgressPercentage gets a reference to the given int32 and assigns it to the ProgressPercentage field.
func (o *FileDownloadJobProcessMonitor) SetProgressPercentage(v int32) {
	o.ProgressPercentage = &v
}

// GetProgressStateInfo returns the ProgressStateInfo field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetProgressStateInfo() string {
	if o == nil || isNil(o.ProgressStateInfo) {
		var ret string
		return ret
	}
	return *o.ProgressStateInfo
}

// GetProgressStateInfoOk returns a tuple with the ProgressStateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetProgressStateInfoOk() (*string, bool) {
	if o == nil || isNil(o.ProgressStateInfo) {
    return nil, false
	}
	return o.ProgressStateInfo, true
}

// HasProgressStateInfo returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasProgressStateInfo() bool {
	if o != nil && !isNil(o.ProgressStateInfo) {
		return true
	}

	return false
}

// SetProgressStateInfo gets a reference to the given string and assigns it to the ProgressStateInfo field.
func (o *FileDownloadJobProcessMonitor) SetProgressStateInfo(v string) {
	o.ProgressStateInfo = &v
}

// GetResultStateInfo returns the ResultStateInfo field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetResultStateInfo() FileDownloadJobProcessMonitorResultStateInfo {
	if o == nil || isNil(o.ResultStateInfo) {
		var ret FileDownloadJobProcessMonitorResultStateInfo
		return ret
	}
	return *o.ResultStateInfo
}

// GetResultStateInfoOk returns a tuple with the ResultStateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetResultStateInfoOk() (*FileDownloadJobProcessMonitorResultStateInfo, bool) {
	if o == nil || isNil(o.ResultStateInfo) {
    return nil, false
	}
	return o.ResultStateInfo, true
}

// HasResultStateInfo returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasResultStateInfo() bool {
	if o != nil && !isNil(o.ResultStateInfo) {
		return true
	}

	return false
}

// SetResultStateInfo gets a reference to the given FileDownloadJobProcessMonitorResultStateInfo and assigns it to the ResultStateInfo field.
func (o *FileDownloadJobProcessMonitor) SetResultStateInfo(v FileDownloadJobProcessMonitorResultStateInfo) {
	o.ResultStateInfo = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *FileDownloadJobProcessMonitor) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetEndTime() time.Time {
	if o == nil || isNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTime) {
    return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *FileDownloadJobProcessMonitor) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetTimer returns the Timer field value if set, zero value otherwise.
func (o *FileDownloadJobProcessMonitor) GetTimer() int32 {
	if o == nil || isNil(o.Timer) {
		var ret int32
		return ret
	}
	return *o.Timer
}

// GetTimerOk returns a tuple with the Timer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobProcessMonitor) GetTimerOk() (*int32, bool) {
	if o == nil || isNil(o.Timer) {
    return nil, false
	}
	return o.Timer, true
}

// HasTimer returns a boolean if a field has been set.
func (o *FileDownloadJobProcessMonitor) HasTimer() bool {
	if o != nil && !isNil(o.Timer) {
		return true
	}

	return false
}

// SetTimer gets a reference to the given int32 and assigns it to the Timer field.
func (o *FileDownloadJobProcessMonitor) SetTimer(v int32) {
	o.Timer = &v
}

func (o FileDownloadJobProcessMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ProgressPercentage) {
		toSerialize["progressPercentage"] = o.ProgressPercentage
	}
	if !isNil(o.ProgressStateInfo) {
		toSerialize["progressStateInfo"] = o.ProgressStateInfo
	}
	if !isNil(o.ResultStateInfo) {
		toSerialize["resultStateInfo"] = o.ResultStateInfo
	}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !isNil(o.Timer) {
		toSerialize["timer"] = o.Timer
	}
	return json.Marshal(toSerialize)
}

type NullableFileDownloadJobProcessMonitor struct {
	value *FileDownloadJobProcessMonitor
	isSet bool
}

func (v NullableFileDownloadJobProcessMonitor) Get() *FileDownloadJobProcessMonitor {
	return v.value
}

func (v *NullableFileDownloadJobProcessMonitor) Set(val *FileDownloadJobProcessMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDownloadJobProcessMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDownloadJobProcessMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDownloadJobProcessMonitor(val *FileDownloadJobProcessMonitor) *NullableFileDownloadJobProcessMonitor {
	return &NullableFileDownloadJobProcessMonitor{value: val, isSet: true}
}

func (v NullableFileDownloadJobProcessMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDownloadJobProcessMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

