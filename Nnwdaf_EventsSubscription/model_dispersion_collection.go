/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_EventsSubscription

import (
	"encoding/json"
)

// DispersionCollection Dispersion collection per UE location or per slice.
type DispersionCollection struct {
	UeLoc *UserLocation `json:"ueLoc,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	Supis []string `json:"supis,omitempty"`
	Gpsis []string `json:"gpsis,omitempty"`
	AppVolumes []ApplicationVolume `json:"appVolumes,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DisperAmount *int32 `json:"disperAmount,omitempty"`
	DisperClass *DispersionClass `json:"disperClass,omitempty"`
	// Integer where the allowed values correspond to 1, 2, 3 only.
	UsageRank *int32 `json:"usageRank,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	PercentileRank *int32 `json:"percentileRank,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	UeRatio *int32 `json:"ueRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// NewDispersionCollection instantiates a new DispersionCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDispersionCollection() *DispersionCollection {
	this := DispersionCollection{}
	return &this
}

// NewDispersionCollectionWithDefaults instantiates a new DispersionCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDispersionCollectionWithDefaults() *DispersionCollection {
	this := DispersionCollection{}
	return &this
}

// GetUeLoc returns the UeLoc field value if set, zero value otherwise.
func (o *DispersionCollection) GetUeLoc() UserLocation {
	if o == nil || isNil(o.UeLoc) {
		var ret UserLocation
		return ret
	}
	return *o.UeLoc
}

// GetUeLocOk returns a tuple with the UeLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetUeLocOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UeLoc) {
    return nil, false
	}
	return o.UeLoc, true
}

// HasUeLoc returns a boolean if a field has been set.
func (o *DispersionCollection) HasUeLoc() bool {
	if o != nil && !isNil(o.UeLoc) {
		return true
	}

	return false
}

// SetUeLoc gets a reference to the given UserLocation and assigns it to the UeLoc field.
func (o *DispersionCollection) SetUeLoc(v UserLocation) {
	o.UeLoc = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *DispersionCollection) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *DispersionCollection) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *DispersionCollection) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *DispersionCollection) GetSupis() []string {
	if o == nil || isNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetSupisOk() ([]string, bool) {
	if o == nil || isNil(o.Supis) {
    return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *DispersionCollection) HasSupis() bool {
	if o != nil && !isNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *DispersionCollection) SetSupis(v []string) {
	o.Supis = v
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *DispersionCollection) GetGpsis() []string {
	if o == nil || isNil(o.Gpsis) {
		var ret []string
		return ret
	}
	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetGpsisOk() ([]string, bool) {
	if o == nil || isNil(o.Gpsis) {
    return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *DispersionCollection) HasGpsis() bool {
	if o != nil && !isNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given []string and assigns it to the Gpsis field.
func (o *DispersionCollection) SetGpsis(v []string) {
	o.Gpsis = v
}

// GetAppVolumes returns the AppVolumes field value if set, zero value otherwise.
func (o *DispersionCollection) GetAppVolumes() []ApplicationVolume {
	if o == nil || isNil(o.AppVolumes) {
		var ret []ApplicationVolume
		return ret
	}
	return o.AppVolumes
}

// GetAppVolumesOk returns a tuple with the AppVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetAppVolumesOk() ([]ApplicationVolume, bool) {
	if o == nil || isNil(o.AppVolumes) {
    return nil, false
	}
	return o.AppVolumes, true
}

// HasAppVolumes returns a boolean if a field has been set.
func (o *DispersionCollection) HasAppVolumes() bool {
	if o != nil && !isNil(o.AppVolumes) {
		return true
	}

	return false
}

// SetAppVolumes gets a reference to the given []ApplicationVolume and assigns it to the AppVolumes field.
func (o *DispersionCollection) SetAppVolumes(v []ApplicationVolume) {
	o.AppVolumes = v
}

// GetDisperAmount returns the DisperAmount field value if set, zero value otherwise.
func (o *DispersionCollection) GetDisperAmount() int32 {
	if o == nil || isNil(o.DisperAmount) {
		var ret int32
		return ret
	}
	return *o.DisperAmount
}

// GetDisperAmountOk returns a tuple with the DisperAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetDisperAmountOk() (*int32, bool) {
	if o == nil || isNil(o.DisperAmount) {
    return nil, false
	}
	return o.DisperAmount, true
}

// HasDisperAmount returns a boolean if a field has been set.
func (o *DispersionCollection) HasDisperAmount() bool {
	if o != nil && !isNil(o.DisperAmount) {
		return true
	}

	return false
}

// SetDisperAmount gets a reference to the given int32 and assigns it to the DisperAmount field.
func (o *DispersionCollection) SetDisperAmount(v int32) {
	o.DisperAmount = &v
}

// GetDisperClass returns the DisperClass field value if set, zero value otherwise.
func (o *DispersionCollection) GetDisperClass() DispersionClass {
	if o == nil || isNil(o.DisperClass) {
		var ret DispersionClass
		return ret
	}
	return *o.DisperClass
}

// GetDisperClassOk returns a tuple with the DisperClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetDisperClassOk() (*DispersionClass, bool) {
	if o == nil || isNil(o.DisperClass) {
    return nil, false
	}
	return o.DisperClass, true
}

// HasDisperClass returns a boolean if a field has been set.
func (o *DispersionCollection) HasDisperClass() bool {
	if o != nil && !isNil(o.DisperClass) {
		return true
	}

	return false
}

// SetDisperClass gets a reference to the given DispersionClass and assigns it to the DisperClass field.
func (o *DispersionCollection) SetDisperClass(v DispersionClass) {
	o.DisperClass = &v
}

// GetUsageRank returns the UsageRank field value if set, zero value otherwise.
func (o *DispersionCollection) GetUsageRank() int32 {
	if o == nil || isNil(o.UsageRank) {
		var ret int32
		return ret
	}
	return *o.UsageRank
}

// GetUsageRankOk returns a tuple with the UsageRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetUsageRankOk() (*int32, bool) {
	if o == nil || isNil(o.UsageRank) {
    return nil, false
	}
	return o.UsageRank, true
}

// HasUsageRank returns a boolean if a field has been set.
func (o *DispersionCollection) HasUsageRank() bool {
	if o != nil && !isNil(o.UsageRank) {
		return true
	}

	return false
}

// SetUsageRank gets a reference to the given int32 and assigns it to the UsageRank field.
func (o *DispersionCollection) SetUsageRank(v int32) {
	o.UsageRank = &v
}

// GetPercentileRank returns the PercentileRank field value if set, zero value otherwise.
func (o *DispersionCollection) GetPercentileRank() int32 {
	if o == nil || isNil(o.PercentileRank) {
		var ret int32
		return ret
	}
	return *o.PercentileRank
}

// GetPercentileRankOk returns a tuple with the PercentileRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetPercentileRankOk() (*int32, bool) {
	if o == nil || isNil(o.PercentileRank) {
    return nil, false
	}
	return o.PercentileRank, true
}

