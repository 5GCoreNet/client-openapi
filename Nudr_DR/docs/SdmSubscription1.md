# SdmSubscription1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**ImplicitUnsubscribe** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**AmfServiceName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**MonitoredResourceUris** | **[]string** |  | 
**SingleNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**PlmnId** | Pointer to [**PlmnId1**](PlmnId1.md) |  | [optional] 
**ImmediateReport** | Pointer to **bool** |  | [optional] [default to false]
**Report** | Pointer to [**ImmediateReport**](ImmediateReport.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ContextInfo** | Pointer to [**ContextInfo**](ContextInfo.md) |  | [optional] 
**NfChangeFilter** | Pointer to **bool** |  | [optional] [default to false]
**UniqueSubscription** | Pointer to **bool** |  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**UeConSmfDataSubFilter** | Pointer to [**UeContextInSmfDataSubFilter1**](UeContextInSmfDataSubFilter1.md) |  | [optional] 

## Methods

### NewSdmSubscription1

`func NewSdmSubscription1(nfInstanceId string, callbackReference string, monitoredResourceUris []string, ) *SdmSubscription1`

NewSdmSubscription1 instantiates a new SdmSubscription1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdmSubscription1WithDefaults

`func NewSdmSubscription1WithDefaults() *SdmSubscription1`

NewSdmSubscription1WithDefaults instantiates a new SdmSubscription1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *SdmSubscription1) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *SdmSubscription1) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *SdmSubscription1) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetImplicitUnsubscribe

`func (o *SdmSubscription1) GetImplicitUnsubscribe() bool`

GetImplicitUnsubscribe returns the ImplicitUnsubscribe field if non-nil, zero value otherwise.

### GetImplicitUnsubscribeOk

`func (o *SdmSubscription1) GetImplicitUnsubscribeOk() (*bool, bool)`

GetImplicitUnsubscribeOk returns a tuple with the ImplicitUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitUnsubscribe

`func (o *SdmSubscription1) SetImplicitUnsubscribe(v bool)`

SetImplicitUnsubscribe sets ImplicitUnsubscribe field to given value.

### HasImplicitUnsubscribe

`func (o *SdmSubscription1) HasImplicitUnsubscribe() bool`

HasImplicitUnsubscribe returns a boolean if a field has been set.

### GetExpires

`func (o *SdmSubscription1) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SdmSubscription1) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SdmSubscription1) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SdmSubscription1) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetCallbackReference

`func (o *SdmSubscription1) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *SdmSubscription1) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *SdmSubscription1) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetAmfServiceName

`func (o *SdmSubscription1) GetAmfServiceName() ServiceName`

GetAmfServiceName returns the AmfServiceName field if non-nil, zero value otherwise.

### GetAmfServiceNameOk

`func (o *SdmSubscription1) GetAmfServiceNameOk() (*ServiceName, bool)`

GetAmfServiceNameOk returns a tuple with the AmfServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfServiceName

`func (o *SdmSubscription1) SetAmfServiceName(v ServiceName)`

SetAmfServiceName sets AmfServiceName field to given value.

### HasAmfServiceName

`func (o *SdmSubscription1) HasAmfServiceName() bool`

HasAmfServiceName returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *SdmSubscription1) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *SdmSubscription1) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *SdmSubscription1) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetSingleNssai

`func (o *SdmSubscription1) GetSingleNssai() Snssai`

GetSingleNssai returns the SingleNssai field if non-nil, zero value otherwise.

### GetSingleNssaiOk

`func (o *SdmSubscription1) GetSingleNssaiOk() (*Snssai, bool)`

GetSingleNssaiOk returns a tuple with the SingleNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNssai

`func (o *SdmSubscription1) SetSingleNssai(v Snssai)`

SetSingleNssai sets SingleNssai field to given value.

### HasSingleNssai

`func (o *SdmSubscription1) HasSingleNssai() bool`

HasSingleNssai returns a boolean if a field has been set.

### GetDnn

