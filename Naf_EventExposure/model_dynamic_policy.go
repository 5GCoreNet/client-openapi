/*
Naf_EventExposure

AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_EventExposure

import (
	"encoding/json"
)

// DynamicPolicy A representation of a Dynamic Policy resource.
type DynamicPolicy struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	DynamicPolicyId string `json:"dynamicPolicyId"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	PolicyTemplateId string `json:"policyTemplateId"`
	ServiceDataFlowDescriptions []ServiceDataFlowDescription `json:"serviceDataFlowDescriptions"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	ProvisioningSessionId string `json:"provisioningSessionId"`
	QosSpecification *M5QoSSpecification `json:"qosSpecification,omitempty"`
	EnforcementMethod *string `json:"enforcementMethod,omitempty"`
	EnforcementBitRate *int32 `json:"enforcementBitRate,omitempty"`
}

// NewDynamicPolicy instantiates a new DynamicPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicPolicy(dynamicPolicyId string, policyTemplateId string, serviceDataFlowDescriptions []ServiceDataFlowDescription, provisioningSessionId string) *DynamicPolicy {
	this := DynamicPolicy{}
	this.DynamicPolicyId = dynamicPolicyId
	this.PolicyTemplateId = policyTemplateId
	this.ServiceDataFlowDescriptions = serviceDataFlowDescriptions
	this.ProvisioningSessionId = provisioningSessionId
	return &this
}

// NewDynamicPolicyWithDefaults instantiates a new DynamicPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicPolicyWithDefaults() *DynamicPolicy {
	this := DynamicPolicy{}
	return &this
}

// GetDynamicPolicyId returns the DynamicPolicyId field value
func (o *DynamicPolicy) GetDynamicPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DynamicPolicyId
}

// GetDynamicPolicyIdOk returns a tuple with the DynamicPolicyId field value
// and a boolean to check if the value has been set.
func (o *DynamicPolicy) GetDynamicPolicyIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DynamicPolicyId, true
}

// SetDynamicPolicyId sets field value
func (o *DynamicPolicy) SetDynamicPolicyId(v string) {
	o.DynamicPolicyId = v
}

// GetPolicyTemplateId returns the PolicyTemplateId field value
func (o *DynamicPolicy) GetPolicyTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyTemplateId
}

// GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field value
// and a boolean to check if the value has been set.
func (o *DynamicPolicy) GetPolicyTemplateIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PolicyTemplateId, true
}

// SetPolicyTemplateId sets field value
func (o *DynamicPolicy) SetPolicyTemplateId(v string) {
	o.PolicyTemplateId = v
}

// GetServiceDataFlowDescriptions returns the ServiceDataFlowDescriptions field value
func (o *DynamicPolicy) GetServiceDataFlowDescriptions() []ServiceDataFlowDescription {
	if o == nil {
		var ret []ServiceDataFlowDescription
		return ret
	}

	return o.ServiceDataFlowDescriptions
}

// GetServiceDataFlowDescriptionsOk returns a tuple with the ServiceDataFlowDescriptions field value
// and a boolean to check if the value has been set.
func (o *DynamicPolicy) GetServiceDataFlowDescriptionsOk() ([]ServiceDataFlowDescription, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServiceDataFlowDescriptions, true
}

// SetServiceDataFlowDescriptions sets field value
func (o *DynamicPolicy) SetServiceDataFlowDescriptions(v []ServiceDataFlowDescription) {
	o.ServiceDataFlowDescriptions = v
}

// GetProvisioningSessionId returns the ProvisioningSessionId field value
func (o *DynamicPolicy) GetProvisioningSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisioningSessionId
}

// GetProvisioningSessionIdOk returns a tuple with the ProvisioningSessionId field value
// and a boolean to check if the value has been set.
func (o *DynamicPolicy) GetProvisioningSessionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProvisioningSessionId, true
}

// SetProvisioningSessionId sets field value
func (o *DynamicPolicy) SetProvisioningSessionId(v string) {
	o.ProvisioningSessionId = v
}

// GetQosSpecification returns the QosSpecification field value if set, zero value otherwise.
func (o *DynamicPolicy) GetQosSpecification() M5QoSSpecification {
	if o == nil || isNil(o.QosSpecification) {
		var ret M5QoSSpecification
		return ret
	}
	return *o.QosSpecification
}

// GetQosSpecificationOk returns a tuple with the QosSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicPolicy) GetQosSpecificationOk() (*M5QoSSpecification, bool) {
	if o == nil || isNil(o.QosSpecification) {
    return nil, false
	}
	return o.QosSpecification, true
}

// HasQosSpecification returns a boolean if a field has been set.
func (o *DynamicPolicy) HasQosSpecification() bool {
	if o != nil && !isNil(o.QosSpecification) {
		return true
	}

	return false
}

// SetQosSpecification gets a reference to the given M5QoSSpecification and assigns it to the QosSpecification field.
func (o *DynamicPolicy) SetQosSpecification(v M5QoSSpecification) {
	o.QosSpecification = &v
}

// GetEnforcementMethod returns the EnforcementMethod field value if set, zero value otherwise.
func (o *DynamicPolicy) GetEnforcementMethod() string {
	if o == nil || isNil(o.EnforcementMethod) {
		var ret string
		return ret
	}
	return *o.EnforcementMethod
}

// GetEnforcementMethodOk returns a tuple with the EnforcementMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicPolicy) GetEnforcementMethodOk() (*string, bool) {
	if o == nil || isNil(o.EnforcementMethod) {
    return nil, false
	}
	return o.EnforcementMethod, true
}

// HasEnforcementMethod returns a boolean if a field has been set.
func (o *DynamicPolicy) HasEnforcementMethod() bool {
	if o != nil && !isNil(o.EnforcementMethod) {
		return true
	}

	return false
}

// SetEnforcementMethod gets a reference to the given string and assigns it to the EnforcementMethod field.
func (o *DynamicPolicy) SetEnforcementMethod(v string) {
	o.EnforcementMethod = &v
}

// GetEnforcementBitRate returns the EnforcementBitRate field value if set, zero value otherwise.
func (o *DynamicPolicy) GetEnforcementBitRate() int32 {
	if o == nil || isNil(o.EnforcementBitRate) {
		var ret int32
		return ret
	}
	return *o.EnforcementBitRate
}

// GetEnforcementBitRateOk returns a tuple with the EnforcementBitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicPolicy) GetEnforcementBitRateOk() (*int32, bool) {
	if o == nil || isNil(o.EnforcementBitRate) {
    return nil, false
	}
	return o.EnforcementBitRate, true
}

// HasEnforcementBitRate returns a boolean if a field has been set.
func (o *DynamicPolicy) HasEnforcementBitRate() bool {
	if o != nil && !isNil(o.EnforcementBitRate) {
		return true
	}

	return false
}

// SetEnforcementBitRate gets a reference to the given int32 and assigns it to the EnforcementBitRate field.
func (o *DynamicPolicy) SetEnforcementBitRate(v int32) {
	o.EnforcementBitRate = &v
}

func (o DynamicPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dynamicPolicyId"] = o.DynamicPolicyId
	}
	if true {
		toSerialize["policyTemplateId"] = o.PolicyTemplateId
	}
	if true {
		toSerialize["serviceDataFlowDescriptions"] = o.ServiceDataFlowDescriptions
	}
	if true {
		toSerialize["provisioningSessionId"] = o.ProvisioningSessionId
	}
	if !isNil(o.QosSpecification) {
		toSerialize["qosSpecification"] = o.QosSpecification
	}
	if !isNil(o.EnforcementMethod) {
		toSerialize["enforcementMethod"] = o.EnforcementMethod
	}
	if !isNil(o.EnforcementBitRate) {
		toSerialize["enforcementBitRate"] = o.EnforcementBitRate
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicPolicy struct {
	value *DynamicPolicy
	isSet bool
}

func (v NullableDynamicPolicy) Get() *DynamicPolicy {
	return v.value
}

func (v *NullableDynamicPolicy) Set(val *DynamicPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicPolicy(val *DynamicPolicy) *NullableDynamicPolicy {
	return &NullableDynamicPolicy{value: val, isSet: true}
}

func (v NullableDynamicPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


