/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
)

// QosFlowAddModifyRequestItem Individual QoS flow requested to be created or modified
type QosFlowAddModifyRequestItem struct {
	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi"`
	// EPS Bearer Identifier
	Ebi *int32 `json:"ebi,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	QosRules *string `json:"qosRules,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	QosFlowDescription *string `json:"qosFlowDescription,omitempty"`
	QosFlowProfile *QosFlowProfile `json:"qosFlowProfile,omitempty"`
	AssociatedAnType *QosFlowAccessType `json:"associatedAnType,omitempty"`
}

// NewQosFlowAddModifyRequestItem instantiates a new QosFlowAddModifyRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosFlowAddModifyRequestItem(qfi int32) *QosFlowAddModifyRequestItem {
	this := QosFlowAddModifyRequestItem{}
	this.Qfi = qfi
	return &this
}

// NewQosFlowAddModifyRequestItemWithDefaults instantiates a new QosFlowAddModifyRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosFlowAddModifyRequestItemWithDefaults() *QosFlowAddModifyRequestItem {
	this := QosFlowAddModifyRequestItem{}
	return &this
}

// GetQfi returns the Qfi field value
func (o *QosFlowAddModifyRequestItem) GetQfi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Qfi
}

// GetQfiOk returns a tuple with the Qfi field value
// and a boolean to check if the value has been set.
func (o *QosFlowAddModifyRequestItem) GetQfiOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Qfi, true
}

// SetQfi sets field value
func (o *QosFlowAddModifyRequestItem) SetQfi(v int32) {
	o.Qfi = v
}

// GetEbi returns the Ebi field value if set, zero value otherwise.
func (o *QosFlowAddModifyRequestItem) GetEbi() int32 {
	if o == nil || isNil(o.Ebi) {
		var ret int32
		return ret
	}
	return *o.Ebi
}

// GetEbiOk returns a tuple with the Ebi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowAddModifyRequestItem) GetEbiOk() (*int32, bool) {
	if o == nil || isNil(o.Ebi) {
    return nil, false
	}
	return o.Ebi, true
}

// HasEbi returns a boolean if a field has been set.
func (o *QosFlowAddModifyRequestItem) HasEbi() bool {
	if o != nil && !isNil(o.Ebi) {
		return true
	}

	return false
}

// SetEbi gets a reference to the given int32 and assigns it to the Ebi field.
func (o *QosFlowAddModifyRequestItem) SetEbi(v int32) {
	o.Ebi = &v
}

// GetQosRules returns the QosRules field value if set, zero value otherwise.
func (o *QosFlowAddModifyRequestItem) GetQosRules() string {
	if o == nil || isNil(o.QosRules) {
		var ret string
		return ret
	}
	return *o.QosRules
}

// GetQosRulesOk returns a tuple with the QosRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowAddModifyRequestItem) GetQosRulesOk() (*string, bool) {
	if o == nil || isNil(o.QosRules) {
    return nil, false
	}
	return o.QosRules, true
}

// HasQosRules returns a boolean if a field has been set.
func (o *QosFlowAddModifyRequestItem) HasQosRules() bool {
	if o != nil && !isNil(o.QosRules) {
		return true
	}

	return false
}

// SetQosRules gets a reference to the given string and assigns it to the QosRules field.
func (o *QosFlowAddModifyRequestItem) SetQosRules(v string) {
	o.QosRules = &v
}

// GetQosFlowDescription returns the QosFlowDescription field value if set, zero value otherwise.
func (o *QosFlowAddModifyRequestItem) GetQosFlowDescription() string {
	if o == nil || isNil(o.QosFlowDescription) {
		var ret string
		return ret
	}
	return *o.QosFlowDescription
}

// GetQosFlowDescriptionOk returns a tuple with the QosFlowDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowAddModifyRequestItem) GetQosFlowDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.QosFlowDescription) {
    return nil, false
	}
	return o.QosFlowDescription, true
}

// HasQosFlowDescription returns a boolean if a field has been set.
func (o *QosFlowAddModifyRequestItem) HasQosFlowDescription() bool {
	if o != nil && !isNil(o.QosFlowDescription) {
		return true
	}

	return false
}

