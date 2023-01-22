# StatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusInfo** | [**StatusInfo**](StatusInfo.md) |  | 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 
**TargetDnaiInfo** | Pointer to [**TargetDnaiInfo**](TargetDnaiInfo.md) |  | [optional] 
**OldPduSessionRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NewSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**EpsPdnCnxInfo** | Pointer to [**EpsPdnCnxInfo**](EpsPdnCnxInfo.md) |  | [optional] 
**InterPlmnApiRoot** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**IntraPlmnApiRoot** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewStatusNotification

`func NewStatusNotification(statusInfo StatusInfo, ) *StatusNotification`

NewStatusNotification instantiates a new StatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusNotificationWithDefaults

`func NewStatusNotificationWithDefaults() *StatusNotification`

NewStatusNotificationWithDefaults instantiates a new StatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusInfo

`func (o *StatusNotification) GetStatusInfo() StatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *StatusNotification) GetStatusInfoOk() (*StatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *StatusNotification) SetStatusInfo(v StatusInfo)`

SetStatusInfo sets StatusInfo field to given value.


### GetSmallDataRateStatus

`func (o *StatusNotification) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *StatusNotification) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *StatusNotification) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *StatusNotification) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *StatusNotification) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *StatusNotification) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *StatusNotification) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *StatusNotification) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.

### GetTargetDnaiInfo

`func (o *StatusNotification) GetTargetDnaiInfo() TargetDnaiInfo`

GetTargetDnaiInfo returns the TargetDnaiInfo field if non-nil, zero value otherwise.

### GetTargetDnaiInfoOk

`func (o *StatusNotification) GetTargetDnaiInfoOk() (*TargetDnaiInfo, bool)`

GetTargetDnaiInfoOk returns a tuple with the TargetDnaiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnaiInfo

`func (o *StatusNotification) SetTargetDnaiInfo(v TargetDnaiInfo)`

SetTargetDnaiInfo sets TargetDnaiInfo field to given value.

### HasTargetDnaiInfo

`func (o *StatusNotification) HasTargetDnaiInfo() bool`

HasTargetDnaiInfo returns a boolean if a field has been set.

### GetOldPduSessionRef

`func (o *StatusNotification) GetOldPduSessionRef() string`

GetOldPduSessionRef returns the OldPduSessionRef field if non-nil, zero value otherwise.

### GetOldPduSessionRefOk

`func (o *StatusNotification) GetOldPduSessionRefOk() (*string, bool)`

GetOldPduSessionRefOk returns a tuple with the OldPduSessionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPduSessionRef

`func (o *StatusNotification) SetOldPduSessionRef(v string)`

SetOldPduSessionRef sets OldPduSessionRef field to given value.

### HasOldPduSessionRef

`func (o *StatusNotification) HasOldPduSessionRef() bool`

HasOldPduSessionRef returns a boolean if a field has been set.

### GetNewSmfId

`func (o *StatusNotification) GetNewSmfId() string`

GetNewSmfId returns the NewSmfId field if non-nil, zero value otherwise.

### GetNewSmfIdOk

`func (o *StatusNotification) GetNewSmfIdOk() (*string, bool)`

GetNewSmfIdOk returns a tuple with the NewSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSmfId

`func (o *StatusNotification) SetNewSmfId(v string)`

SetNewSmfId sets NewSmfId field to given value.

### HasNewSmfId

`func (o *StatusNotification) HasNewSmfId() bool`

HasNewSmfId returns a boolean if a field has been set.

### GetEpsPdnCnxInfo

`func (o *StatusNotification) GetEpsPdnCnxInfo() EpsPdnCnxInfo`

GetEpsPdnCnxInfo returns the EpsPdnCnxInfo field if non-nil, zero value otherwise.

### GetEpsPdnCnxInfoOk

`func (o *StatusNotification) GetEpsPdnCnxInfoOk() (*EpsPdnCnxInfo, bool)`

GetEpsPdnCnxInfoOk returns a tuple with the EpsPdnCnxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsPdnCnxInfo

`func (o *StatusNotification) SetEpsPdnCnxInfo(v EpsPdnCnxInfo)`

SetEpsPdnCnxInfo sets EpsPdnCnxInfo field to given value.

### HasEpsPdnCnxInfo

`func (o *StatusNotification) HasEpsPdnCnxInfo() bool`

HasEpsPdnCnxInfo returns a boolean if a field has been set.

### GetInterPlmnApiRoot

`func (o *StatusNotification) GetInterPlmnApiRoot() string`

GetInterPlmnApiRoot returns the InterPlmnApiRoot field if non-nil, zero value otherwise.

### GetInterPlmnApiRootOk

`func (o *StatusNotification) GetInterPlmnApiRootOk() (*string, bool)`

GetInterPlmnApiRootOk returns a tuple with the InterPlmnApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnApiRoot

`func (o *StatusNotification) SetInterPlmnApiRoot(v string)`

SetInterPlmnApiRoot sets InterPlmnApiRoot field to given value.

### HasInterPlmnApiRoot

`func (o *StatusNotification) HasInterPlmnApiRoot() bool`

HasInterPlmnApiRoot returns a boolean if a field has been set.

### GetIntraPlmnApiRoot

`func (o *StatusNotification) GetIntraPlmnApiRoot() string`

GetIntraPlmnApiRoot returns the IntraPlmnApiRoot field if non-nil, zero value otherwise.

### GetIntraPlmnApiRootOk

`func (o *StatusNotification) GetIntraPlmnApiRootOk() (*string, bool)`

GetIntraPlmnApiRootOk returns a tuple with the IntraPlmnApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraPlmnApiRoot

`func (o *StatusNotification) SetIntraPlmnApiRoot(v string)`

SetIntraPlmnApiRoot sets IntraPlmnApiRoot field to given value.

### HasIntraPlmnApiRoot

`func (o *StatusNotification) HasIntraPlmnApiRoot() bool`

HasIntraPlmnApiRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


