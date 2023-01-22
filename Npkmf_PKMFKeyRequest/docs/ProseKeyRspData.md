# ProseKeyRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Knrp** | **string** | Key for NR PC5 | 
**KnrpFreshness2** | **string** | KNRP Freshness Parameter 2 | 
**Gpi** | Pointer to **string** | GBA Pushing Information | [optional] 

## Methods

### NewProseKeyRspData

`func NewProseKeyRspData(knrp string, knrpFreshness2 string, ) *ProseKeyRspData`

NewProseKeyRspData instantiates a new ProseKeyRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProseKeyRspDataWithDefaults

`func NewProseKeyRspDataWithDefaults() *ProseKeyRspData`

NewProseKeyRspDataWithDefaults instantiates a new ProseKeyRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnrp

`func (o *ProseKeyRspData) GetKnrp() string`

GetKnrp returns the Knrp field if non-nil, zero value otherwise.

### GetKnrpOk

`func (o *ProseKeyRspData) GetKnrpOk() (*string, bool)`

GetKnrpOk returns a tuple with the Knrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnrp

`func (o *ProseKeyRspData) SetKnrp(v string)`

SetKnrp sets Knrp field to given value.


### GetKnrpFreshness2

`func (o *ProseKeyRspData) GetKnrpFreshness2() string`

GetKnrpFreshness2 returns the KnrpFreshness2 field if non-nil, zero value otherwise.

### GetKnrpFreshness2Ok

`func (o *ProseKeyRspData) GetKnrpFreshness2Ok() (*string, bool)`

GetKnrpFreshness2Ok returns a tuple with the KnrpFreshness2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnrpFreshness2

`func (o *ProseKeyRspData) SetKnrpFreshness2(v string)`

SetKnrpFreshness2 sets KnrpFreshness2 field to given value.


### GetGpi

`func (o *ProseKeyRspData) GetGpi() string`

GetGpi returns the Gpi field if non-nil, zero value otherwise.

### GetGpiOk

`func (o *ProseKeyRspData) GetGpiOk() (*string, bool)`

GetGpiOk returns a tuple with the Gpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpi

`func (o *ProseKeyRspData) SetGpi(v string)`

SetGpi sets Gpi field to given value.

### HasGpi

`func (o *ProseKeyRspData) HasGpi() bool`

HasGpi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


