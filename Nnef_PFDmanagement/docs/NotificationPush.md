# NotificationPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | **[]string** |  | 
**AllowedDelay** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**PfdOp** | Pointer to [**PfdOperation**](PfdOperation.md) |  | [optional] 

## Methods

### NewNotificationPush

`func NewNotificationPush(appIds []string, ) *NotificationPush`

NewNotificationPush instantiates a new NotificationPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPushWithDefaults

`func NewNotificationPushWithDefaults() *NotificationPush`

NewNotificationPushWithDefaults instantiates a new NotificationPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *NotificationPush) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *NotificationPush) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *NotificationPush) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.


### GetAllowedDelay

`func (o *NotificationPush) GetAllowedDelay() int32`

GetAllowedDelay returns the AllowedDelay field if non-nil, zero value otherwise.

### GetAllowedDelayOk

`func (o *NotificationPush) GetAllowedDelayOk() (*int32, bool)`

GetAllowedDelayOk returns a tuple with the AllowedDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDelay

`func (o *NotificationPush) SetAllowedDelay(v int32)`

SetAllowedDelay sets AllowedDelay field to given value.

### HasAllowedDelay

`func (o *NotificationPush) HasAllowedDelay() bool`

HasAllowedDelay returns a boolean if a field has been set.

### GetPfdOp

`func (o *NotificationPush) GetPfdOp() PfdOperation`

GetPfdOp returns the PfdOp field if non-nil, zero value otherwise.

### GetPfdOpOk

`func (o *NotificationPush) GetPfdOpOk() (*PfdOperation, bool)`

GetPfdOpOk returns a tuple with the PfdOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdOp

`func (o *NotificationPush) SetPfdOp(v PfdOperation)`

SetPfdOp sets PfdOp field to given value.

### HasPfdOp

`func (o *NotificationPush) HasPfdOp() bool`

HasPfdOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


