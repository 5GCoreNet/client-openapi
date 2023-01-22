# MBSUserDataIngSessionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsDistSessInfos** | Pointer to [**map[string]MBSDistributionSessionInfo**](MBSDistributionSessionInfo.md) | Contains the requested modifications to one or more MBS Distribution Session(s)  composing the MBS User Data Ingest Session.  | [optional] 
**ActPeriods** | Pointer to [**[]TimeWindow**](TimeWindow.md) |  | [optional] 

## Methods

### NewMBSUserDataIngSessionPatch

`func NewMBSUserDataIngSessionPatch() *MBSUserDataIngSessionPatch`

NewMBSUserDataIngSessionPatch instantiates a new MBSUserDataIngSessionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserDataIngSessionPatchWithDefaults

`func NewMBSUserDataIngSessionPatchWithDefaults() *MBSUserDataIngSessionPatch`

NewMBSUserDataIngSessionPatchWithDefaults instantiates a new MBSUserDataIngSessionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsDistSessInfos

`func (o *MBSUserDataIngSessionPatch) GetMbsDistSessInfos() map[string]MBSDistributionSessionInfo`

GetMbsDistSessInfos returns the MbsDistSessInfos field if non-nil, zero value otherwise.

### GetMbsDistSessInfosOk

`func (o *MBSUserDataIngSessionPatch) GetMbsDistSessInfosOk() (*map[string]MBSDistributionSessionInfo, bool)`

GetMbsDistSessInfosOk returns a tuple with the MbsDistSessInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDistSessInfos

`func (o *MBSUserDataIngSessionPatch) SetMbsDistSessInfos(v map[string]MBSDistributionSessionInfo)`

SetMbsDistSessInfos sets MbsDistSessInfos field to given value.

### HasMbsDistSessInfos

`func (o *MBSUserDataIngSessionPatch) HasMbsDistSessInfos() bool`

HasMbsDistSessInfos returns a boolean if a field has been set.

### GetActPeriods

`func (o *MBSUserDataIngSessionPatch) GetActPeriods() []TimeWindow`

GetActPeriods returns the ActPeriods field if non-nil, zero value otherwise.

### GetActPeriodsOk

`func (o *MBSUserDataIngSessionPatch) GetActPeriodsOk() (*[]TimeWindow, bool)`

GetActPeriodsOk returns a tuple with the ActPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActPeriods

`func (o *MBSUserDataIngSessionPatch) SetActPeriods(v []TimeWindow)`

SetActPeriods sets ActPeriods field to given value.

### HasActPeriods

`func (o *MBSUserDataIngSessionPatch) HasActPeriods() bool`

HasActPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


