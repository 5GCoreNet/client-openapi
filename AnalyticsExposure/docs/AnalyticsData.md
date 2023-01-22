# AnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeStampGen** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**UeMobilityInfos** | Pointer to [**[]UeMobilityExposure**](UeMobilityExposure.md) |  | [optional] 
**UeCommInfos** | Pointer to [**[]UeCommunication**](UeCommunication.md) |  | [optional] 
**NwPerfInfos** | Pointer to [**[]NetworkPerfExposure**](NetworkPerfExposure.md) |  | [optional] 
**AbnormalInfos** | Pointer to [**[]AbnormalExposure**](AbnormalExposure.md) |  | [optional] 
**CongestInfos** | Pointer to [**[]CongestInfo**](CongestInfo.md) |  | [optional] 
**QosSustainInfos** | Pointer to [**[]QosSustainabilityExposure**](QosSustainabilityExposure.md) |  | [optional] 
**DisperInfos** | Pointer to [**[]DispersionInfo**](DispersionInfo.md) |  | [optional] 
**DnPerfInfos** | Pointer to [**[]DnPerfInfo**](DnPerfInfo.md) |  | [optional] 
**SvcExps** | Pointer to [**[]ServiceExperienceInfo**](ServiceExperienceInfo.md) |  | [optional] 
**DisperReqs** | Pointer to [**[]DispersionRequirement**](DispersionRequirement.md) |  | [optional] 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 

## Methods

### NewAnalyticsData

`func NewAnalyticsData(suppFeat string, ) *AnalyticsData`

NewAnalyticsData instantiates a new AnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsDataWithDefaults

`func NewAnalyticsDataWithDefaults() *AnalyticsData`

NewAnalyticsDataWithDefaults instantiates a new AnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *AnalyticsData) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AnalyticsData) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AnalyticsData) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *AnalyticsData) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetExpiry

