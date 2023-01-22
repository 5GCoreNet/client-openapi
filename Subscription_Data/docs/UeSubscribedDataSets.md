# UeSubscribedDataSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmData** | Pointer to [**AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md) |  | [optional] 
**SmfSelData** | Pointer to [**SmfSelectionSubscriptionData**](SmfSelectionSubscriptionData.md) |  | [optional] 
**SmsSubsData** | Pointer to [**SmsSubscriptionData**](SmsSubscriptionData.md) |  | [optional] 
**SmData** | Pointer to [**SmSubsData**](SmSubsData.md) |  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**SmsMngData** | Pointer to [**SmsManagementSubscriptionData**](SmsManagementSubscriptionData.md) |  | [optional] 
**LcsPrivacyData** | Pointer to [**LcsPrivacyData**](LcsPrivacyData.md) |  | [optional] 
**LcsMoData** | Pointer to [**LcsMoData**](LcsMoData.md) |  | [optional] 
**LcsBcaData** | Pointer to [**LcsBroadcastAssistanceTypesData**](LcsBroadcastAssistanceTypesData.md) |  | [optional] 
**V2xData** | Pointer to [**V2xSubscriptionData**](V2xSubscriptionData.md) |  | [optional] 
**ProseData** | Pointer to [**ProseSubscriptionData**](ProseSubscriptionData.md) |  | [optional] 
**OdbData** | Pointer to [**OdbData**](OdbData.md) |  | [optional] 
**EeProfileData** | Pointer to [**EeProfileData**](EeProfileData.md) |  | [optional] 
**PpProfileData** | Pointer to [**PpProfileData**](PpProfileData.md) |  | [optional] 
**NiddAuthData** | Pointer to [**AuthorizationData**](AuthorizationData.md) |  | [optional] 
**UcData** | Pointer to [**UcSubscriptionData**](UcSubscriptionData.md) |  | [optional] 
**MbsSubscriptionData** | Pointer to [**MbsSubscriptionData**](MbsSubscriptionData.md) |  | [optional] 

## Methods

### NewUeSubscribedDataSets

`func NewUeSubscribedDataSets() *UeSubscribedDataSets`

NewUeSubscribedDataSets instantiates a new UeSubscribedDataSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeSubscribedDataSetsWithDefaults

`func NewUeSubscribedDataSetsWithDefaults() *UeSubscribedDataSets`

NewUeSubscribedDataSetsWithDefaults instantiates a new UeSubscribedDataSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmData

`func (o *UeSubscribedDataSets) GetAmData() AccessAndMobilitySubscriptionData`

GetAmData returns the AmData field if non-nil, zero value otherwise.

### GetAmDataOk

`func (o *UeSubscribedDataSets) GetAmDataOk() (*AccessAndMobilitySubscriptionData, bool)`

GetAmDataOk returns a tuple with the AmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmData

`func (o *UeSubscribedDataSets) SetAmData(v AccessAndMobilitySubscriptionData)`

SetAmData sets AmData field to given value.

### HasAmData

`func (o *UeSubscribedDataSets) HasAmData() bool`

HasAmData returns a boolean if a field has been set.

### GetSmfSelData

`func (o *UeSubscribedDataSets) GetSmfSelData() SmfSelectionSubscriptionData`

GetSmfSelData returns the SmfSelData field if non-nil, zero value otherwise.

### GetSmfSelDataOk

`func (o *UeSubscribedDataSets) GetSmfSelDataOk() (*SmfSelectionSubscriptionData, bool)`

GetSmfSelDataOk returns a tuple with the SmfSelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelData

`func (o *UeSubscribedDataSets) SetSmfSelData(v SmfSelectionSubscriptionData)`

SetSmfSelData sets SmfSelData field to given value.

### HasSmfSelData

`func (o *UeSubscribedDataSets) HasSmfSelData() bool`

HasSmfSelData returns a boolean if a field has been set.

### GetSmsSubsData

`func (o *UeSubscribedDataSets) GetSmsSubsData() SmsSubscriptionData`

GetSmsSubsData returns the SmsSubsData field if non-nil, zero value otherwise.

### GetSmsSubsDataOk

`func (o *UeSubscribedDataSets) GetSmsSubsDataOk() (*SmsSubscriptionData, bool)`

GetSmsSubsDataOk returns a tuple with the SmsSubsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSubsData

`func (o *UeSubscribedDataSets) SetSmsSubsData(v SmsSubscriptionData)`

SetSmsSubsData sets SmsSubsData field to given value.

### HasSmsSubsData

`func (o *UeSubscribedDataSets) HasSmsSubsData() bool`

HasSmsSubsData returns a boolean if a field has been set.

### GetSmData

`func (o *UeSubscribedDataSets) GetSmData() SmSubsData`

