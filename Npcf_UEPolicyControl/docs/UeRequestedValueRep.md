# UeRequestedValueRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**PraStatuses** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | Contains the UE presence statuses for tracking areas. The praId attribute within the PresenceInfo data type is the key of the map.  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ConnectState** | Pointer to [**CmState**](CmState.md) |  | [optional] 

## Methods

### NewUeRequestedValueRep

`func NewUeRequestedValueRep() *UeRequestedValueRep`

NewUeRequestedValueRep instantiates a new UeRequestedValueRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeRequestedValueRepWithDefaults

`func NewUeRequestedValueRepWithDefaults() *UeRequestedValueRep`

NewUeRequestedValueRepWithDefaults instantiates a new UeRequestedValueRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLoc

`func (o *UeRequestedValueRep) GetUserLoc() UserLocation`

GetUserLoc returns the UserLoc field if non-nil, zero value otherwise.

### GetUserLocOk

`func (o *UeRequestedValueRep) GetUserLocOk() (*UserLocation, bool)`

GetUserLocOk returns a tuple with the UserLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoc

`func (o *UeRequestedValueRep) SetUserLoc(v UserLocation)`

SetUserLoc sets UserLoc field to given value.

### HasUserLoc

`func (o *UeRequestedValueRep) HasUserLoc() bool`

HasUserLoc returns a boolean if a field has been set.

### GetPraStatuses

`func (o *UeRequestedValueRep) GetPraStatuses() map[string]PresenceInfo`

GetPraStatuses returns the PraStatuses field if non-nil, zero value otherwise.

### GetPraStatusesOk

`func (o *UeRequestedValueRep) GetPraStatusesOk() (*map[string]PresenceInfo, bool)`

GetPraStatusesOk returns a tuple with the PraStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraStatuses

`func (o *UeRequestedValueRep) SetPraStatuses(v map[string]PresenceInfo)`

SetPraStatuses sets PraStatuses field to given value.

### HasPraStatuses

`func (o *UeRequestedValueRep) HasPraStatuses() bool`

HasPraStatuses returns a boolean if a field has been set.

### GetPlmnId

`func (o *UeRequestedValueRep) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *UeRequestedValueRep) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *UeRequestedValueRep) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *UeRequestedValueRep) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetConnectState

`func (o *UeRequestedValueRep) GetConnectState() CmState`

GetConnectState returns the ConnectState field if non-nil, zero value otherwise.

### GetConnectStateOk

`func (o *UeRequestedValueRep) GetConnectStateOk() (*CmState, bool)`

GetConnectStateOk returns a tuple with the ConnectState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectState

`func (o *UeRequestedValueRep) SetConnectState(v CmState)`

SetConnectState sets ConnectState field to given value.

### HasConnectState

`func (o *UeRequestedValueRep) HasConnectState() bool`

HasConnectState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


