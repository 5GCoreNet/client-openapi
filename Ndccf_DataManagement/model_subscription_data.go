/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"time"
)

// SubscriptionData Information of a subscription to notifications to NRF events, included in subscription requests and responses 
type SubscriptionData struct {
	NfStatusNotificationUri string `json:"nfStatusNotificationUri"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	ReqNfInstanceId *string `json:"reqNfInstanceId,omitempty"`
	SubscrCond *SubscrCond `json:"subscrCond,omitempty"`
	SubscriptionId string `json:"subscriptionId"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	ReqNotifEvents []NotificationEventType `json:"reqNotifEvents,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
	NotifCondition *NotifCondition `json:"notifCondition,omitempty"`
	ReqNfType *NFType `json:"reqNfType,omitempty"`
	// Fully Qualified Domain Name
	ReqNfFqdn *string `json:"reqNfFqdn,omitempty"`
	ReqSnssais []ExtSnssai `json:"reqSnssais,omitempty"`
	ReqPerPlmnSnssais []PlmnSnssai `json:"reqPerPlmnSnssais,omitempty"`
	ReqPlmnList []PlmnId `json:"reqPlmnList,omitempty"`
	ReqSnpnList []PlmnIdNid `json:"reqSnpnList,omitempty"`
	ServingScope []string `json:"servingScope,omitempty"`
	RequesterFeatures *string `json:"requesterFeatures,omitempty"`
	NrfSupportedFeatures *string `json:"nrfSupportedFeatures,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HnrfUri *string `json:"hnrfUri,omitempty"`
	OnboardingCapability *bool `json:"onboardingCapability,omitempty"`
	// Fully Qualified Domain Name
	TargetHni *string `json:"targetHni,omitempty"`
	PreferredLocality *string `json:"preferredLocality,omitempty"`
	// A map (list of key-value pairs) where the key of the map represents the relative priority, for the requester, of each locality description among the list of locality descriptions in this query parameter, encoded as \"1\" (highest priority\"), \"2\", \"3\", …,  \"n\" (lowest priority) 
	ExtPreferredLocality *map[string][]LocalityDescription `json:"extPreferredLocality,omitempty"`
}

// NewSubscriptionData instantiates a new SubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionData(nfStatusNotificationUri string, subscriptionId string) *SubscriptionData {
	this := SubscriptionData{}
	this.NfStatusNotificationUri = nfStatusNotificationUri
	this.SubscriptionId = subscriptionId
	var onboardingCapability bool = false
	this.OnboardingCapability = &onboardingCapability
	return &this
}

// NewSubscriptionDataWithDefaults instantiates a new SubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDataWithDefaults() *SubscriptionData {
	this := SubscriptionData{}
	var onboardingCapability bool = false
	this.OnboardingCapability = &onboardingCapability
	return &this
}

// GetNfStatusNotificationUri returns the NfStatusNotificationUri field value
func (o *SubscriptionData) GetNfStatusNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfStatusNotificationUri
}

// GetNfStatusNotificationUriOk returns a tuple with the NfStatusNotificationUri field value
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetNfStatusNotificationUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfStatusNotificationUri, true
}

// SetNfStatusNotificationUri sets field value
func (o *SubscriptionData) SetNfStatusNotificationUri(v string) {
	o.NfStatusNotificationUri = v
}

// GetReqNfInstanceId returns the ReqNfInstanceId field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqNfInstanceId() string {
	if o == nil || isNil(o.ReqNfInstanceId) {
		var ret string
		return ret
	}
	return *o.ReqNfInstanceId
}

// GetReqNfInstanceIdOk returns a tuple with the ReqNfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqNfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReqNfInstanceId) {
    return nil, false
	}
	return o.ReqNfInstanceId, true
}

// HasReqNfInstanceId returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqNfInstanceId() bool {
	if o != nil && !isNil(o.ReqNfInstanceId) {
		return true
	}

	return false
}

