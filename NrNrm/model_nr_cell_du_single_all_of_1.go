/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// NrCellDuSingleAllOf1 struct for NrCellDuSingleAllOf1
type NrCellDuSingleAllOf1 struct {
	RRMPolicyRatio []RRMPolicyRatioSingle `json:"RRMPolicyRatio,omitempty"`
	CPCIConfigurationFunction *CPCIConfigurationFunctionSingle `json:"CPCIConfigurationFunction,omitempty"`
	DRACHOptimizationFunction *DRACHOptimizationFunctionSingle `json:"DRACHOptimizationFunction,omitempty"`
	NrOperatorCellDu []NrOperatorCellDuSingle `json:"NrOperatorCellDu,omitempty"`
}

// NewNrCellDuSingleAllOf1 instantiates a new NrCellDuSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrCellDuSingleAllOf1() *NrCellDuSingleAllOf1 {
	this := NrCellDuSingleAllOf1{}
	return &this
}

// NewNrCellDuSingleAllOf1WithDefaults instantiates a new NrCellDuSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrCellDuSingleAllOf1WithDefaults() *NrCellDuSingleAllOf1 {
	this := NrCellDuSingleAllOf1{}
	return &this
}

// GetRRMPolicyRatio returns the RRMPolicyRatio field value if set, zero value otherwise.
func (o *NrCellDuSingleAllOf1) GetRRMPolicyRatio() []RRMPolicyRatioSingle {
	if o == nil || isNil(o.RRMPolicyRatio) {
		var ret []RRMPolicyRatioSingle
		return ret
	}
	return o.RRMPolicyRatio
}

// GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellDuSingleAllOf1) GetRRMPolicyRatioOk() ([]RRMPolicyRatioSingle, bool) {
	if o == nil || isNil(o.RRMPolicyRatio) {
    return nil, false
	}
	return o.RRMPolicyRatio, true
}

// HasRRMPolicyRatio returns a boolean if a field has been set.
func (o *NrCellDuSingleAllOf1) HasRRMPolicyRatio() bool {
	if o != nil && !isNil(o.RRMPolicyRatio) {
		return true
	}

	return false
}

// SetRRMPolicyRatio gets a reference to the given []RRMPolicyRatioSingle and assigns it to the RRMPolicyRatio field.
func (o *NrCellDuSingleAllOf1) SetRRMPolicyRatio(v []RRMPolicyRatioSingle) {
	o.RRMPolicyRatio = v
}

// GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field value if set, zero value otherwise.
func (o *NrCellDuSingleAllOf1) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle {
	if o == nil || isNil(o.CPCIConfigurationFunction) {
		var ret CPCIConfigurationFunctionSingle
		return ret
	}
	return *o.CPCIConfigurationFunction
}

// GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellDuSingleAllOf1) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool) {
	if o == nil || isNil(o.CPCIConfigurationFunction) {
    return nil, false
	}
	return o.CPCIConfigurationFunction, true
}

// HasCPCIConfigurationFunction returns a boolean if a field has been set.
func (o *NrCellDuSingleAllOf1) HasCPCIConfigurationFunction() bool {
	if o != nil && !isNil(o.CPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetCPCIConfigurationFunction gets a reference to the given CPCIConfigurationFunctionSingle and assigns it to the CPCIConfigurationFunction field.
func (o *NrCellDuSingleAllOf1) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle) {
	o.CPCIConfigurationFunction = &v
}

// GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field value if set, zero value otherwise.
func (o *NrCellDuSingleAllOf1) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle {
	if o == nil || isNil(o.DRACHOptimizationFunction) {
		var ret DRACHOptimizationFunctionSingle
		return ret
	}
	return *o.DRACHOptimizationFunction
}

// GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellDuSingleAllOf1) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool) {
	if o == nil || isNil(o.DRACHOptimizationFunction) {
    return nil, false
	}
	return o.DRACHOptimizationFunction, true
}

// HasDRACHOptimizationFunction returns a boolean if a field has been set.
func (o *NrCellDuSingleAllOf1) HasDRACHOptimizationFunction() bool {
	if o != nil && !isNil(o.DRACHOptimizationFunction) {
		return true
	}

	return false
}

// SetDRACHOptimizationFunction gets a reference to the given DRACHOptimizationFunctionSingle and assigns it to the DRACHOptimizationFunction field.
func (o *NrCellDuSingleAllOf1) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle) {
	o.DRACHOptimizationFunction = &v
}

// GetNrOperatorCellDu returns the NrOperatorCellDu field value if set, zero value otherwise.
func (o *NrCellDuSingleAllOf1) GetNrOperatorCellDu() []NrOperatorCellDuSingle {
	if o == nil || isNil(o.NrOperatorCellDu) {
		var ret []NrOperatorCellDuSingle
		return ret
	}
	return o.NrOperatorCellDu
}

// GetNrOperatorCellDuOk returns a tuple with the NrOperatorCellDu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellDuSingleAllOf1) GetNrOperatorCellDuOk() ([]NrOperatorCellDuSingle, bool) {
	if o == nil || isNil(o.NrOperatorCellDu) {
    return nil, false
	}
	return o.NrOperatorCellDu, true
}

// HasNrOperatorCellDu returns a boolean if a field has been set.
func (o *NrCellDuSingleAllOf1) HasNrOperatorCellDu() bool {
	if o != nil && !isNil(o.NrOperatorCellDu) {
		return true
	}

	return false
}

// SetNrOperatorCellDu gets a reference to the given []NrOperatorCellDuSingle and assigns it to the NrOperatorCellDu field.
func (o *NrCellDuSingleAllOf1) SetNrOperatorCellDu(v []NrOperatorCellDuSingle) {
	o.NrOperatorCellDu = v
}

func (o NrCellDuSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RRMPolicyRatio) {
		toSerialize["RRMPolicyRatio"] = o.RRMPolicyRatio
	}
	if !isNil(o.CPCIConfigurationFunction) {
		toSerialize["CPCIConfigurationFunction"] = o.CPCIConfigurationFunction
	}
	if !isNil(o.DRACHOptimizationFunction) {
		toSerialize["DRACHOptimizationFunction"] = o.DRACHOptimizationFunction
	}
	if !isNil(o.NrOperatorCellDu) {
		toSerialize["NrOperatorCellDu"] = o.NrOperatorCellDu
	}
	return json.Marshal(toSerialize)
}

type NullableNrCellDuSingleAllOf1 struct {
	value *NrCellDuSingleAllOf1
	isSet bool
}

func (v NullableNrCellDuSingleAllOf1) Get() *NrCellDuSingleAllOf1 {
	return v.value
}

func (v *NullableNrCellDuSingleAllOf1) Set(val *NrCellDuSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableNrCellDuSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableNrCellDuSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrCellDuSingleAllOf1(val *NrCellDuSingleAllOf1) *NullableNrCellDuSingleAllOf1 {
	return &NullableNrCellDuSingleAllOf1{value: val, isSet: true}
}

func (v NullableNrCellDuSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrCellDuSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


