# PC5DataContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSequenceNumber** | Pointer to **string** |  | [optional] 
**ChangeTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**CoverageStatus** | Pointer to **bool** |  | [optional] 
**UserLocationInformation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**DataVolume** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**ChangeCondition** | Pointer to **string** |  | [optional] 
**RadioResourcesId** | Pointer to [**RadioResourcesId**](RadioResourcesId.md) |  | [optional] 
**RadioFrequency** | Pointer to **string** |  | [optional] 
**PC5RadioTechnology** | Pointer to **string** |  | [optional] 

## Methods

### NewPC5DataContainer

`func NewPC5DataContainer() *PC5DataContainer`

NewPC5DataContainer instantiates a new PC5DataContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPC5DataContainerWithDefaults

`func NewPC5DataContainerWithDefaults() *PC5DataContainer`

NewPC5DataContainerWithDefaults instantiates a new PC5DataContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSequenceNumber

`func (o *PC5DataContainer) GetLocalSequenceNumber() string`

GetLocalSequenceNumber returns the LocalSequenceNumber field if non-nil, zero value otherwise.

### GetLocalSequenceNumberOk

`func (o *PC5DataContainer) GetLocalSequenceNumberOk() (*string, bool)`

GetLocalSequenceNumberOk returns a tuple with the LocalSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSequenceNumber

`func (o *PC5DataContainer) SetLocalSequenceNumber(v string)`

SetLocalSequenceNumber sets LocalSequenceNumber field to given value.

### HasLocalSequenceNumber

`func (o *PC5DataContainer) HasLocalSequenceNumber() bool`

HasLocalSequenceNumber returns a boolean if a field has been set.

### GetChangeTime

`func (o *PC5DataContainer) GetChangeTime() time.Time`

GetChangeTime returns the ChangeTime field if non-nil, zero value otherwise.

### GetChangeTimeOk

`func (o *PC5DataContainer) GetChangeTimeOk() (*time.Time, bool)`

GetChangeTimeOk returns a tuple with the ChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTime

`func (o *PC5DataContainer) SetChangeTime(v time.Time)`

SetChangeTime sets ChangeTime field to given value.

### HasChangeTime

`func (o *PC5DataContainer) HasChangeTime() bool`

HasChangeTime returns a boolean if a field has been set.

### GetCoverageStatus

`func (o *PC5DataContainer) GetCoverageStatus() bool`

GetCoverageStatus returns the CoverageStatus field if non-nil, zero value otherwise.

### GetCoverageStatusOk

`func (o *PC5DataContainer) GetCoverageStatusOk() (*bool, bool)`

GetCoverageStatusOk returns a tuple with the CoverageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageStatus

`func (o *PC5DataContainer) SetCoverageStatus(v bool)`

SetCoverageStatus sets CoverageStatus field to given value.

### HasCoverageStatus

`func (o *PC5DataContainer) HasCoverageStatus() bool`

HasCoverageStatus returns a boolean if a field has been set.

### GetUserLocationInformation

`func (o *PC5DataContainer) GetUserLocationInformation() UserLocation`

GetUserLocationInformation returns the UserLocationInformation field if non-nil, zero value otherwise.

### GetUserLocationInformationOk

`func (o *PC5DataContainer) GetUserLocationInformationOk() (*UserLocation, bool)`

GetUserLocationInformationOk returns a tuple with the UserLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInformation

`func (o *PC5DataContainer) SetUserLocationInformation(v UserLocation)`

SetUserLocationInformation sets UserLocationInformation field to given value.

### HasUserLocationInformation

`func (o *PC5DataContainer) HasUserLocationInformation() bool`

HasUserLocationInformation returns a boolean if a field has been set.

### GetDataVolume

`func (o *PC5DataContainer) GetDataVolume() int32`

GetDataVolume returns the DataVolume field if non-nil, zero value otherwise.

### GetDataVolumeOk

`func (o *PC5DataContainer) GetDataVolumeOk() (*int32, bool)`

GetDataVolumeOk returns a tuple with the DataVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVolume

`func (o *PC5DataContainer) SetDataVolume(v int32)`

SetDataVolume sets DataVolume field to given value.

### HasDataVolume

`func (o *PC5DataContainer) HasDataVolume() bool`

HasDataVolume returns a boolean if a field has been set.

### GetChangeCondition

`func (o *PC5DataContainer) GetChangeCondition() string`

GetChangeCondition returns the ChangeCondition field if non-nil, zero value otherwise.

### GetChangeConditionOk

`func (o *PC5DataContainer) GetChangeConditionOk() (*string, bool)`

GetChangeConditionOk returns a tuple with the ChangeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeCondition

`func (o *PC5DataContainer) SetChangeCondition(v string)`

SetChangeCondition sets ChangeCondition field to given value.

### HasChangeCondition

`func (o *PC5DataContainer) HasChangeCondition() bool`

HasChangeCondition returns a boolean if a field has been set.

### GetRadioResourcesId

`func (o *PC5DataContainer) GetRadioResourcesId() RadioResourcesId`

GetRadioResourcesId returns the RadioResourcesId field if non-nil, zero value otherwise.

### GetRadioResourcesIdOk

`func (o *PC5DataContainer) GetRadioResourcesIdOk() (*RadioResourcesId, bool)`

GetRadioResourcesIdOk returns a tuple with the RadioResourcesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioResourcesId

`func (o *PC5DataContainer) SetRadioResourcesId(v RadioResourcesId)`

SetRadioResourcesId sets RadioResourcesId field to given value.

### HasRadioResourcesId

`func (o *PC5DataContainer) HasRadioResourcesId() bool`

HasRadioResourcesId returns a boolean if a field has been set.

### GetRadioFrequency

`func (o *PC5DataContainer) GetRadioFrequency() string`

GetRadioFrequency returns the RadioFrequency field if non-nil, zero value otherwise.

### GetRadioFrequencyOk

`func (o *PC5DataContainer) GetRadioFrequencyOk() (*string, bool)`

GetRadioFrequencyOk returns a tuple with the RadioFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioFrequency

`func (o *PC5DataContainer) SetRadioFrequency(v string)`

SetRadioFrequency sets RadioFrequency field to given value.

### HasRadioFrequency

`func (o *PC5DataContainer) HasRadioFrequency() bool`

HasRadioFrequency returns a boolean if a field has been set.

### GetPC5RadioTechnology

`func (o *PC5DataContainer) GetPC5RadioTechnology() string`

GetPC5RadioTechnology returns the PC5RadioTechnology field if non-nil, zero value otherwise.

### GetPC5RadioTechnologyOk

`func (o *PC5DataContainer) GetPC5RadioTechnologyOk() (*string, bool)`

GetPC5RadioTechnologyOk returns a tuple with the PC5RadioTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPC5RadioTechnology

`func (o *PC5DataContainer) SetPC5RadioTechnology(v string)`

SetPC5RadioTechnology sets PC5RadioTechnology field to given value.

### HasPC5RadioTechnology

`func (o *PC5DataContainer) HasPC5RadioTechnology() bool`

HasPC5RadioTechnology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


