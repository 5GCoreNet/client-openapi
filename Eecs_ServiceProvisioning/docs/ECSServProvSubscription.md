# ECSServProvSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EecId** | **string** | Represents a unique identifier of the EEC. | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AcProfs** | Pointer to [**[]ACProfile**](ACProfile.md) | Information about services the EEC wants to connect to. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EecSvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.  | [optional] 
**ConnInfo** | Pointer to [**[]ConnectivityInfo**](ConnectivityInfo.md) | List of connectivity information for the UE. | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by Subscriber to request the ECS to send a test notification. Set to  false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewECSServProvSubscription

`func NewECSServProvSubscription(eecId string, ) *ECSServProvSubscription`

NewECSServProvSubscription instantiates a new ECSServProvSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSServProvSubscriptionWithDefaults

`func NewECSServProvSubscriptionWithDefaults() *ECSServProvSubscription`

NewECSServProvSubscriptionWithDefaults instantiates a new ECSServProvSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEecId

`func (o *ECSServProvSubscription) GetEecId() string`

GetEecId returns the EecId field if non-nil, zero value otherwise.

### GetEecIdOk

`func (o *ECSServProvSubscription) GetEecIdOk() (*string, bool)`

GetEecIdOk returns a tuple with the EecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecId

`func (o *ECSServProvSubscription) SetEecId(v string)`

SetEecId sets EecId field to given value.


### GetUeId

`func (o *ECSServProvSubscription) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ECSServProvSubscription) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ECSServProvSubscription) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *ECSServProvSubscription) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetAcProfs

`func (o *ECSServProvSubscription) GetAcProfs() []ACProfile`

GetAcProfs returns the AcProfs field if non-nil, zero value otherwise.

### GetAcProfsOk

`func (o *ECSServProvSubscription) GetAcProfsOk() (*[]ACProfile, bool)`

GetAcProfsOk returns a tuple with the AcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcProfs

`func (o *ECSServProvSubscription) SetAcProfs(v []ACProfile)`

SetAcProfs sets AcProfs field to given value.

### HasAcProfs

`func (o *ECSServProvSubscription) HasAcProfs() bool`

HasAcProfs returns a boolean if a field has been set.

### GetExpTime

`func (o *ECSServProvSubscription) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *ECSServProvSubscription) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *ECSServProvSubscription) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *ECSServProvSubscription) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetEecSvcContSupp

`func (o *ECSServProvSubscription) GetEecSvcContSupp() []ACRScenario`

GetEecSvcContSupp returns the EecSvcContSupp field if non-nil, zero value otherwise.

### GetEecSvcContSuppOk

`func (o *ECSServProvSubscription) GetEecSvcContSuppOk() (*[]ACRScenario, bool)`

GetEecSvcContSuppOk returns a tuple with the EecSvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecSvcContSupp

`func (o *ECSServProvSubscription) SetEecSvcContSupp(v []ACRScenario)`

SetEecSvcContSupp sets EecSvcContSupp field to given value.

### HasEecSvcContSupp

`func (o *ECSServProvSubscription) HasEecSvcContSupp() bool`

HasEecSvcContSupp returns a boolean if a field has been set.

### GetConnInfo

`func (o *ECSServProvSubscription) GetConnInfo() []ConnectivityInfo`

GetConnInfo returns the ConnInfo field if non-nil, zero value otherwise.

### GetConnInfoOk

`func (o *ECSServProvSubscription) GetConnInfoOk() (*[]ConnectivityInfo, bool)`

GetConnInfoOk returns a tuple with the ConnInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnInfo

`func (o *ECSServProvSubscription) SetConnInfo(v []ConnectivityInfo)`

SetConnInfo sets ConnInfo field to given value.

### HasConnInfo

`func (o *ECSServProvSubscription) HasConnInfo() bool`

HasConnInfo returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *ECSServProvSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ECSServProvSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ECSServProvSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *ECSServProvSubscription) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *ECSServProvSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *ECSServProvSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *ECSServProvSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *ECSServProvSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *ECSServProvSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *ECSServProvSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *ECSServProvSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *ECSServProvSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ECSServProvSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ECSServProvSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ECSServProvSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ECSServProvSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


