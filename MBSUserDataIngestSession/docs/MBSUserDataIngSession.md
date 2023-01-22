# MBSUserDataIngSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsUserServId** | **string** |  | 
**MbsDisSessInfos** | [**map[string]MBSDistributionSessionInfo**](MBSDistributionSessionInfo.md) | Represents one or more MBS Distribution Session(s) composing the MBS User Data Ingest  Session. The key of the map shall be set to the value ofthe \&quot;mbsDistSessionId\&quot; attribute  of the MBSDistributionSessionInfo data structure encoding the corresponding map entry.  | 
**ActPeriods** | Pointer to [**[]TimeWindow**](TimeWindow.md) |  | [optional] 
**MbsUserServAnmt** | Pointer to [**MBSUserServAnmt**](MBSUserServAnmt.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMBSUserDataIngSession

`func NewMBSUserDataIngSession(mbsUserServId string, mbsDisSessInfos map[string]MBSDistributionSessionInfo, ) *MBSUserDataIngSession`

NewMBSUserDataIngSession instantiates a new MBSUserDataIngSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserDataIngSessionWithDefaults

`func NewMBSUserDataIngSessionWithDefaults() *MBSUserDataIngSession`

NewMBSUserDataIngSessionWithDefaults instantiates a new MBSUserDataIngSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsUserServId

`func (o *MBSUserDataIngSession) GetMbsUserServId() string`

GetMbsUserServId returns the MbsUserServId field if non-nil, zero value otherwise.

### GetMbsUserServIdOk

`func (o *MBSUserDataIngSession) GetMbsUserServIdOk() (*string, bool)`

GetMbsUserServIdOk returns a tuple with the MbsUserServId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsUserServId

`func (o *MBSUserDataIngSession) SetMbsUserServId(v string)`

SetMbsUserServId sets MbsUserServId field to given value.


### GetMbsDisSessInfos

`func (o *MBSUserDataIngSession) GetMbsDisSessInfos() map[string]MBSDistributionSessionInfo`

GetMbsDisSessInfos returns the MbsDisSessInfos field if non-nil, zero value otherwise.

### GetMbsDisSessInfosOk

`func (o *MBSUserDataIngSession) GetMbsDisSessInfosOk() (*map[string]MBSDistributionSessionInfo, bool)`

GetMbsDisSessInfosOk returns a tuple with the MbsDisSessInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDisSessInfos

`func (o *MBSUserDataIngSession) SetMbsDisSessInfos(v map[string]MBSDistributionSessionInfo)`

SetMbsDisSessInfos sets MbsDisSessInfos field to given value.


### GetActPeriods

`func (o *MBSUserDataIngSession) GetActPeriods() []TimeWindow`

GetActPeriods returns the ActPeriods field if non-nil, zero value otherwise.

### GetActPeriodsOk

`func (o *MBSUserDataIngSession) GetActPeriodsOk() (*[]TimeWindow, bool)`

GetActPeriodsOk returns a tuple with the ActPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActPeriods

`func (o *MBSUserDataIngSession) SetActPeriods(v []TimeWindow)`

SetActPeriods sets ActPeriods field to given value.

### HasActPeriods

`func (o *MBSUserDataIngSession) HasActPeriods() bool`

HasActPeriods returns a boolean if a field has been set.

### GetMbsUserServAnmt

`func (o *MBSUserDataIngSession) GetMbsUserServAnmt() MBSUserServAnmt`

GetMbsUserServAnmt returns the MbsUserServAnmt field if non-nil, zero value otherwise.

### GetMbsUserServAnmtOk

`func (o *MBSUserDataIngSession) GetMbsUserServAnmtOk() (*MBSUserServAnmt, bool)`

GetMbsUserServAnmtOk returns a tuple with the MbsUserServAnmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsUserServAnmt

`func (o *MBSUserDataIngSession) SetMbsUserServAnmt(v MBSUserServAnmt)`

SetMbsUserServAnmt sets MbsUserServAnmt field to given value.

### HasMbsUserServAnmt

`func (o *MBSUserDataIngSession) HasMbsUserServAnmt() bool`

HasMbsUserServAnmt returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MBSUserDataIngSession) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MBSUserDataIngSession) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MBSUserDataIngSession) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MBSUserDataIngSession) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


