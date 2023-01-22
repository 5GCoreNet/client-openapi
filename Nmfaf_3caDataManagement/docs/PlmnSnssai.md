# PlmnSnssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**SNssaiList** | [**[]ExtSnssai**](ExtSnssai.md) |  | 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 

## Methods

### NewPlmnSnssai

`func NewPlmnSnssai(plmnId PlmnId, sNssaiList []ExtSnssai, ) *PlmnSnssai`

NewPlmnSnssai instantiates a new PlmnSnssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlmnSnssaiWithDefaults

`func NewPlmnSnssaiWithDefaults() *PlmnSnssai`

NewPlmnSnssaiWithDefaults instantiates a new PlmnSnssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *PlmnSnssai) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PlmnSnssai) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PlmnSnssai) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSNssaiList

`func (o *PlmnSnssai) GetSNssaiList() []ExtSnssai`

GetSNssaiList returns the SNssaiList field if non-nil, zero value otherwise.

### GetSNssaiListOk

`func (o *PlmnSnssai) GetSNssaiListOk() (*[]ExtSnssai, bool)`

GetSNssaiListOk returns a tuple with the SNssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiList

`func (o *PlmnSnssai) SetSNssaiList(v []ExtSnssai)`

SetSNssaiList sets SNssaiList field to given value.


### GetNid

`func (o *PlmnSnssai) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *PlmnSnssai) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *PlmnSnssai) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *PlmnSnssai) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


