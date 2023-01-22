# MBSUserDataIngStatSubscPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubscs** | Pointer to [**[]SubscribedEvent**](SubscribedEvent.md) |  | [optional] 
**NotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewMBSUserDataIngStatSubscPatch

`func NewMBSUserDataIngStatSubscPatch() *MBSUserDataIngStatSubscPatch`

NewMBSUserDataIngStatSubscPatch instantiates a new MBSUserDataIngStatSubscPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserDataIngStatSubscPatchWithDefaults

`func NewMBSUserDataIngStatSubscPatchWithDefaults() *MBSUserDataIngStatSubscPatch`

NewMBSUserDataIngStatSubscPatchWithDefaults instantiates a new MBSUserDataIngStatSubscPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubscs

`func (o *MBSUserDataIngStatSubscPatch) GetEventSubscs() []SubscribedEvent`

GetEventSubscs returns the EventSubscs field if non-nil, zero value otherwise.

### GetEventSubscsOk

`func (o *MBSUserDataIngStatSubscPatch) GetEventSubscsOk() (*[]SubscribedEvent, bool)`

GetEventSubscsOk returns a tuple with the EventSubscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscs

`func (o *MBSUserDataIngStatSubscPatch) SetEventSubscs(v []SubscribedEvent)`

SetEventSubscs sets EventSubscs field to given value.

### HasEventSubscs

`func (o *MBSUserDataIngStatSubscPatch) HasEventSubscs() bool`

HasEventSubscs returns a boolean if a field has been set.

### GetNotifUri

`func (o *MBSUserDataIngStatSubscPatch) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *MBSUserDataIngStatSubscPatch) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *MBSUserDataIngStatSubscPatch) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *MBSUserDataIngStatSubscPatch) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


