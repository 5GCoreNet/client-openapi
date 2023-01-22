# AnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyEvent** | [**AnalyticsEvent**](AnalyticsEvent.md) |  | 
**AnalyEventFilter** | Pointer to [**AnalyticsEventFilter**](AnalyticsEventFilter.md) |  | [optional] 
**AnalyRep** | Pointer to [**EventReportingRequirement**](EventReportingRequirement.md) |  | [optional] 
**TgtUe** | Pointer to [**TargetUeId**](TargetUeId.md) |  | [optional] 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 

## Methods

### NewAnalyticsRequest

`func NewAnalyticsRequest(analyEvent AnalyticsEvent, suppFeat string, ) *AnalyticsRequest`

NewAnalyticsRequest instantiates a new AnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRequestWithDefaults

`func NewAnalyticsRequestWithDefaults() *AnalyticsRequest`

NewAnalyticsRequestWithDefaults instantiates a new AnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyEvent

`func (o *AnalyticsRequest) GetAnalyEvent() AnalyticsEvent`

GetAnalyEvent returns the AnalyEvent field if non-nil, zero value otherwise.

### GetAnalyEventOk

`func (o *AnalyticsRequest) GetAnalyEventOk() (*AnalyticsEvent, bool)`

GetAnalyEventOk returns a tuple with the AnalyEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyEvent

`func (o *AnalyticsRequest) SetAnalyEvent(v AnalyticsEvent)`

SetAnalyEvent sets AnalyEvent field to given value.


### GetAnalyEventFilter

`func (o *AnalyticsRequest) GetAnalyEventFilter() AnalyticsEventFilter`

GetAnalyEventFilter returns the AnalyEventFilter field if non-nil, zero value otherwise.

### GetAnalyEventFilterOk

`func (o *AnalyticsRequest) GetAnalyEventFilterOk() (*AnalyticsEventFilter, bool)`

GetAnalyEventFilterOk returns a tuple with the AnalyEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyEventFilter

`func (o *AnalyticsRequest) SetAnalyEventFilter(v AnalyticsEventFilter)`

SetAnalyEventFilter sets AnalyEventFilter field to given value.

### HasAnalyEventFilter

`func (o *AnalyticsRequest) HasAnalyEventFilter() bool`

HasAnalyEventFilter returns a boolean if a field has been set.

### GetAnalyRep

`func (o *AnalyticsRequest) GetAnalyRep() EventReportingRequirement`

GetAnalyRep returns the AnalyRep field if non-nil, zero value otherwise.

### GetAnalyRepOk

`func (o *AnalyticsRequest) GetAnalyRepOk() (*EventReportingRequirement, bool)`

GetAnalyRepOk returns a tuple with the AnalyRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyRep

`func (o *AnalyticsRequest) SetAnalyRep(v EventReportingRequirement)`

SetAnalyRep sets AnalyRep field to given value.

### HasAnalyRep

`func (o *AnalyticsRequest) HasAnalyRep() bool`

HasAnalyRep returns a boolean if a field has been set.

### GetTgtUe

`func (o *AnalyticsRequest) GetTgtUe() TargetUeId`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *AnalyticsRequest) GetTgtUeOk() (*TargetUeId, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *AnalyticsRequest) SetTgtUe(v TargetUeId)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *AnalyticsRequest) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AnalyticsRequest) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AnalyticsRequest) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AnalyticsRequest) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


