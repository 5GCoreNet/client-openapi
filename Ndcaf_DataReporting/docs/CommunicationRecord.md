# CommunicationRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**TimeInterval** | [**TimeWindow**](TimeWindow.md) |  | 
**UplinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**DownlinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 

## Methods

### NewCommunicationRecord

`func NewCommunicationRecord(timestamp time.Time, timeInterval TimeWindow, ) *CommunicationRecord`

NewCommunicationRecord instantiates a new CommunicationRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationRecordWithDefaults

`func NewCommunicationRecordWithDefaults() *CommunicationRecord`

NewCommunicationRecordWithDefaults instantiates a new CommunicationRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *CommunicationRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CommunicationRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CommunicationRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTimeInterval

`func (o *CommunicationRecord) GetTimeInterval() TimeWindow`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *CommunicationRecord) GetTimeIntervalOk() (*TimeWindow, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *CommunicationRecord) SetTimeInterval(v TimeWindow)`

SetTimeInterval sets TimeInterval field to given value.


### GetUplinkVolume

`func (o *CommunicationRecord) GetUplinkVolume() int64`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *CommunicationRecord) GetUplinkVolumeOk() (*int64, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *CommunicationRecord) SetUplinkVolume(v int64)`

SetUplinkVolume sets UplinkVolume field to given value.

### HasUplinkVolume

`func (o *CommunicationRecord) HasUplinkVolume() bool`

HasUplinkVolume returns a boolean if a field has been set.

### GetDownlinkVolume

`func (o *CommunicationRecord) GetDownlinkVolume() int64`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *CommunicationRecord) GetDownlinkVolumeOk() (*int64, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *CommunicationRecord) SetDownlinkVolume(v int64)`

SetDownlinkVolume sets DownlinkVolume field to given value.

### HasDownlinkVolume

`func (o *CommunicationRecord) HasDownlinkVolume() bool`

HasDownlinkVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


