# HdMapDynamicInfoData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | **string** | Represents the identifier of the V2X UE. | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Range** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the NF service consumer to request the VAE server to send a test notification as defined in clause 6.3.5.3. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewHdMapDynamicInfoData

`func NewHdMapDynamicInfoData(ueId string, notifUri string, range_ int32, ) *HdMapDynamicInfoData`

NewHdMapDynamicInfoData instantiates a new HdMapDynamicInfoData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdMapDynamicInfoDataWithDefaults

`func NewHdMapDynamicInfoDataWithDefaults() *HdMapDynamicInfoData`

NewHdMapDynamicInfoDataWithDefaults instantiates a new HdMapDynamicInfoData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *HdMapDynamicInfoData) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *HdMapDynamicInfoData) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *HdMapDynamicInfoData) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetNotifUri

`func (o *HdMapDynamicInfoData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *HdMapDynamicInfoData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *HdMapDynamicInfoData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetRange

`func (o *HdMapDynamicInfoData) GetRange() int32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *HdMapDynamicInfoData) GetRangeOk() (*int32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *HdMapDynamicInfoData) SetRange(v int32)`

SetRange sets Range field to given value.


### GetRequestTestNotification

`func (o *HdMapDynamicInfoData) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *HdMapDynamicInfoData) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *HdMapDynamicInfoData) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *HdMapDynamicInfoData) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *HdMapDynamicInfoData) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *HdMapDynamicInfoData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *HdMapDynamicInfoData) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *HdMapDynamicInfoData) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *HdMapDynamicInfoData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *HdMapDynamicInfoData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *HdMapDynamicInfoData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *HdMapDynamicInfoData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


