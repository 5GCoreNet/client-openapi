/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package PerfMeasJobCtrlMnS

import (
	"encoding/json"
)

// MeasJobCreationRequestType struct for MeasJobCreationRequestType
type MeasJobCreationRequestType struct {
	IOCName *string `json:"iOCName,omitempty"`
	IOCInstanceList []string `json:"iOCInstanceList,omitempty"`
	MeasurementCategoryList []string `json:"measurementCategoryList,omitempty"`
	ReportingMethod *ReportingMethodType `json:"reportingMethod,omitempty"`
	GranularityPeriod *int32 `json:"granularityPeriod,omitempty"`
	ReportingPeriod *int32 `json:"reportingPeriod,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
	StopTime *string `json:"stopTime,omitempty"`
	Schedule *ScheduleType `json:"schedule,omitempty"`
	StreamTarget *string `json:"streamTarget,omitempty"`
	Priority *PriorityType `json:"priority,omitempty"`
	Reliability *string `json:"reliability,omitempty"`
}

// NewMeasJobCreationRequestType instantiates a new MeasJobCreationRequestType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasJobCreationRequestType() *MeasJobCreationRequestType {
	this := MeasJobCreationRequestType{}
	return &this
}

// NewMeasJobCreationRequestTypeWithDefaults instantiates a new MeasJobCreationRequestType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasJobCreationRequestTypeWithDefaults() *MeasJobCreationRequestType {
	this := MeasJobCreationRequestType{}
	return &this
}

// GetIOCName returns the IOCName field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetIOCName() string {
	if o == nil || isNil(o.IOCName) {
		var ret string
		return ret
	}
	return *o.IOCName
}

// GetIOCNameOk returns a tuple with the IOCName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetIOCNameOk() (*string, bool) {
	if o == nil || isNil(o.IOCName) {
    return nil, false
	}
	return o.IOCName, true
}

// HasIOCName returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasIOCName() bool {
	if o != nil && !isNil(o.IOCName) {
		return true
	}

	return false
}

// SetIOCName gets a reference to the given string and assigns it to the IOCName field.
func (o *MeasJobCreationRequestType) SetIOCName(v string) {
	o.IOCName = &v
}

// GetIOCInstanceList returns the IOCInstanceList field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetIOCInstanceList() []string {
	if o == nil || isNil(o.IOCInstanceList) {
		var ret []string
		return ret
	}
	return o.IOCInstanceList
}

// GetIOCInstanceListOk returns a tuple with the IOCInstanceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetIOCInstanceListOk() ([]string, bool) {
	if o == nil || isNil(o.IOCInstanceList) {
    return nil, false
	}
	return o.IOCInstanceList, true
}

// HasIOCInstanceList returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasIOCInstanceList() bool {
	if o != nil && !isNil(o.IOCInstanceList) {
		return true
	}

	return false
}

// SetIOCInstanceList gets a reference to the given []string and assigns it to the IOCInstanceList field.
func (o *MeasJobCreationRequestType) SetIOCInstanceList(v []string) {
	o.IOCInstanceList = v
}

// GetMeasurementCategoryList returns the MeasurementCategoryList field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetMeasurementCategoryList() []string {
	if o == nil || isNil(o.MeasurementCategoryList) {
		var ret []string
		return ret
	}
	return o.MeasurementCategoryList
}

// GetMeasurementCategoryListOk returns a tuple with the MeasurementCategoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetMeasurementCategoryListOk() ([]string, bool) {
	if o == nil || isNil(o.MeasurementCategoryList) {
    return nil, false
	}
	return o.MeasurementCategoryList, true
}

// HasMeasurementCategoryList returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasMeasurementCategoryList() bool {
	if o != nil && !isNil(o.MeasurementCategoryList) {
		return true
	}

	return false
}

// SetMeasurementCategoryList gets a reference to the given []string and assigns it to the MeasurementCategoryList field.
func (o *MeasJobCreationRequestType) SetMeasurementCategoryList(v []string) {
	o.MeasurementCategoryList = v
}

// GetReportingMethod returns the ReportingMethod field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetReportingMethod() ReportingMethodType {
	if o == nil || isNil(o.ReportingMethod) {
		var ret ReportingMethodType
		return ret
	}
	return *o.ReportingMethod
}

// GetReportingMethodOk returns a tuple with the ReportingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetReportingMethodOk() (*ReportingMethodType, bool) {
	if o == nil || isNil(o.ReportingMethod) {
    return nil, false
	}
	return o.ReportingMethod, true
}

// HasReportingMethod returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasReportingMethod() bool {
	if o != nil && !isNil(o.ReportingMethod) {
		return true
	}

	return false
}

// SetReportingMethod gets a reference to the given ReportingMethodType and assigns it to the ReportingMethod field.
func (o *MeasJobCreationRequestType) SetReportingMethod(v ReportingMethodType) {
	o.ReportingMethod = &v
}

// GetGranularityPeriod returns the GranularityPeriod field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetGranularityPeriod() int32 {
	if o == nil || isNil(o.GranularityPeriod) {
		var ret int32
		return ret
	}
	return *o.GranularityPeriod
}

// GetGranularityPeriodOk returns a tuple with the GranularityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetGranularityPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.GranularityPeriod) {
    return nil, false
	}
	return o.GranularityPeriod, true
}

// HasGranularityPeriod returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasGranularityPeriod() bool {
	if o != nil && !isNil(o.GranularityPeriod) {
		return true
	}

	return false
}

// SetGranularityPeriod gets a reference to the given int32 and assigns it to the GranularityPeriod field.
func (o *MeasJobCreationRequestType) SetGranularityPeriod(v int32) {
	o.GranularityPeriod = &v
}

// GetReportingPeriod returns the ReportingPeriod field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetReportingPeriod() int32 {
	if o == nil || isNil(o.ReportingPeriod) {
		var ret int32
		return ret
	}
	return *o.ReportingPeriod
}

// GetReportingPeriodOk returns a tuple with the ReportingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetReportingPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.ReportingPeriod) {
    return nil, false
	}
	return o.ReportingPeriod, true
}

// HasReportingPeriod returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasReportingPeriod() bool {
	if o != nil && !isNil(o.ReportingPeriod) {
		return true
	}

	return false
}

// SetReportingPeriod gets a reference to the given int32 and assigns it to the ReportingPeriod field.
func (o *MeasJobCreationRequestType) SetReportingPeriod(v int32) {
	o.ReportingPeriod = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetStartTime() string {
	if o == nil || isNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetStartTimeOk() (*string, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *MeasJobCreationRequestType) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetStopTime() string {
	if o == nil || isNil(o.StopTime) {
		var ret string
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetStopTimeOk() (*string, bool) {
	if o == nil || isNil(o.StopTime) {
    return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasStopTime() bool {
	if o != nil && !isNil(o.StopTime) {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given string and assigns it to the StopTime field.
func (o *MeasJobCreationRequestType) SetStopTime(v string) {
	o.StopTime = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetSchedule() ScheduleType {
	if o == nil || isNil(o.Schedule) {
		var ret ScheduleType
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetScheduleOk() (*ScheduleType, bool) {
	if o == nil || isNil(o.Schedule) {
    return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ScheduleType and assigns it to the Schedule field.
func (o *MeasJobCreationRequestType) SetSchedule(v ScheduleType) {
	o.Schedule = &v
}

// GetStreamTarget returns the StreamTarget field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetStreamTarget() string {
	if o == nil || isNil(o.StreamTarget) {
		var ret string
		return ret
	}
	return *o.StreamTarget
}

// GetStreamTargetOk returns a tuple with the StreamTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetStreamTargetOk() (*string, bool) {
	if o == nil || isNil(o.StreamTarget) {
    return nil, false
	}
	return o.StreamTarget, true
}

// HasStreamTarget returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasStreamTarget() bool {
	if o != nil && !isNil(o.StreamTarget) {
		return true
	}

	return false
}

// SetStreamTarget gets a reference to the given string and assigns it to the StreamTarget field.
func (o *MeasJobCreationRequestType) SetStreamTarget(v string) {
	o.StreamTarget = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetPriority() PriorityType {
	if o == nil || isNil(o.Priority) {
		var ret PriorityType
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetPriorityOk() (*PriorityType, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given PriorityType and assigns it to the Priority field.
func (o *MeasJobCreationRequestType) SetPriority(v PriorityType) {
	o.Priority = &v
}

// GetReliability returns the Reliability field value if set, zero value otherwise.
func (o *MeasJobCreationRequestType) GetReliability() string {
	if o == nil || isNil(o.Reliability) {
		var ret string
		return ret
	}
	return *o.Reliability
}

// GetReliabilityOk returns a tuple with the Reliability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobCreationRequestType) GetReliabilityOk() (*string, bool) {
	if o == nil || isNil(o.Reliability) {
    return nil, false
	}
	return o.Reliability, true
}

// HasReliability returns a boolean if a field has been set.
func (o *MeasJobCreationRequestType) HasReliability() bool {
	if o != nil && !isNil(o.Reliability) {
		return true
	}

	return false
}

// SetReliability gets a reference to the given string and assigns it to the Reliability field.
func (o *MeasJobCreationRequestType) SetReliability(v string) {
	o.Reliability = &v
}

func (o MeasJobCreationRequestType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IOCName) {
		toSerialize["iOCName"] = o.IOCName
	}
	if !isNil(o.IOCInstanceList) {
		toSerialize["iOCInstanceList"] = o.IOCInstanceList
	}
	if !isNil(o.MeasurementCategoryList) {
		toSerialize["measurementCategoryList"] = o.MeasurementCategoryList
	}
	if !isNil(o.ReportingMethod) {
		toSerialize["reportingMethod"] = o.ReportingMethod
	}
	if !isNil(o.GranularityPeriod) {
		toSerialize["granularityPeriod"] = o.GranularityPeriod
	}
	if !isNil(o.ReportingPeriod) {
		toSerialize["reportingPeriod"] = o.ReportingPeriod
	}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.StopTime) {
		toSerialize["stopTime"] = o.StopTime
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !isNil(o.StreamTarget) {
		toSerialize["streamTarget"] = o.StreamTarget
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.Reliability) {
		toSerialize["reliability"] = o.Reliability
	}
	return json.Marshal(toSerialize)
}

type NullableMeasJobCreationRequestType struct {
	value *MeasJobCreationRequestType
	isSet bool
}

func (v NullableMeasJobCreationRequestType) Get() *MeasJobCreationRequestType {
	return v.value
}

func (v *NullableMeasJobCreationRequestType) Set(val *MeasJobCreationRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasJobCreationRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasJobCreationRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasJobCreationRequestType(val *MeasJobCreationRequestType) *NullableMeasJobCreationRequestType {
	return &NullableMeasJobCreationRequestType{value: val, isSet: true}
}

func (v NullableMeasJobCreationRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasJobCreationRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

