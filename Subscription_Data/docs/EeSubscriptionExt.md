# EeSubscriptionExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**MonitoringConfigurations** | [**map[string]MonitoringConfiguration**](MonitoringConfiguration.md) | A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations | 
**ReportingOptions** | Pointer to [**ReportingOptions**](ReportingOptions.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**EpcAppliedInd** | Pointer to **bool** |  | [optional] [default to false]
**ScefDiamHost** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**ScefDiamRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 
**SecondCallbackRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExcludeGpsiList** | Pointer to **[]string** |  | [optional] 
**IncludeGpsiList** | Pointer to **[]string** |  | [optional] 
**AmfSubscriptionInfoList** | Pointer to [**[]AmfSubscriptionInfo**](AmfSubscriptionInfo.md) |  | [optional] 
**SmfSubscriptionInfo** | Pointer to [**SmfSubscriptionInfo**](SmfSubscriptionInfo.md) |  | [optional] 
**HssSubscriptionInfo** | Pointer to [**HssSubscriptionInfo**](HssSubscriptionInfo.md) |  | [optional] 

## Methods

### NewEeSubscriptionExt

`func NewEeSubscriptionExt(callbackReference string, monitoringConfigurations map[string]MonitoringConfiguration, ) *EeSubscriptionExt`

NewEeSubscriptionExt instantiates a new EeSubscriptionExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeSubscriptionExtWithDefaults

`func NewEeSubscriptionExtWithDefaults() *EeSubscriptionExt`

NewEeSubscriptionExtWithDefaults instantiates a new EeSubscriptionExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackReference

`func (o *EeSubscriptionExt) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *EeSubscriptionExt) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *EeSubscriptionExt) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetMonitoringConfigurations

`func (o *EeSubscriptionExt) GetMonitoringConfigurations() map[string]MonitoringConfiguration`

GetMonitoringConfigurations returns the MonitoringConfigurations field if non-nil, zero value otherwise.

### GetMonitoringConfigurationsOk

`func (o *EeSubscriptionExt) GetMonitoringConfigurationsOk() (*map[string]MonitoringConfiguration, bool)`

GetMonitoringConfigurationsOk returns a tuple with the MonitoringConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringConfigurations

`func (o *EeSubscriptionExt) SetMonitoringConfigurations(v map[string]MonitoringConfiguration)`

SetMonitoringConfigurations sets MonitoringConfigurations field to given value.


### GetReportingOptions

`func (o *EeSubscriptionExt) GetReportingOptions() ReportingOptions`

GetReportingOptions returns the ReportingOptions field if non-nil, zero value otherwise.

### GetReportingOptionsOk

`func (o *EeSubscriptionExt) GetReportingOptionsOk() (*ReportingOptions, bool)`

GetReportingOptionsOk returns a tuple with the ReportingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOptions

`func (o *EeSubscriptionExt) SetReportingOptions(v ReportingOptions)`

SetReportingOptions sets ReportingOptions field to given value.

### HasReportingOptions

`func (o *EeSubscriptionExt) HasReportingOptions() bool`

HasReportingOptions returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeSubscriptionExt) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeSubscriptionExt) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeSubscriptionExt) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeSubscriptionExt) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *EeSubscriptionExt) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *EeSubscriptionExt) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *EeSubscriptionExt) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *EeSubscriptionExt) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetContextInfo

`func (o *EeSubscriptionExt) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *EeSubscriptionExt) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *EeSubscriptionExt) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *EeSubscriptionExt) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetEpcAppliedInd

`func (o *EeSubscriptionExt) GetEpcAppliedInd() bool`

GetEpcAppliedInd returns the EpcAppliedInd field if non-nil, zero value otherwise.

### GetEpcAppliedIndOk

`func (o *EeSubscriptionExt) GetEpcAppliedIndOk() (*bool, bool)`

GetEpcAppliedIndOk returns a tuple with the EpcAppliedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpcAppliedInd

`func (o *EeSubscriptionExt) SetEpcAppliedInd(v bool)`

SetEpcAppliedInd sets EpcAppliedInd field to given value.

### HasEpcAppliedInd

`func (o *EeSubscriptionExt) HasEpcAppliedInd() bool`

HasEpcAppliedInd returns a boolean if a field has been set.

### GetScefDiamHost

`func (o *EeSubscriptionExt) GetScefDiamHost() string`

GetScefDiamHost returns the ScefDiamHost field if non-nil, zero value otherwise.

### GetScefDiamHostOk

`func (o *EeSubscriptionExt) GetScefDiamHostOk() (*string, bool)`

GetScefDiamHostOk returns a tuple with the ScefDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefDiamHost

`func (o *EeSubscriptionExt) SetScefDiamHost(v string)`

SetScefDiamHost sets ScefDiamHost field to given value.

### HasScefDiamHost

`func (o *EeSubscriptionExt) HasScefDiamHost() bool`

HasScefDiamHost returns a boolean if a field has been set.

### GetScefDiamRealm

`func (o *EeSubscriptionExt) GetScefDiamRealm() string`

GetScefDiamRealm returns the ScefDiamRealm field if non-nil, zero value otherwise.

### GetScefDiamRealmOk

`func (o *EeSubscriptionExt) GetScefDiamRealmOk() (*string, bool)`

GetScefDiamRealmOk returns a tuple with the ScefDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefDiamRealm

`func (o *EeSubscriptionExt) SetScefDiamRealm(v string)`

SetScefDiamRealm sets ScefDiamRealm field to given value.

