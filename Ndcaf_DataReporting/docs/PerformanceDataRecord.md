# PerformanceDataRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**TimeInterval** | [**TimeWindow**](TimeWindow.md) |  | 
**Location** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**RemoteEndpoint** | Pointer to [**AddrFqdn**](AddrFqdn.md) |  | [optional] 
**PacketDelayBudget** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**PacketLossRate** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 
**UplinkThroughput** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**DownlinkThrougput** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewPerformanceDataRecord

`func NewPerformanceDataRecord(timestamp time.Time, timeInterval TimeWindow, ) *PerformanceDataRecord`

NewPerformanceDataRecord instantiates a new PerformanceDataRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceDataRecordWithDefaults

`func NewPerformanceDataRecordWithDefaults() *PerformanceDataRecord`

NewPerformanceDataRecordWithDefaults instantiates a new PerformanceDataRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *PerformanceDataRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PerformanceDataRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PerformanceDataRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTimeInterval

`func (o *PerformanceDataRecord) GetTimeInterval() TimeWindow`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *PerformanceDataRecord) GetTimeIntervalOk() (*TimeWindow, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *PerformanceDataRecord) SetTimeInterval(v TimeWindow)`

SetTimeInterval sets TimeInterval field to given value.


### GetLocation

`func (o *PerformanceDataRecord) GetLocation() LocationArea5G`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PerformanceDataRecord) GetLocationOk() (*LocationArea5G, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PerformanceDataRecord) SetLocation(v LocationArea5G)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PerformanceDataRecord) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRemoteEndpoint

`func (o *PerformanceDataRecord) GetRemoteEndpoint() AddrFqdn`

GetRemoteEndpoint returns the RemoteEndpoint field if non-nil, zero value otherwise.

### GetRemoteEndpointOk

`func (o *PerformanceDataRecord) GetRemoteEndpointOk() (*AddrFqdn, bool)`

GetRemoteEndpointOk returns a tuple with the RemoteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndpoint

`func (o *PerformanceDataRecord) SetRemoteEndpoint(v AddrFqdn)`

SetRemoteEndpoint sets RemoteEndpoint field to given value.

### HasRemoteEndpoint

`func (o *PerformanceDataRecord) HasRemoteEndpoint() bool`

HasRemoteEndpoint returns a boolean if a field has been set.

### GetPacketDelayBudget

`func (o *PerformanceDataRecord) GetPacketDelayBudget() int32`

GetPacketDelayBudget returns the PacketDelayBudget field if non-nil, zero value otherwise.

### GetPacketDelayBudgetOk

`func (o *PerformanceDataRecord) GetPacketDelayBudgetOk() (*int32, bool)`

GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDelayBudget

`func (o *PerformanceDataRecord) SetPacketDelayBudget(v int32)`

SetPacketDelayBudget sets PacketDelayBudget field to given value.

### HasPacketDelayBudget

`func (o *PerformanceDataRecord) HasPacketDelayBudget() bool`

HasPacketDelayBudget returns a boolean if a field has been set.

### GetPacketLossRate

`func (o *PerformanceDataRecord) GetPacketLossRate() int32`

GetPacketLossRate returns the PacketLossRate field if non-nil, zero value otherwise.

### GetPacketLossRateOk

`func (o *PerformanceDataRecord) GetPacketLossRateOk() (*int32, bool)`

GetPacketLossRateOk returns a tuple with the PacketLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLossRate

`func (o *PerformanceDataRecord) SetPacketLossRate(v int32)`

SetPacketLossRate sets PacketLossRate field to given value.

### HasPacketLossRate

`func (o *PerformanceDataRecord) HasPacketLossRate() bool`

HasPacketLossRate returns a boolean if a field has been set.

### GetUplinkThroughput

`func (o *PerformanceDataRecord) GetUplinkThroughput() string`

GetUplinkThroughput returns the UplinkThroughput field if non-nil, zero value otherwise.

### GetUplinkThroughputOk

`func (o *PerformanceDataRecord) GetUplinkThroughputOk() (*string, bool)`

GetUplinkThroughputOk returns a tuple with the UplinkThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkThroughput

`func (o *PerformanceDataRecord) SetUplinkThroughput(v string)`

SetUplinkThroughput sets UplinkThroughput field to given value.

### HasUplinkThroughput

`func (o *PerformanceDataRecord) HasUplinkThroughput() bool`

HasUplinkThroughput returns a boolean if a field has been set.

### GetDownlinkThrougput

`func (o *PerformanceDataRecord) GetDownlinkThrougput() string`

GetDownlinkThrougput returns the DownlinkThrougput field if non-nil, zero value otherwise.

### GetDownlinkThrougputOk

`func (o *PerformanceDataRecord) GetDownlinkThrougputOk() (*string, bool)`

GetDownlinkThrougputOk returns a tuple with the DownlinkThrougput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkThrougput

`func (o *PerformanceDataRecord) SetDownlinkThrougput(v string)`

SetDownlinkThrougput sets DownlinkThrougput field to given value.

### HasDownlinkThrougput

`func (o *PerformanceDataRecord) HasDownlinkThrougput() bool`

HasDownlinkThrougput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


