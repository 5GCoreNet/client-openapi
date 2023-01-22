# MbsMediaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsMedType** | Pointer to [**MediaType**](MediaType.md) |  | [optional] 
**MaxReqMbsBwDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MinReqMbsBwDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**Codecs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMbsMediaInfo

`func NewMbsMediaInfo() *MbsMediaInfo`

NewMbsMediaInfo instantiates a new MbsMediaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsMediaInfoWithDefaults

`func NewMbsMediaInfoWithDefaults() *MbsMediaInfo`

NewMbsMediaInfoWithDefaults instantiates a new MbsMediaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsMedType

`func (o *MbsMediaInfo) GetMbsMedType() MediaType`

GetMbsMedType returns the MbsMedType field if non-nil, zero value otherwise.

### GetMbsMedTypeOk

`func (o *MbsMediaInfo) GetMbsMedTypeOk() (*MediaType, bool)`

GetMbsMedTypeOk returns a tuple with the MbsMedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsMedType

`func (o *MbsMediaInfo) SetMbsMedType(v MediaType)`

SetMbsMedType sets MbsMedType field to given value.

### HasMbsMedType

`func (o *MbsMediaInfo) HasMbsMedType() bool`

HasMbsMedType returns a boolean if a field has been set.

### GetMaxReqMbsBwDl

`func (o *MbsMediaInfo) GetMaxReqMbsBwDl() string`

GetMaxReqMbsBwDl returns the MaxReqMbsBwDl field if non-nil, zero value otherwise.

### GetMaxReqMbsBwDlOk

`func (o *MbsMediaInfo) GetMaxReqMbsBwDlOk() (*string, bool)`

GetMaxReqMbsBwDlOk returns a tuple with the MaxReqMbsBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReqMbsBwDl

`func (o *MbsMediaInfo) SetMaxReqMbsBwDl(v string)`

SetMaxReqMbsBwDl sets MaxReqMbsBwDl field to given value.

### HasMaxReqMbsBwDl

`func (o *MbsMediaInfo) HasMaxReqMbsBwDl() bool`

HasMaxReqMbsBwDl returns a boolean if a field has been set.

### GetMinReqMbsBwDl

`func (o *MbsMediaInfo) GetMinReqMbsBwDl() string`

GetMinReqMbsBwDl returns the MinReqMbsBwDl field if non-nil, zero value otherwise.

### GetMinReqMbsBwDlOk

`func (o *MbsMediaInfo) GetMinReqMbsBwDlOk() (*string, bool)`

GetMinReqMbsBwDlOk returns a tuple with the MinReqMbsBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReqMbsBwDl

`func (o *MbsMediaInfo) SetMinReqMbsBwDl(v string)`

SetMinReqMbsBwDl sets MinReqMbsBwDl field to given value.

### HasMinReqMbsBwDl

`func (o *MbsMediaInfo) HasMinReqMbsBwDl() bool`

HasMinReqMbsBwDl returns a boolean if a field has been set.

### GetCodecs

`func (o *MbsMediaInfo) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *MbsMediaInfo) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *MbsMediaInfo) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *MbsMediaInfo) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


