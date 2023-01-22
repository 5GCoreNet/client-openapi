# TmgiRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsServiceIdStart** | **string** |  | 
**MbsServiceIdEnd** | **string** |  | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 

## Methods

### NewTmgiRange

`func NewTmgiRange(mbsServiceIdStart string, mbsServiceIdEnd string, plmnId PlmnId, ) *TmgiRange`

NewTmgiRange instantiates a new TmgiRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTmgiRangeWithDefaults

`func NewTmgiRangeWithDefaults() *TmgiRange`

NewTmgiRangeWithDefaults instantiates a new TmgiRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsServiceIdStart

`func (o *TmgiRange) GetMbsServiceIdStart() string`

GetMbsServiceIdStart returns the MbsServiceIdStart field if non-nil, zero value otherwise.

### GetMbsServiceIdStartOk

`func (o *TmgiRange) GetMbsServiceIdStartOk() (*string, bool)`

GetMbsServiceIdStartOk returns a tuple with the MbsServiceIdStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceIdStart

`func (o *TmgiRange) SetMbsServiceIdStart(v string)`

SetMbsServiceIdStart sets MbsServiceIdStart field to given value.


### GetMbsServiceIdEnd

`func (o *TmgiRange) GetMbsServiceIdEnd() string`

GetMbsServiceIdEnd returns the MbsServiceIdEnd field if non-nil, zero value otherwise.

### GetMbsServiceIdEndOk

`func (o *TmgiRange) GetMbsServiceIdEndOk() (*string, bool)`

GetMbsServiceIdEndOk returns a tuple with the MbsServiceIdEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceIdEnd

`func (o *TmgiRange) SetMbsServiceIdEnd(v string)`

SetMbsServiceIdEnd sets MbsServiceIdEnd field to given value.


### GetPlmnId

`func (o *TmgiRange) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *TmgiRange) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *TmgiRange) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetNid

`func (o *TmgiRange) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *TmgiRange) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *TmgiRange) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *TmgiRange) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


