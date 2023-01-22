# PfdManagementPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PfdDatas** | Pointer to [**map[string]PfdData**](PfdData.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewPfdManagementPatch

`func NewPfdManagementPatch() *PfdManagementPatch`

NewPfdManagementPatch instantiates a new PfdManagementPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdManagementPatchWithDefaults

`func NewPfdManagementPatchWithDefaults() *PfdManagementPatch`

NewPfdManagementPatchWithDefaults instantiates a new PfdManagementPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPfdDatas

`func (o *PfdManagementPatch) GetPfdDatas() map[string]PfdData`

GetPfdDatas returns the PfdDatas field if non-nil, zero value otherwise.

### GetPfdDatasOk

`func (o *PfdManagementPatch) GetPfdDatasOk() (*map[string]PfdData, bool)`

GetPfdDatasOk returns a tuple with the PfdDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdDatas

`func (o *PfdManagementPatch) SetPfdDatas(v map[string]PfdData)`

SetPfdDatas sets PfdDatas field to given value.

### HasPfdDatas

`func (o *PfdManagementPatch) HasPfdDatas() bool`

HasPfdDatas returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *PfdManagementPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *PfdManagementPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *PfdManagementPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *PfdManagementPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


