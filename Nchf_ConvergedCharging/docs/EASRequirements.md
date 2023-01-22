# EASRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredEASservingLocation** | Pointer to [**ServingLocation**](ServingLocation.md) |  | [optional] 
**SoftwareImageInfo** | Pointer to [**SoftwareImageInfo**](SoftwareImageInfo.md) |  | [optional] 
**AffinityAntiAffinity** | Pointer to [**AffinityAntiAffinity**](AffinityAntiAffinity.md) |  | [optional] 
**ServiceContinuity** | Pointer to **bool** |  | [optional] 
**VirtualResource** | Pointer to [**VirtualResource**](VirtualResource.md) |  | [optional] 

## Methods

### NewEASRequirements

`func NewEASRequirements() *EASRequirements`

NewEASRequirements instantiates a new EASRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEASRequirementsWithDefaults

`func NewEASRequirementsWithDefaults() *EASRequirements`

NewEASRequirementsWithDefaults instantiates a new EASRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredEASservingLocation

`func (o *EASRequirements) GetRequiredEASservingLocation() ServingLocation`

GetRequiredEASservingLocation returns the RequiredEASservingLocation field if non-nil, zero value otherwise.

### GetRequiredEASservingLocationOk

`func (o *EASRequirements) GetRequiredEASservingLocationOk() (*ServingLocation, bool)`

GetRequiredEASservingLocationOk returns a tuple with the RequiredEASservingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredEASservingLocation

`func (o *EASRequirements) SetRequiredEASservingLocation(v ServingLocation)`

SetRequiredEASservingLocation sets RequiredEASservingLocation field to given value.

### HasRequiredEASservingLocation

`func (o *EASRequirements) HasRequiredEASservingLocation() bool`

HasRequiredEASservingLocation returns a boolean if a field has been set.

### GetSoftwareImageInfo

`func (o *EASRequirements) GetSoftwareImageInfo() SoftwareImageInfo`

GetSoftwareImageInfo returns the SoftwareImageInfo field if non-nil, zero value otherwise.

### GetSoftwareImageInfoOk

`func (o *EASRequirements) GetSoftwareImageInfoOk() (*SoftwareImageInfo, bool)`

GetSoftwareImageInfoOk returns a tuple with the SoftwareImageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageInfo

`func (o *EASRequirements) SetSoftwareImageInfo(v SoftwareImageInfo)`

SetSoftwareImageInfo sets SoftwareImageInfo field to given value.

### HasSoftwareImageInfo

`func (o *EASRequirements) HasSoftwareImageInfo() bool`

HasSoftwareImageInfo returns a boolean if a field has been set.

### GetAffinityAntiAffinity

`func (o *EASRequirements) GetAffinityAntiAffinity() AffinityAntiAffinity`

GetAffinityAntiAffinity returns the AffinityAntiAffinity field if non-nil, zero value otherwise.

### GetAffinityAntiAffinityOk

`func (o *EASRequirements) GetAffinityAntiAffinityOk() (*AffinityAntiAffinity, bool)`

GetAffinityAntiAffinityOk returns a tuple with the AffinityAntiAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityAntiAffinity

`func (o *EASRequirements) SetAffinityAntiAffinity(v AffinityAntiAffinity)`

SetAffinityAntiAffinity sets AffinityAntiAffinity field to given value.

### HasAffinityAntiAffinity

`func (o *EASRequirements) HasAffinityAntiAffinity() bool`

HasAffinityAntiAffinity returns a boolean if a field has been set.

### GetServiceContinuity

`func (o *EASRequirements) GetServiceContinuity() bool`

GetServiceContinuity returns the ServiceContinuity field if non-nil, zero value otherwise.

### GetServiceContinuityOk

`func (o *EASRequirements) GetServiceContinuityOk() (*bool, bool)`

GetServiceContinuityOk returns a tuple with the ServiceContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceContinuity

`func (o *EASRequirements) SetServiceContinuity(v bool)`

SetServiceContinuity sets ServiceContinuity field to given value.

### HasServiceContinuity

`func (o *EASRequirements) HasServiceContinuity() bool`

HasServiceContinuity returns a boolean if a field has been set.

### GetVirtualResource

`func (o *EASRequirements) GetVirtualResource() VirtualResource`

GetVirtualResource returns the VirtualResource field if non-nil, zero value otherwise.

### GetVirtualResourceOk

`func (o *EASRequirements) GetVirtualResourceOk() (*VirtualResource, bool)`

GetVirtualResourceOk returns a tuple with the VirtualResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualResource

`func (o *EASRequirements) SetVirtualResource(v VirtualResource)`

SetVirtualResource sets VirtualResource field to given value.

### HasVirtualResource

`func (o *EASRequirements) HasVirtualResource() bool`

HasVirtualResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


