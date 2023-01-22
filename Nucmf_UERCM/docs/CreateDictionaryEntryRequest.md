# CreateDictionaryEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**DicEntryCreateData**](DicEntryCreateData.md) |  | [optional] 
**BinaryDataUeRadioCapability5GS** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUeRadioCapabilityEPS** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUeRadioCap5GSForPaging** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUeRadioCapEPSForPaging** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewCreateDictionaryEntryRequest

`func NewCreateDictionaryEntryRequest() *CreateDictionaryEntryRequest`

NewCreateDictionaryEntryRequest instantiates a new CreateDictionaryEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDictionaryEntryRequestWithDefaults

`func NewCreateDictionaryEntryRequestWithDefaults() *CreateDictionaryEntryRequest`

NewCreateDictionaryEntryRequestWithDefaults instantiates a new CreateDictionaryEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *CreateDictionaryEntryRequest) GetJsonData() DicEntryCreateData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *CreateDictionaryEntryRequest) GetJsonDataOk() (*DicEntryCreateData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *CreateDictionaryEntryRequest) SetJsonData(v DicEntryCreateData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *CreateDictionaryEntryRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataUeRadioCapability5GS

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapability5GS() *os.File`

GetBinaryDataUeRadioCapability5GS returns the BinaryDataUeRadioCapability5GS field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCapability5GSOk

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapability5GSOk() (**os.File, bool)`

GetBinaryDataUeRadioCapability5GSOk returns a tuple with the BinaryDataUeRadioCapability5GS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCapability5GS

`func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCapability5GS(v *os.File)`

SetBinaryDataUeRadioCapability5GS sets BinaryDataUeRadioCapability5GS field to given value.

### HasBinaryDataUeRadioCapability5GS

`func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCapability5GS() bool`

HasBinaryDataUeRadioCapability5GS returns a boolean if a field has been set.

### GetBinaryDataUeRadioCapabilityEPS

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapabilityEPS() *os.File`

GetBinaryDataUeRadioCapabilityEPS returns the BinaryDataUeRadioCapabilityEPS field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCapabilityEPSOk

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapabilityEPSOk() (**os.File, bool)`

GetBinaryDataUeRadioCapabilityEPSOk returns a tuple with the BinaryDataUeRadioCapabilityEPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCapabilityEPS

`func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCapabilityEPS(v *os.File)`

SetBinaryDataUeRadioCapabilityEPS sets BinaryDataUeRadioCapabilityEPS field to given value.

### HasBinaryDataUeRadioCapabilityEPS

`func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCapabilityEPS() bool`

HasBinaryDataUeRadioCapabilityEPS returns a boolean if a field has been set.

### GetBinaryDataUeRadioCap5GSForPaging

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCap5GSForPaging() *os.File`

GetBinaryDataUeRadioCap5GSForPaging returns the BinaryDataUeRadioCap5GSForPaging field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCap5GSForPagingOk

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCap5GSForPagingOk() (**os.File, bool)`

GetBinaryDataUeRadioCap5GSForPagingOk returns a tuple with the BinaryDataUeRadioCap5GSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCap5GSForPaging

`func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCap5GSForPaging(v *os.File)`

SetBinaryDataUeRadioCap5GSForPaging sets BinaryDataUeRadioCap5GSForPaging field to given value.

### HasBinaryDataUeRadioCap5GSForPaging

`func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCap5GSForPaging() bool`

HasBinaryDataUeRadioCap5GSForPaging returns a boolean if a field has been set.

### GetBinaryDataUeRadioCapEPSForPaging

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapEPSForPaging() *os.File`

GetBinaryDataUeRadioCapEPSForPaging returns the BinaryDataUeRadioCapEPSForPaging field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCapEPSForPagingOk

`func (o *CreateDictionaryEntryRequest) GetBinaryDataUeRadioCapEPSForPagingOk() (**os.File, bool)`

GetBinaryDataUeRadioCapEPSForPagingOk returns a tuple with the BinaryDataUeRadioCapEPSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCapEPSForPaging

`func (o *CreateDictionaryEntryRequest) SetBinaryDataUeRadioCapEPSForPaging(v *os.File)`

SetBinaryDataUeRadioCapEPSForPaging sets BinaryDataUeRadioCapEPSForPaging field to given value.

### HasBinaryDataUeRadioCapEPSForPaging

`func (o *CreateDictionaryEntryRequest) HasBinaryDataUeRadioCapEPSForPaging() bool`

HasBinaryDataUeRadioCapEPSForPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


