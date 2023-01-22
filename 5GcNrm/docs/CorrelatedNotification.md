# CorrelatedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceObjectInstance** | **string** |  | 
**NotificationIds** | **[]int32** |  | 

## Methods

### NewCorrelatedNotification

`func NewCorrelatedNotification(sourceObjectInstance string, notificationIds []int32, ) *CorrelatedNotification`

NewCorrelatedNotification instantiates a new CorrelatedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelatedNotificationWithDefaults

`func NewCorrelatedNotificationWithDefaults() *CorrelatedNotification`

NewCorrelatedNotificationWithDefaults instantiates a new CorrelatedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceObjectInstance

`func (o *CorrelatedNotification) GetSourceObjectInstance() string`

GetSourceObjectInstance returns the SourceObjectInstance field if non-nil, zero value otherwise.

### GetSourceObjectInstanceOk

`func (o *CorrelatedNotification) GetSourceObjectInstanceOk() (*string, bool)`

GetSourceObjectInstanceOk returns a tuple with the SourceObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectInstance

`func (o *CorrelatedNotification) SetSourceObjectInstance(v string)`

SetSourceObjectInstance sets SourceObjectInstance field to given value.


### GetNotificationIds

`func (o *CorrelatedNotification) GetNotificationIds() []int32`

GetNotificationIds returns the NotificationIds field if non-nil, zero value otherwise.

### GetNotificationIdsOk

`func (o *CorrelatedNotification) GetNotificationIdsOk() (*[]int32, bool)`

GetNotificationIdsOk returns a tuple with the NotificationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIds

`func (o *CorrelatedNotification) SetNotificationIds(v []int32)`

SetNotificationIds sets NotificationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


