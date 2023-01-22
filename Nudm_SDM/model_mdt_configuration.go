/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
)

// MdtConfiguration contains contain MDT configuration data.
type MdtConfiguration struct {
	JobType JobType `json:"jobType"`
	ReportType *ReportTypeMdt `json:"reportType,omitempty"`
	AreaScope *AreaScope `json:"areaScope,omitempty"`
	MeasurementLteList []MeasurementLteForMdt `json:"measurementLteList,omitempty"`
	MeasurementNrList []MeasurementNrForMdt `json:"measurementNrList,omitempty"`
	SensorMeasurementList []SensorMeasurement `json:"sensorMeasurementList,omitempty"`
	ReportingTriggerList []ReportingTrigger `json:"reportingTriggerList,omitempty"`
	ReportInterval *ReportIntervalMdt `json:"reportInterval,omitempty"`
	ReportIntervalNr *ReportIntervalNrMdt `json:"reportIntervalNr,omitempty"`
	ReportAmount *ReportAmountMdt `json:"reportAmount,omitempty"`
	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE. When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-97. 
	EventThresholdRsrp *int32 `json:"eventThresholdRsrp,omitempty"`
	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR. When present, this IE shall indicate the Event Threshold for RSRP, and the value shall be between 0-127. 
	EventThresholdRsrpNr *int32 `json:"eventThresholdRsrpNr,omitempty"`
	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in LTE.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-34. 
	EventThresholdRsrq *int32 `json:"eventThresholdRsrq,omitempty"`
	// This IE shall be present if the report trigger parameter is configured for A2 event reporting or A2 event triggered periodic reporting and the job type parameter is configured for Immediate MDT or combined Immediate MDT and Trace in NR.When present, this IE shall indicate the Event Threshold for RSRQ, and the value shall be between 0-127. 
	EventThresholdRsrqNr *int32 `json:"eventThresholdRsrqNr,omitempty"`
	EventList []EventForMdt `json:"eventList,omitempty"`
	LoggingInterval *LoggingIntervalMdt `json:"loggingInterval,omitempty"`
	LoggingIntervalNr *LoggingIntervalNrMdt `json:"loggingIntervalNr,omitempty"`
	LoggingDuration *LoggingDurationMdt `json:"loggingDuration,omitempty"`
	LoggingDurationNr *LoggingDurationNrMdt `json:"loggingDurationNr,omitempty"`
	PositioningMethod *PositioningMethodMdt `json:"positioningMethod,omitempty"`
	AddPositioningMethodList []PositioningMethodMdt `json:"addPositioningMethodList,omitempty"`
	CollectionPeriodRmmLte *CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty"`
	CollectionPeriodRmmNr *CollectionPeriodRmmNrMdt `json:"collectionPeriodRmmNr,omitempty"`
	MeasurementPeriodLte *MeasurementPeriodLteMdt `json:"measurementPeriodLte,omitempty"`
	MdtAllowedPlmnIdList []PlmnId `json:"mdtAllowedPlmnIdList,omitempty"`
	MbsfnAreaList []MbsfnArea `json:"mbsfnAreaList,omitempty"`
	InterFreqTargetList []InterFreqTargetInfo `json:"interFreqTargetList,omitempty"`
}

// NewMdtConfiguration instantiates a new MdtConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMdtConfiguration(jobType JobType) *MdtConfiguration {
	this := MdtConfiguration{}
	this.JobType = jobType
	return &this
}

// NewMdtConfigurationWithDefaults instantiates a new MdtConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMdtConfigurationWithDefaults() *MdtConfiguration {
	this := MdtConfiguration{}
	return &this
}

