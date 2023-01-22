# IdentityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValSvcId** | Pointer to **string** | Identity of the VAL service | [optional] 
**ValTgtUes** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | VAL User IDs or VAL UE IDs that the event subscriber wants to know in the interested event.  | [optional] 
**SuppLoc** | Pointer to **bool** | Set to true by Subscriber to request the supplementary location information. | [optional] 

## Methods

### NewIdentityFilter

`func NewIdentityFilter() *IdentityFilter`

NewIdentityFilter instantiates a new IdentityFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityFilterWithDefaults

`func NewIdentityFilterWithDefaults() *IdentityFilter`

NewIdentityFilterWithDefaults instantiates a new IdentityFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValSvcId

`func (o *IdentityFilter) GetValSvcId() string`

GetValSvcId returns the ValSvcId field if non-nil, zero value otherwise.

### GetValSvcIdOk

`func (o *IdentityFilter) GetValSvcIdOk() (*string, bool)`

GetValSvcIdOk returns a tuple with the ValSvcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValSvcId

`func (o *IdentityFilter) SetValSvcId(v string)`

SetValSvcId sets ValSvcId field to given value.

### HasValSvcId

`func (o *IdentityFilter) HasValSvcId() bool`

HasValSvcId returns a boolean if a field has been set.

### GetValTgtUes

`func (o *IdentityFilter) GetValTgtUes() []ValTargetUe`

GetValTgtUes returns the ValTgtUes field if non-nil, zero value otherwise.

### GetValTgtUesOk

`func (o *IdentityFilter) GetValTgtUesOk() (*[]ValTargetUe, bool)`

GetValTgtUesOk returns a tuple with the ValTgtUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUes

`func (o *IdentityFilter) SetValTgtUes(v []ValTargetUe)`

SetValTgtUes sets ValTgtUes field to given value.

### HasValTgtUes

`func (o *IdentityFilter) HasValTgtUes() bool`

HasValTgtUes returns a boolean if a field has been set.

### GetSuppLoc

`func (o *IdentityFilter) GetSuppLoc() bool`

GetSuppLoc returns the SuppLoc field if non-nil, zero value otherwise.

### GetSuppLocOk

`func (o *IdentityFilter) GetSuppLocOk() (*bool, bool)`

GetSuppLocOk returns a tuple with the SuppLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppLoc

`func (o *IdentityFilter) SetSuppLoc(v bool)`

SetSuppLoc sets SuppLoc field to given value.

### HasSuppLoc

`func (o *IdentityFilter) HasSuppLoc() bool`

HasSuppLoc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


