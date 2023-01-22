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

// FiveQICharacteristicsSingleAllOf struct for FiveQICharacteristicsSingleAllOf
type FiveQICharacteristicsSingleAllOf struct {
	FiveQIValue *int32 `json:"fiveQIValue,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	PriorityLevel *int32 `json:"priorityLevel,omitempty"`
	PacketDelayBudget *int32 `json:"packetDelayBudget,omitempty"`
	PacketErrorRate *PacketErrorRate `json:"packetErrorRate,omitempty"`
	AveragingWindow *int32 `json:"averagingWindow,omitempty"`
	MaximumDataBurstVolume *int32 `json:"maximumDataBurstVolume,omitempty"`
}

// NewFiveQICharacteristicsSingleAllOf instantiates a new FiveQICharacteristicsSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiveQICharacteristicsSingleAllOf() *FiveQICharacteristicsSingleAllOf {
	this := FiveQICharacteristicsSingleAllOf{}
	return &this
}

// NewFiveQICharacteristicsSingleAllOfWithDefaults instantiates a new FiveQICharacteristicsSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiveQICharacteristicsSingleAllOfWithDefaults() *FiveQICharacteristicsSingleAllOf {
	this := FiveQICharacteristicsSingleAllOf{}
	return &this
}

// GetFiveQIValue returns the FiveQIValue field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingleAllOf) GetFiveQIValue() int32 {
	if o == nil || isNil(o.FiveQIValue) {
		var ret int32
		return ret
	}
	return *o.FiveQIValue
}

// GetFiveQIValueOk returns a tuple with the FiveQIValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingleAllOf) GetFiveQIValueOk() (*int32, bool) {
	if o == nil || isNil(o.FiveQIValue) {
    return nil, false
	}
	return o.FiveQIValue, true
}

// HasFiveQIValue returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingleAllOf) HasFiveQIValue() bool {
	if o != nil && !isNil(o.FiveQIValue) {
		return true
	}

	return false
}

// SetFiveQIValue gets a reference to the given int32 and assigns it to the FiveQIValue field.
func (o *FiveQICharacteristicsSingleAllOf) SetFiveQIValue(v int32) {
	o.FiveQIValue = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingleAllOf) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingleAllOf) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
    return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingleAllOf) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *FiveQICharacteristicsSingleAllOf) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetPriorityLevel returns the PriorityLevel field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingleAllOf) GetPriorityLevel() int32 {
	if o == nil || isNil(o.PriorityLevel) {
		var ret int32
		return ret
	}
	return *o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingleAllOf) GetPriorityLevelOk() (*int32, bool) {
	if o == nil || isNil(o.PriorityLevel) {
    return nil, false
	}
	return o.PriorityLevel, true
}

// HasPriorityLevel returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingleAllOf) HasPriorityLevel() bool {
	if o != nil && !isNil(o.PriorityLevel) {
		return true
	}

	return false
}

// SetPriorityLevel gets a reference to the given int32 and assigns it to the PriorityLevel field.
func (o *FiveQICharacteristicsSingleAllOf) SetPriorityLevel(v int32) {
	o.PriorityLevel = &v
}

// GetPacketDelayBudget returns the PacketDelayBudget field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingleAllOf) GetPacketDelayBudget() int32 {
	if o == nil || isNil(o.PacketDelayBudget) {
		var ret int32
		return ret
	}
	return *o.PacketDelayBudget
}

// GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingleAllOf) GetPacketDelayBudgetOk() (*int32, bool) {
	if o == nil || isNil(o.PacketDelayBudget) {
    return nil, false
	}
	return o.PacketDelayBudget, true
}

// HasPacketDelayBudget returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingleAllOf) HasPacketDelayBudget() bool {
	if o != nil && !isNil(o.PacketDelayBudget) {
		return true
	}

	return false
}

// SetPacketDelayBudget gets a reference to the given int32 and assigns it to the PacketDelayBudget field.
func (o *FiveQICharacteristicsSingleAllOf) SetPacketDelayBudget(v int32) {
	o.PacketDelayBudget = &v
}

// GetPacketErrorRate returns the PacketErrorRate field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingleAllOf) GetPacketErrorRate() PacketErrorRate {
	if o == nil || isNil(o.PacketErrorRate) {
		var ret PacketErrorRate
		return ret
	}
	return *o.PacketErrorRate
}

// GetPacketErrorRateOk returns a tuple with the PacketErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingleAllOf) GetPacketErrorRateOk() (*PacketErrorRate, bool) {
	if o == nil || isNil(o.PacketErrorRate) {
    return nil, false
	}
	return o.PacketErrorRate, true
}

// HasPacketErrorRate returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingleAllOf) HasPacketErrorRate() bool {
	if o != nil && !isNil(o.PacketErrorRate) {
		return true
	}

	return false
}

// SetPacketErrorRate gets a reference to the given PacketErrorRate and assigns it to the PacketErrorRate field.
func (o *FiveQICharacteristicsSingleAllOf) SetPacketErrorRate(v PacketErrorRate) {
	o.PacketErrorRate = &v
}

// GetAveragingWindow returns the AveragingWindow field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingleAllOf) GetAveragingWindow() int32 {
	if o == nil || isNil(o.AveragingWindow) {
		var ret int32
		return ret
	}
	return *o.AveragingWindow
}

// GetAveragingWindowOk returns a tuple with the AveragingWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingleAllOf) GetAveragingWindowOk() (*int32, bool) {
	if o == nil || isNil(o.AveragingWindow) {
    return nil, false
	}
	return o.AveragingWindow, true
}

// HasAveragingWindow returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingleAllOf) HasAveragingWindow() bool {
	if o != nil && !isNil(o.AveragingWindow) {
		return true
	}

	return false
}

// SetAveragingWindow gets a reference to the given int32 and assigns it to the AveragingWindow field.
func (o *FiveQICharacteristicsSingleAllOf) SetAveragingWindow(v int32) {
	o.AveragingWindow = &v
}

// GetMaximumDataBurstVolume returns the MaximumDataBurstVolume field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingleAllOf) GetMaximumDataBurstVolume() int32 {
	if o == nil || isNil(o.MaximumDataBurstVolume) {
		var ret int32
		return ret
	}
	return *o.MaximumDataBurstVolume
}

// GetMaximumDataBurstVolumeOk returns a tuple with the MaximumDataBurstVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingleAllOf) GetMaximumDataBurstVolumeOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumDataBurstVolume) {
    return nil, false
	}
	return o.MaximumDataBurstVolume, true
}

// HasMaximumDataBurstVolume returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingleAllOf) HasMaximumDataBurstVolume() bool {
	if o != nil && !isNil(o.MaximumDataBurstVolume) {
		return true
	}

	return false
}

// SetMaximumDataBurstVolume gets a reference to the given int32 and assigns it to the MaximumDataBurstVolume field.
func (o *FiveQICharacteristicsSingleAllOf) SetMaximumDataBurstVolume(v int32) {
	o.MaximumDataBurstVolume = &v
}

func (o FiveQICharacteristicsSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FiveQIValue) {
		toSerialize["fiveQIValue"] = o.FiveQIValue
	}
	if !isNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !isNil(o.PriorityLevel) {
		toSerialize["priorityLevel"] = o.PriorityLevel
	}
	if !isNil(o.PacketDelayBudget) {
		toSerialize["packetDelayBudget"] = o.PacketDelayBudget
	}
	if !isNil(o.PacketErrorRate) {
		toSerialize["packetErrorRate"] = o.PacketErrorRate
	}
	if !isNil(o.AveragingWindow) {
		toSerialize["averagingWindow"] = o.AveragingWindow
	}
	if !isNil(o.MaximumDataBurstVolume) {
		toSerialize["maximumDataBurstVolume"] = o.MaximumDataBurstVolume
	}
	return json.Marshal(toSerialize)
}

type NullableFiveQICharacteristicsSingleAllOf struct {
	value *FiveQICharacteristicsSingleAllOf
	isSet bool
}

func (v NullableFiveQICharacteristicsSingleAllOf) Get() *FiveQICharacteristicsSingleAllOf {
	return v.value
}

func (v *NullableFiveQICharacteristicsSingleAllOf) Set(val *FiveQICharacteristicsSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFiveQICharacteristicsSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFiveQICharacteristicsSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiveQICharacteristicsSingleAllOf(val *FiveQICharacteristicsSingleAllOf) *NullableFiveQICharacteristicsSingleAllOf {
	return &NullableFiveQICharacteristicsSingleAllOf{value: val, isSet: true}
}

func (v NullableFiveQICharacteristicsSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiveQICharacteristicsSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


