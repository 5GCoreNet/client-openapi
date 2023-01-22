# NFProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NFInstanceId** | Pointer to **string** | uuid of NF instance | [optional] 
**NFType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**NFStatus** | Pointer to [**NFStatus**](NFStatus.md) |  | [optional] 
**Plmn** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**SNssais** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**InterPlmnFqdn** | Pointer to **string** |  | [optional] 
**NfServices** | Pointer to [**[]NFService**](NFService.md) |  | [optional] 

## Methods

### NewNFProfile

`func NewNFProfile() *NFProfile`

NewNFProfile instantiates a new NFProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFProfileWithDefaults

`func NewNFProfileWithDefaults() *NFProfile`

NewNFProfileWithDefaults instantiates a new NFProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNFInstanceId

`func (o *NFProfile) GetNFInstanceId() string`

GetNFInstanceId returns the NFInstanceId field if non-nil, zero value otherwise.

### GetNFInstanceIdOk

`func (o *NFProfile) GetNFInstanceIdOk() (*string, bool)`

GetNFInstanceIdOk returns a tuple with the NFInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFInstanceId

`func (o *NFProfile) SetNFInstanceId(v string)`

SetNFInstanceId sets NFInstanceId field to given value.

### HasNFInstanceId

`func (o *NFProfile) HasNFInstanceId() bool`

HasNFInstanceId returns a boolean if a field has been set.

### GetNFType

`func (o *NFProfile) GetNFType() NFType`

GetNFType returns the NFType field if non-nil, zero value otherwise.

### GetNFTypeOk

`func (o *NFProfile) GetNFTypeOk() (*NFType, bool)`

GetNFTypeOk returns a tuple with the NFType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFType

`func (o *NFProfile) SetNFType(v NFType)`

SetNFType sets NFType field to given value.

### HasNFType

`func (o *NFProfile) HasNFType() bool`

HasNFType returns a boolean if a field has been set.

### GetNFStatus

`func (o *NFProfile) GetNFStatus() NFStatus`

GetNFStatus returns the NFStatus field if non-nil, zero value otherwise.

### GetNFStatusOk

`func (o *NFProfile) GetNFStatusOk() (*NFStatus, bool)`

GetNFStatusOk returns a tuple with the NFStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFStatus

`func (o *NFProfile) SetNFStatus(v NFStatus)`

SetNFStatus sets NFStatus field to given value.

### HasNFStatus

`func (o *NFProfile) HasNFStatus() bool`

HasNFStatus returns a boolean if a field has been set.

### GetPlmn

`func (o *NFProfile) GetPlmn() PlmnId`

GetPlmn returns the Plmn field if non-nil, zero value otherwise.

### GetPlmnOk

`func (o *NFProfile) GetPlmnOk() (*PlmnId, bool)`

GetPlmnOk returns a tuple with the Plmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmn

`func (o *NFProfile) SetPlmn(v PlmnId)`

SetPlmn sets Plmn field to given value.

### HasPlmn

`func (o *NFProfile) HasPlmn() bool`

HasPlmn returns a boolean if a field has been set.

### GetSNssais

`func (o *NFProfile) GetSNssais() Snssai`

GetSNssais returns the SNssais field if non-nil, zero value otherwise.

### GetSNssaisOk

`func (o *NFProfile) GetSNssaisOk() (*Snssai, bool)`

GetSNssaisOk returns a tuple with the SNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssais

`func (o *NFProfile) SetSNssais(v Snssai)`

SetSNssais sets SNssais field to given value.

### HasSNssais

`func (o *NFProfile) HasSNssais() bool`

HasSNssais returns a boolean if a field has been set.

### GetFqdn

`func (o *NFProfile) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NFProfile) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NFProfile) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NFProfile) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInterPlmnFqdn

`func (o *NFProfile) GetInterPlmnFqdn() string`

GetInterPlmnFqdn returns the InterPlmnFqdn field if non-nil, zero value otherwise.

### GetInterPlmnFqdnOk

`func (o *NFProfile) GetInterPlmnFqdnOk() (*string, bool)`

GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnFqdn

`func (o *NFProfile) SetInterPlmnFqdn(v string)`

SetInterPlmnFqdn sets InterPlmnFqdn field to given value.

### HasInterPlmnFqdn

`func (o *NFProfile) HasInterPlmnFqdn() bool`

HasInterPlmnFqdn returns a boolean if a field has been set.

### GetNfServices

`func (o *NFProfile) GetNfServices() []NFService`

GetNfServices returns the NfServices field if non-nil, zero value otherwise.

### GetNfServicesOk

`func (o *NFProfile) GetNfServicesOk() (*[]NFService, bool)`

GetNfServicesOk returns a tuple with the NfServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServices

`func (o *NFProfile) SetNfServices(v []NFService)`

SetNfServices sets NfServices field to given value.

### HasNfServices

`func (o *NFProfile) HasNfServices() bool`

HasNfServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


