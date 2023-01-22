# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerReference** | Pointer to **string** |  | [optional] 
**TimeTick** | Pointer to **int32** |  | [optional] 
**Filter** | Pointer to **string** | The filter format shall be compliant to XPath 1.0. | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerReference

`func (o *Subscription) GetConsumerReference() string`

GetConsumerReference returns the ConsumerReference field if non-nil, zero value otherwise.

### GetConsumerReferenceOk

`func (o *Subscription) GetConsumerReferenceOk() (*string, bool)`

GetConsumerReferenceOk returns a tuple with the ConsumerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerReference

`func (o *Subscription) SetConsumerReference(v string)`

SetConsumerReference sets ConsumerReference field to given value.

### HasConsumerReference

`func (o *Subscription) HasConsumerReference() bool`

HasConsumerReference returns a boolean if a field has been set.

### GetTimeTick

`func (o *Subscription) GetTimeTick() int32`

GetTimeTick returns the TimeTick field if non-nil, zero value otherwise.

### GetTimeTickOk

`func (o *Subscription) GetTimeTickOk() (*int32, bool)`

GetTimeTickOk returns a tuple with the TimeTick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTick

`func (o *Subscription) SetTimeTick(v int32)`

SetTimeTick sets TimeTick field to given value.

### HasTimeTick

`func (o *Subscription) HasTimeTick() bool`

HasTimeTick returns a boolean if a field has been set.

### GetFilter

`func (o *Subscription) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Subscription) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Subscription) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Subscription) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


