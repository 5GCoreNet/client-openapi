# DynamicGroupNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**JoinedUeIds** | Pointer to **[]string** |  | [optional] 
**LeftUeIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDynamicGroupNotification

`func NewDynamicGroupNotification(resourceUri string, ) *DynamicGroupNotification`

NewDynamicGroupNotification instantiates a new DynamicGroupNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicGroupNotificationWithDefaults

`func NewDynamicGroupNotificationWithDefaults() *DynamicGroupNotification`

NewDynamicGroupNotificationWithDefaults instantiates a new DynamicGroupNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *DynamicGroupNotification) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *DynamicGroupNotification) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *DynamicGroupNotification) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetJoinedUeIds

`func (o *DynamicGroupNotification) GetJoinedUeIds() []string`

GetJoinedUeIds returns the JoinedUeIds field if non-nil, zero value otherwise.

### GetJoinedUeIdsOk

`func (o *DynamicGroupNotification) GetJoinedUeIdsOk() (*[]string, bool)`

GetJoinedUeIdsOk returns a tuple with the JoinedUeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedUeIds

`func (o *DynamicGroupNotification) SetJoinedUeIds(v []string)`

SetJoinedUeIds sets JoinedUeIds field to given value.

### HasJoinedUeIds

`func (o *DynamicGroupNotification) HasJoinedUeIds() bool`

HasJoinedUeIds returns a boolean if a field has been set.

### GetLeftUeIds

`func (o *DynamicGroupNotification) GetLeftUeIds() []string`

GetLeftUeIds returns the LeftUeIds field if non-nil, zero value otherwise.

### GetLeftUeIdsOk

`func (o *DynamicGroupNotification) GetLeftUeIdsOk() (*[]string, bool)`

GetLeftUeIdsOk returns a tuple with the LeftUeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftUeIds

`func (o *DynamicGroupNotification) SetLeftUeIds(v []string)`

SetLeftUeIds sets LeftUeIds field to given value.

### HasLeftUeIds

`func (o *DynamicGroupNotification) HasLeftUeIds() bool`

HasLeftUeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


