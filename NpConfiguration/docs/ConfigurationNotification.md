# ConfigurationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**ConfigResults** | Pointer to [**[]ConfigResult**](ConfigResult.md) | The grouping configuration result notification provided by the SCEF. | [optional] 
**AppliedParam** | Pointer to [**AppliedParameterConfiguration**](AppliedParameterConfiguration.md) |  | [optional] 

## Methods

### NewConfigurationNotification

`func NewConfigurationNotification(configuration string, ) *ConfigurationNotification`

NewConfigurationNotification instantiates a new ConfigurationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationNotificationWithDefaults

`func NewConfigurationNotificationWithDefaults() *ConfigurationNotification`

NewConfigurationNotificationWithDefaults instantiates a new ConfigurationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ConfigurationNotification) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ConfigurationNotification) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ConfigurationNotification) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.


### GetConfigResults

`func (o *ConfigurationNotification) GetConfigResults() []ConfigResult`

GetConfigResults returns the ConfigResults field if non-nil, zero value otherwise.

### GetConfigResultsOk

`func (o *ConfigurationNotification) GetConfigResultsOk() (*[]ConfigResult, bool)`

GetConfigResultsOk returns a tuple with the ConfigResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResults

`func (o *ConfigurationNotification) SetConfigResults(v []ConfigResult)`

SetConfigResults sets ConfigResults field to given value.

### HasConfigResults

`func (o *ConfigurationNotification) HasConfigResults() bool`

HasConfigResults returns a boolean if a field has been set.

### GetAppliedParam

`func (o *ConfigurationNotification) GetAppliedParam() AppliedParameterConfiguration`

GetAppliedParam returns the AppliedParam field if non-nil, zero value otherwise.

### GetAppliedParamOk

`func (o *ConfigurationNotification) GetAppliedParamOk() (*AppliedParameterConfiguration, bool)`

GetAppliedParamOk returns a tuple with the AppliedParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedParam

`func (o *ConfigurationNotification) SetAppliedParam(v AppliedParameterConfiguration)`

SetAppliedParam sets AppliedParam field to given value.

### HasAppliedParam

`func (o *ConfigurationNotification) HasAppliedParam() bool`

HasAppliedParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


