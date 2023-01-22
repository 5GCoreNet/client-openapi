# ChargingdataPost400Response

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
**InvocationTimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**InvocationSequenceNumber** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**InvocationResult** | Pointer to [**InvocationResult**](InvocationResult.md) |  | [optional] 
**SessionFailover** | Pointer to [**SessionFailover**](SessionFailover.md) |  | [optional] 
**MultipleUnitInformation** | Pointer to [**[]MultipleUnitInformation**](MultipleUnitInformation.md) |  | [optional] 
**Triggers** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**PDUSessionChargingInformation** | Pointer to [**PDUSessionChargingInformation**](PDUSessionChargingInformation.md) |  | [optional] 
**RoamingQBCInformation** | Pointer to [**RoamingQBCInformation**](RoamingQBCInformation.md) |  | [optional] 
**LocationReportingChargingInformation** | Pointer to [**LocationReportingChargingInformation**](LocationReportingChargingInformation.md) |  | [optional] 

## Methods

### NewChargingdataPost400Response

`func NewChargingdataPost400Response(invocationTimeStamp time.Time, invocationSequenceNumber int32, ) *ChargingdataPost400Response`

NewChargingdataPost400Response instantiates a new ChargingdataPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingdataPost400ResponseWithDefaults

`func NewChargingdataPost400ResponseWithDefaults() *ChargingdataPost400Response`

NewChargingdataPost400ResponseWithDefaults instantiates a new ChargingdataPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChargingdataPost400Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargingdataPost400Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargingdataPost400Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChargingdataPost400Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *ChargingdataPost400Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChargingdataPost400Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChargingdataPost400Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChargingdataPost400Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *ChargingdataPost400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargingdataPost400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargingdataPost400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargingdataPost400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *ChargingdataPost400Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ChargingdataPost400Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ChargingdataPost400Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ChargingdataPost400Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *ChargingdataPost400Response) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ChargingdataPost400Response) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ChargingdataPost400Response) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ChargingdataPost400Response) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *ChargingdataPost400Response) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ChargingdataPost400Response) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ChargingdataPost400Response) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ChargingdataPost400Response) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *ChargingdataPost400Response) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *ChargingdataPost400Response) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *ChargingdataPost400Response) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *ChargingdataPost400Response) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ChargingdataPost400Response) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ChargingdataPost400Response) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ChargingdataPost400Response) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ChargingdataPost400Response) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAccessTokenError

`func (o *ChargingdataPost400Response) GetAccessTokenError() AccessTokenErr`

GetAccessTokenError returns the AccessTokenError field if non-nil, zero value otherwise.

### GetAccessTokenErrorOk

`func (o *ChargingdataPost400Response) GetAccessTokenErrorOk() (*AccessTokenErr, bool)`

GetAccessTokenErrorOk returns a tuple with the AccessTokenError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenError

`func (o *ChargingdataPost400Response) SetAccessTokenError(v AccessTokenErr)`

SetAccessTokenError sets AccessTokenError field to given value.

### HasAccessTokenError

`func (o *ChargingdataPost400Response) HasAccessTokenError() bool`

HasAccessTokenError returns a boolean if a field has been set.

### GetAccessTokenRequest

`func (o *ChargingdataPost400Response) GetAccessTokenRequest() AccessTokenReq`

GetAccessTokenRequest returns the AccessTokenRequest field if non-nil, zero value otherwise.

### GetAccessTokenRequestOk

`func (o *ChargingdataPost400Response) GetAccessTokenRequestOk() (*AccessTokenReq, bool)`

GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRequest

`func (o *ChargingdataPost400Response) SetAccessTokenRequest(v AccessTokenReq)`

SetAccessTokenRequest sets AccessTokenRequest field to given value.

### HasAccessTokenRequest

`func (o *ChargingdataPost400Response) HasAccessTokenRequest() bool`

HasAccessTokenRequest returns a boolean if a field has been set.

### GetNrfId

