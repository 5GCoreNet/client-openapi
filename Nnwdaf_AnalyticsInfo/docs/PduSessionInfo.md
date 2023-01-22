# PduSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N4SessId** | Pointer to **string** |  | [optional] 
**SessInactiveTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**PduSessStatus** | Pointer to [**PduSessionStatus**](PduSessionStatus.md) |  | [optional] 

## Methods

### NewPduSessionInfo

`func NewPduSessionInfo() *PduSessionInfo`

NewPduSessionInfo instantiates a new PduSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionInfoWithDefaults

`func NewPduSessionInfoWithDefaults() *PduSessionInfo`

NewPduSessionInfoWithDefaults instantiates a new PduSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN4SessId

`func (o *PduSessionInfo) GetN4SessId() string`

GetN4SessId returns the N4SessId field if non-nil, zero value otherwise.

### GetN4SessIdOk

`func (o *PduSessionInfo) GetN4SessIdOk() (*string, bool)`

GetN4SessIdOk returns a tuple with the N4SessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4SessId

`func (o *PduSessionInfo) SetN4SessId(v string)`

SetN4SessId sets N4SessId field to given value.

### HasN4SessId

`func (o *PduSessionInfo) HasN4SessId() bool`

HasN4SessId returns a boolean if a field has been set.

### GetSessInactiveTimer

`func (o *PduSessionInfo) GetSessInactiveTimer() int32`

GetSessInactiveTimer returns the SessInactiveTimer field if non-nil, zero value otherwise.

### GetSessInactiveTimerOk

`func (o *PduSessionInfo) GetSessInactiveTimerOk() (*int32, bool)`

GetSessInactiveTimerOk returns a tuple with the SessInactiveTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessInactiveTimer

`func (o *PduSessionInfo) SetSessInactiveTimer(v int32)`

SetSessInactiveTimer sets SessInactiveTimer field to given value.

### HasSessInactiveTimer

`func (o *PduSessionInfo) HasSessInactiveTimer() bool`

HasSessInactiveTimer returns a boolean if a field has been set.

### GetPduSessStatus

`func (o *PduSessionInfo) GetPduSessStatus() PduSessionStatus`

GetPduSessStatus returns the PduSessStatus field if non-nil, zero value otherwise.

### GetPduSessStatusOk

`func (o *PduSessionInfo) GetPduSessStatusOk() (*PduSessionStatus, bool)`

GetPduSessStatusOk returns a tuple with the PduSessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessStatus

`func (o *PduSessionInfo) SetPduSessStatus(v PduSessionStatus)`

SetPduSessStatus sets PduSessStatus field to given value.

### HasPduSessStatus

`func (o *PduSessionInfo) HasPduSessStatus() bool`

HasPduSessStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


