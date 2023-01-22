/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package IntentNrm

import (
	"encoding/json"
)

// ServiceSupportExpectation This data type is the \"IntentExpectation\" data type with specialisations to represent MnS consumer's expectations for service deployment    
type ServiceSupportExpectation struct {
	ExpectationId *string `json:"expectationId,omitempty"`
	ExpectationVerb *ExpectationVerb `json:"expectationVerb,omitempty"`
	ExpectationObjects []ServiceSupportExpectationObject `json:"expectationObjects,omitempty"`
	ExpectationTargets []ServiceSupportExpectationExpectationTargetsInner `json:"expectationTargets,omitempty"`
	ExpectationContexts []ServiceSupportExpectationExpectationContextsInner `json:"expectationContexts,omitempty"`
	ExpectationfulfilmentInfo *FulfilmentInfo `json:"expectationfulfilmentInfo,omitempty"`
}

// NewServiceSupportExpectation instantiates a new ServiceSupportExpectation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSupportExpectation() *ServiceSupportExpectation {
	this := ServiceSupportExpectation{}
	return &this
}

// NewServiceSupportExpectationWithDefaults instantiates a new ServiceSupportExpectation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSupportExpectationWithDefaults() *ServiceSupportExpectation {
	this := ServiceSupportExpectation{}
	return &this
}

// GetExpectationId returns the ExpectationId field value if set, zero value otherwise.
func (o *ServiceSupportExpectation) GetExpectationId() string {
	if o == nil || isNil(o.ExpectationId) {
		var ret string
		return ret
	}
	return *o.ExpectationId
}

// GetExpectationIdOk returns a tuple with the ExpectationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectation) GetExpectationIdOk() (*string, bool) {
	if o == nil || isNil(o.ExpectationId) {
    return nil, false
	}
	return o.ExpectationId, true
}

// HasExpectationId returns a boolean if a field has been set.
func (o *ServiceSupportExpectation) HasExpectationId() bool {
	if o != nil && !isNil(o.ExpectationId) {
		return true
	}

	return false
}

// SetExpectationId gets a reference to the given string and assigns it to the ExpectationId field.
func (o *ServiceSupportExpectation) SetExpectationId(v string) {
	o.ExpectationId = &v
}

// GetExpectationVerb returns the ExpectationVerb field value if set, zero value otherwise.
func (o *ServiceSupportExpectation) GetExpectationVerb() ExpectationVerb {
	if o == nil || isNil(o.ExpectationVerb) {
		var ret ExpectationVerb
		return ret
	}
	return *o.ExpectationVerb
}

// GetExpectationVerbOk returns a tuple with the ExpectationVerb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectation) GetExpectationVerbOk() (*ExpectationVerb, bool) {
	if o == nil || isNil(o.ExpectationVerb) {
    return nil, false
	}
	return o.ExpectationVerb, true
}

// HasExpectationVerb returns a boolean if a field has been set.
func (o *ServiceSupportExpectation) HasExpectationVerb() bool {
	if o != nil && !isNil(o.ExpectationVerb) {
		return true
	}

	return false
}

// SetExpectationVerb gets a reference to the given ExpectationVerb and assigns it to the ExpectationVerb field.
func (o *ServiceSupportExpectation) SetExpectationVerb(v ExpectationVerb) {
	o.ExpectationVerb = &v
}

// GetExpectationObjects returns the ExpectationObjects field value if set, zero value otherwise.
func (o *ServiceSupportExpectation) GetExpectationObjects() []ServiceSupportExpectationObject {
	if o == nil || isNil(o.ExpectationObjects) {
		var ret []ServiceSupportExpectationObject
		return ret
	}
	return o.ExpectationObjects
}

// GetExpectationObjectsOk returns a tuple with the ExpectationObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectation) GetExpectationObjectsOk() ([]ServiceSupportExpectationObject, bool) {
	if o == nil || isNil(o.ExpectationObjects) {
    return nil, false
	}
	return o.ExpectationObjects, true
}

// HasExpectationObjects returns a boolean if a field has been set.
func (o *ServiceSupportExpectation) HasExpectationObjects() bool {
	if o != nil && !isNil(o.ExpectationObjects) {
		return true
	}

	return false
}

// SetExpectationObjects gets a reference to the given []ServiceSupportExpectationObject and assigns it to the ExpectationObjects field.
func (o *ServiceSupportExpectation) SetExpectationObjects(v []ServiceSupportExpectationObject) {
	o.ExpectationObjects = v
}

// GetExpectationTargets returns the ExpectationTargets field value if set, zero value otherwise.
func (o *ServiceSupportExpectation) GetExpectationTargets() []ServiceSupportExpectationExpectationTargetsInner {
	if o == nil || isNil(o.ExpectationTargets) {
		var ret []ServiceSupportExpectationExpectationTargetsInner
		return ret
	}
	return o.ExpectationTargets
}

