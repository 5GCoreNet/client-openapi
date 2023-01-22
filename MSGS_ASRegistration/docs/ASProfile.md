# ASProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**AppProviders** | Pointer to **[]string** | The provider of the AS. | [optional] 
**AppSenarios** | Pointer to **[]string** | The application scenario. | [optional] 
**AppCategory** | Pointer to **string** |  | [optional] 
**AsStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewASProfile

`func NewASProfile() *ASProfile`

NewASProfile instantiates a new ASProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASProfileWithDefaults

`func NewASProfileWithDefaults() *ASProfile`

NewASProfileWithDefaults instantiates a new ASProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ASProfile) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ASProfile) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ASProfile) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ASProfile) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAppProviders

`func (o *ASProfile) GetAppProviders() []string`

GetAppProviders returns the AppProviders field if non-nil, zero value otherwise.

### GetAppProvidersOk

`func (o *ASProfile) GetAppProvidersOk() (*[]string, bool)`

GetAppProvidersOk returns a tuple with the AppProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProviders

`func (o *ASProfile) SetAppProviders(v []string)`

SetAppProviders sets AppProviders field to given value.

### HasAppProviders

`func (o *ASProfile) HasAppProviders() bool`

HasAppProviders returns a boolean if a field has been set.

### GetAppSenarios

`func (o *ASProfile) GetAppSenarios() []string`

GetAppSenarios returns the AppSenarios field if non-nil, zero value otherwise.

### GetAppSenariosOk

`func (o *ASProfile) GetAppSenariosOk() (*[]string, bool)`

GetAppSenariosOk returns a tuple with the AppSenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSenarios

`func (o *ASProfile) SetAppSenarios(v []string)`

SetAppSenarios sets AppSenarios field to given value.

### HasAppSenarios

`func (o *ASProfile) HasAppSenarios() bool`

HasAppSenarios returns a boolean if a field has been set.

### GetAppCategory

`func (o *ASProfile) GetAppCategory() string`

GetAppCategory returns the AppCategory field if non-nil, zero value otherwise.

### GetAppCategoryOk

`func (o *ASProfile) GetAppCategoryOk() (*string, bool)`

GetAppCategoryOk returns a tuple with the AppCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCategory

`func (o *ASProfile) SetAppCategory(v string)`

SetAppCategory sets AppCategory field to given value.

### HasAppCategory

`func (o *ASProfile) HasAppCategory() bool`

HasAppCategory returns a boolean if a field has been set.

### GetAsStatus

`func (o *ASProfile) GetAsStatus() string`

GetAsStatus returns the AsStatus field if non-nil, zero value otherwise.

### GetAsStatusOk

`func (o *ASProfile) GetAsStatusOk() (*string, bool)`

GetAsStatusOk returns a tuple with the AsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsStatus

`func (o *ASProfile) SetAsStatus(v string)`

SetAsStatus sets AsStatus field to given value.

### HasAsStatus

`func (o *ASProfile) HasAsStatus() bool`

HasAsStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


