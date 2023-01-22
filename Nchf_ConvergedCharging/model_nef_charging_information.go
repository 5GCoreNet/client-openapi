/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
)

// NEFChargingInformation struct for NEFChargingInformation
type NEFChargingInformation struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	ExternalIndividualIdentifier *string `json:"externalIndividualIdentifier,omitempty"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExternalGroupIdentifier *string `json:"externalGroupIdentifier,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	GroupIdentifier *string `json:"groupIdentifier,omitempty"`
	APIDirection *APIDirection `json:"aPIDirection,omitempty"`
	APITargetNetworkFunction *NFIdentification `json:"aPITargetNetworkFunction,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	APIResultCode *int32 `json:"aPIResultCode,omitempty"`
	APIName string `json:"aPIName"`
	// String providing an URI formatted according to RFC 3986.
	APIReference *string `json:"aPIReference,omitempty"`
	APIContent *string `json:"aPIContent,omitempty"`
}

// NewNEFChargingInformation instantiates a new NEFChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNEFChargingInformation(aPIName string) *NEFChargingInformation {
	this := NEFChargingInformation{}
	this.APIName = aPIName
	return &this
}

// NewNEFChargingInformationWithDefaults instantiates a new NEFChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNEFChargingInformationWithDefaults() *NEFChargingInformation {
	this := NEFChargingInformation{}
	return &this
}

// GetExternalIndividualIdentifier returns the ExternalIndividualIdentifier field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetExternalIndividualIdentifier() string {
	if o == nil || isNil(o.ExternalIndividualIdentifier) {
		var ret string
		return ret
	}
	return *o.ExternalIndividualIdentifier
}

// GetExternalIndividualIdentifierOk returns a tuple with the ExternalIndividualIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetExternalIndividualIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.ExternalIndividualIdentifier) {
    return nil, false
	}
	return o.ExternalIndividualIdentifier, true
}

// HasExternalIndividualIdentifier returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasExternalIndividualIdentifier() bool {
	if o != nil && !isNil(o.ExternalIndividualIdentifier) {
		return true
	}

	return false
}

// SetExternalIndividualIdentifier gets a reference to the given string and assigns it to the ExternalIndividualIdentifier field.
func (o *NEFChargingInformation) SetExternalIndividualIdentifier(v string) {
	o.ExternalIndividualIdentifier = &v
}

// GetExternalGroupIdentifier returns the ExternalGroupIdentifier field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetExternalGroupIdentifier() string {
	if o == nil || isNil(o.ExternalGroupIdentifier) {
		var ret string
		return ret
	}
	return *o.ExternalGroupIdentifier
}

// GetExternalGroupIdentifierOk returns a tuple with the ExternalGroupIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetExternalGroupIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.ExternalGroupIdentifier) {
    return nil, false
	}
	return o.ExternalGroupIdentifier, true
}

// HasExternalGroupIdentifier returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasExternalGroupIdentifier() bool {
	if o != nil && !isNil(o.ExternalGroupIdentifier) {
		return true
	}

	return false
}

// SetExternalGroupIdentifier gets a reference to the given string and assigns it to the ExternalGroupIdentifier field.
func (o *NEFChargingInformation) SetExternalGroupIdentifier(v string) {
	o.ExternalGroupIdentifier = &v
}

// GetGroupIdentifier returns the GroupIdentifier field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetGroupIdentifier() string {
	if o == nil || isNil(o.GroupIdentifier) {
		var ret string
		return ret
	}
	return *o.GroupIdentifier
}

// GetGroupIdentifierOk returns a tuple with the GroupIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetGroupIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.GroupIdentifier) {
    return nil, false
	}
	return o.GroupIdentifier, true
}

// HasGroupIdentifier returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasGroupIdentifier() bool {
	if o != nil && !isNil(o.GroupIdentifier) {
		return true
	}

	return false
}

// SetGroupIdentifier gets a reference to the given string and assigns it to the GroupIdentifier field.
func (o *NEFChargingInformation) SetGroupIdentifier(v string) {
	o.GroupIdentifier = &v
}

// GetAPIDirection returns the APIDirection field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetAPIDirection() APIDirection {
	if o == nil || isNil(o.APIDirection) {
		var ret APIDirection
		return ret
	}
	return *o.APIDirection
}

// GetAPIDirectionOk returns a tuple with the APIDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetAPIDirectionOk() (*APIDirection, bool) {
	if o == nil || isNil(o.APIDirection) {
    return nil, false
	}
	return o.APIDirection, true
}

// HasAPIDirection returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasAPIDirection() bool {
	if o != nil && !isNil(o.APIDirection) {
		return true
	}

	return false
}

// SetAPIDirection gets a reference to the given APIDirection and assigns it to the APIDirection field.
func (o *NEFChargingInformation) SetAPIDirection(v APIDirection) {
	o.APIDirection = &v
}

// GetAPITargetNetworkFunction returns the APITargetNetworkFunction field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetAPITargetNetworkFunction() NFIdentification {
	if o == nil || isNil(o.APITargetNetworkFunction) {
		var ret NFIdentification
		return ret
	}
	return *o.APITargetNetworkFunction
}

// GetAPITargetNetworkFunctionOk returns a tuple with the APITargetNetworkFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetAPITargetNetworkFunctionOk() (*NFIdentification, bool) {
	if o == nil || isNil(o.APITargetNetworkFunction) {
    return nil, false
	}
	return o.APITargetNetworkFunction, true
}

// HasAPITargetNetworkFunction returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasAPITargetNetworkFunction() bool {
	if o != nil && !isNil(o.APITargetNetworkFunction) {
		return true
	}

	return false
}

// SetAPITargetNetworkFunction gets a reference to the given NFIdentification and assigns it to the APITargetNetworkFunction field.
func (o *NEFChargingInformation) SetAPITargetNetworkFunction(v NFIdentification) {
	o.APITargetNetworkFunction = &v
}

// GetAPIResultCode returns the APIResultCode field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetAPIResultCode() int32 {
	if o == nil || isNil(o.APIResultCode) {
		var ret int32
		return ret
	}
	return *o.APIResultCode
}

// GetAPIResultCodeOk returns a tuple with the APIResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetAPIResultCodeOk() (*int32, bool) {
	if o == nil || isNil(o.APIResultCode) {
    return nil, false
	}
	return o.APIResultCode, true
}

// HasAPIResultCode returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasAPIResultCode() bool {
	if o != nil && !isNil(o.APIResultCode) {
		return true
	}

	return false
}

// SetAPIResultCode gets a reference to the given int32 and assigns it to the APIResultCode field.
func (o *NEFChargingInformation) SetAPIResultCode(v int32) {
	o.APIResultCode = &v
}

// GetAPIName returns the APIName field value
func (o *NEFChargingInformation) GetAPIName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.APIName
}

// GetAPINameOk returns a tuple with the APIName field value
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetAPINameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.APIName, true
}

// SetAPIName sets field value
func (o *NEFChargingInformation) SetAPIName(v string) {
	o.APIName = v
}

// GetAPIReference returns the APIReference field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetAPIReference() string {
	if o == nil || isNil(o.APIReference) {
		var ret string
		return ret
	}
	return *o.APIReference
}

// GetAPIReferenceOk returns a tuple with the APIReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetAPIReferenceOk() (*string, bool) {
	if o == nil || isNil(o.APIReference) {
    return nil, false
	}
	return o.APIReference, true
}

// HasAPIReference returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasAPIReference() bool {
	if o != nil && !isNil(o.APIReference) {
		return true
	}

	return false
}

// SetAPIReference gets a reference to the given string and assigns it to the APIReference field.
func (o *NEFChargingInformation) SetAPIReference(v string) {
	o.APIReference = &v
}

// GetAPIContent returns the APIContent field value if set, zero value otherwise.
func (o *NEFChargingInformation) GetAPIContent() string {
	if o == nil || isNil(o.APIContent) {
		var ret string
		return ret
	}
	return *o.APIContent
}

// GetAPIContentOk returns a tuple with the APIContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NEFChargingInformation) GetAPIContentOk() (*string, bool) {
	if o == nil || isNil(o.APIContent) {
    return nil, false
	}
	return o.APIContent, true
}

// HasAPIContent returns a boolean if a field has been set.
func (o *NEFChargingInformation) HasAPIContent() bool {
	if o != nil && !isNil(o.APIContent) {
		return true
	}

	return false
}

// SetAPIContent gets a reference to the given string and assigns it to the APIContent field.
func (o *NEFChargingInformation) SetAPIContent(v string) {
	o.APIContent = &v
}

func (o NEFChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalIndividualIdentifier) {
		toSerialize["externalIndividualIdentifier"] = o.ExternalIndividualIdentifier
	}
	if !isNil(o.ExternalGroupIdentifier) {
		toSerialize["externalGroupIdentifier"] = o.ExternalGroupIdentifier
	}
	if !isNil(o.GroupIdentifier) {
		toSerialize["groupIdentifier"] = o.GroupIdentifier
	}
	if !isNil(o.APIDirection) {
		toSerialize["aPIDirection"] = o.APIDirection
	}
	if !isNil(o.APITargetNetworkFunction) {
		toSerialize["aPITargetNetworkFunction"] = o.APITargetNetworkFunction
	}
	if !isNil(o.APIResultCode) {
		toSerialize["aPIResultCode"] = o.APIResultCode
	}
	if true {
		toSerialize["aPIName"] = o.APIName
	}
	if !isNil(o.APIReference) {
		toSerialize["aPIReference"] = o.APIReference
	}
	if !isNil(o.APIContent) {
		toSerialize["aPIContent"] = o.APIContent
	}
	return json.Marshal(toSerialize)
}

type NullableNEFChargingInformation struct {
	value *NEFChargingInformation
	isSet bool
}

func (v NullableNEFChargingInformation) Get() *NEFChargingInformation {
	return v.value
}

func (v *NullableNEFChargingInformation) Set(val *NEFChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNEFChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNEFChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNEFChargingInformation(val *NEFChargingInformation) *NullableNEFChargingInformation {
	return &NullableNEFChargingInformation{value: val, isSet: true}
}

func (v NullableNEFChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNEFChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


