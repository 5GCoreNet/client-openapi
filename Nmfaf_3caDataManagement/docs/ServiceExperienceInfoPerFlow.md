# ServiceExperienceInfoPerFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SvcExprc** | Pointer to [**SvcExperience**](SvcExperience.md) |  | [optional] 
**TimeIntev** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**Dnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**IpTrafficFilter** | Pointer to [**FlowInfo**](FlowInfo.md) |  | [optional] 
**EthTrafficFilter** | Pointer to [**EthFlowDescription**](EthFlowDescription.md) |  | [optional] 

## Methods

### NewServiceExperienceInfoPerFlow

`func NewServiceExperienceInfoPerFlow() *ServiceExperienceInfoPerFlow`

NewServiceExperienceInfoPerFlow instantiates a new ServiceExperienceInfoPerFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceExperienceInfoPerFlowWithDefaults

`func NewServiceExperienceInfoPerFlowWithDefaults() *ServiceExperienceInfoPerFlow`

NewServiceExperienceInfoPerFlowWithDefaults instantiates a new ServiceExperienceInfoPerFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSvcExprc

`func (o *ServiceExperienceInfoPerFlow) GetSvcExprc() SvcExperience`

GetSvcExprc returns the SvcExprc field if non-nil, zero value otherwise.

### GetSvcExprcOk

`func (o *ServiceExperienceInfoPerFlow) GetSvcExprcOk() (*SvcExperience, bool)`

GetSvcExprcOk returns a tuple with the SvcExprc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExprc

`func (o *ServiceExperienceInfoPerFlow) SetSvcExprc(v SvcExperience)`

SetSvcExprc sets SvcExprc field to given value.

### HasSvcExprc

`func (o *ServiceExperienceInfoPerFlow) HasSvcExprc() bool`

HasSvcExprc returns a boolean if a field has been set.

### GetTimeIntev

`func (o *ServiceExperienceInfoPerFlow) GetTimeIntev() TimeWindow`

GetTimeIntev returns the TimeIntev field if non-nil, zero value otherwise.

### GetTimeIntevOk

`func (o *ServiceExperienceInfoPerFlow) GetTimeIntevOk() (*TimeWindow, bool)`

GetTimeIntevOk returns a tuple with the TimeIntev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntev

`func (o *ServiceExperienceInfoPerFlow) SetTimeIntev(v TimeWindow)`

SetTimeIntev sets TimeIntev field to given value.

### HasTimeIntev

`func (o *ServiceExperienceInfoPerFlow) HasTimeIntev() bool`

HasTimeIntev returns a boolean if a field has been set.

### GetDnai

`func (o *ServiceExperienceInfoPerFlow) GetDnai() string`

GetDnai returns the Dnai field if non-nil, zero value otherwise.

### GetDnaiOk

`func (o *ServiceExperienceInfoPerFlow) GetDnaiOk() (*string, bool)`

GetDnaiOk returns a tuple with the Dnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnai

`func (o *ServiceExperienceInfoPerFlow) SetDnai(v string)`

SetDnai sets Dnai field to given value.

### HasDnai

`func (o *ServiceExperienceInfoPerFlow) HasDnai() bool`

HasDnai returns a boolean if a field has been set.

### GetIpTrafficFilter

`func (o *ServiceExperienceInfoPerFlow) GetIpTrafficFilter() FlowInfo`

GetIpTrafficFilter returns the IpTrafficFilter field if non-nil, zero value otherwise.

### GetIpTrafficFilterOk

`func (o *ServiceExperienceInfoPerFlow) GetIpTrafficFilterOk() (*FlowInfo, bool)`

GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTrafficFilter

`func (o *ServiceExperienceInfoPerFlow) SetIpTrafficFilter(v FlowInfo)`

SetIpTrafficFilter sets IpTrafficFilter field to given value.

### HasIpTrafficFilter

`func (o *ServiceExperienceInfoPerFlow) HasIpTrafficFilter() bool`

HasIpTrafficFilter returns a boolean if a field has been set.

### GetEthTrafficFilter

`func (o *ServiceExperienceInfoPerFlow) GetEthTrafficFilter() EthFlowDescription`

GetEthTrafficFilter returns the EthTrafficFilter field if non-nil, zero value otherwise.

### GetEthTrafficFilterOk

`func (o *ServiceExperienceInfoPerFlow) GetEthTrafficFilterOk() (*EthFlowDescription, bool)`

GetEthTrafficFilterOk returns a tuple with the EthTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthTrafficFilter

`func (o *ServiceExperienceInfoPerFlow) SetEthTrafficFilter(v EthFlowDescription)`

SetEthTrafficFilter sets EthTrafficFilter field to given value.

### HasEthTrafficFilter

`func (o *ServiceExperienceInfoPerFlow) HasEthTrafficFilter() bool`

HasEthTrafficFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


