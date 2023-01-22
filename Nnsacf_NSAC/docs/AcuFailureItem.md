# AcuFailureItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Reason** | Pointer to [**AcuFailureReason**](AcuFailureReason.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**PduSessionId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 

## Methods

### NewAcuFailureItem

`func NewAcuFailureItem(snssai Snssai, ) *AcuFailureItem`

NewAcuFailureItem instantiates a new AcuFailureItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcuFailureItemWithDefaults

`func NewAcuFailureItemWithDefaults() *AcuFailureItem`

NewAcuFailureItemWithDefaults instantiates a new AcuFailureItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *AcuFailureItem) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AcuFailureItem) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AcuFailureItem) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetReason

`func (o *AcuFailureItem) GetReason() AcuFailureReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AcuFailureItem) GetReasonOk() (*AcuFailureReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AcuFailureItem) SetReason(v AcuFailureReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AcuFailureItem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPlmnId

`func (o *AcuFailureItem) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *AcuFailureItem) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *AcuFailureItem) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *AcuFailureItem) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetPduSessionId

`func (o *AcuFailureItem) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *AcuFailureItem) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *AcuFailureItem) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *AcuFailureItem) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