// SetReqNfInstanceId gets a reference to the given string and assigns it to the ReqNfInstanceId field.
func (o *SubscriptionData) SetReqNfInstanceId(v string) {
	o.ReqNfInstanceId = &v
}

// GetSubscrCond returns the SubscrCond field value if set, zero value otherwise.
func (o *SubscriptionData) GetSubscrCond() SubscrCond {
	if o == nil || isNil(o.SubscrCond) {
		var ret SubscrCond
		return ret
	}
	return *o.SubscrCond
}

// GetSubscrCondOk returns a tuple with the SubscrCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetSubscrCondOk() (*SubscrCond, bool) {
	if o == nil || isNil(o.SubscrCond) {
    return nil, false
	}
	return o.SubscrCond, true
}

// HasSubscrCond returns a boolean if a field has been set.
func (o *SubscriptionData) HasSubscrCond() bool {
	if o != nil && !isNil(o.SubscrCond) {
		return true
	}

	return false
}

// SetSubscrCond gets a reference to the given SubscrCond and assigns it to the SubscrCond field.
func (o *SubscriptionData) SetSubscrCond(v SubscrCond) {
	o.SubscrCond = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *SubscriptionData) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *SubscriptionData) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *SubscriptionData) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
    return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *SubscriptionData) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *SubscriptionData) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetReqNotifEvents returns the ReqNotifEvents field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqNotifEvents() []NotificationEventType {
	if o == nil || isNil(o.ReqNotifEvents) {
		var ret []NotificationEventType
		return ret
	}
	return o.ReqNotifEvents
}

// GetReqNotifEventsOk returns a tuple with the ReqNotifEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqNotifEventsOk() ([]NotificationEventType, bool) {
	if o == nil || isNil(o.ReqNotifEvents) {
    return nil, false
	}
	return o.ReqNotifEvents, true
}

// HasReqNotifEvents returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqNotifEvents() bool {
	if o != nil && !isNil(o.ReqNotifEvents) {
		return true
	}

	return false
}

// SetReqNotifEvents gets a reference to the given []NotificationEventType and assigns it to the ReqNotifEvents field.
func (o *SubscriptionData) SetReqNotifEvents(v []NotificationEventType) {
	o.ReqNotifEvents = v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *SubscriptionData) GetPlmnId() PlmnId {
	if o == nil || isNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.PlmnId) {
    return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *SubscriptionData) HasPlmnId() bool {
	if o != nil && !isNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *SubscriptionData) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *SubscriptionData) GetNid() string {
	if o == nil || isNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetNidOk() (*string, bool) {
	if o == nil || isNil(o.Nid) {
    return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *SubscriptionData) HasNid() bool {
	if o != nil && !isNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *SubscriptionData) SetNid(v string) {
	o.Nid = &v
}

// GetNotifCondition returns the NotifCondition field value if set, zero value otherwise.
func (o *SubscriptionData) GetNotifCondition() NotifCondition {
	if o == nil || isNil(o.NotifCondition) {
		var ret NotifCondition
		return ret
	}
	return *o.NotifCondition
}

// GetNotifConditionOk returns a tuple with the NotifCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetNotifConditionOk() (*NotifCondition, bool) {
	if o == nil || isNil(o.NotifCondition) {
    return nil, false
	}
	return o.NotifCondition, true
}

// HasNotifCondition returns a boolean if a field has been set.
func (o *SubscriptionData) HasNotifCondition() bool {
	if o != nil && !isNil(o.NotifCondition) {
		return true
	}

	return false
}

// SetNotifCondition gets a reference to the given NotifCondition and assigns it to the NotifCondition field.
func (o *SubscriptionData) SetNotifCondition(v NotifCondition) {
	o.NotifCondition = &v
}

// GetReqNfType returns the ReqNfType field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqNfType() NFType {
	if o == nil || isNil(o.ReqNfType) {
		var ret NFType
		return ret
	}
	return *o.ReqNfType
}

