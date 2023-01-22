# ACREventsSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EasIds** | Pointer to **[]string** | The list of application identifiers of the EASs. | [optional] 
**EventIds** | Pointer to [**ACREventIDs**](ACREventIDs.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 

## Methods

### NewACREventsSubscriptionPatch

`func NewACREventsSubscriptionPatch() *ACREventsSubscriptionPatch`

NewACREventsSubscriptionPatch instantiates a new ACREventsSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACREventsSubscriptionPatchWithDefaults

`func NewACREventsSubscriptionPatchWithDefaults() *ACREventsSubscriptionPatch`

NewACREventsSubscriptionPatchWithDefaults instantiates a new ACREventsSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpTime

`func (o *ACREventsSubscriptionPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *ACREventsSubscriptionPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *ACREventsSubscriptionPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *ACREventsSubscriptionPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetEasIds

`func (o *ACREventsSubscriptionPatch) GetEasIds() []string`

GetEasIds returns the EasIds field if non-nil, zero value otherwise.

### GetEasIdsOk

`func (o *ACREventsSubscriptionPatch) GetEasIdsOk() (*[]string, bool)`

GetEasIdsOk returns a tuple with the EasIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIds

`func (o *ACREventsSubscriptionPatch) SetEasIds(v []string)`

SetEasIds sets EasIds field to given value.

### HasEasIds

`func (o *ACREventsSubscriptionPatch) HasEasIds() bool`

HasEasIds returns a boolean if a field has been set.

### GetEventIds

`func (o *ACREventsSubscriptionPatch) GetEventIds() ACREventIDs`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *ACREventsSubscriptionPatch) GetEventIdsOk() (*ACREventIDs, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *ACREventsSubscriptionPatch) SetEventIds(v ACREventIDs)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *ACREventsSubscriptionPatch) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *ACREventsSubscriptionPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ACREventsSubscriptionPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ACREventsSubscriptionPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *ACREventsSubscriptionPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


