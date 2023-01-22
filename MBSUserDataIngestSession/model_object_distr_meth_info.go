/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserDataIngestSession

import (
	"encoding/json"
)

// ObjectDistrMethInfo Represents additional MBS Distribution Session parameters for the case of an Object  Distribution Method. 
type ObjectDistrMethInfo struct {
	OperatingMode *ObjDistributionOperatingMode `json:"operatingMode,omitempty"`
	ObjAcqMethod *ObjAcquisitionMethod `json:"objAcqMethod,omitempty"`
	ObjAcqIds []string `json:"objAcqIds"`
	// String providing an URI formatted according to RFC 3986.
	ObjIngUri *string `json:"objIngUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ObjDistrUri *string `json:"objDistrUri,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ObjRepairUri *string `json:"objRepairUri,omitempty"`
}

// NewObjectDistrMethInfo instantiates a new ObjectDistrMethInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectDistrMethInfo(objAcqIds []string) *ObjectDistrMethInfo {
	this := ObjectDistrMethInfo{}
	this.ObjAcqIds = objAcqIds
	return &this
}

// NewObjectDistrMethInfoWithDefaults instantiates a new ObjectDistrMethInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectDistrMethInfoWithDefaults() *ObjectDistrMethInfo {
	this := ObjectDistrMethInfo{}
	return &this
}

// GetOperatingMode returns the OperatingMode field value if set, zero value otherwise.
func (o *ObjectDistrMethInfo) GetOperatingMode() ObjDistributionOperatingMode {
	if o == nil || isNil(o.OperatingMode) {
		var ret ObjDistributionOperatingMode
		return ret
	}
	return *o.OperatingMode
}

// GetOperatingModeOk returns a tuple with the OperatingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDistrMethInfo) GetOperatingModeOk() (*ObjDistributionOperatingMode, bool) {
	if o == nil || isNil(o.OperatingMode) {
    return nil, false
	}
	return o.OperatingMode, true
}

// HasOperatingMode returns a boolean if a field has been set.
func (o *ObjectDistrMethInfo) HasOperatingMode() bool {
	if o != nil && !isNil(o.OperatingMode) {
		return true
	}

	return false
}

// SetOperatingMode gets a reference to the given ObjDistributionOperatingMode and assigns it to the OperatingMode field.
func (o *ObjectDistrMethInfo) SetOperatingMode(v ObjDistributionOperatingMode) {
	o.OperatingMode = &v
}

// GetObjAcqMethod returns the ObjAcqMethod field value if set, zero value otherwise.
func (o *ObjectDistrMethInfo) GetObjAcqMethod() ObjAcquisitionMethod {
	if o == nil || isNil(o.ObjAcqMethod) {
		var ret ObjAcquisitionMethod
		return ret
	}
	return *o.ObjAcqMethod
}

// GetObjAcqMethodOk returns a tuple with the ObjAcqMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDistrMethInfo) GetObjAcqMethodOk() (*ObjAcquisitionMethod, bool) {
	if o == nil || isNil(o.ObjAcqMethod) {
    return nil, false
	}
	return o.ObjAcqMethod, true
}

// HasObjAcqMethod returns a boolean if a field has been set.
func (o *ObjectDistrMethInfo) HasObjAcqMethod() bool {
	if o != nil && !isNil(o.ObjAcqMethod) {
		return true
	}

	return false
}

// SetObjAcqMethod gets a reference to the given ObjAcquisitionMethod and assigns it to the ObjAcqMethod field.
func (o *ObjectDistrMethInfo) SetObjAcqMethod(v ObjAcquisitionMethod) {
	o.ObjAcqMethod = &v
}

// GetObjAcqIds returns the ObjAcqIds field value
func (o *ObjectDistrMethInfo) GetObjAcqIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjAcqIds
}

// GetObjAcqIdsOk returns a tuple with the ObjAcqIds field value
// and a boolean to check if the value has been set.
func (o *ObjectDistrMethInfo) GetObjAcqIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ObjAcqIds, true
}

// SetObjAcqIds sets field value
func (o *ObjectDistrMethInfo) SetObjAcqIds(v []string) {
	o.ObjAcqIds = v
}

