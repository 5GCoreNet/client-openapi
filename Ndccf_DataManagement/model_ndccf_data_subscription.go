/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// NdccfDataSubscription Represents an Individual DCCF Data Subscription.
type NdccfDataSubscription struct {
	DataSub DataSubscription `json:"dataSub"`
	// String providing an URI formatted according to RFC 3986.
	DataNotifUri string `json:"dataNotifUri"`
	// Notification correlation identifier.
	DataNotifCorrId string `json:"dataNotifCorrId"`
	FormatInstruct *FormattingInstruction `json:"formatInstruct,omitempty"`
	// Processing instructions to be used for sending event notifications.
	ProcInstructs []ProcessingInstruction `json:"procInstructs,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	TargetNfId *string `json:"targetNfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	TargetNfSetId *string `json:"targetNfSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AdrfId *string `json:"adrfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	ArdfSetId *string `json:"ardfSetId,omitempty"`
	TimePeriod *TimeWindow `json:"timePeriod,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	// The purposes of data collection. This attribute may only be provided if the consumer has checked user consent. 
	DataCollectPurposes []DataCollectionPurpose `json:"dataCollectPurposes,omitempty"`
}

// NewNdccfDataSubscription instantiates a new NdccfDataSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNdccfDataSubscription(dataSub DataSubscription, dataNotifUri string, dataNotifCorrId string) *NdccfDataSubscription {
	this := NdccfDataSubscription{}
	this.DataSub = dataSub
	this.DataNotifUri = dataNotifUri
	this.DataNotifCorrId = dataNotifCorrId
	return &this
}

// NewNdccfDataSubscriptionWithDefaults instantiates a new NdccfDataSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNdccfDataSubscriptionWithDefaults() *NdccfDataSubscription {
	this := NdccfDataSubscription{}
	return &this
}

// GetDataSub returns the DataSub field value
func (o *NdccfDataSubscription) GetDataSub() DataSubscription {
	if o == nil {
		var ret DataSubscription
		return ret
	}

	return o.DataSub
}

// GetDataSubOk returns a tuple with the DataSub field value
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetDataSubOk() (*DataSubscription, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataSub, true
}

// SetDataSub sets field value
func (o *NdccfDataSubscription) SetDataSub(v DataSubscription) {
	o.DataSub = v
}

// GetDataNotifUri returns the DataNotifUri field value
func (o *NdccfDataSubscription) GetDataNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataNotifUri
}

// GetDataNotifUriOk returns a tuple with the DataNotifUri field value
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetDataNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataNotifUri, true
}

// SetDataNotifUri sets field value
func (o *NdccfDataSubscription) SetDataNotifUri(v string) {
	o.DataNotifUri = v
}

// GetDataNotifCorrId returns the DataNotifCorrId field value
func (o *NdccfDataSubscription) GetDataNotifCorrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataNotifCorrId
}

// GetDataNotifCorrIdOk returns a tuple with the DataNotifCorrId field value
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetDataNotifCorrIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataNotifCorrId, true
}

// SetDataNotifCorrId sets field value
func (o *NdccfDataSubscription) SetDataNotifCorrId(v string) {
	o.DataNotifCorrId = v
}

// GetFormatInstruct returns the FormatInstruct field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetFormatInstruct() FormattingInstruction {
	if o == nil || isNil(o.FormatInstruct) {
		var ret FormattingInstruction
		return ret
	}
	return *o.FormatInstruct
}

// GetFormatInstructOk returns a tuple with the FormatInstruct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetFormatInstructOk() (*FormattingInstruction, bool) {
	if o == nil || isNil(o.FormatInstruct) {
    return nil, false
	}
	return o.FormatInstruct, true
}

// HasFormatInstruct returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasFormatInstruct() bool {
	if o != nil && !isNil(o.FormatInstruct) {
		return true
	}

	return false
}