// GetJobType returns the JobType field value
func (o *MdtConfiguration) GetJobType() JobType {
	if o == nil {
		var ret JobType
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetJobTypeOk() (*JobType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *MdtConfiguration) SetJobType(v JobType) {
	o.JobType = v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *MdtConfiguration) GetReportType() ReportTypeMdt {
	if o == nil || isNil(o.ReportType) {
		var ret ReportTypeMdt
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetReportTypeOk() (*ReportTypeMdt, bool) {
	if o == nil || isNil(o.ReportType) {
    return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *MdtConfiguration) HasReportType() bool {
	if o != nil && !isNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given ReportTypeMdt and assigns it to the ReportType field.
func (o *MdtConfiguration) SetReportType(v ReportTypeMdt) {
	o.ReportType = &v
}

// GetAreaScope returns the AreaScope field value if set, zero value otherwise.
func (o *MdtConfiguration) GetAreaScope() AreaScope {
	if o == nil || isNil(o.AreaScope) {
		var ret AreaScope
		return ret
	}
	return *o.AreaScope
}

// GetAreaScopeOk returns a tuple with the AreaScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetAreaScopeOk() (*AreaScope, bool) {
	if o == nil || isNil(o.AreaScope) {
    return nil, false
	}
	return o.AreaScope, true
}

// HasAreaScope returns a boolean if a field has been set.
func (o *MdtConfiguration) HasAreaScope() bool {
	if o != nil && !isNil(o.AreaScope) {
		return true
	}

	return false
}

// SetAreaScope gets a reference to the given AreaScope and assigns it to the AreaScope field.
func (o *MdtConfiguration) SetAreaScope(v AreaScope) {
	o.AreaScope = &v
}

// GetMeasurementLteList returns the MeasurementLteList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetMeasurementLteList() []MeasurementLteForMdt {
	if o == nil || isNil(o.MeasurementLteList) {
		var ret []MeasurementLteForMdt
		return ret
	}
	return o.MeasurementLteList
}

// GetMeasurementLteListOk returns a tuple with the MeasurementLteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetMeasurementLteListOk() ([]MeasurementLteForMdt, bool) {
	if o == nil || isNil(o.MeasurementLteList) {
    return nil, false
	}
	return o.MeasurementLteList, true
}

// HasMeasurementLteList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasMeasurementLteList() bool {
	if o != nil && !isNil(o.MeasurementLteList) {
		return true
	}

	return false
}

// SetMeasurementLteList gets a reference to the given []MeasurementLteForMdt and assigns it to the MeasurementLteList field.
func (o *MdtConfiguration) SetMeasurementLteList(v []MeasurementLteForMdt) {
	o.MeasurementLteList = v
}

// GetMeasurementNrList returns the MeasurementNrList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetMeasurementNrList() []MeasurementNrForMdt {
	if o == nil || isNil(o.MeasurementNrList) {
		var ret []MeasurementNrForMdt
		return ret
	}
	return o.MeasurementNrList
}

// GetMeasurementNrListOk returns a tuple with the MeasurementNrList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetMeasurementNrListOk() ([]MeasurementNrForMdt, bool) {
	if o == nil || isNil(o.MeasurementNrList) {
    return nil, false
	}
	return o.MeasurementNrList, true
}

// HasMeasurementNrList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasMeasurementNrList() bool {
	if o != nil && !isNil(o.MeasurementNrList) {
		return true
	}

	return false
}

// SetMeasurementNrList gets a reference to the given []MeasurementNrForMdt and assigns it to the MeasurementNrList field.
func (o *MdtConfiguration) SetMeasurementNrList(v []MeasurementNrForMdt) {
	o.MeasurementNrList = v
}

// GetSensorMeasurementList returns the SensorMeasurementList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetSensorMeasurementList() []SensorMeasurement {
	if o == nil || isNil(o.SensorMeasurementList) {
		var ret []SensorMeasurement
		return ret
	}
	return o.SensorMeasurementList
}

// GetSensorMeasurementListOk returns a tuple with the SensorMeasurementList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetSensorMeasurementListOk() ([]SensorMeasurement, bool) {
	if o == nil || isNil(o.SensorMeasurementList) {
    return nil, false
	}
	return o.SensorMeasurementList, true
}

// HasSensorMeasurementList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasSensorMeasurementList() bool {
	if o != nil && !isNil(o.SensorMeasurementList) {
		return true
	}

	return false
}

// SetSensorMeasurementList gets a reference to the given []SensorMeasurement and assigns it to the SensorMeasurementList field.
func (o *MdtConfiguration) SetSensorMeasurementList(v []SensorMeasurement) {
	o.SensorMeasurementList = v
}

// GetReportingTriggerList returns the ReportingTriggerList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetReportingTriggerList() []ReportingTrigger {
	if o == nil || isNil(o.ReportingTriggerList) {
		var ret []ReportingTrigger
		return ret
	}
	return o.ReportingTriggerList
}

// GetReportingTriggerListOk returns a tuple with the ReportingTriggerList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetReportingTriggerListOk() ([]ReportingTrigger, bool) {
	if o == nil || isNil(o.ReportingTriggerList) {
    return nil, false
	}
	return o.ReportingTriggerList, true
}

