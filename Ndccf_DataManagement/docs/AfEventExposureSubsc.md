# AfEventExposureSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataAccProfId** | Pointer to **string** |  | [optional] 
**EventsSubs** | [**[]EventsSubs**](EventsSubs.md) |  | 
**EventsRepInfo** | [**ReportingInformation**](ReportingInformation.md) |  | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifId** | **string** |  | 
**EventNotifs** | Pointer to [**[]AfEventNotification**](AfEventNotification.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAfEventExposureSubsc

`func NewAfEventExposureSubsc(eventsSubs []EventsSubs, eventsRepInfo ReportingInformation, notifUri string, notifId string, ) *AfEventExposureSubsc`

NewAfEventExposureSubsc instantiates a new AfEventExposureSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfEventExposureSubscWithDefaults

`func NewAfEventExposureSubscWithDefaults() *AfEventExposureSubsc`

NewAfEventExposureSubscWithDefaults instantiates a new AfEventExposureSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataAccProfId

`func (o *AfEventExposureSubsc) GetDataAccProfId() string`

GetDataAccProfId returns the DataAccProfId field if non-nil, zero value otherwise.

### GetDataAccProfIdOk

`func (o *AfEventExposureSubsc) GetDataAccProfIdOk() (*string, bool)`

GetDataAccProfIdOk returns a tuple with the DataAccProfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccProfId

`func (o *AfEventExposureSubsc) SetDataAccProfId(v string)`

SetDataAccProfId sets DataAccProfId field to given value.

### HasDataAccProfId

`func (o *AfEventExposureSubsc) HasDataAccProfId() bool`

HasDataAccProfId returns a boolean if a field has been set.

### GetEventsSubs

`func (o *AfEventExposureSubsc) GetEventsSubs() []EventsSubs`

GetEventsSubs returns the EventsSubs field if non-nil, zero value otherwise.

### GetEventsSubsOk

`func (o *AfEventExposureSubsc) GetEventsSubsOk() (*[]EventsSubs, bool)`

GetEventsSubsOk returns a tuple with the EventsSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsSubs

`func (o *AfEventExposureSubsc) SetEventsSubs(v []EventsSubs)`

SetEventsSubs sets EventsSubs field to given value.


### GetEventsRepInfo

`func (o *AfEventExposureSubsc) GetEventsRepInfo() ReportingInformation`

GetEventsRepInfo returns the EventsRepInfo field if non-nil, zero value otherwise.

### GetEventsRepInfoOk

`func (o *AfEventExposureSubsc) GetEventsRepInfoOk() (*ReportingInformation, bool)`

GetEventsRepInfoOk returns a tuple with the EventsRepInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRepInfo

`func (o *AfEventExposureSubsc) SetEventsRepInfo(v ReportingInformation)`

SetEventsRepInfo sets EventsRepInfo field to given value.


### GetNotifUri

`func (o *AfEventExposureSubsc) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *AfEventExposureSubsc) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *AfEventExposureSubsc) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetNotifId

`func (o *AfEventExposureSubsc) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *AfEventExposureSubsc) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *AfEventExposureSubsc) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *AfEventExposureSubsc) GetEventNotifs() []AfEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *AfEventExposureSubsc) GetEventNotifsOk() (*[]AfEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *AfEventExposureSubsc) SetEventNotifs(v []AfEventNotification)`

SetEventNotifs sets EventNotifs field to given value.

### HasEventNotifs

`func (o *AfEventExposureSubsc) HasEventNotifs() bool`

HasEventNotifs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AfEventExposureSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AfEventExposureSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AfEventExposureSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AfEventExposureSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


