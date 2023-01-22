/*
M1_EdgeResourcesProvisioning

5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_EdgeResourcesProvisioning

import (
	"encoding/json"
)

// EdgeResourcesConfiguration A representation of an Edge Resources Configuration resource.
type EdgeResourcesConfiguration struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	EdgeResourcesConfigurationId string `json:"edgeResourcesConfigurationId"`
	EdgeManagementMode EdgeManagementMode `json:"edgeManagementMode"`
	EligibilityCriteria *EdgeProcessingEligibilityCriteria `json:"eligibilityCriteria,omitempty"`
	EasRequirements EASRequirements `json:"easRequirements"`
	EasRelocationRequirements *M1EASRelocationRequirements `json:"easRelocationRequirements,omitempty"`
}

// NewEdgeResourcesConfiguration instantiates a new EdgeResourcesConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeResourcesConfiguration(edgeResourcesConfigurationId string, edgeManagementMode EdgeManagementMode, easRequirements EASRequirements) *EdgeResourcesConfiguration {
	this := EdgeResourcesConfiguration{}
	this.EdgeResourcesConfigurationId = edgeResourcesConfigurationId
	this.EdgeManagementMode = edgeManagementMode
	this.EasRequirements = easRequirements
	return &this
}

// NewEdgeResourcesConfigurationWithDefaults instantiates a new EdgeResourcesConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeResourcesConfigurationWithDefaults() *EdgeResourcesConfiguration {
	this := EdgeResourcesConfiguration{}
	return &this
}

// GetEdgeResourcesConfigurationId returns the EdgeResourcesConfigurationId field value
func (o *EdgeResourcesConfiguration) GetEdgeResourcesConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EdgeResourcesConfigurationId
}

// GetEdgeResourcesConfigurationIdOk returns a tuple with the EdgeResourcesConfigurationId field value
// and a boolean to check if the value has been set.
func (o *EdgeResourcesConfiguration) GetEdgeResourcesConfigurationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EdgeResourcesConfigurationId, true
}

// SetEdgeResourcesConfigurationId sets field value
func (o *EdgeResourcesConfiguration) SetEdgeResourcesConfigurationId(v string) {
	o.EdgeResourcesConfigurationId = v
}

// GetEdgeManagementMode returns the EdgeManagementMode field value
func (o *EdgeResourcesConfiguration) GetEdgeManagementMode() EdgeManagementMode {
	if o == nil {
		var ret EdgeManagementMode
		return ret
	}

	return o.EdgeManagementMode
}

// GetEdgeManagementModeOk returns a tuple with the EdgeManagementMode field value
// and a boolean to check if the value has been set.
func (o *EdgeResourcesConfiguration) GetEdgeManagementModeOk() (*EdgeManagementMode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EdgeManagementMode, true
}

// SetEdgeManagementMode sets field value
func (o *EdgeResourcesConfiguration) SetEdgeManagementMode(v EdgeManagementMode) {
	o.EdgeManagementMode = v
}

// GetEligibilityCriteria returns the EligibilityCriteria field value if set, zero value otherwise.
func (o *EdgeResourcesConfiguration) GetEligibilityCriteria() EdgeProcessingEligibilityCriteria {
	if o == nil || isNil(o.EligibilityCriteria) {
		var ret EdgeProcessingEligibilityCriteria
		return ret
	}
	return *o.EligibilityCriteria
}

// GetEligibilityCriteriaOk returns a tuple with the EligibilityCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeResourcesConfiguration) GetEligibilityCriteriaOk() (*EdgeProcessingEligibilityCriteria, bool) {
	if o == nil || isNil(o.EligibilityCriteria) {
    return nil, false
	}
	return o.EligibilityCriteria, true
}

// HasEligibilityCriteria returns a boolean if a field has been set.
func (o *EdgeResourcesConfiguration) HasEligibilityCriteria() bool {
	if o != nil && !isNil(o.EligibilityCriteria) {
		return true
	}

	return false
}

// SetEligibilityCriteria gets a reference to the given EdgeProcessingEligibilityCriteria and assigns it to the EligibilityCriteria field.
func (o *EdgeResourcesConfiguration) SetEligibilityCriteria(v EdgeProcessingEligibilityCriteria) {
	o.EligibilityCriteria = &v
}

// GetEasRequirements returns the EasRequirements field value
func (o *EdgeResourcesConfiguration) GetEasRequirements() EASRequirements {
	if o == nil {
		var ret EASRequirements
		return ret
	}

	return o.EasRequirements
}

// GetEasRequirementsOk returns a tuple with the EasRequirements field value
// and a boolean to check if the value has been set.
func (o *EdgeResourcesConfiguration) GetEasRequirementsOk() (*EASRequirements, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EasRequirements, true
}

// SetEasRequirements sets field value
func (o *EdgeResourcesConfiguration) SetEasRequirements(v EASRequirements) {
	o.EasRequirements = v
}

// GetEasRelocationRequirements returns the EasRelocationRequirements field value if set, zero value otherwise.
func (o *EdgeResourcesConfiguration) GetEasRelocationRequirements() M1EASRelocationRequirements {
	if o == nil || isNil(o.EasRelocationRequirements) {
		var ret M1EASRelocationRequirements
		return ret
	}
	return *o.EasRelocationRequirements
}

// GetEasRelocationRequirementsOk returns a tuple with the EasRelocationRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeResourcesConfiguration) GetEasRelocationRequirementsOk() (*M1EASRelocationRequirements, bool) {
	if o == nil || isNil(o.EasRelocationRequirements) {
    return nil, false
	}
	return o.EasRelocationRequirements, true
}

// HasEasRelocationRequirements returns a boolean if a field has been set.
func (o *EdgeResourcesConfiguration) HasEasRelocationRequirements() bool {
	if o != nil && !isNil(o.EasRelocationRequirements) {
		return true
	}

	return false
}

// SetEasRelocationRequirements gets a reference to the given M1EASRelocationRequirements and assigns it to the EasRelocationRequirements field.
func (o *EdgeResourcesConfiguration) SetEasRelocationRequirements(v M1EASRelocationRequirements) {
	o.EasRelocationRequirements = &v
}

func (o EdgeResourcesConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["edgeResourcesConfigurationId"] = o.EdgeResourcesConfigurationId
	}
	if true {
		toSerialize["edgeManagementMode"] = o.EdgeManagementMode
	}
	if !isNil(o.EligibilityCriteria) {
		toSerialize["eligibilityCriteria"] = o.EligibilityCriteria
	}
	if true {
		toSerialize["easRequirements"] = o.EasRequirements
	}
	if !isNil(o.EasRelocationRequirements) {
		toSerialize["easRelocationRequirements"] = o.EasRelocationRequirements
	}
	return json.Marshal(toSerialize)
}

type NullableEdgeResourcesConfiguration struct {
	value *EdgeResourcesConfiguration
	isSet bool
}

func (v NullableEdgeResourcesConfiguration) Get() *EdgeResourcesConfiguration {
	return v.value
}

func (v *NullableEdgeResourcesConfiguration) Set(val *EdgeResourcesConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeResourcesConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeResourcesConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeResourcesConfiguration(val *EdgeResourcesConfiguration) *NullableEdgeResourcesConfiguration {
	return &NullableEdgeResourcesConfiguration{value: val, isSet: true}
}

func (v NullableEdgeResourcesConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeResourcesConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


