/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_UECM

import (
	"encoding/json"
	"time"
)

// AmfNon3GppAccessRegistration struct for AmfNon3GppAccessRegistration
type AmfNon3GppAccessRegistration struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfInstanceId string `json:"amfInstanceId"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	PurgeFlag *bool `json:"purgeFlag,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	Pei *string `json:"pei,omitempty"`
	ImsVoPs ImsVoPs `json:"imsVoPs"`
	// String providing an URI formatted according to RFC 3986.
	DeregCallbackUri string `json:"deregCallbackUri"`
	AmfServiceNameDereg *ServiceName `json:"amfServiceNameDereg,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	PcscfRestorationCallbackUri *string `json:"pcscfRestorationCallbackUri,omitempty"`
	AmfServiceNamePcscfRest *ServiceName `json:"amfServiceNamePcscfRest,omitempty"`
	Guami Guami `json:"guami"`
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	RatType RatType `json:"ratType"`
	UrrpIndicator *bool `json:"urrpIndicator,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AmfEeSubscriptionId *string `json:"amfEeSubscriptionId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTime *time.Time `json:"registrationTime,omitempty"`
	VgmlcAddress *VgmlcAddress `json:"vgmlcAddress,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	NoEeSubscriptionInd *bool `json:"noEeSubscriptionInd,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	ReRegistrationRequired *bool `json:"reRegistrationRequired,omitempty"`
	AdminDeregSubWithdrawn *bool `json:"adminDeregSubWithdrawn,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri *string `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
	DisasterRoamingInd *bool `json:"disasterRoamingInd,omitempty"`
	SorSnpnSiSupported *bool `json:"sorSnpnSiSupported,omitempty"`
	UdrRestartInd *bool `json:"udrRestartInd,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LastSynchronizationTime *time.Time `json:"lastSynchronizationTime,omitempty"`
}

// NewAmfNon3GppAccessRegistration instantiates a new AmfNon3GppAccessRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfNon3GppAccessRegistration(amfInstanceId string, imsVoPs ImsVoPs, deregCallbackUri string, guami Guami, ratType RatType) *AmfNon3GppAccessRegistration {
	this := AmfNon3GppAccessRegistration{}
	this.AmfInstanceId = amfInstanceId
	this.ImsVoPs = imsVoPs
	this.DeregCallbackUri = deregCallbackUri
	this.Guami = guami
	this.RatType = ratType
	var disasterRoamingInd bool = false
	this.DisasterRoamingInd = &disasterRoamingInd
	var sorSnpnSiSupported bool = false
	this.SorSnpnSiSupported = &sorSnpnSiSupported
	var udrRestartInd bool = false
	this.UdrRestartInd = &udrRestartInd
	return &this
}

// NewAmfNon3GppAccessRegistrationWithDefaults instantiates a new AmfNon3GppAccessRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfNon3GppAccessRegistrationWithDefaults() *AmfNon3GppAccessRegistration {
	this := AmfNon3GppAccessRegistration{}
	var disasterRoamingInd bool = false
	this.DisasterRoamingInd = &disasterRoamingInd
	var sorSnpnSiSupported bool = false
	this.SorSnpnSiSupported = &sorSnpnSiSupported
	var udrRestartInd bool = false
	this.UdrRestartInd = &udrRestartInd
	return &this
}

// GetAmfInstanceId returns the AmfInstanceId field value
func (o *AmfNon3GppAccessRegistration) GetAmfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfInstanceId
}

// GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field value
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetAmfInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AmfInstanceId, true
}

// SetAmfInstanceId sets field value
func (o *AmfNon3GppAccessRegistration) SetAmfInstanceId(v string) {
	o.AmfInstanceId = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AmfNon3GppAccessRegistration) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetPurgeFlag returns the PurgeFlag field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetPurgeFlag() bool {
	if o == nil || isNil(o.PurgeFlag) {
		var ret bool
		return ret
	}
	return *o.PurgeFlag
}

// GetPurgeFlagOk returns a tuple with the PurgeFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetPurgeFlagOk() (*bool, bool) {
	if o == nil || isNil(o.PurgeFlag) {
    return nil, false
	}
	return o.PurgeFlag, true
}

// HasPurgeFlag returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasPurgeFlag() bool {
	if o != nil && !isNil(o.PurgeFlag) {
		return true
	}

	return false
}

// SetPurgeFlag gets a reference to the given bool and assigns it to the PurgeFlag field.
func (o *AmfNon3GppAccessRegistration) SetPurgeFlag(v bool) {
	o.PurgeFlag = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetPei() string {
	if o == nil || isNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetPeiOk() (*string, bool) {
	if o == nil || isNil(o.Pei) {
    return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasPei() bool {
	if o != nil && !isNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *AmfNon3GppAccessRegistration) SetPei(v string) {
	o.Pei = &v
}

// GetImsVoPs returns the ImsVoPs field value
func (o *AmfNon3GppAccessRegistration) GetImsVoPs() ImsVoPs {
	if o == nil {
		var ret ImsVoPs
		return ret
	}

	return o.ImsVoPs
}

// GetImsVoPsOk returns a tuple with the ImsVoPs field value
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetImsVoPsOk() (*ImsVoPs, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ImsVoPs, true
}

// SetImsVoPs sets field value
func (o *AmfNon3GppAccessRegistration) SetImsVoPs(v ImsVoPs) {
	o.ImsVoPs = v
}

// GetDeregCallbackUri returns the DeregCallbackUri field value
func (o *AmfNon3GppAccessRegistration) GetDeregCallbackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeregCallbackUri
}

// GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field value
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetDeregCallbackUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeregCallbackUri, true
}

// SetDeregCallbackUri sets field value
func (o *AmfNon3GppAccessRegistration) SetDeregCallbackUri(v string) {
	o.DeregCallbackUri = v
}

// GetAmfServiceNameDereg returns the AmfServiceNameDereg field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetAmfServiceNameDereg() ServiceName {
	if o == nil || isNil(o.AmfServiceNameDereg) {
		var ret ServiceName
		return ret
	}
	return *o.AmfServiceNameDereg
}

// GetAmfServiceNameDeregOk returns a tuple with the AmfServiceNameDereg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetAmfServiceNameDeregOk() (*ServiceName, bool) {
	if o == nil || isNil(o.AmfServiceNameDereg) {
    return nil, false
	}
	return o.AmfServiceNameDereg, true
}

// HasAmfServiceNameDereg returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasAmfServiceNameDereg() bool {
	if o != nil && !isNil(o.AmfServiceNameDereg) {
		return true
	}

	return false
}

// SetAmfServiceNameDereg gets a reference to the given ServiceName and assigns it to the AmfServiceNameDereg field.
func (o *AmfNon3GppAccessRegistration) SetAmfServiceNameDereg(v ServiceName) {
	o.AmfServiceNameDereg = &v
}

// GetPcscfRestorationCallbackUri returns the PcscfRestorationCallbackUri field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetPcscfRestorationCallbackUri() string {
	if o == nil || isNil(o.PcscfRestorationCallbackUri) {
		var ret string
		return ret
	}
	return *o.PcscfRestorationCallbackUri
}

// GetPcscfRestorationCallbackUriOk returns a tuple with the PcscfRestorationCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetPcscfRestorationCallbackUriOk() (*string, bool) {
	if o == nil || isNil(o.PcscfRestorationCallbackUri) {
    return nil, false
	}
	return o.PcscfRestorationCallbackUri, true
}

// HasPcscfRestorationCallbackUri returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasPcscfRestorationCallbackUri() bool {
	if o != nil && !isNil(o.PcscfRestorationCallbackUri) {
		return true
	}

	return false
}

// SetPcscfRestorationCallbackUri gets a reference to the given string and assigns it to the PcscfRestorationCallbackUri field.
func (o *AmfNon3GppAccessRegistration) SetPcscfRestorationCallbackUri(v string) {
	o.PcscfRestorationCallbackUri = &v
}

// GetAmfServiceNamePcscfRest returns the AmfServiceNamePcscfRest field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetAmfServiceNamePcscfRest() ServiceName {
	if o == nil || isNil(o.AmfServiceNamePcscfRest) {
		var ret ServiceName
		return ret
	}
	return *o.AmfServiceNamePcscfRest
}

// GetAmfServiceNamePcscfRestOk returns a tuple with the AmfServiceNamePcscfRest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetAmfServiceNamePcscfRestOk() (*ServiceName, bool) {
	if o == nil || isNil(o.AmfServiceNamePcscfRest) {
    return nil, false
	}
	return o.AmfServiceNamePcscfRest, true
}

// HasAmfServiceNamePcscfRest returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasAmfServiceNamePcscfRest() bool {
	if o != nil && !isNil(o.AmfServiceNamePcscfRest) {
		return true
	}

	return false
}

// SetAmfServiceNamePcscfRest gets a reference to the given ServiceName and assigns it to the AmfServiceNamePcscfRest field.
func (o *AmfNon3GppAccessRegistration) SetAmfServiceNamePcscfRest(v ServiceName) {
	o.AmfServiceNamePcscfRest = &v
}

// GetGuami returns the Guami field value
func (o *AmfNon3GppAccessRegistration) GetGuami() Guami {
	if o == nil {
		var ret Guami
		return ret
	}

	return o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetGuamiOk() (*Guami, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Guami, true
}

// SetGuami sets field value
func (o *AmfNon3GppAccessRegistration) SetGuami(v Guami) {
	o.Guami = v
}

// GetBackupAmfInfo returns the BackupAmfInfo field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetBackupAmfInfo() []BackupAmfInfo {
	if o == nil || isNil(o.BackupAmfInfo) {
		var ret []BackupAmfInfo
		return ret
	}
	return o.BackupAmfInfo
}

// GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetBackupAmfInfoOk() ([]BackupAmfInfo, bool) {
	if o == nil || isNil(o.BackupAmfInfo) {
    return nil, false
	}
	return o.BackupAmfInfo, true
}

// HasBackupAmfInfo returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasBackupAmfInfo() bool {
	if o != nil && !isNil(o.BackupAmfInfo) {
		return true
	}

	return false
}

// SetBackupAmfInfo gets a reference to the given []BackupAmfInfo and assigns it to the BackupAmfInfo field.
func (o *AmfNon3GppAccessRegistration) SetBackupAmfInfo(v []BackupAmfInfo) {
	o.BackupAmfInfo = v
}

// GetRatType returns the RatType field value
func (o *AmfNon3GppAccessRegistration) GetRatType() RatType {
	if o == nil {
		var ret RatType
		return ret
	}

	return o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetRatTypeOk() (*RatType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RatType, true
}

// SetRatType sets field value
func (o *AmfNon3GppAccessRegistration) SetRatType(v RatType) {
	o.RatType = v
}

// GetUrrpIndicator returns the UrrpIndicator field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetUrrpIndicator() bool {
	if o == nil || isNil(o.UrrpIndicator) {
		var ret bool
		return ret
	}
	return *o.UrrpIndicator
}

// GetUrrpIndicatorOk returns a tuple with the UrrpIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetUrrpIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.UrrpIndicator) {
    return nil, false
	}
	return o.UrrpIndicator, true
}

// HasUrrpIndicator returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasUrrpIndicator() bool {
	if o != nil && !isNil(o.UrrpIndicator) {
		return true
	}

	return false
}

// SetUrrpIndicator gets a reference to the given bool and assigns it to the UrrpIndicator field.
func (o *AmfNon3GppAccessRegistration) SetUrrpIndicator(v bool) {
	o.UrrpIndicator = &v
}

// GetAmfEeSubscriptionId returns the AmfEeSubscriptionId field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetAmfEeSubscriptionId() string {
	if o == nil || isNil(o.AmfEeSubscriptionId) {
		var ret string
		return ret
	}
	return *o.AmfEeSubscriptionId
}

// GetAmfEeSubscriptionIdOk returns a tuple with the AmfEeSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetAmfEeSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfEeSubscriptionId) {
    return nil, false
	}
	return o.AmfEeSubscriptionId, true
}

// HasAmfEeSubscriptionId returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasAmfEeSubscriptionId() bool {
	if o != nil && !isNil(o.AmfEeSubscriptionId) {
		return true
	}

	return false
}

// SetAmfEeSubscriptionId gets a reference to the given string and assigns it to the AmfEeSubscriptionId field.
func (o *AmfNon3GppAccessRegistration) SetAmfEeSubscriptionId(v string) {
	o.AmfEeSubscriptionId = &v
}

// GetRegistrationTime returns the RegistrationTime field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetRegistrationTime() time.Time {
	if o == nil || isNil(o.RegistrationTime) {
		var ret time.Time
		return ret
	}
	return *o.RegistrationTime
}

// GetRegistrationTimeOk returns a tuple with the RegistrationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetRegistrationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RegistrationTime) {
    return nil, false
	}
	return o.RegistrationTime, true
}

// HasRegistrationTime returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasRegistrationTime() bool {
	if o != nil && !isNil(o.RegistrationTime) {
		return true
	}

	return false
}

// SetRegistrationTime gets a reference to the given time.Time and assigns it to the RegistrationTime field.
func (o *AmfNon3GppAccessRegistration) SetRegistrationTime(v time.Time) {
	o.RegistrationTime = &v
}

// GetVgmlcAddress returns the VgmlcAddress field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetVgmlcAddress() VgmlcAddress {
	if o == nil || isNil(o.VgmlcAddress) {
		var ret VgmlcAddress
		return ret
	}
	return *o.VgmlcAddress
}

// GetVgmlcAddressOk returns a tuple with the VgmlcAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetVgmlcAddressOk() (*VgmlcAddress, bool) {
	if o == nil || isNil(o.VgmlcAddress) {
    return nil, false
	}
	return o.VgmlcAddress, true
}

// HasVgmlcAddress returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasVgmlcAddress() bool {
	if o != nil && !isNil(o.VgmlcAddress) {
		return true
	}

	return false
}

// SetVgmlcAddress gets a reference to the given VgmlcAddress and assigns it to the VgmlcAddress field.
func (o *AmfNon3GppAccessRegistration) SetVgmlcAddress(v VgmlcAddress) {
	o.VgmlcAddress = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetContextInfo() ContextInfo {
	if o == nil || isNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || isNil(o.ContextInfo) {
    return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasContextInfo() bool {
	if o != nil && !isNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *AmfNon3GppAccessRegistration) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

// GetNoEeSubscriptionInd returns the NoEeSubscriptionInd field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetNoEeSubscriptionInd() bool {
	if o == nil || isNil(o.NoEeSubscriptionInd) {
		var ret bool
		return ret
	}
	return *o.NoEeSubscriptionInd
}

// GetNoEeSubscriptionIndOk returns a tuple with the NoEeSubscriptionInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetNoEeSubscriptionIndOk() (*bool, bool) {
	if o == nil || isNil(o.NoEeSubscriptionInd) {
    return nil, false
	}
	return o.NoEeSubscriptionInd, true
}

// HasNoEeSubscriptionInd returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasNoEeSubscriptionInd() bool {
	if o != nil && !isNil(o.NoEeSubscriptionInd) {
		return true
	}

	return false
}

// SetNoEeSubscriptionInd gets a reference to the given bool and assigns it to the NoEeSubscriptionInd field.
func (o *AmfNon3GppAccessRegistration) SetNoEeSubscriptionInd(v bool) {
	o.NoEeSubscriptionInd = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *AmfNon3GppAccessRegistration) SetSupi(v string) {
	o.Supi = &v
}

// GetReRegistrationRequired returns the ReRegistrationRequired field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetReRegistrationRequired() bool {
	if o == nil || isNil(o.ReRegistrationRequired) {
		var ret bool
		return ret
	}
	return *o.ReRegistrationRequired
}

// GetReRegistrationRequiredOk returns a tuple with the ReRegistrationRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetReRegistrationRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.ReRegistrationRequired) {
    return nil, false
	}
	return o.ReRegistrationRequired, true
}

// HasReRegistrationRequired returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasReRegistrationRequired() bool {
	if o != nil && !isNil(o.ReRegistrationRequired) {
		return true
	}

	return false
}

// SetReRegistrationRequired gets a reference to the given bool and assigns it to the ReRegistrationRequired field.
func (o *AmfNon3GppAccessRegistration) SetReRegistrationRequired(v bool) {
	o.ReRegistrationRequired = &v
}

// GetAdminDeregSubWithdrawn returns the AdminDeregSubWithdrawn field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetAdminDeregSubWithdrawn() bool {
	if o == nil || isNil(o.AdminDeregSubWithdrawn) {
		var ret bool
		return ret
	}
	return *o.AdminDeregSubWithdrawn
}

// GetAdminDeregSubWithdrawnOk returns a tuple with the AdminDeregSubWithdrawn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetAdminDeregSubWithdrawnOk() (*bool, bool) {
	if o == nil || isNil(o.AdminDeregSubWithdrawn) {
    return nil, false
	}
	return o.AdminDeregSubWithdrawn, true
}

// HasAdminDeregSubWithdrawn returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasAdminDeregSubWithdrawn() bool {
	if o != nil && !isNil(o.AdminDeregSubWithdrawn) {
		return true
	}

	return false
}

// SetAdminDeregSubWithdrawn gets a reference to the given bool and assigns it to the AdminDeregSubWithdrawn field.
func (o *AmfNon3GppAccessRegistration) SetAdminDeregSubWithdrawn(v bool) {
	o.AdminDeregSubWithdrawn = &v
}

// GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetDataRestorationCallbackUri() string {
	if o == nil || isNil(o.DataRestorationCallbackUri) {
		var ret string
		return ret
	}
	return *o.DataRestorationCallbackUri
}

// GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetDataRestorationCallbackUriOk() (*string, bool) {
	if o == nil || isNil(o.DataRestorationCallbackUri) {
    return nil, false
	}
	return o.DataRestorationCallbackUri, true
}

// HasDataRestorationCallbackUri returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasDataRestorationCallbackUri() bool {
	if o != nil && !isNil(o.DataRestorationCallbackUri) {
		return true
	}

	return false
}

// SetDataRestorationCallbackUri gets a reference to the given string and assigns it to the DataRestorationCallbackUri field.
func (o *AmfNon3GppAccessRegistration) SetDataRestorationCallbackUri(v string) {
	o.DataRestorationCallbackUri = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *AmfNon3GppAccessRegistration) SetResetIds(v []string) {
	o.ResetIds = v
}

// GetDisasterRoamingInd returns the DisasterRoamingInd field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetDisasterRoamingInd() bool {
	if o == nil || isNil(o.DisasterRoamingInd) {
		var ret bool
		return ret
	}
	return *o.DisasterRoamingInd
}

// GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetDisasterRoamingIndOk() (*bool, bool) {
	if o == nil || isNil(o.DisasterRoamingInd) {
    return nil, false
	}
	return o.DisasterRoamingInd, true
}

// HasDisasterRoamingInd returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasDisasterRoamingInd() bool {
	if o != nil && !isNil(o.DisasterRoamingInd) {
		return true
	}

	return false
}

// SetDisasterRoamingInd gets a reference to the given bool and assigns it to the DisasterRoamingInd field.
func (o *AmfNon3GppAccessRegistration) SetDisasterRoamingInd(v bool) {
	o.DisasterRoamingInd = &v
}

// GetSorSnpnSiSupported returns the SorSnpnSiSupported field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetSorSnpnSiSupported() bool {
	if o == nil || isNil(o.SorSnpnSiSupported) {
		var ret bool
		return ret
	}
	return *o.SorSnpnSiSupported
}

// GetSorSnpnSiSupportedOk returns a tuple with the SorSnpnSiSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetSorSnpnSiSupportedOk() (*bool, bool) {
	if o == nil || isNil(o.SorSnpnSiSupported) {
    return nil, false
	}
	return o.SorSnpnSiSupported, true
}

// HasSorSnpnSiSupported returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasSorSnpnSiSupported() bool {
	if o != nil && !isNil(o.SorSnpnSiSupported) {
		return true
	}

	return false
}

// SetSorSnpnSiSupported gets a reference to the given bool and assigns it to the SorSnpnSiSupported field.
func (o *AmfNon3GppAccessRegistration) SetSorSnpnSiSupported(v bool) {
	o.SorSnpnSiSupported = &v
}

// GetUdrRestartInd returns the UdrRestartInd field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetUdrRestartInd() bool {
	if o == nil || isNil(o.UdrRestartInd) {
		var ret bool
		return ret
	}
	return *o.UdrRestartInd
}

// GetUdrRestartIndOk returns a tuple with the UdrRestartInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetUdrRestartIndOk() (*bool, bool) {
	if o == nil || isNil(o.UdrRestartInd) {
    return nil, false
	}
	return o.UdrRestartInd, true
}

// HasUdrRestartInd returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasUdrRestartInd() bool {
	if o != nil && !isNil(o.UdrRestartInd) {
		return true
	}

	return false
}

// SetUdrRestartInd gets a reference to the given bool and assigns it to the UdrRestartInd field.
func (o *AmfNon3GppAccessRegistration) SetUdrRestartInd(v bool) {
	o.UdrRestartInd = &v
}

// GetLastSynchronizationTime returns the LastSynchronizationTime field value if set, zero value otherwise.
func (o *AmfNon3GppAccessRegistration) GetLastSynchronizationTime() time.Time {
	if o == nil || isNil(o.LastSynchronizationTime) {
		var ret time.Time
		return ret
	}
	return *o.LastSynchronizationTime
}

// GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfNon3GppAccessRegistration) GetLastSynchronizationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastSynchronizationTime) {
    return nil, false
	}
	return o.LastSynchronizationTime, true
}

// HasLastSynchronizationTime returns a boolean if a field has been set.
func (o *AmfNon3GppAccessRegistration) HasLastSynchronizationTime() bool {
	if o != nil && !isNil(o.LastSynchronizationTime) {
		return true
	}

	return false
}

// SetLastSynchronizationTime gets a reference to the given time.Time and assigns it to the LastSynchronizationTime field.
func (o *AmfNon3GppAccessRegistration) SetLastSynchronizationTime(v time.Time) {
	o.LastSynchronizationTime = &v
}

func (o AmfNon3GppAccessRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amfInstanceId"] = o.AmfInstanceId
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.PurgeFlag) {
		toSerialize["purgeFlag"] = o.PurgeFlag
	}
	if !isNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if true {
		toSerialize["imsVoPs"] = o.ImsVoPs
	}
	if true {
		toSerialize["deregCallbackUri"] = o.DeregCallbackUri
	}
	if !isNil(o.AmfServiceNameDereg) {
		toSerialize["amfServiceNameDereg"] = o.AmfServiceNameDereg
	}
	if !isNil(o.PcscfRestorationCallbackUri) {
		toSerialize["pcscfRestorationCallbackUri"] = o.PcscfRestorationCallbackUri
	}
	if !isNil(o.AmfServiceNamePcscfRest) {
		toSerialize["amfServiceNamePcscfRest"] = o.AmfServiceNamePcscfRest
	}
	if true {
		toSerialize["guami"] = o.Guami
	}
	if !isNil(o.BackupAmfInfo) {
		toSerialize["backupAmfInfo"] = o.BackupAmfInfo
	}
	if true {
		toSerialize["ratType"] = o.RatType
	}
	if !isNil(o.UrrpIndicator) {
		toSerialize["urrpIndicator"] = o.UrrpIndicator
	}
	if !isNil(o.AmfEeSubscriptionId) {
		toSerialize["amfEeSubscriptionId"] = o.AmfEeSubscriptionId
	}
	if !isNil(o.RegistrationTime) {
		toSerialize["registrationTime"] = o.RegistrationTime
	}
	if !isNil(o.VgmlcAddress) {
		toSerialize["vgmlcAddress"] = o.VgmlcAddress
	}
	if !isNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	if !isNil(o.NoEeSubscriptionInd) {
		toSerialize["noEeSubscriptionInd"] = o.NoEeSubscriptionInd
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.ReRegistrationRequired) {
		toSerialize["reRegistrationRequired"] = o.ReRegistrationRequired
	}
	if !isNil(o.AdminDeregSubWithdrawn) {
		toSerialize["adminDeregSubWithdrawn"] = o.AdminDeregSubWithdrawn
	}
	if !isNil(o.DataRestorationCallbackUri) {
		toSerialize["dataRestorationCallbackUri"] = o.DataRestorationCallbackUri
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	if !isNil(o.DisasterRoamingInd) {
		toSerialize["disasterRoamingInd"] = o.DisasterRoamingInd
	}
	if !isNil(o.SorSnpnSiSupported) {
		toSerialize["sorSnpnSiSupported"] = o.SorSnpnSiSupported
	}
	if !isNil(o.UdrRestartInd) {
		toSerialize["udrRestartInd"] = o.UdrRestartInd
	}
	if !isNil(o.LastSynchronizationTime) {
		toSerialize["lastSynchronizationTime"] = o.LastSynchronizationTime
	}
	return json.Marshal(toSerialize)
}

type NullableAmfNon3GppAccessRegistration struct {
	value *AmfNon3GppAccessRegistration
	isSet bool
}

func (v NullableAmfNon3GppAccessRegistration) Get() *AmfNon3GppAccessRegistration {
	return v.value
}

func (v *NullableAmfNon3GppAccessRegistration) Set(val *AmfNon3GppAccessRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfNon3GppAccessRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfNon3GppAccessRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfNon3GppAccessRegistration(val *AmfNon3GppAccessRegistration) *NullableAmfNon3GppAccessRegistration {
	return &NullableAmfNon3GppAccessRegistration{value: val, isSet: true}
}

func (v NullableAmfNon3GppAccessRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfNon3GppAccessRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


