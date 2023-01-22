# SmPolicyDnnDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**BdtRefIds** | Pointer to **map[string]string** | Contains updated transfer policies of background data transfer. Any string value can be used as a key of the map.  | [optional] 

## Methods

### NewSmPolicyDnnDataPatch

`func NewSmPolicyDnnDataPatch(dnn string, ) *SmPolicyDnnDataPatch`

NewSmPolicyDnnDataPatch instantiates a new SmPolicyDnnDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyDnnDataPatchWithDefaults

`func NewSmPolicyDnnDataPatchWithDefaults() *SmPolicyDnnDataPatch`

NewSmPolicyDnnDataPatchWithDefaults instantiates a new SmPolicyDnnDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *SmPolicyDnnDataPatch) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmPolicyDnnDataPatch) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmPolicyDnnDataPatch) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetBdtRefIds

`func (o *SmPolicyDnnDataPatch) GetBdtRefIds() map[string]string`

GetBdtRefIds returns the BdtRefIds field if non-nil, zero value otherwise.

### GetBdtRefIdsOk

`func (o *SmPolicyDnnDataPatch) GetBdtRefIdsOk() (*map[string]string, bool)`

GetBdtRefIdsOk returns a tuple with the BdtRefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefIds

`func (o *SmPolicyDnnDataPatch) SetBdtRefIds(v map[string]string)`

SetBdtRefIds sets BdtRefIds field to given value.

### HasBdtRefIds

`func (o *SmPolicyDnnDataPatch) HasBdtRefIds() bool`

HasBdtRefIds returns a boolean if a field has been set.

### SetBdtRefIdsNil

`func (o *SmPolicyDnnDataPatch) SetBdtRefIdsNil(b bool)`

 SetBdtRefIdsNil sets the value for BdtRefIds to be an explicit nil

### UnsetBdtRefIds
`func (o *SmPolicyDnnDataPatch) UnsetBdtRefIds()`

UnsetBdtRefIds ensures that no value is present for BdtRefIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