GetSmData returns the SmData field if non-nil, zero value otherwise.

### GetSmDataOk

`func (o *UeSubscribedDataSets) GetSmDataOk() (*SmSubsData, bool)`

GetSmDataOk returns a tuple with the SmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmData

`func (o *UeSubscribedDataSets) SetSmData(v SmSubsData)`

SetSmData sets SmData field to given value.

### HasSmData

`func (o *UeSubscribedDataSets) HasSmData() bool`

HasSmData returns a boolean if a field has been set.

### GetTraceData

`func (o *UeSubscribedDataSets) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *UeSubscribedDataSets) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *UeSubscribedDataSets) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *UeSubscribedDataSets) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *UeSubscribedDataSets) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *UeSubscribedDataSets) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetSmsMngData

`func (o *UeSubscribedDataSets) GetSmsMngData() SmsManagementSubscriptionData`

GetSmsMngData returns the SmsMngData field if non-nil, zero value otherwise.

### GetSmsMngDataOk

`func (o *UeSubscribedDataSets) GetSmsMngDataOk() (*SmsManagementSubscriptionData, bool)`

GetSmsMngDataOk returns a tuple with the SmsMngData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMngData

`func (o *UeSubscribedDataSets) SetSmsMngData(v SmsManagementSubscriptionData)`

SetSmsMngData sets SmsMngData field to given value.

### HasSmsMngData

`func (o *UeSubscribedDataSets) HasSmsMngData() bool`

HasSmsMngData returns a boolean if a field has been set.

### GetLcsPrivacyData

`func (o *UeSubscribedDataSets) GetLcsPrivacyData() LcsPrivacyData`

GetLcsPrivacyData returns the LcsPrivacyData field if non-nil, zero value otherwise.

### GetLcsPrivacyDataOk

`func (o *UeSubscribedDataSets) GetLcsPrivacyDataOk() (*LcsPrivacyData, bool)`

GetLcsPrivacyDataOk returns a tuple with the LcsPrivacyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsPrivacyData

`func (o *UeSubscribedDataSets) SetLcsPrivacyData(v LcsPrivacyData)`

SetLcsPrivacyData sets LcsPrivacyData field to given value.

### HasLcsPrivacyData

`func (o *UeSubscribedDataSets) HasLcsPrivacyData() bool`

HasLcsPrivacyData returns a boolean if a field has been set.

### GetLcsMoData

`func (o *UeSubscribedDataSets) GetLcsMoData() LcsMoData`

GetLcsMoData returns the LcsMoData field if non-nil, zero value otherwise.

### GetLcsMoDataOk

`func (o *UeSubscribedDataSets) GetLcsMoDataOk() (*LcsMoData, bool)`

GetLcsMoDataOk returns a tuple with the LcsMoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsMoData

`func (o *UeSubscribedDataSets) SetLcsMoData(v LcsMoData)`

SetLcsMoData sets LcsMoData field to given value.

### HasLcsMoData

`func (o *UeSubscribedDataSets) HasLcsMoData() bool`

HasLcsMoData returns a boolean if a field has been set.

### GetLcsBcaData

`func (o *UeSubscribedDataSets) GetLcsBcaData() LcsBroadcastAssistanceTypesData`

GetLcsBcaData returns the LcsBcaData field if non-nil, zero value otherwise.

### GetLcsBcaDataOk

`func (o *UeSubscribedDataSets) GetLcsBcaDataOk() (*LcsBroadcastAssistanceTypesData, bool)`

GetLcsBcaDataOk returns a tuple with the LcsBcaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsBcaData

`func (o *UeSubscribedDataSets) SetLcsBcaData(v LcsBroadcastAssistanceTypesData)`

SetLcsBcaData sets LcsBcaData field to given value.

### HasLcsBcaData

`func (o *UeSubscribedDataSets) HasLcsBcaData() bool`

HasLcsBcaData returns a boolean if a field has been set.

### GetV2xData

`func (o *UeSubscribedDataSets) GetV2xData() V2xSubscriptionData`

GetV2xData returns the V2xData field if non-nil, zero value otherwise.

### GetV2xDataOk

`func (o *UeSubscribedDataSets) GetV2xDataOk() (*V2xSubscriptionData, bool)`

GetV2xDataOk returns a tuple with the V2xData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xData

`func (o *UeSubscribedDataSets) SetV2xData(v V2xSubscriptionData)`

SetV2xData sets V2xData field to given value.

### HasV2xData

`func (o *UeSubscribedDataSets) HasV2xData() bool`

HasV2xData returns a boolean if a field has been set.

### GetProseData

`func (o *UeSubscribedDataSets) GetProseData() ProseSubscriptionData`

GetProseData returns the ProseData field if non-nil, zero value otherwise.

### GetProseDataOk

`func (o *UeSubscribedDataSets) GetProseDataOk() (*ProseSubscriptionData, bool)`

GetProseDataOk returns a tuple with the ProseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseData

`func (o *UeSubscribedDataSets) SetProseData(v ProseSubscriptionData)`

SetProseData sets ProseData field to given value.

### HasProseData

`func (o *UeSubscribedDataSets) HasProseData() bool`

HasProseData returns a boolean if a field has been set.

### GetOdbData

`func (o *UeSubscribedDataSets) GetOdbData() OdbData`

GetOdbData returns the OdbData field if non-nil, zero value otherwise.

### GetOdbDataOk

`func (o *UeSubscribedDataSets) GetOdbDataOk() (*OdbData, bool)`

GetOdbDataOk returns a tuple with the OdbData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdbData

`func (o *UeSubscribedDataSets) SetOdbData(v OdbData)`

SetOdbData sets OdbData field to given value.

### HasOdbData

`func (o *UeSubscribedDataSets) HasOdbData() bool`

HasOdbData returns a boolean if a field has been set.

### GetEeProfileData

`func (o *UeSubscribedDataSets) GetEeProfileData() EeProfileData`

GetEeProfileData returns the EeProfileData field if non-nil, zero value otherwise.

### GetEeProfileDataOk

`func (o *UeSubscribedDataSets) GetEeProfileDataOk() (*EeProfileData, bool)`

GetEeProfileDataOk returns a tuple with the EeProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEeProfileData

`func (o *UeSubscribedDataSets) SetEeProfileData(v EeProfileData)`

SetEeProfileData sets EeProfileData field to given value.

### HasEeProfileData

`func (o *UeSubscribedDataSets) HasEeProfileData() bool`

HasEeProfileData returns a boolean if a field has been set.

### GetPpProfileData

`func (o *UeSubscribedDataSets) GetPpProfileData() PpProfileData`

GetPpProfileData returns the PpProfileData field if non-nil, zero value otherwise.

### GetPpProfileDataOk

`func (o *UeSubscribedDataSets) GetPpProfileDataOk() (*PpProfileData, bool)`

GetPpProfileDataOk returns a tuple with the PpProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpProfileData

`func (o *UeSubscribedDataSets) SetPpProfileData(v PpProfileData)`

SetPpProfileData sets PpProfileData field to given value.

### HasPpProfileData

`func (o *UeSubscribedDataSets) HasPpProfileData() bool`

HasPpProfileData returns a boolean if a field has been set.

### GetNiddAuthData

`func (o *UeSubscribedDataSets) GetNiddAuthData() AuthorizationData`

GetNiddAuthData returns the NiddAuthData field if non-nil, zero value otherwise.

### GetNiddAuthDataOk

`func (o *UeSubscribedDataSets) GetNiddAuthDataOk() (*AuthorizationData, bool)`

GetNiddAuthDataOk returns a tuple with the NiddAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddAuthData

`func (o *UeSubscribedDataSets) SetNiddAuthData(v AuthorizationData)`

SetNiddAuthData sets NiddAuthData field to given value.

### HasNiddAuthData

`func (o *UeSubscribedDataSets) HasNiddAuthData() bool`

HasNiddAuthData returns a boolean if a field has been set.

### GetUcData

`func (o *UeSubscribedDataSets) GetUcData() UcSubscriptionData`

GetUcData returns the UcData field if non-nil, zero value otherwise.

### GetUcDataOk

`func (o *UeSubscribedDataSets) GetUcDataOk() (*UcSubscriptionData, bool)`

GetUcDataOk returns a tuple with the UcData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcData

`func (o *UeSubscribedDataSets) SetUcData(v UcSubscriptionData)`

SetUcData sets UcData field to given value.

### HasUcData

`func (o *UeSubscribedDataSets) HasUcData() bool`

HasUcData returns a boolean if a field has been set.

### GetMbsSubscriptionData

`func (o *UeSubscribedDataSets) GetMbsSubscriptionData() MbsSubscriptionData`

GetMbsSubscriptionData returns the MbsSubscriptionData field if non-nil, zero value otherwise.

### GetMbsSubscriptionDataOk

`func (o *UeSubscribedDataSets) GetMbsSubscriptionDataOk() (*MbsSubscriptionData, bool)`

GetMbsSubscriptionDataOk returns a tuple with the MbsSubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSubscriptionData

`func (o *UeSubscribedDataSets) SetMbsSubscriptionData(v MbsSubscriptionData)`

SetMbsSubscriptionData sets MbsSubscriptionData field to given value.

### HasMbsSubscriptionData

`func (o *UeSubscribedDataSets) HasMbsSubscriptionData() bool`

HasMbsSubscriptionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


