/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsUECM

import (
	"encoding/json"
)

// ScscfRegistration Scscf Registration
type ScscfRegistration struct {
	// IMS Private Identity of the UE
	Impi *string `json:"impi,omitempty"`
	ImsRegistrationType ImsRegistrationType `json:"imsRegistrationType"`
	CscfServerName string `json:"cscfServerName"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	ScscfInstanceId *string `json:"scscfInstanceId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DeregCallbackUri *string `json:"deregCallbackUri,omitempty"`
	AssociatedImpis []string `json:"associatedImpis,omitempty"`
	AssociatedRegisteredImpis []string `json:"associatedRegisteredImpis,omitempty"`
	IrsImpus []string `json:"irsImpus,omitempty"`
	// IMS Public Identity of the UE (sip URI or tel URI)
	WildcardedPui *string `json:"wildcardedPui,omitempty"`
	LooseRouteIndicator *LooseRouteIndication `json:"looseRouteIndicator,omitempty"`
	// IMS Public Identity of the UE (sip URI or tel URI)
	WildcardedPsi *string `json:"wildcardedPsi,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	MultipleRegistrationIndicator *bool `json:"multipleRegistrationIndicator,omitempty"`
	PcscfRestorationIndicator *bool `json:"pcscfRestorationIndicator,omitempty"`
	ScscfReselectionIndicator *bool `json:"scscfReselectionIndicator,omitempty"`
}

// NewScscfRegistration instantiates a new ScscfRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScscfRegistration(imsRegistrationType ImsRegistrationType, cscfServerName string) *ScscfRegistration {
	this := ScscfRegistration{}
	this.ImsRegistrationType = imsRegistrationType
	this.CscfServerName = cscfServerName
	var pcscfRestorationIndicator bool = false
	this.PcscfRestorationIndicator = &pcscfRestorationIndicator
	var scscfReselectionIndicator bool = false
	this.ScscfReselectionIndicator = &scscfReselectionIndicator
	return &this
}

// NewScscfRegistrationWithDefaults instantiates a new ScscfRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScscfRegistrationWithDefaults() *ScscfRegistration {
	this := ScscfRegistration{}
	var pcscfRestorationIndicator bool = false
	this.PcscfRestorationIndicator = &pcscfRestorationIndicator
	var scscfReselectionIndicator bool = false
	this.ScscfReselectionIndicator = &scscfReselectionIndicator
	return &this
}

// GetImpi returns the Impi field value if set, zero value otherwise.
func (o *ScscfRegistration) GetImpi() string {
	if o == nil || isNil(o.Impi) {
		var ret string
		return ret
	}
	return *o.Impi
}

// GetImpiOk returns a tuple with the Impi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetImpiOk() (*string, bool) {
	if o == nil || isNil(o.Impi) {
    return nil, false
	}
	return o.Impi, true
}

// HasImpi returns a boolean if a field has been set.
func (o *ScscfRegistration) HasImpi() bool {
	if o != nil && !isNil(o.Impi) {
		return true
	}

	return false
}

// SetImpi gets a reference to the given string and assigns it to the Impi field.
func (o *ScscfRegistration) SetImpi(v string) {
	o.Impi = &v
}

// GetImsRegistrationType returns the ImsRegistrationType field value
func (o *ScscfRegistration) GetImsRegistrationType() ImsRegistrationType {
	if o == nil {
		var ret ImsRegistrationType
		return ret
	}

	return o.ImsRegistrationType
}

// GetImsRegistrationTypeOk returns a tuple with the ImsRegistrationType field value
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetImsRegistrationTypeOk() (*ImsRegistrationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ImsRegistrationType, true
}

// SetImsRegistrationType sets field value
func (o *ScscfRegistration) SetImsRegistrationType(v ImsRegistrationType) {
	o.ImsRegistrationType = v
}

// GetCscfServerName returns the CscfServerName field value
func (o *ScscfRegistration) GetCscfServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CscfServerName
}

// GetCscfServerNameOk returns a tuple with the CscfServerName field value
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetCscfServerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CscfServerName, true
}

// SetCscfServerName sets field value
func (o *ScscfRegistration) SetCscfServerName(v string) {
	o.CscfServerName = v
}

// GetScscfInstanceId returns the ScscfInstanceId field value if set, zero value otherwise.
func (o *ScscfRegistration) GetScscfInstanceId() string {
	if o == nil || isNil(o.ScscfInstanceId) {
		var ret string
		return ret
	}
	return *o.ScscfInstanceId
}

// GetScscfInstanceIdOk returns a tuple with the ScscfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetScscfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.ScscfInstanceId) {
    return nil, false
	}
	return o.ScscfInstanceId, true
}

// HasScscfInstanceId returns a boolean if a field has been set.
func (o *ScscfRegistration) HasScscfInstanceId() bool {
	if o != nil && !isNil(o.ScscfInstanceId) {
		return true
	}

	return false
}

// SetScscfInstanceId gets a reference to the given string and assigns it to the ScscfInstanceId field.
func (o *ScscfRegistration) SetScscfInstanceId(v string) {
	o.ScscfInstanceId = &v
}

// GetDeregCallbackUri returns the DeregCallbackUri field value if set, zero value otherwise.
func (o *ScscfRegistration) GetDeregCallbackUri() string {
	if o == nil || isNil(o.DeregCallbackUri) {
		var ret string
		return ret
	}
	return *o.DeregCallbackUri
}

// GetDeregCallbackUriOk returns a tuple with the DeregCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetDeregCallbackUriOk() (*string, bool) {
	if o == nil || isNil(o.DeregCallbackUri) {
    return nil, false
	}
	return o.DeregCallbackUri, true
}

// HasDeregCallbackUri returns a boolean if a field has been set.
func (o *ScscfRegistration) HasDeregCallbackUri() bool {
	if o != nil && !isNil(o.DeregCallbackUri) {
		return true
	}

	return false
}

// SetDeregCallbackUri gets a reference to the given string and assigns it to the DeregCallbackUri field.
func (o *ScscfRegistration) SetDeregCallbackUri(v string) {
	o.DeregCallbackUri = &v
}

// GetAssociatedImpis returns the AssociatedImpis field value if set, zero value otherwise.
func (o *ScscfRegistration) GetAssociatedImpis() []string {
	if o == nil || isNil(o.AssociatedImpis) {
		var ret []string
		return ret
	}
	return o.AssociatedImpis
}

// GetAssociatedImpisOk returns a tuple with the AssociatedImpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetAssociatedImpisOk() ([]string, bool) {
	if o == nil || isNil(o.AssociatedImpis) {
    return nil, false
	}
	return o.AssociatedImpis, true
}

// HasAssociatedImpis returns a boolean if a field has been set.
func (o *ScscfRegistration) HasAssociatedImpis() bool {
	if o != nil && !isNil(o.AssociatedImpis) {
		return true
	}

	return false
}

// SetAssociatedImpis gets a reference to the given []string and assigns it to the AssociatedImpis field.
func (o *ScscfRegistration) SetAssociatedImpis(v []string) {
	o.AssociatedImpis = v
}

// GetAssociatedRegisteredImpis returns the AssociatedRegisteredImpis field value if set, zero value otherwise.
func (o *ScscfRegistration) GetAssociatedRegisteredImpis() []string {
	if o == nil || isNil(o.AssociatedRegisteredImpis) {
		var ret []string
		return ret
	}
	return o.AssociatedRegisteredImpis
}

// GetAssociatedRegisteredImpisOk returns a tuple with the AssociatedRegisteredImpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetAssociatedRegisteredImpisOk() ([]string, bool) {
	if o == nil || isNil(o.AssociatedRegisteredImpis) {
    return nil, false
	}
	return o.AssociatedRegisteredImpis, true
}

// HasAssociatedRegisteredImpis returns a boolean if a field has been set.
func (o *ScscfRegistration) HasAssociatedRegisteredImpis() bool {
	if o != nil && !isNil(o.AssociatedRegisteredImpis) {
		return true
	}

	return false
}

// SetAssociatedRegisteredImpis gets a reference to the given []string and assigns it to the AssociatedRegisteredImpis field.
func (o *ScscfRegistration) SetAssociatedRegisteredImpis(v []string) {
	o.AssociatedRegisteredImpis = v
}

// GetIrsImpus returns the IrsImpus field value if set, zero value otherwise.
func (o *ScscfRegistration) GetIrsImpus() []string {
	if o == nil || isNil(o.IrsImpus) {
		var ret []string
		return ret
	}
	return o.IrsImpus
}

// GetIrsImpusOk returns a tuple with the IrsImpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetIrsImpusOk() ([]string, bool) {
	if o == nil || isNil(o.IrsImpus) {
    return nil, false
	}
	return o.IrsImpus, true
}

// HasIrsImpus returns a boolean if a field has been set.
func (o *ScscfRegistration) HasIrsImpus() bool {
	if o != nil && !isNil(o.IrsImpus) {
		return true
	}

	return false
}

// SetIrsImpus gets a reference to the given []string and assigns it to the IrsImpus field.
func (o *ScscfRegistration) SetIrsImpus(v []string) {
	o.IrsImpus = v
}

// GetWildcardedPui returns the WildcardedPui field value if set, zero value otherwise.
func (o *ScscfRegistration) GetWildcardedPui() string {
	if o == nil || isNil(o.WildcardedPui) {
		var ret string
		return ret
	}
	return *o.WildcardedPui
}

// GetWildcardedPuiOk returns a tuple with the WildcardedPui field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetWildcardedPuiOk() (*string, bool) {
	if o == nil || isNil(o.WildcardedPui) {
    return nil, false
	}
	return o.WildcardedPui, true
}

// HasWildcardedPui returns a boolean if a field has been set.
func (o *ScscfRegistration) HasWildcardedPui() bool {
	if o != nil && !isNil(o.WildcardedPui) {
		return true
	}

	return false
}

// SetWildcardedPui gets a reference to the given string and assigns it to the WildcardedPui field.
func (o *ScscfRegistration) SetWildcardedPui(v string) {
	o.WildcardedPui = &v
}

// GetLooseRouteIndicator returns the LooseRouteIndicator field value if set, zero value otherwise.
func (o *ScscfRegistration) GetLooseRouteIndicator() LooseRouteIndication {
	if o == nil || isNil(o.LooseRouteIndicator) {
		var ret LooseRouteIndication
		return ret
	}
	return *o.LooseRouteIndicator
}

// GetLooseRouteIndicatorOk returns a tuple with the LooseRouteIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetLooseRouteIndicatorOk() (*LooseRouteIndication, bool) {
	if o == nil || isNil(o.LooseRouteIndicator) {
    return nil, false
	}
	return o.LooseRouteIndicator, true
}

// HasLooseRouteIndicator returns a boolean if a field has been set.
func (o *ScscfRegistration) HasLooseRouteIndicator() bool {
	if o != nil && !isNil(o.LooseRouteIndicator) {
		return true
	}

	return false
}

// SetLooseRouteIndicator gets a reference to the given LooseRouteIndication and assigns it to the LooseRouteIndicator field.
func (o *ScscfRegistration) SetLooseRouteIndicator(v LooseRouteIndication) {
	o.LooseRouteIndicator = &v
}

// GetWildcardedPsi returns the WildcardedPsi field value if set, zero value otherwise.
func (o *ScscfRegistration) GetWildcardedPsi() string {
	if o == nil || isNil(o.WildcardedPsi) {
		var ret string
		return ret
	}
	return *o.WildcardedPsi
}

// GetWildcardedPsiOk returns a tuple with the WildcardedPsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetWildcardedPsiOk() (*string, bool) {
	if o == nil || isNil(o.WildcardedPsi) {
    return nil, false
	}
	return o.WildcardedPsi, true
}

// HasWildcardedPsi returns a boolean if a field has been set.
func (o *ScscfRegistration) HasWildcardedPsi() bool {
	if o != nil && !isNil(o.WildcardedPsi) {
		return true
	}

	return false
}

// SetWildcardedPsi gets a reference to the given string and assigns it to the WildcardedPsi field.
func (o *ScscfRegistration) SetWildcardedPsi(v string) {
	o.WildcardedPsi = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ScscfRegistration) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ScscfRegistration) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ScscfRegistration) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetMultipleRegistrationIndicator returns the MultipleRegistrationIndicator field value if set, zero value otherwise.
func (o *ScscfRegistration) GetMultipleRegistrationIndicator() bool {
	if o == nil || isNil(o.MultipleRegistrationIndicator) {
		var ret bool
		return ret
	}
	return *o.MultipleRegistrationIndicator
}

// GetMultipleRegistrationIndicatorOk returns a tuple with the MultipleRegistrationIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetMultipleRegistrationIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.MultipleRegistrationIndicator) {
    return nil, false
	}
	return o.MultipleRegistrationIndicator, true
}

// HasMultipleRegistrationIndicator returns a boolean if a field has been set.
func (o *ScscfRegistration) HasMultipleRegistrationIndicator() bool {
	if o != nil && !isNil(o.MultipleRegistrationIndicator) {
		return true
	}

	return false
}

// SetMultipleRegistrationIndicator gets a reference to the given bool and assigns it to the MultipleRegistrationIndicator field.
func (o *ScscfRegistration) SetMultipleRegistrationIndicator(v bool) {
	o.MultipleRegistrationIndicator = &v
}

// GetPcscfRestorationIndicator returns the PcscfRestorationIndicator field value if set, zero value otherwise.
func (o *ScscfRegistration) GetPcscfRestorationIndicator() bool {
	if o == nil || isNil(o.PcscfRestorationIndicator) {
		var ret bool
		return ret
	}
	return *o.PcscfRestorationIndicator
}

// GetPcscfRestorationIndicatorOk returns a tuple with the PcscfRestorationIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetPcscfRestorationIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.PcscfRestorationIndicator) {
    return nil, false
	}
	return o.PcscfRestorationIndicator, true
}

// HasPcscfRestorationIndicator returns a boolean if a field has been set.
func (o *ScscfRegistration) HasPcscfRestorationIndicator() bool {
	if o != nil && !isNil(o.PcscfRestorationIndicator) {
		return true
	}

	return false
}

// SetPcscfRestorationIndicator gets a reference to the given bool and assigns it to the PcscfRestorationIndicator field.
func (o *ScscfRegistration) SetPcscfRestorationIndicator(v bool) {
	o.PcscfRestorationIndicator = &v
}

// GetScscfReselectionIndicator returns the ScscfReselectionIndicator field value if set, zero value otherwise.
func (o *ScscfRegistration) GetScscfReselectionIndicator() bool {
	if o == nil || isNil(o.ScscfReselectionIndicator) {
		var ret bool
		return ret
	}
	return *o.ScscfReselectionIndicator
}

// GetScscfReselectionIndicatorOk returns a tuple with the ScscfReselectionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRegistration) GetScscfReselectionIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.ScscfReselectionIndicator) {
    return nil, false
	}
	return o.ScscfReselectionIndicator, true
}

// HasScscfReselectionIndicator returns a boolean if a field has been set.
func (o *ScscfRegistration) HasScscfReselectionIndicator() bool {
	if o != nil && !isNil(o.ScscfReselectionIndicator) {
		return true
	}

	return false
}

// SetScscfReselectionIndicator gets a reference to the given bool and assigns it to the ScscfReselectionIndicator field.
func (o *ScscfRegistration) SetScscfReselectionIndicator(v bool) {
	o.ScscfReselectionIndicator = &v
}

func (o ScscfRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Impi) {
		toSerialize["impi"] = o.Impi
	}
	if true {
		toSerialize["imsRegistrationType"] = o.ImsRegistrationType
	}
	if true {
		toSerialize["cscfServerName"] = o.CscfServerName
	}
	if !isNil(o.ScscfInstanceId) {
		toSerialize["scscfInstanceId"] = o.ScscfInstanceId
	}
	if !isNil(o.DeregCallbackUri) {
		toSerialize["deregCallbackUri"] = o.DeregCallbackUri
	}
	if !isNil(o.AssociatedImpis) {
		toSerialize["associatedImpis"] = o.AssociatedImpis
	}
	if !isNil(o.AssociatedRegisteredImpis) {
		toSerialize["associatedRegisteredImpis"] = o.AssociatedRegisteredImpis
	}
	if !isNil(o.IrsImpus) {
		toSerialize["irsImpus"] = o.IrsImpus
	}
	if !isNil(o.WildcardedPui) {
		toSerialize["wildcardedPui"] = o.WildcardedPui
	}
	if !isNil(o.LooseRouteIndicator) {
		toSerialize["looseRouteIndicator"] = o.LooseRouteIndicator
	}
	if !isNil(o.WildcardedPsi) {
		toSerialize["wildcardedPsi"] = o.WildcardedPsi
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.MultipleRegistrationIndicator) {
		toSerialize["multipleRegistrationIndicator"] = o.MultipleRegistrationIndicator
	}
	if !isNil(o.PcscfRestorationIndicator) {
		toSerialize["pcscfRestorationIndicator"] = o.PcscfRestorationIndicator
	}
	if !isNil(o.ScscfReselectionIndicator) {
		toSerialize["scscfReselectionIndicator"] = o.ScscfReselectionIndicator
	}
	return json.Marshal(toSerialize)
}

type NullableScscfRegistration struct {
	value *ScscfRegistration
	isSet bool
}

func (v NullableScscfRegistration) Get() *ScscfRegistration {
	return v.value
}

func (v *NullableScscfRegistration) Set(val *ScscfRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableScscfRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableScscfRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScscfRegistration(val *ScscfRegistration) *NullableScscfRegistration {
	return &NullableScscfRegistration{value: val, isSet: true}
}

func (v NullableScscfRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScscfRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


