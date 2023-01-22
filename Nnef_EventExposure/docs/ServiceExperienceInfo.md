# ServiceExperienceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**Supis** | Pointer to **[]string** |  | [optional] 
**SvcExpPerFlows** | [**[]ServiceExperienceInfoPerFlow**](ServiceExperienceInfoPerFlow.md) |  | 

## Methods

### NewServiceExperienceInfo

`func NewServiceExperienceInfo(svcExpPerFlows []ServiceExperienceInfoPerFlow, ) *ServiceExperienceInfo`

NewServiceExperienceInfo instantiates a new ServiceExperienceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceExperienceInfoWithDefaults

`func NewServiceExperienceInfoWithDefaults() *ServiceExperienceInfo`

NewServiceExperienceInfoWithDefaults instantiates a new ServiceExperienceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ServiceExperienceInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceExperienceInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceExperienceInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceExperienceInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetSupis

`func (o *ServiceExperienceInfo) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *ServiceExperienceInfo) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *ServiceExperienceInfo) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *ServiceExperienceInfo) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetSvcExpPerFlows

`func (o *ServiceExperienceInfo) GetSvcExpPerFlows() []ServiceExperienceInfoPerFlow`

GetSvcExpPerFlows returns the SvcExpPerFlows field if non-nil, zero value otherwise.

### GetSvcExpPerFlowsOk

`func (o *ServiceExperienceInfo) GetSvcExpPerFlowsOk() (*[]ServiceExperienceInfoPerFlow, bool)`

GetSvcExpPerFlowsOk returns a tuple with the SvcExpPerFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExpPerFlows

`func (o *ServiceExperienceInfo) SetSvcExpPerFlows(v []ServiceExperienceInfoPerFlow)`

SetSvcExpPerFlows sets SvcExpPerFlows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


