# NSPAContainerInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latency** | Pointer to **int32** |  | [optional] 
**Throughput** | Pointer to [**Throughput**](Throughput.md) |  | [optional] 
**MaximumPacketLossRate** | Pointer to **string** |  | [optional] 
**ServiceExperienceStatisticsData** | Pointer to [**ServiceExperienceInfo**](ServiceExperienceInfo.md) |  | [optional] 
**TheNumberOfPDUSessions** | Pointer to **int32** |  | [optional] 
**TheNumberOfRegisteredSubscribers** | Pointer to **int32** |  | [optional] 
**LoadLevel** | Pointer to [**NsiLoadLevelInfo**](NsiLoadLevelInfo.md) |  | [optional] 

## Methods

### NewNSPAContainerInformation

`func NewNSPAContainerInformation() *NSPAContainerInformation`

NewNSPAContainerInformation instantiates a new NSPAContainerInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSPAContainerInformationWithDefaults

`func NewNSPAContainerInformationWithDefaults() *NSPAContainerInformation`

NewNSPAContainerInformationWithDefaults instantiates a new NSPAContainerInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatency

`func (o *NSPAContainerInformation) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *NSPAContainerInformation) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *NSPAContainerInformation) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *NSPAContainerInformation) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetThroughput

`func (o *NSPAContainerInformation) GetThroughput() Throughput`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *NSPAContainerInformation) GetThroughputOk() (*Throughput, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *NSPAContainerInformation) SetThroughput(v Throughput)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *NSPAContainerInformation) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetMaximumPacketLossRate

`func (o *NSPAContainerInformation) GetMaximumPacketLossRate() string`

GetMaximumPacketLossRate returns the MaximumPacketLossRate field if non-nil, zero value otherwise.

### GetMaximumPacketLossRateOk

`func (o *NSPAContainerInformation) GetMaximumPacketLossRateOk() (*string, bool)`

GetMaximumPacketLossRateOk returns a tuple with the MaximumPacketLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPacketLossRate

`func (o *NSPAContainerInformation) SetMaximumPacketLossRate(v string)`

SetMaximumPacketLossRate sets MaximumPacketLossRate field to given value.

### HasMaximumPacketLossRate

`func (o *NSPAContainerInformation) HasMaximumPacketLossRate() bool`

HasMaximumPacketLossRate returns a boolean if a field has been set.

### GetServiceExperienceStatisticsData

`func (o *NSPAContainerInformation) GetServiceExperienceStatisticsData() ServiceExperienceInfo`

GetServiceExperienceStatisticsData returns the ServiceExperienceStatisticsData field if non-nil, zero value otherwise.

### GetServiceExperienceStatisticsDataOk

`func (o *NSPAContainerInformation) GetServiceExperienceStatisticsDataOk() (*ServiceExperienceInfo, bool)`

GetServiceExperienceStatisticsDataOk returns a tuple with the ServiceExperienceStatisticsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExperienceStatisticsData

`func (o *NSPAContainerInformation) SetServiceExperienceStatisticsData(v ServiceExperienceInfo)`

SetServiceExperienceStatisticsData sets ServiceExperienceStatisticsData field to given value.

### HasServiceExperienceStatisticsData

`func (o *NSPAContainerInformation) HasServiceExperienceStatisticsData() bool`

HasServiceExperienceStatisticsData returns a boolean if a field has been set.

### GetTheNumberOfPDUSessions

`func (o *NSPAContainerInformation) GetTheNumberOfPDUSessions() int32`

GetTheNumberOfPDUSessions returns the TheNumberOfPDUSessions field if non-nil, zero value otherwise.

### GetTheNumberOfPDUSessionsOk

`func (o *NSPAContainerInformation) GetTheNumberOfPDUSessionsOk() (*int32, bool)`

GetTheNumberOfPDUSessionsOk returns a tuple with the TheNumberOfPDUSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheNumberOfPDUSessions

`func (o *NSPAContainerInformation) SetTheNumberOfPDUSessions(v int32)`

SetTheNumberOfPDUSessions sets TheNumberOfPDUSessions field to given value.

### HasTheNumberOfPDUSessions

`func (o *NSPAContainerInformation) HasTheNumberOfPDUSessions() bool`

HasTheNumberOfPDUSessions returns a boolean if a field has been set.

### GetTheNumberOfRegisteredSubscribers

`func (o *NSPAContainerInformation) GetTheNumberOfRegisteredSubscribers() int32`

GetTheNumberOfRegisteredSubscribers returns the TheNumberOfRegisteredSubscribers field if non-nil, zero value otherwise.

### GetTheNumberOfRegisteredSubscribersOk

`func (o *NSPAContainerInformation) GetTheNumberOfRegisteredSubscribersOk() (*int32, bool)`

GetTheNumberOfRegisteredSubscribersOk returns a tuple with the TheNumberOfRegisteredSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheNumberOfRegisteredSubscribers

`func (o *NSPAContainerInformation) SetTheNumberOfRegisteredSubscribers(v int32)`

SetTheNumberOfRegisteredSubscribers sets TheNumberOfRegisteredSubscribers field to given value.

### HasTheNumberOfRegisteredSubscribers

`func (o *NSPAContainerInformation) HasTheNumberOfRegisteredSubscribers() bool`

HasTheNumberOfRegisteredSubscribers returns a boolean if a field has been set.

### GetLoadLevel

`func (o *NSPAContainerInformation) GetLoadLevel() NsiLoadLevelInfo`

GetLoadLevel returns the LoadLevel field if non-nil, zero value otherwise.

### GetLoadLevelOk

`func (o *NSPAContainerInformation) GetLoadLevelOk() (*NsiLoadLevelInfo, bool)`

GetLoadLevelOk returns a tuple with the LoadLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLevel

`func (o *NSPAContainerInformation) SetLoadLevel(v NsiLoadLevelInfo)`

SetLoadLevel sets LoadLevel field to given value.

### HasLoadLevel

`func (o *NSPAContainerInformation) HasLoadLevel() bool`

HasLoadLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


