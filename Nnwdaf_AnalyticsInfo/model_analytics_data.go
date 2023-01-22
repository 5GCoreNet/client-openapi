/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"time"
)

// AnalyticsData Represents the description of analytics with parameters as relevant for the requesting NF  service consumer. 
type AnalyticsData struct {
	// string with format 'date-time' as defined in OpenAPI.
	Start *time.Time `json:"start,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStampGen *time.Time `json:"timeStampGen,omitempty"`
	AnaMetaInfo *AnalyticsMetadataInfo `json:"anaMetaInfo,omitempty"`
	// The slices and their load level information.
	SliceLoadLevelInfos []SliceLoadLevelInformation `json:"sliceLoadLevelInfos,omitempty"`
	NsiLoadLevelInfos []NsiLoadLevelInfo `json:"nsiLoadLevelInfos,omitempty"`
	NfLoadLevelInfos []NfLoadLevelInformation `json:"nfLoadLevelInfos,omitempty"`
	NwPerfs []NetworkPerfInfo `json:"nwPerfs,omitempty"`
	SvcExps []ServiceExperienceInfo `json:"svcExps,omitempty"`
	QosSustainInfos []QosSustainabilityInfo `json:"qosSustainInfos,omitempty"`
	UeMobs []UeMobility `json:"ueMobs,omitempty"`
	UeComms []UeCommunication `json:"ueComms,omitempty"`
	UserDataCongInfos []UserDataCongestionInfo `json:"userDataCongInfos,omitempty"`
	AbnorBehavrs []AbnormalBehaviour `json:"abnorBehavrs,omitempty"`
	SmccExps []SmcceInfo `json:"smccExps,omitempty"`
	DisperInfos []DispersionInfo `json:"disperInfos,omitempty"`
	RedTransInfos []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
	WlanInfos []WlanPerformanceInfo `json:"wlanInfos,omitempty"`
	DnPerfInfos []DnPerfInfo `json:"dnPerfInfos,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewAnalyticsData instantiates a new AnalyticsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsData() *AnalyticsData {
	this := AnalyticsData{}
	return &this
}

// NewAnalyticsDataWithDefaults instantiates a new AnalyticsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsDataWithDefaults() *AnalyticsData {
	this := AnalyticsData{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *AnalyticsData) GetStart() time.Time {
	if o == nil || isNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *AnalyticsData) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *AnalyticsData) SetStart(v time.Time) {
	o.Start = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *AnalyticsData) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *AnalyticsData) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *AnalyticsData) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetTimeStampGen returns the TimeStampGen field value if set, zero value otherwise.
func (o *AnalyticsData) GetTimeStampGen() time.Time {
	if o == nil || isNil(o.TimeStampGen) {
		var ret time.Time
		return ret
	}
	return *o.TimeStampGen
}

// GetTimeStampGenOk returns a tuple with the TimeStampGen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetTimeStampGenOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeStampGen) {
    return nil, false
	}
	return o.TimeStampGen, true
}

// HasTimeStampGen returns a boolean if a field has been set.
func (o *AnalyticsData) HasTimeStampGen() bool {
	if o != nil && !isNil(o.TimeStampGen) {
		return true
	}

	return false
}

// SetTimeStampGen gets a reference to the given time.Time and assigns it to the TimeStampGen field.
func (o *AnalyticsData) SetTimeStampGen(v time.Time) {
	o.TimeStampGen = &v
}

// GetAnaMetaInfo returns the AnaMetaInfo field value if set, zero value otherwise.
func (o *AnalyticsData) GetAnaMetaInfo() AnalyticsMetadataInfo {
	if o == nil || isNil(o.AnaMetaInfo) {
		var ret AnalyticsMetadataInfo
		return ret
	}
	return *o.AnaMetaInfo
}

// GetAnaMetaInfoOk returns a tuple with the AnaMetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetAnaMetaInfoOk() (*AnalyticsMetadataInfo, bool) {
	if o == nil || isNil(o.AnaMetaInfo) {
    return nil, false
	}
	return o.AnaMetaInfo, true
}

// HasAnaMetaInfo returns a boolean if a field has been set.
func (o *AnalyticsData) HasAnaMetaInfo() bool {
	if o != nil && !isNil(o.AnaMetaInfo) {
		return true
	}

	return false
}

// SetAnaMetaInfo gets a reference to the given AnalyticsMetadataInfo and assigns it to the AnaMetaInfo field.
func (o *AnalyticsData) SetAnaMetaInfo(v AnalyticsMetadataInfo) {
	o.AnaMetaInfo = &v
}

