# LcsMoData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedServiceClasses** | [**[]LcsMoServiceClass**](LcsMoServiceClass.md) |  | 
**MoAssistanceDataTypes** | Pointer to [**LcsBroadcastAssistanceTypesData**](LcsBroadcastAssistanceTypesData.md) |  | [optional] 

## Methods

### NewLcsMoData

`func NewLcsMoData(allowedServiceClasses []LcsMoServiceClass, ) *LcsMoData`

NewLcsMoData instantiates a new LcsMoData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLcsMoDataWithDefaults

`func NewLcsMoDataWithDefaults() *LcsMoData`

NewLcsMoDataWithDefaults instantiates a new LcsMoData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedServiceClasses

`func (o *LcsMoData) GetAllowedServiceClasses() []LcsMoServiceClass`

GetAllowedServiceClasses returns the AllowedServiceClasses field if non-nil, zero value otherwise.

### GetAllowedServiceClassesOk

`func (o *LcsMoData) GetAllowedServiceClassesOk() (*[]LcsMoServiceClass, bool)`

GetAllowedServiceClassesOk returns a tuple with the AllowedServiceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServiceClasses

`func (o *LcsMoData) SetAllowedServiceClasses(v []LcsMoServiceClass)`

SetAllowedServiceClasses sets AllowedServiceClasses field to given value.


### GetMoAssistanceDataTypes

`func (o *LcsMoData) GetMoAssistanceDataTypes() LcsBroadcastAssistanceTypesData`

GetMoAssistanceDataTypes returns the MoAssistanceDataTypes field if non-nil, zero value otherwise.

### GetMoAssistanceDataTypesOk

`func (o *LcsMoData) GetMoAssistanceDataTypesOk() (*LcsBroadcastAssistanceTypesData, bool)`

GetMoAssistanceDataTypesOk returns a tuple with the MoAssistanceDataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoAssistanceDataTypes

`func (o *LcsMoData) SetMoAssistanceDataTypes(v LcsBroadcastAssistanceTypesData)`

SetMoAssistanceDataTypes sets MoAssistanceDataTypes field to given value.

### HasMoAssistanceDataTypes

`func (o *LcsMoData) HasMoAssistanceDataTypes() bool`

HasMoAssistanceDataTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


