# ConsentRevocNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**ConsentsRevoked** | [**[]ConsentRevoked**](ConsentRevoked.md) |  | 

## Methods

### NewConsentRevocNotif

`func NewConsentRevocNotif(subscriptionId string, consentsRevoked []ConsentRevoked, ) *ConsentRevocNotif`

NewConsentRevocNotif instantiates a new ConsentRevocNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentRevocNotifWithDefaults

`func NewConsentRevocNotifWithDefaults() *ConsentRevocNotif`

NewConsentRevocNotifWithDefaults instantiates a new ConsentRevocNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *ConsentRevocNotif) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ConsentRevocNotif) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ConsentRevocNotif) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetConsentsRevoked

`func (o *ConsentRevocNotif) GetConsentsRevoked() []ConsentRevoked`

GetConsentsRevoked returns the ConsentsRevoked field if non-nil, zero value otherwise.

### GetConsentsRevokedOk

`func (o *ConsentRevocNotif) GetConsentsRevokedOk() (*[]ConsentRevoked, bool)`

GetConsentsRevokedOk returns a tuple with the ConsentsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentsRevoked

`func (o *ConsentRevocNotif) SetConsentsRevoked(v []ConsentRevoked)`

SetConsentsRevoked sets ConsentsRevoked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