// GetSliceLoadLevelInfos returns the SliceLoadLevelInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetSliceLoadLevelInfos() []SliceLoadLevelInformation {
	if o == nil || isNil(o.SliceLoadLevelInfos) {
		var ret []SliceLoadLevelInformation
		return ret
	}
	return o.SliceLoadLevelInfos
}

// GetSliceLoadLevelInfosOk returns a tuple with the SliceLoadLevelInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetSliceLoadLevelInfosOk() ([]SliceLoadLevelInformation, bool) {
	if o == nil || isNil(o.SliceLoadLevelInfos) {
    return nil, false
	}
	return o.SliceLoadLevelInfos, true
}

// HasSliceLoadLevelInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasSliceLoadLevelInfos() bool {
	if o != nil && !isNil(o.SliceLoadLevelInfos) {
		return true
	}

	return false
}

// SetSliceLoadLevelInfos gets a reference to the given []SliceLoadLevelInformation and assigns it to the SliceLoadLevelInfos field.
func (o *AnalyticsData) SetSliceLoadLevelInfos(v []SliceLoadLevelInformation) {
	o.SliceLoadLevelInfos = v
}

// GetNsiLoadLevelInfos returns the NsiLoadLevelInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetNsiLoadLevelInfos() []NsiLoadLevelInfo {
	if o == nil || isNil(o.NsiLoadLevelInfos) {
		var ret []NsiLoadLevelInfo
		return ret
	}
	return o.NsiLoadLevelInfos
}

// GetNsiLoadLevelInfosOk returns a tuple with the NsiLoadLevelInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetNsiLoadLevelInfosOk() ([]NsiLoadLevelInfo, bool) {
	if o == nil || isNil(o.NsiLoadLevelInfos) {
    return nil, false
	}
	return o.NsiLoadLevelInfos, true
}

// HasNsiLoadLevelInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasNsiLoadLevelInfos() bool {
	if o != nil && !isNil(o.NsiLoadLevelInfos) {
		return true
	}

	return false
}

// SetNsiLoadLevelInfos gets a reference to the given []NsiLoadLevelInfo and assigns it to the NsiLoadLevelInfos field.
func (o *AnalyticsData) SetNsiLoadLevelInfos(v []NsiLoadLevelInfo) {
	o.NsiLoadLevelInfos = v
}

// GetNfLoadLevelInfos returns the NfLoadLevelInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetNfLoadLevelInfos() []NfLoadLevelInformation {
	if o == nil || isNil(o.NfLoadLevelInfos) {
		var ret []NfLoadLevelInformation
		return ret
	}
	return o.NfLoadLevelInfos
}

// GetNfLoadLevelInfosOk returns a tuple with the NfLoadLevelInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetNfLoadLevelInfosOk() ([]NfLoadLevelInformation, bool) {
	if o == nil || isNil(o.NfLoadLevelInfos) {
    return nil, false
	}
	return o.NfLoadLevelInfos, true
}

// HasNfLoadLevelInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasNfLoadLevelInfos() bool {
	if o != nil && !isNil(o.NfLoadLevelInfos) {
		return true
	}

	return false
}

// SetNfLoadLevelInfos gets a reference to the given []NfLoadLevelInformation and assigns it to the NfLoadLevelInfos field.
func (o *AnalyticsData) SetNfLoadLevelInfos(v []NfLoadLevelInformation) {
	o.NfLoadLevelInfos = v
}

// GetNwPerfs returns the NwPerfs field value if set, zero value otherwise.
func (o *AnalyticsData) GetNwPerfs() []NetworkPerfInfo {
	if o == nil || isNil(o.NwPerfs) {
		var ret []NetworkPerfInfo
		return ret
	}
	return o.NwPerfs
}

// GetNwPerfsOk returns a tuple with the NwPerfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetNwPerfsOk() ([]NetworkPerfInfo, bool) {
	if o == nil || isNil(o.NwPerfs) {
    return nil, false
	}
	return o.NwPerfs, true
}

// HasNwPerfs returns a boolean if a field has been set.
func (o *AnalyticsData) HasNwPerfs() bool {
	if o != nil && !isNil(o.NwPerfs) {
		return true
	}

	return false
}

// SetNwPerfs gets a reference to the given []NetworkPerfInfo and assigns it to the NwPerfs field.
func (o *AnalyticsData) SetNwPerfs(v []NetworkPerfInfo) {
	o.NwPerfs = v
}

