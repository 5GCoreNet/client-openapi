# UpdateSmContext400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**SmContextUpdateError**](SmContextUpdateError.md) |  | [optional] 
**BinaryDataN1SmMessage** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN2SmInformation** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewUpdateSmContext400Response

`func NewUpdateSmContext400Response() *UpdateSmContext400Response`

NewUpdateSmContext400Response instantiates a new UpdateSmContext400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSmContext400ResponseWithDefaults

`func NewUpdateSmContext400ResponseWithDefaults() *UpdateSmContext400Response`

NewUpdateSmContext400ResponseWithDefaults instantiates a new UpdateSmContext400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *UpdateSmContext400Response) GetJsonData() SmContextUpdateError`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *UpdateSmContext400Response) GetJsonDataOk() (*SmContextUpdateError, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *UpdateSmContext400Response) SetJsonData(v SmContextUpdateError)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *UpdateSmContext400Response) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1SmMessage

`func (o *UpdateSmContext400Response) GetBinaryDataN1SmMessage() *os.File`

GetBinaryDataN1SmMessage returns the BinaryDataN1SmMessage field if non-nil, zero value otherwise.

### GetBinaryDataN1SmMessageOk

`func (o *UpdateSmContext400Response) GetBinaryDataN1SmMessageOk() (**os.File, bool)`

GetBinaryDataN1SmMessageOk returns a tuple with the BinaryDataN1SmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1SmMessage

`func (o *UpdateSmContext400Response) SetBinaryDataN1SmMessage(v *os.File)`

SetBinaryDataN1SmMessage sets BinaryDataN1SmMessage field to given value.

### HasBinaryDataN1SmMessage

`func (o *UpdateSmContext400Response) HasBinaryDataN1SmMessage() bool`

HasBinaryDataN1SmMessage returns a boolean if a field has been set.

### GetBinaryDataN2SmInformation

`func (o *UpdateSmContext400Response) GetBinaryDataN2SmInformation() *os.File`

GetBinaryDataN2SmInformation returns the BinaryDataN2SmInformation field if non-nil, zero value otherwise.

### GetBinaryDataN2SmInformationOk

`func (o *UpdateSmContext400Response) GetBinaryDataN2SmInformationOk() (**os.File, bool)`

GetBinaryDataN2SmInformationOk returns a tuple with the BinaryDataN2SmInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2SmInformation

`func (o *UpdateSmContext400Response) SetBinaryDataN2SmInformation(v *os.File)`

SetBinaryDataN2SmInformation sets BinaryDataN2SmInformation field to given value.

### HasBinaryDataN2SmInformation

`func (o *UpdateSmContext400Response) HasBinaryDataN2SmInformation() bool`

HasBinaryDataN2SmInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


