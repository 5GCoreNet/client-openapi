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

// Configurable5QISetSingle struct for Configurable5QISetSingle
type Configurable5QISetSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewConfigurable5QISetSingle instantiates a new Configurable5QISetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurable5QISetSingle(id NullableString) *Configurable5QISetSingle {
	this := Configurable5QISetSingle{}
	this.Id = id
	return &this
}

// NewConfigurable5QISetSingleWithDefaults instantiates a new Configurable5QISetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurable5QISetSingleWithDefaults() *Configurable5QISetSingle {
	this := Configurable5QISetSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Configurable5QISetSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Configurable5QISetSingle) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *Configurable5QISetSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *Configurable5QISetSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configurable5QISetSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *Configurable5QISetSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *Configurable5QISetSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *Configurable5QISetSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configurable5QISetSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *Configurable5QISetSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *Configurable5QISetSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *Configurable5QISetSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configurable5QISetSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *Configurable5QISetSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *Configurable5QISetSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Configurable5QISetSingle) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configurable5QISetSingle) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Configurable5QISetSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *Configurable5QISetSingle) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o Configurable5QISetSingle) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurable5QISetSingle struct {
	value *Configurable5QISetSingle
	isSet bool
}

func (v NullableConfigurable5QISetSingle) Get() *Configurable5QISetSingle {
	return v.value
}

func (v *NullableConfigurable5QISetSingle) Set(val *Configurable5QISetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurable5QISetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurable5QISetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurable5QISetSingle(val *Configurable5QISetSingle) *NullableConfigurable5QISetSingle {
	return &NullableConfigurable5QISetSingle{value: val, isSet: true}
}

func (v NullableConfigurable5QISetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurable5QISetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


