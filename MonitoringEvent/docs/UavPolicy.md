# UavPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UavMoveInd** | **bool** |  | 
**RevokeInd** | **bool** |  | 

## Methods

### NewUavPolicy

`func NewUavPolicy(uavMoveInd bool, revokeInd bool, ) *UavPolicy`

NewUavPolicy instantiates a new UavPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUavPolicyWithDefaults

`func NewUavPolicyWithDefaults() *UavPolicy`

NewUavPolicyWithDefaults instantiates a new UavPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUavMoveInd

`func (o *UavPolicy) GetUavMoveInd() bool`

GetUavMoveInd returns the UavMoveInd field if non-nil, zero value otherwise.

### GetUavMoveIndOk

`func (o *UavPolicy) GetUavMoveIndOk() (*bool, bool)`

GetUavMoveIndOk returns a tuple with the UavMoveInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavMoveInd

`func (o *UavPolicy) SetUavMoveInd(v bool)`

SetUavMoveInd sets UavMoveInd field to given value.


### GetRevokeInd

`func (o *UavPolicy) GetRevokeInd() bool`

GetRevokeInd returns the RevokeInd field if non-nil, zero value otherwise.

### GetRevokeIndOk

`func (o *UavPolicy) GetRevokeIndOk() (*bool, bool)`

GetRevokeIndOk returns a tuple with the RevokeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeInd

`func (o *UavPolicy) SetRevokeInd(v bool)`

SetRevokeInd sets RevokeInd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


