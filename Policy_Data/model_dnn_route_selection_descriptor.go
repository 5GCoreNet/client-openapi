/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Policy_Data

import (
	"encoding/json"
)

// DnnRouteSelectionDescriptor Contains the route selector parameters (PDU session types, SSC modes and ATSSS information) per DNN 
type DnnRouteSelectionDescriptor struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	SscModes []SscMode `json:"sscModes,omitempty"`
	PduSessTypes []PduSessionType `json:"pduSessTypes,omitempty"`
	// Indicates whether MA PDU session establishment is allowed for this DNN. When set to value true MA PDU session establishment is allowed for this DNN. 
	AtsssInfo *bool `json:"atsssInfo,omitempty"`
}

// NewDnnRouteSelectionDescriptor instantiates a new DnnRouteSelectionDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnRouteSelectionDescriptor(dnn string) *DnnRouteSelectionDescriptor {
	this := DnnRouteSelectionDescriptor{}
	this.Dnn = dnn
	var atsssInfo bool = false
	this.AtsssInfo = &atsssInfo
	return &this
}

// NewDnnRouteSelectionDescriptorWithDefaults instantiates a new DnnRouteSelectionDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnRouteSelectionDescriptorWithDefaults() *DnnRouteSelectionDescriptor {
	this := DnnRouteSelectionDescriptor{}
	var atsssInfo bool = false
	this.AtsssInfo = &atsssInfo
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnRouteSelectionDescriptor) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnRouteSelectionDescriptor) GetDnnOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnRouteSelectionDescriptor) SetDnn(v string) {
	o.Dnn = v
}

// GetSscModes returns the SscModes field value if set, zero value otherwise.
func (o *DnnRouteSelectionDescriptor) GetSscModes() []SscMode {
	if o == nil || isNil(o.SscModes) {
		var ret []SscMode
		return ret
	}
	return o.SscModes
}

// GetSscModesOk returns a tuple with the SscModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnRouteSelectionDescriptor) GetSscModesOk() ([]SscMode, bool) {
	if o == nil || isNil(o.SscModes) {
    return nil, false
	}
	return o.SscModes, true
}

// HasSscModes returns a boolean if a field has been set.
func (o *DnnRouteSelectionDescriptor) HasSscModes() bool {
	if o != nil && !isNil(o.SscModes) {
		return true
	}

	return false
}

// SetSscModes gets a reference to the given []SscMode and assigns it to the SscModes field.
func (o *DnnRouteSelectionDescriptor) SetSscModes(v []SscMode) {
	o.SscModes = v
}

// GetPduSessTypes returns the PduSessTypes field value if set, zero value otherwise.
func (o *DnnRouteSelectionDescriptor) GetPduSessTypes() []PduSessionType {
	if o == nil || isNil(o.PduSessTypes) {
		var ret []PduSessionType
		return ret
	}
	return o.PduSessTypes
}

// GetPduSessTypesOk returns a tuple with the PduSessTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnRouteSelectionDescriptor) GetPduSessTypesOk() ([]PduSessionType, bool) {
	if o == nil || isNil(o.PduSessTypes) {
    return nil, false
	}
	return o.PduSessTypes, true
}

// HasPduSessTypes returns a boolean if a field has been set.
func (o *DnnRouteSelectionDescriptor) HasPduSessTypes() bool {
	if o != nil && !isNil(o.PduSessTypes) {
		return true
	}

	return false
}

// SetPduSessTypes gets a reference to the given []PduSessionType and assigns it to the PduSessTypes field.
func (o *DnnRouteSelectionDescriptor) SetPduSessTypes(v []PduSessionType) {
	o.PduSessTypes = v
}

// GetAtsssInfo returns the AtsssInfo field value if set, zero value otherwise.
func (o *DnnRouteSelectionDescriptor) GetAtsssInfo() bool {
	if o == nil || isNil(o.AtsssInfo) {
		var ret bool
		return ret
	}
	return *o.AtsssInfo
}

// GetAtsssInfoOk returns a tuple with the AtsssInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnRouteSelectionDescriptor) GetAtsssInfoOk() (*bool, bool) {
	if o == nil || isNil(o.AtsssInfo) {
    return nil, false
	}
	return o.AtsssInfo, true
}

// HasAtsssInfo returns a boolean if a field has been set.
func (o *DnnRouteSelectionDescriptor) HasAtsssInfo() bool {
	if o != nil && !isNil(o.AtsssInfo) {
		return true
	}

	return false
}

// SetAtsssInfo gets a reference to the given bool and assigns it to the AtsssInfo field.
func (o *DnnRouteSelectionDescriptor) SetAtsssInfo(v bool) {
	o.AtsssInfo = &v
}

func (o DnnRouteSelectionDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.SscModes) {
		toSerialize["sscModes"] = o.SscModes
	}
	if !isNil(o.PduSessTypes) {
		toSerialize["pduSessTypes"] = o.PduSessTypes
	}
	if !isNil(o.AtsssInfo) {
		toSerialize["atsssInfo"] = o.AtsssInfo
	}
	return json.Marshal(toSerialize)
}

type NullableDnnRouteSelectionDescriptor struct {
	value *DnnRouteSelectionDescriptor
	isSet bool
}

func (v NullableDnnRouteSelectionDescriptor) Get() *DnnRouteSelectionDescriptor {
	return v.value
}

func (v *NullableDnnRouteSelectionDescriptor) Set(val *DnnRouteSelectionDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnRouteSelectionDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnRouteSelectionDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnRouteSelectionDescriptor(val *DnnRouteSelectionDescriptor) *NullableDnnRouteSelectionDescriptor {
	return &NullableDnnRouteSelectionDescriptor{value: val, isSet: true}
}

func (v NullableDnnRouteSelectionDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnRouteSelectionDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


