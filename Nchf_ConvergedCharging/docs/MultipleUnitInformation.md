# MultipleUnitInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to [**ResultCode**](ResultCode.md) |  | [optional] 
**RatingGroup** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**GrantedUnit** | Pointer to [**GrantedUnit**](GrantedUnit.md) |  | [optional] 
**Triggers** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**ValidityTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**QuotaHoldingTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**FinalUnitIndication** | Pointer to [**FinalUnitIndication**](FinalUnitIndication.md) |  | [optional] 
**TimeQuotaThreshold** | Pointer to **int32** |  | [optional] 
**VolumeQuotaThreshold** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**UnitQuotaThreshold** | Pointer to **int32** |  | [optional] 
**UPFID** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**AnnouncementInformation** | Pointer to [**AnnouncementInformation**](AnnouncementInformation.md) |  | [optional] 

## Methods

### NewMultipleUnitInformation

`func NewMultipleUnitInformation(ratingGroup int32, ) *MultipleUnitInformation`

NewMultipleUnitInformation instantiates a new MultipleUnitInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleUnitInformationWithDefaults

`func NewMultipleUnitInformationWithDefaults() *MultipleUnitInformation`

NewMultipleUnitInformationWithDefaults instantiates a new MultipleUnitInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *MultipleUnitInformation) GetResultCode() ResultCode`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *MultipleUnitInformation) GetResultCodeOk() (*ResultCode, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *MultipleUnitInformation) SetResultCode(v ResultCode)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *MultipleUnitInformation) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetRatingGroup

`func (o *MultipleUnitInformation) GetRatingGroup() int32`

GetRatingGroup returns the RatingGroup field if non-nil, zero value otherwise.

### GetRatingGroupOk

`func (o *MultipleUnitInformation) GetRatingGroupOk() (*int32, bool)`

GetRatingGroupOk returns a tuple with the RatingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingGroup

`func (o *MultipleUnitInformation) SetRatingGroup(v int32)`

SetRatingGroup sets RatingGroup field to given value.


### GetGrantedUnit

`func (o *MultipleUnitInformation) GetGrantedUnit() GrantedUnit`

GetGrantedUnit returns the GrantedUnit field if non-nil, zero value otherwise.

### GetGrantedUnitOk

`func (o *MultipleUnitInformation) GetGrantedUnitOk() (*GrantedUnit, bool)`

GetGrantedUnitOk returns a tuple with the GrantedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedUnit

`func (o *MultipleUnitInformation) SetGrantedUnit(v GrantedUnit)`

SetGrantedUnit sets GrantedUnit field to given value.

### HasGrantedUnit

`func (o *MultipleUnitInformation) HasGrantedUnit() bool`

HasGrantedUnit returns a boolean if a field has been set.

### GetTriggers

`func (o *MultipleUnitInformation) GetTriggers() []Trigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MultipleUnitInformation) GetTriggersOk() (*[]Trigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MultipleUnitInformation) SetTriggers(v []Trigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *MultipleUnitInformation) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetValidityTime

`func (o *MultipleUnitInformation) GetValidityTime() int32`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *MultipleUnitInformation) GetValidityTimeOk() (*int32, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *MultipleUnitInformation) SetValidityTime(v int32)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *MultipleUnitInformation) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetQuotaHoldingTime

`func (o *MultipleUnitInformation) GetQuotaHoldingTime() int32`

GetQuotaHoldingTime returns the QuotaHoldingTime field if non-nil, zero value otherwise.

### GetQuotaHoldingTimeOk

`func (o *MultipleUnitInformation) GetQuotaHoldingTimeOk() (*int32, bool)`

GetQuotaHoldingTimeOk returns a tuple with the QuotaHoldingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaHoldingTime

`func (o *MultipleUnitInformation) SetQuotaHoldingTime(v int32)`

SetQuotaHoldingTime sets QuotaHoldingTime field to given value.

### HasQuotaHoldingTime

`func (o *MultipleUnitInformation) HasQuotaHoldingTime() bool`

HasQuotaHoldingTime returns a boolean if a field has been set.

### GetFinalUnitIndication

`func (o *MultipleUnitInformation) GetFinalUnitIndication() FinalUnitIndication`

GetFinalUnitIndication returns the FinalUnitIndication field if non-nil, zero value otherwise.

### GetFinalUnitIndicationOk

`func (o *MultipleUnitInformation) GetFinalUnitIndicationOk() (*FinalUnitIndication, bool)`

