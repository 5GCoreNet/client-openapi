/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Application_Data

import (
	"encoding/json"
)

// DataFilter Identifies a data filter.
type DataFilter struct {
	DataInd DataInd `json:"dataInd"`
	Dnns []string `json:"dnns,omitempty"`
	Snssais []Snssai `json:"snssais,omitempty"`
	InternalGroupIds []string `json:"internalGroupIds,omitempty"`
	Supis []string `json:"supis,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
	UeIpv4s []string `json:"ueIpv4s,omitempty"`
	UeIpv6s []Ipv6Addr `json:"ueIpv6s,omitempty"`
	UeMacs []string `json:"ueMacs,omitempty"`
	// Indicates the request is for any UE.
	AnyUeInd *bool `json:"anyUeInd,omitempty"`
	// Indicates the request is for any DNN and S-NSSAI combination present in the array.
	DnnSnssaiInfos []DnnSnssaiInformation `json:"dnnSnssaiInfos,omitempty"`
}

// NewDataFilter instantiates a new DataFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataFilter(dataInd DataInd) *DataFilter {
	this := DataFilter{}
	this.DataInd = dataInd
	return &this
}

// NewDataFilterWithDefaults instantiates a new DataFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataFilterWithDefaults() *DataFilter {
	this := DataFilter{}
	return &this
}

// GetDataInd returns the DataInd field value
func (o *DataFilter) GetDataInd() DataInd {
	if o == nil {
		var ret DataInd
		return ret
	}

	return o.DataInd
}

// GetDataIndOk returns a tuple with the DataInd field value
// and a boolean to check if the value has been set.
func (o *DataFilter) GetDataIndOk() (*DataInd, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataInd, true
}

// SetDataInd sets field value
func (o *DataFilter) SetDataInd(v DataInd) {
	o.DataInd = v
}

// GetDnns returns the Dnns field value if set, zero value otherwise.
func (o *DataFilter) GetDnns() []string {
	if o == nil || isNil(o.Dnns) {
		var ret []string
		return ret
	}
	return o.Dnns
}

// GetDnnsOk returns a tuple with the Dnns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetDnnsOk() ([]string, bool) {
	if o == nil || isNil(o.Dnns) {
    return nil, false
	}
	return o.Dnns, true
}

// HasDnns returns a boolean if a field has been set.
func (o *DataFilter) HasDnns() bool {
	if o != nil && !isNil(o.Dnns) {
		return true
	}

	return false
}

// SetDnns gets a reference to the given []string and assigns it to the Dnns field.
func (o *DataFilter) SetDnns(v []string) {
	o.Dnns = v
}

// GetSnssais returns the Snssais field value if set, zero value otherwise.
func (o *DataFilter) GetSnssais() []Snssai {
	if o == nil || isNil(o.Snssais) {
		var ret []Snssai
		return ret
	}
	return o.Snssais
}

// GetSnssaisOk returns a tuple with the Snssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetSnssaisOk() ([]Snssai, bool) {
	if o == nil || isNil(o.Snssais) {
    return nil, false
	}
	return o.Snssais, true
}

// HasSnssais returns a boolean if a field has been set.
func (o *DataFilter) HasSnssais() bool {
	if o != nil && !isNil(o.Snssais) {
		return true
	}

	return false
}

// SetSnssais gets a reference to the given []Snssai and assigns it to the Snssais field.
func (o *DataFilter) SetSnssais(v []Snssai) {
	o.Snssais = v
}

// GetInternalGroupIds returns the InternalGroupIds field value if set, zero value otherwise.
func (o *DataFilter) GetInternalGroupIds() []string {
	if o == nil || isNil(o.InternalGroupIds) {
		var ret []string
		return ret
	}
	return o.InternalGroupIds
}

// GetInternalGroupIdsOk returns a tuple with the InternalGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetInternalGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.InternalGroupIds) {
    return nil, false
	}
	return o.InternalGroupIds, true
}

// HasInternalGroupIds returns a boolean if a field has been set.
func (o *DataFilter) HasInternalGroupIds() bool {
	if o != nil && !isNil(o.InternalGroupIds) {
		return true
	}

	return false
}

// SetInternalGroupIds gets a reference to the given []string and assigns it to the InternalGroupIds field.
func (o *DataFilter) SetInternalGroupIds(v []string) {
	o.InternalGroupIds = v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *DataFilter) GetSupis() []string {
	if o == nil || isNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetSupisOk() ([]string, bool) {
	if o == nil || isNil(o.Supis) {
    return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *DataFilter) HasSupis() bool {
	if o != nil && !isNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *DataFilter) SetSupis(v []string) {
	o.Supis = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *DataFilter) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
    return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *DataFilter) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *DataFilter) SetAppIds(v []string) {
	o.AppIds = v
}

// GetUeIpv4s returns the UeIpv4s field value if set, zero value otherwise.
func (o *DataFilter) GetUeIpv4s() []string {
	if o == nil || isNil(o.UeIpv4s) {
		var ret []string
		return ret
	}
	return o.UeIpv4s
}

// GetUeIpv4sOk returns a tuple with the UeIpv4s field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetUeIpv4sOk() ([]string, bool) {
	if o == nil || isNil(o.UeIpv4s) {
    return nil, false
	}
	return o.UeIpv4s, true
}

// HasUeIpv4s returns a boolean if a field has been set.
func (o *DataFilter) HasUeIpv4s() bool {
	if o != nil && !isNil(o.UeIpv4s) {
		return true
	}

	return false
}

// SetUeIpv4s gets a reference to the given []string and assigns it to the UeIpv4s field.
func (o *DataFilter) SetUeIpv4s(v []string) {
	o.UeIpv4s = v
}

// GetUeIpv6s returns the UeIpv6s field value if set, zero value otherwise.
func (o *DataFilter) GetUeIpv6s() []Ipv6Addr {
	if o == nil || isNil(o.UeIpv6s) {
		var ret []Ipv6Addr
		return ret
	}
	return o.UeIpv6s
}

// GetUeIpv6sOk returns a tuple with the UeIpv6s field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetUeIpv6sOk() ([]Ipv6Addr, bool) {
	if o == nil || isNil(o.UeIpv6s) {
    return nil, false
	}
	return o.UeIpv6s, true
}

// HasUeIpv6s returns a boolean if a field has been set.
func (o *DataFilter) HasUeIpv6s() bool {
	if o != nil && !isNil(o.UeIpv6s) {
		return true
	}

	return false
}

// SetUeIpv6s gets a reference to the given []Ipv6Addr and assigns it to the UeIpv6s field.
func (o *DataFilter) SetUeIpv6s(v []Ipv6Addr) {
	o.UeIpv6s = v
}

// GetUeMacs returns the UeMacs field value if set, zero value otherwise.
func (o *DataFilter) GetUeMacs() []string {
	if o == nil || isNil(o.UeMacs) {
		var ret []string
		return ret
	}
	return o.UeMacs
}

// GetUeMacsOk returns a tuple with the UeMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetUeMacsOk() ([]string, bool) {
	if o == nil || isNil(o.UeMacs) {
    return nil, false
	}
	return o.UeMacs, true
}

// HasUeMacs returns a boolean if a field has been set.
func (o *DataFilter) HasUeMacs() bool {
	if o != nil && !isNil(o.UeMacs) {
		return true
	}

	return false
}

// SetUeMacs gets a reference to the given []string and assigns it to the UeMacs field.
func (o *DataFilter) SetUeMacs(v []string) {
	o.UeMacs = v
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *DataFilter) GetAnyUeInd() bool {
	if o == nil || isNil(o.AnyUeInd) {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || isNil(o.AnyUeInd) {
    return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *DataFilter) HasAnyUeInd() bool {
	if o != nil && !isNil(o.AnyUeInd) {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *DataFilter) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetDnnSnssaiInfos returns the DnnSnssaiInfos field value if set, zero value otherwise.
func (o *DataFilter) GetDnnSnssaiInfos() []DnnSnssaiInformation {
	if o == nil || isNil(o.DnnSnssaiInfos) {
		var ret []DnnSnssaiInformation
		return ret
	}
	return o.DnnSnssaiInfos
}

// GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFilter) GetDnnSnssaiInfosOk() ([]DnnSnssaiInformation, bool) {
	if o == nil || isNil(o.DnnSnssaiInfos) {
    return nil, false
	}
	return o.DnnSnssaiInfos, true
}

// HasDnnSnssaiInfos returns a boolean if a field has been set.
func (o *DataFilter) HasDnnSnssaiInfos() bool {
	if o != nil && !isNil(o.DnnSnssaiInfos) {
		return true
	}

	return false
}

// SetDnnSnssaiInfos gets a reference to the given []DnnSnssaiInformation and assigns it to the DnnSnssaiInfos field.
func (o *DataFilter) SetDnnSnssaiInfos(v []DnnSnssaiInformation) {
	o.DnnSnssaiInfos = v
}

func (o DataFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataInd"] = o.DataInd
	}
	if !isNil(o.Dnns) {
		toSerialize["dnns"] = o.Dnns
	}
	if !isNil(o.Snssais) {
		toSerialize["snssais"] = o.Snssais
	}
	if !isNil(o.InternalGroupIds) {
		toSerialize["internalGroupIds"] = o.InternalGroupIds
	}
	if !isNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.UeIpv4s) {
		toSerialize["ueIpv4s"] = o.UeIpv4s
	}
	if !isNil(o.UeIpv6s) {
		toSerialize["ueIpv6s"] = o.UeIpv6s
	}
	if !isNil(o.UeMacs) {
		toSerialize["ueMacs"] = o.UeMacs
	}
	if !isNil(o.AnyUeInd) {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if !isNil(o.DnnSnssaiInfos) {
		toSerialize["dnnSnssaiInfos"] = o.DnnSnssaiInfos
	}
	return json.Marshal(toSerialize)
}

type NullableDataFilter struct {
	value *DataFilter
	isSet bool
}

func (v NullableDataFilter) Get() *DataFilter {
	return v.value
}

func (v *NullableDataFilter) Set(val *DataFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFilter(val *DataFilter) *NullableDataFilter {
	return &NullableDataFilter{value: val, isSet: true}
}

func (v NullableDataFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


