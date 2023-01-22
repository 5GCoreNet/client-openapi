# ACTStatusNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**ActStatus** | [**ACTResult**](ACTResult.md) |  | 

## Methods

### NewACTStatusNotif

`func NewACTStatusNotif(subscriptionId string, actStatus ACTResult, ) *ACTStatusNotif`

NewACTStatusNotif instantiates a new ACTStatusNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACTStatusNotifWithDefaults

`func NewACTStatusNotifWithDefaults() *ACTStatusNotif`

NewACTStatusNotifWithDefaults instantiates a new ACTStatusNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *ACTStatusNotif) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ACTStatusNotif) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ACTStatusNotif) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetActStatus

`func (o *ACTStatusNotif) GetActStatus() ACTResult`

GetActStatus returns the ActStatus field if non-nil, zero value otherwise.

### GetActStatusOk

`func (o *ACTStatusNotif) GetActStatusOk() (*ACTResult, bool)`

GetActStatusOk returns a tuple with the ActStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActStatus

`func (o *ACTStatusNotif) SetActStatus(v ACTResult)`

SetActStatus sets ActStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


