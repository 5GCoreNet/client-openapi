# NrfInfoServedUdmInfoValue

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

### NewNrfInfoServedUdmInfoValue

`func NewNrfInfoServedUdmInfoValue() *NrfInfoServedUdmInfoValue`

NewNrfInfoServedUdmInfoValue instantiates a new NrfInfoServedUdmInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedUdmInfoValueWithDefaults

`func NewNrfInfoServedUdmInfoValueWithDefaults() *NrfInfoServedUdmInfoValue`

NewNrfInfoServedUdmInfoValueWithDefaults instantiates a new NrfInfoServedUdmInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *NrfInfoServedUdmInfoValue) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NrfInfoServedUdmInfoValue) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NrfInfoServedUdmInfoValue) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NrfInfoServedUdmInfoValue) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *NrfInfoServedUdmInfoValue) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *NrfInfoServedUdmInfoValue) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *NrfInfoServedUdmInfoValue) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *NrfInfoServedUdmInfoValue) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *NrfInfoServedUdmInfoValue) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *NrfInfoServedUdmInfoValue) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *NrfInfoServedUdmInfoValue) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *NrfInfoServedUdmInfoValue) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdmInfoValue) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *NrfInfoServedUdmInfoValue) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdmInfoValue) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdmInfoValue) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetRoutingIndicators

`func (o *NrfInfoServedUdmInfoValue) GetRoutingIndicators() []string`

GetRoutingIndicators returns the RoutingIndicators field if non-nil, zero value otherwise.

### GetRoutingIndicatorsOk

`func (o *NrfInfoServedUdmInfoValue) GetRoutingIndicatorsOk() (*[]string, bool)`

GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicators

`func (o *NrfInfoServedUdmInfoValue) SetRoutingIndicators(v []string)`

SetRoutingIndicators sets RoutingIndicators field to given value.

### HasRoutingIndicators

`func (o *NrfInfoServedUdmInfoValue) HasRoutingIndicators() bool`

HasRoutingIndicators returns a boolean if a field has been set.

### GetInternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdmInfoValue) GetInternalGroupIdentifiersRanges() []InternalGroupIdRange`

GetInternalGroupIdentifiersRanges returns the InternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetInternalGroupIdentifiersRangesOk

`func (o *NrfInfoServedUdmInfoValue) GetInternalGroupIdentifiersRangesOk() (*[]InternalGroupIdRange, bool)`

GetInternalGroupIdentifiersRangesOk returns a tuple with the InternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdmInfoValue) SetInternalGroupIdentifiersRanges(v []InternalGroupIdRange)`

SetInternalGroupIdentifiersRanges sets InternalGroupIdentifiersRanges field to given value.

### HasInternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdmInfoValue) HasInternalGroupIdentifiersRanges() bool`

HasInternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetSuciInfos

`func (o *NrfInfoServedUdmInfoValue) GetSuciInfos() []SuciInfo`

GetSuciInfos returns the SuciInfos field if non-nil, zero value otherwise.

### GetSuciInfosOk

`func (o *NrfInfoServedUdmInfoValue) GetSuciInfosOk() (*[]SuciInfo, bool)`

GetSuciInfosOk returns a tuple with the SuciInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuciInfos

`func (o *NrfInfoServedUdmInfoValue) SetSuciInfos(v []SuciInfo)`

SetSuciInfos sets SuciInfos field to given value.

### HasSuciInfos

`func (o *NrfInfoServedUdmInfoValue) HasSuciInfos() bool`

HasSuciInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


