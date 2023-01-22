# PcfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**DnnList** | Pointer to **[]string** |  | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**RxDiamHost** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**RxDiamRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**V2xSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**ProseSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**ProseCapability** | Pointer to [**ProSeCapability**](ProSeCapability.md) |  | [optional] 
**V2xCapability** | Pointer to [**V2xCapability**](V2xCapability.md) |  | [optional] 

## Methods

### NewPcfInfo

`func NewPcfInfo() *PcfInfo`

NewPcfInfo instantiates a new PcfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfInfoWithDefaults

`func NewPcfInfoWithDefaults() *PcfInfo`

NewPcfInfoWithDefaults instantiates a new PcfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *PcfInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PcfInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PcfInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PcfInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDnnList

`func (o *PcfInfo) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *PcfInfo) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *PcfInfo) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *PcfInfo) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetSupiRanges

`func (o *PcfInfo) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *PcfInfo) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *PcfInfo) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *PcfInfo) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *PcfInfo) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *PcfInfo) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *PcfInfo) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *PcfInfo) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetRxDiamHost

`func (o *PcfInfo) GetRxDiamHost() string`

GetRxDiamHost returns the RxDiamHost field if non-nil, zero value otherwise.

### GetRxDiamHostOk

`func (o *PcfInfo) GetRxDiamHostOk() (*string, bool)`

GetRxDiamHostOk returns a tuple with the RxDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamHost

`func (o *PcfInfo) SetRxDiamHost(v string)`

SetRxDiamHost sets RxDiamHost field to given value.

### HasRxDiamHost

`func (o *PcfInfo) HasRxDiamHost() bool`

HasRxDiamHost returns a boolean if a field has been set.

### GetRxDiamRealm

`func (o *PcfInfo) GetRxDiamRealm() string`

GetRxDiamRealm returns the RxDiamRealm field if non-nil, zero value otherwise.

### GetRxDiamRealmOk

`func (o *PcfInfo) GetRxDiamRealmOk() (*string, bool)`

GetRxDiamRealmOk returns a tuple with the RxDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamRealm

`func (o *PcfInfo) SetRxDiamRealm(v string)`

SetRxDiamRealm sets RxDiamRealm field to given value.

### HasRxDiamRealm

`func (o *PcfInfo) HasRxDiamRealm() bool`

HasRxDiamRealm returns a boolean if a field has been set.

### GetV2xSupportInd

`func (o *PcfInfo) GetV2xSupportInd() bool`

GetV2xSupportInd returns the V2xSupportInd field if non-nil, zero value otherwise.

### GetV2xSupportIndOk

`func (o *PcfInfo) GetV2xSupportIndOk() (*bool, bool)`

GetV2xSupportIndOk returns a tuple with the V2xSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xSupportInd

`func (o *PcfInfo) SetV2xSupportInd(v bool)`

SetV2xSupportInd sets V2xSupportInd field to given value.

### HasV2xSupportInd

`func (o *PcfInfo) HasV2xSupportInd() bool`

HasV2xSupportInd returns a boolean if a field has been set.

### GetProseSupportInd

`func (o *PcfInfo) GetProseSupportInd() bool`

GetProseSupportInd returns the ProseSupportInd field if non-nil, zero value otherwise.

### GetProseSupportIndOk

`func (o *PcfInfo) GetProseSupportIndOk() (*bool, bool)`

GetProseSupportIndOk returns a tuple with the ProseSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseSupportInd

`func (o *PcfInfo) SetProseSupportInd(v bool)`

SetProseSupportInd sets ProseSupportInd field to given value.

### HasProseSupportInd

`func (o *PcfInfo) HasProseSupportInd() bool`

HasProseSupportInd returns a boolean if a field has been set.

### GetProseCapability

`func (o *PcfInfo) GetProseCapability() ProSeCapability`

GetProseCapability returns the ProseCapability field if non-nil, zero value otherwise.

### GetProseCapabilityOk

`func (o *PcfInfo) GetProseCapabilityOk() (*ProSeCapability, bool)`

GetProseCapabilityOk returns a tuple with the ProseCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseCapability

`func (o *PcfInfo) SetProseCapability(v ProSeCapability)`

SetProseCapability sets ProseCapability field to given value.

### HasProseCapability

`func (o *PcfInfo) HasProseCapability() bool`

HasProseCapability returns a boolean if a field has been set.

### GetV2xCapability

`func (o *PcfInfo) GetV2xCapability() V2xCapability`

GetV2xCapability returns the V2xCapability field if non-nil, zero value otherwise.

### GetV2xCapabilityOk

`func (o *PcfInfo) GetV2xCapabilityOk() (*V2xCapability, bool)`

GetV2xCapabilityOk returns a tuple with the V2xCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xCapability

`func (o *PcfInfo) SetV2xCapability(v V2xCapability)`

SetV2xCapability sets V2xCapability field to given value.

### HasV2xCapability

`func (o *PcfInfo) HasV2xCapability() bool`

HasV2xCapability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


