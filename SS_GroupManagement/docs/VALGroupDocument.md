# VALGroupDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValGroupId** | **string** | The VAL group idenitity. | 
**GrpDesc** | Pointer to **string** | The text description of the VAL group. | [optional] 
**Members** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | The list of VAL User IDs or VAL UE IDs, which are members of the VAL group. | [optional] 
**ValGrpConf** | Pointer to **string** | Configuration data for the VAL group. | [optional] 
**ValServiceIds** | Pointer to **[]string** | The list of VAL services enabled on the group. | [optional] 
**ValSvcInf** | Pointer to **string** | VAL service specific information. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResUri** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**LocInfo** | Pointer to [**LocationInfo**](LocationInfo.md) |  | [optional] 
**AddLocInfo** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**ExtGrpId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**Com5GLanType** | Pointer to [**PduSessionType**](PduSessionType.md) |  | [optional] 

## Methods

### NewVALGroupDocument

`func NewVALGroupDocument(valGroupId string, ) *VALGroupDocument`

NewVALGroupDocument instantiates a new VALGroupDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVALGroupDocumentWithDefaults

`func NewVALGroupDocumentWithDefaults() *VALGroupDocument`

NewVALGroupDocumentWithDefaults instantiates a new VALGroupDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValGroupId

`func (o *VALGroupDocument) GetValGroupId() string`

GetValGroupId returns the ValGroupId field if non-nil, zero value otherwise.

### GetValGroupIdOk

`func (o *VALGroupDocument) GetValGroupIdOk() (*string, bool)`

GetValGroupIdOk returns a tuple with the ValGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGroupId

`func (o *VALGroupDocument) SetValGroupId(v string)`

SetValGroupId sets ValGroupId field to given value.


### GetGrpDesc

`func (o *VALGroupDocument) GetGrpDesc() string`

GetGrpDesc returns the GrpDesc field if non-nil, zero value otherwise.

### GetGrpDescOk

`func (o *VALGroupDocument) GetGrpDescOk() (*string, bool)`

GetGrpDescOk returns a tuple with the GrpDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpDesc

`func (o *VALGroupDocument) SetGrpDesc(v string)`

SetGrpDesc sets GrpDesc field to given value.

### HasGrpDesc

`func (o *VALGroupDocument) HasGrpDesc() bool`

HasGrpDesc returns a boolean if a field has been set.

### GetMembers

`func (o *VALGroupDocument) GetMembers() []ValTargetUe`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *VALGroupDocument) GetMembersOk() (*[]ValTargetUe, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *VALGroupDocument) SetMembers(v []ValTargetUe)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *VALGroupDocument) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetValGrpConf

`func (o *VALGroupDocument) GetValGrpConf() string`

GetValGrpConf returns the ValGrpConf field if non-nil, zero value otherwise.

### GetValGrpConfOk

`func (o *VALGroupDocument) GetValGrpConfOk() (*string, bool)`

GetValGrpConfOk returns a tuple with the ValGrpConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGrpConf

`func (o *VALGroupDocument) SetValGrpConf(v string)`

SetValGrpConf sets ValGrpConf field to given value.

### HasValGrpConf

`func (o *VALGroupDocument) HasValGrpConf() bool`

HasValGrpConf returns a boolean if a field has been set.

### GetValServiceIds

`func (o *VALGroupDocument) GetValServiceIds() []string`

GetValServiceIds returns the ValServiceIds field if non-nil, zero value otherwise.

### GetValServiceIdsOk

`func (o *VALGroupDocument) GetValServiceIdsOk() (*[]string, bool)`

GetValServiceIdsOk returns a tuple with the ValServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValServiceIds

`func (o *VALGroupDocument) SetValServiceIds(v []string)`

SetValServiceIds sets ValServiceIds field to given value.

### HasValServiceIds

`func (o *VALGroupDocument) HasValServiceIds() bool`

HasValServiceIds returns a boolean if a field has been set.

### GetValSvcInf

`func (o *VALGroupDocument) GetValSvcInf() string`

GetValSvcInf returns the ValSvcInf field if non-nil, zero value otherwise.

### GetValSvcInfOk

`func (o *VALGroupDocument) GetValSvcInfOk() (*string, bool)`

GetValSvcInfOk returns a tuple with the ValSvcInf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValSvcInf

`func (o *VALGroupDocument) SetValSvcInf(v string)`

SetValSvcInf sets ValSvcInf field to given value.

### HasValSvcInf

`func (o *VALGroupDocument) HasValSvcInf() bool`

HasValSvcInf returns a boolean if a field has been set.

### GetSuppFeat

`func (o *VALGroupDocument) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *VALGroupDocument) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *VALGroupDocument) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *VALGroupDocument) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResUri

`func (o *VALGroupDocument) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *VALGroupDocument) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *VALGroupDocument) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *VALGroupDocument) HasResUri() bool`

HasResUri returns a boolean if a field has been set.

### GetLocInfo

`func (o *VALGroupDocument) GetLocInfo() LocationInfo`

GetLocInfo returns the LocInfo field if non-nil, zero value otherwise.

### GetLocInfoOk

`func (o *VALGroupDocument) GetLocInfoOk() (*LocationInfo, bool)`

GetLocInfoOk returns a tuple with the LocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfo

`func (o *VALGroupDocument) SetLocInfo(v LocationInfo)`

SetLocInfo sets LocInfo field to given value.

### HasLocInfo

`func (o *VALGroupDocument) HasLocInfo() bool`

HasLocInfo returns a boolean if a field has been set.

### GetAddLocInfo

`func (o *VALGroupDocument) GetAddLocInfo() LocationArea5G`

GetAddLocInfo returns the AddLocInfo field if non-nil, zero value otherwise.

### GetAddLocInfoOk

`func (o *VALGroupDocument) GetAddLocInfoOk() (*LocationArea5G, bool)`

GetAddLocInfoOk returns a tuple with the AddLocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLocInfo

`func (o *VALGroupDocument) SetAddLocInfo(v LocationArea5G)`

SetAddLocInfo sets AddLocInfo field to given value.

### HasAddLocInfo

`func (o *VALGroupDocument) HasAddLocInfo() bool`

HasAddLocInfo returns a boolean if a field has been set.

### GetExtGrpId

`func (o *VALGroupDocument) GetExtGrpId() string`

GetExtGrpId returns the ExtGrpId field if non-nil, zero value otherwise.

### GetExtGrpIdOk

`func (o *VALGroupDocument) GetExtGrpIdOk() (*string, bool)`

GetExtGrpIdOk returns a tuple with the ExtGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGrpId

`func (o *VALGroupDocument) SetExtGrpId(v string)`

SetExtGrpId sets ExtGrpId field to given value.

### HasExtGrpId

`func (o *VALGroupDocument) HasExtGrpId() bool`

HasExtGrpId returns a boolean if a field has been set.

### GetCom5GLanType

`func (o *VALGroupDocument) GetCom5GLanType() PduSessionType`

GetCom5GLanType returns the Com5GLanType field if non-nil, zero value otherwise.

### GetCom5GLanTypeOk

`func (o *VALGroupDocument) GetCom5GLanTypeOk() (*PduSessionType, bool)`

GetCom5GLanTypeOk returns a tuple with the Com5GLanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCom5GLanType

`func (o *VALGroupDocument) SetCom5GLanType(v PduSessionType)`

SetCom5GLanType sets Com5GLanType field to given value.

### HasCom5GLanType

`func (o *VALGroupDocument) HasCom5GLanType() bool`

HasCom5GLanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


