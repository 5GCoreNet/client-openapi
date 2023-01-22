# ExtendedSteeringContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SteeringContainer** | Pointer to [**[]SteeringInfo**](SteeringInfo.md) |  | [optional] 
**SorCmci** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**StoreSorCmciInMe** | Pointer to **bool** |  | [optional] 

## Methods

### NewExtendedSteeringContainer

`func NewExtendedSteeringContainer() *ExtendedSteeringContainer`

NewExtendedSteeringContainer instantiates a new ExtendedSteeringContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedSteeringContainerWithDefaults

`func NewExtendedSteeringContainerWithDefaults() *ExtendedSteeringContainer`

NewExtendedSteeringContainerWithDefaults instantiates a new ExtendedSteeringContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteeringContainer

`func (o *ExtendedSteeringContainer) GetSteeringContainer() []SteeringInfo`

GetSteeringContainer returns the SteeringContainer field if non-nil, zero value otherwise.

### GetSteeringContainerOk

`func (o *ExtendedSteeringContainer) GetSteeringContainerOk() (*[]SteeringInfo, bool)`

GetSteeringContainerOk returns a tuple with the SteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteeringContainer

`func (o *ExtendedSteeringContainer) SetSteeringContainer(v []SteeringInfo)`

SetSteeringContainer sets SteeringContainer field to given value.

### HasSteeringContainer

`func (o *ExtendedSteeringContainer) HasSteeringContainer() bool`

HasSteeringContainer returns a boolean if a field has been set.

### GetSorCmci

`func (o *ExtendedSteeringContainer) GetSorCmci() string`

GetSorCmci returns the SorCmci field if non-nil, zero value otherwise.

### GetSorCmciOk

`func (o *ExtendedSteeringContainer) GetSorCmciOk() (*string, bool)`

GetSorCmciOk returns a tuple with the SorCmci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorCmci

`func (o *ExtendedSteeringContainer) SetSorCmci(v string)`

SetSorCmci sets SorCmci field to given value.

### HasSorCmci

`func (o *ExtendedSteeringContainer) HasSorCmci() bool`

HasSorCmci returns a boolean if a field has been set.

### GetStoreSorCmciInMe

`func (o *ExtendedSteeringContainer) GetStoreSorCmciInMe() bool`

GetStoreSorCmciInMe returns the StoreSorCmciInMe field if non-nil, zero value otherwise.

### GetStoreSorCmciInMeOk

`func (o *ExtendedSteeringContainer) GetStoreSorCmciInMeOk() (*bool, bool)`

GetStoreSorCmciInMeOk returns a tuple with the StoreSorCmciInMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSorCmciInMe

`func (o *ExtendedSteeringContainer) SetStoreSorCmciInMe(v bool)`

SetStoreSorCmciInMe sets StoreSorCmciInMe field to given value.

### HasStoreSorCmciInMe

`func (o *ExtendedSteeringContainer) HasStoreSorCmciInMe() bool`

HasStoreSorCmciInMe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


