/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"time"
)

// ChargingDataRequest struct for ChargingDataRequest
type ChargingDataRequest struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	SubscriberIdentifier *string `json:"subscriberIdentifier,omitempty"`
	NfConsumerIdentification NFIdentification `json:"nfConsumerIdentification"`
	// string with format 'date-time' as defined in OpenAPI.
	InvocationTimeStamp time.Time `json:"invocationTimeStamp"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	InvocationSequenceNumber int32 `json:"invocationSequenceNumber"`
	RetransmissionIndicator *bool `json:"retransmissionIndicator,omitempty"`
	ServiceSpecificationInfo *string `json:"serviceSpecificationInfo,omitempty"`
	MultipleUnitUsage []MultipleUnitUsage `json:"multipleUnitUsage,omitempty"`
	Triggers []Trigger `json:"triggers,omitempty"`
	PDUSessionChargingInformation *PDUSessionChargingInformation `json:"pDUSessionChargingInformation,omitempty"`
	RoamingQBCInformation *RoamingQBCInformation `json:"roamingQBCInformation,omitempty"`
}

// NewChargingDataRequest instantiates a new ChargingDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargingDataRequest(nfConsumerIdentification NFIdentification, invocationTimeStamp time.Time, invocationSequenceNumber int32) *ChargingDataRequest {
	this := ChargingDataRequest{}
	this.NfConsumerIdentification = nfConsumerIdentification
	this.InvocationTimeStamp = invocationTimeStamp
	this.InvocationSequenceNumber = invocationSequenceNumber
	return &this
}

// NewChargingDataRequestWithDefaults instantiates a new ChargingDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargingDataRequestWithDefaults() *ChargingDataRequest {
	this := ChargingDataRequest{}
	return &this
}

// GetSubscriberIdentifier returns the SubscriberIdentifier field value if set, zero value otherwise.
func (o *ChargingDataRequest) GetSubscriberIdentifier() string {
	if o == nil || isNil(o.SubscriberIdentifier) {
		var ret string
		return ret
	}
	return *o.SubscriberIdentifier
}

// GetSubscriberIdentifierOk returns a tuple with the SubscriberIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetSubscriberIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.SubscriberIdentifier) {
    return nil, false
	}
	return o.SubscriberIdentifier, true
}

// HasSubscriberIdentifier returns a boolean if a field has been set.
func (o *ChargingDataRequest) HasSubscriberIdentifier() bool {
	if o != nil && !isNil(o.SubscriberIdentifier) {
		return true
	}

	return false
}

// SetSubscriberIdentifier gets a reference to the given string and assigns it to the SubscriberIdentifier field.
func (o *ChargingDataRequest) SetSubscriberIdentifier(v string) {
	o.SubscriberIdentifier = &v
}

// GetNfConsumerIdentification returns the NfConsumerIdentification field value
func (o *ChargingDataRequest) GetNfConsumerIdentification() NFIdentification {
	if o == nil {
		var ret NFIdentification
		return ret
	}

	return o.NfConsumerIdentification
}

// GetNfConsumerIdentificationOk returns a tuple with the NfConsumerIdentification field value
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetNfConsumerIdentificationOk() (*NFIdentification, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfConsumerIdentification, true
}

// SetNfConsumerIdentification sets field value
func (o *ChargingDataRequest) SetNfConsumerIdentification(v NFIdentification) {
	o.NfConsumerIdentification = v
}

// GetInvocationTimeStamp returns the InvocationTimeStamp field value
func (o *ChargingDataRequest) GetInvocationTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.InvocationTimeStamp
}

// GetInvocationTimeStampOk returns a tuple with the InvocationTimeStamp field value
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetInvocationTimeStampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InvocationTimeStamp, true
}

// SetInvocationTimeStamp sets field value
func (o *ChargingDataRequest) SetInvocationTimeStamp(v time.Time) {
	o.InvocationTimeStamp = v
}

// GetInvocationSequenceNumber returns the InvocationSequenceNumber field value
func (o *ChargingDataRequest) GetInvocationSequenceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InvocationSequenceNumber
}

// GetInvocationSequenceNumberOk returns a tuple with the InvocationSequenceNumber field value
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetInvocationSequenceNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InvocationSequenceNumber, true
}

// SetInvocationSequenceNumber sets field value
func (o *ChargingDataRequest) SetInvocationSequenceNumber(v int32) {
	o.InvocationSequenceNumber = v
}

// GetRetransmissionIndicator returns the RetransmissionIndicator field value if set, zero value otherwise.
func (o *ChargingDataRequest) GetRetransmissionIndicator() bool {
	if o == nil || isNil(o.RetransmissionIndicator) {
		var ret bool
		return ret
	}
	return *o.RetransmissionIndicator
}

// GetRetransmissionIndicatorOk returns a tuple with the RetransmissionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetRetransmissionIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.RetransmissionIndicator) {
    return nil, false
	}
	return o.RetransmissionIndicator, true
}

// HasRetransmissionIndicator returns a boolean if a field has been set.
func (o *ChargingDataRequest) HasRetransmissionIndicator() bool {
	if o != nil && !isNil(o.RetransmissionIndicator) {
		return true
	}

	return false
}

// SetRetransmissionIndicator gets a reference to the given bool and assigns it to the RetransmissionIndicator field.
func (o *ChargingDataRequest) SetRetransmissionIndicator(v bool) {
	o.RetransmissionIndicator = &v
}

// GetServiceSpecificationInfo returns the ServiceSpecificationInfo field value if set, zero value otherwise.
func (o *ChargingDataRequest) GetServiceSpecificationInfo() string {
	if o == nil || isNil(o.ServiceSpecificationInfo) {
		var ret string
		return ret
	}
	return *o.ServiceSpecificationInfo
}

// GetServiceSpecificationInfoOk returns a tuple with the ServiceSpecificationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetServiceSpecificationInfoOk() (*string, bool) {
	if o == nil || isNil(o.ServiceSpecificationInfo) {
    return nil, false
	}
	return o.ServiceSpecificationInfo, true
}

// HasServiceSpecificationInfo returns a boolean if a field has been set.
func (o *ChargingDataRequest) HasServiceSpecificationInfo() bool {
	if o != nil && !isNil(o.ServiceSpecificationInfo) {
		return true
	}

	return false
}

// SetServiceSpecificationInfo gets a reference to the given string and assigns it to the ServiceSpecificationInfo field.
func (o *ChargingDataRequest) SetServiceSpecificationInfo(v string) {
	o.ServiceSpecificationInfo = &v
}

// GetMultipleUnitUsage returns the MultipleUnitUsage field value if set, zero value otherwise.
func (o *ChargingDataRequest) GetMultipleUnitUsage() []MultipleUnitUsage {
	if o == nil || isNil(o.MultipleUnitUsage) {
		var ret []MultipleUnitUsage
		return ret
	}
	return o.MultipleUnitUsage
}

// GetMultipleUnitUsageOk returns a tuple with the MultipleUnitUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetMultipleUnitUsageOk() ([]MultipleUnitUsage, bool) {
	if o == nil || isNil(o.MultipleUnitUsage) {
    return nil, false
	}
	return o.MultipleUnitUsage, true
}

// HasMultipleUnitUsage returns a boolean if a field has been set.
func (o *ChargingDataRequest) HasMultipleUnitUsage() bool {
	if o != nil && !isNil(o.MultipleUnitUsage) {
		return true
	}

	return false
}

// SetMultipleUnitUsage gets a reference to the given []MultipleUnitUsage and assigns it to the MultipleUnitUsage field.
func (o *ChargingDataRequest) SetMultipleUnitUsage(v []MultipleUnitUsage) {
	o.MultipleUnitUsage = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *ChargingDataRequest) GetTriggers() []Trigger {
	if o == nil || isNil(o.Triggers) {
		var ret []Trigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetTriggersOk() ([]Trigger, bool) {
	if o == nil || isNil(o.Triggers) {
    return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *ChargingDataRequest) HasTriggers() bool {
	if o != nil && !isNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []Trigger and assigns it to the Triggers field.
func (o *ChargingDataRequest) SetTriggers(v []Trigger) {
	o.Triggers = v
}

// GetPDUSessionChargingInformation returns the PDUSessionChargingInformation field value if set, zero value otherwise.
func (o *ChargingDataRequest) GetPDUSessionChargingInformation() PDUSessionChargingInformation {
	if o == nil || isNil(o.PDUSessionChargingInformation) {
		var ret PDUSessionChargingInformation
		return ret
	}
	return *o.PDUSessionChargingInformation
}

// GetPDUSessionChargingInformationOk returns a tuple with the PDUSessionChargingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetPDUSessionChargingInformationOk() (*PDUSessionChargingInformation, bool) {
	if o == nil || isNil(o.PDUSessionChargingInformation) {
    return nil, false
	}
	return o.PDUSessionChargingInformation, true
}

// HasPDUSessionChargingInformation returns a boolean if a field has been set.
func (o *ChargingDataRequest) HasPDUSessionChargingInformation() bool {
	if o != nil && !isNil(o.PDUSessionChargingInformation) {
		return true
	}

	return false
}

// SetPDUSessionChargingInformation gets a reference to the given PDUSessionChargingInformation and assigns it to the PDUSessionChargingInformation field.
func (o *ChargingDataRequest) SetPDUSessionChargingInformation(v PDUSessionChargingInformation) {
	o.PDUSessionChargingInformation = &v
}

// GetRoamingQBCInformation returns the RoamingQBCInformation field value if set, zero value otherwise.
func (o *ChargingDataRequest) GetRoamingQBCInformation() RoamingQBCInformation {
	if o == nil || isNil(o.RoamingQBCInformation) {
		var ret RoamingQBCInformation
		return ret
	}
	return *o.RoamingQBCInformation
}

// GetRoamingQBCInformationOk returns a tuple with the RoamingQBCInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingDataRequest) GetRoamingQBCInformationOk() (*RoamingQBCInformation, bool) {
	if o == nil || isNil(o.RoamingQBCInformation) {
    return nil, false
	}
	return o.RoamingQBCInformation, true
}

// HasRoamingQBCInformation returns a boolean if a field has been set.
func (o *ChargingDataRequest) HasRoamingQBCInformation() bool {
	if o != nil && !isNil(o.RoamingQBCInformation) {
		return true
	}

	return false
}

// SetRoamingQBCInformation gets a reference to the given RoamingQBCInformation and assigns it to the RoamingQBCInformation field.
func (o *ChargingDataRequest) SetRoamingQBCInformation(v RoamingQBCInformation) {
	o.RoamingQBCInformation = &v
}

func (o ChargingDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubscriberIdentifier) {
		toSerialize["subscriberIdentifier"] = o.SubscriberIdentifier
	}
	if true {
		toSerialize["nfConsumerIdentification"] = o.NfConsumerIdentification
	}
	if true {
		toSerialize["invocationTimeStamp"] = o.InvocationTimeStamp
	}
	if true {
		toSerialize["invocationSequenceNumber"] = o.InvocationSequenceNumber
	}
	if !isNil(o.RetransmissionIndicator) {
		toSerialize["retransmissionIndicator"] = o.RetransmissionIndicator
	}
	if !isNil(o.ServiceSpecificationInfo) {
		toSerialize["serviceSpecificationInfo"] = o.ServiceSpecificationInfo
	}
	if !isNil(o.MultipleUnitUsage) {
		toSerialize["multipleUnitUsage"] = o.MultipleUnitUsage
	}
	if !isNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !isNil(o.PDUSessionChargingInformation) {
		toSerialize["pDUSessionChargingInformation"] = o.PDUSessionChargingInformation
	}
	if !isNil(o.RoamingQBCInformation) {
		toSerialize["roamingQBCInformation"] = o.RoamingQBCInformation
	}
	return json.Marshal(toSerialize)
}

type NullableChargingDataRequest struct {
	value *ChargingDataRequest
	isSet bool
}

func (v NullableChargingDataRequest) Get() *ChargingDataRequest {
	return v.value
}

func (v *NullableChargingDataRequest) Set(val *ChargingDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingDataRequest(val *ChargingDataRequest) *NullableChargingDataRequest {
	return &NullableChargingDataRequest{value: val, isSet: true}
}

func (v NullableChargingDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


