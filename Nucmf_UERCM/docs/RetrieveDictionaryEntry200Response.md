# RetrieveDictionaryEntry200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**DicEntryData**](DicEntryData.md) |  | [optional] 
**BinaryDataUeRadioCapability5GS** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUeRadioCapabilityEPS** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUeRadioCap5GSForPaging** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUeRadioCapEPSForPaging** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewRetrieveDictionaryEntry200Response

`func NewRetrieveDictionaryEntry200Response() *RetrieveDictionaryEntry200Response`

NewRetrieveDictionaryEntry200Response instantiates a new RetrieveDictionaryEntry200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveDictionaryEntry200ResponseWithDefaults

`func NewRetrieveDictionaryEntry200ResponseWithDefaults() *RetrieveDictionaryEntry200Response`

NewRetrieveDictionaryEntry200ResponseWithDefaults instantiates a new RetrieveDictionaryEntry200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *RetrieveDictionaryEntry200Response) GetJsonData() DicEntryData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *RetrieveDictionaryEntry200Response) GetJsonDataOk() (*DicEntryData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *RetrieveDictionaryEntry200Response) SetJsonData(v DicEntryData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *RetrieveDictionaryEntry200Response) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataUeRadioCapability5GS

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapability5GS() *os.File`

GetBinaryDataUeRadioCapability5GS returns the BinaryDataUeRadioCapability5GS field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCapability5GSOk

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapability5GSOk() (**os.File, bool)`

GetBinaryDataUeRadioCapability5GSOk returns a tuple with the BinaryDataUeRadioCapability5GS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCapability5GS

`func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCapability5GS(v *os.File)`

SetBinaryDataUeRadioCapability5GS sets BinaryDataUeRadioCapability5GS field to given value.

### HasBinaryDataUeRadioCapability5GS

`func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCapability5GS() bool`

HasBinaryDataUeRadioCapability5GS returns a boolean if a field has been set.

### GetBinaryDataUeRadioCapabilityEPS

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapabilityEPS() *os.File`

GetBinaryDataUeRadioCapabilityEPS returns the BinaryDataUeRadioCapabilityEPS field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCapabilityEPSOk

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapabilityEPSOk() (**os.File, bool)`

GetBinaryDataUeRadioCapabilityEPSOk returns a tuple with the BinaryDataUeRadioCapabilityEPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCapabilityEPS

`func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCapabilityEPS(v *os.File)`

SetBinaryDataUeRadioCapabilityEPS sets BinaryDataUeRadioCapabilityEPS field to given value.

### HasBinaryDataUeRadioCapabilityEPS

`func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCapabilityEPS() bool`

HasBinaryDataUeRadioCapabilityEPS returns a boolean if a field has been set.

### GetBinaryDataUeRadioCap5GSForPaging

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCap5GSForPaging() *os.File`

GetBinaryDataUeRadioCap5GSForPaging returns the BinaryDataUeRadioCap5GSForPaging field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCap5GSForPagingOk

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCap5GSForPagingOk() (**os.File, bool)`

GetBinaryDataUeRadioCap5GSForPagingOk returns a tuple with the BinaryDataUeRadioCap5GSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCap5GSForPaging

`func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCap5GSForPaging(v *os.File)`

SetBinaryDataUeRadioCap5GSForPaging sets BinaryDataUeRadioCap5GSForPaging field to given value.

### HasBinaryDataUeRadioCap5GSForPaging

`func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCap5GSForPaging() bool`

HasBinaryDataUeRadioCap5GSForPaging returns a boolean if a field has been set.

### GetBinaryDataUeRadioCapEPSForPaging

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapEPSForPaging() *os.File`

GetBinaryDataUeRadioCapEPSForPaging returns the BinaryDataUeRadioCapEPSForPaging field if non-nil, zero value otherwise.

### GetBinaryDataUeRadioCapEPSForPagingOk

`func (o *RetrieveDictionaryEntry200Response) GetBinaryDataUeRadioCapEPSForPagingOk() (**os.File, bool)`

GetBinaryDataUeRadioCapEPSForPagingOk returns a tuple with the BinaryDataUeRadioCapEPSForPaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUeRadioCapEPSForPaging

`func (o *RetrieveDictionaryEntry200Response) SetBinaryDataUeRadioCapEPSForPaging(v *os.File)`

SetBinaryDataUeRadioCapEPSForPaging sets BinaryDataUeRadioCapEPSForPaging field to given value.

### HasBinaryDataUeRadioCapEPSForPaging

`func (o *RetrieveDictionaryEntry200Response) HasBinaryDataUeRadioCapEPSForPaging() bool`

HasBinaryDataUeRadioCapEPSForPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


