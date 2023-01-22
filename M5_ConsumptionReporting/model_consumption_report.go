/*
M5_ConsumptionReporting

5GMS AF M5 Consumption Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M5_ConsumptionReporting

import (
	"encoding/json"
)

// ConsumptionReport A representation of a Consumption Report.
type ConsumptionReport struct {
	MediaPlayerEntry string `json:"mediaPlayerEntry"`
	ReportingClientId string `json:"reportingClientId"`
	ConsumptionReportingUnits []ConsumptionReportingUnit `json:"consumptionReportingUnits"`
}

// NewConsumptionReport instantiates a new ConsumptionReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumptionReport(mediaPlayerEntry string, reportingClientId string, consumptionReportingUnits []ConsumptionReportingUnit) *ConsumptionReport {
	this := ConsumptionReport{}
	this.MediaPlayerEntry = mediaPlayerEntry
	this.ReportingClientId = reportingClientId
	this.ConsumptionReportingUnits = consumptionReportingUnits
	return &this
}

// NewConsumptionReportWithDefaults instantiates a new ConsumptionReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumptionReportWithDefaults() *ConsumptionReport {
	this := ConsumptionReport{}
	return &this
}

// GetMediaPlayerEntry returns the MediaPlayerEntry field value
func (o *ConsumptionReport) GetMediaPlayerEntry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaPlayerEntry
}

// GetMediaPlayerEntryOk returns a tuple with the MediaPlayerEntry field value
// and a boolean to check if the value has been set.
func (o *ConsumptionReport) GetMediaPlayerEntryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MediaPlayerEntry, true
}

// SetMediaPlayerEntry sets field value
func (o *ConsumptionReport) SetMediaPlayerEntry(v string) {
	o.MediaPlayerEntry = v
}

// GetReportingClientId returns the ReportingClientId field value
func (o *ConsumptionReport) GetReportingClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportingClientId
}

// GetReportingClientIdOk returns a tuple with the ReportingClientId field value
// and a boolean to check if the value has been set.
func (o *ConsumptionReport) GetReportingClientIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReportingClientId, true
}

// SetReportingClientId sets field value
func (o *ConsumptionReport) SetReportingClientId(v string) {
	o.ReportingClientId = v
}

// GetConsumptionReportingUnits returns the ConsumptionReportingUnits field value
func (o *ConsumptionReport) GetConsumptionReportingUnits() []ConsumptionReportingUnit {
	if o == nil {
		var ret []ConsumptionReportingUnit
		return ret
	}

	return o.ConsumptionReportingUnits
}

// GetConsumptionReportingUnitsOk returns a tuple with the ConsumptionReportingUnits field value
// and a boolean to check if the value has been set.
func (o *ConsumptionReport) GetConsumptionReportingUnitsOk() ([]ConsumptionReportingUnit, bool) {
	if o == nil {
    return nil, false
	}
	return o.ConsumptionReportingUnits, true
}

// SetConsumptionReportingUnits sets field value
func (o *ConsumptionReport) SetConsumptionReportingUnits(v []ConsumptionReportingUnit) {
	o.ConsumptionReportingUnits = v
}

func (o ConsumptionReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mediaPlayerEntry"] = o.MediaPlayerEntry
	}
	if true {
		toSerialize["reportingClientId"] = o.ReportingClientId
	}
	if true {
		toSerialize["consumptionReportingUnits"] = o.ConsumptionReportingUnits
	}
	return json.Marshal(toSerialize)
}

type NullableConsumptionReport struct {
	value *ConsumptionReport
	isSet bool
}

func (v NullableConsumptionReport) Get() *ConsumptionReport {
	return v.value
}

func (v *NullableConsumptionReport) Set(val *ConsumptionReport) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumptionReport) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumptionReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumptionReport(val *ConsumptionReport) *NullableConsumptionReport {
	return &NullableConsumptionReport{value: val, isSet: true}
}

func (v NullableConsumptionReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumptionReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