// SetFormatInstruct gets a reference to the given FormattingInstruction and assigns it to the FormatInstruct field.
func (o *NdccfDataSubscription) SetFormatInstruct(v FormattingInstruction) {
	o.FormatInstruct = &v
}

// GetProcInstructs returns the ProcInstructs field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetProcInstructs() []ProcessingInstruction {
	if o == nil || isNil(o.ProcInstructs) {
		var ret []ProcessingInstruction
		return ret
	}
	return o.ProcInstructs
}

// GetProcInstructsOk returns a tuple with the ProcInstructs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetProcInstructsOk() ([]ProcessingInstruction, bool) {
	if o == nil || isNil(o.ProcInstructs) {
    return nil, false
	}
	return o.ProcInstructs, true
}

// HasProcInstructs returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasProcInstructs() bool {
	if o != nil && !isNil(o.ProcInstructs) {
		return true
	}

	return false
}

// SetProcInstructs gets a reference to the given []ProcessingInstruction and assigns it to the ProcInstructs field.
func (o *NdccfDataSubscription) SetProcInstructs(v []ProcessingInstruction) {
	o.ProcInstructs = v
}

// GetTargetNfId returns the TargetNfId field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetTargetNfId() string {
	if o == nil || isNil(o.TargetNfId) {
		var ret string
		return ret
	}
	return *o.TargetNfId
}

// GetTargetNfIdOk returns a tuple with the TargetNfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetTargetNfIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetNfId) {
    return nil, false
	}
	return o.TargetNfId, true
}

// HasTargetNfId returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasTargetNfId() bool {
	if o != nil && !isNil(o.TargetNfId) {
		return true
	}

	return false
}

// SetTargetNfId gets a reference to the given string and assigns it to the TargetNfId field.
func (o *NdccfDataSubscription) SetTargetNfId(v string) {
	o.TargetNfId = &v
}

// GetTargetNfSetId returns the TargetNfSetId field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetTargetNfSetId() string {
	if o == nil || isNil(o.TargetNfSetId) {
		var ret string
		return ret
	}
	return *o.TargetNfSetId
}

// GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetTargetNfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetNfSetId) {
    return nil, false
	}
	return o.TargetNfSetId, true
}

// HasTargetNfSetId returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasTargetNfSetId() bool {
	if o != nil && !isNil(o.TargetNfSetId) {
		return true
	}

	return false
}

// SetTargetNfSetId gets a reference to the given string and assigns it to the TargetNfSetId field.
func (o *NdccfDataSubscription) SetTargetNfSetId(v string) {
	o.TargetNfSetId = &v
}

// GetAdrfId returns the AdrfId field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetAdrfId() string {
	if o == nil || isNil(o.AdrfId) {
		var ret string
		return ret
	}
	return *o.AdrfId
}

// GetAdrfIdOk returns a tuple with the AdrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetAdrfIdOk() (*string, bool) {
	if o == nil || isNil(o.AdrfId) {
    return nil, false
	}
	return o.AdrfId, true
}

// HasAdrfId returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasAdrfId() bool {
	if o != nil && !isNil(o.AdrfId) {
		return true
	}

	return false
}

// SetAdrfId gets a reference to the given string and assigns it to the AdrfId field.
func (o *NdccfDataSubscription) SetAdrfId(v string) {
	o.AdrfId = &v
}

// GetArdfSetId returns the ArdfSetId field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetArdfSetId() string {
	if o == nil || isNil(o.ArdfSetId) {
		var ret string
		return ret
	}
	return *o.ArdfSetId
}

// GetArdfSetIdOk returns a tuple with the ArdfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetArdfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.ArdfSetId) {
    return nil, false
	}
	return o.ArdfSetId, true
}

// HasArdfSetId returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasArdfSetId() bool {
	if o != nil && !isNil(o.ArdfSetId) {
		return true
	}

	return false
}

