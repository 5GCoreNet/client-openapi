# ServiceExperienceInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**Supis** | Pointer to **[]string** |  | [optional] 
**SvcExpPerFlows** | [**[]ServiceExperienceInfoPerFlow**](ServiceExperienceInfoPerFlow.md) |  | 

## Methods

### NewServiceExperienceInfo1

`func NewServiceExperienceInfo1(svcExpPerFlows []ServiceExperienceInfoPerFlow, ) *ServiceExperienceInfo1`

NewServiceExperienceInfo1 instantiates a new ServiceExperienceInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceExperienceInfo1WithDefaults

`func NewServiceExperienceInfo1WithDefaults() *ServiceExperienceInfo1`

NewServiceExperienceInfo1WithDefaults instantiates a new ServiceExperienceInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ServiceExperienceInfo1) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceExperienceInfo1) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceExperienceInfo1) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceExperienceInfo1) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetSupis

`func (o *ServiceExperienceInfo1) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *ServiceExperienceInfo1) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *ServiceExperienceInfo1) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *ServiceExperienceInfo1) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetSvcExpPerFlows

`func (o *ServiceExperienceInfo1) GetSvcExpPerFlows() []ServiceExperienceInfoPerFlow`

GetSvcExpPerFlows returns the SvcExpPerFlows field if non-nil, zero value otherwise.

### GetSvcExpPerFlowsOk

`func (o *ServiceExperienceInfo1) GetSvcExpPerFlowsOk() (*[]ServiceExperienceInfoPerFlow, bool)`

GetSvcExpPerFlowsOk returns a tuple with the SvcExpPerFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExpPerFlows

`func (o *ServiceExperienceInfo1) SetSvcExpPerFlows(v []ServiceExperienceInfoPerFlow)`

SetSvcExpPerFlows sets SvcExpPerFlows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