// GetExpectationTargetsOk returns a tuple with the ExpectationTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectation) GetExpectationTargetsOk() ([]ServiceSupportExpectationExpectationTargetsInner, bool) {
	if o == nil || isNil(o.ExpectationTargets) {
    return nil, false
	}
	return o.ExpectationTargets, true
}

// HasExpectationTargets returns a boolean if a field has been set.
func (o *ServiceSupportExpectation) HasExpectationTargets() bool {
	if o != nil && !isNil(o.ExpectationTargets) {
		return true
	}

	return false
}

// SetExpectationTargets gets a reference to the given []ServiceSupportExpectationExpectationTargetsInner and assigns it to the ExpectationTargets field.
func (o *ServiceSupportExpectation) SetExpectationTargets(v []ServiceSupportExpectationExpectationTargetsInner) {
	o.ExpectationTargets = v
}

// GetExpectationContexts returns the ExpectationContexts field value if set, zero value otherwise.
func (o *ServiceSupportExpectation) GetExpectationContexts() []ServiceSupportExpectationExpectationContextsInner {
	if o == nil || isNil(o.ExpectationContexts) {
		var ret []ServiceSupportExpectationExpectationContextsInner
		return ret
	}
	return o.ExpectationContexts
}

// GetExpectationContextsOk returns a tuple with the ExpectationContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectation) GetExpectationContextsOk() ([]ServiceSupportExpectationExpectationContextsInner, bool) {
	if o == nil || isNil(o.ExpectationContexts) {
    return nil, false
	}
	return o.ExpectationContexts, true
}

// HasExpectationContexts returns a boolean if a field has been set.
func (o *ServiceSupportExpectation) HasExpectationContexts() bool {
	if o != nil && !isNil(o.ExpectationContexts) {
		return true
	}

	return false
}

// SetExpectationContexts gets a reference to the given []ServiceSupportExpectationExpectationContextsInner and assigns it to the ExpectationContexts field.
func (o *ServiceSupportExpectation) SetExpectationContexts(v []ServiceSupportExpectationExpectationContextsInner) {
	o.ExpectationContexts = v
}

// GetExpectationfulfilmentInfo returns the ExpectationfulfilmentInfo field value if set, zero value otherwise.
func (o *ServiceSupportExpectation) GetExpectationfulfilmentInfo() FulfilmentInfo {
	if o == nil || isNil(o.ExpectationfulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.ExpectationfulfilmentInfo
}

// GetExpectationfulfilmentInfoOk returns a tuple with the ExpectationfulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSupportExpectation) GetExpectationfulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || isNil(o.ExpectationfulfilmentInfo) {
    return nil, false
	}
	return o.ExpectationfulfilmentInfo, true
}

// HasExpectationfulfilmentInfo returns a boolean if a field has been set.
func (o *ServiceSupportExpectation) HasExpectationfulfilmentInfo() bool {
	if o != nil && !isNil(o.ExpectationfulfilmentInfo) {
		return true
	}

	return false
}

// SetExpectationfulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the ExpectationfulfilmentInfo field.
func (o *ServiceSupportExpectation) SetExpectationfulfilmentInfo(v FulfilmentInfo) {
	o.ExpectationfulfilmentInfo = &v
}

func (o ServiceSupportExpectation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpectationId) {
		toSerialize["expectationId"] = o.ExpectationId
	}
	if !isNil(o.ExpectationVerb) {
		toSerialize["expectationVerb"] = o.ExpectationVerb
	}
	if !isNil(o.ExpectationObjects) {
		toSerialize["expectationObjects"] = o.ExpectationObjects
	}
	if !isNil(o.ExpectationTargets) {
		toSerialize["expectationTargets"] = o.ExpectationTargets
	}
	if !isNil(o.ExpectationContexts) {
		toSerialize["expectationContexts"] = o.ExpectationContexts
	}
	if !isNil(o.ExpectationfulfilmentInfo) {
		toSerialize["expectationfulfilmentInfo"] = o.ExpectationfulfilmentInfo
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSupportExpectation struct {
	value *ServiceSupportExpectation
	isSet bool
}

func (v NullableServiceSupportExpectation) Get() *ServiceSupportExpectation {
	return v.value
}

func (v *NullableServiceSupportExpectation) Set(val *ServiceSupportExpectation) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSupportExpectation) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSupportExpectation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSupportExpectation(val *ServiceSupportExpectation) *NullableServiceSupportExpectation {
	return &NullableServiceSupportExpectation{value: val, isSet: true}
}

func (v NullableServiceSupportExpectation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSupportExpectation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


