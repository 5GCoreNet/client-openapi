# PerFlowServiceExperienceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceExperience** | [**SvcExperience**](SvcExperience.md) |  | 
**TimeInterval** | [**TimeWindow**](TimeWindow.md) |  | 
**RemoteEndpoint** | [**AddrFqdn**](AddrFqdn.md) |  | 

## Methods

### NewPerFlowServiceExperienceInfo

`func NewPerFlowServiceExperienceInfo(serviceExperience SvcExperience, timeInterval TimeWindow, remoteEndpoint AddrFqdn, ) *PerFlowServiceExperienceInfo`

NewPerFlowServiceExperienceInfo instantiates a new PerFlowServiceExperienceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerFlowServiceExperienceInfoWithDefaults

`func NewPerFlowServiceExperienceInfoWithDefaults() *PerFlowServiceExperienceInfo`

NewPerFlowServiceExperienceInfoWithDefaults instantiates a new PerFlowServiceExperienceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceExperience

`func (o *PerFlowServiceExperienceInfo) GetServiceExperience() SvcExperience`

GetServiceExperience returns the ServiceExperience field if non-nil, zero value otherwise.

### GetServiceExperienceOk

`func (o *PerFlowServiceExperienceInfo) GetServiceExperienceOk() (*SvcExperience, bool)`

GetServiceExperienceOk returns a tuple with the ServiceExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExperience

`func (o *PerFlowServiceExperienceInfo) SetServiceExperience(v SvcExperience)`

SetServiceExperience sets ServiceExperience field to given value.


### GetTimeInterval

`func (o *PerFlowServiceExperienceInfo) GetTimeInterval() TimeWindow`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *PerFlowServiceExperienceInfo) GetTimeIntervalOk() (*TimeWindow, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *PerFlowServiceExperienceInfo) SetTimeInterval(v TimeWindow)`

SetTimeInterval sets TimeInterval field to given value.


### GetRemoteEndpoint

`func (o *PerFlowServiceExperienceInfo) GetRemoteEndpoint() AddrFqdn`

GetRemoteEndpoint returns the RemoteEndpoint field if non-nil, zero value otherwise.

### GetRemoteEndpointOk

`func (o *PerFlowServiceExperienceInfo) GetRemoteEndpointOk() (*AddrFqdn, bool)`

GetRemoteEndpointOk returns a tuple with the RemoteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndpoint

`func (o *PerFlowServiceExperienceInfo) SetRemoteEndpoint(v AddrFqdn)`

SetRemoteEndpoint sets RemoteEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


