/*
M5_ServiceAccessInformation

5GMS AF M5 Service Access Information API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_ServiceAccessInformation

import (
	"encoding/json"
)

// ServiceAccessInformationResourceClientConsumptionReportingConfiguration struct for ServiceAccessInformationResourceClientConsumptionReportingConfiguration
type ServiceAccessInformationResourceClientConsumptionReportingConfiguration struct {
	// indicating a time in seconds.
	ReportingInterval *int32 `json:"reportingInterval,omitempty"`
	// A set of application endpoint addresses.
	ServerAddresses []string `json:"serverAddresses"`
	LocationReporting bool `json:"locationReporting"`
	AccessReporting *bool `json:"accessReporting,omitempty"`
	SamplePercentage float32 `json:"samplePercentage"`
}

// NewServiceAccessInformationResourceClientConsumptionReportingConfiguration instantiates a new ServiceAccessInformationResourceClientConsumptionReportingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccessInformationResourceClientConsumptionReportingConfiguration(serverAddresses []string, locationReporting bool, samplePercentage float32) *ServiceAccessInformationResourceClientConsumptionReportingConfiguration {
	this := ServiceAccessInformationResourceClientConsumptionReportingConfiguration{}
	this.ServerAddresses = serverAddresses
	this.LocationReporting = locationReporting
	this.SamplePercentage = samplePercentage
	return &this
}

// NewServiceAccessInformationResourceClientConsumptionReportingConfigurationWithDefaults instantiates a new ServiceAccessInformationResourceClientConsumptionReportingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccessInformationResourceClientConsumptionReportingConfigurationWithDefaults() *ServiceAccessInformationResourceClientConsumptionReportingConfiguration {
	this := ServiceAccessInformationResourceClientConsumptionReportingConfiguration{}
	return &this
}

// GetReportingInterval returns the ReportingInterval field value if set, zero value otherwise.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetReportingInterval() int32 {
	if o == nil || isNil(o.ReportingInterval) {
		var ret int32
		return ret
	}
	return *o.ReportingInterval
}

// GetReportingIntervalOk returns a tuple with the ReportingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetReportingIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.ReportingInterval) {
    return nil, false
	}
	return o.ReportingInterval, true
}

// HasReportingInterval returns a boolean if a field has been set.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) HasReportingInterval() bool {
	if o != nil && !isNil(o.ReportingInterval) {
		return true
	}

	return false
}

// SetReportingInterval gets a reference to the given int32 and assigns it to the ReportingInterval field.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetReportingInterval(v int32) {
	o.ReportingInterval = &v
}

// GetServerAddresses returns the ServerAddresses field value
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetServerAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServerAddresses
}

// GetServerAddressesOk returns a tuple with the ServerAddresses field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetServerAddressesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServerAddresses, true
}

// SetServerAddresses sets field value
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetServerAddresses(v []string) {
	o.ServerAddresses = v
}

// GetLocationReporting returns the LocationReporting field value
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetLocationReporting() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LocationReporting
}

// GetLocationReportingOk returns a tuple with the LocationReporting field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetLocationReportingOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocationReporting, true
}

// SetLocationReporting sets field value
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetLocationReporting(v bool) {
	o.LocationReporting = v
}

// GetAccessReporting returns the AccessReporting field value if set, zero value otherwise.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetAccessReporting() bool {
	if o == nil || isNil(o.AccessReporting) {
		var ret bool
		return ret
	}
	return *o.AccessReporting
}

// GetAccessReportingOk returns a tuple with the AccessReporting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetAccessReportingOk() (*bool, bool) {
	if o == nil || isNil(o.AccessReporting) {
    return nil, false
	}
	return o.AccessReporting, true
}

// HasAccessReporting returns a boolean if a field has been set.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) HasAccessReporting() bool {
	if o != nil && !isNil(o.AccessReporting) {
		return true
	}

	return false
}

// SetAccessReporting gets a reference to the given bool and assigns it to the AccessReporting field.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetAccessReporting(v bool) {
	o.AccessReporting = &v
}

// GetSamplePercentage returns the SamplePercentage field value
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetSamplePercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SamplePercentage
}

// GetSamplePercentageOk returns a tuple with the SamplePercentage field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetSamplePercentageOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SamplePercentage, true
}

// SetSamplePercentage sets field value
func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetSamplePercentage(v float32) {
	o.SamplePercentage = v
}

func (o ServiceAccessInformationResourceClientConsumptionReportingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReportingInterval) {
		toSerialize["reportingInterval"] = o.ReportingInterval
	}
	if true {
		toSerialize["serverAddresses"] = o.ServerAddresses
	}
	if true {
		toSerialize["locationReporting"] = o.LocationReporting
	}
	if !isNil(o.AccessReporting) {
		toSerialize["accessReporting"] = o.AccessReporting
	}
	if true {
		toSerialize["samplePercentage"] = o.SamplePercentage
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration struct {
	value *ServiceAccessInformationResourceClientConsumptionReportingConfiguration
	isSet bool
}

func (v NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration) Get() *ServiceAccessInformationResourceClientConsumptionReportingConfiguration {
	return v.value
}

func (v *NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration) Set(val *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccessInformationResourceClientConsumptionReportingConfiguration(val *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) *NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration {
	return &NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration{value: val, isSet: true}
}

func (v NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccessInformationResourceClientConsumptionReportingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


