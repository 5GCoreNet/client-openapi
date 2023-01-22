# APIProviderFunctionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiProvFuncId** | Pointer to **string** | API provider domain functionID assigned by the CAPIF core function to the API provider domain function while registering/updating the API provider domain. Shall not be present in the HTTP POST request from the API management function to the CAPIF core function, to register itself. Shall be present in all other HTTP requests and responses.  | [optional] 
**RegInfo** | [**RegistrationInformation**](RegistrationInformation.md) |  | 
**ApiProvFuncRole** | [**ApiProviderFuncRole**](ApiProviderFuncRole.md) |  | 
**ApiProvFuncInfo** | Pointer to **string** | Generic information related to the API provider domain function such as details of the API provider applications.  | [optional] 

## Methods

### NewAPIProviderFunctionDetails

`func NewAPIProviderFunctionDetails(regInfo RegistrationInformation, apiProvFuncRole ApiProviderFuncRole, ) *APIProviderFunctionDetails`

NewAPIProviderFunctionDetails instantiates a new APIProviderFunctionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIProviderFunctionDetailsWithDefaults

`func NewAPIProviderFunctionDetailsWithDefaults() *APIProviderFunctionDetails`

NewAPIProviderFunctionDetailsWithDefaults instantiates a new APIProviderFunctionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiProvFuncId

`func (o *APIProviderFunctionDetails) GetApiProvFuncId() string`

GetApiProvFuncId returns the ApiProvFuncId field if non-nil, zero value otherwise.

### GetApiProvFuncIdOk

`func (o *APIProviderFunctionDetails) GetApiProvFuncIdOk() (*string, bool)`

GetApiProvFuncIdOk returns a tuple with the ApiProvFuncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvFuncId

`func (o *APIProviderFunctionDetails) SetApiProvFuncId(v string)`

SetApiProvFuncId sets ApiProvFuncId field to given value.

### HasApiProvFuncId

`func (o *APIProviderFunctionDetails) HasApiProvFuncId() bool`

HasApiProvFuncId returns a boolean if a field has been set.

### GetRegInfo

`func (o *APIProviderFunctionDetails) GetRegInfo() RegistrationInformation`

GetRegInfo returns the RegInfo field if non-nil, zero value otherwise.

### GetRegInfoOk

`func (o *APIProviderFunctionDetails) GetRegInfoOk() (*RegistrationInformation, bool)`

GetRegInfoOk returns a tuple with the RegInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegInfo

`func (o *APIProviderFunctionDetails) SetRegInfo(v RegistrationInformation)`

SetRegInfo sets RegInfo field to given value.


### GetApiProvFuncRole

`func (o *APIProviderFunctionDetails) GetApiProvFuncRole() ApiProviderFuncRole`

GetApiProvFuncRole returns the ApiProvFuncRole field if non-nil, zero value otherwise.

### GetApiProvFuncRoleOk

`func (o *APIProviderFunctionDetails) GetApiProvFuncRoleOk() (*ApiProviderFuncRole, bool)`

GetApiProvFuncRoleOk returns a tuple with the ApiProvFuncRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvFuncRole

`func (o *APIProviderFunctionDetails) SetApiProvFuncRole(v ApiProviderFuncRole)`

SetApiProvFuncRole sets ApiProvFuncRole field to given value.


### GetApiProvFuncInfo

`func (o *APIProviderFunctionDetails) GetApiProvFuncInfo() string`

GetApiProvFuncInfo returns the ApiProvFuncInfo field if non-nil, zero value otherwise.

### GetApiProvFuncInfoOk

`func (o *APIProviderFunctionDetails) GetApiProvFuncInfoOk() (*string, bool)`

GetApiProvFuncInfoOk returns a tuple with the ApiProvFuncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvFuncInfo

`func (o *APIProviderFunctionDetails) SetApiProvFuncInfo(v string)`

SetApiProvFuncInfo sets ApiProvFuncInfo field to given value.

### HasApiProvFuncInfo

`func (o *APIProviderFunctionDetails) HasApiProvFuncInfo() bool`

HasApiProvFuncInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