GetFinalUnitIndicationOk returns a tuple with the FinalUnitIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalUnitIndication

`func (o *MultipleUnitInformation) SetFinalUnitIndication(v FinalUnitIndication)`

SetFinalUnitIndication sets FinalUnitIndication field to given value.

### HasFinalUnitIndication

`func (o *MultipleUnitInformation) HasFinalUnitIndication() bool`

HasFinalUnitIndication returns a boolean if a field has been set.

### GetTimeQuotaThreshold

`func (o *MultipleUnitInformation) GetTimeQuotaThreshold() int32`

GetTimeQuotaThreshold returns the TimeQuotaThreshold field if non-nil, zero value otherwise.

### GetTimeQuotaThresholdOk

`func (o *MultipleUnitInformation) GetTimeQuotaThresholdOk() (*int32, bool)`

GetTimeQuotaThresholdOk returns a tuple with the TimeQuotaThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeQuotaThreshold

`func (o *MultipleUnitInformation) SetTimeQuotaThreshold(v int32)`

SetTimeQuotaThreshold sets TimeQuotaThreshold field to given value.

### HasTimeQuotaThreshold

`func (o *MultipleUnitInformation) HasTimeQuotaThreshold() bool`

HasTimeQuotaThreshold returns a boolean if a field has been set.

### GetVolumeQuotaThreshold

`func (o *MultipleUnitInformation) GetVolumeQuotaThreshold() int32`

GetVolumeQuotaThreshold returns the VolumeQuotaThreshold field if non-nil, zero value otherwise.

### GetVolumeQuotaThresholdOk

`func (o *MultipleUnitInformation) GetVolumeQuotaThresholdOk() (*int32, bool)`

GetVolumeQuotaThresholdOk returns a tuple with the VolumeQuotaThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeQuotaThreshold

`func (o *MultipleUnitInformation) SetVolumeQuotaThreshold(v int32)`

SetVolumeQuotaThreshold sets VolumeQuotaThreshold field to given value.

### HasVolumeQuotaThreshold

`func (o *MultipleUnitInformation) HasVolumeQuotaThreshold() bool`

HasVolumeQuotaThreshold returns a boolean if a field has been set.

### GetUnitQuotaThreshold

`func (o *MultipleUnitInformation) GetUnitQuotaThreshold() int32`

GetUnitQuotaThreshold returns the UnitQuotaThreshold field if non-nil, zero value otherwise.

### GetUnitQuotaThresholdOk

`func (o *MultipleUnitInformation) GetUnitQuotaThresholdOk() (*int32, bool)`

GetUnitQuotaThresholdOk returns a tuple with the UnitQuotaThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitQuotaThreshold

`func (o *MultipleUnitInformation) SetUnitQuotaThreshold(v int32)`

SetUnitQuotaThreshold sets UnitQuotaThreshold field to given value.

### HasUnitQuotaThreshold

`func (o *MultipleUnitInformation) HasUnitQuotaThreshold() bool`

HasUnitQuotaThreshold returns a boolean if a field has been set.

### GetUPFID

`func (o *MultipleUnitInformation) GetUPFID() string`

GetUPFID returns the UPFID field if non-nil, zero value otherwise.

### GetUPFIDOk

`func (o *MultipleUnitInformation) GetUPFIDOk() (*string, bool)`

GetUPFIDOk returns a tuple with the UPFID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPFID

`func (o *MultipleUnitInformation) SetUPFID(v string)`

SetUPFID sets UPFID field to given value.

### HasUPFID

`func (o *MultipleUnitInformation) HasUPFID() bool`

HasUPFID returns a boolean if a field has been set.

### GetAnnouncementInformation

`func (o *MultipleUnitInformation) GetAnnouncementInformation() AnnouncementInformation`

GetAnnouncementInformation returns the AnnouncementInformation field if non-nil, zero value otherwise.

### GetAnnouncementInformationOk

`func (o *MultipleUnitInformation) GetAnnouncementInformationOk() (*AnnouncementInformation, bool)`

GetAnnouncementInformationOk returns a tuple with the AnnouncementInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementInformation

`func (o *MultipleUnitInformation) SetAnnouncementInformation(v AnnouncementInformation)`

SetAnnouncementInformation sets AnnouncementInformation field to given value.

### HasAnnouncementInformation

`func (o *MultipleUnitInformation) HasAnnouncementInformation() bool`

HasAnnouncementInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


