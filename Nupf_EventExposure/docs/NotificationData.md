# NotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationItems** | [**[]NotificationItem**](NotificationItem.md) |  | 
**CorrelationId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationData

`func NewNotificationData(notificationItems []NotificationItem, ) *NotificationData`

NewNotificationData instantiates a new NotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDataWithDefaults

`func NewNotificationDataWithDefaults() *NotificationData`

NewNotificationDataWithDefaults instantiates a new NotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationItems

`func (o *NotificationData) GetNotificationItems() []NotificationItem`

GetNotificationItems returns the NotificationItems field if non-nil, zero value otherwise.

### GetNotificationItemsOk

`func (o *NotificationData) GetNotificationItemsOk() (*[]NotificationItem, bool)`

GetNotificationItemsOk returns a tuple with the NotificationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationItems

`func (o *NotificationData) SetNotificationItems(v []NotificationItem)`

SetNotificationItems sets NotificationItems field to given value.


### GetCorrelationId

`func (o *NotificationData) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *NotificationData) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *NotificationData) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *NotificationData) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


