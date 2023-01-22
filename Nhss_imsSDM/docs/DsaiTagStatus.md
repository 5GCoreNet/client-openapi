# DsaiTagStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DsaiTag** | **string** |  | 
**DsaiState** | [**ActivationState**](ActivationState.md) |  | 

## Methods

### NewDsaiTagStatus

`func NewDsaiTagStatus(dsaiTag string, dsaiState ActivationState, ) *DsaiTagStatus`

NewDsaiTagStatus instantiates a new DsaiTagStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsaiTagStatusWithDefaults

`func NewDsaiTagStatusWithDefaults() *DsaiTagStatus`

NewDsaiTagStatusWithDefaults instantiates a new DsaiTagStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDsaiTag

`func (o *DsaiTagStatus) GetDsaiTag() string`

GetDsaiTag returns the DsaiTag field if non-nil, zero value otherwise.

### GetDsaiTagOk

`func (o *DsaiTagStatus) GetDsaiTagOk() (*string, bool)`

GetDsaiTagOk returns a tuple with the DsaiTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaiTag

`func (o *DsaiTagStatus) SetDsaiTag(v string)`

SetDsaiTag sets DsaiTag field to given value.


### GetDsaiState

`func (o *DsaiTagStatus) GetDsaiState() ActivationState`

GetDsaiState returns the DsaiState field if non-nil, zero value otherwise.

### GetDsaiStateOk

`func (o *DsaiTagStatus) GetDsaiStateOk() (*ActivationState, bool)`

GetDsaiStateOk returns a tuple with the DsaiState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaiState

`func (o *DsaiTagStatus) SetDsaiState(v ActivationState)`

SetDsaiState sets DsaiState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