// HasPercentileRank returns a boolean if a field has been set.
func (o *DispersionCollection) HasPercentileRank() bool {
	if o != nil && !isNil(o.PercentileRank) {
		return true
	}

	return false
}

// SetPercentileRank gets a reference to the given int32 and assigns it to the PercentileRank field.
func (o *DispersionCollection) SetPercentileRank(v int32) {
	o.PercentileRank = &v
}

// GetUeRatio returns the UeRatio field value if set, zero value otherwise.
func (o *DispersionCollection) GetUeRatio() int32 {
	if o == nil || isNil(o.UeRatio) {
		var ret int32
		return ret
	}
	return *o.UeRatio
}

// GetUeRatioOk returns a tuple with the UeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetUeRatioOk() (*int32, bool) {
	if o == nil || isNil(o.UeRatio) {
    return nil, false
	}
	return o.UeRatio, true
}

// HasUeRatio returns a boolean if a field has been set.
func (o *DispersionCollection) HasUeRatio() bool {
	if o != nil && !isNil(o.UeRatio) {
		return true
	}

	return false
}

// SetUeRatio gets a reference to the given int32 and assigns it to the UeRatio field.
func (o *DispersionCollection) SetUeRatio(v int32) {
	o.UeRatio = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *DispersionCollection) GetConfidence() int32 {
	if o == nil || isNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispersionCollection) GetConfidenceOk() (*int32, bool) {
	if o == nil || isNil(o.Confidence) {
    return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *DispersionCollection) HasConfidence() bool {
	if o != nil && !isNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *DispersionCollection) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o DispersionCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UeLoc) {
		toSerialize["ueLoc"] = o.UeLoc
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !isNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !isNil(o.AppVolumes) {
		toSerialize["appVolumes"] = o.AppVolumes
	}
	if !isNil(o.DisperAmount) {
		toSerialize["disperAmount"] = o.DisperAmount
	}
	if !isNil(o.DisperClass) {
		toSerialize["disperClass"] = o.DisperClass
	}
	if !isNil(o.UsageRank) {
		toSerialize["usageRank"] = o.UsageRank
	}
	if !isNil(o.PercentileRank) {
		toSerialize["percentileRank"] = o.PercentileRank
	}
	if !isNil(o.UeRatio) {
		toSerialize["ueRatio"] = o.UeRatio
	}
	if !isNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullableDispersionCollection struct {
	value *DispersionCollection
	isSet bool
}

func (v NullableDispersionCollection) Get() *DispersionCollection {
	return v.value
}

func (v *NullableDispersionCollection) Set(val *DispersionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionCollection(val *DispersionCollection) *NullableDispersionCollection {
	return &NullableDispersionCollection{value: val, isSet: true}
}

func (v NullableDispersionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


