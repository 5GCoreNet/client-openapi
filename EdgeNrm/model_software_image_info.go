/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package EdgeNrm

import (
	"encoding/json"
)

// SoftwareImageInfo struct for SoftwareImageInfo
type SoftwareImageInfo struct {
	MinimumDisk *int32 `json:"minimumDisk,omitempty"`
	MinimumRAM *int32 `json:"minimumRAM,omitempty"`
	DiscFormat *string `json:"discFormat,omitempty"`
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// indicates the reference to the actual software image that is represented by URL (see clause 7.1.6.5 in ETSI NFV IFA-011 [7]).
	SwImageRef *string `json:"swImageRef,omitempty"`
}

// NewSoftwareImageInfo instantiates a new SoftwareImageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareImageInfo() *SoftwareImageInfo {
	this := SoftwareImageInfo{}
	return &this
}

// NewSoftwareImageInfoWithDefaults instantiates a new SoftwareImageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareImageInfoWithDefaults() *SoftwareImageInfo {
	this := SoftwareImageInfo{}
	return &this
}

// GetMinimumDisk returns the MinimumDisk field value if set, zero value otherwise.
func (o *SoftwareImageInfo) GetMinimumDisk() int32 {
	if o == nil || isNil(o.MinimumDisk) {
		var ret int32
		return ret
	}
	return *o.MinimumDisk
}

// GetMinimumDiskOk returns a tuple with the MinimumDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageInfo) GetMinimumDiskOk() (*int32, bool) {
	if o == nil || isNil(o.MinimumDisk) {
    return nil, false
	}
	return o.MinimumDisk, true
}

// HasMinimumDisk returns a boolean if a field has been set.
func (o *SoftwareImageInfo) HasMinimumDisk() bool {
	if o != nil && !isNil(o.MinimumDisk) {
		return true
	}

	return false
}

// SetMinimumDisk gets a reference to the given int32 and assigns it to the MinimumDisk field.
func (o *SoftwareImageInfo) SetMinimumDisk(v int32) {
	o.MinimumDisk = &v
}

// GetMinimumRAM returns the MinimumRAM field value if set, zero value otherwise.
func (o *SoftwareImageInfo) GetMinimumRAM() int32 {
	if o == nil || isNil(o.MinimumRAM) {
		var ret int32
		return ret
	}
	return *o.MinimumRAM
}

// GetMinimumRAMOk returns a tuple with the MinimumRAM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageInfo) GetMinimumRAMOk() (*int32, bool) {
	if o == nil || isNil(o.MinimumRAM) {
    return nil, false
	}
	return o.MinimumRAM, true
}

// HasMinimumRAM returns a boolean if a field has been set.
func (o *SoftwareImageInfo) HasMinimumRAM() bool {
	if o != nil && !isNil(o.MinimumRAM) {
		return true
	}

	return false
}

// SetMinimumRAM gets a reference to the given int32 and assigns it to the MinimumRAM field.
func (o *SoftwareImageInfo) SetMinimumRAM(v int32) {
	o.MinimumRAM = &v
}

// GetDiscFormat returns the DiscFormat field value if set, zero value otherwise.
func (o *SoftwareImageInfo) GetDiscFormat() string {
	if o == nil || isNil(o.DiscFormat) {
		var ret string
		return ret
	}
	return *o.DiscFormat
}

// GetDiscFormatOk returns a tuple with the DiscFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageInfo) GetDiscFormatOk() (*string, bool) {
	if o == nil || isNil(o.DiscFormat) {
    return nil, false
	}
	return o.DiscFormat, true
}

// HasDiscFormat returns a boolean if a field has been set.
func (o *SoftwareImageInfo) HasDiscFormat() bool {
	if o != nil && !isNil(o.DiscFormat) {
		return true
	}

	return false
}

// SetDiscFormat gets a reference to the given string and assigns it to the DiscFormat field.
func (o *SoftwareImageInfo) SetDiscFormat(v string) {
	o.DiscFormat = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *SoftwareImageInfo) GetOperatingSystem() string {
	if o == nil || isNil(o.OperatingSystem) {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil || isNil(o.OperatingSystem) {
    return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *SoftwareImageInfo) HasOperatingSystem() bool {
	if o != nil && !isNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *SoftwareImageInfo) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetSwImageRef returns the SwImageRef field value if set, zero value otherwise.
func (o *SoftwareImageInfo) GetSwImageRef() string {
	if o == nil || isNil(o.SwImageRef) {
		var ret string
		return ret
	}
	return *o.SwImageRef
}

// GetSwImageRefOk returns a tuple with the SwImageRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageInfo) GetSwImageRefOk() (*string, bool) {
	if o == nil || isNil(o.SwImageRef) {
    return nil, false
	}
	return o.SwImageRef, true
}

// HasSwImageRef returns a boolean if a field has been set.
func (o *SoftwareImageInfo) HasSwImageRef() bool {
	if o != nil && !isNil(o.SwImageRef) {
		return true
	}

	return false
}

// SetSwImageRef gets a reference to the given string and assigns it to the SwImageRef field.
func (o *SoftwareImageInfo) SetSwImageRef(v string) {
	o.SwImageRef = &v
}

func (o SoftwareImageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinimumDisk) {
		toSerialize["minimumDisk"] = o.MinimumDisk
	}
	if !isNil(o.MinimumRAM) {
		toSerialize["minimumRAM"] = o.MinimumRAM
	}
	if !isNil(o.DiscFormat) {
		toSerialize["discFormat"] = o.DiscFormat
	}
	if !isNil(o.OperatingSystem) {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if !isNil(o.SwImageRef) {
		toSerialize["swImageRef"] = o.SwImageRef
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwareImageInfo struct {
	value *SoftwareImageInfo
	isSet bool
}

func (v NullableSoftwareImageInfo) Get() *SoftwareImageInfo {
	return v.value
}

func (v *NullableSoftwareImageInfo) Set(val *SoftwareImageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareImageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareImageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareImageInfo(val *SoftwareImageInfo) *NullableSoftwareImageInfo {
	return &NullableSoftwareImageInfo{value: val, isSet: true}
}

func (v NullableSoftwareImageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareImageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


