# ACInfoSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasId** | **string** | Identifier of the EAS subscribing for AC information report. | 
**AcFltrs** | Pointer to [**[]ACFilters**](ACFilters.md) | Filters to retrieve the information about specific ACs. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the EAS to request the EES to send a test notification. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewACInfoSubscription

`func NewACInfoSubscription(easId string, ) *ACInfoSubscription`

NewACInfoSubscription instantiates a new ACInfoSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACInfoSubscriptionWithDefaults

`func NewACInfoSubscriptionWithDefaults() *ACInfoSubscription`

NewACInfoSubscriptionWithDefaults instantiates a new ACInfoSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasId

`func (o *ACInfoSubscription) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *ACInfoSubscription) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *ACInfoSubscription) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetAcFltrs

`func (o *ACInfoSubscription) GetAcFltrs() []ACFilters`

GetAcFltrs returns the AcFltrs field if non-nil, zero value otherwise.

### GetAcFltrsOk

`func (o *ACInfoSubscription) GetAcFltrsOk() (*[]ACFilters, bool)`

GetAcFltrsOk returns a tuple with the AcFltrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcFltrs

`func (o *ACInfoSubscription) SetAcFltrs(v []ACFilters)`

SetAcFltrs sets AcFltrs field to given value.

### HasAcFltrs

`func (o *ACInfoSubscription) HasAcFltrs() bool`

HasAcFltrs returns a boolean if a field has been set.

### GetExpTime

`func (o *ACInfoSubscription) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *ACInfoSubscription) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *ACInfoSubscription) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *ACInfoSubscription) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetEventReq

`func (o *ACInfoSubscription) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *ACInfoSubscription) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *ACInfoSubscription) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *ACInfoSubscription) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *ACInfoSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ACInfoSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ACInfoSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *ACInfoSubscription) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *ACInfoSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *ACInfoSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *ACInfoSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *ACInfoSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *ACInfoSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *ACInfoSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *ACInfoSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *ACInfoSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ACInfoSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ACInfoSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ACInfoSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ACInfoSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


