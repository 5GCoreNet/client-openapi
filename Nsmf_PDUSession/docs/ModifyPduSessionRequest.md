# ModifyPduSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**VsmfUpdateData**](VsmfUpdateData.md) |  | [optional] 
**BinaryDataN1SmInfoToUe** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4Information** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt1** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt2** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt3** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewModifyPduSessionRequest

`func NewModifyPduSessionRequest() *ModifyPduSessionRequest`

NewModifyPduSessionRequest instantiates a new ModifyPduSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyPduSessionRequestWithDefaults

`func NewModifyPduSessionRequestWithDefaults() *ModifyPduSessionRequest`

NewModifyPduSessionRequestWithDefaults instantiates a new ModifyPduSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *ModifyPduSessionRequest) GetJsonData() VsmfUpdateData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *ModifyPduSessionRequest) GetJsonDataOk() (*VsmfUpdateData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *ModifyPduSessionRequest) SetJsonData(v VsmfUpdateData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *ModifyPduSessionRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1SmInfoToUe

`func (o *ModifyPduSessionRequest) GetBinaryDataN1SmInfoToUe() *os.File`

GetBinaryDataN1SmInfoToUe returns the BinaryDataN1SmInfoToUe field if non-nil, zero value otherwise.

### GetBinaryDataN1SmInfoToUeOk

`func (o *ModifyPduSessionRequest) GetBinaryDataN1SmInfoToUeOk() (**os.File, bool)`

GetBinaryDataN1SmInfoToUeOk returns a tuple with the BinaryDataN1SmInfoToUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1SmInfoToUe

`func (o *ModifyPduSessionRequest) SetBinaryDataN1SmInfoToUe(v *os.File)`

SetBinaryDataN1SmInfoToUe sets BinaryDataN1SmInfoToUe field to given value.

### HasBinaryDataN1SmInfoToUe

`func (o *ModifyPduSessionRequest) HasBinaryDataN1SmInfoToUe() bool`

HasBinaryDataN1SmInfoToUe returns a boolean if a field has been set.

### GetBinaryDataN4Information

`func (o *ModifyPduSessionRequest) GetBinaryDataN4Information() *os.File`

GetBinaryDataN4Information returns the BinaryDataN4Information field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationOk

`func (o *ModifyPduSessionRequest) GetBinaryDataN4InformationOk() (**os.File, bool)`

GetBinaryDataN4InformationOk returns a tuple with the BinaryDataN4Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4Information

`func (o *ModifyPduSessionRequest) SetBinaryDataN4Information(v *os.File)`

SetBinaryDataN4Information sets BinaryDataN4Information field to given value.

### HasBinaryDataN4Information

`func (o *ModifyPduSessionRequest) HasBinaryDataN4Information() bool`

HasBinaryDataN4Information returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt1

`func (o *ModifyPduSessionRequest) GetBinaryDataN4InformationExt1() *os.File`

GetBinaryDataN4InformationExt1 returns the BinaryDataN4InformationExt1 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt1Ok

`func (o *ModifyPduSessionRequest) GetBinaryDataN4InformationExt1Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt1Ok returns a tuple with the BinaryDataN4InformationExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt1

`func (o *ModifyPduSessionRequest) SetBinaryDataN4InformationExt1(v *os.File)`

SetBinaryDataN4InformationExt1 sets BinaryDataN4InformationExt1 field to given value.

### HasBinaryDataN4InformationExt1

`func (o *ModifyPduSessionRequest) HasBinaryDataN4InformationExt1() bool`

HasBinaryDataN4InformationExt1 returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt2

`func (o *ModifyPduSessionRequest) GetBinaryDataN4InformationExt2() *os.File`

GetBinaryDataN4InformationExt2 returns the BinaryDataN4InformationExt2 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt2Ok

`func (o *ModifyPduSessionRequest) GetBinaryDataN4InformationExt2Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt2Ok returns a tuple with the BinaryDataN4InformationExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt2

`func (o *ModifyPduSessionRequest) SetBinaryDataN4InformationExt2(v *os.File)`

SetBinaryDataN4InformationExt2 sets BinaryDataN4InformationExt2 field to given value.

### HasBinaryDataN4InformationExt2

`func (o *ModifyPduSessionRequest) HasBinaryDataN4InformationExt2() bool`

HasBinaryDataN4InformationExt2 returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt3

`func (o *ModifyPduSessionRequest) GetBinaryDataN4InformationExt3() *os.File`

GetBinaryDataN4InformationExt3 returns the BinaryDataN4InformationExt3 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt3Ok

`func (o *ModifyPduSessionRequest) GetBinaryDataN4InformationExt3Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt3Ok returns a tuple with the BinaryDataN4InformationExt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt3

`func (o *ModifyPduSessionRequest) SetBinaryDataN4InformationExt3(v *os.File)`

SetBinaryDataN4InformationExt3 sets BinaryDataN4InformationExt3 field to given value.

### HasBinaryDataN4InformationExt3

`func (o *ModifyPduSessionRequest) HasBinaryDataN4InformationExt3() bool`

HasBinaryDataN4InformationExt3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