### HasScefDiamRealm

`func (o *EeSubscriptionExt) HasScefDiamRealm() bool`

HasScefDiamRealm returns a boolean if a field has been set.

### GetNotifyCorrelationId

`func (o *EeSubscriptionExt) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *EeSubscriptionExt) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *EeSubscriptionExt) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *EeSubscriptionExt) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetSecondCallbackRef

`func (o *EeSubscriptionExt) GetSecondCallbackRef() string`

GetSecondCallbackRef returns the SecondCallbackRef field if non-nil, zero value otherwise.

### GetSecondCallbackRefOk

`func (o *EeSubscriptionExt) GetSecondCallbackRefOk() (*string, bool)`

GetSecondCallbackRefOk returns a tuple with the SecondCallbackRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondCallbackRef

`func (o *EeSubscriptionExt) SetSecondCallbackRef(v string)`

SetSecondCallbackRef sets SecondCallbackRef field to given value.

### HasSecondCallbackRef

`func (o *EeSubscriptionExt) HasSecondCallbackRef() bool`

HasSecondCallbackRef returns a boolean if a field has been set.

### GetGpsi

`func (o *EeSubscriptionExt) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *EeSubscriptionExt) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *EeSubscriptionExt) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *EeSubscriptionExt) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetExcludeGpsiList

`func (o *EeSubscriptionExt) GetExcludeGpsiList() []string`

GetExcludeGpsiList returns the ExcludeGpsiList field if non-nil, zero value otherwise.

### GetExcludeGpsiListOk

`func (o *EeSubscriptionExt) GetExcludeGpsiListOk() (*[]string, bool)`

GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGpsiList

`func (o *EeSubscriptionExt) SetExcludeGpsiList(v []string)`

SetExcludeGpsiList sets ExcludeGpsiList field to given value.

### HasExcludeGpsiList

`func (o *EeSubscriptionExt) HasExcludeGpsiList() bool`

HasExcludeGpsiList returns a boolean if a field has been set.

### GetIncludeGpsiList

`func (o *EeSubscriptionExt) GetIncludeGpsiList() []string`

GetIncludeGpsiList returns the IncludeGpsiList field if non-nil, zero value otherwise.

### GetIncludeGpsiListOk

`func (o *EeSubscriptionExt) GetIncludeGpsiListOk() (*[]string, bool)`

GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGpsiList

`func (o *EeSubscriptionExt) SetIncludeGpsiList(v []string)`

SetIncludeGpsiList sets IncludeGpsiList field to given value.

### HasIncludeGpsiList

`func (o *EeSubscriptionExt) HasIncludeGpsiList() bool`

HasIncludeGpsiList returns a boolean if a field has been set.

### GetAmfSubscriptionInfoList

`func (o *EeSubscriptionExt) GetAmfSubscriptionInfoList() []AmfSubscriptionInfo`

GetAmfSubscriptionInfoList returns the AmfSubscriptionInfoList field if non-nil, zero value otherwise.

### GetAmfSubscriptionInfoListOk

`func (o *EeSubscriptionExt) GetAmfSubscriptionInfoListOk() (*[]AmfSubscriptionInfo, bool)`

GetAmfSubscriptionInfoListOk returns a tuple with the AmfSubscriptionInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSubscriptionInfoList

`func (o *EeSubscriptionExt) SetAmfSubscriptionInfoList(v []AmfSubscriptionInfo)`

SetAmfSubscriptionInfoList sets AmfSubscriptionInfoList field to given value.

### HasAmfSubscriptionInfoList

`func (o *EeSubscriptionExt) HasAmfSubscriptionInfoList() bool`

HasAmfSubscriptionInfoList returns a boolean if a field has been set.

### GetSmfSubscriptionInfo

`func (o *EeSubscriptionExt) GetSmfSubscriptionInfo() SmfSubscriptionInfo`

GetSmfSubscriptionInfo returns the SmfSubscriptionInfo field if non-nil, zero value otherwise.

### GetSmfSubscriptionInfoOk

`func (o *EeSubscriptionExt) GetSmfSubscriptionInfoOk() (*SmfSubscriptionInfo, bool)`

GetSmfSubscriptionInfoOk returns a tuple with the SmfSubscriptionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSubscriptionInfo

`func (o *EeSubscriptionExt) SetSmfSubscriptionInfo(v SmfSubscriptionInfo)`

SetSmfSubscriptionInfo sets SmfSubscriptionInfo field to given value.

### HasSmfSubscriptionInfo

`func (o *EeSubscriptionExt) HasSmfSubscriptionInfo() bool`

HasSmfSubscriptionInfo returns a boolean if a field has been set.

### GetHssSubscriptionInfo

`func (o *EeSubscriptionExt) GetHssSubscriptionInfo() HssSubscriptionInfo`

GetHssSubscriptionInfo returns the HssSubscriptionInfo field if non-nil, zero value otherwise.

### GetHssSubscriptionInfoOk

`func (o *EeSubscriptionExt) GetHssSubscriptionInfoOk() (*HssSubscriptionInfo, bool)`

GetHssSubscriptionInfoOk returns a tuple with the HssSubscriptionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssSubscriptionInfo

`func (o *EeSubscriptionExt) SetHssSubscriptionInfo(v HssSubscriptionInfo)`

SetHssSubscriptionInfo sets HssSubscriptionInfo field to given value.

### HasHssSubscriptionInfo

`func (o *EeSubscriptionExt) HasHssSubscriptionInfo() bool`

HasHssSubscriptionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