// HasReportingTriggerList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasReportingTriggerList() bool {
	if o != nil && !isNil(o.ReportingTriggerList) {
		return true
	}

	return false
}

// SetReportingTriggerList gets a reference to the given []ReportingTrigger and assigns it to the ReportingTriggerList field.
func (o *MdtConfiguration) SetReportingTriggerList(v []ReportingTrigger) {
	o.ReportingTriggerList = v
}

// GetReportInterval returns the ReportInterval field value if set, zero value otherwise.
func (o *MdtConfiguration) GetReportInterval() ReportIntervalMdt {
	if o == nil || isNil(o.ReportInterval) {
		var ret ReportIntervalMdt
		return ret
	}
	return *o.ReportInterval
}

// GetReportIntervalOk returns a tuple with the ReportInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetReportIntervalOk() (*ReportIntervalMdt, bool) {
	if o == nil || isNil(o.ReportInterval) {
    return nil, false
	}
	return o.ReportInterval, true
}

// HasReportInterval returns a boolean if a field has been set.
func (o *MdtConfiguration) HasReportInterval() bool {
	if o != nil && !isNil(o.ReportInterval) {
		return true
	}

	return false
}

// SetReportInterval gets a reference to the given ReportIntervalMdt and assigns it to the ReportInterval field.
func (o *MdtConfiguration) SetReportInterval(v ReportIntervalMdt) {
	o.ReportInterval = &v
}

// GetReportIntervalNr returns the ReportIntervalNr field value if set, zero value otherwise.
func (o *MdtConfiguration) GetReportIntervalNr() ReportIntervalNrMdt {
	if o == nil || isNil(o.ReportIntervalNr) {
		var ret ReportIntervalNrMdt
		return ret
	}
	return *o.ReportIntervalNr
}

// GetReportIntervalNrOk returns a tuple with the ReportIntervalNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetReportIntervalNrOk() (*ReportIntervalNrMdt, bool) {
	if o == nil || isNil(o.ReportIntervalNr) {
    return nil, false
	}
	return o.ReportIntervalNr, true
}

// HasReportIntervalNr returns a boolean if a field has been set.
func (o *MdtConfiguration) HasReportIntervalNr() bool {
	if o != nil && !isNil(o.ReportIntervalNr) {
		return true
	}

	return false
}

// SetReportIntervalNr gets a reference to the given ReportIntervalNrMdt and assigns it to the ReportIntervalNr field.
func (o *MdtConfiguration) SetReportIntervalNr(v ReportIntervalNrMdt) {
	o.ReportIntervalNr = &v
}

// GetReportAmount returns the ReportAmount field value if set, zero value otherwise.
func (o *MdtConfiguration) GetReportAmount() ReportAmountMdt {
	if o == nil || isNil(o.ReportAmount) {
		var ret ReportAmountMdt
		return ret
	}
	return *o.ReportAmount
}

// GetReportAmountOk returns a tuple with the ReportAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetReportAmountOk() (*ReportAmountMdt, bool) {
	if o == nil || isNil(o.ReportAmount) {
    return nil, false
	}
	return o.ReportAmount, true
}

// HasReportAmount returns a boolean if a field has been set.
func (o *MdtConfiguration) HasReportAmount() bool {
	if o != nil && !isNil(o.ReportAmount) {
		return true
	}

	return false
}

// SetReportAmount gets a reference to the given ReportAmountMdt and assigns it to the ReportAmount field.
func (o *MdtConfiguration) SetReportAmount(v ReportAmountMdt) {
	o.ReportAmount = &v
}

// GetEventThresholdRsrp returns the EventThresholdRsrp field value if set, zero value otherwise.
func (o *MdtConfiguration) GetEventThresholdRsrp() int32 {
	if o == nil || isNil(o.EventThresholdRsrp) {
		var ret int32
		return ret
	}
	return *o.EventThresholdRsrp
}

// GetEventThresholdRsrpOk returns a tuple with the EventThresholdRsrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetEventThresholdRsrpOk() (*int32, bool) {
	if o == nil || isNil(o.EventThresholdRsrp) {
    return nil, false
	}
	return o.EventThresholdRsrp, true
}

// HasEventThresholdRsrp returns a boolean if a field has been set.
func (o *MdtConfiguration) HasEventThresholdRsrp() bool {
	if o != nil && !isNil(o.EventThresholdRsrp) {
		return true
	}

	return false
}

