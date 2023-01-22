# FileDistributionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Represents the group ID for which a V2X message is addressed. | [optional] 
**FileLists** | [**[]FileList**](FileList.md) |  | 
**ServiceClass** | Pointer to **string** |  | [optional] 
**GeoArea** | [**GeographicArea**](GeographicArea.md) |  | 
**MaxBitrate** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**MaxDelay** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**LocalMbmsInfo** | Pointer to [**LocalMbmsInfo**](LocalMbmsInfo.md) |  | [optional] 
**LocalMbmsActInd** | Pointer to **bool** |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewFileDistributionData

`func NewFileDistributionData(fileLists []FileList, geoArea GeographicArea, maxBitrate string, maxDelay int32, ) *FileDistributionData`

NewFileDistributionData instantiates a new FileDistributionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDistributionDataWithDefaults

`func NewFileDistributionDataWithDefaults() *FileDistributionData`

NewFileDistributionDataWithDefaults instantiates a new FileDistributionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *FileDistributionData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FileDistributionData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FileDistributionData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *FileDistributionData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetFileLists

`func (o *FileDistributionData) GetFileLists() []FileList`

GetFileLists returns the FileLists field if non-nil, zero value otherwise.

### GetFileListsOk

`func (o *FileDistributionData) GetFileListsOk() (*[]FileList, bool)`

GetFileListsOk returns a tuple with the FileLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLists

`func (o *FileDistributionData) SetFileLists(v []FileList)`

SetFileLists sets FileLists field to given value.


### GetServiceClass

`func (o *FileDistributionData) GetServiceClass() string`

GetServiceClass returns the ServiceClass field if non-nil, zero value otherwise.

### GetServiceClassOk

`func (o *FileDistributionData) GetServiceClassOk() (*string, bool)`

GetServiceClassOk returns a tuple with the ServiceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClass

`func (o *FileDistributionData) SetServiceClass(v string)`

SetServiceClass sets ServiceClass field to given value.

### HasServiceClass

`func (o *FileDistributionData) HasServiceClass() bool`

HasServiceClass returns a boolean if a field has been set.

### GetGeoArea

`func (o *FileDistributionData) GetGeoArea() GeographicArea`

GetGeoArea returns the GeoArea field if non-nil, zero value otherwise.

### GetGeoAreaOk

`func (o *FileDistributionData) GetGeoAreaOk() (*GeographicArea, bool)`

GetGeoAreaOk returns a tuple with the GeoArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoArea

`func (o *FileDistributionData) SetGeoArea(v GeographicArea)`

SetGeoArea sets GeoArea field to given value.


### GetMaxBitrate

`func (o *FileDistributionData) GetMaxBitrate() string`

GetMaxBitrate returns the MaxBitrate field if non-nil, zero value otherwise.

### GetMaxBitrateOk

`func (o *FileDistributionData) GetMaxBitrateOk() (*string, bool)`

GetMaxBitrateOk returns a tuple with the MaxBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitrate

`func (o *FileDistributionData) SetMaxBitrate(v string)`

SetMaxBitrate sets MaxBitrate field to given value.


### GetMaxDelay

`func (o *FileDistributionData) GetMaxDelay() int32`

GetMaxDelay returns the MaxDelay field if non-nil, zero value otherwise.

### GetMaxDelayOk

`func (o *FileDistributionData) GetMaxDelayOk() (*int32, bool)`

GetMaxDelayOk returns a tuple with the MaxDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDelay

`func (o *FileDistributionData) SetMaxDelay(v int32)`

SetMaxDelay sets MaxDelay field to given value.


### GetLocalMbmsInfo

`func (o *FileDistributionData) GetLocalMbmsInfo() LocalMbmsInfo`

GetLocalMbmsInfo returns the LocalMbmsInfo field if non-nil, zero value otherwise.

### GetLocalMbmsInfoOk

`func (o *FileDistributionData) GetLocalMbmsInfoOk() (*LocalMbmsInfo, bool)`

GetLocalMbmsInfoOk returns a tuple with the LocalMbmsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMbmsInfo

`func (o *FileDistributionData) SetLocalMbmsInfo(v LocalMbmsInfo)`

SetLocalMbmsInfo sets LocalMbmsInfo field to given value.

### HasLocalMbmsInfo

`func (o *FileDistributionData) HasLocalMbmsInfo() bool`

HasLocalMbmsInfo returns a boolean if a field has been set.

### GetLocalMbmsActInd

`func (o *FileDistributionData) GetLocalMbmsActInd() bool`

GetLocalMbmsActInd returns the LocalMbmsActInd field if non-nil, zero value otherwise.

### GetLocalMbmsActIndOk

`func (o *FileDistributionData) GetLocalMbmsActIndOk() (*bool, bool)`

GetLocalMbmsActIndOk returns a tuple with the LocalMbmsActInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMbmsActInd

`func (o *FileDistributionData) SetLocalMbmsActInd(v bool)`

SetLocalMbmsActInd sets LocalMbmsActInd field to given value.

### HasLocalMbmsActInd

`func (o *FileDistributionData) HasLocalMbmsActInd() bool`

HasLocalMbmsActInd returns a boolean if a field has been set.

### GetSuppFeat

`func (o *FileDistributionData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *FileDistributionData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *FileDistributionData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *FileDistributionData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


