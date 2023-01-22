# AuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupiOrSuci** | **string** | String identifying a SUPI or a SUCI. | 
**ServingNetworkName** | **string** |  | 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**UdmGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**RoutingIndicator** | Pointer to **string** |  | [optional] 
**CellCagInfo** | Pointer to **[]string** |  | [optional] 
**N5gcInd** | Pointer to **bool** |  | [optional] [default to false]
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**NswoInd** | Pointer to **bool** |  | [optional] [default to false]
**DisasterRoamingInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAuthenticationInfo

`func NewAuthenticationInfo(supiOrSuci string, servingNetworkName string, ) *AuthenticationInfo`

NewAuthenticationInfo instantiates a new AuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationInfoWithDefaults

`func NewAuthenticationInfoWithDefaults() *AuthenticationInfo`

NewAuthenticationInfoWithDefaults instantiates a new AuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupiOrSuci

`func (o *AuthenticationInfo) GetSupiOrSuci() string`

GetSupiOrSuci returns the SupiOrSuci field if non-nil, zero value otherwise.

### GetSupiOrSuciOk

`func (o *AuthenticationInfo) GetSupiOrSuciOk() (*string, bool)`

GetSupiOrSuciOk returns a tuple with the SupiOrSuci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiOrSuci

`func (o *AuthenticationInfo) SetSupiOrSuci(v string)`

SetSupiOrSuci sets SupiOrSuci field to given value.


### GetServingNetworkName

`func (o *AuthenticationInfo) GetServingNetworkName() string`

GetServingNetworkName returns the ServingNetworkName field if non-nil, zero value otherwise.

### GetServingNetworkNameOk

`func (o *AuthenticationInfo) GetServingNetworkNameOk() (*string, bool)`

GetServingNetworkNameOk returns a tuple with the ServingNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetworkName

`func (o *AuthenticationInfo) SetServingNetworkName(v string)`

SetServingNetworkName sets ServingNetworkName field to given value.


### GetResynchronizationInfo

`func (o *AuthenticationInfo) GetResynchronizationInfo() ResynchronizationInfo`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *AuthenticationInfo) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *AuthenticationInfo) SetResynchronizationInfo(v ResynchronizationInfo)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *AuthenticationInfo) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.

### GetPei

`func (o *AuthenticationInfo) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *AuthenticationInfo) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *AuthenticationInfo) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *AuthenticationInfo) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetTraceData

`func (o *AuthenticationInfo) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *AuthenticationInfo) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *AuthenticationInfo) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *AuthenticationInfo) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *AuthenticationInfo) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *AuthenticationInfo) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetUdmGroupId

`func (o *AuthenticationInfo) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *AuthenticationInfo) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *AuthenticationInfo) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *AuthenticationInfo) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.

### GetRoutingIndicator

`func (o *AuthenticationInfo) GetRoutingIndicator() string`

GetRoutingIndicator returns the RoutingIndicator field if non-nil, zero value otherwise.

### GetRoutingIndicatorOk

`func (o *AuthenticationInfo) GetRoutingIndicatorOk() (*string, bool)`

GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicator

`func (o *AuthenticationInfo) SetRoutingIndicator(v string)`

SetRoutingIndicator sets RoutingIndicator field to given value.

### HasRoutingIndicator

`func (o *AuthenticationInfo) HasRoutingIndicator() bool`

HasRoutingIndicator returns a boolean if a field has been set.

### GetCellCagInfo

`func (o *AuthenticationInfo) GetCellCagInfo() []string`

GetCellCagInfo returns the CellCagInfo field if non-nil, zero value otherwise.

### GetCellCagInfoOk

`func (o *AuthenticationInfo) GetCellCagInfoOk() (*[]string, bool)`

GetCellCagInfoOk returns a tuple with the CellCagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellCagInfo

`func (o *AuthenticationInfo) SetCellCagInfo(v []string)`

SetCellCagInfo sets CellCagInfo field to given value.

### HasCellCagInfo

`func (o *AuthenticationInfo) HasCellCagInfo() bool`

HasCellCagInfo returns a boolean if a field has been set.

### GetN5gcInd

`func (o *AuthenticationInfo) GetN5gcInd() bool`

GetN5gcInd returns the N5gcInd field if non-nil, zero value otherwise.

### GetN5gcIndOk

`func (o *AuthenticationInfo) GetN5gcIndOk() (*bool, bool)`

GetN5gcIndOk returns a tuple with the N5gcInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN5gcInd

`func (o *AuthenticationInfo) SetN5gcInd(v bool)`

SetN5gcInd sets N5gcInd field to given value.

### HasN5gcInd

`func (o *AuthenticationInfo) HasN5gcInd() bool`

HasN5gcInd returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *AuthenticationInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *AuthenticationInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *AuthenticationInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *AuthenticationInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetNswoInd

`func (o *AuthenticationInfo) GetNswoInd() bool`

GetNswoInd returns the NswoInd field if non-nil, zero value otherwise.

### GetNswoIndOk

`func (o *AuthenticationInfo) GetNswoIndOk() (*bool, bool)`

GetNswoIndOk returns a tuple with the NswoInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNswoInd

`func (o *AuthenticationInfo) SetNswoInd(v bool)`

SetNswoInd sets NswoInd field to given value.

### HasNswoInd

`func (o *AuthenticationInfo) HasNswoInd() bool`

HasNswoInd returns a boolean if a field has been set.

### GetDisasterRoamingInd

`func (o *AuthenticationInfo) GetDisasterRoamingInd() bool`

GetDisasterRoamingInd returns the DisasterRoamingInd field if non-nil, zero value otherwise.

### GetDisasterRoamingIndOk

`func (o *AuthenticationInfo) GetDisasterRoamingIndOk() (*bool, bool)`

GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRoamingInd

`func (o *AuthenticationInfo) SetDisasterRoamingInd(v bool)`

SetDisasterRoamingInd sets DisasterRoamingInd field to given value.

### HasDisasterRoamingInd

`func (o *AuthenticationInfo) HasDisasterRoamingInd() bool`

HasDisasterRoamingInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


