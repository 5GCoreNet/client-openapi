# EcsAddressProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**EcsServerAddr** | [**EcsServerAddr**](EcsServerAddr.md) |  | 
**SpatialValidityCond** | Pointer to [**SpatialValidityCond**](SpatialValidityCond.md) |  | [optional] 
**TgtUe** | Pointer to [**TargetUeId**](TargetUeId.md) |  | [optional] 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 

## Methods

### NewEcsAddressProvision

`func NewEcsAddressProvision(ecsServerAddr EcsServerAddr, suppFeat string, ) *EcsAddressProvision`

NewEcsAddressProvision instantiates a new EcsAddressProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcsAddressProvisionWithDefaults

`func NewEcsAddressProvisionWithDefaults() *EcsAddressProvision`

NewEcsAddressProvisionWithDefaults instantiates a new EcsAddressProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *EcsAddressProvision) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *EcsAddressProvision) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *EcsAddressProvision) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *EcsAddressProvision) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetEcsServerAddr

`func (o *EcsAddressProvision) GetEcsServerAddr() EcsServerAddr`

GetEcsServerAddr returns the EcsServerAddr field if non-nil, zero value otherwise.

### GetEcsServerAddrOk

`func (o *EcsAddressProvision) GetEcsServerAddrOk() (*EcsServerAddr, bool)`

GetEcsServerAddrOk returns a tuple with the EcsServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsServerAddr

`func (o *EcsAddressProvision) SetEcsServerAddr(v EcsServerAddr)`

SetEcsServerAddr sets EcsServerAddr field to given value.


### GetSpatialValidityCond

`func (o *EcsAddressProvision) GetSpatialValidityCond() SpatialValidityCond`

GetSpatialValidityCond returns the SpatialValidityCond field if non-nil, zero value otherwise.

### GetSpatialValidityCondOk

`func (o *EcsAddressProvision) GetSpatialValidityCondOk() (*SpatialValidityCond, bool)`

GetSpatialValidityCondOk returns a tuple with the SpatialValidityCond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidityCond

`func (o *EcsAddressProvision) SetSpatialValidityCond(v SpatialValidityCond)`

SetSpatialValidityCond sets SpatialValidityCond field to given value.

### HasSpatialValidityCond

`func (o *EcsAddressProvision) HasSpatialValidityCond() bool`

HasSpatialValidityCond returns a boolean if a field has been set.

### GetTgtUe

`func (o *EcsAddressProvision) GetTgtUe() TargetUeId`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *EcsAddressProvision) GetTgtUeOk() (*TargetUeId, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *EcsAddressProvision) SetTgtUe(v TargetUeId)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *EcsAddressProvision) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.

### GetSuppFeat

`func (o *EcsAddressProvision) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *EcsAddressProvision) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *EcsAddressProvision) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


