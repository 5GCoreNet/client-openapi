# UcmfNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DicEntryId** | **int32** |  | 
**EventType** | [**EventType**](EventType.md) |  | 
**ManAssOpRequestlist** | Pointer to [**ManAssOpRequestlist**](ManAssOpRequestlist.md) |  | [optional] 
**VersionId** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewUcmfNotification

`func NewUcmfNotification(dicEntryId int32, eventType EventType, ) *UcmfNotification`

NewUcmfNotification instantiates a new UcmfNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUcmfNotificationWithDefaults

`func NewUcmfNotificationWithDefaults() *UcmfNotification`

NewUcmfNotificationWithDefaults instantiates a new UcmfNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDicEntryId

`func (o *UcmfNotification) GetDicEntryId() int32`

GetDicEntryId returns the DicEntryId field if non-nil, zero value otherwise.

### GetDicEntryIdOk

`func (o *UcmfNotification) GetDicEntryIdOk() (*int32, bool)`

GetDicEntryIdOk returns a tuple with the DicEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDicEntryId

`func (o *UcmfNotification) SetDicEntryId(v int32)`

SetDicEntryId sets DicEntryId field to given value.


### GetEventType

`func (o *UcmfNotification) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *UcmfNotification) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *UcmfNotification) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetManAssOpRequestlist

`func (o *UcmfNotification) GetManAssOpRequestlist() ManAssOpRequestlist`

GetManAssOpRequestlist returns the ManAssOpRequestlist field if non-nil, zero value otherwise.

### GetManAssOpRequestlistOk

`func (o *UcmfNotification) GetManAssOpRequestlistOk() (*ManAssOpRequestlist, bool)`

GetManAssOpRequestlistOk returns a tuple with the ManAssOpRequestlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManAssOpRequestlist

`func (o *UcmfNotification) SetManAssOpRequestlist(v ManAssOpRequestlist)`

SetManAssOpRequestlist sets ManAssOpRequestlist field to given value.

### HasManAssOpRequestlist

`func (o *UcmfNotification) HasManAssOpRequestlist() bool`

HasManAssOpRequestlist returns a boolean if a field has been set.

### GetVersionId

`func (o *UcmfNotification) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UcmfNotification) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UcmfNotification) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *UcmfNotification) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


