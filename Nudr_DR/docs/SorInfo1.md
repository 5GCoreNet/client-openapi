# SorInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SteeringContainer** | Pointer to [**SteeringContainer**](SteeringContainer.md) |  | [optional] 
**AckInd** | **bool** | Contains indication whether the acknowledgement from UE is needed. | 
**SorMacIausf** | Pointer to **string** | MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE). | [optional] 
**Countersor** | Pointer to **string** | CounterSoR. | [optional] 
**ProvisioningTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SorTransparentContainer** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SorCmci** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**StoreSorCmciInMe** | Pointer to **bool** |  | [optional] 
**UsimSupportOfSorCmci** | Pointer to **bool** |  | [optional] 

## Methods

### NewSorInfo1

`func NewSorInfo1(ackInd bool, provisioningTime time.Time, ) *SorInfo1`

NewSorInfo1 instantiates a new SorInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSorInfo1WithDefaults

`func NewSorInfo1WithDefaults() *SorInfo1`

NewSorInfo1WithDefaults instantiates a new SorInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteeringContainer

`func (o *SorInfo1) GetSteeringContainer() SteeringContainer`

GetSteeringContainer returns the SteeringContainer field if non-nil, zero value otherwise.

### GetSteeringContainerOk

`func (o *SorInfo1) GetSteeringContainerOk() (*SteeringContainer, bool)`

GetSteeringContainerOk returns a tuple with the SteeringContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteeringContainer

`func (o *SorInfo1) SetSteeringContainer(v SteeringContainer)`

SetSteeringContainer sets SteeringContainer field to given value.

### HasSteeringContainer

`func (o *SorInfo1) HasSteeringContainer() bool`

HasSteeringContainer returns a boolean if a field has been set.

### GetAckInd

`func (o *SorInfo1) GetAckInd() bool`

GetAckInd returns the AckInd field if non-nil, zero value otherwise.

### GetAckIndOk

`func (o *SorInfo1) GetAckIndOk() (*bool, bool)`

GetAckIndOk returns a tuple with the AckInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckInd

`func (o *SorInfo1) SetAckInd(v bool)`

SetAckInd sets AckInd field to given value.


### GetSorMacIausf

`func (o *SorInfo1) GetSorMacIausf() string`

GetSorMacIausf returns the SorMacIausf field if non-nil, zero value otherwise.

### GetSorMacIausfOk

`func (o *SorInfo1) GetSorMacIausfOk() (*string, bool)`

GetSorMacIausfOk returns a tuple with the SorMacIausf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorMacIausf

`func (o *SorInfo1) SetSorMacIausf(v string)`

SetSorMacIausf sets SorMacIausf field to given value.

### HasSorMacIausf

`func (o *SorInfo1) HasSorMacIausf() bool`

HasSorMacIausf returns a boolean if a field has been set.

### GetCountersor

`func (o *SorInfo1) GetCountersor() string`

GetCountersor returns the Countersor field if non-nil, zero value otherwise.

### GetCountersorOk

`func (o *SorInfo1) GetCountersorOk() (*string, bool)`

GetCountersorOk returns a tuple with the Countersor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountersor

`func (o *SorInfo1) SetCountersor(v string)`

SetCountersor sets Countersor field to given value.

### HasCountersor

`func (o *SorInfo1) HasCountersor() bool`

HasCountersor returns a boolean if a field has been set.

### GetProvisioningTime

`func (o *SorInfo1) GetProvisioningTime() time.Time`

GetProvisioningTime returns the ProvisioningTime field if non-nil, zero value otherwise.

### GetProvisioningTimeOk

`func (o *SorInfo1) GetProvisioningTimeOk() (*time.Time, bool)`

GetProvisioningTimeOk returns a tuple with the ProvisioningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTime

`func (o *SorInfo1) SetProvisioningTime(v time.Time)`

SetProvisioningTime sets ProvisioningTime field to given value.


### GetSorTransparentContainer

`func (o *SorInfo1) GetSorTransparentContainer() string`

GetSorTransparentContainer returns the SorTransparentContainer field if non-nil, zero value otherwise.

### GetSorTransparentContainerOk

`func (o *SorInfo1) GetSorTransparentContainerOk() (*string, bool)`

GetSorTransparentContainerOk returns a tuple with the SorTransparentContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorTransparentContainer

`func (o *SorInfo1) SetSorTransparentContainer(v string)`

SetSorTransparentContainer sets SorTransparentContainer field to given value.

### HasSorTransparentContainer

`func (o *SorInfo1) HasSorTransparentContainer() bool`

HasSorTransparentContainer returns a boolean if a field has been set.

### GetSorCmci

`func (o *SorInfo1) GetSorCmci() string`

GetSorCmci returns the SorCmci field if non-nil, zero value otherwise.

### GetSorCmciOk

`func (o *SorInfo1) GetSorCmciOk() (*string, bool)`

GetSorCmciOk returns a tuple with the SorCmci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorCmci

`func (o *SorInfo1) SetSorCmci(v string)`

SetSorCmci sets SorCmci field to given value.

### HasSorCmci

`func (o *SorInfo1) HasSorCmci() bool`

HasSorCmci returns a boolean if a field has been set.

### GetStoreSorCmciInMe

`func (o *SorInfo1) GetStoreSorCmciInMe() bool`

GetStoreSorCmciInMe returns the StoreSorCmciInMe field if non-nil, zero value otherwise.

### GetStoreSorCmciInMeOk

`func (o *SorInfo1) GetStoreSorCmciInMeOk() (*bool, bool)`

GetStoreSorCmciInMeOk returns a tuple with the StoreSorCmciInMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSorCmciInMe

`func (o *SorInfo1) SetStoreSorCmciInMe(v bool)`

SetStoreSorCmciInMe sets StoreSorCmciInMe field to given value.

### HasStoreSorCmciInMe

`func (o *SorInfo1) HasStoreSorCmciInMe() bool`

HasStoreSorCmciInMe returns a boolean if a field has been set.

### GetUsimSupportOfSorCmci

`func (o *SorInfo1) GetUsimSupportOfSorCmci() bool`

GetUsimSupportOfSorCmci returns the UsimSupportOfSorCmci field if non-nil, zero value otherwise.

### GetUsimSupportOfSorCmciOk

`func (o *SorInfo1) GetUsimSupportOfSorCmciOk() (*bool, bool)`

GetUsimSupportOfSorCmciOk returns a tuple with the UsimSupportOfSorCmci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsimSupportOfSorCmci

`func (o *SorInfo1) SetUsimSupportOfSorCmci(v bool)`

SetUsimSupportOfSorCmci sets UsimSupportOfSorCmci field to given value.

### HasUsimSupportOfSorCmci

`func (o *SorInfo1) HasUsimSupportOfSorCmci() bool`

HasUsimSupportOfSorCmci returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


