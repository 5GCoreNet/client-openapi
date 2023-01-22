# UpuInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpuDataList** | Pointer to [**[]UpuData2**](UpuData2.md) |  | [optional] 
**UpuRegInd** | Pointer to **bool** |  | [optional] 
**UpuAckInd** | Pointer to **bool** | Contains the indication of whether the acknowledgement from UE is needed. | [optional] 
**UpuMacIausf** | Pointer to **string** | MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE). | [optional] 
**CounterUpu** | Pointer to **string** | CounterUPU. | [optional] 
**ProvisioningTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**UpuTransparentContainer** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 

## Methods

### NewUpuInfo1

`func NewUpuInfo1(provisioningTime time.Time, ) *UpuInfo1`

NewUpuInfo1 instantiates a new UpuInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpuInfo1WithDefaults

`func NewUpuInfo1WithDefaults() *UpuInfo1`

NewUpuInfo1WithDefaults instantiates a new UpuInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpuDataList

`func (o *UpuInfo1) GetUpuDataList() []UpuData2`

GetUpuDataList returns the UpuDataList field if non-nil, zero value otherwise.

### GetUpuDataListOk

`func (o *UpuInfo1) GetUpuDataListOk() (*[]UpuData2, bool)`

GetUpuDataListOk returns a tuple with the UpuDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuDataList

`func (o *UpuInfo1) SetUpuDataList(v []UpuData2)`

SetUpuDataList sets UpuDataList field to given value.

### HasUpuDataList

`func (o *UpuInfo1) HasUpuDataList() bool`

HasUpuDataList returns a boolean if a field has been set.

### GetUpuRegInd

`func (o *UpuInfo1) GetUpuRegInd() bool`

GetUpuRegInd returns the UpuRegInd field if non-nil, zero value otherwise.

### GetUpuRegIndOk

`func (o *UpuInfo1) GetUpuRegIndOk() (*bool, bool)`

GetUpuRegIndOk returns a tuple with the UpuRegInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuRegInd

`func (o *UpuInfo1) SetUpuRegInd(v bool)`

SetUpuRegInd sets UpuRegInd field to given value.

### HasUpuRegInd

`func (o *UpuInfo1) HasUpuRegInd() bool`

HasUpuRegInd returns a boolean if a field has been set.

### GetUpuAckInd

`func (o *UpuInfo1) GetUpuAckInd() bool`

GetUpuAckInd returns the UpuAckInd field if non-nil, zero value otherwise.

### GetUpuAckIndOk

`func (o *UpuInfo1) GetUpuAckIndOk() (*bool, bool)`

GetUpuAckIndOk returns a tuple with the UpuAckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuAckInd

`func (o *UpuInfo1) SetUpuAckInd(v bool)`

SetUpuAckInd sets UpuAckInd field to given value.

### HasUpuAckInd

`func (o *UpuInfo1) HasUpuAckInd() bool`

HasUpuAckInd returns a boolean if a field has been set.

### GetUpuMacIausf

`func (o *UpuInfo1) GetUpuMacIausf() string`

GetUpuMacIausf returns the UpuMacIausf field if non-nil, zero value otherwise.

### GetUpuMacIausfOk

`func (o *UpuInfo1) GetUpuMacIausfOk() (*string, bool)`

GetUpuMacIausfOk returns a tuple with the UpuMacIausf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuMacIausf

`func (o *UpuInfo1) SetUpuMacIausf(v string)`

SetUpuMacIausf sets UpuMacIausf field to given value.

### HasUpuMacIausf

`func (o *UpuInfo1) HasUpuMacIausf() bool`

HasUpuMacIausf returns a boolean if a field has been set.

### GetCounterUpu

`func (o *UpuInfo1) GetCounterUpu() string`

GetCounterUpu returns the CounterUpu field if non-nil, zero value otherwise.

### GetCounterUpuOk

`func (o *UpuInfo1) GetCounterUpuOk() (*string, bool)`

GetCounterUpuOk returns a tuple with the CounterUpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterUpu

`func (o *UpuInfo1) SetCounterUpu(v string)`

SetCounterUpu sets CounterUpu field to given value.

### HasCounterUpu

`func (o *UpuInfo1) HasCounterUpu() bool`

HasCounterUpu returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *UpuInfo1) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *UpuInfo1) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *UpuInfo1) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.


### GetUpuTransparentContainer

`func (o *UpuInfo1) GetUpuTransparentContainer() string`

GetUpuTransparentContainer returns the UpuTransparentContainer field if non-nil, zero value otherwise.

### GetUpuTransparentContainerOk

`func (o *UpuInfo1) GetUpuTransparentContainerOk() (*string, bool)`

GetUpuTransparentContainerOk returns a tuple with the UpuTransparentContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpuTransparentContainer

`func (o *UpuInfo1) SetUpuTransparentContainer(v string)`

SetUpuTransparentContainer sets UpuTransparentContainer field to given value.

### HasUpuTransparentContainer

`func (o *UpuInfo1) HasUpuTransparentContainer() bool`

HasUpuTransparentContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


