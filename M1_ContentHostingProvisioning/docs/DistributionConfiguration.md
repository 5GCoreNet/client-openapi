# DistributionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentPreparationTemplateId** | Pointer to **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | [optional] 
**CanonicalDomainName** | **string** |  | 
**DomainNameAlias** | **string** |  | 
**PathRewriteRules** | Pointer to [**[]PathRewriteRule**](PathRewriteRule.md) |  | [optional] 
**CachingConfigurations** | Pointer to [**[]CachingConfiguration**](CachingConfiguration.md) |  | [optional] 
**GeoFencing** | Pointer to [**DistributionConfigurationGeoFencing**](DistributionConfigurationGeoFencing.md) |  | [optional] 
**UrlSignature** | Pointer to [**DistributionConfigurationUrlSignature**](DistributionConfigurationUrlSignature.md) |  | [optional] 
**CertificateId** | Pointer to **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | [optional] 
**SupplementaryDistributionNetworks** | Pointer to [**DistributionConfigurationSupplementaryDistributionNetworks**](DistributionConfigurationSupplementaryDistributionNetworks.md) |  | [optional] 

## Methods

### NewDistributionConfiguration

`func NewDistributionConfiguration(canonicalDomainName string, domainNameAlias string, ) *DistributionConfiguration`

NewDistributionConfiguration instantiates a new DistributionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionConfigurationWithDefaults

`func NewDistributionConfigurationWithDefaults() *DistributionConfiguration`

NewDistributionConfigurationWithDefaults instantiates a new DistributionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentPreparationTemplateId

`func (o *DistributionConfiguration) GetContentPreparationTemplateId() string`

GetContentPreparationTemplateId returns the ContentPreparationTemplateId field if non-nil, zero value otherwise.

### GetContentPreparationTemplateIdOk

`func (o *DistributionConfiguration) GetContentPreparationTemplateIdOk() (*string, bool)`

GetContentPreparationTemplateIdOk returns a tuple with the ContentPreparationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPreparationTemplateId

`func (o *DistributionConfiguration) SetContentPreparationTemplateId(v string)`

SetContentPreparationTemplateId sets ContentPreparationTemplateId field to given value.

### HasContentPreparationTemplateId

`func (o *DistributionConfiguration) HasContentPreparationTemplateId() bool`

HasContentPreparationTemplateId returns a boolean if a field has been set.

### GetCanonicalDomainName

`func (o *DistributionConfiguration) GetCanonicalDomainName() string`

GetCanonicalDomainName returns the CanonicalDomainName field if non-nil, zero value otherwise.

### GetCanonicalDomainNameOk

`func (o *DistributionConfiguration) GetCanonicalDomainNameOk() (*string, bool)`

GetCanonicalDomainNameOk returns a tuple with the CanonicalDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalDomainName

`func (o *DistributionConfiguration) SetCanonicalDomainName(v string)`

SetCanonicalDomainName sets CanonicalDomainName field to given value.


### GetDomainNameAlias

`func (o *DistributionConfiguration) GetDomainNameAlias() string`

GetDomainNameAlias returns the DomainNameAlias field if non-nil, zero value otherwise.

### GetDomainNameAliasOk

`func (o *DistributionConfiguration) GetDomainNameAliasOk() (*string, bool)`

GetDomainNameAliasOk returns a tuple with the DomainNameAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameAlias

`func (o *DistributionConfiguration) SetDomainNameAlias(v string)`

SetDomainNameAlias sets DomainNameAlias field to given value.


### GetPathRewriteRules

`func (o *DistributionConfiguration) GetPathRewriteRules() []PathRewriteRule`

GetPathRewriteRules returns the PathRewriteRules field if non-nil, zero value otherwise.

### GetPathRewriteRulesOk

`func (o *DistributionConfiguration) GetPathRewriteRulesOk() (*[]PathRewriteRule, bool)`

