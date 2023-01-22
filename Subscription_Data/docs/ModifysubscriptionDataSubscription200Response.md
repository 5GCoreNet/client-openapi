# ModifysubscriptionDataSubscription200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Report** | [**[]ReportItem**](ReportItem.md) | The execution report contains an array of report items. Each report item indicates one  failed modification.  | 
**UeId** | Pointer to **string** | String represents the SUPI or GPSI | [optional] 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**OriginalCallbackReference** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**MonitoredResourceUris** | **[]string** |  | 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SdmSubscription** | Pointer to [**SdmSubscription**](SdmSubscription.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**UniqueSubscription** | Pointer to **bool** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewModifysubscriptionDataSubscription200Response

`func NewModifysubscriptionDataSubscription200Response(report []ReportItem, callbackReference string, monitoredResourceUris []string, ) *ModifysubscriptionDataSubscription200Response`

NewModifysubscriptionDataSubscription200Response instantiates a new ModifysubscriptionDataSubscription200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifysubscriptionDataSubscription200ResponseWithDefaults

`func NewModifysubscriptionDataSubscription200ResponseWithDefaults() *ModifysubscriptionDataSubscription200Response`

NewModifysubscriptionDataSubscription200ResponseWithDefaults instantiates a new ModifysubscriptionDataSubscription200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReport

`func (o *ModifysubscriptionDataSubscription200Response) GetReport() []ReportItem`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ModifysubscriptionDataSubscription200Response) GetReportOk() (*[]ReportItem, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ModifysubscriptionDataSubscription200Response) SetReport(v []ReportItem)`

SetReport sets Report field to given value.


### GetUeId

`func (o *ModifysubscriptionDataSubscription200Response) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ModifysubscriptionDataSubscription200Response) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ModifysubscriptionDataSubscription200Response) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *ModifysubscriptionDataSubscription200Response) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetCallbackReference

`func (o *ModifysubscriptionDataSubscription200Response) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *ModifysubscriptionDataSubscription200Response) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *ModifysubscriptionDataSubscription200Response) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetOriginalCallbackReference

`func (o *ModifysubscriptionDataSubscription200Response) GetOriginalCallbackReference() string`

GetOriginalCallbackReference returns the OriginalCallbackReference field if non-nil, zero value otherwise.

### GetOriginalCallbackReferenceOk

`func (o *ModifysubscriptionDataSubscription200Response) GetOriginalCallbackReferenceOk() (*string, bool)`

GetOriginalCallbackReferenceOk returns a tuple with the OriginalCallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCallbackReference

`func (o *ModifysubscriptionDataSubscription200Response) SetOriginalCallbackReference(v string)`

SetOriginalCallbackReference sets OriginalCallbackReference field to given value.

### HasOriginalCallbackReference

`func (o *ModifysubscriptionDataSubscription200Response) HasOriginalCallbackReference() bool`

HasOriginalCallbackReference returns a boolean if a field has been set.

### GetMonitoredResourceUris

`func (o *ModifysubscriptionDataSubscription200Response) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *ModifysubscriptionDataSubscription200Response) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *ModifysubscriptionDataSubscription200Response) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetExpiry

`func (o *ModifysubscriptionDataSubscription200Response) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ModifysubscriptionDataSubscription200Response) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ModifysubscriptionDataSubscription200Response) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ModifysubscriptionDataSubscription200Response) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSdmSubscription

`func (o *ModifysubscriptionDataSubscription200Response) GetSdmSubscription() SdmSubscription`

GetSdmSubscription returns the SdmSubscription field if non-nil, zero value otherwise.

### GetSdmSubscriptionOk

`func (o *ModifysubscriptionDataSubscription200Response) GetSdmSubscriptionOk() (*SdmSubscription, bool)`

GetSdmSubscriptionOk returns a tuple with the SdmSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdmSubscription

`func (o *ModifysubscriptionDataSubscription200Response) SetSdmSubscription(v SdmSubscription)`

SetSdmSubscription sets SdmSubscription field to given value.

### HasSdmSubscription

`func (o *ModifysubscriptionDataSubscription200Response) HasSdmSubscription() bool`

HasSdmSubscription returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *ModifysubscriptionDataSubscription200Response) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ModifysubscriptionDataSubscription200Response) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ModifysubscriptionDataSubscription200Response) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ModifysubscriptionDataSubscription200Response) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUniqueSubscription

`func (o *ModifysubscriptionDataSubscription200Response) GetUniqueSubscription() bool`

GetUniqueSubscription returns the UniqueSubscription field if non-nil, zero value otherwise.

### GetUniqueSubscriptionOk

`func (o *ModifysubscriptionDataSubscription200Response) GetUniqueSubscriptionOk() (*bool, bool)`

GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSubscription

`func (o *ModifysubscriptionDataSubscription200Response) SetUniqueSubscription(v bool)`

SetUniqueSubscription sets UniqueSubscription field to given value.

### HasUniqueSubscription

`func (o *ModifysubscriptionDataSubscription200Response) HasUniqueSubscription() bool`

HasUniqueSubscription returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ModifysubscriptionDataSubscription200Response) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ModifysubscriptionDataSubscription200Response) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ModifysubscriptionDataSubscription200Response) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ModifysubscriptionDataSubscription200Response) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


