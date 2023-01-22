# LocationSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasId** | **string** | Identifier of the EAS subscribing for location information report. | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**IntGrpId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**ExtGrpId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**LocGran** | Pointer to [**Accuracy**](Accuracy.md) |  | [optional] 
**LocQos** | Pointer to [**LocationQoS**](LocationQoS.md) |  | [optional] 
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the EAS to request the EES to send a test notification. Set to false or omitted otherwise.  | [optional] 
**RevocationNotifUri** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewLocationSubscription

`func NewLocationSubscription(easId string, ) *LocationSubscription`

NewLocationSubscription instantiates a new LocationSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationSubscriptionWithDefaults

`func NewLocationSubscriptionWithDefaults() *LocationSubscription`

NewLocationSubscriptionWithDefaults instantiates a new LocationSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasId

`func (o *LocationSubscription) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *LocationSubscription) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *LocationSubscription) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetUeId

`func (o *LocationSubscription) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *LocationSubscription) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *LocationSubscription) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *LocationSubscription) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetIntGrpId

`func (o *LocationSubscription) GetIntGrpId() string`

GetIntGrpId returns the IntGrpId field if non-nil, zero value otherwise.

### GetIntGrpIdOk

`func (o *LocationSubscription) GetIntGrpIdOk() (*string, bool)`

GetIntGrpIdOk returns a tuple with the IntGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGrpId

`func (o *LocationSubscription) SetIntGrpId(v string)`

SetIntGrpId sets IntGrpId field to given value.

### HasIntGrpId

`func (o *LocationSubscription) HasIntGrpId() bool`

HasIntGrpId returns a boolean if a field has been set.

### GetExtGrpId

`func (o *LocationSubscription) GetExtGrpId() string`

GetExtGrpId returns the ExtGrpId field if non-nil, zero value otherwise.

### GetExtGrpIdOk

`func (o *LocationSubscription) GetExtGrpIdOk() (*string, bool)`

GetExtGrpIdOk returns a tuple with the ExtGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGrpId

`func (o *LocationSubscription) SetExtGrpId(v string)`

SetExtGrpId sets ExtGrpId field to given value.

### HasExtGrpId

`func (o *LocationSubscription) HasExtGrpId() bool`

HasExtGrpId returns a boolean if a field has been set.

### GetExpTime

`func (o *LocationSubscription) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *LocationSubscription) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *LocationSubscription) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *LocationSubscription) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetLocGran

`func (o *LocationSubscription) GetLocGran() Accuracy`

GetLocGran returns the LocGran field if non-nil, zero value otherwise.

### GetLocGranOk

`func (o *LocationSubscription) GetLocGranOk() (*Accuracy, bool)`

GetLocGranOk returns a tuple with the LocGran field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocGran

`func (o *LocationSubscription) SetLocGran(v Accuracy)`

SetLocGran sets LocGran field to given value.

### HasLocGran

`func (o *LocationSubscription) HasLocGran() bool`

HasLocGran returns a boolean if a field has been set.

### GetLocQos

`func (o *LocationSubscription) GetLocQos() LocationQoS`

GetLocQos returns the LocQos field if non-nil, zero value otherwise.

### GetLocQosOk

`func (o *LocationSubscription) GetLocQosOk() (*LocationQoS, bool)`

GetLocQosOk returns a tuple with the LocQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocQos

`func (o *LocationSubscription) SetLocQos(v LocationQoS)`

SetLocQos sets LocQos field to given value.

### HasLocQos

`func (o *LocationSubscription) HasLocQos() bool`

HasLocQos returns a boolean if a field has been set.

### GetEventReq

`func (o *LocationSubscription) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *LocationSubscription) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *LocationSubscription) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *LocationSubscription) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *LocationSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *LocationSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *LocationSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *LocationSubscription) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *LocationSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *LocationSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *LocationSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *LocationSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetRevocationNotifUri

`func (o *LocationSubscription) GetRevocationNotifUri() string`

GetRevocationNotifUri returns the RevocationNotifUri field if non-nil, zero value otherwise.

### GetRevocationNotifUriOk

`func (o *LocationSubscription) GetRevocationNotifUriOk() (*string, bool)`

GetRevocationNotifUriOk returns a tuple with the RevocationNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationNotifUri

`func (o *LocationSubscription) SetRevocationNotifUri(v string)`

SetRevocationNotifUri sets RevocationNotifUri field to given value.

### HasRevocationNotifUri

`func (o *LocationSubscription) HasRevocationNotifUri() bool`

HasRevocationNotifUri returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *LocationSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *LocationSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *LocationSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *LocationSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *LocationSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *LocationSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *LocationSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *LocationSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


