# UpdateSmContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**SmContextUpdateData**](SmContextUpdateData.md) |  | [optional] 
**BinaryDataN1SmMessage** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN2SmInformation** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN2SmInformationExt1** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewUpdateSmContextRequest

`func NewUpdateSmContextRequest() *UpdateSmContextRequest`

NewUpdateSmContextRequest instantiates a new UpdateSmContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSmContextRequestWithDefaults

`func NewUpdateSmContextRequestWithDefaults() *UpdateSmContextRequest`

NewUpdateSmContextRequestWithDefaults instantiates a new UpdateSmContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *UpdateSmContextRequest) GetJsonData() SmContextUpdateData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *UpdateSmContextRequest) GetJsonDataOk() (*SmContextUpdateData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *UpdateSmContextRequest) SetJsonData(v SmContextUpdateData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *UpdateSmContextRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1SmMessage

`func (o *UpdateSmContextRequest) GetBinaryDataN1SmMessage() *os.File`

GetBinaryDataN1SmMessage returns the BinaryDataN1SmMessage field if non-nil, zero value otherwise.

### GetBinaryDataN1SmMessageOk

`func (o *UpdateSmContextRequest) GetBinaryDataN1SmMessageOk() (**os.File, bool)`

GetBinaryDataN1SmMessageOk returns a tuple with the BinaryDataN1SmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1SmMessage

`func (o *UpdateSmContextRequest) SetBinaryDataN1SmMessage(v *os.File)`

SetBinaryDataN1SmMessage sets BinaryDataN1SmMessage field to given value.

### HasBinaryDataN1SmMessage

`func (o *UpdateSmContextRequest) HasBinaryDataN1SmMessage() bool`

HasBinaryDataN1SmMessage returns a boolean if a field has been set.

### GetBinaryDataN2SmInformation

`func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformation() *os.File`

GetBinaryDataN2SmInformation returns the BinaryDataN2SmInformation field if non-nil, zero value otherwise.

### GetBinaryDataN2SmInformationOk

`func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformationOk() (**os.File, bool)`

GetBinaryDataN2SmInformationOk returns a tuple with the BinaryDataN2SmInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2SmInformation

`func (o *UpdateSmContextRequest) SetBinaryDataN2SmInformation(v *os.File)`

SetBinaryDataN2SmInformation sets BinaryDataN2SmInformation field to given value.

### HasBinaryDataN2SmInformation

`func (o *UpdateSmContextRequest) HasBinaryDataN2SmInformation() bool`

HasBinaryDataN2SmInformation returns a boolean if a field has been set.

### GetBinaryDataN2SmInformationExt1

`func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformationExt1() *os.File`

GetBinaryDataN2SmInformationExt1 returns the BinaryDataN2SmInformationExt1 field if non-nil, zero value otherwise.

### GetBinaryDataN2SmInformationExt1Ok

`func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformationExt1Ok() (**os.File, bool)`

GetBinaryDataN2SmInformationExt1Ok returns a tuple with the BinaryDataN2SmInformationExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2SmInformationExt1

`func (o *UpdateSmContextRequest) SetBinaryDataN2SmInformationExt1(v *os.File)`

SetBinaryDataN2SmInformationExt1 sets BinaryDataN2SmInformationExt1 field to given value.

### HasBinaryDataN2SmInformationExt1

`func (o *UpdateSmContextRequest) HasBinaryDataN2SmInformationExt1() bool`

HasBinaryDataN2SmInformationExt1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


