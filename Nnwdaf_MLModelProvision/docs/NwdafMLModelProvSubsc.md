# NwdafMLModelProvSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MLEventSubscs** | [**[]MLEventSubscription**](MLEventSubscription.md) | Subscribed events | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**MLEventNotifs** | Pointer to [**[]MLEventNotif**](MLEventNotif.md) |  | [optional] 
**SuppFeats** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**NotifCorreId** | Pointer to **string** |  | [optional] 
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**FailEventReports** | Pointer to [**[]FailureEventInfoForMLModel**](FailureEventInfoForMLModel.md) |  | [optional] 

## Methods

### NewNwdafMLModelProvSubsc

`func NewNwdafMLModelProvSubsc(mLEventSubscs []MLEventSubscription, notifUri string, ) *NwdafMLModelProvSubsc`

NewNwdafMLModelProvSubsc instantiates a new NwdafMLModelProvSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwdafMLModelProvSubscWithDefaults

`func NewNwdafMLModelProvSubscWithDefaults() *NwdafMLModelProvSubsc`

NewNwdafMLModelProvSubscWithDefaults instantiates a new NwdafMLModelProvSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMLEventSubscs

`func (o *NwdafMLModelProvSubsc) GetMLEventSubscs() []MLEventSubscription`

GetMLEventSubscs returns the MLEventSubscs field if non-nil, zero value otherwise.

### GetMLEventSubscsOk

`func (o *NwdafMLModelProvSubsc) GetMLEventSubscsOk() (*[]MLEventSubscription, bool)`

GetMLEventSubscsOk returns a tuple with the MLEventSubscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLEventSubscs

`func (o *NwdafMLModelProvSubsc) SetMLEventSubscs(v []MLEventSubscription)`

SetMLEventSubscs sets MLEventSubscs field to given value.


### GetNotifUri

`func (o *NwdafMLModelProvSubsc) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *NwdafMLModelProvSubsc) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *NwdafMLModelProvSubsc) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetMLEventNotifs

`func (o *NwdafMLModelProvSubsc) GetMLEventNotifs() []MLEventNotif`

GetMLEventNotifs returns the MLEventNotifs field if non-nil, zero value otherwise.

### GetMLEventNotifsOk

`func (o *NwdafMLModelProvSubsc) GetMLEventNotifsOk() (*[]MLEventNotif, bool)`

GetMLEventNotifsOk returns a tuple with the MLEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLEventNotifs

`func (o *NwdafMLModelProvSubsc) SetMLEventNotifs(v []MLEventNotif)`

SetMLEventNotifs sets MLEventNotifs field to given value.

### HasMLEventNotifs

`func (o *NwdafMLModelProvSubsc) HasMLEventNotifs() bool`

HasMLEventNotifs returns a boolean if a field has been set.

### GetSuppFeats

`func (o *NwdafMLModelProvSubsc) GetSuppFeats() string`

GetSuppFeats returns the SuppFeats field if non-nil, zero value otherwise.

### GetSuppFeatsOk

`func (o *NwdafMLModelProvSubsc) GetSuppFeatsOk() (*string, bool)`

GetSuppFeatsOk returns a tuple with the SuppFeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeats

`func (o *NwdafMLModelProvSubsc) SetSuppFeats(v string)`

SetSuppFeats sets SuppFeats field to given value.

### HasSuppFeats

`func (o *NwdafMLModelProvSubsc) HasSuppFeats() bool`

HasSuppFeats returns a boolean if a field has been set.

### GetNotifCorreId

`func (o *NwdafMLModelProvSubsc) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *NwdafMLModelProvSubsc) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *NwdafMLModelProvSubsc) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.

### HasNotifCorreId

`func (o *NwdafMLModelProvSubsc) HasNotifCorreId() bool`

HasNotifCorreId returns a boolean if a field has been set.

### GetEventReq

`func (o *NwdafMLModelProvSubsc) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *NwdafMLModelProvSubsc) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *NwdafMLModelProvSubsc) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *NwdafMLModelProvSubsc) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.

### GetFailEventReports

`func (o *NwdafMLModelProvSubsc) GetFailEventReports() []FailureEventInfoForMLModel`

GetFailEventReports returns the FailEventReports field if non-nil, zero value otherwise.

### GetFailEventReportsOk

`func (o *NwdafMLModelProvSubsc) GetFailEventReportsOk() (*[]FailureEventInfoForMLModel, bool)`

GetFailEventReportsOk returns a tuple with the FailEventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailEventReports

`func (o *NwdafMLModelProvSubsc) SetFailEventReports(v []FailureEventInfoForMLModel)`

SetFailEventReports sets FailEventReports field to given value.

### HasFailEventReports

`func (o *NwdafMLModelProvSubsc) HasFailEventReports() bool`

HasFailEventReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


