/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
)

// ReportingArea Indicates an area for event reporting.
type ReportingArea struct {
	AreaType ReportingAreaType `json:"areaType"`
	Tai *Tai `json:"tai,omitempty"`
	Ecgi *Ecgi `json:"ecgi,omitempty"`
	Ncgi *Ncgi `json:"ncgi,omitempty"`
}

// NewReportingArea instantiates a new ReportingArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingArea(areaType ReportingAreaType) *ReportingArea {
	this := ReportingArea{}
	this.AreaType = areaType
	return &this
}

// NewReportingAreaWithDefaults instantiates a new ReportingArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingAreaWithDefaults() *ReportingArea {
	this := ReportingArea{}
	return &this
}

// GetAreaType returns the AreaType field value
func (o *ReportingArea) GetAreaType() ReportingAreaType {
	if o == nil {
		var ret ReportingAreaType
		return ret
	}

	return o.AreaType
}

// GetAreaTypeOk returns a tuple with the AreaType field value
// and a boolean to check if the value has been set.
func (o *ReportingArea) GetAreaTypeOk() (*ReportingAreaType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AreaType, true
}

// SetAreaType sets field value
func (o *ReportingArea) SetAreaType(v ReportingAreaType) {
	o.AreaType = v
}

// GetTai returns the Tai field value if set, zero value otherwise.
func (o *ReportingArea) GetTai() Tai {
	if o == nil || isNil(o.Tai) {
		var ret Tai
		return ret
	}
	return *o.Tai
}

// GetTaiOk returns a tuple with the Tai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingArea) GetTaiOk() (*Tai, bool) {
	if o == nil || isNil(o.Tai) {
    return nil, false
	}
	return o.Tai, true
}

// HasTai returns a boolean if a field has been set.
func (o *ReportingArea) HasTai() bool {
	if o != nil && !isNil(o.Tai) {
		return true
	}

	return false
}

// SetTai gets a reference to the given Tai and assigns it to the Tai field.
func (o *ReportingArea) SetTai(v Tai) {
	o.Tai = &v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *ReportingArea) GetEcgi() Ecgi {
	if o == nil || isNil(o.Ecgi) {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingArea) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || isNil(o.Ecgi) {
    return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *ReportingArea) HasEcgi() bool {
	if o != nil && !isNil(o.Ecgi) {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *ReportingArea) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

// GetNcgi returns the Ncgi field value if set, zero value otherwise.
func (o *ReportingArea) GetNcgi() Ncgi {
	if o == nil || isNil(o.Ncgi) {
		var ret Ncgi
		return ret
	}
	return *o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingArea) GetNcgiOk() (*Ncgi, bool) {
	if o == nil || isNil(o.Ncgi) {
    return nil, false
	}
	return o.Ncgi, true
}

// HasNcgi returns a boolean if a field has been set.
func (o *ReportingArea) HasNcgi() bool {
	if o != nil && !isNil(o.Ncgi) {
		return true
	}

	return false
}

// SetNcgi gets a reference to the given Ncgi and assigns it to the Ncgi field.
func (o *ReportingArea) SetNcgi(v Ncgi) {
	o.Ncgi = &v
}

func (o ReportingArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["areaType"] = o.AreaType
	}
	if !isNil(o.Tai) {
		toSerialize["tai"] = o.Tai
	}
	if !isNil(o.Ecgi) {
		toSerialize["ecgi"] = o.Ecgi
	}
	if !isNil(o.Ncgi) {
		toSerialize["ncgi"] = o.Ncgi
	}
	return json.Marshal(toSerialize)
}

type NullableReportingArea struct {
	value *ReportingArea
	isSet bool
}

func (v NullableReportingArea) Get() *ReportingArea {
	return v.value
}

func (v *NullableReportingArea) Set(val *ReportingArea) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingArea) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingArea(val *ReportingArea) *NullableReportingArea {
	return &NullableReportingArea{value: val, isSet: true}
}

func (v NullableReportingArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


