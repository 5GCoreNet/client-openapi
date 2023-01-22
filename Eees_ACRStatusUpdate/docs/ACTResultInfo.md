# ACTResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActResult** | [**ACTResult**](ACTResult.md) |  | 
**ActFailureCause** | Pointer to [**ACTFailureCause**](ACTFailureCause.md) |  | [optional] 
**UeId** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**EasEndPoint** | [**EndPoint**](EndPoint.md) |  | 

## Methods

### NewACTResultInfo

`func NewACTResultInfo(actResult ACTResult, ueId string, easEndPoint EndPoint, ) *ACTResultInfo`

NewACTResultInfo instantiates a new ACTResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACTResultInfoWithDefaults

`func NewACTResultInfoWithDefaults() *ACTResultInfo`

NewACTResultInfoWithDefaults instantiates a new ACTResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActResult

`func (o *ACTResultInfo) GetActResult() ACTResult`

GetActResult returns the ActResult field if non-nil, zero value otherwise.

### GetActResultOk

`func (o *ACTResultInfo) GetActResultOk() (*ACTResult, bool)`

GetActResultOk returns a tuple with the ActResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActResult

`func (o *ACTResultInfo) SetActResult(v ACTResult)`

SetActResult sets ActResult field to given value.


### GetActFailureCause

`func (o *ACTResultInfo) GetActFailureCause() ACTFailureCause`

GetActFailureCause returns the ActFailureCause field if non-nil, zero value otherwise.

### GetActFailureCauseOk

`func (o *ACTResultInfo) GetActFailureCauseOk() (*ACTFailureCause, bool)`

GetActFailureCauseOk returns a tuple with the ActFailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActFailureCause

`func (o *ACTResultInfo) SetActFailureCause(v ACTFailureCause)`

SetActFailureCause sets ActFailureCause field to given value.

### HasActFailureCause

`func (o *ACTResultInfo) HasActFailureCause() bool`

HasActFailureCause returns a boolean if a field has been set.

### GetUeId

`func (o *ACTResultInfo) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ACTResultInfo) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ACTResultInfo) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetEasEndPoint

`func (o *ACTResultInfo) GetEasEndPoint() EndPoint`

GetEasEndPoint returns the EasEndPoint field if non-nil, zero value otherwise.

### GetEasEndPointOk

`func (o *ACTResultInfo) GetEasEndPointOk() (*EndPoint, bool)`

GetEasEndPointOk returns a tuple with the EasEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasEndPoint

`func (o *ACTResultInfo) SetEasEndPoint(v EndPoint)`

SetEasEndPoint sets EasEndPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


