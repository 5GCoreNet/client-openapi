# AusfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**RoutingIndicators** | Pointer to **[]string** |  | [optional] 
**SuciInfos** | Pointer to [**[]SuciInfo**](SuciInfo.md) |  | [optional] 

## Methods

### NewAusfInfo

`func NewAusfInfo() *AusfInfo`

NewAusfInfo instantiates a new AusfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAusfInfoWithDefaults

`func NewAusfInfoWithDefaults() *AusfInfo`

NewAusfInfoWithDefaults instantiates a new AusfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *AusfInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AusfInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AusfInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AusfInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *AusfInfo) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *AusfInfo) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *AusfInfo) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *AusfInfo) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetRoutingIndicators

`func (o *AusfInfo) GetRoutingIndicators() []string`

GetRoutingIndicators returns the RoutingIndicators field if non-nil, zero value otherwise.

### GetRoutingIndicatorsOk

`func (o *AusfInfo) GetRoutingIndicatorsOk() (*[]string, bool)`

GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicators

`func (o *AusfInfo) SetRoutingIndicators(v []string)`

SetRoutingIndicators sets RoutingIndicators field to given value.

### HasRoutingIndicators

`func (o *AusfInfo) HasRoutingIndicators() bool`

HasRoutingIndicators returns a boolean if a field has been set.

### GetSuciInfos

`func (o *AusfInfo) GetSuciInfos() []SuciInfo`

GetSuciInfos returns the SuciInfos field if non-nil, zero value otherwise.

### GetSuciInfosOk

`func (o *AusfInfo) GetSuciInfosOk() (*[]SuciInfo, bool)`

GetSuciInfosOk returns a tuple with the SuciInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuciInfos

`func (o *AusfInfo) SetSuciInfos(v []SuciInfo)`

SetSuciInfos sets SuciInfos field to given value.

### HasSuciInfos

`func (o *AusfInfo) HasSuciInfos() bool`

HasSuciInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


