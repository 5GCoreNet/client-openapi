# EasDiscoveryNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubId** | **string** | Identifier of the individual service provisioning subscription for which the service provisioning notification is delivered. | 
**EventType** | [**EASDiscEventIDs**](EASDiscEventIDs.md) |  | 
**DiscoveredEas** | [**[]DiscoveredEas**](DiscoveredEas.md) | List of EAS discovery information. | 

## Methods

### NewEasDiscoveryNotification

`func NewEasDiscoveryNotification(subId string, eventType EASDiscEventIDs, discoveredEas []DiscoveredEas, ) *EasDiscoveryNotification`

NewEasDiscoveryNotification instantiates a new EasDiscoveryNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDiscoveryNotificationWithDefaults

`func NewEasDiscoveryNotificationWithDefaults() *EasDiscoveryNotification`

NewEasDiscoveryNotificationWithDefaults instantiates a new EasDiscoveryNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubId

`func (o *EasDiscoveryNotification) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *EasDiscoveryNotification) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *EasDiscoveryNotification) SetSubId(v string)`

SetSubId sets SubId field to given value.


### GetEventType

`func (o *EasDiscoveryNotification) GetEventType() EASDiscEventIDs`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EasDiscoveryNotification) GetEventTypeOk() (*EASDiscEventIDs, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EasDiscoveryNotification) SetEventType(v EASDiscEventIDs)`

SetEventType sets EventType field to given value.


### GetDiscoveredEas

`func (o *EasDiscoveryNotification) GetDiscoveredEas() []DiscoveredEas`

GetDiscoveredEas returns the DiscoveredEas field if non-nil, zero value otherwise.

### GetDiscoveredEasOk

`func (o *EasDiscoveryNotification) GetDiscoveredEasOk() (*[]DiscoveredEas, bool)`

GetDiscoveredEasOk returns a tuple with the DiscoveredEas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredEas

`func (o *EasDiscoveryNotification) SetDiscoveredEas(v []DiscoveredEas)`

SetDiscoveredEas sets DiscoveredEas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


