/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// EASRequirements struct for EASRequirements
type EASRequirements struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	RequiredEASservingLocation *ServingLocation `json:"requiredEASservingLocation,omitempty"`
	AffinityAntiAffinity *AffinityAntiAffinity `json:"affinityAntiAffinity,omitempty"`
	ServiceContinuity *bool `json:"serviceContinuity,omitempty"`
	VirtualResource *VirtualResource `json:"virtualResource,omitempty"`
	SoftwareImageInfo *SoftwareImageInfo `json:"softwareImageInfo,omitempty"`
}

// NewEASRequirements instantiates a new EASRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASRequirements(id NullableString) *EASRequirements {
	this := EASRequirements{}
	this.Id = id
	return &this
}

// NewEASRequirementsWithDefaults instantiates a new EASRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASRequirementsWithDefaults() *EASRequirements {
	this := EASRequirements{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EASRequirements) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EASRequirements) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *EASRequirements) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *EASRequirements) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *EASRequirements) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *EASRequirements) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *EASRequirements) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *EASRequirements) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *EASRequirements) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *EASRequirements) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *EASRequirements) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *EASRequirements) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetRequiredEASservingLocation returns the RequiredEASservingLocation field value if set, zero value otherwise.
func (o *EASRequirements) GetRequiredEASservingLocation() ServingLocation {
	if o == nil || isNil(o.RequiredEASservingLocation) {
		var ret ServingLocation
		return ret
	}
	return *o.RequiredEASservingLocation
}

// GetRequiredEASservingLocationOk returns a tuple with the RequiredEASservingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetRequiredEASservingLocationOk() (*ServingLocation, bool) {
	if o == nil || isNil(o.RequiredEASservingLocation) {
    return nil, false
	}
	return o.RequiredEASservingLocation, true
}

// HasRequiredEASservingLocation returns a boolean if a field has been set.
func (o *EASRequirements) HasRequiredEASservingLocation() bool {
	if o != nil && !isNil(o.RequiredEASservingLocation) {
		return true
	}

	return false
}

// SetRequiredEASservingLocation gets a reference to the given ServingLocation and assigns it to the RequiredEASservingLocation field.
func (o *EASRequirements) SetRequiredEASservingLocation(v ServingLocation) {
	o.RequiredEASservingLocation = &v
}

// GetAffinityAntiAffinity returns the AffinityAntiAffinity field value if set, zero value otherwise.
func (o *EASRequirements) GetAffinityAntiAffinity() AffinityAntiAffinity {
	if o == nil || isNil(o.AffinityAntiAffinity) {
		var ret AffinityAntiAffinity
		return ret
	}
	return *o.AffinityAntiAffinity
}

// GetAffinityAntiAffinityOk returns a tuple with the AffinityAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetAffinityAntiAffinityOk() (*AffinityAntiAffinity, bool) {
	if o == nil || isNil(o.AffinityAntiAffinity) {
    return nil, false
	}
	return o.AffinityAntiAffinity, true
}

// HasAffinityAntiAffinity returns a boolean if a field has been set.
func (o *EASRequirements) HasAffinityAntiAffinity() bool {
	if o != nil && !isNil(o.AffinityAntiAffinity) {
		return true
	}

	return false
}

// SetAffinityAntiAffinity gets a reference to the given AffinityAntiAffinity and assigns it to the AffinityAntiAffinity field.
func (o *EASRequirements) SetAffinityAntiAffinity(v AffinityAntiAffinity) {
	o.AffinityAntiAffinity = &v
}

// GetServiceContinuity returns the ServiceContinuity field value if set, zero value otherwise.
func (o *EASRequirements) GetServiceContinuity() bool {
	if o == nil || isNil(o.ServiceContinuity) {
		var ret bool
		return ret
	}
	return *o.ServiceContinuity
}

