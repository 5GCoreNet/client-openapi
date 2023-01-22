# HssSubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HssInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SubscriptionId** | **string** | String providing an URI formatted according to RFC 3986. | 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 

## Methods

### NewHssSubscriptionItem

`func NewHssSubscriptionItem(hssInstanceId string, subscriptionId string, ) *HssSubscriptionItem`

NewHssSubscriptionItem instantiates a new HssSubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHssSubscriptionItemWithDefaults

`func NewHssSubscriptionItemWithDefaults() *HssSubscriptionItem`

NewHssSubscriptionItemWithDefaults instantiates a new HssSubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHssInstanceId

`func (o *HssSubscriptionItem) GetHssInstanceId() string`

GetHssInstanceId returns the HssInstanceId field if non-nil, zero value otherwise.

### GetHssInstanceIdOk

`func (o *HssSubscriptionItem) GetHssInstanceIdOk() (*string, bool)`

GetHssInstanceIdOk returns a tuple with the HssInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssInstanceId

`func (o *HssSubscriptionItem) SetHssInstanceId(v string)`

SetHssInstanceId sets HssInstanceId field to given value.


### GetSubscriptionId

`func (o *HssSubscriptionItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *HssSubscriptionItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *HssSubscriptionItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetContextInfo

`func (o *HssSubscriptionItem) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *HssSubscriptionItem) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *HssSubscriptionItem) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *HssSubscriptionItem) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