// SetEventThresholdRsrp gets a reference to the given int32 and assigns it to the EventThresholdRsrp field.
func (o *MdtConfiguration) SetEventThresholdRsrp(v int32) {
	o.EventThresholdRsrp = &v
}

// GetEventThresholdRsrpNr returns the EventThresholdRsrpNr field value if set, zero value otherwise.
func (o *MdtConfiguration) GetEventThresholdRsrpNr() int32 {
	if o == nil || isNil(o.EventThresholdRsrpNr) {
		var ret int32
		return ret
	}
	return *o.EventThresholdRsrpNr
}

// GetEventThresholdRsrpNrOk returns a tuple with the EventThresholdRsrpNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetEventThresholdRsrpNrOk() (*int32, bool) {
	if o == nil || isNil(o.EventThresholdRsrpNr) {
    return nil, false
	}
	return o.EventThresholdRsrpNr, true
}

// HasEventThresholdRsrpNr returns a boolean if a field has been set.
func (o *MdtConfiguration) HasEventThresholdRsrpNr() bool {
	if o != nil && !isNil(o.EventThresholdRsrpNr) {
		return true
	}

	return false
}

// SetEventThresholdRsrpNr gets a reference to the given int32 and assigns it to the EventThresholdRsrpNr field.
func (o *MdtConfiguration) SetEventThresholdRsrpNr(v int32) {
	o.EventThresholdRsrpNr = &v
}

// GetEventThresholdRsrq returns the EventThresholdRsrq field value if set, zero value otherwise.
func (o *MdtConfiguration) GetEventThresholdRsrq() int32 {
	if o == nil || isNil(o.EventThresholdRsrq) {
		var ret int32
		return ret
	}
	return *o.EventThresholdRsrq
}

// GetEventThresholdRsrqOk returns a tuple with the EventThresholdRsrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetEventThresholdRsrqOk() (*int32, bool) {
	if o == nil || isNil(o.EventThresholdRsrq) {
    return nil, false
	}
	return o.EventThresholdRsrq, true
}

// HasEventThresholdRsrq returns a boolean if a field has been set.
func (o *MdtConfiguration) HasEventThresholdRsrq() bool {
	if o != nil && !isNil(o.EventThresholdRsrq) {
		return true
	}

	return false
}

// SetEventThresholdRsrq gets a reference to the given int32 and assigns it to the EventThresholdRsrq field.
func (o *MdtConfiguration) SetEventThresholdRsrq(v int32) {
	o.EventThresholdRsrq = &v
}

// GetEventThresholdRsrqNr returns the EventThresholdRsrqNr field value if set, zero value otherwise.
func (o *MdtConfiguration) GetEventThresholdRsrqNr() int32 {
	if o == nil || isNil(o.EventThresholdRsrqNr) {
		var ret int32
		return ret
	}
	return *o.EventThresholdRsrqNr
}

// GetEventThresholdRsrqNrOk returns a tuple with the EventThresholdRsrqNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetEventThresholdRsrqNrOk() (*int32, bool) {
	if o == nil || isNil(o.EventThresholdRsrqNr) {
    return nil, false
	}
	return o.EventThresholdRsrqNr, true
}

// HasEventThresholdRsrqNr returns a boolean if a field has been set.
func (o *MdtConfiguration) HasEventThresholdRsrqNr() bool {
	if o != nil && !isNil(o.EventThresholdRsrqNr) {
		return true
	}

	return false
}

// SetEventThresholdRsrqNr gets a reference to the given int32 and assigns it to the EventThresholdRsrqNr field.
func (o *MdtConfiguration) SetEventThresholdRsrqNr(v int32) {
	o.EventThresholdRsrqNr = &v
}

// GetEventList returns the EventList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetEventList() []EventForMdt {
	if o == nil || isNil(o.EventList) {
		var ret []EventForMdt
		return ret
	}
	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetEventListOk() ([]EventForMdt, bool) {
	if o == nil || isNil(o.EventList) {
    return nil, false
	}
	return o.EventList, true
}

// HasEventList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasEventList() bool {
	if o != nil && !isNil(o.EventList) {
		return true
	}

	return false
}

// SetEventList gets a reference to the given []EventForMdt and assigns it to the EventList field.
func (o *MdtConfiguration) SetEventList(v []EventForMdt) {
	o.EventList = v
}

// GetLoggingInterval returns the LoggingInterval field value if set, zero value otherwise.
func (o *MdtConfiguration) GetLoggingInterval() LoggingIntervalMdt {
	if o == nil || isNil(o.LoggingInterval) {
		var ret LoggingIntervalMdt
		return ret
	}
	return *o.LoggingInterval
}