// GetSvcExps returns the SvcExps field value if set, zero value otherwise.
func (o *AnalyticsData) GetSvcExps() []ServiceExperienceInfo {
	if o == nil || isNil(o.SvcExps) {
		var ret []ServiceExperienceInfo
		return ret
	}
	return o.SvcExps
}

// GetSvcExpsOk returns a tuple with the SvcExps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetSvcExpsOk() ([]ServiceExperienceInfo, bool) {
	if o == nil || isNil(o.SvcExps) {
    return nil, false
	}
	return o.SvcExps, true
}

// HasSvcExps returns a boolean if a field has been set.
func (o *AnalyticsData) HasSvcExps() bool {
	if o != nil && !isNil(o.SvcExps) {
		return true
	}

	return false
}

// SetSvcExps gets a reference to the given []ServiceExperienceInfo and assigns it to the SvcExps field.
func (o *AnalyticsData) SetSvcExps(v []ServiceExperienceInfo) {
	o.SvcExps = v
}

// GetQosSustainInfos returns the QosSustainInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetQosSustainInfos() []QosSustainabilityInfo {
	if o == nil || isNil(o.QosSustainInfos) {
		var ret []QosSustainabilityInfo
		return ret
	}
	return o.QosSustainInfos
}

// GetQosSustainInfosOk returns a tuple with the QosSustainInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetQosSustainInfosOk() ([]QosSustainabilityInfo, bool) {
	if o == nil || isNil(o.QosSustainInfos) {
    return nil, false
	}
	return o.QosSustainInfos, true
}

// HasQosSustainInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasQosSustainInfos() bool {
	if o != nil && !isNil(o.QosSustainInfos) {
		return true
	}

	return false
}

// SetQosSustainInfos gets a reference to the given []QosSustainabilityInfo and assigns it to the QosSustainInfos field.
func (o *AnalyticsData) SetQosSustainInfos(v []QosSustainabilityInfo) {
	o.QosSustainInfos = v
}

// GetUeMobs returns the UeMobs field value if set, zero value otherwise.
func (o *AnalyticsData) GetUeMobs() []UeMobility {
	if o == nil || isNil(o.UeMobs) {
		var ret []UeMobility
		return ret
	}
	return o.UeMobs
}

// GetUeMobsOk returns a tuple with the UeMobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetUeMobsOk() ([]UeMobility, bool) {
	if o == nil || isNil(o.UeMobs) {
    return nil, false
	}
	return o.UeMobs, true
}

// HasUeMobs returns a boolean if a field has been set.
func (o *AnalyticsData) HasUeMobs() bool {
	if o != nil && !isNil(o.UeMobs) {
		return true
	}

	return false
}

// SetUeMobs gets a reference to the given []UeMobility and assigns it to the UeMobs field.
func (o *AnalyticsData) SetUeMobs(v []UeMobility) {
	o.UeMobs = v
}

// GetUeComms returns the UeComms field value if set, zero value otherwise.
func (o *AnalyticsData) GetUeComms() []UeCommunication {
	if o == nil || isNil(o.UeComms) {
		var ret []UeCommunication
		return ret
	}
	return o.UeComms
}

// GetUeCommsOk returns a tuple with the UeComms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetUeCommsOk() ([]UeCommunication, bool) {
	if o == nil || isNil(o.UeComms) {
    return nil, false
	}
	return o.UeComms, true
}

// HasUeComms returns a boolean if a field has been set.
func (o *AnalyticsData) HasUeComms() bool {
	if o != nil && !isNil(o.UeComms) {
		return true
	}

	return false
}

// SetUeComms gets a reference to the given []UeCommunication and assigns it to the UeComms field.
func (o *AnalyticsData) SetUeComms(v []UeCommunication) {
	o.UeComms = v
}

// GetUserDataCongInfos returns the UserDataCongInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetUserDataCongInfos() []UserDataCongestionInfo {
	if o == nil || isNil(o.UserDataCongInfos) {
		var ret []UserDataCongestionInfo
		return ret
	}
	return o.UserDataCongInfos
}

// GetUserDataCongInfosOk returns a tuple with the UserDataCongInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetUserDataCongInfosOk() ([]UserDataCongestionInfo, bool) {
	if o == nil || isNil(o.UserDataCongInfos) {
    return nil, false
	}
	return o.UserDataCongInfos, true
}

// HasUserDataCongInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasUserDataCongInfos() bool {
	if o != nil && !isNil(o.UserDataCongInfos) {
		return true
	}

	return false
}

// SetUserDataCongInfos gets a reference to the given []UserDataCongestionInfo and assigns it to the UserDataCongInfos field.
func (o *AnalyticsData) SetUserDataCongInfos(v []UserDataCongestionInfo) {
	o.UserDataCongInfos = v
}

// GetAbnorBehavrs returns the AbnorBehavrs field value if set, zero value otherwise.
func (o *AnalyticsData) GetAbnorBehavrs() []AbnormalBehaviour {
	if o == nil || isNil(o.AbnorBehavrs) {
		var ret []AbnormalBehaviour
		return ret
	}
	return o.AbnorBehavrs
}

// GetAbnorBehavrsOk returns a tuple with the AbnorBehavrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetAbnorBehavrsOk() ([]AbnormalBehaviour, bool) {
	if o == nil || isNil(o.AbnorBehavrs) {
    return nil, false
	}
	return o.AbnorBehavrs, true
}

// HasAbnorBehavrs returns a boolean if a field has been set.
func (o *AnalyticsData) HasAbnorBehavrs() bool {
	if o != nil && !isNil(o.AbnorBehavrs) {
		return true
	}

	return false
}

// SetAbnorBehavrs gets a reference to the given []AbnormalBehaviour and assigns it to the AbnorBehavrs field.
func (o *AnalyticsData) SetAbnorBehavrs(v []AbnormalBehaviour) {
	o.AbnorBehavrs = v
}

// GetSmccExps returns the SmccExps field value if set, zero value otherwise.
func (o *AnalyticsData) GetSmccExps() []SmcceInfo {
	if o == nil || isNil(o.SmccExps) {
		var ret []SmcceInfo
		return ret
	}
	return o.SmccExps
}

// GetSmccExpsOk returns a tuple with the SmccExps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetSmccExpsOk() ([]SmcceInfo, bool) {
	if o == nil || isNil(o.SmccExps) {
    return nil, false
	}
	return o.SmccExps, true
}

// HasSmccExps returns a boolean if a field has been set.
func (o *AnalyticsData) HasSmccExps() bool {
	if o != nil && !isNil(o.SmccExps) {
		return true
	}

	return false
}

// SetSmccExps gets a reference to the given []SmcceInfo and assigns it to the SmccExps field.
func (o *AnalyticsData) SetSmccExps(v []SmcceInfo) {
	o.SmccExps = v
}

// GetDisperInfos returns the DisperInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetDisperInfos() []DispersionInfo {
	if o == nil || isNil(o.DisperInfos) {
		var ret []DispersionInfo
		return ret
	}
	return o.DisperInfos
}

// GetDisperInfosOk returns a tuple with the DisperInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetDisperInfosOk() ([]DispersionInfo, bool) {
	if o == nil || isNil(o.DisperInfos) {
    return nil, false
	}
	return o.DisperInfos, true
}

// HasDisperInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasDisperInfos() bool {
	if o != nil && !isNil(o.DisperInfos) {
		return true
	}

	return false
}

// SetDisperInfos gets a reference to the given []DispersionInfo and assigns it to the DisperInfos field.
func (o *AnalyticsData) SetDisperInfos(v []DispersionInfo) {
	o.DisperInfos = v
}

// GetRedTransInfos returns the RedTransInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetRedTransInfos() []RedundantTransmissionExpInfo {
	if o == nil || isNil(o.RedTransInfos) {
		var ret []RedundantTransmissionExpInfo
		return ret
	}
	return o.RedTransInfos
}

// GetRedTransInfosOk returns a tuple with the RedTransInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetRedTransInfosOk() ([]RedundantTransmissionExpInfo, bool) {
	if o == nil || isNil(o.RedTransInfos) {
    return nil, false
	}
	return o.RedTransInfos, true
}

// HasRedTransInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasRedTransInfos() bool {
	if o != nil && !isNil(o.RedTransInfos) {
		return true
	}

	return false
}

// SetRedTransInfos gets a reference to the given []RedundantTransmissionExpInfo and assigns it to the RedTransInfos field.
func (o *AnalyticsData) SetRedTransInfos(v []RedundantTransmissionExpInfo) {
	o.RedTransInfos = v
}

// GetWlanInfos returns the WlanInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetWlanInfos() []WlanPerformanceInfo {
	if o == nil || isNil(o.WlanInfos) {
		var ret []WlanPerformanceInfo
		return ret
	}
	return o.WlanInfos
}

