# AnalyticsExposureSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyEventsSubs** | [**[]AnalyticsEventSubsc**](AnalyticsEventSubsc.md) |  | 
**AnalyRepInfo** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifId** | **string** |  | 
**EventNotifis** | Pointer to [**[]AnalyticsEventNotif**](AnalyticsEventNotif.md) |  | [optional] 
**FailEventReports** | Pointer to [**[]AnalyticsFailureEventInfo**](AnalyticsFailureEventInfo.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the AF to request the NEF to send a test notification as defined in clause 5.2.5.3 of 3GPP TS 29.122. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 

## Methods

### NewAnalyticsExposureSubsc

`func NewAnalyticsExposureSubsc(analyEventsSubs []AnalyticsEventSubsc, notifUri string, notifId string, ) *AnalyticsExposureSubsc`

NewAnalyticsExposureSubsc instantiates a new AnalyticsExposureSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsExposureSubscWithDefaults

`func NewAnalyticsExposureSubscWithDefaults() *AnalyticsExposureSubsc`

NewAnalyticsExposureSubscWithDefaults instantiates a new AnalyticsExposureSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyEventsSubs

`func (o *AnalyticsExposureSubsc) GetAnalyEventsSubs() []AnalyticsEventSubsc`

GetAnalyEventsSubs returns the AnalyEventsSubs field if non-nil, zero value otherwise.

### GetAnalyEventsSubsOk

`func (o *AnalyticsExposureSubsc) GetAnalyEventsSubsOk() (*[]AnalyticsEventSubsc, bool)`

GetAnalyEventsSubsOk returns a tuple with the AnalyEventsSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyEventsSubs

`func (o *AnalyticsExposureSubsc) SetAnalyEventsSubs(v []AnalyticsEventSubsc)`

SetAnalyEventsSubs sets AnalyEventsSubs field to given value.


### GetAnalyRepInfo

`func (o *AnalyticsExposureSubsc) GetAnalyRepInfo() ReportingInformation`

GetAnalyRepInfo returns the AnalyRepInfo field if non-nil, zero value otherwise.

### GetAnalyRepInfoOk

`func (o *AnalyticsExposureSubsc) GetAnalyRepInfoOk() (*ReportingInformation, bool)`

GetAnalyRepInfoOk returns a tuple with the AnalyRepInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyRepInfo

`func (o *AnalyticsExposureSubsc) SetAnalyRepInfo(v ReportingInformation)`

SetAnalyRepInfo sets AnalyRepInfo field to given value.

### HasAnalyRepInfo

`func (o *AnalyticsExposureSubsc) HasAnalyRepInfo() bool`

HasAnalyRepInfo returns a boolean if a field has been set.

### GetNotifUri

`func (o *AnalyticsExposureSubsc) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *AnalyticsExposureSubsc) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *AnalyticsExposureSubsc) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetNotifId

`func (o *AnalyticsExposureSubsc) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *AnalyticsExposureSubsc) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *AnalyticsExposureSubsc) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifis

`func (o *AnalyticsExposureSubsc) GetEventNotifis() []AnalyticsEventNotif`

GetEventNotifis returns the EventNotifis field if non-nil, zero value otherwise.

### GetEventNotifisOk

`func (o *AnalyticsExposureSubsc) GetEventNotifisOk() (*[]AnalyticsEventNotif, bool)`

GetEventNotifisOk returns a tuple with the EventNotifis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifis

`func (o *AnalyticsExposureSubsc) SetEventNotifis(v []AnalyticsEventNotif)`

SetEventNotifis sets EventNotifis field to given value.

### HasEventNotifis

`func (o *AnalyticsExposureSubsc) HasEventNotifis() bool`

HasEventNotifis returns a boolean if a field has been set.

### GetFailEventReports

`func (o *AnalyticsExposureSubsc) GetFailEventReports() []AnalyticsFailureEventInfo`

GetFailEventReports returns the FailEventReports field if non-nil, zero value otherwise.

### GetFailEventReportsOk

`func (o *AnalyticsExposureSubsc) GetFailEventReportsOk() (*[]AnalyticsFailureEventInfo, bool)`

GetFailEventReportsOk returns a tuple with the FailEventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailEventReports

`func (o *AnalyticsExposureSubsc) SetFailEventReports(v []AnalyticsFailureEventInfo)`

SetFailEventReports sets FailEventReports field to given value.

### HasFailEventReports

`func (o *AnalyticsExposureSubsc) HasFailEventReports() bool`

HasFailEventReports returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AnalyticsExposureSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AnalyticsExposureSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AnalyticsExposureSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AnalyticsExposureSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetSelf

`func (o *AnalyticsExposureSubsc) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AnalyticsExposureSubsc) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AnalyticsExposureSubsc) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AnalyticsExposureSubsc) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *AnalyticsExposureSubsc) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *AnalyticsExposureSubsc) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *AnalyticsExposureSubsc) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *AnalyticsExposureSubsc) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *AnalyticsExposureSubsc) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *AnalyticsExposureSubsc) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *AnalyticsExposureSubsc) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *AnalyticsExposureSubsc) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


