# CorrelatedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**NotificationIds** | **[]int32** |  | 

## Methods

### NewCorrelatedNotification

`func NewCorrelatedNotification(source string, notificationIds []int32, ) *CorrelatedNotification`

NewCorrelatedNotification instantiates a new CorrelatedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelatedNotificationWithDefaults

`func NewCorrelatedNotificationWithDefaults() *CorrelatedNotification`

NewCorrelatedNotificationWithDefaults instantiates a new CorrelatedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *CorrelatedNotification) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CorrelatedNotification) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CorrelatedNotification) SetSource(v string)`

SetSource sets Source field to given value.


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


