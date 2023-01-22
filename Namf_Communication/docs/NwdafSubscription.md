# NwdafSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NwdafEvtSubsServiceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NwdafEventsSubscription** | [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | 

## Methods

### NewNwdafSubscription

`func NewNwdafSubscription(nwdafEvtSubsServiceUri string, nwdafEventsSubscription NnwdafEventsSubscription, ) *NwdafSubscription`

NewNwdafSubscription instantiates a new NwdafSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwdafSubscriptionWithDefaults

`func NewNwdafSubscriptionWithDefaults() *NwdafSubscription`

NewNwdafSubscriptionWithDefaults instantiates a new NwdafSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNwdafEvtSubsServiceUri

`func (o *NwdafSubscription) GetNwdafEvtSubsServiceUri() string`

GetNwdafEvtSubsServiceUri returns the NwdafEvtSubsServiceUri field if non-nil, zero value otherwise.

### GetNwdafEvtSubsServiceUriOk

`func (o *NwdafSubscription) GetNwdafEvtSubsServiceUriOk() (*string, bool)`

GetNwdafEvtSubsServiceUriOk returns a tuple with the NwdafEvtSubsServiceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEvtSubsServiceUri

`func (o *NwdafSubscription) SetNwdafEvtSubsServiceUri(v string)`

SetNwdafEvtSubsServiceUri sets NwdafEvtSubsServiceUri field to given value.


### GetNwdafEventsSubscription

`func (o *NwdafSubscription) GetNwdafEventsSubscription() NnwdafEventsSubscription`

GetNwdafEventsSubscription returns the NwdafEventsSubscription field if non-nil, zero value otherwise.

### GetNwdafEventsSubscriptionOk

`func (o *NwdafSubscription) GetNwdafEventsSubscriptionOk() (*NnwdafEventsSubscription, bool)`

GetNwdafEventsSubscriptionOk returns a tuple with the NwdafEventsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEventsSubscription

`func (o *NwdafSubscription) SetNwdafEventsSubscription(v NnwdafEventsSubscription)`

SetNwdafEventsSubscription sets NwdafEventsSubscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


