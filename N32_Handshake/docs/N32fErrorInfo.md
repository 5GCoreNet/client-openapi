# N32fErrorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N32fMessageId** | **string** |  | 
**N32fErrorType** | [**N32fErrorType**](N32fErrorType.md) |  | 
**N32fContextId** | Pointer to **string** |  | [optional] 
**FailedModificationList** | Pointer to [**[]FailedModificationInfo**](FailedModificationInfo.md) |  | [optional] 
**ErrorDetailsList** | Pointer to [**[]N32fErrorDetail**](N32fErrorDetail.md) |  | [optional] 

## Methods

### NewN32fErrorInfo

`func NewN32fErrorInfo(n32fMessageId string, n32fErrorType N32fErrorType, ) *N32fErrorInfo`

NewN32fErrorInfo instantiates a new N32fErrorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN32fErrorInfoWithDefaults

`func NewN32fErrorInfoWithDefaults() *N32fErrorInfo`

NewN32fErrorInfoWithDefaults instantiates a new N32fErrorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN32fMessageId

`func (o *N32fErrorInfo) GetN32fMessageId() string`

GetN32fMessageId returns the N32fMessageId field if non-nil, zero value otherwise.

### GetN32fMessageIdOk

`func (o *N32fErrorInfo) GetN32fMessageIdOk() (*string, bool)`

GetN32fMessageIdOk returns a tuple with the N32fMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN32fMessageId

`func (o *N32fErrorInfo) SetN32fMessageId(v string)`

SetN32fMessageId sets N32fMessageId field to given value.


### GetN32fErrorType

`func (o *N32fErrorInfo) GetN32fErrorType() N32fErrorType`

GetN32fErrorType returns the N32fErrorType field if non-nil, zero value otherwise.

### GetN32fErrorTypeOk

`func (o *N32fErrorInfo) GetN32fErrorTypeOk() (*N32fErrorType, bool)`

GetN32fErrorTypeOk returns a tuple with the N32fErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN32fErrorType

`func (o *N32fErrorInfo) SetN32fErrorType(v N32fErrorType)`

SetN32fErrorType sets N32fErrorType field to given value.


### GetN32fContextId

`func (o *N32fErrorInfo) GetN32fContextId() string`

GetN32fContextId returns the N32fContextId field if non-nil, zero value otherwise.

### GetN32fContextIdOk

`func (o *N32fErrorInfo) GetN32fContextIdOk() (*string, bool)`

GetN32fContextIdOk returns a tuple with the N32fContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN32fContextId

`func (o *N32fErrorInfo) SetN32fContextId(v string)`

SetN32fContextId sets N32fContextId field to given value.

### HasN32fContextId

`func (o *N32fErrorInfo) HasN32fContextId() bool`

HasN32fContextId returns a boolean if a field has been set.

### GetFailedModificationList

`func (o *N32fErrorInfo) GetFailedModificationList() []FailedModificationInfo`

GetFailedModificationList returns the FailedModificationList field if non-nil, zero value otherwise.

### GetFailedModificationListOk

`func (o *N32fErrorInfo) GetFailedModificationListOk() (*[]FailedModificationInfo, bool)`

GetFailedModificationListOk returns a tuple with the FailedModificationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedModificationList

`func (o *N32fErrorInfo) SetFailedModificationList(v []FailedModificationInfo)`

SetFailedModificationList sets FailedModificationList field to given value.

### HasFailedModificationList

`func (o *N32fErrorInfo) HasFailedModificationList() bool`

HasFailedModificationList returns a boolean if a field has been set.

### GetErrorDetailsList

`func (o *N32fErrorInfo) GetErrorDetailsList() []N32fErrorDetail`

GetErrorDetailsList returns the ErrorDetailsList field if non-nil, zero value otherwise.

### GetErrorDetailsListOk

`func (o *N32fErrorInfo) GetErrorDetailsListOk() (*[]N32fErrorDetail, bool)`

GetErrorDetailsListOk returns a tuple with the ErrorDetailsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetailsList

`func (o *N32fErrorInfo) SetErrorDetailsList(v []N32fErrorDetail)`

SetErrorDetailsList sets ErrorDetailsList field to given value.

### HasErrorDetailsList

`func (o *N32fErrorInfo) HasErrorDetailsList() bool`

HasErrorDetailsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


