# DicEntryCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeAllocationCode** | **string** | Type Allocation Code (TAC) of the UE, comprising the initial eight-digit portion of the 15-digit IMEI and 16-digit IMEISV codes. See clause 6.2 of 3GPP TS 23.003.  | 
**UeRadioCapability5GS** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeRadioCapabilityEPS** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**UeRadioCap5GSForPaging** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeRadioCapEPSForPaging** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 

## Methods

### NewDicEntryCreateData

`func NewDicEntryCreateData(typeAllocationCode string, ) *DicEntryCreateData`

NewDicEntryCreateData instantiates a new DicEntryCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDicEntryCreateDataWithDefaults

`func NewDicEntryCreateDataWithDefaults() *DicEntryCreateData`

NewDicEntryCreateDataWithDefaults instantiates a new DicEntryCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeAllocationCode

`func (o *DicEntryCreateData) GetTypeAllocationCode() string`

GetTypeAllocationCode returns the TypeAllocationCode field if non-nil, zero value otherwise.

### GetTypeAllocationCodeOk

`func (o *DicEntryCreateData) GetTypeAllocationCodeOk() (*string, bool)`

GetTypeAllocationCodeOk returns a tuple with the TypeAllocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAllocationCode

`func (o *DicEntryCreateData) SetTypeAllocationCode(v string)`

SetTypeAllocationCode sets TypeAllocationCode field to given value.


### GetUeRadioCapability5GS

`func (o *DicEntryCreateData) GetUeRadioCapability5GS() RefToBinaryData`

GetUeRadioCapability5GS returns the UeRadioCapability5GS field if non-nil, zero value otherwise.

### GetUeRadioCapability5GSOk

`func (o *DicEntryCreateData) GetUeRadioCapability5GSOk() (*RefToBinaryData, bool)`

GetUeRadioCapability5GSOk returns a tuple with the UeRadioCapability5GS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapability5GS

`func (o *DicEntryCreateData) SetUeRadioCapability5GS(v RefToBinaryData)`

SetUeRadioCapability5GS sets UeRadioCapability5GS field to given value.

### HasUeRadioCapability5GS

`func (o *DicEntryCreateData) HasUeRadioCapability5GS() bool`

HasUeRadioCapability5GS returns a boolean if a field has been set.

### GetUeRadioCapabilityEPS

`func (o *DicEntryCreateData) GetUeRadioCapabilityEPS() RefToBinaryData`

GetUeRadioCapabilityEPS returns the UeRadioCapabilityEPS field if non-nil, zero value otherwise.

### GetUeRadioCapabilityEPSOk

`func (o *DicEntryCreateData) GetUeRadioCapabilityEPSOk() (*RefToBinaryData, bool)`

GetUeRadioCapabilityEPSOk returns a tuple with the UeRadioCapabilityEPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapabilityEPS

`func (o *DicEntryCreateData) SetUeRadioCapabilityEPS(v RefToBinaryData)`

SetUeRadioCapabilityEPS sets UeRadioCapabilityEPS field to given value.

### HasUeRadioCapabilityEPS

`func (o *DicEntryCreateData) HasUeRadioCapabilityEPS() bool`

HasUeRadioCapabilityEPS returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *DicEntryCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *DicEntryCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *DicEntryCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *DicEntryCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetUeRadioCap5GSForPaging

`func (o *DicEntryCreateData) GetUeRadioCap5GSForPaging() RefToBinaryData`

GetUeRadioCap5GSForPaging returns the UeRadioCap5GSForPaging field if non-nil, zero value otherwise.

### GetUeRadioCap5GSForPagingOk

`func (o *DicEntryCreateData) GetUeRadioCap5GSForPagingOk() (*RefToBinaryData, bool)`

GetUeRadioCap5GSForPagingOk returns a tuple with the UeRadioCap5GSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCap5GSForPaging

`func (o *DicEntryCreateData) SetUeRadioCap5GSForPaging(v RefToBinaryData)`

SetUeRadioCap5GSForPaging sets UeRadioCap5GSForPaging field to given value.

### HasUeRadioCap5GSForPaging

`func (o *DicEntryCreateData) HasUeRadioCap5GSForPaging() bool`

HasUeRadioCap5GSForPaging returns a boolean if a field has been set.

### GetUeRadioCapEPSForPaging

`func (o *DicEntryCreateData) GetUeRadioCapEPSForPaging() RefToBinaryData`

GetUeRadioCapEPSForPaging returns the UeRadioCapEPSForPaging field if non-nil, zero value otherwise.

### GetUeRadioCapEPSForPagingOk

`func (o *DicEntryCreateData) GetUeRadioCapEPSForPagingOk() (*RefToBinaryData, bool)`

GetUeRadioCapEPSForPagingOk returns a tuple with the UeRadioCapEPSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapEPSForPaging

`func (o *DicEntryCreateData) SetUeRadioCapEPSForPaging(v RefToBinaryData)`

SetUeRadioCapEPSForPaging sets UeRadioCapEPSForPaging field to given value.

### HasUeRadioCapEPSForPaging

`func (o *DicEntryCreateData) HasUeRadioCapEPSForPaging() bool`

HasUeRadioCapEPSForPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


