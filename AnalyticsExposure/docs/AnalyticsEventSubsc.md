# AnalyticsEventSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyEvent** | [**AnalyticsEvent**](AnalyticsEvent.md) |  | 
**AnalyEventFilter** | Pointer to [**AnalyticsEventFilterSubsc**](AnalyticsEventFilterSubsc.md) |  | [optional] 
**TgtUe** | Pointer to [**TargetUeId**](TargetUeId.md) |  | [optional] 

## Methods

### NewAnalyticsEventSubsc

`func NewAnalyticsEventSubsc(analyEvent AnalyticsEvent, ) *AnalyticsEventSubsc`

NewAnalyticsEventSubsc instantiates a new AnalyticsEventSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsEventSubscWithDefaults

`func NewAnalyticsEventSubscWithDefaults() *AnalyticsEventSubsc`

NewAnalyticsEventSubscWithDefaults instantiates a new AnalyticsEventSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyEvent

`func (o *AnalyticsEventSubsc) GetAnalyEvent() AnalyticsEvent`

GetAnalyEvent returns the AnalyEvent field if non-nil, zero value otherwise.

### GetAnalyEventOk

`func (o *AnalyticsEventSubsc) GetAnalyEventOk() (*AnalyticsEvent, bool)`

GetAnalyEventOk returns a tuple with the AnalyEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyEvent

`func (o *AnalyticsEventSubsc) SetAnalyEvent(v AnalyticsEvent)`

SetAnalyEvent sets AnalyEvent field to given value.


### GetAnalyEventFilter

`func (o *AnalyticsEventSubsc) GetAnalyEventFilter() AnalyticsEventFilterSubsc`

GetAnalyEventFilter returns the AnalyEventFilter field if non-nil, zero value otherwise.

### GetAnalyEventFilterOk

`func (o *AnalyticsEventSubsc) GetAnalyEventFilterOk() (*AnalyticsEventFilterSubsc, bool)`

GetAnalyEventFilterOk returns a tuple with the AnalyEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyEventFilter

`func (o *AnalyticsEventSubsc) SetAnalyEventFilter(v AnalyticsEventFilterSubsc)`

SetAnalyEventFilter sets AnalyEventFilter field to given value.

### HasAnalyEventFilter

`func (o *AnalyticsEventSubsc) HasAnalyEventFilter() bool`

HasAnalyEventFilter returns a boolean if a field has been set.

### GetTgtUe

`func (o *AnalyticsEventSubsc) GetTgtUe() TargetUeId`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *AnalyticsEventSubsc) GetTgtUeOk() (*TargetUeId, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *AnalyticsEventSubsc) SetTgtUe(v TargetUeId)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *AnalyticsEventSubsc) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


