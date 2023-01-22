# ChargingNotifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**ReauthorizationDetails** | Pointer to [**[]ReauthorizationDetails**](ReauthorizationDetails.md) |  | [optional] 

## Methods

### NewChargingNotifyRequest

`func NewChargingNotifyRequest(notificationType NotificationType, ) *ChargingNotifyRequest`

NewChargingNotifyRequest instantiates a new ChargingNotifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingNotifyRequestWithDefaults

`func NewChargingNotifyRequestWithDefaults() *ChargingNotifyRequest`

NewChargingNotifyRequestWithDefaults instantiates a new ChargingNotifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *ChargingNotifyRequest) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *ChargingNotifyRequest) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *ChargingNotifyRequest) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetReauthorizationDetails

`func (o *ChargingNotifyRequest) GetReauthorizationDetails() []ReauthorizationDetails`

GetReauthorizationDetails returns the ReauthorizationDetails field if non-nil, zero value otherwise.

### GetReauthorizationDetailsOk

`func (o *ChargingNotifyRequest) GetReauthorizationDetailsOk() (*[]ReauthorizationDetails, bool)`

GetReauthorizationDetailsOk returns a tuple with the ReauthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthorizationDetails

`func (o *ChargingNotifyRequest) SetReauthorizationDetails(v []ReauthorizationDetails)`

SetReauthorizationDetails sets ReauthorizationDetails field to given value.

### HasReauthorizationDetails

`func (o *ChargingNotifyRequest) HasReauthorizationDetails() bool`

HasReauthorizationDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


