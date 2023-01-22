# OperatorSpecificDataContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  | 
**DataTypeDefinition** | Pointer to **string** |  | [optional] 
**Value** | [**OperatorSpecificDataContainerValue**](OperatorSpecificDataContainerValue.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOperatorSpecificDataContainer

`func NewOperatorSpecificDataContainer(dataType string, value OperatorSpecificDataContainerValue, ) *OperatorSpecificDataContainer`

NewOperatorSpecificDataContainer instantiates a new OperatorSpecificDataContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorSpecificDataContainerWithDefaults

`func NewOperatorSpecificDataContainerWithDefaults() *OperatorSpecificDataContainer`

NewOperatorSpecificDataContainerWithDefaults instantiates a new OperatorSpecificDataContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *OperatorSpecificDataContainer) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *OperatorSpecificDataContainer) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *OperatorSpecificDataContainer) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDataTypeDefinition

`func (o *OperatorSpecificDataContainer) GetDataTypeDefinition() string`

GetDataTypeDefinition returns the DataTypeDefinition field if non-nil, zero value otherwise.

### GetDataTypeDefinitionOk

`func (o *OperatorSpecificDataContainer) GetDataTypeDefinitionOk() (*string, bool)`

GetDataTypeDefinitionOk returns a tuple with the DataTypeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypeDefinition

`func (o *OperatorSpecificDataContainer) SetDataTypeDefinition(v string)`

SetDataTypeDefinition sets DataTypeDefinition field to given value.

### HasDataTypeDefinition

`func (o *OperatorSpecificDataContainer) HasDataTypeDefinition() bool`

HasDataTypeDefinition returns a boolean if a field has been set.

### GetValue

`func (o *OperatorSpecificDataContainer) GetValue() OperatorSpecificDataContainerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OperatorSpecificDataContainer) GetValueOk() (*OperatorSpecificDataContainerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OperatorSpecificDataContainer) SetValue(v OperatorSpecificDataContainerValue)`

SetValue sets Value field to given value.


### GetSupportedFeatures

`func (o *OperatorSpecificDataContainer) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *OperatorSpecificDataContainer) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *OperatorSpecificDataContainer) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *OperatorSpecificDataContainer) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetResetIds

`func (o *OperatorSpecificDataContainer) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *OperatorSpecificDataContainer) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *OperatorSpecificDataContainer) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *OperatorSpecificDataContainer) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


