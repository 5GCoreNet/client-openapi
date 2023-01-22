# DataReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalApplicationId** | **string** | String providing an application identifier. | 
**ServiceExperienceRecords** | Pointer to [**[]ServiceExperienceRecord**](ServiceExperienceRecord.md) |  | [optional] 
**LocationRecords** | Pointer to [**[]LocationRecord**](LocationRecord.md) |  | [optional] 
**CommunicationRecords** | Pointer to [**[]CommunicationRecord**](CommunicationRecord.md) |  | [optional] 
**PerformanceDataRecords** | Pointer to [**[]PerformanceDataRecord**](PerformanceDataRecord.md) |  | [optional] 
**ApplicationSpecificRecords** | Pointer to [**[]ApplicationSpecificRecord**](ApplicationSpecificRecord.md) |  | [optional] 
**TripPlanRecords** | Pointer to [**[]TripPlanRecord**](TripPlanRecord.md) |  | [optional] 
**MediaStreamingAccessRecords** | Pointer to [**[]MediaStreamingAccessRecord**](MediaStreamingAccessRecord.md) |  | [optional] 

## Methods

### NewDataReport

`func NewDataReport(externalApplicationId string, ) *DataReport`

NewDataReport instantiates a new DataReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataReportWithDefaults

`func NewDataReportWithDefaults() *DataReport`

NewDataReportWithDefaults instantiates a new DataReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalApplicationId

`func (o *DataReport) GetExternalApplicationId() string`

GetExternalApplicationId returns the ExternalApplicationId field if non-nil, zero value otherwise.

### GetExternalApplicationIdOk

`func (o *DataReport) GetExternalApplicationIdOk() (*string, bool)`

GetExternalApplicationIdOk returns a tuple with the ExternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApplicationId

`func (o *DataReport) SetExternalApplicationId(v string)`

SetExternalApplicationId sets ExternalApplicationId field to given value.


### GetServiceExperienceRecords

`func (o *DataReport) GetServiceExperienceRecords() []ServiceExperienceRecord`

GetServiceExperienceRecords returns the ServiceExperienceRecords field if non-nil, zero value otherwise.

### GetServiceExperienceRecordsOk

`func (o *DataReport) GetServiceExperienceRecordsOk() (*[]ServiceExperienceRecord, bool)`

GetServiceExperienceRecordsOk returns a tuple with the ServiceExperienceRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExperienceRecords

`func (o *DataReport) SetServiceExperienceRecords(v []ServiceExperienceRecord)`

SetServiceExperienceRecords sets ServiceExperienceRecords field to given value.

### HasServiceExperienceRecords

`func (o *DataReport) HasServiceExperienceRecords() bool`

HasServiceExperienceRecords returns a boolean if a field has been set.

### GetLocationRecords

`func (o *DataReport) GetLocationRecords() []LocationRecord`

GetLocationRecords returns the LocationRecords field if non-nil, zero value otherwise.

### GetLocationRecordsOk

`func (o *DataReport) GetLocationRecordsOk() (*[]LocationRecord, bool)`

GetLocationRecordsOk returns a tuple with the LocationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationRecords

`func (o *DataReport) SetLocationRecords(v []LocationRecord)`

SetLocationRecords sets LocationRecords field to given value.

### HasLocationRecords

`func (o *DataReport) HasLocationRecords() bool`

HasLocationRecords returns a boolean if a field has been set.

### GetCommunicationRecords

`func (o *DataReport) GetCommunicationRecords() []CommunicationRecord`

GetCommunicationRecords returns the CommunicationRecords field if non-nil, zero value otherwise.

### GetCommunicationRecordsOk

`func (o *DataReport) GetCommunicationRecordsOk() (*[]CommunicationRecord, bool)`

GetCommunicationRecordsOk returns a tuple with the CommunicationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationRecords

`func (o *DataReport) SetCommunicationRecords(v []CommunicationRecord)`

SetCommunicationRecords sets CommunicationRecords field to given value.

### HasCommunicationRecords

`func (o *DataReport) HasCommunicationRecords() bool`

HasCommunicationRecords returns a boolean if a field has been set.

### GetPerformanceDataRecords

`func (o *DataReport) GetPerformanceDataRecords() []PerformanceDataRecord`

GetPerformanceDataRecords returns the PerformanceDataRecords field if non-nil, zero value otherwise.

### GetPerformanceDataRecordsOk

`func (o *DataReport) GetPerformanceDataRecordsOk() (*[]PerformanceDataRecord, bool)`

GetPerformanceDataRecordsOk returns a tuple with the PerformanceDataRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceDataRecords

`func (o *DataReport) SetPerformanceDataRecords(v []PerformanceDataRecord)`

SetPerformanceDataRecords sets PerformanceDataRecords field to given value.

### HasPerformanceDataRecords

`func (o *DataReport) HasPerformanceDataRecords() bool`

HasPerformanceDataRecords returns a boolean if a field has been set.

### GetApplicationSpecificRecords

`func (o *DataReport) GetApplicationSpecificRecords() []ApplicationSpecificRecord`

GetApplicationSpecificRecords returns the ApplicationSpecificRecords field if non-nil, zero value otherwise.

### GetApplicationSpecificRecordsOk

`func (o *DataReport) GetApplicationSpecificRecordsOk() (*[]ApplicationSpecificRecord, bool)`

GetApplicationSpecificRecordsOk returns a tuple with the ApplicationSpecificRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSpecificRecords

`func (o *DataReport) SetApplicationSpecificRecords(v []ApplicationSpecificRecord)`

SetApplicationSpecificRecords sets ApplicationSpecificRecords field to given value.

### HasApplicationSpecificRecords

`func (o *DataReport) HasApplicationSpecificRecords() bool`

HasApplicationSpecificRecords returns a boolean if a field has been set.

### GetTripPlanRecords

`func (o *DataReport) GetTripPlanRecords() []TripPlanRecord`

GetTripPlanRecords returns the TripPlanRecords field if non-nil, zero value otherwise.

### GetTripPlanRecordsOk

`func (o *DataReport) GetTripPlanRecordsOk() (*[]TripPlanRecord, bool)`

GetTripPlanRecordsOk returns a tuple with the TripPlanRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripPlanRecords

`func (o *DataReport) SetTripPlanRecords(v []TripPlanRecord)`

SetTripPlanRecords sets TripPlanRecords field to given value.

### HasTripPlanRecords

`func (o *DataReport) HasTripPlanRecords() bool`

HasTripPlanRecords returns a boolean if a field has been set.

### GetMediaStreamingAccessRecords

`func (o *DataReport) GetMediaStreamingAccessRecords() []MediaStreamingAccessRecord`

GetMediaStreamingAccessRecords returns the MediaStreamingAccessRecords field if non-nil, zero value otherwise.

### GetMediaStreamingAccessRecordsOk

`func (o *DataReport) GetMediaStreamingAccessRecordsOk() (*[]MediaStreamingAccessRecord, bool)`

GetMediaStreamingAccessRecordsOk returns a tuple with the MediaStreamingAccessRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreamingAccessRecords

`func (o *DataReport) SetMediaStreamingAccessRecords(v []MediaStreamingAccessRecord)`

SetMediaStreamingAccessRecords sets MediaStreamingAccessRecords field to given value.

### HasMediaStreamingAccessRecords

`func (o *DataReport) HasMediaStreamingAccessRecords() bool`

HasMediaStreamingAccessRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


