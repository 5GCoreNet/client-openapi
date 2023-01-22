# AnalyticsEventNotif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyEvent** | [**AnalyticsEvent**](AnalyticsEvent.md) |  | 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeStamp** | **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 
**FailNotifyCode** | Pointer to [**NwdafFailureCode**](NwdafFailureCode.md) |  | [optional] 
**RvWaitTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**UeMobilityInfos** | Pointer to [**[]UeMobilityExposure**](UeMobilityExposure.md) |  | [optional] 
**UeCommInfos** | Pointer to [**[]UeCommunication**](UeCommunication.md) |  | [optional] 
**AbnormalInfos** | Pointer to [**[]AbnormalExposure**](AbnormalExposure.md) |  | [optional] 
**CongestInfos** | Pointer to [**[]CongestInfo**](CongestInfo.md) |  | [optional] 
**NwPerfInfos** | Pointer to [**[]NetworkPerfExposure**](NetworkPerfExposure.md) |  | [optional] 
**QosSustainInfos** | Pointer to [**[]QosSustainabilityExposure**](QosSustainabilityExposure.md) |  | [optional] 
**DisperInfos** | Pointer to [**[]DispersionInfo**](DispersionInfo.md) |  | [optional] 
**DnPerfInfos** | Pointer to [**[]DnPerfInfo**](DnPerfInfo.md) |  | [optional] 
**SvcExps** | Pointer to [**[]ServiceExperienceInfo**](ServiceExperienceInfo.md) |  | [optional] 
**Start** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeStampGen** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewAnalyticsEventNotif

`func NewAnalyticsEventNotif(analyEvent AnalyticsEvent, timeStamp time.Time, ) *AnalyticsEventNotif`

NewAnalyticsEventNotif instantiates a new AnalyticsEventNotif object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsEventNotifWithDefaults

`func NewAnalyticsEventNotifWithDefaults() *AnalyticsEventNotif`

NewAnalyticsEventNotifWithDefaults instantiates a new AnalyticsEventNotif object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyEvent

`func (o *AnalyticsEventNotif) GetAnalyEvent() AnalyticsEvent`

GetAnalyEvent returns the AnalyEvent field if non-nil, zero value otherwise.

### GetAnalyEventOk

`func (o *AnalyticsEventNotif) GetAnalyEventOk() (*AnalyticsEvent, bool)`

GetAnalyEventOk returns a tuple with the AnalyEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyEvent

`func (o *AnalyticsEventNotif) SetAnalyEvent(v AnalyticsEvent)`

SetAnalyEvent sets AnalyEvent field to given value.


### GetExpiry

`func (o *AnalyticsEventNotif) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AnalyticsEventNotif) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AnalyticsEventNotif) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AnalyticsEventNotif) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetTimeStamp

`func (o *AnalyticsEventNotif) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AnalyticsEventNotif) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AnalyticsEventNotif) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetFailNotifyCode

`func (o *AnalyticsEventNotif) GetFailNotifyCode() NwdafFailureCode`

GetFailNotifyCode returns the FailNotifyCode field if non-nil, zero value otherwise.

### GetFailNotifyCodeOk

`func (o *AnalyticsEventNotif) GetFailNotifyCodeOk() (*NwdafFailureCode, bool)`

GetFailNotifyCodeOk returns a tuple with the FailNotifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailNotifyCode

`func (o *AnalyticsEventNotif) SetFailNotifyCode(v NwdafFailureCode)`

SetFailNotifyCode sets FailNotifyCode field to given value.

### HasFailNotifyCode

`func (o *AnalyticsEventNotif) HasFailNotifyCode() bool`

HasFailNotifyCode returns a boolean if a field has been set.

### GetRvWaitTime

`func (o *AnalyticsEventNotif) GetRvWaitTime() int32`

GetRvWaitTime returns the RvWaitTime field if non-nil, zero value otherwise.

### GetRvWaitTimeOk

`func (o *AnalyticsEventNotif) GetRvWaitTimeOk() (*int32, bool)`

GetRvWaitTimeOk returns a tuple with the RvWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRvWaitTime

