# ServiceProfileChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceProfileIdentifier** | Pointer to **string** |  | [optional] 
**SNSSAIList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**Latency** | Pointer to **int32** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**Jitter** | Pointer to **int32** |  | [optional] 
**Reliability** | Pointer to **string** |  | [optional] 
**MaxNumberofUEs** | Pointer to **int32** |  | [optional] 
**CoverageArea** | Pointer to **string** |  | [optional] 
**DLThptPerSlice** | Pointer to [**Throughput**](Throughput.md) |  | [optional] 
**DLThptPerUE** | Pointer to [**Throughput**](Throughput.md) |  | [optional] 
**ULThptPerSlice** | Pointer to [**Throughput**](Throughput.md) |  | [optional] 
**ULThptPerUE** | Pointer to [**Throughput**](Throughput.md) |  | [optional] 
**MaxNumberofPDUsessions** | Pointer to **int32** |  | [optional] 
**KPIMonitoringList** | Pointer to **string** |  | [optional] 
**SupportedAccessTechnology** | Pointer to **int32** |  | [optional] 
**AddServiceProfileInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceProfileChargingInformation

`func NewServiceProfileChargingInformation() *ServiceProfileChargingInformation`

NewServiceProfileChargingInformation instantiates a new ServiceProfileChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileChargingInformationWithDefaults

`func NewServiceProfileChargingInformationWithDefaults() *ServiceProfileChargingInformation`

NewServiceProfileChargingInformationWithDefaults instantiates a new ServiceProfileChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceProfileIdentifier

`func (o *ServiceProfileChargingInformation) GetServiceProfileIdentifier() string`

GetServiceProfileIdentifier returns the ServiceProfileIdentifier field if non-nil, zero value otherwise.

### GetServiceProfileIdentifierOk

`func (o *ServiceProfileChargingInformation) GetServiceProfileIdentifierOk() (*string, bool)`

GetServiceProfileIdentifierOk returns a tuple with the ServiceProfileIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfileIdentifier

`func (o *ServiceProfileChargingInformation) SetServiceProfileIdentifier(v string)`

SetServiceProfileIdentifier sets ServiceProfileIdentifier field to given value.

### HasServiceProfileIdentifier

`func (o *ServiceProfileChargingInformation) HasServiceProfileIdentifier() bool`

HasServiceProfileIdentifier returns a boolean if a field has been set.

### GetSNSSAIList

`func (o *ServiceProfileChargingInformation) GetSNSSAIList() []Snssai`

GetSNSSAIList returns the SNSSAIList field if non-nil, zero value otherwise.

### GetSNSSAIListOk

`func (o *ServiceProfileChargingInformation) GetSNSSAIListOk() (*[]Snssai, bool)`

GetSNSSAIListOk returns a tuple with the SNSSAIList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNSSAIList

`func (o *ServiceProfileChargingInformation) SetSNSSAIList(v []Snssai)`

SetSNSSAIList sets SNSSAIList field to given value.

### HasSNSSAIList

`func (o *ServiceProfileChargingInformation) HasSNSSAIList() bool`

HasSNSSAIList returns a boolean if a field has been set.

### GetLatency

`func (o *ServiceProfileChargingInformation) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *ServiceProfileChargingInformation) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *ServiceProfileChargingInformation) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *ServiceProfileChargingInformation) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetAvailability

`func (o *ServiceProfileChargingInformation) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ServiceProfileChargingInformation) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ServiceProfileChargingInformation) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ServiceProfileChargingInformation) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetJitter

`func (o *ServiceProfileChargingInformation) GetJitter() int32`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *ServiceProfileChargingInformation) GetJitterOk() (*int32, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *ServiceProfileChargingInformation) SetJitter(v int32)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *ServiceProfileChargingInformation) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetReliability

`func (o *ServiceProfileChargingInformation) GetReliability() string`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *ServiceProfileChargingInformation) GetReliabilityOk() (*string, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *ServiceProfileChargingInformation) SetReliability(v string)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *ServiceProfileChargingInformation) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetMaxNumberofUEs

`func (o *ServiceProfileChargingInformation) GetMaxNumberofUEs() int32`

GetMaxNumberofUEs returns the MaxNumberofUEs field if non-nil, zero value otherwise.

### GetMaxNumberofUEsOk

`func (o *ServiceProfileChargingInformation) GetMaxNumberofUEsOk() (*int32, bool)`

GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofUEs

`func (o *ServiceProfileChargingInformation) SetMaxNumberofUEs(v int32)`

SetMaxNumberofUEs sets MaxNumberofUEs field to given value.

### HasMaxNumberofUEs

`func (o *ServiceProfileChargingInformation) HasMaxNumberofUEs() bool`

HasMaxNumberofUEs returns a boolean if a field has been set.

### GetCoverageArea

`func (o *ServiceProfileChargingInformation) GetCoverageArea() string`

GetCoverageArea returns the CoverageArea field if non-nil, zero value otherwise.

### GetCoverageAreaOk

`func (o *ServiceProfileChargingInformation) GetCoverageAreaOk() (*string, bool)`

GetCoverageAreaOk returns a tuple with the CoverageArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageArea

`func (o *ServiceProfileChargingInformation) SetCoverageArea(v string)`

SetCoverageArea sets CoverageArea field to given value.

### HasCoverageArea

`func (o *ServiceProfileChargingInformation) HasCoverageArea() bool`

HasCoverageArea returns a boolean if a field has been set.

### GetDLThptPerSlice

`func (o *ServiceProfileChargingInformation) GetDLThptPerSlice() Throughput`

GetDLThptPerSlice returns the DLThptPerSlice field if non-nil, zero value otherwise.

### GetDLThptPerSliceOk

`func (o *ServiceProfileChargingInformation) GetDLThptPerSliceOk() (*Throughput, bool)`

GetDLThptPerSliceOk returns a tuple with the DLThptPerSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerSlice

`func (o *ServiceProfileChargingInformation) SetDLThptPerSlice(v Throughput)`

SetDLThptPerSlice sets DLThptPerSlice field to given value.

### HasDLThptPerSlice

`func (o *ServiceProfileChargingInformation) HasDLThptPerSlice() bool`

HasDLThptPerSlice returns a boolean if a field has been set.

### GetDLThptPerUE

`func (o *ServiceProfileChargingInformation) GetDLThptPerUE() Throughput`

GetDLThptPerUE returns the DLThptPerUE field if non-nil, zero value otherwise.

### GetDLThptPerUEOk

`func (o *ServiceProfileChargingInformation) GetDLThptPerUEOk() (*Throughput, bool)`

GetDLThptPerUEOk returns a tuple with the DLThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerUE

`func (o *ServiceProfileChargingInformation) SetDLThptPerUE(v Throughput)`

SetDLThptPerUE sets DLThptPerUE field to given value.

### HasDLThptPerUE

`func (o *ServiceProfileChargingInformation) HasDLThptPerUE() bool`

HasDLThptPerUE returns a boolean if a field has been set.

### GetULThptPerSlice

`func (o *ServiceProfileChargingInformation) GetULThptPerSlice() Throughput`

GetULThptPerSlice returns the ULThptPerSlice field if non-nil, zero value otherwise.

### GetULThptPerSliceOk

`func (o *ServiceProfileChargingInformation) GetULThptPerSliceOk() (*Throughput, bool)`

GetULThptPerSliceOk returns a tuple with the ULThptPerSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerSlice

`func (o *ServiceProfileChargingInformation) SetULThptPerSlice(v Throughput)`

SetULThptPerSlice sets ULThptPerSlice field to given value.

### HasULThptPerSlice

`func (o *ServiceProfileChargingInformation) HasULThptPerSlice() bool`

HasULThptPerSlice returns a boolean if a field has been set.

### GetULThptPerUE

`func (o *ServiceProfileChargingInformation) GetULThptPerUE() Throughput`

GetULThptPerUE returns the ULThptPerUE field if non-nil, zero value otherwise.

