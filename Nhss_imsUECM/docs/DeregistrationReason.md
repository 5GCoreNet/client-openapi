# DeregistrationReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | [**DeregistrationReasonCode**](DeregistrationReasonCode.md) |  | 
**ReasonText** | **string** |  | 

## Methods

### NewDeregistrationReason

`func NewDeregistrationReason(reasonCode DeregistrationReasonCode, reasonText string, ) *DeregistrationReason`

NewDeregistrationReason instantiates a new DeregistrationReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeregistrationReasonWithDefaults

`func NewDeregistrationReasonWithDefaults() *DeregistrationReason`

NewDeregistrationReasonWithDefaults instantiates a new DeregistrationReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonCode

`func (o *DeregistrationReason) GetReasonCode() DeregistrationReasonCode`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DeregistrationReason) GetReasonCodeOk() (*DeregistrationReasonCode, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DeregistrationReason) SetReasonCode(v DeregistrationReasonCode)`

SetReasonCode sets ReasonCode field to given value.


### GetReasonText

`func (o *DeregistrationReason) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *DeregistrationReason) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *DeregistrationReason) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


