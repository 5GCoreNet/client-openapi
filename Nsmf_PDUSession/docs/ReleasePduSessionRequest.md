# ReleasePduSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**ReleaseData**](ReleaseData.md) |  | [optional] 
**BinaryDataN4Information** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt1** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt2** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewReleasePduSessionRequest

`func NewReleasePduSessionRequest() *ReleasePduSessionRequest`

NewReleasePduSessionRequest instantiates a new ReleasePduSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePduSessionRequestWithDefaults

`func NewReleasePduSessionRequestWithDefaults() *ReleasePduSessionRequest`

NewReleasePduSessionRequestWithDefaults instantiates a new ReleasePduSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *ReleasePduSessionRequest) GetJsonData() ReleaseData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *ReleasePduSessionRequest) GetJsonDataOk() (*ReleaseData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *ReleasePduSessionRequest) SetJsonData(v ReleaseData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *ReleasePduSessionRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN4Information

`func (o *ReleasePduSessionRequest) GetBinaryDataN4Information() *os.File`

GetBinaryDataN4Information returns the BinaryDataN4Information field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationOk

`func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationOk() (**os.File, bool)`

GetBinaryDataN4InformationOk returns a tuple with the BinaryDataN4Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4Information

`func (o *ReleasePduSessionRequest) SetBinaryDataN4Information(v *os.File)`

SetBinaryDataN4Information sets BinaryDataN4Information field to given value.

### HasBinaryDataN4Information

`func (o *ReleasePduSessionRequest) HasBinaryDataN4Information() bool`

HasBinaryDataN4Information returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt1

`func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt1() *os.File`

GetBinaryDataN4InformationExt1 returns the BinaryDataN4InformationExt1 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt1Ok

`func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt1Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt1Ok returns a tuple with the BinaryDataN4InformationExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt1

`func (o *ReleasePduSessionRequest) SetBinaryDataN4InformationExt1(v *os.File)`

SetBinaryDataN4InformationExt1 sets BinaryDataN4InformationExt1 field to given value.

### HasBinaryDataN4InformationExt1

`func (o *ReleasePduSessionRequest) HasBinaryDataN4InformationExt1() bool`

HasBinaryDataN4InformationExt1 returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt2

`func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt2() *os.File`

GetBinaryDataN4InformationExt2 returns the BinaryDataN4InformationExt2 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt2Ok

`func (o *ReleasePduSessionRequest) GetBinaryDataN4InformationExt2Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt2Ok returns a tuple with the BinaryDataN4InformationExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt2

`func (o *ReleasePduSessionRequest) SetBinaryDataN4InformationExt2(v *os.File)`

SetBinaryDataN4InformationExt2 sets BinaryDataN4InformationExt2 field to given value.

### HasBinaryDataN4InformationExt2

`func (o *ReleasePduSessionRequest) HasBinaryDataN4InformationExt2() bool`

HasBinaryDataN4InformationExt2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


