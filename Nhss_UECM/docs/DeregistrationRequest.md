# DeregistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imsi** | **string** |  | 
**DeregReason** | [**DeregistrationReason**](DeregistrationReason.md) |  | 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 

## Methods

### NewDeregistrationRequest

`func NewDeregistrationRequest(imsi string, deregReason DeregistrationReason, ) *DeregistrationRequest`

NewDeregistrationRequest instantiates a new DeregistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeregistrationRequestWithDefaults

`func NewDeregistrationRequestWithDefaults() *DeregistrationRequest`

NewDeregistrationRequestWithDefaults instantiates a new DeregistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImsi

`func (o *DeregistrationRequest) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *DeregistrationRequest) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *DeregistrationRequest) SetImsi(v string)`

SetImsi sets Imsi field to given value.


### GetDeregReason

`func (o *DeregistrationRequest) GetDeregReason() DeregistrationReason`

GetDeregReason returns the DeregReason field if non-nil, zero value otherwise.

### GetDeregReasonOk

`func (o *DeregistrationRequest) GetDeregReasonOk() (*DeregistrationReason, bool)`

GetDeregReasonOk returns a tuple with the DeregReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregReason

`func (o *DeregistrationRequest) SetDeregReason(v DeregistrationReason)`

SetDeregReason sets DeregReason field to given value.


### GetGuami

`func (o *DeregistrationRequest) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *DeregistrationRequest) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *DeregistrationRequest) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *DeregistrationRequest) HasGuami() bool`

HasGuami returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