`func (o *AnalyticsEventNotif) SetRvWaitTime(v int32)`

SetRvWaitTime sets RvWaitTime field to given value.

### HasRvWaitTime

`func (o *AnalyticsEventNotif) HasRvWaitTime() bool`

HasRvWaitTime returns a boolean if a field has been set.

### GetUeMobilityInfos

`func (o *AnalyticsEventNotif) GetUeMobilityInfos() []UeMobilityExposure`

GetUeMobilityInfos returns the UeMobilityInfos field if non-nil, zero value otherwise.

### GetUeMobilityInfosOk

`func (o *AnalyticsEventNotif) GetUeMobilityInfosOk() (*[]UeMobilityExposure, bool)`

GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobilityInfos

`func (o *AnalyticsEventNotif) SetUeMobilityInfos(v []UeMobilityExposure)`

SetUeMobilityInfos sets UeMobilityInfos field to given value.

### HasUeMobilityInfos

`func (o *AnalyticsEventNotif) HasUeMobilityInfos() bool`

HasUeMobilityInfos returns a boolean if a field has been set.

### GetUeCommInfos

`func (o *AnalyticsEventNotif) GetUeCommInfos() []UeCommunication`

GetUeCommInfos returns the UeCommInfos field if non-nil, zero value otherwise.

### GetUeCommInfosOk

`func (o *AnalyticsEventNotif) GetUeCommInfosOk() (*[]UeCommunication, bool)`

GetUeCommInfosOk returns a tuple with the UeCommInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCommInfos

`func (o *AnalyticsEventNotif) SetUeCommInfos(v []UeCommunication)`

SetUeCommInfos sets UeCommInfos field to given value.

### HasUeCommInfos

`func (o *AnalyticsEventNotif) HasUeCommInfos() bool`

HasUeCommInfos returns a boolean if a field has been set.

### GetAbnormalInfos

`func (o *AnalyticsEventNotif) GetAbnormalInfos() []AbnormalExposure`

GetAbnormalInfos returns the AbnormalInfos field if non-nil, zero value otherwise.

### GetAbnormalInfosOk

`func (o *AnalyticsEventNotif) GetAbnormalInfosOk() (*[]AbnormalExposure, bool)`

GetAbnormalInfosOk returns a tuple with the AbnormalInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalInfos

`func (o *AnalyticsEventNotif) SetAbnormalInfos(v []AbnormalExposure)`

SetAbnormalInfos sets AbnormalInfos field to given value.

### HasAbnormalInfos

`func (o *AnalyticsEventNotif) HasAbnormalInfos() bool`

HasAbnormalInfos returns a boolean if a field has been set.

### GetCongestInfos

`func (o *AnalyticsEventNotif) GetCongestInfos() []CongestInfo`

GetCongestInfos returns the CongestInfos field if non-nil, zero value otherwise.

### GetCongestInfosOk

`func (o *AnalyticsEventNotif) GetCongestInfosOk() (*[]CongestInfo, bool)`

GetCongestInfosOk returns a tuple with the CongestInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestInfos

`func (o *AnalyticsEventNotif) SetCongestInfos(v []CongestInfo)`

SetCongestInfos sets CongestInfos field to given value.

### HasCongestInfos

`func (o *AnalyticsEventNotif) HasCongestInfos() bool`

HasCongestInfos returns a boolean if a field has been set.

### GetNwPerfInfos

`func (o *AnalyticsEventNotif) GetNwPerfInfos() []NetworkPerfExposure`

GetNwPerfInfos returns the NwPerfInfos field if non-nil, zero value otherwise.

### GetNwPerfInfosOk

`func (o *AnalyticsEventNotif) GetNwPerfInfosOk() (*[]NetworkPerfExposure, bool)`

GetNwPerfInfosOk returns a tuple with the NwPerfInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfInfos

`func (o *AnalyticsEventNotif) SetNwPerfInfos(v []NetworkPerfExposure)`

SetNwPerfInfos sets NwPerfInfos field to given value.

### HasNwPerfInfos