### GetULThptPerUEOk

`func (o *ServiceProfileChargingInformation) GetULThptPerUEOk() (*Throughput, bool)`

GetULThptPerUEOk returns a tuple with the ULThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerUE

`func (o *ServiceProfileChargingInformation) SetULThptPerUE(v Throughput)`

SetULThptPerUE sets ULThptPerUE field to given value.

### HasULThptPerUE

`func (o *ServiceProfileChargingInformation) HasULThptPerUE() bool`

HasULThptPerUE returns a boolean if a field has been set.

### GetMaxNumberofPDUsessions

`func (o *ServiceProfileChargingInformation) GetMaxNumberofPDUsessions() int32`

GetMaxNumberofPDUsessions returns the MaxNumberofPDUsessions field if non-nil, zero value otherwise.

### GetMaxNumberofPDUsessionsOk

`func (o *ServiceProfileChargingInformation) GetMaxNumberofPDUsessionsOk() (*int32, bool)`

GetMaxNumberofPDUsessionsOk returns a tuple with the MaxNumberofPDUsessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofPDUsessions

`func (o *ServiceProfileChargingInformation) SetMaxNumberofPDUsessions(v int32)`

SetMaxNumberofPDUsessions sets MaxNumberofPDUsessions field to given value.

### HasMaxNumberofPDUsessions

`func (o *ServiceProfileChargingInformation) HasMaxNumberofPDUsessions() bool`

HasMaxNumberofPDUsessions returns a boolean if a field has been set.

### GetKPIMonitoringList

`func (o *ServiceProfileChargingInformation) GetKPIMonitoringList() string`

GetKPIMonitoringList returns the KPIMonitoringList field if non-nil, zero value otherwise.

### GetKPIMonitoringListOk

`func (o *ServiceProfileChargingInformation) GetKPIMonitoringListOk() (*string, bool)`

GetKPIMonitoringListOk returns a tuple with the KPIMonitoringList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKPIMonitoringList

`func (o *ServiceProfileChargingInformation) SetKPIMonitoringList(v string)`

SetKPIMonitoringList sets KPIMonitoringList field to given value.

### HasKPIMonitoringList

`func (o *ServiceProfileChargingInformation) HasKPIMonitoringList() bool`

HasKPIMonitoringList returns a boolean if a field has been set.

### GetSupportedAccessTechnology

`func (o *ServiceProfileChargingInformation) GetSupportedAccessTechnology() int32`

GetSupportedAccessTechnology returns the SupportedAccessTechnology field if non-nil, zero value otherwise.

### GetSupportedAccessTechnologyOk

`func (o *ServiceProfileChargingInformation) GetSupportedAccessTechnologyOk() (*int32, bool)`

GetSupportedAccessTechnologyOk returns a tuple with the SupportedAccessTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAccessTechnology

`func (o *ServiceProfileChargingInformation) SetSupportedAccessTechnology(v int32)`

SetSupportedAccessTechnology sets SupportedAccessTechnology field to given value.

### HasSupportedAccessTechnology

`func (o *ServiceProfileChargingInformation) HasSupportedAccessTechnology() bool`

HasSupportedAccessTechnology returns a boolean if a field has been set.

### GetAddServiceProfileInfo

`func (o *ServiceProfileChargingInformation) GetAddServiceProfileInfo() string`

GetAddServiceProfileInfo returns the AddServiceProfileInfo field if non-nil, zero value otherwise.

### GetAddServiceProfileInfoOk

`func (o *ServiceProfileChargingInformation) GetAddServiceProfileInfoOk() (*string, bool)`

GetAddServiceProfileInfoOk returns a tuple with the AddServiceProfileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddServiceProfileInfo

`func (o *ServiceProfileChargingInformation) SetAddServiceProfileInfo(v string)`

SetAddServiceProfileInfo sets AddServiceProfileInfo field to given value.

### HasAddServiceProfileInfo

`func (o *ServiceProfileChargingInformation) HasAddServiceProfileInfo() bool`

HasAddServiceProfileInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


