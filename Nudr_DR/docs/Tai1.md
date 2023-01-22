# Tai1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId1**](PlmnId1.md) |  | 
**Tac** | **string** | 2 or 3-octet string identifying a tracking area code as specified in clause 9.3.3.10  of 3GPP TS 38.413, in hexadecimal representation. Each character in the string shall  take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits of the TAC shall  appear first in the string, and the character representing the 4 least significant bit  of the TAC shall appear last in the string.   | 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 

## Methods

### NewTai1

`func NewTai1(plmnId PlmnId1, tac string, ) *Tai1`

NewTai1 instantiates a new Tai1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTai1WithDefaults

`func NewTai1WithDefaults() *Tai1`

NewTai1WithDefaults instantiates a new Tai1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *Tai1) GetPlmnId() PlmnId1`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *Tai1) GetPlmnIdOk() (*PlmnId1, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *Tai1) SetPlmnId(v PlmnId1)`

SetPlmnId sets PlmnId field to given value.


### GetTac

`func (o *Tai1) GetTac() string`

GetTac returns the Tac field if non-nil, zero value otherwise.

### GetTacOk

`func (o *Tai1) GetTacOk() (*string, bool)`

GetTacOk returns a tuple with the Tac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTac

`func (o *Tai1) SetTac(v string)`

SetTac sets Tac field to given value.


### GetNid

`func (o *Tai1) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *Tai1) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *Tai1) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *Tai1) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


