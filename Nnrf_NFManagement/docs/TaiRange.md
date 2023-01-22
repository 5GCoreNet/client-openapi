# TaiRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**TacRangeList** | [**[]TacRange**](TacRange.md) |  | 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 

## Methods

### NewTaiRange

`func NewTaiRange(plmnId PlmnId, tacRangeList []TacRange, ) *TaiRange`

NewTaiRange instantiates a new TaiRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaiRangeWithDefaults

`func NewTaiRangeWithDefaults() *TaiRange`

NewTaiRangeWithDefaults instantiates a new TaiRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *TaiRange) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *TaiRange) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *TaiRange) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetTacRangeList

`func (o *TaiRange) GetTacRangeList() []TacRange`

GetTacRangeList returns the TacRangeList field if non-nil, zero value otherwise.

### GetTacRangeListOk

`func (o *TaiRange) GetTacRangeListOk() (*[]TacRange, bool)`

GetTacRangeListOk returns a tuple with the TacRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacRangeList

`func (o *TaiRange) SetTacRangeList(v []TacRange)`

SetTacRangeList sets TacRangeList field to given value.


### GetNid

`func (o *TaiRange) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *TaiRange) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *TaiRange) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *TaiRange) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


