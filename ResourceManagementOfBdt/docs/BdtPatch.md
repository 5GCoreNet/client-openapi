# BdtPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedPolicy** | **int32** | Identity of the selected background data transfer policy. | 
**WarnNotifEnabled** | Pointer to **bool** | Indicates whether the BDT warning notification is enabled (true) or not (false).  | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewBdtPatch

`func NewBdtPatch(selectedPolicy int32, ) *BdtPatch`

NewBdtPatch instantiates a new BdtPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtPatchWithDefaults

`func NewBdtPatchWithDefaults() *BdtPatch`

NewBdtPatchWithDefaults instantiates a new BdtPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedPolicy

`func (o *BdtPatch) GetSelectedPolicy() int32`

GetSelectedPolicy returns the SelectedPolicy field if non-nil, zero value otherwise.

### GetSelectedPolicyOk

`func (o *BdtPatch) GetSelectedPolicyOk() (*int32, bool)`

GetSelectedPolicyOk returns a tuple with the SelectedPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPolicy

`func (o *BdtPatch) SetSelectedPolicy(v int32)`

SetSelectedPolicy sets SelectedPolicy field to given value.


### GetWarnNotifEnabled

`func (o *BdtPatch) GetWarnNotifEnabled() bool`

GetWarnNotifEnabled returns the WarnNotifEnabled field if non-nil, zero value otherwise.

### GetWarnNotifEnabledOk

`func (o *BdtPatch) GetWarnNotifEnabledOk() (*bool, bool)`

GetWarnNotifEnabledOk returns a tuple with the WarnNotifEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnNotifEnabled

`func (o *BdtPatch) SetWarnNotifEnabled(v bool)`

SetWarnNotifEnabled sets WarnNotifEnabled field to given value.

### HasWarnNotifEnabled

`func (o *BdtPatch) HasWarnNotifEnabled() bool`

HasWarnNotifEnabled returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *BdtPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *BdtPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *BdtPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *BdtPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


