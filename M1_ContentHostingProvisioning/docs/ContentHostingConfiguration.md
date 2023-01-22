# ContentHostingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IngestConfiguration** | [**IngestConfiguration**](IngestConfiguration.md) |  | 
**DistributionConfigurations** | [**[]DistributionConfiguration**](DistributionConfiguration.md) |  | 

## Methods

### NewContentHostingConfiguration

`func NewContentHostingConfiguration(name string, ingestConfiguration IngestConfiguration, distributionConfigurations []DistributionConfiguration, ) *ContentHostingConfiguration`

NewContentHostingConfiguration instantiates a new ContentHostingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentHostingConfigurationWithDefaults

`func NewContentHostingConfigurationWithDefaults() *ContentHostingConfiguration`

NewContentHostingConfigurationWithDefaults instantiates a new ContentHostingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContentHostingConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentHostingConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentHostingConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetIngestConfiguration

`func (o *ContentHostingConfiguration) GetIngestConfiguration() IngestConfiguration`

GetIngestConfiguration returns the IngestConfiguration field if non-nil, zero value otherwise.

### GetIngestConfigurationOk

`func (o *ContentHostingConfiguration) GetIngestConfigurationOk() (*IngestConfiguration, bool)`

GetIngestConfigurationOk returns a tuple with the IngestConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestConfiguration

`func (o *ContentHostingConfiguration) SetIngestConfiguration(v IngestConfiguration)`

SetIngestConfiguration sets IngestConfiguration field to given value.


### GetDistributionConfigurations

`func (o *ContentHostingConfiguration) GetDistributionConfigurations() []DistributionConfiguration`

GetDistributionConfigurations returns the DistributionConfigurations field if non-nil, zero value otherwise.

### GetDistributionConfigurationsOk

`func (o *ContentHostingConfiguration) GetDistributionConfigurationsOk() (*[]DistributionConfiguration, bool)`

GetDistributionConfigurationsOk returns a tuple with the DistributionConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionConfigurations

`func (o *ContentHostingConfiguration) SetDistributionConfigurations(v []DistributionConfiguration)`

SetDistributionConfigurations sets DistributionConfigurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


