# MoiChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationId** | **int32** |  | 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**SourceIndicator** | Pointer to [**SourceIndicator**](SourceIndicator.md) |  | [optional] 
**Op** | [**Operation**](Operation.md) |  | 
**Path** | **string** |  | 
**Value** | Pointer to **interface{}** |  | [optional] 
**OldValue** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewMoiChange

`func NewMoiChange(notificationId int32, op Operation, path string, ) *MoiChange`

NewMoiChange instantiates a new MoiChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiChangeWithDefaults

`func NewMoiChangeWithDefaults() *MoiChange`

NewMoiChangeWithDefaults instantiates a new MoiChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationId

`func (o *MoiChange) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *MoiChange) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *MoiChange) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetCorrelatedNotifications

`func (o *MoiChange) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *MoiChange) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *MoiChange) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *MoiChange) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *MoiChange) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *MoiChange) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *MoiChange) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *MoiChange) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetSourceIndicator

`func (o *MoiChange) GetSourceIndicator() SourceIndicator`

GetSourceIndicator returns the SourceIndicator field if non-nil, zero value otherwise.

### GetSourceIndicatorOk

`func (o *MoiChange) GetSourceIndicatorOk() (*SourceIndicator, bool)`

GetSourceIndicatorOk returns a tuple with the SourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIndicator

`func (o *MoiChange) SetSourceIndicator(v SourceIndicator)`

SetSourceIndicator sets SourceIndicator field to given value.

### HasSourceIndicator

`func (o *MoiChange) HasSourceIndicator() bool`

HasSourceIndicator returns a boolean if a field has been set.

### GetOp

`func (o *MoiChange) GetOp() Operation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *MoiChange) GetOpOk() (*Operation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *MoiChange) SetOp(v Operation)`

SetOp sets Op field to given value.


### GetPath

`func (o *MoiChange) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MoiChange) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MoiChange) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *MoiChange) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MoiChange) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MoiChange) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *MoiChange) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MoiChange) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MoiChange) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOldValue

`func (o *MoiChange) GetOldValue() interface{}`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *MoiChange) GetOldValueOk() (*interface{}, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *MoiChange) SetOldValue(v interface{})`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *MoiChange) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### SetOldValueNil

`func (o *MoiChange) SetOldValueNil(b bool)`

 SetOldValueNil sets the value for OldValue to be an explicit nil

### UnsetOldValue
`func (o *MoiChange) UnsetOldValue()`

UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


