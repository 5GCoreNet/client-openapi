# VirtualResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMemory** | Pointer to **int32** |  | [optional] 
**VirtualDisk** | Pointer to **int32** |  | [optional] 
**VirtualCPU** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualResource

`func NewVirtualResource() *VirtualResource`

NewVirtualResource instantiates a new VirtualResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualResourceWithDefaults

`func NewVirtualResourceWithDefaults() *VirtualResource`

NewVirtualResourceWithDefaults instantiates a new VirtualResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMemory

`func (o *VirtualResource) GetVirtualMemory() int32`

GetVirtualMemory returns the VirtualMemory field if non-nil, zero value otherwise.

### GetVirtualMemoryOk

`func (o *VirtualResource) GetVirtualMemoryOk() (*int32, bool)`

GetVirtualMemoryOk returns a tuple with the VirtualMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMemory

`func (o *VirtualResource) SetVirtualMemory(v int32)`

SetVirtualMemory sets VirtualMemory field to given value.

### HasVirtualMemory

`func (o *VirtualResource) HasVirtualMemory() bool`

HasVirtualMemory returns a boolean if a field has been set.

### GetVirtualDisk

`func (o *VirtualResource) GetVirtualDisk() int32`

GetVirtualDisk returns the VirtualDisk field if non-nil, zero value otherwise.

### GetVirtualDiskOk

`func (o *VirtualResource) GetVirtualDiskOk() (*int32, bool)`

GetVirtualDiskOk returns a tuple with the VirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisk

`func (o *VirtualResource) SetVirtualDisk(v int32)`

SetVirtualDisk sets VirtualDisk field to given value.

### HasVirtualDisk

`func (o *VirtualResource) HasVirtualDisk() bool`

HasVirtualDisk returns a boolean if a field has been set.

### GetVirtualCPU

`func (o *VirtualResource) GetVirtualCPU() string`

GetVirtualCPU returns the VirtualCPU field if non-nil, zero value otherwise.

### GetVirtualCPUOk

`func (o *VirtualResource) GetVirtualCPUOk() (*string, bool)`

GetVirtualCPUOk returns a tuple with the VirtualCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCPU

`func (o *VirtualResource) SetVirtualCPU(v string)`

SetVirtualCPU sets VirtualCPU field to given value.

### HasVirtualCPU

`func (o *VirtualResource) HasVirtualCPU() bool`

HasVirtualCPU returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


