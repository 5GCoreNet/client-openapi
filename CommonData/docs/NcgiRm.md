# NcgiRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**NrCellId** | **string** | 36-bit string identifying an NR Cell Id as specified in clause 9.3.1.7 of 3GPP TS 38.413,  in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character  representing the 4 most significant bits of the Cell Id shall appear first in the string, and  the character representing the 4 least significant bit of the Cell Id shall appear last in the  string.   | 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 

## Methods

### NewNcgiRm

`func NewNcgiRm(plmnId PlmnId, nrCellId string, ) *NcgiRm`

NewNcgiRm instantiates a new NcgiRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNcgiRmWithDefaults

`func NewNcgiRmWithDefaults() *NcgiRm`

NewNcgiRmWithDefaults instantiates a new NcgiRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *NcgiRm) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *NcgiRm) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *NcgiRm) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetNrCellId

`func (o *NcgiRm) GetNrCellId() string`

GetNrCellId returns the NrCellId field if non-nil, zero value otherwise.

### GetNrCellIdOk

`func (o *NcgiRm) GetNrCellIdOk() (*string, bool)`

GetNrCellIdOk returns a tuple with the NrCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrCellId

`func (o *NcgiRm) SetNrCellId(v string)`

SetNrCellId sets NrCellId field to given value.


### GetNid

`func (o *NcgiRm) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *NcgiRm) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *NcgiRm) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *NcgiRm) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


