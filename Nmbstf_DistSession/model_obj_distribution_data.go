/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbstf_DistSession

import (
	"encoding/json"
)

// ObjDistributionData Info for Object Distribution Method
type ObjDistributionData struct {
	ObjDistributionOperatingMode ObjDistributionOperatingMode `json:"objDistributionOperatingMode"`
	ObjAcquisitionMethod ObjAcquisitionMethod `json:"objAcquisitionMethod"`
	ObjAcquisitionIdsPull []string `json:"objAcquisitionIdsPull,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ObjAcquisitionIdPush *string `json:"objAcquisitionIdPush,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ObjIngestBaseUrl *string `json:"objIngestBaseUrl,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ObjDistributionBaseUrl *string `json:"objDistributionBaseUrl,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ObjRepairBaseUrl *string `json:"objRepairBaseUrl,omitempty"`
}

// NewObjDistributionData instantiates a new ObjDistributionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjDistributionData(objDistributionOperatingMode ObjDistributionOperatingMode, objAcquisitionMethod ObjAcquisitionMethod) *ObjDistributionData {
	this := ObjDistributionData{}
	this.ObjDistributionOperatingMode = objDistributionOperatingMode
	this.ObjAcquisitionMethod = objAcquisitionMethod
	return &this
}

// NewObjDistributionDataWithDefaults instantiates a new ObjDistributionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjDistributionDataWithDefaults() *ObjDistributionData {
	this := ObjDistributionData{}
	return &this
}

// GetObjDistributionOperatingMode returns the ObjDistributionOperatingMode field value
func (o *ObjDistributionData) GetObjDistributionOperatingMode() ObjDistributionOperatingMode {
	if o == nil {
		var ret ObjDistributionOperatingMode
		return ret
	}

	return o.ObjDistributionOperatingMode
}

// GetObjDistributionOperatingModeOk returns a tuple with the ObjDistributionOperatingMode field value
// and a boolean to check if the value has been set.
func (o *ObjDistributionData) GetObjDistributionOperatingModeOk() (*ObjDistributionOperatingMode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ObjDistributionOperatingMode, true
}

// SetObjDistributionOperatingMode sets field value
func (o *ObjDistributionData) SetObjDistributionOperatingMode(v ObjDistributionOperatingMode) {
	o.ObjDistributionOperatingMode = v
}

// GetObjAcquisitionMethod returns the ObjAcquisitionMethod field value
func (o *ObjDistributionData) GetObjAcquisitionMethod() ObjAcquisitionMethod {
	if o == nil {
		var ret ObjAcquisitionMethod
		return ret
	}

	return o.ObjAcquisitionMethod
}

// GetObjAcquisitionMethodOk returns a tuple with the ObjAcquisitionMethod field value
// and a boolean to check if the value has been set.
func (o *ObjDistributionData) GetObjAcquisitionMethodOk() (*ObjAcquisitionMethod, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ObjAcquisitionMethod, true
}

// SetObjAcquisitionMethod sets field value
func (o *ObjDistributionData) SetObjAcquisitionMethod(v ObjAcquisitionMethod) {
	o.ObjAcquisitionMethod = v
}

// GetObjAcquisitionIdsPull returns the ObjAcquisitionIdsPull field value if set, zero value otherwise.
func (o *ObjDistributionData) GetObjAcquisitionIdsPull() []string {
	if o == nil || isNil(o.ObjAcquisitionIdsPull) {
		var ret []string
		return ret
	}
	return o.ObjAcquisitionIdsPull
}

// GetObjAcquisitionIdsPullOk returns a tuple with the ObjAcquisitionIdsPull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjDistributionData) GetObjAcquisitionIdsPullOk() ([]string, bool) {
	if o == nil || isNil(o.ObjAcquisitionIdsPull) {
    return nil, false
	}
	return o.ObjAcquisitionIdsPull, true
}

// HasObjAcquisitionIdsPull returns a boolean if a field has been set.
func (o *ObjDistributionData) HasObjAcquisitionIdsPull() bool {
	if o != nil && !isNil(o.ObjAcquisitionIdsPull) {
		return true
	}

	return false
}

// SetObjAcquisitionIdsPull gets a reference to the given []string and assigns it to the ObjAcquisitionIdsPull field.
func (o *ObjDistributionData) SetObjAcquisitionIdsPull(v []string) {
	o.ObjAcquisitionIdsPull = v
}

// GetObjAcquisitionIdPush returns the ObjAcquisitionIdPush field value if set, zero value otherwise.
func (o *ObjDistributionData) GetObjAcquisitionIdPush() string {
	if o == nil || isNil(o.ObjAcquisitionIdPush) {
		var ret string
		return ret
	}
	return *o.ObjAcquisitionIdPush
}

// GetObjAcquisitionIdPushOk returns a tuple with the ObjAcquisitionIdPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjDistributionData) GetObjAcquisitionIdPushOk() (*string, bool) {
	if o == nil || isNil(o.ObjAcquisitionIdPush) {
    return nil, false
	}
	return o.ObjAcquisitionIdPush, true
}

// HasObjAcquisitionIdPush returns a boolean if a field has been set.
func (o *ObjDistributionData) HasObjAcquisitionIdPush() bool {
	if o != nil && !isNil(o.ObjAcquisitionIdPush) {
		return true
	}

	return false
}

// SetObjAcquisitionIdPush gets a reference to the given string and assigns it to the ObjAcquisitionIdPush field.
func (o *ObjDistributionData) SetObjAcquisitionIdPush(v string) {
	o.ObjAcquisitionIdPush = &v
}

// GetObjIngestBaseUrl returns the ObjIngestBaseUrl field value if set, zero value otherwise.
func (o *ObjDistributionData) GetObjIngestBaseUrl() string {
	if o == nil || isNil(o.ObjIngestBaseUrl) {
		var ret string
		return ret
	}
	return *o.ObjIngestBaseUrl
}

// GetObjIngestBaseUrlOk returns a tuple with the ObjIngestBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjDistributionData) GetObjIngestBaseUrlOk() (*string, bool) {
	if o == nil || isNil(o.ObjIngestBaseUrl) {
    return nil, false
	}
	return o.ObjIngestBaseUrl, true
}

// HasObjIngestBaseUrl returns a boolean if a field has been set.
func (o *ObjDistributionData) HasObjIngestBaseUrl() bool {
	if o != nil && !isNil(o.ObjIngestBaseUrl) {
		return true
	}

	return false
}

// SetObjIngestBaseUrl gets a reference to the given string and assigns it to the ObjIngestBaseUrl field.
func (o *ObjDistributionData) SetObjIngestBaseUrl(v string) {
	o.ObjIngestBaseUrl = &v
}

// GetObjDistributionBaseUrl returns the ObjDistributionBaseUrl field value if set, zero value otherwise.
func (o *ObjDistributionData) GetObjDistributionBaseUrl() string {
	if o == nil || isNil(o.ObjDistributionBaseUrl) {
		var ret string
		return ret
	}
	return *o.ObjDistributionBaseUrl
}

// GetObjDistributionBaseUrlOk returns a tuple with the ObjDistributionBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjDistributionData) GetObjDistributionBaseUrlOk() (*string, bool) {
	if o == nil || isNil(o.ObjDistributionBaseUrl) {
    return nil, false
	}
	return o.ObjDistributionBaseUrl, true
}

// HasObjDistributionBaseUrl returns a boolean if a field has been set.
func (o *ObjDistributionData) HasObjDistributionBaseUrl() bool {
	if o != nil && !isNil(o.ObjDistributionBaseUrl) {
		return true
	}

	return false
}

// SetObjDistributionBaseUrl gets a reference to the given string and assigns it to the ObjDistributionBaseUrl field.
func (o *ObjDistributionData) SetObjDistributionBaseUrl(v string) {
	o.ObjDistributionBaseUrl = &v
}

// GetObjRepairBaseUrl returns the ObjRepairBaseUrl field value if set, zero value otherwise.
func (o *ObjDistributionData) GetObjRepairBaseUrl() string {
	if o == nil || isNil(o.ObjRepairBaseUrl) {
		var ret string
		return ret
	}
	return *o.ObjRepairBaseUrl
}

// GetObjRepairBaseUrlOk returns a tuple with the ObjRepairBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjDistributionData) GetObjRepairBaseUrlOk() (*string, bool) {
	if o == nil || isNil(o.ObjRepairBaseUrl) {
    return nil, false
	}
	return o.ObjRepairBaseUrl, true
}

// HasObjRepairBaseUrl returns a boolean if a field has been set.
func (o *ObjDistributionData) HasObjRepairBaseUrl() bool {
	if o != nil && !isNil(o.ObjRepairBaseUrl) {
		return true
	}

	return false
}

// SetObjRepairBaseUrl gets a reference to the given string and assigns it to the ObjRepairBaseUrl field.
func (o *ObjDistributionData) SetObjRepairBaseUrl(v string) {
	o.ObjRepairBaseUrl = &v
}

func (o ObjDistributionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objDistributionOperatingMode"] = o.ObjDistributionOperatingMode
	}
	if true {
		toSerialize["objAcquisitionMethod"] = o.ObjAcquisitionMethod
	}
	if !isNil(o.ObjAcquisitionIdsPull) {
		toSerialize["objAcquisitionIdsPull"] = o.ObjAcquisitionIdsPull
	}
	if !isNil(o.ObjAcquisitionIdPush) {
		toSerialize["objAcquisitionIdPush"] = o.ObjAcquisitionIdPush
	}
	if !isNil(o.ObjIngestBaseUrl) {
		toSerialize["objIngestBaseUrl"] = o.ObjIngestBaseUrl
	}
	if !isNil(o.ObjDistributionBaseUrl) {
		toSerialize["objDistributionBaseUrl"] = o.ObjDistributionBaseUrl
	}
	if !isNil(o.ObjRepairBaseUrl) {
		toSerialize["objRepairBaseUrl"] = o.ObjRepairBaseUrl
	}
	return json.Marshal(toSerialize)
}

type NullableObjDistributionData struct {
	value *ObjDistributionData
	isSet bool
}

func (v NullableObjDistributionData) Get() *ObjDistributionData {
	return v.value
}

func (v *NullableObjDistributionData) Set(val *ObjDistributionData) {
	v.value = val
	v.isSet = true
}

func (v NullableObjDistributionData) IsSet() bool {
	return v.isSet
}

func (v *NullableObjDistributionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjDistributionData(val *ObjDistributionData) *NullableObjDistributionData {
	return &NullableObjDistributionData{value: val, isSet: true}
}

func (v NullableObjDistributionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjDistributionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


