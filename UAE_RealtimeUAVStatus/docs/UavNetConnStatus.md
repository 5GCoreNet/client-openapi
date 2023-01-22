# UavNetConnStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusInfo** | [**MonitoringType**](MonitoringType.md) |  | 
**Timestamp** | **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 

## Methods

### NewUavNetConnStatus

`func NewUavNetConnStatus(statusInfo MonitoringType, timestamp time.Time, ) *UavNetConnStatus`

NewUavNetConnStatus instantiates a new UavNetConnStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUavNetConnStatusWithDefaults

`func NewUavNetConnStatusWithDefaults() *UavNetConnStatus`

NewUavNetConnStatusWithDefaults instantiates a new UavNetConnStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusInfo

`func (o *UavNetConnStatus) GetStatusInfo() MonitoringType`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *UavNetConnStatus) GetStatusInfoOk() (*MonitoringType, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *UavNetConnStatus) SetStatusInfo(v MonitoringType)`

SetStatusInfo sets StatusInfo field to given value.


### GetTimestamp

`func (o *UavNetConnStatus) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UavNetConnStatus) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UavNetConnStatus) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


