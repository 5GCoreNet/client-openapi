/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// FiveQICharacteristicsSingle struct for FiveQICharacteristicsSingle
type FiveQICharacteristicsSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	FiveQIValue *int32 `json:"fiveQIValue,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	PriorityLevel *int32 `json:"priorityLevel,omitempty"`
	PacketDelayBudget *int32 `json:"packetDelayBudget,omitempty"`
	PacketErrorRate *PacketErrorRate `json:"packetErrorRate,omitempty"`
	AveragingWindow *int32 `json:"averagingWindow,omitempty"`
	MaximumDataBurstVolume *int32 `json:"maximumDataBurstVolume,omitempty"`
}

// NewFiveQICharacteristicsSingle instantiates a new FiveQICharacteristicsSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiveQICharacteristicsSingle(id NullableString) *FiveQICharacteristicsSingle {
	this := FiveQICharacteristicsSingle{}
	this.Id = id
	return &this
}

// NewFiveQICharacteristicsSingleWithDefaults instantiates a new FiveQICharacteristicsSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiveQICharacteristicsSingleWithDefaults() *FiveQICharacteristicsSingle {
	this := FiveQICharacteristicsSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FiveQICharacteristicsSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FiveQICharacteristicsSingle) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *FiveQICharacteristicsSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *FiveQICharacteristicsSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *FiveQICharacteristicsSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *FiveQICharacteristicsSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetFiveQIValue returns the FiveQIValue field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetFiveQIValue() int32 {
	if o == nil || isNil(o.FiveQIValue) {
		var ret int32
		return ret
	}
	return *o.FiveQIValue
}

// GetFiveQIValueOk returns a tuple with the FiveQIValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetFiveQIValueOk() (*int32, bool) {
	if o == nil || isNil(o.FiveQIValue) {
    return nil, false
	}
	return o.FiveQIValue, true
}

// HasFiveQIValue returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasFiveQIValue() bool {
	if o != nil && !isNil(o.FiveQIValue) {
		return true
	}

	return false
}

// SetFiveQIValue gets a reference to the given int32 and assigns it to the FiveQIValue field.
func (o *FiveQICharacteristicsSingle) SetFiveQIValue(v int32) {
	o.FiveQIValue = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetResourceType() string {
	if o == nil || isNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetResourceTypeOk() (*string, bool) {
	if o == nil || isNil(o.ResourceType) {
    return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *FiveQICharacteristicsSingle) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetPriorityLevel returns the PriorityLevel field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetPriorityLevel() int32 {
	if o == nil || isNil(o.PriorityLevel) {
		var ret int32
		return ret
	}
	return *o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetPriorityLevelOk() (*int32, bool) {
	if o == nil || isNil(o.PriorityLevel) {
    return nil, false
	}
	return o.PriorityLevel, true
}

// HasPriorityLevel returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasPriorityLevel() bool {
	if o != nil && !isNil(o.PriorityLevel) {
		return true
	}

	return false
}

// SetPriorityLevel gets a reference to the given int32 and assigns it to the PriorityLevel field.
func (o *FiveQICharacteristicsSingle) SetPriorityLevel(v int32) {
	o.PriorityLevel = &v
}

// GetPacketDelayBudget returns the PacketDelayBudget field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetPacketDelayBudget() int32 {
	if o == nil || isNil(o.PacketDelayBudget) {
		var ret int32
		return ret
	}
	return *o.PacketDelayBudget
}

// GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetPacketDelayBudgetOk() (*int32, bool) {
	if o == nil || isNil(o.PacketDelayBudget) {
    return nil, false
	}
	return o.PacketDelayBudget, true
}

// HasPacketDelayBudget returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasPacketDelayBudget() bool {
	if o != nil && !isNil(o.PacketDelayBudget) {
		return true
	}

	return false
}

// SetPacketDelayBudget gets a reference to the given int32 and assigns it to the PacketDelayBudget field.
func (o *FiveQICharacteristicsSingle) SetPacketDelayBudget(v int32) {
	o.PacketDelayBudget = &v
}

// GetPacketErrorRate returns the PacketErrorRate field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetPacketErrorRate() PacketErrorRate {
	if o == nil || isNil(o.PacketErrorRate) {
		var ret PacketErrorRate
		return ret
	}
	return *o.PacketErrorRate
}

// GetPacketErrorRateOk returns a tuple with the PacketErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetPacketErrorRateOk() (*PacketErrorRate, bool) {
	if o == nil || isNil(o.PacketErrorRate) {
    return nil, false
	}
	return o.PacketErrorRate, true
}

// HasPacketErrorRate returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasPacketErrorRate() bool {
	if o != nil && !isNil(o.PacketErrorRate) {
		return true
	}

	return false
}

// SetPacketErrorRate gets a reference to the given PacketErrorRate and assigns it to the PacketErrorRate field.
func (o *FiveQICharacteristicsSingle) SetPacketErrorRate(v PacketErrorRate) {
	o.PacketErrorRate = &v
}

// GetAveragingWindow returns the AveragingWindow field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetAveragingWindow() int32 {
	if o == nil || isNil(o.AveragingWindow) {
		var ret int32
		return ret
	}
	return *o.AveragingWindow
}

// GetAveragingWindowOk returns a tuple with the AveragingWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetAveragingWindowOk() (*int32, bool) {
	if o == nil || isNil(o.AveragingWindow) {
    return nil, false
	}
	return o.AveragingWindow, true
}

// HasAveragingWindow returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasAveragingWindow() bool {
	if o != nil && !isNil(o.AveragingWindow) {
		return true
	}

	return false
}

// SetAveragingWindow gets a reference to the given int32 and assigns it to the AveragingWindow field.
func (o *FiveQICharacteristicsSingle) SetAveragingWindow(v int32) {
	o.AveragingWindow = &v
}

// GetMaximumDataBurstVolume returns the MaximumDataBurstVolume field value if set, zero value otherwise.
func (o *FiveQICharacteristicsSingle) GetMaximumDataBurstVolume() int32 {
	if o == nil || isNil(o.MaximumDataBurstVolume) {
		var ret int32
		return ret
	}
	return *o.MaximumDataBurstVolume
}

// GetMaximumDataBurstVolumeOk returns a tuple with the MaximumDataBurstVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQICharacteristicsSingle) GetMaximumDataBurstVolumeOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumDataBurstVolume) {
    return nil, false
	}
	return o.MaximumDataBurstVolume, true
}

// HasMaximumDataBurstVolume returns a boolean if a field has been set.
func (o *FiveQICharacteristicsSingle) HasMaximumDataBurstVolume() bool {
	if o != nil && !isNil(o.MaximumDataBurstVolume) {
		return true
	}

	return false
}

// SetMaximumDataBurstVolume gets a reference to the given int32 and assigns it to the MaximumDataBurstVolume field.
func (o *FiveQICharacteristicsSingle) SetMaximumDataBurstVolume(v int32) {
	o.MaximumDataBurstVolume = &v
}

func (o FiveQICharacteristicsSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if !isNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !isNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !isNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
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

type NullableFiveQICharacteristicsSingle struct {
	value *FiveQICharacteristicsSingle
	isSet bool
}

func (v NullableFiveQICharacteristicsSingle) Get() *FiveQICharacteristicsSingle {
	return v.value
}

func (v *NullableFiveQICharacteristicsSingle) Set(val *FiveQICharacteristicsSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableFiveQICharacteristicsSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableFiveQICharacteristicsSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiveQICharacteristicsSingle(val *FiveQICharacteristicsSingle) *NullableFiveQICharacteristicsSingle {
	return &NullableFiveQICharacteristicsSingle{value: val, isSet: true}
}

func (v NullableFiveQICharacteristicsSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiveQICharacteristicsSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


