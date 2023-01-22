# SmfSubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SubscriptionId** | **string** | String providing an URI formatted according to RFC 3986. | 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 

## Methods

### NewSmfSubscriptionItem

`func NewSmfSubscriptionItem(smfInstanceId string, subscriptionId string, ) *SmfSubscriptionItem`

NewSmfSubscriptionItem instantiates a new SmfSubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfSubscriptionItemWithDefaults

`func NewSmfSubscriptionItemWithDefaults() *SmfSubscriptionItem`

NewSmfSubscriptionItemWithDefaults instantiates a new SmfSubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmfInstanceId

`func (o *SmfSubscriptionItem) GetSmfInstanceId() string`

GetSmfInstanceId returns the SmfInstanceId field if non-nil, zero value otherwise.

### GetSmfInstanceIdOk

`func (o *SmfSubscriptionItem) GetSmfInstanceIdOk() (*string, bool)`

GetSmfInstanceIdOk returns a tuple with the SmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInstanceId

`func (o *SmfSubscriptionItem) SetSmfInstanceId(v string)`

SetSmfInstanceId sets SmfInstanceId field to given value.


### GetSubscriptionId

`func (o *SmfSubscriptionItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SmfSubscriptionItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SmfSubscriptionItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetContextInfo

`func (o *SmfSubscriptionItem) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *SmfSubscriptionItem) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *SmfSubscriptionItem) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *SmfSubscriptionItem) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


