/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// CCOFunctionSingleAllOfAttributes struct for CCOFunctionSingleAllOfAttributes
type CCOFunctionSingleAllOfAttributes struct {
	CCOControl *bool `json:"cCOControl,omitempty"`
	CCOWeakCoverageParameters *CCOWeakCoverageParametersSingle `json:"cCOWeakCoverageParameters,omitempty"`
	CCOPilotPollutionParameters *CCOPilotPollutionParametersSingle `json:"cCOPilotPollutionParameters,omitempty"`
	CCOOvershootCoverageParametersSingle *CCOOvershootCoverageParametersSingle `json:"cCOOvershootCoverageParameters-Single,omitempty"`
}

// NewCCOFunctionSingleAllOfAttributes instantiates a new CCOFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCOFunctionSingleAllOfAttributes() *CCOFunctionSingleAllOfAttributes {
	this := CCOFunctionSingleAllOfAttributes{}
	return &this
}

// NewCCOFunctionSingleAllOfAttributesWithDefaults instantiates a new CCOFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCOFunctionSingleAllOfAttributesWithDefaults() *CCOFunctionSingleAllOfAttributes {
	this := CCOFunctionSingleAllOfAttributes{}
	return &this
}

// GetCCOControl returns the CCOControl field value if set, zero value otherwise.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOControl() bool {
	if o == nil || isNil(o.CCOControl) {
		var ret bool
		return ret
	}
	return *o.CCOControl
}

// GetCCOControlOk returns a tuple with the CCOControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOControlOk() (*bool, bool) {
	if o == nil || isNil(o.CCOControl) {
    return nil, false
	}
	return o.CCOControl, true
}

// HasCCOControl returns a boolean if a field has been set.
func (o *CCOFunctionSingleAllOfAttributes) HasCCOControl() bool {
	if o != nil && !isNil(o.CCOControl) {
		return true
	}

	return false
}

// SetCCOControl gets a reference to the given bool and assigns it to the CCOControl field.
func (o *CCOFunctionSingleAllOfAttributes) SetCCOControl(v bool) {
	o.CCOControl = &v
}

// GetCCOWeakCoverageParameters returns the CCOWeakCoverageParameters field value if set, zero value otherwise.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOWeakCoverageParameters() CCOWeakCoverageParametersSingle {
	if o == nil || isNil(o.CCOWeakCoverageParameters) {
		var ret CCOWeakCoverageParametersSingle
		return ret
	}
	return *o.CCOWeakCoverageParameters
}

// GetCCOWeakCoverageParametersOk returns a tuple with the CCOWeakCoverageParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOWeakCoverageParametersOk() (*CCOWeakCoverageParametersSingle, bool) {
	if o == nil || isNil(o.CCOWeakCoverageParameters) {
    return nil, false
	}
	return o.CCOWeakCoverageParameters, true
}

// HasCCOWeakCoverageParameters returns a boolean if a field has been set.
func (o *CCOFunctionSingleAllOfAttributes) HasCCOWeakCoverageParameters() bool {
	if o != nil && !isNil(o.CCOWeakCoverageParameters) {
		return true
	}

	return false
}

// SetCCOWeakCoverageParameters gets a reference to the given CCOWeakCoverageParametersSingle and assigns it to the CCOWeakCoverageParameters field.
func (o *CCOFunctionSingleAllOfAttributes) SetCCOWeakCoverageParameters(v CCOWeakCoverageParametersSingle) {
	o.CCOWeakCoverageParameters = &v
}

// GetCCOPilotPollutionParameters returns the CCOPilotPollutionParameters field value if set, zero value otherwise.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOPilotPollutionParameters() CCOPilotPollutionParametersSingle {
	if o == nil || isNil(o.CCOPilotPollutionParameters) {
		var ret CCOPilotPollutionParametersSingle
		return ret
	}
	return *o.CCOPilotPollutionParameters
}

// GetCCOPilotPollutionParametersOk returns a tuple with the CCOPilotPollutionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOPilotPollutionParametersOk() (*CCOPilotPollutionParametersSingle, bool) {
	if o == nil || isNil(o.CCOPilotPollutionParameters) {
    return nil, false
	}
	return o.CCOPilotPollutionParameters, true
}

// HasCCOPilotPollutionParameters returns a boolean if a field has been set.
func (o *CCOFunctionSingleAllOfAttributes) HasCCOPilotPollutionParameters() bool {
	if o != nil && !isNil(o.CCOPilotPollutionParameters) {
		return true
	}

	return false
}

// SetCCOPilotPollutionParameters gets a reference to the given CCOPilotPollutionParametersSingle and assigns it to the CCOPilotPollutionParameters field.
func (o *CCOFunctionSingleAllOfAttributes) SetCCOPilotPollutionParameters(v CCOPilotPollutionParametersSingle) {
	o.CCOPilotPollutionParameters = &v
}

// GetCCOOvershootCoverageParametersSingle returns the CCOOvershootCoverageParametersSingle field value if set, zero value otherwise.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOOvershootCoverageParametersSingle() CCOOvershootCoverageParametersSingle {
	if o == nil || isNil(o.CCOOvershootCoverageParametersSingle) {
		var ret CCOOvershootCoverageParametersSingle
		return ret
	}
	return *o.CCOOvershootCoverageParametersSingle
}

// GetCCOOvershootCoverageParametersSingleOk returns a tuple with the CCOOvershootCoverageParametersSingle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOFunctionSingleAllOfAttributes) GetCCOOvershootCoverageParametersSingleOk() (*CCOOvershootCoverageParametersSingle, bool) {
	if o == nil || isNil(o.CCOOvershootCoverageParametersSingle) {
    return nil, false
	}
	return o.CCOOvershootCoverageParametersSingle, true
}

// HasCCOOvershootCoverageParametersSingle returns a boolean if a field has been set.
func (o *CCOFunctionSingleAllOfAttributes) HasCCOOvershootCoverageParametersSingle() bool {
	if o != nil && !isNil(o.CCOOvershootCoverageParametersSingle) {
		return true
	}

	return false
}

// SetCCOOvershootCoverageParametersSingle gets a reference to the given CCOOvershootCoverageParametersSingle and assigns it to the CCOOvershootCoverageParametersSingle field.
func (o *CCOFunctionSingleAllOfAttributes) SetCCOOvershootCoverageParametersSingle(v CCOOvershootCoverageParametersSingle) {
	o.CCOOvershootCoverageParametersSingle = &v
}

func (o CCOFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CCOControl) {
		toSerialize["cCOControl"] = o.CCOControl
	}
	if !isNil(o.CCOWeakCoverageParameters) {
		toSerialize["cCOWeakCoverageParameters"] = o.CCOWeakCoverageParameters
	}
	if !isNil(o.CCOPilotPollutionParameters) {
		toSerialize["cCOPilotPollutionParameters"] = o.CCOPilotPollutionParameters
	}
	if !isNil(o.CCOOvershootCoverageParametersSingle) {
		toSerialize["cCOOvershootCoverageParameters-Single"] = o.CCOOvershootCoverageParametersSingle
	}
	return json.Marshal(toSerialize)
}

type NullableCCOFunctionSingleAllOfAttributes struct {
	value *CCOFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableCCOFunctionSingleAllOfAttributes) Get() *CCOFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableCCOFunctionSingleAllOfAttributes) Set(val *CCOFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCCOFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCCOFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCOFunctionSingleAllOfAttributes(val *CCOFunctionSingleAllOfAttributes) *NullableCCOFunctionSingleAllOfAttributes {
	return &NullableCCOFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableCCOFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCOFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


