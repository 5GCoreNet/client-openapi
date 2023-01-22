# MbsKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyDomainId** | **string** | string with format &#39;bytes&#39; as defined in OpenAPI | 
**MskId** | **string** | string with format &#39;bytes&#39; as defined in OpenAPI | 
**Msk** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**MskLifetime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**MtkId** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**Mtk** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 

## Methods

### NewMbsKeyInfo

`func NewMbsKeyInfo(keyDomainId string, mskId string, ) *MbsKeyInfo`

NewMbsKeyInfo instantiates a new MbsKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsKeyInfoWithDefaults

`func NewMbsKeyInfoWithDefaults() *MbsKeyInfo`

NewMbsKeyInfoWithDefaults instantiates a new MbsKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyDomainId

`func (o *MbsKeyInfo) GetKeyDomainId() string`

GetKeyDomainId returns the KeyDomainId field if non-nil, zero value otherwise.

### GetKeyDomainIdOk

`func (o *MbsKeyInfo) GetKeyDomainIdOk() (*string, bool)`

GetKeyDomainIdOk returns a tuple with the KeyDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDomainId

`func (o *MbsKeyInfo) SetKeyDomainId(v string)`

SetKeyDomainId sets KeyDomainId field to given value.


### GetMskId

`func (o *MbsKeyInfo) GetMskId() string`

GetMskId returns the MskId field if non-nil, zero value otherwise.

### GetMskIdOk

`func (o *MbsKeyInfo) GetMskIdOk() (*string, bool)`

GetMskIdOk returns a tuple with the MskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMskId

`func (o *MbsKeyInfo) SetMskId(v string)`

SetMskId sets MskId field to given value.


### GetMsk

`func (o *MbsKeyInfo) GetMsk() string`

GetMsk returns the Msk field if non-nil, zero value otherwise.

### GetMskOk

`func (o *MbsKeyInfo) GetMskOk() (*string, bool)`

GetMskOk returns a tuple with the Msk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsk

`func (o *MbsKeyInfo) SetMsk(v string)`

SetMsk sets Msk field to given value.

### HasMsk

`func (o *MbsKeyInfo) HasMsk() bool`

HasMsk returns a boolean if a field has been set.

### GetMskLifetime

`func (o *MbsKeyInfo) GetMskLifetime() time.Time`

GetMskLifetime returns the MskLifetime field if non-nil, zero value otherwise.

### GetMskLifetimeOk

`func (o *MbsKeyInfo) GetMskLifetimeOk() (*time.Time, bool)`

GetMskLifetimeOk returns a tuple with the MskLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMskLifetime

`func (o *MbsKeyInfo) SetMskLifetime(v time.Time)`

SetMskLifetime sets MskLifetime field to given value.

### HasMskLifetime

`func (o *MbsKeyInfo) HasMskLifetime() bool`

HasMskLifetime returns a boolean if a field has been set.

### GetMtkId

`func (o *MbsKeyInfo) GetMtkId() string`

GetMtkId returns the MtkId field if non-nil, zero value otherwise.

### GetMtkIdOk

`func (o *MbsKeyInfo) GetMtkIdOk() (*string, bool)`

GetMtkIdOk returns a tuple with the MtkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtkId

`func (o *MbsKeyInfo) SetMtkId(v string)`

SetMtkId sets MtkId field to given value.

### HasMtkId

`func (o *MbsKeyInfo) HasMtkId() bool`

HasMtkId returns a boolean if a field has been set.

### GetMtk

`func (o *MbsKeyInfo) GetMtk() string`

GetMtk returns the Mtk field if non-nil, zero value otherwise.

### GetMtkOk

`func (o *MbsKeyInfo) GetMtkOk() (*string, bool)`

GetMtkOk returns a tuple with the Mtk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtk

`func (o *MbsKeyInfo) SetMtk(v string)`

SetMtk sets Mtk field to given value.

### HasMtk

`func (o *MbsKeyInfo) HasMtk() bool`

HasMtk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


