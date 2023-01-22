# BootstrappingInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtId** | **string** | Bootstrapping Transaction Identifier | 
**NafId** | [**NafId**](NafId.md) |  | 
**GbaUAware** | Pointer to **bool** |  | [optional] [default to false]
**GsIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewBootstrappingInfoRequest

`func NewBootstrappingInfoRequest(btId string, nafId NafId, ) *BootstrappingInfoRequest`

NewBootstrappingInfoRequest instantiates a new BootstrappingInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootstrappingInfoRequestWithDefaults

`func NewBootstrappingInfoRequestWithDefaults() *BootstrappingInfoRequest`

NewBootstrappingInfoRequestWithDefaults instantiates a new BootstrappingInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtId

`func (o *BootstrappingInfoRequest) GetBtId() string`

GetBtId returns the BtId field if non-nil, zero value otherwise.

### GetBtIdOk

`func (o *BootstrappingInfoRequest) GetBtIdOk() (*string, bool)`

GetBtIdOk returns a tuple with the BtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtId

`func (o *BootstrappingInfoRequest) SetBtId(v string)`

SetBtId sets BtId field to given value.


### GetNafId

`func (o *BootstrappingInfoRequest) GetNafId() NafId`

GetNafId returns the NafId field if non-nil, zero value otherwise.

### GetNafIdOk

`func (o *BootstrappingInfoRequest) GetNafIdOk() (*NafId, bool)`

GetNafIdOk returns a tuple with the NafId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNafId

`func (o *BootstrappingInfoRequest) SetNafId(v NafId)`

SetNafId sets NafId field to given value.


### GetGbaUAware

`func (o *BootstrappingInfoRequest) GetGbaUAware() bool`

GetGbaUAware returns the GbaUAware field if non-nil, zero value otherwise.

### GetGbaUAwareOk

`func (o *BootstrappingInfoRequest) GetGbaUAwareOk() (*bool, bool)`

GetGbaUAwareOk returns a tuple with the GbaUAware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbaUAware

`func (o *BootstrappingInfoRequest) SetGbaUAware(v bool)`

SetGbaUAware sets GbaUAware field to given value.

### HasGbaUAware

`func (o *BootstrappingInfoRequest) HasGbaUAware() bool`

HasGbaUAware returns a boolean if a field has been set.

### GetGsIds

`func (o *BootstrappingInfoRequest) GetGsIds() []int32`

GetGsIds returns the GsIds field if non-nil, zero value otherwise.

### GetGsIdsOk

`func (o *BootstrappingInfoRequest) GetGsIdsOk() (*[]int32, bool)`

GetGsIdsOk returns a tuple with the GsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGsIds

`func (o *BootstrappingInfoRequest) SetGsIds(v []int32)`

SetGsIds sets GsIds field to given value.

### HasGsIds

`func (o *BootstrappingInfoRequest) HasGsIds() bool`

HasGsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


