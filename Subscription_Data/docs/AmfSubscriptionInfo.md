# AmfSubscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SubscriptionId** | **string** | String providing an URI formatted according to RFC 3986. | 
**SubsChangeNotifyCorrelationId** | Pointer to **string** |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 

## Methods

### NewAmfSubscriptionInfo

`func NewAmfSubscriptionInfo(amfInstanceId string, subscriptionId string, ) *AmfSubscriptionInfo`

NewAmfSubscriptionInfo instantiates a new AmfSubscriptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfSubscriptionInfoWithDefaults

`func NewAmfSubscriptionInfoWithDefaults() *AmfSubscriptionInfo`

NewAmfSubscriptionInfoWithDefaults instantiates a new AmfSubscriptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfInstanceId

`func (o *AmfSubscriptionInfo) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *AmfSubscriptionInfo) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *AmfSubscriptionInfo) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.


### GetSubscriptionId

`func (o *AmfSubscriptionInfo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AmfSubscriptionInfo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AmfSubscriptionInfo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSubsChangeNotifyCorrelationId

`func (o *AmfSubscriptionInfo) GetSubsChangeNotifyCorrelationId() string`

GetSubsChangeNotifyCorrelationId returns the SubsChangeNotifyCorrelationId field if non-nil, zero value otherwise.

### GetSubsChangeNotifyCorrelationIdOk

`func (o *AmfSubscriptionInfo) GetSubsChangeNotifyCorrelationIdOk() (*string, bool)`

GetSubsChangeNotifyCorrelationIdOk returns a tuple with the SubsChangeNotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsChangeNotifyCorrelationId

`func (o *AmfSubscriptionInfo) SetSubsChangeNotifyCorrelationId(v string)`

SetSubsChangeNotifyCorrelationId sets SubsChangeNotifyCorrelationId field to given value.

### HasSubsChangeNotifyCorrelationId

`func (o *AmfSubscriptionInfo) HasSubsChangeNotifyCorrelationId() bool`

HasSubsChangeNotifyCorrelationId returns a boolean if a field has been set.

### GetContextInfo

`func (o *AmfSubscriptionInfo) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *AmfSubscriptionInfo) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *AmfSubscriptionInfo) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *AmfSubscriptionInfo) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


