/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M1_ContentHostingProvisioning

import (
	"encoding/json"
)

// CachingConfigurationCachingDirectives struct for CachingConfigurationCachingDirectives
type CachingConfigurationCachingDirectives struct {
	StatusCodeFilters []int32 `json:"statusCodeFilters,omitempty"`
	NoCache bool `json:"noCache"`
	MaxAge *int32 `json:"maxAge,omitempty"`
}

// NewCachingConfigurationCachingDirectives instantiates a new CachingConfigurationCachingDirectives object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCachingConfigurationCachingDirectives(noCache bool) *CachingConfigurationCachingDirectives {
	this := CachingConfigurationCachingDirectives{}
	this.NoCache = noCache
	return &this
}

// NewCachingConfigurationCachingDirectivesWithDefaults instantiates a new CachingConfigurationCachingDirectives object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCachingConfigurationCachingDirectivesWithDefaults() *CachingConfigurationCachingDirectives {
	this := CachingConfigurationCachingDirectives{}
	return &this
}

// GetStatusCodeFilters returns the StatusCodeFilters field value if set, zero value otherwise.
func (o *CachingConfigurationCachingDirectives) GetStatusCodeFilters() []int32 {
	if o == nil || isNil(o.StatusCodeFilters) {
		var ret []int32
		return ret
	}
	return o.StatusCodeFilters
}

// GetStatusCodeFiltersOk returns a tuple with the StatusCodeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CachingConfigurationCachingDirectives) GetStatusCodeFiltersOk() ([]int32, bool) {
	if o == nil || isNil(o.StatusCodeFilters) {
    return nil, false
	}
	return o.StatusCodeFilters, true
}

// HasStatusCodeFilters returns a boolean if a field has been set.
func (o *CachingConfigurationCachingDirectives) HasStatusCodeFilters() bool {
	if o != nil && !isNil(o.StatusCodeFilters) {
		return true
	}

	return false
}

// SetStatusCodeFilters gets a reference to the given []int32 and assigns it to the StatusCodeFilters field.
func (o *CachingConfigurationCachingDirectives) SetStatusCodeFilters(v []int32) {
	o.StatusCodeFilters = v
}

// GetNoCache returns the NoCache field value
func (o *CachingConfigurationCachingDirectives) GetNoCache() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NoCache
}

// GetNoCacheOk returns a tuple with the NoCache field value
// and a boolean to check if the value has been set.
func (o *CachingConfigurationCachingDirectives) GetNoCacheOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NoCache, true
}

// SetNoCache sets field value
func (o *CachingConfigurationCachingDirectives) SetNoCache(v bool) {
	o.NoCache = v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *CachingConfigurationCachingDirectives) GetMaxAge() int32 {
	if o == nil || isNil(o.MaxAge) {
		var ret int32
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CachingConfigurationCachingDirectives) GetMaxAgeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxAge) {
    return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *CachingConfigurationCachingDirectives) HasMaxAge() bool {
	if o != nil && !isNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int32 and assigns it to the MaxAge field.
func (o *CachingConfigurationCachingDirectives) SetMaxAge(v int32) {
	o.MaxAge = &v
}

func (o CachingConfigurationCachingDirectives) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StatusCodeFilters) {
		toSerialize["statusCodeFilters"] = o.StatusCodeFilters
	}
	if true {
		toSerialize["noCache"] = o.NoCache
	}
	if !isNil(o.MaxAge) {
		toSerialize["maxAge"] = o.MaxAge
	}
	return json.Marshal(toSerialize)
}

type NullableCachingConfigurationCachingDirectives struct {
	value *CachingConfigurationCachingDirectives
	isSet bool
}

func (v NullableCachingConfigurationCachingDirectives) Get() *CachingConfigurationCachingDirectives {
	return v.value
}

func (v *NullableCachingConfigurationCachingDirectives) Set(val *CachingConfigurationCachingDirectives) {
	v.value = val
	v.isSet = true
}

func (v NullableCachingConfigurationCachingDirectives) IsSet() bool {
	return v.isSet
}

func (v *NullableCachingConfigurationCachingDirectives) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCachingConfigurationCachingDirectives(val *CachingConfigurationCachingDirectives) *NullableCachingConfigurationCachingDirectives {
	return &NullableCachingConfigurationCachingDirectives{value: val, isSet: true}
}

func (v NullableCachingConfigurationCachingDirectives) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCachingConfigurationCachingDirectives) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