`func (o *AnalyticsData) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AnalyticsData) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AnalyticsData) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AnalyticsData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetTimeStampGen

`func (o *AnalyticsData) GetTimeStampGen() time.Time`

GetTimeStampGen returns the TimeStampGen field if non-nil, zero value otherwise.

### GetTimeStampGenOk

`func (o *AnalyticsData) GetTimeStampGenOk() (*time.Time, bool)`

GetTimeStampGenOk returns a tuple with the TimeStampGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampGen

`func (o *AnalyticsData) SetTimeStampGen(v time.Time)`

SetTimeStampGen sets TimeStampGen field to given value.

### HasTimeStampGen

`func (o *AnalyticsData) HasTimeStampGen() bool`

HasTimeStampGen returns a boolean if a field has been set.

### GetUeMobilityInfos

`func (o *AnalyticsData) GetUeMobilityInfos() []UeMobilityExposure`

GetUeMobilityInfos returns the UeMobilityInfos field if non-nil, zero value otherwise.

### GetUeMobilityInfosOk

`func (o *AnalyticsData) GetUeMobilityInfosOk() (*[]UeMobilityExposure, bool)`

GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobilityInfos

`func (o *AnalyticsData) SetUeMobilityInfos(v []UeMobilityExposure)`

SetUeMobilityInfos sets UeMobilityInfos field to given value.

### HasUeMobilityInfos

`func (o *AnalyticsData) HasUeMobilityInfos() bool`

HasUeMobilityInfos returns a boolean if a field has been set.

### GetUeCommInfos

`func (o *AnalyticsData) GetUeCommInfos() []UeCommunication`

GetUeCommInfos returns the UeCommInfos field if non-nil, zero value otherwise.

### GetUeCommInfosOk

`func (o *AnalyticsData) GetUeCommInfosOk() (*[]UeCommunication, bool)`

GetUeCommInfosOk returns a tuple with the UeCommInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCommInfos

`func (o *AnalyticsData) SetUeCommInfos(v []UeCommunication)`

SetUeCommInfos sets UeCommInfos field to given value.

### HasUeCommInfos

`func (o *AnalyticsData) HasUeCommInfos() bool`

HasUeCommInfos returns a boolean if a field has been set.

### GetNwPerfInfos

`func (o *AnalyticsData) GetNwPerfInfos() []NetworkPerfExposure`

GetNwPerfInfos returns the NwPerfInfos field if non-nil, zero value otherwise.

### GetNwPerfInfosOk

`func (o *AnalyticsData) GetNwPerfInfosOk() (*[]NetworkPerfExposure, bool)`

GetNwPerfInfosOk returns a tuple with the NwPerfInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfInfos

`func (o *AnalyticsData) SetNwPerfInfos(v []NetworkPerfExposure)`

SetNwPerfInfos sets NwPerfInfos field to given value.

### HasNwPerfInfos

`func (o *AnalyticsData) HasNwPerfInfos() bool`

HasNwPerfInfos returns a boolean if a field has been set.

### GetAbnormalInfos

`func (o *AnalyticsData) GetAbnormalInfos() []AbnormalExposure`

GetAbnormalInfos returns the AbnormalInfos field if non-nil, zero value otherwise.

### GetAbnormalInfosOk

`func (o *AnalyticsData) GetAbnormalInfosOk() (*[]AbnormalExposure, bool)`

GetAbnormalInfosOk returns a tuple with the AbnormalInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalInfos

`func (o *AnalyticsData) SetAbnormalInfos(v []AbnormalExposure)`

SetAbnormalInfos sets AbnormalInfos field to given value.

### HasAbnormalInfos

`func (o *AnalyticsData) HasAbnormalInfos() bool`

HasAbnormalInfos returns a boolean if a field has been set.

### GetCongestInfos

`func (o *AnalyticsData) GetCongestInfos() []CongestInfo`

GetCongestInfos returns the CongestInfos field if non-nil, zero value otherwise.

### GetCongestInfosOk

`func (o *AnalyticsData) GetCongestInfosOk() (*[]CongestInfo, bool)`

GetCongestInfosOk returns a tuple with the CongestInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestInfos

`func (o *AnalyticsData) SetCongestInfos(v []CongestInfo)`

SetCongestInfos sets CongestInfos field to given value.

### HasCongestInfos

`func (o *AnalyticsData) HasCongestInfos() bool`

HasCongestInfos returns a boolean if a field has been set.

### GetQosSustainInfos

`func (o *AnalyticsData) GetQosSustainInfos() []QosSustainabilityExposure`

GetQosSustainInfos returns the QosSustainInfos field if non-nil, zero value otherwise.

### GetQosSustainInfosOk

`func (o *AnalyticsData) GetQosSustainInfosOk() (*[]QosSustainabilityExposure, bool)`

GetQosSustainInfosOk returns a tuple with the QosSustainInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSustainInfos

`func (o *AnalyticsData) SetQosSustainInfos(v []QosSustainabilityExposure)`

SetQosSustainInfos sets QosSustainInfos field to given value.

### HasQosSustainInfos

`func (o *AnalyticsData) HasQosSustainInfos() bool`

HasQosSustainInfos returns a boolean if a field has been set.

### GetDisperInfos

`func (o *AnalyticsData) GetDisperInfos() []DispersionInfo`

GetDisperInfos returns the DisperInfos field if non-nil, zero value otherwise.

### GetDisperInfosOk

`func (o *AnalyticsData) GetDisperInfosOk() (*[]DispersionInfo, bool)`

GetDisperInfosOk returns a tuple with the DisperInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperInfos

`func (o *AnalyticsData) SetDisperInfos(v []DispersionInfo)`

SetDisperInfos sets DisperInfos field to given value.

### HasDisperInfos

`func (o *AnalyticsData) HasDisperInfos() bool`

HasDisperInfos returns a boolean if a field has been set.

### GetDnPerfInfos

`func (o *AnalyticsData) GetDnPerfInfos() []DnPerfInfo`

GetDnPerfInfos returns the DnPerfInfos field if non-nil, zero value otherwise.

### GetDnPerfInfosOk

`func (o *AnalyticsData) GetDnPerfInfosOk() (*[]DnPerfInfo, bool)`

GetDnPerfInfosOk returns a tuple with the DnPerfInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfInfos

`func (o *AnalyticsData) SetDnPerfInfos(v []DnPerfInfo)`

SetDnPerfInfos sets DnPerfInfos field to given value.

### HasDnPerfInfos

`func (o *AnalyticsData) HasDnPerfInfos() bool`

HasDnPerfInfos returns a boolean if a field has been set.

### GetSvcExps

`func (o *AnalyticsData) GetSvcExps() []ServiceExperienceInfo`

GetSvcExps returns the SvcExps field if non-nil, zero value otherwise.

### GetSvcExpsOk

`func (o *AnalyticsData) GetSvcExpsOk() (*[]ServiceExperienceInfo, bool)`

GetSvcExpsOk returns a tuple with the SvcExps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExps

`func (o *AnalyticsData) SetSvcExps(v []ServiceExperienceInfo)`

SetSvcExps sets SvcExps field to given value.

### HasSvcExps

`func (o *AnalyticsData) HasSvcExps() bool`

HasSvcExps returns a boolean if a field has been set.

### GetDisperReqs

`func (o *AnalyticsData) GetDisperReqs() []DispersionRequirement`

GetDisperReqs returns the DisperReqs field if non-nil, zero value otherwise.

### GetDisperReqsOk

`func (o *AnalyticsData) GetDisperReqsOk() (*[]DispersionRequirement, bool)`

GetDisperReqsOk returns a tuple with the DisperReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperReqs

`func (o *AnalyticsData) SetDisperReqs(v []DispersionRequirement)`

SetDisperReqs sets DisperReqs field to given value.

### HasDisperReqs

`func (o *AnalyticsData) HasDisperReqs() bool`

HasDisperReqs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AnalyticsData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AnalyticsData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AnalyticsData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


