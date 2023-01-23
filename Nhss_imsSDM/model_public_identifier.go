/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// PublicIdentifier Distinct or wildcarded public identity and its associated priority, trace and barring information 
type PublicIdentifier struct {
	PublicIdentity PublicIdentity `json:"publicIdentity"`
	DisplayName *string `json:"displayName,omitempty"`
	ImsServicePriority *PriorityLevels `json:"imsServicePriority,omitempty"`
	ServiceLevelTraceInfo *ServiceLevelTraceInformation `json:"serviceLevelTraceInfo,omitempty"`
	BarringIndicator *bool `json:"barringIndicator,omitempty"`
	WildcardedImpu *string `json:"wildcardedImpu,omitempty"`
}

// NewPublicIdentifier instantiates a new PublicIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIdentifier(publicIdentity PublicIdentity) *PublicIdentifier {
	this := PublicIdentifier{}
	this.PublicIdentity = publicIdentity
	return &this
}

// NewPublicIdentifierWithDefaults instantiates a new PublicIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIdentifierWithDefaults() *PublicIdentifier {
	this := PublicIdentifier{}
	return &this
}

// GetPublicIdentity returns the PublicIdentity field value
func (o *PublicIdentifier) GetPublicIdentity() PublicIdentity {
	if o == nil {
		var ret PublicIdentity
		return ret
	}

	return o.PublicIdentity
}

// GetPublicIdentityOk returns a tuple with the PublicIdentity field value
// and a boolean to check if the value has been set.
func (o *PublicIdentifier) GetPublicIdentityOk() (*PublicIdentity, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublicIdentity, true
}

// SetPublicIdentity sets field value
func (o *PublicIdentifier) SetPublicIdentity(v PublicIdentity) {
	o.PublicIdentity = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PublicIdentifier) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentifier) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PublicIdentifier) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PublicIdentifier) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetImsServicePriority returns the ImsServicePriority field value if set, zero value otherwise.
func (o *PublicIdentifier) GetImsServicePriority() PriorityLevels {
	if o == nil || isNil(o.ImsServicePriority) {
		var ret PriorityLevels
		return ret
	}
	return *o.ImsServicePriority
}

// GetImsServicePriorityOk returns a tuple with the ImsServicePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentifier) GetImsServicePriorityOk() (*PriorityLevels, bool) {
	if o == nil || isNil(o.ImsServicePriority) {
    return nil, false
	}
	return o.ImsServicePriority, true
}

// HasImsServicePriority returns a boolean if a field has been set.
func (o *PublicIdentifier) HasImsServicePriority() bool {
	if o != nil && !isNil(o.ImsServicePriority) {
		return true
	}

	return false
}

// SetImsServicePriority gets a reference to the given PriorityLevels and assigns it to the ImsServicePriority field.
func (o *PublicIdentifier) SetImsServicePriority(v PriorityLevels) {
	o.ImsServicePriority = &v
}

// GetServiceLevelTraceInfo returns the ServiceLevelTraceInfo field value if set, zero value otherwise.
func (o *PublicIdentifier) GetServiceLevelTraceInfo() ServiceLevelTraceInformation {
	if o == nil || isNil(o.ServiceLevelTraceInfo) {
		var ret ServiceLevelTraceInformation
		return ret
	}
	return *o.ServiceLevelTraceInfo
}

// GetServiceLevelTraceInfoOk returns a tuple with the ServiceLevelTraceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentifier) GetServiceLevelTraceInfoOk() (*ServiceLevelTraceInformation, bool) {
	if o == nil || isNil(o.ServiceLevelTraceInfo) {
    return nil, false
	}
	return o.ServiceLevelTraceInfo, true
}

// HasServiceLevelTraceInfo returns a boolean if a field has been set.
func (o *PublicIdentifier) HasServiceLevelTraceInfo() bool {
	if o != nil && !isNil(o.ServiceLevelTraceInfo) {
		return true
	}

	return false
}

// SetServiceLevelTraceInfo gets a reference to the given ServiceLevelTraceInformation and assigns it to the ServiceLevelTraceInfo field.
func (o *PublicIdentifier) SetServiceLevelTraceInfo(v ServiceLevelTraceInformation) {
	o.ServiceLevelTraceInfo = &v
}

// GetBarringIndicator returns the BarringIndicator field value if set, zero value otherwise.
func (o *PublicIdentifier) GetBarringIndicator() bool {
	if o == nil || isNil(o.BarringIndicator) {
		var ret bool
		return ret
	}
	return *o.BarringIndicator
}

// GetBarringIndicatorOk returns a tuple with the BarringIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentifier) GetBarringIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.BarringIndicator) {
    return nil, false
	}
	return o.BarringIndicator, true
}

// HasBarringIndicator returns a boolean if a field has been set.
func (o *PublicIdentifier) HasBarringIndicator() bool {
	if o != nil && !isNil(o.BarringIndicator) {
		return true
	}

	return false
}

// SetBarringIndicator gets a reference to the given bool and assigns it to the BarringIndicator field.
func (o *PublicIdentifier) SetBarringIndicator(v bool) {
	o.BarringIndicator = &v
}

// GetWildcardedImpu returns the WildcardedImpu field value if set, zero value otherwise.
func (o *PublicIdentifier) GetWildcardedImpu() string {
	if o == nil || isNil(o.WildcardedImpu) {
		var ret string
		return ret
	}
	return *o.WildcardedImpu
}

// GetWildcardedImpuOk returns a tuple with the WildcardedImpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentifier) GetWildcardedImpuOk() (*string, bool) {
	if o == nil || isNil(o.WildcardedImpu) {
    return nil, false
	}
	return o.WildcardedImpu, true
}

// HasWildcardedImpu returns a boolean if a field has been set.
func (o *PublicIdentifier) HasWildcardedImpu() bool {
	if o != nil && !isNil(o.WildcardedImpu) {
		return true
	}

	return false
}

// SetWildcardedImpu gets a reference to the given string and assigns it to the WildcardedImpu field.
func (o *PublicIdentifier) SetWildcardedImpu(v string) {
	o.WildcardedImpu = &v
}

func (o PublicIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publicIdentity"] = o.PublicIdentity
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.ImsServicePriority) {
		toSerialize["imsServicePriority"] = o.ImsServicePriority
	}
	if !isNil(o.ServiceLevelTraceInfo) {
		toSerialize["serviceLevelTraceInfo"] = o.ServiceLevelTraceInfo
	}
	if !isNil(o.BarringIndicator) {
		toSerialize["barringIndicator"] = o.BarringIndicator
	}
	if !isNil(o.WildcardedImpu) {
		toSerialize["wildcardedImpu"] = o.WildcardedImpu
	}
	return json.Marshal(toSerialize)
}

type NullablePublicIdentifier struct {
	value *PublicIdentifier
	isSet bool
}

func (v NullablePublicIdentifier) Get() *PublicIdentifier {
	return v.value
}

func (v *NullablePublicIdentifier) Set(val *PublicIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIdentifier(val *PublicIdentifier) *NullablePublicIdentifier {
	return &NullablePublicIdentifier{value: val, isSet: true}
}

func (v NullablePublicIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


