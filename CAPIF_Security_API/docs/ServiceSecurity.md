# ServiceSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityInfo** | [**[]SecurityInformation**](SecurityInformation.md) |  | 
**NotificationDestination** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by API invoker to request the CAPIF core function to send a test notification as defined in in clause 7.6. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewServiceSecurity

`func NewServiceSecurity(securityInfo []SecurityInformation, notificationDestination string, ) *ServiceSecurity`

NewServiceSecurity instantiates a new ServiceSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSecurityWithDefaults

`func NewServiceSecurityWithDefaults() *ServiceSecurity`

NewServiceSecurityWithDefaults instantiates a new ServiceSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityInfo

`func (o *ServiceSecurity) GetSecurityInfo() []SecurityInformation`

GetSecurityInfo returns the SecurityInfo field if non-nil, zero value otherwise.

### GetSecurityInfoOk

`func (o *ServiceSecurity) GetSecurityInfoOk() (*[]SecurityInformation, bool)`

GetSecurityInfoOk returns a tuple with the SecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityInfo

`func (o *ServiceSecurity) SetSecurityInfo(v []SecurityInformation)`

SetSecurityInfo sets SecurityInfo field to given value.


### GetNotificationDestination

`func (o *ServiceSecurity) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ServiceSecurity) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ServiceSecurity) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *ServiceSecurity) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *ServiceSecurity) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *ServiceSecurity) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *ServiceSecurity) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *ServiceSecurity) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *ServiceSecurity) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *ServiceSecurity) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *ServiceSecurity) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ServiceSecurity) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ServiceSecurity) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ServiceSecurity) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ServiceSecurity) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


