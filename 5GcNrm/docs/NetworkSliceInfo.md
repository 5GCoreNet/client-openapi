# NetworkSliceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNSSAI** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**CNSIId** | Pointer to **string** | CNSI Id is defined in TS 29.531, only for Core Network | [optional] 
**NetworkSliceRef** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkSliceInfo

`func NewNetworkSliceInfo() *NetworkSliceInfo`

NewNetworkSliceInfo instantiates a new NetworkSliceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSliceInfoWithDefaults

`func NewNetworkSliceInfoWithDefaults() *NetworkSliceInfo`

NewNetworkSliceInfoWithDefaults instantiates a new NetworkSliceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNSSAI

`func (o *NetworkSliceInfo) GetSNSSAI() Snssai`

GetSNSSAI returns the SNSSAI field if non-nil, zero value otherwise.

### GetSNSSAIOk

`func (o *NetworkSliceInfo) GetSNSSAIOk() (*Snssai, bool)`

GetSNSSAIOk returns a tuple with the SNSSAI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNSSAI

`func (o *NetworkSliceInfo) SetSNSSAI(v Snssai)`

SetSNSSAI sets SNSSAI field to given value.

### HasSNSSAI

`func (o *NetworkSliceInfo) HasSNSSAI() bool`

HasSNSSAI returns a boolean if a field has been set.

### GetCNSIId

`func (o *NetworkSliceInfo) GetCNSIId() string`

GetCNSIId returns the CNSIId field if non-nil, zero value otherwise.

### GetCNSIIdOk

`func (o *NetworkSliceInfo) GetCNSIIdOk() (*string, bool)`

GetCNSIIdOk returns a tuple with the CNSIId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCNSIId

`func (o *NetworkSliceInfo) SetCNSIId(v string)`

SetCNSIId sets CNSIId field to given value.

### HasCNSIId

`func (o *NetworkSliceInfo) HasCNSIId() bool`

HasCNSIId returns a boolean if a field has been set.

### GetNetworkSliceRef

`func (o *NetworkSliceInfo) GetNetworkSliceRef() []string`

GetNetworkSliceRef returns the NetworkSliceRef field if non-nil, zero value otherwise.

### GetNetworkSliceRefOk

`func (o *NetworkSliceInfo) GetNetworkSliceRefOk() (*[]string, bool)`

GetNetworkSliceRefOk returns a tuple with the NetworkSliceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceRef

`func (o *NetworkSliceInfo) SetNetworkSliceRef(v []string)`

SetNetworkSliceRef sets NetworkSliceRef field to given value.

### HasNetworkSliceRef

`func (o *NetworkSliceInfo) HasNetworkSliceRef() bool`

HasNetworkSliceRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


