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

// DistributionConfiguration A content distribution configuration.
type DistributionConfiguration struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	ContentPreparationTemplateId *string `json:"contentPreparationTemplateId,omitempty"`
	CanonicalDomainName string `json:"canonicalDomainName"`
	DomainNameAlias string `json:"domainNameAlias"`
	PathRewriteRules []PathRewriteRule `json:"pathRewriteRules,omitempty"`
	CachingConfigurations []CachingConfiguration `json:"cachingConfigurations,omitempty"`
	GeoFencing *DistributionConfigurationGeoFencing `json:"geoFencing,omitempty"`
	UrlSignature *DistributionConfigurationUrlSignature `json:"urlSignature,omitempty"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	CertificateId *string `json:"certificateId,omitempty"`
	SupplementaryDistributionNetworks *DistributionConfigurationSupplementaryDistributionNetworks `json:"supplementaryDistributionNetworks,omitempty"`
}

// NewDistributionConfiguration instantiates a new DistributionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionConfiguration(canonicalDomainName string, domainNameAlias string) *DistributionConfiguration {
	this := DistributionConfiguration{}
	this.CanonicalDomainName = canonicalDomainName
	this.DomainNameAlias = domainNameAlias
	return &this
}

// NewDistributionConfigurationWithDefaults instantiates a new DistributionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionConfigurationWithDefaults() *DistributionConfiguration {
	this := DistributionConfiguration{}
	return &this
}

// GetContentPreparationTemplateId returns the ContentPreparationTemplateId field value if set, zero value otherwise.
func (o *DistributionConfiguration) GetContentPreparationTemplateId() string {
	if o == nil || isNil(o.ContentPreparationTemplateId) {
		var ret string
		return ret
	}
	return *o.ContentPreparationTemplateId
}

// GetContentPreparationTemplateIdOk returns a tuple with the ContentPreparationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetContentPreparationTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.ContentPreparationTemplateId) {
    return nil, false
	}
	return o.ContentPreparationTemplateId, true
}

// HasContentPreparationTemplateId returns a boolean if a field has been set.
func (o *DistributionConfiguration) HasContentPreparationTemplateId() bool {
	if o != nil && !isNil(o.ContentPreparationTemplateId) {
		return true
	}

	return false
}

// SetContentPreparationTemplateId gets a reference to the given string and assigns it to the ContentPreparationTemplateId field.
func (o *DistributionConfiguration) SetContentPreparationTemplateId(v string) {
	o.ContentPreparationTemplateId = &v
}

// GetCanonicalDomainName returns the CanonicalDomainName field value
func (o *DistributionConfiguration) GetCanonicalDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CanonicalDomainName
}

// GetCanonicalDomainNameOk returns a tuple with the CanonicalDomainName field value
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetCanonicalDomainNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CanonicalDomainName, true
}

// SetCanonicalDomainName sets field value
func (o *DistributionConfiguration) SetCanonicalDomainName(v string) {
	o.CanonicalDomainName = v
}

// GetDomainNameAlias returns the DomainNameAlias field value
func (o *DistributionConfiguration) GetDomainNameAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainNameAlias
}

// GetDomainNameAliasOk returns a tuple with the DomainNameAlias field value
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetDomainNameAliasOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DomainNameAlias, true
}

// SetDomainNameAlias sets field value
func (o *DistributionConfiguration) SetDomainNameAlias(v string) {
	o.DomainNameAlias = v
}

// GetPathRewriteRules returns the PathRewriteRules field value if set, zero value otherwise.
func (o *DistributionConfiguration) GetPathRewriteRules() []PathRewriteRule {
	if o == nil || isNil(o.PathRewriteRules) {
		var ret []PathRewriteRule
		return ret
	}
	return o.PathRewriteRules
}

// GetPathRewriteRulesOk returns a tuple with the PathRewriteRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetPathRewriteRulesOk() ([]PathRewriteRule, bool) {
	if o == nil || isNil(o.PathRewriteRules) {
    return nil, false
	}
	return o.PathRewriteRules, true
}

// HasPathRewriteRules returns a boolean if a field has been set.
func (o *DistributionConfiguration) HasPathRewriteRules() bool {
	if o != nil && !isNil(o.PathRewriteRules) {
		return true
	}

	return false
}

// SetPathRewriteRules gets a reference to the given []PathRewriteRule and assigns it to the PathRewriteRules field.
func (o *DistributionConfiguration) SetPathRewriteRules(v []PathRewriteRule) {
	o.PathRewriteRules = v
}

// GetCachingConfigurations returns the CachingConfigurations field value if set, zero value otherwise.
func (o *DistributionConfiguration) GetCachingConfigurations() []CachingConfiguration {
	if o == nil || isNil(o.CachingConfigurations) {
		var ret []CachingConfiguration
		return ret
	}
	return o.CachingConfigurations
}

// GetCachingConfigurationsOk returns a tuple with the CachingConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetCachingConfigurationsOk() ([]CachingConfiguration, bool) {
	if o == nil || isNil(o.CachingConfigurations) {
    return nil, false
	}
	return o.CachingConfigurations, true
}

// HasCachingConfigurations returns a boolean if a field has been set.
func (o *DistributionConfiguration) HasCachingConfigurations() bool {
	if o != nil && !isNil(o.CachingConfigurations) {
		return true
	}

	return false
}

// SetCachingConfigurations gets a reference to the given []CachingConfiguration and assigns it to the CachingConfigurations field.
func (o *DistributionConfiguration) SetCachingConfigurations(v []CachingConfiguration) {
	o.CachingConfigurations = v
}

// GetGeoFencing returns the GeoFencing field value if set, zero value otherwise.
func (o *DistributionConfiguration) GetGeoFencing() DistributionConfigurationGeoFencing {
	if o == nil || isNil(o.GeoFencing) {
		var ret DistributionConfigurationGeoFencing
		return ret
	}
	return *o.GeoFencing
}

// GetGeoFencingOk returns a tuple with the GeoFencing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetGeoFencingOk() (*DistributionConfigurationGeoFencing, bool) {
	if o == nil || isNil(o.GeoFencing) {
    return nil, false
	}
	return o.GeoFencing, true
}

// HasGeoFencing returns a boolean if a field has been set.
func (o *DistributionConfiguration) HasGeoFencing() bool {
	if o != nil && !isNil(o.GeoFencing) {
		return true
	}

	return false
}

// SetGeoFencing gets a reference to the given DistributionConfigurationGeoFencing and assigns it to the GeoFencing field.
func (o *DistributionConfiguration) SetGeoFencing(v DistributionConfigurationGeoFencing) {
	o.GeoFencing = &v
}

// GetUrlSignature returns the UrlSignature field value if set, zero value otherwise.
func (o *DistributionConfiguration) GetUrlSignature() DistributionConfigurationUrlSignature {
	if o == nil || isNil(o.UrlSignature) {
		var ret DistributionConfigurationUrlSignature
		return ret
	}
	return *o.UrlSignature
}

// GetUrlSignatureOk returns a tuple with the UrlSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetUrlSignatureOk() (*DistributionConfigurationUrlSignature, bool) {
	if o == nil || isNil(o.UrlSignature) {
    return nil, false
	}
	return o.UrlSignature, true
}

// HasUrlSignature returns a boolean if a field has been set.
func (o *DistributionConfiguration) HasUrlSignature() bool {
	if o != nil && !isNil(o.UrlSignature) {
		return true
	}

	return false
}

// SetUrlSignature gets a reference to the given DistributionConfigurationUrlSignature and assigns it to the UrlSignature field.
func (o *DistributionConfiguration) SetUrlSignature(v DistributionConfigurationUrlSignature) {
	o.UrlSignature = &v
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *DistributionConfiguration) GetCertificateId() string {
	if o == nil || isNil(o.CertificateId) {
		var ret string
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetCertificateIdOk() (*string, bool) {
	if o == nil || isNil(o.CertificateId) {
    return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *DistributionConfiguration) HasCertificateId() bool {
	if o != nil && !isNil(o.CertificateId) {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given string and assigns it to the CertificateId field.
func (o *DistributionConfiguration) SetCertificateId(v string) {
	o.CertificateId = &v
}

// GetSupplementaryDistributionNetworks returns the SupplementaryDistributionNetworks field value if set, zero value otherwise.
func (o *DistributionConfiguration) GetSupplementaryDistributionNetworks() DistributionConfigurationSupplementaryDistributionNetworks {
	if o == nil || isNil(o.SupplementaryDistributionNetworks) {
		var ret DistributionConfigurationSupplementaryDistributionNetworks
		return ret
	}
	return *o.SupplementaryDistributionNetworks
}

// GetSupplementaryDistributionNetworksOk returns a tuple with the SupplementaryDistributionNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfiguration) GetSupplementaryDistributionNetworksOk() (*DistributionConfigurationSupplementaryDistributionNetworks, bool) {
	if o == nil || isNil(o.SupplementaryDistributionNetworks) {
    return nil, false
	}
	return o.SupplementaryDistributionNetworks, true
}

// HasSupplementaryDistributionNetworks returns a boolean if a field has been set.
func (o *DistributionConfiguration) HasSupplementaryDistributionNetworks() bool {
	if o != nil && !isNil(o.SupplementaryDistributionNetworks) {
		return true
	}

	return false
}

// SetSupplementaryDistributionNetworks gets a reference to the given DistributionConfigurationSupplementaryDistributionNetworks and assigns it to the SupplementaryDistributionNetworks field.
func (o *DistributionConfiguration) SetSupplementaryDistributionNetworks(v DistributionConfigurationSupplementaryDistributionNetworks) {
	o.SupplementaryDistributionNetworks = &v
}

func (o DistributionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContentPreparationTemplateId) {
		toSerialize["contentPreparationTemplateId"] = o.ContentPreparationTemplateId
	}
	if true {
		toSerialize["canonicalDomainName"] = o.CanonicalDomainName
	}
	if true {
		toSerialize["domainNameAlias"] = o.DomainNameAlias
	}
	if !isNil(o.PathRewriteRules) {
		toSerialize["pathRewriteRules"] = o.PathRewriteRules
	}
	if !isNil(o.CachingConfigurations) {
		toSerialize["cachingConfigurations"] = o.CachingConfigurations
	}
	if !isNil(o.GeoFencing) {
		toSerialize["geoFencing"] = o.GeoFencing
	}
	if !isNil(o.UrlSignature) {
		toSerialize["urlSignature"] = o.UrlSignature
	}
	if !isNil(o.CertificateId) {
		toSerialize["certificateId"] = o.CertificateId
	}
	if !isNil(o.SupplementaryDistributionNetworks) {
		toSerialize["supplementaryDistributionNetworks"] = o.SupplementaryDistributionNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableDistributionConfiguration struct {
	value *DistributionConfiguration
	isSet bool
}

func (v NullableDistributionConfiguration) Get() *DistributionConfiguration {
	return v.value
}

func (v *NullableDistributionConfiguration) Set(val *DistributionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionConfiguration(val *DistributionConfiguration) *NullableDistributionConfiguration {
	return &NullableDistributionConfiguration{value: val, isSet: true}
}

func (v NullableDistributionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


