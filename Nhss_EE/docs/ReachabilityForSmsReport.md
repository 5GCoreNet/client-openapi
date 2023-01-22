# ReachabilityForSmsReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReachabilitySmsStatus** | **bool** |  | 
**MaxAvailabilityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewReachabilityForSmsReport

`func NewReachabilityForSmsReport(reachabilitySmsStatus bool, ) *ReachabilityForSmsReport`

NewReachabilityForSmsReport instantiates a new ReachabilityForSmsReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachabilityForSmsReportWithDefaults

`func NewReachabilityForSmsReportWithDefaults() *ReachabilityForSmsReport`

NewReachabilityForSmsReportWithDefaults instantiates a new ReachabilityForSmsReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachabilitySmsStatus

`func (o *ReachabilityForSmsReport) GetReachabilitySmsStatus() bool`

GetReachabilitySmsStatus returns the ReachabilitySmsStatus field if non-nil, zero value otherwise.

### GetReachabilitySmsStatusOk

`func (o *ReachabilityForSmsReport) GetReachabilitySmsStatusOk() (*bool, bool)`

GetReachabilitySmsStatusOk returns a tuple with the ReachabilitySmsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilitySmsStatus

`func (o *ReachabilityForSmsReport) SetReachabilitySmsStatus(v bool)`

SetReachabilitySmsStatus sets ReachabilitySmsStatus field to given value.


### GetMaxAvailabilityTime

`func (o *ReachabilityForSmsReport) GetMaxAvailabilityTime() time.Time`

GetMaxAvailabilityTime returns the MaxAvailabilityTime field if non-nil, zero value otherwise.

### GetMaxAvailabilityTimeOk

`func (o *ReachabilityForSmsReport) GetMaxAvailabilityTimeOk() (*time.Time, bool)`

GetMaxAvailabilityTimeOk returns a tuple with the MaxAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAvailabilityTime

`func (o *ReachabilityForSmsReport) SetMaxAvailabilityTime(v time.Time)`

SetMaxAvailabilityTime sets MaxAvailabilityTime field to given value.

### HasMaxAvailabilityTime

`func (o *ReachabilityForSmsReport) HasMaxAvailabilityTime() bool`

HasMaxAvailabilityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


