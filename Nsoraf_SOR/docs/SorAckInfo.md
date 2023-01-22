# SorAckInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SorAckStatus** | [**SorAckStatus**](SorAckStatus.md) |  | 
**SorSendingTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**MeSupportOfSorCmci** | Pointer to **bool** |  | [optional] 

## Methods

### NewSorAckInfo

`func NewSorAckInfo(sorAckStatus SorAckStatus, sorSendingTime time.Time, ) *SorAckInfo`

NewSorAckInfo instantiates a new SorAckInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorAckInfoWithDefaults

`func NewSorAckInfoWithDefaults() *SorAckInfo`

NewSorAckInfoWithDefaults instantiates a new SorAckInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSorAckStatus

`func (o *SorAckInfo) GetSorAckStatus() SorAckStatus`

GetSorAckStatus returns the SorAckStatus field if non-nil, zero value otherwise.

### GetSorAckStatusOk

`func (o *SorAckInfo) GetSorAckStatusOk() (*SorAckStatus, bool)`

GetSorAckStatusOk returns a tuple with the SorAckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorAckStatus

`func (o *SorAckInfo) SetSorAckStatus(v SorAckStatus)`

SetSorAckStatus sets SorAckStatus field to given value.


### GetSorSendingTime

`func (o *SorAckInfo) GetSorSendingTime() time.Time`

GetSorSendingTime returns the SorSendingTime field if non-nil, zero value otherwise.

### GetSorSendingTimeOk

`func (o *SorAckInfo) GetSorSendingTimeOk() (*time.Time, bool)`

GetSorSendingTimeOk returns a tuple with the SorSendingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorSendingTime

`func (o *SorAckInfo) SetSorSendingTime(v time.Time)`

SetSorSendingTime sets SorSendingTime field to given value.


### GetMeSupportOfSorCmci

`func (o *SorAckInfo) GetMeSupportOfSorCmci() bool`

GetMeSupportOfSorCmci returns the MeSupportOfSorCmci field if non-nil, zero value otherwise.

### GetMeSupportOfSorCmciOk

`func (o *SorAckInfo) GetMeSupportOfSorCmciOk() (*bool, bool)`

GetMeSupportOfSorCmciOk returns a tuple with the MeSupportOfSorCmci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeSupportOfSorCmci

`func (o *SorAckInfo) SetMeSupportOfSorCmci(v bool)`

SetMeSupportOfSorCmci sets MeSupportOfSorCmci field to given value.

### HasMeSupportOfSorCmci

`func (o *SorAckInfo) HasMeSupportOfSorCmci() bool`

HasMeSupportOfSorCmci returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


