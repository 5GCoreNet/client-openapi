# NotificationDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordRef** | **string** | String providing an URI formatted according to RFC 3986. | 
**OperationType** | [**RecordOperation**](RecordOperation.md) |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationDescription

`func NewNotificationDescription(recordRef string, operationType RecordOperation, ) *NotificationDescription`

NewNotificationDescription instantiates a new NotificationDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDescriptionWithDefaults

`func NewNotificationDescriptionWithDefaults() *NotificationDescription`

NewNotificationDescriptionWithDefaults instantiates a new NotificationDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordRef

`func (o *NotificationDescription) GetRecordRef() string`

GetRecordRef returns the RecordRef field if non-nil, zero value otherwise.

### GetRecordRefOk

`func (o *NotificationDescription) GetRecordRefOk() (*string, bool)`

GetRecordRefOk returns a tuple with the RecordRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordRef

`func (o *NotificationDescription) SetRecordRef(v string)`

SetRecordRef sets RecordRef field to given value.


### GetOperationType

`func (o *NotificationDescription) GetOperationType() RecordOperation`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *NotificationDescription) GetOperationTypeOk() (*RecordOperation, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *NotificationDescription) SetOperationType(v RecordOperation)`

SetOperationType sets OperationType field to given value.


### GetSubscriptionId

`func (o *NotificationDescription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *NotificationDescription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *NotificationDescription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *NotificationDescription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


