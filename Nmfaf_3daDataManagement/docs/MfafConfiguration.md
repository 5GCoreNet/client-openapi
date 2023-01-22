# MfafConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageConfigurations** | [**[]MessageConfiguration**](MessageConfiguration.md) | The configuration of the MFAF for mapping data or analytics. | 

## Methods

### NewMfafConfiguration

`func NewMfafConfiguration(messageConfigurations []MessageConfiguration, ) *MfafConfiguration`

NewMfafConfiguration instantiates a new MfafConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfafConfigurationWithDefaults

`func NewMfafConfigurationWithDefaults() *MfafConfiguration`

NewMfafConfigurationWithDefaults instantiates a new MfafConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageConfigurations

`func (o *MfafConfiguration) GetMessageConfigurations() []MessageConfiguration`

GetMessageConfigurations returns the MessageConfigurations field if non-nil, zero value otherwise.

### GetMessageConfigurationsOk

`func (o *MfafConfiguration) GetMessageConfigurationsOk() (*[]MessageConfiguration, bool)`

GetMessageConfigurationsOk returns a tuple with the MessageConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageConfigurations

`func (o *MfafConfiguration) SetMessageConfigurations(v []MessageConfiguration)`

SetMessageConfigurations sets MessageConfigurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


