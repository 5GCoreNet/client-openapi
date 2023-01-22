# UeMobilityExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**RecurringTime** | Pointer to [**ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | [optional] 
**Duration** | **int32** | Unsigned integer identifying a period of time in units of seconds. | 
**DurationVariance** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**LocInfo** | [**[]UeLocationInfo**](UeLocationInfo.md) |  | 

## Methods

### NewUeMobilityExposure

`func NewUeMobilityExposure(duration int32, locInfo []UeLocationInfo, ) *UeMobilityExposure`

NewUeMobilityExposure instantiates a new UeMobilityExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeMobilityExposureWithDefaults

`func NewUeMobilityExposureWithDefaults() *UeMobilityExposure`

NewUeMobilityExposureWithDefaults instantiates a new UeMobilityExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *UeMobilityExposure) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UeMobilityExposure) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UeMobilityExposure) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *UeMobilityExposure) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetRecurringTime

`func (o *UeMobilityExposure) GetRecurringTime() ScheduledCommunicationTime`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *UeMobilityExposure) GetRecurringTimeOk() (*ScheduledCommunicationTime, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *UeMobilityExposure) SetRecurringTime(v ScheduledCommunicationTime)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *UeMobilityExposure) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetDuration

`func (o *UeMobilityExposure) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UeMobilityExposure) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UeMobilityExposure) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetDurationVariance

`func (o *UeMobilityExposure) GetDurationVariance() float32`

GetDurationVariance returns the DurationVariance field if non-nil, zero value otherwise.

### GetDurationVarianceOk

`func (o *UeMobilityExposure) GetDurationVarianceOk() (*float32, bool)`

GetDurationVarianceOk returns a tuple with the DurationVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationVariance

`func (o *UeMobilityExposure) SetDurationVariance(v float32)`

SetDurationVariance sets DurationVariance field to given value.

### HasDurationVariance

`func (o *UeMobilityExposure) HasDurationVariance() bool`

HasDurationVariance returns a boolean if a field has been set.

### GetLocInfo

`func (o *UeMobilityExposure) GetLocInfo() []UeLocationInfo`

GetLocInfo returns the LocInfo field if non-nil, zero value otherwise.

### GetLocInfoOk

`func (o *UeMobilityExposure) GetLocInfoOk() (*[]UeLocationInfo, bool)`

GetLocInfoOk returns a tuple with the LocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfo

`func (o *UeMobilityExposure) SetLocInfo(v []UeLocationInfo)`

SetLocInfo sets LocInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


