# APIProviderEnrolmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiProvDomId** | Pointer to **string** | API provider domain ID assigned by the CAPIF core function to the API management function while registering the API provider domain. Shall not be present in the HTTP POST request from the API Management function to the CAPIF core function, to on-board itself. Shall be present in all other HTTP requests and responses.  | [optional] [readonly] 
**RegSec** | **string** | Security information necessary for the CAPIF core function to validate the registration of the API provider domain. Shall be present in HTTP POST request from API management function to CAPIF core function for API provider domain registration.  | 
**ApiProvFuncs** | Pointer to [**[]APIProviderFunctionDetails**](APIProviderFunctionDetails.md) | A list of individual API provider domain functions details. When included by the API management function in the HTTP request message, it lists the API provider domain functions that the API management function intends to register/update in registration or update registration procedure. When included by the CAPIF core function in the HTTP response message, it lists the API domain functions details that are registered or updated successfully.  | [optional] 
**ApiProvDomInfo** | Pointer to **string** | Generic information related to the API provider domain such as details of the API provider applications.  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**FailReason** | Pointer to **string** | Registration or update specific failure information of failed API provider domain function registrations.Shall be present in the HTTP response body if atleast one of the API provider domain function registration or update registration fails.  | [optional] 

## Methods

### NewAPIProviderEnrolmentDetails

`func NewAPIProviderEnrolmentDetails(regSec string, ) *APIProviderEnrolmentDetails`

NewAPIProviderEnrolmentDetails instantiates a new APIProviderEnrolmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIProviderEnrolmentDetailsWithDefaults

`func NewAPIProviderEnrolmentDetailsWithDefaults() *APIProviderEnrolmentDetails`

NewAPIProviderEnrolmentDetailsWithDefaults instantiates a new APIProviderEnrolmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiProvDomId

`func (o *APIProviderEnrolmentDetails) GetApiProvDomId() string`

GetApiProvDomId returns the ApiProvDomId field if non-nil, zero value otherwise.

### GetApiProvDomIdOk

`func (o *APIProviderEnrolmentDetails) GetApiProvDomIdOk() (*string, bool)`

GetApiProvDomIdOk returns a tuple with the ApiProvDomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvDomId

`func (o *APIProviderEnrolmentDetails) SetApiProvDomId(v string)`

SetApiProvDomId sets ApiProvDomId field to given value.

### HasApiProvDomId

`func (o *APIProviderEnrolmentDetails) HasApiProvDomId() bool`

HasApiProvDomId returns a boolean if a field has been set.

### GetRegSec

`func (o *APIProviderEnrolmentDetails) GetRegSec() string`

GetRegSec returns the RegSec field if non-nil, zero value otherwise.

### GetRegSecOk

`func (o *APIProviderEnrolmentDetails) GetRegSecOk() (*string, bool)`

GetRegSecOk returns a tuple with the RegSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegSec

`func (o *APIProviderEnrolmentDetails) SetRegSec(v string)`

SetRegSec sets RegSec field to given value.


### GetApiProvFuncs

`func (o *APIProviderEnrolmentDetails) GetApiProvFuncs() []APIProviderFunctionDetails`

GetApiProvFuncs returns the ApiProvFuncs field if non-nil, zero value otherwise.

### GetApiProvFuncsOk

`func (o *APIProviderEnrolmentDetails) GetApiProvFuncsOk() (*[]APIProviderFunctionDetails, bool)`

GetApiProvFuncsOk returns a tuple with the ApiProvFuncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvFuncs

`func (o *APIProviderEnrolmentDetails) SetApiProvFuncs(v []APIProviderFunctionDetails)`

SetApiProvFuncs sets ApiProvFuncs field to given value.

### HasApiProvFuncs

`func (o *APIProviderEnrolmentDetails) HasApiProvFuncs() bool`

HasApiProvFuncs returns a boolean if a field has been set.

### GetApiProvDomInfo

`func (o *APIProviderEnrolmentDetails) GetApiProvDomInfo() string`

GetApiProvDomInfo returns the ApiProvDomInfo field if non-nil, zero value otherwise.

### GetApiProvDomInfoOk

`func (o *APIProviderEnrolmentDetails) GetApiProvDomInfoOk() (*string, bool)`

GetApiProvDomInfoOk returns a tuple with the ApiProvDomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvDomInfo

`func (o *APIProviderEnrolmentDetails) SetApiProvDomInfo(v string)`

SetApiProvDomInfo sets ApiProvDomInfo field to given value.

### HasApiProvDomInfo

`func (o *APIProviderEnrolmentDetails) HasApiProvDomInfo() bool`

HasApiProvDomInfo returns a boolean if a field has been set.

### GetSuppFeat

`func (o *APIProviderEnrolmentDetails) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *APIProviderEnrolmentDetails) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *APIProviderEnrolmentDetails) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *APIProviderEnrolmentDetails) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetFailReason

`func (o *APIProviderEnrolmentDetails) GetFailReason() string`

GetFailReason returns the FailReason field if non-nil, zero value otherwise.

### GetFailReasonOk

`func (o *APIProviderEnrolmentDetails) GetFailReasonOk() (*string, bool)`

GetFailReasonOk returns a tuple with the FailReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailReason

`func (o *APIProviderEnrolmentDetails) SetFailReason(v string)`

SetFailReason sets FailReason field to given value.

### HasFailReason

`func (o *APIProviderEnrolmentDetails) HasFailReason() bool`

HasFailReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


