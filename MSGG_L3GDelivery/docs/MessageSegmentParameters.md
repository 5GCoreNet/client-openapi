# MessageSegmentParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegId** | Pointer to **string** |  | [optional] 
**TotalSegCount** | Pointer to **int32** |  | [optional] 
**SegNumb** | Pointer to **int32** |  | [optional] 
**LastSegFlag** | Pointer to **bool** |  | [optional] 

## Methods

### NewMessageSegmentParameters

`func NewMessageSegmentParameters() *MessageSegmentParameters`

NewMessageSegmentParameters instantiates a new MessageSegmentParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSegmentParametersWithDefaults

`func NewMessageSegmentParametersWithDefaults() *MessageSegmentParameters`

NewMessageSegmentParametersWithDefaults instantiates a new MessageSegmentParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegId

`func (o *MessageSegmentParameters) GetSegId() string`

GetSegId returns the SegId field if non-nil, zero value otherwise.

### GetSegIdOk

`func (o *MessageSegmentParameters) GetSegIdOk() (*string, bool)`

GetSegIdOk returns a tuple with the SegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegId

`func (o *MessageSegmentParameters) SetSegId(v string)`

SetSegId sets SegId field to given value.

### HasSegId

`func (o *MessageSegmentParameters) HasSegId() bool`

HasSegId returns a boolean if a field has been set.

### GetTotalSegCount

`func (o *MessageSegmentParameters) GetTotalSegCount() int32`

GetTotalSegCount returns the TotalSegCount field if non-nil, zero value otherwise.

### GetTotalSegCountOk

`func (o *MessageSegmentParameters) GetTotalSegCountOk() (*int32, bool)`

GetTotalSegCountOk returns a tuple with the TotalSegCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSegCount

`func (o *MessageSegmentParameters) SetTotalSegCount(v int32)`

SetTotalSegCount sets TotalSegCount field to given value.

### HasTotalSegCount

`func (o *MessageSegmentParameters) HasTotalSegCount() bool`

HasTotalSegCount returns a boolean if a field has been set.

### GetSegNumb

`func (o *MessageSegmentParameters) GetSegNumb() int32`

GetSegNumb returns the SegNumb field if non-nil, zero value otherwise.

### GetSegNumbOk

`func (o *MessageSegmentParameters) GetSegNumbOk() (*int32, bool)`

GetSegNumbOk returns a tuple with the SegNumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegNumb

`func (o *MessageSegmentParameters) SetSegNumb(v int32)`

SetSegNumb sets SegNumb field to given value.

### HasSegNumb

`func (o *MessageSegmentParameters) HasSegNumb() bool`

HasSegNumb returns a boolean if a field has been set.

### GetLastSegFlag

`func (o *MessageSegmentParameters) GetLastSegFlag() bool`

GetLastSegFlag returns the LastSegFlag field if non-nil, zero value otherwise.

### GetLastSegFlagOk

`func (o *MessageSegmentParameters) GetLastSegFlagOk() (*bool, bool)`

GetLastSegFlagOk returns a tuple with the LastSegFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSegFlag

`func (o *MessageSegmentParameters) SetLastSegFlag(v bool)`

SetLastSegFlag sets LastSegFlag field to given value.

### HasLastSegFlag

`func (o *MessageSegmentParameters) HasLastSegFlag() bool`

HasLastSegFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