// GetLoggingIntervalOk returns a tuple with the LoggingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetLoggingIntervalOk() (*LoggingIntervalMdt, bool) {
	if o == nil || isNil(o.LoggingInterval) {
    return nil, false
	}
	return o.LoggingInterval, true
}

// HasLoggingInterval returns a boolean if a field has been set.
func (o *MdtConfiguration) HasLoggingInterval() bool {
	if o != nil && !isNil(o.LoggingInterval) {
		return true
	}

	return false
}

// SetLoggingInterval gets a reference to the given LoggingIntervalMdt and assigns it to the LoggingInterval field.
func (o *MdtConfiguration) SetLoggingInterval(v LoggingIntervalMdt) {
	o.LoggingInterval = &v
}

// GetLoggingIntervalNr returns the LoggingIntervalNr field value if set, zero value otherwise.
func (o *MdtConfiguration) GetLoggingIntervalNr() LoggingIntervalNrMdt {
	if o == nil || isNil(o.LoggingIntervalNr) {
		var ret LoggingIntervalNrMdt
		return ret
	}
	return *o.LoggingIntervalNr
}

// GetLoggingIntervalNrOk returns a tuple with the LoggingIntervalNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetLoggingIntervalNrOk() (*LoggingIntervalNrMdt, bool) {
	if o == nil || isNil(o.LoggingIntervalNr) {
    return nil, false
	}
	return o.LoggingIntervalNr, true
}

// HasLoggingIntervalNr returns a boolean if a field has been set.
func (o *MdtConfiguration) HasLoggingIntervalNr() bool {
	if o != nil && !isNil(o.LoggingIntervalNr) {
		return true
	}

	return false
}

// SetLoggingIntervalNr gets a reference to the given LoggingIntervalNrMdt and assigns it to the LoggingIntervalNr field.
func (o *MdtConfiguration) SetLoggingIntervalNr(v LoggingIntervalNrMdt) {
	o.LoggingIntervalNr = &v
}

// GetLoggingDuration returns the LoggingDuration field value if set, zero value otherwise.
func (o *MdtConfiguration) GetLoggingDuration() LoggingDurationMdt {
	if o == nil || isNil(o.LoggingDuration) {
		var ret LoggingDurationMdt
		return ret
	}
	return *o.LoggingDuration
}

// GetLoggingDurationOk returns a tuple with the LoggingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetLoggingDurationOk() (*LoggingDurationMdt, bool) {
	if o == nil || isNil(o.LoggingDuration) {
    return nil, false
	}
	return o.LoggingDuration, true
}

// HasLoggingDuration returns a boolean if a field has been set.
func (o *MdtConfiguration) HasLoggingDuration() bool {
	if o != nil && !isNil(o.LoggingDuration) {
		return true
	}

	return false
}

// SetLoggingDuration gets a reference to the given LoggingDurationMdt and assigns it to the LoggingDuration field.
func (o *MdtConfiguration) SetLoggingDuration(v LoggingDurationMdt) {
	o.LoggingDuration = &v
}

// GetLoggingDurationNr returns the LoggingDurationNr field value if set, zero value otherwise.
func (o *MdtConfiguration) GetLoggingDurationNr() LoggingDurationNrMdt {
	if o == nil || isNil(o.LoggingDurationNr) {
		var ret LoggingDurationNrMdt
		return ret
	}
	return *o.LoggingDurationNr
}

// GetLoggingDurationNrOk returns a tuple with the LoggingDurationNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetLoggingDurationNrOk() (*LoggingDurationNrMdt, bool) {
	if o == nil || isNil(o.LoggingDurationNr) {
    return nil, false
	}
	return o.LoggingDurationNr, true
}

// HasLoggingDurationNr returns a boolean if a field has been set.
func (o *MdtConfiguration) HasLoggingDurationNr() bool {
	if o != nil && !isNil(o.LoggingDurationNr) {
		return true
	}

	return false
}

// SetLoggingDurationNr gets a reference to the given LoggingDurationNrMdt and assigns it to the LoggingDurationNr field.
func (o *MdtConfiguration) SetLoggingDurationNr(v LoggingDurationNrMdt) {
	o.LoggingDurationNr = &v
}

