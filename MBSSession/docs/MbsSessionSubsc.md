# MbsSessionSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | **string** |  | 
**Subscription** | [**MbsSessionSubscription**](MbsSessionSubscription.md) |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewMbsSessionSubsc

`func NewMbsSessionSubsc(afId string, subscription MbsSessionSubscription, ) *MbsSessionSubsc`

NewMbsSessionSubsc instantiates a new MbsSessionSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionSubscWithDefaults

`func NewMbsSessionSubscWithDefaults() *MbsSessionSubsc`

NewMbsSessionSubscWithDefaults instantiates a new MbsSessionSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *MbsSessionSubsc) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *MbsSessionSubsc) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *MbsSessionSubsc) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetSubscription

`func (o *MbsSessionSubsc) GetSubscription() MbsSessionSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MbsSessionSubsc) GetSubscriptionOk() (*MbsSessionSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MbsSessionSubsc) SetSubscription(v MbsSessionSubscription)`

SetSubscription sets Subscription field to given value.


### GetSubscriptionId

`func (o *MbsSessionSubsc) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MbsSessionSubsc) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MbsSessionSubsc) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MbsSessionSubsc) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


