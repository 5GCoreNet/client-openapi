# TargetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetRpauid** | **string** | Contains the RPAUID. | 
**Pduid** | **string** | Contains the PDUID. | 
**MetadataIndic** | Pointer to [**MetadataIndic**](MetadataIndic.md) |  | [optional] 

## Methods

### NewTargetData

`func NewTargetData(targetRpauid string, pduid string, ) *TargetData`

NewTargetData instantiates a new TargetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetDataWithDefaults

`func NewTargetDataWithDefaults() *TargetData`

NewTargetDataWithDefaults instantiates a new TargetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetRpauid

`func (o *TargetData) GetTargetRpauid() string`

GetTargetRpauid returns the TargetRpauid field if non-nil, zero value otherwise.

### GetTargetRpauidOk

`func (o *TargetData) GetTargetRpauidOk() (*string, bool)`

GetTargetRpauidOk returns a tuple with the TargetRpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRpauid

`func (o *TargetData) SetTargetRpauid(v string)`

SetTargetRpauid sets TargetRpauid field to given value.


### GetPduid

`func (o *TargetData) GetPduid() string`

GetPduid returns the Pduid field if non-nil, zero value otherwise.

### GetPduidOk

`func (o *TargetData) GetPduidOk() (*string, bool)`

GetPduidOk returns a tuple with the Pduid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduid

`func (o *TargetData) SetPduid(v string)`

SetPduid sets Pduid field to given value.


### GetMetadataIndic

`func (o *TargetData) GetMetadataIndic() MetadataIndic`

GetMetadataIndic returns the MetadataIndic field if non-nil, zero value otherwise.

### GetMetadataIndicOk

`func (o *TargetData) GetMetadataIndicOk() (*MetadataIndic, bool)`

GetMetadataIndicOk returns a tuple with the MetadataIndic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataIndic

`func (o *TargetData) SetMetadataIndic(v MetadataIndic)`

SetMetadataIndic sets MetadataIndic field to given value.

### HasMetadataIndic

`func (o *TargetData) HasMetadataIndic() bool`

HasMetadataIndic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


