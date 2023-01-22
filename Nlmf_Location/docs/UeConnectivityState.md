# UeConnectivityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**AccessType**](AccessType.md) |  | 
**Connectivitystate** | Pointer to [**CmState**](CmState.md) |  | [optional] 

## Methods

### NewUeConnectivityState

`func NewUeConnectivityState(accessType AccessType, ) *UeConnectivityState`

NewUeConnectivityState instantiates a new UeConnectivityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeConnectivityStateWithDefaults

`func NewUeConnectivityStateWithDefaults() *UeConnectivityState`

NewUeConnectivityStateWithDefaults instantiates a new UeConnectivityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *UeConnectivityState) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UeConnectivityState) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UeConnectivityState) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetConnectivitystate

`func (o *UeConnectivityState) GetConnectivitystate() CmState`

GetConnectivitystate returns the Connectivitystate field if non-nil, zero value otherwise.

### GetConnectivitystateOk

`func (o *UeConnectivityState) GetConnectivitystateOk() (*CmState, bool)`

GetConnectivitystateOk returns a tuple with the Connectivitystate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivitystate

`func (o *UeConnectivityState) SetConnectivitystate(v CmState)`

SetConnectivitystate sets Connectivitystate field to given value.

### HasConnectivitystate

`func (o *UeConnectivityState) HasConnectivitystate() bool`

HasConnectivitystate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


