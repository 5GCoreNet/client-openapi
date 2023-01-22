# PcEventExposureSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubs** | [**[]PcEvent**](PcEvent.md) |  | 
**EventsRepInfo** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**GroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**FilterDnns** | Pointer to **[]string** |  | [optional] 
**FilterSnssais** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**SnssaiDnns** | Pointer to [**[]SnssaiDnnCombination**](SnssaiDnnCombination.md) |  | [optional] 
**FilterServices** | Pointer to [**[]ServiceIdentification**](ServiceIdentification.md) |  | [optional] 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifId** | **string** |  | 
**EventNotifs** | Pointer to [**[]PcEventNotification**](PcEventNotification.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewPcEventExposureSubsc

`func NewPcEventExposureSubsc(eventSubs []PcEvent, notifUri string, notifId string, ) *PcEventExposureSubsc`

NewPcEventExposureSubsc instantiates a new PcEventExposureSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcEventExposureSubscWithDefaults

`func NewPcEventExposureSubscWithDefaults() *PcEventExposureSubsc`

NewPcEventExposureSubscWithDefaults instantiates a new PcEventExposureSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubs

`func (o *PcEventExposureSubsc) GetEventSubs() []PcEvent`

GetEventSubs returns the EventSubs field if non-nil, zero value otherwise.

### GetEventSubsOk

`func (o *PcEventExposureSubsc) GetEventSubsOk() (*[]PcEvent, bool)`

GetEventSubsOk returns a tuple with the EventSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubs

`func (o *PcEventExposureSubsc) SetEventSubs(v []PcEvent)`

SetEventSubs sets EventSubs field to given value.


### GetEventsRepInfo

`func (o *PcEventExposureSubsc) GetEventsRepInfo() ReportingInformation`

GetEventsRepInfo returns the EventsRepInfo field if non-nil, zero value otherwise.

### GetEventsRepInfoOk

`func (o *PcEventExposureSubsc) GetEventsRepInfoOk() (*ReportingInformation, bool)`

GetEventsRepInfoOk returns a tuple with the EventsRepInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRepInfo

`func (o *PcEventExposureSubsc) SetEventsRepInfo(v ReportingInformation)`

SetEventsRepInfo sets EventsRepInfo field to given value.

### HasEventsRepInfo

`func (o *PcEventExposureSubsc) HasEventsRepInfo() bool`

HasEventsRepInfo returns a boolean if a field has been set.

### GetGroupId

`func (o *PcEventExposureSubsc) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PcEventExposureSubsc) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PcEventExposureSubsc) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PcEventExposureSubsc) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetFilterDnns

`func (o *PcEventExposureSubsc) GetFilterDnns() []string`

GetFilterDnns returns the FilterDnns field if non-nil, zero value otherwise.

### GetFilterDnnsOk

`func (o *PcEventExposureSubsc) GetFilterDnnsOk() (*[]string, bool)`

GetFilterDnnsOk returns a tuple with the FilterDnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDnns

`func (o *PcEventExposureSubsc) SetFilterDnns(v []string)`

SetFilterDnns sets FilterDnns field to given value.

### HasFilterDnns

`func (o *PcEventExposureSubsc) HasFilterDnns() bool`

HasFilterDnns returns a boolean if a field has been set.

### GetFilterSnssais

`func (o *PcEventExposureSubsc) GetFilterSnssais() []Snssai`

GetFilterSnssais returns the FilterSnssais field if non-nil, zero value otherwise.

### GetFilterSnssaisOk

`func (o *PcEventExposureSubsc) GetFilterSnssaisOk() (*[]Snssai, bool)`

GetFilterSnssaisOk returns a tuple with the FilterSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSnssais

`func (o *PcEventExposureSubsc) SetFilterSnssais(v []Snssai)`

SetFilterSnssais sets FilterSnssais field to given value.

### HasFilterSnssais

`func (o *PcEventExposureSubsc) HasFilterSnssais() bool`

HasFilterSnssais returns a boolean if a field has been set.

### GetSnssaiDnns

`func (o *PcEventExposureSubsc) GetSnssaiDnns() []SnssaiDnnCombination`

GetSnssaiDnns returns the SnssaiDnns field if non-nil, zero value otherwise.

### GetSnssaiDnnsOk

`func (o *PcEventExposureSubsc) GetSnssaiDnnsOk() (*[]SnssaiDnnCombination, bool)`

GetSnssaiDnnsOk returns a tuple with the SnssaiDnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiDnns

`func (o *PcEventExposureSubsc) SetSnssaiDnns(v []SnssaiDnnCombination)`

SetSnssaiDnns sets SnssaiDnns field to given value.

### HasSnssaiDnns

`func (o *PcEventExposureSubsc) HasSnssaiDnns() bool`

HasSnssaiDnns returns a boolean if a field has been set.

### GetFilterServices

`func (o *PcEventExposureSubsc) GetFilterServices() []ServiceIdentification`

GetFilterServices returns the FilterServices field if non-nil, zero value otherwise.

### GetFilterServicesOk

`func (o *PcEventExposureSubsc) GetFilterServicesOk() (*[]ServiceIdentification, bool)`

GetFilterServicesOk returns a tuple with the FilterServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterServices

`func (o *PcEventExposureSubsc) SetFilterServices(v []ServiceIdentification)`

SetFilterServices sets FilterServices field to given value.

### HasFilterServices

`func (o *PcEventExposureSubsc) HasFilterServices() bool`

HasFilterServices returns a boolean if a field has been set.

### GetNotifUri

`func (o *PcEventExposureSubsc) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *PcEventExposureSubsc) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *PcEventExposureSubsc) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetNotifId

`func (o *PcEventExposureSubsc) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *PcEventExposureSubsc) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *PcEventExposureSubsc) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *PcEventExposureSubsc) GetEventNotifs() []PcEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *PcEventExposureSubsc) GetEventNotifsOk() (*[]PcEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *PcEventExposureSubsc) SetEventNotifs(v []PcEventNotification)`

SetEventNotifs sets EventNotifs field to given value.

### HasEventNotifs

`func (o *PcEventExposureSubsc) HasEventNotifs() bool`

HasEventNotifs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PcEventExposureSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PcEventExposureSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PcEventExposureSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *PcEventExposureSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


