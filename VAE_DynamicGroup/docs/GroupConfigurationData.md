# GroupConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Represents the group ID for which a V2X message is addressed. | 
**Definition** | **string** |  | 
**LeaderId** | **string** | Represents the identifier of the V2X UE. | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Duration** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the NF service consumer to request the VAE server to test a notification connection. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewGroupConfigurationData

`func NewGroupConfigurationData(groupId string, definition string, leaderId string, notifUri string, ) *GroupConfigurationData`

NewGroupConfigurationData instantiates a new GroupConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupConfigurationDataWithDefaults

`func NewGroupConfigurationDataWithDefaults() *GroupConfigurationData`

NewGroupConfigurationDataWithDefaults instantiates a new GroupConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupConfigurationData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupConfigurationData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupConfigurationData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetDefinition

`func (o *GroupConfigurationData) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *GroupConfigurationData) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *GroupConfigurationData) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetLeaderId

`func (o *GroupConfigurationData) GetLeaderId() string`

GetLeaderId returns the LeaderId field if non-nil, zero value otherwise.

### GetLeaderIdOk

`func (o *GroupConfigurationData) GetLeaderIdOk() (*string, bool)`

GetLeaderIdOk returns a tuple with the LeaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderId

`func (o *GroupConfigurationData) SetLeaderId(v string)`

SetLeaderId sets LeaderId field to given value.


### GetNotifUri

`func (o *GroupConfigurationData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *GroupConfigurationData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *GroupConfigurationData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetDuration

`func (o *GroupConfigurationData) GetDuration() time.Time`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GroupConfigurationData) GetDurationOk() (*time.Time, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GroupConfigurationData) SetDuration(v time.Time)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GroupConfigurationData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *GroupConfigurationData) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *GroupConfigurationData) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *GroupConfigurationData) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *GroupConfigurationData) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *GroupConfigurationData) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *GroupConfigurationData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *GroupConfigurationData) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *GroupConfigurationData) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *GroupConfigurationData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *GroupConfigurationData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *GroupConfigurationData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *GroupConfigurationData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


