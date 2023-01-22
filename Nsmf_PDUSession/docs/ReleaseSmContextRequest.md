# ReleaseSmContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**SmContextReleaseData**](SmContextReleaseData.md) |  | [optional] 
**BinaryDataN2SmInformation** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewReleaseSmContextRequest

`func NewReleaseSmContextRequest() *ReleaseSmContextRequest`

NewReleaseSmContextRequest instantiates a new ReleaseSmContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseSmContextRequestWithDefaults

`func NewReleaseSmContextRequestWithDefaults() *ReleaseSmContextRequest`

NewReleaseSmContextRequestWithDefaults instantiates a new ReleaseSmContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *ReleaseSmContextRequest) GetJsonData() SmContextReleaseData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *ReleaseSmContextRequest) GetJsonDataOk() (*SmContextReleaseData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *ReleaseSmContextRequest) SetJsonData(v SmContextReleaseData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *ReleaseSmContextRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN2SmInformation

`func (o *ReleaseSmContextRequest) GetBinaryDataN2SmInformation() *os.File`

GetBinaryDataN2SmInformation returns the BinaryDataN2SmInformation field if non-nil, zero value otherwise.

### GetBinaryDataN2SmInformationOk

`func (o *ReleaseSmContextRequest) GetBinaryDataN2SmInformationOk() (**os.File, bool)`

GetBinaryDataN2SmInformationOk returns a tuple with the BinaryDataN2SmInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2SmInformation

`func (o *ReleaseSmContextRequest) SetBinaryDataN2SmInformation(v *os.File)`

SetBinaryDataN2SmInformation sets BinaryDataN2SmInformation field to given value.

### HasBinaryDataN2SmInformation

`func (o *ReleaseSmContextRequest) HasBinaryDataN2SmInformation() bool`

HasBinaryDataN2SmInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


