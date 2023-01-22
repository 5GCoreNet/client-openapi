# DetermineLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**InputData**](InputData.md) |  | [optional] 
**BinaryDataLppMessage** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewDetermineLocationRequest

`func NewDetermineLocationRequest() *DetermineLocationRequest`

NewDetermineLocationRequest instantiates a new DetermineLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetermineLocationRequestWithDefaults

`func NewDetermineLocationRequestWithDefaults() *DetermineLocationRequest`

NewDetermineLocationRequestWithDefaults instantiates a new DetermineLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *DetermineLocationRequest) GetJsonData() InputData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *DetermineLocationRequest) GetJsonDataOk() (*InputData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *DetermineLocationRequest) SetJsonData(v InputData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *DetermineLocationRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataLppMessage

`func (o *DetermineLocationRequest) GetBinaryDataLppMessage() *os.File`

GetBinaryDataLppMessage returns the BinaryDataLppMessage field if non-nil, zero value otherwise.

### GetBinaryDataLppMessageOk

`func (o *DetermineLocationRequest) GetBinaryDataLppMessageOk() (**os.File, bool)`

GetBinaryDataLppMessageOk returns a tuple with the BinaryDataLppMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataLppMessage

`func (o *DetermineLocationRequest) SetBinaryDataLppMessage(v *os.File)`

SetBinaryDataLppMessage sets BinaryDataLppMessage field to given value.

### HasBinaryDataLppMessage

`func (o *DetermineLocationRequest) HasBinaryDataLppMessage() bool`

HasBinaryDataLppMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


