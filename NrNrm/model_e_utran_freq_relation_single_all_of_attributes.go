/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// EUtranFreqRelationSingleAllOfAttributes struct for EUtranFreqRelationSingleAllOfAttributes
type EUtranFreqRelationSingleAllOfAttributes struct {
	CellIndividualOffset *CellIndividualOffset `json:"cellIndividualOffset,omitempty"`
	BlackListEntry []int32 `json:"blackListEntry,omitempty"`
	BlackListEntryIdleMode *int32 `json:"blackListEntryIdleMode,omitempty"`
	CellReselectionPriority *int32 `json:"cellReselectionPriority,omitempty"`
	CellReselectionSubPriority *float32 `json:"cellReselectionSubPriority,omitempty"`
	PMax *int32 `json:"pMax,omitempty"`
	QOffsetFreq *float32 `json:"qOffsetFreq,omitempty"`
	QQualMin *float32 `json:"qQualMin,omitempty"`
	QRxLevMin *int32 `json:"qRxLevMin,omitempty"`
	ThreshXHighP *int32 `json:"threshXHighP,omitempty"`
	ThreshXHighQ *int32 `json:"threshXHighQ,omitempty"`
	ThreshXLowP *int32 `json:"threshXLowP,omitempty"`
	ThreshXLowQ *int32 `json:"threshXLowQ,omitempty"`
	TReselectionEutran *int32 `json:"tReselectionEutran,omitempty"`
	TReselectionNRSfHigh *TReselectionNRSf `json:"tReselectionNRSfHigh,omitempty"`
	TReselectionNRSfMedium *TReselectionNRSf `json:"tReselectionNRSfMedium,omitempty"`
	EUTranFrequencyRef *string `json:"eUTranFrequencyRef,omitempty"`
}

// NewEUtranFreqRelationSingleAllOfAttributes instantiates a new EUtranFreqRelationSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEUtranFreqRelationSingleAllOfAttributes() *EUtranFreqRelationSingleAllOfAttributes {
	this := EUtranFreqRelationSingleAllOfAttributes{}
	return &this
}

// NewEUtranFreqRelationSingleAllOfAttributesWithDefaults instantiates a new EUtranFreqRelationSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEUtranFreqRelationSingleAllOfAttributesWithDefaults() *EUtranFreqRelationSingleAllOfAttributes {
	this := EUtranFreqRelationSingleAllOfAttributes{}
	return &this
}

// GetCellIndividualOffset returns the CellIndividualOffset field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellIndividualOffset() CellIndividualOffset {
	if o == nil || isNil(o.CellIndividualOffset) {
		var ret CellIndividualOffset
		return ret
	}
	return *o.CellIndividualOffset
}

// GetCellIndividualOffsetOk returns a tuple with the CellIndividualOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellIndividualOffsetOk() (*CellIndividualOffset, bool) {
	if o == nil || isNil(o.CellIndividualOffset) {
    return nil, false
	}
	return o.CellIndividualOffset, true
}

// HasCellIndividualOffset returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasCellIndividualOffset() bool {
	if o != nil && !isNil(o.CellIndividualOffset) {
		return true
	}

	return false
}

// SetCellIndividualOffset gets a reference to the given CellIndividualOffset and assigns it to the CellIndividualOffset field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetCellIndividualOffset(v CellIndividualOffset) {
	o.CellIndividualOffset = &v
}

// GetBlackListEntry returns the BlackListEntry field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntry() []int32 {
	if o == nil || isNil(o.BlackListEntry) {
		var ret []int32
		return ret
	}
	return o.BlackListEntry
}

// GetBlackListEntryOk returns a tuple with the BlackListEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntryOk() ([]int32, bool) {
	if o == nil || isNil(o.BlackListEntry) {
    return nil, false
	}
	return o.BlackListEntry, true
}

// HasBlackListEntry returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasBlackListEntry() bool {
	if o != nil && !isNil(o.BlackListEntry) {
		return true
	}

	return false
}

