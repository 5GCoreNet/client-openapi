# NadrfDataStoreRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataNotif** | Pointer to [**DataNotification**](DataNotification.md) |  | [optional] 
**AnaNotifications** | Pointer to [**[]NnwdafEventsSubscriptionNotification**](NnwdafEventsSubscriptionNotification.md) | List of analytics subscription notifications. | [optional] 
**AnaSub** | Pointer to [**[]NnwdafEventsSubscription**](NnwdafEventsSubscription.md) | Represents the subscription information of the corresponding analytics notification.  | [optional] 
**DataSub** | Pointer to [**[]DataSubscription**](DataSubscription.md) | Represents the subscription information of the corresponding data notification.  | [optional] 

## Methods

### NewNadrfDataStoreRecord

`func NewNadrfDataStoreRecord() *NadrfDataStoreRecord`

NewNadrfDataStoreRecord instantiates a new NadrfDataStoreRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNadrfDataStoreRecordWithDefaults

`func NewNadrfDataStoreRecordWithDefaults() *NadrfDataStoreRecord`

NewNadrfDataStoreRecordWithDefaults instantiates a new NadrfDataStoreRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataNotif

`func (o *NadrfDataStoreRecord) GetDataNotif() DataNotification`

GetDataNotif returns the DataNotif field if non-nil, zero value otherwise.

### GetDataNotifOk

`func (o *NadrfDataStoreRecord) GetDataNotifOk() (*DataNotification, bool)`

GetDataNotifOk returns a tuple with the DataNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNotif

`func (o *NadrfDataStoreRecord) SetDataNotif(v DataNotification)`

SetDataNotif sets DataNotif field to given value.

### HasDataNotif

`func (o *NadrfDataStoreRecord) HasDataNotif() bool`

HasDataNotif returns a boolean if a field has been set.

### GetAnaNotifications

`func (o *NadrfDataStoreRecord) GetAnaNotifications() []NnwdafEventsSubscriptionNotification`

GetAnaNotifications returns the AnaNotifications field if non-nil, zero value otherwise.

### GetAnaNotificationsOk

`func (o *NadrfDataStoreRecord) GetAnaNotificationsOk() (*[]NnwdafEventsSubscriptionNotification, bool)`

GetAnaNotificationsOk returns a tuple with the AnaNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaNotifications

`func (o *NadrfDataStoreRecord) SetAnaNotifications(v []NnwdafEventsSubscriptionNotification)`

SetAnaNotifications sets AnaNotifications field to given value.

### HasAnaNotifications

`func (o *NadrfDataStoreRecord) HasAnaNotifications() bool`

HasAnaNotifications returns a boolean if a field has been set.

### GetAnaSub

`func (o *NadrfDataStoreRecord) GetAnaSub() []NnwdafEventsSubscription`

GetAnaSub returns the AnaSub field if non-nil, zero value otherwise.

### GetAnaSubOk

`func (o *NadrfDataStoreRecord) GetAnaSubOk() (*[]NnwdafEventsSubscription, bool)`

GetAnaSubOk returns a tuple with the AnaSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaSub

`func (o *NadrfDataStoreRecord) SetAnaSub(v []NnwdafEventsSubscription)`

SetAnaSub sets AnaSub field to given value.

### HasAnaSub

`func (o *NadrfDataStoreRecord) HasAnaSub() bool`

HasAnaSub returns a boolean if a field has been set.

### GetDataSub

`func (o *NadrfDataStoreRecord) GetDataSub() []DataSubscription`

GetDataSub returns the DataSub field if non-nil, zero value otherwise.

### GetDataSubOk

`func (o *NadrfDataStoreRecord) GetDataSubOk() (*[]DataSubscription, bool)`

GetDataSubOk returns a tuple with the DataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSub

`func (o *NadrfDataStoreRecord) SetDataSub(v []DataSubscription)`

SetDataSub sets DataSub field to given value.

### HasDataSub

`func (o *NadrfDataStoreRecord) HasDataSub() bool`

HasDataSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


