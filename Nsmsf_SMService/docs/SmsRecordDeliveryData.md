# SmsRecordDeliveryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsRecordId** | **string** | Represents a record ID, used to identify a message carrying SMS payload. | 
**DeliveryStatus** | [**SmsDeliveryStatus**](SmsDeliveryStatus.md) |  | 

## Methods

### NewSmsRecordDeliveryData

`func NewSmsRecordDeliveryData(smsRecordId string, deliveryStatus SmsDeliveryStatus, ) *SmsRecordDeliveryData`

NewSmsRecordDeliveryData instantiates a new SmsRecordDeliveryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsRecordDeliveryDataWithDefaults

`func NewSmsRecordDeliveryDataWithDefaults() *SmsRecordDeliveryData`

NewSmsRecordDeliveryDataWithDefaults instantiates a new SmsRecordDeliveryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsRecordId

`func (o *SmsRecordDeliveryData) GetSmsRecordId() string`

GetSmsRecordId returns the SmsRecordId field if non-nil, zero value otherwise.

### GetSmsRecordIdOk

`func (o *SmsRecordDeliveryData) GetSmsRecordIdOk() (*string, bool)`

GetSmsRecordIdOk returns a tuple with the SmsRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsRecordId

`func (o *SmsRecordDeliveryData) SetSmsRecordId(v string)`

SetSmsRecordId sets SmsRecordId field to given value.


### GetDeliveryStatus

`func (o *SmsRecordDeliveryData) GetDeliveryStatus() SmsDeliveryStatus`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *SmsRecordDeliveryData) GetDeliveryStatusOk() (*SmsDeliveryStatus, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *SmsRecordDeliveryData) SetDeliveryStatus(v SmsDeliveryStatus)`

SetDeliveryStatus sets DeliveryStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


