# SnssaiUpfInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssai** | [**ExtSnssai**](ExtSnssai.md) |  | 
**DnnUpfInfoList** | [**[]DnnUpfInfoItem**](DnnUpfInfoItem.md) |  | 
**RedundantTransport** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSnssaiUpfInfoItem

`func NewSnssaiUpfInfoItem(sNssai ExtSnssai, dnnUpfInfoList []DnnUpfInfoItem, ) *SnssaiUpfInfoItem`

NewSnssaiUpfInfoItem instantiates a new SnssaiUpfInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnssaiUpfInfoItemWithDefaults

`func NewSnssaiUpfInfoItemWithDefaults() *SnssaiUpfInfoItem`

NewSnssaiUpfInfoItemWithDefaults instantiates a new SnssaiUpfInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssai

`func (o *SnssaiUpfInfoItem) GetSNssai() ExtSnssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SnssaiUpfInfoItem) GetSNssaiOk() (*ExtSnssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SnssaiUpfInfoItem) SetSNssai(v ExtSnssai)`

SetSNssai sets SNssai field to given value.


### GetDnnUpfInfoList

`func (o *SnssaiUpfInfoItem) GetDnnUpfInfoList() []DnnUpfInfoItem`

GetDnnUpfInfoList returns the DnnUpfInfoList field if non-nil, zero value otherwise.

### GetDnnUpfInfoListOk

`func (o *SnssaiUpfInfoItem) GetDnnUpfInfoListOk() (*[]DnnUpfInfoItem, bool)`

GetDnnUpfInfoListOk returns a tuple with the DnnUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnUpfInfoList

`func (o *SnssaiUpfInfoItem) SetDnnUpfInfoList(v []DnnUpfInfoItem)`

SetDnnUpfInfoList sets DnnUpfInfoList field to given value.


### GetRedundantTransport

`func (o *SnssaiUpfInfoItem) GetRedundantTransport() bool`

GetRedundantTransport returns the RedundantTransport field if non-nil, zero value otherwise.

### GetRedundantTransportOk

`func (o *SnssaiUpfInfoItem) GetRedundantTransportOk() (*bool, bool)`

GetRedundantTransportOk returns a tuple with the RedundantTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantTransport

`func (o *SnssaiUpfInfoItem) SetRedundantTransport(v bool)`

SetRedundantTransport sets RedundantTransport field to given value.

### HasRedundantTransport

`func (o *SnssaiUpfInfoItem) HasRedundantTransport() bool`

HasRedundantTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


