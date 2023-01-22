# UeTrajectoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**Location** | [**UserLocation**](UserLocation.md) |  | 

## Methods

### NewUeTrajectoryInfo

`func NewUeTrajectoryInfo(ts time.Time, location UserLocation, ) *UeTrajectoryInfo`

NewUeTrajectoryInfo instantiates a new UeTrajectoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeTrajectoryInfoWithDefaults

`func NewUeTrajectoryInfoWithDefaults() *UeTrajectoryInfo`

NewUeTrajectoryInfoWithDefaults instantiates a new UeTrajectoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *UeTrajectoryInfo) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UeTrajectoryInfo) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UeTrajectoryInfo) SetTs(v time.Time)`

SetTs sets Ts field to given value.


### GetLocation

`func (o *UeTrajectoryInfo) GetLocation() UserLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UeTrajectoryInfo) GetLocationOk() (*UserLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UeTrajectoryInfo) SetLocation(v UserLocation)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


