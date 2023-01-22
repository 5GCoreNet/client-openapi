# MbmsLocArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellId** | Pointer to **[]string** | Indicates a Cell Global Identification of the user which identifies the cell the UE is registered. | [optional] 
**EnodeBId** | Pointer to **[]string** | Indicates an eNodeB in which the UE is currently located. | [optional] 
**GeographicArea** | Pointer to [**[]GeographicArea**](GeographicArea.md) | Identifies a geographic area of the user where the UE is located. | [optional] 
**MbmsServiceAreaId** | Pointer to **[]string** | Identifies an MBMS Service Area Identity of the user where the UE is located. | [optional] 
**CivicAddress** | Pointer to [**[]CivicAddress**](CivicAddress.md) | Identifies a civic address of the user where the UE is located. | [optional] 

## Methods

### NewMbmsLocArea

`func NewMbmsLocArea() *MbmsLocArea`

NewMbmsLocArea instantiates a new MbmsLocArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbmsLocAreaWithDefaults

`func NewMbmsLocAreaWithDefaults() *MbmsLocArea`

NewMbmsLocAreaWithDefaults instantiates a new MbmsLocArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellId

`func (o *MbmsLocArea) GetCellId() []string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *MbmsLocArea) GetCellIdOk() (*[]string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *MbmsLocArea) SetCellId(v []string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *MbmsLocArea) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### GetEnodeBId

`func (o *MbmsLocArea) GetEnodeBId() []string`

GetEnodeBId returns the EnodeBId field if non-nil, zero value otherwise.

### GetEnodeBIdOk

`func (o *MbmsLocArea) GetEnodeBIdOk() (*[]string, bool)`

GetEnodeBIdOk returns a tuple with the EnodeBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnodeBId

`func (o *MbmsLocArea) SetEnodeBId(v []string)`

SetEnodeBId sets EnodeBId field to given value.

### HasEnodeBId

`func (o *MbmsLocArea) HasEnodeBId() bool`

HasEnodeBId returns a boolean if a field has been set.

### GetGeographicArea

`func (o *MbmsLocArea) GetGeographicArea() []GeographicArea`

GetGeographicArea returns the GeographicArea field if non-nil, zero value otherwise.

### GetGeographicAreaOk

`func (o *MbmsLocArea) GetGeographicAreaOk() (*[]GeographicArea, bool)`

GetGeographicAreaOk returns a tuple with the GeographicArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicArea

`func (o *MbmsLocArea) SetGeographicArea(v []GeographicArea)`

SetGeographicArea sets GeographicArea field to given value.

### HasGeographicArea

`func (o *MbmsLocArea) HasGeographicArea() bool`

HasGeographicArea returns a boolean if a field has been set.

### GetMbmsServiceAreaId

`func (o *MbmsLocArea) GetMbmsServiceAreaId() []string`

GetMbmsServiceAreaId returns the MbmsServiceAreaId field if non-nil, zero value otherwise.

### GetMbmsServiceAreaIdOk

`func (o *MbmsLocArea) GetMbmsServiceAreaIdOk() (*[]string, bool)`

GetMbmsServiceAreaIdOk returns a tuple with the MbmsServiceAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsServiceAreaId

`func (o *MbmsLocArea) SetMbmsServiceAreaId(v []string)`

SetMbmsServiceAreaId sets MbmsServiceAreaId field to given value.

### HasMbmsServiceAreaId

`func (o *MbmsLocArea) HasMbmsServiceAreaId() bool`

HasMbmsServiceAreaId returns a boolean if a field has been set.

### GetCivicAddress

`func (o *MbmsLocArea) GetCivicAddress() []CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *MbmsLocArea) GetCivicAddressOk() (*[]CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *MbmsLocArea) SetCivicAddress(v []CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *MbmsLocArea) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


