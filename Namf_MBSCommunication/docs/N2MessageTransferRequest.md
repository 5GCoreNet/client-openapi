# N2MessageTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**MbsN2MessageTransferReqData**](MbsN2MessageTransferReqData.md) |  | [optional] 
**BinaryDataN2Information** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewN2MessageTransferRequest

`func NewN2MessageTransferRequest() *N2MessageTransferRequest`

NewN2MessageTransferRequest instantiates a new N2MessageTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2MessageTransferRequestWithDefaults

`func NewN2MessageTransferRequestWithDefaults() *N2MessageTransferRequest`

NewN2MessageTransferRequestWithDefaults instantiates a new N2MessageTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *N2MessageTransferRequest) GetJsonData() MbsN2MessageTransferReqData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *N2MessageTransferRequest) GetJsonDataOk() (*MbsN2MessageTransferReqData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *N2MessageTransferRequest) SetJsonData(v MbsN2MessageTransferReqData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *N2MessageTransferRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryDataN2Information

`func (o *N2MessageTransferRequest) GetBinaryDataN2Information() *os.File`

GetBinaryDataN2Information returns the BinaryDataN2Information field if non-nil, zero value otherwise.

### GetBinaryDataN2InformationOk

`func (o *N2MessageTransferRequest) GetBinaryDataN2InformationOk() (**os.File, bool)`

GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDataN2Information

`func (o *N2MessageTransferRequest) SetBinaryDataN2Information(v *os.File)`

SetBinaryDataN2Information sets BinaryDataN2Information field to given value.

### HasBinaryDataN2Information

`func (o *N2MessageTransferRequest) HasBinaryDataN2Information() bool`

HasBinaryDataN2Information returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


