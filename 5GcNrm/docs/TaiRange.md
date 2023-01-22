# TaiRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**NRTACRangelist** | Pointer to [**[]NRTACRange**](NRTACRange.md) |  | [optional] 

## Methods

### NewTaiRange

`func NewTaiRange() *TaiRange`

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

### HasPlmnId

`func (o *TaiRange) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetNRTACRangelist

`func (o *TaiRange) GetNRTACRangelist() []NRTACRange`

GetNRTACRangelist returns the NRTACRangelist field if non-nil, zero value otherwise.

### GetNRTACRangelistOk

`func (o *TaiRange) GetNRTACRangelistOk() (*[]NRTACRange, bool)`

GetNRTACRangelistOk returns a tuple with the NRTACRangelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRTACRangelist

`func (o *TaiRange) SetNRTACRangelist(v []NRTACRange)`

SetNRTACRangelist sets NRTACRangelist field to given value.

### HasNRTACRangelist

`func (o *TaiRange) HasNRTACRangelist() bool`

HasNRTACRangelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


