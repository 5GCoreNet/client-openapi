# SessInactTimerForUeComm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N4SessId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**SessInactiveTimer** | **int32** | indicating a time in seconds. | 

## Methods

### NewSessInactTimerForUeComm

`func NewSessInactTimerForUeComm(n4SessId int32, sessInactiveTimer int32, ) *SessInactTimerForUeComm`

NewSessInactTimerForUeComm instantiates a new SessInactTimerForUeComm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessInactTimerForUeCommWithDefaults

`func NewSessInactTimerForUeCommWithDefaults() *SessInactTimerForUeComm`

NewSessInactTimerForUeCommWithDefaults instantiates a new SessInactTimerForUeComm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN4SessId

`func (o *SessInactTimerForUeComm) GetN4SessId() int32`

GetN4SessId returns the N4SessId field if non-nil, zero value otherwise.

### GetN4SessIdOk

`func (o *SessInactTimerForUeComm) GetN4SessIdOk() (*int32, bool)`

GetN4SessIdOk returns a tuple with the N4SessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN4SessId

`func (o *SessInactTimerForUeComm) SetN4SessId(v int32)`

SetN4SessId sets N4SessId field to given value.


### GetSessInactiveTimer

`func (o *SessInactTimerForUeComm) GetSessInactiveTimer() int32`

GetSessInactiveTimer returns the SessInactiveTimer field if non-nil, zero value otherwise.

### GetSessInactiveTimerOk

`func (o *SessInactTimerForUeComm) GetSessInactiveTimerOk() (*int32, bool)`

GetSessInactiveTimerOk returns a tuple with the SessInactiveTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessInactiveTimer

`func (o *SessInactTimerForUeComm) SetSessInactiveTimer(v int32)`

SetSessInactiveTimer sets SessInactiveTimer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


