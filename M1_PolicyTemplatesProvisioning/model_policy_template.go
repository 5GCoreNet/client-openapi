/*
M1_PolicyTemplatesProvisioning

5GMS AF M1 Policy Templates Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_PolicyTemplatesProvisioning

import (
	"encoding/json"
)

// PolicyTemplate A representation of a Policy Template resource.
type PolicyTemplate struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	PolicyTemplateId string `json:"policyTemplateId"`
	State PolicyTemplateState `json:"state"`
	ApiEndPoint string `json:"apiEndPoint"`
	ApiType PolicyTemplateApiType `json:"apiType"`
	ExternalReference string `json:"externalReference"`
	QoSSpecification *M1QoSSpecification `json:"qoSSpecification,omitempty"`
	ApplicationSessionContext PolicyTemplateApplicationSessionContext `json:"applicationSessionContext"`
	ChargingSpecification *ChargingSpecification `json:"chargingSpecification,omitempty"`
}

// NewPolicyTemplate instantiates a new PolicyTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyTemplate(policyTemplateId string, state PolicyTemplateState, apiEndPoint string, apiType PolicyTemplateApiType, externalReference string, applicationSessionContext PolicyTemplateApplicationSessionContext) *PolicyTemplate {
	this := PolicyTemplate{}
	this.PolicyTemplateId = policyTemplateId
	this.State = state
	this.ApiEndPoint = apiEndPoint
	this.ApiType = apiType
	this.ExternalReference = externalReference
	this.ApplicationSessionContext = applicationSessionContext
	return &this
}

// NewPolicyTemplateWithDefaults instantiates a new PolicyTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyTemplateWithDefaults() *PolicyTemplate {
	this := PolicyTemplate{}
	return &this
}

// GetPolicyTemplateId returns the PolicyTemplateId field value
func (o *PolicyTemplate) GetPolicyTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyTemplateId
}

// GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field value
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetPolicyTemplateIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PolicyTemplateId, true
}

// SetPolicyTemplateId sets field value
func (o *PolicyTemplate) SetPolicyTemplateId(v string) {
	o.PolicyTemplateId = v
}

// GetState returns the State field value
func (o *PolicyTemplate) GetState() PolicyTemplateState {
	if o == nil {
		var ret PolicyTemplateState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetStateOk() (*PolicyTemplateState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PolicyTemplate) SetState(v PolicyTemplateState) {
	o.State = v
}

// GetApiEndPoint returns the ApiEndPoint field value
func (o *PolicyTemplate) GetApiEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiEndPoint
}

// GetApiEndPointOk returns a tuple with the ApiEndPoint field value
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetApiEndPointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiEndPoint, true
}

// SetApiEndPoint sets field value
func (o *PolicyTemplate) SetApiEndPoint(v string) {
	o.ApiEndPoint = v
}

// GetApiType returns the ApiType field value
func (o *PolicyTemplate) GetApiType() PolicyTemplateApiType {
	if o == nil {
		var ret PolicyTemplateApiType
		return ret
	}

	return o.ApiType
}

// GetApiTypeOk returns a tuple with the ApiType field value
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetApiTypeOk() (*PolicyTemplateApiType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiType, true
}

// SetApiType sets field value
func (o *PolicyTemplate) SetApiType(v PolicyTemplateApiType) {
	o.ApiType = v
}

// GetExternalReference returns the ExternalReference field value
func (o *PolicyTemplate) GetExternalReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalReference
}

// GetExternalReferenceOk returns a tuple with the ExternalReference field value
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetExternalReferenceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExternalReference, true
}

// SetExternalReference sets field value
func (o *PolicyTemplate) SetExternalReference(v string) {
	o.ExternalReference = v
}

// GetQoSSpecification returns the QoSSpecification field value if set, zero value otherwise.
func (o *PolicyTemplate) GetQoSSpecification() M1QoSSpecification {
	if o == nil || isNil(o.QoSSpecification) {
		var ret M1QoSSpecification
		return ret
	}
	return *o.QoSSpecification
}

// GetQoSSpecificationOk returns a tuple with the QoSSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetQoSSpecificationOk() (*M1QoSSpecification, bool) {
	if o == nil || isNil(o.QoSSpecification) {
    return nil, false
	}
	return o.QoSSpecification, true
}

// HasQoSSpecification returns a boolean if a field has been set.
func (o *PolicyTemplate) HasQoSSpecification() bool {
	if o != nil && !isNil(o.QoSSpecification) {
		return true
	}

	return false
}

// SetQoSSpecification gets a reference to the given M1QoSSpecification and assigns it to the QoSSpecification field.
func (o *PolicyTemplate) SetQoSSpecification(v M1QoSSpecification) {
	o.QoSSpecification = &v
}

// GetApplicationSessionContext returns the ApplicationSessionContext field value
func (o *PolicyTemplate) GetApplicationSessionContext() PolicyTemplateApplicationSessionContext {
	if o == nil {
		var ret PolicyTemplateApplicationSessionContext
		return ret
	}

	return o.ApplicationSessionContext
}

// GetApplicationSessionContextOk returns a tuple with the ApplicationSessionContext field value
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetApplicationSessionContextOk() (*PolicyTemplateApplicationSessionContext, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApplicationSessionContext, true
}

// SetApplicationSessionContext sets field value
func (o *PolicyTemplate) SetApplicationSessionContext(v PolicyTemplateApplicationSessionContext) {
	o.ApplicationSessionContext = v
}

// GetChargingSpecification returns the ChargingSpecification field value if set, zero value otherwise.
func (o *PolicyTemplate) GetChargingSpecification() ChargingSpecification {
	if o == nil || isNil(o.ChargingSpecification) {
		var ret ChargingSpecification
		return ret
	}
	return *o.ChargingSpecification
}

// GetChargingSpecificationOk returns a tuple with the ChargingSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTemplate) GetChargingSpecificationOk() (*ChargingSpecification, bool) {
	if o == nil || isNil(o.ChargingSpecification) {
    return nil, false
	}
	return o.ChargingSpecification, true
}

// HasChargingSpecification returns a boolean if a field has been set.
func (o *PolicyTemplate) HasChargingSpecification() bool {
	if o != nil && !isNil(o.ChargingSpecification) {
		return true
	}

	return false
}

// SetChargingSpecification gets a reference to the given ChargingSpecification and assigns it to the ChargingSpecification field.
func (o *PolicyTemplate) SetChargingSpecification(v ChargingSpecification) {
	o.ChargingSpecification = &v
}

func (o PolicyTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policyTemplateId"] = o.PolicyTemplateId
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["apiEndPoint"] = o.ApiEndPoint
	}
	if true {
		toSerialize["apiType"] = o.ApiType
	}
	if true {
		toSerialize["externalReference"] = o.ExternalReference
	}
	if !isNil(o.QoSSpecification) {
		toSerialize["qoSSpecification"] = o.QoSSpecification
	}
	if true {
		toSerialize["applicationSessionContext"] = o.ApplicationSessionContext
	}
	if !isNil(o.ChargingSpecification) {
		toSerialize["chargingSpecification"] = o.ChargingSpecification
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyTemplate struct {
	value *PolicyTemplate
	isSet bool
}

func (v NullablePolicyTemplate) Get() *PolicyTemplate {
	return v.value
}

func (v *NullablePolicyTemplate) Set(val *PolicyTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyTemplate(val *PolicyTemplate) *NullablePolicyTemplate {
	return &NullablePolicyTemplate{value: val, isSet: true}
}

func (v NullablePolicyTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