`func (o *ChargingdataPost400Response) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *ChargingdataPost400Response) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *ChargingdataPost400Response) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.

### HasNrfId

`func (o *ChargingdataPost400Response) HasNrfId() bool`

HasNrfId returns a boolean if a field has been set.

### GetInvocationTimeStamp

`func (o *ChargingdataPost400Response) GetInvocationTimeStamp() time.Time`

GetInvocationTimeStamp returns the InvocationTimeStamp field if non-nil, zero value otherwise.

### GetInvocationTimeStampOk

`func (o *ChargingdataPost400Response) GetInvocationTimeStampOk() (*time.Time, bool)`

GetInvocationTimeStampOk returns a tuple with the InvocationTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationTimeStamp

`func (o *ChargingdataPost400Response) SetInvocationTimeStamp(v time.Time)`

SetInvocationTimeStamp sets InvocationTimeStamp field to given value.


### GetInvocationSequenceNumber

`func (o *ChargingdataPost400Response) GetInvocationSequenceNumber() int32`

GetInvocationSequenceNumber returns the InvocationSequenceNumber field if non-nil, zero value otherwise.

### GetInvocationSequenceNumberOk

`func (o *ChargingdataPost400Response) GetInvocationSequenceNumberOk() (*int32, bool)`

GetInvocationSequenceNumberOk returns a tuple with the InvocationSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationSequenceNumber

`func (o *ChargingdataPost400Response) SetInvocationSequenceNumber(v int32)`

SetInvocationSequenceNumber sets InvocationSequenceNumber field to given value.


### GetInvocationResult

`func (o *ChargingdataPost400Response) GetInvocationResult() InvocationResult`

GetInvocationResult returns the InvocationResult field if non-nil, zero value otherwise.

### GetInvocationResultOk

`func (o *ChargingdataPost400Response) GetInvocationResultOk() (*InvocationResult, bool)`

GetInvocationResultOk returns a tuple with the InvocationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationResult

`func (o *ChargingdataPost400Response) SetInvocationResult(v InvocationResult)`

SetInvocationResult sets InvocationResult field to given value.

### HasInvocationResult

`func (o *ChargingdataPost400Response) HasInvocationResult() bool`

HasInvocationResult returns a boolean if a field has been set.

### GetSessionFailover

`func (o *ChargingdataPost400Response) GetSessionFailover() SessionFailover`

GetSessionFailover returns the SessionFailover field if non-nil, zero value otherwise.

### GetSessionFailoverOk

`func (o *ChargingdataPost400Response) GetSessionFailoverOk() (*SessionFailover, bool)`

GetSessionFailoverOk returns a tuple with the SessionFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionFailover

`func (o *ChargingdataPost400Response) SetSessionFailover(v SessionFailover)`

SetSessionFailover sets SessionFailover field to given value.

### HasSessionFailover

`func (o *ChargingdataPost400Response) HasSessionFailover() bool`

HasSessionFailover returns a boolean if a field has been set.

### GetMultipleUnitInformation

`func (o *ChargingdataPost400Response) GetMultipleUnitInformation() []MultipleUnitInformation`

GetMultipleUnitInformation returns the MultipleUnitInformation field if non-nil, zero value otherwise.

### GetMultipleUnitInformationOk

`func (o *ChargingdataPost400Response) GetMultipleUnitInformationOk() (*[]MultipleUnitInformation, bool)`

GetMultipleUnitInformationOk returns a tuple with the MultipleUnitInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleUnitInformation

`func (o *ChargingdataPost400Response) SetMultipleUnitInformation(v []MultipleUnitInformation)`

SetMultipleUnitInformation sets MultipleUnitInformation field to given value.

### HasMultipleUnitInformation

`func (o *ChargingdataPost400Response) HasMultipleUnitInformation() bool`

HasMultipleUnitInformation returns a boolean if a field has been set.

### GetTriggers

`func (o *ChargingdataPost400Response) GetTriggers() []Trigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ChargingdataPost400Response) GetTriggersOk() (*[]Trigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ChargingdataPost400Response) SetTriggers(v []Trigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ChargingdataPost400Response) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetPDUSessionChargingInformation

`func (o *ChargingdataPost400Response) GetPDUSessionChargingInformation() PDUSessionChargingInformation`

GetPDUSessionChargingInformation returns the PDUSessionChargingInformation field if non-nil, zero value otherwise.

### GetPDUSessionChargingInformationOk

`func (o *ChargingdataPost400Response) GetPDUSessionChargingInformationOk() (*PDUSessionChargingInformation, bool)`

GetPDUSessionChargingInformationOk returns a tuple with the PDUSessionChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPDUSessionChargingInformation

`func (o *ChargingdataPost400Response) SetPDUSessionChargingInformation(v PDUSessionChargingInformation)`

SetPDUSessionChargingInformation sets PDUSessionChargingInformation field to given value.

### HasPDUSessionChargingInformation

`func (o *ChargingdataPost400Response) HasPDUSessionChargingInformation() bool`

HasPDUSessionChargingInformation returns a boolean if a field has been set.

### GetRoamingQBCInformation

`func (o *ChargingdataPost400Response) GetRoamingQBCInformation() RoamingQBCInformation`

GetRoamingQBCInformation returns the RoamingQBCInformation field if non-nil, zero value otherwise.

### GetRoamingQBCInformationOk

`func (o *ChargingdataPost400Response) GetRoamingQBCInformationOk() (*RoamingQBCInformation, bool)`

GetRoamingQBCInformationOk returns a tuple with the RoamingQBCInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingQBCInformation

`func (o *ChargingdataPost400Response) SetRoamingQBCInformation(v RoamingQBCInformation)`

SetRoamingQBCInformation sets RoamingQBCInformation field to given value.

### HasRoamingQBCInformation

`func (o *ChargingdataPost400Response) HasRoamingQBCInformation() bool`

HasRoamingQBCInformation returns a boolean if a field has been set.

### GetLocationReportingChargingInformation

`func (o *ChargingdataPost400Response) GetLocationReportingChargingInformation() LocationReportingChargingInformation`

GetLocationReportingChargingInformation returns the LocationReportingChargingInformation field if non-nil, zero value otherwise.

### GetLocationReportingChargingInformationOk

`func (o *ChargingdataPost400Response) GetLocationReportingChargingInformationOk() (*LocationReportingChargingInformation, bool)`

GetLocationReportingChargingInformationOk returns a tuple with the LocationReportingChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReportingChargingInformation

`func (o *ChargingdataPost400Response) SetLocationReportingChargingInformation(v LocationReportingChargingInformation)`

SetLocationReportingChargingInformation sets LocationReportingChargingInformation field to given value.

### HasLocationReportingChargingInformation

`func (o *ChargingdataPost400Response) HasLocationReportingChargingInformation() bool`

HasLocationReportingChargingInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


