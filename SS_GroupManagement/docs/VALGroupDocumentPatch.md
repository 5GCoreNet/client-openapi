# VALGroupDocumentPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrpDesc** | Pointer to **string** | The text description of the VAL group. | [optional] 
**Members** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | The list of VAL User IDs or VAL UE IDs, which are members of the VAL group. | [optional] 
**ValGrpConf** | Pointer to **string** | Configuration data for the VAL group. | [optional] 
**ValServiceIds** | Pointer to **[]string** | The list of VAL services enabled on the group. | [optional] 
**LocInfo** | Pointer to [**LocationInfo**](LocationInfo.md) |  | [optional] 
**AddLocInfo** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**ExtGrpId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**Com5GLanType** | Pointer to [**PduSessionType**](PduSessionType.md) |  | [optional] 

## Methods

### NewVALGroupDocumentPatch

`func NewVALGroupDocumentPatch() *VALGroupDocumentPatch`

NewVALGroupDocumentPatch instantiates a new VALGroupDocumentPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVALGroupDocumentPatchWithDefaults

`func NewVALGroupDocumentPatchWithDefaults() *VALGroupDocumentPatch`

NewVALGroupDocumentPatchWithDefaults instantiates a new VALGroupDocumentPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrpDesc

`func (o *VALGroupDocumentPatch) GetGrpDesc() string`

GetGrpDesc returns the GrpDesc field if non-nil, zero value otherwise.

### GetGrpDescOk

`func (o *VALGroupDocumentPatch) GetGrpDescOk() (*string, bool)`

GetGrpDescOk returns a tuple with the GrpDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpDesc

`func (o *VALGroupDocumentPatch) SetGrpDesc(v string)`

SetGrpDesc sets GrpDesc field to given value.

### HasGrpDesc

`func (o *VALGroupDocumentPatch) HasGrpDesc() bool`

HasGrpDesc returns a boolean if a field has been set.

### GetMembers

`func (o *VALGroupDocumentPatch) GetMembers() []ValTargetUe`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *VALGroupDocumentPatch) GetMembersOk() (*[]ValTargetUe, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *VALGroupDocumentPatch) SetMembers(v []ValTargetUe)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *VALGroupDocumentPatch) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetValGrpConf

`func (o *VALGroupDocumentPatch) GetValGrpConf() string`

GetValGrpConf returns the ValGrpConf field if non-nil, zero value otherwise.

### GetValGrpConfOk

`func (o *VALGroupDocumentPatch) GetValGrpConfOk() (*string, bool)`

GetValGrpConfOk returns a tuple with the ValGrpConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGrpConf

`func (o *VALGroupDocumentPatch) SetValGrpConf(v string)`

SetValGrpConf sets ValGrpConf field to given value.

### HasValGrpConf

`func (o *VALGroupDocumentPatch) HasValGrpConf() bool`

HasValGrpConf returns a boolean if a field has been set.

### GetValServiceIds

`func (o *VALGroupDocumentPatch) GetValServiceIds() []string`

GetValServiceIds returns the ValServiceIds field if non-nil, zero value otherwise.

### GetValServiceIdsOk

`func (o *VALGroupDocumentPatch) GetValServiceIdsOk() (*[]string, bool)`

GetValServiceIdsOk returns a tuple with the ValServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValServiceIds

`func (o *VALGroupDocumentPatch) SetValServiceIds(v []string)`

SetValServiceIds sets ValServiceIds field to given value.

### HasValServiceIds

`func (o *VALGroupDocumentPatch) HasValServiceIds() bool`

HasValServiceIds returns a boolean if a field has been set.

### GetLocInfo

`func (o *VALGroupDocumentPatch) GetLocInfo() LocationInfo`

GetLocInfo returns the LocInfo field if non-nil, zero value otherwise.

### GetLocInfoOk

`func (o *VALGroupDocumentPatch) GetLocInfoOk() (*LocationInfo, bool)`

GetLocInfoOk returns a tuple with the LocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfo

`func (o *VALGroupDocumentPatch) SetLocInfo(v LocationInfo)`

SetLocInfo sets LocInfo field to given value.

### HasLocInfo

`func (o *VALGroupDocumentPatch) HasLocInfo() bool`

HasLocInfo returns a boolean if a field has been set.

### GetAddLocInfo

`func (o *VALGroupDocumentPatch) GetAddLocInfo() LocationArea5G`

GetAddLocInfo returns the AddLocInfo field if non-nil, zero value otherwise.

### GetAddLocInfoOk

`func (o *VALGroupDocumentPatch) GetAddLocInfoOk() (*LocationArea5G, bool)`

GetAddLocInfoOk returns a tuple with the AddLocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLocInfo

`func (o *VALGroupDocumentPatch) SetAddLocInfo(v LocationArea5G)`

SetAddLocInfo sets AddLocInfo field to given value.

### HasAddLocInfo

`func (o *VALGroupDocumentPatch) HasAddLocInfo() bool`

HasAddLocInfo returns a boolean if a field has been set.

### GetExtGrpId

`func (o *VALGroupDocumentPatch) GetExtGrpId() string`

GetExtGrpId returns the ExtGrpId field if non-nil, zero value otherwise.

### GetExtGrpIdOk

`func (o *VALGroupDocumentPatch) GetExtGrpIdOk() (*string, bool)`

GetExtGrpIdOk returns a tuple with the ExtGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGrpId

`func (o *VALGroupDocumentPatch) SetExtGrpId(v string)`

SetExtGrpId sets ExtGrpId field to given value.

### HasExtGrpId

`func (o *VALGroupDocumentPatch) HasExtGrpId() bool`

HasExtGrpId returns a boolean if a field has been set.

### GetCom5GLanType

`func (o *VALGroupDocumentPatch) GetCom5GLanType() PduSessionType`

GetCom5GLanType returns the Com5GLanType field if non-nil, zero value otherwise.

### GetCom5GLanTypeOk

`func (o *VALGroupDocumentPatch) GetCom5GLanTypeOk() (*PduSessionType, bool)`

GetCom5GLanTypeOk returns a tuple with the Com5GLanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCom5GLanType

`func (o *VALGroupDocumentPatch) SetCom5GLanType(v PduSessionType)`

SetCom5GLanType sets Com5GLanType field to given value.

### HasCom5GLanType

`func (o *VALGroupDocumentPatch) HasCom5GLanType() bool`

HasCom5GLanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


