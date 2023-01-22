# EasDepNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasDepInfo** | [**EasDeployInfoData**](EasDeployInfoData.md) |  | 
**EventId** | [**EasEvent**](EasEvent.md) |  | 

## Methods

### NewEasDepNotification

`func NewEasDepNotification(easDepInfo EasDeployInfoData, eventId EasEvent, ) *EasDepNotification`

NewEasDepNotification instantiates a new EasDepNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDepNotificationWithDefaults

`func NewEasDepNotificationWithDefaults() *EasDepNotification`

NewEasDepNotificationWithDefaults instantiates a new EasDepNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasDepInfo

`func (o *EasDepNotification) GetEasDepInfo() EasDeployInfoData`

GetEasDepInfo returns the EasDepInfo field if non-nil, zero value otherwise.

### GetEasDepInfoOk

`func (o *EasDepNotification) GetEasDepInfoOk() (*EasDeployInfoData, bool)`

GetEasDepInfoOk returns a tuple with the EasDepInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDepInfo

`func (o *EasDepNotification) SetEasDepInfo(v EasDeployInfoData)`

SetEasDepInfo sets EasDepInfo field to given value.


### GetEventId

`func (o *EasDepNotification) GetEventId() EasEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EasDepNotification) GetEventIdOk() (*EasEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EasDepNotification) SetEventId(v EasEvent)`

SetEventId sets EventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


