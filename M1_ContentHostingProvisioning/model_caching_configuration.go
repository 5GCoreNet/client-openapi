/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentHostingProvisioning

import (
	"encoding/json"
)

// CachingConfiguration A content caching configuration.
type CachingConfiguration struct {
	UrlPatternFilter string `json:"urlPatternFilter"`
	CachingDirectives *CachingConfigurationCachingDirectives `json:"cachingDirectives,omitempty"`
}

// NewCachingConfiguration instantiates a new CachingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCachingConfiguration(urlPatternFilter string) *CachingConfiguration {
	this := CachingConfiguration{}
	this.UrlPatternFilter = urlPatternFilter
	return &this
}

// NewCachingConfigurationWithDefaults instantiates a new CachingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCachingConfigurationWithDefaults() *CachingConfiguration {
	this := CachingConfiguration{}
	return &this
}

// GetUrlPatternFilter returns the UrlPatternFilter field value
func (o *CachingConfiguration) GetUrlPatternFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlPatternFilter
}

// GetUrlPatternFilterOk returns a tuple with the UrlPatternFilter field value
// and a boolean to check if the value has been set.
func (o *CachingConfiguration) GetUrlPatternFilterOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UrlPatternFilter, true
}

// SetUrlPatternFilter sets field value
func (o *CachingConfiguration) SetUrlPatternFilter(v string) {
	o.UrlPatternFilter = v
}

// GetCachingDirectives returns the CachingDirectives field value if set, zero value otherwise.
func (o *CachingConfiguration) GetCachingDirectives() CachingConfigurationCachingDirectives {
	if o == nil || isNil(o.CachingDirectives) {
		var ret CachingConfigurationCachingDirectives
		return ret
	}
	return *o.CachingDirectives
}

// GetCachingDirectivesOk returns a tuple with the CachingDirectives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CachingConfiguration) GetCachingDirectivesOk() (*CachingConfigurationCachingDirectives, bool) {
	if o == nil || isNil(o.CachingDirectives) {
    return nil, false
	}
	return o.CachingDirectives, true
}

// HasCachingDirectives returns a boolean if a field has been set.
func (o *CachingConfiguration) HasCachingDirectives() bool {
	if o != nil && !isNil(o.CachingDirectives) {
		return true
	}

	return false
}

// SetCachingDirectives gets a reference to the given CachingConfigurationCachingDirectives and assigns it to the CachingDirectives field.
func (o *CachingConfiguration) SetCachingDirectives(v CachingConfigurationCachingDirectives) {
	o.CachingDirectives = &v
}

func (o CachingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["urlPatternFilter"] = o.UrlPatternFilter
	}
	if !isNil(o.CachingDirectives) {
		toSerialize["cachingDirectives"] = o.CachingDirectives
	}
	return json.Marshal(toSerialize)
}

type NullableCachingConfiguration struct {
	value *CachingConfiguration
	isSet bool
}

func (v NullableCachingConfiguration) Get() *CachingConfiguration {
	return v.value
}

func (v *NullableCachingConfiguration) Set(val *CachingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCachingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCachingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCachingConfiguration(val *CachingConfiguration) *NullableCachingConfiguration {
	return &NullableCachingConfiguration{value: val, isSet: true}
}

func (v NullableCachingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCachingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


