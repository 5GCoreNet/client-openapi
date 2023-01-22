# ServProvNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubId** | **string** | Identifier of the individual service provisioning subscription for which the service provisioning notification is delivered.  | 
**EdnCnfgInfo** | [**[]EDNConfigInfo**](EDNConfigInfo.md) | List of EDN configuration information. | 

## Methods

### NewServProvNotification

`func NewServProvNotification(subId string, ednCnfgInfo []EDNConfigInfo, ) *ServProvNotification`

NewServProvNotification instantiates a new ServProvNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServProvNotificationWithDefaults

`func NewServProvNotificationWithDefaults() *ServProvNotification`

NewServProvNotificationWithDefaults instantiates a new ServProvNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubId

`func (o *ServProvNotification) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ServProvNotification) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ServProvNotification) SetSubId(v string)`

SetSubId sets SubId field to given value.


### GetEdnCnfgInfo

`func (o *ServProvNotification) GetEdnCnfgInfo() []EDNConfigInfo`

GetEdnCnfgInfo returns the EdnCnfgInfo field if non-nil, zero value otherwise.

### GetEdnCnfgInfoOk

`func (o *ServProvNotification) GetEdnCnfgInfoOk() (*[]EDNConfigInfo, bool)`

GetEdnCnfgInfoOk returns a tuple with the EdnCnfgInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnCnfgInfo

`func (o *ServProvNotification) SetEdnCnfgInfo(v []EDNConfigInfo)`

SetEdnCnfgInfo sets EdnCnfgInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


