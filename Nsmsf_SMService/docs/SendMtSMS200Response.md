# SendMtSMS200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonData** | Pointer to [**SmsDeliveryData**](SmsDeliveryData.md) |  | [optional] 
**BinaryPayload** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewSendMtSMS200Response

`func NewSendMtSMS200Response() *SendMtSMS200Response`

NewSendMtSMS200Response instantiates a new SendMtSMS200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMtSMS200ResponseWithDefaults

`func NewSendMtSMS200ResponseWithDefaults() *SendMtSMS200Response`

NewSendMtSMS200ResponseWithDefaults instantiates a new SendMtSMS200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonData

`func (o *SendMtSMS200Response) GetJsonData() SmsDeliveryData`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *SendMtSMS200Response) GetJsonDataOk() (*SmsDeliveryData, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *SendMtSMS200Response) SetJsonData(v SmsDeliveryData)`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *SendMtSMS200Response) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetBinaryPayload

`func (o *SendMtSMS200Response) GetBinaryPayload() *os.File`

GetBinaryPayload returns the BinaryPayload field if non-nil, zero value otherwise.

### GetBinaryPayloadOk

`func (o *SendMtSMS200Response) GetBinaryPayloadOk() (**os.File, bool)`

GetBinaryPayloadOk returns a tuple with the BinaryPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryPayload

`func (o *SendMtSMS200Response) SetBinaryPayload(v *os.File)`

SetBinaryPayload sets BinaryPayload field to given value.

### HasBinaryPayload

`func (o *SendMtSMS200Response) HasBinaryPayload() bool`

HasBinaryPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


