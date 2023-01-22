# ScpDomainCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScpDomains** | **[]string** |  | 
**NfTypeList** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 

## Methods

### NewScpDomainCond

`func NewScpDomainCond(scpDomains []string, ) *ScpDomainCond`

NewScpDomainCond instantiates a new ScpDomainCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScpDomainCondWithDefaults

`func NewScpDomainCondWithDefaults() *ScpDomainCond`

NewScpDomainCondWithDefaults instantiates a new ScpDomainCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScpDomains

`func (o *ScpDomainCond) GetScpDomains() []string`

GetScpDomains returns the ScpDomains field if non-nil, zero value otherwise.

### GetScpDomainsOk

`func (o *ScpDomainCond) GetScpDomainsOk() (*[]string, bool)`

GetScpDomainsOk returns a tuple with the ScpDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpDomains

`func (o *ScpDomainCond) SetScpDomains(v []string)`

SetScpDomains sets ScpDomains field to given value.


### GetNfTypeList

`func (o *ScpDomainCond) GetNfTypeList() []NFType`

GetNfTypeList returns the NfTypeList field if non-nil, zero value otherwise.

### GetNfTypeListOk

`func (o *ScpDomainCond) GetNfTypeListOk() (*[]NFType, bool)`

GetNfTypeListOk returns a tuple with the NfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypeList

`func (o *ScpDomainCond) SetNfTypeList(v []NFType)`

SetNfTypeList sets NfTypeList field to given value.

### HasNfTypeList

`func (o *ScpDomainCond) HasNfTypeList() bool`

HasNfTypeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


