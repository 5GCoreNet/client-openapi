# StatusResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InactiveUes** | Pointer to **[]string** |  | [optional] 
**ActiveUes** | Pointer to [**[]ActiveUe**](ActiveUe.md) |  | [optional] 

## Methods

### NewStatusResponseData

`func NewStatusResponseData() *StatusResponseData`

NewStatusResponseData instantiates a new StatusResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseDataWithDefaults

`func NewStatusResponseDataWithDefaults() *StatusResponseData`

NewStatusResponseDataWithDefaults instantiates a new StatusResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInactiveUes

`func (o *StatusResponseData) GetInactiveUes() []string`

GetInactiveUes returns the InactiveUes field if non-nil, zero value otherwise.

### GetInactiveUesOk

`func (o *StatusResponseData) GetInactiveUesOk() (*[]string, bool)`

GetInactiveUesOk returns a tuple with the InactiveUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveUes

`func (o *StatusResponseData) SetInactiveUes(v []string)`

SetInactiveUes sets InactiveUes field to given value.

### HasInactiveUes

`func (o *StatusResponseData) HasInactiveUes() bool`

HasInactiveUes returns a boolean if a field has been set.

### GetActiveUes

`func (o *StatusResponseData) GetActiveUes() []ActiveUe`

GetActiveUes returns the ActiveUes field if non-nil, zero value otherwise.

### GetActiveUesOk

`func (o *StatusResponseData) GetActiveUesOk() (*[]ActiveUe, bool)`

GetActiveUesOk returns a tuple with the ActiveUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUes

`func (o *StatusResponseData) SetActiveUes(v []ActiveUe)`

SetActiveUes sets ActiveUes field to given value.

### HasActiveUes

`func (o *StatusResponseData) HasActiveUes() bool`

HasActiveUes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