// GetPositioningMethod returns the PositioningMethod field value if set, zero value otherwise.
func (o *MdtConfiguration) GetPositioningMethod() PositioningMethodMdt {
	if o == nil || isNil(o.PositioningMethod) {
		var ret PositioningMethodMdt
		return ret
	}
	return *o.PositioningMethod
}

// GetPositioningMethodOk returns a tuple with the PositioningMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetPositioningMethodOk() (*PositioningMethodMdt, bool) {
	if o == nil || isNil(o.PositioningMethod) {
    return nil, false
	}
	return o.PositioningMethod, true
}

// HasPositioningMethod returns a boolean if a field has been set.
func (o *MdtConfiguration) HasPositioningMethod() bool {
	if o != nil && !isNil(o.PositioningMethod) {
		return true
	}

	return false
}

// SetPositioningMethod gets a reference to the given PositioningMethodMdt and assigns it to the PositioningMethod field.
func (o *MdtConfiguration) SetPositioningMethod(v PositioningMethodMdt) {
	o.PositioningMethod = &v
}

// GetAddPositioningMethodList returns the AddPositioningMethodList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetAddPositioningMethodList() []PositioningMethodMdt {
	if o == nil || isNil(o.AddPositioningMethodList) {
		var ret []PositioningMethodMdt
		return ret
	}
	return o.AddPositioningMethodList
}

// GetAddPositioningMethodListOk returns a tuple with the AddPositioningMethodList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetAddPositioningMethodListOk() ([]PositioningMethodMdt, bool) {
	if o == nil || isNil(o.AddPositioningMethodList) {
    return nil, false
	}
	return o.AddPositioningMethodList, true
}

// HasAddPositioningMethodList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasAddPositioningMethodList() bool {
	if o != nil && !isNil(o.AddPositioningMethodList) {
		return true
	}

	return false
}

// SetAddPositioningMethodList gets a reference to the given []PositioningMethodMdt and assigns it to the AddPositioningMethodList field.
func (o *MdtConfiguration) SetAddPositioningMethodList(v []PositioningMethodMdt) {
	o.AddPositioningMethodList = v
}

// GetCollectionPeriodRmmLte returns the CollectionPeriodRmmLte field value if set, zero value otherwise.
func (o *MdtConfiguration) GetCollectionPeriodRmmLte() CollectionPeriodRmmLteMdt {
	if o == nil || isNil(o.CollectionPeriodRmmLte) {
		var ret CollectionPeriodRmmLteMdt
		return ret
	}
	return *o.CollectionPeriodRmmLte
}

// GetCollectionPeriodRmmLteOk returns a tuple with the CollectionPeriodRmmLte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetCollectionPeriodRmmLteOk() (*CollectionPeriodRmmLteMdt, bool) {
	if o == nil || isNil(o.CollectionPeriodRmmLte) {
    return nil, false
	}
	return o.CollectionPeriodRmmLte, true
}

// HasCollectionPeriodRmmLte returns a boolean if a field has been set.
func (o *MdtConfiguration) HasCollectionPeriodRmmLte() bool {
	if o != nil && !isNil(o.CollectionPeriodRmmLte) {
		return true
	}

	return false
}

// SetCollectionPeriodRmmLte gets a reference to the given CollectionPeriodRmmLteMdt and assigns it to the CollectionPeriodRmmLte field.
func (o *MdtConfiguration) SetCollectionPeriodRmmLte(v CollectionPeriodRmmLteMdt) {
	o.CollectionPeriodRmmLte = &v
}

// GetCollectionPeriodRmmNr returns the CollectionPeriodRmmNr field value if set, zero value otherwise.
func (o *MdtConfiguration) GetCollectionPeriodRmmNr() CollectionPeriodRmmNrMdt {
	if o == nil || isNil(o.CollectionPeriodRmmNr) {
		var ret CollectionPeriodRmmNrMdt
		return ret
	}
	return *o.CollectionPeriodRmmNr
}

// GetCollectionPeriodRmmNrOk returns a tuple with the CollectionPeriodRmmNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetCollectionPeriodRmmNrOk() (*CollectionPeriodRmmNrMdt, bool) {
	if o == nil || isNil(o.CollectionPeriodRmmNr) {
    return nil, false
	}
	return o.CollectionPeriodRmmNr, true
}

// HasCollectionPeriodRmmNr returns a boolean if a field has been set.
func (o *MdtConfiguration) HasCollectionPeriodRmmNr() bool {
	if o != nil && !isNil(o.CollectionPeriodRmmNr) {
		return true
	}

	return false
}

