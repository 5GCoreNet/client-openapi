# ProseKeyReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayServCode** | **int32** | Relay Service Code to identify a connectivity service provided by the UE-to-Network relay.  | 
**KnrpFreshness1** | **string** | KNRP Freshness Parameter 1 | 
**ResyncInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 
**PrukId** | Pointer to **string** | Prose Remote User Key ID | [optional] 
**Suci** | Pointer to **string** | Contains the SUCI. | [optional] 

## Methods

### NewProseKeyReqData

`func NewProseKeyReqData(relayServCode int32, knrpFreshness1 string, ) *ProseKeyReqData`

NewProseKeyReqData instantiates a new ProseKeyReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProseKeyReqDataWithDefaults

`func NewProseKeyReqDataWithDefaults() *ProseKeyReqData`

NewProseKeyReqDataWithDefaults instantiates a new ProseKeyReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayServCode

`func (o *ProseKeyReqData) GetRelayServCode() int32`

GetRelayServCode returns the RelayServCode field if non-nil, zero value otherwise.

### GetRelayServCodeOk

`func (o *ProseKeyReqData) GetRelayServCodeOk() (*int32, bool)`

GetRelayServCodeOk returns a tuple with the RelayServCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServCode

`func (o *ProseKeyReqData) SetRelayServCode(v int32)`

SetRelayServCode sets RelayServCode field to given value.


### GetKnrpFreshness1

`func (o *ProseKeyReqData) GetKnrpFreshness1() string`

GetKnrpFreshness1 returns the KnrpFreshness1 field if non-nil, zero value otherwise.

### GetKnrpFreshness1Ok

`func (o *ProseKeyReqData) GetKnrpFreshness1Ok() (*string, bool)`

GetKnrpFreshness1Ok returns a tuple with the KnrpFreshness1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnrpFreshness1

`func (o *ProseKeyReqData) SetKnrpFreshness1(v string)`

SetKnrpFreshness1 sets KnrpFreshness1 field to given value.


### GetResyncInfo

`func (o *ProseKeyReqData) GetResyncInfo() ResynchronizationInfo`

GetResyncInfo returns the ResyncInfo field if non-nil, zero value otherwise.

### GetResyncInfoOk

`func (o *ProseKeyReqData) GetResyncInfoOk() (*ResynchronizationInfo, bool)`

GetResyncInfoOk returns a tuple with the ResyncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncInfo

`func (o *ProseKeyReqData) SetResyncInfo(v ResynchronizationInfo)`

SetResyncInfo sets ResyncInfo field to given value.

### HasResyncInfo

`func (o *ProseKeyReqData) HasResyncInfo() bool`

HasResyncInfo returns a boolean if a field has been set.

### GetPrukId

`func (o *ProseKeyReqData) GetPrukId() string`

GetPrukId returns the PrukId field if non-nil, zero value otherwise.

### GetPrukIdOk

`func (o *ProseKeyReqData) GetPrukIdOk() (*string, bool)`

GetPrukIdOk returns a tuple with the PrukId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrukId

`func (o *ProseKeyReqData) SetPrukId(v string)`

SetPrukId sets PrukId field to given value.

### HasPrukId

`func (o *ProseKeyReqData) HasPrukId() bool`

HasPrukId returns a boolean if a field has been set.

### GetSuci

`func (o *ProseKeyReqData) GetSuci() string`

GetSuci returns the Suci field if non-nil, zero value otherwise.

### GetSuciOk

`func (o *ProseKeyReqData) GetSuciOk() (*string, bool)`

GetSuciOk returns a tuple with the Suci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuci

`func (o *ProseKeyReqData) SetSuci(v string)`

SetSuci sets Suci field to given value.

### HasSuci

`func (o *ProseKeyReqData) HasSuci() bool`

HasSuci returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


