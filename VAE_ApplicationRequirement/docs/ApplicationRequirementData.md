# ApplicationRequirementData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | Pointer to **string** | Represents the identifier of the V2X UE. | [optional] 
**GroupId** | Pointer to **string** | Represents the group ID for which a V2X message is addressed. | [optional] 
**Duration** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ServiceId** | **string** | Represents the V2X service ID to which a V2X message belongs. | 
**AppRequirement** | [**ApplicationRequirement**](ApplicationRequirement.md) |  | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the NF service consumer to request the VAE server to send a test notification as defined in clause 6.3.5.3. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewApplicationRequirementData

`func NewApplicationRequirementData(serviceId string, appRequirement ApplicationRequirement, notifUri string, ) *ApplicationRequirementData`

NewApplicationRequirementData instantiates a new ApplicationRequirementData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequirementDataWithDefaults

`func NewApplicationRequirementDataWithDefaults() *ApplicationRequirementData`

NewApplicationRequirementDataWithDefaults instantiates a new ApplicationRequirementData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *ApplicationRequirementData) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ApplicationRequirementData) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ApplicationRequirementData) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *ApplicationRequirementData) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetGroupId

`func (o *ApplicationRequirementData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApplicationRequirementData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApplicationRequirementData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApplicationRequirementData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDuration

`func (o *ApplicationRequirementData) GetDuration() time.Time`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApplicationRequirementData) GetDurationOk() (*time.Time, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApplicationRequirementData) SetDuration(v time.Time)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApplicationRequirementData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetServiceId

`func (o *ApplicationRequirementData) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ApplicationRequirementData) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ApplicationRequirementData) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetAppRequirement

`func (o *ApplicationRequirementData) GetAppRequirement() ApplicationRequirement`

GetAppRequirement returns the AppRequirement field if non-nil, zero value otherwise.

### GetAppRequirementOk

`func (o *ApplicationRequirementData) GetAppRequirementOk() (*ApplicationRequirement, bool)`

GetAppRequirementOk returns a tuple with the AppRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRequirement

`func (o *ApplicationRequirementData) SetAppRequirement(v ApplicationRequirement)`

SetAppRequirement sets AppRequirement field to given value.


### GetNotifUri

`func (o *ApplicationRequirementData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *ApplicationRequirementData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *ApplicationRequirementData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetRequestTestNotification

`func (o *ApplicationRequirementData) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *ApplicationRequirementData) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *ApplicationRequirementData) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *ApplicationRequirementData) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *ApplicationRequirementData) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *ApplicationRequirementData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *ApplicationRequirementData) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *ApplicationRequirementData) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ApplicationRequirementData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ApplicationRequirementData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ApplicationRequirementData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ApplicationRequirementData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


