/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// SubNetworkSingle1AllOf1 struct for SubNetworkSingle1AllOf1
type SubNetworkSingle1AllOf1 struct {
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`
	ExternalAmfFunction []ExternalAmfFunctionSingle `json:"ExternalAmfFunction,omitempty"`
	ExternalNrfFunction []ExternalNrfFunctionSingle `json:"ExternalNrfFunction,omitempty"`
	ExternalNssfFunction []ExternalNssfFunctionSingle `json:"ExternalNssfFunction,omitempty"`
	AmfSet []AmfSetSingle `json:"AmfSet,omitempty"`
	AmfRegion []AmfRegionSingle `json:"AmfRegion,omitempty"`
	Configurable5QISet []Configurable5QISetSingle `json:"Configurable5QISet,omitempty"`
	Dynamic5QISet []Dynamic5QISetSingle `json:"Dynamic5QISet,omitempty"`
	EcmConnectionInfo []EcmConnectionInfoSingle `json:"EcmConnectionInfo,omitempty"`
}

// NewSubNetworkSingle1AllOf1 instantiates a new SubNetworkSingle1AllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle1AllOf1() *SubNetworkSingle1AllOf1 {
	this := SubNetworkSingle1AllOf1{}
	return &this
}

// NewSubNetworkSingle1AllOf1WithDefaults instantiates a new SubNetworkSingle1AllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingle1AllOf1WithDefaults() *SubNetworkSingle1AllOf1 {
	this := SubNetworkSingle1AllOf1{}
	return &this
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetSubNetwork() []SubNetworkSingle {
	if o == nil || isNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.SubNetwork) {
    return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasSubNetwork() bool {
	if o != nil && !isNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingle1AllOf1) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetManagedElement() []ManagedElementSingle {
	if o == nil || isNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || isNil(o.ManagedElement) {
    return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasManagedElement() bool {
	if o != nil && !isNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *SubNetworkSingle1AllOf1) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

// GetExternalAmfFunction returns the ExternalAmfFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetExternalAmfFunction() []ExternalAmfFunctionSingle {
	if o == nil || isNil(o.ExternalAmfFunction) {
		var ret []ExternalAmfFunctionSingle
		return ret
	}
	return o.ExternalAmfFunction
}

// GetExternalAmfFunctionOk returns a tuple with the ExternalAmfFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetExternalAmfFunctionOk() ([]ExternalAmfFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalAmfFunction) {
    return nil, false
	}
	return o.ExternalAmfFunction, true
}

// HasExternalAmfFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasExternalAmfFunction() bool {
	if o != nil && !isNil(o.ExternalAmfFunction) {
		return true
	}

	return false
}

// SetExternalAmfFunction gets a reference to the given []ExternalAmfFunctionSingle and assigns it to the ExternalAmfFunction field.
func (o *SubNetworkSingle1AllOf1) SetExternalAmfFunction(v []ExternalAmfFunctionSingle) {
	o.ExternalAmfFunction = v
}

// GetExternalNrfFunction returns the ExternalNrfFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetExternalNrfFunction() []ExternalNrfFunctionSingle {
	if o == nil || isNil(o.ExternalNrfFunction) {
		var ret []ExternalNrfFunctionSingle
		return ret
	}
	return o.ExternalNrfFunction
}

// GetExternalNrfFunctionOk returns a tuple with the ExternalNrfFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetExternalNrfFunctionOk() ([]ExternalNrfFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalNrfFunction) {
    return nil, false
	}
	return o.ExternalNrfFunction, true
}

// HasExternalNrfFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasExternalNrfFunction() bool {
	if o != nil && !isNil(o.ExternalNrfFunction) {
		return true
	}

	return false
}

// SetExternalNrfFunction gets a reference to the given []ExternalNrfFunctionSingle and assigns it to the ExternalNrfFunction field.
func (o *SubNetworkSingle1AllOf1) SetExternalNrfFunction(v []ExternalNrfFunctionSingle) {
	o.ExternalNrfFunction = v
}

// GetExternalNssfFunction returns the ExternalNssfFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetExternalNssfFunction() []ExternalNssfFunctionSingle {
	if o == nil || isNil(o.ExternalNssfFunction) {
		var ret []ExternalNssfFunctionSingle
		return ret
	}
	return o.ExternalNssfFunction
}

// GetExternalNssfFunctionOk returns a tuple with the ExternalNssfFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetExternalNssfFunctionOk() ([]ExternalNssfFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalNssfFunction) {
    return nil, false
	}
	return o.ExternalNssfFunction, true
}

// HasExternalNssfFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasExternalNssfFunction() bool {
	if o != nil && !isNil(o.ExternalNssfFunction) {
		return true
	}

	return false
}

// SetExternalNssfFunction gets a reference to the given []ExternalNssfFunctionSingle and assigns it to the ExternalNssfFunction field.
func (o *SubNetworkSingle1AllOf1) SetExternalNssfFunction(v []ExternalNssfFunctionSingle) {
	o.ExternalNssfFunction = v
}

// GetAmfSet returns the AmfSet field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetAmfSet() []AmfSetSingle {
	if o == nil || isNil(o.AmfSet) {
		var ret []AmfSetSingle
		return ret
	}
	return o.AmfSet
}

// GetAmfSetOk returns a tuple with the AmfSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetAmfSetOk() ([]AmfSetSingle, bool) {
	if o == nil || isNil(o.AmfSet) {
    return nil, false
	}
	return o.AmfSet, true
}

// HasAmfSet returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasAmfSet() bool {
	if o != nil && !isNil(o.AmfSet) {
		return true
	}

	return false
}

// SetAmfSet gets a reference to the given []AmfSetSingle and assigns it to the AmfSet field.
func (o *SubNetworkSingle1AllOf1) SetAmfSet(v []AmfSetSingle) {
	o.AmfSet = v
}

// GetAmfRegion returns the AmfRegion field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetAmfRegion() []AmfRegionSingle {
	if o == nil || isNil(o.AmfRegion) {
		var ret []AmfRegionSingle
		return ret
	}
	return o.AmfRegion
}

// GetAmfRegionOk returns a tuple with the AmfRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetAmfRegionOk() ([]AmfRegionSingle, bool) {
	if o == nil || isNil(o.AmfRegion) {
    return nil, false
	}
	return o.AmfRegion, true
}

// HasAmfRegion returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasAmfRegion() bool {
	if o != nil && !isNil(o.AmfRegion) {
		return true
	}

	return false
}

// SetAmfRegion gets a reference to the given []AmfRegionSingle and assigns it to the AmfRegion field.
func (o *SubNetworkSingle1AllOf1) SetAmfRegion(v []AmfRegionSingle) {
	o.AmfRegion = v
}

// GetConfigurable5QISet returns the Configurable5QISet field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetConfigurable5QISet() []Configurable5QISetSingle {
	if o == nil || isNil(o.Configurable5QISet) {
		var ret []Configurable5QISetSingle
		return ret
	}
	return o.Configurable5QISet
}

// GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetConfigurable5QISetOk() ([]Configurable5QISetSingle, bool) {
	if o == nil || isNil(o.Configurable5QISet) {
    return nil, false
	}
	return o.Configurable5QISet, true
}

// HasConfigurable5QISet returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasConfigurable5QISet() bool {
	if o != nil && !isNil(o.Configurable5QISet) {
		return true
	}

	return false
}

// SetConfigurable5QISet gets a reference to the given []Configurable5QISetSingle and assigns it to the Configurable5QISet field.
func (o *SubNetworkSingle1AllOf1) SetConfigurable5QISet(v []Configurable5QISetSingle) {
	o.Configurable5QISet = v
}

// GetDynamic5QISet returns the Dynamic5QISet field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetDynamic5QISet() []Dynamic5QISetSingle {
	if o == nil || isNil(o.Dynamic5QISet) {
		var ret []Dynamic5QISetSingle
		return ret
	}
	return o.Dynamic5QISet
}

// GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetDynamic5QISetOk() ([]Dynamic5QISetSingle, bool) {
	if o == nil || isNil(o.Dynamic5QISet) {
    return nil, false
	}
	return o.Dynamic5QISet, true
}

// HasDynamic5QISet returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasDynamic5QISet() bool {
	if o != nil && !isNil(o.Dynamic5QISet) {
		return true
	}

	return false
}

// SetDynamic5QISet gets a reference to the given []Dynamic5QISetSingle and assigns it to the Dynamic5QISet field.
func (o *SubNetworkSingle1AllOf1) SetDynamic5QISet(v []Dynamic5QISetSingle) {
	o.Dynamic5QISet = v
}

// GetEcmConnectionInfo returns the EcmConnectionInfo field value if set, zero value otherwise.
func (o *SubNetworkSingle1AllOf1) GetEcmConnectionInfo() []EcmConnectionInfoSingle {
	if o == nil || isNil(o.EcmConnectionInfo) {
		var ret []EcmConnectionInfoSingle
		return ret
	}
	return o.EcmConnectionInfo
}

// GetEcmConnectionInfoOk returns a tuple with the EcmConnectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1AllOf1) GetEcmConnectionInfoOk() ([]EcmConnectionInfoSingle, bool) {
	if o == nil || isNil(o.EcmConnectionInfo) {
    return nil, false
	}
	return o.EcmConnectionInfo, true
}

// HasEcmConnectionInfo returns a boolean if a field has been set.
func (o *SubNetworkSingle1AllOf1) HasEcmConnectionInfo() bool {
	if o != nil && !isNil(o.EcmConnectionInfo) {
		return true
	}

	return false
}

// SetEcmConnectionInfo gets a reference to the given []EcmConnectionInfoSingle and assigns it to the EcmConnectionInfo field.
func (o *SubNetworkSingle1AllOf1) SetEcmConnectionInfo(v []EcmConnectionInfoSingle) {
	o.EcmConnectionInfo = v
}

func (o SubNetworkSingle1AllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !isNil(o.ManagedElement) {
		toSerialize["ManagedElement"] = o.ManagedElement
	}
	if !isNil(o.ExternalAmfFunction) {
		toSerialize["ExternalAmfFunction"] = o.ExternalAmfFunction
	}
	if !isNil(o.ExternalNrfFunction) {
		toSerialize["ExternalNrfFunction"] = o.ExternalNrfFunction
	}
	if !isNil(o.ExternalNssfFunction) {
		toSerialize["ExternalNssfFunction"] = o.ExternalNssfFunction
	}
	if !isNil(o.AmfSet) {
		toSerialize["AmfSet"] = o.AmfSet
	}
	if !isNil(o.AmfRegion) {
		toSerialize["AmfRegion"] = o.AmfRegion
	}
	if !isNil(o.Configurable5QISet) {
		toSerialize["Configurable5QISet"] = o.Configurable5QISet
	}
	if !isNil(o.Dynamic5QISet) {
		toSerialize["Dynamic5QISet"] = o.Dynamic5QISet
	}
	if !isNil(o.EcmConnectionInfo) {
		toSerialize["EcmConnectionInfo"] = o.EcmConnectionInfo
	}
	return json.Marshal(toSerialize)
}

type NullableSubNetworkSingle1AllOf1 struct {
	value *SubNetworkSingle1AllOf1
	isSet bool
}

func (v NullableSubNetworkSingle1AllOf1) Get() *SubNetworkSingle1AllOf1 {
	return v.value
}

func (v *NullableSubNetworkSingle1AllOf1) Set(val *SubNetworkSingle1AllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle1AllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle1AllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle1AllOf1(val *SubNetworkSingle1AllOf1) *NullableSubNetworkSingle1AllOf1 {
	return &NullableSubNetworkSingle1AllOf1{value: val, isSet: true}
}

func (v NullableSubNetworkSingle1AllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle1AllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


