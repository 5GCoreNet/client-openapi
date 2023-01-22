# PolicyTemplateApplicationSessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfAppId** | Pointer to **string** | Contains an AF application identifier. | [optional] 
**SliceInfo** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**AspId** | Pointer to **string** | Contains an identity of an application service provider. | [optional] 

## Methods

### NewPolicyTemplateApplicationSessionContext

`func NewPolicyTemplateApplicationSessionContext() *PolicyTemplateApplicationSessionContext`

NewPolicyTemplateApplicationSessionContext instantiates a new PolicyTemplateApplicationSessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateApplicationSessionContextWithDefaults

`func NewPolicyTemplateApplicationSessionContextWithDefaults() *PolicyTemplateApplicationSessionContext`

NewPolicyTemplateApplicationSessionContextWithDefaults instantiates a new PolicyTemplateApplicationSessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfAppId

`func (o *PolicyTemplateApplicationSessionContext) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *PolicyTemplateApplicationSessionContext) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *PolicyTemplateApplicationSessionContext) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *PolicyTemplateApplicationSessionContext) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.

### GetSliceInfo

`func (o *PolicyTemplateApplicationSessionContext) GetSliceInfo() Snssai`

GetSliceInfo returns the SliceInfo field if non-nil, zero value otherwise.

### GetSliceInfoOk

`func (o *PolicyTemplateApplicationSessionContext) GetSliceInfoOk() (*Snssai, bool)`

GetSliceInfoOk returns a tuple with the SliceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceInfo

`func (o *PolicyTemplateApplicationSessionContext) SetSliceInfo(v Snssai)`

SetSliceInfo sets SliceInfo field to given value.

### HasSliceInfo

`func (o *PolicyTemplateApplicationSessionContext) HasSliceInfo() bool`

HasSliceInfo returns a boolean if a field has been set.

### GetDnn

`func (o *PolicyTemplateApplicationSessionContext) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PolicyTemplateApplicationSessionContext) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PolicyTemplateApplicationSessionContext) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *PolicyTemplateApplicationSessionContext) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetAspId

`func (o *PolicyTemplateApplicationSessionContext) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *PolicyTemplateApplicationSessionContext) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *PolicyTemplateApplicationSessionContext) SetAspId(v string)`

SetAspId sets AspId field to given value.

### HasAspId

`func (o *PolicyTemplateApplicationSessionContext) HasAspId() bool`

HasAspId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


