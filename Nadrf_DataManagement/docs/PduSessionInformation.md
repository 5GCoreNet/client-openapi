# PduSessionInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**SessInfo** | Pointer to [**PduSessionInfo**](PduSessionInfo.md) |  | [optional] 

## Methods

### NewPduSessionInformation

`func NewPduSessionInformation() *PduSessionInformation`

NewPduSessionInformation instantiates a new PduSessionInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionInformationWithDefaults

`func NewPduSessionInformationWithDefaults() *PduSessionInformation`

NewPduSessionInformationWithDefaults instantiates a new PduSessionInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessId

`func (o *PduSessionInformation) GetPduSessId() int32`

GetPduSessId returns the PduSessId field if non-nil, zero value otherwise.

### GetPduSessIdOk

`func (o *PduSessionInformation) GetPduSessIdOk() (*int32, bool)`

GetPduSessIdOk returns a tuple with the PduSessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessId

`func (o *PduSessionInformation) SetPduSessId(v int32)`

SetPduSessId sets PduSessId field to given value.

### HasPduSessId

`func (o *PduSessionInformation) HasPduSessId() bool`

HasPduSessId returns a boolean if a field has been set.

### GetSessInfo

`func (o *PduSessionInformation) GetSessInfo() PduSessionInfo`

GetSessInfo returns the SessInfo field if non-nil, zero value otherwise.

### GetSessInfoOk

`func (o *PduSessionInformation) GetSessInfoOk() (*PduSessionInfo, bool)`

GetSessInfoOk returns a tuple with the SessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessInfo

`func (o *PduSessionInformation) SetSessInfo(v PduSessionInfo)`

SetSessInfo sets SessInfo field to given value.

### HasSessInfo

`func (o *PduSessionInformation) HasSessInfo() bool`

HasSessInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


