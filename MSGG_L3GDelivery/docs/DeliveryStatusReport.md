# DeliveryStatusReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriAddr** | [**Address1**](Address1.md) |  | 
**DestAddr** | [**Address1**](Address1.md) |  | 
**MsgId** | **string** |  | 
**SecCred** | Pointer to **string** |  | [optional] 
**FailureCause** | Pointer to **string** |  | [optional] 
**DelivSt** | [**ReportDeliveryStatus**](ReportDeliveryStatus.md) |  | 

## Methods

### NewDeliveryStatusReport

`func NewDeliveryStatusReport(oriAddr Address1, destAddr Address1, msgId string, delivSt ReportDeliveryStatus, ) *DeliveryStatusReport`

NewDeliveryStatusReport instantiates a new DeliveryStatusReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryStatusReportWithDefaults

`func NewDeliveryStatusReportWithDefaults() *DeliveryStatusReport`

NewDeliveryStatusReportWithDefaults instantiates a new DeliveryStatusReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriAddr

`func (o *DeliveryStatusReport) GetOriAddr() Address1`

GetOriAddr returns the OriAddr field if non-nil, zero value otherwise.

### GetOriAddrOk

`func (o *DeliveryStatusReport) GetOriAddrOk() (*Address1, bool)`

GetOriAddrOk returns a tuple with the OriAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriAddr

`func (o *DeliveryStatusReport) SetOriAddr(v Address1)`

SetOriAddr sets OriAddr field to given value.


### GetDestAddr

`func (o *DeliveryStatusReport) GetDestAddr() Address1`

GetDestAddr returns the DestAddr field if non-nil, zero value otherwise.

### GetDestAddrOk

`func (o *DeliveryStatusReport) GetDestAddrOk() (*Address1, bool)`

GetDestAddrOk returns a tuple with the DestAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAddr

`func (o *DeliveryStatusReport) SetDestAddr(v Address1)`

SetDestAddr sets DestAddr field to given value.


### GetMsgId

`func (o *DeliveryStatusReport) GetMsgId() string`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *DeliveryStatusReport) GetMsgIdOk() (*string, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *DeliveryStatusReport) SetMsgId(v string)`

SetMsgId sets MsgId field to given value.


### GetSecCred

`func (o *DeliveryStatusReport) GetSecCred() string`

GetSecCred returns the SecCred field if non-nil, zero value otherwise.

### GetSecCredOk

`func (o *DeliveryStatusReport) GetSecCredOk() (*string, bool)`

GetSecCredOk returns a tuple with the SecCred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecCred

`func (o *DeliveryStatusReport) SetSecCred(v string)`

SetSecCred sets SecCred field to given value.

### HasSecCred

`func (o *DeliveryStatusReport) HasSecCred() bool`

HasSecCred returns a boolean if a field has been set.

### GetFailureCause

`func (o *DeliveryStatusReport) GetFailureCause() string`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *DeliveryStatusReport) GetFailureCauseOk() (*string, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *DeliveryStatusReport) SetFailureCause(v string)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *DeliveryStatusReport) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetDelivSt

`func (o *DeliveryStatusReport) GetDelivSt() ReportDeliveryStatus`

GetDelivSt returns the DelivSt field if non-nil, zero value otherwise.

### GetDelivStOk

`func (o *DeliveryStatusReport) GetDelivStOk() (*ReportDeliveryStatus, bool)`

GetDelivStOk returns a tuple with the DelivSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivSt

`func (o *DeliveryStatusReport) SetDelivSt(v ReportDeliveryStatus)`

SetDelivSt sets DelivSt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


