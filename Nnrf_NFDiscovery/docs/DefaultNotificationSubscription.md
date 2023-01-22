# DefaultNotificationSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**CallbackUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**N1MessageClass** | Pointer to [**N1MessageClass**](N1MessageClass.md) |  | [optional] 
**N2InformationClass** | Pointer to [**N2InformationClass**](N2InformationClass.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 
**Binding** | Pointer to **string** |  | [optional] 
**AcceptedEncoding** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ServiceInfoList** | Pointer to [**map[string]DefSubServiceInfo**](DefSubServiceInfo.md) | A map of service specific information. The name of the corresponding service (as specified in ServiceName data type) is the key.  | [optional] 

## Methods

### NewDefaultNotificationSubscription

`func NewDefaultNotificationSubscription(notificationType NotificationType, callbackUri string, ) *DefaultNotificationSubscription`

NewDefaultNotificationSubscription instantiates a new DefaultNotificationSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultNotificationSubscriptionWithDefaults

`func NewDefaultNotificationSubscriptionWithDefaults() *DefaultNotificationSubscription`

NewDefaultNotificationSubscriptionWithDefaults instantiates a new DefaultNotificationSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *DefaultNotificationSubscription) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *DefaultNotificationSubscription) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *DefaultNotificationSubscription) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetCallbackUri

`func (o *DefaultNotificationSubscription) GetCallbackUri() string`

GetCallbackUri returns the CallbackUri field if non-nil, zero value otherwise.

### GetCallbackUriOk

`func (o *DefaultNotificationSubscription) GetCallbackUriOk() (*string, bool)`

GetCallbackUriOk returns a tuple with the CallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUri

`func (o *DefaultNotificationSubscription) SetCallbackUri(v string)`

SetCallbackUri sets CallbackUri field to given value.


### GetN1MessageClass

`func (o *DefaultNotificationSubscription) GetN1MessageClass() N1MessageClass`

GetN1MessageClass returns the N1MessageClass field if non-nil, zero value otherwise.

### GetN1MessageClassOk

`func (o *DefaultNotificationSubscription) GetN1MessageClassOk() (*N1MessageClass, bool)`

GetN1MessageClassOk returns a tuple with the N1MessageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1MessageClass

`func (o *DefaultNotificationSubscription) SetN1MessageClass(v N1MessageClass)`

SetN1MessageClass sets N1MessageClass field to given value.

### HasN1MessageClass

`func (o *DefaultNotificationSubscription) HasN1MessageClass() bool`

HasN1MessageClass returns a boolean if a field has been set.

### GetN2InformationClass

`func (o *DefaultNotificationSubscription) GetN2InformationClass() N2InformationClass`

GetN2InformationClass returns the N2InformationClass field if non-nil, zero value otherwise.

### GetN2InformationClassOk

`func (o *DefaultNotificationSubscription) GetN2InformationClassOk() (*N2InformationClass, bool)`

GetN2InformationClassOk returns a tuple with the N2InformationClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InformationClass

`func (o *DefaultNotificationSubscription) SetN2InformationClass(v N2InformationClass)`

SetN2InformationClass sets N2InformationClass field to given value.

### HasN2InformationClass

`func (o *DefaultNotificationSubscription) HasN2InformationClass() bool`

HasN2InformationClass returns a boolean if a field has been set.

### GetVersions

`func (o *DefaultNotificationSubscription) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *DefaultNotificationSubscription) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *DefaultNotificationSubscription) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *DefaultNotificationSubscription) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetBinding

`func (o *DefaultNotificationSubscription) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *DefaultNotificationSubscription) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *DefaultNotificationSubscription) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *DefaultNotificationSubscription) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetAcceptedEncoding

`func (o *DefaultNotificationSubscription) GetAcceptedEncoding() string`

GetAcceptedEncoding returns the AcceptedEncoding field if non-nil, zero value otherwise.

### GetAcceptedEncodingOk

`func (o *DefaultNotificationSubscription) GetAcceptedEncodingOk() (*string, bool)`

GetAcceptedEncodingOk returns a tuple with the AcceptedEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedEncoding

`func (o *DefaultNotificationSubscription) SetAcceptedEncoding(v string)`

SetAcceptedEncoding sets AcceptedEncoding field to given value.

### HasAcceptedEncoding

`func (o *DefaultNotificationSubscription) HasAcceptedEncoding() bool`

HasAcceptedEncoding returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *DefaultNotificationSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *DefaultNotificationSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *DefaultNotificationSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *DefaultNotificationSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetServiceInfoList

`func (o *DefaultNotificationSubscription) GetServiceInfoList() map[string]DefSubServiceInfo`

GetServiceInfoList returns the ServiceInfoList field if non-nil, zero value otherwise.

### GetServiceInfoListOk

`func (o *DefaultNotificationSubscription) GetServiceInfoListOk() (*map[string]DefSubServiceInfo, bool)`

GetServiceInfoListOk returns a tuple with the ServiceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInfoList

`func (o *DefaultNotificationSubscription) SetServiceInfoList(v map[string]DefSubServiceInfo)`

SetServiceInfoList sets ServiceInfoList field to given value.

### HasServiceInfoList

`func (o *DefaultNotificationSubscription) HasServiceInfoList() bool`

HasServiceInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


