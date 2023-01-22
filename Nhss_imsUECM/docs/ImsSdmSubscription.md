# ImsSdmSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**MonitoredResourceUris** | **[]string** |  | 
**Expires** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewImsSdmSubscription

`func NewImsSdmSubscription(nfInstanceId string, callbackReference string, monitoredResourceUris []string, ) *ImsSdmSubscription`

NewImsSdmSubscription instantiates a new ImsSdmSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImsSdmSubscriptionWithDefaults

`func NewImsSdmSubscriptionWithDefaults() *ImsSdmSubscription`

NewImsSdmSubscriptionWithDefaults instantiates a new ImsSdmSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *ImsSdmSubscription) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *ImsSdmSubscription) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *ImsSdmSubscription) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetCallbackReference

`func (o *ImsSdmSubscription) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *ImsSdmSubscription) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *ImsSdmSubscription) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetMonitoredResourceUris

`func (o *ImsSdmSubscription) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *ImsSdmSubscription) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *ImsSdmSubscription) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetExpires

`func (o *ImsSdmSubscription) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ImsSdmSubscription) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ImsSdmSubscription) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ImsSdmSubscription) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


