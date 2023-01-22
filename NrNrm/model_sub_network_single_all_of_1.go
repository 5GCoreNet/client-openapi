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

// SubNetworkSingleAllOf1 struct for SubNetworkSingleAllOf1
type SubNetworkSingleAllOf1 struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`
	NRFrequency []NRFrequencySingle `json:"NRFrequency,omitempty"`
	ExternalGnbCuCpFunction []ExternalGnbCuCpFunctionSingle `json:"ExternalGnbCuCpFunction,omitempty"`
	ExternalENBFunction []ExternalENBFunctionSingle `json:"ExternalENBFunction,omitempty"`
	EUtranFrequency []EUtranFrequencySingle `json:"EUtranFrequency,omitempty"`
	DESManagementFunction *DESManagementFunctionSingle `json:"DESManagementFunction,omitempty"`
	DRACHOptimizationFunction *DRACHOptimizationFunctionSingle `json:"DRACHOptimizationFunction,omitempty"`
	DMROFunction *DMROFunctionSingle `json:"DMROFunction,omitempty"`
	DLBOFunction *DLBOFunctionSingle `json:"DLBOFunction,omitempty"`
	DPCIConfigurationFunction *DPCIConfigurationFunctionSingle `json:"DPCIConfigurationFunction,omitempty"`
	CPCIConfigurationFunction *CPCIConfigurationFunctionSingle `json:"CPCIConfigurationFunction,omitempty"`
	CESManagementFunction *CESManagementFunctionSingle `json:"CESManagementFunction,omitempty"`
	Configurable5QISet []Configurable5QISetSingle `json:"Configurable5QISet,omitempty"`
	RimRSGlobal *RimRSGlobalSingle `json:"RimRSGlobal,omitempty"`
	Dynamic5QISet []Dynamic5QISetSingle `json:"Dynamic5QISet,omitempty"`
	CCOFunction *CCOFunctionSingle `json:"CCOFunction,omitempty"`
}

// NewSubNetworkSingleAllOf1 instantiates a new SubNetworkSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingleAllOf1() *SubNetworkSingleAllOf1 {
	this := SubNetworkSingleAllOf1{}
	return &this
}

// NewSubNetworkSingleAllOf1WithDefaults instantiates a new SubNetworkSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleAllOf1WithDefaults() *SubNetworkSingleAllOf1 {
	this := SubNetworkSingleAllOf1{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetSubNetwork() []SubNetworkSingle {
	if o == nil || isNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.SubNetwork) {
    return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasSubNetwork() bool {
	if o != nil && !isNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingleAllOf1) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetManagedElement() []ManagedElementSingle {
	if o == nil || isNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || isNil(o.ManagedElement) {
    return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasManagedElement() bool {
	if o != nil && !isNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *SubNetworkSingleAllOf1) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

// GetNRFrequency returns the NRFrequency field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetNRFrequency() []NRFrequencySingle {
	if o == nil || isNil(o.NRFrequency) {
		var ret []NRFrequencySingle
		return ret
	}
	return o.NRFrequency
}

// GetNRFrequencyOk returns a tuple with the NRFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetNRFrequencyOk() ([]NRFrequencySingle, bool) {
	if o == nil || isNil(o.NRFrequency) {
    return nil, false
	}
	return o.NRFrequency, true
}

// HasNRFrequency returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasNRFrequency() bool {
	if o != nil && !isNil(o.NRFrequency) {
		return true
	}

	return false
}

// SetNRFrequency gets a reference to the given []NRFrequencySingle and assigns it to the NRFrequency field.
func (o *SubNetworkSingleAllOf1) SetNRFrequency(v []NRFrequencySingle) {
	o.NRFrequency = v
}

// GetExternalGnbCuCpFunction returns the ExternalGnbCuCpFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetExternalGnbCuCpFunction() []ExternalGnbCuCpFunctionSingle {
	if o == nil || isNil(o.ExternalGnbCuCpFunction) {
		var ret []ExternalGnbCuCpFunctionSingle
		return ret
	}
	return o.ExternalGnbCuCpFunction
}

// GetExternalGnbCuCpFunctionOk returns a tuple with the ExternalGnbCuCpFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetExternalGnbCuCpFunctionOk() ([]ExternalGnbCuCpFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalGnbCuCpFunction) {
    return nil, false
	}
	return o.ExternalGnbCuCpFunction, true
}

// HasExternalGnbCuCpFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasExternalGnbCuCpFunction() bool {
	if o != nil && !isNil(o.ExternalGnbCuCpFunction) {
		return true
	}

	return false
}

// SetExternalGnbCuCpFunction gets a reference to the given []ExternalGnbCuCpFunctionSingle and assigns it to the ExternalGnbCuCpFunction field.
func (o *SubNetworkSingleAllOf1) SetExternalGnbCuCpFunction(v []ExternalGnbCuCpFunctionSingle) {
	o.ExternalGnbCuCpFunction = v
}

// GetExternalENBFunction returns the ExternalENBFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetExternalENBFunction() []ExternalENBFunctionSingle {
	if o == nil || isNil(o.ExternalENBFunction) {
		var ret []ExternalENBFunctionSingle
		return ret
	}
	return o.ExternalENBFunction
}

// GetExternalENBFunctionOk returns a tuple with the ExternalENBFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetExternalENBFunctionOk() ([]ExternalENBFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalENBFunction) {
    return nil, false
	}
	return o.ExternalENBFunction, true
}

// HasExternalENBFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasExternalENBFunction() bool {
	if o != nil && !isNil(o.ExternalENBFunction) {
		return true
	}

	return false
}

// SetExternalENBFunction gets a reference to the given []ExternalENBFunctionSingle and assigns it to the ExternalENBFunction field.
func (o *SubNetworkSingleAllOf1) SetExternalENBFunction(v []ExternalENBFunctionSingle) {
	o.ExternalENBFunction = v
}

// GetEUtranFrequency returns the EUtranFrequency field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetEUtranFrequency() []EUtranFrequencySingle {
	if o == nil || isNil(o.EUtranFrequency) {
		var ret []EUtranFrequencySingle
		return ret
	}
	return o.EUtranFrequency
}

// GetEUtranFrequencyOk returns a tuple with the EUtranFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetEUtranFrequencyOk() ([]EUtranFrequencySingle, bool) {
	if o == nil || isNil(o.EUtranFrequency) {
    return nil, false
	}
	return o.EUtranFrequency, true
}

// HasEUtranFrequency returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasEUtranFrequency() bool {
	if o != nil && !isNil(o.EUtranFrequency) {
		return true
	}

	return false
}

// SetEUtranFrequency gets a reference to the given []EUtranFrequencySingle and assigns it to the EUtranFrequency field.
func (o *SubNetworkSingleAllOf1) SetEUtranFrequency(v []EUtranFrequencySingle) {
	o.EUtranFrequency = v
}

// GetDESManagementFunction returns the DESManagementFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetDESManagementFunction() DESManagementFunctionSingle {
	if o == nil || isNil(o.DESManagementFunction) {
		var ret DESManagementFunctionSingle
		return ret
	}
	return *o.DESManagementFunction
}

// GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool) {
	if o == nil || isNil(o.DESManagementFunction) {
    return nil, false
	}
	return o.DESManagementFunction, true
}

// HasDESManagementFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasDESManagementFunction() bool {
	if o != nil && !isNil(o.DESManagementFunction) {
		return true
	}

	return false
}

// SetDESManagementFunction gets a reference to the given DESManagementFunctionSingle and assigns it to the DESManagementFunction field.
func (o *SubNetworkSingleAllOf1) SetDESManagementFunction(v DESManagementFunctionSingle) {
	o.DESManagementFunction = &v
}

// GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle {
	if o == nil || isNil(o.DRACHOptimizationFunction) {
		var ret DRACHOptimizationFunctionSingle
		return ret
	}
	return *o.DRACHOptimizationFunction
}

// GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool) {
	if o == nil || isNil(o.DRACHOptimizationFunction) {
    return nil, false
	}
	return o.DRACHOptimizationFunction, true
}

// HasDRACHOptimizationFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasDRACHOptimizationFunction() bool {
	if o != nil && !isNil(o.DRACHOptimizationFunction) {
		return true
	}

	return false
}

// SetDRACHOptimizationFunction gets a reference to the given DRACHOptimizationFunctionSingle and assigns it to the DRACHOptimizationFunction field.
func (o *SubNetworkSingleAllOf1) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle) {
	o.DRACHOptimizationFunction = &v
}

// GetDMROFunction returns the DMROFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetDMROFunction() DMROFunctionSingle {
	if o == nil || isNil(o.DMROFunction) {
		var ret DMROFunctionSingle
		return ret
	}
	return *o.DMROFunction
}

// GetDMROFunctionOk returns a tuple with the DMROFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetDMROFunctionOk() (*DMROFunctionSingle, bool) {
	if o == nil || isNil(o.DMROFunction) {
    return nil, false
	}
	return o.DMROFunction, true
}

// HasDMROFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasDMROFunction() bool {
	if o != nil && !isNil(o.DMROFunction) {
		return true
	}

	return false
}

// SetDMROFunction gets a reference to the given DMROFunctionSingle and assigns it to the DMROFunction field.
func (o *SubNetworkSingleAllOf1) SetDMROFunction(v DMROFunctionSingle) {
	o.DMROFunction = &v
}

// GetDLBOFunction returns the DLBOFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetDLBOFunction() DLBOFunctionSingle {
	if o == nil || isNil(o.DLBOFunction) {
		var ret DLBOFunctionSingle
		return ret
	}
	return *o.DLBOFunction
}

// GetDLBOFunctionOk returns a tuple with the DLBOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool) {
	if o == nil || isNil(o.DLBOFunction) {
    return nil, false
	}
	return o.DLBOFunction, true
}

// HasDLBOFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasDLBOFunction() bool {
	if o != nil && !isNil(o.DLBOFunction) {
		return true
	}

	return false
}

// SetDLBOFunction gets a reference to the given DLBOFunctionSingle and assigns it to the DLBOFunction field.
func (o *SubNetworkSingleAllOf1) SetDLBOFunction(v DLBOFunctionSingle) {
	o.DLBOFunction = &v
}

// GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle {
	if o == nil || isNil(o.DPCIConfigurationFunction) {
		var ret DPCIConfigurationFunctionSingle
		return ret
	}
	return *o.DPCIConfigurationFunction
}

// GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool) {
	if o == nil || isNil(o.DPCIConfigurationFunction) {
    return nil, false
	}
	return o.DPCIConfigurationFunction, true
}

// HasDPCIConfigurationFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasDPCIConfigurationFunction() bool {
	if o != nil && !isNil(o.DPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetDPCIConfigurationFunction gets a reference to the given DPCIConfigurationFunctionSingle and assigns it to the DPCIConfigurationFunction field.
func (o *SubNetworkSingleAllOf1) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle) {
	o.DPCIConfigurationFunction = &v
}

// GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle {
	if o == nil || isNil(o.CPCIConfigurationFunction) {
		var ret CPCIConfigurationFunctionSingle
		return ret
	}
	return *o.CPCIConfigurationFunction
}

// GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool) {
	if o == nil || isNil(o.CPCIConfigurationFunction) {
    return nil, false
	}
	return o.CPCIConfigurationFunction, true
}

// HasCPCIConfigurationFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasCPCIConfigurationFunction() bool {
	if o != nil && !isNil(o.CPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetCPCIConfigurationFunction gets a reference to the given CPCIConfigurationFunctionSingle and assigns it to the CPCIConfigurationFunction field.
func (o *SubNetworkSingleAllOf1) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle) {
	o.CPCIConfigurationFunction = &v
}

// GetCESManagementFunction returns the CESManagementFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetCESManagementFunction() CESManagementFunctionSingle {
	if o == nil || isNil(o.CESManagementFunction) {
		var ret CESManagementFunctionSingle
		return ret
	}
	return *o.CESManagementFunction
}

// GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool) {
	if o == nil || isNil(o.CESManagementFunction) {
    return nil, false
	}
	return o.CESManagementFunction, true
}

// HasCESManagementFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasCESManagementFunction() bool {
	if o != nil && !isNil(o.CESManagementFunction) {
		return true
	}

	return false
}

// SetCESManagementFunction gets a reference to the given CESManagementFunctionSingle and assigns it to the CESManagementFunction field.
func (o *SubNetworkSingleAllOf1) SetCESManagementFunction(v CESManagementFunctionSingle) {
	o.CESManagementFunction = &v
}

// GetConfigurable5QISet returns the Configurable5QISet field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetConfigurable5QISet() []Configurable5QISetSingle {
	if o == nil || isNil(o.Configurable5QISet) {
		var ret []Configurable5QISetSingle
		return ret
	}
	return o.Configurable5QISet
}

// GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetConfigurable5QISetOk() ([]Configurable5QISetSingle, bool) {
	if o == nil || isNil(o.Configurable5QISet) {
    return nil, false
	}
	return o.Configurable5QISet, true
}

// HasConfigurable5QISet returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasConfigurable5QISet() bool {
	if o != nil && !isNil(o.Configurable5QISet) {
		return true
	}

	return false
}

// SetConfigurable5QISet gets a reference to the given []Configurable5QISetSingle and assigns it to the Configurable5QISet field.
func (o *SubNetworkSingleAllOf1) SetConfigurable5QISet(v []Configurable5QISetSingle) {
	o.Configurable5QISet = v
}

// GetRimRSGlobal returns the RimRSGlobal field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetRimRSGlobal() RimRSGlobalSingle {
	if o == nil || isNil(o.RimRSGlobal) {
		var ret RimRSGlobalSingle
		return ret
	}
	return *o.RimRSGlobal
}

// GetRimRSGlobalOk returns a tuple with the RimRSGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetRimRSGlobalOk() (*RimRSGlobalSingle, bool) {
	if o == nil || isNil(o.RimRSGlobal) {
    return nil, false
	}
	return o.RimRSGlobal, true
}

// HasRimRSGlobal returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasRimRSGlobal() bool {
	if o != nil && !isNil(o.RimRSGlobal) {
		return true
	}

	return false
}

// SetRimRSGlobal gets a reference to the given RimRSGlobalSingle and assigns it to the RimRSGlobal field.
func (o *SubNetworkSingleAllOf1) SetRimRSGlobal(v RimRSGlobalSingle) {
	o.RimRSGlobal = &v
}

// GetDynamic5QISet returns the Dynamic5QISet field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetDynamic5QISet() []Dynamic5QISetSingle {
	if o == nil || isNil(o.Dynamic5QISet) {
		var ret []Dynamic5QISetSingle
		return ret
	}
	return o.Dynamic5QISet
}

// GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetDynamic5QISetOk() ([]Dynamic5QISetSingle, bool) {
	if o == nil || isNil(o.Dynamic5QISet) {
    return nil, false
	}
	return o.Dynamic5QISet, true
}

// HasDynamic5QISet returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasDynamic5QISet() bool {
	if o != nil && !isNil(o.Dynamic5QISet) {
		return true
	}

	return false
}

// SetDynamic5QISet gets a reference to the given []Dynamic5QISetSingle and assigns it to the Dynamic5QISet field.
func (o *SubNetworkSingleAllOf1) SetDynamic5QISet(v []Dynamic5QISetSingle) {
	o.Dynamic5QISet = v
}

// GetCCOFunction returns the CCOFunction field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf1) GetCCOFunction() CCOFunctionSingle {
	if o == nil || isNil(o.CCOFunction) {
		var ret CCOFunctionSingle
		return ret
	}
	return *o.CCOFunction
}

// GetCCOFunctionOk returns a tuple with the CCOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf1) GetCCOFunctionOk() (*CCOFunctionSingle, bool) {
	if o == nil || isNil(o.CCOFunction) {
    return nil, false
	}
	return o.CCOFunction, true
}

// HasCCOFunction returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf1) HasCCOFunction() bool {
	if o != nil && !isNil(o.CCOFunction) {
		return true
	}

	return false
}

// SetCCOFunction gets a reference to the given CCOFunctionSingle and assigns it to the CCOFunction field.
func (o *SubNetworkSingleAllOf1) SetCCOFunction(v CCOFunctionSingle) {
	o.CCOFunction = &v
}

func (o SubNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !isNil(o.ManagedElement) {
		toSerialize["ManagedElement"] = o.ManagedElement
	}
	if !isNil(o.NRFrequency) {
		toSerialize["NRFrequency"] = o.NRFrequency
	}
	if !isNil(o.ExternalGnbCuCpFunction) {
		toSerialize["ExternalGnbCuCpFunction"] = o.ExternalGnbCuCpFunction
	}
	if !isNil(o.ExternalENBFunction) {
		toSerialize["ExternalENBFunction"] = o.ExternalENBFunction
	}
	if !isNil(o.EUtranFrequency) {
		toSerialize["EUtranFrequency"] = o.EUtranFrequency
	}
	if !isNil(o.DESManagementFunction) {
		toSerialize["DESManagementFunction"] = o.DESManagementFunction
	}
	if !isNil(o.DRACHOptimizationFunction) {
		toSerialize["DRACHOptimizationFunction"] = o.DRACHOptimizationFunction
	}
	if !isNil(o.DMROFunction) {
		toSerialize["DMROFunction"] = o.DMROFunction
	}
	if !isNil(o.DLBOFunction) {
		toSerialize["DLBOFunction"] = o.DLBOFunction
	}
	if !isNil(o.DPCIConfigurationFunction) {
		toSerialize["DPCIConfigurationFunction"] = o.DPCIConfigurationFunction
	}
	if !isNil(o.CPCIConfigurationFunction) {
		toSerialize["CPCIConfigurationFunction"] = o.CPCIConfigurationFunction
	}
	if !isNil(o.CESManagementFunction) {
		toSerialize["CESManagementFunction"] = o.CESManagementFunction
	}
	if !isNil(o.Configurable5QISet) {
		toSerialize["Configurable5QISet"] = o.Configurable5QISet
	}
	if !isNil(o.RimRSGlobal) {
		toSerialize["RimRSGlobal"] = o.RimRSGlobal
	}
	if !isNil(o.Dynamic5QISet) {
		toSerialize["Dynamic5QISet"] = o.Dynamic5QISet
	}
	if !isNil(o.CCOFunction) {
		toSerialize["CCOFunction"] = o.CCOFunction
	}
	return json.Marshal(toSerialize)
}

type NullableSubNetworkSingleAllOf1 struct {
	value *SubNetworkSingleAllOf1
	isSet bool
}

func (v NullableSubNetworkSingleAllOf1) Get() *SubNetworkSingleAllOf1 {
	return v.value
}

func (v *NullableSubNetworkSingleAllOf1) Set(val *SubNetworkSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingleAllOf1(val *SubNetworkSingleAllOf1) *NullableSubNetworkSingleAllOf1 {
	return &NullableSubNetworkSingleAllOf1{value: val, isSet: true}
}

func (v NullableSubNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