`func (o *SdmSubscription1) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SdmSubscription1) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SdmSubscription1) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SdmSubscription1) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SdmSubscription1) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SdmSubscription1) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SdmSubscription1) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SdmSubscription1) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPlmnId

`func (o *SdmSubscription1) GetPlmnId() PlmnId1`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SdmSubscription1) GetPlmnIdOk() (*PlmnId1, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SdmSubscription1) SetPlmnId(v PlmnId1)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *SdmSubscription1) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetImmediateReport

`func (o *SdmSubscription1) GetImmediateReport() bool`

GetImmediateReport returns the ImmediateReport field if non-nil, zero value otherwise.

### GetImmediateReportOk

`func (o *SdmSubscription1) GetImmediateReportOk() (*bool, bool)`

GetImmediateReportOk returns a tuple with the ImmediateReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateReport

`func (o *SdmSubscription1) SetImmediateReport(v bool)`

SetImmediateReport sets ImmediateReport field to given value.

### HasImmediateReport

`func (o *SdmSubscription1) HasImmediateReport() bool`

HasImmediateReport returns a boolean if a field has been set.

### GetReport

`func (o *SdmSubscription1) GetReport() ImmediateReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *SdmSubscription1) GetReportOk() (*ImmediateReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *SdmSubscription1) SetReport(v ImmediateReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *SdmSubscription1) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SdmSubscription1) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SdmSubscription1) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SdmSubscription1) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SdmSubscription1) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetContextInfo

`func (o *SdmSubscription1) GetContextInfo() ContextInfo`

GetContextInfo returns the ContextInfo field if non-nil, zero value otherwise.

### GetContextInfoOk

`func (o *SdmSubscription1) GetContextInfoOk() (*ContextInfo, bool)`

GetContextInfoOk returns a tuple with the ContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextInfo

`func (o *SdmSubscription1) SetContextInfo(v ContextInfo)`

SetContextInfo sets ContextInfo field to given value.

### HasContextInfo

`func (o *SdmSubscription1) HasContextInfo() bool`

HasContextInfo returns a boolean if a field has been set.

### GetNfChangeFilter

`func (o *SdmSubscription1) GetNfChangeFilter() bool`

GetNfChangeFilter returns the NfChangeFilter field if non-nil, zero value otherwise.

### GetNfChangeFilterOk

`func (o *SdmSubscription1) GetNfChangeFilterOk() (*bool, bool)`

GetNfChangeFilterOk returns a tuple with the NfChangeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfChangeFilter

`func (o *SdmSubscription1) SetNfChangeFilter(v bool)`

SetNfChangeFilter sets NfChangeFilter field to given value.

### HasNfChangeFilter

`func (o *SdmSubscription1) HasNfChangeFilter() bool`

HasNfChangeFilter returns a boolean if a field has been set.

### GetUniqueSubscription

`func (o *SdmSubscription1) GetUniqueSubscription() bool`

GetUniqueSubscription returns the UniqueSubscription field if non-nil, zero value otherwise.

### GetUniqueSubscriptionOk

`func (o *SdmSubscription1) GetUniqueSubscriptionOk() (*bool, bool)`

GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSubscription

`func (o *SdmSubscription1) SetUniqueSubscription(v bool)`

SetUniqueSubscription sets UniqueSubscription field to given value.

### HasUniqueSubscription

`func (o *SdmSubscription1) HasUniqueSubscription() bool`

HasUniqueSubscription returns a boolean if a field has been set.

### GetResetIds

`func (o *SdmSubscription1) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *SdmSubscription1) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *SdmSubscription1) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *SdmSubscription1) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetUeConSmfDataSubFilter

`func (o *SdmSubscription1) GetUeConSmfDataSubFilter() UeContextInSmfDataSubFilter1`

GetUeConSmfDataSubFilter returns the UeConSmfDataSubFilter field if non-nil, zero value otherwise.

### GetUeConSmfDataSubFilterOk

`func (o *SdmSubscription1) GetUeConSmfDataSubFilterOk() (*UeContextInSmfDataSubFilter1, bool)`

GetUeConSmfDataSubFilterOk returns a tuple with the UeConSmfDataSubFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeConSmfDataSubFilter

`func (o *SdmSubscription1) SetUeConSmfDataSubFilter(v UeContextInSmfDataSubFilter1)`

SetUeConSmfDataSubFilter sets UeConSmfDataSubFilter field to given value.

### HasUeConSmfDataSubFilter

`func (o *SdmSubscription1) HasUeConSmfDataSubFilter() bool`

HasUeConSmfDataSubFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


