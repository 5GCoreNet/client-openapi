# GlobalRanNodeId1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId1**](PlmnId1.md) |  | 
**N3IwfId** | Pointer to **string** | This represents the identifier of the N3IWF ID as specified in clause 9.3.1.57 of  3GPP TS 38.413 in hexadecimal representation. Each character in the string shall take a value  of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant  character representing the 4 most significant bits of the N3IWF ID shall appear first in the  string, and the character representing the 4 least significant bit of the N3IWF ID shall  appear last in the string.   | [optional] 
**GNbId** | Pointer to [**GNbId**](GNbId.md) |  | [optional] 
**NgeNbId** | Pointer to **string** | This represents the identifier of the ng-eNB ID as specified in clause 9.3.1.8 of  3GPP TS 38.413. The value of the ng-eNB ID shall be encoded in hexadecimal representation.  Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and  shall represent 4 bits. The padding 0 shall be added to make multiple nibbles, so the most  significant character representing the padding 0 if required together with the 4 most  significant bits of the ng-eNB ID shall appear first in the string, and the character  representing the 4 least significant bit of the ng-eNB ID (to form a nibble) shall appear last  in the string.   | [optional] 
**WagfId** | Pointer to **string** | This represents the identifier of the W-AGF ID as specified in clause 9.3.1.162 of  3GPP TS 38.413 in hexadecimal representation. Each character in the string shall take a value  of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant  character representing the 4 most significant bits of the W-AGF ID shall appear first in the  string, and the character representing the 4 least significant bit of the W-AGF ID shall  appear last in the string.   | [optional] 
**TngfId** | Pointer to **string** | This represents the identifier of the TNGF ID as specified in clause 9.3.1.161 of  3GPP TS 38.413  in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot;  to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the  4 most significant bits of the TNGF ID shall appear first in the string, and the character  representing the 4 least significant bit of the TNGF ID shall appear last in the string.   | [optional] 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 
**ENbId** | Pointer to **string** | This represents the identifier of the eNB ID as specified in clause 9.2.1.37 of  3GPP TS 36.413. The string shall be formatted with the following pattern  &#39;^(&#39;MacroeNB-[A-Fa-f0-9]{5}|LMacroeNB-[A-Fa-f0-9]{6}|SMacroeNB-[A-Fa-f0-9]{5} |HomeeNB-[A-Fa-f0-9]{7})$&#39;. The value of the eNB ID shall be encoded in hexadecimal representation. Each character in the  string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits.  The padding 0 shall be added to make multiple nibbles, so the most significant character  representing the padding 0 if required together with the 4 most significant bits of the eNB ID  shall appear first in the string, and the character representing the 4 least significant bit  of the eNB ID (to form a nibble) shall appear last in the string.  | [optional] 

## Methods

### NewGlobalRanNodeId1

`func NewGlobalRanNodeId1(plmnId PlmnId1, ) *GlobalRanNodeId1`

NewGlobalRanNodeId1 instantiates a new GlobalRanNodeId1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalRanNodeId1WithDefaults

`func NewGlobalRanNodeId1WithDefaults() *GlobalRanNodeId1`

NewGlobalRanNodeId1WithDefaults instantiates a new GlobalRanNodeId1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *GlobalRanNodeId1) GetPlmnId() PlmnId1`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *GlobalRanNodeId1) GetPlmnIdOk() (*PlmnId1, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *GlobalRanNodeId1) SetPlmnId(v PlmnId1)`

SetPlmnId sets PlmnId field to given value.


### GetN3IwfId

`func (o *GlobalRanNodeId1) GetN3IwfId() string`

GetN3IwfId returns the N3IwfId field if non-nil, zero value otherwise.

### GetN3IwfIdOk

`func (o *GlobalRanNodeId1) GetN3IwfIdOk() (*string, bool)`

GetN3IwfIdOk returns a tuple with the N3IwfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3IwfId

`func (o *GlobalRanNodeId1) SetN3IwfId(v string)`

SetN3IwfId sets N3IwfId field to given value.

### HasN3IwfId

`func (o *GlobalRanNodeId1) HasN3IwfId() bool`

HasN3IwfId returns a boolean if a field has been set.

### GetGNbId

`func (o *GlobalRanNodeId1) GetGNbId() GNbId`

GetGNbId returns the GNbId field if non-nil, zero value otherwise.

### GetGNbIdOk

`func (o *GlobalRanNodeId1) GetGNbIdOk() (*GNbId, bool)`

GetGNbIdOk returns a tuple with the GNbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNbId

`func (o *GlobalRanNodeId1) SetGNbId(v GNbId)`

SetGNbId sets GNbId field to given value.

### HasGNbId

`func (o *GlobalRanNodeId1) HasGNbId() bool`

HasGNbId returns a boolean if a field has been set.

### GetNgeNbId

`func (o *GlobalRanNodeId1) GetNgeNbId() string`

GetNgeNbId returns the NgeNbId field if non-nil, zero value otherwise.

### GetNgeNbIdOk

`func (o *GlobalRanNodeId1) GetNgeNbIdOk() (*string, bool)`

GetNgeNbIdOk returns a tuple with the NgeNbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgeNbId

`func (o *GlobalRanNodeId1) SetNgeNbId(v string)`

SetNgeNbId sets NgeNbId field to given value.

### HasNgeNbId

`func (o *GlobalRanNodeId1) HasNgeNbId() bool`

HasNgeNbId returns a boolean if a field has been set.

### GetWagfId

`func (o *GlobalRanNodeId1) GetWagfId() string`

GetWagfId returns the WagfId field if non-nil, zero value otherwise.

### GetWagfIdOk

`func (o *GlobalRanNodeId1) GetWagfIdOk() (*string, bool)`

GetWagfIdOk returns a tuple with the WagfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagfId

`func (o *GlobalRanNodeId1) SetWagfId(v string)`

SetWagfId sets WagfId field to given value.

### HasWagfId

`func (o *GlobalRanNodeId1) HasWagfId() bool`

HasWagfId returns a boolean if a field has been set.

### GetTngfId

`func (o *GlobalRanNodeId1) GetTngfId() string`

GetTngfId returns the TngfId field if non-nil, zero value otherwise.

### GetTngfIdOk

`func (o *GlobalRanNodeId1) GetTngfIdOk() (*string, bool)`

GetTngfIdOk returns a tuple with the TngfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfId

`func (o *GlobalRanNodeId1) SetTngfId(v string)`

SetTngfId sets TngfId field to given value.

### HasTngfId

`func (o *GlobalRanNodeId1) HasTngfId() bool`

HasTngfId returns a boolean if a field has been set.

### GetNid

`func (o *GlobalRanNodeId1) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *GlobalRanNodeId1) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *GlobalRanNodeId1) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *GlobalRanNodeId1) HasNid() bool`

HasNid returns a boolean if a field has been set.

### GetENbId

`func (o *GlobalRanNodeId1) GetENbId() string`

GetENbId returns the ENbId field if non-nil, zero value otherwise.

### GetENbIdOk

`func (o *GlobalRanNodeId1) GetENbIdOk() (*string, bool)`

GetENbIdOk returns a tuple with the ENbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENbId

`func (o *GlobalRanNodeId1) SetENbId(v string)`

SetENbId sets ENbId field to given value.

### HasENbId

`func (o *GlobalRanNodeId1) HasENbId() bool`

HasENbId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


