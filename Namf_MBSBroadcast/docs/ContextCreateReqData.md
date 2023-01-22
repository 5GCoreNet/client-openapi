# ContextCreateReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**MbsServiceAreaInfoList** | Pointer to [**[]MbsServiceAreaInfo**](MbsServiceAreaInfo.md) |  | [optional] 
**MbsServiceArea** | Pointer to [**MbsServiceArea**](MbsServiceArea.md) |  | [optional] 
**N2MbsSmInfo** | [**N2MbsSmInfo**](N2MbsSmInfo.md) |  | 
**NotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**MaxResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**MbsmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**MbsmfServiceInstId** | Pointer to **string** |  | [optional] 

## Methods

### NewContextCreateReqData

`func NewContextCreateReqData(mbsSessionId MbsSessionId, n2MbsSmInfo N2MbsSmInfo, notifyUri string, snssai Snssai, ) *ContextCreateReqData`

NewContextCreateReqData instantiates a new ContextCreateReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextCreateReqDataWithDefaults

`func NewContextCreateReqDataWithDefaults() *ContextCreateReqData`

NewContextCreateReqDataWithDefaults instantiates a new ContextCreateReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *ContextCreateReqData) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *ContextCreateReqData) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *ContextCreateReqData) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetMbsServiceAreaInfoList

`func (o *ContextCreateReqData) GetMbsServiceAreaInfoList() []MbsServiceAreaInfo`

GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field if non-nil, zero value otherwise.

### GetMbsServiceAreaInfoListOk

`func (o *ContextCreateReqData) GetMbsServiceAreaInfoListOk() (*[]MbsServiceAreaInfo, bool)`

GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceAreaInfoList

`func (o *ContextCreateReqData) SetMbsServiceAreaInfoList(v []MbsServiceAreaInfo)`

SetMbsServiceAreaInfoList sets MbsServiceAreaInfoList field to given value.

### HasMbsServiceAreaInfoList

`func (o *ContextCreateReqData) HasMbsServiceAreaInfoList() bool`

HasMbsServiceAreaInfoList returns a boolean if a field has been set.

### GetMbsServiceArea

`func (o *ContextCreateReqData) GetMbsServiceArea() MbsServiceArea`

GetMbsServiceArea returns the MbsServiceArea field if non-nil, zero value otherwise.

### GetMbsServiceAreaOk

`func (o *ContextCreateReqData) GetMbsServiceAreaOk() (*MbsServiceArea, bool)`

GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceArea

`func (o *ContextCreateReqData) SetMbsServiceArea(v MbsServiceArea)`

SetMbsServiceArea sets MbsServiceArea field to given value.

### HasMbsServiceArea

`func (o *ContextCreateReqData) HasMbsServiceArea() bool`

HasMbsServiceArea returns a boolean if a field has been set.

### GetN2MbsSmInfo

`func (o *ContextCreateReqData) GetN2MbsSmInfo() N2MbsSmInfo`

GetN2MbsSmInfo returns the N2MbsSmInfo field if non-nil, zero value otherwise.

### GetN2MbsSmInfoOk

`func (o *ContextCreateReqData) GetN2MbsSmInfoOk() (*N2MbsSmInfo, bool)`

GetN2MbsSmInfoOk returns a tuple with the N2MbsSmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsSmInfo

`func (o *ContextCreateReqData) SetN2MbsSmInfo(v N2MbsSmInfo)`

SetN2MbsSmInfo sets N2MbsSmInfo field to given value.


### GetNotifyUri

`func (o *ContextCreateReqData) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *ContextCreateReqData) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *ContextCreateReqData) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.


### GetMaxResponseTime

`func (o *ContextCreateReqData) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *ContextCreateReqData) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *ContextCreateReqData) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *ContextCreateReqData) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetSnssai

`func (o *ContextCreateReqData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *ContextCreateReqData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *ContextCreateReqData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetMbsmfId

`func (o *ContextCreateReqData) GetMbsmfId() string`

GetMbsmfId returns the MbsmfId field if non-nil, zero value otherwise.

### GetMbsmfIdOk

`func (o *ContextCreateReqData) GetMbsmfIdOk() (*string, bool)`

GetMbsmfIdOk returns a tuple with the MbsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsmfId

`func (o *ContextCreateReqData) SetMbsmfId(v string)`

SetMbsmfId sets MbsmfId field to given value.

### HasMbsmfId

`func (o *ContextCreateReqData) HasMbsmfId() bool`

HasMbsmfId returns a boolean if a field has been set.

### GetMbsmfServiceInstId

`func (o *ContextCreateReqData) GetMbsmfServiceInstId() string`

GetMbsmfServiceInstId returns the MbsmfServiceInstId field if non-nil, zero value otherwise.

### GetMbsmfServiceInstIdOk

`func (o *ContextCreateReqData) GetMbsmfServiceInstIdOk() (*string, bool)`

GetMbsmfServiceInstIdOk returns a tuple with the MbsmfServiceInstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsmfServiceInstId

`func (o *ContextCreateReqData) SetMbsmfServiceInstId(v string)`

SetMbsmfServiceInstId sets MbsmfServiceInstId field to given value.

### HasMbsmfServiceInstId

`func (o *ContextCreateReqData) HasMbsmfServiceInstId() bool`

HasMbsmfServiceInstId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


