/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsUECM

import (
	"encoding/json"
)

// AuthorizationRequest Ims authorization request data
type AuthorizationRequest struct {
	// IMS Private Identity of the UE
	Impi *string `json:"impi,omitempty"`
	AuthorizationType AuthorizationType `json:"authorizationType"`
	VisitedNetworkIdentifier *string `json:"visitedNetworkIdentifier,omitempty"`
	EmergencyIndicator *bool `json:"emergencyIndicator,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewAuthorizationRequest instantiates a new AuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationRequest(authorizationType AuthorizationType) *AuthorizationRequest {
	this := AuthorizationRequest{}
	this.AuthorizationType = authorizationType
	return &this
}

// NewAuthorizationRequestWithDefaults instantiates a new AuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationRequestWithDefaults() *AuthorizationRequest {
	this := AuthorizationRequest{}
	return &this
}

// GetImpi returns the Impi field value if set, zero value otherwise.
func (o *AuthorizationRequest) GetImpi() string {
	if o == nil || isNil(o.Impi) {
		var ret string
		return ret
	}
	return *o.Impi
}

// GetImpiOk returns a tuple with the Impi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetImpiOk() (*string, bool) {
	if o == nil || isNil(o.Impi) {
    return nil, false
	}
	return o.Impi, true
}

// HasImpi returns a boolean if a field has been set.
func (o *AuthorizationRequest) HasImpi() bool {
	if o != nil && !isNil(o.Impi) {
		return true
	}

	return false
}

// SetImpi gets a reference to the given string and assigns it to the Impi field.
func (o *AuthorizationRequest) SetImpi(v string) {
	o.Impi = &v
}

// GetAuthorizationType returns the AuthorizationType field value
func (o *AuthorizationRequest) GetAuthorizationType() AuthorizationType {
	if o == nil {
		var ret AuthorizationType
		return ret
	}

	return o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetAuthorizationTypeOk() (*AuthorizationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthorizationType, true
}

// SetAuthorizationType sets field value
func (o *AuthorizationRequest) SetAuthorizationType(v AuthorizationType) {
	o.AuthorizationType = v
}

// GetVisitedNetworkIdentifier returns the VisitedNetworkIdentifier field value if set, zero value otherwise.
func (o *AuthorizationRequest) GetVisitedNetworkIdentifier() string {
	if o == nil || isNil(o.VisitedNetworkIdentifier) {
		var ret string
		return ret
	}
	return *o.VisitedNetworkIdentifier
}

// GetVisitedNetworkIdentifierOk returns a tuple with the VisitedNetworkIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetVisitedNetworkIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.VisitedNetworkIdentifier) {
    return nil, false
	}
	return o.VisitedNetworkIdentifier, true
}

// HasVisitedNetworkIdentifier returns a boolean if a field has been set.
func (o *AuthorizationRequest) HasVisitedNetworkIdentifier() bool {
	if o != nil && !isNil(o.VisitedNetworkIdentifier) {
		return true
	}

	return false
}

// SetVisitedNetworkIdentifier gets a reference to the given string and assigns it to the VisitedNetworkIdentifier field.
func (o *AuthorizationRequest) SetVisitedNetworkIdentifier(v string) {
	o.VisitedNetworkIdentifier = &v
}

// GetEmergencyIndicator returns the EmergencyIndicator field value if set, zero value otherwise.
func (o *AuthorizationRequest) GetEmergencyIndicator() bool {
	if o == nil || isNil(o.EmergencyIndicator) {
		var ret bool
		return ret
	}
	return *o.EmergencyIndicator
}

// GetEmergencyIndicatorOk returns a tuple with the EmergencyIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetEmergencyIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.EmergencyIndicator) {
    return nil, false
	}
	return o.EmergencyIndicator, true
}

// HasEmergencyIndicator returns a boolean if a field has been set.
func (o *AuthorizationRequest) HasEmergencyIndicator() bool {
	if o != nil && !isNil(o.EmergencyIndicator) {
		return true
	}

	return false
}

// SetEmergencyIndicator gets a reference to the given bool and assigns it to the EmergencyIndicator field.
func (o *AuthorizationRequest) SetEmergencyIndicator(v bool) {
	o.EmergencyIndicator = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AuthorizationRequest) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AuthorizationRequest) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AuthorizationRequest) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o AuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Impi) {
		toSerialize["impi"] = o.Impi
	}
	if true {
		toSerialize["authorizationType"] = o.AuthorizationType
	}
	if !isNil(o.VisitedNetworkIdentifier) {
		toSerialize["visitedNetworkIdentifier"] = o.VisitedNetworkIdentifier
	}
	if !isNil(o.EmergencyIndicator) {
		toSerialize["emergencyIndicator"] = o.EmergencyIndicator
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationRequest struct {
	value *AuthorizationRequest
	isSet bool
}

func (v NullableAuthorizationRequest) Get() *AuthorizationRequest {
	return v.value
}

func (v *NullableAuthorizationRequest) Set(val *AuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationRequest(val *AuthorizationRequest) *NullableAuthorizationRequest {
	return &NullableAuthorizationRequest{value: val, isSet: true}
}

func (v NullableAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