// SetCollectionPeriodRmmNr gets a reference to the given CollectionPeriodRmmNrMdt and assigns it to the CollectionPeriodRmmNr field.
func (o *MdtConfiguration) SetCollectionPeriodRmmNr(v CollectionPeriodRmmNrMdt) {
	o.CollectionPeriodRmmNr = &v
}

// GetMeasurementPeriodLte returns the MeasurementPeriodLte field value if set, zero value otherwise.
func (o *MdtConfiguration) GetMeasurementPeriodLte() MeasurementPeriodLteMdt {
	if o == nil || isNil(o.MeasurementPeriodLte) {
		var ret MeasurementPeriodLteMdt
		return ret
	}
	return *o.MeasurementPeriodLte
}

// GetMeasurementPeriodLteOk returns a tuple with the MeasurementPeriodLte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetMeasurementPeriodLteOk() (*MeasurementPeriodLteMdt, bool) {
	if o == nil || isNil(o.MeasurementPeriodLte) {
    return nil, false
	}
	return o.MeasurementPeriodLte, true
}

// HasMeasurementPeriodLte returns a boolean if a field has been set.
func (o *MdtConfiguration) HasMeasurementPeriodLte() bool {
	if o != nil && !isNil(o.MeasurementPeriodLte) {
		return true
	}

	return false
}

// SetMeasurementPeriodLte gets a reference to the given MeasurementPeriodLteMdt and assigns it to the MeasurementPeriodLte field.
func (o *MdtConfiguration) SetMeasurementPeriodLte(v MeasurementPeriodLteMdt) {
	o.MeasurementPeriodLte = &v
}

// GetMdtAllowedPlmnIdList returns the MdtAllowedPlmnIdList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetMdtAllowedPlmnIdList() []PlmnId {
	if o == nil || isNil(o.MdtAllowedPlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.MdtAllowedPlmnIdList
}

// GetMdtAllowedPlmnIdListOk returns a tuple with the MdtAllowedPlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetMdtAllowedPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.MdtAllowedPlmnIdList) {
    return nil, false
	}
	return o.MdtAllowedPlmnIdList, true
}

// HasMdtAllowedPlmnIdList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasMdtAllowedPlmnIdList() bool {
	if o != nil && !isNil(o.MdtAllowedPlmnIdList) {
		return true
	}

	return false
}

// SetMdtAllowedPlmnIdList gets a reference to the given []PlmnId and assigns it to the MdtAllowedPlmnIdList field.
func (o *MdtConfiguration) SetMdtAllowedPlmnIdList(v []PlmnId) {
	o.MdtAllowedPlmnIdList = v
}

// GetMbsfnAreaList returns the MbsfnAreaList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetMbsfnAreaList() []MbsfnArea {
	if o == nil || isNil(o.MbsfnAreaList) {
		var ret []MbsfnArea
		return ret
	}
	return o.MbsfnAreaList
}

// GetMbsfnAreaListOk returns a tuple with the MbsfnAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetMbsfnAreaListOk() ([]MbsfnArea, bool) {
	if o == nil || isNil(o.MbsfnAreaList) {
    return nil, false
	}
	return o.MbsfnAreaList, true
}

// HasMbsfnAreaList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasMbsfnAreaList() bool {
	if o != nil && !isNil(o.MbsfnAreaList) {
		return true
	}

	return false
}

// SetMbsfnAreaList gets a reference to the given []MbsfnArea and assigns it to the MbsfnAreaList field.
func (o *MdtConfiguration) SetMbsfnAreaList(v []MbsfnArea) {
	o.MbsfnAreaList = v
}

// GetInterFreqTargetList returns the InterFreqTargetList field value if set, zero value otherwise.
func (o *MdtConfiguration) GetInterFreqTargetList() []InterFreqTargetInfo {
	if o == nil || isNil(o.InterFreqTargetList) {
		var ret []InterFreqTargetInfo
		return ret
	}
	return o.InterFreqTargetList
}

// GetInterFreqTargetListOk returns a tuple with the InterFreqTargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdtConfiguration) GetInterFreqTargetListOk() ([]InterFreqTargetInfo, bool) {
	if o == nil || isNil(o.InterFreqTargetList) {
    return nil, false
	}
	return o.InterFreqTargetList, true
}

// HasInterFreqTargetList returns a boolean if a field has been set.
func (o *MdtConfiguration) HasInterFreqTargetList() bool {
	if o != nil && !isNil(o.InterFreqTargetList) {
		return true
	}

	return false
}

