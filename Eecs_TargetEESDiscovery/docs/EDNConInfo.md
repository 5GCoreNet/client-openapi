# EDNConInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**EdnTopoSrvArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 

## Methods

### NewEDNConInfo

`func NewEDNConInfo() *EDNConInfo`

NewEDNConInfo instantiates a new EDNConInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEDNConInfoWithDefaults

`func NewEDNConInfoWithDefaults() *EDNConInfo`

NewEDNConInfoWithDefaults instantiates a new EDNConInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *EDNConInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *EDNConInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *EDNConInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *EDNConInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *EDNConInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *EDNConInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *EDNConInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *EDNConInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetEdnTopoSrvArea

`func (o *EDNConInfo) GetEdnTopoSrvArea() LocationArea5G`

GetEdnTopoSrvArea returns the EdnTopoSrvArea field if non-nil, zero value otherwise.

### GetEdnTopoSrvAreaOk

`func (o *EDNConInfo) GetEdnTopoSrvAreaOk() (*LocationArea5G, bool)`

GetEdnTopoSrvAreaOk returns a tuple with the EdnTopoSrvArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnTopoSrvArea

`func (o *EDNConInfo) SetEdnTopoSrvArea(v LocationArea5G)`

SetEdnTopoSrvArea sets EdnTopoSrvArea field to given value.

### HasEdnTopoSrvArea

`func (o *EDNConInfo) HasEdnTopoSrvArea() bool`

HasEdnTopoSrvArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


