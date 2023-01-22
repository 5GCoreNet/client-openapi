# BackhaulAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GnbId** | Pointer to **string** |  | [optional] 
**Tai** | Pointer to [**Tai**](Tai.md) |  | [optional] 

## Methods

### NewBackhaulAddress

`func NewBackhaulAddress() *BackhaulAddress`

NewBackhaulAddress instantiates a new BackhaulAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackhaulAddressWithDefaults

`func NewBackhaulAddressWithDefaults() *BackhaulAddress`

NewBackhaulAddressWithDefaults instantiates a new BackhaulAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGnbId

`func (o *BackhaulAddress) GetGnbId() string`

GetGnbId returns the GnbId field if non-nil, zero value otherwise.

### GetGnbIdOk

`func (o *BackhaulAddress) GetGnbIdOk() (*string, bool)`

GetGnbIdOk returns a tuple with the GnbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbId

`func (o *BackhaulAddress) SetGnbId(v string)`

SetGnbId sets GnbId field to given value.

### HasGnbId

`func (o *BackhaulAddress) HasGnbId() bool`

HasGnbId returns a boolean if a field has been set.

### GetTai

`func (o *BackhaulAddress) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *BackhaulAddress) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *BackhaulAddress) SetTai(v Tai)`

SetTai sets Tai field to given value.

### HasTai

`func (o *BackhaulAddress) HasTai() bool`

HasTai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


