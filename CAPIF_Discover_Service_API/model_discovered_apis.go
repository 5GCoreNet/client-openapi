/*
CAPIF_Discover_Service_API

API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Discover_Service_API

import (
	"encoding/json"
)

// DiscoveredAPIs Represents a list of APIs currently registered in the CAPIF core function and satisfying a number of filter criteria provided by the API consumer. 
type DiscoveredAPIs struct {
	// Description of the service API as published by the service. Each service API description shall include AEF profiles matching the filter criteria. 
	ServiceAPIDescriptions []ServiceAPIDescription `json:"serviceAPIDescriptions,omitempty"`
}

// NewDiscoveredAPIs instantiates a new DiscoveredAPIs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredAPIs() *DiscoveredAPIs {
	this := DiscoveredAPIs{}
	return &this
}

// NewDiscoveredAPIsWithDefaults instantiates a new DiscoveredAPIs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredAPIsWithDefaults() *DiscoveredAPIs {
	this := DiscoveredAPIs{}
	return &this
}

// GetServiceAPIDescriptions returns the ServiceAPIDescriptions field value if set, zero value otherwise.
func (o *DiscoveredAPIs) GetServiceAPIDescriptions() []ServiceAPIDescription {
	if o == nil || isNil(o.ServiceAPIDescriptions) {
		var ret []ServiceAPIDescription
		return ret
	}
	return o.ServiceAPIDescriptions
}

// GetServiceAPIDescriptionsOk returns a tuple with the ServiceAPIDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredAPIs) GetServiceAPIDescriptionsOk() ([]ServiceAPIDescription, bool) {
	if o == nil || isNil(o.ServiceAPIDescriptions) {
    return nil, false
	}
	return o.ServiceAPIDescriptions, true
}

// HasServiceAPIDescriptions returns a boolean if a field has been set.
func (o *DiscoveredAPIs) HasServiceAPIDescriptions() bool {
	if o != nil && !isNil(o.ServiceAPIDescriptions) {
		return true
	}

	return false
}

// SetServiceAPIDescriptions gets a reference to the given []ServiceAPIDescription and assigns it to the ServiceAPIDescriptions field.
func (o *DiscoveredAPIs) SetServiceAPIDescriptions(v []ServiceAPIDescription) {
	o.ServiceAPIDescriptions = v
}

func (o DiscoveredAPIs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceAPIDescriptions) {
		toSerialize["serviceAPIDescriptions"] = o.ServiceAPIDescriptions
	}
	return json.Marshal(toSerialize)
}

type NullableDiscoveredAPIs struct {
	value *DiscoveredAPIs
	isSet bool
}

func (v NullableDiscoveredAPIs) Get() *DiscoveredAPIs {
	return v.value
}

func (v *NullableDiscoveredAPIs) Set(val *DiscoveredAPIs) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredAPIs) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredAPIs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredAPIs(val *DiscoveredAPIs) *NullableDiscoveredAPIs {
	return &NullableDiscoveredAPIs{value: val, isSet: true}
}

func (v NullableDiscoveredAPIs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredAPIs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