`func (o *AnalyticsEventNotif) HasNwPerfInfos() bool`

HasNwPerfInfos returns a boolean if a field has been set.

### GetQosSustainInfos

`func (o *AnalyticsEventNotif) GetQosSustainInfos() []QosSustainabilityExposure`

GetQosSustainInfos returns the QosSustainInfos field if non-nil, zero value otherwise.

### GetQosSustainInfosOk

`func (o *AnalyticsEventNotif) GetQosSustainInfosOk() (*[]QosSustainabilityExposure, bool)`

GetQosSustainInfosOk returns a tuple with the QosSustainInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSustainInfos

`func (o *AnalyticsEventNotif) SetQosSustainInfos(v []QosSustainabilityExposure)`

SetQosSustainInfos sets QosSustainInfos field to given value.

### HasQosSustainInfos

`func (o *AnalyticsEventNotif) HasQosSustainInfos() bool`

HasQosSustainInfos returns a boolean if a field has been set.

### GetDisperInfos

`func (o *AnalyticsEventNotif) GetDisperInfos() []DispersionInfo`

GetDisperInfos returns the DisperInfos field if non-nil, zero value otherwise.

### GetDisperInfosOk

`func (o *AnalyticsEventNotif) GetDisperInfosOk() (*[]DispersionInfo, bool)`

GetDisperInfosOk returns a tuple with the DisperInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperInfos

`func (o *AnalyticsEventNotif) SetDisperInfos(v []DispersionInfo)`

SetDisperInfos sets DisperInfos field to given value.

### HasDisperInfos

`func (o *AnalyticsEventNotif) HasDisperInfos() bool`

HasDisperInfos returns a boolean if a field has been set.

### GetDnPerfInfos

`func (o *AnalyticsEventNotif) GetDnPerfInfos() []DnPerfInfo`

GetDnPerfInfos returns the DnPerfInfos field if non-nil, zero value otherwise.

### GetDnPerfInfosOk

`func (o *AnalyticsEventNotif) GetDnPerfInfosOk() (*[]DnPerfInfo, bool)`

GetDnPerfInfosOk returns a tuple with the DnPerfInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfInfos

`func (o *AnalyticsEventNotif) SetDnPerfInfos(v []DnPerfInfo)`

SetDnPerfInfos sets DnPerfInfos field to given value.

### HasDnPerfInfos

`func (o *AnalyticsEventNotif) HasDnPerfInfos() bool`

HasDnPerfInfos returns a boolean if a field has been set.

### GetSvcExps

`func (o *AnalyticsEventNotif) GetSvcExps() []ServiceExperienceInfo`

GetSvcExps returns the SvcExps field if non-nil, zero value otherwise.

### GetSvcExpsOk

`func (o *AnalyticsEventNotif) GetSvcExpsOk() (*[]ServiceExperienceInfo, bool)`

GetSvcExpsOk returns a tuple with the SvcExps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExps

`func (o *AnalyticsEventNotif) SetSvcExps(v []ServiceExperienceInfo)`

SetSvcExps sets SvcExps field to given value.

### HasSvcExps

`func (o *AnalyticsEventNotif) HasSvcExps() bool`

HasSvcExps returns a boolean if a field has been set.

### GetStart

`func (o *AnalyticsEventNotif) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AnalyticsEventNotif) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AnalyticsEventNotif) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *AnalyticsEventNotif) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTimeStampGen

`func (o *AnalyticsEventNotif) GetTimeStampGen() time.Time`

GetTimeStampGen returns the TimeStampGen field if non-nil, zero value otherwise.

### GetTimeStampGenOk

`func (o *AnalyticsEventNotif) GetTimeStampGenOk() (*time.Time, bool)`

GetTimeStampGenOk returns a tuple with the TimeStampGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampGen

`func (o *AnalyticsEventNotif) SetTimeStampGen(v time.Time)`

SetTimeStampGen sets TimeStampGen field to given value.

### HasTimeStampGen

`func (o *AnalyticsEventNotif) HasTimeStampGen() bool`

HasTimeStampGen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


