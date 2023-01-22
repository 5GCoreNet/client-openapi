# EECContextPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EesId** | **string** | Identifier of the S-EES pushing the EEC context information. | 
**EecCntx** | [**EECContext**](EECContext.md) |  | 

## Methods

### NewEECContextPush

`func NewEECContextPush(eesId string, eecCntx EECContext, ) *EECContextPush`

NewEECContextPush instantiates a new EECContextPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEECContextPushWithDefaults

`func NewEECContextPushWithDefaults() *EECContextPush`

NewEECContextPushWithDefaults instantiates a new EECContextPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEesId

`func (o *EECContextPush) GetEesId() string`

GetEesId returns the EesId field if non-nil, zero value otherwise.

### GetEesIdOk

`func (o *EECContextPush) GetEesIdOk() (*string, bool)`

GetEesIdOk returns a tuple with the EesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEesId

`func (o *EECContextPush) SetEesId(v string)`

SetEesId sets EesId field to given value.


### GetEecCntx

`func (o *EECContextPush) GetEecCntx() EECContext`

GetEecCntx returns the EecCntx field if non-nil, zero value otherwise.

### GetEecCntxOk

`func (o *EECContextPush) GetEecCntxOk() (*EECContext, bool)`

GetEecCntxOk returns a tuple with the EecCntx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecCntx

`func (o *EECContextPush) SetEecCntx(v EECContext)`

SetEecCntx sets EecCntx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


