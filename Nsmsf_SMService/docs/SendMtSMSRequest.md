# SendMtSMSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**SmsData**](SmsData.md) |  | [optional] 
**BinaryPayload** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewSendMtSMSRequest

`func NewSendMtSMSRequest() *SendMtSMSRequest`

NewSendMtSMSRequest instantiates a new SendMtSMSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMtSMSRequestWithDefaults

`func NewSendMtSMSRequestWithDefaults() *SendMtSMSRequest`

NewSendMtSMSRequestWithDefaults instantiates a new SendMtSMSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *SendMtSMSRequest) GetJsonData() SmsData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *SendMtSMSRequest) GetJsonDataOk() (*SmsData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *SendMtSMSRequest) SetJsonData(v SmsData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *SendMtSMSRequest) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryPayload

`func (o *SendMtSMSRequest) GetBinaryPayload() *os.File`

GetBinaryPayload returns the BinaryPayload field if non-nil, zero value otherwise.

### GetBinaryPayloadOk

`func (o *SendMtSMSRequest) GetBinaryPayloadOk() (**os.File, bool)`

GetBinaryPayloadOk returns a tuple with the BinaryPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryPayload

`func (o *SendMtSMSRequest) SetBinaryPayload(v *os.File)`

SetBinaryPayload sets BinaryPayload field to given value.

### HasBinaryPayload

`func (o *SendMtSMSRequest) HasBinaryPayload() bool`

HasBinaryPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


