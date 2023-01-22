# ServiceAccessInformationResourceClientEdgeResourcesConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EligibilityCriteria** | Pointer to [**EdgeProcessingEligibilityCriteria**](EdgeProcessingEligibilityCriteria.md) |  | [optional] 
**EasDiscoveryTemplate** | [**EASDiscoveryTemplate**](EASDiscoveryTemplate.md) |  | 
**EasRelocationRequirements** | Pointer to [**M5EASRelocationRequirements**](M5EASRelocationRequirements.md) |  | [optional] 

## Methods

### NewServiceAccessInformationResourceClientEdgeResourcesConfiguration

`func NewServiceAccessInformationResourceClientEdgeResourcesConfiguration(easDiscoveryTemplate EASDiscoveryTemplate, ) *ServiceAccessInformationResourceClientEdgeResourcesConfiguration`

NewServiceAccessInformationResourceClientEdgeResourcesConfiguration instantiates a new ServiceAccessInformationResourceClientEdgeResourcesConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessInformationResourceClientEdgeResourcesConfigurationWithDefaults

`func NewServiceAccessInformationResourceClientEdgeResourcesConfigurationWithDefaults() *ServiceAccessInformationResourceClientEdgeResourcesConfiguration`

NewServiceAccessInformationResourceClientEdgeResourcesConfigurationWithDefaults instantiates a new ServiceAccessInformationResourceClientEdgeResourcesConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligibilityCriteria

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) GetEligibilityCriteria() EdgeProcessingEligibilityCriteria`

GetEligibilityCriteria returns the EligibilityCriteria field if non-nil, zero value otherwise.

### GetEligibilityCriteriaOk

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) GetEligibilityCriteriaOk() (*EdgeProcessingEligibilityCriteria, bool)`

GetEligibilityCriteriaOk returns a tuple with the EligibilityCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityCriteria

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) SetEligibilityCriteria(v EdgeProcessingEligibilityCriteria)`

SetEligibilityCriteria sets EligibilityCriteria field to given value.

### HasEligibilityCriteria

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) HasEligibilityCriteria() bool`

HasEligibilityCriteria returns a boolean if a field has been set.

### GetEasDiscoveryTemplate

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) GetEasDiscoveryTemplate() EASDiscoveryTemplate`

GetEasDiscoveryTemplate returns the EasDiscoveryTemplate field if non-nil, zero value otherwise.

### GetEasDiscoveryTemplateOk

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) GetEasDiscoveryTemplateOk() (*EASDiscoveryTemplate, bool)`

GetEasDiscoveryTemplateOk returns a tuple with the EasDiscoveryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDiscoveryTemplate

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) SetEasDiscoveryTemplate(v EASDiscoveryTemplate)`

SetEasDiscoveryTemplate sets EasDiscoveryTemplate field to given value.


### GetEasRelocationRequirements

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) GetEasRelocationRequirements() M5EASRelocationRequirements`

GetEasRelocationRequirements returns the EasRelocationRequirements field if non-nil, zero value otherwise.

### GetEasRelocationRequirementsOk

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) GetEasRelocationRequirementsOk() (*M5EASRelocationRequirements, bool)`

GetEasRelocationRequirementsOk returns a tuple with the EasRelocationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasRelocationRequirements

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) SetEasRelocationRequirements(v M5EASRelocationRequirements)`

SetEasRelocationRequirements sets EasRelocationRequirements field to given value.

### HasEasRelocationRequirements

`func (o *ServiceAccessInformationResourceClientEdgeResourcesConfiguration) HasEasRelocationRequirements() bool`

HasEasRelocationRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


