# TscQosRequirementRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReqGbrDl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**ReqGbrUl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**ReqMbrDl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**ReqMbrUl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**MaxTscBurstSize** | Pointer to **NullableInt32** | This data type is defined in the same way as the &#39;ExtMaxDataBurstVol&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**Req5Gsdelay** | Pointer to **NullableInt32** | This data type is defined in the same way as the &#39;PacketDelBudget&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**Priority** | Pointer to **NullableInt32** | This data type is defined in the same way as the TscPriorityLevel data type, but with the OpenAPI nullable property set to true. | [optional] 
**TscaiTimeDom** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**TscaiInputDl** | Pointer to [**NullableTscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 
**TscaiInputUl** | Pointer to [**NullableTscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 

## Methods

### NewTscQosRequirementRm

`func NewTscQosRequirementRm() *TscQosRequirementRm`

NewTscQosRequirementRm instantiates a new TscQosRequirementRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscQosRequirementRmWithDefaults

`func NewTscQosRequirementRmWithDefaults() *TscQosRequirementRm`

NewTscQosRequirementRmWithDefaults instantiates a new TscQosRequirementRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReqGbrDl

`func (o *TscQosRequirementRm) GetReqGbrDl() string`

GetReqGbrDl returns the ReqGbrDl field if non-nil, zero value otherwise.

### GetReqGbrDlOk

`func (o *TscQosRequirementRm) GetReqGbrDlOk() (*string, bool)`

GetReqGbrDlOk returns a tuple with the ReqGbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqGbrDl

`func (o *TscQosRequirementRm) SetReqGbrDl(v string)`

SetReqGbrDl sets ReqGbrDl field to given value.

### HasReqGbrDl

`func (o *TscQosRequirementRm) HasReqGbrDl() bool`

HasReqGbrDl returns a boolean if a field has been set.

### SetReqGbrDlNil

`func (o *TscQosRequirementRm) SetReqGbrDlNil(b bool)`

 SetReqGbrDlNil sets the value for ReqGbrDl to be an explicit nil

### UnsetReqGbrDl
`func (o *TscQosRequirementRm) UnsetReqGbrDl()`

UnsetReqGbrDl ensures that no value is present for ReqGbrDl, not even an explicit nil
### GetReqGbrUl

`func (o *TscQosRequirementRm) GetReqGbrUl() string`

GetReqGbrUl returns the ReqGbrUl field if non-nil, zero value otherwise.

### GetReqGbrUlOk

`func (o *TscQosRequirementRm) GetReqGbrUlOk() (*string, bool)`

GetReqGbrUlOk returns a tuple with the ReqGbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqGbrUl

`func (o *TscQosRequirementRm) SetReqGbrUl(v string)`

SetReqGbrUl sets ReqGbrUl field to given value.

### HasReqGbrUl

`func (o *TscQosRequirementRm) HasReqGbrUl() bool`

HasReqGbrUl returns a boolean if a field has been set.

### SetReqGbrUlNil

`func (o *TscQosRequirementRm) SetReqGbrUlNil(b bool)`

 SetReqGbrUlNil sets the value for ReqGbrUl to be an explicit nil

### UnsetReqGbrUl
`func (o *TscQosRequirementRm) UnsetReqGbrUl()`

UnsetReqGbrUl ensures that no value is present for ReqGbrUl, not even an explicit nil
### GetReqMbrDl

`func (o *TscQosRequirementRm) GetReqMbrDl() string`

GetReqMbrDl returns the ReqMbrDl field if non-nil, zero value otherwise.

### GetReqMbrDlOk

`func (o *TscQosRequirementRm) GetReqMbrDlOk() (*string, bool)`

GetReqMbrDlOk returns a tuple with the ReqMbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMbrDl

`func (o *TscQosRequirementRm) SetReqMbrDl(v string)`

SetReqMbrDl sets ReqMbrDl field to given value.

### HasReqMbrDl

`func (o *TscQosRequirementRm) HasReqMbrDl() bool`

HasReqMbrDl returns a boolean if a field has been set.

### SetReqMbrDlNil

`func (o *TscQosRequirementRm) SetReqMbrDlNil(b bool)`

 SetReqMbrDlNil sets the value for ReqMbrDl to be an explicit nil

### UnsetReqMbrDl
`func (o *TscQosRequirementRm) UnsetReqMbrDl()`

UnsetReqMbrDl ensures that no value is present for ReqMbrDl, not even an explicit nil
### GetReqMbrUl

`func (o *TscQosRequirementRm) GetReqMbrUl() string`

GetReqMbrUl returns the ReqMbrUl field if non-nil, zero value otherwise.

### GetReqMbrUlOk

`func (o *TscQosRequirementRm) GetReqMbrUlOk() (*string, bool)`

GetReqMbrUlOk returns a tuple with the ReqMbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMbrUl

`func (o *TscQosRequirementRm) SetReqMbrUl(v string)`

SetReqMbrUl sets ReqMbrUl field to given value.

### HasReqMbrUl

`func (o *TscQosRequirementRm) HasReqMbrUl() bool`

HasReqMbrUl returns a boolean if a field has been set.

### SetReqMbrUlNil

`func (o *TscQosRequirementRm) SetReqMbrUlNil(b bool)`

 SetReqMbrUlNil sets the value for ReqMbrUl to be an explicit nil

### UnsetReqMbrUl
`func (o *TscQosRequirementRm) UnsetReqMbrUl()`

UnsetReqMbrUl ensures that no value is present for ReqMbrUl, not even an explicit nil
### GetMaxTscBurstSize

`func (o *TscQosRequirementRm) GetMaxTscBurstSize() int32`

GetMaxTscBurstSize returns the MaxTscBurstSize field if non-nil, zero value otherwise.

### GetMaxTscBurstSizeOk

`func (o *TscQosRequirementRm) GetMaxTscBurstSizeOk() (*int32, bool)`

GetMaxTscBurstSizeOk returns a tuple with the MaxTscBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTscBurstSize

`func (o *TscQosRequirementRm) SetMaxTscBurstSize(v int32)`

SetMaxTscBurstSize sets MaxTscBurstSize field to given value.

### HasMaxTscBurstSize

`func (o *TscQosRequirementRm) HasMaxTscBurstSize() bool`

HasMaxTscBurstSize returns a boolean if a field has been set.

### SetMaxTscBurstSizeNil

`func (o *TscQosRequirementRm) SetMaxTscBurstSizeNil(b bool)`

 SetMaxTscBurstSizeNil sets the value for MaxTscBurstSize to be an explicit nil

### UnsetMaxTscBurstSize
`func (o *TscQosRequirementRm) UnsetMaxTscBurstSize()`

UnsetMaxTscBurstSize ensures that no value is present for MaxTscBurstSize, not even an explicit nil
### GetReq5Gsdelay

`func (o *TscQosRequirementRm) GetReq5Gsdelay() int32`

GetReq5Gsdelay returns the Req5Gsdelay field if non-nil, zero value otherwise.

### GetReq5GsdelayOk

`func (o *TscQosRequirementRm) GetReq5GsdelayOk() (*int32, bool)`

GetReq5GsdelayOk returns a tuple with the Req5Gsdelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReq5Gsdelay

`func (o *TscQosRequirementRm) SetReq5Gsdelay(v int32)`

SetReq5Gsdelay sets Req5Gsdelay field to given value.

### HasReq5Gsdelay

`func (o *TscQosRequirementRm) HasReq5Gsdelay() bool`

HasReq5Gsdelay returns a boolean if a field has been set.

### SetReq5GsdelayNil

`func (o *TscQosRequirementRm) SetReq5GsdelayNil(b bool)`

 SetReq5GsdelayNil sets the value for Req5Gsdelay to be an explicit nil

### UnsetReq5Gsdelay
`func (o *TscQosRequirementRm) UnsetReq5Gsdelay()`

UnsetReq5Gsdelay ensures that no value is present for Req5Gsdelay, not even an explicit nil
### GetPriority

`func (o *TscQosRequirementRm) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TscQosRequirementRm) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TscQosRequirementRm) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TscQosRequirementRm) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TscQosRequirementRm) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TscQosRequirementRm) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetTscaiTimeDom

`func (o *TscQosRequirementRm) GetTscaiTimeDom() int32`

GetTscaiTimeDom returns the TscaiTimeDom field if non-nil, zero value otherwise.

### GetTscaiTimeDomOk

`func (o *TscQosRequirementRm) GetTscaiTimeDomOk() (*int32, bool)`

GetTscaiTimeDomOk returns a tuple with the TscaiTimeDom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiTimeDom

`func (o *TscQosRequirementRm) SetTscaiTimeDom(v int32)`

SetTscaiTimeDom sets TscaiTimeDom field to given value.

### HasTscaiTimeDom

`func (o *TscQosRequirementRm) HasTscaiTimeDom() bool`

HasTscaiTimeDom returns a boolean if a field has been set.

### SetTscaiTimeDomNil

`func (o *TscQosRequirementRm) SetTscaiTimeDomNil(b bool)`

 SetTscaiTimeDomNil sets the value for TscaiTimeDom to be an explicit nil

### UnsetTscaiTimeDom
`func (o *TscQosRequirementRm) UnsetTscaiTimeDom()`

UnsetTscaiTimeDom ensures that no value is present for TscaiTimeDom, not even an explicit nil
### GetTscaiInputDl

`func (o *TscQosRequirementRm) GetTscaiInputDl() TscaiInputContainer`

GetTscaiInputDl returns the TscaiInputDl field if non-nil, zero value otherwise.

### GetTscaiInputDlOk

`func (o *TscQosRequirementRm) GetTscaiInputDlOk() (*TscaiInputContainer, bool)`

GetTscaiInputDlOk returns a tuple with the TscaiInputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputDl

`func (o *TscQosRequirementRm) SetTscaiInputDl(v TscaiInputContainer)`

SetTscaiInputDl sets TscaiInputDl field to given value.

### HasTscaiInputDl

`func (o *TscQosRequirementRm) HasTscaiInputDl() bool`

HasTscaiInputDl returns a boolean if a field has been set.

### SetTscaiInputDlNil

`func (o *TscQosRequirementRm) SetTscaiInputDlNil(b bool)`

 SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil

### UnsetTscaiInputDl
`func (o *TscQosRequirementRm) UnsetTscaiInputDl()`

UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
### GetTscaiInputUl

`func (o *TscQosRequirementRm) GetTscaiInputUl() TscaiInputContainer`

GetTscaiInputUl returns the TscaiInputUl field if non-nil, zero value otherwise.

### GetTscaiInputUlOk

`func (o *TscQosRequirementRm) GetTscaiInputUlOk() (*TscaiInputContainer, bool)`

GetTscaiInputUlOk returns a tuple with the TscaiInputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputUl

`func (o *TscQosRequirementRm) SetTscaiInputUl(v TscaiInputContainer)`

SetTscaiInputUl sets TscaiInputUl field to given value.

### HasTscaiInputUl

`func (o *TscQosRequirementRm) HasTscaiInputUl() bool`

HasTscaiInputUl returns a boolean if a field has been set.

### SetTscaiInputUlNil

`func (o *TscQosRequirementRm) SetTscaiInputUlNil(b bool)`

 SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil

### UnsetTscaiInputUl
`func (o *TscQosRequirementRm) UnsetTscaiInputUl()`

UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