// SetBlackListEntry gets a reference to the given []int32 and assigns it to the BlackListEntry field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetBlackListEntry(v []int32) {
	o.BlackListEntry = v
}

// GetBlackListEntryIdleMode returns the BlackListEntryIdleMode field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntryIdleMode() int32 {
	if o == nil || isNil(o.BlackListEntryIdleMode) {
		var ret int32
		return ret
	}
	return *o.BlackListEntryIdleMode
}

// GetBlackListEntryIdleModeOk returns a tuple with the BlackListEntryIdleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetBlackListEntryIdleModeOk() (*int32, bool) {
	if o == nil || isNil(o.BlackListEntryIdleMode) {
    return nil, false
	}
	return o.BlackListEntryIdleMode, true
}

// HasBlackListEntryIdleMode returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasBlackListEntryIdleMode() bool {
	if o != nil && !isNil(o.BlackListEntryIdleMode) {
		return true
	}

	return false
}

// SetBlackListEntryIdleMode gets a reference to the given int32 and assigns it to the BlackListEntryIdleMode field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetBlackListEntryIdleMode(v int32) {
	o.BlackListEntryIdleMode = &v
}

// GetCellReselectionPriority returns the CellReselectionPriority field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionPriority() int32 {
	if o == nil || isNil(o.CellReselectionPriority) {
		var ret int32
		return ret
	}
	return *o.CellReselectionPriority
}

// GetCellReselectionPriorityOk returns a tuple with the CellReselectionPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.CellReselectionPriority) {
    return nil, false
	}
	return o.CellReselectionPriority, true
}

// HasCellReselectionPriority returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasCellReselectionPriority() bool {
	if o != nil && !isNil(o.CellReselectionPriority) {
		return true
	}

	return false
}

// SetCellReselectionPriority gets a reference to the given int32 and assigns it to the CellReselectionPriority field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetCellReselectionPriority(v int32) {
	o.CellReselectionPriority = &v
}

// GetCellReselectionSubPriority returns the CellReselectionSubPriority field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionSubPriority() float32 {
	if o == nil || isNil(o.CellReselectionSubPriority) {
		var ret float32
		return ret
	}
	return *o.CellReselectionSubPriority
}

// GetCellReselectionSubPriorityOk returns a tuple with the CellReselectionSubPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetCellReselectionSubPriorityOk() (*float32, bool) {
	if o == nil || isNil(o.CellReselectionSubPriority) {
    return nil, false
	}
	return o.CellReselectionSubPriority, true
}

// HasCellReselectionSubPriority returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasCellReselectionSubPriority() bool {
	if o != nil && !isNil(o.CellReselectionSubPriority) {
		return true
	}

	return false
}

// SetCellReselectionSubPriority gets a reference to the given float32 and assigns it to the CellReselectionSubPriority field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetCellReselectionSubPriority(v float32) {
	o.CellReselectionSubPriority = &v
}

// GetPMax returns the PMax field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetPMax() int32 {
	if o == nil || isNil(o.PMax) {
		var ret int32
		return ret
	}
	return *o.PMax
}

// GetPMaxOk returns a tuple with the PMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetPMaxOk() (*int32, bool) {
	if o == nil || isNil(o.PMax) {
    return nil, false
	}
	return o.PMax, true
}

// HasPMax returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasPMax() bool {
	if o != nil && !isNil(o.PMax) {
		return true
	}

	return false
}

// SetPMax gets a reference to the given int32 and assigns it to the PMax field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetPMax(v int32) {
	o.PMax = &v
}

// GetQOffsetFreq returns the QOffsetFreq field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetQOffsetFreq() float32 {
	if o == nil || isNil(o.QOffsetFreq) {
		var ret float32
		return ret
	}
	return *o.QOffsetFreq
}

// GetQOffsetFreqOk returns a tuple with the QOffsetFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetQOffsetFreqOk() (*float32, bool) {
	if o == nil || isNil(o.QOffsetFreq) {
    return nil, false
	}
	return o.QOffsetFreq, true
}

