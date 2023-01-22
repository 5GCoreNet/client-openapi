# MoveInOutUEDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveInUEs** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of identities of VAL UEs who moved in to given location area since previous notification.  | [optional] 
**MoveOutUEs** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of identities of VAL UEs who moved out of the given location area since previous notification.  | [optional] 

## Methods

### NewMoveInOutUEDetails

`func NewMoveInOutUEDetails() *MoveInOutUEDetails`

NewMoveInOutUEDetails instantiates a new MoveInOutUEDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveInOutUEDetailsWithDefaults

`func NewMoveInOutUEDetailsWithDefaults() *MoveInOutUEDetails`

NewMoveInOutUEDetailsWithDefaults instantiates a new MoveInOutUEDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveInUEs

`func (o *MoveInOutUEDetails) GetMoveInUEs() []ValTargetUe`

GetMoveInUEs returns the MoveInUEs field if non-nil, zero value otherwise.

### GetMoveInUEsOk

`func (o *MoveInOutUEDetails) GetMoveInUEsOk() (*[]ValTargetUe, bool)`

GetMoveInUEsOk returns a tuple with the MoveInUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveInUEs

`func (o *MoveInOutUEDetails) SetMoveInUEs(v []ValTargetUe)`

SetMoveInUEs sets MoveInUEs field to given value.

### HasMoveInUEs

`func (o *MoveInOutUEDetails) HasMoveInUEs() bool`

HasMoveInUEs returns a boolean if a field has been set.

### GetMoveOutUEs

`func (o *MoveInOutUEDetails) GetMoveOutUEs() []ValTargetUe`

GetMoveOutUEs returns the MoveOutUEs field if non-nil, zero value otherwise.

### GetMoveOutUEsOk

`func (o *MoveInOutUEDetails) GetMoveOutUEsOk() (*[]ValTargetUe, bool)`

GetMoveOutUEsOk returns a tuple with the MoveOutUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveOutUEs

`func (o *MoveInOutUEDetails) SetMoveOutUEs(v []ValTargetUe)`

SetMoveOutUEs sets MoveOutUEs field to given value.

### HasMoveOutUEs

`func (o *MoveInOutUEDetails) HasMoveOutUEs() bool`

HasMoveOutUEs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


