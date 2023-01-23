/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentHostingProvisioning

import (
	"encoding/json"
)

// ContentHostingConfiguration A representation of a Content Hosting Configuration resource.
type ContentHostingConfiguration struct {
	Name string `json:"name"`
	IngestConfiguration IngestConfiguration `json:"ingestConfiguration"`
	DistributionConfigurations []DistributionConfiguration `json:"distributionConfigurations"`
}

// NewContentHostingConfiguration instantiates a new ContentHostingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentHostingConfiguration(name string, ingestConfiguration IngestConfiguration, distributionConfigurations []DistributionConfiguration) *ContentHostingConfiguration {
	this := ContentHostingConfiguration{}
	this.Name = name
	this.IngestConfiguration = ingestConfiguration
	this.DistributionConfigurations = distributionConfigurations
	return &this
}

// NewContentHostingConfigurationWithDefaults instantiates a new ContentHostingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentHostingConfigurationWithDefaults() *ContentHostingConfiguration {
	this := ContentHostingConfiguration{}
	return &this
}

// GetName returns the Name field value
func (o *ContentHostingConfiguration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContentHostingConfiguration) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContentHostingConfiguration) SetName(v string) {
	o.Name = v
}

// GetIngestConfiguration returns the IngestConfiguration field value
func (o *ContentHostingConfiguration) GetIngestConfiguration() IngestConfiguration {
	if o == nil {
		var ret IngestConfiguration
		return ret
	}

	return o.IngestConfiguration
}

// GetIngestConfigurationOk returns a tuple with the IngestConfiguration field value
// and a boolean to check if the value has been set.
func (o *ContentHostingConfiguration) GetIngestConfigurationOk() (*IngestConfiguration, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IngestConfiguration, true
}

// SetIngestConfiguration sets field value
func (o *ContentHostingConfiguration) SetIngestConfiguration(v IngestConfiguration) {
	o.IngestConfiguration = v
}

// GetDistributionConfigurations returns the DistributionConfigurations field value
func (o *ContentHostingConfiguration) GetDistributionConfigurations() []DistributionConfiguration {
	if o == nil {
		var ret []DistributionConfiguration
		return ret
	}

	return o.DistributionConfigurations
}

// GetDistributionConfigurationsOk returns a tuple with the DistributionConfigurations field value
// and a boolean to check if the value has been set.
func (o *ContentHostingConfiguration) GetDistributionConfigurationsOk() ([]DistributionConfiguration, bool) {
	if o == nil {
    return nil, false
	}
	return o.DistributionConfigurations, true
}

// SetDistributionConfigurations sets field value
func (o *ContentHostingConfiguration) SetDistributionConfigurations(v []DistributionConfiguration) {
	o.DistributionConfigurations = v
}

func (o ContentHostingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ingestConfiguration"] = o.IngestConfiguration
	}
	if true {
		toSerialize["distributionConfigurations"] = o.DistributionConfigurations
	}
	return json.Marshal(toSerialize)
}

type NullableContentHostingConfiguration struct {
	value *ContentHostingConfiguration
	isSet bool
}

func (v NullableContentHostingConfiguration) Get() *ContentHostingConfiguration {
	return v.value
}

func (v *NullableContentHostingConfiguration) Set(val *ContentHostingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableContentHostingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableContentHostingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentHostingConfiguration(val *ContentHostingConfiguration) *NullableContentHostingConfiguration {
	return &NullableContentHostingConfiguration{value: val, isSet: true}
}

func (v NullableContentHostingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentHostingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


