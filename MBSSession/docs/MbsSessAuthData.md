# MbsSessAuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtGroupId** | **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | 
**GpsisList** | Pointer to **map[string]string** | Represents the list of the GPSI(s) of the member UE(s) constituting the multicast MBS group. Any value of type can be used as a key of the map.  | [optional] 
**MbsSessionIdList** | [**NullableModel5MbsAuthorizationInfo**](Model5MbsAuthorizationInfo.md) |  | 

## Methods

### NewMbsSessAuthData

`func NewMbsSessAuthData(extGroupId string, mbsSessionIdList NullableModel5MbsAuthorizationInfo, ) *MbsSessAuthData`

NewMbsSessAuthData instantiates a new MbsSessAuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessAuthDataWithDefaults

`func NewMbsSessAuthDataWithDefaults() *MbsSessAuthData`

NewMbsSessAuthDataWithDefaults instantiates a new MbsSessAuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtGroupId

`func (o *MbsSessAuthData) GetExtGroupId() string`

GetExtGroupId returns the ExtGroupId field if non-nil, zero value otherwise.

### GetExtGroupIdOk

`func (o *MbsSessAuthData) GetExtGroupIdOk() (*string, bool)`

GetExtGroupIdOk returns a tuple with the ExtGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGroupId

`func (o *MbsSessAuthData) SetExtGroupId(v string)`

SetExtGroupId sets ExtGroupId field to given value.


### GetGpsisList

`func (o *MbsSessAuthData) GetGpsisList() map[string]string`

GetGpsisList returns the GpsisList field if non-nil, zero value otherwise.

### GetGpsisListOk

`func (o *MbsSessAuthData) GetGpsisListOk() (*map[string]string, bool)`

GetGpsisListOk returns a tuple with the GpsisList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsisList

`func (o *MbsSessAuthData) SetGpsisList(v map[string]string)`

SetGpsisList sets GpsisList field to given value.

### HasGpsisList

`func (o *MbsSessAuthData) HasGpsisList() bool`

HasGpsisList returns a boolean if a field has been set.

### GetMbsSessionIdList

`func (o *MbsSessAuthData) GetMbsSessionIdList() Model5MbsAuthorizationInfo`

GetMbsSessionIdList returns the MbsSessionIdList field if non-nil, zero value otherwise.

### GetMbsSessionIdListOk

`func (o *MbsSessAuthData) GetMbsSessionIdListOk() (*Model5MbsAuthorizationInfo, bool)`

GetMbsSessionIdListOk returns a tuple with the MbsSessionIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionIdList

`func (o *MbsSessAuthData) SetMbsSessionIdList(v Model5MbsAuthorizationInfo)`

SetMbsSessionIdList sets MbsSessionIdList field to given value.


### SetMbsSessionIdListNil

`func (o *MbsSessAuthData) SetMbsSessionIdListNil(b bool)`

 SetMbsSessionIdListNil sets the value for MbsSessionIdList to be an explicit nil

### UnsetMbsSessionIdList
`func (o *MbsSessAuthData) UnsetMbsSessionIdList()`

UnsetMbsSessionIdList ensures that no value is present for MbsSessionIdList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


