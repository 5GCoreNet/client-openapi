# SoftwareImageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumDisk** | Pointer to **int32** |  | [optional] 
**MinimumRAM** | Pointer to **int32** |  | [optional] 
**DiscFormat** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**SwImageRef** | Pointer to **string** | indicates the reference to the actual software image that is represented by URL (see clause 7.1.6.5 in ETSI NFV IFA-011 [7]). | [optional] 

## Methods

### NewSoftwareImageInfo

`func NewSoftwareImageInfo() *SoftwareImageInfo`

NewSoftwareImageInfo instantiates a new SoftwareImageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareImageInfoWithDefaults

`func NewSoftwareImageInfoWithDefaults() *SoftwareImageInfo`

NewSoftwareImageInfoWithDefaults instantiates a new SoftwareImageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumDisk

`func (o *SoftwareImageInfo) GetMinimumDisk() int32`

GetMinimumDisk returns the MinimumDisk field if non-nil, zero value otherwise.

### GetMinimumDiskOk

`func (o *SoftwareImageInfo) GetMinimumDiskOk() (*int32, bool)`

GetMinimumDiskOk returns a tuple with the MinimumDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumDisk

`func (o *SoftwareImageInfo) SetMinimumDisk(v int32)`

SetMinimumDisk sets MinimumDisk field to given value.

### HasMinimumDisk

`func (o *SoftwareImageInfo) HasMinimumDisk() bool`

HasMinimumDisk returns a boolean if a field has been set.

### GetMinimumRAM

`func (o *SoftwareImageInfo) GetMinimumRAM() int32`

GetMinimumRAM returns the MinimumRAM field if non-nil, zero value otherwise.

### GetMinimumRAMOk

`func (o *SoftwareImageInfo) GetMinimumRAMOk() (*int32, bool)`

GetMinimumRAMOk returns a tuple with the MinimumRAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRAM

`func (o *SoftwareImageInfo) SetMinimumRAM(v int32)`

SetMinimumRAM sets MinimumRAM field to given value.

### HasMinimumRAM

`func (o *SoftwareImageInfo) HasMinimumRAM() bool`

HasMinimumRAM returns a boolean if a field has been set.

### GetDiscFormat

`func (o *SoftwareImageInfo) GetDiscFormat() string`

GetDiscFormat returns the DiscFormat field if non-nil, zero value otherwise.

### GetDiscFormatOk

`func (o *SoftwareImageInfo) GetDiscFormatOk() (*string, bool)`

GetDiscFormatOk returns a tuple with the DiscFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscFormat

`func (o *SoftwareImageInfo) SetDiscFormat(v string)`

SetDiscFormat sets DiscFormat field to given value.

### HasDiscFormat

`func (o *SoftwareImageInfo) HasDiscFormat() bool`

HasDiscFormat returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *SoftwareImageInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *SoftwareImageInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *SoftwareImageInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *SoftwareImageInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetSwImageRef

`func (o *SoftwareImageInfo) GetSwImageRef() string`

GetSwImageRef returns the SwImageRef field if non-nil, zero value otherwise.

### GetSwImageRefOk

`func (o *SoftwareImageInfo) GetSwImageRefOk() (*string, bool)`

GetSwImageRefOk returns a tuple with the SwImageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwImageRef

`func (o *SoftwareImageInfo) SetSwImageRef(v string)`

SetSwImageRef sets SwImageRef field to given value.

### HasSwImageRef

`func (o *SoftwareImageInfo) HasSwImageRef() bool`

HasSwImageRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


