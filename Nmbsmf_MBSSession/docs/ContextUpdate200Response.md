# ContextUpdate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**ContextUpdateRspData**](ContextUpdateRspData.md) |  | [optional] 
**BinaryDataN2Information** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewContextUpdate200Response

`func NewContextUpdate200Response() *ContextUpdate200Response`

NewContextUpdate200Response instantiates a new ContextUpdate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextUpdate200ResponseWithDefaults

`func NewContextUpdate200ResponseWithDefaults() *ContextUpdate200Response`

NewContextUpdate200ResponseWithDefaults instantiates a new ContextUpdate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *ContextUpdate200Response) GetJsonData() ContextUpdateRspData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *ContextUpdate200Response) GetJsonDataOk() (*ContextUpdateRspData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *ContextUpdate200Response) SetJsonData(v ContextUpdateRspData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *ContextUpdate200Response) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN2Information

`func (o *ContextUpdate200Response) GetBinaryDataN2Information() *os.File`

GetBinaryDataN2Information returns the BinaryDataN2Information field if non-nil, zero value otherwise.

### GetBinaryDataN2InformationOk

`func (o *ContextUpdate200Response) GetBinaryDataN2InformationOk() (**os.File, bool)`

GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2Information

`func (o *ContextUpdate200Response) SetBinaryDataN2Information(v *os.File)`

SetBinaryDataN2Information sets BinaryDataN2Information field to given value.

### HasBinaryDataN2Information

`func (o *ContextUpdate200Response) HasBinaryDataN2Information() bool`

HasBinaryDataN2Information returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


