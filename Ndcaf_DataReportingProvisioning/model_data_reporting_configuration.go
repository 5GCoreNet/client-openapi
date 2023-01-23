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

// DataReportingConfiguration A Data Reporting Configuration subresource.
type DataReportingConfiguration struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	DataReportingConfigurationId string `json:"dataReportingConfigurationId"`
	DataCollectionClientType DataCollectionClientType `json:"dataCollectionClientType"`
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	AuthorizationURL *string `json:"authorizationURL,omitempty"`
	DataAccessProfiles []DataAccessProfile `json:"dataAccessProfiles"`
}

// NewDataReportingConfiguration instantiates a new DataReportingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataReportingConfiguration(dataReportingConfigurationId string, dataCollectionClientType DataCollectionClientType, dataAccessProfiles []DataAccessProfile) *DataReportingConfiguration {
	this := DataReportingConfiguration{}
	this.DataReportingConfigurationId = dataReportingConfigurationId
	this.DataCollectionClientType = dataCollectionClientType
	this.DataAccessProfiles = dataAccessProfiles
	return &this
}

// NewDataReportingConfigurationWithDefaults instantiates a new DataReportingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataReportingConfigurationWithDefaults() *DataReportingConfiguration {
	this := DataReportingConfiguration{}
	return &this
}

// GetDataReportingConfigurationId returns the DataReportingConfigurationId field value
func (o *DataReportingConfiguration) GetDataReportingConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataReportingConfigurationId
}

// GetDataReportingConfigurationIdOk returns a tuple with the DataReportingConfigurationId field value
// and a boolean to check if the value has been set.
func (o *DataReportingConfiguration) GetDataReportingConfigurationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataReportingConfigurationId, true
}

// SetDataReportingConfigurationId sets field value
func (o *DataReportingConfiguration) SetDataReportingConfigurationId(v string) {
	o.DataReportingConfigurationId = v
}

// GetDataCollectionClientType returns the DataCollectionClientType field value
func (o *DataReportingConfiguration) GetDataCollectionClientType() DataCollectionClientType {
	if o == nil {
		var ret DataCollectionClientType
		return ret
	}

	return o.DataCollectionClientType
}

// GetDataCollectionClientTypeOk returns a tuple with the DataCollectionClientType field value
// and a boolean to check if the value has been set.
func (o *DataReportingConfiguration) GetDataCollectionClientTypeOk() (*DataCollectionClientType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataCollectionClientType, true
}

// SetDataCollectionClientType sets field value
func (o *DataReportingConfiguration) SetDataCollectionClientType(v DataCollectionClientType) {
	o.DataCollectionClientType = v
}

// GetAuthorizationURL returns the AuthorizationURL field value if set, zero value otherwise.
func (o *DataReportingConfiguration) GetAuthorizationURL() string {
	if o == nil || isNil(o.AuthorizationURL) {
		var ret string
		return ret
	}
	return *o.AuthorizationURL
}

// GetAuthorizationURLOk returns a tuple with the AuthorizationURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReportingConfiguration) GetAuthorizationURLOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationURL) {
    return nil, false
	}
	return o.AuthorizationURL, true
}

// HasAuthorizationURL returns a boolean if a field has been set.
func (o *DataReportingConfiguration) HasAuthorizationURL() bool {
	if o != nil && !isNil(o.AuthorizationURL) {
		return true
	}

	return false
}

// SetAuthorizationURL gets a reference to the given string and assigns it to the AuthorizationURL field.
func (o *DataReportingConfiguration) SetAuthorizationURL(v string) {
	o.AuthorizationURL = &v
}

// GetDataAccessProfiles returns the DataAccessProfiles field value
func (o *DataReportingConfiguration) GetDataAccessProfiles() []DataAccessProfile {
	if o == nil {
		var ret []DataAccessProfile
		return ret
	}

	return o.DataAccessProfiles
}

// GetDataAccessProfilesOk returns a tuple with the DataAccessProfiles field value
// and a boolean to check if the value has been set.
func (o *DataReportingConfiguration) GetDataAccessProfilesOk() ([]DataAccessProfile, bool) {
	if o == nil {
    return nil, false
	}
	return o.DataAccessProfiles, true
}

// SetDataAccessProfiles sets field value
func (o *DataReportingConfiguration) SetDataAccessProfiles(v []DataAccessProfile) {
	o.DataAccessProfiles = v
}

func (o DataReportingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataReportingConfigurationId"] = o.DataReportingConfigurationId
	}
	if true {
		toSerialize["dataCollectionClientType"] = o.DataCollectionClientType
	}
	if !isNil(o.AuthorizationURL) {
		toSerialize["authorizationURL"] = o.AuthorizationURL
	}
	if true {
		toSerialize["dataAccessProfiles"] = o.DataAccessProfiles
	}
	return json.Marshal(toSerialize)
}

type NullableDataReportingConfiguration struct {
	value *DataReportingConfiguration
	isSet bool
}

func (v NullableDataReportingConfiguration) Get() *DataReportingConfiguration {
	return v.value
}

func (v *NullableDataReportingConfiguration) Set(val *DataReportingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDataReportingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDataReportingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataReportingConfiguration(val *DataReportingConfiguration) *NullableDataReportingConfiguration {
	return &NullableDataReportingConfiguration{value: val, isSet: true}
}

func (v NullableDataReportingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataReportingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


