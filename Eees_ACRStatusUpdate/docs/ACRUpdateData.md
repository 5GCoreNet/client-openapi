# ACRUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasId** | **string** |  | 
**AcId** | Pointer to **string** |  | [optional] 
**ActResultInfo** | Pointer to [**ACTResultInfo**](ACTResultInfo.md) |  | [optional] 
**E3SubscIds** | Pointer to **[]string** |  | [optional] 
**E3NotificationUri** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 

## Methods

### NewACRUpdateData

`func NewACRUpdateData(easId string, ) *ACRUpdateData`

NewACRUpdateData instantiates a new ACRUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACRUpdateDataWithDefaults

`func NewACRUpdateDataWithDefaults() *ACRUpdateData`

NewACRUpdateDataWithDefaults instantiates a new ACRUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasId

`func (o *ACRUpdateData) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *ACRUpdateData) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *ACRUpdateData) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetAcId

`func (o *ACRUpdateData) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *ACRUpdateData) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *ACRUpdateData) SetAcId(v string)`

SetAcId sets AcId field to given value.

### HasAcId

`func (o *ACRUpdateData) HasAcId() bool`

HasAcId returns a boolean if a field has been set.

### GetActResultInfo

`func (o *ACRUpdateData) GetActResultInfo() ACTResultInfo`

GetActResultInfo returns the ActResultInfo field if non-nil, zero value otherwise.

### GetActResultInfoOk

`func (o *ACRUpdateData) GetActResultInfoOk() (*ACTResultInfo, bool)`

GetActResultInfoOk returns a tuple with the ActResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActResultInfo

`func (o *ACRUpdateData) SetActResultInfo(v ACTResultInfo)`

SetActResultInfo sets ActResultInfo field to given value.

### HasActResultInfo

`func (o *ACRUpdateData) HasActResultInfo() bool`

HasActResultInfo returns a boolean if a field has been set.

### GetE3SubscIds

`func (o *ACRUpdateData) GetE3SubscIds() []string`

GetE3SubscIds returns the E3SubscIds field if non-nil, zero value otherwise.

### GetE3SubscIdsOk

`func (o *ACRUpdateData) GetE3SubscIdsOk() (*[]string, bool)`

GetE3SubscIdsOk returns a tuple with the E3SubscIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE3SubscIds

`func (o *ACRUpdateData) SetE3SubscIds(v []string)`

SetE3SubscIds sets E3SubscIds field to given value.

### HasE3SubscIds

`func (o *ACRUpdateData) HasE3SubscIds() bool`

HasE3SubscIds returns a boolean if a field has been set.

### GetE3NotificationUri

`func (o *ACRUpdateData) GetE3NotificationUri() string`

GetE3NotificationUri returns the E3NotificationUri field if non-nil, zero value otherwise.

### GetE3NotificationUriOk

`func (o *ACRUpdateData) GetE3NotificationUriOk() (*string, bool)`

GetE3NotificationUriOk returns a tuple with the E3NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE3NotificationUri

`func (o *ACRUpdateData) SetE3NotificationUri(v string)`

SetE3NotificationUri sets E3NotificationUri field to given value.

### HasE3NotificationUri

`func (o *ACRUpdateData) HasE3NotificationUri() bool`

HasE3NotificationUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


