# DicEntryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DicEntryId** | Pointer to **int32** |  | [optional] 
**TypeAllocationCode** | **string** | Type Allocation Code (TAC) of the UE, comprising the initial eight-digit portion of the 15-digit IMEI and 16-digit IMEISV codes. See clause 6.2 of 3GPP TS 23.003.  | 
**PlmnAssiUeRadioCapId** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**ManAssiUeRadioCapId** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**UeRadioCapability5GS** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeRadioCapabilityEPS** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeRadioCap5GSForPaging** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UeRadioCapEPSForPaging** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 

## Methods

### NewDicEntryData

`func NewDicEntryData(typeAllocationCode string, ) *DicEntryData`

NewDicEntryData instantiates a new DicEntryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDicEntryDataWithDefaults

`func NewDicEntryDataWithDefaults() *DicEntryData`

NewDicEntryDataWithDefaults instantiates a new DicEntryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDicEntryId

`func (o *DicEntryData) GetDicEntryId() int32`

GetDicEntryId returns the DicEntryId field if non-nil, zero value otherwise.

### GetDicEntryIdOk

`func (o *DicEntryData) GetDicEntryIdOk() (*int32, bool)`

GetDicEntryIdOk returns a tuple with the DicEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDicEntryId

`func (o *DicEntryData) SetDicEntryId(v int32)`

SetDicEntryId sets DicEntryId field to given value.

### HasDicEntryId

`func (o *DicEntryData) HasDicEntryId() bool`

HasDicEntryId returns a boolean if a field has been set.

### GetTypeAllocationCode

`func (o *DicEntryData) GetTypeAllocationCode() string`

GetTypeAllocationCode returns the TypeAllocationCode field if non-nil, zero value otherwise.

### GetTypeAllocationCodeOk

`func (o *DicEntryData) GetTypeAllocationCodeOk() (*string, bool)`

GetTypeAllocationCodeOk returns a tuple with the TypeAllocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAllocationCode

`func (o *DicEntryData) SetTypeAllocationCode(v string)`

SetTypeAllocationCode sets TypeAllocationCode field to given value.


### GetPlmnAssiUeRadioCapId

`func (o *DicEntryData) GetPlmnAssiUeRadioCapId() string`

GetPlmnAssiUeRadioCapId returns the PlmnAssiUeRadioCapId field if non-nil, zero value otherwise.

### GetPlmnAssiUeRadioCapIdOk

`func (o *DicEntryData) GetPlmnAssiUeRadioCapIdOk() (*string, bool)`

GetPlmnAssiUeRadioCapIdOk returns a tuple with the PlmnAssiUeRadioCapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnAssiUeRadioCapId

`func (o *DicEntryData) SetPlmnAssiUeRadioCapId(v string)`

SetPlmnAssiUeRadioCapId sets PlmnAssiUeRadioCapId field to given value.

### HasPlmnAssiUeRadioCapId

`func (o *DicEntryData) HasPlmnAssiUeRadioCapId() bool`

HasPlmnAssiUeRadioCapId returns a boolean if a field has been set.

### GetManAssiUeRadioCapId

`func (o *DicEntryData) GetManAssiUeRadioCapId() string`

GetManAssiUeRadioCapId returns the ManAssiUeRadioCapId field if non-nil, zero value otherwise.

### GetManAssiUeRadioCapIdOk

`func (o *DicEntryData) GetManAssiUeRadioCapIdOk() (*string, bool)`

GetManAssiUeRadioCapIdOk returns a tuple with the ManAssiUeRadioCapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManAssiUeRadioCapId

`func (o *DicEntryData) SetManAssiUeRadioCapId(v string)`

SetManAssiUeRadioCapId sets ManAssiUeRadioCapId field to given value.

### HasManAssiUeRadioCapId

`func (o *DicEntryData) HasManAssiUeRadioCapId() bool`

HasManAssiUeRadioCapId returns a boolean if a field has been set.

### GetUeRadioCapability5GS

`func (o *DicEntryData) GetUeRadioCapability5GS() RefToBinaryData`

GetUeRadioCapability5GS returns the UeRadioCapability5GS field if non-nil, zero value otherwise.

### GetUeRadioCapability5GSOk

`func (o *DicEntryData) GetUeRadioCapability5GSOk() (*RefToBinaryData, bool)`

GetUeRadioCapability5GSOk returns a tuple with the UeRadioCapability5GS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapability5GS

`func (o *DicEntryData) SetUeRadioCapability5GS(v RefToBinaryData)`

SetUeRadioCapability5GS sets UeRadioCapability5GS field to given value.

### HasUeRadioCapability5GS

`func (o *DicEntryData) HasUeRadioCapability5GS() bool`

HasUeRadioCapability5GS returns a boolean if a field has been set.

### GetUeRadioCapabilityEPS

`func (o *DicEntryData) GetUeRadioCapabilityEPS() RefToBinaryData`

GetUeRadioCapabilityEPS returns the UeRadioCapabilityEPS field if non-nil, zero value otherwise.

### GetUeRadioCapabilityEPSOk

`func (o *DicEntryData) GetUeRadioCapabilityEPSOk() (*RefToBinaryData, bool)`

GetUeRadioCapabilityEPSOk returns a tuple with the UeRadioCapabilityEPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapabilityEPS

`func (o *DicEntryData) SetUeRadioCapabilityEPS(v RefToBinaryData)`

SetUeRadioCapabilityEPS sets UeRadioCapabilityEPS field to given value.

### HasUeRadioCapabilityEPS

`func (o *DicEntryData) HasUeRadioCapabilityEPS() bool`

HasUeRadioCapabilityEPS returns a boolean if a field has been set.

### GetUeRadioCap5GSForPaging

`func (o *DicEntryData) GetUeRadioCap5GSForPaging() RefToBinaryData`

GetUeRadioCap5GSForPaging returns the UeRadioCap5GSForPaging field if non-nil, zero value otherwise.

### GetUeRadioCap5GSForPagingOk

`func (o *DicEntryData) GetUeRadioCap5GSForPagingOk() (*RefToBinaryData, bool)`

GetUeRadioCap5GSForPagingOk returns a tuple with the UeRadioCap5GSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCap5GSForPaging

`func (o *DicEntryData) SetUeRadioCap5GSForPaging(v RefToBinaryData)`

SetUeRadioCap5GSForPaging sets UeRadioCap5GSForPaging field to given value.

### HasUeRadioCap5GSForPaging

`func (o *DicEntryData) HasUeRadioCap5GSForPaging() bool`

HasUeRadioCap5GSForPaging returns a boolean if a field has been set.

### GetUeRadioCapEPSForPaging

`func (o *DicEntryData) GetUeRadioCapEPSForPaging() RefToBinaryData`

GetUeRadioCapEPSForPaging returns the UeRadioCapEPSForPaging field if non-nil, zero value otherwise.

### GetUeRadioCapEPSForPagingOk

`func (o *DicEntryData) GetUeRadioCapEPSForPagingOk() (*RefToBinaryData, bool)`

GetUeRadioCapEPSForPagingOk returns a tuple with the UeRadioCapEPSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRadioCapEPSForPaging

`func (o *DicEntryData) SetUeRadioCapEPSForPaging(v RefToBinaryData)`

SetUeRadioCapEPSForPaging sets UeRadioCapEPSForPaging field to given value.

### HasUeRadioCapEPSForPaging

`func (o *DicEntryData) HasUeRadioCapEPSForPaging() bool`

HasUeRadioCapEPSForPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


