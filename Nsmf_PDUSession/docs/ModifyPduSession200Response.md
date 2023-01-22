# ModifyPduSession200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**VsmfUpdatedData**](VsmfUpdatedData.md) |  | [optional] 
**BinaryDataN1SmInfoFromUe** | Pointer to ***os.File** |  | [optional] 
**BinaryDataUnknownN1SmInfo** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4Information** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt1** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt2** | Pointer to ***os.File** |  | [optional] 
**BinaryDataN4InformationExt3** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewModifyPduSession200Response

`func NewModifyPduSession200Response() *ModifyPduSession200Response`

NewModifyPduSession200Response instantiates a new ModifyPduSession200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyPduSession200ResponseWithDefaults

`func NewModifyPduSession200ResponseWithDefaults() *ModifyPduSession200Response`

NewModifyPduSession200ResponseWithDefaults instantiates a new ModifyPduSession200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *ModifyPduSession200Response) GetJsonData() VsmfUpdatedData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *ModifyPduSession200Response) GetJsonDataOk() (*VsmfUpdatedData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *ModifyPduSession200Response) SetJsonData(v VsmfUpdatedData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *ModifyPduSession200Response) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1SmInfoFromUe

`func (o *ModifyPduSession200Response) GetBinaryDataN1SmInfoFromUe() *os.File`

GetBinaryDataN1SmInfoFromUe returns the BinaryDataN1SmInfoFromUe field if non-nil, zero value otherwise.

### GetBinaryDataN1SmInfoFromUeOk

`func (o *ModifyPduSession200Response) GetBinaryDataN1SmInfoFromUeOk() (**os.File, bool)`

GetBinaryDataN1SmInfoFromUeOk returns a tuple with the BinaryDataN1SmInfoFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1SmInfoFromUe

`func (o *ModifyPduSession200Response) SetBinaryDataN1SmInfoFromUe(v *os.File)`

SetBinaryDataN1SmInfoFromUe sets BinaryDataN1SmInfoFromUe field to given value.

### HasBinaryDataN1SmInfoFromUe

`func (o *ModifyPduSession200Response) HasBinaryDataN1SmInfoFromUe() bool`

HasBinaryDataN1SmInfoFromUe returns a boolean if a field has been set.

### GetBinaryDataUnknownN1SmInfo

`func (o *ModifyPduSession200Response) GetBinaryDataUnknownN1SmInfo() *os.File`

GetBinaryDataUnknownN1SmInfo returns the BinaryDataUnknownN1SmInfo field if non-nil, zero value otherwise.

### GetBinaryDataUnknownN1SmInfoOk

`func (o *ModifyPduSession200Response) GetBinaryDataUnknownN1SmInfoOk() (**os.File, bool)`

GetBinaryDataUnknownN1SmInfoOk returns a tuple with the BinaryDataUnknownN1SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataUnknownN1SmInfo

`func (o *ModifyPduSession200Response) SetBinaryDataUnknownN1SmInfo(v *os.File)`

SetBinaryDataUnknownN1SmInfo sets BinaryDataUnknownN1SmInfo field to given value.

### HasBinaryDataUnknownN1SmInfo

`func (o *ModifyPduSession200Response) HasBinaryDataUnknownN1SmInfo() bool`

HasBinaryDataUnknownN1SmInfo returns a boolean if a field has been set.

### GetBinaryDataN4Information

`func (o *ModifyPduSession200Response) GetBinaryDataN4Information() *os.File`

GetBinaryDataN4Information returns the BinaryDataN4Information field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationOk

`func (o *ModifyPduSession200Response) GetBinaryDataN4InformationOk() (**os.File, bool)`

GetBinaryDataN4InformationOk returns a tuple with the BinaryDataN4Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4Information

`func (o *ModifyPduSession200Response) SetBinaryDataN4Information(v *os.File)`

SetBinaryDataN4Information sets BinaryDataN4Information field to given value.

### HasBinaryDataN4Information

`func (o *ModifyPduSession200Response) HasBinaryDataN4Information() bool`

HasBinaryDataN4Information returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt1

`func (o *ModifyPduSession200Response) GetBinaryDataN4InformationExt1() *os.File`

GetBinaryDataN4InformationExt1 returns the BinaryDataN4InformationExt1 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt1Ok

`func (o *ModifyPduSession200Response) GetBinaryDataN4InformationExt1Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt1Ok returns a tuple with the BinaryDataN4InformationExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt1

`func (o *ModifyPduSession200Response) SetBinaryDataN4InformationExt1(v *os.File)`

SetBinaryDataN4InformationExt1 sets BinaryDataN4InformationExt1 field to given value.

### HasBinaryDataN4InformationExt1

`func (o *ModifyPduSession200Response) HasBinaryDataN4InformationExt1() bool`

HasBinaryDataN4InformationExt1 returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt2

`func (o *ModifyPduSession200Response) GetBinaryDataN4InformationExt2() *os.File`

GetBinaryDataN4InformationExt2 returns the BinaryDataN4InformationExt2 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt2Ok

`func (o *ModifyPduSession200Response) GetBinaryDataN4InformationExt2Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt2Ok returns a tuple with the BinaryDataN4InformationExt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt2

`func (o *ModifyPduSession200Response) SetBinaryDataN4InformationExt2(v *os.File)`

SetBinaryDataN4InformationExt2 sets BinaryDataN4InformationExt2 field to given value.

### HasBinaryDataN4InformationExt2

`func (o *ModifyPduSession200Response) HasBinaryDataN4InformationExt2() bool`

HasBinaryDataN4InformationExt2 returns a boolean if a field has been set.

### GetBinaryDataN4InformationExt3

`func (o *ModifyPduSession200Response) GetBinaryDataN4InformationExt3() *os.File`

GetBinaryDataN4InformationExt3 returns the BinaryDataN4InformationExt3 field if non-nil, zero value otherwise.

### GetBinaryDataN4InformationExt3Ok

`func (o *ModifyPduSession200Response) GetBinaryDataN4InformationExt3Ok() (**os.File, bool)`

GetBinaryDataN4InformationExt3Ok returns a tuple with the BinaryDataN4InformationExt3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN4InformationExt3

`func (o *ModifyPduSession200Response) SetBinaryDataN4InformationExt3(v *os.File)`

SetBinaryDataN4InformationExt3 sets BinaryDataN4InformationExt3 field to given value.

### HasBinaryDataN4InformationExt3

`func (o *ModifyPduSession200Response) HasBinaryDataN4InformationExt3() bool`

HasBinaryDataN4InformationExt3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


