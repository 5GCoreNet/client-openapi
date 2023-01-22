# ReachabilityNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReachableUeList** | Pointer to [**[]ReachableUeInfo**](ReachableUeInfo.md) |  | [optional] 
**UnreachableUeList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewReachabilityNotificationData

`func NewReachabilityNotificationData() *ReachabilityNotificationData`

NewReachabilityNotificationData instantiates a new ReachabilityNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachabilityNotificationDataWithDefaults

`func NewReachabilityNotificationDataWithDefaults() *ReachabilityNotificationData`

NewReachabilityNotificationDataWithDefaults instantiates a new ReachabilityNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachableUeList

`func (o *ReachabilityNotificationData) GetReachableUeList() []ReachableUeInfo`

GetReachableUeList returns the ReachableUeList field if non-nil, zero value otherwise.

### GetReachableUeListOk

`func (o *ReachabilityNotificationData) GetReachableUeListOk() (*[]ReachableUeInfo, bool)`

GetReachableUeListOk returns a tuple with the ReachableUeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableUeList

`func (o *ReachabilityNotificationData) SetReachableUeList(v []ReachableUeInfo)`

SetReachableUeList sets ReachableUeList field to given value.

### HasReachableUeList

`func (o *ReachabilityNotificationData) HasReachableUeList() bool`

HasReachableUeList returns a boolean if a field has been set.

### GetUnreachableUeList

`func (o *ReachabilityNotificationData) GetUnreachableUeList() []string`

GetUnreachableUeList returns the UnreachableUeList field if non-nil, zero value otherwise.

### GetUnreachableUeListOk

`func (o *ReachabilityNotificationData) GetUnreachableUeListOk() (*[]string, bool)`

GetUnreachableUeListOk returns a tuple with the UnreachableUeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreachableUeList

`func (o *ReachabilityNotificationData) SetUnreachableUeList(v []string)`

SetUnreachableUeList sets UnreachableUeList field to given value.

### HasUnreachableUeList

`func (o *ReachabilityNotificationData) HasUnreachableUeList() bool`

HasUnreachableUeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


