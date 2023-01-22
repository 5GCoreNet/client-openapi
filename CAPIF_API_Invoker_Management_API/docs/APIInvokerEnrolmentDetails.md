# APIInvokerEnrolmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInvokerId** | Pointer to **string** | API invoker ID assigned by the CAPIF core function to the API invoker while on-boarding the API invoker. Shall not be present in the HTTP POST request from the API invoker to the CAPIF core function, to on-board itself. Shall be present in all other HTTP requests and responses.  | [optional] [readonly] 
**OnboardingInformation** | [**OnboardingInformation**](OnboardingInformation.md) |  | 
**NotificationDestination** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by Subscriber to request the CAPIF core function to send a test notification as defined in in clause 7.6. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**ApiList** | Pointer to [**[]ServiceAPIDescription**](ServiceAPIDescription.md) | The list of service APIs that the API Invoker is allowed to invoke | [optional] 
**ApiInvokerInformation** | Pointer to **string** | Generic information related to the API invoker such as details of the device or the application.  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAPIInvokerEnrolmentDetails

`func NewAPIInvokerEnrolmentDetails(onboardingInformation OnboardingInformation, notificationDestination string, ) *APIInvokerEnrolmentDetails`

NewAPIInvokerEnrolmentDetails instantiates a new APIInvokerEnrolmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIInvokerEnrolmentDetailsWithDefaults

`func NewAPIInvokerEnrolmentDetailsWithDefaults() *APIInvokerEnrolmentDetails`

NewAPIInvokerEnrolmentDetailsWithDefaults instantiates a new APIInvokerEnrolmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInvokerId

`func (o *APIInvokerEnrolmentDetails) GetApiInvokerId() string`

GetApiInvokerId returns the ApiInvokerId field if non-nil, zero value otherwise.

### GetApiInvokerIdOk

`func (o *APIInvokerEnrolmentDetails) GetApiInvokerIdOk() (*string, bool)`

GetApiInvokerIdOk returns a tuple with the ApiInvokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerId

`func (o *APIInvokerEnrolmentDetails) SetApiInvokerId(v string)`

SetApiInvokerId sets ApiInvokerId field to given value.

### HasApiInvokerId

`func (o *APIInvokerEnrolmentDetails) HasApiInvokerId() bool`

HasApiInvokerId returns a boolean if a field has been set.

### GetOnboardingInformation

`func (o *APIInvokerEnrolmentDetails) GetOnboardingInformation() OnboardingInformation`

GetOnboardingInformation returns the OnboardingInformation field if non-nil, zero value otherwise.

### GetOnboardingInformationOk

`func (o *APIInvokerEnrolmentDetails) GetOnboardingInformationOk() (*OnboardingInformation, bool)`

GetOnboardingInformationOk returns a tuple with the OnboardingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingInformation

`func (o *APIInvokerEnrolmentDetails) SetOnboardingInformation(v OnboardingInformation)`

SetOnboardingInformation sets OnboardingInformation field to given value.


### GetNotificationDestination

`func (o *APIInvokerEnrolmentDetails) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *APIInvokerEnrolmentDetails) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *APIInvokerEnrolmentDetails) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *APIInvokerEnrolmentDetails) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *APIInvokerEnrolmentDetails) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *APIInvokerEnrolmentDetails) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *APIInvokerEnrolmentDetails) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *APIInvokerEnrolmentDetails) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *APIInvokerEnrolmentDetails) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *APIInvokerEnrolmentDetails) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *APIInvokerEnrolmentDetails) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetApiList

`func (o *APIInvokerEnrolmentDetails) GetApiList() []ServiceAPIDescription`

GetApiList returns the ApiList field if non-nil, zero value otherwise.

### GetApiListOk

`func (o *APIInvokerEnrolmentDetails) GetApiListOk() (*[]ServiceAPIDescription, bool)`

GetApiListOk returns a tuple with the ApiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiList

`func (o *APIInvokerEnrolmentDetails) SetApiList(v []ServiceAPIDescription)`

SetApiList sets ApiList field to given value.

### HasApiList

`func (o *APIInvokerEnrolmentDetails) HasApiList() bool`

HasApiList returns a boolean if a field has been set.

### GetApiInvokerInformation

`func (o *APIInvokerEnrolmentDetails) GetApiInvokerInformation() string`

GetApiInvokerInformation returns the ApiInvokerInformation field if non-nil, zero value otherwise.

### GetApiInvokerInformationOk

`func (o *APIInvokerEnrolmentDetails) GetApiInvokerInformationOk() (*string, bool)`

GetApiInvokerInformationOk returns a tuple with the ApiInvokerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerInformation

`func (o *APIInvokerEnrolmentDetails) SetApiInvokerInformation(v string)`

SetApiInvokerInformation sets ApiInvokerInformation field to given value.

### HasApiInvokerInformation

`func (o *APIInvokerEnrolmentDetails) HasApiInvokerInformation() bool`

HasApiInvokerInformation returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *APIInvokerEnrolmentDetails) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *APIInvokerEnrolmentDetails) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *APIInvokerEnrolmentDetails) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *APIInvokerEnrolmentDetails) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