// GetReqNfTypeOk returns a tuple with the ReqNfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqNfTypeOk() (*NFType, bool) {
	if o == nil || isNil(o.ReqNfType) {
    return nil, false
	}
	return o.ReqNfType, true
}

// HasReqNfType returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqNfType() bool {
	if o != nil && !isNil(o.ReqNfType) {
		return true
	}

	return false
}

// SetReqNfType gets a reference to the given NFType and assigns it to the ReqNfType field.
func (o *SubscriptionData) SetReqNfType(v NFType) {
	o.ReqNfType = &v
}

// GetReqNfFqdn returns the ReqNfFqdn field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqNfFqdn() string {
	if o == nil || isNil(o.ReqNfFqdn) {
		var ret string
		return ret
	}
	return *o.ReqNfFqdn
}

// GetReqNfFqdnOk returns a tuple with the ReqNfFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqNfFqdnOk() (*string, bool) {
	if o == nil || isNil(o.ReqNfFqdn) {
    return nil, false
	}
	return o.ReqNfFqdn, true
}

// HasReqNfFqdn returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqNfFqdn() bool {
	if o != nil && !isNil(o.ReqNfFqdn) {
		return true
	}

	return false
}

// SetReqNfFqdn gets a reference to the given string and assigns it to the ReqNfFqdn field.
func (o *SubscriptionData) SetReqNfFqdn(v string) {
	o.ReqNfFqdn = &v
}

// GetReqSnssais returns the ReqSnssais field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqSnssais() []ExtSnssai {
	if o == nil || isNil(o.ReqSnssais) {
		var ret []ExtSnssai
		return ret
	}
	return o.ReqSnssais
}

// GetReqSnssaisOk returns a tuple with the ReqSnssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqSnssaisOk() ([]ExtSnssai, bool) {
	if o == nil || isNil(o.ReqSnssais) {
    return nil, false
	}
	return o.ReqSnssais, true
}

// HasReqSnssais returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqSnssais() bool {
	if o != nil && !isNil(o.ReqSnssais) {
		return true
	}

	return false
}

// SetReqSnssais gets a reference to the given []ExtSnssai and assigns it to the ReqSnssais field.
func (o *SubscriptionData) SetReqSnssais(v []ExtSnssai) {
	o.ReqSnssais = v
}

// GetReqPerPlmnSnssais returns the ReqPerPlmnSnssais field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqPerPlmnSnssais() []PlmnSnssai {
	if o == nil || isNil(o.ReqPerPlmnSnssais) {
		var ret []PlmnSnssai
		return ret
	}
	return o.ReqPerPlmnSnssais
}

// GetReqPerPlmnSnssaisOk returns a tuple with the ReqPerPlmnSnssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqPerPlmnSnssaisOk() ([]PlmnSnssai, bool) {
	if o == nil || isNil(o.ReqPerPlmnSnssais) {
    return nil, false
	}
	return o.ReqPerPlmnSnssais, true
}

// HasReqPerPlmnSnssais returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqPerPlmnSnssais() bool {
	if o != nil && !isNil(o.ReqPerPlmnSnssais) {
		return true
	}

	return false
}

// SetReqPerPlmnSnssais gets a reference to the given []PlmnSnssai and assigns it to the ReqPerPlmnSnssais field.
func (o *SubscriptionData) SetReqPerPlmnSnssais(v []PlmnSnssai) {
	o.ReqPerPlmnSnssais = v
}

// GetReqPlmnList returns the ReqPlmnList field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqPlmnList() []PlmnId {
	if o == nil || isNil(o.ReqPlmnList) {
		var ret []PlmnId
		return ret
	}
	return o.ReqPlmnList
}

// GetReqPlmnListOk returns a tuple with the ReqPlmnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqPlmnListOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.ReqPlmnList) {
    return nil, false
	}
	return o.ReqPlmnList, true
}

// HasReqPlmnList returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqPlmnList() bool {
	if o != nil && !isNil(o.ReqPlmnList) {
		return true
	}

	return false
}

