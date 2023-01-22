# MBSDistSessionAnmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | Pointer to [**MbsSessionId**](MbsSessionId.md) |  | [optional] 
**MbsFSAId** | Pointer to **string** | MBS Frequency Selection Area Identifier | [optional] 
**DistrMethod** | [**DistributionMethod**](DistributionMethod.md) |  | 
**ObjDistrAnnInfo** | Pointer to [**ObjectDistMethAnmtInfo**](ObjectDistMethAnmtInfo.md) |  | [optional] 
**SesDesInfo** | **[]string** |  | 

## Methods

### NewMBSDistSessionAnmt

`func NewMBSDistSessionAnmt(distrMethod DistributionMethod, sesDesInfo []string, ) *MBSDistSessionAnmt`

NewMBSDistSessionAnmt instantiates a new MBSDistSessionAnmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSDistSessionAnmtWithDefaults

`func NewMBSDistSessionAnmtWithDefaults() *MBSDistSessionAnmt`

NewMBSDistSessionAnmtWithDefaults instantiates a new MBSDistSessionAnmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *MBSDistSessionAnmt) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MBSDistSessionAnmt) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MBSDistSessionAnmt) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.

### HasMbsSessionId

`func (o *MBSDistSessionAnmt) HasMbsSessionId() bool`

HasMbsSessionId returns a boolean if a field has been set.

### GetMbsFSAId

`func (o *MBSDistSessionAnmt) GetMbsFSAId() string`

GetMbsFSAId returns the MbsFSAId field if non-nil, zero value otherwise.

### GetMbsFSAIdOk

`func (o *MBSDistSessionAnmt) GetMbsFSAIdOk() (*string, bool)`

GetMbsFSAIdOk returns a tuple with the MbsFSAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsFSAId

`func (o *MBSDistSessionAnmt) SetMbsFSAId(v string)`

SetMbsFSAId sets MbsFSAId field to given value.

### HasMbsFSAId

`func (o *MBSDistSessionAnmt) HasMbsFSAId() bool`

HasMbsFSAId returns a boolean if a field has been set.

### GetDistrMethod

`func (o *MBSDistSessionAnmt) GetDistrMethod() DistributionMethod`

GetDistrMethod returns the DistrMethod field if non-nil, zero value otherwise.

### GetDistrMethodOk

`func (o *MBSDistSessionAnmt) GetDistrMethodOk() (*DistributionMethod, bool)`

GetDistrMethodOk returns a tuple with the DistrMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrMethod

`func (o *MBSDistSessionAnmt) SetDistrMethod(v DistributionMethod)`

SetDistrMethod sets DistrMethod field to given value.


### GetObjDistrAnnInfo

`func (o *MBSDistSessionAnmt) GetObjDistrAnnInfo() ObjectDistMethAnmtInfo`

GetObjDistrAnnInfo returns the ObjDistrAnnInfo field if non-nil, zero value otherwise.

### GetObjDistrAnnInfoOk

`func (o *MBSDistSessionAnmt) GetObjDistrAnnInfoOk() (*ObjectDistMethAnmtInfo, bool)`

GetObjDistrAnnInfoOk returns a tuple with the ObjDistrAnnInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDistrAnnInfo

`func (o *MBSDistSessionAnmt) SetObjDistrAnnInfo(v ObjectDistMethAnmtInfo)`

SetObjDistrAnnInfo sets ObjDistrAnnInfo field to given value.

### HasObjDistrAnnInfo

`func (o *MBSDistSessionAnmt) HasObjDistrAnnInfo() bool`

HasObjDistrAnnInfo returns a boolean if a field has been set.

### GetSesDesInfo

`func (o *MBSDistSessionAnmt) GetSesDesInfo() []string`

GetSesDesInfo returns the SesDesInfo field if non-nil, zero value otherwise.

### GetSesDesInfoOk

`func (o *MBSDistSessionAnmt) GetSesDesInfoOk() (*[]string, bool)`

GetSesDesInfoOk returns a tuple with the SesDesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSesDesInfo

`func (o *MBSDistSessionAnmt) SetSesDesInfo(v []string)`

SetSesDesInfo sets SesDesInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


