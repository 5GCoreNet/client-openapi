/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
	"time"
)

// DataReportingSession A representation of a Data Reporting Session.
type DataReportingSession struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	SessionId *string `json:"sessionId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// String providing an application identifier.
	ExternalApplicationId string `json:"externalApplicationId"`
	SupportedDomains []DataDomain `json:"supportedDomains"`
	ReportingConditions *DataReportingSessionReportingConditions `json:"reportingConditions,omitempty"`
}

// NewDataReportingSession instantiates a new DataReportingSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataReportingSession(externalApplicationId string, supportedDomains []DataDomain) *DataReportingSession {
	this := DataReportingSession{}
	this.ExternalApplicationId = externalApplicationId
	this.SupportedDomains = supportedDomains
	return &this
}

// NewDataReportingSessionWithDefaults instantiates a new DataReportingSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataReportingSessionWithDefaults() *DataReportingSession {
	this := DataReportingSession{}
	return &this
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DataReportingSession) GetSessionId() string {
	if o == nil || isNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReportingSession) GetSessionIdOk() (*string, bool) {
	if o == nil || isNil(o.SessionId) {
    return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DataReportingSession) HasSessionId() bool {
	if o != nil && !isNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DataReportingSession) SetSessionId(v string) {
	o.SessionId = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *DataReportingSession) GetValidUntil() time.Time {
	if o == nil || isNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReportingSession) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidUntil) {
    return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *DataReportingSession) HasValidUntil() bool {
	if o != nil && !isNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *DataReportingSession) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetExternalApplicationId returns the ExternalApplicationId field value
func (o *DataReportingSession) GetExternalApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalApplicationId
}

// GetExternalApplicationIdOk returns a tuple with the ExternalApplicationId field value
// and a boolean to check if the value has been set.
func (o *DataReportingSession) GetExternalApplicationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExternalApplicationId, true
}

// SetExternalApplicationId sets field value
func (o *DataReportingSession) SetExternalApplicationId(v string) {
	o.ExternalApplicationId = v
}

// GetSupportedDomains returns the SupportedDomains field value
func (o *DataReportingSession) GetSupportedDomains() []DataDomain {
	if o == nil {
		var ret []DataDomain
		return ret
	}

	return o.SupportedDomains
}

// GetSupportedDomainsOk returns a tuple with the SupportedDomains field value
// and a boolean to check if the value has been set.
func (o *DataReportingSession) GetSupportedDomainsOk() ([]DataDomain, bool) {
	if o == nil {
    return nil, false
	}
	return o.SupportedDomains, true
}

// SetSupportedDomains sets field value
func (o *DataReportingSession) SetSupportedDomains(v []DataDomain) {
	o.SupportedDomains = v
}

// GetReportingConditions returns the ReportingConditions field value if set, zero value otherwise.
func (o *DataReportingSession) GetReportingConditions() DataReportingSessionReportingConditions {
	if o == nil || isNil(o.ReportingConditions) {
		var ret DataReportingSessionReportingConditions
		return ret
	}
	return *o.ReportingConditions
}

// GetReportingConditionsOk returns a tuple with the ReportingConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReportingSession) GetReportingConditionsOk() (*DataReportingSessionReportingConditions, bool) {
	if o == nil || isNil(o.ReportingConditions) {
    return nil, false
	}
	return o.ReportingConditions, true
}

// HasReportingConditions returns a boolean if a field has been set.
func (o *DataReportingSession) HasReportingConditions() bool {
	if o != nil && !isNil(o.ReportingConditions) {
		return true
	}

	return false
}

// SetReportingConditions gets a reference to the given DataReportingSessionReportingConditions and assigns it to the ReportingConditions field.
func (o *DataReportingSession) SetReportingConditions(v DataReportingSessionReportingConditions) {
	o.ReportingConditions = &v
}

func (o DataReportingSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	if !isNil(o.ValidUntil) {
		toSerialize["validUntil"] = o.ValidUntil
	}
	if true {
		toSerialize["externalApplicationId"] = o.ExternalApplicationId
	}
	if true {
		toSerialize["supportedDomains"] = o.SupportedDomains
	}
	if !isNil(o.ReportingConditions) {
		toSerialize["reportingConditions"] = o.ReportingConditions
	}
	return json.Marshal(toSerialize)
}

type NullableDataReportingSession struct {
	value *DataReportingSession
	isSet bool
}

func (v NullableDataReportingSession) Get() *DataReportingSession {
	return v.value
}

func (v *NullableDataReportingSession) Set(val *DataReportingSession) {
	v.value = val
	v.isSet = true
}

func (v NullableDataReportingSession) IsSet() bool {
	return v.isSet
}

func (v *NullableDataReportingSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataReportingSession(val *DataReportingSession) *NullableDataReportingSession {
	return &NullableDataReportingSession{value: val, isSet: true}
}

func (v NullableDataReportingSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataReportingSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


