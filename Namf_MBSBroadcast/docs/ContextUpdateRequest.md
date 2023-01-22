# ContextUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**ContextUpdateReqData**](ContextUpdateReqData.md) |  | [optional] 
**BinaryDataN2Information** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewContextUpdateRequest

`func NewContextUpdateRequest() *ContextUpdateRequest`

NewContextUpdateRequest instantiates a new ContextUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextUpdateRequestWithDefaults

`func NewContextUpdateRequestWithDefaults() *ContextUpdateRequest`

NewContextUpdateRequestWithDefaults instantiates a new ContextUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *ContextUpdateRequest) GetJsonData() ContextUpdateReqData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *ContextUpdateRequest) GetJsonDataOk() (*ContextUpdateReqData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *ContextUpdateRequest) SetJsonData(v ContextUpdateReqData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *ContextUpdateRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN2Information

`func (o *ContextUpdateRequest) GetBinaryDataN2Information() *os.File`

GetBinaryDataN2Information returns the BinaryDataN2Information field if non-nil, zero value otherwise.

### GetBinaryDataN2InformationOk

`func (o *ContextUpdateRequest) GetBinaryDataN2InformationOk() (**os.File, bool)`

GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2Information

`func (o *ContextUpdateRequest) SetBinaryDataN2Information(v *os.File)`

SetBinaryDataN2Information sets BinaryDataN2Information field to given value.

### HasBinaryDataN2Information

`func (o *ContextUpdateRequest) HasBinaryDataN2Information() bool`

HasBinaryDataN2Information returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


