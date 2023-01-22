# ACInfoNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubId** | **string** | Identifier of the AC information subscription for which this notification is related to.  | 
**AcInfs** | [**[]ACInformation**](ACInformation.md) | Notifications that include the ACs information matching filter criteria. | 

## Methods

### NewACInfoNotification

`func NewACInfoNotification(subId string, acInfs []ACInformation, ) *ACInfoNotification`

NewACInfoNotification instantiates a new ACInfoNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACInfoNotificationWithDefaults

`func NewACInfoNotificationWithDefaults() *ACInfoNotification`

NewACInfoNotificationWithDefaults instantiates a new ACInfoNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubId

`func (o *ACInfoNotification) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ACInfoNotification) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ACInfoNotification) SetSubId(v string)`

SetSubId sets SubId field to given value.


### GetAcInfs

`func (o *ACInfoNotification) GetAcInfs() []ACInformation`

GetAcInfs returns the AcInfs field if non-nil, zero value otherwise.

### GetAcInfsOk

`func (o *ACInfoNotification) GetAcInfsOk() (*[]ACInformation, bool)`

GetAcInfsOk returns a tuple with the AcInfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcInfs

`func (o *ACInfoNotification) SetAcInfs(v []ACInformation)`

SetAcInfs sets AcInfs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


