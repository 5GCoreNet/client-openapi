/*
EES Session with QoS API

API for EES Session with Qos service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_SessionWithQoS

import (
	"encoding/json"
)

// SessionWithQoS Represents an Individual Session with QoS Subscription.
type SessionWithQoS struct {
	// string providing an URI formatted according to IETF RFC 3986.
	Self *string `json:"self,omitempty"`
	// Identifier of an EAS.
	EasId string `json:"easId"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	UeIpv4Addr *string `json:"ueIpv4Addr,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	UeIpv6Addr *string `json:"ueIpv6Addr,omitempty"`
	IpDomain *string `json:"ipDomain,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	UeId *string `json:"ueId,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	IntGrpId *string `json:"intGrpId,omitempty"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExtGrpId *string `json:"extGrpId,omitempty"`
	// Contains the flow description for the Uplink and/or Downlink IP flows.
	IpFlows []string `json:"ipFlows"`
	// Identifies a pre-defined QoS information.
	QosReference *string `json:"qosReference,omitempty"`
	// Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority. 
	AltQosReference []string `json:"altQosReference,omitempty"`
	// Indicates the events subscribed by the EAS.
	Events []UserPlaneEvent `json:"events,omitempty"`
	SponsorInformation *SponsorInformation `json:"sponsorInformation,omitempty"`
	QosMonInfo *QosMonitoringInformation `json:"qosMonInfo,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxbrUl *string `json:"maxbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxbrDl *string `json:"maxbrDl,omitempty"`
	DisUeNotif *bool `json:"disUeNotif,omitempty"`
	// Set to true by Subscriber to request the EES to send a test notification as defined in 3GPP TS 29.122 [6]. Set to false or omitted otherwise. 
	RequestTestNotification *bool `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewSessionWithQoS instantiates a new SessionWithQoS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionWithQoS(easId string, ipFlows []string) *SessionWithQoS {
	this := SessionWithQoS{}
	this.EasId = easId
	this.IpFlows = ipFlows
	return &this
}

// NewSessionWithQoSWithDefaults instantiates a new SessionWithQoS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithQoSWithDefaults() *SessionWithQoS {
	this := SessionWithQoS{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SessionWithQoS) GetSelf() string {
	if o == nil || isNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetSelfOk() (*string, bool) {
	if o == nil || isNil(o.Self) {
    return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SessionWithQoS) HasSelf() bool {
	if o != nil && !isNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *SessionWithQoS) SetSelf(v string) {
	o.Self = &v
}

// GetEasId returns the EasId field value
func (o *SessionWithQoS) GetEasId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetEasIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EasId, true
}

// SetEasId sets field value
func (o *SessionWithQoS) SetEasId(v string) {
	o.EasId = v
}

// GetUeIpv4Addr returns the UeIpv4Addr field value if set, zero value otherwise.
func (o *SessionWithQoS) GetUeIpv4Addr() string {
	if o == nil || isNil(o.UeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.UeIpv4Addr
}

// GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetUeIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.UeIpv4Addr) {
    return nil, false
	}
	return o.UeIpv4Addr, true
}

// HasUeIpv4Addr returns a boolean if a field has been set.
func (o *SessionWithQoS) HasUeIpv4Addr() bool {
	if o != nil && !isNil(o.UeIpv4Addr) {
		return true
	}

	return false
}

// SetUeIpv4Addr gets a reference to the given string and assigns it to the UeIpv4Addr field.
func (o *SessionWithQoS) SetUeIpv4Addr(v string) {
	o.UeIpv4Addr = &v
}

// GetUeIpv6Addr returns the UeIpv6Addr field value if set, zero value otherwise.
func (o *SessionWithQoS) GetUeIpv6Addr() string {
	if o == nil || isNil(o.UeIpv6Addr) {
		var ret string
		return ret
	}
	return *o.UeIpv6Addr
}

// GetUeIpv6AddrOk returns a tuple with the UeIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetUeIpv6AddrOk() (*string, bool) {
	if o == nil || isNil(o.UeIpv6Addr) {
    return nil, false
	}
	return o.UeIpv6Addr, true
}

// HasUeIpv6Addr returns a boolean if a field has been set.
func (o *SessionWithQoS) HasUeIpv6Addr() bool {
	if o != nil && !isNil(o.UeIpv6Addr) {
		return true
	}

	return false
}

// SetUeIpv6Addr gets a reference to the given string and assigns it to the UeIpv6Addr field.
func (o *SessionWithQoS) SetUeIpv6Addr(v string) {
	o.UeIpv6Addr = &v
}

// GetIpDomain returns the IpDomain field value if set, zero value otherwise.
func (o *SessionWithQoS) GetIpDomain() string {
	if o == nil || isNil(o.IpDomain) {
		var ret string
		return ret
	}
	return *o.IpDomain
}

// GetIpDomainOk returns a tuple with the IpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetIpDomainOk() (*string, bool) {
	if o == nil || isNil(o.IpDomain) {
    return nil, false
	}
	return o.IpDomain, true
}

// HasIpDomain returns a boolean if a field has been set.
func (o *SessionWithQoS) HasIpDomain() bool {
	if o != nil && !isNil(o.IpDomain) {
		return true
	}

	return false
}

// SetIpDomain gets a reference to the given string and assigns it to the IpDomain field.
func (o *SessionWithQoS) SetIpDomain(v string) {
	o.IpDomain = &v
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *SessionWithQoS) GetUeId() string {
	if o == nil || isNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetUeIdOk() (*string, bool) {
	if o == nil || isNil(o.UeId) {
    return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *SessionWithQoS) HasUeId() bool {
	if o != nil && !isNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *SessionWithQoS) SetUeId(v string) {
	o.UeId = &v
}

// GetIntGrpId returns the IntGrpId field value if set, zero value otherwise.
func (o *SessionWithQoS) GetIntGrpId() string {
	if o == nil || isNil(o.IntGrpId) {
		var ret string
		return ret
	}
	return *o.IntGrpId
}

// GetIntGrpIdOk returns a tuple with the IntGrpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetIntGrpIdOk() (*string, bool) {
	if o == nil || isNil(o.IntGrpId) {
    return nil, false
	}
	return o.IntGrpId, true
}

// HasIntGrpId returns a boolean if a field has been set.
func (o *SessionWithQoS) HasIntGrpId() bool {
	if o != nil && !isNil(o.IntGrpId) {
		return true
	}

	return false
}

// SetIntGrpId gets a reference to the given string and assigns it to the IntGrpId field.
func (o *SessionWithQoS) SetIntGrpId(v string) {
	o.IntGrpId = &v
}

// GetExtGrpId returns the ExtGrpId field value if set, zero value otherwise.
func (o *SessionWithQoS) GetExtGrpId() string {
	if o == nil || isNil(o.ExtGrpId) {
		var ret string
		return ret
	}
	return *o.ExtGrpId
}

// GetExtGrpIdOk returns a tuple with the ExtGrpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetExtGrpIdOk() (*string, bool) {
	if o == nil || isNil(o.ExtGrpId) {
    return nil, false
	}
	return o.ExtGrpId, true
}

// HasExtGrpId returns a boolean if a field has been set.
func (o *SessionWithQoS) HasExtGrpId() bool {
	if o != nil && !isNil(o.ExtGrpId) {
		return true
	}

	return false
}

// SetExtGrpId gets a reference to the given string and assigns it to the ExtGrpId field.
func (o *SessionWithQoS) SetExtGrpId(v string) {
	o.ExtGrpId = &v
}

// GetIpFlows returns the IpFlows field value
func (o *SessionWithQoS) GetIpFlows() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IpFlows
}

// GetIpFlowsOk returns a tuple with the IpFlows field value
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetIpFlowsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.IpFlows, true
}

// SetIpFlows sets field value
func (o *SessionWithQoS) SetIpFlows(v []string) {
	o.IpFlows = v
}

// GetQosReference returns the QosReference field value if set, zero value otherwise.
func (o *SessionWithQoS) GetQosReference() string {
	if o == nil || isNil(o.QosReference) {
		var ret string
		return ret
	}
	return *o.QosReference
}

// GetQosReferenceOk returns a tuple with the QosReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetQosReferenceOk() (*string, bool) {
	if o == nil || isNil(o.QosReference) {
    return nil, false
	}
	return o.QosReference, true
}

// HasQosReference returns a boolean if a field has been set.
func (o *SessionWithQoS) HasQosReference() bool {
	if o != nil && !isNil(o.QosReference) {
		return true
	}

	return false
}

// SetQosReference gets a reference to the given string and assigns it to the QosReference field.
func (o *SessionWithQoS) SetQosReference(v string) {
	o.QosReference = &v
}

// GetAltQosReference returns the AltQosReference field value if set, zero value otherwise.
func (o *SessionWithQoS) GetAltQosReference() []string {
	if o == nil || isNil(o.AltQosReference) {
		var ret []string
		return ret
	}
	return o.AltQosReference
}

// GetAltQosReferenceOk returns a tuple with the AltQosReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetAltQosReferenceOk() ([]string, bool) {
	if o == nil || isNil(o.AltQosReference) {
    return nil, false
	}
	return o.AltQosReference, true
}

// HasAltQosReference returns a boolean if a field has been set.
func (o *SessionWithQoS) HasAltQosReference() bool {
	if o != nil && !isNil(o.AltQosReference) {
		return true
	}

	return false
}

// SetAltQosReference gets a reference to the given []string and assigns it to the AltQosReference field.
func (o *SessionWithQoS) SetAltQosReference(v []string) {
	o.AltQosReference = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SessionWithQoS) GetEvents() []UserPlaneEvent {
	if o == nil || isNil(o.Events) {
		var ret []UserPlaneEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetEventsOk() ([]UserPlaneEvent, bool) {
	if o == nil || isNil(o.Events) {
    return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SessionWithQoS) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []UserPlaneEvent and assigns it to the Events field.
func (o *SessionWithQoS) SetEvents(v []UserPlaneEvent) {
	o.Events = v
}

// GetSponsorInformation returns the SponsorInformation field value if set, zero value otherwise.
func (o *SessionWithQoS) GetSponsorInformation() SponsorInformation {
	if o == nil || isNil(o.SponsorInformation) {
		var ret SponsorInformation
		return ret
	}
	return *o.SponsorInformation
}

// GetSponsorInformationOk returns a tuple with the SponsorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetSponsorInformationOk() (*SponsorInformation, bool) {
	if o == nil || isNil(o.SponsorInformation) {
    return nil, false
	}
	return o.SponsorInformation, true
}

// HasSponsorInformation returns a boolean if a field has been set.
func (o *SessionWithQoS) HasSponsorInformation() bool {
	if o != nil && !isNil(o.SponsorInformation) {
		return true
	}

	return false
}

// SetSponsorInformation gets a reference to the given SponsorInformation and assigns it to the SponsorInformation field.
func (o *SessionWithQoS) SetSponsorInformation(v SponsorInformation) {
	o.SponsorInformation = &v
}

// GetQosMonInfo returns the QosMonInfo field value if set, zero value otherwise.
func (o *SessionWithQoS) GetQosMonInfo() QosMonitoringInformation {
	if o == nil || isNil(o.QosMonInfo) {
		var ret QosMonitoringInformation
		return ret
	}
	return *o.QosMonInfo
}

// GetQosMonInfoOk returns a tuple with the QosMonInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetQosMonInfoOk() (*QosMonitoringInformation, bool) {
	if o == nil || isNil(o.QosMonInfo) {
    return nil, false
	}
	return o.QosMonInfo, true
}

// HasQosMonInfo returns a boolean if a field has been set.
func (o *SessionWithQoS) HasQosMonInfo() bool {
	if o != nil && !isNil(o.QosMonInfo) {
		return true
	}

	return false
}

// SetQosMonInfo gets a reference to the given QosMonitoringInformation and assigns it to the QosMonInfo field.
func (o *SessionWithQoS) SetQosMonInfo(v QosMonitoringInformation) {
	o.QosMonInfo = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *SessionWithQoS) GetNotificationDestination() string {
	if o == nil || isNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || isNil(o.NotificationDestination) {
    return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *SessionWithQoS) HasNotificationDestination() bool {
	if o != nil && !isNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *SessionWithQoS) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *SessionWithQoS) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *SessionWithQoS) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *SessionWithQoS) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *SessionWithQoS) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *SessionWithQoS) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *SessionWithQoS) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetMaxbrUl returns the MaxbrUl field value if set, zero value otherwise.
func (o *SessionWithQoS) GetMaxbrUl() string {
	if o == nil || isNil(o.MaxbrUl) {
		var ret string
		return ret
	}
	return *o.MaxbrUl
}

// GetMaxbrUlOk returns a tuple with the MaxbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetMaxbrUlOk() (*string, bool) {
	if o == nil || isNil(o.MaxbrUl) {
    return nil, false
	}
	return o.MaxbrUl, true
}

// HasMaxbrUl returns a boolean if a field has been set.
func (o *SessionWithQoS) HasMaxbrUl() bool {
	if o != nil && !isNil(o.MaxbrUl) {
		return true
	}

	return false
}

// SetMaxbrUl gets a reference to the given string and assigns it to the MaxbrUl field.
func (o *SessionWithQoS) SetMaxbrUl(v string) {
	o.MaxbrUl = &v
}

// GetMaxbrDl returns the MaxbrDl field value if set, zero value otherwise.
func (o *SessionWithQoS) GetMaxbrDl() string {
	if o == nil || isNil(o.MaxbrDl) {
		var ret string
		return ret
	}
	return *o.MaxbrDl
}

// GetMaxbrDlOk returns a tuple with the MaxbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetMaxbrDlOk() (*string, bool) {
	if o == nil || isNil(o.MaxbrDl) {
    return nil, false
	}
	return o.MaxbrDl, true
}

// HasMaxbrDl returns a boolean if a field has been set.
func (o *SessionWithQoS) HasMaxbrDl() bool {
	if o != nil && !isNil(o.MaxbrDl) {
		return true
	}

	return false
}

// SetMaxbrDl gets a reference to the given string and assigns it to the MaxbrDl field.
func (o *SessionWithQoS) SetMaxbrDl(v string) {
	o.MaxbrDl = &v
}

// GetDisUeNotif returns the DisUeNotif field value if set, zero value otherwise.
func (o *SessionWithQoS) GetDisUeNotif() bool {
	if o == nil || isNil(o.DisUeNotif) {
		var ret bool
		return ret
	}
	return *o.DisUeNotif
}

// GetDisUeNotifOk returns a tuple with the DisUeNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetDisUeNotifOk() (*bool, bool) {
	if o == nil || isNil(o.DisUeNotif) {
    return nil, false
	}
	return o.DisUeNotif, true
}

// HasDisUeNotif returns a boolean if a field has been set.
func (o *SessionWithQoS) HasDisUeNotif() bool {
	if o != nil && !isNil(o.DisUeNotif) {
		return true
	}

	return false
}

// SetDisUeNotif gets a reference to the given bool and assigns it to the DisUeNotif field.
func (o *SessionWithQoS) SetDisUeNotif(v bool) {
	o.DisUeNotif = &v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *SessionWithQoS) GetRequestTestNotification() bool {
	if o == nil || isNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || isNil(o.RequestTestNotification) {
    return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *SessionWithQoS) HasRequestTestNotification() bool {
	if o != nil && !isNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *SessionWithQoS) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *SessionWithQoS) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || isNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || isNil(o.WebsockNotifConfig) {
    return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *SessionWithQoS) HasWebsockNotifConfig() bool {
	if o != nil && !isNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *SessionWithQoS) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *SessionWithQoS) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoS) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *SessionWithQoS) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *SessionWithQoS) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o SessionWithQoS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if true {
		toSerialize["easId"] = o.EasId
	}
	if !isNil(o.UeIpv4Addr) {
		toSerialize["ueIpv4Addr"] = o.UeIpv4Addr
	}
	if !isNil(o.UeIpv6Addr) {
		toSerialize["ueIpv6Addr"] = o.UeIpv6Addr
	}
	if !isNil(o.IpDomain) {
		toSerialize["ipDomain"] = o.IpDomain
	}
	if !isNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	if !isNil(o.IntGrpId) {
		toSerialize["intGrpId"] = o.IntGrpId
	}
	if !isNil(o.ExtGrpId) {
		toSerialize["extGrpId"] = o.ExtGrpId
	}
	if true {
		toSerialize["ipFlows"] = o.IpFlows
	}
	if !isNil(o.QosReference) {
		toSerialize["qosReference"] = o.QosReference
	}
	if !isNil(o.AltQosReference) {
		toSerialize["altQosReference"] = o.AltQosReference
	}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !isNil(o.SponsorInformation) {
		toSerialize["sponsorInformation"] = o.SponsorInformation
	}
	if !isNil(o.QosMonInfo) {
		toSerialize["qosMonInfo"] = o.QosMonInfo
	}
	if !isNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.MaxbrUl) {
		toSerialize["maxbrUl"] = o.MaxbrUl
	}
	if !isNil(o.MaxbrDl) {
		toSerialize["maxbrDl"] = o.MaxbrDl
	}
	if !isNil(o.DisUeNotif) {
		toSerialize["disUeNotif"] = o.DisUeNotif
	}
	if !isNil(o.RequestTestNotification) {
		toSerialize["requestTestNotification"] = o.RequestTestNotification
	}
	if !isNil(o.WebsockNotifConfig) {
		toSerialize["websockNotifConfig"] = o.WebsockNotifConfig
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableSessionWithQoS struct {
	value *SessionWithQoS
	isSet bool
}

func (v NullableSessionWithQoS) Get() *SessionWithQoS {
	return v.value
}

func (v *NullableSessionWithQoS) Set(val *SessionWithQoS) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionWithQoS) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionWithQoS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionWithQoS(val *SessionWithQoS) *NullableSessionWithQoS {
	return &NullableSessionWithQoS{value: val, isSet: true}
}

func (v NullableSessionWithQoS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionWithQoS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