// SetReqPlmnList gets a reference to the given []PlmnId and assigns it to the ReqPlmnList field.
func (o *SubscriptionData) SetReqPlmnList(v []PlmnId) {
	o.ReqPlmnList = v
}

// GetReqSnpnList returns the ReqSnpnList field value if set, zero value otherwise.
func (o *SubscriptionData) GetReqSnpnList() []PlmnIdNid {
	if o == nil || isNil(o.ReqSnpnList) {
		var ret []PlmnIdNid
		return ret
	}
	return o.ReqSnpnList
}

// GetReqSnpnListOk returns a tuple with the ReqSnpnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetReqSnpnListOk() ([]PlmnIdNid, bool) {
	if o == nil || isNil(o.ReqSnpnList) {
    return nil, false
	}
	return o.ReqSnpnList, true
}

// HasReqSnpnList returns a boolean if a field has been set.
func (o *SubscriptionData) HasReqSnpnList() bool {
	if o != nil && !isNil(o.ReqSnpnList) {
		return true
	}

	return false
}

// SetReqSnpnList gets a reference to the given []PlmnIdNid and assigns it to the ReqSnpnList field.
func (o *SubscriptionData) SetReqSnpnList(v []PlmnIdNid) {
	o.ReqSnpnList = v
}

// GetServingScope returns the ServingScope field value if set, zero value otherwise.
func (o *SubscriptionData) GetServingScope() []string {
	if o == nil || isNil(o.ServingScope) {
		var ret []string
		return ret
	}
	return o.ServingScope
}

// GetServingScopeOk returns a tuple with the ServingScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetServingScopeOk() ([]string, bool) {
	if o == nil || isNil(o.ServingScope) {
    return nil, false
	}
	return o.ServingScope, true
}

// HasServingScope returns a boolean if a field has been set.
func (o *SubscriptionData) HasServingScope() bool {
	if o != nil && !isNil(o.ServingScope) {
		return true
	}

	return false
}

// SetServingScope gets a reference to the given []string and assigns it to the ServingScope field.
func (o *SubscriptionData) SetServingScope(v []string) {
	o.ServingScope = v
}

// GetRequesterFeatures returns the RequesterFeatures field value if set, zero value otherwise.
func (o *SubscriptionData) GetRequesterFeatures() string {
	if o == nil || isNil(o.RequesterFeatures) {
		var ret string
		return ret
	}
	return *o.RequesterFeatures
}

// GetRequesterFeaturesOk returns a tuple with the RequesterFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetRequesterFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.RequesterFeatures) {
    return nil, false
	}
	return o.RequesterFeatures, true
}

// HasRequesterFeatures returns a boolean if a field has been set.
func (o *SubscriptionData) HasRequesterFeatures() bool {
	if o != nil && !isNil(o.RequesterFeatures) {
		return true
	}

	return false
}

// SetRequesterFeatures gets a reference to the given string and assigns it to the RequesterFeatures field.
func (o *SubscriptionData) SetRequesterFeatures(v string) {
	o.RequesterFeatures = &v
}

// GetNrfSupportedFeatures returns the NrfSupportedFeatures field value if set, zero value otherwise.
func (o *SubscriptionData) GetNrfSupportedFeatures() string {
	if o == nil || isNil(o.NrfSupportedFeatures) {
		var ret string
		return ret
	}
	return *o.NrfSupportedFeatures
}

// GetNrfSupportedFeaturesOk returns a tuple with the NrfSupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetNrfSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.NrfSupportedFeatures) {
    return nil, false
	}
	return o.NrfSupportedFeatures, true
}

// HasNrfSupportedFeatures returns a boolean if a field has been set.
func (o *SubscriptionData) HasNrfSupportedFeatures() bool {
	if o != nil && !isNil(o.NrfSupportedFeatures) {
		return true
	}

	return false
}

