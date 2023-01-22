/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// SMSChargingInformation struct for SMSChargingInformation
type SMSChargingInformation struct {
	OriginatorInfo *OriginatorInfo `json:"originatorInfo,omitempty"`
	RecipientInfo []RecipientInfo `json:"recipientInfo,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	UserEquipmentInfo *string `json:"userEquipmentInfo,omitempty"`
	RoamerInOut *RoamerInOut `json:"roamerInOut,omitempty"`
	UserLocationinfo *UserLocation `json:"userLocationinfo,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UetimeZone *string `json:"uetimeZone,omitempty"`
	RATType *RatType `json:"rATType,omitempty"`
	SMSCAddress *string `json:"sMSCAddress,omitempty"`
	SMDataCodingScheme *int32 `json:"sMDataCodingScheme,omitempty"`
	SMMessageType *SMMessageType `json:"sMMessageType,omitempty"`
	SMReplyPathRequested *ReplyPathRequested `json:"sMReplyPathRequested,omitempty"`
	SMUserDataHeader *string `json:"sMUserDataHeader,omitempty"`
	SMStatus *string `json:"sMStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SMDischargeTime *time.Time `json:"sMDischargeTime,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	NumberofMessagesSent *int32 `json:"numberofMessagesSent,omitempty"`
	SMServiceType *SMServiceType `json:"sMServiceType,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	SMSequenceNumber *int32 `json:"sMSequenceNumber,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	SMSresult *int32 `json:"sMSresult,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SubmissionTime *time.Time `json:"submissionTime,omitempty"`
	SMPriority *SMPriority `json:"sMPriority,omitempty"`
	MessageReference *string `json:"messageReference,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	MessageSize *int32 `json:"messageSize,omitempty"`
	MessageClass *MessageClass `json:"messageClass,omitempty"`
	DeliveryReportRequested *DeliveryReportRequested `json:"deliveryReportRequested,omitempty"`
}

// NewSMSChargingInformation instantiates a new SMSChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSChargingInformation() *SMSChargingInformation {
	this := SMSChargingInformation{}
	return &this
}

// NewSMSChargingInformationWithDefaults instantiates a new SMSChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSChargingInformationWithDefaults() *SMSChargingInformation {
	this := SMSChargingInformation{}
	return &this
}

// GetOriginatorInfo returns the OriginatorInfo field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetOriginatorInfo() OriginatorInfo {
	if o == nil || isNil(o.OriginatorInfo) {
		var ret OriginatorInfo
		return ret
	}
	return *o.OriginatorInfo
}

// GetOriginatorInfoOk returns a tuple with the OriginatorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetOriginatorInfoOk() (*OriginatorInfo, bool) {
	if o == nil || isNil(o.OriginatorInfo) {
    return nil, false
	}
	return o.OriginatorInfo, true
}

// HasOriginatorInfo returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasOriginatorInfo() bool {
	if o != nil && !isNil(o.OriginatorInfo) {
		return true
	}

	return false
}

// SetOriginatorInfo gets a reference to the given OriginatorInfo and assigns it to the OriginatorInfo field.
func (o *SMSChargingInformation) SetOriginatorInfo(v OriginatorInfo) {
	o.OriginatorInfo = &v
}

// GetRecipientInfo returns the RecipientInfo field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetRecipientInfo() []RecipientInfo {
	if o == nil || isNil(o.RecipientInfo) {
		var ret []RecipientInfo
		return ret
	}
	return o.RecipientInfo
}

// GetRecipientInfoOk returns a tuple with the RecipientInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetRecipientInfoOk() ([]RecipientInfo, bool) {
	if o == nil || isNil(o.RecipientInfo) {
    return nil, false
	}
	return o.RecipientInfo, true
}

// HasRecipientInfo returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasRecipientInfo() bool {
	if o != nil && !isNil(o.RecipientInfo) {
		return true
	}

	return false
}

// SetRecipientInfo gets a reference to the given []RecipientInfo and assigns it to the RecipientInfo field.
func (o *SMSChargingInformation) SetRecipientInfo(v []RecipientInfo) {
	o.RecipientInfo = v
}

// GetUserEquipmentInfo returns the UserEquipmentInfo field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetUserEquipmentInfo() string {
	if o == nil || isNil(o.UserEquipmentInfo) {
		var ret string
		return ret
	}
	return *o.UserEquipmentInfo
}

// GetUserEquipmentInfoOk returns a tuple with the UserEquipmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetUserEquipmentInfoOk() (*string, bool) {
	if o == nil || isNil(o.UserEquipmentInfo) {
    return nil, false
	}
	return o.UserEquipmentInfo, true
}

// HasUserEquipmentInfo returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasUserEquipmentInfo() bool {
	if o != nil && !isNil(o.UserEquipmentInfo) {
		return true
	}

	return false
}

// SetUserEquipmentInfo gets a reference to the given string and assigns it to the UserEquipmentInfo field.
func (o *SMSChargingInformation) SetUserEquipmentInfo(v string) {
	o.UserEquipmentInfo = &v
}

// GetRoamerInOut returns the RoamerInOut field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetRoamerInOut() RoamerInOut {
	if o == nil || isNil(o.RoamerInOut) {
		var ret RoamerInOut
		return ret
	}
	return *o.RoamerInOut
}

// GetRoamerInOutOk returns a tuple with the RoamerInOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetRoamerInOutOk() (*RoamerInOut, bool) {
	if o == nil || isNil(o.RoamerInOut) {
    return nil, false
	}
	return o.RoamerInOut, true
}

// HasRoamerInOut returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasRoamerInOut() bool {
	if o != nil && !isNil(o.RoamerInOut) {
		return true
	}

	return false
}

// SetRoamerInOut gets a reference to the given RoamerInOut and assigns it to the RoamerInOut field.
func (o *SMSChargingInformation) SetRoamerInOut(v RoamerInOut) {
	o.RoamerInOut = &v
}

// GetUserLocationinfo returns the UserLocationinfo field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetUserLocationinfo() UserLocation {
	if o == nil || isNil(o.UserLocationinfo) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationinfo
}

// GetUserLocationinfoOk returns a tuple with the UserLocationinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UserLocationinfo) {
    return nil, false
	}
	return o.UserLocationinfo, true
}

// HasUserLocationinfo returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasUserLocationinfo() bool {
	if o != nil && !isNil(o.UserLocationinfo) {
		return true
	}

	return false
}

// SetUserLocationinfo gets a reference to the given UserLocation and assigns it to the UserLocationinfo field.
func (o *SMSChargingInformation) SetUserLocationinfo(v UserLocation) {
	o.UserLocationinfo = &v
}

// GetUetimeZone returns the UetimeZone field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetUetimeZone() string {
	if o == nil || isNil(o.UetimeZone) {
		var ret string
		return ret
	}
	return *o.UetimeZone
}

// GetUetimeZoneOk returns a tuple with the UetimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetUetimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.UetimeZone) {
    return nil, false
	}
	return o.UetimeZone, true
}

// HasUetimeZone returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasUetimeZone() bool {
	if o != nil && !isNil(o.UetimeZone) {
		return true
	}

	return false
}

// SetUetimeZone gets a reference to the given string and assigns it to the UetimeZone field.
func (o *SMSChargingInformation) SetUetimeZone(v string) {
	o.UetimeZone = &v
}

// GetRATType returns the RATType field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetRATType() RatType {
	if o == nil || isNil(o.RATType) {
		var ret RatType
		return ret
	}
	return *o.RATType
}

// GetRATTypeOk returns a tuple with the RATType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetRATTypeOk() (*RatType, bool) {
	if o == nil || isNil(o.RATType) {
    return nil, false
	}
	return o.RATType, true
}

// HasRATType returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasRATType() bool {
	if o != nil && !isNil(o.RATType) {
		return true
	}

	return false
}

// SetRATType gets a reference to the given RatType and assigns it to the RATType field.
func (o *SMSChargingInformation) SetRATType(v RatType) {
	o.RATType = &v
}

// GetSMSCAddress returns the SMSCAddress field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMSCAddress() string {
	if o == nil || isNil(o.SMSCAddress) {
		var ret string
		return ret
	}
	return *o.SMSCAddress
}

// GetSMSCAddressOk returns a tuple with the SMSCAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMSCAddressOk() (*string, bool) {
	if o == nil || isNil(o.SMSCAddress) {
    return nil, false
	}
	return o.SMSCAddress, true
}

// HasSMSCAddress returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMSCAddress() bool {
	if o != nil && !isNil(o.SMSCAddress) {
		return true
	}

	return false
}

// SetSMSCAddress gets a reference to the given string and assigns it to the SMSCAddress field.
func (o *SMSChargingInformation) SetSMSCAddress(v string) {
	o.SMSCAddress = &v
}

// GetSMDataCodingScheme returns the SMDataCodingScheme field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMDataCodingScheme() int32 {
	if o == nil || isNil(o.SMDataCodingScheme) {
		var ret int32
		return ret
	}
	return *o.SMDataCodingScheme
}

// GetSMDataCodingSchemeOk returns a tuple with the SMDataCodingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMDataCodingSchemeOk() (*int32, bool) {
	if o == nil || isNil(o.SMDataCodingScheme) {
    return nil, false
	}
	return o.SMDataCodingScheme, true
}

// HasSMDataCodingScheme returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMDataCodingScheme() bool {
	if o != nil && !isNil(o.SMDataCodingScheme) {
		return true
	}

	return false
}

// SetSMDataCodingScheme gets a reference to the given int32 and assigns it to the SMDataCodingScheme field.
func (o *SMSChargingInformation) SetSMDataCodingScheme(v int32) {
	o.SMDataCodingScheme = &v
}

// GetSMMessageType returns the SMMessageType field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMMessageType() SMMessageType {
	if o == nil || isNil(o.SMMessageType) {
		var ret SMMessageType
		return ret
	}
	return *o.SMMessageType
}

// GetSMMessageTypeOk returns a tuple with the SMMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMMessageTypeOk() (*SMMessageType, bool) {
	if o == nil || isNil(o.SMMessageType) {
    return nil, false
	}
	return o.SMMessageType, true
}

// HasSMMessageType returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMMessageType() bool {
	if o != nil && !isNil(o.SMMessageType) {
		return true
	}

	return false
}

// SetSMMessageType gets a reference to the given SMMessageType and assigns it to the SMMessageType field.
func (o *SMSChargingInformation) SetSMMessageType(v SMMessageType) {
	o.SMMessageType = &v
}

// GetSMReplyPathRequested returns the SMReplyPathRequested field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMReplyPathRequested() ReplyPathRequested {
	if o == nil || isNil(o.SMReplyPathRequested) {
		var ret ReplyPathRequested
		return ret
	}
	return *o.SMReplyPathRequested
}

// GetSMReplyPathRequestedOk returns a tuple with the SMReplyPathRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMReplyPathRequestedOk() (*ReplyPathRequested, bool) {
	if o == nil || isNil(o.SMReplyPathRequested) {
    return nil, false
	}
	return o.SMReplyPathRequested, true
}

// HasSMReplyPathRequested returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMReplyPathRequested() bool {
	if o != nil && !isNil(o.SMReplyPathRequested) {
		return true
	}

	return false
}

// SetSMReplyPathRequested gets a reference to the given ReplyPathRequested and assigns it to the SMReplyPathRequested field.
func (o *SMSChargingInformation) SetSMReplyPathRequested(v ReplyPathRequested) {
	o.SMReplyPathRequested = &v
}

// GetSMUserDataHeader returns the SMUserDataHeader field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMUserDataHeader() string {
	if o == nil || isNil(o.SMUserDataHeader) {
		var ret string
		return ret
	}
	return *o.SMUserDataHeader
}

// GetSMUserDataHeaderOk returns a tuple with the SMUserDataHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMUserDataHeaderOk() (*string, bool) {
	if o == nil || isNil(o.SMUserDataHeader) {
    return nil, false
	}
	return o.SMUserDataHeader, true
}

// HasSMUserDataHeader returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMUserDataHeader() bool {
	if o != nil && !isNil(o.SMUserDataHeader) {
		return true
	}

	return false
}

// SetSMUserDataHeader gets a reference to the given string and assigns it to the SMUserDataHeader field.
func (o *SMSChargingInformation) SetSMUserDataHeader(v string) {
	o.SMUserDataHeader = &v
}

// GetSMStatus returns the SMStatus field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMStatus() string {
	if o == nil || isNil(o.SMStatus) {
		var ret string
		return ret
	}
	return *o.SMStatus
}

// GetSMStatusOk returns a tuple with the SMStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMStatusOk() (*string, bool) {
	if o == nil || isNil(o.SMStatus) {
    return nil, false
	}
	return o.SMStatus, true
}

// HasSMStatus returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMStatus() bool {
	if o != nil && !isNil(o.SMStatus) {
		return true
	}

	return false
}

// SetSMStatus gets a reference to the given string and assigns it to the SMStatus field.
func (o *SMSChargingInformation) SetSMStatus(v string) {
	o.SMStatus = &v
}

// GetSMDischargeTime returns the SMDischargeTime field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMDischargeTime() time.Time {
	if o == nil || isNil(o.SMDischargeTime) {
		var ret time.Time
		return ret
	}
	return *o.SMDischargeTime
}

// GetSMDischargeTimeOk returns a tuple with the SMDischargeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMDischargeTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.SMDischargeTime) {
    return nil, false
	}
	return o.SMDischargeTime, true
}

// HasSMDischargeTime returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMDischargeTime() bool {
	if o != nil && !isNil(o.SMDischargeTime) {
		return true
	}

	return false
}

// SetSMDischargeTime gets a reference to the given time.Time and assigns it to the SMDischargeTime field.
func (o *SMSChargingInformation) SetSMDischargeTime(v time.Time) {
	o.SMDischargeTime = &v
}

// GetNumberofMessagesSent returns the NumberofMessagesSent field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetNumberofMessagesSent() int32 {
	if o == nil || isNil(o.NumberofMessagesSent) {
		var ret int32
		return ret
	}
	return *o.NumberofMessagesSent
}

// GetNumberofMessagesSentOk returns a tuple with the NumberofMessagesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetNumberofMessagesSentOk() (*int32, bool) {
	if o == nil || isNil(o.NumberofMessagesSent) {
    return nil, false
	}
	return o.NumberofMessagesSent, true
}

// HasNumberofMessagesSent returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasNumberofMessagesSent() bool {
	if o != nil && !isNil(o.NumberofMessagesSent) {
		return true
	}

	return false
}

// SetNumberofMessagesSent gets a reference to the given int32 and assigns it to the NumberofMessagesSent field.
func (o *SMSChargingInformation) SetNumberofMessagesSent(v int32) {
	o.NumberofMessagesSent = &v
}

// GetSMServiceType returns the SMServiceType field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMServiceType() SMServiceType {
	if o == nil || isNil(o.SMServiceType) {
		var ret SMServiceType
		return ret
	}
	return *o.SMServiceType
}

// GetSMServiceTypeOk returns a tuple with the SMServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMServiceTypeOk() (*SMServiceType, bool) {
	if o == nil || isNil(o.SMServiceType) {
    return nil, false
	}
	return o.SMServiceType, true
}

// HasSMServiceType returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMServiceType() bool {
	if o != nil && !isNil(o.SMServiceType) {
		return true
	}

	return false
}

// SetSMServiceType gets a reference to the given SMServiceType and assigns it to the SMServiceType field.
func (o *SMSChargingInformation) SetSMServiceType(v SMServiceType) {
	o.SMServiceType = &v
}

// GetSMSequenceNumber returns the SMSequenceNumber field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMSequenceNumber() int32 {
	if o == nil || isNil(o.SMSequenceNumber) {
		var ret int32
		return ret
	}
	return *o.SMSequenceNumber
}

// GetSMSequenceNumberOk returns a tuple with the SMSequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMSequenceNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SMSequenceNumber) {
    return nil, false
	}
	return o.SMSequenceNumber, true
}

// HasSMSequenceNumber returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMSequenceNumber() bool {
	if o != nil && !isNil(o.SMSequenceNumber) {
		return true
	}

	return false
}

// SetSMSequenceNumber gets a reference to the given int32 and assigns it to the SMSequenceNumber field.
func (o *SMSChargingInformation) SetSMSequenceNumber(v int32) {
	o.SMSequenceNumber = &v
}

// GetSMSresult returns the SMSresult field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMSresult() int32 {
	if o == nil || isNil(o.SMSresult) {
		var ret int32
		return ret
	}
	return *o.SMSresult
}

// GetSMSresultOk returns a tuple with the SMSresult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMSresultOk() (*int32, bool) {
	if o == nil || isNil(o.SMSresult) {
    return nil, false
	}
	return o.SMSresult, true
}

// HasSMSresult returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMSresult() bool {
	if o != nil && !isNil(o.SMSresult) {
		return true
	}

	return false
}

// SetSMSresult gets a reference to the given int32 and assigns it to the SMSresult field.
func (o *SMSChargingInformation) SetSMSresult(v int32) {
	o.SMSresult = &v
}

// GetSubmissionTime returns the SubmissionTime field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSubmissionTime() time.Time {
	if o == nil || isNil(o.SubmissionTime) {
		var ret time.Time
		return ret
	}
	return *o.SubmissionTime
}

// GetSubmissionTimeOk returns a tuple with the SubmissionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSubmissionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.SubmissionTime) {
    return nil, false
	}
	return o.SubmissionTime, true
}

// HasSubmissionTime returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSubmissionTime() bool {
	if o != nil && !isNil(o.SubmissionTime) {
		return true
	}

	return false
}

// SetSubmissionTime gets a reference to the given time.Time and assigns it to the SubmissionTime field.
func (o *SMSChargingInformation) SetSubmissionTime(v time.Time) {
	o.SubmissionTime = &v
}

// GetSMPriority returns the SMPriority field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetSMPriority() SMPriority {
	if o == nil || isNil(o.SMPriority) {
		var ret SMPriority
		return ret
	}
	return *o.SMPriority
}

// GetSMPriorityOk returns a tuple with the SMPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetSMPriorityOk() (*SMPriority, bool) {
	if o == nil || isNil(o.SMPriority) {
    return nil, false
	}
	return o.SMPriority, true
}

// HasSMPriority returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasSMPriority() bool {
	if o != nil && !isNil(o.SMPriority) {
		return true
	}

	return false
}

// SetSMPriority gets a reference to the given SMPriority and assigns it to the SMPriority field.
func (o *SMSChargingInformation) SetSMPriority(v SMPriority) {
	o.SMPriority = &v
}

// GetMessageReference returns the MessageReference field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetMessageReference() string {
	if o == nil || isNil(o.MessageReference) {
		var ret string
		return ret
	}
	return *o.MessageReference
}

// GetMessageReferenceOk returns a tuple with the MessageReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetMessageReferenceOk() (*string, bool) {
	if o == nil || isNil(o.MessageReference) {
    return nil, false
	}
	return o.MessageReference, true
}

// HasMessageReference returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasMessageReference() bool {
	if o != nil && !isNil(o.MessageReference) {
		return true
	}

	return false
}

// SetMessageReference gets a reference to the given string and assigns it to the MessageReference field.
func (o *SMSChargingInformation) SetMessageReference(v string) {
	o.MessageReference = &v
}

// GetMessageSize returns the MessageSize field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetMessageSize() int32 {
	if o == nil || isNil(o.MessageSize) {
		var ret int32
		return ret
	}
	return *o.MessageSize
}

// GetMessageSizeOk returns a tuple with the MessageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetMessageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.MessageSize) {
    return nil, false
	}
	return o.MessageSize, true
}

// HasMessageSize returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasMessageSize() bool {
	if o != nil && !isNil(o.MessageSize) {
		return true
	}

	return false
}

// SetMessageSize gets a reference to the given int32 and assigns it to the MessageSize field.
func (o *SMSChargingInformation) SetMessageSize(v int32) {
	o.MessageSize = &v
}

// GetMessageClass returns the MessageClass field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetMessageClass() MessageClass {
	if o == nil || isNil(o.MessageClass) {
		var ret MessageClass
		return ret
	}
	return *o.MessageClass
}

// GetMessageClassOk returns a tuple with the MessageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetMessageClassOk() (*MessageClass, bool) {
	if o == nil || isNil(o.MessageClass) {
    return nil, false
	}
	return o.MessageClass, true
}

// HasMessageClass returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasMessageClass() bool {
	if o != nil && !isNil(o.MessageClass) {
		return true
	}

	return false
}

// SetMessageClass gets a reference to the given MessageClass and assigns it to the MessageClass field.
func (o *SMSChargingInformation) SetMessageClass(v MessageClass) {
	o.MessageClass = &v
}

// GetDeliveryReportRequested returns the DeliveryReportRequested field value if set, zero value otherwise.
func (o *SMSChargingInformation) GetDeliveryReportRequested() DeliveryReportRequested {
	if o == nil || isNil(o.DeliveryReportRequested) {
		var ret DeliveryReportRequested
		return ret
	}
	return *o.DeliveryReportRequested
}

// GetDeliveryReportRequestedOk returns a tuple with the DeliveryReportRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSChargingInformation) GetDeliveryReportRequestedOk() (*DeliveryReportRequested, bool) {
	if o == nil || isNil(o.DeliveryReportRequested) {
    return nil, false
	}
	return o.DeliveryReportRequested, true
}

// HasDeliveryReportRequested returns a boolean if a field has been set.
func (o *SMSChargingInformation) HasDeliveryReportRequested() bool {
	if o != nil && !isNil(o.DeliveryReportRequested) {
		return true
	}

	return false
}

// SetDeliveryReportRequested gets a reference to the given DeliveryReportRequested and assigns it to the DeliveryReportRequested field.
func (o *SMSChargingInformation) SetDeliveryReportRequested(v DeliveryReportRequested) {
	o.DeliveryReportRequested = &v
}

func (o SMSChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OriginatorInfo) {
		toSerialize["originatorInfo"] = o.OriginatorInfo
	}
	if !isNil(o.RecipientInfo) {
		toSerialize["recipientInfo"] = o.RecipientInfo
	}
	if !isNil(o.UserEquipmentInfo) {
		toSerialize["userEquipmentInfo"] = o.UserEquipmentInfo
	}
	if !isNil(o.RoamerInOut) {
		toSerialize["roamerInOut"] = o.RoamerInOut
	}
	if !isNil(o.UserLocationinfo) {
		toSerialize["userLocationinfo"] = o.UserLocationinfo
	}
	if !isNil(o.UetimeZone) {
		toSerialize["uetimeZone"] = o.UetimeZone
	}
	if !isNil(o.RATType) {
		toSerialize["rATType"] = o.RATType
	}
	if !isNil(o.SMSCAddress) {
		toSerialize["sMSCAddress"] = o.SMSCAddress
	}
	if !isNil(o.SMDataCodingScheme) {
		toSerialize["sMDataCodingScheme"] = o.SMDataCodingScheme
	}
	if !isNil(o.SMMessageType) {
		toSerialize["sMMessageType"] = o.SMMessageType
	}
	if !isNil(o.SMReplyPathRequested) {
		toSerialize["sMReplyPathRequested"] = o.SMReplyPathRequested
	}
	if !isNil(o.SMUserDataHeader) {
		toSerialize["sMUserDataHeader"] = o.SMUserDataHeader
	}
	if !isNil(o.SMStatus) {
		toSerialize["sMStatus"] = o.SMStatus
	}
	if !isNil(o.SMDischargeTime) {
		toSerialize["sMDischargeTime"] = o.SMDischargeTime
	}
	if !isNil(o.NumberofMessagesSent) {
		toSerialize["numberofMessagesSent"] = o.NumberofMessagesSent
	}
	if !isNil(o.SMServiceType) {
		toSerialize["sMServiceType"] = o.SMServiceType
	}
	if !isNil(o.SMSequenceNumber) {
		toSerialize["sMSequenceNumber"] = o.SMSequenceNumber
	}
	if !isNil(o.SMSresult) {
		toSerialize["sMSresult"] = o.SMSresult
	}
	if !isNil(o.SubmissionTime) {
		toSerialize["submissionTime"] = o.SubmissionTime
	}
	if !isNil(o.SMPriority) {
		toSerialize["sMPriority"] = o.SMPriority
	}
	if !isNil(o.MessageReference) {
		toSerialize["messageReference"] = o.MessageReference
	}
	if !isNil(o.MessageSize) {
		toSerialize["messageSize"] = o.MessageSize
	}
	if !isNil(o.MessageClass) {
		toSerialize["messageClass"] = o.MessageClass
	}
	if !isNil(o.DeliveryReportRequested) {
		toSerialize["deliveryReportRequested"] = o.DeliveryReportRequested
	}
	return json.Marshal(toSerialize)
}

type NullableSMSChargingInformation struct {
	value *SMSChargingInformation
	isSet bool
}

func (v NullableSMSChargingInformation) Get() *SMSChargingInformation {
	return v.value
}

func (v *NullableSMSChargingInformation) Set(val *SMSChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSChargingInformation(val *SMSChargingInformation) *NullableSMSChargingInformation {
	return &NullableSMSChargingInformation{value: val, isSet: true}
}

func (v NullableSMSChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

