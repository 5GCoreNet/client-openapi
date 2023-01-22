# EeSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**ScefId** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**ScefDiamRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**MonitoringConfigurations** | Pointer to [**map[string]MonitoringConfiguration**](MonitoringConfiguration.md) | A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ReportingOptions** | Pointer to [**ReportingOptions**](ReportingOptions.md) |  | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**ExternalIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewEeSubscription

`func NewEeSubscription(callbackReference string, ) *EeSubscription`

NewEeSubscription instantiates a new EeSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeSubscriptionWithDefaults

`func NewEeSubscriptionWithDefaults() *EeSubscription`

NewEeSubscriptionWithDefaults instantiates a new EeSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackReference

`func (o *EeSubscription) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *EeSubscription) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *EeSubscription) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetScefId

`func (o *EeSubscription) GetScefId() string`

GetScefId returns the ScefId field if non-nil, zero value otherwise.

### GetScefIdOk

`func (o *EeSubscription) GetScefIdOk() (*string, bool)`

GetScefIdOk returns a tuple with the ScefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefId

`func (o *EeSubscription) SetScefId(v string)`

SetScefId sets ScefId field to given value.

### HasScefId

`func (o *EeSubscription) HasScefId() bool`

HasScefId returns a boolean if a field has been set.

### GetScefDiamRealm

`func (o *EeSubscription) GetScefDiamRealm() string`

GetScefDiamRealm returns the ScefDiamRealm field if non-nil, zero value otherwise.

### GetScefDiamRealmOk

`func (o *EeSubscription) GetScefDiamRealmOk() (*string, bool)`

GetScefDiamRealmOk returns a tuple with the ScefDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefDiamRealm

`func (o *EeSubscription) SetScefDiamRealm(v string)`

SetScefDiamRealm sets ScefDiamRealm field to given value.

### HasScefDiamRealm

`func (o *EeSubscription) HasScefDiamRealm() bool`

HasScefDiamRealm returns a boolean if a field has been set.

### GetMonitoringConfigurations

`func (o *EeSubscription) GetMonitoringConfigurations() map[string]MonitoringConfiguration`

GetMonitoringConfigurations returns the MonitoringConfigurations field if non-nil, zero value otherwise.

### GetMonitoringConfigurationsOk

`func (o *EeSubscription) GetMonitoringConfigurationsOk() (*map[string]MonitoringConfiguration, bool)`

GetMonitoringConfigurationsOk returns a tuple with the MonitoringConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringConfigurations

`func (o *EeSubscription) SetMonitoringConfigurations(v map[string]MonitoringConfiguration)`

SetMonitoringConfigurations sets MonitoringConfigurations field to given value.

### HasMonitoringConfigurations

`func (o *EeSubscription) HasMonitoringConfigurations() bool`

HasMonitoringConfigurations returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetReportingOptions

`func (o *EeSubscription) GetReportingOptions() ReportingOptions`

GetReportingOptions returns the ReportingOptions field if non-nil, zero value otherwise.

### GetReportingOptionsOk

`func (o *EeSubscription) GetReportingOptionsOk() (*ReportingOptions, bool)`

GetReportingOptionsOk returns a tuple with the ReportingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOptions

`func (o *EeSubscription) SetReportingOptions(v ReportingOptions)`

SetReportingOptions sets ReportingOptions field to given value.

### HasReportingOptions

`func (o *EeSubscription) HasReportingOptions() bool`

HasReportingOptions returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *EeSubscription) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *EeSubscription) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *EeSubscription) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *EeSubscription) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *EeSubscription) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *EeSubscription) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *EeSubscription) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *EeSubscription) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


