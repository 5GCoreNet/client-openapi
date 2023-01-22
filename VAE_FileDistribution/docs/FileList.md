# FileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**FileDisplayUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**FileEarFetchTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**FileLatFetchTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**FileSize** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**FileStatus** | [**FileStatus**](FileStatus.md) |  | 
**CompletionTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**KeepUpdateInterval** | **int32** | indicating a time in seconds. | 
**UniAvailability** | Pointer to **bool** |  | [optional] 
**FileRepetition** | Pointer to **int32** |  | [optional] 

## Methods

### NewFileList

`func NewFileList(fileUri string, fileDisplayUri string, fileEarFetchTime time.Time, fileLatFetchTime time.Time, fileStatus FileStatus, completionTime time.Time, keepUpdateInterval int32, ) *FileList`

NewFileList instantiates a new FileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileListWithDefaults

`func NewFileListWithDefaults() *FileList`

NewFileListWithDefaults instantiates a new FileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUri

`func (o *FileList) GetFileUri() string`

GetFileUri returns the FileUri field if non-nil, zero value otherwise.

### GetFileUriOk

`func (o *FileList) GetFileUriOk() (*string, bool)`

GetFileUriOk returns a tuple with the FileUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUri

`func (o *FileList) SetFileUri(v string)`

SetFileUri sets FileUri field to given value.


### GetFileDisplayUri

`func (o *FileList) GetFileDisplayUri() string`

GetFileDisplayUri returns the FileDisplayUri field if non-nil, zero value otherwise.

### GetFileDisplayUriOk

`func (o *FileList) GetFileDisplayUriOk() (*string, bool)`

GetFileDisplayUriOk returns a tuple with the FileDisplayUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDisplayUri

`func (o *FileList) SetFileDisplayUri(v string)`

SetFileDisplayUri sets FileDisplayUri field to given value.


### GetFileEarFetchTime

`func (o *FileList) GetFileEarFetchTime() time.Time`

GetFileEarFetchTime returns the FileEarFetchTime field if non-nil, zero value otherwise.

### GetFileEarFetchTimeOk

`func (o *FileList) GetFileEarFetchTimeOk() (*time.Time, bool)`

GetFileEarFetchTimeOk returns a tuple with the FileEarFetchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEarFetchTime

`func (o *FileList) SetFileEarFetchTime(v time.Time)`

SetFileEarFetchTime sets FileEarFetchTime field to given value.


### GetFileLatFetchTime

`func (o *FileList) GetFileLatFetchTime() time.Time`

GetFileLatFetchTime returns the FileLatFetchTime field if non-nil, zero value otherwise.

### GetFileLatFetchTimeOk

`func (o *FileList) GetFileLatFetchTimeOk() (*time.Time, bool)`

GetFileLatFetchTimeOk returns a tuple with the FileLatFetchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLatFetchTime

`func (o *FileList) SetFileLatFetchTime(v time.Time)`

SetFileLatFetchTime sets FileLatFetchTime field to given value.


### GetFileSize

`func (o *FileList) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileList) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileList) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FileList) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileStatus

`func (o *FileList) GetFileStatus() FileStatus`

GetFileStatus returns the FileStatus field if non-nil, zero value otherwise.

### GetFileStatusOk

`func (o *FileList) GetFileStatusOk() (*FileStatus, bool)`

GetFileStatusOk returns a tuple with the FileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStatus

`func (o *FileList) SetFileStatus(v FileStatus)`

SetFileStatus sets FileStatus field to given value.


### GetCompletionTime

`func (o *FileList) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *FileList) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *FileList) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.


### GetKeepUpdateInterval

`func (o *FileList) GetKeepUpdateInterval() int32`

GetKeepUpdateInterval returns the KeepUpdateInterval field if non-nil, zero value otherwise.

### GetKeepUpdateIntervalOk

`func (o *FileList) GetKeepUpdateIntervalOk() (*int32, bool)`

GetKeepUpdateIntervalOk returns a tuple with the KeepUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUpdateInterval

`func (o *FileList) SetKeepUpdateInterval(v int32)`

SetKeepUpdateInterval sets KeepUpdateInterval field to given value.


### GetUniAvailability

`func (o *FileList) GetUniAvailability() bool`

GetUniAvailability returns the UniAvailability field if non-nil, zero value otherwise.

### GetUniAvailabilityOk

`func (o *FileList) GetUniAvailabilityOk() (*bool, bool)`

GetUniAvailabilityOk returns a tuple with the UniAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniAvailability

`func (o *FileList) SetUniAvailability(v bool)`

SetUniAvailability sets UniAvailability field to given value.

### HasUniAvailability

`func (o *FileList) HasUniAvailability() bool`

HasUniAvailability returns a boolean if a field has been set.

### GetFileRepetition

`func (o *FileList) GetFileRepetition() int32`

GetFileRepetition returns the FileRepetition field if non-nil, zero value otherwise.

### GetFileRepetitionOk

`func (o *FileList) GetFileRepetitionOk() (*int32, bool)`

GetFileRepetitionOk returns a tuple with the FileRepetition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRepetition

`func (o *FileList) SetFileRepetition(v int32)`

SetFileRepetition sets FileRepetition field to given value.

### HasFileRepetition

`func (o *FileList) HasFileRepetition() bool`

HasFileRepetition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


