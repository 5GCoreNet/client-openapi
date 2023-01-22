# TmgiAllocRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | **string** |  | 
**TmgiParams** | [**TmgiAllocate**](TmgiAllocate.md) |  | 
**NotificationUri** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**RequestTestNotification** | Pointer to **bool** |  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewTmgiAllocRequest

`func NewTmgiAllocRequest(afId string, tmgiParams TmgiAllocate, ) *TmgiAllocRequest`

NewTmgiAllocRequest instantiates a new TmgiAllocRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTmgiAllocRequestWithDefaults

`func NewTmgiAllocRequestWithDefaults() *TmgiAllocRequest`

NewTmgiAllocRequestWithDefaults instantiates a new TmgiAllocRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *TmgiAllocRequest) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *TmgiAllocRequest) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *TmgiAllocRequest) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetTmgiParams

`func (o *TmgiAllocRequest) GetTmgiParams() TmgiAllocate`

GetTmgiParams returns the TmgiParams field if non-nil, zero value otherwise.

### GetTmgiParamsOk

`func (o *TmgiAllocRequest) GetTmgiParamsOk() (*TmgiAllocate, bool)`

GetTmgiParamsOk returns a tuple with the TmgiParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgiParams

`func (o *TmgiAllocRequest) SetTmgiParams(v TmgiAllocate)`

SetTmgiParams sets TmgiParams field to given value.


### GetNotificationUri

`func (o *TmgiAllocRequest) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *TmgiAllocRequest) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *TmgiAllocRequest) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.

### HasNotificationUri

`func (o *TmgiAllocRequest) HasNotificationUri() bool`

HasNotificationUri returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *TmgiAllocRequest) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *TmgiAllocRequest) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *TmgiAllocRequest) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *TmgiAllocRequest) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *TmgiAllocRequest) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *TmgiAllocRequest) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *TmgiAllocRequest) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *TmgiAllocRequest) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *TmgiAllocRequest) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *TmgiAllocRequest) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *TmgiAllocRequest) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *TmgiAllocRequest) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


