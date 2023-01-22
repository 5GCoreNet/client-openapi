# SEALEventSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubs** | Pointer to [**[]EventSubscription**](EventSubscription.md) | Subscribed events. | [optional] 
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 

## Methods

### NewSEALEventSubscriptionPatch

`func NewSEALEventSubscriptionPatch() *SEALEventSubscriptionPatch`

NewSEALEventSubscriptionPatch instantiates a new SEALEventSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSEALEventSubscriptionPatchWithDefaults

`func NewSEALEventSubscriptionPatchWithDefaults() *SEALEventSubscriptionPatch`

NewSEALEventSubscriptionPatchWithDefaults instantiates a new SEALEventSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubs

`func (o *SEALEventSubscriptionPatch) GetEventSubs() []EventSubscription`

GetEventSubs returns the EventSubs field if non-nil, zero value otherwise.

### GetEventSubsOk

`func (o *SEALEventSubscriptionPatch) GetEventSubsOk() (*[]EventSubscription, bool)`

GetEventSubsOk returns a tuple with the EventSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubs

`func (o *SEALEventSubscriptionPatch) SetEventSubs(v []EventSubscription)`

SetEventSubs sets EventSubs field to given value.

### HasEventSubs

`func (o *SEALEventSubscriptionPatch) HasEventSubs() bool`

HasEventSubs returns a boolean if a field has been set.

### GetEventReq

`func (o *SEALEventSubscriptionPatch) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *SEALEventSubscriptionPatch) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *SEALEventSubscriptionPatch) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *SEALEventSubscriptionPatch) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *SEALEventSubscriptionPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *SEALEventSubscriptionPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *SEALEventSubscriptionPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *SEALEventSubscriptionPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


