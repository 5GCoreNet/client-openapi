# PostPduSessions201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**PduSessionCreatedData**](PduSessionCreatedData.md) |  | [optional] 
**BinaryDataN1SmInfoToUe** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewPostPduSessions201Response

`func NewPostPduSessions201Response() *PostPduSessions201Response`

NewPostPduSessions201Response instantiates a new PostPduSessions201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPduSessions201ResponseWithDefaults

`func NewPostPduSessions201ResponseWithDefaults() *PostPduSessions201Response`

NewPostPduSessions201ResponseWithDefaults instantiates a new PostPduSessions201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *PostPduSessions201Response) GetJsonData() PduSessionCreatedData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *PostPduSessions201Response) GetJsonDataOk() (*PduSessionCreatedData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *PostPduSessions201Response) SetJsonData(v PduSessionCreatedData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *PostPduSessions201Response) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN1SmInfoToUe

`func (o *PostPduSessions201Response) GetBinaryDataN1SmInfoToUe() *os.File`

GetBinaryDataN1SmInfoToUe returns the BinaryDataN1SmInfoToUe field if non-nil, zero value otherwise.

### GetBinaryDataN1SmInfoToUeOk

`func (o *PostPduSessions201Response) GetBinaryDataN1SmInfoToUeOk() (**os.File, bool)`

GetBinaryDataN1SmInfoToUeOk returns a tuple with the BinaryDataN1SmInfoToUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN1SmInfoToUe

`func (o *PostPduSessions201Response) SetBinaryDataN1SmInfoToUe(v *os.File)`

SetBinaryDataN1SmInfoToUe sets BinaryDataN1SmInfoToUe field to given value.

### HasBinaryDataN1SmInfoToUe

`func (o *PostPduSessions201Response) HasBinaryDataN1SmInfoToUe() bool`

HasBinaryDataN1SmInfoToUe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


