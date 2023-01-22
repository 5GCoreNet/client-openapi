# AcrMgntEventsSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubscs** | Pointer to [**[]AcrMgntEventSubsc**](AcrMgntEventSubsc.md) | The subscribed ACR management events. | [optional] 
**EvtReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 

## Methods

### NewAcrMgntEventsSubscriptionPatch

`func NewAcrMgntEventsSubscriptionPatch() *AcrMgntEventsSubscriptionPatch`

NewAcrMgntEventsSubscriptionPatch instantiates a new AcrMgntEventsSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrMgntEventsSubscriptionPatchWithDefaults

`func NewAcrMgntEventsSubscriptionPatchWithDefaults() *AcrMgntEventsSubscriptionPatch`

NewAcrMgntEventsSubscriptionPatchWithDefaults instantiates a new AcrMgntEventsSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubscs

`func (o *AcrMgntEventsSubscriptionPatch) GetEventSubscs() []AcrMgntEventSubsc`

GetEventSubscs returns the EventSubscs field if non-nil, zero value otherwise.

### GetEventSubscsOk

`func (o *AcrMgntEventsSubscriptionPatch) GetEventSubscsOk() (*[]AcrMgntEventSubsc, bool)`

GetEventSubscsOk returns a tuple with the EventSubscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscs

`func (o *AcrMgntEventsSubscriptionPatch) SetEventSubscs(v []AcrMgntEventSubsc)`

SetEventSubscs sets EventSubscs field to given value.

### HasEventSubscs

`func (o *AcrMgntEventsSubscriptionPatch) HasEventSubscs() bool`

HasEventSubscs returns a boolean if a field has been set.

### GetEvtReq

`func (o *AcrMgntEventsSubscriptionPatch) GetEvtReq() ReportingInformation`

GetEvtReq returns the EvtReq field if non-nil, zero value otherwise.

### GetEvtReqOk

`func (o *AcrMgntEventsSubscriptionPatch) GetEvtReqOk() (*ReportingInformation, bool)`

GetEvtReqOk returns a tuple with the EvtReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtReq

`func (o *AcrMgntEventsSubscriptionPatch) SetEvtReq(v ReportingInformation)`

SetEvtReq sets EvtReq field to given value.

### HasEvtReq

`func (o *AcrMgntEventsSubscriptionPatch) HasEvtReq() bool`

HasEvtReq returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *AcrMgntEventsSubscriptionPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *AcrMgntEventsSubscriptionPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *AcrMgntEventsSubscriptionPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *AcrMgntEventsSubscriptionPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