// GetWlanInfosOk returns a tuple with the WlanInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetWlanInfosOk() ([]WlanPerformanceInfo, bool) {
	if o == nil || isNil(o.WlanInfos) {
    return nil, false
	}
	return o.WlanInfos, true
}

// HasWlanInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasWlanInfos() bool {
	if o != nil && !isNil(o.WlanInfos) {
		return true
	}

	return false
}

// SetWlanInfos gets a reference to the given []WlanPerformanceInfo and assigns it to the WlanInfos field.
func (o *AnalyticsData) SetWlanInfos(v []WlanPerformanceInfo) {
	o.WlanInfos = v
}

// GetDnPerfInfos returns the DnPerfInfos field value if set, zero value otherwise.
func (o *AnalyticsData) GetDnPerfInfos() []DnPerfInfo {
	if o == nil || isNil(o.DnPerfInfos) {
		var ret []DnPerfInfo
		return ret
	}
	return o.DnPerfInfos
}

// GetDnPerfInfosOk returns a tuple with the DnPerfInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetDnPerfInfosOk() ([]DnPerfInfo, bool) {
	if o == nil || isNil(o.DnPerfInfos) {
    return nil, false
	}
	return o.DnPerfInfos, true
}

// HasDnPerfInfos returns a boolean if a field has been set.
func (o *AnalyticsData) HasDnPerfInfos() bool {
	if o != nil && !isNil(o.DnPerfInfos) {
		return true
	}

	return false
}

// SetDnPerfInfos gets a reference to the given []DnPerfInfo and assigns it to the DnPerfInfos field.
func (o *AnalyticsData) SetDnPerfInfos(v []DnPerfInfo) {
	o.DnPerfInfos = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *AnalyticsData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *AnalyticsData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *AnalyticsData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o AnalyticsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !isNil(o.TimeStampGen) {
		toSerialize["timeStampGen"] = o.TimeStampGen
	}
	if !isNil(o.AnaMetaInfo) {
		toSerialize["anaMetaInfo"] = o.AnaMetaInfo
	}
	if !isNil(o.SliceLoadLevelInfos) {
		toSerialize["sliceLoadLevelInfos"] = o.SliceLoadLevelInfos
	}
	if !isNil(o.NsiLoadLevelInfos) {
		toSerialize["nsiLoadLevelInfos"] = o.NsiLoadLevelInfos
	}
	if !isNil(o.NfLoadLevelInfos) {
		toSerialize["nfLoadLevelInfos"] = o.NfLoadLevelInfos
	}
	if !isNil(o.NwPerfs) {
		toSerialize["nwPerfs"] = o.NwPerfs
	}
	if !isNil(o.SvcExps) {
		toSerialize["svcExps"] = o.SvcExps
	}
	if !isNil(o.QosSustainInfos) {
		toSerialize["qosSustainInfos"] = o.QosSustainInfos
	}
	if !isNil(o.UeMobs) {
		toSerialize["ueMobs"] = o.UeMobs
	}
	if !isNil(o.UeComms) {
		toSerialize["ueComms"] = o.UeComms
	}
	if !isNil(o.UserDataCongInfos) {
		toSerialize["userDataCongInfos"] = o.UserDataCongInfos
	}
	if !isNil(o.AbnorBehavrs) {
		toSerialize["abnorBehavrs"] = o.AbnorBehavrs
	}
	if !isNil(o.SmccExps) {
		toSerialize["smccExps"] = o.SmccExps
	}
	if !isNil(o.DisperInfos) {
		toSerialize["disperInfos"] = o.DisperInfos
	}
	if !isNil(o.RedTransInfos) {
		toSerialize["redTransInfos"] = o.RedTransInfos
	}
	if !isNil(o.WlanInfos) {
		toSerialize["wlanInfos"] = o.WlanInfos
	}
	if !isNil(o.DnPerfInfos) {
		toSerialize["dnPerfInfos"] = o.DnPerfInfos
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableAnalyticsData struct {
	value *AnalyticsData
	isSet bool
}

func (v NullableAnalyticsData) Get() *AnalyticsData {
	return v.value
}

func (v *NullableAnalyticsData) Set(val *AnalyticsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsData(val *AnalyticsData) *NullableAnalyticsData {
	return &NullableAnalyticsData{value: val, isSet: true}
}

func (v NullableAnalyticsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


