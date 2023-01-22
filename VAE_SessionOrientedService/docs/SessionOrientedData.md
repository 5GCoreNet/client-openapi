# SessionOrientedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | **string** | Represents the identifier of the V2X UE. | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**ServiceId** | **string** | Represents the V2X service ID to which a V2X message belongs. | 
**AppSerId** | **string** | Represents the V2X application specific server identifier. | 
**AppQosReq** | Pointer to [**AppplicationQosRequirement**](AppplicationQosRequirement.md) |  | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the NF service consumer to request the VAE server to send a test notification as defined in clause 6.3.5.3. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSessionOrientedData

`func NewSessionOrientedData(ueId string, notifUri string, serviceId string, appSerId string, ) *SessionOrientedData`

NewSessionOrientedData instantiates a new SessionOrientedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionOrientedDataWithDefaults

`func NewSessionOrientedDataWithDefaults() *SessionOrientedData`

NewSessionOrientedDataWithDefaults instantiates a new SessionOrientedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *SessionOrientedData) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *SessionOrientedData) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *SessionOrientedData) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetNotifUri

`func (o *SessionOrientedData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *SessionOrientedData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *SessionOrientedData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetServiceId

`func (o *SessionOrientedData) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SessionOrientedData) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SessionOrientedData) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetAppSerId

`func (o *SessionOrientedData) GetAppSerId() string`

GetAppSerId returns the AppSerId field if non-nil, zero value otherwise.

### GetAppSerIdOk

`func (o *SessionOrientedData) GetAppSerIdOk() (*string, bool)`

GetAppSerIdOk returns a tuple with the AppSerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSerId

`func (o *SessionOrientedData) SetAppSerId(v string)`

SetAppSerId sets AppSerId field to given value.


### GetAppQosReq

`func (o *SessionOrientedData) GetAppQosReq() AppplicationQosRequirement`

GetAppQosReq returns the AppQosReq field if non-nil, zero value otherwise.

### GetAppQosReqOk

`func (o *SessionOrientedData) GetAppQosReqOk() (*AppplicationQosRequirement, bool)`

GetAppQosReqOk returns a tuple with the AppQosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppQosReq

`func (o *SessionOrientedData) SetAppQosReq(v AppplicationQosRequirement)`

SetAppQosReq sets AppQosReq field to given value.

### HasAppQosReq

`func (o *SessionOrientedData) HasAppQosReq() bool`

HasAppQosReq returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *SessionOrientedData) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *SessionOrientedData) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *SessionOrientedData) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *SessionOrientedData) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *SessionOrientedData) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *SessionOrientedData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *SessionOrientedData) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *SessionOrientedData) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *SessionOrientedData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *SessionOrientedData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *SessionOrientedData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *SessionOrientedData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