// SetInterFreqTargetList gets a reference to the given []InterFreqTargetInfo and assigns it to the InterFreqTargetList field.
func (o *MdtConfiguration) SetInterFreqTargetList(v []InterFreqTargetInfo) {
	o.InterFreqTargetList = v
}

func (o MdtConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobType"] = o.JobType
	}
	if !isNil(o.ReportType) {
		toSerialize["reportType"] = o.ReportType
	}
	if !isNil(o.AreaScope) {
		toSerialize["areaScope"] = o.AreaScope
	}
	if !isNil(o.MeasurementLteList) {
		toSerialize["measurementLteList"] = o.MeasurementLteList
	}
	if !isNil(o.MeasurementNrList) {
		toSerialize["measurementNrList"] = o.MeasurementNrList
	}
	if !isNil(o.SensorMeasurementList) {
		toSerialize["sensorMeasurementList"] = o.SensorMeasurementList
	}
	if !isNil(o.ReportingTriggerList) {
		toSerialize["reportingTriggerList"] = o.ReportingTriggerList
	}
	if !isNil(o.ReportInterval) {
		toSerialize["reportInterval"] = o.ReportInterval
	}
	if !isNil(o.ReportIntervalNr) {
		toSerialize["reportIntervalNr"] = o.ReportIntervalNr
	}
	if !isNil(o.ReportAmount) {
		toSerialize["reportAmount"] = o.ReportAmount
	}
	if !isNil(o.EventThresholdRsrp) {
		toSerialize["eventThresholdRsrp"] = o.EventThresholdRsrp
	}
	if !isNil(o.EventThresholdRsrpNr) {
		toSerialize["eventThresholdRsrpNr"] = o.EventThresholdRsrpNr
	}
	if !isNil(o.EventThresholdRsrq) {
		toSerialize["eventThresholdRsrq"] = o.EventThresholdRsrq
	}
	if !isNil(o.EventThresholdRsrqNr) {
		toSerialize["eventThresholdRsrqNr"] = o.EventThresholdRsrqNr
	}
	if !isNil(o.EventList) {
		toSerialize["eventList"] = o.EventList
	}
	if !isNil(o.LoggingInterval) {
		toSerialize["loggingInterval"] = o.LoggingInterval
	}
	if !isNil(o.LoggingIntervalNr) {
		toSerialize["loggingIntervalNr"] = o.LoggingIntervalNr
	}
	if !isNil(o.LoggingDuration) {
		toSerialize["loggingDuration"] = o.LoggingDuration
	}
	if !isNil(o.LoggingDurationNr) {
		toSerialize["loggingDurationNr"] = o.LoggingDurationNr
	}
	if !isNil(o.PositioningMethod) {
		toSerialize["positioningMethod"] = o.PositioningMethod
	}
	if !isNil(o.AddPositioningMethodList) {
		toSerialize["addPositioningMethodList"] = o.AddPositioningMethodList
	}
	if !isNil(o.CollectionPeriodRmmLte) {
		toSerialize["collectionPeriodRmmLte"] = o.CollectionPeriodRmmLte
	}
	if !isNil(o.CollectionPeriodRmmNr) {
		toSerialize["collectionPeriodRmmNr"] = o.CollectionPeriodRmmNr
	}
	if !isNil(o.MeasurementPeriodLte) {
		toSerialize["measurementPeriodLte"] = o.MeasurementPeriodLte
	}
	if !isNil(o.MdtAllowedPlmnIdList) {
		toSerialize["mdtAllowedPlmnIdList"] = o.MdtAllowedPlmnIdList
	}
	if !isNil(o.MbsfnAreaList) {
		toSerialize["mbsfnAreaList"] = o.MbsfnAreaList
	}
	if !isNil(o.InterFreqTargetList) {
		toSerialize["interFreqTargetList"] = o.InterFreqTargetList
	}
	return json.Marshal(toSerialize)
}

type NullableMdtConfiguration struct {
	value *MdtConfiguration
	isSet bool
}

func (v NullableMdtConfiguration) Get() *MdtConfiguration {
	return v.value
}

func (v *NullableMdtConfiguration) Set(val *MdtConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMdtConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMdtConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMdtConfiguration(val *MdtConfiguration) *NullableMdtConfiguration {
	return &NullableMdtConfiguration{value: val, isSet: true}
}

func (v NullableMdtConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMdtConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


