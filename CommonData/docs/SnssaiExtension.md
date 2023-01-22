# SnssaiExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdRanges** | Pointer to [**[]SdRange**](SdRange.md) | When present, it shall contain the range(s) of Slice Differentiator values supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type  | [optional] 
**WildcardSd** | Pointer to **bool** | When present, it shall be set to true, to indicate that all SD values are supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type.  | [optional] 

## Methods

### NewSnssaiExtension

`func NewSnssaiExtension() *SnssaiExtension`

NewSnssaiExtension instantiates a new SnssaiExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnssaiExtensionWithDefaults

`func NewSnssaiExtensionWithDefaults() *SnssaiExtension`

NewSnssaiExtensionWithDefaults instantiates a new SnssaiExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdRanges

`func (o *SnssaiExtension) GetSdRanges() []SdRange`

GetSdRanges returns the SdRanges field if non-nil, zero value otherwise.

### GetSdRangesOk

`func (o *SnssaiExtension) GetSdRangesOk() (*[]SdRange, bool)`

GetSdRangesOk returns a tuple with the SdRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdRanges

`func (o *SnssaiExtension) SetSdRanges(v []SdRange)`

SetSdRanges sets SdRanges field to given value.

### HasSdRanges

`func (o *SnssaiExtension) HasSdRanges() bool`

HasSdRanges returns a boolean if a field has been set.

### GetWildcardSd

`func (o *SnssaiExtension) GetWildcardSd() bool`

GetWildcardSd returns the WildcardSd field if non-nil, zero value otherwise.

### GetWildcardSdOk

`func (o *SnssaiExtension) GetWildcardSdOk() (*bool, bool)`

GetWildcardSdOk returns a tuple with the WildcardSd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardSd

`func (o *SnssaiExtension) SetWildcardSd(v bool)`

SetWildcardSd sets WildcardSd field to given value.

### HasWildcardSd

`func (o *SnssaiExtension) HasWildcardSd() bool`

HasWildcardSd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