// HasQOffsetFreq returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasQOffsetFreq() bool {
	if o != nil && !isNil(o.QOffsetFreq) {
		return true
	}

	return false
}

// SetQOffsetFreq gets a reference to the given float32 and assigns it to the QOffsetFreq field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetQOffsetFreq(v float32) {
	o.QOffsetFreq = &v
}

// GetQQualMin returns the QQualMin field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetQQualMin() float32 {
	if o == nil || isNil(o.QQualMin) {
		var ret float32
		return ret
	}
	return *o.QQualMin
}

// GetQQualMinOk returns a tuple with the QQualMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetQQualMinOk() (*float32, bool) {
	if o == nil || isNil(o.QQualMin) {
    return nil, false
	}
	return o.QQualMin, true
}

// HasQQualMin returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasQQualMin() bool {
	if o != nil && !isNil(o.QQualMin) {
		return true
	}

	return false
}

// SetQQualMin gets a reference to the given float32 and assigns it to the QQualMin field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetQQualMin(v float32) {
	o.QQualMin = &v
}

// GetQRxLevMin returns the QRxLevMin field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetQRxLevMin() int32 {
	if o == nil || isNil(o.QRxLevMin) {
		var ret int32
		return ret
	}
	return *o.QRxLevMin
}

// GetQRxLevMinOk returns a tuple with the QRxLevMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetQRxLevMinOk() (*int32, bool) {
	if o == nil || isNil(o.QRxLevMin) {
    return nil, false
	}
	return o.QRxLevMin, true
}

// HasQRxLevMin returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasQRxLevMin() bool {
	if o != nil && !isNil(o.QRxLevMin) {
		return true
	}

	return false
}

// SetQRxLevMin gets a reference to the given int32 and assigns it to the QRxLevMin field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetQRxLevMin(v int32) {
	o.QRxLevMin = &v
}

// GetThreshXHighP returns the ThreshXHighP field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighP() int32 {
	if o == nil || isNil(o.ThreshXHighP) {
		var ret int32
		return ret
	}
	return *o.ThreshXHighP
}

// GetThreshXHighPOk returns a tuple with the ThreshXHighP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighPOk() (*int32, bool) {
	if o == nil || isNil(o.ThreshXHighP) {
    return nil, false
	}
	return o.ThreshXHighP, true
}

// HasThreshXHighP returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXHighP() bool {
	if o != nil && !isNil(o.ThreshXHighP) {
		return true
	}

	return false
}

// SetThreshXHighP gets a reference to the given int32 and assigns it to the ThreshXHighP field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXHighP(v int32) {
	o.ThreshXHighP = &v
}

// GetThreshXHighQ returns the ThreshXHighQ field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighQ() int32 {
	if o == nil || isNil(o.ThreshXHighQ) {
		var ret int32
		return ret
	}
	return *o.ThreshXHighQ
}

// GetThreshXHighQOk returns a tuple with the ThreshXHighQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXHighQOk() (*int32, bool) {
	if o == nil || isNil(o.ThreshXHighQ) {
    return nil, false
	}
	return o.ThreshXHighQ, true
}

// HasThreshXHighQ returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXHighQ() bool {
	if o != nil && !isNil(o.ThreshXHighQ) {
		return true
	}

	return false
}

// SetThreshXHighQ gets a reference to the given int32 and assigns it to the ThreshXHighQ field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXHighQ(v int32) {
	o.ThreshXHighQ = &v
}

// GetThreshXLowP returns the ThreshXLowP field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowP() int32 {
	if o == nil || isNil(o.ThreshXLowP) {
		var ret int32
		return ret
	}
	return *o.ThreshXLowP
}

// GetThreshXLowPOk returns a tuple with the ThreshXLowP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowPOk() (*int32, bool) {
	if o == nil || isNil(o.ThreshXLowP) {
    return nil, false
	}
	return o.ThreshXLowP, true
}

// HasThreshXLowP returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXLowP() bool {
	if o != nil && !isNil(o.ThreshXLowP) {
		return true
	}

	return false
}

