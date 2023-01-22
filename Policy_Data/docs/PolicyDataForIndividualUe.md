# PolicyDataForIndividualUe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UePolicyDataSet** | Pointer to [**UePolicySet**](UePolicySet.md) |  | [optional] 
**SmPolicyDataSet** | Pointer to [**SmPolicyData**](SmPolicyData.md) |  | [optional] 
**AmPolicyDataSet** | Pointer to [**AmPolicyData**](AmPolicyData.md) |  | [optional] 
**UmData** | Pointer to [**map[string]UsageMonData**](UsageMonData.md) | Contains UM policies. The value of the limit identifier is used as the key of the map.  | [optional] 
**OperatorSpecificDataSet** | Pointer to [**map[string]OperatorSpecificDataContainer**](OperatorSpecificDataContainer.md) | Contains Operator Specific Data resource data. The key of the map is operator specific data element name and the value is the operator specific data of the UE.  | [optional] 

## Methods

### NewPolicyDataForIndividualUe

`func NewPolicyDataForIndividualUe() *PolicyDataForIndividualUe`

NewPolicyDataForIndividualUe instantiates a new PolicyDataForIndividualUe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDataForIndividualUeWithDefaults

`func NewPolicyDataForIndividualUeWithDefaults() *PolicyDataForIndividualUe`

NewPolicyDataForIndividualUeWithDefaults instantiates a new PolicyDataForIndividualUe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUePolicyDataSet

`func (o *PolicyDataForIndividualUe) GetUePolicyDataSet() UePolicySet`

GetUePolicyDataSet returns the UePolicyDataSet field if non-nil, zero value otherwise.

### GetUePolicyDataSetOk

`func (o *PolicyDataForIndividualUe) GetUePolicyDataSetOk() (*UePolicySet, bool)`

GetUePolicyDataSetOk returns a tuple with the UePolicyDataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolicyDataSet

`func (o *PolicyDataForIndividualUe) SetUePolicyDataSet(v UePolicySet)`

SetUePolicyDataSet sets UePolicyDataSet field to given value.

### HasUePolicyDataSet

`func (o *PolicyDataForIndividualUe) HasUePolicyDataSet() bool`

HasUePolicyDataSet returns a boolean if a field has been set.

### GetSmPolicyDataSet

`func (o *PolicyDataForIndividualUe) GetSmPolicyDataSet() SmPolicyData`

GetSmPolicyDataSet returns the SmPolicyDataSet field if non-nil, zero value otherwise.

### GetSmPolicyDataSetOk

`func (o *PolicyDataForIndividualUe) GetSmPolicyDataSetOk() (*SmPolicyData, bool)`

GetSmPolicyDataSetOk returns a tuple with the SmPolicyDataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyDataSet

`func (o *PolicyDataForIndividualUe) SetSmPolicyDataSet(v SmPolicyData)`

SetSmPolicyDataSet sets SmPolicyDataSet field to given value.

### HasSmPolicyDataSet

`func (o *PolicyDataForIndividualUe) HasSmPolicyDataSet() bool`

HasSmPolicyDataSet returns a boolean if a field has been set.

### GetAmPolicyDataSet

`func (o *PolicyDataForIndividualUe) GetAmPolicyDataSet() AmPolicyData`

GetAmPolicyDataSet returns the AmPolicyDataSet field if non-nil, zero value otherwise.

### GetAmPolicyDataSetOk

`func (o *PolicyDataForIndividualUe) GetAmPolicyDataSetOk() (*AmPolicyData, bool)`

GetAmPolicyDataSetOk returns a tuple with the AmPolicyDataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmPolicyDataSet

`func (o *PolicyDataForIndividualUe) SetAmPolicyDataSet(v AmPolicyData)`

SetAmPolicyDataSet sets AmPolicyDataSet field to given value.

### HasAmPolicyDataSet

`func (o *PolicyDataForIndividualUe) HasAmPolicyDataSet() bool`

HasAmPolicyDataSet returns a boolean if a field has been set.

### GetUmData

`func (o *PolicyDataForIndividualUe) GetUmData() map[string]UsageMonData`

GetUmData returns the UmData field if non-nil, zero value otherwise.

### GetUmDataOk

`func (o *PolicyDataForIndividualUe) GetUmDataOk() (*map[string]UsageMonData, bool)`

GetUmDataOk returns a tuple with the UmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmData

`func (o *PolicyDataForIndividualUe) SetUmData(v map[string]UsageMonData)`

SetUmData sets UmData field to given value.

### HasUmData

`func (o *PolicyDataForIndividualUe) HasUmData() bool`

HasUmData returns a boolean if a field has been set.

### GetOperatorSpecificDataSet

`func (o *PolicyDataForIndividualUe) GetOperatorSpecificDataSet() map[string]OperatorSpecificDataContainer`

GetOperatorSpecificDataSet returns the OperatorSpecificDataSet field if non-nil, zero value otherwise.

### GetOperatorSpecificDataSetOk

`func (o *PolicyDataForIndividualUe) GetOperatorSpecificDataSetOk() (*map[string]OperatorSpecificDataContainer, bool)`

GetOperatorSpecificDataSetOk returns a tuple with the OperatorSpecificDataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorSpecificDataSet

`func (o *PolicyDataForIndividualUe) SetOperatorSpecificDataSet(v map[string]OperatorSpecificDataContainer)`

SetOperatorSpecificDataSet sets OperatorSpecificDataSet field to given value.

### HasOperatorSpecificDataSet

`func (o *PolicyDataForIndividualUe) HasOperatorSpecificDataSet() bool`

HasOperatorSpecificDataSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


