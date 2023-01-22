# NrOperatorCellDuSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**CellLocalId** | Pointer to **int32** |  | [optional] 
**AdministrativeState** | Pointer to [**AdministrativeState**](AdministrativeState.md) |  | [optional] 
**PlmnInfoList** | Pointer to [**[]PlmnInfo**](PlmnInfo.md) |  | [optional] 
**NrTac** | Pointer to **int32** |  | [optional] 

## Methods

### NewNrOperatorCellDuSingle

`func NewNrOperatorCellDuSingle(id NullableString, ) *NrOperatorCellDuSingle`

NewNrOperatorCellDuSingle instantiates a new NrOperatorCellDuSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrOperatorCellDuSingleWithDefaults

`func NewNrOperatorCellDuSingleWithDefaults() *NrOperatorCellDuSingle`

NewNrOperatorCellDuSingleWithDefaults instantiates a new NrOperatorCellDuSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NrOperatorCellDuSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NrOperatorCellDuSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NrOperatorCellDuSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *NrOperatorCellDuSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NrOperatorCellDuSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *NrOperatorCellDuSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *NrOperatorCellDuSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *NrOperatorCellDuSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *NrOperatorCellDuSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *NrOperatorCellDuSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *NrOperatorCellDuSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *NrOperatorCellDuSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *NrOperatorCellDuSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *NrOperatorCellDuSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *NrOperatorCellDuSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *NrOperatorCellDuSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *NrOperatorCellDuSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetCellLocalId

`func (o *NrOperatorCellDuSingle) GetCellLocalId() int32`

GetCellLocalId returns the CellLocalId field if non-nil, zero value otherwise.

### GetCellLocalIdOk

`func (o *NrOperatorCellDuSingle) GetCellLocalIdOk() (*int32, bool)`

GetCellLocalIdOk returns a tuple with the CellLocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellLocalId

`func (o *NrOperatorCellDuSingle) SetCellLocalId(v int32)`

SetCellLocalId sets CellLocalId field to given value.

### HasCellLocalId

`func (o *NrOperatorCellDuSingle) HasCellLocalId() bool`

HasCellLocalId returns a boolean if a field has been set.

### GetAdministrativeState

`func (o *NrOperatorCellDuSingle) GetAdministrativeState() AdministrativeState`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *NrOperatorCellDuSingle) GetAdministrativeStateOk() (*AdministrativeState, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *NrOperatorCellDuSingle) SetAdministrativeState(v AdministrativeState)`

SetAdministrativeState sets AdministrativeState field to given value.

### HasAdministrativeState

`func (o *NrOperatorCellDuSingle) HasAdministrativeState() bool`

HasAdministrativeState returns a boolean if a field has been set.

### GetPlmnInfoList

`func (o *NrOperatorCellDuSingle) GetPlmnInfoList() []PlmnInfo`

GetPlmnInfoList returns the PlmnInfoList field if non-nil, zero value otherwise.

### GetPlmnInfoListOk

`func (o *NrOperatorCellDuSingle) GetPlmnInfoListOk() (*[]PlmnInfo, bool)`

GetPlmnInfoListOk returns a tuple with the PlmnInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnInfoList

`func (o *NrOperatorCellDuSingle) SetPlmnInfoList(v []PlmnInfo)`

SetPlmnInfoList sets PlmnInfoList field to given value.

### HasPlmnInfoList

`func (o *NrOperatorCellDuSingle) HasPlmnInfoList() bool`

HasPlmnInfoList returns a boolean if a field has been set.

### GetNrTac

`func (o *NrOperatorCellDuSingle) GetNrTac() int32`

GetNrTac returns the NrTac field if non-nil, zero value otherwise.

### GetNrTacOk

`func (o *NrOperatorCellDuSingle) GetNrTacOk() (*int32, bool)`

GetNrTacOk returns a tuple with the NrTac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrTac

`func (o *NrOperatorCellDuSingle) SetNrTac(v int32)`

SetNrTac sets NrTac field to given value.

### HasNrTac

`func (o *NrOperatorCellDuSingle) HasNrTac() bool`

HasNrTac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


