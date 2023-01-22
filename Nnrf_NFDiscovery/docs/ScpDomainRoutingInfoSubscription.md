# ScpDomainRoutingInfoSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**ValidityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ReqInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**LocalInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewScpDomainRoutingInfoSubscription

`func NewScpDomainRoutingInfoSubscription(callbackUri string, ) *ScpDomainRoutingInfoSubscription`

NewScpDomainRoutingInfoSubscription instantiates a new ScpDomainRoutingInfoSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScpDomainRoutingInfoSubscriptionWithDefaults

`func NewScpDomainRoutingInfoSubscriptionWithDefaults() *ScpDomainRoutingInfoSubscription`

NewScpDomainRoutingInfoSubscriptionWithDefaults instantiates a new ScpDomainRoutingInfoSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUri

`func (o *ScpDomainRoutingInfoSubscription) GetCallbackUri() string`

GetCallbackUri returns the CallbackUri field if non-nil, zero value otherwise.

### GetCallbackUriOk

`func (o *ScpDomainRoutingInfoSubscription) GetCallbackUriOk() (*string, bool)`

GetCallbackUriOk returns a tuple with the CallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUri

`func (o *ScpDomainRoutingInfoSubscription) SetCallbackUri(v string)`

SetCallbackUri sets CallbackUri field to given value.


### GetValidityTime

`func (o *ScpDomainRoutingInfoSubscription) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *ScpDomainRoutingInfoSubscription) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *ScpDomainRoutingInfoSubscription) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *ScpDomainRoutingInfoSubscription) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetReqInstanceId

`func (o *ScpDomainRoutingInfoSubscription) GetReqInstanceId() string`

GetReqInstanceId returns the ReqInstanceId field if non-nil, zero value otherwise.

### GetReqInstanceIdOk

`func (o *ScpDomainRoutingInfoSubscription) GetReqInstanceIdOk() (*string, bool)`

GetReqInstanceIdOk returns a tuple with the ReqInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqInstanceId

`func (o *ScpDomainRoutingInfoSubscription) SetReqInstanceId(v string)`

SetReqInstanceId sets ReqInstanceId field to given value.

### HasReqInstanceId

`func (o *ScpDomainRoutingInfoSubscription) HasReqInstanceId() bool`

HasReqInstanceId returns a boolean if a field has been set.

### GetLocalInd

`func (o *ScpDomainRoutingInfoSubscription) GetLocalInd() bool`

GetLocalInd returns the LocalInd field if non-nil, zero value otherwise.

### GetLocalIndOk

`func (o *ScpDomainRoutingInfoSubscription) GetLocalIndOk() (*bool, bool)`

GetLocalIndOk returns a tuple with the LocalInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInd

`func (o *ScpDomainRoutingInfoSubscription) SetLocalInd(v bool)`

SetLocalInd sets LocalInd field to given value.

### HasLocalInd

`func (o *ScpDomainRoutingInfoSubscription) HasLocalInd() bool`

HasLocalInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


