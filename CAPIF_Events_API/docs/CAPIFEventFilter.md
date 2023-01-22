# CAPIFEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiIds** | Pointer to **[]string** | Identifier of the service API | [optional] 
**ApiInvokerIds** | Pointer to **[]string** | Identity of the API invoker | [optional] 
**AefIds** | Pointer to **[]string** | Identifier of the API exposing function | [optional] 

## Methods

### NewCAPIFEventFilter

`func NewCAPIFEventFilter() *CAPIFEventFilter`

NewCAPIFEventFilter instantiates a new CAPIFEventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCAPIFEventFilterWithDefaults

`func NewCAPIFEventFilterWithDefaults() *CAPIFEventFilter`

NewCAPIFEventFilterWithDefaults instantiates a new CAPIFEventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiIds

`func (o *CAPIFEventFilter) GetApiIds() []string`

GetApiIds returns the ApiIds field if non-nil, zero value otherwise.

### GetApiIdsOk

`func (o *CAPIFEventFilter) GetApiIdsOk() (*[]string, bool)`

GetApiIdsOk returns a tuple with the ApiIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiIds

`func (o *CAPIFEventFilter) SetApiIds(v []string)`

SetApiIds sets ApiIds field to given value.

### HasApiIds

`func (o *CAPIFEventFilter) HasApiIds() bool`

HasApiIds returns a boolean if a field has been set.

### GetApiInvokerIds

`func (o *CAPIFEventFilter) GetApiInvokerIds() []string`

GetApiInvokerIds returns the ApiInvokerIds field if non-nil, zero value otherwise.

### GetApiInvokerIdsOk

`func (o *CAPIFEventFilter) GetApiInvokerIdsOk() (*[]string, bool)`

GetApiInvokerIdsOk returns a tuple with the ApiInvokerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerIds

`func (o *CAPIFEventFilter) SetApiInvokerIds(v []string)`

SetApiInvokerIds sets ApiInvokerIds field to given value.

### HasApiInvokerIds

`func (o *CAPIFEventFilter) HasApiInvokerIds() bool`

HasApiInvokerIds returns a boolean if a field has been set.

### GetAefIds

`func (o *CAPIFEventFilter) GetAefIds() []string`

GetAefIds returns the AefIds field if non-nil, zero value otherwise.

### GetAefIdsOk

`func (o *CAPIFEventFilter) GetAefIdsOk() (*[]string, bool)`

GetAefIdsOk returns a tuple with the AefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefIds

`func (o *CAPIFEventFilter) SetAefIds(v []string)`

SetAefIds sets AefIds field to given value.

### HasAefIds

`func (o *CAPIFEventFilter) HasAefIds() bool`

HasAefIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


