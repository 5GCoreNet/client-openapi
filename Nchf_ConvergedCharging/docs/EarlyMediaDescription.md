# EarlyMediaDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SDPTimeStamps** | Pointer to [**SDPTimeStamps**](SDPTimeStamps.md) |  | [optional] 
**SDPMediaComponent** | Pointer to [**[]SDPMediaComponent**](SDPMediaComponent.md) |  | [optional] 
**SDPSessionDescription** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEarlyMediaDescription

`func NewEarlyMediaDescription() *EarlyMediaDescription`

NewEarlyMediaDescription instantiates a new EarlyMediaDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarlyMediaDescriptionWithDefaults

`func NewEarlyMediaDescriptionWithDefaults() *EarlyMediaDescription`

NewEarlyMediaDescriptionWithDefaults instantiates a new EarlyMediaDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSDPTimeStamps

`func (o *EarlyMediaDescription) GetSDPTimeStamps() SDPTimeStamps`

GetSDPTimeStamps returns the SDPTimeStamps field if non-nil, zero value otherwise.

### GetSDPTimeStampsOk

`func (o *EarlyMediaDescription) GetSDPTimeStampsOk() (*SDPTimeStamps, bool)`

GetSDPTimeStampsOk returns a tuple with the SDPTimeStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPTimeStamps

`func (o *EarlyMediaDescription) SetSDPTimeStamps(v SDPTimeStamps)`

SetSDPTimeStamps sets SDPTimeStamps field to given value.

### HasSDPTimeStamps

`func (o *EarlyMediaDescription) HasSDPTimeStamps() bool`

HasSDPTimeStamps returns a boolean if a field has been set.

### GetSDPMediaComponent

`func (o *EarlyMediaDescription) GetSDPMediaComponent() []SDPMediaComponent`

GetSDPMediaComponent returns the SDPMediaComponent field if non-nil, zero value otherwise.

### GetSDPMediaComponentOk

`func (o *EarlyMediaDescription) GetSDPMediaComponentOk() (*[]SDPMediaComponent, bool)`

GetSDPMediaComponentOk returns a tuple with the SDPMediaComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPMediaComponent

`func (o *EarlyMediaDescription) SetSDPMediaComponent(v []SDPMediaComponent)`

SetSDPMediaComponent sets SDPMediaComponent field to given value.

### HasSDPMediaComponent

`func (o *EarlyMediaDescription) HasSDPMediaComponent() bool`

HasSDPMediaComponent returns a boolean if a field has been set.

### GetSDPSessionDescription

`func (o *EarlyMediaDescription) GetSDPSessionDescription() []string`

GetSDPSessionDescription returns the SDPSessionDescription field if non-nil, zero value otherwise.

### GetSDPSessionDescriptionOk

`func (o *EarlyMediaDescription) GetSDPSessionDescriptionOk() (*[]string, bool)`

GetSDPSessionDescriptionOk returns a tuple with the SDPSessionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPSessionDescription

`func (o *EarlyMediaDescription) SetSDPSessionDescription(v []string)`

SetSDPSessionDescription sets SDPSessionDescription field to given value.

### HasSDPSessionDescription

`func (o *EarlyMediaDescription) HasSDPSessionDescription() bool`

HasSDPSessionDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