// SetThreshXLowP gets a reference to the given int32 and assigns it to the ThreshXLowP field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXLowP(v int32) {
	o.ThreshXLowP = &v
}

// GetThreshXLowQ returns the ThreshXLowQ field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowQ() int32 {
	if o == nil || isNil(o.ThreshXLowQ) {
		var ret int32
		return ret
	}
	return *o.ThreshXLowQ
}

// GetThreshXLowQOk returns a tuple with the ThreshXLowQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetThreshXLowQOk() (*int32, bool) {
	if o == nil || isNil(o.ThreshXLowQ) {
    return nil, false
	}
	return o.ThreshXLowQ, true
}

// HasThreshXLowQ returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasThreshXLowQ() bool {
	if o != nil && !isNil(o.ThreshXLowQ) {
		return true
	}

	return false
}

// SetThreshXLowQ gets a reference to the given int32 and assigns it to the ThreshXLowQ field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetThreshXLowQ(v int32) {
	o.ThreshXLowQ = &v
}

// GetTReselectionEutran returns the TReselectionEutran field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionEutran() int32 {
	if o == nil || isNil(o.TReselectionEutran) {
		var ret int32
		return ret
	}
	return *o.TReselectionEutran
}

// GetTReselectionEutranOk returns a tuple with the TReselectionEutran field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionEutranOk() (*int32, bool) {
	if o == nil || isNil(o.TReselectionEutran) {
    return nil, false
	}
	return o.TReselectionEutran, true
}

// HasTReselectionEutran returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasTReselectionEutran() bool {
	if o != nil && !isNil(o.TReselectionEutran) {
		return true
	}

	return false
}

// SetTReselectionEutran gets a reference to the given int32 and assigns it to the TReselectionEutran field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetTReselectionEutran(v int32) {
	o.TReselectionEutran = &v
}

// GetTReselectionNRSfHigh returns the TReselectionNRSfHigh field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfHigh() TReselectionNRSf {
	if o == nil || isNil(o.TReselectionNRSfHigh) {
		var ret TReselectionNRSf
		return ret
	}
	return *o.TReselectionNRSfHigh
}

// GetTReselectionNRSfHighOk returns a tuple with the TReselectionNRSfHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfHighOk() (*TReselectionNRSf, bool) {
	if o == nil || isNil(o.TReselectionNRSfHigh) {
    return nil, false
	}
	return o.TReselectionNRSfHigh, true
}

// HasTReselectionNRSfHigh returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasTReselectionNRSfHigh() bool {
	if o != nil && !isNil(o.TReselectionNRSfHigh) {
		return true
	}

	return false
}

// SetTReselectionNRSfHigh gets a reference to the given TReselectionNRSf and assigns it to the TReselectionNRSfHigh field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetTReselectionNRSfHigh(v TReselectionNRSf) {
	o.TReselectionNRSfHigh = &v
}

// GetTReselectionNRSfMedium returns the TReselectionNRSfMedium field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfMedium() TReselectionNRSf {
	if o == nil || isNil(o.TReselectionNRSfMedium) {
		var ret TReselectionNRSf
		return ret
	}
	return *o.TReselectionNRSfMedium
}

// GetTReselectionNRSfMediumOk returns a tuple with the TReselectionNRSfMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetTReselectionNRSfMediumOk() (*TReselectionNRSf, bool) {
	if o == nil || isNil(o.TReselectionNRSfMedium) {
    return nil, false
	}
	return o.TReselectionNRSfMedium, true
}

// HasTReselectionNRSfMedium returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasTReselectionNRSfMedium() bool {
	if o != nil && !isNil(o.TReselectionNRSfMedium) {
		return true
	}

	return false
}

// SetTReselectionNRSfMedium gets a reference to the given TReselectionNRSf and assigns it to the TReselectionNRSfMedium field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetTReselectionNRSfMedium(v TReselectionNRSf) {
	o.TReselectionNRSfMedium = &v
}

// GetEUTranFrequencyRef returns the EUTranFrequencyRef field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetEUTranFrequencyRef() string {
	if o == nil || isNil(o.EUTranFrequencyRef) {
		var ret string
		return ret
	}
	return *o.EUTranFrequencyRef
}