// GetServiceContinuityOk returns a tuple with the ServiceContinuity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetServiceContinuityOk() (*bool, bool) {
	if o == nil || isNil(o.ServiceContinuity) {
    return nil, false
	}
	return o.ServiceContinuity, true
}

// HasServiceContinuity returns a boolean if a field has been set.
func (o *EASRequirements) HasServiceContinuity() bool {
	if o != nil && !isNil(o.ServiceContinuity) {
		return true
	}

	return false
}

// SetServiceContinuity gets a reference to the given bool and assigns it to the ServiceContinuity field.
func (o *EASRequirements) SetServiceContinuity(v bool) {
	o.ServiceContinuity = &v
}

// GetVirtualResource returns the VirtualResource field value if set, zero value otherwise.
func (o *EASRequirements) GetVirtualResource() VirtualResource {
	if o == nil || isNil(o.VirtualResource) {
		var ret VirtualResource
		return ret
	}
	return *o.VirtualResource
}

// GetVirtualResourceOk returns a tuple with the VirtualResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetVirtualResourceOk() (*VirtualResource, bool) {
	if o == nil || isNil(o.VirtualResource) {
    return nil, false
	}
	return o.VirtualResource, true
}

// HasVirtualResource returns a boolean if a field has been set.
func (o *EASRequirements) HasVirtualResource() bool {
	if o != nil && !isNil(o.VirtualResource) {
		return true
	}

	return false
}

// SetVirtualResource gets a reference to the given VirtualResource and assigns it to the VirtualResource field.
func (o *EASRequirements) SetVirtualResource(v VirtualResource) {
	o.VirtualResource = &v
}

// GetSoftwareImageInfo returns the SoftwareImageInfo field value if set, zero value otherwise.
func (o *EASRequirements) GetSoftwareImageInfo() SoftwareImageInfo {
	if o == nil || isNil(o.SoftwareImageInfo) {
		var ret SoftwareImageInfo
		return ret
	}
	return *o.SoftwareImageInfo
}

// GetSoftwareImageInfoOk returns a tuple with the SoftwareImageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirements) GetSoftwareImageInfoOk() (*SoftwareImageInfo, bool) {
	if o == nil || isNil(o.SoftwareImageInfo) {
    return nil, false
	}
	return o.SoftwareImageInfo, true
}

// HasSoftwareImageInfo returns a boolean if a field has been set.
func (o *EASRequirements) HasSoftwareImageInfo() bool {
	if o != nil && !isNil(o.SoftwareImageInfo) {
		return true
	}

	return false
}

// SetSoftwareImageInfo gets a reference to the given SoftwareImageInfo and assigns it to the SoftwareImageInfo field.
func (o *EASRequirements) SetSoftwareImageInfo(v SoftwareImageInfo) {
	o.SoftwareImageInfo = &v
}

func (o EASRequirements) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.RequiredEASservingLocation) {
		toSerialize["requiredEASservingLocation"] = o.RequiredEASservingLocation
	}
	if !isNil(o.AffinityAntiAffinity) {
		toSerialize["affinityAntiAffinity"] = o.AffinityAntiAffinity
	}
	if !isNil(o.ServiceContinuity) {
		toSerialize["serviceContinuity"] = o.ServiceContinuity
	}
	if !isNil(o.VirtualResource) {
		toSerialize["virtualResource"] = o.VirtualResource
	}
	if !isNil(o.SoftwareImageInfo) {
		toSerialize["softwareImageInfo"] = o.SoftwareImageInfo
	}
	return json.Marshal(toSerialize)
}

type NullableEASRequirements struct {
	value *EASRequirements
	isSet bool
}

func (v NullableEASRequirements) Get() *EASRequirements {
	return v.value
}

func (v *NullableEASRequirements) Set(val *EASRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableEASRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableEASRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASRequirements(val *EASRequirements) *NullableEASRequirements {
	return &NullableEASRequirements{value: val, isSet: true}
}

func (v NullableEASRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


