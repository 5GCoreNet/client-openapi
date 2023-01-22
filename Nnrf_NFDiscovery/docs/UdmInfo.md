# UdmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**RoutingIndicators** | Pointer to **[]string** |  | [optional] 
**InternalGroupIdentifiersRanges** | Pointer to [**[]InternalGroupIdRange**](InternalGroupIdRange.md) |  | [optional] 
**SuciInfos** | Pointer to [**[]SuciInfo**](SuciInfo.md) |  | [optional] 

## Methods

### NewUdmInfo

`func NewUdmInfo() *UdmInfo`

NewUdmInfo instantiates a new UdmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdmInfoWithDefaults

`func NewUdmInfoWithDefaults() *UdmInfo`

NewUdmInfoWithDefaults instantiates a new UdmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UdmInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UdmInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UdmInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UdmInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *UdmInfo) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *UdmInfo) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *UdmInfo) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *UdmInfo) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *UdmInfo) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *UdmInfo) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *UdmInfo) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *UdmInfo) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *UdmInfo) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *UdmInfo) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *UdmInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *UdmInfo) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetRoutingIndicators

`func (o *UdmInfo) GetRoutingIndicators() []string`

GetRoutingIndicators returns the RoutingIndicators field if non-nil, zero value otherwise.

### GetRoutingIndicatorsOk

`func (o *UdmInfo) GetRoutingIndicatorsOk() (*[]string, bool)`

GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicators

`func (o *UdmInfo) SetRoutingIndicators(v []string)`

SetRoutingIndicators sets RoutingIndicators field to given value.

### HasRoutingIndicators

`func (o *UdmInfo) HasRoutingIndicators() bool`

HasRoutingIndicators returns a boolean if a field has been set.

### GetInternalGroupIdentifiersRanges

`func (o *UdmInfo) GetInternalGroupIdentifiersRanges() []InternalGroupIdRange`

GetInternalGroupIdentifiersRanges returns the InternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetInternalGroupIdentifiersRangesOk

`func (o *UdmInfo) GetInternalGroupIdentifiersRangesOk() (*[]InternalGroupIdRange, bool)`

GetInternalGroupIdentifiersRangesOk returns a tuple with the InternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIdentifiersRanges

`func (o *UdmInfo) SetInternalGroupIdentifiersRanges(v []InternalGroupIdRange)`

SetInternalGroupIdentifiersRanges sets InternalGroupIdentifiersRanges field to given value.

### HasInternalGroupIdentifiersRanges

`func (o *UdmInfo) HasInternalGroupIdentifiersRanges() bool`

HasInternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetSuciInfos

`func (o *UdmInfo) GetSuciInfos() []SuciInfo`

GetSuciInfos returns the SuciInfos field if non-nil, zero value otherwise.

### GetSuciInfosOk

`func (o *UdmInfo) GetSuciInfosOk() (*[]SuciInfo, bool)`

GetSuciInfosOk returns a tuple with the SuciInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuciInfos

`func (o *UdmInfo) SetSuciInfos(v []SuciInfo)`

SetSuciInfos sets SuciInfos field to given value.

### HasSuciInfos

`func (o *UdmInfo) HasSuciInfos() bool`

HasSuciInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


