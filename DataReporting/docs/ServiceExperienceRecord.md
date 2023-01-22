# ServiceExperienceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**ServiceExperienceInfos** | [**PerFlowServiceExperienceInfo**](PerFlowServiceExperienceInfo.md) |  | 

## Methods

### NewServiceExperienceRecord

`func NewServiceExperienceRecord(timestamp time.Time, serviceExperienceInfos PerFlowServiceExperienceInfo, ) *ServiceExperienceRecord`

NewServiceExperienceRecord instantiates a new ServiceExperienceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceExperienceRecordWithDefaults

`func NewServiceExperienceRecordWithDefaults() *ServiceExperienceRecord`

NewServiceExperienceRecordWithDefaults instantiates a new ServiceExperienceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ServiceExperienceRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceExperienceRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceExperienceRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetServiceExperienceInfos

`func (o *ServiceExperienceRecord) GetServiceExperienceInfos() PerFlowServiceExperienceInfo`

GetServiceExperienceInfos returns the ServiceExperienceInfos field if non-nil, zero value otherwise.

### GetServiceExperienceInfosOk

`func (o *ServiceExperienceRecord) GetServiceExperienceInfosOk() (*PerFlowServiceExperienceInfo, bool)`

GetServiceExperienceInfosOk returns a tuple with the ServiceExperienceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExperienceInfos

`func (o *ServiceExperienceRecord) SetServiceExperienceInfos(v PerFlowServiceExperienceInfo)`

SetServiceExperienceInfos sets ServiceExperienceInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


