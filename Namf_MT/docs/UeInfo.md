# UeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeList** | **[]string** |  | 
**PduSessionId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 

## Methods

### NewUeInfo

`func NewUeInfo(ueList []string, ) *UeInfo`

NewUeInfo instantiates a new UeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeInfoWithDefaults

`func NewUeInfoWithDefaults() *UeInfo`

NewUeInfoWithDefaults instantiates a new UeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeList

`func (o *UeInfo) GetUeList() []string`

GetUeList returns the UeList field if non-nil, zero value otherwise.

### GetUeListOk

`func (o *UeInfo) GetUeListOk() (*[]string, bool)`

GetUeListOk returns a tuple with the UeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeList

`func (o *UeInfo) SetUeList(v []string)`

SetUeList sets UeList field to given value.


### GetPduSessionId

`func (o *UeInfo) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *UeInfo) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *UeInfo) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *UeInfo) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


