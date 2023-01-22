# RedundantPduSessionInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rsn** | [**Rsn**](Rsn.md) |  | 
**PduSessionPairId** | Pointer to **int32** |  | [optional] 

## Methods

### NewRedundantPduSessionInformation

`func NewRedundantPduSessionInformation(rsn Rsn, ) *RedundantPduSessionInformation`

NewRedundantPduSessionInformation instantiates a new RedundantPduSessionInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedundantPduSessionInformationWithDefaults

`func NewRedundantPduSessionInformationWithDefaults() *RedundantPduSessionInformation`

NewRedundantPduSessionInformationWithDefaults instantiates a new RedundantPduSessionInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRsn

`func (o *RedundantPduSessionInformation) GetRsn() Rsn`

GetRsn returns the Rsn field if non-nil, zero value otherwise.

### GetRsnOk

`func (o *RedundantPduSessionInformation) GetRsnOk() (*Rsn, bool)`

GetRsnOk returns a tuple with the Rsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsn

`func (o *RedundantPduSessionInformation) SetRsn(v Rsn)`

SetRsn sets Rsn field to given value.


### GetPduSessionPairId

`func (o *RedundantPduSessionInformation) GetPduSessionPairId() int32`

GetPduSessionPairId returns the PduSessionPairId field if non-nil, zero value otherwise.

### GetPduSessionPairIdOk

`func (o *RedundantPduSessionInformation) GetPduSessionPairIdOk() (*int32, bool)`

GetPduSessionPairIdOk returns a tuple with the PduSessionPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionPairId

`func (o *RedundantPduSessionInformation) SetPduSessionPairId(v int32)`

SetPduSessionPairId sets PduSessionPairId field to given value.

### HasPduSessionPairId

`func (o *RedundantPduSessionInformation) HasPduSessionPairId() bool`

HasPduSessionPairId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


