# APIProviderEnrolmentDetailsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiProvFuncs** | Pointer to [**[]APIProviderFunctionDetails**](APIProviderFunctionDetails.md) | A list of individual API provider domain functions details. When included by the API management function in the HTTP request message, it lists the API provider domain functions that the API management function intends to register/update in registration or update registration procedure.   | [optional] 
**ApiProvDomInfo** | Pointer to **string** | Generic information related to the API provider domain such as details of the API provider applications.  | [optional] 

## Methods

### NewAPIProviderEnrolmentDetailsPatch

`func NewAPIProviderEnrolmentDetailsPatch() *APIProviderEnrolmentDetailsPatch`

NewAPIProviderEnrolmentDetailsPatch instantiates a new APIProviderEnrolmentDetailsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIProviderEnrolmentDetailsPatchWithDefaults

`func NewAPIProviderEnrolmentDetailsPatchWithDefaults() *APIProviderEnrolmentDetailsPatch`

NewAPIProviderEnrolmentDetailsPatchWithDefaults instantiates a new APIProviderEnrolmentDetailsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiProvFuncs

`func (o *APIProviderEnrolmentDetailsPatch) GetApiProvFuncs() []APIProviderFunctionDetails`

GetApiProvFuncs returns the ApiProvFuncs field if non-nil, zero value otherwise.

### GetApiProvFuncsOk

`func (o *APIProviderEnrolmentDetailsPatch) GetApiProvFuncsOk() (*[]APIProviderFunctionDetails, bool)`

GetApiProvFuncsOk returns a tuple with the ApiProvFuncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvFuncs

`func (o *APIProviderEnrolmentDetailsPatch) SetApiProvFuncs(v []APIProviderFunctionDetails)`

SetApiProvFuncs sets ApiProvFuncs field to given value.

### HasApiProvFuncs

`func (o *APIProviderEnrolmentDetailsPatch) HasApiProvFuncs() bool`

HasApiProvFuncs returns a boolean if a field has been set.

### GetApiProvDomInfo

`func (o *APIProviderEnrolmentDetailsPatch) GetApiProvDomInfo() string`

GetApiProvDomInfo returns the ApiProvDomInfo field if non-nil, zero value otherwise.

### GetApiProvDomInfoOk

`func (o *APIProviderEnrolmentDetailsPatch) GetApiProvDomInfoOk() (*string, bool)`

GetApiProvDomInfoOk returns a tuple with the ApiProvDomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvDomInfo

`func (o *APIProviderEnrolmentDetailsPatch) SetApiProvDomInfo(v string)`

SetApiProvDomInfo sets ApiProvDomInfo field to given value.

### HasApiProvDomInfo

`func (o *APIProviderEnrolmentDetailsPatch) HasApiProvDomInfo() bool`

HasApiProvDomInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


