/*
Ndcaf_DataReportingProvisioning

Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReportingProvisioning

import (
	"encoding/json"
)

// DataReportingProvisioningSession A representation of a Data Reporting Provisioning Session.
type DataReportingProvisioningSession struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	ProvisioningSessionId string `json:"provisioningSessionId"`
	// Contains an identity of an application service provider.
	AspId string `json:"aspId"`
	// String providing an application identifier.
	ExternalApplicationId string `json:"externalApplicationId"`
	// String providing an application identifier.
	InternalApplicationId *string `json:"internalApplicationId,omitempty"`
	EventId AfEvent `json:"eventId"`
	DataReportingConfigurationIds []string `json:"dataReportingConfigurationIds"`
}

// NewDataReportingProvisioningSession instantiates a new DataReportingProvisioningSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataReportingProvisioningSession(provisioningSessionId string, aspId string, externalApplicationId string, eventId AfEvent, dataReportingConfigurationIds []string) *DataReportingProvisioningSession {
	this := DataReportingProvisioningSession{}
	this.ProvisioningSessionId = provisioningSessionId
	this.AspId = aspId
	this.ExternalApplicationId = externalApplicationId
	this.EventId = eventId
	this.DataReportingConfigurationIds = dataReportingConfigurationIds
	return &this
}

// NewDataReportingProvisioningSessionWithDefaults instantiates a new DataReportingProvisioningSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataReportingProvisioningSessionWithDefaults() *DataReportingProvisioningSession {
	this := DataReportingProvisioningSession{}
	return &this
}

// GetProvisioningSessionId returns the ProvisioningSessionId field value
func (o *DataReportingProvisioningSession) GetProvisioningSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisioningSessionId
}

// GetProvisioningSessionIdOk returns a tuple with the ProvisioningSessionId field value
// and a boolean to check if the value has been set.
func (o *DataReportingProvisioningSession) GetProvisioningSessionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProvisioningSessionId, true
}

// SetProvisioningSessionId sets field value
func (o *DataReportingProvisioningSession) SetProvisioningSessionId(v string) {
	o.ProvisioningSessionId = v
}

// GetAspId returns the AspId field value
func (o *DataReportingProvisioningSession) GetAspId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AspId
}

// GetAspIdOk returns a tuple with the AspId field value
// and a boolean to check if the value has been set.
func (o *DataReportingProvisioningSession) GetAspIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AspId, true
}

// SetAspId sets field value
func (o *DataReportingProvisioningSession) SetAspId(v string) {
	o.AspId = v
}

// GetExternalApplicationId returns the ExternalApplicationId field value
func (o *DataReportingProvisioningSession) GetExternalApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalApplicationId
}

// GetExternalApplicationIdOk returns a tuple with the ExternalApplicationId field value
// and a boolean to check if the value has been set.
func (o *DataReportingProvisioningSession) GetExternalApplicationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExternalApplicationId, true
}

// SetExternalApplicationId sets field value
func (o *DataReportingProvisioningSession) SetExternalApplicationId(v string) {
	o.ExternalApplicationId = v
}

// GetInternalApplicationId returns the InternalApplicationId field value if set, zero value otherwise.
func (o *DataReportingProvisioningSession) GetInternalApplicationId() string {
	if o == nil || isNil(o.InternalApplicationId) {
		var ret string
		return ret
	}
	return *o.InternalApplicationId
}

// GetInternalApplicationIdOk returns a tuple with the InternalApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReportingProvisioningSession) GetInternalApplicationIdOk() (*string, bool) {
	if o == nil || isNil(o.InternalApplicationId) {
    return nil, false
	}
	return o.InternalApplicationId, true
}

// HasInternalApplicationId returns a boolean if a field has been set.
func (o *DataReportingProvisioningSession) HasInternalApplicationId() bool {
	if o != nil && !isNil(o.InternalApplicationId) {
		return true
	}

	return false
}

// SetInternalApplicationId gets a reference to the given string and assigns it to the InternalApplicationId field.
func (o *DataReportingProvisioningSession) SetInternalApplicationId(v string) {
	o.InternalApplicationId = &v
}

// GetEventId returns the EventId field value
func (o *DataReportingProvisioningSession) GetEventId() AfEvent {
	if o == nil {
		var ret AfEvent
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *DataReportingProvisioningSession) GetEventIdOk() (*AfEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *DataReportingProvisioningSession) SetEventId(v AfEvent) {
	o.EventId = v
}

// GetDataReportingConfigurationIds returns the DataReportingConfigurationIds field value
func (o *DataReportingProvisioningSession) GetDataReportingConfigurationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DataReportingConfigurationIds
}

// GetDataReportingConfigurationIdsOk returns a tuple with the DataReportingConfigurationIds field value
// and a boolean to check if the value has been set.
func (o *DataReportingProvisioningSession) GetDataReportingConfigurationIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DataReportingConfigurationIds, true
}

// SetDataReportingConfigurationIds sets field value
func (o *DataReportingProvisioningSession) SetDataReportingConfigurationIds(v []string) {
	o.DataReportingConfigurationIds = v
}

func (o DataReportingProvisioningSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provisioningSessionId"] = o.ProvisioningSessionId
	}
	if true {
		toSerialize["aspId"] = o.AspId
	}
	if true {
		toSerialize["externalApplicationId"] = o.ExternalApplicationId
	}
	if !isNil(o.InternalApplicationId) {
		toSerialize["internalApplicationId"] = o.InternalApplicationId
	}
	if true {
		toSerialize["eventId"] = o.EventId
	}
	if true {
		toSerialize["dataReportingConfigurationIds"] = o.DataReportingConfigurationIds
	}
	return json.Marshal(toSerialize)
}

type NullableDataReportingProvisioningSession struct {
	value *DataReportingProvisioningSession
	isSet bool
}

func (v NullableDataReportingProvisioningSession) Get() *DataReportingProvisioningSession {
	return v.value
}

func (v *NullableDataReportingProvisioningSession) Set(val *DataReportingProvisioningSession) {
	v.value = val
	v.isSet = true
}

func (v NullableDataReportingProvisioningSession) IsSet() bool {
	return v.isSet
}

func (v *NullableDataReportingProvisioningSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataReportingProvisioningSession(val *DataReportingProvisioningSession) *NullableDataReportingProvisioningSession {
	return &NullableDataReportingProvisioningSession{value: val, isSet: true}
}

func (v NullableDataReportingProvisioningSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataReportingProvisioningSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


