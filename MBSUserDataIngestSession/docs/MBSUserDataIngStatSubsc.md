# MBSUserDataIngStatSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsIngSessionId** | **string** |  | 
**EventSubscs** | [**[]SubscribedEvent**](SubscribedEvent.md) |  | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 

## Methods

### NewMBSUserDataIngStatSubsc

`func NewMBSUserDataIngStatSubsc(mbsIngSessionId string, eventSubscs []SubscribedEvent, notifUri string, ) *MBSUserDataIngStatSubsc`

NewMBSUserDataIngStatSubsc instantiates a new MBSUserDataIngStatSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserDataIngStatSubscWithDefaults

`func NewMBSUserDataIngStatSubscWithDefaults() *MBSUserDataIngStatSubsc`

NewMBSUserDataIngStatSubscWithDefaults instantiates a new MBSUserDataIngStatSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsIngSessionId

`func (o *MBSUserDataIngStatSubsc) GetMbsIngSessionId() string`

GetMbsIngSessionId returns the MbsIngSessionId field if non-nil, zero value otherwise.

### GetMbsIngSessionIdOk

`func (o *MBSUserDataIngStatSubsc) GetMbsIngSessionIdOk() (*string, bool)`

GetMbsIngSessionIdOk returns a tuple with the MbsIngSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsIngSessionId

`func (o *MBSUserDataIngStatSubsc) SetMbsIngSessionId(v string)`

SetMbsIngSessionId sets MbsIngSessionId field to given value.


### GetEventSubscs

`func (o *MBSUserDataIngStatSubsc) GetEventSubscs() []SubscribedEvent`

GetEventSubscs returns the EventSubscs field if non-nil, zero value otherwise.

### GetEventSubscsOk

`func (o *MBSUserDataIngStatSubsc) GetEventSubscsOk() (*[]SubscribedEvent, bool)`

GetEventSubscsOk returns a tuple with the EventSubscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscs

`func (o *MBSUserDataIngStatSubsc) SetEventSubscs(v []SubscribedEvent)`

SetEventSubscs sets EventSubscs field to given value.


### GetNotifUri

`func (o *MBSUserDataIngStatSubsc) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *MBSUserDataIngStatSubsc) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *MBSUserDataIngStatSubsc) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


