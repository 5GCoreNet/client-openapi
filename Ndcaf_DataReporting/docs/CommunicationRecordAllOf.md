# CommunicationRecordAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeInterval** | [**TimeWindow**](TimeWindow.md) |  | 
**UplinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**DownlinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 

## Methods

### NewCommunicationRecordAllOf

`func NewCommunicationRecordAllOf(timeInterval TimeWindow, ) *CommunicationRecordAllOf`

NewCommunicationRecordAllOf instantiates a new CommunicationRecordAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationRecordAllOfWithDefaults

`func NewCommunicationRecordAllOfWithDefaults() *CommunicationRecordAllOf`

NewCommunicationRecordAllOfWithDefaults instantiates a new CommunicationRecordAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeInterval

`func (o *CommunicationRecordAllOf) GetTimeInterval() TimeWindow`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *CommunicationRecordAllOf) GetTimeIntervalOk() (*TimeWindow, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *CommunicationRecordAllOf) SetTimeInterval(v TimeWindow)`

SetTimeInterval sets TimeInterval field to given value.


### GetUplinkVolume

`func (o *CommunicationRecordAllOf) GetUplinkVolume() int64`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *CommunicationRecordAllOf) GetUplinkVolumeOk() (*int64, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *CommunicationRecordAllOf) SetUplinkVolume(v int64)`

SetUplinkVolume sets UplinkVolume field to given value.

### HasUplinkVolume

`func (o *CommunicationRecordAllOf) HasUplinkVolume() bool`

HasUplinkVolume returns a boolean if a field has been set.

### GetDownlinkVolume

`func (o *CommunicationRecordAllOf) GetDownlinkVolume() int64`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *CommunicationRecordAllOf) GetDownlinkVolumeOk() (*int64, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *CommunicationRecordAllOf) SetDownlinkVolume(v int64)`

SetDownlinkVolume sets DownlinkVolume field to given value.

### HasDownlinkVolume

`func (o *CommunicationRecordAllOf) HasDownlinkVolume() bool`

HasDownlinkVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


