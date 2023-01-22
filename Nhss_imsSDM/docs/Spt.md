# Spt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionNegated** | **bool** |  | 
**SptGroup** | **[]int32** |  | 
**RegType** | Pointer to [**[]RegistrationType**](RegistrationType.md) |  | [optional] 
**RequestUri** | Pointer to **string** |  | [optional] 
**SipMethod** | Pointer to **string** |  | [optional] 
**SipHeader** | Pointer to [**HeaderSipRequest**](HeaderSipRequest.md) |  | [optional] 
**SessionCase** | Pointer to [**RequestDirection**](RequestDirection.md) |  | [optional] 
**SessionDescription** | Pointer to [**SdpDescription**](SdpDescription.md) |  | [optional] 

## Methods

### NewSpt

`func NewSpt(conditionNegated bool, sptGroup []int32, ) *Spt`

NewSpt instantiates a new Spt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSptWithDefaults

`func NewSptWithDefaults() *Spt`

NewSptWithDefaults instantiates a new Spt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionNegated

`func (o *Spt) GetConditionNegated() bool`

GetConditionNegated returns the ConditionNegated field if non-nil, zero value otherwise.

### GetConditionNegatedOk

`func (o *Spt) GetConditionNegatedOk() (*bool, bool)`

GetConditionNegatedOk returns a tuple with the ConditionNegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionNegated

`func (o *Spt) SetConditionNegated(v bool)`

SetConditionNegated sets ConditionNegated field to given value.


### GetSptGroup

`func (o *Spt) GetSptGroup() []int32`

GetSptGroup returns the SptGroup field if non-nil, zero value otherwise.

### GetSptGroupOk

`func (o *Spt) GetSptGroupOk() (*[]int32, bool)`

GetSptGroupOk returns a tuple with the SptGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSptGroup

`func (o *Spt) SetSptGroup(v []int32)`

SetSptGroup sets SptGroup field to given value.


### GetRegType

`func (o *Spt) GetRegType() []RegistrationType`

GetRegType returns the RegType field if non-nil, zero value otherwise.

### GetRegTypeOk

`func (o *Spt) GetRegTypeOk() (*[]RegistrationType, bool)`

GetRegTypeOk returns a tuple with the RegType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegType

`func (o *Spt) SetRegType(v []RegistrationType)`

SetRegType sets RegType field to given value.

### HasRegType

`func (o *Spt) HasRegType() bool`

HasRegType returns a boolean if a field has been set.

### GetRequestUri

`func (o *Spt) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *Spt) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *Spt) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *Spt) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetSipMethod

`func (o *Spt) GetSipMethod() string`

GetSipMethod returns the SipMethod field if non-nil, zero value otherwise.

### GetSipMethodOk

`func (o *Spt) GetSipMethodOk() (*string, bool)`

GetSipMethodOk returns a tuple with the SipMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipMethod

`func (o *Spt) SetSipMethod(v string)`

SetSipMethod sets SipMethod field to given value.

### HasSipMethod

`func (o *Spt) HasSipMethod() bool`

HasSipMethod returns a boolean if a field has been set.

### GetSipHeader

`func (o *Spt) GetSipHeader() HeaderSipRequest`

GetSipHeader returns the SipHeader field if non-nil, zero value otherwise.

### GetSipHeaderOk

`func (o *Spt) GetSipHeaderOk() (*HeaderSipRequest, bool)`

GetSipHeaderOk returns a tuple with the SipHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeader

`func (o *Spt) SetSipHeader(v HeaderSipRequest)`

SetSipHeader sets SipHeader field to given value.

### HasSipHeader

`func (o *Spt) HasSipHeader() bool`

HasSipHeader returns a boolean if a field has been set.

### GetSessionCase

`func (o *Spt) GetSessionCase() RequestDirection`

GetSessionCase returns the SessionCase field if non-nil, zero value otherwise.

### GetSessionCaseOk

`func (o *Spt) GetSessionCaseOk() (*RequestDirection, bool)`

GetSessionCaseOk returns a tuple with the SessionCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCase

`func (o *Spt) SetSessionCase(v RequestDirection)`

SetSessionCase sets SessionCase field to given value.

### HasSessionCase

`func (o *Spt) HasSessionCase() bool`

HasSessionCase returns a boolean if a field has been set.

### GetSessionDescription

`func (o *Spt) GetSessionDescription() SdpDescription`

GetSessionDescription returns the SessionDescription field if non-nil, zero value otherwise.

### GetSessionDescriptionOk

`func (o *Spt) GetSessionDescriptionOk() (*SdpDescription, bool)`

GetSessionDescriptionOk returns a tuple with the SessionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDescription

`func (o *Spt) SetSessionDescription(v SdpDescription)`

SetSessionDescription sets SessionDescription field to given value.

### HasSessionDescription

`func (o *Spt) HasSessionDescription() bool`

HasSessionDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


