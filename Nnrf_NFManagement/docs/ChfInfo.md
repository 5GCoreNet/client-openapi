# ChfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupiRangeList** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRangeList** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**PlmnRangeList** | Pointer to [**[]PlmnRange**](PlmnRange.md) |  | [optional] 
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**PrimaryChfInstance** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SecondaryChfInstance** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewChfInfo

`func NewChfInfo() *ChfInfo`

NewChfInfo instantiates a new ChfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChfInfoWithDefaults

`func NewChfInfoWithDefaults() *ChfInfo`

NewChfInfoWithDefaults instantiates a new ChfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupiRangeList

`func (o *ChfInfo) GetSupiRangeList() []SupiRange`

GetSupiRangeList returns the SupiRangeList field if non-nil, zero value otherwise.

### GetSupiRangeListOk

`func (o *ChfInfo) GetSupiRangeListOk() (*[]SupiRange, bool)`

GetSupiRangeListOk returns a tuple with the SupiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRangeList

`func (o *ChfInfo) SetSupiRangeList(v []SupiRange)`

SetSupiRangeList sets SupiRangeList field to given value.

### HasSupiRangeList

`func (o *ChfInfo) HasSupiRangeList() bool`

HasSupiRangeList returns a boolean if a field has been set.

### GetGpsiRangeList

`func (o *ChfInfo) GetGpsiRangeList() []IdentityRange`

GetGpsiRangeList returns the GpsiRangeList field if non-nil, zero value otherwise.

### GetGpsiRangeListOk

`func (o *ChfInfo) GetGpsiRangeListOk() (*[]IdentityRange, bool)`

GetGpsiRangeListOk returns a tuple with the GpsiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRangeList

`func (o *ChfInfo) SetGpsiRangeList(v []IdentityRange)`

SetGpsiRangeList sets GpsiRangeList field to given value.

### HasGpsiRangeList

`func (o *ChfInfo) HasGpsiRangeList() bool`

HasGpsiRangeList returns a boolean if a field has been set.

### GetPlmnRangeList

`func (o *ChfInfo) GetPlmnRangeList() []PlmnRange`

GetPlmnRangeList returns the PlmnRangeList field if non-nil, zero value otherwise.

### GetPlmnRangeListOk

`func (o *ChfInfo) GetPlmnRangeListOk() (*[]PlmnRange, bool)`

GetPlmnRangeListOk returns a tuple with the PlmnRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnRangeList

`func (o *ChfInfo) SetPlmnRangeList(v []PlmnRange)`

SetPlmnRangeList sets PlmnRangeList field to given value.

### HasPlmnRangeList

`func (o *ChfInfo) HasPlmnRangeList() bool`

HasPlmnRangeList returns a boolean if a field has been set.

### GetGroupId

`func (o *ChfInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ChfInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ChfInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ChfInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPrimaryChfInstance

`func (o *ChfInfo) GetPrimaryChfInstance() string`

GetPrimaryChfInstance returns the PrimaryChfInstance field if non-nil, zero value otherwise.

### GetPrimaryChfInstanceOk

`func (o *ChfInfo) GetPrimaryChfInstanceOk() (*string, bool)`

GetPrimaryChfInstanceOk returns a tuple with the PrimaryChfInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryChfInstance

`func (o *ChfInfo) SetPrimaryChfInstance(v string)`

SetPrimaryChfInstance sets PrimaryChfInstance field to given value.

### HasPrimaryChfInstance

`func (o *ChfInfo) HasPrimaryChfInstance() bool`

HasPrimaryChfInstance returns a boolean if a field has been set.

### GetSecondaryChfInstance

`func (o *ChfInfo) GetSecondaryChfInstance() string`

GetSecondaryChfInstance returns the SecondaryChfInstance field if non-nil, zero value otherwise.

### GetSecondaryChfInstanceOk

`func (o *ChfInfo) GetSecondaryChfInstanceOk() (*string, bool)`

GetSecondaryChfInstanceOk returns a tuple with the SecondaryChfInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryChfInstance

`func (o *ChfInfo) SetSecondaryChfInstance(v string)`

SetSecondaryChfInstance sets SecondaryChfInstance field to given value.

### HasSecondaryChfInstance

`func (o *ChfInfo) HasSecondaryChfInstance() bool`

HasSecondaryChfInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


