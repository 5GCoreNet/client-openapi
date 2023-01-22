# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileLocation** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** |  | [optional] 
**FileReadyTime** | Pointer to **time.Time** |  | [optional] 
**FileExpirationTime** | Pointer to **time.Time** |  | [optional] 
**FileCompression** | Pointer to **string** |  | [optional] 
**FileFormat** | Pointer to **string** |  | [optional] 
**FileDataType** | Pointer to [**FileDataType**](FileDataType.md) |  | [optional] 

## Methods

### NewFileInfo

`func NewFileInfo() *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileLocation

`func (o *FileInfo) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *FileInfo) GetFileLocationOk() (*string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *FileInfo) SetFileLocation(v string)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *FileInfo) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetFileSize

`func (o *FileInfo) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileInfo) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileInfo) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FileInfo) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileReadyTime

`func (o *FileInfo) GetFileReadyTime() time.Time`

GetFileReadyTime returns the FileReadyTime field if non-nil, zero value otherwise.

### GetFileReadyTimeOk

`func (o *FileInfo) GetFileReadyTimeOk() (*time.Time, bool)`

GetFileReadyTimeOk returns a tuple with the FileReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileReadyTime

`func (o *FileInfo) SetFileReadyTime(v time.Time)`

SetFileReadyTime sets FileReadyTime field to given value.

### HasFileReadyTime

`func (o *FileInfo) HasFileReadyTime() bool`

HasFileReadyTime returns a boolean if a field has been set.

### GetFileExpirationTime

`func (o *FileInfo) GetFileExpirationTime() time.Time`

GetFileExpirationTime returns the FileExpirationTime field if non-nil, zero value otherwise.

### GetFileExpirationTimeOk

`func (o *FileInfo) GetFileExpirationTimeOk() (*time.Time, bool)`

GetFileExpirationTimeOk returns a tuple with the FileExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExpirationTime

`func (o *FileInfo) SetFileExpirationTime(v time.Time)`

SetFileExpirationTime sets FileExpirationTime field to given value.

### HasFileExpirationTime

`func (o *FileInfo) HasFileExpirationTime() bool`

HasFileExpirationTime returns a boolean if a field has been set.

### GetFileCompression

`func (o *FileInfo) GetFileCompression() string`

GetFileCompression returns the FileCompression field if non-nil, zero value otherwise.

### GetFileCompressionOk

`func (o *FileInfo) GetFileCompressionOk() (*string, bool)`

GetFileCompressionOk returns a tuple with the FileCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCompression

`func (o *FileInfo) SetFileCompression(v string)`

SetFileCompression sets FileCompression field to given value.

### HasFileCompression

`func (o *FileInfo) HasFileCompression() bool`

HasFileCompression returns a boolean if a field has been set.

### GetFileFormat

`func (o *FileInfo) GetFileFormat() string`

GetFileFormat returns the FileFormat field if non-nil, zero value otherwise.

### GetFileFormatOk

`func (o *FileInfo) GetFileFormatOk() (*string, bool)`

GetFileFormatOk returns a tuple with the FileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFormat

`func (o *FileInfo) SetFileFormat(v string)`

SetFileFormat sets FileFormat field to given value.

### HasFileFormat

`func (o *FileInfo) HasFileFormat() bool`

HasFileFormat returns a boolean if a field has been set.

### GetFileDataType

`func (o *FileInfo) GetFileDataType() FileDataType`

GetFileDataType returns the FileDataType field if non-nil, zero value otherwise.

### GetFileDataTypeOk

`func (o *FileInfo) GetFileDataTypeOk() (*FileDataType, bool)`

GetFileDataTypeOk returns a tuple with the FileDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDataType

`func (o *FileInfo) SetFileDataType(v FileDataType)`

SetFileDataType sets FileDataType field to given value.

### HasFileDataType

`func (o *FileInfo) HasFileDataType() bool`

HasFileDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


