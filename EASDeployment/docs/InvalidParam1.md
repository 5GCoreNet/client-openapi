# InvalidParam1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Param** | **string** | If the invalid parameter is an attribute in a JSON body, this IE shall contain the  attribute&#39;s name and shall be encoded as a JSON Pointer. If the invalid parameter is  an HTTP header, this IE shall be formatted as the concatenation of the string \&quot;header \&quot;  plus the name of such header. If the invalid parameter is a query parameter, this IE  shall be formatted as the concatenation of the string \&quot;query \&quot; plus the name of such  query parameter. If the invalid parameter is a variable part in the path of a resource  URI, this IE shall contain the name of the variable, including the symbols \&quot;{\&quot; and \&quot;}\&quot;  used in OpenAPI specification as the notation to represent variable path segments.  | 
**Reason** | Pointer to **string** | A human-readable reason, e.g. \&quot;must be a positive integer\&quot;. In cases involving failed  operations in a PATCH request, the reason string should identify the operation that  failed using the operation&#39;s array index to assist in correlation of the invalid  parameter with the failed operation, e.g.\&quot; Replacement value invalid for attribute  (failed operation index&#x3D; 4)\&quot;  | [optional] 

## Methods

### NewInvalidParam1

`func NewInvalidParam1(param string, ) *InvalidParam1`

NewInvalidParam1 instantiates a new InvalidParam1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidParam1WithDefaults

`func NewInvalidParam1WithDefaults() *InvalidParam1`

NewInvalidParam1WithDefaults instantiates a new InvalidParam1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParam

`func (o *InvalidParam1) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *InvalidParam1) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *InvalidParam1) SetParam(v string)`

SetParam sets Param field to given value.


### GetReason

`func (o *InvalidParam1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvalidParam1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvalidParam1) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvalidParam1) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


