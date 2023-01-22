# MbsExtProblemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Instance** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Cause** | Pointer to **string** | A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available.  | [optional] 
**InvalidParams** | Pointer to [**[]InvalidParam**](InvalidParam.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AccessTokenError** | Pointer to [**AccessTokenErr**](AccessTokenErr.md) |  | [optional] 
**AccessTokenRequest** | Pointer to [**AccessTokenReq**](AccessTokenReq.md) |  | [optional] 
**NrfId** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewMbsExtProblemDetails

`func NewMbsExtProblemDetails() *MbsExtProblemDetails`

NewMbsExtProblemDetails instantiates a new MbsExtProblemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsExtProblemDetailsWithDefaults

`func NewMbsExtProblemDetailsWithDefaults() *MbsExtProblemDetails`

NewMbsExtProblemDetailsWithDefaults instantiates a new MbsExtProblemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MbsExtProblemDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MbsExtProblemDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MbsExtProblemDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MbsExtProblemDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *MbsExtProblemDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MbsExtProblemDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MbsExtProblemDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MbsExtProblemDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *MbsExtProblemDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MbsExtProblemDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MbsExtProblemDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MbsExtProblemDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *MbsExtProblemDetails) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *MbsExtProblemDetails) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *MbsExtProblemDetails) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *MbsExtProblemDetails) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *MbsExtProblemDetails) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *MbsExtProblemDetails) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *MbsExtProblemDetails) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *MbsExtProblemDetails) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *MbsExtProblemDetails) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *MbsExtProblemDetails) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *MbsExtProblemDetails) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *MbsExtProblemDetails) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *MbsExtProblemDetails) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *MbsExtProblemDetails) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *MbsExtProblemDetails) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *MbsExtProblemDetails) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *MbsExtProblemDetails) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *MbsExtProblemDetails) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *MbsExtProblemDetails) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *MbsExtProblemDetails) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAccessTokenError

`func (o *MbsExtProblemDetails) GetAccessTokenError() AccessTokenErr`

GetAccessTokenError returns the AccessTokenError field if non-nil, zero value otherwise.

### GetAccessTokenErrorOk

`func (o *MbsExtProblemDetails) GetAccessTokenErrorOk() (*AccessTokenErr, bool)`

GetAccessTokenErrorOk returns a tuple with the AccessTokenError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenError

`func (o *MbsExtProblemDetails) SetAccessTokenError(v AccessTokenErr)`

SetAccessTokenError sets AccessTokenError field to given value.

### HasAccessTokenError

`func (o *MbsExtProblemDetails) HasAccessTokenError() bool`

HasAccessTokenError returns a boolean if a field has been set.

### GetAccessTokenRequest

`func (o *MbsExtProblemDetails) GetAccessTokenRequest() AccessTokenReq`

GetAccessTokenRequest returns the AccessTokenRequest field if non-nil, zero value otherwise.

### GetAccessTokenRequestOk

`func (o *MbsExtProblemDetails) GetAccessTokenRequestOk() (*AccessTokenReq, bool)`

GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRequest

`func (o *MbsExtProblemDetails) SetAccessTokenRequest(v AccessTokenReq)`

SetAccessTokenRequest sets AccessTokenRequest field to given value.

### HasAccessTokenRequest

`func (o *MbsExtProblemDetails) HasAccessTokenRequest() bool`

HasAccessTokenRequest returns a boolean if a field has been set.

### GetNrfId

`func (o *MbsExtProblemDetails) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *MbsExtProblemDetails) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *MbsExtProblemDetails) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.

### HasNrfId

`func (o *MbsExtProblemDetails) HasNrfId() bool`

HasNrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


