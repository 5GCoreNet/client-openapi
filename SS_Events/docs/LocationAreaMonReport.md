# LocationAreaMonReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurPreUEs** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of identities of all VAL UEs present in the given location area. | [optional] 
**MoveInOutUEs** | Pointer to [**MoveInOutUEDetails**](MoveInOutUEDetails.md) |  | [optional] 
**TrigEvnt** | Pointer to [**MonLocTriggerEvent**](MonLocTriggerEvent.md) |  | [optional] 

## Methods

### NewLocationAreaMonReport

`func NewLocationAreaMonReport() *LocationAreaMonReport`

NewLocationAreaMonReport instantiates a new LocationAreaMonReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAreaMonReportWithDefaults

`func NewLocationAreaMonReportWithDefaults() *LocationAreaMonReport`

NewLocationAreaMonReportWithDefaults instantiates a new LocationAreaMonReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurPreUEs

`func (o *LocationAreaMonReport) GetCurPreUEs() []ValTargetUe`

GetCurPreUEs returns the CurPreUEs field if non-nil, zero value otherwise.

### GetCurPreUEsOk

`func (o *LocationAreaMonReport) GetCurPreUEsOk() (*[]ValTargetUe, bool)`

GetCurPreUEsOk returns a tuple with the CurPreUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurPreUEs

`func (o *LocationAreaMonReport) SetCurPreUEs(v []ValTargetUe)`

SetCurPreUEs sets CurPreUEs field to given value.

### HasCurPreUEs

`func (o *LocationAreaMonReport) HasCurPreUEs() bool`

HasCurPreUEs returns a boolean if a field has been set.

### GetMoveInOutUEs

`func (o *LocationAreaMonReport) GetMoveInOutUEs() MoveInOutUEDetails`

GetMoveInOutUEs returns the MoveInOutUEs field if non-nil, zero value otherwise.

### GetMoveInOutUEsOk

`func (o *LocationAreaMonReport) GetMoveInOutUEsOk() (*MoveInOutUEDetails, bool)`

GetMoveInOutUEsOk returns a tuple with the MoveInOutUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveInOutUEs

`func (o *LocationAreaMonReport) SetMoveInOutUEs(v MoveInOutUEDetails)`

SetMoveInOutUEs sets MoveInOutUEs field to given value.

### HasMoveInOutUEs

`func (o *LocationAreaMonReport) HasMoveInOutUEs() bool`

HasMoveInOutUEs returns a boolean if a field has been set.

### GetTrigEvnt

`func (o *LocationAreaMonReport) GetTrigEvnt() MonLocTriggerEvent`

GetTrigEvnt returns the TrigEvnt field if non-nil, zero value otherwise.

### GetTrigEvntOk

`func (o *LocationAreaMonReport) GetTrigEvntOk() (*MonLocTriggerEvent, bool)`

GetTrigEvntOk returns a tuple with the TrigEvnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigEvnt

`func (o *LocationAreaMonReport) SetTrigEvnt(v MonLocTriggerEvent)`

SetTrigEvnt sets TrigEvnt field to given value.

### HasTrigEvnt

`func (o *LocationAreaMonReport) HasTrigEvnt() bool`

HasTrigEvnt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


