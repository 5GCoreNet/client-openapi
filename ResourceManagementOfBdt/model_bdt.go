/*
3gpp-bdt

API for BDT resouce management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ResourceManagementOfBdt

import (
	"encoding/json"
)

// Bdt Represents a Background Data Transfer subscription.
type Bdt struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	VolumePerUE UsageThreshold `json:"volumePerUE"`
	// Identifies the number of UEs.
	NumberOfUEs int32 `json:"numberOfUEs"`
	DesiredTimeWindow TimeWindow `json:"desiredTimeWindow"`
	LocationArea *LocationArea `json:"locationArea,omitempty"`
	LocationArea5G *LocationArea5G `json:"locationArea5G,omitempty"`
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	ReferenceId *string `json:"referenceId,omitempty"`
	// Identifies an offered transfer policy.
	TransferPolicies []TransferPolicy `json:"transferPolicies,omitempty"`
	// Identity of the selected background data transfer policy. Shall not be present in initial message exchange, can be provided by NF service consumer in a subsequent message exchange.
	SelectedPolicy *int32 `json:"selectedPolicy,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId *string `json:"externalGroupId,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// Indicates whether the BDT warning notification is enabled (true) or not (false). Default value is false. 
	WarnNotifEnabled *bool `json:"warnNotifEnabled,omitempty"`
	// Identify a traffic descriptor as defined in Figure 5.2.2 of 3GPP TS 24.526, octets v+5 to w.
	TrafficDes *string `json:"trafficDes,omitempty"`
}

// NewBdt instantiates a new Bdt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBdt(volumePerUE UsageThreshold, numberOfUEs int32, desiredTimeWindow TimeWindow) *Bdt {
	this := Bdt{}
	this.VolumePerUE = volumePerUE
	this.NumberOfUEs = numberOfUEs
	this.DesiredTimeWindow = desiredTimeWindow
	return &this
}

// NewBdtWithDefaults instantiates a new Bdt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBdtWithDefaults() *Bdt {
	this := Bdt{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Bdt) GetSelf() string {
	if o == nil || isNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetSelfOk() (*string, bool) {
	if o == nil || isNil(o.Self) {
    return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Bdt) HasSelf() bool {
	if o != nil && !isNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *Bdt) SetSelf(v string) {
	o.Self = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *Bdt) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *Bdt) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *Bdt) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetVolumePerUE returns the VolumePerUE field value
func (o *Bdt) GetVolumePerUE() UsageThreshold {
	if o == nil {
		var ret UsageThreshold
		return ret
	}

	return o.VolumePerUE
}

// GetVolumePerUEOk returns a tuple with the VolumePerUE field value
// and a boolean to check if the value has been set.
func (o *Bdt) GetVolumePerUEOk() (*UsageThreshold, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VolumePerUE, true
}

// SetVolumePerUE sets field value
func (o *Bdt) SetVolumePerUE(v UsageThreshold) {
	o.VolumePerUE = v
}

// GetNumberOfUEs returns the NumberOfUEs field value
func (o *Bdt) GetNumberOfUEs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfUEs
}

// GetNumberOfUEsOk returns a tuple with the NumberOfUEs field value
// and a boolean to check if the value has been set.
func (o *Bdt) GetNumberOfUEsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NumberOfUEs, true
}

// SetNumberOfUEs sets field value
func (o *Bdt) SetNumberOfUEs(v int32) {
	o.NumberOfUEs = v
}

// GetDesiredTimeWindow returns the DesiredTimeWindow field value
func (o *Bdt) GetDesiredTimeWindow() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.DesiredTimeWindow
}

// GetDesiredTimeWindowOk returns a tuple with the DesiredTimeWindow field value
// and a boolean to check if the value has been set.
func (o *Bdt) GetDesiredTimeWindowOk() (*TimeWindow, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DesiredTimeWindow, true
}

// SetDesiredTimeWindow sets field value
func (o *Bdt) SetDesiredTimeWindow(v TimeWindow) {
	o.DesiredTimeWindow = v
}

// GetLocationArea returns the LocationArea field value if set, zero value otherwise.
func (o *Bdt) GetLocationArea() LocationArea {
	if o == nil || isNil(o.LocationArea) {
		var ret LocationArea
		return ret
	}
	return *o.LocationArea
}

// GetLocationAreaOk returns a tuple with the LocationArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetLocationAreaOk() (*LocationArea, bool) {
	if o == nil || isNil(o.LocationArea) {
    return nil, false
	}
	return o.LocationArea, true
}

// HasLocationArea returns a boolean if a field has been set.
func (o *Bdt) HasLocationArea() bool {
	if o != nil && !isNil(o.LocationArea) {
		return true
	}

	return false
}

// SetLocationArea gets a reference to the given LocationArea and assigns it to the LocationArea field.
func (o *Bdt) SetLocationArea(v LocationArea) {
	o.LocationArea = &v
}

// GetLocationArea5G returns the LocationArea5G field value if set, zero value otherwise.
func (o *Bdt) GetLocationArea5G() LocationArea5G {
	if o == nil || isNil(o.LocationArea5G) {
		var ret LocationArea5G
		return ret
	}
	return *o.LocationArea5G
}

// GetLocationArea5GOk returns a tuple with the LocationArea5G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetLocationArea5GOk() (*LocationArea5G, bool) {
	if o == nil || isNil(o.LocationArea5G) {
    return nil, false
	}
	return o.LocationArea5G, true
}

// HasLocationArea5G returns a boolean if a field has been set.
func (o *Bdt) HasLocationArea5G() bool {
	if o != nil && !isNil(o.LocationArea5G) {
		return true
	}

	return false
}

// SetLocationArea5G gets a reference to the given LocationArea5G and assigns it to the LocationArea5G field.
func (o *Bdt) SetLocationArea5G(v LocationArea5G) {
	o.LocationArea5G = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *Bdt) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *Bdt) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *Bdt) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetTransferPolicies returns the TransferPolicies field value if set, zero value otherwise.
func (o *Bdt) GetTransferPolicies() []TransferPolicy {
	if o == nil || isNil(o.TransferPolicies) {
		var ret []TransferPolicy
		return ret
	}
	return o.TransferPolicies
}

// GetTransferPoliciesOk returns a tuple with the TransferPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetTransferPoliciesOk() ([]TransferPolicy, bool) {
	if o == nil || isNil(o.TransferPolicies) {
    return nil, false
	}
	return o.TransferPolicies, true
}

// HasTransferPolicies returns a boolean if a field has been set.
func (o *Bdt) HasTransferPolicies() bool {
	if o != nil && !isNil(o.TransferPolicies) {
		return true
	}

	return false
}

// SetTransferPolicies gets a reference to the given []TransferPolicy and assigns it to the TransferPolicies field.
func (o *Bdt) SetTransferPolicies(v []TransferPolicy) {
	o.TransferPolicies = v
}

// GetSelectedPolicy returns the SelectedPolicy field value if set, zero value otherwise.
func (o *Bdt) GetSelectedPolicy() int32 {
	if o == nil || isNil(o.SelectedPolicy) {
		var ret int32
		return ret
	}
	return *o.SelectedPolicy
}

// GetSelectedPolicyOk returns a tuple with the SelectedPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetSelectedPolicyOk() (*int32, bool) {
	if o == nil || isNil(o.SelectedPolicy) {
    return nil, false
	}
	return o.SelectedPolicy, true
}

// HasSelectedPolicy returns a boolean if a field has been set.
func (o *Bdt) HasSelectedPolicy() bool {
	if o != nil && !isNil(o.SelectedPolicy) {
		return true
	}

	return false
}

// SetSelectedPolicy gets a reference to the given int32 and assigns it to the SelectedPolicy field.
func (o *Bdt) SetSelectedPolicy(v int32) {
	o.SelectedPolicy = &v
}

// GetExternalGroupId returns the ExternalGroupId field value if set, zero value otherwise.
func (o *Bdt) GetExternalGroupId() string {
	if o == nil || isNil(o.ExternalGroupId) {
		var ret string
		return ret
	}
	return *o.ExternalGroupId
}

// GetExternalGroupIdOk returns a tuple with the ExternalGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetExternalGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalGroupId) {
    return nil, false
	}
	return o.ExternalGroupId, true
}

// HasExternalGroupId returns a boolean if a field has been set.
func (o *Bdt) HasExternalGroupId() bool {
	if o != nil && !isNil(o.ExternalGroupId) {
		return true
	}

	return false
}

// SetExternalGroupId gets a reference to the given string and assigns it to the ExternalGroupId field.
func (o *Bdt) SetExternalGroupId(v string) {
	o.ExternalGroupId = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *Bdt) GetNotificationDestination() string {
	if o == nil || isNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || isNil(o.NotificationDestination) {
    return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *Bdt) HasNotificationDestination() bool {
	if o != nil && !isNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *Bdt) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetWarnNotifEnabled returns the WarnNotifEnabled field value if set, zero value otherwise.
func (o *Bdt) GetWarnNotifEnabled() bool {
	if o == nil || isNil(o.WarnNotifEnabled) {
		var ret bool
		return ret
	}
	return *o.WarnNotifEnabled
}

// GetWarnNotifEnabledOk returns a tuple with the WarnNotifEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetWarnNotifEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.WarnNotifEnabled) {
    return nil, false
	}
	return o.WarnNotifEnabled, true
}

// HasWarnNotifEnabled returns a boolean if a field has been set.
func (o *Bdt) HasWarnNotifEnabled() bool {
	if o != nil && !isNil(o.WarnNotifEnabled) {
		return true
	}

	return false
}

// SetWarnNotifEnabled gets a reference to the given bool and assigns it to the WarnNotifEnabled field.
func (o *Bdt) SetWarnNotifEnabled(v bool) {
	o.WarnNotifEnabled = &v
}

// GetTrafficDes returns the TrafficDes field value if set, zero value otherwise.
func (o *Bdt) GetTrafficDes() string {
	if o == nil || isNil(o.TrafficDes) {
		var ret string
		return ret
	}
	return *o.TrafficDes
}

// GetTrafficDesOk returns a tuple with the TrafficDes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bdt) GetTrafficDesOk() (*string, bool) {
	if o == nil || isNil(o.TrafficDes) {
    return nil, false
	}
	return o.TrafficDes, true
}

// HasTrafficDes returns a boolean if a field has been set.
func (o *Bdt) HasTrafficDes() bool {
	if o != nil && !isNil(o.TrafficDes) {
		return true
	}

	return false
}

// SetTrafficDes gets a reference to the given string and assigns it to the TrafficDes field.
func (o *Bdt) SetTrafficDes(v string) {
	o.TrafficDes = &v
}

func (o Bdt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if true {
		toSerialize["volumePerUE"] = o.VolumePerUE
	}
	if true {
		toSerialize["numberOfUEs"] = o.NumberOfUEs
	}
	if true {
		toSerialize["desiredTimeWindow"] = o.DesiredTimeWindow
	}
	if !isNil(o.LocationArea) {
		toSerialize["locationArea"] = o.LocationArea
	}
	if !isNil(o.LocationArea5G) {
		toSerialize["locationArea5G"] = o.LocationArea5G
	}
	if !isNil(o.ReferenceId) {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if !isNil(o.TransferPolicies) {
		toSerialize["transferPolicies"] = o.TransferPolicies
	}
	if !isNil(o.SelectedPolicy) {
		toSerialize["selectedPolicy"] = o.SelectedPolicy
	}
	if !isNil(o.ExternalGroupId) {
		toSerialize["externalGroupId"] = o.ExternalGroupId
	}
	if !isNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !isNil(o.WarnNotifEnabled) {
		toSerialize["warnNotifEnabled"] = o.WarnNotifEnabled
	}
	if !isNil(o.TrafficDes) {
		toSerialize["trafficDes"] = o.TrafficDes
	}
	return json.Marshal(toSerialize)
}

type NullableBdt struct {
	value *Bdt
	isSet bool
}

func (v NullableBdt) Get() *Bdt {
	return v.value
}

func (v *NullableBdt) Set(val *Bdt) {
	v.value = val
	v.isSet = true
}

func (v NullableBdt) IsSet() bool {
	return v.isSet
}

func (v *NullableBdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBdt(val *Bdt) *NullableBdt {
	return &NullableBdt{value: val, isSet: true}
}

func (v NullableBdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