// SetArdfSetId gets a reference to the given string and assigns it to the ArdfSetId field.
func (o *NdccfDataSubscription) SetArdfSetId(v string) {
	o.ArdfSetId = &v
}

// GetTimePeriod returns the TimePeriod field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetTimePeriod() TimeWindow {
	if o == nil || isNil(o.TimePeriod) {
		var ret TimeWindow
		return ret
	}
	return *o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetTimePeriodOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.TimePeriod) {
    return nil, false
	}
	return o.TimePeriod, true
}

// HasTimePeriod returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasTimePeriod() bool {
	if o != nil && !isNil(o.TimePeriod) {
		return true
	}

	return false
}

// SetTimePeriod gets a reference to the given TimeWindow and assigns it to the TimePeriod field.
func (o *NdccfDataSubscription) SetTimePeriod(v TimeWindow) {
	o.TimePeriod = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *NdccfDataSubscription) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetDataCollectPurposes returns the DataCollectPurposes field value if set, zero value otherwise.
func (o *NdccfDataSubscription) GetDataCollectPurposes() []DataCollectionPurpose {
	if o == nil || isNil(o.DataCollectPurposes) {
		var ret []DataCollectionPurpose
		return ret
	}
	return o.DataCollectPurposes
}

// GetDataCollectPurposesOk returns a tuple with the DataCollectPurposes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfDataSubscription) GetDataCollectPurposesOk() ([]DataCollectionPurpose, bool) {
	if o == nil || isNil(o.DataCollectPurposes) {
    return nil, false
	}
	return o.DataCollectPurposes, true
}

// HasDataCollectPurposes returns a boolean if a field has been set.
func (o *NdccfDataSubscription) HasDataCollectPurposes() bool {
	if o != nil && !isNil(o.DataCollectPurposes) {
		return true
	}

	return false
}

// SetDataCollectPurposes gets a reference to the given []DataCollectionPurpose and assigns it to the DataCollectPurposes field.
func (o *NdccfDataSubscription) SetDataCollectPurposes(v []DataCollectionPurpose) {
	o.DataCollectPurposes = v
}

func (o NdccfDataSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSub"] = o.DataSub
	}
	if true {
		toSerialize["dataNotifUri"] = o.DataNotifUri
	}
	if true {
		toSerialize["dataNotifCorrId"] = o.DataNotifCorrId
	}
	if !isNil(o.FormatInstruct) {
		toSerialize["formatInstruct"] = o.FormatInstruct
	}
	if !isNil(o.ProcInstructs) {
		toSerialize["procInstructs"] = o.ProcInstructs
	}
	if !isNil(o.TargetNfId) {
		toSerialize["targetNfId"] = o.TargetNfId
	}
	if !isNil(o.TargetNfSetId) {
		toSerialize["targetNfSetId"] = o.TargetNfSetId
	}
	if !isNil(o.AdrfId) {
		toSerialize["adrfId"] = o.AdrfId
	}
	if !isNil(o.ArdfSetId) {
		toSerialize["ardfSetId"] = o.ArdfSetId
	}
	if !isNil(o.TimePeriod) {
		toSerialize["timePeriod"] = o.TimePeriod
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !isNil(o.DataCollectPurposes) {
		toSerialize["dataCollectPurposes"] = o.DataCollectPurposes
	}
	return json.Marshal(toSerialize)
}

type NullableNdccfDataSubscription struct {
	value *NdccfDataSubscription
	isSet bool
}

func (v NullableNdccfDataSubscription) Get() *NdccfDataSubscription {
	return v.value
}

func (v *NullableNdccfDataSubscription) Set(val *NdccfDataSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableNdccfDataSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableNdccfDataSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNdccfDataSubscription(val *NdccfDataSubscription) *NullableNdccfDataSubscription {
	return &NullableNdccfDataSubscription{value: val, isSet: true}
}

func (v NullableNdccfDataSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNdccfDataSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