// GetObjIngUri returns the ObjIngUri field value if set, zero value otherwise.
func (o *ObjectDistrMethInfo) GetObjIngUri() string {
	if o == nil || isNil(o.ObjIngUri) {
		var ret string
		return ret
	}
	return *o.ObjIngUri
}

// GetObjIngUriOk returns a tuple with the ObjIngUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDistrMethInfo) GetObjIngUriOk() (*string, bool) {
	if o == nil || isNil(o.ObjIngUri) {
    return nil, false
	}
	return o.ObjIngUri, true
}

// HasObjIngUri returns a boolean if a field has been set.
func (o *ObjectDistrMethInfo) HasObjIngUri() bool {
	if o != nil && !isNil(o.ObjIngUri) {
		return true
	}

	return false
}

// SetObjIngUri gets a reference to the given string and assigns it to the ObjIngUri field.
func (o *ObjectDistrMethInfo) SetObjIngUri(v string) {
	o.ObjIngUri = &v
}

// GetObjDistrUri returns the ObjDistrUri field value if set, zero value otherwise.
func (o *ObjectDistrMethInfo) GetObjDistrUri() string {
	if o == nil || isNil(o.ObjDistrUri) {
		var ret string
		return ret
	}
	return *o.ObjDistrUri
}

// GetObjDistrUriOk returns a tuple with the ObjDistrUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDistrMethInfo) GetObjDistrUriOk() (*string, bool) {
	if o == nil || isNil(o.ObjDistrUri) {
    return nil, false
	}
	return o.ObjDistrUri, true
}

// HasObjDistrUri returns a boolean if a field has been set.
func (o *ObjectDistrMethInfo) HasObjDistrUri() bool {
	if o != nil && !isNil(o.ObjDistrUri) {
		return true
	}

	return false
}

// SetObjDistrUri gets a reference to the given string and assigns it to the ObjDistrUri field.
func (o *ObjectDistrMethInfo) SetObjDistrUri(v string) {
	o.ObjDistrUri = &v
}

// GetObjRepairUri returns the ObjRepairUri field value if set, zero value otherwise.
func (o *ObjectDistrMethInfo) GetObjRepairUri() string {
	if o == nil || isNil(o.ObjRepairUri) {
		var ret string
		return ret
	}
	return *o.ObjRepairUri
}

// GetObjRepairUriOk returns a tuple with the ObjRepairUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDistrMethInfo) GetObjRepairUriOk() (*string, bool) {
	if o == nil || isNil(o.ObjRepairUri) {
    return nil, false
	}
	return o.ObjRepairUri, true
}

// HasObjRepairUri returns a boolean if a field has been set.
func (o *ObjectDistrMethInfo) HasObjRepairUri() bool {
	if o != nil && !isNil(o.ObjRepairUri) {
		return true
	}

	return false
}

// SetObjRepairUri gets a reference to the given string and assigns it to the ObjRepairUri field.
func (o *ObjectDistrMethInfo) SetObjRepairUri(v string) {
	o.ObjRepairUri = &v
}

func (o ObjectDistrMethInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OperatingMode) {
		toSerialize["operatingMode"] = o.OperatingMode
	}
	if !isNil(o.ObjAcqMethod) {
		toSerialize["objAcqMethod"] = o.ObjAcqMethod
	}
	if true {
		toSerialize["objAcqIds"] = o.ObjAcqIds
	}
	if !isNil(o.ObjIngUri) {
		toSerialize["objIngUri"] = o.ObjIngUri
	}
	if !isNil(o.ObjDistrUri) {
		toSerialize["objDistrUri"] = o.ObjDistrUri
	}
	if !isNil(o.ObjRepairUri) {
		toSerialize["objRepairUri"] = o.ObjRepairUri
	}
	return json.Marshal(toSerialize)
}

type NullableObjectDistrMethInfo struct {
	value *ObjectDistrMethInfo
	isSet bool
}

func (v NullableObjectDistrMethInfo) Get() *ObjectDistrMethInfo {
	return v.value
}

func (v *NullableObjectDistrMethInfo) Set(val *ObjectDistrMethInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectDistrMethInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectDistrMethInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectDistrMethInfo(val *ObjectDistrMethInfo) *NullableObjectDistrMethInfo {
	return &NullableObjectDistrMethInfo{value: val, isSet: true}
}

func (v NullableObjectDistrMethInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectDistrMethInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


