# CustomOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommType** | [**CommunicationType**](CommunicationType.md) |  | 
**CustOpName** | **string** | it is set as {custOpName} part of the URI structure for a custom operation without resource association as defined in clause 5.2.4 of 3GPP TS 29.122.  | 
**Operations** | Pointer to [**[]Operation**](Operation.md) | Supported HTTP methods for the API resource. Only applicable when the protocol in AefProfile indicates HTTP.  | [optional] 
**Description** | Pointer to **string** | Text description of the custom operation | [optional] 

## Methods

### NewCustomOperation

`func NewCustomOperation(commType CommunicationType, custOpName string, ) *CustomOperation`

NewCustomOperation instantiates a new CustomOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomOperationWithDefaults

`func NewCustomOperationWithDefaults() *CustomOperation`

NewCustomOperationWithDefaults instantiates a new CustomOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommType

`func (o *CustomOperation) GetCommType() CommunicationType`

GetCommType returns the CommType field if non-nil, zero value otherwise.

### GetCommTypeOk

`func (o *CustomOperation) GetCommTypeOk() (*CommunicationType, bool)`

GetCommTypeOk returns a tuple with the CommType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommType

`func (o *CustomOperation) SetCommType(v CommunicationType)`

SetCommType sets CommType field to given value.


### GetCustOpName

`func (o *CustomOperation) GetCustOpName() string`

GetCustOpName returns the CustOpName field if non-nil, zero value otherwise.

### GetCustOpNameOk

`func (o *CustomOperation) GetCustOpNameOk() (*string, bool)`

GetCustOpNameOk returns a tuple with the CustOpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustOpName

`func (o *CustomOperation) SetCustOpName(v string)`

SetCustOpName sets CustOpName field to given value.


### GetOperations

`func (o *CustomOperation) GetOperations() []Operation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CustomOperation) GetOperationsOk() (*[]Operation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CustomOperation) SetOperations(v []Operation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CustomOperation) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetDescription

`func (o *CustomOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomOperation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomOperation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