// SetNrfSupportedFeatures gets a reference to the given string and assigns it to the NrfSupportedFeatures field.
func (o *SubscriptionData) SetNrfSupportedFeatures(v string) {
	o.NrfSupportedFeatures = &v
}

// GetHnrfUri returns the HnrfUri field value if set, zero value otherwise.
func (o *SubscriptionData) GetHnrfUri() string {
	if o == nil || isNil(o.HnrfUri) {
		var ret string
		return ret
	}
	return *o.HnrfUri
}

// GetHnrfUriOk returns a tuple with the HnrfUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetHnrfUriOk() (*string, bool) {
	if o == nil || isNil(o.HnrfUri) {
    return nil, false
	}
	return o.HnrfUri, true
}

// HasHnrfUri returns a boolean if a field has been set.
func (o *SubscriptionData) HasHnrfUri() bool {
	if o != nil && !isNil(o.HnrfUri) {
		return true
	}

	return false
}

// SetHnrfUri gets a reference to the given string and assigns it to the HnrfUri field.
func (o *SubscriptionData) SetHnrfUri(v string) {
	o.HnrfUri = &v
}

// GetOnboardingCapability returns the OnboardingCapability field value if set, zero value otherwise.
func (o *SubscriptionData) GetOnboardingCapability() bool {
	if o == nil || isNil(o.OnboardingCapability) {
		var ret bool
		return ret
	}
	return *o.OnboardingCapability
}

// GetOnboardingCapabilityOk returns a tuple with the OnboardingCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetOnboardingCapabilityOk() (*bool, bool) {
	if o == nil || isNil(o.OnboardingCapability) {
    return nil, false
	}
	return o.OnboardingCapability, true
}

// HasOnboardingCapability returns a boolean if a field has been set.
func (o *SubscriptionData) HasOnboardingCapability() bool {
	if o != nil && !isNil(o.OnboardingCapability) {
		return true
	}

	return false
}

// SetOnboardingCapability gets a reference to the given bool and assigns it to the OnboardingCapability field.
func (o *SubscriptionData) SetOnboardingCapability(v bool) {
	o.OnboardingCapability = &v
}

// GetTargetHni returns the TargetHni field value if set, zero value otherwise.
func (o *SubscriptionData) GetTargetHni() string {
	if o == nil || isNil(o.TargetHni) {
		var ret string
		return ret
	}
	return *o.TargetHni
}

// GetTargetHniOk returns a tuple with the TargetHni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetTargetHniOk() (*string, bool) {
	if o == nil || isNil(o.TargetHni) {
    return nil, false
	}
	return o.TargetHni, true
}

// HasTargetHni returns a boolean if a field has been set.
func (o *SubscriptionData) HasTargetHni() bool {
	if o != nil && !isNil(o.TargetHni) {
		return true
	}

	return false
}

// SetTargetHni gets a reference to the given string and assigns it to the TargetHni field.
func (o *SubscriptionData) SetTargetHni(v string) {
	o.TargetHni = &v
}

// GetPreferredLocality returns the PreferredLocality field value if set, zero value otherwise.
func (o *SubscriptionData) GetPreferredLocality() string {
	if o == nil || isNil(o.PreferredLocality) {
		var ret string
		return ret
	}
	return *o.PreferredLocality
}

// GetPreferredLocalityOk returns a tuple with the PreferredLocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetPreferredLocalityOk() (*string, bool) {
	if o == nil || isNil(o.PreferredLocality) {
    return nil, false
	}
	return o.PreferredLocality, true
}

// HasPreferredLocality returns a boolean if a field has been set.
func (o *SubscriptionData) HasPreferredLocality() bool {
	if o != nil && !isNil(o.PreferredLocality) {
		return true
	}

	return false
}

// SetPreferredLocality gets a reference to the given string and assigns it to the PreferredLocality field.
func (o *SubscriptionData) SetPreferredLocality(v string) {
	o.PreferredLocality = &v
}

