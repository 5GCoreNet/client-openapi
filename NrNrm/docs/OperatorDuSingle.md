# OperatorDuSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**GnbId** | Pointer to **string** |  | [optional] 
**GnbIdLength** | Pointer to **int32** |  | [optional] 
**EPF1C** | Pointer to [**EPF1CSingle**](EPF1CSingle.md) |  | [optional] 
**EPF1U** | Pointer to [**[]EPF1USingle**](EPF1USingle.md) |  | [optional] 

## Methods

### NewOperatorDuSingle

`func NewOperatorDuSingle(id NullableString, ) *OperatorDuSingle`

NewOperatorDuSingle instantiates a new OperatorDuSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorDuSingleWithDefaults

`func NewOperatorDuSingleWithDefaults() *OperatorDuSingle`

NewOperatorDuSingleWithDefaults instantiates a new OperatorDuSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatorDuSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatorDuSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatorDuSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *OperatorDuSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OperatorDuSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *OperatorDuSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *OperatorDuSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *OperatorDuSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *OperatorDuSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *OperatorDuSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *OperatorDuSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *OperatorDuSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *OperatorDuSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *OperatorDuSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *OperatorDuSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *OperatorDuSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *OperatorDuSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetGnbId

`func (o *OperatorDuSingle) GetGnbId() string`

GetGnbId returns the GnbId field if non-nil, zero value otherwise.

### GetGnbIdOk

`func (o *OperatorDuSingle) GetGnbIdOk() (*string, bool)`

GetGnbIdOk returns a tuple with the GnbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbId

`func (o *OperatorDuSingle) SetGnbId(v string)`

SetGnbId sets GnbId field to given value.

### HasGnbId

`func (o *OperatorDuSingle) HasGnbId() bool`

HasGnbId returns a boolean if a field has been set.

### GetGnbIdLength

`func (o *OperatorDuSingle) GetGnbIdLength() int32`

GetGnbIdLength returns the GnbIdLength field if non-nil, zero value otherwise.

### GetGnbIdLengthOk

`func (o *OperatorDuSingle) GetGnbIdLengthOk() (*int32, bool)`

GetGnbIdLengthOk returns a tuple with the GnbIdLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbIdLength

`func (o *OperatorDuSingle) SetGnbIdLength(v int32)`

SetGnbIdLength sets GnbIdLength field to given value.

### HasGnbIdLength

`func (o *OperatorDuSingle) HasGnbIdLength() bool`

HasGnbIdLength returns a boolean if a field has been set.

### GetEPF1C

`func (o *OperatorDuSingle) GetEPF1C() EPF1CSingle`

GetEPF1C returns the EPF1C field if non-nil, zero value otherwise.

### GetEPF1COk

`func (o *OperatorDuSingle) GetEPF1COk() (*EPF1CSingle, bool)`

GetEPF1COk returns a tuple with the EPF1C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1C

`func (o *OperatorDuSingle) SetEPF1C(v EPF1CSingle)`

SetEPF1C sets EPF1C field to given value.

### HasEPF1C

`func (o *OperatorDuSingle) HasEPF1C() bool`

HasEPF1C returns a boolean if a field has been set.

### GetEPF1U

`func (o *OperatorDuSingle) GetEPF1U() []EPF1USingle`

GetEPF1U returns the EPF1U field if non-nil, zero value otherwise.

### GetEPF1UOk

`func (o *OperatorDuSingle) GetEPF1UOk() (*[]EPF1USingle, bool)`

GetEPF1UOk returns a tuple with the EPF1U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1U

`func (o *OperatorDuSingle) SetEPF1U(v []EPF1USingle)`

SetEPF1U sets EPF1U field to given value.

### HasEPF1U

`func (o *OperatorDuSingle) HasEPF1U() bool`

HasEPF1U returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


