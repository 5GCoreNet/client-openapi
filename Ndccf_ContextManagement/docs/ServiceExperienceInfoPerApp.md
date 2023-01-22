# ServiceExperienceInfoPerApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**AppServerIns** | Pointer to [**AddrFqdn**](AddrFqdn.md) |  | [optional] 
**SvcExpPerFlows** | [**[]ServiceExperienceInfoPerFlow**](ServiceExperienceInfoPerFlow.md) |  | 
**Gpsis** | Pointer to **[]string** |  | [optional] 
**Supis** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceExperienceInfoPerApp

`func NewServiceExperienceInfoPerApp(svcExpPerFlows []ServiceExperienceInfoPerFlow, ) *ServiceExperienceInfoPerApp`

NewServiceExperienceInfoPerApp instantiates a new ServiceExperienceInfoPerApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceExperienceInfoPerAppWithDefaults

`func NewServiceExperienceInfoPerAppWithDefaults() *ServiceExperienceInfoPerApp`

NewServiceExperienceInfoPerAppWithDefaults instantiates a new ServiceExperienceInfoPerApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ServiceExperienceInfoPerApp) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceExperienceInfoPerApp) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceExperienceInfoPerApp) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceExperienceInfoPerApp) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppServerIns

`func (o *ServiceExperienceInfoPerApp) GetAppServerIns() AddrFqdn`

GetAppServerIns returns the AppServerIns field if non-nil, zero value otherwise.

### GetAppServerInsOk

`func (o *ServiceExperienceInfoPerApp) GetAppServerInsOk() (*AddrFqdn, bool)`

GetAppServerInsOk returns a tuple with the AppServerIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerIns

`func (o *ServiceExperienceInfoPerApp) SetAppServerIns(v AddrFqdn)`

SetAppServerIns sets AppServerIns field to given value.

### HasAppServerIns

`func (o *ServiceExperienceInfoPerApp) HasAppServerIns() bool`

HasAppServerIns returns a boolean if a field has been set.

### GetSvcExpPerFlows

`func (o *ServiceExperienceInfoPerApp) GetSvcExpPerFlows() []ServiceExperienceInfoPerFlow`

GetSvcExpPerFlows returns the SvcExpPerFlows field if non-nil, zero value otherwise.

### GetSvcExpPerFlowsOk

`func (o *ServiceExperienceInfoPerApp) GetSvcExpPerFlowsOk() (*[]ServiceExperienceInfoPerFlow, bool)`

GetSvcExpPerFlowsOk returns a tuple with the SvcExpPerFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExpPerFlows

`func (o *ServiceExperienceInfoPerApp) SetSvcExpPerFlows(v []ServiceExperienceInfoPerFlow)`

SetSvcExpPerFlows sets SvcExpPerFlows field to given value.


### GetGpsis

`func (o *ServiceExperienceInfoPerApp) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *ServiceExperienceInfoPerApp) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *ServiceExperienceInfoPerApp) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *ServiceExperienceInfoPerApp) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetSupis

`func (o *ServiceExperienceInfoPerApp) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *ServiceExperienceInfoPerApp) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *ServiceExperienceInfoPerApp) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *ServiceExperienceInfoPerApp) HasSupis() bool`

HasSupis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