GetPathRewriteRulesOk returns a tuple with the PathRewriteRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathRewriteRules

`func (o *DistributionConfiguration) SetPathRewriteRules(v []PathRewriteRule)`

SetPathRewriteRules sets PathRewriteRules field to given value.

### HasPathRewriteRules

`func (o *DistributionConfiguration) HasPathRewriteRules() bool`

HasPathRewriteRules returns a boolean if a field has been set.

### GetCachingConfigurations

`func (o *DistributionConfiguration) GetCachingConfigurations() []CachingConfiguration`

GetCachingConfigurations returns the CachingConfigurations field if non-nil, zero value otherwise.

### GetCachingConfigurationsOk

`func (o *DistributionConfiguration) GetCachingConfigurationsOk() (*[]CachingConfiguration, bool)`

GetCachingConfigurationsOk returns a tuple with the CachingConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingConfigurations

`func (o *DistributionConfiguration) SetCachingConfigurations(v []CachingConfiguration)`

SetCachingConfigurations sets CachingConfigurations field to given value.

### HasCachingConfigurations

`func (o *DistributionConfiguration) HasCachingConfigurations() bool`

HasCachingConfigurations returns a boolean if a field has been set.

### GetGeoFencing

`func (o *DistributionConfiguration) GetGeoFencing() DistributionConfigurationGeoFencing`

GetGeoFencing returns the GeoFencing field if non-nil, zero value otherwise.

### GetGeoFencingOk

`func (o *DistributionConfiguration) GetGeoFencingOk() (*DistributionConfigurationGeoFencing, bool)`

GetGeoFencingOk returns a tuple with the GeoFencing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoFencing

`func (o *DistributionConfiguration) SetGeoFencing(v DistributionConfigurationGeoFencing)`

SetGeoFencing sets GeoFencing field to given value.

### HasGeoFencing

`func (o *DistributionConfiguration) HasGeoFencing() bool`

HasGeoFencing returns a boolean if a field has been set.

### GetUrlSignature

`func (o *DistributionConfiguration) GetUrlSignature() DistributionConfigurationUrlSignature`

GetUrlSignature returns the UrlSignature field if non-nil, zero value otherwise.

### GetUrlSignatureOk

`func (o *DistributionConfiguration) GetUrlSignatureOk() (*DistributionConfigurationUrlSignature, bool)`

GetUrlSignatureOk returns a tuple with the UrlSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSignature

`func (o *DistributionConfiguration) SetUrlSignature(v DistributionConfigurationUrlSignature)`

SetUrlSignature sets UrlSignature field to given value.

### HasUrlSignature

`func (o *DistributionConfiguration) HasUrlSignature() bool`

HasUrlSignature returns a boolean if a field has been set.

### GetCertificateId

`func (o *DistributionConfiguration) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *DistributionConfiguration) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *DistributionConfiguration) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *DistributionConfiguration) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetSupplementaryDistributionNetworks

`func (o *DistributionConfiguration) GetSupplementaryDistributionNetworks() DistributionConfigurationSupplementaryDistributionNetworks`

GetSupplementaryDistributionNetworks returns the SupplementaryDistributionNetworks field if non-nil, zero value otherwise.

### GetSupplementaryDistributionNetworksOk

`func (o *DistributionConfiguration) GetSupplementaryDistributionNetworksOk() (*DistributionConfigurationSupplementaryDistributionNetworks, bool)`

GetSupplementaryDistributionNetworksOk returns a tuple with the SupplementaryDistributionNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryDistributionNetworks

`func (o *DistributionConfiguration) SetSupplementaryDistributionNetworks(v DistributionConfigurationSupplementaryDistributionNetworks)`

SetSupplementaryDistributionNetworks sets SupplementaryDistributionNetworks field to given value.

### HasSupplementaryDistributionNetworks

`func (o *DistributionConfiguration) HasSupplementaryDistributionNetworks() bool`

HasSupplementaryDistributionNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


