# ReachabilityForDataReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReachabilityDataStatus** | **bool** |  | 
**MaxAvailabilityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewReachabilityForDataReport

`func NewReachabilityForDataReport(reachabilityDataStatus bool, ) *ReachabilityForDataReport`

NewReachabilityForDataReport instantiates a new ReachabilityForDataReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachabilityForDataReportWithDefaults

`func NewReachabilityForDataReportWithDefaults() *ReachabilityForDataReport`

NewReachabilityForDataReportWithDefaults instantiates a new ReachabilityForDataReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachabilityDataStatus

`func (o *ReachabilityForDataReport) GetReachabilityDataStatus() bool`

GetReachabilityDataStatus returns the ReachabilityDataStatus field if non-nil, zero value otherwise.

### GetReachabilityDataStatusOk

`func (o *ReachabilityForDataReport) GetReachabilityDataStatusOk() (*bool, bool)`

GetReachabilityDataStatusOk returns a tuple with the ReachabilityDataStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityDataStatus

`func (o *ReachabilityForDataReport) SetReachabilityDataStatus(v bool)`

SetReachabilityDataStatus sets ReachabilityDataStatus field to given value.


### GetMaxAvailabilityTime

`func (o *ReachabilityForDataReport) GetMaxAvailabilityTime() time.Time`

GetMaxAvailabilityTime returns the MaxAvailabilityTime field if non-nil, zero value otherwise.

### GetMaxAvailabilityTimeOk

`func (o *ReachabilityForDataReport) GetMaxAvailabilityTimeOk() (*time.Time, bool)`

GetMaxAvailabilityTimeOk returns a tuple with the MaxAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailabilityTime

`func (o *ReachabilityForDataReport) SetMaxAvailabilityTime(v time.Time)`

SetMaxAvailabilityTime sets MaxAvailabilityTime field to given value.

### HasMaxAvailabilityTime

`func (o *ReachabilityForDataReport) HasMaxAvailabilityTime() bool`

HasMaxAvailabilityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


