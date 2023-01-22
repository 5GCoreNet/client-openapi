# MBSUserDataIngStatNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsIngSessionId** | **string** |  | 
**EventNotifs** | [**[]EventNotification**](EventNotification.md) |  | 

## Methods

### NewMBSUserDataIngStatNotif

`func NewMBSUserDataIngStatNotif(mbsIngSessionId string, eventNotifs []EventNotification, ) *MBSUserDataIngStatNotif`

NewMBSUserDataIngStatNotif instantiates a new MBSUserDataIngStatNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserDataIngStatNotifWithDefaults

`func NewMBSUserDataIngStatNotifWithDefaults() *MBSUserDataIngStatNotif`

NewMBSUserDataIngStatNotifWithDefaults instantiates a new MBSUserDataIngStatNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsIngSessionId

`func (o *MBSUserDataIngStatNotif) GetMbsIngSessionId() string`

GetMbsIngSessionId returns the MbsIngSessionId field if non-nil, zero value otherwise.

### GetMbsIngSessionIdOk

`func (o *MBSUserDataIngStatNotif) GetMbsIngSessionIdOk() (*string, bool)`

GetMbsIngSessionIdOk returns a tuple with the MbsIngSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsIngSessionId

`func (o *MBSUserDataIngStatNotif) SetMbsIngSessionId(v string)`

SetMbsIngSessionId sets MbsIngSessionId field to given value.


### GetEventNotifs

`func (o *MBSUserDataIngStatNotif) GetEventNotifs() []EventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *MBSUserDataIngStatNotif) GetEventNotifsOk() (*[]EventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *MBSUserDataIngStatNotif) SetEventNotifs(v []EventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