// SetQosFlowDescription gets a reference to the given string and assigns it to the QosFlowDescription field.
func (o *QosFlowAddModifyRequestItem) SetQosFlowDescription(v string) {
	o.QosFlowDescription = &v
}

// GetQosFlowProfile returns the QosFlowProfile field value if set, zero value otherwise.
func (o *QosFlowAddModifyRequestItem) GetQosFlowProfile() QosFlowProfile {
	if o == nil || isNil(o.QosFlowProfile) {
		var ret QosFlowProfile
		return ret
	}
	return *o.QosFlowProfile
}

// GetQosFlowProfileOk returns a tuple with the QosFlowProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowAddModifyRequestItem) GetQosFlowProfileOk() (*QosFlowProfile, bool) {
	if o == nil || isNil(o.QosFlowProfile) {
    return nil, false
	}
	return o.QosFlowProfile, true
}

// HasQosFlowProfile returns a boolean if a field has been set.
func (o *QosFlowAddModifyRequestItem) HasQosFlowProfile() bool {
	if o != nil && !isNil(o.QosFlowProfile) {
		return true
	}

	return false
}

// SetQosFlowProfile gets a reference to the given QosFlowProfile and assigns it to the QosFlowProfile field.
func (o *QosFlowAddModifyRequestItem) SetQosFlowProfile(v QosFlowProfile) {
	o.QosFlowProfile = &v
}

// GetAssociatedAnType returns the AssociatedAnType field value if set, zero value otherwise.
func (o *QosFlowAddModifyRequestItem) GetAssociatedAnType() QosFlowAccessType {
	if o == nil || isNil(o.AssociatedAnType) {
		var ret QosFlowAccessType
		return ret
	}
	return *o.AssociatedAnType
}

// GetAssociatedAnTypeOk returns a tuple with the AssociatedAnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowAddModifyRequestItem) GetAssociatedAnTypeOk() (*QosFlowAccessType, bool) {
	if o == nil || isNil(o.AssociatedAnType) {
    return nil, false
	}
	return o.AssociatedAnType, true
}

// HasAssociatedAnType returns a boolean if a field has been set.
func (o *QosFlowAddModifyRequestItem) HasAssociatedAnType() bool {
	if o != nil && !isNil(o.AssociatedAnType) {
		return true
	}

	return false
}

// SetAssociatedAnType gets a reference to the given QosFlowAccessType and assigns it to the AssociatedAnType field.
func (o *QosFlowAddModifyRequestItem) SetAssociatedAnType(v QosFlowAccessType) {
	o.AssociatedAnType = &v
}

func (o QosFlowAddModifyRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["qfi"] = o.Qfi
	}
	if !isNil(o.Ebi) {
		toSerialize["ebi"] = o.Ebi
	}
	if !isNil(o.QosRules) {
		toSerialize["qosRules"] = o.QosRules
	}
	if !isNil(o.QosFlowDescription) {
		toSerialize["qosFlowDescription"] = o.QosFlowDescription
	}
	if !isNil(o.QosFlowProfile) {
		toSerialize["qosFlowProfile"] = o.QosFlowProfile
	}
	if !isNil(o.AssociatedAnType) {
		toSerialize["associatedAnType"] = o.AssociatedAnType
	}
	return json.Marshal(toSerialize)
}

type NullableQosFlowAddModifyRequestItem struct {
	value *QosFlowAddModifyRequestItem
	isSet bool
}

func (v NullableQosFlowAddModifyRequestItem) Get() *QosFlowAddModifyRequestItem {
	return v.value
}

func (v *NullableQosFlowAddModifyRequestItem) Set(val *QosFlowAddModifyRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowAddModifyRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowAddModifyRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowAddModifyRequestItem(val *QosFlowAddModifyRequestItem) *NullableQosFlowAddModifyRequestItem {
	return &NullableQosFlowAddModifyRequestItem{value: val, isSet: true}
}

func (v NullableQosFlowAddModifyRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowAddModifyRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


