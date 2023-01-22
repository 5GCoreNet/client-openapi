# MbsSessionId1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tmgi** | Pointer to [**Tmgi1**](Tmgi1.md) |  | [optional] 
**Ssm** | Pointer to [**Ssm1**](Ssm1.md) |  | [optional] 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 

## Methods

### NewMbsSessionId1

`func NewMbsSessionId1() *MbsSessionId1`

NewMbsSessionId1 instantiates a new MbsSessionId1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionId1WithDefaults

`func NewMbsSessionId1WithDefaults() *MbsSessionId1`

NewMbsSessionId1WithDefaults instantiates a new MbsSessionId1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmgi

`func (o *MbsSessionId1) GetTmgi() Tmgi1`

GetTmgi returns the Tmgi field if non-nil, zero value otherwise.

### GetTmgiOk

`func (o *MbsSessionId1) GetTmgiOk() (*Tmgi1, bool)`

GetTmgiOk returns a tuple with the Tmgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgi

`func (o *MbsSessionId1) SetTmgi(v Tmgi1)`

SetTmgi sets Tmgi field to given value.

### HasTmgi

`func (o *MbsSessionId1) HasTmgi() bool`

HasTmgi returns a boolean if a field has been set.

### GetSsm

`func (o *MbsSessionId1) GetSsm() Ssm1`

GetSsm returns the Ssm field if non-nil, zero value otherwise.

### GetSsmOk

`func (o *MbsSessionId1) GetSsmOk() (*Ssm1, bool)`

GetSsmOk returns a tuple with the Ssm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsm

`func (o *MbsSessionId1) SetSsm(v Ssm1)`

SetSsm sets Ssm field to given value.

### HasSsm

`func (o *MbsSessionId1) HasSsm() bool`

HasSsm returns a boolean if a field has been set.

### GetNid

`func (o *MbsSessionId1) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *MbsSessionId1) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *MbsSessionId1) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *MbsSessionId1) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


