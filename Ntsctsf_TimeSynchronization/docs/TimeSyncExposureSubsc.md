# TimeSyncExposureSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supis** | Pointer to **[]string** |  | [optional] 
**Gpsis** | Pointer to **[]string** |  | [optional] 
**InterGrpId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**ExterGrpId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**AnyUeInd** | Pointer to **bool** | Identifies whether the request applies to any UE. This attribute shall set to \&quot;true\&quot; if  applicable for any UE, otherwise, set to \&quot;false\&quot;.  | [optional] 
**NotifMethod** | Pointer to [**NotificationMethod**](NotificationMethod.md) |  | [optional] 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**SubscribedEvents** | [**[]SubscribedEvent**](SubscribedEvent.md) |  | 
**EventFilters** | Pointer to [**[]EventFilter**](EventFilter.md) |  | [optional] 
**SubsNotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**SubsNotifId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**MaxReportNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewTimeSyncExposureSubsc

`func NewTimeSyncExposureSubsc(dnn string, snssai Snssai, subscribedEvents []SubscribedEvent, subsNotifUri string, subsNotifId string, ) *TimeSyncExposureSubsc`

NewTimeSyncExposureSubsc instantiates a new TimeSyncExposureSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncExposureSubscWithDefaults

`func NewTimeSyncExposureSubscWithDefaults() *TimeSyncExposureSubsc`

NewTimeSyncExposureSubscWithDefaults instantiates a new TimeSyncExposureSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupis

`func (o *TimeSyncExposureSubsc) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *TimeSyncExposureSubsc) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *TimeSyncExposureSubsc) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *TimeSyncExposureSubsc) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetGpsis

`func (o *TimeSyncExposureSubsc) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *TimeSyncExposureSubsc) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *TimeSyncExposureSubsc) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *TimeSyncExposureSubsc) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetInterGrpId

`func (o *TimeSyncExposureSubsc) GetInterGrpId() string`

GetInterGrpId returns the InterGrpId field if non-nil, zero value otherwise.

### GetInterGrpIdOk

`func (o *TimeSyncExposureSubsc) GetInterGrpIdOk() (*string, bool)`

GetInterGrpIdOk returns a tuple with the InterGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGrpId

`func (o *TimeSyncExposureSubsc) SetInterGrpId(v string)`

SetInterGrpId sets InterGrpId field to given value.

### HasInterGrpId

`func (o *TimeSyncExposureSubsc) HasInterGrpId() bool`

HasInterGrpId returns a boolean if a field has been set.

### GetExterGrpId

`func (o *TimeSyncExposureSubsc) GetExterGrpId() string`

GetExterGrpId returns the ExterGrpId field if non-nil, zero value otherwise.

### GetExterGrpIdOk

`func (o *TimeSyncExposureSubsc) GetExterGrpIdOk() (*string, bool)`

GetExterGrpIdOk returns a tuple with the ExterGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGrpId

`func (o *TimeSyncExposureSubsc) SetExterGrpId(v string)`

SetExterGrpId sets ExterGrpId field to given value.

### HasExterGrpId

`func (o *TimeSyncExposureSubsc) HasExterGrpId() bool`

HasExterGrpId returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *TimeSyncExposureSubsc) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *TimeSyncExposureSubsc) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *TimeSyncExposureSubsc) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *TimeSyncExposureSubsc) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetNotifMethod

`func (o *TimeSyncExposureSubsc) GetNotifMethod() NotificationMethod`

GetNotifMethod returns the NotifMethod field if non-nil, zero value otherwise.

### GetNotifMethodOk

`func (o *TimeSyncExposureSubsc) GetNotifMethodOk() (*NotificationMethod, bool)`

GetNotifMethodOk returns a tuple with the NotifMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifMethod

`func (o *TimeSyncExposureSubsc) SetNotifMethod(v NotificationMethod)`

SetNotifMethod sets NotifMethod field to given value.

### HasNotifMethod