// GetExtPreferredLocality returns the ExtPreferredLocality field value if set, zero value otherwise.
func (o *SubscriptionData) GetExtPreferredLocality() map[string][]LocalityDescription {
	if o == nil || isNil(o.ExtPreferredLocality) {
		var ret map[string][]LocalityDescription
		return ret
	}
	return *o.ExtPreferredLocality
}

// GetExtPreferredLocalityOk returns a tuple with the ExtPreferredLocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionData) GetExtPreferredLocalityOk() (*map[string][]LocalityDescription, bool) {
	if o == nil || isNil(o.ExtPreferredLocality) {
    return nil, false
	}
	return o.ExtPreferredLocality, true
}

// HasExtPreferredLocality returns a boolean if a field has been set.
func (o *SubscriptionData) HasExtPreferredLocality() bool {
	if o != nil && !isNil(o.ExtPreferredLocality) {
		return true
	}

	return false
}

// SetExtPreferredLocality gets a reference to the given map[string][]LocalityDescription and assigns it to the ExtPreferredLocality field.
func (o *SubscriptionData) SetExtPreferredLocality(v map[string][]LocalityDescription) {
	o.ExtPreferredLocality = &v
}

func (o SubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfStatusNotificationUri"] = o.NfStatusNotificationUri
	}
	if !isNil(o.ReqNfInstanceId) {
		toSerialize["reqNfInstanceId"] = o.ReqNfInstanceId
	}
	if !isNil(o.SubscrCond) {
		toSerialize["subscrCond"] = o.SubscrCond
	}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !isNil(o.ReqNotifEvents) {
		toSerialize["reqNotifEvents"] = o.ReqNotifEvents
	}
	if !isNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	if !isNil(o.NotifCondition) {
		toSerialize["notifCondition"] = o.NotifCondition
	}
	if !isNil(o.ReqNfType) {
		toSerialize["reqNfType"] = o.ReqNfType
	}
	if !isNil(o.ReqNfFqdn) {
		toSerialize["reqNfFqdn"] = o.ReqNfFqdn
	}
	if !isNil(o.ReqSnssais) {
		toSerialize["reqSnssais"] = o.ReqSnssais
	}
	if !isNil(o.ReqPerPlmnSnssais) {
		toSerialize["reqPerPlmnSnssais"] = o.ReqPerPlmnSnssais
	}
	if !isNil(o.ReqPlmnList) {
		toSerialize["reqPlmnList"] = o.ReqPlmnList
	}
	if !isNil(o.ReqSnpnList) {
		toSerialize["reqSnpnList"] = o.ReqSnpnList
	}
	if !isNil(o.ServingScope) {
		toSerialize["servingScope"] = o.ServingScope
	}
	if !isNil(o.RequesterFeatures) {
		toSerialize["requesterFeatures"] = o.RequesterFeatures
	}
	if !isNil(o.NrfSupportedFeatures) {
		toSerialize["nrfSupportedFeatures"] = o.NrfSupportedFeatures
	}
	if !isNil(o.HnrfUri) {
		toSerialize["hnrfUri"] = o.HnrfUri
	}
	if !isNil(o.OnboardingCapability) {
		toSerialize["onboardingCapability"] = o.OnboardingCapability
	}
	if !isNil(o.TargetHni) {
		toSerialize["targetHni"] = o.TargetHni
	}
	if !isNil(o.PreferredLocality) {
		toSerialize["preferredLocality"] = o.PreferredLocality
	}
	if !isNil(o.ExtPreferredLocality) {
		toSerialize["extPreferredLocality"] = o.ExtPreferredLocality
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionData struct {
	value *SubscriptionData
	isSet bool
}

func (v NullableSubscriptionData) Get() *SubscriptionData {
	return v.value
}

func (v *NullableSubscriptionData) Set(val *SubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionData(val *SubscriptionData) *NullableSubscriptionData {
	return &NullableSubscriptionData{value: val, isSet: true}
}

func (v NullableSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


