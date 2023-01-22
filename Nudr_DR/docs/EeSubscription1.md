# EeSubscription1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**MonitoringConfigurations** | [**map[string]MonitoringConfiguration1**](MonitoringConfiguration1.md) | A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations | 
**ReportingOptions** | Pointer to [**ReportingOptions1**](ReportingOptions1.md) |  | [optional] 
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

## Methods

### NewEeSubscription1

`func NewEeSubscription1(callbackReference string, monitoringConfigurations map[string]MonitoringConfiguration1, ) *EeSubscription1`

NewEeSubscription1 instantiates a new EeSubscription1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEeSubscription1WithDefaults

`func NewEeSubscription1WithDefaults() *EeSubscription1`

NewEeSubscription1WithDefaults instantiates a new EeSubscription1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackReference

`func (o *EeSubscription1) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *EeSubscription1) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *EeSubscription1) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetMonitoringConfigurations

`func (o *EeSubscription1) GetMonitoringConfigurations() map[string]MonitoringConfiguration1`

GetMonitoringConfigurations returns the MonitoringConfigurations field if non-nil, zero value otherwise.

### GetMonitoringConfigurationsOk

`func (o *EeSubscription1) GetMonitoringConfigurationsOk() (*map[string]MonitoringConfiguration1, bool)`

GetMonitoringConfigurationsOk returns a tuple with the MonitoringConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringConfigurations

`func (o *EeSubscription1) SetMonitoringConfigurations(v map[string]MonitoringConfiguration1)`

SetMonitoringConfigurations sets MonitoringConfigurations field to given value.


### GetReportingOptions

`func (o *EeSubscription1) GetReportingOptions() ReportingOptions1`

GetReportingOptions returns the ReportingOptions field if non-nil, zero value otherwise.

### GetReportingOptionsOk

`func (o *EeSubscription1) GetReportingOptionsOk() (*ReportingOptions1, bool)`

GetReportingOptionsOk returns a tuple with the ReportingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOptions

`func (o *EeSubscription1) SetReportingOptions(v ReportingOptions1)`

SetReportingOptions sets ReportingOptions field to given value.

### HasReportingOptions

`func (o *EeSubscription1) HasReportingOptions() bool`

HasReportingOptions returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EeSubscription1) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EeSubscription1) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EeSubscription1) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EeSubscription1) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *EeSubscription1) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *EeSubscription1) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *EeSubscription1) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *EeSubscription1) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetContextInfo

`func (o *EeSubscription1) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *EeSubscription1) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *EeSubscription1) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *EeSubscription1) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetEpcAppliedInd

`func (o *EeSubscription1) GetEpcAppliedInd() bool`

GetEpcAppliedInd returns the EpcAppliedInd field if non-nil, zero value otherwise.

### GetEpcAppliedIndOk

`func (o *EeSubscription1) GetEpcAppliedIndOk() (*bool, bool)`

GetEpcAppliedIndOk returns a tuple with the EpcAppliedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpcAppliedInd

`func (o *EeSubscription1) SetEpcAppliedInd(v bool)`

SetEpcAppliedInd sets EpcAppliedInd field to given value.

### HasEpcAppliedInd

`func (o *EeSubscription1) HasEpcAppliedInd() bool`

HasEpcAppliedInd returns a boolean if a field has been set.

### GetScefDiamHost

`func (o *EeSubscription1) GetScefDiamHost() string`

GetScefDiamHost returns the ScefDiamHost field if non-nil, zero value otherwise.

### GetScefDiamHostOk

`func (o *EeSubscription1) GetScefDiamHostOk() (*string, bool)`

GetScefDiamHostOk returns a tuple with the ScefDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefDiamHost

`func (o *EeSubscription1) SetScefDiamHost(v string)`

SetScefDiamHost sets ScefDiamHost field to given value.

### HasScefDiamHost

`func (o *EeSubscription1) HasScefDiamHost() bool`

HasScefDiamHost returns a boolean if a field has been set.

### GetScefDiamRealm

`func (o *EeSubscription1) GetScefDiamRealm() string`

GetScefDiamRealm returns the ScefDiamRealm field if non-nil, zero value otherwise.

### GetScefDiamRealmOk

`func (o *EeSubscription1) GetScefDiamRealmOk() (*string, bool)`

GetScefDiamRealmOk returns a tuple with the ScefDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefDiamRealm

`func (o *EeSubscription1) SetScefDiamRealm(v string)`

SetScefDiamRealm sets ScefDiamRealm field to given value.

### HasScefDiamRealm

`func (o *EeSubscription1) HasScefDiamRealm() bool`

HasScefDiamRealm returns a boolean if a field has been set.

### GetNotifyCorrelationId

`func (o *EeSubscription1) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *EeSubscription1) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *EeSubscription1) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *EeSubscription1) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetSecondCallbackRef

`func (o *EeSubscription1) GetSecondCallbackRef() string`

GetSecondCallbackRef returns the SecondCallbackRef field if non-nil, zero value otherwise.

### GetSecondCallbackRefOk

`func (o *EeSubscription1) GetSecondCallbackRefOk() (*string, bool)`

GetSecondCallbackRefOk returns a tuple with the SecondCallbackRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondCallbackRef

`func (o *EeSubscription1) SetSecondCallbackRef(v string)`

SetSecondCallbackRef sets SecondCallbackRef field to given value.

### HasSecondCallbackRef

`func (o *EeSubscription1) HasSecondCallbackRef() bool`

HasSecondCallbackRef returns a boolean if a field has been set.

### GetGpsi

`func (o *EeSubscription1) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *EeSubscription1) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *EeSubscription1) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *EeSubscription1) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetExcludeGpsiList

`func (o *EeSubscription1) GetExcludeGpsiList() []string`

GetExcludeGpsiList returns the ExcludeGpsiList field if non-nil, zero value otherwise.

### GetExcludeGpsiListOk

`func (o *EeSubscription1) GetExcludeGpsiListOk() (*[]string, bool)`

GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGpsiList

`func (o *EeSubscription1) SetExcludeGpsiList(v []string)`

SetExcludeGpsiList sets ExcludeGpsiList field to given value.

### HasExcludeGpsiList

`func (o *EeSubscription1) HasExcludeGpsiList() bool`

HasExcludeGpsiList returns a boolean if a field has been set.

### GetIncludeGpsiList

`func (o *EeSubscription1) GetIncludeGpsiList() []string`

GetIncludeGpsiList returns the IncludeGpsiList field if non-nil, zero value otherwise.

### GetIncludeGpsiListOk

`func (o *EeSubscription1) GetIncludeGpsiListOk() (*[]string, bool)`

GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGpsiList

`func (o *EeSubscription1) SetIncludeGpsiList(v []string)`

SetIncludeGpsiList sets IncludeGpsiList field to given value.

### HasIncludeGpsiList

`func (o *EeSubscription1) HasIncludeGpsiList() bool`

HasIncludeGpsiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