`func (o *TimeSyncExposureSubsc) HasNotifMethod() bool`

HasNotifMethod returns a boolean if a field has been set.

### GetDnn

`func (o *TimeSyncExposureSubsc) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TimeSyncExposureSubsc) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TimeSyncExposureSubsc) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSnssai

`func (o *TimeSyncExposureSubsc) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TimeSyncExposureSubsc) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TimeSyncExposureSubsc) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetSubscribedEvents

`func (o *TimeSyncExposureSubsc) GetSubscribedEvents() []SubscribedEvent`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *TimeSyncExposureSubsc) GetSubscribedEventsOk() (*[]SubscribedEvent, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *TimeSyncExposureSubsc) SetSubscribedEvents(v []SubscribedEvent)`

SetSubscribedEvents sets SubscribedEvents field to given value.


### GetEventFilters

`func (o *TimeSyncExposureSubsc) GetEventFilters() []EventFilter`

GetEventFilters returns the EventFilters field if non-nil, zero value otherwise.

### GetEventFiltersOk

`func (o *TimeSyncExposureSubsc) GetEventFiltersOk() (*[]EventFilter, bool)`

GetEventFiltersOk returns a tuple with the EventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilters

`func (o *TimeSyncExposureSubsc) SetEventFilters(v []EventFilter)`

SetEventFilters sets EventFilters field to given value.

### HasEventFilters

`func (o *TimeSyncExposureSubsc) HasEventFilters() bool`

HasEventFilters returns a boolean if a field has been set.

### GetSubsNotifUri

`func (o *TimeSyncExposureSubsc) GetSubsNotifUri() string`

GetSubsNotifUri returns the SubsNotifUri field if non-nil, zero value otherwise.

### GetSubsNotifUriOk

`func (o *TimeSyncExposureSubsc) GetSubsNotifUriOk() (*string, bool)`

GetSubsNotifUriOk returns a tuple with the SubsNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsNotifUri

`func (o *TimeSyncExposureSubsc) SetSubsNotifUri(v string)`

SetSubsNotifUri sets SubsNotifUri field to given value.


### GetSubsNotifId

`func (o *TimeSyncExposureSubsc) GetSubsNotifId() string`

GetSubsNotifId returns the SubsNotifId field if non-nil, zero value otherwise.

### GetSubsNotifIdOk

`func (o *TimeSyncExposureSubsc) GetSubsNotifIdOk() (*string, bool)`

GetSubsNotifIdOk returns a tuple with the SubsNotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsNotifId

`func (o *TimeSyncExposureSubsc) SetSubsNotifId(v string)`

SetSubsNotifId sets SubsNotifId field to given value.


### GetMaxReportNbr

`func (o *TimeSyncExposureSubsc) GetMaxReportNbr() int32`

GetMaxReportNbr returns the MaxReportNbr field if non-nil, zero value otherwise.

### GetMaxReportNbrOk

`func (o *TimeSyncExposureSubsc) GetMaxReportNbrOk() (*int32, bool)`

GetMaxReportNbrOk returns a tuple with the MaxReportNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReportNbr

`func (o *TimeSyncExposureSubsc) SetMaxReportNbr(v int32)`

SetMaxReportNbr sets MaxReportNbr field to given value.

### HasMaxReportNbr

`func (o *TimeSyncExposureSubsc) HasMaxReportNbr() bool`

HasMaxReportNbr returns a boolean if a field has been set.

### GetExpiry

`func (o *TimeSyncExposureSubsc) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *TimeSyncExposureSubsc) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *TimeSyncExposureSubsc) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *TimeSyncExposureSubsc) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRepPeriod

`func (o *TimeSyncExposureSubsc) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *TimeSyncExposureSubsc) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *TimeSyncExposureSubsc) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *TimeSyncExposureSubsc) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetSuppFeat

`func (o *TimeSyncExposureSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *TimeSyncExposureSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *TimeSyncExposureSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *TimeSyncExposureSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


