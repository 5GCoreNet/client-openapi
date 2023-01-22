# CreatedEeSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EeSubscription** | [**EeSubscription**](EeSubscription.md) |  | 
**EventReports** | Pointer to [**[]MonitoringReport**](MonitoringReport.md) |  | [optional] 
**FailedMonitoringConfigs** | Pointer to [**map[string]FailedMonitoringConfiguration**](FailedMonitoringConfiguration.md) | A map (list of key-value pairs where referenceId converted from integer to string serves as key; see clause 6.4.6.3.2) of FailedMonitoringConfiguration | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewCreatedEeSubscription

`func NewCreatedEeSubscription(eeSubscription EeSubscription, ) *CreatedEeSubscription`

NewCreatedEeSubscription instantiates a new CreatedEeSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedEeSubscriptionWithDefaults

`func NewCreatedEeSubscriptionWithDefaults() *CreatedEeSubscription`

NewCreatedEeSubscriptionWithDefaults instantiates a new CreatedEeSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEeSubscription

`func (o *CreatedEeSubscription) GetEeSubscription() EeSubscription`

GetEeSubscription returns the EeSubscription field if non-nil, zero value otherwise.

### GetEeSubscriptionOk

`func (o *CreatedEeSubscription) GetEeSubscriptionOk() (*EeSubscription, bool)`

GetEeSubscriptionOk returns a tuple with the EeSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeSubscription

`func (o *CreatedEeSubscription) SetEeSubscription(v EeSubscription)`

SetEeSubscription sets EeSubscription field to given value.


### GetEventReports

`func (o *CreatedEeSubscription) GetEventReports() []MonitoringReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *CreatedEeSubscription) GetEventReportsOk() (*[]MonitoringReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *CreatedEeSubscription) SetEventReports(v []MonitoringReport)`

SetEventReports sets EventReports field to given value.

### HasEventReports

`func (o *CreatedEeSubscription) HasEventReports() bool`

HasEventReports returns a boolean if a field has been set.

### GetFailedMonitoringConfigs

`func (o *CreatedEeSubscription) GetFailedMonitoringConfigs() map[string]FailedMonitoringConfiguration`

GetFailedMonitoringConfigs returns the FailedMonitoringConfigs field if non-nil, zero value otherwise.

### GetFailedMonitoringConfigsOk

`func (o *CreatedEeSubscription) GetFailedMonitoringConfigsOk() (*map[string]FailedMonitoringConfiguration, bool)`

GetFailedMonitoringConfigsOk returns a tuple with the FailedMonitoringConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMonitoringConfigs

`func (o *CreatedEeSubscription) SetFailedMonitoringConfigs(v map[string]FailedMonitoringConfiguration)`

SetFailedMonitoringConfigs sets FailedMonitoringConfigs field to given value.

### HasFailedMonitoringConfigs

`func (o *CreatedEeSubscription) HasFailedMonitoringConfigs() bool`

HasFailedMonitoringConfigs returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *CreatedEeSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *CreatedEeSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *CreatedEeSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *CreatedEeSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


