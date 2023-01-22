# SmcceInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**SmcceUeList** | [**SmcceUeList1**](SmcceUeList1.md) |  | 

## Methods

### NewSmcceInfo1

`func NewSmcceInfo1(smcceUeList SmcceUeList1, ) *SmcceInfo1`

NewSmcceInfo1 instantiates a new SmcceInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmcceInfo1WithDefaults

`func NewSmcceInfo1WithDefaults() *SmcceInfo1`

NewSmcceInfo1WithDefaults instantiates a new SmcceInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *SmcceInfo1) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmcceInfo1) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmcceInfo1) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SmcceInfo1) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *SmcceInfo1) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SmcceInfo1) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SmcceInfo1) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *SmcceInfo1) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetSmcceUeList

`func (o *SmcceInfo1) GetSmcceUeList() SmcceUeList1`

GetSmcceUeList returns the SmcceUeList field if non-nil, zero value otherwise.

### GetSmcceUeListOk

`func (o *SmcceInfo1) GetSmcceUeListOk() (*SmcceUeList1, bool)`

GetSmcceUeListOk returns a tuple with the SmcceUeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmcceUeList

`func (o *SmcceInfo1) SetSmcceUeList(v SmcceUeList1)`

SetSmcceUeList sets SmcceUeList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


