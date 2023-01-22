# DnnUpfInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**DnaiList** | Pointer to **[]string** |  | [optional] 
**PduSessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) |  | [optional] 
**Ipv4AddressRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**Ipv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 
**Ipv4IndexList** | Pointer to [**[]IpIndex**](IpIndex.md) |  | [optional] 
**Ipv6IndexList** | Pointer to [**[]IpIndex**](IpIndex.md) |  | [optional] 
**DnaiNwInstanceList** | Pointer to **map[string]string** | Map of network instance per DNAI for the DNN, where the key of the map is the DNAI. When present, the value of each entry of the map shall contain a N6 network instance that is configured for the DNAI indicated by the key.  | [optional] 

## Methods

### NewDnnUpfInfoItem

`func NewDnnUpfInfoItem(dnn string, ) *DnnUpfInfoItem`

NewDnnUpfInfoItem instantiates a new DnnUpfInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnUpfInfoItemWithDefaults

`func NewDnnUpfInfoItemWithDefaults() *DnnUpfInfoItem`

NewDnnUpfInfoItemWithDefaults instantiates a new DnnUpfInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *DnnUpfInfoItem) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnnUpfInfoItem) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnnUpfInfoItem) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetDnaiList

`func (o *DnnUpfInfoItem) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *DnnUpfInfoItem) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *DnnUpfInfoItem) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *DnnUpfInfoItem) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetPduSessionTypes

`func (o *DnnUpfInfoItem) GetPduSessionTypes() []PduSessionType`

GetPduSessionTypes returns the PduSessionTypes field if non-nil, zero value otherwise.

### GetPduSessionTypesOk

`func (o *DnnUpfInfoItem) GetPduSessionTypesOk() (*[]PduSessionType, bool)`

GetPduSessionTypesOk returns a tuple with the PduSessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionTypes

`func (o *DnnUpfInfoItem) SetPduSessionTypes(v []PduSessionType)`

SetPduSessionTypes sets PduSessionTypes field to given value.

### HasPduSessionTypes

`func (o *DnnUpfInfoItem) HasPduSessionTypes() bool`

HasPduSessionTypes returns a boolean if a field has been set.

### GetIpv4AddressRanges

`func (o *DnnUpfInfoItem) GetIpv4AddressRanges() []Ipv4AddressRange`

GetIpv4AddressRanges returns the Ipv4AddressRanges field if non-nil, zero value otherwise.

### GetIpv4AddressRangesOk

`func (o *DnnUpfInfoItem) GetIpv4AddressRangesOk() (*[]Ipv4AddressRange, bool)`

GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressRanges

`func (o *DnnUpfInfoItem) SetIpv4AddressRanges(v []Ipv4AddressRange)`

SetIpv4AddressRanges sets Ipv4AddressRanges field to given value.

### HasIpv4AddressRanges

`func (o *DnnUpfInfoItem) HasIpv4AddressRanges() bool`

HasIpv4AddressRanges returns a boolean if a field has been set.

### GetIpv6PrefixRanges

`func (o *DnnUpfInfoItem) GetIpv6PrefixRanges() []Ipv6PrefixRange`

GetIpv6PrefixRanges returns the Ipv6PrefixRanges field if non-nil, zero value otherwise.

### GetIpv6PrefixRangesOk

`func (o *DnnUpfInfoItem) GetIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixRanges

`func (o *DnnUpfInfoItem) SetIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetIpv6PrefixRanges sets Ipv6PrefixRanges field to given value.

### HasIpv6PrefixRanges

`func (o *DnnUpfInfoItem) HasIpv6PrefixRanges() bool`

HasIpv6PrefixRanges returns a boolean if a field has been set.

### GetIpv4IndexList

`func (o *DnnUpfInfoItem) GetIpv4IndexList() []IpIndex`

GetIpv4IndexList returns the Ipv4IndexList field if non-nil, zero value otherwise.

### GetIpv4IndexListOk

`func (o *DnnUpfInfoItem) GetIpv4IndexListOk() (*[]IpIndex, bool)`

GetIpv4IndexListOk returns a tuple with the Ipv4IndexList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4IndexList

`func (o *DnnUpfInfoItem) SetIpv4IndexList(v []IpIndex)`

SetIpv4IndexList sets Ipv4IndexList field to given value.

### HasIpv4IndexList

`func (o *DnnUpfInfoItem) HasIpv4IndexList() bool`

HasIpv4IndexList returns a boolean if a field has been set.

### GetIpv6IndexList

`func (o *DnnUpfInfoItem) GetIpv6IndexList() []IpIndex`

GetIpv6IndexList returns the Ipv6IndexList field if non-nil, zero value otherwise.

### GetIpv6IndexListOk

`func (o *DnnUpfInfoItem) GetIpv6IndexListOk() (*[]IpIndex, bool)`

GetIpv6IndexListOk returns a tuple with the Ipv6IndexList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6IndexList

`func (o *DnnUpfInfoItem) SetIpv6IndexList(v []IpIndex)`

SetIpv6IndexList sets Ipv6IndexList field to given value.

### HasIpv6IndexList

`func (o *DnnUpfInfoItem) HasIpv6IndexList() bool`

HasIpv6IndexList returns a boolean if a field has been set.

### GetDnaiNwInstanceList

`func (o *DnnUpfInfoItem) GetDnaiNwInstanceList() map[string]string`

GetDnaiNwInstanceList returns the DnaiNwInstanceList field if non-nil, zero value otherwise.

### GetDnaiNwInstanceListOk

`func (o *DnnUpfInfoItem) GetDnaiNwInstanceListOk() (*map[string]string, bool)`

GetDnaiNwInstanceListOk returns a tuple with the DnaiNwInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiNwInstanceList

`func (o *DnnUpfInfoItem) SetDnaiNwInstanceList(v map[string]string)`

SetDnaiNwInstanceList sets DnaiNwInstanceList field to given value.

### HasDnaiNwInstanceList

`func (o *DnnUpfInfoItem) HasDnaiNwInstanceList() bool`

HasDnaiNwInstanceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