// GetEUTranFrequencyRefOk returns a tuple with the EUTranFrequencyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) GetEUTranFrequencyRefOk() (*string, bool) {
	if o == nil || isNil(o.EUTranFrequencyRef) {
    return nil, false
	}
	return o.EUTranFrequencyRef, true
}

// HasEUTranFrequencyRef returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOfAttributes) HasEUTranFrequencyRef() bool {
	if o != nil && !isNil(o.EUTranFrequencyRef) {
		return true
	}

	return false
}

// SetEUTranFrequencyRef gets a reference to the given string and assigns it to the EUTranFrequencyRef field.
func (o *EUtranFreqRelationSingleAllOfAttributes) SetEUTranFrequencyRef(v string) {
	o.EUTranFrequencyRef = &v
}

func (o EUtranFreqRelationSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CellIndividualOffset) {
		toSerialize["cellIndividualOffset"] = o.CellIndividualOffset
	}
	if !isNil(o.BlackListEntry) {
		toSerialize["blackListEntry"] = o.BlackListEntry
	}
	if !isNil(o.BlackListEntryIdleMode) {
		toSerialize["blackListEntryIdleMode"] = o.BlackListEntryIdleMode
	}
	if !isNil(o.CellReselectionPriority) {
		toSerialize["cellReselectionPriority"] = o.CellReselectionPriority
	}
	if !isNil(o.CellReselectionSubPriority) {
		toSerialize["cellReselectionSubPriority"] = o.CellReselectionSubPriority
	}
	if !isNil(o.PMax) {
		toSerialize["pMax"] = o.PMax
	}
	if !isNil(o.QOffsetFreq) {
		toSerialize["qOffsetFreq"] = o.QOffsetFreq
	}
	if !isNil(o.QQualMin) {
		toSerialize["qQualMin"] = o.QQualMin
	}
	if !isNil(o.QRxLevMin) {
		toSerialize["qRxLevMin"] = o.QRxLevMin
	}
	if !isNil(o.ThreshXHighP) {
		toSerialize["threshXHighP"] = o.ThreshXHighP
	}
	if !isNil(o.ThreshXHighQ) {
		toSerialize["threshXHighQ"] = o.ThreshXHighQ
	}
	if !isNil(o.ThreshXLowP) {
		toSerialize["threshXLowP"] = o.ThreshXLowP
	}
	if !isNil(o.ThreshXLowQ) {
		toSerialize["threshXLowQ"] = o.ThreshXLowQ
	}
	if !isNil(o.TReselectionEutran) {
		toSerialize["tReselectionEutran"] = o.TReselectionEutran
	}
	if !isNil(o.TReselectionNRSfHigh) {
		toSerialize["tReselectionNRSfHigh"] = o.TReselectionNRSfHigh
	}
	if !isNil(o.TReselectionNRSfMedium) {
		toSerialize["tReselectionNRSfMedium"] = o.TReselectionNRSfMedium
	}
	if !isNil(o.EUTranFrequencyRef) {
		toSerialize["eUTranFrequencyRef"] = o.EUTranFrequencyRef
	}
	return json.Marshal(toSerialize)
}

type NullableEUtranFreqRelationSingleAllOfAttributes struct {
	value *EUtranFreqRelationSingleAllOfAttributes
	isSet bool
}

func (v NullableEUtranFreqRelationSingleAllOfAttributes) Get() *EUtranFreqRelationSingleAllOfAttributes {
	return v.value
}

func (v *NullableEUtranFreqRelationSingleAllOfAttributes) Set(val *EUtranFreqRelationSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEUtranFreqRelationSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEUtranFreqRelationSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEUtranFreqRelationSingleAllOfAttributes(val *EUtranFreqRelationSingleAllOfAttributes) *NullableEUtranFreqRelationSingleAllOfAttributes {
	return &NullableEUtranFreqRelationSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableEUtranFreqRelationSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEUtranFreqRelationSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


